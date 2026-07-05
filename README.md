# OGCR-chain-cache

Reads data points off the **OGCR chain** (private EVM, chain id `2025`) and
caches them into **OBP dynamic entities**. The reverse direction of the
[`OGCR-Chain/tokenizer`](https://github.com/OpenBankProject/OGCR-Chain) Go
service, which reads OBP and *mints* NFTs.

```
chain 2025   ──read──▶  OGCR-chain-cache  ──write──▶  OBP dynamic entities
ParcelNFT                                              parcel_on_chain
ActivityNFT                                            activity_on_chain
CertificationNFT                                       certification_on_chain
```

Written in Go, mirroring the tokenizer's stack. Kept as a separate repo (the
tokenizer mints; this mirrors — opposite jobs), and it reuses the tokenizer's
generated contract bindings (`internal/contract`).

## What it does

The tokenizer mints a family of tokens from OBP data: **ParcelNFT**,
**ActivityNFT**, **CertificationNFT** (plus CarbonCreditBatch / CarbonCredit,
which are out of scope here). This tool reads the three NFT types back and
upserts one `*_on_chain` record per token into OBP — a 1:1 mirror, no
aggregation. The source `parcel` / `activity` / `certificate_of_compliance`
entities are **not** mutated.

Each record holds the token's on-chain fields plus provenance
(`token_id`, `owner_address`, `token_uri`, `contract_address`, `chain_id`,
`tx_hash`, `block_number`):

| Entity | Business key | On-chain payload mirrored |
|---|---|---|
| `parcel_on_chain` | `parcel_id` | `parcel_uri`, `parcel_hash` |
| `activity_on_chain` | `activity_id` | `operator_id`, `name`, `activity_type`, `start_date`, `end_date`, `activity_url`, `activity_hash` |
| `certification_on_chain` | `certification_of_compliance_id` | `activity_nft_id`, `certification_scheme_id`, `certification_body_id`, `issue_date`, `expiry_date`, `certification_status`, `activity_url/hash`, `certification_url/hash` |

It only ever reads the chain (no private key) and writes to OBP via DirectLogin.

## Setup

```bash
cp .env.example .env        # then fill in OBP credentials + the 3 contract addresses
go build ./...              # or: make build
```

`.env` needs `PARCEL_CONTRACT_ADDRESS`, `ACTIVITY_CONTRACT_ADDRESS` and
`CERTIFICATION_CONTRACT_ADDRESS` (the tokenizer deployment addresses).

### One-time: create/update the `*_on_chain` entities

Needs `CanCreateSystemLevelDynamicEntity` / `CanUpdateSystemLevelDynamicEntity`
on the OBP user. Idempotent — an entity that already exists is updated in place
(PUT), so the previously-deployed `parcel_on_chain` (old CarbonProjectNFT shape)
is migrated to the new schema.

```bash
make setup-entity        # go run ./cmd/setup-entity
```

The OBP user also needs the per-entity `CanGetDynamicEntity_System*`,
`CanCreateDynamicEntity_System*` and `CanUpdateDynamicEntity_System*` roles for
each `*_on_chain` entity before the cacher can read/write records.

## Run

```bash
# Mirror all three token types (scans *Minted events)
make run                 # or: go run ./cmd/cacher

# Or limit to one or more types
go run ./cmd/cacher parcel
go run ./cmd/cacher activity certification
```

Re-running is safe: existing records are updated in place (matched by business
key), new ones are created.

## Layout

```
entities/*.json               OBP dynamic-entity definitions (for setup)
config/                       env/.env loading (3 contract addresses)
internal/contract/            ParcelNFT / ActivityNFT / CertificationNFT bindings
internal/obp/                 OBP DirectLogin auth + dynamic-entity read/write + management
internal/eth/                 chain reader (3 NFTs → OnChain* structs)
internal/cache/               upsert chain data into the *_on_chain entities
cmd/cacher/                   entry point: chain → *_on_chain
cmd/setup-entity/             one-time entity create/update
```
