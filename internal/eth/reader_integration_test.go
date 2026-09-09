package eth

import (
	"context"
	"os"
	"testing"
)

// Integration tests against a real chain. Skipped unless the environment names
// one, so `go test ./...` stays offline by default.
//
// The OGCR chain is a private Besu network, but any EVM node works. The quickest
// fixture is a local anvil seeded by the deploy script in OGCR-Smart-Contracts,
// which stands up all five contracts and mints a coherent set of records:
//
//	anvil --chain-id 2025
//	# in OGCR-Smart-Contracts:
//	forge script script/local/DeployLocalStack.s.sol \
//	  --rpc-url http://127.0.0.1:8545 --broadcast --legacy
//	# it writes deployments/local-anvil.env with every address:
//	set -a; . deployments/local-anvil.env; set +a
//	OGCR_TEST_RPC_URL=$RPC_URL \
//	OGCR_TEST_PARCEL_ADDRESS=$PARCEL_CONTRACT_ADDRESS \
//	OGCR_TEST_ACTIVITY_ADDRESS=$ACTIVITY_CONTRACT_ADDRESS \
//	OGCR_TEST_CERTIFICATION_ADDRESS=$CERTIFICATION_CONTRACT_ADDRESS \
//	OGCR_TEST_CREDIT_BATCH_ADDRESS=$CREDIT_BATCH_CONTRACT_ADDRESS \
//	OGCR_TEST_CREDIT_ADDRESS=$CREDIT_CONTRACT_ADDRESS \
//	go test ./internal/eth -run TestScan -v

// placeholder address for a contract a given test does not exercise. Binding is
// lazy, so an address with no code is harmless as long as it is never scanned.
const unusedAddr = "0x0000000000000000000000000000000000000001"

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// TestScanNFTs covers the three registry-mirroring NFTs.
func TestScanNFTs(t *testing.T) {
	rpcURL := os.Getenv("OGCR_TEST_RPC_URL")
	parcelAddr := os.Getenv("OGCR_TEST_PARCEL_ADDRESS")
	activityAddr := os.Getenv("OGCR_TEST_ACTIVITY_ADDRESS")
	certAddr := os.Getenv("OGCR_TEST_CERTIFICATION_ADDRESS")
	if rpcURL == "" || parcelAddr == "" || activityAddr == "" || certAddr == "" {
		t.Skip("set OGCR_TEST_RPC_URL and the three NFT addresses to run")
	}

	r, err := NewReader(rpcURL, Addresses{
		Parcel:        parcelAddr,
		Activity:      activityAddr,
		Certification: certAddr,
	})
	if err != nil {
		t.Fatalf("NewReader: %v", err)
	}
	ctx := context.Background()

	parcels, err := r.ScanParcels(ctx, 0)
	if err != nil {
		t.Fatalf("ScanParcels: %v", err)
	}
	if len(parcels) == 0 {
		t.Error("no parcels found; the chain needs at least one minted ParcelNFT")
	}
	for _, p := range parcels {
		if p.ParcelID == "" {
			t.Errorf("parcel token %d has no parcel_id, which is its upsert key", p.TokenID)
		}
	}

	activities, err := r.ScanActivities(ctx, 0)
	if err != nil {
		t.Fatalf("ScanActivities: %v", err)
	}
	if len(activities) == 0 {
		t.Error("no activities found; the chain needs at least one minted ActivityNFT")
	}
	for _, a := range activities {
		if a.ActivityID == "" {
			t.Errorf("activity token %d has no activity_id, which is its upsert key", a.TokenID)
		}
	}

	certs, err := r.ScanCertifications(ctx, 0)
	if err != nil {
		t.Fatalf("ScanCertifications: %v", err)
	}
	if len(certs) == 0 {
		t.Error("no certifications found; the chain needs at least one minted CertificationNFT")
	}

	// A certificate points at an activity by NFT id. That join is how the
	// marketplace gets from a certificate back to an OBP activity_id, so it is
	// worth asserting the fixture actually holds together.
	activityByToken := map[uint64]bool{}
	for _, a := range activities {
		activityByToken[a.TokenID] = true
	}
	for _, c := range certs {
		if c.CertificationOfComplianceID == "" {
			t.Errorf("certification token %d has no compliance id, which is its upsert key", c.TokenID)
		}
		if !activityByToken[c.ActivityNftID] {
			t.Errorf("certification %s references activity token %d, which was not found",
				c.CertificationOfComplianceID, c.ActivityNftID)
		}
	}
}

