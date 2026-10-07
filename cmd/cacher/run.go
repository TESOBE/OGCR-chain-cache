package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/TESOBE/OGCR-chain-cache/internal/cache"
	"github.com/TESOBE/OGCR-chain-cache/internal/eth"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// mirrorNames are the mirrors in the order they run and are reported.
var mirrorNames = []string{"parcel", "activity", "certification", "credit_batch", "credit_balance"}

// RunStatusNoChain is a run that could not read the chain head, so it wrote no
// chain_sync_status record. ok and partial are cache.RunStatusOK/Partial.
const RunStatusNoChain = "no_chain"

// Report is what one run did, for the status page.
type Report struct {
	Started, Finished time.Time
	ChainID           uint64
	HeadBlock         uint64
	RunStatus         string
	Mirrors           []MirrorLine
	Errors            []string
	SyncRecorded      bool
	Chain             eth.ChainInfo
}

// MirrorLine is one mirror's part of a Report.
type MirrorLine struct {
	Name   string
	Ran    bool
	Count  int
	Errors int
}

// runner holds what every run needs.
type runner struct {
	reader          *eth.Reader
	client          *obp.Client
	fromBlock       uint64
	intervalSeconds int
	want            map[string]bool
	// For the status page only.
	chainID        uint64
	obpURL, rpcURL string
}

// once mirrors every wanted token type and records the run in chain_sync_status.
func (rn *runner) once(ctx context.Context) *Report {
	r := &run{ctx: ctx, reader: rn.reader, client: rn.client, fromBlock: rn.fromBlock}
	rep := &Report{Started: time.Now(), ChainID: rn.reader.ChainID()}
	slog.Info("cacher run start", "chain_id", rn.reader.ChainID(), "from_block", rn.fromBlock)
	rep.Chain = rn.reader.Info(ctx)
	for _, c := range rep.Chain.Contracts {
		if c.Configured && c.Err == "" && c.CodeBytes == 0 {
			slog.Warn("no contract code at the configured address on this chain", "contract", c.Name, "address", c.Address, "chain_id", rep.ChainID)
		}
	}

	results := map[string]mirrorResult{}
	if rn.want["parcel"] {
		results["parcel"] = r.mirrorParcels()
	}
	if rn.want["activity"] {
		results["activity"] = r.mirrorActivities()
	}
	if rn.want["certification"] {
		results["certification"] = r.mirrorCertifications()
	}
	if rn.want["credit"] {
		results["credit_batch"], results["credit_balance"] = r.mirrorCredits()
	}

	rep.HeadBlock, rep.RunStatus, rep.SyncRecorded = r.recordSyncStatus(rn.intervalSeconds, results)
	for _, name := range mirrorNames {
		res := results[name]
		rep.Mirrors = append(rep.Mirrors, MirrorLine{Name: name, Ran: res.ran, Count: res.count, Errors: res.errors})
	}
	rep.Errors = r.errors
	rep.Finished = time.Now()
	return rep
}

// mirrorResult summarises one mirror so the run can be recorded in
// chain_sync_status. `ran` distinguishes a genuine count of zero from a mirror
// that never executed.
type mirrorResult struct {
	count  int
	errors int
	ran    bool
}

// run is one pass over the chain. It keeps the errors it logs, so the status
// page can show them next to the counts.
type run struct {
	ctx       context.Context
	reader    *eth.Reader
	client    *obp.Client
	fromBlock uint64
	errors    []string
}

// maxRunErrors caps how many error lines a report keeps; the log has them all.
const maxRunErrors = 50

// fail logs an error and keeps it, as one line, for the report.
func (r *run) fail(msg string, args ...any) {
	slog.Error(msg, args...)
	if len(r.errors) >= maxRunErrors {
		return
	}
	line := msg
	for i := 0; i+1 < len(args); i += 2 {
		line += fmt.Sprintf(" %v=%v", args[i], args[i+1])
	}
	r.errors = append(r.errors, line)
}

func (r *run) mirrorParcels() mirrorResult {
	res := mirrorResult{ran: true}
	parcels, err := r.reader.ScanParcels(r.ctx, r.fromBlock)
	if err != nil {
		r.fail("scan parcels failed", "err", err)
		res.errors++
		return res
	}
	slog.Info("parcels read", "count", len(parcels))
	for _, p := range parcels {
		msg, err := cache.UpsertParcel(r.client, p)
		if err != nil {
			r.fail("upsert parcel failed", "parcel_id", p.ParcelID, "err", err)
			res.errors++
			continue
		}
		res.count++
		slog.Info("parcel "+msg, "parcel_id", p.ParcelID, "token_id", p.TokenID)
	}
	return res
}

