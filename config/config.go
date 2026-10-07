package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds everything the cacher needs. It only ever reads the chain (no
// private key) and writes to OBP as a Platform App, with an OAuth2 client
// credentials token from OBP-OIDC. The chain side is the OGCR
// token family; this tool mirrors every deployed contract in it into the
// matching `*_on_chain` OBP dynamic entity.
//
// The three NFT addresses are required. The two carbon-credit addresses are
// optional so the cacher still runs against a chain where the credit contracts
// have not been deployed yet; the credit mirrors are skipped when they are
// unset.
type Config struct {
	OBPURL string

	// OIDCIssuer is the OBP-OIDC issuer URL; its token endpoint is discovered
	// from <issuer>/.well-known/openid-configuration. OIDCClientID and
	// OIDCClientSecret are the app's OIDC client; OBP makes it a Consumer on
	// its first call.
	OIDCIssuer       string
	OIDCClientID     string
	OIDCClientSecret string

	// EntitySpaceID is the bank (aka Space) the *_on_chain entities live in:
	// OBP_ENTITY_SPACE_ID, the same variable OGCR-DynamicEntities and OGCR-App
	// use, so all three agree on where the entities are. Unset means "ogcr";
	// set it to the empty string for system level entities.
	EntitySpaceID string

	RPCURL                       string
	ParcelContractAddress        string
	ActivityContractAddress      string
	CertificationContractAddress string

	// Optional: the carbon-credit half of the token family.
	CreditBatchContractAddress string
	CreditContractAddress      string

	// FromBlock is the block to start scanning *Minted events from (default 0).
	FromBlock uint64

	// IntervalSeconds is how often a supervising loop intends to re-run this
	// tool. It is recorded in chain_sync_status so a consumer can decide what
	// counts as stale without hardcoding the schedule. 0 means "run by hand".
	IntervalSeconds int

	// AutoSetupEntities makes the cacher create missing entity definitions and
	// update changed ones before it mirrors, like a migration
	// (AUTO_SETUP_ENTITIES, default true). Turn it off where an administrator
	// manages the definitions.
	AutoSetupEntities bool
}

// HasCreditBatch reports whether the CarbonCreditBatchNFT address is configured.
func (c *Config) HasCreditBatch() bool { return c.CreditBatchContractAddress != "" }

// HasCredit reports whether the CarbonCredit ERC-20 address is configured.
func (c *Config) HasCredit() bool { return c.CreditContractAddress != "" }

func Load() (*Config, error) {
	_ = godotenv.Load() // already-set env vars take precedence

	required := []string{
		"OBP_URL",
		"OIDC_ISSUER",
		"OIDC_CLIENT_ID",
		"OIDC_CLIENT_SECRET",
		"RPC_URL",
		"PARCEL_CONTRACT_ADDRESS",
		"ACTIVITY_CONTRACT_ADDRESS",
		"CERTIFICATION_CONTRACT_ADDRESS",
	}
	var missing []string
	for _, k := range required {
		if os.Getenv(k) == "" {
			missing = append(missing, k)
		}
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missing)
	}

	fromBlock := uint64(0)
	if s := os.Getenv("FROM_BLOCK"); s != "" {
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid FROM_BLOCK: %w", err)
		}
		fromBlock = n
	}

	interval := 0
	if s := os.Getenv("SYNC_INTERVAL_SECONDS"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil || n < 0 {
			return nil, fmt.Errorf("invalid SYNC_INTERVAL_SECONDS: %q", s)
		}
		interval = n
	}

	autoSetup := true
	if s := os.Getenv("AUTO_SETUP_ENTITIES"); s != "" {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, fmt.Errorf("invalid AUTO_SETUP_ENTITIES: %q", s)
		}
		autoSetup = b
	}

	space, ok := os.LookupEnv("OBP_ENTITY_SPACE_ID")
	if !ok {
		space = "ogcr"
	}

	return &Config{
		OBPURL:                       os.Getenv("OBP_URL"),
		OIDCIssuer:                   os.Getenv("OIDC_ISSUER"),
		OIDCClientID:                 os.Getenv("OIDC_CLIENT_ID"),
		OIDCClientSecret:             os.Getenv("OIDC_CLIENT_SECRET"),
		EntitySpaceID:                strings.TrimSpace(space),
		RPCURL:                       os.Getenv("RPC_URL"),
		ParcelContractAddress:        os.Getenv("PARCEL_CONTRACT_ADDRESS"),
		ActivityContractAddress:      os.Getenv("ACTIVITY_CONTRACT_ADDRESS"),
		CertificationContractAddress: os.Getenv("CERTIFICATION_CONTRACT_ADDRESS"),
		CreditBatchContractAddress:   os.Getenv("CREDIT_BATCH_CONTRACT_ADDRESS"),
		CreditContractAddress:        os.Getenv("CREDIT_CONTRACT_ADDRESS"),
		FromBlock:                    fromBlock,
		IntervalSeconds:              interval,
		AutoSetupEntities:            autoSetup,
	}, nil
}