// TestScanCredits covers the carbon-credit layer: batch tokens and the ERC-20
// balances held by their bound accounts and by ordinary wallets.
func TestScanCredits(t *testing.T) {
	rpcURL := os.Getenv("OGCR_TEST_RPC_URL")
	batchAddr := os.Getenv("OGCR_TEST_CREDIT_BATCH_ADDRESS")
	creditAddr := os.Getenv("OGCR_TEST_CREDIT_ADDRESS")
	if rpcURL == "" || batchAddr == "" || creditAddr == "" {
		t.Skip("set OGCR_TEST_RPC_URL, OGCR_TEST_CREDIT_BATCH_ADDRESS and OGCR_TEST_CREDIT_ADDRESS to run")
	}

	r, err := NewReader(rpcURL, Addresses{
		// The NFTs are not exercised here; fall back to a placeholder so the
		// reader still constructs when only the credit addresses are given.
		Parcel:        envOr("OGCR_TEST_PARCEL_ADDRESS", unusedAddr),
		Activity:      envOr("OGCR_TEST_ACTIVITY_ADDRESS", unusedAddr),
		Certification: envOr("OGCR_TEST_CERTIFICATION_ADDRESS", unusedAddr),
		CreditBatch:   batchAddr,
		Credit:        creditAddr,
	})
	if err != nil {
		t.Fatalf("NewReader: %v", err)
	}
	if !r.HasCreditBatch() || !r.HasCredit() {
		t.Fatal("credit contracts should be configured")
	}

	ctx := context.Background()

	batches, err := r.ScanCreditBatches(ctx, 0)
	if err != nil {
		t.Fatalf("ScanCreditBatches: %v", err)
	}
	if len(batches) == 0 {
		t.Fatal("no credit batches found; the chain needs at least one minted batch")
	}

	seen := map[string]bool{}
	for _, b := range batches {
		if b.BatchKey == "" || b.CreditType == "" {
			t.Errorf("batch %d: empty business key or credit type", b.TokenID)
		}
		if seen[b.BatchKey] {
			t.Errorf("batch key %q is not unique; it is the upsert key", b.BatchKey)
		}
		seen[b.BatchKey] = true

		if b.TokenBoundAccount == "" {
			t.Errorf("batch %s: no token-bound account", b.BatchKey)
		}
		if b.CreditBalance == "" {
			t.Errorf("batch %s: balance not read despite a configured credit contract", b.BatchKey)
		}
		if b.BlockNumberBalance == 0 {
			t.Errorf("batch %s: balance has no block height", b.BatchKey)
		}
		if b.TokenURI == "" {
			t.Errorf("batch %s: no token URI", b.BatchKey)
		}
		if b.ChainID != r.ChainID() {
			t.Errorf("batch %s: chain id %d, want %d", b.BatchKey, b.ChainID, r.ChainID())
		}
	}

	balances, err := r.ScanCreditBalances(ctx, 0, batches)
	if err != nil {
		t.Fatalf("ScanCreditBalances: %v", err)
	}
	if len(balances) == 0 {
		t.Fatal("no credit balances found; the chain needs at least one credit holder")
	}

	// Every batch's token-bound account must appear, classified as such. This is
	// what lets the marketplace tell credits still held in a batch apart from
	// credits withdrawn into a pooled wallet balance.
	byAddress := map[string]*OnChainCreditBalance{}
	for _, b := range balances {
		if b.Balance == "" {
			t.Errorf("balance for %s is empty", b.OwnerAddress)
		}
		if b.BlockNumber == 0 {
			t.Errorf("balance for %s has no block height", b.OwnerAddress)
		}
		byAddress[b.OwnerAddress] = b
	}
	for _, batch := range batches {
		got, ok := byAddress[batch.TokenBoundAccount]
		if !ok {
			t.Errorf("batch %s: token-bound account %s missing from balances", batch.BatchKey, batch.TokenBoundAccount)
			continue
		}
		if got.HolderType != HolderTokenBoundAccount {
			t.Errorf("batch %s: account classified %q, want %q", batch.BatchKey, got.HolderType, HolderTokenBoundAccount)
		}
		if got.BatchTokenID != batch.TokenID {
			t.Errorf("batch %s: account maps to batch %d, want %d", batch.BatchKey, got.BatchTokenID, batch.TokenID)
		}
		if got.Balance != batch.CreditBalance {
			t.Errorf("batch %s: batch balance %s but account balance %s", batch.BatchKey, batch.CreditBalance, got.Balance)
		}
	}
}
