// Command setup-entity creates (or updates) the `*_on_chain` dynamic entities in
// OBP from their JSON definitions, in the space named by OBP_ENTITY_SPACE_ID
// (default the `ogcr` bank; empty for system level), via the v7.0.0 management
// API. Runs as the Platform App, whose Consumer needs the
// CanGetDynamicEntityDefinitions, CanCreateDynamicEntityDefinition and
// CanUpdateDynamicEntityDefinition Scopes at that bank id (SYS for system level).
// It declares the app's Scopes when it is done, including the record Roles of
// the entities it has just created, so an administrator can grant them.
// Idempotent: an entity that already exists is updated in place (PUT) when its
// definition differs, so a stale schema (e.g. the old CarbonProjectNFT-shaped
// parcel_on_chain) is fixed.
//
// The cacher does the same on start, so this is only needed to apply the
// definitions without running the cacher. By default it applies the
// definitions built into the binary; -dir reads them from a directory instead.
//
//	setup-entity [-dir entities]
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/entities"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
	"github.com/TESOBE/OGCR-chain-cache/internal/setup"
)

func main() {
	dir := flag.String("dir", "", "directory holding the entity definition JSON files (default: the ones built into this binary)")
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OIDCIssuer, cfg.OIDCClientID, cfg.OIDCClientSecret, cfg.EntitySpaceID)

	var fsys fs.FS = entities.FS
	if *dir != "" {
		fsys = os.DirFS(*dir)
	}
	defs, err := setup.Load(fsys)
	if err != nil {
		slog.Error("failed to load entity definitions", "err", err)
		os.Exit(1)
	}

	slog.Info("applying entity definitions", "space", client.Space())
	res, err := setup.Apply(client, defs)
	if err != nil {
		slog.Error("failed to list dynamic entities", "space", client.Space(), "err", err)
		declarePlatformApp(client)
		os.Exit(1)
	}
	for _, o := range res.Outcomes {
		if o.Action == setup.Unchanged {
			slog.Info("entity up to date", "entity", o.Entity)
		}
	}

	declarePlatformApp(client)

	if failed := res.Failed(); len(failed) > 0 {
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
	app, err := client.DeclareFor(setup.Names())
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
