package main

import (
	"log/slog"
	"time"

	"github.com/TESOBE/OGCR-chain-cache/entities"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
	"github.com/TESOBE/OGCR-chain-cache/internal/setup"
)

// entitySetup is the last attempt to bring the entity definitions in line with
// the ones built into this binary, for the status page.
type entitySetup struct {
	At     time.Time
	Result setup.Result
	// Err is set when the existing definitions could not even be listed.
	Err string
}

// done reports whether nothing is left to apply. Definitions only change with a
// new binary, so once they are all applied there is no need to look again.
func (e *entitySetup) done() bool {
	return e != nil && e.Err == "" && len(e.Result.Failed()) == 0
}

// ensureEntities creates the entities that are missing in OBP and updates the
// ones whose definition has changed, like a migration on start. It needs the
// definition Scopes (CanGet/CanCreate/CanUpdateDynamicEntityDefinition), which
// declarePlatformApp declares whether or not any entity exists yet. A failure
// only logs: the mirrors of entities that do exist still run.
func ensureEntities(client *obp.Client, defs []obp.EntityDefinition) *entitySetup {
	slog.Info("applying entity definitions", "space", client.Space())
	res, err := setup.Apply(client, defs)
	es := &entitySetup{At: time.Now(), Result: res}
	if err != nil {
		slog.Warn("could not apply entity definitions; the Platform App may lack CanGetDynamicEntityDefinitions", "err", err)
		es.Err = err.Error()
		return es
	}
	if failed := res.Failed(); len(failed) > 0 {
		slog.Warn("some entity definitions were not applied", "failed", failed)
	}
	return es
}

// loadDefinitions reads the definitions built into the binary. They are checked
// by the tests, so an error here is a broken build.
func loadDefinitions() ([]obp.EntityDefinition, error) {
	return setup.Load(entities.FS)
}
