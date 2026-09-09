// Command cacher reads the OGCR token family off the chain (chain id 2025) and
// upserts each token into its OBP `*_on_chain` dynamic entity.
//
//	cacher                       # mirror everything that is configured
//	cacher parcel                # only ParcelNFT            -> parcel_on_chain
//	cacher activity              # only ActivityNFT          -> activity_on_chain
//	cacher certification         # only CertificationNFT     -> certification_on_chain
//	cacher credit                # CarbonCreditBatchNFT      -> carbon_credit_batch_on_chain
//	                             # and CarbonCredit balances -> carbon_credit_balance_on_chain
//
// Multiple types may be listed. Re-running is safe (records are upserted by
// business key).
//
// The credit contracts are optional. With neither CREDIT_BATCH_CONTRACT_ADDRESS
// nor CREDIT_CONTRACT_ADDRESS set, a default run skips them and says so; asking
// for `credit` explicitly is then an error, so a forgotten address is not
// mistaken for "nothing to mirror".
package main

import (
	"context"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/internal/cache"
	"github.com/TESOBE/OGCR-chain-cache/internal/eth"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// tokenTypes are the values accepted as command-line arguments.
var tokenTypes = []string{"parcel", "activity", "certification", "credit"}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}

	reader, err := eth.NewReader(cfg.RPCURL, eth.Addresses{
		Parcel:        cfg.ParcelContractAddress,
		Activity:      cfg.ActivityContractAddress,
		Certification: cfg.CertificationContractAddress,
		CreditBatch:   cfg.CreditBatchContractAddress,
		Credit:        cfg.CreditContractAddress,
	})
	if err != nil {
		slog.Error("failed to init chain reader", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey)

	// Which token types to mirror (default: all).
	want := map[string]bool{}
	explicit := len(os.Args) > 1
	if explicit {
		for _, a := range os.Args[1:] {
			if !slices.Contains(tokenTypes, a) {
				slog.Error("unknown token type", "arg", a, "want", tokenTypes)
				os.Exit(1)
			}
			want[a] = true
		}
	} else {
		for _, t := range tokenTypes {
			want[t] = true
		}
	}

	// An explicit `credit` request with no contract configured is a mistake
	// worth failing on; the same gap on a default run is just a narrower run.
	if want["credit"] && !reader.HasCreditBatch() && !reader.HasCredit() {
		if explicit {
			slog.Error("credit mirroring requested but neither CREDIT_BATCH_CONTRACT_ADDRESS nor CREDIT_CONTRACT_ADDRESS is set")
			os.Exit(1)
		}
		slog.Info("skipping credit mirror: no credit contract addresses configured")
		want["credit"] = false
	}

	ctx := context.Background()
	slog.Info("cacher start", "chain_id", reader.ChainID(), "from_block", cfg.FromBlock)

	results := map[string]mirrorResult{}
	if want["parcel"] {
		results["parcel"] = mirrorParcels(ctx, reader, client, cfg.FromBlock)
	}
	if want["activity"] {
		results["activity"] = mirrorActivities(ctx, reader, client, cfg.FromBlock)
	}
	if want["certification"] {
		results["certification"] = mirrorCertifications(ctx, reader, client, cfg.FromBlock)
	}
	if want["credit"] {
		results["credit_batch"], results["credit_balance"] = mirrorCredits(ctx, reader, client, cfg.FromBlock)
	}

	recordSyncStatus(ctx, reader, client, cfg.IntervalSeconds, results)
}

// mirrorResult summarises one mirror so the run can be recorded in
// chain_sync_status. `ran` distinguishes a genuine count of zero from a mirror
// that never executed.
type mirrorResult struct {
	count  int
	errors int
	ran    bool
}

func mirrorParcels(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) mirrorResult {
	res := mirrorResult{ran: true}
	parcels, err := reader.ScanParcels(ctx, fromBlock)
	if err != nil {
		slog.Error("scan parcels failed", "err", err)
		res.errors++
		return res
	}
	slog.Info("parcels read", "count", len(parcels))
	for _, p := range parcels {
		msg, err := cache.UpsertParcel(client, p)
		if err != nil {
			slog.Error("upsert parcel failed", "parcel_id", p.ParcelID, "err", err)
			res.errors++
			continue
		}
		res.count++
		slog.Info("parcel "+msg, "parcel_id", p.ParcelID, "token_id", p.TokenID)
	}
	return res
}

func mirrorActivities(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) mirrorResult {
	res := mirrorResult{ran: true}
	activities, err := reader.ScanActivities(ctx, fromBlock)
	if err != nil {
		slog.Error("scan activities failed", "err", err)
		res.errors++
		return res
	}
	slog.Info("activities read", "count", len(activities))
	for _, a := range activities {
		msg, err := cache.UpsertActivity(client, a)
		if err != nil {
			slog.Error("upsert activity failed", "activity_id", a.ActivityID, "err", err)
			res.errors++
			continue
		}
		res.count++
		slog.Info("activity "+msg, "activity_id", a.ActivityID, "token_id", a.TokenID)
	}
	return res
}

