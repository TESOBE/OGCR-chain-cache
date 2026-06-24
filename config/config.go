package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds everything the cacher needs. Unlike the tokenizer it never
// signs transactions, so there is no private key — it only reads chain and
// writes to OBP via DirectLogin.
type Config struct {
	OBPURL          string
	OBPUsername     string
	OBPPassword     string
	OBPConsumerKey  string
	RPCURL          string
	ContractAddress string
	FromBlock       uint64
}

func Load() (*Config, error) {
	_ = godotenv.Load() // already-set env vars take precedence

	required := []string{
		"OBP_URL",
		"OBP_USERNAME",
		"OBP_PASSWORD",
		"OBP_CONSUMER_KEY",
		"RPC_URL",
		"CONTRACT_ADDRESS",
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
		OBPURL:          os.Getenv("OBP_URL"),
		OBPUsername:     os.Getenv("OBP_USERNAME"),
		OBPPassword:     os.Getenv("OBP_PASSWORD"),
		OBPConsumerKey:  os.Getenv("OBP_CONSUMER_KEY"),
		RPCURL:          os.Getenv("RPC_URL"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		FromBlock:       fromBlock,
	}, nil
}
