// Package obp is a minimal OBP client: DirectLogin auth plus the dynamic-entity
// read and write calls the cacher needs. The auth flow mirrors the tokenizer's
// client (sibling repo OGCR-Chain); the write methods (POST/PUT and entity
// creation) are new — the tokenizer only ever reads.
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

type Client struct {
	baseURL     string
	username    string
	password    string
	consumerKey string
	http        *http.Client

	mu    sync.Mutex
	token string
}

func NewClient(baseURL, username, password, consumerKey string) *Client {
	return &Client{
		baseURL:     strings.TrimRight(baseURL, "/"),
		username:    username,
		password:    password,
		consumerKey: consumerKey,
		http:        &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) authenticate() error {
	req, err := http.NewRequest("POST", c.baseURL+"/my/logins/direct", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf(
		`DirectLogin username="%s",password="%s",consumer_key="%s"`,
		c.username, c.password, c.consumerKey,
	))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("directlogin request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("directlogin failed (%d): %s", resp.StatusCode, body)
	}

	var result struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode token: %w", err)
	}
	if result.Token == "" {
		return fmt.Errorf("empty token in directlogin response")
	}

	c.mu.Lock()
	c.token = result.Token
	c.mu.Unlock()
	return nil
}

func (c *Client) getToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.token
}

// do performs an authenticated request, re-authenticating once on 401. A nil
// body is allowed (GET). out may be nil to discard the response body.
func (c *Client) do(method, fullURL string, body any, out any) error {
	token := c.getToken()
	if token == "" {
		if err := c.authenticate(); err != nil {
			return err
		}
		token = c.getToken()
	}

	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, fullURL, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "DirectLogin token="+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("%s %s failed: %w", method, fullURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		if err := c.authenticate(); err != nil {
			return err
		}
		return c.do(method, fullURL, body, out)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s %s returned %d: %s", method, fullURL, resp.StatusCode, b)
	}

	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) entityURL(entity, suffix string, params url.Values) string {
	u := c.baseURL + "/obp/dynamic-entity/" + entity + suffix
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}

// GetRecords returns dynamic-entity records, optionally filtered by query
// params. OBP wraps lists as {"<entity>_list": [...]}.
func (c *Client) GetRecords(entity string, params url.Values) ([]map[string]any, error) {
	var raw map[string]json.RawMessage
	if err := c.do("GET", c.entityURL(entity, "", params), nil, &raw); err != nil {
		return nil, err
	}
	listJSON, ok := raw[entity+"_list"]
	if !ok {
		return nil, nil
	}
	var list []map[string]any
	if err := json.Unmarshal(listJSON, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) CreateRecord(entity string, record any) (map[string]any, error) {
	var out map[string]any
	err := c.do("POST", c.entityURL(entity, "", nil), record, &out)
	return out, err
}

func (c *Client) UpdateRecord(entity, recordID string, record any) (map[string]any, error) {
	var out map[string]any
	err := c.do("PUT", c.entityURL(entity, "/"+recordID, nil), record, &out)
	return out, err
}

// ── dynamic-entity definitions (one-time management) ────────────────────────

func (c *Client) managementURL(version string) string {
	return fmt.Sprintf("%s/obp/%s/management/system-dynamic-entities", c.baseURL, version)
}

// ListSystemDynamicEntityNames returns the set of existing system dynamic
// entity names, so creation can be made idempotent.
func (c *Client) ListSystemDynamicEntityNames(version string) (map[string]bool, error) {
	var raw struct {
		DynamicEntities []map[string]any `json:"dynamic_entities"`
	}
	if err := c.do("GET", c.managementURL(version), nil, &raw); err != nil {
		return nil, err
	}
	names := make(map[string]bool)
	for _, e := range raw.DynamicEntities {
		for _, k := range []string{"entityName", "entity_name"} {
			if v, ok := e[k].(string); ok && v != "" {
				names[v] = true
			}
		}
	}
	return names, nil
}

func (c *Client) CreateSystemDynamicEntity(definition any, version string) (map[string]any, error) {
	var out map[string]any
	err := c.do("POST", c.managementURL(version), definition, &out)
	return out, err
}
