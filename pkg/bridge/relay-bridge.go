// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
)

// RelayBridgeMetaData contains all meta data concerning the RelayBridge contract.
var RelayBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"SentData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"TransmittedData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"dataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sendData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RelayBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayBridgeMetaData.ABI instead.
var RelayBridgeABI = RelayBridgeMetaData.ABI

// RelayBridge is an auto generated Go binding around an Ethereum contract.
type RelayBridge struct {
	RelayBridgeCaller     // Read-only binding to the contract
	RelayBridgeTransactor // Write-only binding to the contract
	RelayBridgeFilterer   // Log filterer for contract events
}

// RelayBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayBridgeSession struct {
	Contract     *RelayBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayBridgeCallerSession struct {
	Contract *RelayBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RelayBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayBridgeTransactorSession struct {
	Contract     *RelayBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RelayBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayBridgeRaw struct {
	Contract *RelayBridge // Generic contract binding to access the raw methods on
}

// RelayBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayBridgeCallerRaw struct {
	Contract *RelayBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RelayBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayBridgeTransactorRaw struct {
	Contract *RelayBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayBridge creates a new instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridge(address common.Address, backend bind.ContractBackend) (*RelayBridge, error) {
	contract, err := bindRelayBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayBridge{RelayBridgeCaller: RelayBridgeCaller{contract: contract}, RelayBridgeTransactor: RelayBridgeTransactor{contract: contract}, RelayBridgeFilterer: RelayBridgeFilterer{contract: contract}}, nil
}

// NewRelayBridgeCaller creates a new read-only instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridgeCaller(address common.Address, caller bind.ContractCaller) (*RelayBridgeCaller, error) {
	contract, err := bindRelayBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeCaller{contract: contract}, nil
}

// NewRelayBridgeTransactor creates a new write-only instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayBridgeTransactor, error) {
	contract, err := bindRelayBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeTransactor{contract: contract}, nil
}

// NewRelayBridgeFilterer creates a new log filterer instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayBridgeFilterer, error) {
	contract, err := bindRelayBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeFilterer{contract: contract}, nil
}

// bindRelayBridge binds a generic wrapper to an already deployed contract.
func bindRelayBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayBridge *RelayBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayBridge.Contract.RelayBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayBridge *RelayBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayBridge.Contract.RelayBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayBridge *RelayBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayBridge.Contract.RelayBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayBridge *RelayBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayBridge *RelayBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayBridge *RelayBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayBridge.Contract.contract.Transact(opts, method, params...)
}

