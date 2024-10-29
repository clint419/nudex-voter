// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// VotingManagerContractMetaData contains all meta data concerning the VotingManagerContract contract.
var VotingManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"accountManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAccountManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addBonusPoints\",\"inputs\":[{\"name\":\"_voters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addParticipant\",\"inputs\":[{\"name\":\"newParticipant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAssetManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chooseNewSubmitter\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmTasks\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delistAsset\",\"inputs\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDepositManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forcedRotationWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_accountManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_assetManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_participantManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nuDexOperations\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nuvoLock\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastSubmissionTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listAsset\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nuDexName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nuDexOperations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINuDexOperations\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nuvoLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINuvoLock\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participantManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIParticipantManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmTask\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAccount\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chain\",\"type\":\"uint8\",\"internalType\":\"enumIAccountManager.Chain\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeParticipant\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardPerPeriod\",\"inputs\":[{\"name\":\"newRewardPerPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDepositInfo\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTaskReceipt\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitWithdrawalInfo\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskCompletionThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPerPeriodVoted\",\"inputs\":[{\"name\":\"newRewardPerPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubmitterChosen\",\"inputs\":[{\"name\":\"newSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubmitterRotationRequested\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"currentSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IncorrectSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RotationWindowNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaskAlreadyCompleted\",\"inputs\":[]}]",
}

// VotingManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingManagerContractMetaData.ABI instead.
var VotingManagerContractABI = VotingManagerContractMetaData.ABI

// VotingManagerContract is an auto generated Go binding around an Ethereum contract.
type VotingManagerContract struct {
	VotingManagerContractCaller     // Read-only binding to the contract
	VotingManagerContractTransactor // Write-only binding to the contract
	VotingManagerContractFilterer   // Log filterer for contract events
}

// VotingManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingManagerContractSession struct {
	Contract     *VotingManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VotingManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingManagerContractCallerSession struct {
	Contract *VotingManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// VotingManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingManagerContractTransactorSession struct {
	Contract     *VotingManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// VotingManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingManagerContractRaw struct {
	Contract *VotingManagerContract // Generic contract binding to access the raw methods on
}

// VotingManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingManagerContractCallerRaw struct {
	Contract *VotingManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// VotingManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingManagerContractTransactorRaw struct {
	Contract *VotingManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingManagerContract creates a new instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContract(address common.Address, backend bind.ContractBackend) (*VotingManagerContract, error) {
	contract, err := bindVotingManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContract{VotingManagerContractCaller: VotingManagerContractCaller{contract: contract}, VotingManagerContractTransactor: VotingManagerContractTransactor{contract: contract}, VotingManagerContractFilterer: VotingManagerContractFilterer{contract: contract}}, nil
}

// NewVotingManagerContractCaller creates a new read-only instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContractCaller(address common.Address, caller bind.ContractCaller) (*VotingManagerContractCaller, error) {
	contract, err := bindVotingManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractCaller{contract: contract}, nil
}

// NewVotingManagerContractTransactor creates a new write-only instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingManagerContractTransactor, error) {
	contract, err := bindVotingManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractTransactor{contract: contract}, nil
}

// NewVotingManagerContractFilterer creates a new log filterer instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingManagerContractFilterer, error) {
	contract, err := bindVotingManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractFilterer{contract: contract}, nil
}

// bindVotingManagerContract binds a generic wrapper to an already deployed contract.
func bindVotingManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingManagerContract *VotingManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingManagerContract.Contract.VotingManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingManagerContract *VotingManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.VotingManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingManagerContract *VotingManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.VotingManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingManagerContract *VotingManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingManagerContract *VotingManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingManagerContract *VotingManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.contract.Transact(opts, method, params...)
}

// AccountManager is a free data retrieval call binding the contract method 0x91c2c469.
//
// Solidity: function accountManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) AccountManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "accountManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountManager is a free data retrieval call binding the contract method 0x91c2c469.
//
// Solidity: function accountManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) AccountManager() (common.Address, error) {
	return _VotingManagerContract.Contract.AccountManager(&_VotingManagerContract.CallOpts)
}

