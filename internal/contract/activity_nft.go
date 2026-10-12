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

// ActivityNFTActivityData is an auto generated low-level Go binding around an user-defined struct.
type ActivityNFTActivityData struct {
	ActivityId            string
	OperatorId            string
	OperatorName          string
	Name                  string
	Summary               string
	Website               string
	ActivityType          string
	UnitTypes             string
	City                  string
	CountryId             string
	CertificationSchemeId string
	ParcelIds             []string
	StartDate             string
	EndDate               string
	ActivityUrl           string
	ActivityHash          string
}

// ActivityNFTMetaData contains all meta data concerning the ActivityNFT contract.
var ActivityNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActivity\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structActivityNFT.ActivityData\",\"components\":[{\"name\":\"activity_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_type\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"unit_types\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"city\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"country_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_ids\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"start_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"end_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_hash\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structActivityNFT.ActivityData\",\"components\":[{\"name\":\"activity_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_type\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"unit_types\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"city\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"country_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parcel_ids\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"start_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"end_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_hash\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenIdByActivityId\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivityMinted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activity_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"operator_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterUpdated\",\"inputs\":[{\"name\":\"oldMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ActivityNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use ActivityNFTMetaData.ABI instead.
var ActivityNFTABI = ActivityNFTMetaData.ABI

// ActivityNFT is an auto generated Go binding around an Ethereum contract.
type ActivityNFT struct {
	ActivityNFTCaller     // Read-only binding to the contract
	ActivityNFTTransactor // Write-only binding to the contract
	ActivityNFTFilterer   // Log filterer for contract events
}

// ActivityNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityNFTSession struct {
	Contract     *ActivityNFT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivityNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityNFTCallerSession struct {
	Contract *ActivityNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ActivityNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityNFTTransactorSession struct {
	Contract     *ActivityNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ActivityNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityNFTRaw struct {
	Contract *ActivityNFT // Generic contract binding to access the raw methods on
}

// ActivityNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityNFTCallerRaw struct {
	Contract *ActivityNFTCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityNFTTransactorRaw struct {
	Contract *ActivityNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityNFT creates a new instance of ActivityNFT, bound to a specific deployed contract.
func NewActivityNFT(address common.Address, backend bind.ContractBackend) (*ActivityNFT, error) {
	contract, err := bindActivityNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityNFT{ActivityNFTCaller: ActivityNFTCaller{contract: contract}, ActivityNFTTransactor: ActivityNFTTransactor{contract: contract}, ActivityNFTFilterer: ActivityNFTFilterer{contract: contract}}, nil
}

// NewActivityNFTCaller creates a new read-only instance of ActivityNFT, bound to a specific deployed contract.
func NewActivityNFTCaller(address common.Address, caller bind.ContractCaller) (*ActivityNFTCaller, error) {
	contract, err := bindActivityNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTCaller{contract: contract}, nil
}

// NewActivityNFTTransactor creates a new write-only instance of ActivityNFT, bound to a specific deployed contract.
func NewActivityNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityNFTTransactor, error) {
	contract, err := bindActivityNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTTransactor{contract: contract}, nil
}

// NewActivityNFTFilterer creates a new log filterer instance of ActivityNFT, bound to a specific deployed contract.
func NewActivityNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityNFTFilterer, error) {
	contract, err := bindActivityNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTFilterer{contract: contract}, nil
}

// bindActivityNFT binds a generic wrapper to an already deployed contract.
func bindActivityNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ActivityNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityNFT *ActivityNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActivityNFT.Contract.ActivityNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityNFT *ActivityNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityNFT.Contract.ActivityNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityNFT *ActivityNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityNFT.Contract.ActivityNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityNFT *ActivityNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActivityNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityNFT *ActivityNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityNFT *ActivityNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ActivityNFT *ActivityNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ActivityNFT *ActivityNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ActivityNFT.Contract.BalanceOf(&_ActivityNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ActivityNFT *ActivityNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ActivityNFT.Contract.BalanceOf(&_ActivityNFT.CallOpts, owner)
}

// GetActivity is a free data retrieval call binding the contract method 0x922635d8.
//
// Solidity: function getActivity(uint256 tokenId) view returns((string,string,string,string,string,string,string,string,string,string,string,string[],string,string,string,string))
func (_ActivityNFT *ActivityNFTCaller) GetActivity(opts *bind.CallOpts, tokenId *big.Int) (ActivityNFTActivityData, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "getActivity", tokenId)

	if err != nil {
		return *new(ActivityNFTActivityData), err
	}

	out0 := *abi.ConvertType(out[0], new(ActivityNFTActivityData)).(*ActivityNFTActivityData)

	return out0, err

}

