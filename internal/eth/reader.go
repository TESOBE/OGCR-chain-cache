// Package eth reads the OGCR token family off the OGCR chain (chain id 2025).
// The tokenizer mints ParcelNFT / ActivityNFT / CertificationNFT (plus credit
// batches, which are out of scope here); this reads those tokens back so they
// can be mirrored into the OBP `*_on_chain` dynamic entities.
package eth

import (
	"context"
	"fmt"
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

// Reader binds the three NFT contracts on one RPC connection.
type Reader struct {
	client       *ethclient.Client
	chainID      uint64
	parcel       *contract.ParcelNFT
	parcelAddr   common.Address
	activity     *contract.ActivityNFT
	activityAddr common.Address
	cert         *contract.CertificationNFT
	certAddr     common.Address
}

func NewReader(rpcURL, parcelAddr, activityAddr, certAddr string) (*Reader, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC: %w", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	pAddr := common.HexToAddress(parcelAddr)
	parcel, err := contract.NewParcelNFT(pAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind ParcelNFT: %w", err)
	}
	aAddr := common.HexToAddress(activityAddr)
	activity, err := contract.NewActivityNFT(aAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind ActivityNFT: %w", err)
	}
	cAddr := common.HexToAddress(certAddr)
	cert, err := contract.NewCertificationNFT(cAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind CertificationNFT: %w", err)
	}

	return &Reader{
		client:       client,
		chainID:      chainID.Uint64(),
		parcel:       parcel,
		parcelAddr:   pAddr,
		activity:     activity,
		activityAddr: aAddr,
		cert:         cert,
		certAddr:     cAddr,
	}, nil
}

func (r *Reader) ChainID() uint64 { return r.chainID }

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

// eachWindow calls fn for each [start,end] block window from fromBlock to the
// current head, in scanChunk-sized steps (the RPC caps eth_getLogs range).
func (r *Reader) eachWindow(ctx context.Context, fromBlock uint64, fn func(start, end uint64) error) error {
	var head uint64
	if err := retry(func() error {
		var e error
		head, e = r.client.BlockNumber(ctx)
		return e
	}); err != nil {
		return fmt.Errorf("get head block: %w", err)
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
