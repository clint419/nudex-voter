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

// AccountManagerContractMetaData contains all meta data concerning the AccountManagerContract contract.
var AccountManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addressRecord\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressRecord\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chain\",\"type\":\"uint8\",\"internalType\":\"enumIAccountManager.Chain\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerNewAddress\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chain\",\"type\":\"uint8\",\"internalType\":\"enumIAccountManager.Chain\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userMapping\",\"inputs\":[{\"name\":\"depositAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAccountManager.Chain\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddressRegistered\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIAccountManager.Chain\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAccountNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RegisteredAccount\",\"inputs\":[]}]",
}

// AccountManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountManagerContractMetaData.ABI instead.
var AccountManagerContractABI = AccountManagerContractMetaData.ABI

// AccountManagerContract is an auto generated Go binding around an Ethereum contract.
type AccountManagerContract struct {
	AccountManagerContractCaller     // Read-only binding to the contract
	AccountManagerContractTransactor // Write-only binding to the contract
	AccountManagerContractFilterer   // Log filterer for contract events
}

// AccountManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerContractSession struct {
	Contract     *AccountManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AccountManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerContractCallerSession struct {
	Contract *AccountManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AccountManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerContractTransactorSession struct {
	Contract     *AccountManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AccountManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerContractRaw struct {
	Contract *AccountManagerContract // Generic contract binding to access the raw methods on
}

// AccountManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerContractCallerRaw struct {
	Contract *AccountManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerContractTransactorRaw struct {
	Contract *AccountManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManagerContract creates a new instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContract(address common.Address, backend bind.ContractBackend) (*AccountManagerContract, error) {
	contract, err := bindAccountManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContract{AccountManagerContractCaller: AccountManagerContractCaller{contract: contract}, AccountManagerContractTransactor: AccountManagerContractTransactor{contract: contract}, AccountManagerContractFilterer: AccountManagerContractFilterer{contract: contract}}, nil
}

// NewAccountManagerContractCaller creates a new read-only instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContractCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerContractCaller, error) {
	contract, err := bindAccountManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractCaller{contract: contract}, nil
}

// NewAccountManagerContractTransactor creates a new write-only instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerContractTransactor, error) {
	contract, err := bindAccountManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractTransactor{contract: contract}, nil
}

// NewAccountManagerContractFilterer creates a new log filterer instance of AccountManagerContract, bound to a specific deployed contract.
func NewAccountManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerContractFilterer, error) {
	contract, err := bindAccountManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractFilterer{contract: contract}, nil
}

// bindAccountManagerContract binds a generic wrapper to an already deployed contract.
func bindAccountManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerContract *AccountManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerContract.Contract.AccountManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerContract *AccountManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.AccountManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerContract *AccountManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.AccountManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerContract *AccountManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerContract *AccountManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerContract *AccountManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.contract.Transact(opts, method, params...)
}