// DataHash is a free data retrieval call binding the contract method 0xaea02932.
//
// Solidity: function dataHash(uint256 chainId, bytes data) pure returns(bytes32)
func (_RelayBridge *RelayBridgeCaller) DataHash(opts *bind.CallOpts, chainId *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "dataHash", chainId, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataHash is a free data retrieval call binding the contract method 0xaea02932.
//
// Solidity: function dataHash(uint256 chainId, bytes data) pure returns(bytes32)
func (_RelayBridge *RelayBridgeSession) DataHash(chainId *big.Int, data []byte) ([32]byte, error) {
	return _RelayBridge.Contract.DataHash(&_RelayBridge.CallOpts, chainId, data)
}

// DataHash is a free data retrieval call binding the contract method 0xaea02932.
//
// Solidity: function dataHash(uint256 chainId, bytes data) pure returns(bytes32)
func (_RelayBridge *RelayBridgeCallerSession) DataHash(chainId *big.Int, data []byte) ([32]byte, error) {
	return _RelayBridge.Contract.DataHash(&_RelayBridge.CallOpts, chainId, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayBridge *RelayBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayBridge *RelayBridgeSession) Owner() (common.Address, error) {
	return _RelayBridge.Contract.Owner(&_RelayBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayBridge *RelayBridgeCallerSession) Owner() (common.Address, error) {
	return _RelayBridge.Contract.Owner(&_RelayBridge.CallOpts)
}

// SendData is a free data retrieval call binding the contract method 0xbb7a7425.
//
// Solidity: function sendData(bytes32 ) view returns(bytes)
func (_RelayBridge *RelayBridgeCaller) SendData(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "sendData", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SendData is a free data retrieval call binding the contract method 0xbb7a7425.
//
// Solidity: function sendData(bytes32 ) view returns(bytes)
func (_RelayBridge *RelayBridgeSession) SendData(arg0 [32]byte) ([]byte, error) {
	return _RelayBridge.Contract.SendData(&_RelayBridge.CallOpts, arg0)
}

// SendData is a free data retrieval call binding the contract method 0xbb7a7425.
//
// Solidity: function sendData(bytes32 ) view returns(bytes)
func (_RelayBridge *RelayBridgeCallerSession) SendData(arg0 [32]byte) ([]byte, error) {
	return _RelayBridge.Contract.SendData(&_RelayBridge.CallOpts, arg0)
}

// Transmitted is a free data retrieval call binding the contract method 0xc84b3b2f.
//
// Solidity: function transmitted(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCaller) Transmitted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "transmitted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transmitted is a free data retrieval call binding the contract method 0xc84b3b2f.
//
// Solidity: function transmitted(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeSession) Transmitted(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Transmitted(&_RelayBridge.CallOpts, arg0)
}

// Transmitted is a free data retrieval call binding the contract method 0xc84b3b2f.
//
// Solidity: function transmitted(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCallerSession) Transmitted(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Transmitted(&_RelayBridge.CallOpts, arg0)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_RelayBridge *RelayBridgeCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_RelayBridge *RelayBridgeSession) Validator() (common.Address, error) {
	return _RelayBridge.Contract.Validator(&_RelayBridge.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_RelayBridge *RelayBridgeCallerSession) Validator() (common.Address, error) {
	return _RelayBridge.Contract.Validator(&_RelayBridge.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validator) returns()
func (_RelayBridge *RelayBridgeTransactor) Initialize(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "initialize", _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validator) returns()
func (_RelayBridge *RelayBridgeSession) Initialize(_validator common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.Initialize(&_RelayBridge.TransactOpts, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validator) returns()
func (_RelayBridge *RelayBridgeTransactorSession) Initialize(_validator common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.Initialize(&_RelayBridge.TransactOpts, _validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RelayBridge *RelayBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RelayBridge *RelayBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _RelayBridge.Contract.RenounceOwnership(&_RelayBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RelayBridge *RelayBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RelayBridge.Contract.RenounceOwnership(&_RelayBridge.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x69bf1944.
//
// Solidity: function send(uint256 chainId, bytes data) returns()
func (_RelayBridge *RelayBridgeTransactor) Send(opts *bind.TransactOpts, chainId *big.Int, data []byte) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "send", chainId, data)
}

// Send is a paid mutator transaction binding the contract method 0x69bf1944.
//
// Solidity: function send(uint256 chainId, bytes data) returns()
func (_RelayBridge *RelayBridgeSession) Send(chainId *big.Int, data []byte) (*types.Transaction, error) {
	return _RelayBridge.Contract.Send(&_RelayBridge.TransactOpts, chainId, data)
}

// Send is a paid mutator transaction binding the contract method 0x69bf1944.
//
// Solidity: function send(uint256 chainId, bytes data) returns()
func (_RelayBridge *RelayBridgeTransactorSession) Send(chainId *big.Int, data []byte) (*types.Transaction, error) {
	return _RelayBridge.Contract.Send(&_RelayBridge.TransactOpts, chainId, data)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _validator) returns()
func (_RelayBridge *RelayBridgeTransactor) SetValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "setValidator", _validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _validator) returns()
func (_RelayBridge *RelayBridgeSession) SetValidator(_validator common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.SetValidator(&_RelayBridge.TransactOpts, _validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _validator) returns()
func (_RelayBridge *RelayBridgeTransactorSession) SetValidator(_validator common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.SetValidator(&_RelayBridge.TransactOpts, _validator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RelayBridge *RelayBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RelayBridge *RelayBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.TransferOwnership(&_RelayBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RelayBridge *RelayBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.TransferOwnership(&_RelayBridge.TransactOpts, newOwner)
}

// Transmit is a paid mutator transaction binding the contract method 0x816f3430.
//
// Solidity: function transmit(uint256 fromChainId, bytes data) returns()
func (_RelayBridge *RelayBridgeTransactor) Transmit(opts *bind.TransactOpts, fromChainId *big.Int, data []byte) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "transmit", fromChainId, data)
}

// Transmit is a paid mutator transaction binding the contract method 0x816f3430.
//
// Solidity: function transmit(uint256 fromChainId, bytes data) returns()
func (_RelayBridge *RelayBridgeSession) Transmit(fromChainId *big.Int, data []byte) (*types.Transaction, error) {
	return _RelayBridge.Contract.Transmit(&_RelayBridge.TransactOpts, fromChainId, data)
}

// Transmit is a paid mutator transaction binding the contract method 0x816f3430.
//
// Solidity: function transmit(uint256 fromChainId, bytes data) returns()
func (_RelayBridge *RelayBridgeTransactorSession) Transmit(fromChainId *big.Int, data []byte) (*types.Transaction, error) {
	return _RelayBridge.Contract.Transmit(&_RelayBridge.TransactOpts, fromChainId, data)
}

// RelayBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RelayBridge contract.
type RelayBridgeInitializedIterator struct {
	Event *RelayBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *RelayBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeInitialized)
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
		it.Event = new(RelayBridgeInitialized)
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
func (it *RelayBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeInitialized represents a Initialized event raised by the RelayBridge contract.
type RelayBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayBridge *RelayBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*RelayBridgeInitializedIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeInitializedIterator{contract: _RelayBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayBridge *RelayBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RelayBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeInitialized)
				if err := _RelayBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayBridge *RelayBridgeFilterer) ParseInitialized(log types.Log) (*RelayBridgeInitialized, error) {
	event := new(RelayBridgeInitialized)
	if err := _RelayBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RelayBridge contract.
type RelayBridgeOwnershipTransferredIterator struct {
	Event *RelayBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RelayBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeOwnershipTransferred)
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
		it.Event = new(RelayBridgeOwnershipTransferred)
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
func (it *RelayBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the RelayBridge contract.
type RelayBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RelayBridge *RelayBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RelayBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeOwnershipTransferredIterator{contract: _RelayBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RelayBridge *RelayBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RelayBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeOwnershipTransferred)
				if err := _RelayBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RelayBridge *RelayBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*RelayBridgeOwnershipTransferred, error) {
	event := new(RelayBridgeOwnershipTransferred)
	if err := _RelayBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeSentDataIterator is returned from FilterSentData and is used to iterate over the raw logs and unpacked data for SentData events raised by the RelayBridge contract.
type RelayBridgeSentDataIterator struct {
	Event *RelayBridgeSentData // Event containing the contract specifics and raw log

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
func (it *RelayBridgeSentDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeSentData)
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
		it.Event = new(RelayBridgeSentData)
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
func (it *RelayBridgeSentDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeSentDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeSentData represents a SentData event raised by the RelayBridge contract.
type RelayBridgeSentData struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSentData is a free log retrieval operation binding the contract event 0x0772e5782827b3340e53f44c38640372637d26e3e45ca5506d50692f3d736bc6.
//
// Solidity: event SentData(bytes32 hash)
func (_RelayBridge *RelayBridgeFilterer) FilterSentData(opts *bind.FilterOpts) (*RelayBridgeSentDataIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "SentData")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeSentDataIterator{contract: _RelayBridge.contract, event: "SentData", logs: logs, sub: sub}, nil
}

// WatchSentData is a free log subscription operation binding the contract event 0x0772e5782827b3340e53f44c38640372637d26e3e45ca5506d50692f3d736bc6.
//
// Solidity: event SentData(bytes32 hash)
func (_RelayBridge *RelayBridgeFilterer) WatchSentData(opts *bind.WatchOpts, sink chan<- *RelayBridgeSentData) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "SentData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeSentData)
				if err := _RelayBridge.contract.UnpackLog(event, "SentData", log); err != nil {
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

// ParseSentData is a log parse operation binding the contract event 0x0772e5782827b3340e53f44c38640372637d26e3e45ca5506d50692f3d736bc6.
//
// Solidity: event SentData(bytes32 hash)
func (_RelayBridge *RelayBridgeFilterer) ParseSentData(log types.Log) (*RelayBridgeSentData, error) {
	event := new(RelayBridgeSentData)
	if err := _RelayBridge.contract.UnpackLog(event, "SentData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeTransmittedDataIterator is returned from FilterTransmittedData and is used to iterate over the raw logs and unpacked data for TransmittedData events raised by the RelayBridge contract.
type RelayBridgeTransmittedDataIterator struct {
	Event *RelayBridgeTransmittedData // Event containing the contract specifics and raw log

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
func (it *RelayBridgeTransmittedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeTransmittedData)
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
		it.Event = new(RelayBridgeTransmittedData)
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
func (it *RelayBridgeTransmittedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeTransmittedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeTransmittedData represents a TransmittedData event raised by the RelayBridge contract.
type RelayBridgeTransmittedData struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransmittedData is a free log retrieval operation binding the contract event 0x9fe4b61f465b1e0093aada1ea2c53d139a58503717d5c6d8224581c278515743.
//
// Solidity: event TransmittedData(bytes32 hash)
func (_RelayBridge *RelayBridgeFilterer) FilterTransmittedData(opts *bind.FilterOpts) (*RelayBridgeTransmittedDataIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "TransmittedData")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeTransmittedDataIterator{contract: _RelayBridge.contract, event: "TransmittedData", logs: logs, sub: sub}, nil
}

// WatchTransmittedData is a free log subscription operation binding the contract event 0x9fe4b61f465b1e0093aada1ea2c53d139a58503717d5c6d8224581c278515743.
//
// Solidity: event TransmittedData(bytes32 hash)
func (_RelayBridge *RelayBridgeFilterer) WatchTransmittedData(opts *bind.WatchOpts, sink chan<- *RelayBridgeTransmittedData) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "TransmittedData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeTransmittedData)
				if err := _RelayBridge.contract.UnpackLog(event, "TransmittedData", log); err != nil {
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

// ParseTransmittedData is a log parse operation binding the contract event 0x9fe4b61f465b1e0093aada1ea2c53d139a58503717d5c6d8224581c278515743.
//
// Solidity: event TransmittedData(bytes32 hash)
func (_RelayBridge *RelayBridgeFilterer) ParseTransmittedData(log types.Log) (*RelayBridgeTransmittedData, error) {
	event := new(RelayBridgeTransmittedData)
	if err := _RelayBridge.contract.UnpackLog(event, "TransmittedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
