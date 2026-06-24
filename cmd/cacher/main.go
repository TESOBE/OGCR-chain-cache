// Command cacher reads CarbonProjectNFT data off the OGCR chain (chain id 2025)
// and upserts it into the OBP `parcel_on_chain` dynamic entity.
//
//	cacher                          # mirror every minted parcel (event scan)
//	cacher <parcel_id> [parcel_id…] # only the given parcel_id(s)
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

	reader, err := eth.NewReader(cfg.RPCURL, cfg.ContractAddress)
	if err != nil {
		slog.Error("failed to init chain reader", "err", err)
		os.Exit(1)
	}
	client := obp.NewClient(cfg.OBPURL, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey)

	ctx := context.Background()
	parcelIDs := os.Args[1:]

	var parcels []*eth.OnChainParcel
	if len(parcelIDs) > 0 {
		for _, pid := range parcelIDs {
			ps, err := reader.ByParcelID(ctx, pid)
			if err != nil {
				slog.Error("read by parcel_id failed", "parcel_id", pid, "err", err)
				continue
			}
			parcels = append(parcels, ps...)
		}
	} else {
		parcels, err = reader.ScanMinted(ctx, cfg.FromBlock)
		if err != nil {
			slog.Error("scan minted failed", "err", err)
			os.Exit(1)
		}
	}

	slog.Info("read complete", "chain_id", reader.ChainID(), "parcels", len(parcels))
	for _, p := range parcels {
		msg, err := cache.Upsert(client, p)
		if err != nil {
			slog.Error("upsert failed", "parcel_id", p.ParcelID, "err", err)
			continue
		}
		slog.Info(msg, "parcel_id", p.ParcelID, "token_id", p.TokenID)
	}
}