func mirrorCertifications(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) mirrorResult {
	res := mirrorResult{ran: true}
	certs, err := reader.ScanCertifications(ctx, fromBlock)
	if err != nil {
		slog.Error("scan certifications failed", "err", err)
		res.errors++
		return res
	}
	slog.Info("certifications read", "count", len(certs))
	for _, c := range certs {
		msg, err := cache.UpsertCertification(client, c)
		if err != nil {
			slog.Error("upsert certification failed", "certification_of_compliance_id", c.CertificationOfComplianceID, "err", err)
			res.errors++
			continue
		}
		res.count++
		slog.Info("certification "+msg, "certification_of_compliance_id", c.CertificationOfComplianceID, "token_id", c.TokenID)
	}
	return res
}

// mirrorCredits handles both halves of the carbon-credit layer. Batches are
// scanned first and passed to the balance scan, which uses them to tell a
// batch's token-bound account apart from an ordinary wallet without walking the
// batch log a second time.
func mirrorCredits(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) (batchRes, balanceRes mirrorResult) {
	var batches []*eth.OnChainCreditBatch

	if reader.HasCreditBatch() {
		batchRes.ran = true
		var err error
		batches, err = reader.ScanCreditBatches(ctx, fromBlock)
		if err != nil {
			// Stop rather than fall through. Without the batch list every
			// token-bound account would be mirrored as an ordinary wallet,
			// overwriting a previously correct holder_type with a wrong one.
			slog.Error("scan credit batches failed, skipping credit balances too", "err", err)
			batchRes.errors++
			return batchRes, balanceRes
		}
		slog.Info("credit batches read", "count", len(batches))
		for _, b := range batches {
			msg, err := cache.UpsertCreditBatch(client, b)
			if err != nil {
				slog.Error("upsert credit batch failed", "batch_key", b.BatchKey, "err", err)
				batchRes.errors++
				continue
			}
			batchRes.count++
			slog.Info("credit batch "+msg, "batch_key", b.BatchKey, "token_id", b.TokenID, "credit_balance", b.CreditBalance)
		}
	} else {
		slog.Info("skipping credit batches: CREDIT_BATCH_CONTRACT_ADDRESS not set")
	}

	if !reader.HasCredit() {
		slog.Info("skipping credit balances: CREDIT_CONTRACT_ADDRESS not set")
		return batchRes, balanceRes
	}
	balanceRes.ran = true
	balances, err := reader.ScanCreditBalances(ctx, fromBlock, batches)
	if err != nil {
		slog.Error("scan credit balances failed", "err", err)
		balanceRes.errors++
		return batchRes, balanceRes
	}
	slog.Info("credit balances read", "count", len(balances))
	for _, b := range balances {
		msg, err := cache.UpsertCreditBalance(client, b)
		if err != nil {
			slog.Error("upsert credit balance failed", "owner_address", b.OwnerAddress, "err", err)
			balanceRes.errors++
			continue
		}
		balanceRes.count++
		slog.Info("credit balance "+msg, "owner_address", b.OwnerAddress, "balance", b.Balance, "holder_type", b.HolderType)
	}
	return batchRes, balanceRes
}

// recordSyncStatus writes the liveness record for this run. It is what lets a
// consumer tell a quiet chain from a dead mirror, so it is written even when
// some mirrors failed: an honest "partial" is more useful than no record, which
// would look identical to the cacher never having run.
func recordSyncStatus(
	ctx context.Context,
	reader *eth.Reader,
	client *obp.Client,
	intervalSeconds int,
	results map[string]mirrorResult,
) {
	head, err := reader.HeadBlock(ctx)
	if err != nil {
		// No head means the chain went away mid-run. Leave the previous record
		// alone so it ages visibly, rather than writing a status that claims a
		// successful look at a chain we could not reach.
		slog.Error("could not read chain head, not recording sync status", "err", err)
		return
	}

	var ran []string
	errors := 0
	for _, name := range []string{"parcel", "activity", "certification", "credit_batch", "credit_balance"} {
		r := results[name]
		if r.ran {
			ran = append(ran, name)
		}
		errors += r.errors
	}

	status := cache.RunStatusOK
	if errors > 0 {
		status = cache.RunStatusPartial
	}

	st := &cache.SyncStatus{
		SyncKey:            cache.SyncKeyForChain(reader.ChainID()),
		ChainID:            reader.ChainID(),
		HeadBlock:          head,
		SyncedAt:           time.Now().UTC().Format(time.RFC3339),
		RunStatus:          status,
		MirroredTypes:      strings.Join(ran, ","),
		IntervalSeconds:    intervalSeconds,
		ErrorCount:         errors,
		ParcelCount:        results["parcel"].count,
		ActivityCount:      results["activity"].count,
		CertificationCount: results["certification"].count,
		CreditBatchCount:   results["credit_batch"].count,
		CreditBalanceCount: results["credit_balance"].count,
	}

	msg, err := cache.UpsertSyncStatus(client, st)
	if err != nil {
		slog.Error("upsert sync status failed", "sync_key", st.SyncKey, "err", err)
		return
	}
	slog.Info("sync status "+msg, "sync_key", st.SyncKey, "head_block", head, "run_status", status, "errors", errors)
}
