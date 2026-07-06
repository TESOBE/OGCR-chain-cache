package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds everything the cacher needs. It only ever reads the chain (no
// private key) and writes to OBP via DirectLogin. The chain side is the OGCR
// token family; this tool mirrors the three NFT types (Parcel, Activity,
// Certification) into their `*_on_chain` OBP dynamic entities.
type Config struct {
	OBPURL         string
	OBPUsername    string
	OBPPassword    string
	OBPConsumerKey string

	RPCURL                       string
	ParcelContractAddress        string
	ActivityContractAddress      string
	CertificationContractAddress string

	// FromBlock is the block to start scanning *Minted events from (default 0).
	FromBlock uint64
}

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

	return &Config{
		OBPURL:                       os.Getenv("OBP_URL"),
		OBPUsername:                  os.Getenv("OBP_USERNAME"),
		OBPPassword:                  os.Getenv("OBP_PASSWORD"),
		OBPConsumerKey:               os.Getenv("OBP_CONSUMER_KEY"),
		RPCURL:                       os.Getenv("RPC_URL"),
		ParcelContractAddress:        os.Getenv("PARCEL_CONTRACT_ADDRESS"),
		ActivityContractAddress:      os.Getenv("ACTIVITY_CONTRACT_ADDRESS"),
		CertificationContractAddress: os.Getenv("CERTIFICATION_CONTRACT_ADDRESS"),
		FromBlock:                    fromBlock,
	}, nil
}
