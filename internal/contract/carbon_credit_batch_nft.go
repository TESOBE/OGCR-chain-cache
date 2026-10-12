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

// CarbonCreditBatchNFTBatchData is an auto generated low-level Go binding around an user-defined struct.
type CarbonCreditBatchNFTBatchData struct {
	CarbonCreditBatchId         string
	BatchName                   string
	ActivityId                  string
	ActivityNftId               *big.Int
	ActivityName                string
	ActivityUrl                 string
	ActivityHash                string
	CountryId                   string
	OperatorUrl                 string
	OperatorHash                string
	CertificateOfComplianceUri  string
	CertificateOfComplianceHash string
	UnitType                    string
	MonitoringPeriodStartDate   string
	MonitoringPeriodEndDate     string
	Amount                      *big.Int
	StatusCode                  string
	IssueDate                   string
	ExpiryDate                  string
	CertificationSchemeUrl      string
	CertificationSchemeHash     string
	CertificationBodyUrl        string
	CertificationBodyHash       string
}

// CarbonCreditBatchNFTMetaData contains all meta data concerning the CarbonCreditBatchNFT contract.
var CarbonCreditBatchNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_accountImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACCOUNT_IMPLEMENTATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC6551_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatch\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCarbonCreditBatchNFT.BatchData\",\"components\":[{\"name\":\"carbon_credit_batch_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"batch_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_nft_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activity_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"country_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"unit_type\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"monitoring_period_start_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"monitoring_period_end_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status_code\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"issue_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expiry_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_hash\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTBA\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structCarbonCreditBatchNFT.BatchData\",\"components\":[{\"name\":\"carbon_credit_batch_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"batch_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_nft_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activity_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"country_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operator_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"unit_type\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"monitoring_period_start_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"monitoring_period_end_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status_code\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"issue_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expiry_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_hash\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tba\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBoundAccount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenIdByBatchId\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMinted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activity_nft_id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"carbon_credit_batch_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"unit_type\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tba\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterUpdated\",\"inputs\":[{\"name\":\"oldMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// CarbonCreditBatchNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CarbonCreditBatchNFTMetaData.ABI instead.
var CarbonCreditBatchNFTABI = CarbonCreditBatchNFTMetaData.ABI

// CarbonCreditBatchNFT is an auto generated Go binding around an Ethereum contract.
type CarbonCreditBatchNFT struct {
	CarbonCreditBatchNFTCaller     // Read-only binding to the contract
	CarbonCreditBatchNFTTransactor // Write-only binding to the contract
	CarbonCreditBatchNFTFilterer   // Log filterer for contract events
}

// CarbonCreditBatchNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarbonCreditBatchNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarbonCreditBatchNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarbonCreditBatchNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarbonCreditBatchNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CarbonCreditBatchNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarbonCreditBatchNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarbonCreditBatchNFTSession struct {
	Contract     *CarbonCreditBatchNFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CarbonCreditBatchNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarbonCreditBatchNFTCallerSession struct {
	Contract *CarbonCreditBatchNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CarbonCreditBatchNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarbonCreditBatchNFTTransactorSession struct {
	Contract     *CarbonCreditBatchNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CarbonCreditBatchNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarbonCreditBatchNFTRaw struct {
	Contract *CarbonCreditBatchNFT // Generic contract binding to access the raw methods on
}

// CarbonCreditBatchNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarbonCreditBatchNFTCallerRaw struct {
	Contract *CarbonCreditBatchNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CarbonCreditBatchNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarbonCreditBatchNFTTransactorRaw struct {
	Contract *CarbonCreditBatchNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCarbonCreditBatchNFT creates a new instance of CarbonCreditBatchNFT, bound to a specific deployed contract.
func NewCarbonCreditBatchNFT(address common.Address, backend bind.ContractBackend) (*CarbonCreditBatchNFT, error) {
	contract, err := bindCarbonCreditBatchNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFT{CarbonCreditBatchNFTCaller: CarbonCreditBatchNFTCaller{contract: contract}, CarbonCreditBatchNFTTransactor: CarbonCreditBatchNFTTransactor{contract: contract}, CarbonCreditBatchNFTFilterer: CarbonCreditBatchNFTFilterer{contract: contract}}, nil
}

// NewCarbonCreditBatchNFTCaller creates a new read-only instance of CarbonCreditBatchNFT, bound to a specific deployed contract.
func NewCarbonCreditBatchNFTCaller(address common.Address, caller bind.ContractCaller) (*CarbonCreditBatchNFTCaller, error) {
	contract, err := bindCarbonCreditBatchNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTCaller{contract: contract}, nil
}

// NewCarbonCreditBatchNFTTransactor creates a new write-only instance of CarbonCreditBatchNFT, bound to a specific deployed contract.
func NewCarbonCreditBatchNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CarbonCreditBatchNFTTransactor, error) {
	contract, err := bindCarbonCreditBatchNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTTransactor{contract: contract}, nil
}

// NewCarbonCreditBatchNFTFilterer creates a new log filterer instance of CarbonCreditBatchNFT, bound to a specific deployed contract.
func NewCarbonCreditBatchNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CarbonCreditBatchNFTFilterer, error) {
	contract, err := bindCarbonCreditBatchNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTFilterer{contract: contract}, nil
}

// bindCarbonCreditBatchNFT binds a generic wrapper to an already deployed contract.
func bindCarbonCreditBatchNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CarbonCreditBatchNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarbonCreditBatchNFT.Contract.CarbonCreditBatchNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.CarbonCreditBatchNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.CarbonCreditBatchNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarbonCreditBatchNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTIMPLEMENTATION is a free data retrieval call binding the contract method 0x290ab984.
//
// Solidity: function ACCOUNT_IMPLEMENTATION() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) ACCOUNTIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "ACCOUNT_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTIMPLEMENTATION is a free data retrieval call binding the contract method 0x290ab984.
//
// Solidity: function ACCOUNT_IMPLEMENTATION() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) ACCOUNTIMPLEMENTATION() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.ACCOUNTIMPLEMENTATION(&_CarbonCreditBatchNFT.CallOpts)
}