// AccountManager is a free data retrieval call binding the contract method 0x91c2c469.
//
// Solidity: function accountManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) AccountManager() (common.Address, error) {
	return _VotingManagerContract.Contract.AccountManager(&_VotingManagerContract.CallOpts)
}

// AssetManager is a free data retrieval call binding the contract method 0x94217ad1.
//
// Solidity: function assetManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) AssetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "assetManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetManager is a free data retrieval call binding the contract method 0x94217ad1.
//
// Solidity: function assetManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) AssetManager() (common.Address, error) {
	return _VotingManagerContract.Contract.AssetManager(&_VotingManagerContract.CallOpts)
}

// AssetManager is a free data retrieval call binding the contract method 0x94217ad1.
//
// Solidity: function assetManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) AssetManager() (common.Address, error) {
	return _VotingManagerContract.Contract.AssetManager(&_VotingManagerContract.CallOpts)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) DepositManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "depositManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) DepositManager() (common.Address, error) {
	return _VotingManagerContract.Contract.DepositManager(&_VotingManagerContract.CallOpts)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) DepositManager() (common.Address, error) {
	return _VotingManagerContract.Contract.DepositManager(&_VotingManagerContract.CallOpts)
}

// ForcedRotationWindow is a free data retrieval call binding the contract method 0xc858379d.
//
// Solidity: function forcedRotationWindow() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) ForcedRotationWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "forcedRotationWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForcedRotationWindow is a free data retrieval call binding the contract method 0xc858379d.
//
// Solidity: function forcedRotationWindow() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) ForcedRotationWindow() (*big.Int, error) {
	return _VotingManagerContract.Contract.ForcedRotationWindow(&_VotingManagerContract.CallOpts)
}

// ForcedRotationWindow is a free data retrieval call binding the contract method 0xc858379d.
//
// Solidity: function forcedRotationWindow() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) ForcedRotationWindow() (*big.Int, error) {
	return _VotingManagerContract.Contract.ForcedRotationWindow(&_VotingManagerContract.CallOpts)
}

// LastSubmissionTime is a free data retrieval call binding the contract method 0x4f70104e.
//
// Solidity: function lastSubmissionTime() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) LastSubmissionTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "lastSubmissionTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSubmissionTime is a free data retrieval call binding the contract method 0x4f70104e.
//
// Solidity: function lastSubmissionTime() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) LastSubmissionTime() (*big.Int, error) {
	return _VotingManagerContract.Contract.LastSubmissionTime(&_VotingManagerContract.CallOpts)
}

// LastSubmissionTime is a free data retrieval call binding the contract method 0x4f70104e.
//
// Solidity: function lastSubmissionTime() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) LastSubmissionTime() (*big.Int, error) {
	return _VotingManagerContract.Contract.LastSubmissionTime(&_VotingManagerContract.CallOpts)
}

// NextSubmitter is a free data retrieval call binding the contract method 0x249c94e8.
//
// Solidity: function nextSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) NextSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "nextSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextSubmitter is a free data retrieval call binding the contract method 0x249c94e8.
//
// Solidity: function nextSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) NextSubmitter() (common.Address, error) {
	return _VotingManagerContract.Contract.NextSubmitter(&_VotingManagerContract.CallOpts)
}

// NextSubmitter is a free data retrieval call binding the contract method 0x249c94e8.
//
// Solidity: function nextSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) NextSubmitter() (common.Address, error) {
	return _VotingManagerContract.Contract.NextSubmitter(&_VotingManagerContract.CallOpts)
}

// NuDexOperations is a free data retrieval call binding the contract method 0xb02ae352.
//
// Solidity: function nuDexOperations() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) NuDexOperations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "nuDexOperations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NuDexOperations is a free data retrieval call binding the contract method 0xb02ae352.
//
// Solidity: function nuDexOperations() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) NuDexOperations() (common.Address, error) {
	return _VotingManagerContract.Contract.NuDexOperations(&_VotingManagerContract.CallOpts)
}

