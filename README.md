# OGCR-chain-cache

Reads data points off the **OGCR chain** (private EVM, chain id `2025`) and
caches them into **OBP dynamic entities**. The reverse direction of the
[`OGCR-Chain/tokenizer`](https://github.com/OpenBankProject/OGCR-Chain) Go
service, which reads OBP and *mints* NFTs.

```
chain 2025  ──read──▶  OGCR-chain-cache  ──write──▶  OBP dynamic entity
(CarbonProjectNFT)                                   (parcel_on_chain)
```

Written in Go, mirroring the tokenizer's stack. Kept as a separate repo (the
tokenizer mints; this mirrors — opposite jobs), but it reuses the tokenizer's
generated `CarbonProjectNFT` binding (`internal/contract`).

## What it does (first/simplest case)

The tokenizer mints one `CarbonProjectNFT` per verified `parcel_id`. This tool
reads those tokens back and upserts one `parcel_on_chain` record per parcel into
OBP — a 1:1 mirror, no aggregation. The source `parcel` entity is **not**
mutated (kept stable so it can later be hashed and anchored on-chain).

Each `parcel_on_chain` record holds: `parcel_id`, `token_id`, `carbon_amount`,
`project_type`, `methodology`, `vintage`, `minted_at`, `owner_address`,
`token_uri`, `contract_address`, `chain_id`, `tx_hash`, `block_number`.

It only ever reads the chain (no private key) and writes to OBP via DirectLogin.

## Setup

```bash
cp .env.example .env        # then fill in OBP credentials
go build ./...              # or: make build
```

Confirm `CONTRACT_ADDRESS` in `.env` points at the live CarbonProjectNFT — the
deployment doc and the tokenizer's `.env.example` currently disagree
(`0xeec918d7…` vs `0x30753E4A…`).

### One-time: create the `parcel_on_chain` entity

Needs the `CanCreateSystemLevelDynamicEntity` role on the OBP user.

```bash
make setup-entity        # go run ./cmd/setup-entity
```

## Run

```bash
# Mirror every minted parcel (scans CarbonProjectMinted events)
make run                 # or: go run ./cmd/cacher

# Or just specific parcel_id(s)
make run ARGS="parcel_AB12CD34EF parcel_99XYZ"
```

Re-running is safe: existing `parcel_on_chain` records are updated in place
(matched by `parcel_id`), new ones are created.

## Layout

```
abi/CarbonProjectNFT.json     ABI (boundary with the chain half)
entities/parcel_on_chain.json OBP dynamic-entity definition (for setup)
config/                       env/.env loading
internal/contract/            CarbonProjectNFT binding (from OGCR-Chain)
internal/obp/                 OBP DirectLogin auth + dynamic-entity read/write
internal/eth/                 chain reader (CarbonProjectNFT → OnChainParcel)
internal/cache/               upsert chain data into parcel_on_chain
cmd/cacher/                   entry point: chain → parcel_on_chain
cmd/setup-entity/             one-time entity creation
```
