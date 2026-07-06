// Command cacher reads the OGCR NFT family off the chain (chain id 2025) and
// upserts each token into its OBP `*_on_chain` dynamic entity.
//
//	cacher                       # mirror all three token types
//	cacher parcel                # only ParcelNFT       -> parcel_on_chain
//	cacher activity              # only ActivityNFT     -> activity_on_chain
//	cacher certification         # only CertificationNFT-> certification_on_chain
//
// Multiple types may be listed. Re-running is safe (records are upserted by
// business key).
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/internal/cache"
	"github.com/TESOBE/OGCR-chain-cache/internal/eth"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}

	reader, err := eth.NewReader(cfg.RPCURL, cfg.ParcelContractAddress, cfg.ActivityContractAddress, cfg.CertificationContractAddress)
	if err != nil {
		slog.Error("failed to init chain reader", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey)

	// Which token types to mirror (default: all).
	want := map[string]bool{"parcel": true, "activity": true, "certification": true}
	if args := os.Args[1:]; len(args) > 0 {
		want = map[string]bool{}
		for _, a := range args {
			switch a {
			case "parcel", "activity", "certification":
				want[a] = true
			default:
				slog.Error("unknown token type (want parcel|activity|certification)", "arg", a)
				os.Exit(1)
			}
		}
	}

	ctx := context.Background()
	slog.Info("cacher start", "chain_id", reader.ChainID(), "from_block", cfg.FromBlock)

	if want["parcel"] {
		mirrorParcels(ctx, reader, client, cfg.FromBlock)
	}
	if want["activity"] {
		mirrorActivities(ctx, reader, client, cfg.FromBlock)
	}
	if want["certification"] {
		mirrorCertifications(ctx, reader, client, cfg.FromBlock)
	}
}

func mirrorParcels(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) {
	parcels, err := reader.ScanParcels(ctx, fromBlock)
	if err != nil {
		slog.Error("scan parcels failed", "err", err)
		return
	}
	slog.Info("parcels read", "count", len(parcels))
	for _, p := range parcels {
		msg, err := cache.UpsertParcel(client, p)
		if err != nil {
			slog.Error("upsert parcel failed", "parcel_id", p.ParcelID, "err", err)
			continue
		}
		slog.Info("parcel "+msg, "parcel_id", p.ParcelID, "token_id", p.TokenID)
	}
}

func mirrorActivities(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) {
	activities, err := reader.ScanActivities(ctx, fromBlock)
	if err != nil {
		slog.Error("scan activities failed", "err", err)
		return
	}
	slog.Info("activities read", "count", len(activities))
	for _, a := range activities {
		msg, err := cache.UpsertActivity(client, a)
		if err != nil {
			slog.Error("upsert activity failed", "activity_id", a.ActivityID, "err", err)
			continue
		}
		slog.Info("activity "+msg, "activity_id", a.ActivityID, "token_id", a.TokenID)
	}
}

func mirrorCertifications(ctx context.Context, reader *eth.Reader, client *obp.Client, fromBlock uint64) {
	certs, err := reader.ScanCertifications(ctx, fromBlock)
	if err != nil {
		slog.Error("scan certifications failed", "err", err)
		return
	}
	slog.Info("certifications read", "count", len(certs))
	for _, c := range certs {
		msg, err := cache.UpsertCertification(client, c)
		if err != nil {
			slog.Error("upsert certification failed", "certification_of_compliance_id", c.CertificationOfComplianceID, "err", err)
			continue
		}
		slog.Info("certification "+msg, "certification_of_compliance_id", c.CertificationOfComplianceID, "token_id", c.TokenID)
	}
}