func (r *run) mirrorActivities() mirrorResult {
	res := mirrorResult{ran: true}
	activities, err := r.reader.ScanActivities(r.ctx, r.fromBlock)
	if err != nil {
		r.fail("scan activities failed", "err", err)
		res.errors++
		return res
	}
	slog.Info("activities read", "count", len(activities))
	for _, a := range activities {
		msg, err := cache.UpsertActivity(r.client, a)
		if err != nil {
			r.fail("upsert activity failed", "activity_id", a.ActivityID, "err", err)
			res.errors++
			continue
		}
		res.count++
		slog.Info("activity "+msg, "activity_id", a.ActivityID, "token_id", a.TokenID)
	}
	return res
}

func (r *run) mirrorCertifications() mirrorResult {
	res := mirrorResult{ran: true}
	certs, err := r.reader.ScanCertifications(r.ctx, r.fromBlock)
	if err != nil {
		r.fail("scan certifications failed", "err", err)
		res.errors++
		return res
	}
	slog.Info("certifications read", "count", len(certs))
	for _, c := range certs {
		msg, err := cache.UpsertCertification(r.client, c)
		if err != nil {
			r.fail("upsert certification failed", "certification_of_compliance_id", c.CertificationOfComplianceID, "err", err)
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
func (r *run) mirrorCredits() (batchRes, balanceRes mirrorResult) {
	var batches []*eth.OnChainCreditBatch

	if r.reader.HasCreditBatch() {
		batchRes.ran = true
		var err error
		batches, err = r.reader.ScanCreditBatches(r.ctx, r.fromBlock)
		if err != nil {
			// Stop rather than fall through. Without the batch list every
			// token-bound account would be mirrored as an ordinary wallet,
			// overwriting a previously correct holder_type with a wrong one.
			r.fail("scan credit batches failed, skipping credit balances too", "err", err)
			batchRes.errors++
			return batchRes, balanceRes
		}
		slog.Info("credit batches read", "count", len(batches))
		for _, b := range batches {
			msg, err := cache.UpsertCreditBatch(r.client, b)
			if err != nil {
				r.fail("upsert credit batch failed", "batch_key", b.BatchKey, "err", err)
				batchRes.errors++
				continue
			}
			batchRes.count++
			slog.Info("credit batch "+msg, "batch_key", b.BatchKey, "token_id", b.TokenID, "credit_balance", b.CreditBalance)
		}
	} else {
		slog.Info("skipping credit batches: CREDIT_BATCH_CONTRACT_ADDRESS not set")
	}

	if !r.reader.HasCredit() {
		slog.Info("skipping credit balances: CREDIT_CONTRACT_ADDRESS not set")
		return batchRes, balanceRes
	}
	balanceRes.ran = true
	balances, err := r.reader.ScanCreditBalances(r.ctx, r.fromBlock, batches)
	if err != nil {
		r.fail("scan credit balances failed", "err", err)
		balanceRes.errors++
		return batchRes, balanceRes
	}
	slog.Info("credit balances read", "count", len(balances))
	for _, b := range balances {
		msg, err := cache.UpsertCreditBalance(r.client, b)
		if err != nil {
			r.fail("upsert credit balance failed", "owner_address", b.OwnerAddress, "err", err)
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
//
// It returns the head block, the run status, and whether the record was written.
func (r *run) recordSyncStatus(intervalSeconds int, results map[string]mirrorResult) (uint64, string, bool) {
	head, err := r.reader.HeadBlock(r.ctx)
	if err != nil {
		// No head means the chain went away mid-run. Leave the previous record
		// alone so it ages visibly, rather than writing a status that claims a
		// successful look at a chain we could not reach.
		r.fail("could not read chain head, not recording sync status", "err", err)
		return 0, RunStatusNoChain, false
	}

	var ran []string
	errors := 0
	for _, name := range mirrorNames {
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
		SyncKey:            cache.SyncKeyForChain(r.reader.ChainID()),
		ChainID:            r.reader.ChainID(),
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

	msg, err := cache.UpsertSyncStatus(r.client, st)
	if err != nil {
		r.fail("upsert sync status failed", "sync_key", st.SyncKey, "err", err)
		return head, status, false
	}
	slog.Info("sync status "+msg, "sync_key", st.SyncKey, "head_block", head, "run_status", status, "errors", errors)
	return head, status, true
}
