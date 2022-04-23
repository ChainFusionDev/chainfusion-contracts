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

// TokenManagerMetaData contains all meta data concerning the TokenManager contract.
var TokenManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenManagerMetaData.ABI instead.
var TokenManagerABI = TokenManagerMetaData.ABI

// TokenManager is an auto generated Go binding around an Ethereum contract.
type TokenManager struct {
	TokenManagerCaller     // Read-only binding to the contract
	TokenManagerTransactor // Write-only binding to the contract
	TokenManagerFilterer   // Log filterer for contract events
}

// TokenManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenManagerSession struct {
	Contract     *TokenManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenManagerCallerSession struct {
	Contract *TokenManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenManagerTransactorSession struct {
	Contract     *TokenManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenManagerRaw struct {
	Contract *TokenManager // Generic contract binding to access the raw methods on
}

// TokenManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenManagerCallerRaw struct {
	Contract *TokenManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenManagerTransactorRaw struct {
	Contract *TokenManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenManager creates a new instance of TokenManager, bound to a specific deployed contract.
func NewTokenManager(address common.Address, backend bind.ContractBackend) (*TokenManager, error) {
	contract, err := bindTokenManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenManager{TokenManagerCaller: TokenManagerCaller{contract: contract}, TokenManagerTransactor: TokenManagerTransactor{contract: contract}, TokenManagerFilterer: TokenManagerFilterer{contract: contract}}, nil
}

// NewTokenManagerCaller creates a new read-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerCaller(address common.Address, caller bind.ContractCaller) (*TokenManagerCaller, error) {
	contract, err := bindTokenManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerCaller{contract: contract}, nil
}

// NewTokenManagerTransactor creates a new write-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenManagerTransactor, error) {
	contract, err := bindTokenManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerTransactor{contract: contract}, nil
}

// NewTokenManagerFilterer creates a new log filterer instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenManagerFilterer, error) {
	contract, err := bindTokenManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenManagerFilterer{contract: contract}, nil
}

// bindTokenManager binds a generic wrapper to an already deployed contract.
func bindTokenManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.TokenManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transact(opts, method, params...)
}

// GetDestinationToken is a free data retrieval call binding the contract method 0xabf1535f.
//
// Solidity: function getDestinationToken(address _token, uint256 _chainId) view returns(address)
func (_TokenManager *TokenManagerCaller) GetDestinationToken(opts *bind.CallOpts, _token common.Address, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "getDestinationToken", _token, _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDestinationToken is a free data retrieval call binding the contract method 0xabf1535f.
//
// Solidity: function getDestinationToken(address _token, uint256 _chainId) view returns(address)
func (_TokenManager *TokenManagerSession) GetDestinationToken(_token common.Address, _chainId *big.Int) (common.Address, error) {
	return _TokenManager.Contract.GetDestinationToken(&_TokenManager.CallOpts, _token, _chainId)
}

// GetDestinationToken is a free data retrieval call binding the contract method 0xabf1535f.
//
// Solidity: function getDestinationToken(address _token, uint256 _chainId) view returns(address)
func (_TokenManager *TokenManagerCallerSession) GetDestinationToken(_token common.Address, _chainId *big.Int) (common.Address, error) {
	return _TokenManager.Contract.GetDestinationToken(&_TokenManager.CallOpts, _token, _chainId)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_TokenManager *TokenManagerCaller) IsTokenSupported(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "isTokenSupported", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_TokenManager *TokenManagerSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _TokenManager.Contract.IsTokenSupported(&_TokenManager.CallOpts, _token)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _TokenManager.Contract.IsTokenSupported(&_TokenManager.CallOpts, _token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerSession) Owner() (common.Address, error) {
	return _TokenManager.Contract.Owner(&_TokenManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerCallerSession) Owner() (common.Address, error) {
	return _TokenManager.Contract.Owner(&_TokenManager.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool isSupported)
func (_TokenManager *TokenManagerCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool isSupported)
func (_TokenManager *TokenManagerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _TokenManager.Contract.SupportedTokens(&_TokenManager.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool isSupported)
func (_TokenManager *TokenManagerCallerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _TokenManager.Contract.SupportedTokens(&_TokenManager.CallOpts, arg0)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x49be7387.
//
// Solidity: function addSupportedToken(uint256 _chainId, address _token, address _destinationToken) returns()
func (_TokenManager *TokenManagerTransactor) AddSupportedToken(opts *bind.TransactOpts, _chainId *big.Int, _token common.Address, _destinationToken common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "addSupportedToken", _chainId, _token, _destinationToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x49be7387.
//
// Solidity: function addSupportedToken(uint256 _chainId, address _token, address _destinationToken) returns()
func (_TokenManager *TokenManagerSession) AddSupportedToken(_chainId *big.Int, _token common.Address, _destinationToken common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.AddSupportedToken(&_TokenManager.TransactOpts, _chainId, _token, _destinationToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x49be7387.
//
// Solidity: function addSupportedToken(uint256 _chainId, address _token, address _destinationToken) returns()
func (_TokenManager *TokenManagerTransactorSession) AddSupportedToken(_chainId *big.Int, _token common.Address, _destinationToken common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.AddSupportedToken(&_TokenManager.TransactOpts, _chainId, _token, _destinationToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_TokenManager *TokenManagerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_TokenManager *TokenManagerSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Initialize(&_TokenManager.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_TokenManager *TokenManagerTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Initialize(&_TokenManager.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenManager.Contract.RenounceOwnership(&_TokenManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenManager.Contract.RenounceOwnership(&_TokenManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.TransferOwnership(&_TokenManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.TransferOwnership(&_TokenManager.TransactOpts, newOwner)
}

// TokenManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenManager contract.
type TokenManagerOwnershipTransferredIterator struct {
	Event *TokenManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerOwnershipTransferred)
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
		it.Event = new(TokenManagerOwnershipTransferred)
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
func (it *TokenManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TokenManager contract.
type TokenManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenManager *TokenManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenManagerOwnershipTransferredIterator{contract: _TokenManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenManager *TokenManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerOwnershipTransferred)
				if err := _TokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenManager *TokenManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TokenManagerOwnershipTransferred, error) {
	event := new(TokenManagerOwnershipTransferred)
	if err := _TokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