// NuDexOperations is a free data retrieval call binding the contract method 0xb02ae352.
//
// Solidity: function nuDexOperations() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) NuDexOperations() (common.Address, error) {
	return _VotingManagerContract.Contract.NuDexOperations(&_VotingManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) NuvoLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "nuvoLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) NuvoLock() (common.Address, error) {
	return _VotingManagerContract.Contract.NuvoLock(&_VotingManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) NuvoLock() (common.Address, error) {
	return _VotingManagerContract.Contract.NuvoLock(&_VotingManagerContract.CallOpts)
}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) ParticipantManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "participantManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) ParticipantManager() (common.Address, error) {
	return _VotingManagerContract.Contract.ParticipantManager(&_VotingManagerContract.CallOpts)
}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) ParticipantManager() (common.Address, error) {
	return _VotingManagerContract.Contract.ParticipantManager(&_VotingManagerContract.CallOpts)
}

// TaskCompletionThreshold is a free data retrieval call binding the contract method 0xc6a1f28b.
//
// Solidity: function taskCompletionThreshold() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) TaskCompletionThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "taskCompletionThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCompletionThreshold is a free data retrieval call binding the contract method 0xc6a1f28b.
//
// Solidity: function taskCompletionThreshold() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) TaskCompletionThreshold() (*big.Int, error) {
	return _VotingManagerContract.Contract.TaskCompletionThreshold(&_VotingManagerContract.CallOpts)
}

// TaskCompletionThreshold is a free data retrieval call binding the contract method 0xc6a1f28b.
//
// Solidity: function taskCompletionThreshold() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) TaskCompletionThreshold() (*big.Int, error) {
	return _VotingManagerContract.Contract.TaskCompletionThreshold(&_VotingManagerContract.CallOpts)
}

// AddBonusPoints is a paid mutator transaction binding the contract method 0x57b61412.
//
// Solidity: function addBonusPoints(address[] _voters, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) AddBonusPoints(opts *bind.TransactOpts, _voters []common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "addBonusPoints", _voters, _signature)
}

// AddBonusPoints is a paid mutator transaction binding the contract method 0x57b61412.
//
// Solidity: function addBonusPoints(address[] _voters, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) AddBonusPoints(_voters []common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.AddBonusPoints(&_VotingManagerContract.TransactOpts, _voters, _signature)
}

// AddBonusPoints is a paid mutator transaction binding the contract method 0x57b61412.
//
// Solidity: function addBonusPoints(address[] _voters, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) AddBonusPoints(_voters []common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.AddBonusPoints(&_VotingManagerContract.TransactOpts, _voters, _signature)
}

// AddParticipant is a paid mutator transaction binding the contract method 0x8d240523.
//
// Solidity: function addParticipant(address newParticipant, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) AddParticipant(opts *bind.TransactOpts, newParticipant common.Address, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "addParticipant", newParticipant, signature)
}

// AddParticipant is a paid mutator transaction binding the contract method 0x8d240523.
//
// Solidity: function addParticipant(address newParticipant, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) AddParticipant(newParticipant common.Address, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.AddParticipant(&_VotingManagerContract.TransactOpts, newParticipant, signature)
}

// AddParticipant is a paid mutator transaction binding the contract method 0x8d240523.
//
// Solidity: function addParticipant(address newParticipant, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) AddParticipant(newParticipant common.Address, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.AddParticipant(&_VotingManagerContract.TransactOpts, newParticipant, signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0xfe22e244.
//
// Solidity: function chooseNewSubmitter(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) ChooseNewSubmitter(opts *bind.TransactOpts, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "chooseNewSubmitter", _signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0xfe22e244.
//
// Solidity: function chooseNewSubmitter(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) ChooseNewSubmitter(_signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ChooseNewSubmitter(&_VotingManagerContract.TransactOpts, _signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0xfe22e244.
//
// Solidity: function chooseNewSubmitter(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) ChooseNewSubmitter(_signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ChooseNewSubmitter(&_VotingManagerContract.TransactOpts, _signature)
}

// ConfirmTasks is a paid mutator transaction binding the contract method 0xbc85d659.
//
// Solidity: function confirmTasks(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) ConfirmTasks(opts *bind.TransactOpts, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "confirmTasks", _signature)
}

// ConfirmTasks is a paid mutator transaction binding the contract method 0xbc85d659.
//
// Solidity: function confirmTasks(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) ConfirmTasks(_signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ConfirmTasks(&_VotingManagerContract.TransactOpts, _signature)
}

