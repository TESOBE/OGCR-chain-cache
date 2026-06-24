// Package eth reads CarbonProjectNFT data off the OGCR chain (chain id 2025).
// The tokenizer mints one NFT per verified parcel_id; this reads those tokens
// back so they can be mirrored into the OBP `parcel_on_chain` dynamic entity.
package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/TESOBE/OGCR-chain-cache/internal/contract"
)

// OnChainParcel is the chain-side view of one minted parcel, shaped for the
// OBP `parcel_on_chain` entity. JSON tags are the entity field names.
type OnChainParcel struct {
	ParcelID        string `json:"parcel_id"`
	TokenID         uint64 `json:"token_id"`
	CarbonAmount    uint64 `json:"carbon_amount"`
	ProjectType     string `json:"project_type"`
	Methodology     string `json:"methodology"`
	Vintage         string `json:"vintage"`
	MintedAt        uint64 `json:"minted_at"`
	OwnerAddress    string `json:"owner_address"`
	TokenURI        string `json:"token_uri"`
	ContractAddress string `json:"contract_address"`
	ChainID         uint64 `json:"chain_id"`
	TxHash          string `json:"tx_hash"`
	BlockNumber     uint64 `json:"block_number"`
}

type Reader struct {
	client  *ethclient.Client
	nft     *contract.CarbonProjectNFT
	address common.Address
	chainID uint64
}

func NewReader(rpcURL, contractAddress string) (*Reader, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC: %w", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	addr := common.HexToAddress(contractAddress)
	nft, err := contract.NewCarbonProjectNFT(addr, client)
	if err != nil {
		return nil, fmt.Errorf("bind contract: %w", err)
	}
	return &Reader{client: client, nft: nft, address: addr, chainID: chainID.Uint64()}, nil
}

func (r *Reader) ChainID() uint64 { return r.chainID }

// build assembles an OnChainParcel from a token id. txHash/blockNumber come
// from the mint event when available (empty/zero otherwise).
func (r *Reader) build(ctx context.Context, tokenID *big.Int, parcelID, txHash string, blockNumber uint64) (*OnChainParcel, error) {
	opts := &bind.CallOpts{Context: ctx}
	data, err := r.nft.GetProjectData(opts, tokenID)
	if err != nil {
		return nil, fmt.Errorf("getProjectData(%s): %w", tokenID, err)
	}
	owner, err := r.nft.OwnerOf(opts, tokenID)
	if err != nil {
		return nil, fmt.Errorf("ownerOf(%s): %w", tokenID, err)
	}
	uri, err := r.nft.TokenURI(opts, tokenID)
	if err != nil {
		return nil, fmt.Errorf("tokenURI(%s): %w", tokenID, err)
	}
	if parcelID == "" {
		parcelID = data.ParcelId
	}
	return &OnChainParcel{
		ParcelID:        parcelID,
		TokenID:         tokenID.Uint64(),
		CarbonAmount:    data.CarbonAmount.Uint64(),
		ProjectType:     data.ProjectType,
		Methodology:     data.Methodology,
		Vintage:         data.Vintage,
		MintedAt:        data.MintedAt.Uint64(),
		OwnerAddress:    owner.Hex(),
		TokenURI:        uri,
		ContractAddress: r.address.Hex(),
		ChainID:         r.chainID,
		TxHash:          txHash,
		BlockNumber:     blockNumber,
	}, nil
}

// ScanMinted yields every parcel minted, via the CarbonProjectMinted event log.
func (r *Reader) ScanMinted(ctx context.Context, fromBlock uint64) ([]*OnChainParcel, error) {
	it, err := r.nft.FilterCarbonProjectMinted(
		&bind.FilterOpts{Start: fromBlock, Context: ctx}, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("filter CarbonProjectMinted: %w", err)
	}
	defer it.Close()

	var out []*OnChainParcel
	for it.Next() {
		ev := it.Event
		p, err := r.build(ctx, ev.ProjectId, ev.ParcelId,
			ev.Raw.TxHash.Hex(), ev.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, it.Error()
}

// ByParcelID reads the token(s) minted for a single parcel_id (no event scan,
// so tx_hash/block_number are left empty).
func (r *Reader) ByParcelID(ctx context.Context, parcelID string) ([]*OnChainParcel, error) {
	ids, err := r.nft.GetProjectsByParcelId(&bind.CallOpts{Context: ctx}, parcelID)
	if err != nil {
		return nil, fmt.Errorf("getProjectsByParcelId(%s): %w", parcelID, err)
	}
	var out []*OnChainParcel
	for _, id := range ids {
		p, err := r.build(ctx, id, parcelID, "", 0)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}
