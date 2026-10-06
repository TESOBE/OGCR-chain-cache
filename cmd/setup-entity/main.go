// Command setup-entity creates (or updates) the `*_on_chain` dynamic entities in
// OBP from their JSON definitions, in the space named by OBP_ENTITY_SPACE_ID
// (default the `ogcr` bank; empty for system level), via the v7.0.0 management
// API. Runs as the Platform App, whose Consumer needs the
// CanGetDynamicEntityDefinitions, CanCreateDynamicEntityDefinition and
// CanUpdateDynamicEntityDefinition Scopes at that bank id (SYS for system level).
// It declares the app's Scopes when it is done, including the record Roles of
// the entities it has just created, so an administrator can grant them.
// Idempotent: an entity that already exists is updated in place (PUT), so a
// stale schema (e.g. the old CarbonProjectNFT-shaped parcel_on_chain) is fixed.
//
// Each file keeps the `{"<entity_name>": {<schema>}}` shape; it is turned into
// the v7.0.0 body (entity_name + schema) here. An access flag such as
// `has_public_access` or `auth_mode` may sit in the schema object; it is moved
// out to the top level of the body, where v7.0.0 expects it.
//
//	setup-entity [-dir entities]
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// defFiles are the entity definitions to apply, in dependency order.
var defFiles = []string{
	"parcel_on_chain.json",
	"activity_on_chain.json",
	"certification_on_chain.json",
	"carbon_credit_batch_on_chain.json",
	"carbon_credit_balance_on_chain.json",
	"chain_sync_status.json",
}

func main() {
	dir := flag.String("dir", "entities", "directory holding the entity definition JSON files")
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OIDCIssuer, cfg.OIDCClientID, cfg.OIDCClientSecret, cfg.EntitySpaceID)

	slog.Info("applying entity definitions", "space", client.Space())
	existing, err := client.DynamicEntityIDs()
	if err != nil {
		slog.Error("failed to list dynamic entities", "space", client.Space(), "err", err)
		declarePlatformApp(client)
		os.Exit(1)
	}

	// Resilient: apply every definition independently and keep going on failure,
	// so a blocked entity (e.g. parcel_on_chain needing CanUpdateDynamicEntity-
	// Definition) doesn't stop the others from being created. Report a summary
	// and exit non-zero if any failed, so the blocker isn't silently lost.
	var failed []string
	for _, file := range defFiles {
		path := filepath.Join(*dir, file)
		raw, err := os.ReadFile(path)
		if err != nil {
			slog.Error("failed to read entity definition", "path", path, "err", err)
			failed = append(failed, file)
			continue
		}
		var parsed map[string]map[string]any
		if err := json.Unmarshal(raw, &parsed); err != nil {
			slog.Error("invalid entity definition JSON", "path", path, "err", err)
			failed = append(failed, file)
			continue
		}
		if len(parsed) != 1 {
			slog.Error("definition must have exactly one top-level entity name", "path", path)
			failed = append(failed, file)
			continue
		}
		var definition obp.EntityDefinition
		for name, schema := range parsed {
			definition = obp.EntityDefinition{EntityName: name, Schema: schema}
		}
		name := definition.EntityName
		if err := liftAccessFlags(&definition); err != nil {
			slog.Error("invalid entity definition", "path", path, "err", err)
			failed = append(failed, file)
			continue
		}

		if id, ok := existing[name]; ok {
			if _, err := client.UpdateDynamicEntity(id, definition); err != nil {
				slog.Error("failed to update entity, skipping", "entity", name, "err", err)
				failed = append(failed, name)
				continue
			}
			slog.Info("updated existing entity", "entity", name, "id", id)
			continue
		}
		if _, err := client.CreateDynamicEntity(definition); err != nil {
			slog.Error("failed to create entity, skipping", "entity", name, "err", err)
			failed = append(failed, name)
			continue
		}
		slog.Info("created entity", "entity", name)
	}

	declarePlatformApp(client)

	if len(failed) > 0 {
		slog.Error("some entities were not applied", "failed", failed)
		fmt.Printf("done with errors: %v\n", failed)
		os.Exit(1)
	}
	fmt.Println("done")
}

// declarePlatformApp declares the Scopes the Platform App needs for every
// entity that now exists, and says which are not granted yet. A failure only
// warns: the entities were applied or not regardless.
func declarePlatformApp(client *obp.Client) {
	var entities []string
	for _, f := range defFiles {
		entities = append(entities, strings.TrimSuffix(f, ".json"))
	}
	app, err := client.DeclareFor(entities)
	if err != nil {
		slog.Warn("could not declare the Platform App's Scopes; an administrator must mark this Consumer as a Platform App", "err", err)
		return
	}
	if missing := app.Missing(); len(missing) > 0 {
		slog.Warn("Platform App is missing Scopes; an administrator must grant them to its Consumer", "consumer_id", app.ConsumerID, "missing", obp.ScopeNames(missing))
		return
	}
	slog.Info("Platform App holds every Scope it declared", "consumer_id", app.ConsumerID)
}

// liftAccessFlags moves has_public_access and auth_mode out of the definition's
// schema, where the entities/*.json files keep them, into the body fields
// v7.0.0 reads.
func liftAccessFlags(d *obp.EntityDefinition) error {
	if v, ok := d.Schema["has_public_access"]; ok {
		b, ok := v.(bool)
		if !ok {
			return fmt.Errorf("has_public_access must be true or false, got %v", v)
		}
		d.HasPublicAccess = b
		delete(d.Schema, "has_public_access")
	}
	if v, ok := d.Schema["auth_mode"]; ok {
		m, ok := v.(string)
		if !ok || (m != "ApplicationOnly" && m != "UserOrApplication") {
			return fmt.Errorf("auth_mode must be ApplicationOnly or UserOrApplication, so the Platform App can write records; got %v", v)
		}
		d.AuthMode = m
		delete(d.Schema, "auth_mode")
	}
	return nil
}
