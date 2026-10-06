package obp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestURLsFollowTheSpace(t *testing.T) {
	bank := NewClient("http://obp.test/", "http://oidc.test", "id", "secret", "ogcr")
	if got, want := bank.entityURL("chain_sync_status", "", nil), "http://obp.test/obp/v7.0.0/banks/ogcr/dynamic-entities/chain_sync_status"; got != want {
		t.Errorf("bank entityURL = %q, want %q", got, want)
	}
	if got, want := bank.managementURL(), "http://obp.test/obp/v7.0.0/management/banks/ogcr/dynamic-entities"; got != want {
		t.Errorf("bank managementURL = %q, want %q", got, want)
	}

	sys := NewClient("http://obp.test", "http://oidc.test", "id", "secret", "")
	if got, want := sys.entityURL("parcel_on_chain", "/r1", nil), "http://obp.test/obp/v7.0.0/banks/SYS/dynamic-entities/parcel_on_chain/r1"; got != want {
		t.Errorf("system entityURL = %q, want %q", got, want)
	}
	if got, want := sys.managementURL(), "http://obp.test/obp/v7.0.0/management/banks/SYS/dynamic-entities"; got != want {
		t.Errorf("system managementURL = %q, want %q", got, want)
	}
}

func TestUnwrap(t *testing.T) {
	wrapped := map[string]any{"parcel_on_chain": map[string]any{"parcel_on_chain_id": "r1"}, "metadata": map[string]any{}}
	if got := unwrap("parcel_on_chain", wrapped)["parcel_on_chain_id"]; got != "r1" {
		t.Errorf("unwrap(wrapped) id = %v, want r1", got)
	}
	flat := map[string]any{"parcel_on_chain_id": "r2"}
	if got := unwrap("parcel_on_chain", flat)["parcel_on_chain_id"]; got != "r2" {
		t.Errorf("unwrap(flat) id = %v, want r2", got)
	}
}

// fakeOBP is an OBP-OIDC issuer and an OBP API in one server. It issues tokens
// tok-1, tok-2, ... and answers API calls with apiStatus (200 if unset).
type fakeOBP struct {
	*httptest.Server
	tokens      int
	discoveries int
	apiStatus   []int // per API call, in order; then 200
	apiCalls    []string
}

func newFakeOBP(t *testing.T) *fakeOBP {
	f := &fakeOBP{}
	mux := http.NewServeMux()
	mux.HandleFunc("/oidc/.well-known/openid-configuration", func(w http.ResponseWriter, r *http.Request) {
		f.discoveries++
		json.NewEncoder(w).Encode(map[string]string{"token_endpoint": f.URL + "/oidc/token"})
	})
	mux.HandleFunc("/oidc/token", func(w http.ResponseWriter, r *http.Request) {
		id, secret, ok := r.BasicAuth()
		if !ok || id != "chain-cache" || secret != "s3cret" {
			t.Errorf("token request basic auth = %q/%q (ok=%v)", id, secret, ok)
		}
		if r.FormValue("grant_type") != "client_credentials" || r.FormValue("scope") != "openid" {
			t.Errorf("token request form = %v", r.Form)
		}
		f.tokens++
		json.NewEncoder(w).Encode(map[string]any{"access_token": "tok-" + string(rune('0'+f.tokens)), "expires_in": 300})
	})
	mux.HandleFunc("/obp/", func(w http.ResponseWriter, r *http.Request) {
		f.apiCalls = append(f.apiCalls, r.Header.Get("Authorization"))
		status := http.StatusOK
		if n := len(f.apiCalls) - 1; n < len(f.apiStatus) {
			status = f.apiStatus[n]
		}
		w.WriteHeader(status)
		w.Write([]byte(`{}`))
	})
	f.Server = httptest.NewServer(mux)
	t.Cleanup(f.Close)
	return f
}

func (f *fakeOBP) client() *Client {
	return NewClient(f.URL, f.URL+"/oidc/", "chain-cache", "s3cret", "ogcr")
}

func TestTokenIsDiscoveredFetchedAndReused(t *testing.T) {
	f := newFakeOBP(t)
	c := f.client()
	for i := 0; i < 3; i++ {
		if err := c.do("GET", f.URL+"/obp/x", nil, nil); err != nil {
			t.Fatal(err)
		}
	}
	if f.discoveries != 1 || f.tokens != 1 {
		t.Errorf("discoveries=%d tokens=%d, want 1 and 1", f.discoveries, f.tokens)
	}
	for _, h := range f.apiCalls {
		if h != "Bearer tok-1" {
			t.Errorf("Authorization = %q, want Bearer tok-1", h)
		}
	}
}

func TestExpiredTokenIsReplaced(t *testing.T) {
	f := newFakeOBP(t)
	c := f.client()
	if err := c.do("GET", f.URL+"/obp/x", nil, nil); err != nil {
		t.Fatal(err)
	}
	c.expiresAt = c.expiresAt.Add(-300 * time.Second)
	if err := c.do("GET", f.URL+"/obp/x", nil, nil); err != nil {
		t.Fatal(err)
	}
	if f.tokens != 2 || f.apiCalls[1] != "Bearer tok-2" {
		t.Errorf("tokens=%d calls=%v, want a second token on the second call", f.tokens, f.apiCalls)
	}
}

func TestUnauthorizedRetriesOnceWithANewToken(t *testing.T) {
	f := newFakeOBP(t)
	f.apiStatus = []int{http.StatusUnauthorized}
	if err := f.client().do("GET", f.URL+"/obp/x", nil, nil); err != nil {
		t.Fatal(err)
	}
	if want := []string{"Bearer tok-1", "Bearer tok-2"}; len(f.apiCalls) != 2 || f.apiCalls[0] != want[0] || f.apiCalls[1] != want[1] {
		t.Errorf("calls = %v, want %v", f.apiCalls, want)
	}
}

func TestUnauthorizedTwiceFails(t *testing.T) {
	f := newFakeOBP(t)
	f.apiStatus = []int{http.StatusUnauthorized, http.StatusUnauthorized, http.StatusUnauthorized}
	err := f.client().do("GET", f.URL+"/obp/x", nil, nil)
	if apiErr, ok := err.(*APIError); !ok || apiErr.Status != http.StatusUnauthorized {
		t.Fatalf("err = %v, want a 401 APIError", err)
	}
	if len(f.apiCalls) != 2 {
		t.Errorf("%d calls, want 2 (one retry)", len(f.apiCalls))
	}
}

func TestRequiredScopes(t *testing.T) {
	scopes := NewClient("http://obp.test", "http://oidc.test", "id", "secret", "").RequiredScopes([]string{"parcel_on_chain"})
	if len(scopes) != 7 {
		t.Fatalf("%d scopes, want 3 definition + 4 record", len(scopes))
	}
	for _, s := range scopes {
		if s.BankID != "SYS" {
			t.Errorf("%s at %q, want SYS for system level", s.RoleName, s.BankID)
		}
		if s.NeededFor == "" || len(s.NeededFor) > 1000 {
			t.Errorf("%s needed_for length %d, OBP wants 1 to 1000", s.RoleName, len(s.NeededFor))
		}
	}
	if got := scopes[4].RoleName; got != "CanCreateDynamicEntityRecord_parcel_on_chain" {
		t.Errorf("scopes[4] = %s", got)
	}
}
