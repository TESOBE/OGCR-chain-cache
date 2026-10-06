# OGCR-chain-cache

Reads data points off the **OGCR chain** (private EVM, chain id `2025`) and
caches them into **OBP dynamic entities**. The reverse direction of the
[`OGCR-Chain/tokenizer`](https://github.com/OpenBankProject/OGCR-Chain) Go
service, which reads OBP and *mints* NFTs.

```
chain 2025            ──read──▶  OGCR-chain-cache  ──write──▶  OBP dynamic entities
ParcelNFT                                                       parcel_on_chain
ActivityNFT                                                     activity_on_chain
CertificationNFT                                                certification_on_chain
CarbonCreditBatchNFT                                            carbon_credit_batch_on_chain
CarbonCredit (ERC-20)                                           carbon_credit_balance_on_chain
```

Written in Go, mirroring the tokenizer's stack. Kept as a separate repo (the
tokenizer mints; this mirrors — opposite jobs), and it reuses the tokenizer's
generated contract bindings (`internal/contract`).

## What it does

The goal is that **every on-chain component of OGCR has an OBP-side mirror**, so
that a consumer of the registry (the marketplace, a report, a dashboard) can read
chain state through the OBP API without speaking to a node.

This tool reads each token back and upserts one `*_on_chain` record into OBP — a
1:1 mirror, no aggregation. The source `parcel` / `activity` /
`certificate_of_compliance` entities are **not** mutated.

Most records hold the token's on-chain fields plus mint provenance (`token_id`,
`owner_address`, `token_uri`, `contract_address`, `chain_id`, `tx_hash`,
`block_number`):

| Entity | Business key | On-chain payload mirrored |
|---|---|---|
| `parcel_on_chain` | `parcel_id` | `parcel_uri`, `parcel_hash` |
| `activity_on_chain` | `activity_id` (a `reference:activity`) | `operator_id`, `name`, `activity_type`, `start_date`, `end_date`, `activity_url`, `activity_hash`, and `minted_at` (Unix time of the mint block) |
| `certification_on_chain` | `certification_of_compliance_id` | `activity_nft_id`, `certification_scheme_id`, `certification_body_id`, `issue_date`, `expiry_date`, `certification_status`, `activity_url/hash`, `certification_url/hash` |
| `carbon_credit_batch_on_chain` | `batch_key` | `token_id`, `credit_type`, `activity_nft_id`, `token_bound_account`, `credit_balance`, url/hash pairs for activity, operator, certification, scheme and body |
| `carbon_credit_balance_on_chain` | `owner_address` | `balance`, `decimals`, `symbol`, `holder_type`, `batch_token_id` |
| `chain_sync_status` | `sync_key` | not a mirror: the liveness record for this tool, see below |

