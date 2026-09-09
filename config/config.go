package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds everything the cacher needs. It only ever reads the chain (no
// private key) and writes to OBP via DirectLogin. The chain side is the OGCR
// token family; this tool mirrors every deployed contract in it into the
// matching `*_on_chain` OBP dynamic entity.
//
// The three NFT addresses are required. The two carbon-credit addresses are
// optional so the cacher still runs against a chain where the credit contracts
// have not been deployed yet; the credit mirrors are skipped when they are
// unset.
type Config struct {
	OBPURL         string
	OBPUsername    string
	OBPPassword    string
	OBPConsumerKey string

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
}

// HasCreditBatch reports whether the CarbonCreditBatchNFT address is configured.
func (c *Config) HasCreditBatch() bool { return c.CreditBatchContractAddress != "" }

// HasCredit reports whether the CarbonCredit ERC-20 address is configured.
func (c *Config) HasCredit() bool { return c.CreditContractAddress != "" }

func Load() (*Config, error) {
	_ = godotenv.Load() // already-set env vars take precedence

	required := []string{
		"OBP_URL",
		"OBP_USERNAME",
		"OBP_PASSWORD",
		"OBP_CONSUMER_KEY",
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

	return &Config{
		OBPURL:                       os.Getenv("OBP_URL"),
		OBPUsername:                  os.Getenv("OBP_USERNAME"),
		OBPPassword:                  os.Getenv("OBP_PASSWORD"),
		OBPConsumerKey:               os.Getenv("OBP_CONSUMER_KEY"),
		RPCURL:                       os.Getenv("RPC_URL"),
		ParcelContractAddress:        os.Getenv("PARCEL_CONTRACT_ADDRESS"),
		ActivityContractAddress:      os.Getenv("ACTIVITY_CONTRACT_ADDRESS"),
		CertificationContractAddress: os.Getenv("CERTIFICATION_CONTRACT_ADDRESS"),
		CreditBatchContractAddress:   os.Getenv("CREDIT_BATCH_CONTRACT_ADDRESS"),
		CreditContractAddress:        os.Getenv("CREDIT_CONTRACT_ADDRESS"),
		FromBlock:                    fromBlock,
		IntervalSeconds:              interval,
	}, nil
}
