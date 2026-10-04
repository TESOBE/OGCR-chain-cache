#!/usr/bin/env bash
# Create (or update in place) the *_on_chain dynamic entities, and
# chain_sync_status, in OBP from their definitions in entities/, in the
# OBP_ENTITY_SPACE_ID space. Safe to re-run. After scripts/delete-records.sh
# this applies schema changes OBP only allows on an empty entity; then
# `make run` refills the records from the chain.
#
#   scripts/setup-entities.sh [-dir entities]
set -euo pipefail

cd "$(dirname "$0")/.."
go run ./cmd/setup-entity "$@"
