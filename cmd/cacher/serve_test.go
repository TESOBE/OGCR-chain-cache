package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

func testServer() *server {
	client := obp.NewClient("http://obp.test", "http://oidc.test", "id", "secret", "ogcr")
	return &server{
		rn: &runner{
			client:  client,
			chainID: 2025,
			obpURL:  "http://localhost:8080",
			rpcURL:  "https://rpc.example/v3/SECRET-KEY?x=1",
		},
		interval: 30 * time.Second,
		runNow:   make(chan struct{}, 1),
	}
}

func (s *server) addRun(status string, errs ...string) {
	finished := time.Now()
	s.history = append([]*Report{{
		Started:      finished.Add(-2 * time.Second),
		Finished:     finished,
		ChainID:      2025,
		HeadBlock:    1234,
		RunStatus:    status,
		Mirrors:      []MirrorLine{{Name: "parcel", Ran: true, Count: 3}, {Name: "credit_batch"}},
		Errors:       errs,
		SyncRecorded: status != RunStatusNoChain,
	}}, s.history...)
}

func get(t *testing.T, s *server, h http.HandlerFunc, path string) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	h(w, httptest.NewRequest("GET", path, nil))
	if w.Code != http.StatusOK {
		t.Fatalf("GET %s = %d", path, w.Code)
	}
	return w
}

func TestHealthBeforeAnyRun(t *testing.T) {
	s := testServer()
	var h health
	if err := json.Unmarshal(get(t, s, s.handleHealth, "/health").Body.Bytes(), &h); err != nil {
		t.Fatal(err)
	}
	if h.Status != "starting" || h.LastRunAt != nil || h.MissingScopes == nil {
		t.Errorf("health = %+v, want starting with no last run and an empty missing list", h)
	}
}

func TestHealthReportsLastRunAndMissingScopes(t *testing.T) {
	s := testServer()
	s.addRun("partial", "upsert parcel failed parcel_id=p1 err=403")
	s.app = &obp.PlatformApp{ConsumerID: "c1", RequiredScopes: []obp.DeclaredScope{
		{Scope: obp.Scope{RoleName: "CanGetDynamicEntityDefinitions", BankID: "ogcr"}, Held: true},
		{Scope: obp.Scope{RoleName: "CanCreateDynamicEntityRecord_parcel_on_chain", BankID: "ogcr"}},
		{Scope: obp.Scope{RoleName: "CanDeleteDynamicEntityRecord_parcel_on_chain", BankID: "ogcr", Optional: true}},
	}}
	w := get(t, s, s.handleHealth, "/health")
	var h health
	if err := json.Unmarshal(w.Body.Bytes(), &h); err != nil {
		t.Fatal(err)
	}
	if h.Status != "partial" || h.HeadBlock != 1234 || h.ErrorCount != 1 {
		t.Errorf("health = %+v", h)
	}
	if len(h.MissingScopes) != 1 || h.MissingScopes[0] != "CanCreateDynamicEntityRecord_parcel_on_chain@ogcr" {
		t.Errorf("missing = %v, want only the required ungranted Scope", h.MissingScopes)
	}
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("health should be readable cross-origin by the dev-env page")
	}
}

func TestStatusPageRenders(t *testing.T) {
	s := testServer()
	s.addRun("ok")
	s.addRun("partial", "upsert parcel failed parcel_id=<script>")
	s.appErr = "PUT .../platform-app returned 404: OBP-xxxx This Consumer's CONSUMER_ID is c1"
	body := get(t, s, s.handleStatus, "/").Body.String()
	for _, want := range []string{"Partial: some records failed", "<code>parcel</code>", "skipped", "CONSUMER_ID is c1", "Recent runs", "https://rpc.example"} {
		if !strings.Contains(body, want) {
			t.Errorf("page is missing %q", want)
		}
	}
	if strings.Contains(body, "SECRET-KEY") {
		t.Error("page shows the RPC URL's path, which can hold an API key")
	}
	if strings.Contains(body, "parcel_id=<script>") {
		t.Error("error text is not escaped")
	}
}

func TestRunNowQueuesOneRun(t *testing.T) {
	s := testServer()
	for i := 0; i < 3; i++ {
		w := httptest.NewRecorder()
		s.handleRun(w, httptest.NewRequest("POST", "/run", nil))
		if w.Code != http.StatusSeeOther {
			t.Fatalf("POST /run = %d, want 303", w.Code)
		}
	}
	if len(s.runNow) != 1 {
		t.Errorf("%d runs queued, want 1", len(s.runNow))
	}
}

func TestRoundDuration(t *testing.T) {
	for d, want := range map[time.Duration]string{
		-time.Second:                         "0s",
		4*time.Second + 300*time.Millisecond: "4s",
		2*time.Minute + 10*time.Second:       "2m10s",
		3*time.Hour + 5*time.Minute:          "3h5m0s",
	} {
		if got := roundDuration(d); got != want {
			t.Errorf("roundDuration(%v) = %q, want %q", d, got, want)
		}
	}
}