// AddressRecord is a free data retrieval call binding the contract method 0x5192a407.
//
// Solidity: function addressRecord(bytes ) view returns(address)
func (_AccountManagerContract *AccountManagerContractCaller) AddressRecord(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "addressRecord", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRecord is a free data retrieval call binding the contract method 0x5192a407.
//
// Solidity: function addressRecord(bytes ) view returns(address)
func (_AccountManagerContract *AccountManagerContractSession) AddressRecord(arg0 []byte) (common.Address, error) {
	return _AccountManagerContract.Contract.AddressRecord(&_AccountManagerContract.CallOpts, arg0)
}

// AddressRecord is a free data retrieval call binding the contract method 0x5192a407.
//
// Solidity: function addressRecord(bytes ) view returns(address)
func (_AccountManagerContract *AccountManagerContractCallerSession) AddressRecord(arg0 []byte) (common.Address, error) {
	return _AccountManagerContract.Contract.AddressRecord(&_AccountManagerContract.CallOpts, arg0)
}

// GetAddressRecord is a free data retrieval call binding the contract method 0x38652e68.
//
// Solidity: function getAddressRecord(address _user, uint256 _account, uint8 _chain, uint256 _index) view returns(address)
func (_AccountManagerContract *AccountManagerContractCaller) GetAddressRecord(opts *bind.CallOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "getAddressRecord", _user, _account, _chain, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressRecord is a free data retrieval call binding the contract method 0x38652e68.
//
// Solidity: function getAddressRecord(address _user, uint256 _account, uint8 _chain, uint256 _index) view returns(address)
func (_AccountManagerContract *AccountManagerContractSession) GetAddressRecord(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (common.Address, error) {
	return _AccountManagerContract.Contract.GetAddressRecord(&_AccountManagerContract.CallOpts, _user, _account, _chain, _index)
}

// GetAddressRecord is a free data retrieval call binding the contract method 0x38652e68.
//
// Solidity: function getAddressRecord(address _user, uint256 _account, uint8 _chain, uint256 _index) view returns(address)
func (_AccountManagerContract *AccountManagerContractCallerSession) GetAddressRecord(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (common.Address, error) {
	return _AccountManagerContract.Contract.GetAddressRecord(&_AccountManagerContract.CallOpts, _user, _account, _chain, _index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManagerContract *AccountManagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManagerContract *AccountManagerContractSession) Owner() (common.Address, error) {
	return _AccountManagerContract.Contract.Owner(&_AccountManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManagerContract *AccountManagerContractCallerSession) Owner() (common.Address, error) {
	return _AccountManagerContract.Contract.Owner(&_AccountManagerContract.CallOpts)
}

// UserMapping is a free data retrieval call binding the contract method 0x39835c17.
//
// Solidity: function userMapping(address depositAddress, uint8 ) view returns(address user)
func (_AccountManagerContract *AccountManagerContractCaller) UserMapping(opts *bind.CallOpts, depositAddress common.Address, arg1 uint8) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerContract.contract.Call(opts, &out, "userMapping", depositAddress, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserMapping is a free data retrieval call binding the contract method 0x39835c17.
//
// Solidity: function userMapping(address depositAddress, uint8 ) view returns(address user)
func (_AccountManagerContract *AccountManagerContractSession) UserMapping(depositAddress common.Address, arg1 uint8) (common.Address, error) {
	return _AccountManagerContract.Contract.UserMapping(&_AccountManagerContract.CallOpts, depositAddress, arg1)
}

// UserMapping is a free data retrieval call binding the contract method 0x39835c17.
//
// Solidity: function userMapping(address depositAddress, uint8 ) view returns(address user)
func (_AccountManagerContract *AccountManagerContractCallerSession) UserMapping(depositAddress common.Address, arg1 uint8) (common.Address, error) {
	return _AccountManagerContract.Contract.UserMapping(&_AccountManagerContract.CallOpts, depositAddress, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AccountManagerContract *AccountManagerContractSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.Initialize(&_AccountManagerContract.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.Initialize(&_AccountManagerContract.TransactOpts, _owner)
}

// RegisterNewAddress is a paid mutator transaction binding the contract method 0x1a98a163.
//
// Solidity: function registerNewAddress(address _user, uint256 _account, uint8 _chain, uint256 _index, address _address) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) RegisterNewAddress(opts *bind.TransactOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "registerNewAddress", _user, _account, _chain, _index, _address)
}

// RegisterNewAddress is a paid mutator transaction binding the contract method 0x1a98a163.
//
// Solidity: function registerNewAddress(address _user, uint256 _account, uint8 _chain, uint256 _index, address _address) returns()
func (_AccountManagerContract *AccountManagerContractSession) RegisterNewAddress(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RegisterNewAddress(&_AccountManagerContract.TransactOpts, _user, _account, _chain, _index, _address)
}

// RegisterNewAddress is a paid mutator transaction binding the contract method 0x1a98a163.
//
// Solidity: function registerNewAddress(address _user, uint256 _account, uint8 _chain, uint256 _index, address _address) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) RegisterNewAddress(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RegisterNewAddress(&_AccountManagerContract.TransactOpts, _user, _account, _chain, _index, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManagerContract *AccountManagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManagerContract *AccountManagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RenounceOwnership(&_AccountManagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountManagerContract.Contract.RenounceOwnership(&_AccountManagerContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManagerContract *AccountManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManagerContract *AccountManagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.TransferOwnership(&_AccountManagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManagerContract *AccountManagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManagerContract.Contract.TransferOwnership(&_AccountManagerContract.TransactOpts, newOwner)
}

// AccountManagerContractAddressRegisteredIterator is returned from FilterAddressRegistered and is used to iterate over the raw logs and unpacked data for AddressRegistered events raised by the AccountManagerContract contract.
type AccountManagerContractAddressRegisteredIterator struct {
	Event *AccountManagerContractAddressRegistered // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractAddressRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractAddressRegistered)
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
		it.Event = new(AccountManagerContractAddressRegistered)
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
func (it *AccountManagerContractAddressRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractAddressRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractAddressRegistered represents a AddressRegistered event raised by the AccountManagerContract contract.
type AccountManagerContractAddressRegistered struct {
	User       common.Address
	Account    *big.Int
	ChainId    uint8
	Index      *big.Int
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressRegistered is a free log retrieval operation binding the contract event 0x10d3f3c3d0c7da2f6751b14c10b9dbc6e04f5ebc6b798a6e220f3857ba1cd454.
//
// Solidity: event AddressRegistered(address indexed user, uint256 account, uint8 indexed chainId, uint256 index, address indexed newAddress)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterAddressRegistered(opts *bind.FilterOpts, user []common.Address, chainId []uint8, newAddress []common.Address) (*AccountManagerContractAddressRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "AddressRegistered", userRule, chainIdRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractAddressRegisteredIterator{contract: _AccountManagerContract.contract, event: "AddressRegistered", logs: logs, sub: sub}, nil
}

// WatchAddressRegistered is a free log subscription operation binding the contract event 0x10d3f3c3d0c7da2f6751b14c10b9dbc6e04f5ebc6b798a6e220f3857ba1cd454.
//
// Solidity: event AddressRegistered(address indexed user, uint256 account, uint8 indexed chainId, uint256 index, address indexed newAddress)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchAddressRegistered(opts *bind.WatchOpts, sink chan<- *AccountManagerContractAddressRegistered, user []common.Address, chainId []uint8, newAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "AddressRegistered", userRule, chainIdRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractAddressRegistered)
				if err := _AccountManagerContract.contract.UnpackLog(event, "AddressRegistered", log); err != nil {
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

// ParseAddressRegistered is a log parse operation binding the contract event 0x10d3f3c3d0c7da2f6751b14c10b9dbc6e04f5ebc6b798a6e220f3857ba1cd454.
//
// Solidity: event AddressRegistered(address indexed user, uint256 account, uint8 indexed chainId, uint256 index, address indexed newAddress)
func (_AccountManagerContract *AccountManagerContractFilterer) ParseAddressRegistered(log types.Log) (*AccountManagerContractAddressRegistered, error) {
	event := new(AccountManagerContractAddressRegistered)
	if err := _AccountManagerContract.contract.UnpackLog(event, "AddressRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccountManagerContract contract.
type AccountManagerContractInitializedIterator struct {
	Event *AccountManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractInitialized)
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
		it.Event = new(AccountManagerContractInitialized)
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
func (it *AccountManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractInitialized represents a Initialized event raised by the AccountManagerContract contract.
type AccountManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountManagerContractInitializedIterator, error) {

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractInitializedIterator{contract: _AccountManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractInitialized)
				if err := _AccountManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AccountManagerContract *AccountManagerContractFilterer) ParseInitialized(log types.Log) (*AccountManagerContractInitialized, error) {
	event := new(AccountManagerContractInitialized)
	if err := _AccountManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountManagerContract contract.
type AccountManagerContractOwnershipTransferredIterator struct {
	Event *AccountManagerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountManagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerContractOwnershipTransferred)
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
		it.Event = new(AccountManagerContractOwnershipTransferred)
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
func (it *AccountManagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the AccountManagerContract contract.
type AccountManagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountManagerContract *AccountManagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountManagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountManagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerContractOwnershipTransferredIterator{contract: _AccountManagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountManagerContract *AccountManagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountManagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountManagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerContractOwnershipTransferred)
				if err := _AccountManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountManagerContract *AccountManagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*AccountManagerContractOwnershipTransferred, error) {
	event := new(AccountManagerContractOwnershipTransferred)
	if err := _AccountManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
