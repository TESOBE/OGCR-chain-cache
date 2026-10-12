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

// CertificationNFTCertificationData is an auto generated low-level Go binding around an user-defined struct.
type CertificationNFTCertificationData struct {
	CertificateOfComplianceId                     string
	CertificateOfComplianceUri                    string
	CertificateOfComplianceHash                   string
	ActivityId                                    string
	ActivityNftId                                 *big.Int
	ActivityUri                                   string
	ActivityHash                                  string
	ActivityName                                  string
	CertificationSchemeName                       string
	CertificationSchemeUri                        string
	CertificationSchemeHash                       string
	CertificationBodyName                         string
	CertificationBodyUri                          string
	CertificationBodyHash                         string
	CountryId                                     string
	IssueDate                                     string
	ExpiryDate                                    string
	AuditReportUri                                string
	AuditReportHash                               string
	CarbonRemovalsUnderBaseline                   string
	SoilEmissionsUnderBaseline                    string
	PermanentNetCarbonRemovalBenefit              string
	CarbonFarmingTemporaryNetCarbonRemovalBenefit string
	CarbonFarmingNetSoilEmissionReductionBenefit  string
	CarbonStorageTemporaryNetCarbonRemovalBenefit string
	UnitTypes                                     string
	CertificationStatus                           string
}

// CertificationNFTMetaData contains all meta data concerning the CertificationNFT contract.
var CertificationNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertification\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCertificationNFT.CertificationData\",\"components\":[{\"name\":\"certificate_of_compliance_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_nft_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activity_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"country_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"issue_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expiry_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"audit_report_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"audit_report_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_removals_under_baseline\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"soil_emissions_under_baseline\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"permanent_net_carbon_removal_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_farming_temporary_net_carbon_removal_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_farming_net_soil_emission_reduction_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_storage_temporary_net_carbon_removal_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"unit_types\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_status\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structCertificationNFT.CertificationData\",\"components\":[{\"name\":\"certificate_of_compliance_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certificate_of_compliance_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_nft_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activity_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activity_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_scheme_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_body_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"country_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"issue_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expiry_date\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"audit_report_uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"audit_report_hash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_removals_under_baseline\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"soil_emissions_under_baseline\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"permanent_net_carbon_removal_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_farming_temporary_net_carbon_removal_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_farming_net_soil_emission_reduction_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"carbon_storage_temporary_net_carbon_removal_benefit\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"unit_types\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"certification_status\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenIdByComplianceId\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CertificationMinted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activity_nft_id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"certificate_of_compliance_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterUpdated\",\"inputs\":[{\"name\":\"oldMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// CertificationNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificationNFTMetaData.ABI instead.
var CertificationNFTABI = CertificationNFTMetaData.ABI

// CertificationNFT is an auto generated Go binding around an Ethereum contract.
type CertificationNFT struct {
	CertificationNFTCaller     // Read-only binding to the contract
	CertificationNFTTransactor // Write-only binding to the contract
	CertificationNFTFilterer   // Log filterer for contract events
}

// CertificationNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificationNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificationNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificationNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificationNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificationNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificationNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificationNFTSession struct {
	Contract     *CertificationNFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertificationNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificationNFTCallerSession struct {
	Contract *CertificationNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CertificationNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificationNFTTransactorSession struct {
	Contract     *CertificationNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CertificationNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificationNFTRaw struct {
	Contract *CertificationNFT // Generic contract binding to access the raw methods on
}

// CertificationNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificationNFTCallerRaw struct {
	Contract *CertificationNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CertificationNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificationNFTTransactorRaw struct {
	Contract *CertificationNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificationNFT creates a new instance of CertificationNFT, bound to a specific deployed contract.
func NewCertificationNFT(address common.Address, backend bind.ContractBackend) (*CertificationNFT, error) {
	contract, err := bindCertificationNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificationNFT{CertificationNFTCaller: CertificationNFTCaller{contract: contract}, CertificationNFTTransactor: CertificationNFTTransactor{contract: contract}, CertificationNFTFilterer: CertificationNFTFilterer{contract: contract}}, nil
}

// NewCertificationNFTCaller creates a new read-only instance of CertificationNFT, bound to a specific deployed contract.
func NewCertificationNFTCaller(address common.Address, caller bind.ContractCaller) (*CertificationNFTCaller, error) {
	contract, err := bindCertificationNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTCaller{contract: contract}, nil
}

// NewCertificationNFTTransactor creates a new write-only instance of CertificationNFT, bound to a specific deployed contract.
func NewCertificationNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificationNFTTransactor, error) {
	contract, err := bindCertificationNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTTransactor{contract: contract}, nil
}

// NewCertificationNFTFilterer creates a new log filterer instance of CertificationNFT, bound to a specific deployed contract.
func NewCertificationNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificationNFTFilterer, error) {
	contract, err := bindCertificationNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTFilterer{contract: contract}, nil
}

// bindCertificationNFT binds a generic wrapper to an already deployed contract.
func bindCertificationNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertificationNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificationNFT *CertificationNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificationNFT.Contract.CertificationNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificationNFT *CertificationNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificationNFT.Contract.CertificationNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificationNFT *CertificationNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificationNFT.Contract.CertificationNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificationNFT *CertificationNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificationNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificationNFT *CertificationNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificationNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificationNFT *CertificationNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificationNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CertificationNFT *CertificationNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CertificationNFT *CertificationNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CertificationNFT.Contract.BalanceOf(&_CertificationNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CertificationNFT *CertificationNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CertificationNFT.Contract.BalanceOf(&_CertificationNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CertificationNFT *CertificationNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CertificationNFT *CertificationNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CertificationNFT.Contract.GetApproved(&_CertificationNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CertificationNFT *CertificationNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CertificationNFT.Contract.GetApproved(&_CertificationNFT.CallOpts, tokenId)
}

// GetCertification is a free data retrieval call binding the contract method 0xec981eac.
//
// Solidity: function getCertification(uint256 tokenId) view returns((string,string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string))
func (_CertificationNFT *CertificationNFTCaller) GetCertification(opts *bind.CallOpts, tokenId *big.Int) (CertificationNFTCertificationData, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "getCertification", tokenId)

	if err != nil {
		return *new(CertificationNFTCertificationData), err
	}

	out0 := *abi.ConvertType(out[0], new(CertificationNFTCertificationData)).(*CertificationNFTCertificationData)

	return out0, err

}

// GetCertification is a free data retrieval call binding the contract method 0xec981eac.
//
// Solidity: function getCertification(uint256 tokenId) view returns((string,string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string))
func (_CertificationNFT *CertificationNFTSession) GetCertification(tokenId *big.Int) (CertificationNFTCertificationData, error) {
	return _CertificationNFT.Contract.GetCertification(&_CertificationNFT.CallOpts, tokenId)
}

// GetCertification is a free data retrieval call binding the contract method 0xec981eac.
//
// Solidity: function getCertification(uint256 tokenId) view returns((string,string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string))
func (_CertificationNFT *CertificationNFTCallerSession) GetCertification(tokenId *big.Int) (CertificationNFTCertificationData, error) {
	return _CertificationNFT.Contract.GetCertification(&_CertificationNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CertificationNFT *CertificationNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CertificationNFT *CertificationNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CertificationNFT.Contract.IsApprovedForAll(&_CertificationNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CertificationNFT *CertificationNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CertificationNFT.Contract.IsApprovedForAll(&_CertificationNFT.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CertificationNFT *CertificationNFTCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CertificationNFT *CertificationNFTSession) Minter() (common.Address, error) {
	return _CertificationNFT.Contract.Minter(&_CertificationNFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CertificationNFT *CertificationNFTCallerSession) Minter() (common.Address, error) {
	return _CertificationNFT.Contract.Minter(&_CertificationNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CertificationNFT *CertificationNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CertificationNFT *CertificationNFTSession) Name() (string, error) {
	return _CertificationNFT.Contract.Name(&_CertificationNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CertificationNFT *CertificationNFTCallerSession) Name() (string, error) {
	return _CertificationNFT.Contract.Name(&_CertificationNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_CertificationNFT *CertificationNFTCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_CertificationNFT *CertificationNFTSession) NextId() (*big.Int, error) {
	return _CertificationNFT.Contract.NextId(&_CertificationNFT.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_CertificationNFT *CertificationNFTCallerSession) NextId() (*big.Int, error) {
	return _CertificationNFT.Contract.NextId(&_CertificationNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CertificationNFT *CertificationNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CertificationNFT *CertificationNFTSession) Owner() (common.Address, error) {
	return _CertificationNFT.Contract.Owner(&_CertificationNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CertificationNFT *CertificationNFTCallerSession) Owner() (common.Address, error) {
	return _CertificationNFT.Contract.Owner(&_CertificationNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CertificationNFT *CertificationNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CertificationNFT *CertificationNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CertificationNFT.Contract.OwnerOf(&_CertificationNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CertificationNFT *CertificationNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CertificationNFT.Contract.OwnerOf(&_CertificationNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CertificationNFT *CertificationNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CertificationNFT *CertificationNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CertificationNFT.Contract.SupportsInterface(&_CertificationNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CertificationNFT *CertificationNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CertificationNFT.Contract.SupportsInterface(&_CertificationNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CertificationNFT *CertificationNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CertificationNFT *CertificationNFTSession) Symbol() (string, error) {
	return _CertificationNFT.Contract.Symbol(&_CertificationNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CertificationNFT *CertificationNFTCallerSession) Symbol() (string, error) {
	return _CertificationNFT.Contract.Symbol(&_CertificationNFT.CallOpts)
}

// TokenIdByComplianceId is a free data retrieval call binding the contract method 0x62c2ea10.
//
// Solidity: function tokenIdByComplianceId(string ) view returns(uint256)
func (_CertificationNFT *CertificationNFTCaller) TokenIdByComplianceId(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "tokenIdByComplianceId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdByComplianceId is a free data retrieval call binding the contract method 0x62c2ea10.
//
// Solidity: function tokenIdByComplianceId(string ) view returns(uint256)
func (_CertificationNFT *CertificationNFTSession) TokenIdByComplianceId(arg0 string) (*big.Int, error) {
	return _CertificationNFT.Contract.TokenIdByComplianceId(&_CertificationNFT.CallOpts, arg0)
}

// TokenIdByComplianceId is a free data retrieval call binding the contract method 0x62c2ea10.
//
// Solidity: function tokenIdByComplianceId(string ) view returns(uint256)
func (_CertificationNFT *CertificationNFTCallerSession) TokenIdByComplianceId(arg0 string) (*big.Int, error) {
	return _CertificationNFT.Contract.TokenIdByComplianceId(&_CertificationNFT.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CertificationNFT *CertificationNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CertificationNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CertificationNFT *CertificationNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CertificationNFT.Contract.TokenURI(&_CertificationNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CertificationNFT *CertificationNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CertificationNFT.Contract.TokenURI(&_CertificationNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.Contract.Approve(&_CertificationNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.Contract.Approve(&_CertificationNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xcd7c1918.
//
// Solidity: function mint(address to, (string,string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string) data) returns(uint256 tokenId)
func (_CertificationNFT *CertificationNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, data CertificationNFTCertificationData) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "mint", to, data)
}

// Mint is a paid mutator transaction binding the contract method 0xcd7c1918.
//
// Solidity: function mint(address to, (string,string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string) data) returns(uint256 tokenId)
func (_CertificationNFT *CertificationNFTSession) Mint(to common.Address, data CertificationNFTCertificationData) (*types.Transaction, error) {
	return _CertificationNFT.Contract.Mint(&_CertificationNFT.TransactOpts, to, data)
}

// Mint is a paid mutator transaction binding the contract method 0xcd7c1918.
//
// Solidity: function mint(address to, (string,string,string,string,uint256,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string,string) data) returns(uint256 tokenId)
func (_CertificationNFT *CertificationNFTTransactorSession) Mint(to common.Address, data CertificationNFTCertificationData) (*types.Transaction, error) {
	return _CertificationNFT.Contract.Mint(&_CertificationNFT.TransactOpts, to, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertificationNFT *CertificationNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertificationNFT *CertificationNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _CertificationNFT.Contract.RenounceOwnership(&_CertificationNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertificationNFT *CertificationNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CertificationNFT.Contract.RenounceOwnership(&_CertificationNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SafeTransferFrom(&_CertificationNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SafeTransferFrom(&_CertificationNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CertificationNFT *CertificationNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CertificationNFT *CertificationNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SafeTransferFrom0(&_CertificationNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SafeTransferFrom0(&_CertificationNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CertificationNFT *CertificationNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CertificationNFT *CertificationNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SetApprovalForAll(&_CertificationNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SetApprovalForAll(&_CertificationNFT.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CertificationNFT *CertificationNFTTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CertificationNFT *CertificationNFTSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SetMinter(&_CertificationNFT.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CertificationNFT.Contract.SetMinter(&_CertificationNFT.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.Contract.TransferFrom(&_CertificationNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificationNFT.Contract.TransferFrom(&_CertificationNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CertificationNFT *CertificationNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CertificationNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CertificationNFT *CertificationNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CertificationNFT.Contract.TransferOwnership(&_CertificationNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CertificationNFT *CertificationNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CertificationNFT.Contract.TransferOwnership(&_CertificationNFT.TransactOpts, newOwner)
}

// CertificationNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CertificationNFT contract.
type CertificationNFTApprovalIterator struct {
	Event *CertificationNFTApproval // Event containing the contract specifics and raw log

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
func (it *CertificationNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificationNFTApproval)
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
		it.Event = new(CertificationNFTApproval)
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
func (it *CertificationNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificationNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificationNFTApproval represents a Approval event raised by the CertificationNFT contract.
type CertificationNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CertificationNFT *CertificationNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CertificationNFTApprovalIterator, error) {

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

	logs, sub, err := _CertificationNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTApprovalIterator{contract: _CertificationNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CertificationNFT *CertificationNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CertificationNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CertificationNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificationNFTApproval)
				if err := _CertificationNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CertificationNFT *CertificationNFTFilterer) ParseApproval(log types.Log) (*CertificationNFTApproval, error) {
	event := new(CertificationNFTApproval)
	if err := _CertificationNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificationNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CertificationNFT contract.
type CertificationNFTApprovalForAllIterator struct {
	Event *CertificationNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CertificationNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificationNFTApprovalForAll)
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
		it.Event = new(CertificationNFTApprovalForAll)
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
func (it *CertificationNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificationNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificationNFTApprovalForAll represents a ApprovalForAll event raised by the CertificationNFT contract.
type CertificationNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CertificationNFT *CertificationNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CertificationNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CertificationNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTApprovalForAllIterator{contract: _CertificationNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CertificationNFT *CertificationNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CertificationNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CertificationNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificationNFTApprovalForAll)
				if err := _CertificationNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_CertificationNFT *CertificationNFTFilterer) ParseApprovalForAll(log types.Log) (*CertificationNFTApprovalForAll, error) {
	event := new(CertificationNFTApprovalForAll)
	if err := _CertificationNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificationNFTCertificationMintedIterator is returned from FilterCertificationMinted and is used to iterate over the raw logs and unpacked data for CertificationMinted events raised by the CertificationNFT contract.
type CertificationNFTCertificationMintedIterator struct {
	Event *CertificationNFTCertificationMinted // Event containing the contract specifics and raw log

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
func (it *CertificationNFTCertificationMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificationNFTCertificationMinted)
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
		it.Event = new(CertificationNFTCertificationMinted)
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
func (it *CertificationNFTCertificationMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificationNFTCertificationMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificationNFTCertificationMinted represents a CertificationMinted event raised by the CertificationNFT contract.
type CertificationNFTCertificationMinted struct {
	TokenId                   *big.Int
	To                        common.Address
	ActivityNftId             *big.Int
	CertificateOfComplianceId string
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterCertificationMinted is a free log retrieval operation binding the contract event 0x7f56d08b4d05f9b7fa105af1225f8682744308792fb4d52a8055d9c377e23179.
//
// Solidity: event CertificationMinted(uint256 indexed tokenId, address indexed to, uint256 indexed activity_nft_id, string certificate_of_compliance_id)
func (_CertificationNFT *CertificationNFTFilterer) FilterCertificationMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address, activity_nft_id []*big.Int) (*CertificationNFTCertificationMintedIterator, error) {

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

	logs, sub, err := _CertificationNFT.contract.FilterLogs(opts, "CertificationMinted", tokenIdRule, toRule, activity_nft_idRule)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTCertificationMintedIterator{contract: _CertificationNFT.contract, event: "CertificationMinted", logs: logs, sub: sub}, nil
}

// WatchCertificationMinted is a free log subscription operation binding the contract event 0x7f56d08b4d05f9b7fa105af1225f8682744308792fb4d52a8055d9c377e23179.
//
// Solidity: event CertificationMinted(uint256 indexed tokenId, address indexed to, uint256 indexed activity_nft_id, string certificate_of_compliance_id)
func (_CertificationNFT *CertificationNFTFilterer) WatchCertificationMinted(opts *bind.WatchOpts, sink chan<- *CertificationNFTCertificationMinted, tokenId []*big.Int, to []common.Address, activity_nft_id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CertificationNFT.contract.WatchLogs(opts, "CertificationMinted", tokenIdRule, toRule, activity_nft_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificationNFTCertificationMinted)
				if err := _CertificationNFT.contract.UnpackLog(event, "CertificationMinted", log); err != nil {
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

// ParseCertificationMinted is a log parse operation binding the contract event 0x7f56d08b4d05f9b7fa105af1225f8682744308792fb4d52a8055d9c377e23179.
//
// Solidity: event CertificationMinted(uint256 indexed tokenId, address indexed to, uint256 indexed activity_nft_id, string certificate_of_compliance_id)
func (_CertificationNFT *CertificationNFTFilterer) ParseCertificationMinted(log types.Log) (*CertificationNFTCertificationMinted, error) {
	event := new(CertificationNFTCertificationMinted)
	if err := _CertificationNFT.contract.UnpackLog(event, "CertificationMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificationNFTMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the CertificationNFT contract.
type CertificationNFTMinterUpdatedIterator struct {
	Event *CertificationNFTMinterUpdated // Event containing the contract specifics and raw log

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
func (it *CertificationNFTMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificationNFTMinterUpdated)
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
		it.Event = new(CertificationNFTMinterUpdated)
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
func (it *CertificationNFTMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificationNFTMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificationNFTMinterUpdated represents a MinterUpdated event raised by the CertificationNFT contract.
type CertificationNFTMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CertificationNFT *CertificationNFTFilterer) FilterMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*CertificationNFTMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CertificationNFT.contract.FilterLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTMinterUpdatedIterator{contract: _CertificationNFT.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CertificationNFT *CertificationNFTFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *CertificationNFTMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CertificationNFT.contract.WatchLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificationNFTMinterUpdated)
				if err := _CertificationNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
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
func (_CertificationNFT *CertificationNFTFilterer) ParseMinterUpdated(log types.Log) (*CertificationNFTMinterUpdated, error) {
	event := new(CertificationNFTMinterUpdated)
	if err := _CertificationNFT.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificationNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CertificationNFT contract.
type CertificationNFTOwnershipTransferredIterator struct {
	Event *CertificationNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CertificationNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificationNFTOwnershipTransferred)
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
		it.Event = new(CertificationNFTOwnershipTransferred)
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
func (it *CertificationNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificationNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificationNFTOwnershipTransferred represents a OwnershipTransferred event raised by the CertificationNFT contract.
type CertificationNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CertificationNFT *CertificationNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CertificationNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CertificationNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTOwnershipTransferredIterator{contract: _CertificationNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CertificationNFT *CertificationNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CertificationNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CertificationNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificationNFTOwnershipTransferred)
				if err := _CertificationNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CertificationNFT *CertificationNFTFilterer) ParseOwnershipTransferred(log types.Log) (*CertificationNFTOwnershipTransferred, error) {
	event := new(CertificationNFTOwnershipTransferred)
	if err := _CertificationNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificationNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CertificationNFT contract.
type CertificationNFTTransferIterator struct {
	Event *CertificationNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CertificationNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificationNFTTransfer)
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
		it.Event = new(CertificationNFTTransfer)
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
func (it *CertificationNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificationNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificationNFTTransfer represents a Transfer event raised by the CertificationNFT contract.
type CertificationNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CertificationNFT *CertificationNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CertificationNFTTransferIterator, error) {

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

	logs, sub, err := _CertificationNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CertificationNFTTransferIterator{contract: _CertificationNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CertificationNFT *CertificationNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CertificationNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CertificationNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificationNFTTransfer)
				if err := _CertificationNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CertificationNFT *CertificationNFTFilterer) ParseTransfer(log types.Log) (*CertificationNFTTransfer, error) {
	event := new(CertificationNFTTransfer)
	if err := _CertificationNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
