// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ParcelNFTParcelData is an auto generated low-level Go binding around an user-defined struct.
type ParcelNFTParcelData struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}

// ParcelNFTMetaData contains all meta data concerning the ParcelNFT contract.
var ParcelNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParcel\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structParcelNFT.ParcelData\",\"components\":[{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_hash\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenIdByParcelId\",\"inputs\":[{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_hash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"parcels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"parcel_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_hash\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenIdByParcelId\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterUpdated\",\"inputs\":[{\"name\":\"oldMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParcelMinted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"parcel_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ParcelNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use ParcelNFTMetaData.ABI instead.
var ParcelNFTABI = ParcelNFTMetaData.ABI

// ParcelNFT is an auto generated Go binding around an Ethereum contract.
type ParcelNFT struct {
	ParcelNFTCaller     // Read-only binding to the contract
	ParcelNFTTransactor // Write-only binding to the contract
	ParcelNFTFilterer   // Log filterer for contract events
}

// ParcelNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParcelNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParcelNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParcelNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParcelNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParcelNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParcelNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParcelNFTSession struct {
	Contract     *ParcelNFT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParcelNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParcelNFTCallerSession struct {
	Contract *ParcelNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ParcelNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParcelNFTTransactorSession struct {
	Contract     *ParcelNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ParcelNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParcelNFTRaw struct {
	Contract *ParcelNFT // Generic contract binding to access the raw methods on
}

// ParcelNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParcelNFTCallerRaw struct {
	Contract *ParcelNFTCaller // Generic read-only contract binding to access the raw methods on
}

// ParcelNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParcelNFTTransactorRaw struct {
	Contract *ParcelNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParcelNFT creates a new instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFT(address common.Address, backend bind.ContractBackend) (*ParcelNFT, error) {
	contract, err := bindParcelNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParcelNFT{ParcelNFTCaller: ParcelNFTCaller{contract: contract}, ParcelNFTTransactor: ParcelNFTTransactor{contract: contract}, ParcelNFTFilterer: ParcelNFTFilterer{contract: contract}}, nil
}

// NewParcelNFTCaller creates a new read-only instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFTCaller(address common.Address, caller bind.ContractCaller) (*ParcelNFTCaller, error) {
	contract, err := bindParcelNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTCaller{contract: contract}, nil
}

// NewParcelNFTTransactor creates a new write-only instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*ParcelNFTTransactor, error) {
	contract, err := bindParcelNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTTransactor{contract: contract}, nil
}

// NewParcelNFTFilterer creates a new log filterer instance of ParcelNFT, bound to a specific deployed contract.
func NewParcelNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*ParcelNFTFilterer, error) {
	contract, err := bindParcelNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTFilterer{contract: contract}, nil
}

// bindParcelNFT binds a generic wrapper to an already deployed contract.
func bindParcelNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParcelNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParcelNFT *ParcelNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParcelNFT.Contract.ParcelNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParcelNFT *ParcelNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParcelNFT.Contract.ParcelNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParcelNFT *ParcelNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParcelNFT.Contract.ParcelNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParcelNFT *ParcelNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParcelNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParcelNFT *ParcelNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParcelNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParcelNFT *ParcelNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParcelNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ParcelNFT.Contract.BalanceOf(&_ParcelNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ParcelNFT.Contract.BalanceOf(&_ParcelNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.GetApproved(&_ParcelNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.GetApproved(&_ParcelNFT.CallOpts, tokenId)
}

// GetParcel is a free data retrieval call binding the contract method 0x4d0434f2.
//
// Solidity: function getParcel(uint256 tokenId) view returns((string,string,string))
func (_ParcelNFT *ParcelNFTCaller) GetParcel(opts *bind.CallOpts, tokenId *big.Int) (ParcelNFTParcelData, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "getParcel", tokenId)

	if err != nil {
		return *new(ParcelNFTParcelData), err
	}

	out0 := *abi.ConvertType(out[0], new(ParcelNFTParcelData)).(*ParcelNFTParcelData)

	return out0, err

}

// GetParcel is a free data retrieval call binding the contract method 0x4d0434f2.
//
// Solidity: function getParcel(uint256 tokenId) view returns((string,string,string))
func (_ParcelNFT *ParcelNFTSession) GetParcel(tokenId *big.Int) (ParcelNFTParcelData, error) {
	return _ParcelNFT.Contract.GetParcel(&_ParcelNFT.CallOpts, tokenId)
}

// GetParcel is a free data retrieval call binding the contract method 0x4d0434f2.
//
// Solidity: function getParcel(uint256 tokenId) view returns((string,string,string))
func (_ParcelNFT *ParcelNFTCallerSession) GetParcel(tokenId *big.Int) (ParcelNFTParcelData, error) {
	return _ParcelNFT.Contract.GetParcel(&_ParcelNFT.CallOpts, tokenId)
}

// GetTokenIdByParcelId is a free data retrieval call binding the contract method 0x8dfa037f.
//
// Solidity: function getTokenIdByParcelId(string parcel_id) view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) GetTokenIdByParcelId(opts *bind.CallOpts, parcel_id string) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "getTokenIdByParcelId", parcel_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenIdByParcelId is a free data retrieval call binding the contract method 0x8dfa037f.
//
// Solidity: function getTokenIdByParcelId(string parcel_id) view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) GetTokenIdByParcelId(parcel_id string) (*big.Int, error) {
	return _ParcelNFT.Contract.GetTokenIdByParcelId(&_ParcelNFT.CallOpts, parcel_id)
}

