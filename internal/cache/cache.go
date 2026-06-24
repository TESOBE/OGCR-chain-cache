// Package cache mirrors on-chain parcels into the OBP `parcel_on_chain`
// dynamic entity (one record per parcel_id, matched for upsert on parcel_id).
package cache

import (
	"fmt"
	"net/url"

	"github.com/TESOBE/OGCR-chain-cache/internal/eth"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

const (
	Entity  = "parcel_on_chain"
	idField = Entity + "_id"
)

// Upsert creates or updates the parcel_on_chain record for a parcel. Returns a
// short human-readable description of what happened.
func Upsert(client *obp.Client, p *eth.OnChainParcel) (string, error) {
	existing, err := client.GetRecords(Entity, url.Values{"parcel_id": {p.ParcelID}})
	if err != nil {
		return "", fmt.Errorf("lookup %s: %w", p.ParcelID, err)
	}
	if len(existing) > 0 {
		recordID, _ := existing[0][idField].(string)
		if recordID == "" {
			return "", fmt.Errorf("existing record for %s has no %s", p.ParcelID, idField)
		}
		if _, err := client.UpdateRecord(Entity, recordID, p); err != nil {
			return "", err
		}
		return fmt.Sprintf("updated %s (token %d)", p.ParcelID, p.TokenID), nil
	}
	if _, err := client.CreateRecord(Entity, p); err != nil {
		return "", err
	}
	return fmt.Sprintf("created %s (token %d)", p.ParcelID, p.TokenID), nil
}
