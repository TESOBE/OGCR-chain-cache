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
// With -serve it keeps running instead: it mirrors every SYNC_INTERVAL_SECONDS
// (default 30) and serves a status page and /health on the given address.
//
//	cacher -serve 127.0.0.1:8766 [types...]
//
// Before mirroring it creates any of its entities that are missing in OBP and
// updates any whose definition has changed (AUTO_SETUP_ENTITIES=false turns
// this off), so a new version brings its own schema with it.
//
// The credit contracts are optional. With neither CREDIT_BATCH_CONTRACT_ADDRESS
// nor CREDIT_CONTRACT_ADDRESS set, a default run skips them and says so; asking
// for `credit` explicitly is then an error, so a forgotten address is not
// mistaken for "nothing to mirror".
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"

	"github.com/TESOBE/OGCR-chain-cache/config"
	"github.com/TESOBE/OGCR-chain-cache/internal/cache"
	"github.com/TESOBE/OGCR-chain-cache/internal/eth"
	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// tokenTypes are the values accepted as command-line arguments.
var tokenTypes = []string{"parcel", "activity", "certification", "credit"}

func main() {
	serveAddr := flag.String("serve", "", "keep running: mirror on a loop and serve a status page on this address (e.g. 127.0.0.1:8766)")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: cacher [-serve addr] [%s ...]\n", strings.Join(tokenTypes, "|"))
		flag.PrintDefaults()
	}
	flag.Parse()

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
	client := obp.NewClient(cfg.OBPURL, cfg.OIDCIssuer, cfg.OIDCClientID, cfg.OIDCClientSecret, cfg.EntitySpaceID)

	// Which token types to mirror (default: all).
	want := map[string]bool{}
	explicit := flag.NArg() > 0
	if explicit {
		for _, a := range flag.Args() {
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

	var defs []obp.EntityDefinition
	if cfg.AutoSetupEntities {
		if defs, err = loadDefinitions(); err != nil {
			slog.Error("built-in entity definitions are broken", "err", err)
			os.Exit(1)
		}
	}

	rn := &runner{
		reader:          reader,
		client:          client,
		fromBlock:       cfg.FromBlock,
		intervalSeconds: cfg.IntervalSeconds,
		want:            want,
		chainID:         reader.ChainID(),
		obpURL:          cfg.OBPURL,
		rpcURL:          cfg.RPCURL,
		defs:            defs,
	}

	if *serveAddr == "" {
		// Declared first, so the definition Scopes are asked for even when no
		// entity exists yet, and again after creating one, so its record Scopes
		// are asked for too.
		declarePlatformApp(client)
		if defs != nil && ensureEntities(client, defs).Result.Created() {
			declarePlatformApp(client)
		}
		rn.once(context.Background())
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := serve(ctx, rn, *serveAddr); err != nil {
		slog.Error("serve failed", "err", err)
		os.Exit(1)
	}
}

// declarePlatformApp declares the Scopes this app needs and says which are not
// granted yet. A failure only warns: the run goes on, and any write a missing
// Scope blocks fails on its own with OBP's message.
func declarePlatformApp(client *obp.Client) (*obp.PlatformApp, error) {
	app, err := client.DeclareFor(cache.Entities)
	if err != nil {
		slog.Warn("could not declare the Platform App's Scopes; an administrator must mark this Consumer as a Platform App", "err", err)
		return nil, err
	}
	if missing := app.Missing(); len(missing) > 0 {
		slog.Warn("Platform App is missing Scopes; an administrator must grant them to its Consumer", "consumer_id", app.ConsumerID, "missing", obp.ScopeNames(missing))
	}
	return app, nil
}
