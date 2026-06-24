// Command setup-entity creates the `parcel_on_chain` system dynamic entity in
// OBP from its JSON definition. Requires the calling user to have the
// CanCreateSystemLevelDynamicEntity role. Idempotent: skips if it already
// exists.
//
//	setup-entity [-def entities/parcel_on_chain.json] [-version v4.0.0]
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/internal/cache"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

func main() {
	defPath := flag.String("def", "entities/parcel_on_chain.json", "path to the entity definition JSON")
	version := flag.String("version", "v4.0.0", "OBP API version for the management endpoint")
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey)

	names, err := client.ListSystemDynamicEntityNames(*version)
	if err != nil {
		slog.Error("failed to list dynamic entities", "err", err)
		os.Exit(1)
	}
	if names[cache.Entity] {
		slog.Info("entity already exists, nothing to do", "entity", cache.Entity)
		return
	}

	raw, err := os.ReadFile(*defPath)
	if err != nil {
		slog.Error("failed to read entity definition", "path", *defPath, "err", err)
		os.Exit(1)
	}
	var definition map[string]any
	if err := json.Unmarshal(raw, &definition); err != nil {
		slog.Error("invalid entity definition JSON", "err", err)
		os.Exit(1)
	}

	result, err := client.CreateSystemDynamicEntity(definition, *version)
	if err != nil {
		slog.Error("failed to create entity", "err", err)
		os.Exit(1)
	}
	pretty, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("created dynamic entity %q:\n%s\n", cache.Entity, pretty)
}
