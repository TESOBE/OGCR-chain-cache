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

run: ## mirror chain -> OBP (limit types via ARGS="parcel activity certification")
	go run ./cmd/cacher $(ARGS)

clean: ## remove build artifacts
	rm -rf bin

.PHONY: help build test vet setup-entity run clean