// GetActivity is a free data retrieval call binding the contract method 0x922635d8.
//
// Solidity: function getActivity(uint256 tokenId) view returns((string,string,string,string,string,string,string,string,string,string,string,string[],string,string,string,string))
func (_ActivityNFT *ActivityNFTSession) GetActivity(tokenId *big.Int) (ActivityNFTActivityData, error) {
	return _ActivityNFT.Contract.GetActivity(&_ActivityNFT.CallOpts, tokenId)
}

// GetActivity is a free data retrieval call binding the contract method 0x922635d8.
//
// Solidity: function getActivity(uint256 tokenId) view returns((string,string,string,string,string,string,string,string,string,string,string,string[],string,string,string,string))
func (_ActivityNFT *ActivityNFTCallerSession) GetActivity(tokenId *big.Int) (ActivityNFTActivityData, error) {
	return _ActivityNFT.Contract.GetActivity(&_ActivityNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ActivityNFT *ActivityNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ActivityNFT *ActivityNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ActivityNFT.Contract.GetApproved(&_ActivityNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ActivityNFT *ActivityNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ActivityNFT.Contract.GetApproved(&_ActivityNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ActivityNFT *ActivityNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ActivityNFT *ActivityNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ActivityNFT.Contract.IsApprovedForAll(&_ActivityNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ActivityNFT *ActivityNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ActivityNFT.Contract.IsApprovedForAll(&_ActivityNFT.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ActivityNFT *ActivityNFTCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ActivityNFT *ActivityNFTSession) Minter() (common.Address, error) {
	return _ActivityNFT.Contract.Minter(&_ActivityNFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ActivityNFT *ActivityNFTCallerSession) Minter() (common.Address, error) {
	return _ActivityNFT.Contract.Minter(&_ActivityNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ActivityNFT *ActivityNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ActivityNFT *ActivityNFTSession) Name() (string, error) {
	return _ActivityNFT.Contract.Name(&_ActivityNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ActivityNFT *ActivityNFTCallerSession) Name() (string, error) {
	return _ActivityNFT.Contract.Name(&_ActivityNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ActivityNFT *ActivityNFTCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ActivityNFT *ActivityNFTSession) NextId() (*big.Int, error) {
	return _ActivityNFT.Contract.NextId(&_ActivityNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_ActivityNFT *ActivityNFTCallerSession) NextId() (*big.Int, error) {
	return _ActivityNFT.Contract.NextId(&_ActivityNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ActivityNFT *ActivityNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ActivityNFT *ActivityNFTSession) Owner() (common.Address, error) {
	return _ActivityNFT.Contract.Owner(&_ActivityNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ActivityNFT *ActivityNFTCallerSession) Owner() (common.Address, error) {
	return _ActivityNFT.Contract.Owner(&_ActivityNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ActivityNFT *ActivityNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ActivityNFT *ActivityNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ActivityNFT.Contract.OwnerOf(&_ActivityNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ActivityNFT *ActivityNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ActivityNFT.Contract.OwnerOf(&_ActivityNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ActivityNFT *ActivityNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ActivityNFT *ActivityNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ActivityNFT.Contract.SupportsInterface(&_ActivityNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ActivityNFT *ActivityNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ActivityNFT.Contract.SupportsInterface(&_ActivityNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ActivityNFT *ActivityNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ActivityNFT *ActivityNFTSession) Symbol() (string, error) {
	return _ActivityNFT.Contract.Symbol(&_ActivityNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ActivityNFT *ActivityNFTCallerSession) Symbol() (string, error) {
	return _ActivityNFT.Contract.Symbol(&_ActivityNFT.CallOpts)
}

// TokenIdByActivityId is a free data retrieval call binding the contract method 0x948029ba.
//
// Solidity: function tokenIdByActivityId(string ) view returns(uint256)
func (_ActivityNFT *ActivityNFTCaller) TokenIdByActivityId(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "tokenIdByActivityId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdByActivityId is a free data retrieval call binding the contract method 0x948029ba.
//
// Solidity: function tokenIdByActivityId(string ) view returns(uint256)
func (_ActivityNFT *ActivityNFTSession) TokenIdByActivityId(arg0 string) (*big.Int, error) {
	return _ActivityNFT.Contract.TokenIdByActivityId(&_ActivityNFT.CallOpts, arg0)
}

// TokenIdByActivityId is a free data retrieval call binding the contract method 0x948029ba.
//
// Solidity: function tokenIdByActivityId(string ) view returns(uint256)
func (_ActivityNFT *ActivityNFTCallerSession) TokenIdByActivityId(arg0 string) (*big.Int, error) {
	return _ActivityNFT.Contract.TokenIdByActivityId(&_ActivityNFT.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ActivityNFT *ActivityNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ActivityNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ActivityNFT *ActivityNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ActivityNFT.Contract.TokenURI(&_ActivityNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ActivityNFT *ActivityNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ActivityNFT.Contract.TokenURI(&_ActivityNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.Contract.Approve(&_ActivityNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.Contract.Approve(&_ActivityNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x644ab4d4.
//
// Solidity: function mint(address to, (string,string,string,string,string,string,string,string,string,string,string,string[],string,string,string,string) data) returns(uint256 tokenId)
func (_ActivityNFT *ActivityNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, data ActivityNFTActivityData) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "mint", to, data)
}

// Mint is a paid mutator transaction binding the contract method 0x644ab4d4.
//
// Solidity: function mint(address to, (string,string,string,string,string,string,string,string,string,string,string,string[],string,string,string,string) data) returns(uint256 tokenId)
func (_ActivityNFT *ActivityNFTSession) Mint(to common.Address, data ActivityNFTActivityData) (*types.Transaction, error) {
	return _ActivityNFT.Contract.Mint(&_ActivityNFT.TransactOpts, to, data)
}

// Mint is a paid mutator transaction binding the contract method 0x644ab4d4.
//
// Solidity: function mint(address to, (string,string,string,string,string,string,string,string,string,string,string,string[],string,string,string,string) data) returns(uint256 tokenId)
func (_ActivityNFT *ActivityNFTTransactorSession) Mint(to common.Address, data ActivityNFTActivityData) (*types.Transaction, error) {
	return _ActivityNFT.Contract.Mint(&_ActivityNFT.TransactOpts, to, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ActivityNFT *ActivityNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ActivityNFT *ActivityNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _ActivityNFT.Contract.RenounceOwnership(&_ActivityNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ActivityNFT *ActivityNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ActivityNFT.Contract.RenounceOwnership(&_ActivityNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SafeTransferFrom(&_ActivityNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SafeTransferFrom(&_ActivityNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ActivityNFT *ActivityNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ActivityNFT *ActivityNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SafeTransferFrom0(&_ActivityNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SafeTransferFrom0(&_ActivityNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ActivityNFT *ActivityNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ActivityNFT *ActivityNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SetApprovalForAll(&_ActivityNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SetApprovalForAll(&_ActivityNFT.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ActivityNFT *ActivityNFTTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ActivityNFT *ActivityNFTSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SetMinter(&_ActivityNFT.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ActivityNFT.Contract.SetMinter(&_ActivityNFT.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.Contract.TransferFrom(&_ActivityNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ActivityNFT.Contract.TransferFrom(&_ActivityNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ActivityNFT *ActivityNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ActivityNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ActivityNFT *ActivityNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ActivityNFT.Contract.TransferOwnership(&_ActivityNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ActivityNFT *ActivityNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ActivityNFT.Contract.TransferOwnership(&_ActivityNFT.TransactOpts, newOwner)
}

// ActivityNFTActivityMintedIterator is returned from FilterActivityMinted and is used to iterate over the raw logs and unpacked data for ActivityMinted events raised by the ActivityNFT contract.
type ActivityNFTActivityMintedIterator struct {
	Event *ActivityNFTActivityMinted // Event containing the contract specifics and raw log

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
func (it *ActivityNFTActivityMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityNFTActivityMinted)
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
		it.Event = new(ActivityNFTActivityMinted)
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
func (it *ActivityNFTActivityMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityNFTActivityMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityNFTActivityMinted represents a ActivityMinted event raised by the ActivityNFT contract.
type ActivityNFTActivityMinted struct {
	TokenId    *big.Int
	To         common.Address
	ActivityId string
	OperatorId string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActivityMinted is a free log retrieval operation binding the contract event 0xa1cb90ced1d6a6ec670e2ff58ee7ac737228a167ba9a7a6f77d6edb1bf4ffa63.
//
// Solidity: event ActivityMinted(uint256 indexed tokenId, address indexed to, string activity_id, string operator_id)
func (_ActivityNFT *ActivityNFTFilterer) FilterActivityMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address) (*ActivityNFTActivityMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ActivityNFT.contract.FilterLogs(opts, "ActivityMinted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTActivityMintedIterator{contract: _ActivityNFT.contract, event: "ActivityMinted", logs: logs, sub: sub}, nil
}

// WatchActivityMinted is a free log subscription operation binding the contract event 0xa1cb90ced1d6a6ec670e2ff58ee7ac737228a167ba9a7a6f77d6edb1bf4ffa63.
//
// Solidity: event ActivityMinted(uint256 indexed tokenId, address indexed to, string activity_id, string operator_id)
func (_ActivityNFT *ActivityNFTFilterer) WatchActivityMinted(opts *bind.WatchOpts, sink chan<- *ActivityNFTActivityMinted, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ActivityNFT.contract.WatchLogs(opts, "ActivityMinted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityNFTActivityMinted)
				if err := _ActivityNFT.contract.UnpackLog(event, "ActivityMinted", log); err != nil {
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

// ParseActivityMinted is a log parse operation binding the contract event 0xa1cb90ced1d6a6ec670e2ff58ee7ac737228a167ba9a7a6f77d6edb1bf4ffa63.
//
// Solidity: event ActivityMinted(uint256 indexed tokenId, address indexed to, string activity_id, string operator_id)
func (_ActivityNFT *ActivityNFTFilterer) ParseActivityMinted(log types.Log) (*ActivityNFTActivityMinted, error) {
	event := new(ActivityNFTActivityMinted)
	if err := _ActivityNFT.contract.UnpackLog(event, "ActivityMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivityNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ActivityNFT contract.
type ActivityNFTApprovalIterator struct {
	Event *ActivityNFTApproval // Event containing the contract specifics and raw log

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
func (it *ActivityNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityNFTApproval)
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
		it.Event = new(ActivityNFTApproval)
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
func (it *ActivityNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityNFTApproval represents a Approval event raised by the ActivityNFT contract.
type ActivityNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ActivityNFT *ActivityNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ActivityNFTApprovalIterator, error) {

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

	logs, sub, err := _ActivityNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTApprovalIterator{contract: _ActivityNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ActivityNFT *ActivityNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ActivityNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ActivityNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityNFTApproval)
				if err := _ActivityNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ActivityNFT *ActivityNFTFilterer) ParseApproval(log types.Log) (*ActivityNFTApproval, error) {
	event := new(ActivityNFTApproval)
	if err := _ActivityNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivityNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ActivityNFT contract.
type ActivityNFTApprovalForAllIterator struct {
	Event *ActivityNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ActivityNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityNFTApprovalForAll)
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
		it.Event = new(ActivityNFTApprovalForAll)
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
func (it *ActivityNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityNFTApprovalForAll represents a ApprovalForAll event raised by the ActivityNFT contract.
type ActivityNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ActivityNFT *ActivityNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ActivityNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ActivityNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTApprovalForAllIterator{contract: _ActivityNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ActivityNFT *ActivityNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ActivityNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ActivityNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityNFTApprovalForAll)
				if err := _ActivityNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ActivityNFT *ActivityNFTFilterer) ParseApprovalForAll(log types.Log) (*ActivityNFTApprovalForAll, error) {
	event := new(ActivityNFTApprovalForAll)
	if err := _ActivityNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivityNFTMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the ActivityNFT contract.
type ActivityNFTMinterUpdatedIterator struct {
	Event *ActivityNFTMinterUpdated // Event containing the contract specifics and raw log

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
func (it *ActivityNFTMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityNFTMinterUpdated)
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
		it.Event = new(ActivityNFTMinterUpdated)
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
func (it *ActivityNFTMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityNFTMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityNFTMinterUpdated represents a MinterUpdated event raised by the ActivityNFT contract.
type ActivityNFTMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ActivityNFT *ActivityNFTFilterer) FilterMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*ActivityNFTMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _ActivityNFT.contract.FilterLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTMinterUpdatedIterator{contract: _ActivityNFT.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_ActivityNFT *ActivityNFTFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *ActivityNFTMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _ActivityNFT.contract.WatchLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityNFTMinterUpdated)
				if err := _ActivityNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
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
func (_ActivityNFT *ActivityNFTFilterer) ParseMinterUpdated(log types.Log) (*ActivityNFTMinterUpdated, error) {
	event := new(ActivityNFTMinterUpdated)
	if err := _ActivityNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivityNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ActivityNFT contract.
type ActivityNFTOwnershipTransferredIterator struct {
	Event *ActivityNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ActivityNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityNFTOwnershipTransferred)
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
		it.Event = new(ActivityNFTOwnershipTransferred)
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
func (it *ActivityNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityNFTOwnershipTransferred represents a OwnershipTransferred event raised by the ActivityNFT contract.
type ActivityNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActivityNFT *ActivityNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ActivityNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ActivityNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTOwnershipTransferredIterator{contract: _ActivityNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActivityNFT *ActivityNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ActivityNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ActivityNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityNFTOwnershipTransferred)
				if err := _ActivityNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ActivityNFT *ActivityNFTFilterer) ParseOwnershipTransferred(log types.Log) (*ActivityNFTOwnershipTransferred, error) {
	event := new(ActivityNFTOwnershipTransferred)
	if err := _ActivityNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivityNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ActivityNFT contract.
type ActivityNFTTransferIterator struct {
	Event *ActivityNFTTransfer // Event containing the contract specifics and raw log

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
func (it *ActivityNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityNFTTransfer)
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
		it.Event = new(ActivityNFTTransfer)
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
func (it *ActivityNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityNFTTransfer represents a Transfer event raised by the ActivityNFT contract.
type ActivityNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ActivityNFT *ActivityNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ActivityNFTTransferIterator, error) {

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

	logs, sub, err := _ActivityNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ActivityNFTTransferIterator{contract: _ActivityNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ActivityNFT *ActivityNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ActivityNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ActivityNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityNFTTransfer)
				if err := _ActivityNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ActivityNFT *ActivityNFTFilterer) ParseTransfer(log types.Log) (*ActivityNFTTransfer, error) {
	event := new(ActivityNFTTransfer)
	if err := _ActivityNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
