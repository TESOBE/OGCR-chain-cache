// Package eth reads the OGCR token family off the OGCR chain (chain id 2025):
// ParcelNFT, ActivityNFT, CertificationNFT, CarbonCreditBatchNFT and the
// CarbonCredit ERC-20. It reads those tokens back so they can be mirrored into
// the OBP `*_on_chain` dynamic entities.
//
// The three NFT types and the credit batch are historical: each is discovered
// from its *Minted event, so a record is a faithful snapshot of a mint. Credit
// balances are the exception. ERC-20 balances are current state, so every
// balance this package returns carries the block height it was read at.
package eth

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/TESOBE/OGCR-chain-cache/internal/contract"
)

// scanChunk is the largest block window this RPC allows per eth_getLogs call
// (it rejects "Requested range exceeds maximum RPC range limit" above ~5000).
const scanChunk = uint64(5000)

// OnChainParcel is the chain-side view of one ParcelNFT, shaped for the OBP
// `parcel_on_chain` entity. JSON tags are the entity field names.
type OnChainParcel struct {
	ParcelID        string `json:"parcel_id"`
	TokenID         uint64 `json:"token_id"`
	ParcelURI       string `json:"parcel_uri"`
	ParcelHash      string `json:"parcel_hash"`
	OwnerAddress    string `json:"owner_address"`
	TokenURI        string `json:"token_uri"`
	ContractAddress string `json:"contract_address"`
	ChainID         uint64 `json:"chain_id"`
	TxHash          string `json:"tx_hash"`
	BlockNumber     uint64 `json:"block_number"`
}

