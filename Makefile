.DEFAULT_GOAL := help

help: ## list targets
	@awk 'BEGIN{FS=":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  %-14s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## build both binaries into bin/
	go build -o bin/cacher ./cmd/cacher
	go build -o bin/setup-entity ./cmd/setup-entity

test: ## run tests
	go test ./...

vet: ## go vet
	go vet ./...

setup-entity: ## create/update the *_on_chain dynamic entities in OBP (one-time)
	go run ./cmd/setup-entity

run: ## mirror chain -> OBP (limit types via ARGS="parcel activity certification credit")
	go run ./cmd/cacher $(ARGS)

# ── Local development ───────────────────────────────────────────────────────
# Run against a throwaway anvil chain and a local OBP. Contract addresses come
# from the deploy script's output, so they cannot go stale relative to what is
# actually deployed. An optional .env.local (gitignored) can override anything
# in .env, which is useful when .env points at shared infrastructure. See the
# README.
LOCAL_CHAIN_ENV ?= ../OGCR-Smart-Contracts/deployments/local-anvil.env

define load_local_env
set -a; \
if [ -f $(LOCAL_CHAIN_ENV) ]; then . $(LOCAL_CHAIN_ENV); else echo "warning: no $(LOCAL_CHAIN_ENV); start the local chain first"; fi; \
if [ -f .env.local ]; then . ./.env.local; fi; \
set +a;
endef

setup-entity-local: ## create/update the entities on the LOCAL OBP
	@$(load_local_env) go run ./cmd/setup-entity

run-local: ## mirror the LOCAL chain into the LOCAL OBP (limit types via ARGS=)
	@$(load_local_env) go run ./cmd/cacher $(ARGS)

clean: ## remove build artifacts
	rm -rf bin

.PHONY: help build test vet setup-entity run setup-entity-local run-local clean