It only ever reads the chain (no private key) and writes to OBP as a Platform App
(see [Authentication](#authentication-a-platform-app)).

This tool is the source of truth for these entities: their definitions live here, in
`entities/`, and OGCR-DynamicEntities ignores them in its spreadsheet. In
`activity_on_chain`, `activity_id` is a `reference:activity` and indexed, so OBP can join
an activity to its NFT (the registry query does, picking the latest by `minted_at`). The
cost: OBP refuses an `activity_on_chain` record whose activity isn't in the same space, so
such a token is logged as a failed upsert and skipped.

### Liveness: `chain_sync_status`

Every run ends by writing one record per chain to `chain_sync_status`, holding
the head block it reached, the time it finished, how many records each mirror
wrote, and how many failed.

It exists because the mirrored tokens cannot answer "is this connection alive".
They record when a token was *minted*, so on a chain where nothing has been
minted for a week, a mirror that ran a second ago and one that died a week ago
produce exactly the same newest record. Freshness has to be recorded separately
or it cannot be known.

Two details worth knowing:

- **A run that cannot reach the chain writes nothing.** The previous record then
  ages visibly, which is the correct signal. Writing a status that claims a
  successful look at an unreachable chain would be worse than writing none.
- **`interval_seconds` comes from `SYNC_INTERVAL_SECONDS`**, set by whatever
  supervises the loop. It is recorded so a consumer can decide what counts as
  stale without hardcoding the schedule. Unset means the tool was run by hand.

`run_status` is `ok` when every configured mirror completed cleanly and
`partial` when the chain was readable but some records failed to write.

### The carbon-credit layer

A **CarbonCreditBatchNFT** is one `(activity_nft_id, credit_type)` pair, and the
contract permits exactly one batch per pair. Each batch owns an ERC-6551
token-bound account, and that account is where the batch's **CarbonCredit**
(ERC-20) balance sits. Owning the batch NFT is what gates moving those credits
out into an ordinary wallet.

That gives two questions a consumer wants answered, and one entity each:

- **What batches exist, and what is still in them?** `carbon_credit_batch_on_chain`.
  Every batch, its credit type, its full provenance chain, and the balance held
  in its token-bound account.
- **Who holds credits right now?** `carbon_credit_balance_on_chain`. One record
  per holding address, with `holder_type` distinguishing credits still inside a
  batch (`token_bound_account`) from credits withdrawn into a pooled wallet
  balance (`wallet`). Match `owner_address` against `operator.ogcr_wallet_address`
  or `land_manager.ogcr_wallet_address` to find a party's holdings.

Two caveats worth knowing:

- **Balances are current state, not history.** The NFT mirrors are built from
  `*Minted` events, so they are faithful snapshots of a mint. An ERC-20 balance
  has no such event to anchor to, so every balance carries the `block_number` it
  was read at and is only true as of that height. Re-run the cacher to refresh.
- **The holder set comes from the `Transfer` log.** ERC-20 keeps no list of
  holders on-chain, so the only way to know whom to ask `balanceOf` about is to
  replay transfers. Zero balances are written rather than skipped, so an address
  that has moved everything out has its cached record corrected to zero.

`batch_key` is `"<activity_nft_id>:<credit_type>"`. It exists because records are
matched for upsert with a server-side query filter, which is dependable for a
string field and not for an integer one; it is exactly as unique as `token_id`.

## Setup

```bash
cp .env.example .env        # then fill in the OIDC client + the 3 contract addresses
go build ./...              # or: make build
```

`.env` needs `PARCEL_CONTRACT_ADDRESS`, `ACTIVITY_CONTRACT_ADDRESS` and
`CERTIFICATION_CONTRACT_ADDRESS` (the tokenizer deployment addresses).

`CREDIT_BATCH_CONTRACT_ADDRESS` and `CREDIT_CONTRACT_ADDRESS` are optional: leave
either unset and the cacher skips that mirror and says so, rather than failing.
That keeps it runnable against a chain where the credit contracts are not
deployed yet.

### Authentication: a Platform App

All three tools (`cacher`, `setup-entity`, `delete-records`) call OBP as their own
application, not as a User, the way OBP-Sentinel does. There is no DirectLogin.
They get a token from OBP-OIDC with the OAuth2 client credentials grant (the
token endpoint is discovered from `OIDC_ISSUER/.well-known/openid-configuration`)
and send it as `Authorization: Bearer`. A token is replaced 30 seconds before it
expires; a 401 gets one retry with a new token.

The Roles they need are **Scopes granted to the app's Consumer**, not
Entitlements of a User. Setting it up:

1. Create an OIDC client for this app in OBP-OIDC and put its id and secret in
   `OIDC_CLIENT_ID` / `OIDC_CLIENT_SECRET`.
2. Run any of the tools once. OBP creates a Consumer for the client on its first
   call. The tool's attempt to declare its Scopes is refused, and the log names
   the Consumer's `CONSUMER_ID`.
3. An administrator marks that Consumer as a Platform App:
   `POST /obp/v7.0.0/management/platform-apps`.
4. Run `make setup-entity`. Each tool declares the Scopes it needs
   (`PUT /obp/v7.0.0/consumers/current/platform-app`) and logs any it does not
   hold yet.
5. The administrator sees them in `GET /obp/v7.0.0/management/platform-apps`
   and grants them with `POST /obp/v7.0.0/consumers/CONSUMER_ID/scopes`.

The Scopes, all at the `OBP_ENTITY_SPACE_ID` bank id (`SYS` for system level):

| Scope | Used by |
|---|---|
| `CanGetDynamicEntityDefinitions` | `setup-entity` |
| `CanCreateDynamicEntityDefinition` | `setup-entity` |
| `CanUpdateDynamicEntityDefinition` | `setup-entity` |
| `CanGetDynamicEntityRecord_<entity>` | `cacher`, `delete-records` |
| `CanCreateDynamicEntityRecord_<entity>` | `cacher` |
| `CanUpdateDynamicEntityRecord_<entity>` | `cacher` |
| `CanDeleteDynamicEntityRecord_<entity>` | `delete-records` (declared optional) |

where `<entity>` is each `*_on_chain` entity and `chain_sync_status`.

OBP refuses a whole declaration that names a Role that does not exist, and a
record Role only exists once its entity does. So the tools declare record Roles
only for the entities that already exist, and the definition Roles alone until
`CanGetDynamicEntityDefinitions` is granted. On a fresh install that takes two
rounds: grant the definition Scopes, run `make setup-entity` again (it creates
the entities, then declares their record Roles), and grant those.

Every definition in `entities/` sets `"auth_mode": "UserOrApplication"`. OBP's
default, `UserOnly`, would refuse the app's token on the record endpoints; with
`UserOrApplication` a User holding the matching Entitlements still has access.

### Running against a local OBP

`.env.example` ships with the shared DCR and the shared chain. For local
development, point `OBP_URL` at your local OBP and `RPC_URL` at a local node,
with an OIDC client on the local OBP-OIDC:

```bash
OBP_URL=http://localhost:8080
OIDC_ISSUER=http://localhost:9000/obp-oidc
OIDC_CLIENT_ID=...          # an OIDC client on the LOCAL OBP-OIDC
OIDC_CLIENT_SECRET=...
RPC_URL=http://127.0.0.1:8545
```

The entities live in the bank (aka Space) named by `OBP_ENTITY_SPACE_ID`,
default `ogcr`: the same variable OGCR-DynamicEntities and OGCR-App read, so set
it to the same value in all three. Set it to the empty string for system level.

The app's Consumer needs the Scopes listed under
[Authentication](#authentication-a-platform-app) at that bank id (`SYS` for
system level).

**Keep the two halves consistent.** A local chain paired with a shared OBP would
mirror throwaway fixture tokens into the real registry. If your `.env` has to
stay pointed at shared infrastructure, put local values in `.env.local`
(gitignored) instead of editing `.env`, and use the local targets:

```bash
make setup-entity-local   # create/update the entities on the local OBP
make run-local            # mirror the local chain into the local OBP
```

Those read `../OGCR-Smart-Contracts/deployments/local-anvil.env` for the contract
addresses, so they cannot go stale against what is actually deployed, then apply
`.env.local` on top. Override the path with `LOCAL_CHAIN_ENV=...` if your
checkout is laid out differently. With a fully local `.env` they behave the same
as `make setup-entity` and `make run`.

Note that OBP refuses a structural change to a dynamic entity that already holds
data. If `setup-entity` reports `OBP-09023` for an entity, its records must be
deleted before its schema can be migrated.

### One-time: create/update the `*_on_chain` entities

Uses the v7.0.0 management API in the `OBP_ENTITY_SPACE_ID` space, and needs the
definition Scopes at that bank id on the app's Consumer. Idempotent — an entity that already exists is updated in place
(PUT), so the previously-deployed `parcel_on_chain` (old CarbonProjectNFT shape)
is migrated to the new schema.

Every definition in `entities/` sets `"has_public_access": true`, so the cached
chain data and `chain_sync_status` can be read without a login, and
`"auth_mode": "UserOrApplication"`, so the Platform App can write them.
`setup-entity` moves both out of the schema to the top level of the v7.0.0
request body.

```bash
make setup-entity        # go run ./cmd/setup-entity
```

When it is done, `setup-entity` declares the record Scopes of the entities it has
created; the cacher can read and write records once they are granted.

### Emptying an entity

OBP only allows some definition changes, such as retyping a field, on an entity
with no records. The records are a cache of the chain, so empty the entity, apply
the definition, and let the cacher write them again:

```bash
go run ./cmd/delete-records activity_on_chain         # list the records
go run ./cmd/delete-records -yes activity_on_chain    # delete them (definition and Roles stay)
make setup-entity
go run ./cmd/cacher activity
```

It only accepts entities defined in `entities/`. Deleting needs
`CanDeleteDynamicEntityRecord_<entity>` at the space's bank id. To empty every
`*_on_chain` entity at once, use `scripts/delete-records.sh` (see Scripts).

## Run

```bash
# Mirror everything configured (scans *Minted events)
make run                 # or: go run ./cmd/cacher

# Or limit to one or more types: parcel, activity, certification, credit
go run ./cmd/cacher parcel
go run ./cmd/cacher activity certification
go run ./cmd/cacher credit
```

`credit` covers both credit entities: batches are scanned first and reused to
classify which balance holders are token-bound accounts, so the batch log is only
walked once. Asking for `credit` explicitly with no credit address configured is
an error, so a forgotten address is not mistaken for "nothing to mirror".

Re-running is safe: existing records are updated in place (matched by business
key), new ones are created.

## Scripts

`scripts/` wraps the commands above for the common cases. Each runs from any
directory, reads `.env`, and works in the `OBP_ENTITY_SPACE_ID` space.

| Script | Runs | Does |
|---|---|---|
| `scripts/delete-records.sh [ENTITY...]` | `cmd/delete-records` | Lists the records of every `*_on_chain` entity in `entities/` (or just the ones named), asks once, then deletes them. Definitions and Roles stay. |
| `scripts/setup-entities.sh` | `cmd/setup-entity` | Creates or updates the entities from `entities/` (same as `make setup-entity`). |
| `scripts/sync-from-ogcr-chain.sh [TYPE...]` | `cmd/cacher` | Mirrors the OGCR chain set by `RPC_URL` in `.env` into OBP once (same as `make run`). |
| `scripts/sync-from-local-anvil-chain.sh [TYPE...]` | `cmd/cacher` | Mirrors the local anvil chain into OBP once, loading `local-anvil.env` and then `.env.local` over `.env` (same as `make run-local`). |

`TYPE` is any of `parcel`, `activity`, `certification`, `credit`. A full reset
of the cache, e.g. after changing a definition in `entities/`:

```bash
scripts/delete-records.sh
scripts/setup-entities.sh
scripts/sync-from-ogcr-chain.sh     # or scripts/sync-from-local-anvil-chain.sh
```

## Layout

```
entities/*.json               OBP dynamic-entity definitions (for setup)
config/                       env/.env loading (contract addresses)
internal/contract/            abigen bindings for the five contracts
internal/obp/                 Platform App auth + dynamic-entity read/write + management
internal/eth/                 chain reader (tokens + balances → OnChain* structs)
internal/cache/               upsert chain data into the *_on_chain entities
cmd/cacher/                   entry point: chain → *_on_chain
cmd/setup-entity/             one-time entity create/update
cmd/delete-records/           empty one *_on_chain entity (the cacher refills it)
scripts/                      shell wrappers: delete, set up, sync (see Scripts)
```

Bindings are generated from the contract ABIs in
[OGCR-Smart-Contracts](https://github.com/OpenBankProject/OGCR-Smart-Contracts):

```bash
forge build --out /tmp/ogcr-out                       # in OGCR-Smart-Contracts
jq -c '.abi' /tmp/ogcr-out/CarbonCredit.sol/CarbonCredit.json > /tmp/CarbonCredit.json
abigen --abi /tmp/CarbonCredit.json --pkg contract --type CarbonCredit \
  --out internal/contract/carbon_credit.go
```

## Testing against a chain

`go test ./...` is offline by default. The integration tests are skipped unless
the environment names a chain to talk to.

Any EVM node works, so a local [anvil](https://getfoundry.sh) stands in for the
Besu network. `OGCR-Smart-Contracts` carries a fixture that deploys all five
contracts and mints a coherent set of records, parcel through activity and
certificate to credit batch:

```bash
anvil --chain-id 2025

# in OGCR-Smart-Contracts
forge script script/local/DeployLocalStack.s.sol \
  --rpc-url http://127.0.0.1:8545 --broadcast --legacy

# it writes deployments/local-anvil.env with every address
set -a; . deployments/local-anvil.env; set +a

# in this repo
OGCR_TEST_RPC_URL=$RPC_URL \
OGCR_TEST_PARCEL_ADDRESS=$PARCEL_CONTRACT_ADDRESS \
OGCR_TEST_ACTIVITY_ADDRESS=$ACTIVITY_CONTRACT_ADDRESS \
OGCR_TEST_CERTIFICATION_ADDRESS=$CERTIFICATION_CONTRACT_ADDRESS \
OGCR_TEST_CREDIT_BATCH_ADDRESS=$CREDIT_BATCH_CONTRACT_ADDRESS \
OGCR_TEST_CREDIT_ADDRESS=$CREDIT_CONTRACT_ADDRESS \
go test ./internal/eth -run TestScan -v
```

`TestScanNFTs` checks that each NFT mirror has the business key it is upserted
on, and that a certificate's activity reference resolves to an activity that was
actually scanned. `TestScanCredits` checks that batch keys are unique, that
balances are read and carry a block height, and that every batch's token-bound
account is classified as one and agrees with the batch's own balance.

Because a fresh anvil assigns addresses deterministically, those addresses are
the same on every run, so the generated env file can be reused rather than
re-read each time.

## Funding

The OGCR Project has received funding from the European Union's Horizon Europe programme under grant agreement 101218854.
