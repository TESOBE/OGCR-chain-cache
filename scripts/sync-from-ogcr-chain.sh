#!/usr/bin/env bash
# Mirror the OGCR chain named in .env (RPC_URL) into the OBP *_on_chain dynamic
# entities once (same as `make run`). Re-running is safe: records are upserted
# by business key. Name token types to limit it to those.
#
#   scripts/sync-from-ogcr-chain.sh [parcel] [activity] [certification] [credit]
set -euo pipefail

cd "$(dirname "$0")/.."
go run ./cmd/cacher "$@"