// ConfirmTasks is a paid mutator transaction binding the contract method 0xbc85d659.
//
// Solidity: function confirmTasks(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) ConfirmTasks(_signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ConfirmTasks(&_VotingManagerContract.TransactOpts, _signature)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe7a36191.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) DelistAsset(opts *bind.TransactOpts, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "delistAsset", assetType, contractAddress, chainId, signature)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe7a36191.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) DelistAsset(assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.DelistAsset(&_VotingManagerContract.TransactOpts, assetType, contractAddress, chainId, signature)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe7a36191.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) DelistAsset(assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.DelistAsset(&_VotingManagerContract.TransactOpts, assetType, contractAddress, chainId, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _accountManager, address _assetManager, address _depositManager, address _participantManager, address _nuDexOperations, address _nuvoLock) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) Initialize(opts *bind.TransactOpts, _accountManager common.Address, _assetManager common.Address, _depositManager common.Address, _participantManager common.Address, _nuDexOperations common.Address, _nuvoLock common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "initialize", _accountManager, _assetManager, _depositManager, _participantManager, _nuDexOperations, _nuvoLock)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _accountManager, address _assetManager, address _depositManager, address _participantManager, address _nuDexOperations, address _nuvoLock) returns()
func (_VotingManagerContract *VotingManagerContractSession) Initialize(_accountManager common.Address, _assetManager common.Address, _depositManager common.Address, _participantManager common.Address, _nuDexOperations common.Address, _nuvoLock common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.Initialize(&_VotingManagerContract.TransactOpts, _accountManager, _assetManager, _depositManager, _participantManager, _nuDexOperations, _nuvoLock)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _accountManager, address _assetManager, address _depositManager, address _participantManager, address _nuDexOperations, address _nuvoLock) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) Initialize(_accountManager common.Address, _assetManager common.Address, _depositManager common.Address, _participantManager common.Address, _nuDexOperations common.Address, _nuvoLock common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.Initialize(&_VotingManagerContract.TransactOpts, _accountManager, _assetManager, _depositManager, _participantManager, _nuDexOperations, _nuvoLock)
}

// ListAsset is a paid mutator transaction binding the contract method 0xd1c0a28a.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) ListAsset(opts *bind.TransactOpts, name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "listAsset", name, nuDexName, assetType, contractAddress, chainId, signature)
}

// ListAsset is a paid mutator transaction binding the contract method 0xd1c0a28a.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) ListAsset(name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ListAsset(&_VotingManagerContract.TransactOpts, name, nuDexName, assetType, contractAddress, chainId, signature)
}

// ListAsset is a paid mutator transaction binding the contract method 0xd1c0a28a.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) ListAsset(name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ListAsset(&_VotingManagerContract.TransactOpts, name, nuDexName, assetType, contractAddress, chainId, signature)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x5f05cb59.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) PreconfirmTask(opts *bind.TransactOpts, _taskId *big.Int, _result []byte, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "preconfirmTask", _taskId, _result, _signature)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x5f05cb59.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) PreconfirmTask(_taskId *big.Int, _result []byte, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.PreconfirmTask(&_VotingManagerContract.TransactOpts, _taskId, _result, _signature)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x5f05cb59.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) PreconfirmTask(_taskId *big.Int, _result []byte, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.PreconfirmTask(&_VotingManagerContract.TransactOpts, _taskId, _result, _signature)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x4e4a3a44.
//
// Solidity: function registerAccount(address _user, uint256 _account, uint8 _chain, uint256 _index, address _address, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) RegisterAccount(opts *bind.TransactOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "registerAccount", _user, _account, _chain, _index, _address, _signature)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x4e4a3a44.
//
// Solidity: function registerAccount(address _user, uint256 _account, uint8 _chain, uint256 _index, address _address, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) RegisterAccount(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RegisterAccount(&_VotingManagerContract.TransactOpts, _user, _account, _chain, _index, _address, _signature)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x4e4a3a44.
//
// Solidity: function registerAccount(address _user, uint256 _account, uint8 _chain, uint256 _index, address _address, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) RegisterAccount(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RegisterAccount(&_VotingManagerContract.TransactOpts, _user, _account, _chain, _index, _address, _signature)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x144aa686.
//
// Solidity: function removeParticipant(address participant, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) RemoveParticipant(opts *bind.TransactOpts, participant common.Address, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "removeParticipant", participant, signature)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x144aa686.
//
// Solidity: function removeParticipant(address participant, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) RemoveParticipant(participant common.Address, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RemoveParticipant(&_VotingManagerContract.TransactOpts, participant, signature)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x144aa686.
//
// Solidity: function removeParticipant(address participant, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) RemoveParticipant(participant common.Address, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RemoveParticipant(&_VotingManagerContract.TransactOpts, participant, signature)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0x77833f18.
//
// Solidity: function setRewardPerPeriod(uint256 newRewardPerPeriod, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SetRewardPerPeriod(opts *bind.TransactOpts, newRewardPerPeriod *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "setRewardPerPeriod", newRewardPerPeriod, signature)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0x77833f18.
//
// Solidity: function setRewardPerPeriod(uint256 newRewardPerPeriod, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SetRewardPerPeriod(newRewardPerPeriod *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SetRewardPerPeriod(&_VotingManagerContract.TransactOpts, newRewardPerPeriod, signature)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0x77833f18.
//
// Solidity: function setRewardPerPeriod(uint256 newRewardPerPeriod, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SetRewardPerPeriod(newRewardPerPeriod *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SetRewardPerPeriod(&_VotingManagerContract.TransactOpts, newRewardPerPeriod, signature)
}

