// Package setup brings the OBP dynamic entity definitions in line with the
// JSON files in entities/: it creates the ones that are missing and updates the
// ones whose definition has changed, leaving the rest alone. It never deletes a
// definition or a record. setup-entity runs it on demand, and the cacher on
// start, like a migration.
//
// Each file keeps the `{"<entity_name>": {<schema>}}` shape; it is turned into
// the v7.0.0 body (entity_name + schema) here. An access flag such as
// `has_public_access` or `auth_mode` may sit in the schema object; it is moved
// out to the top level of the body, where v7.0.0 expects it.
package setup

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"reflect"
	"strings"

	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// Files are the entity definitions to apply, in dependency order.
var Files = []string{
	"parcel_on_chain.json",
	"activity_on_chain.json",
	"certification_on_chain.json",
	"carbon_credit_batch_on_chain.json",
	"carbon_credit_balance_on_chain.json",
	"chain_sync_status.json",
}

// Names are the entity names the Files define, in the same order.
func Names() []string {
	names := make([]string, len(Files))
	for i, f := range Files {
		names[i] = strings.TrimSuffix(f, ".json")
	}
	return names
}

// Load reads and parses every file in Files from fsys.
func Load(fsys fs.FS) ([]obp.EntityDefinition, error) {
	defs := make([]obp.EntityDefinition, 0, len(Files))
	for _, file := range Files {
		raw, err := fs.ReadFile(fsys, file)
		if err != nil {
			return nil, fmt.Errorf("read entity definition: %w", err)
		}
		def, err := parse(raw)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		defs = append(defs, def)
	}
	return defs, nil
}

func parse(raw []byte) (obp.EntityDefinition, error) {
	var parsed map[string]map[string]any
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return obp.EntityDefinition{}, fmt.Errorf("invalid entity definition JSON: %w", err)
	}
	if len(parsed) != 1 {
		return obp.EntityDefinition{}, fmt.Errorf("definition must have exactly one top-level entity name")
	}
	var def obp.EntityDefinition
	for name, schema := range parsed {
		def = obp.EntityDefinition{EntityName: name, Schema: schema}
	}
	if err := liftAccessFlags(&def); err != nil {
		return obp.EntityDefinition{}, err
	}
	return def, nil
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

// Actions an Outcome reports.
const (
	Created   = "created"
	Updated   = "updated"
	Unchanged = "unchanged"
	Failed    = "failed"
)

// Outcome is what Apply did with one entity.
type Outcome struct {
	Entity string
	Action string
	// Err is set when Action is Failed.
	Err string
}

// Result is what Apply did with every entity.
type Result struct {
	Outcomes []Outcome
}

// Failed returns the names of the entities that could not be applied.
func (r Result) Failed() []string {
	var failed []string
	for _, o := range r.Outcomes {
		if o.Action == Failed {
			failed = append(failed, o.Entity)
		}
	}
	return failed
}

// Created reports whether any entity was created, which adds record Roles the
// Platform App should declare.
func (r Result) Created() bool {
	for _, o := range r.Outcomes {
		if o.Action == Created {
			return true
		}
	}
	return false
}

// Apply creates each definition that does not exist in the client's space and
// updates each one that differs from what OBP holds. It goes on past a failed
// entity, so one blocked entity does not stop the others; the error is only for
// not being able to list the existing definitions at all, which needs
// CanGetDynamicEntityDefinitions.
func Apply(client *obp.Client, defs []obp.EntityDefinition) (Result, error) {
	live, err := client.DynamicEntities()
	if err != nil {
		return Result{}, fmt.Errorf("list dynamic entities in %s: %w", client.Space(), err)
	}
	var res Result
	for _, def := range defs {
		name := def.EntityName
		existing, ok := live[name]
		switch {
		case !ok:
			if _, err := client.CreateDynamicEntity(def); err != nil {
				slog.Error("failed to create entity", "entity", name, "err", err)
				res.Outcomes = append(res.Outcomes, Outcome{Entity: name, Action: Failed, Err: err.Error()})
				continue
			}
			slog.Info("created entity", "entity", name)
			res.Outcomes = append(res.Outcomes, Outcome{Entity: name, Action: Created})
		case UpToDate(def, existing):
			res.Outcomes = append(res.Outcomes, Outcome{Entity: name, Action: Unchanged})
		default:
			if _, err := client.UpdateDynamicEntity(existing.DynamicEntityID, def); err != nil {
				// OBP allows some schema changes only on an empty entity;
				// scripts/delete-records.sh empties it, then `make run` refills it.
				slog.Error("failed to update entity", "entity", name, "id", existing.DynamicEntityID, "records", existing.RecordCount, "err", err)
				res.Outcomes = append(res.Outcomes, Outcome{Entity: name, Action: Failed, Err: err.Error()})
				continue
			}
			slog.Info("updated entity", "entity", name, "id", existing.DynamicEntityID)
			res.Outcomes = append(res.Outcomes, Outcome{Entity: name, Action: Updated})
		}
	}
	return res, nil
}

// UpToDate reports whether OBP already holds def: the same access flags and the
// same schema, compared as JSON values so key order does not matter.
func UpToDate(def obp.EntityDefinition, live obp.LiveEntity) bool {
	authMode := def.AuthMode
	if authMode == "" {
		authMode = "UserOnly" // OBP's default when the body leaves it out
	}
	if def.HasPersonalEntity != live.HasPersonalEntity || def.HasPublicAccess != live.HasPublicAccess || authMode != live.AuthMode {
		return false
	}
	return reflect.DeepEqual(normalize(def.Schema), normalize(live.Schema))
}

// normalize round-trips v through JSON so numbers and nested values have the
// same Go types on both sides of a comparison.
func normalize(v any) any {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var out any
	if err := json.Unmarshal(b, &out); err != nil {
		return nil
	}
	return out
}