// GetTokenIdByParcelId is a free data retrieval call binding the contract method 0x8dfa037f.
//
// Solidity: function getTokenIdByParcelId(string parcel_id) view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) GetTokenIdByParcelId(parcel_id string) (*big.Int, error) {
	return _ParcelNFT.Contract.GetTokenIdByParcelId(&_ParcelNFT.CallOpts, parcel_id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParcelNFT *ParcelNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParcelNFT *ParcelNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ParcelNFT.Contract.IsApprovedForAll(&_ParcelNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParcelNFT *ParcelNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ParcelNFT.Contract.IsApprovedForAll(&_ParcelNFT.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ParcelNFT *ParcelNFTCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ParcelNFT *ParcelNFTSession) Minter() (common.Address, error) {
	return _ParcelNFT.Contract.Minter(&_ParcelNFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) Minter() (common.Address, error) {
	return _ParcelNFT.Contract.Minter(&_ParcelNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParcelNFT *ParcelNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParcelNFT *ParcelNFTSession) Name() (string, error) {
	return _ParcelNFT.Contract.Name(&_ParcelNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParcelNFT *ParcelNFTCallerSession) Name() (string, error) {
	return _ParcelNFT.Contract.Name(&_ParcelNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) NextId() (*big.Int, error) {
	return _ParcelNFT.Contract.NextId(&_ParcelNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) NextId() (*big.Int, error) {
	return _ParcelNFT.Contract.NextId(&_ParcelNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParcelNFT *ParcelNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParcelNFT *ParcelNFTSession) Owner() (common.Address, error) {
	return _ParcelNFT.Contract.Owner(&_ParcelNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) Owner() (common.Address, error) {
	return _ParcelNFT.Contract.Owner(&_ParcelNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.OwnerOf(&_ParcelNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParcelNFT *ParcelNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ParcelNFT.Contract.OwnerOf(&_ParcelNFT.CallOpts, tokenId)
}

// Parcels is a free data retrieval call binding the contract method 0xa0fa8ede.
//
// Solidity: function parcels(uint256 ) view returns(string parcel_id, string parcel_uri, string parcel_hash)
func (_ParcelNFT *ParcelNFTCaller) Parcels(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "parcels", arg0)

	outstruct := new(struct {
		ParcelId   string
		ParcelUri  string
		ParcelHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParcelId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ParcelUri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ParcelHash = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Parcels is a free data retrieval call binding the contract method 0xa0fa8ede.
//
// Solidity: function parcels(uint256 ) view returns(string parcel_id, string parcel_uri, string parcel_hash)
func (_ParcelNFT *ParcelNFTSession) Parcels(arg0 *big.Int) (struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}, error) {
	return _ParcelNFT.Contract.Parcels(&_ParcelNFT.CallOpts, arg0)
}

// Parcels is a free data retrieval call binding the contract method 0xa0fa8ede.
//
// Solidity: function parcels(uint256 ) view returns(string parcel_id, string parcel_uri, string parcel_hash)
func (_ParcelNFT *ParcelNFTCallerSession) Parcels(arg0 *big.Int) (struct {
	ParcelId   string
	ParcelUri  string
	ParcelHash string
}, error) {
	return _ParcelNFT.Contract.Parcels(&_ParcelNFT.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParcelNFT *ParcelNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParcelNFT *ParcelNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParcelNFT.Contract.SupportsInterface(&_ParcelNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParcelNFT *ParcelNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParcelNFT.Contract.SupportsInterface(&_ParcelNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParcelNFT *ParcelNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParcelNFT *ParcelNFTSession) Symbol() (string, error) {
	return _ParcelNFT.Contract.Symbol(&_ParcelNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParcelNFT *ParcelNFTCallerSession) Symbol() (string, error) {
	return _ParcelNFT.Contract.Symbol(&_ParcelNFT.CallOpts)
}

// TokenIdByParcelId is a free data retrieval call binding the contract method 0x6a94eb85.
//
// Solidity: function tokenIdByParcelId(string ) view returns(uint256)
func (_ParcelNFT *ParcelNFTCaller) TokenIdByParcelId(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "tokenIdByParcelId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdByParcelId is a free data retrieval call binding the contract method 0x6a94eb85.
//
// Solidity: function tokenIdByParcelId(string ) view returns(uint256)
func (_ParcelNFT *ParcelNFTSession) TokenIdByParcelId(arg0 string) (*big.Int, error) {
	return _ParcelNFT.Contract.TokenIdByParcelId(&_ParcelNFT.CallOpts, arg0)
}

// TokenIdByParcelId is a free data retrieval call binding the contract method 0x6a94eb85.
//
// Solidity: function tokenIdByParcelId(string ) view returns(uint256)
func (_ParcelNFT *ParcelNFTCallerSession) TokenIdByParcelId(arg0 string) (*big.Int, error) {
	return _ParcelNFT.Contract.TokenIdByParcelId(&_ParcelNFT.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParcelNFT *ParcelNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ParcelNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParcelNFT *ParcelNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ParcelNFT.Contract.TokenURI(&_ParcelNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParcelNFT *ParcelNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ParcelNFT.Contract.TokenURI(&_ParcelNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Approve(&_ParcelNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Approve(&_ParcelNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x4a0706be.
//
// Solidity: function mint(address to, string parcel_id, string parcel_uri, string parcel_hash) returns(uint256 tokenId)
func (_ParcelNFT *ParcelNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, parcel_id string, parcel_uri string, parcel_hash string) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "mint", to, parcel_id, parcel_uri, parcel_hash)
}

// Mint is a paid mutator transaction binding the contract method 0x4a0706be.
//
// Solidity: function mint(address to, string parcel_id, string parcel_uri, string parcel_hash) returns(uint256 tokenId)
func (_ParcelNFT *ParcelNFTSession) Mint(to common.Address, parcel_id string, parcel_uri string, parcel_hash string) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Mint(&_ParcelNFT.TransactOpts, to, parcel_id, parcel_uri, parcel_hash)
}

// Mint is a paid mutator transaction binding the contract method 0x4a0706be.
//
// Solidity: function mint(address to, string parcel_id, string parcel_uri, string parcel_hash) returns(uint256 tokenId)
func (_ParcelNFT *ParcelNFTTransactorSession) Mint(to common.Address, parcel_id string, parcel_uri string, parcel_hash string) (*types.Transaction, error) {
	return _ParcelNFT.Contract.Mint(&_ParcelNFT.TransactOpts, to, parcel_id, parcel_uri, parcel_hash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParcelNFT *ParcelNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParcelNFT *ParcelNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParcelNFT.Contract.RenounceOwnership(&_ParcelNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParcelNFT *ParcelNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParcelNFT.Contract.RenounceOwnership(&_ParcelNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParcelNFT *ParcelNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParcelNFT *ParcelNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom0(&_ParcelNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SafeTransferFrom0(&_ParcelNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParcelNFT *ParcelNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParcelNFT *ParcelNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetApprovalForAll(&_ParcelNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetApprovalForAll(&_ParcelNFT.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ParcelNFT *ParcelNFTTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ParcelNFT *ParcelNFTSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetMinter(&_ParcelNFT.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.SetMinter(&_ParcelNFT.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferFrom(&_ParcelNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParcelNFT *ParcelNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ParcelNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParcelNFT *ParcelNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferOwnership(&_ParcelNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParcelNFT *ParcelNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParcelNFT.Contract.TransferOwnership(&_ParcelNFT.TransactOpts, newOwner)
}

// ParcelNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ParcelNFT contract.
type ParcelNFTApprovalIterator struct {
	Event *ParcelNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTApproval represents a Approval event raised by the ParcelNFT contract.
type ParcelNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ParcelNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTApprovalIterator{contract: _ParcelNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ParcelNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTApproval)
				if err := _ParcelNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) ParseApproval(log types.Log) (*ParcelNFTApproval, error) {
	event := new(ParcelNFTApproval)
	if err := _ParcelNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ParcelNFT contract.
type ParcelNFTApprovalForAllIterator struct {
	Event *ParcelNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTApprovalForAll represents a ApprovalForAll event raised by the ParcelNFT contract.
type ParcelNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParcelNFT *ParcelNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ParcelNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTApprovalForAllIterator{contract: _ParcelNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParcelNFT *ParcelNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ParcelNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTApprovalForAll)
				if err := _ParcelNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParcelNFT *ParcelNFTFilterer) ParseApprovalForAll(log types.Log) (*ParcelNFTApprovalForAll, error) {
	event := new(ParcelNFTApprovalForAll)
	if err := _ParcelNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the ParcelNFT contract.
type ParcelNFTMinterUpdatedIterator struct {
	Event *ParcelNFTMinterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTMinterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTMinterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTMinterUpdated represents a MinterUpdated event raised by the ParcelNFT contract.
type ParcelNFTMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ParcelNFT *ParcelNFTFilterer) FilterMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*ParcelNFTMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTMinterUpdatedIterator{contract: _ParcelNFT.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ParcelNFT *ParcelNFTFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *ParcelNFTMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTMinterUpdated)
				if err := _ParcelNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterUpdated is a log parse operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ParcelNFT *ParcelNFTFilterer) ParseMinterUpdated(log types.Log) (*ParcelNFTMinterUpdated, error) {
	event := new(ParcelNFTMinterUpdated)
	if err := _ParcelNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParcelNFT contract.
type ParcelNFTOwnershipTransferredIterator struct {
	Event *ParcelNFTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTOwnershipTransferred represents a OwnershipTransferred event raised by the ParcelNFT contract.
type ParcelNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParcelNFT *ParcelNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParcelNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTOwnershipTransferredIterator{contract: _ParcelNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParcelNFT *ParcelNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParcelNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTOwnershipTransferred)
				if err := _ParcelNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParcelNFT *ParcelNFTFilterer) ParseOwnershipTransferred(log types.Log) (*ParcelNFTOwnershipTransferred, error) {
	event := new(ParcelNFTOwnershipTransferred)
	if err := _ParcelNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTParcelMintedIterator is returned from FilterParcelMinted and is used to iterate over the raw logs and unpacked data for ParcelMinted events raised by the ParcelNFT contract.
type ParcelNFTParcelMintedIterator struct {
	Event *ParcelNFTParcelMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTParcelMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTParcelMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTParcelMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTParcelMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTParcelMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTParcelMinted represents a ParcelMinted event raised by the ParcelNFT contract.
type ParcelNFTParcelMinted struct {
	TokenId  *big.Int
	To       common.Address
	ParcelId string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterParcelMinted is a free log retrieval operation binding the contract event 0xe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364.
//
// Solidity: event ParcelMinted(uint256 indexed tokenId, address indexed to, string parcel_id)
func (_ParcelNFT *ParcelNFTFilterer) FilterParcelMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address) (*ParcelNFTParcelMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "ParcelMinted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTParcelMintedIterator{contract: _ParcelNFT.contract, event: "ParcelMinted", logs: logs, sub: sub}, nil
}

// WatchParcelMinted is a free log subscription operation binding the contract event 0xe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364.
//
// Solidity: event ParcelMinted(uint256 indexed tokenId, address indexed to, string parcel_id)
func (_ParcelNFT *ParcelNFTFilterer) WatchParcelMinted(opts *bind.WatchOpts, sink chan<- *ParcelNFTParcelMinted, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "ParcelMinted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTParcelMinted)
				if err := _ParcelNFT.contract.UnpackLog(event, "ParcelMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParcelMinted is a log parse operation binding the contract event 0xe44c0e9e19a3e72b64f7ab511884d971c495593ea485e76b76271a4cef0e1364.
//
// Solidity: event ParcelMinted(uint256 indexed tokenId, address indexed to, string parcel_id)
func (_ParcelNFT *ParcelNFTFilterer) ParseParcelMinted(log types.Log) (*ParcelNFTParcelMinted, error) {
	event := new(ParcelNFTParcelMinted)
	if err := _ParcelNFT.contract.UnpackLog(event, "ParcelMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParcelNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ParcelNFT contract.
type ParcelNFTTransferIterator struct {
	Event *ParcelNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParcelNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParcelNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParcelNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParcelNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParcelNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParcelNFTTransfer represents a Transfer event raised by the ParcelNFT contract.
type ParcelNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ParcelNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ParcelNFTTransferIterator{contract: _ParcelNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ParcelNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParcelNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParcelNFTTransfer)
				if err := _ParcelNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParcelNFT *ParcelNFTFilterer) ParseTransfer(log types.Log) (*ParcelNFTTransfer, error) {
	event := new(ParcelNFTTransfer)
	if err := _ParcelNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
