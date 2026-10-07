package setup

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/TESOBE/OGCR-chain-cache/entities"
	"github.com/TESOBE/OGCR-chain-cache/internal/cache"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

func TestBuiltInDefinitionsLoad(t *testing.T) {
	defs, err := Load(entities.FS)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, d := range defs {
		names = append(names, d.EntityName)
		if _, ok := d.Schema["has_public_access"]; ok {
			t.Errorf("%s: has_public_access left in the schema", d.EntityName)
		}
		if d.AuthMode == "" {
			t.Errorf("%s: no auth_mode, so OBP would refuse the Platform App's writes", d.EntityName)
		}
	}
	if !slices.Equal(names, Names()) {
		t.Errorf("definition names = %v, want %v", names, Names())
	}
	// Every entity the cacher writes must be one it sets up.
	if !slices.Equal(Names(), cache.Entities) {
		t.Errorf("Names() = %v, cache.Entities = %v", Names(), cache.Entities)
	}
}

func TestUpToDate(t *testing.T) {
	def, err := parse([]byte(`{"e": {"has_public_access": true, "auth_mode": "UserOrApplication", "required": ["k"], "properties": {"k": {"type": "string", "example": "x"}, "n": {"type": "integer", "example": 1}}}}`))
	if err != nil {
		t.Fatal(err)
	}
	live := obp.LiveEntity{
		EntityName:      "e",
		HasPublicAccess: true,
		AuthMode:        "UserOrApplication",
		// Same schema, different key order and a float for the number, as JSON decoding gives.
		Schema: map[string]any{"properties": map[string]any{"n": map[string]any{"example": 1.0, "type": "integer"}, "k": map[string]any{"example": "x", "type": "string"}}, "required": []any{"k"}},
	}
	if !UpToDate(def, live) {
		t.Error("identical definition reported as changed")
	}

	notPublic := live
	notPublic.HasPublicAccess = false
	if UpToDate(def, notPublic) {
		t.Error("public access difference not noticed")
	}

	otherSchema := live
	otherSchema.Schema = map[string]any{"properties": map[string]any{"k": map[string]any{"example": "x", "type": "string"}}, "required": []any{"k"}}
	if UpToDate(def, otherSchema) {
		t.Error("removed property not noticed")
	}
}

// fakeManagement is OBP-OIDC plus the v7.0.0 dynamic entity management API.
type fakeManagement struct {
	*httptest.Server
	live    []obp.LiveEntity
	failPut bool
	calls   []string // "POST name" / "PUT id"
}

func newFakeManagement(t *testing.T) *fakeManagement {
	f := &fakeManagement{}
	mux := http.NewServeMux()
	mux.HandleFunc("/oidc/.well-known/openid-configuration", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token_endpoint": f.URL + "/oidc/token"})
	})
	mux.HandleFunc("/oidc/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"access_token": "tok", "expires_in": 300})
	})
	base := "/obp/v7.0.0/management/banks/ogcr/dynamic-entities"
	mux.HandleFunc("GET "+base, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"dynamic_entities": f.live})
	})
	mux.HandleFunc("POST "+base, func(w http.ResponseWriter, r *http.Request) {
		var def obp.EntityDefinition
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &def)
		f.calls = append(f.calls, "POST "+def.EntityName)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{}`))
	})
	mux.HandleFunc("PUT "+base+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		f.calls = append(f.calls, "PUT "+r.PathValue("id"))
		if f.failPut {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"code":400,"message":"OBP-09xxx: cannot change the schema of an entity with records"}`))
			return
		}
		w.Write([]byte(`{}`))
	})
	f.Server = httptest.NewServer(mux)
	t.Cleanup(f.Close)
	return f
}

func TestApplyCreatesMissingAndUpdatesOnlyChanged(t *testing.T) {
	defs, err := Load(entities.FS)
	if err != nil {
		t.Fatal(err)
	}
	f := newFakeManagement(t)
	asLive := func(d obp.EntityDefinition, id string) obp.LiveEntity {
		return obp.LiveEntity{DynamicEntityID: id, EntityName: d.EntityName, HasPublicAccess: d.HasPublicAccess, AuthMode: d.AuthMode, Schema: d.Schema}
	}
	stale := asLive(defs[1], "id-activity")
	stale.HasPublicAccess = !stale.HasPublicAccess
	f.live = []obp.LiveEntity{asLive(defs[0], "id-parcel"), stale}

	client := obp.NewClient(f.URL, f.URL+"/oidc", "id", "secret", "ogcr")
	res, err := Apply(client, defs)
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"PUT id-activity"}
	for _, d := range defs[2:] {
		want = append(want, "POST "+d.EntityName)
	}
	if !slices.Equal(f.calls, want) {
		t.Errorf("calls = %v, want %v", f.calls, want)
	}
	if res.Outcomes[0].Action != Unchanged || res.Outcomes[1].Action != Updated || res.Outcomes[2].Action != Created {
		t.Errorf("outcomes = %+v", res.Outcomes)
	}
	if !res.Created() || len(res.Failed()) != 0 {
		t.Errorf("Created() = %v, Failed() = %v", res.Created(), res.Failed())
	}
}

func TestApplyGoesOnPastAFailure(t *testing.T) {
	defs, err := Load(entities.FS)
	if err != nil {
		t.Fatal(err)
	}
	f := newFakeManagement(t)
	f.failPut = true
	stale := obp.LiveEntity{DynamicEntityID: "id-parcel", EntityName: defs[0].EntityName, AuthMode: "UserOnly", Schema: map[string]any{}}
	f.live = []obp.LiveEntity{stale}

	client := obp.NewClient(f.URL, f.URL+"/oidc", "id", "secret", "ogcr")
	res, err := Apply(client, defs)
	if err != nil {
		t.Fatal(err)
	}
	if got := res.Failed(); !slices.Equal(got, []string{defs[0].EntityName}) {
		t.Errorf("Failed() = %v", got)
	}
	if !strings.Contains(res.Outcomes[0].Err, "entity with records") {
		t.Errorf("failure kept OBP's message? got %q", res.Outcomes[0].Err)
	}
	if n := len(f.calls); n != len(defs) {
		t.Errorf("%d calls, want one per definition (%d): %v", n, len(defs), f.calls)
	}
}
