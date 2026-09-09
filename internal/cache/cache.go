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
	CreditBatchEntity   = "carbon_credit_batch_on_chain"
	CreditBalanceEntity = "carbon_credit_balance_on_chain"
	SyncStatusEntity    = "chain_sync_status"
)

// Run status values for SyncStatus.RunStatus.
const (
	RunStatusOK      = "ok"
	RunStatusPartial = "partial"
)

// SyncStatus is the liveness record written at the end of every run. The other
// entities describe the chain; this one describes the mirror. It exists because
// mint data alone cannot answer "is this connection alive": on a chain where
// nothing has been minted for a week, the newest record looks identical whether
// the mirror ran a second ago or died a week ago.
type SyncStatus struct {
	SyncKey         string `json:"sync_key"`
	ChainID         uint64 `json:"chain_id"`
	HeadBlock       uint64 `json:"head_block"`
	SyncedAt        string `json:"synced_at"`
	RunStatus       string `json:"run_status"`
	MirroredTypes   string `json:"mirrored_types"`
	IntervalSeconds int    `json:"interval_seconds"`
	ErrorCount      int    `json:"error_count"`

	ParcelCount        int `json:"parcel_count"`
	ActivityCount      int `json:"activity_count"`
	CertificationCount int `json:"certification_count"`
	CreditBatchCount   int `json:"credit_batch_count"`
	CreditBalanceCount int `json:"credit_balance_count"`
}

// SyncKeyForChain is the business key: one record per chain, overwritten each run.
func SyncKeyForChain(chainID uint64) string {
	return fmt.Sprintf("chain-%d", chainID)
}

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

// UpsertCreditBatch mirrors a CarbonCreditBatchNFT into
// carbon_credit_batch_on_chain, keyed by batch_key ("<activity_nft_id>:<credit_type>").
// The token_id would be the more obvious key, but it is an integer field and the
// match is done with a server-side query filter, so a string key is the safe
// choice. batch_key is equally unique: the contract permits one batch per pair.
func UpsertCreditBatch(client *obp.Client, b *eth.OnChainCreditBatch) (string, error) {
	return upsert(client, CreditBatchEntity, "batch_key", b.BatchKey, b)
}

// UpsertCreditBalance mirrors one address's CarbonCredit balance into
// carbon_credit_balance_on_chain (keyed by owner_address).
func UpsertCreditBalance(client *obp.Client, b *eth.OnChainCreditBalance) (string, error) {
	return upsert(client, CreditBalanceEntity, "owner_address", b.OwnerAddress, b)
}

// UpsertSyncStatus records the outcome of a run in chain_sync_status, keyed by
// sync_key so each chain keeps exactly one current record.
func UpsertSyncStatus(client *obp.Client, st *SyncStatus) (string, error) {
	return upsert(client, SyncStatusEntity, "sync_key", st.SyncKey, st)
}
