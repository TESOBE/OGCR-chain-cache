// Package cache mirrors on-chain tokens into their OBP `*_on_chain` dynamic
// entities. One record per business key, matched for upsert on that key.
package cache

import (
	"fmt"
	"net/url"

	"github.com/TESOBE/OGCR-chain-cache/internal/eth"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// Entity names and their business keys.
const (
	ParcelEntity        = "parcel_on_chain"
	ActivityEntity      = "activity_on_chain"
	CertificationEntity = "certification_on_chain"
)

// Upsert creates or updates one record in `entity`, matching an existing record
// on keyField == keyValue. record is the full body to write. Returns a short
// human-readable description of what happened.
func upsert(client *obp.Client, entity, keyField, keyValue string, record any) (string, error) {
	idField := entity + "_id"
	existing, err := client.GetRecords(entity, url.Values{keyField: {keyValue}})
	if err != nil {
		return "", fmt.Errorf("lookup %s=%s: %w", keyField, keyValue, err)
	}
	if len(existing) > 0 {
		recordID, _ := existing[0][idField].(string)
		if recordID == "" {
			return "", fmt.Errorf("existing %s record for %s=%s has no %s", entity, keyField, keyValue, idField)
		}
		if _, err := client.UpdateRecord(entity, recordID, record); err != nil {
			return "", err
		}
		return "updated " + keyValue, nil
	}
	if _, err := client.CreateRecord(entity, record); err != nil {
		return "", err
	}
	return "created " + keyValue, nil
}

// UpsertParcel mirrors a ParcelNFT into parcel_on_chain (keyed by parcel_id).
func UpsertParcel(client *obp.Client, p *eth.OnChainParcel) (string, error) {
	return upsert(client, ParcelEntity, "parcel_id", p.ParcelID, p)
}

// UpsertActivity mirrors an ActivityNFT into activity_on_chain (keyed by activity_id).
func UpsertActivity(client *obp.Client, a *eth.OnChainActivity) (string, error) {
	return upsert(client, ActivityEntity, "activity_id", a.ActivityID, a)
}

// UpsertCertification mirrors a CertificationNFT into certification_on_chain
// (keyed by certification_of_compliance_id).
func UpsertCertification(client *obp.Client, c *eth.OnChainCertification) (string, error) {
	return upsert(client, CertificationEntity, "certification_of_compliance_id", c.CertificationOfComplianceID, c)
}