// SubmitDepositInfo is a paid mutator transaction binding the contract method 0xf9cb78de.
//
// Solidity: function submitDepositInfo(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SubmitDepositInfo(opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int, chainId *big.Int, txInfo []byte, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "submitDepositInfo", targetAddress, amount, chainId, txInfo, extraInfo, signature)
}

// SubmitDepositInfo is a paid mutator transaction binding the contract method 0xf9cb78de.
//
// Solidity: function submitDepositInfo(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SubmitDepositInfo(targetAddress common.Address, amount *big.Int, chainId *big.Int, txInfo []byte, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitDepositInfo(&_VotingManagerContract.TransactOpts, targetAddress, amount, chainId, txInfo, extraInfo, signature)
}

// SubmitDepositInfo is a paid mutator transaction binding the contract method 0xf9cb78de.
//
// Solidity: function submitDepositInfo(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SubmitDepositInfo(targetAddress common.Address, amount *big.Int, chainId *big.Int, txInfo []byte, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitDepositInfo(&_VotingManagerContract.TransactOpts, targetAddress, amount, chainId, txInfo, extraInfo, signature)
}

// SubmitTaskReceipt is a paid mutator transaction binding the contract method 0x980d16c9.
//
// Solidity: function submitTaskReceipt(uint256 taskId, bytes result, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SubmitTaskReceipt(opts *bind.TransactOpts, taskId *big.Int, result []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "submitTaskReceipt", taskId, result, signature)
}

// SubmitTaskReceipt is a paid mutator transaction binding the contract method 0x980d16c9.
//
// Solidity: function submitTaskReceipt(uint256 taskId, bytes result, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SubmitTaskReceipt(taskId *big.Int, result []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitTaskReceipt(&_VotingManagerContract.TransactOpts, taskId, result, signature)
}

// SubmitTaskReceipt is a paid mutator transaction binding the contract method 0x980d16c9.
//
// Solidity: function submitTaskReceipt(uint256 taskId, bytes result, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SubmitTaskReceipt(taskId *big.Int, result []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitTaskReceipt(&_VotingManagerContract.TransactOpts, taskId, result, signature)
}

// SubmitWithdrawalInfo is a paid mutator transaction binding the contract method 0x455274a9.
//
// Solidity: function submitWithdrawalInfo(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SubmitWithdrawalInfo(opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int, chainId *big.Int, txInfo []byte, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "submitWithdrawalInfo", targetAddress, amount, chainId, txInfo, extraInfo, signature)
}

// SubmitWithdrawalInfo is a paid mutator transaction binding the contract method 0x455274a9.
//
// Solidity: function submitWithdrawalInfo(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SubmitWithdrawalInfo(targetAddress common.Address, amount *big.Int, chainId *big.Int, txInfo []byte, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitWithdrawalInfo(&_VotingManagerContract.TransactOpts, targetAddress, amount, chainId, txInfo, extraInfo, signature)
}

