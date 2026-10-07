// Package obp is a minimal OBP client: Platform App auth (OAuth2 client
// credentials from OBP-OIDC) plus the v7.0.0 dynamic-entity read and write
// calls the cacher needs. It calls OBP as its own application, not as a User,
// the way OBP-Sentinel does: the roles it needs are Scopes granted to its
// Consumer, which it declares with DeclarePlatformApp.
package obp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// tokenMargin is how long before its expiry a token is replaced, so a request
// never goes out with one about to lapse.
const tokenMargin = 30 * time.Second

type Client struct {
	baseURL      string
	issuer       string
	clientID     string
	clientSecret string
	// spaceID is the bank the dynamic entities live in; empty for system level.
	spaceID string
	http    *http.Client

	mu            sync.Mutex
	tokenEndpoint string
	token         string
	expiresAt     time.Time
}

func NewClient(baseURL, issuer, clientID, clientSecret, spaceID string) *Client {
	return &Client{
		baseURL:      strings.TrimRight(baseURL, "/"),
		issuer:       strings.TrimRight(issuer, "/"),
		clientID:     clientID,
		clientSecret: clientSecret,
		spaceID:      spaceID,
		http:         &http.Client{Timeout: 30 * time.Second},
	}
}

// discoverTokenEndpoint reads the issuer's token_endpoint from its OpenID
// configuration, once.
func (c *Client) discoverTokenEndpoint() (string, error) {
	if c.tokenEndpoint != "" {
		return c.tokenEndpoint, nil
	}
	resp, err := c.http.Get(c.issuer + "/.well-known/openid-configuration")
	if err != nil {
		return "", fmt.Errorf("oidc discovery failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("oidc discovery failed (%d): %s", resp.StatusCode, body)
	}
	var conf struct {
		TokenEndpoint string `json:"token_endpoint"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&conf); err != nil {
		return "", fmt.Errorf("failed to decode openid configuration: %w", err)
	}
	if conf.TokenEndpoint == "" {
		return "", fmt.Errorf("no token_endpoint in %s/.well-known/openid-configuration", c.issuer)
	}
	c.tokenEndpoint = conf.TokenEndpoint
	return c.tokenEndpoint, nil
}

// fetchToken gets a new access token with the client credentials grant.
// The caller holds c.mu.
func (c *Client) fetchToken() error {
	endpoint, err := c.discoverTokenEndpoint()
	if err != nil {
		return err
	}
	form := url.Values{"grant_type": {"client_credentials"}, "scope": {"openid"}}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	// Unescaped, as OBP-Sentinel sends them: OBP-OIDC reads them as given.
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("token request failed (%d): %s", resp.StatusCode, body)
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode token: %w", err)
	}
	if result.AccessToken == "" {
		return fmt.Errorf("empty access_token in token response")
	}
	expiresIn := time.Duration(result.ExpiresIn) * time.Second
	if expiresIn <= 0 {
		expiresIn = 300 * time.Second
	}
	c.token = result.AccessToken
	c.expiresAt = time.Now().Add(expiresIn - tokenMargin)
	return nil
}

// getToken returns a token that is still valid, fetching a new one if needed.
func (c *Client) getToken() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token == "" || !time.Now().Before(c.expiresAt) {
		if err := c.fetchToken(); err != nil {
			return "", err
		}
	}
	return c.token, nil
}

// dropToken forgets the cached token, so the next request fetches a new one.
func (c *Client) dropToken() {
	c.mu.Lock()
	c.token = ""
	c.mu.Unlock()
}

// do performs an authenticated request. On a 401 it fetches a new token and
// retries once (the token may have been revoked). A nil body is allowed (GET).
// out may be nil to discard the response body.
func (c *Client) do(method, fullURL string, body any, out any) error {
	var payload []byte
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		payload = b
	}

	for attempt := 0; ; attempt++ {
		token, err := c.getToken()
		if err != nil {
			return err
		}
		var reader io.Reader
		if payload != nil {
			reader = bytes.NewReader(payload)
		}
		req, err := http.NewRequest(method, fullURL, reader)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.http.Do(req)
		if err != nil {
			return fmt.Errorf("%s %s failed: %w", method, fullURL, err)
		}
		if resp.StatusCode == http.StatusUnauthorized && attempt == 0 {
			resp.Body.Close()
			c.dropToken()
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			b, _ := io.ReadAll(resp.Body)
			return &APIError{Method: method, URL: fullURL, Status: resp.StatusCode, Body: string(b)}
		}
		if out == nil {
			return nil
		}
		return json.NewDecoder(resp.Body).Decode(out)
	}
}

// APIError is a non-2xx response from OBP.
type APIError struct {
	Method, URL string
	Status      int
	Body        string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s %s returned %d: %s", e.Method, e.URL, e.Status, e.Body)
}

// space is the BANK_ID path segment for the client's space: the bank id, or the
// literal SYS for system level entities (OBP_ENTITY_SPACE_ID set to "").
func (c *Client) space() string {
	if c.spaceID == "" {
		return "SYS"
	}
	return c.spaceID
}

// entityURL is a v7.0.0 dynamic-entity record URL in the client's space.
func (c *Client) entityURL(entity, suffix string, params url.Values) string {
	u := fmt.Sprintf("%s/obp/v7.0.0/banks/%s/dynamic-entities/%s%s", c.baseURL, url.PathEscape(c.space()), entity, suffix)
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}

// unwrap returns the record fields of a v7.0.0 record body, which nests them
// under the entity name next to a `metadata` object.
func unwrap(entity string, body map[string]any) map[string]any {
	if inner, ok := body[entity].(map[string]any); ok {
		return inner
	}
	return body
}

// pageSize is how many records GetRecords asks for per request.
const pageSize = 500

// GetRecords returns dynamic-entity records, optionally filtered by query
// params, reading every page. OBP wraps lists as {"<entity>_list": [...]}.
func (c *Client) GetRecords(entity string, params url.Values) ([]map[string]any, error) {
	var all []map[string]any
	for offset := 0; ; offset += pageSize {
		q := url.Values{}
		for k, v := range params {
			q[k] = v
		}
		q.Set("obp_limit", fmt.Sprint(pageSize))
		q.Set("obp_offset", fmt.Sprint(offset))

		var raw map[string]json.RawMessage
		if err := c.do("GET", c.entityURL(entity, "", q), nil, &raw); err != nil {
			return nil, err
		}
		var page []map[string]any
		if listJSON, ok := raw[entity+"_list"]; ok {
			if err := json.Unmarshal(listJSON, &page); err != nil {
				return nil, err
			}
		}
		for _, r := range page {
			all = append(all, unwrap(entity, r))
		}
		if len(page) != pageSize {
			return all, nil
		}
	}
}

func (c *Client) CreateRecord(entity string, record any) (map[string]any, error) {
	var out map[string]any
	err := c.do("POST", c.entityURL(entity, "", nil), record, &out)
	return unwrap(entity, out), err
}

func (c *Client) UpdateRecord(entity, recordID string, record any) (map[string]any, error) {
	var out map[string]any
	err := c.do("PUT", c.entityURL(entity, "/"+url.PathEscape(recordID), nil), record, &out)
	return unwrap(entity, out), err
}

func (c *Client) DeleteRecord(entity, recordID string) error {
	return c.do("DELETE", c.entityURL(entity, "/"+url.PathEscape(recordID), nil), nil, nil)
}

// ── dynamic-entity definitions (one-time management) ────────────────────────
//
// These use the v7.0.0 management API, which is the same for every space: the
// BANK_ID segment is a bank id, or the literal SYS for the system space (older
// versions reject SYS with OBP-30001). Its request and response bodies name the
// entity in `entity_name` and carry the schema under `schema`.

func (c *Client) managementURL() string {
	return fmt.Sprintf("%s/obp/v7.0.0/management/banks/%s/dynamic-entities", c.baseURL, url.PathEscape(c.space()))
}

// Space describes where the entities live, for log messages.
func (c *Client) Space() string {
	if c.spaceID == "" {
		return "system level"
	}
	return "bank " + c.spaceID
}

// DynamicEntityIDs returns a map of entity name -> dynamic_entity_id for every
// dynamic entity defined in the client's space. The id is needed to update
// (PUT) an entity; a missing name means it must be created (POST).
func (c *Client) DynamicEntityIDs() (map[string]string, error) {
	live, err := c.DynamicEntities()
	if err != nil {
		return nil, err
	}
	ids := make(map[string]string, len(live))
	for name, e := range live {
		ids[name] = e.DynamicEntityID
	}
	return ids, nil
}

// LiveEntity is a dynamic entity definition as OBP holds it.
type LiveEntity struct {
	DynamicEntityID   string         `json:"dynamic_entity_id"`
	EntityName        string         `json:"entity_name"`
	HasPersonalEntity bool           `json:"has_personal_entity"`
	HasPublicAccess   bool           `json:"has_public_access"`
	AuthMode          string         `json:"auth_mode"`
	Schema            map[string]any `json:"schema"`
	RecordCount       int64          `json:"record_count"`
}

// DynamicEntities returns every dynamic entity defined in the client's space,
// by name, with its definition, so a caller can tell whether it is up to date.
func (c *Client) DynamicEntities() (map[string]LiveEntity, error) {
	var raw struct {
		DynamicEntities []LiveEntity `json:"dynamic_entities"`
	}
	if err := c.do("GET", c.managementURL(), nil, &raw); err != nil {
		return nil, err
	}
	live := make(map[string]LiveEntity)
	for _, e := range raw.DynamicEntities {
		if e.DynamicEntityID != "" && e.EntityName != "" {
			live[e.EntityName] = e
		}
	}
	return live, nil
}

// EntityDefinition is a v7.0.0 dynamic entity definition request body.
type EntityDefinition struct {
	EntityName        string `json:"entity_name"`
	HasPersonalEntity bool   `json:"has_personal_entity"`
	HasPublicAccess   bool   `json:"has_public_access"`
	// AuthMode says who may hold the roles that guard the records: the cacher
	// writes as an application, so it must be ApplicationOnly or
	// UserOrApplication (OBP's default, UserOnly, refuses an app token).
	AuthMode string         `json:"auth_mode,omitempty"`
	Schema   map[string]any `json:"schema"`
}

func (c *Client) CreateDynamicEntity(definition EntityDefinition) (map[string]any, error) {
	var out map[string]any
	err := c.do("POST", c.managementURL(), definition, &out)
	return out, err
}

// UpdateDynamicEntity replaces the definition of an existing dynamic entity in
// the client's space.
func (c *Client) UpdateDynamicEntity(entityID string, definition EntityDefinition) (map[string]any, error) {
	var out map[string]any
	err := c.do("PUT", c.managementURL()+"/"+url.PathEscape(entityID), definition, &out)
	return out, err
}

// ── Platform App ────────────────────────────────────────────────────────────
//
// An administrator marks this app's Consumer as a Platform App
// (POST /obp/v7.0.0/management/platform-apps); the app then declares the
// Scopes it needs, and the administrator grants them to its Consumer.

// Scope is one Role the app needs, at a bank id (SYS for the system space of
// dynamic entities).
type Scope struct {
	RoleName  string `json:"role_name"`
	BankID    string `json:"bank_id"`
	NeededFor string `json:"needed_for"`
	Optional  bool   `json:"optional"`
}

// DeclaredScope is a declared Scope as OBP reports it back.
type DeclaredScope struct {
	Scope
	Held bool `json:"held"`
}

// PlatformApp is the app as OBP sees it after a declaration.
type PlatformApp struct {
	ConsumerID     string          `json:"consumer_id"`
	Label          string          `json:"label"`
	RequiredScopes []DeclaredScope `json:"required_scopes"`
}

// Missing returns the required Scopes the app's Consumer does not hold.
func (a *PlatformApp) Missing() []Scope {
	var missing []Scope
	for _, s := range a.RequiredScopes {
		if !s.Held && !s.Optional {
			missing = append(missing, s.Scope)
		}
	}
	return missing
}

// BankID is the bank id the client's Scopes are granted at: the space, or SYS
// for system level entities.
func (c *Client) BankID() string { return c.space() }

// RequiredScopes lists every Scope the cacher, setup-entity and delete-records
// need in the client's space: the definition Roles, and the record Roles of each
// named entity. All three tools declare the same list, since a declaration
// replaces the previous one.
func (c *Client) RequiredScopes(entities []string) []Scope {
	bank := c.space()
	scopes := []Scope{
		{RoleName: "CanGetDynamicEntityDefinitions", BankID: bank, NeededFor: "setup-entity lists the *_on_chain entity definitions to decide whether to create or update each."},
		{RoleName: "CanCreateDynamicEntityDefinition", BankID: bank, NeededFor: "setup-entity creates the *_on_chain entities that mirror the OGCR chain."},
		{RoleName: "CanUpdateDynamicEntityDefinition", BankID: bank, NeededFor: "setup-entity updates the *_on_chain entity definitions when their schema changes."},
	}
	for _, e := range entities {
		scopes = append(scopes,
			Scope{RoleName: "CanGetDynamicEntityRecord_" + e, BankID: bank, NeededFor: "The cacher reads " + e + " records to update the existing record for a token instead of adding another."},
			Scope{RoleName: "CanCreateDynamicEntityRecord_" + e, BankID: bank, NeededFor: "The cacher writes a " + e + " record for each token it finds on the OGCR chain."},
			Scope{RoleName: "CanUpdateDynamicEntityRecord_" + e, BankID: bank, NeededFor: "The cacher updates a " + e + " record when the token's on-chain state changes."},
			Scope{RoleName: "CanDeleteDynamicEntityRecord_" + e, BankID: bank, NeededFor: "delete-records empties " + e + " so setup-entity can make a schema change OBP only allows on an empty entity.", Optional: true},
		)
	}
	return scopes
}

// DeclarePlatformApp declares the Scopes the app needs. OBP refuses it (404)
// until an administrator has marked the app's Consumer as a Platform App, and
// refuses the whole declaration (400) if any Role does not exist, which a record
// Role does not until its entity has been created.
func (c *Client) DeclarePlatformApp(scopes []Scope) (*PlatformApp, error) {
	body := struct {
		RequiredScopes []Scope `json:"required_scopes"`
	}{scopes}
	var out PlatformApp
	if err := c.do("PUT", c.baseURL+"/obp/v7.0.0/consumers/current/platform-app", body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeclareFor declares the Scopes for those of entities that already exist in
// the client's space, so that a record Role for an entity setup-entity has not
// created yet does not make OBP refuse the whole declaration. Listing the
// entities needs CanGetDynamicEntityDefinitions; until that is granted only the
// definition Roles are declared.
func (c *Client) DeclareFor(entities []string) (*PlatformApp, error) {
	ids, err := c.DynamicEntityIDs()
	if err != nil {
		ids = nil
	}
	var existing []string
	for _, e := range entities {
		if _, ok := ids[e]; ok {
			existing = append(existing, e)
		}
	}
	return c.DeclarePlatformApp(c.RequiredScopes(existing))
}

// ScopeNames formats Scopes as Role@bank for log messages.
func ScopeNames(scopes []Scope) []string {
	names := make([]string, len(scopes))
	for i, s := range scopes {
		names[i] = s.RoleName + "@" + s.BankID
	}
	return names
}