// ACCOUNTIMPLEMENTATION is a free data retrieval call binding the contract method 0x290ab984.
//
// Solidity: function ACCOUNT_IMPLEMENTATION() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) ACCOUNTIMPLEMENTATION() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.ACCOUNTIMPLEMENTATION(&_CarbonCreditBatchNFT.CallOpts)
}

// ERC6551REGISTRY is a free data retrieval call binding the contract method 0xedc16084.
//
// Solidity: function ERC6551_REGISTRY() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) ERC6551REGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "ERC6551_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC6551REGISTRY is a free data retrieval call binding the contract method 0xedc16084.
//
// Solidity: function ERC6551_REGISTRY() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) ERC6551REGISTRY() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.ERC6551REGISTRY(&_CarbonCreditBatchNFT.CallOpts)
}

// ERC6551REGISTRY is a free data retrieval call binding the contract method 0xedc16084.
//
// Solidity: function ERC6551_REGISTRY() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) ERC6551REGISTRY() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.ERC6551REGISTRY(&_CarbonCreditBatchNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CarbonCreditBatchNFT.Contract.BalanceOf(&_CarbonCreditBatchNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CarbonCreditBatchNFT.Contract.BalanceOf(&_CarbonCreditBatchNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.GetApproved(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.GetApproved(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 tokenId) view returns((string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,uint256,string,string,string,string,string,string,string))
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) GetBatch(opts *bind.CallOpts, tokenId *big.Int) (CarbonCreditBatchNFTBatchData, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "getBatch", tokenId)

	if err != nil {
		return *new(CarbonCreditBatchNFTBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(CarbonCreditBatchNFTBatchData)).(*CarbonCreditBatchNFTBatchData)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 tokenId) view returns((string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,uint256,string,string,string,string,string,string,string))
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) GetBatch(tokenId *big.Int) (CarbonCreditBatchNFTBatchData, error) {
	return _CarbonCreditBatchNFT.Contract.GetBatch(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 tokenId) view returns((string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,uint256,string,string,string,string,string,string,string))
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) GetBatch(tokenId *big.Int) (CarbonCreditBatchNFTBatchData, error) {
	return _CarbonCreditBatchNFT.Contract.GetBatch(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// GetTBA is a free data retrieval call binding the contract method 0x15d9756e.
//
// Solidity: function getTBA(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) GetTBA(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "getTBA", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTBA is a free data retrieval call binding the contract method 0x15d9756e.
//
// Solidity: function getTBA(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) GetTBA(tokenId *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.GetTBA(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// GetTBA is a free data retrieval call binding the contract method 0x15d9756e.
//
// Solidity: function getTBA(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) GetTBA(tokenId *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.GetTBA(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CarbonCreditBatchNFT.Contract.IsApprovedForAll(&_CarbonCreditBatchNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CarbonCreditBatchNFT.Contract.IsApprovedForAll(&_CarbonCreditBatchNFT.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) Minter() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.Minter(&_CarbonCreditBatchNFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) Minter() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.Minter(&_CarbonCreditBatchNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) Name() (string, error) {
	return _CarbonCreditBatchNFT.Contract.Name(&_CarbonCreditBatchNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) Name() (string, error) {
	return _CarbonCreditBatchNFT.Contract.Name(&_CarbonCreditBatchNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) NextId() (*big.Int, error) {
	return _CarbonCreditBatchNFT.Contract.NextId(&_CarbonCreditBatchNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) NextId() (*big.Int, error) {
	return _CarbonCreditBatchNFT.Contract.NextId(&_CarbonCreditBatchNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) Owner() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.Owner(&_CarbonCreditBatchNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) Owner() (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.Owner(&_CarbonCreditBatchNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.OwnerOf(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.OwnerOf(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CarbonCreditBatchNFT.Contract.SupportsInterface(&_CarbonCreditBatchNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CarbonCreditBatchNFT.Contract.SupportsInterface(&_CarbonCreditBatchNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) Symbol() (string, error) {
	return _CarbonCreditBatchNFT.Contract.Symbol(&_CarbonCreditBatchNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) Symbol() (string, error) {
	return _CarbonCreditBatchNFT.Contract.Symbol(&_CarbonCreditBatchNFT.CallOpts)
}

// TokenBoundAccount is a free data retrieval call binding the contract method 0x0be76ed6.
//
// Solidity: function tokenBoundAccount(uint256 ) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) TokenBoundAccount(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "tokenBoundAccount", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenBoundAccount is a free data retrieval call binding the contract method 0x0be76ed6.
//
// Solidity: function tokenBoundAccount(uint256 ) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) TokenBoundAccount(arg0 *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.TokenBoundAccount(&_CarbonCreditBatchNFT.CallOpts, arg0)
}

// TokenBoundAccount is a free data retrieval call binding the contract method 0x0be76ed6.
//
// Solidity: function tokenBoundAccount(uint256 ) view returns(address)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) TokenBoundAccount(arg0 *big.Int) (common.Address, error) {
	return _CarbonCreditBatchNFT.Contract.TokenBoundAccount(&_CarbonCreditBatchNFT.CallOpts, arg0)
}

// TokenIdByBatchId is a free data retrieval call binding the contract method 0x790a67e0.
//
// Solidity: function tokenIdByBatchId(string ) view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) TokenIdByBatchId(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "tokenIdByBatchId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdByBatchId is a free data retrieval call binding the contract method 0x790a67e0.
//
// Solidity: function tokenIdByBatchId(string ) view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) TokenIdByBatchId(arg0 string) (*big.Int, error) {
	return _CarbonCreditBatchNFT.Contract.TokenIdByBatchId(&_CarbonCreditBatchNFT.CallOpts, arg0)
}

// TokenIdByBatchId is a free data retrieval call binding the contract method 0x790a67e0.
//
// Solidity: function tokenIdByBatchId(string ) view returns(uint256)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) TokenIdByBatchId(arg0 string) (*big.Int, error) {
	return _CarbonCreditBatchNFT.Contract.TokenIdByBatchId(&_CarbonCreditBatchNFT.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CarbonCreditBatchNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CarbonCreditBatchNFT.Contract.TokenURI(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CarbonCreditBatchNFT.Contract.TokenURI(&_CarbonCreditBatchNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.Approve(&_CarbonCreditBatchNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.Approve(&_CarbonCreditBatchNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x348b2425.
//
// Solidity: function mint(address to, (string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,uint256,string,string,string,string,string,string,string) data) returns(uint256 tokenId, address tba)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, data CarbonCreditBatchNFTBatchData) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "mint", to, data)
}

// Mint is a paid mutator transaction binding the contract method 0x348b2425.
//
// Solidity: function mint(address to, (string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,uint256,string,string,string,string,string,string,string) data) returns(uint256 tokenId, address tba)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) Mint(to common.Address, data CarbonCreditBatchNFTBatchData) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.Mint(&_CarbonCreditBatchNFT.TransactOpts, to, data)
}

// Mint is a paid mutator transaction binding the contract method 0x348b2425.
//
// Solidity: function mint(address to, (string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,uint256,string,string,string,string,string,string,string) data) returns(uint256 tokenId, address tba)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) Mint(to common.Address, data CarbonCreditBatchNFTBatchData) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.Mint(&_CarbonCreditBatchNFT.TransactOpts, to, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.RenounceOwnership(&_CarbonCreditBatchNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.RenounceOwnership(&_CarbonCreditBatchNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SafeTransferFrom(&_CarbonCreditBatchNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SafeTransferFrom(&_CarbonCreditBatchNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SafeTransferFrom0(&_CarbonCreditBatchNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SafeTransferFrom0(&_CarbonCreditBatchNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SetApprovalForAll(&_CarbonCreditBatchNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SetApprovalForAll(&_CarbonCreditBatchNFT.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SetMinter(&_CarbonCreditBatchNFT.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.SetMinter(&_CarbonCreditBatchNFT.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.TransferFrom(&_CarbonCreditBatchNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.TransferFrom(&_CarbonCreditBatchNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.TransferOwnership(&_CarbonCreditBatchNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CarbonCreditBatchNFT.Contract.TransferOwnership(&_CarbonCreditBatchNFT.TransactOpts, newOwner)
}

// CarbonCreditBatchNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTApprovalIterator struct {
	Event *CarbonCreditBatchNFTApproval // Event containing the contract specifics and raw log

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
func (it *CarbonCreditBatchNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditBatchNFTApproval)
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
		it.Event = new(CarbonCreditBatchNFTApproval)
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
func (it *CarbonCreditBatchNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditBatchNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditBatchNFTApproval represents a Approval event raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CarbonCreditBatchNFTApprovalIterator, error) {

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

	logs, sub, err := _CarbonCreditBatchNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTApprovalIterator{contract: _CarbonCreditBatchNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CarbonCreditBatchNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CarbonCreditBatchNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditBatchNFTApproval)
				if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) ParseApproval(log types.Log) (*CarbonCreditBatchNFTApproval, error) {
	event := new(CarbonCreditBatchNFTApproval)
	if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditBatchNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTApprovalForAllIterator struct {
	Event *CarbonCreditBatchNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CarbonCreditBatchNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditBatchNFTApprovalForAll)
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
		it.Event = new(CarbonCreditBatchNFTApprovalForAll)
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
func (it *CarbonCreditBatchNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditBatchNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditBatchNFTApprovalForAll represents a ApprovalForAll event raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CarbonCreditBatchNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTApprovalForAllIterator{contract: _CarbonCreditBatchNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CarbonCreditBatchNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditBatchNFTApprovalForAll)
				if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) ParseApprovalForAll(log types.Log) (*CarbonCreditBatchNFTApprovalForAll, error) {
	event := new(CarbonCreditBatchNFTApprovalForAll)
	if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditBatchNFTBatchMintedIterator is returned from FilterBatchMinted and is used to iterate over the raw logs and unpacked data for BatchMinted events raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTBatchMintedIterator struct {
	Event *CarbonCreditBatchNFTBatchMinted // Event containing the contract specifics and raw log

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
func (it *CarbonCreditBatchNFTBatchMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditBatchNFTBatchMinted)
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
		it.Event = new(CarbonCreditBatchNFTBatchMinted)
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
func (it *CarbonCreditBatchNFTBatchMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditBatchNFTBatchMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditBatchNFTBatchMinted represents a BatchMinted event raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTBatchMinted struct {
	TokenId             *big.Int
	To                  common.Address
	ActivityNftId       *big.Int
	CarbonCreditBatchId string
	UnitType            string
	Amount              *big.Int
	Tba                 common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBatchMinted is a free log retrieval operation binding the contract event 0x9a02b37dc081a82ad5d7e0ae0cfc87757774288e78889fb9f2899006c6c1ed3f.
//
// Solidity: event BatchMinted(uint256 indexed tokenId, address indexed to, uint256 indexed activity_nft_id, string carbon_credit_batch_id, string unit_type, uint256 amount, address tba)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) FilterBatchMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address, activity_nft_id []*big.Int) (*CarbonCreditBatchNFTBatchMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var activity_nft_idRule []interface{}
	for _, activity_nft_idItem := range activity_nft_id {
		activity_nft_idRule = append(activity_nft_idRule, activity_nft_idItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.FilterLogs(opts, "BatchMinted", tokenIdRule, toRule, activity_nft_idRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTBatchMintedIterator{contract: _CarbonCreditBatchNFT.contract, event: "BatchMinted", logs: logs, sub: sub}, nil
}

// WatchBatchMinted is a free log subscription operation binding the contract event 0x9a02b37dc081a82ad5d7e0ae0cfc87757774288e78889fb9f2899006c6c1ed3f.
//
// Solidity: event BatchMinted(uint256 indexed tokenId, address indexed to, uint256 indexed activity_nft_id, string carbon_credit_batch_id, string unit_type, uint256 amount, address tba)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) WatchBatchMinted(opts *bind.WatchOpts, sink chan<- *CarbonCreditBatchNFTBatchMinted, tokenId []*big.Int, to []common.Address, activity_nft_id []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var activity_nft_idRule []interface{}
	for _, activity_nft_idItem := range activity_nft_id {
		activity_nft_idRule = append(activity_nft_idRule, activity_nft_idItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.WatchLogs(opts, "BatchMinted", tokenIdRule, toRule, activity_nft_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditBatchNFTBatchMinted)
				if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "BatchMinted", log); err != nil {
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

// ParseBatchMinted is a log parse operation binding the contract event 0x9a02b37dc081a82ad5d7e0ae0cfc87757774288e78889fb9f2899006c6c1ed3f.
//
// Solidity: event BatchMinted(uint256 indexed tokenId, address indexed to, uint256 indexed activity_nft_id, string carbon_credit_batch_id, string unit_type, uint256 amount, address tba)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) ParseBatchMinted(log types.Log) (*CarbonCreditBatchNFTBatchMinted, error) {
	event := new(CarbonCreditBatchNFTBatchMinted)
	if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "BatchMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditBatchNFTMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTMinterUpdatedIterator struct {
	Event *CarbonCreditBatchNFTMinterUpdated // Event containing the contract specifics and raw log

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
func (it *CarbonCreditBatchNFTMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditBatchNFTMinterUpdated)
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
		it.Event = new(CarbonCreditBatchNFTMinterUpdated)
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
func (it *CarbonCreditBatchNFTMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditBatchNFTMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditBatchNFTMinterUpdated represents a MinterUpdated event raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) FilterMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*CarbonCreditBatchNFTMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.FilterLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTMinterUpdatedIterator{contract: _CarbonCreditBatchNFT.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *CarbonCreditBatchNFTMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.WatchLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditBatchNFTMinterUpdated)
				if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
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
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) ParseMinterUpdated(log types.Log) (*CarbonCreditBatchNFTMinterUpdated, error) {
	event := new(CarbonCreditBatchNFTMinterUpdated)
	if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditBatchNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTOwnershipTransferredIterator struct {
	Event *CarbonCreditBatchNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CarbonCreditBatchNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditBatchNFTOwnershipTransferred)
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
		it.Event = new(CarbonCreditBatchNFTOwnershipTransferred)
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
func (it *CarbonCreditBatchNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditBatchNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditBatchNFTOwnershipTransferred represents a OwnershipTransferred event raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CarbonCreditBatchNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTOwnershipTransferredIterator{contract: _CarbonCreditBatchNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CarbonCreditBatchNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CarbonCreditBatchNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditBatchNFTOwnershipTransferred)
				if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) ParseOwnershipTransferred(log types.Log) (*CarbonCreditBatchNFTOwnershipTransferred, error) {
	event := new(CarbonCreditBatchNFTOwnershipTransferred)
	if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditBatchNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTTransferIterator struct {
	Event *CarbonCreditBatchNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CarbonCreditBatchNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditBatchNFTTransfer)
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
		it.Event = new(CarbonCreditBatchNFTTransfer)
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
func (it *CarbonCreditBatchNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditBatchNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditBatchNFTTransfer represents a Transfer event raised by the CarbonCreditBatchNFT contract.
type CarbonCreditBatchNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CarbonCreditBatchNFTTransferIterator, error) {

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

	logs, sub, err := _CarbonCreditBatchNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditBatchNFTTransferIterator{contract: _CarbonCreditBatchNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CarbonCreditBatchNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CarbonCreditBatchNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditBatchNFTTransfer)
				if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CarbonCreditBatchNFT *CarbonCreditBatchNFTFilterer) ParseTransfer(log types.Log) (*CarbonCreditBatchNFTTransfer, error) {
	event := new(CarbonCreditBatchNFTTransfer)
	if err := _CarbonCreditBatchNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