// SubmitWithdrawalInfo is a paid mutator transaction binding the contract method 0x455274a9.
//
// Solidity: function submitWithdrawalInfo(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SubmitWithdrawalInfo(targetAddress common.Address, amount *big.Int, chainId *big.Int, txInfo []byte, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitWithdrawalInfo(&_VotingManagerContract.TransactOpts, targetAddress, amount, chainId, txInfo, extraInfo, signature)
}

// VotingManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VotingManagerContract contract.
type VotingManagerContractInitializedIterator struct {
	Event *VotingManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractInitialized)
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
		it.Event = new(VotingManagerContractInitialized)
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
func (it *VotingManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractInitialized represents a Initialized event raised by the VotingManagerContract contract.
type VotingManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotingManagerContractInitializedIterator, error) {

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractInitializedIterator{contract: _VotingManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotingManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractInitialized)
				if err := _VotingManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseInitialized(log types.Log) (*VotingManagerContractInitialized, error) {
	event := new(VotingManagerContractInitialized)
	if err := _VotingManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractRewardPerPeriodVotedIterator is returned from FilterRewardPerPeriodVoted and is used to iterate over the raw logs and unpacked data for RewardPerPeriodVoted events raised by the VotingManagerContract contract.
type VotingManagerContractRewardPerPeriodVotedIterator struct {
	Event *VotingManagerContractRewardPerPeriodVoted // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractRewardPerPeriodVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractRewardPerPeriodVoted)
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
		it.Event = new(VotingManagerContractRewardPerPeriodVoted)
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
func (it *VotingManagerContractRewardPerPeriodVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractRewardPerPeriodVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractRewardPerPeriodVoted represents a RewardPerPeriodVoted event raised by the VotingManagerContract contract.
type VotingManagerContractRewardPerPeriodVoted struct {
	NewRewardPerPeriod *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardPerPeriodVoted is a free log retrieval operation binding the contract event 0x49c1ed88b0c8d97e2b2c5ed61e83a948d5ceb786bc26227eaab0cd748878cc66.
//
// Solidity: event RewardPerPeriodVoted(uint256 newRewardPerPeriod)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterRewardPerPeriodVoted(opts *bind.FilterOpts) (*VotingManagerContractRewardPerPeriodVotedIterator, error) {

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "RewardPerPeriodVoted")
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractRewardPerPeriodVotedIterator{contract: _VotingManagerContract.contract, event: "RewardPerPeriodVoted", logs: logs, sub: sub}, nil
}

// WatchRewardPerPeriodVoted is a free log subscription operation binding the contract event 0x49c1ed88b0c8d97e2b2c5ed61e83a948d5ceb786bc26227eaab0cd748878cc66.
//
// Solidity: event RewardPerPeriodVoted(uint256 newRewardPerPeriod)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchRewardPerPeriodVoted(opts *bind.WatchOpts, sink chan<- *VotingManagerContractRewardPerPeriodVoted) (event.Subscription, error) {

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "RewardPerPeriodVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractRewardPerPeriodVoted)
				if err := _VotingManagerContract.contract.UnpackLog(event, "RewardPerPeriodVoted", log); err != nil {
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

// ParseRewardPerPeriodVoted is a log parse operation binding the contract event 0x49c1ed88b0c8d97e2b2c5ed61e83a948d5ceb786bc26227eaab0cd748878cc66.
//
// Solidity: event RewardPerPeriodVoted(uint256 newRewardPerPeriod)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseRewardPerPeriodVoted(log types.Log) (*VotingManagerContractRewardPerPeriodVoted, error) {
	event := new(VotingManagerContractRewardPerPeriodVoted)
	if err := _VotingManagerContract.contract.UnpackLog(event, "RewardPerPeriodVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractSubmitterChosenIterator is returned from FilterSubmitterChosen and is used to iterate over the raw logs and unpacked data for SubmitterChosen events raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterChosenIterator struct {
	Event *VotingManagerContractSubmitterChosen // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractSubmitterChosenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractSubmitterChosen)
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
		it.Event = new(VotingManagerContractSubmitterChosen)
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
func (it *VotingManagerContractSubmitterChosenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractSubmitterChosenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractSubmitterChosen represents a SubmitterChosen event raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterChosen struct {
	NewSubmitter common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubmitterChosen is a free log retrieval operation binding the contract event 0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d.
//
// Solidity: event SubmitterChosen(address indexed newSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterSubmitterChosen(opts *bind.FilterOpts, newSubmitter []common.Address) (*VotingManagerContractSubmitterChosenIterator, error) {

	var newSubmitterRule []interface{}
	for _, newSubmitterItem := range newSubmitter {
		newSubmitterRule = append(newSubmitterRule, newSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "SubmitterChosen", newSubmitterRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractSubmitterChosenIterator{contract: _VotingManagerContract.contract, event: "SubmitterChosen", logs: logs, sub: sub}, nil
}

// WatchSubmitterChosen is a free log subscription operation binding the contract event 0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d.
//
// Solidity: event SubmitterChosen(address indexed newSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchSubmitterChosen(opts *bind.WatchOpts, sink chan<- *VotingManagerContractSubmitterChosen, newSubmitter []common.Address) (event.Subscription, error) {

	var newSubmitterRule []interface{}
	for _, newSubmitterItem := range newSubmitter {
		newSubmitterRule = append(newSubmitterRule, newSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "SubmitterChosen", newSubmitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractSubmitterChosen)
				if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterChosen", log); err != nil {
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

// ParseSubmitterChosen is a log parse operation binding the contract event 0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d.
//
// Solidity: event SubmitterChosen(address indexed newSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseSubmitterChosen(log types.Log) (*VotingManagerContractSubmitterChosen, error) {
	event := new(VotingManagerContractSubmitterChosen)
	if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterChosen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractSubmitterRotationRequestedIterator is returned from FilterSubmitterRotationRequested and is used to iterate over the raw logs and unpacked data for SubmitterRotationRequested events raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterRotationRequestedIterator struct {
	Event *VotingManagerContractSubmitterRotationRequested // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractSubmitterRotationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractSubmitterRotationRequested)
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
		it.Event = new(VotingManagerContractSubmitterRotationRequested)
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
func (it *VotingManagerContractSubmitterRotationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractSubmitterRotationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractSubmitterRotationRequested represents a SubmitterRotationRequested event raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterRotationRequested struct {
	Requester        common.Address
	CurrentSubmitter common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSubmitterRotationRequested is a free log retrieval operation binding the contract event 0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584.
//
// Solidity: event SubmitterRotationRequested(address indexed requester, address indexed currentSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterSubmitterRotationRequested(opts *bind.FilterOpts, requester []common.Address, currentSubmitter []common.Address) (*VotingManagerContractSubmitterRotationRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var currentSubmitterRule []interface{}
	for _, currentSubmitterItem := range currentSubmitter {
		currentSubmitterRule = append(currentSubmitterRule, currentSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "SubmitterRotationRequested", requesterRule, currentSubmitterRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractSubmitterRotationRequestedIterator{contract: _VotingManagerContract.contract, event: "SubmitterRotationRequested", logs: logs, sub: sub}, nil
}

// WatchSubmitterRotationRequested is a free log subscription operation binding the contract event 0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584.
//
// Solidity: event SubmitterRotationRequested(address indexed requester, address indexed currentSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchSubmitterRotationRequested(opts *bind.WatchOpts, sink chan<- *VotingManagerContractSubmitterRotationRequested, requester []common.Address, currentSubmitter []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var currentSubmitterRule []interface{}
	for _, currentSubmitterItem := range currentSubmitter {
		currentSubmitterRule = append(currentSubmitterRule, currentSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "SubmitterRotationRequested", requesterRule, currentSubmitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractSubmitterRotationRequested)
				if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterRotationRequested", log); err != nil {
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

// ParseSubmitterRotationRequested is a log parse operation binding the contract event 0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584.
//
// Solidity: event SubmitterRotationRequested(address indexed requester, address indexed currentSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseSubmitterRotationRequested(log types.Log) (*VotingManagerContractSubmitterRotationRequested, error) {
	event := new(VotingManagerContractSubmitterRotationRequested)
	if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterRotationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
