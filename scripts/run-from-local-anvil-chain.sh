#!/usr/bin/env bash
# Mirror the LOCAL anvil chain into the LOCAL OBP once (same as `make run-local`).
# Contract addresses come from the deploy script's output in OGCR-Smart-Contracts,
# loaded over .env, then .env.local (gitignored) over that. Override the deploy
# file's path with LOCAL_CHAIN_ENV=... Name token types to limit it to those.
#
#   scripts/run-from-local-anvil-chain.sh [parcel] [activity] [certification] [credit]
set -euo pipefail

cd "$(dirname "$0")/.."
LOCAL_CHAIN_ENV="${LOCAL_CHAIN_ENV:-../OGCR-Smart-Contracts/deployments/local-anvil.env}"

set -a
if [ -f "$LOCAL_CHAIN_ENV" ]; then
	. "$LOCAL_CHAIN_ENV"
else
	echo "warning: no $LOCAL_CHAIN_ENV; start the local chain first" >&2
fi
if [ -f .env.local ]; then
	. ./.env.local
fi
set +a

go run ./cmd/cacher "$@"