// OnChainActivity is the chain-side view of one ActivityNFT, shaped for the OBP
// `activity_on_chain` entity.
type OnChainActivity struct {
	ActivityID      string `json:"activity_id"`
	TokenID         uint64 `json:"token_id"`
	OperatorID      string `json:"operator_id"`
	Name            string `json:"name"`
	ActivityType    string `json:"activity_type"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	ActivityURL     string `json:"activity_url"`
	ActivityHash    string `json:"activity_hash"`
	OwnerAddress    string `json:"owner_address"`
	TokenURI        string `json:"token_uri"`
	ContractAddress string `json:"contract_address"`
	ChainID         uint64 `json:"chain_id"`
	TxHash          string `json:"tx_hash"`
	BlockNumber     uint64 `json:"block_number"`
}

// OnChainCertification is the chain-side view of one CertificationNFT, shaped
// for the OBP `certification_on_chain` entity.
type OnChainCertification struct {
	CertificationOfComplianceID string `json:"certification_of_compliance_id"`
	TokenID                     uint64 `json:"token_id"`
	ActivityNftID               uint64 `json:"activity_nft_id"`
	CertificationSchemeID       string `json:"certification_scheme_id"`
	CertificationBodyID         string `json:"certification_body_id"`
	IssueDate                   string `json:"issue_date"`
	ExpiryDate                  string `json:"expiry_date"`
	CertificationStatus         string `json:"certification_status"`
	ActivityURL                 string `json:"activity_url"`
	ActivityHash                string `json:"activity_hash"`
	CertificationURL            string `json:"certification_url"`
	CertificationHash           string `json:"certification_hash"`
	OwnerAddress                string `json:"owner_address"`
	TokenURI                    string `json:"token_uri"`
	ContractAddress             string `json:"contract_address"`
	ChainID                     uint64 `json:"chain_id"`
	TxHash                      string `json:"tx_hash"`
	BlockNumber                 uint64 `json:"block_number"`
}

// OnChainCreditBatch is the chain-side view of one CarbonCreditBatchNFT, shaped
// for the OBP `carbon_credit_batch_on_chain` entity. A batch is one
// (activity_nft_id, credit_type) pair and owns an ERC-6551 token-bound account
// that holds its CarbonCredit balance.
type OnChainCreditBatch struct {
	// BatchKey is "<activity_nft_id>:<credit_type>", the contract's own
	// uniqueness constraint rendered as a string so it can serve as the
	// business key for the OBP upsert.
	BatchKey                string `json:"batch_key"`
	TokenID                 uint64 `json:"token_id"`
	CreditType              string `json:"credit_type"`
	ActivityNftID           uint64 `json:"activity_nft_id"`
	ActivityURL             string `json:"activity_url"`
	ActivityHash            string `json:"activity_hash"`
	OperatorURL             string `json:"operator_url"`
	OperatorHash            string `json:"operator_hash"`
	CertificationURL        string `json:"certification_url"`
	CertificationHash       string `json:"certification_hash"`
	CertificationSchemeURL  string `json:"certification_scheme_url"`
	CertificationSchemeHash string `json:"certification_scheme_hash"`
	CertificationBodyURL    string `json:"certification_body_url"`
	CertificationBodyHash   string `json:"certification_body_hash"`
	TokenBoundAccount       string `json:"token_bound_account"`

	// Balance fields are only populated when a CarbonCredit address is
	// configured. They are current state as of BlockNumberBalance, not part of
	// the mint snapshot the other fields describe.
	CreditBalance         string `json:"credit_balance"`
	CreditDecimals        uint8  `json:"credit_decimals"`
	CreditContractAddress string `json:"credit_contract_address"`
	BlockNumberBalance    uint64 `json:"block_number_balance"`

	OwnerAddress    string `json:"owner_address"`
	TokenURI        string `json:"token_uri"`
	ContractAddress string `json:"contract_address"`
	ChainID         uint64 `json:"chain_id"`
	TxHash          string `json:"tx_hash"`
	BlockNumber     uint64 `json:"block_number"`
}

// OnChainCreditBalance is one address's CarbonCredit (ERC-20) balance, shaped
// for the OBP `carbon_credit_balance_on_chain` entity.
type OnChainCreditBalance struct {
	OwnerAddress    string `json:"owner_address"`
	Balance         string `json:"balance"`
	Decimals        uint8  `json:"decimals"`
	Symbol          string `json:"symbol"`
	HolderType      string `json:"holder_type"`
	BatchTokenID    uint64 `json:"batch_token_id"`
	ContractAddress string `json:"contract_address"`
	ChainID         uint64 `json:"chain_id"`
	BlockNumber     uint64 `json:"block_number"`
}

// Holder types for OnChainCreditBalance.
const (
	HolderWallet            = "wallet"
	HolderTokenBoundAccount = "token_bound_account"
)

// Addresses are the deployed contracts to read. Parcel, Activity and
// Certification are required; CreditBatch and Credit are optional so the cacher
// still runs against a chain where the credit half is not deployed yet.
type Addresses struct {
	Parcel        string
	Activity      string
	Certification string
	CreditBatch   string
	Credit        string
}

// Reader binds the OGCR contracts on one RPC connection. creditBatch and credit
// are nil when their addresses were not configured.
type Reader struct {
	client          *ethclient.Client
	chainID         uint64
	parcel          *contract.ParcelNFT
	parcelAddr      common.Address
	activity        *contract.ActivityNFT
	activityAddr    common.Address
	cert            *contract.CertificationNFT
	certAddr        common.Address
	creditBatch     *contract.CarbonCreditBatchNFT
	creditBatchAddr common.Address
	credit          *contract.CarbonCredit
	creditAddr      common.Address
}

func NewReader(rpcURL string, addrs Addresses) (*Reader, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC: %w", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	pAddr := common.HexToAddress(addrs.Parcel)
	parcel, err := contract.NewParcelNFT(pAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind ParcelNFT: %w", err)
	}
	aAddr := common.HexToAddress(addrs.Activity)
	activity, err := contract.NewActivityNFT(aAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind ActivityNFT: %w", err)
	}
	cAddr := common.HexToAddress(addrs.Certification)
	cert, err := contract.NewCertificationNFT(cAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind CertificationNFT: %w", err)
	}

	r := &Reader{
		client:       client,
		chainID:      chainID.Uint64(),
		parcel:       parcel,
		parcelAddr:   pAddr,
		activity:     activity,
		activityAddr: aAddr,
		cert:         cert,
		certAddr:     cAddr,
	}

	if addrs.CreditBatch != "" {
		bAddr := common.HexToAddress(addrs.CreditBatch)
		creditBatch, err := contract.NewCarbonCreditBatchNFT(bAddr, client)
		if err != nil {
			return nil, fmt.Errorf("bind CarbonCreditBatchNFT: %w", err)
		}
		r.creditBatch, r.creditBatchAddr = creditBatch, bAddr
	}
	if addrs.Credit != "" {
		ccAddr := common.HexToAddress(addrs.Credit)
		credit, err := contract.NewCarbonCredit(ccAddr, client)
		if err != nil {
			return nil, fmt.Errorf("bind CarbonCredit: %w", err)
		}
		r.credit, r.creditAddr = credit, ccAddr
	}

	return r, nil
}

func (r *Reader) ChainID() uint64 { return r.chainID }

// HasCreditBatch reports whether a CarbonCreditBatchNFT address was configured.
func (r *Reader) HasCreditBatch() bool { return r.creditBatch != nil }

// HasCredit reports whether a CarbonCredit ERC-20 address was configured.
func (r *Reader) HasCredit() bool { return r.credit != nil }

// retry runs fn up to 4 times with a short backoff, to ride out a flaky RPC
// connection. Returns the last error if all attempts fail.
func retry(fn func() error) error {
	var err error
	for i := 0; i < 4; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(i+1) * 500 * time.Millisecond)
	}
	return err
}

// HeadBlock returns the current chain head. Exported so a caller can record the
// height a run reached, which is what tells a reader the chain itself is moving
// rather than only the mirror still running.
func (r *Reader) HeadBlock(ctx context.Context) (uint64, error) { return r.head(ctx) }

// head returns the current chain head, retried for a flaky RPC.
func (r *Reader) head(ctx context.Context) (uint64, error) {
	var head uint64
	if err := retry(func() error {
		var e error
		head, e = r.client.BlockNumber(ctx)
		return e
	}); err != nil {
		return 0, fmt.Errorf("get head block: %w", err)
	}
	return head, nil
}

// eachWindow calls fn for each [start,end] block window from fromBlock to the
// current head, in scanChunk-sized steps (the RPC caps eth_getLogs range).
func (r *Reader) eachWindow(ctx context.Context, fromBlock uint64, fn func(start, end uint64) error) error {
	head, err := r.head(ctx)
	if err != nil {
		return err
	}
	for start := fromBlock; start <= head; {
		end := start + scanChunk
		if end > head {
			end = head
		}
		if err := fn(start, end); err != nil {
			return err
		}
		start = end + 1
	}
	return nil
}

// ScanParcels yields every minted ParcelNFT via the ParcelMinted event log,
// enriched with getParcel / ownerOf / tokenURI. The scan is chunked across the
// full chain and retried per window for flaky connections.
func (r *Reader) ScanParcels(ctx context.Context, fromBlock uint64) ([]*OnChainParcel, error) {
	opts := &bind.CallOpts{Context: ctx}
	var out []*OnChainParcel
	err := r.eachWindow(ctx, fromBlock, func(start, end uint64) error {
		var it *contract.ParcelNFTParcelMintedIterator
		if err := retry(func() error {
			var e error
			it, e = r.parcel.FilterParcelMinted(&bind.FilterOpts{Start: start, End: &end, Context: ctx}, nil, nil)
			return e
		}); err != nil {
			return fmt.Errorf("filter ParcelMinted [%d,%d]: %w", start, end, err)
		}
		defer it.Close()
		for it.Next() {
			ev := it.Event
			var data contract.ParcelNFTParcelData
			var owner common.Address
			var uri string
			if err := retry(func() error {
				var e error
				if data, e = r.parcel.GetParcel(opts, ev.TokenId); e != nil {
					return e
				}
				if owner, e = r.parcel.OwnerOf(opts, ev.TokenId); e != nil {
					return e
				}
				uri, e = r.parcel.TokenURI(opts, ev.TokenId)
				return e
			}); err != nil {
				return fmt.Errorf("enrich parcel token %s: %w", ev.TokenId, err)
			}
			out = append(out, &OnChainParcel{
				ParcelID:        data.ParcelId,
				TokenID:         ev.TokenId.Uint64(),
				ParcelURI:       data.ParcelUri,
				ParcelHash:      data.ParcelHash,
				OwnerAddress:    owner.Hex(),
				TokenURI:        uri,
				ContractAddress: r.parcelAddr.Hex(),
				ChainID:         r.chainID,
				TxHash:          ev.Raw.TxHash.Hex(),
				BlockNumber:     ev.Raw.BlockNumber,
			})
		}
		return it.Error()
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanActivities yields every minted ActivityNFT via the ActivityMinted event log.
func (r *Reader) ScanActivities(ctx context.Context, fromBlock uint64) ([]*OnChainActivity, error) {
	opts := &bind.CallOpts{Context: ctx}
	var out []*OnChainActivity
	err := r.eachWindow(ctx, fromBlock, func(start, end uint64) error {
		var it *contract.ActivityNFTActivityMintedIterator
		if err := retry(func() error {
			var e error
			it, e = r.activity.FilterActivityMinted(&bind.FilterOpts{Start: start, End: &end, Context: ctx}, nil, nil)
			return e
		}); err != nil {
			return fmt.Errorf("filter ActivityMinted [%d,%d]: %w", start, end, err)
		}
		defer it.Close()
		for it.Next() {
			ev := it.Event
			var data contract.ActivityNFTActivityData
			var owner common.Address
			var uri string
			if err := retry(func() error {
				var e error
				if data, e = r.activity.GetActivity(opts, ev.TokenId); e != nil {
					return e
				}
				if owner, e = r.activity.OwnerOf(opts, ev.TokenId); e != nil {
					return e
				}
				uri, e = r.activity.TokenURI(opts, ev.TokenId)
				return e
			}); err != nil {
				return fmt.Errorf("enrich activity token %s: %w", ev.TokenId, err)
			}
			out = append(out, &OnChainActivity{
				ActivityID:      data.ActivityId,
				TokenID:         ev.TokenId.Uint64(),
				OperatorID:      data.OperatorId,
				Name:            data.Name,
				ActivityType:    data.ActivityType,
				StartDate:       data.StartDate,
				EndDate:         data.EndDate,
				ActivityURL:     data.ActivityUrl,
				ActivityHash:    data.ActivityHash,
				OwnerAddress:    owner.Hex(),
				TokenURI:        uri,
				ContractAddress: r.activityAddr.Hex(),
				ChainID:         r.chainID,
				TxHash:          ev.Raw.TxHash.Hex(),
				BlockNumber:     ev.Raw.BlockNumber,
			})
		}
		return it.Error()
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanCertifications yields every minted CertificationNFT via the
// CertificationMinted event log.
func (r *Reader) ScanCertifications(ctx context.Context, fromBlock uint64) ([]*OnChainCertification, error) {
	opts := &bind.CallOpts{Context: ctx}
	var out []*OnChainCertification
	err := r.eachWindow(ctx, fromBlock, func(start, end uint64) error {
		var it *contract.CertificationNFTCertificationMintedIterator
		if err := retry(func() error {
			var e error
			it, e = r.cert.FilterCertificationMinted(&bind.FilterOpts{Start: start, End: &end, Context: ctx}, nil, nil, nil)
			return e
		}); err != nil {
			return fmt.Errorf("filter CertificationMinted [%d,%d]: %w", start, end, err)
		}
		defer it.Close()
		for it.Next() {
			ev := it.Event
			var data contract.CertificationNFTCertificationData
			var owner common.Address
			var uri string
			if err := retry(func() error {
				var e error
				if data, e = r.cert.GetCertification(opts, ev.TokenId); e != nil {
					return e
				}
				if owner, e = r.cert.OwnerOf(opts, ev.TokenId); e != nil {
					return e
				}
				uri, e = r.cert.TokenURI(opts, ev.TokenId)
				return e
			}); err != nil {
				return fmt.Errorf("enrich certification token %s: %w", ev.TokenId, err)
			}
			out = append(out, &OnChainCertification{
				CertificationOfComplianceID: data.CertificationOfComplianceId,
				TokenID:                     ev.TokenId.Uint64(),
				ActivityNftID:               data.ActivityNftId.Uint64(),
				CertificationSchemeID:       data.CertificationSchemeId,
				CertificationBodyID:         data.CertificationBodyId,
				IssueDate:                   data.IssueDate,
				ExpiryDate:                  data.ExpiryDate,
				CertificationStatus:         data.CertificationStatus,
				ActivityURL:                 data.ActivityUrl,
				ActivityHash:                data.ActivityHash,
				CertificationURL:            data.CertificationUrl,
				CertificationHash:           data.CertificationHash,
				OwnerAddress:                owner.Hex(),
				TokenURI:                    uri,
				ContractAddress:             r.certAddr.Hex(),
				ChainID:                     r.chainID,
				TxHash:                      ev.Raw.TxHash.Hex(),
				BlockNumber:                 ev.Raw.BlockNumber,
			})
		}
		return it.Error()
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanCreditBatches yields every minted CarbonCreditBatchNFT via the BatchMinted
// event log, enriched with getBatch / ownerOf / getTBA / tokenURI. When a
// CarbonCredit address is configured, each batch also carries the ERC-20 balance
// currently held by its token-bound account, read at a single block height so
// every batch in one run is balanced as of the same moment.
func (r *Reader) ScanCreditBatches(ctx context.Context, fromBlock uint64) ([]*OnChainCreditBatch, error) {
	if r.creditBatch == nil {
		return nil, fmt.Errorf("no CarbonCreditBatchNFT address configured")
	}
	opts := &bind.CallOpts{Context: ctx}

	// Balances are current state, not history. Pin a height up front and read
	// every balance at it, so one run is a consistent snapshot rather than a
	// smear across blocks as the scan progresses.
	var decimals uint8
	var balanceBlock uint64
	balanceOpts := &bind.CallOpts{Context: ctx}
	if r.credit != nil {
		if err := retry(func() error {
			var e error
			decimals, e = r.credit.Decimals(opts)
			return e
		}); err != nil {
			return nil, fmt.Errorf("read CarbonCredit decimals: %w", err)
		}
		var err error
		if balanceBlock, err = r.head(ctx); err != nil {
			return nil, err
		}
		balanceOpts.BlockNumber = new(big.Int).SetUint64(balanceBlock)
	}

	var out []*OnChainCreditBatch
	err := r.eachWindow(ctx, fromBlock, func(start, end uint64) error {
		var it *contract.CarbonCreditBatchNFTBatchMintedIterator
		if err := retry(func() error {
			var e error
			it, e = r.creditBatch.FilterBatchMinted(&bind.FilterOpts{Start: start, End: &end, Context: ctx}, nil, nil, nil)
			return e
		}); err != nil {
			return fmt.Errorf("filter BatchMinted [%d,%d]: %w", start, end, err)
		}
		defer it.Close()
		for it.Next() {
			ev := it.Event
			var data contract.CarbonCreditBatchNFTBatchData
			var owner, tba common.Address
			var uri string
			if err := retry(func() error {
				var e error
				if data, e = r.creditBatch.GetBatch(opts, ev.TokenId); e != nil {
					return e
				}
				if owner, e = r.creditBatch.OwnerOf(opts, ev.TokenId); e != nil {
					return e
				}
				if tba, e = r.creditBatch.GetTBA(opts, ev.TokenId); e != nil {
					return e
				}
				uri, e = r.creditBatch.TokenURI(opts, ev.TokenId)
				return e
			}); err != nil {
				return fmt.Errorf("enrich credit batch token %s: %w", ev.TokenId, err)
			}

			batch := &OnChainCreditBatch{
				BatchKey:                fmt.Sprintf("%s:%s", data.ActivityNftId, data.CreditType),
				TokenID:                 ev.TokenId.Uint64(),
				CreditType:              data.CreditType,
				ActivityNftID:           data.ActivityNftId.Uint64(),
				ActivityURL:             data.ActivityUrl,
				ActivityHash:            data.ActivityHash,
				OperatorURL:             data.OperatorUrl,
				OperatorHash:            data.OperatorHash,
				CertificationURL:        data.CertificationUrl,
				CertificationHash:       data.CertificationHash,
				CertificationSchemeURL:  data.CertificationSchemeUrl,
				CertificationSchemeHash: data.CertificationSchemeHash,
				CertificationBodyURL:    data.CertificationBodyUrl,
				CertificationBodyHash:   data.CertificationBodyHash,
				TokenBoundAccount:       tba.Hex(),
				OwnerAddress:            owner.Hex(),
				TokenURI:                uri,
				ContractAddress:         r.creditBatchAddr.Hex(),
				ChainID:                 r.chainID,
				TxHash:                  ev.Raw.TxHash.Hex(),
				BlockNumber:             ev.Raw.BlockNumber,
			}

			if r.credit != nil {
				var bal *big.Int
				if err := retry(func() error {
					var e error
					bal, e = r.credit.BalanceOf(balanceOpts, tba)
					return e
				}); err != nil {
					return fmt.Errorf("read credit balance of batch %s account %s: %w", ev.TokenId, tba.Hex(), err)
				}
				batch.CreditBalance = bal.String()
				batch.CreditDecimals = decimals
				batch.CreditContractAddress = r.creditAddr.Hex()
				batch.BlockNumberBalance = balanceBlock
			}

			out = append(out, batch)
		}
		return it.Error()
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanCreditBalances returns the CarbonCredit (ERC-20) balance of every address
// that has ever held the token. ERC-20 keeps no holder list on-chain, so the
// holder set is derived from the Transfer log and each balance is then read at
// the current head.
//
// Zero balances are returned, not skipped: an address that has moved everything
// out still needs its cached record corrected to zero.
//
// batches, when non-nil, classifies which holders are the token-bound accounts
// of credit batches. Pass the result of ScanCreditBatches to avoid scanning the
// batch log twice; pass nil to label every holder a plain wallet.
func (r *Reader) ScanCreditBalances(ctx context.Context, fromBlock uint64, batches []*OnChainCreditBatch) ([]*OnChainCreditBalance, error) {
	if r.credit == nil {
		return nil, fmt.Errorf("no CarbonCredit address configured")
	}
	opts := &bind.CallOpts{Context: ctx}

	batchByAccount := make(map[common.Address]uint64, len(batches))
	for _, b := range batches {
		if b.TokenBoundAccount != "" {
			batchByAccount[common.HexToAddress(b.TokenBoundAccount)] = b.TokenID
		}
	}

	holders := map[common.Address]struct{}{}
	err := r.eachWindow(ctx, fromBlock, func(start, end uint64) error {
		var it *contract.CarbonCreditTransferIterator
		if err := retry(func() error {
			var e error
			it, e = r.credit.FilterTransfer(&bind.FilterOpts{Start: start, End: &end, Context: ctx}, nil, nil)
			return e
		}); err != nil {
			return fmt.Errorf("filter Transfer [%d,%d]: %w", start, end, err)
		}
		defer it.Close()
		for it.Next() {
			// The zero address is the mint/burn counterparty, never a holder.
			for _, addr := range []common.Address{it.Event.From, it.Event.To} {
				if addr != (common.Address{}) {
					holders[addr] = struct{}{}
				}
			}
		}
		return it.Error()
	})
	if err != nil {
		return nil, err
	}
	if len(holders) == 0 {
		return nil, nil
	}

	var decimals uint8
	var symbol string
	if err := retry(func() error {
		var e error
		if decimals, e = r.credit.Decimals(opts); e != nil {
			return e
		}
		symbol, e = r.credit.Symbol(opts)
		return e
	}); err != nil {
		return nil, fmt.Errorf("read CarbonCredit metadata: %w", err)
	}

	balanceBlock, err := r.head(ctx)
	if err != nil {
		return nil, err
	}
	balanceOpts := &bind.CallOpts{Context: ctx, BlockNumber: new(big.Int).SetUint64(balanceBlock)}

	// Map iteration is randomised; sort so runs are reproducible and logs diff.
	addrs := make([]common.Address, 0, len(holders))
	for addr := range holders {
		addrs = append(addrs, addr)
	}
	sort.Slice(addrs, func(i, j int) bool { return addrs[i].Hex() < addrs[j].Hex() })

	out := make([]*OnChainCreditBalance, 0, len(addrs))
	for _, addr := range addrs {
		var bal *big.Int
		if err := retry(func() error {
			var e error
			bal, e = r.credit.BalanceOf(balanceOpts, addr)
			return e
		}); err != nil {
			return nil, fmt.Errorf("read credit balance of %s: %w", addr.Hex(), err)
		}

		holderType := HolderWallet
		batchTokenID := uint64(0)
		if id, ok := batchByAccount[addr]; ok {
			holderType = HolderTokenBoundAccount
			batchTokenID = id
		}

		out = append(out, &OnChainCreditBalance{
			OwnerAddress:    addr.Hex(),
			Balance:         bal.String(),
			Decimals:        decimals,
			Symbol:          symbol,
			HolderType:      holderType,
			BatchTokenID:    batchTokenID,
			ContractAddress: r.creditAddr.Hex(),
			ChainID:         r.chainID,
			BlockNumber:     balanceBlock,
		})
	}
	return out, nil
}
