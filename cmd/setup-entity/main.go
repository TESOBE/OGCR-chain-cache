// Command setup-entity creates (or updates) the three `*_on_chain` system
// dynamic entities in OBP from their JSON definitions. Requires the calling user
// to have the CanCreateSystemLevelDynamicEntity / CanUpdateSystemLevelDynamicEntity
// roles. Idempotent: an entity that already exists is updated in place (PUT), so
// a stale schema (e.g. the old CarbonProjectNFT-shaped parcel_on_chain) is fixed.
//
//	setup-entity [-dir entities] [-version v4.0.0]
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
}

func main() {
	dir := flag.String("dir", "entities", "directory holding the entity definition JSON files")
	version := flag.String("version", "v4.0.0", "OBP API version for the management endpoint")
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey)

	existing, err := client.SystemDynamicEntityIDs(*version)
	if err != nil {
		slog.Error("failed to list dynamic entities", "err", err)
		os.Exit(1)
	}

	// Resilient: apply every definition independently and keep going on failure,
	// so a blocked entity (e.g. parcel_on_chain needing CanUpdateSystemLevel-
	// DynamicEntity) doesn't stop the others from being created. Report a summary
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
		var definition map[string]any
		if err := json.Unmarshal(raw, &definition); err != nil {
			slog.Error("invalid entity definition JSON", "path", path, "err", err)
			failed = append(failed, file)
			continue
		}
		name := entityName(definition)
		if name == "" {
			slog.Error("definition has no top-level entity name", "path", path)
			failed = append(failed, file)
			continue
		}

		if id, ok := existing[name]; ok {
			if _, err := client.UpdateSystemDynamicEntity(id, definition, *version); err != nil {
				slog.Error("failed to update entity, skipping", "entity", name, "err", err)
				failed = append(failed, name)
				continue
			}
			slog.Info("updated existing entity", "entity", name, "id", id)
			continue
		}
		if _, err := client.CreateSystemDynamicEntity(definition, *version); err != nil {
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

// entityName returns the single top-level key of an entity definition.
func entityName(def map[string]any) string {
	for k := range def {
		return k
	}
	return ""
}
