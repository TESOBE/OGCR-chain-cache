// Command delete-records deletes every record of one `*_on_chain` dynamic entity
// in OBP, keeping its definition and the Role grants for it. The records are a
// cache of the chain, so the cacher writes them again; emptying an entity is what
// lets setup-entity make a change OBP only allows on an empty entity, such as
// retyping a field.
//
// Only entities defined in the -dir directory (entities/*.json) are accepted, so
// it can't be pointed at an entity this tool doesn't own. Without -yes it only
// lists the records.
//
//	delete-records [-dir entities] [-yes] ENTITY
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

func main() {
	dir := flag.String("dir", "entities", "directory holding the entity definition JSON files")
	yes := flag.Bool("yes", false, "delete the records (default: only list them)")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: delete-records [-dir entities] [-yes] ENTITY")
		os.Exit(2)
	}
	entity := flag.Arg(0)

	owned, err := ownedEntities(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ %v\n", err)
		os.Exit(1)
	}
	if !owned[entity] {
		fmt.Fprintf(os.Stderr, "✗ %s is not defined in %s/, so this tool doesn't own it\n", entity, *dir)
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ config error: %v\n", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey, cfg.EntitySpaceID)

	fmt.Printf("%s at %s, %s\n", entity, cfg.OBPURL, client.Space())
	records, err := client.GetRecords(entity, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ could not read the records: %v\n", err)
		os.Exit(1)
	}
	if len(records) == 0 {
		fmt.Println("No records; nothing to delete.")
		return
	}

	idField := entity + "_id"
	for _, r := range records {
		fmt.Printf("  - %v  (%s)\n", r[idField], about(r, idField))
	}
	if !*yes {
		fmt.Printf("%d record(s). Nothing was deleted. Run with -yes to delete them.\n", len(records))
		return
	}

	failed := 0
	for _, r := range records {
		id := fmt.Sprint(r[idField])
		if err := client.DeleteRecord(entity, id); err != nil {
			failed++
			fmt.Printf("✗ %s: %v\n", id, err)
			continue
		}
		fmt.Printf("✓ deleted %s\n", id)
	}
	fmt.Printf("%d deleted, %d failed\n", len(records)-failed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}

// ownedEntities is the set of entity names defined by the JSON files in dir.
func ownedEntities(dir string) (map[string]bool, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil || len(files) == 0 {
		return nil, fmt.Errorf("no entity definitions found in %s/", dir)
	}
	owned := map[string]bool{}
	for _, f := range files {
		owned[strings.TrimSuffix(filepath.Base(f), ".json")] = true
	}
	return owned, nil
}

// about lists a record's other *_id fields, which say what it is about (e.g. its activity_id).
func about(r map[string]any, idField string) string {
	var parts []string
	for k, v := range r {
		if strings.HasSuffix(k, "_id") && k != idField {
			parts = append(parts, fmt.Sprintf("%s=%v", k, v))
		}
	}
	sort.Strings(parts)
	return strings.Join(parts, ", ")
}
