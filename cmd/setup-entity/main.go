// Command setup-entity creates (or updates) the `*_on_chain` dynamic entities in
// OBP from their JSON definitions, in the space named by OBP_ENTITY_SPACE_ID
// (default the `ogcr` bank; empty for system level), via the v7.0.0 management
// API. Requires the calling user to have the CanCreateDynamicEntityDefinition /
// CanUpdateDynamicEntityDefinition roles at that bank id (SYS for system level).
// Idempotent: an entity that already exists is updated in place (PUT), so a
// stale schema (e.g. the old CarbonProjectNFT-shaped parcel_on_chain) is fixed.
//
// Each file keeps the `{"<entity_name>": {<schema>}}` shape; it is turned into
// the v7.0.0 body (entity_name + schema) here.
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
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey, cfg.EntitySpaceID)

	slog.Info("applying entity definitions", "space", client.Space())
	existing, err := client.DynamicEntityIDs()
	if err != nil {
		slog.Error("failed to list dynamic entities", "space", client.Space(), "err", err)
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

	if len(failed) > 0 {
		slog.Error("some entities were not applied", "failed", failed)
		fmt.Printf("done with errors: %v\n", failed)
		os.Exit(1)
	}
	fmt.Println("done")
}
