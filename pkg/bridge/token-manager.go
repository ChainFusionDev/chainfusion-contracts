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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"deployToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"enumTokenType\",\"name\":\"_tokenType\",\"type\":\"uint8\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractSignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetType is a free data retrieval call binding the contract method 0xf429f19b.
//
// Solidity: function getType(address _token) view returns(uint8)
func (_TokenManager *TokenManagerCaller) GetType(opts *bind.CallOpts, _token common.Address) (uint8, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "getType", _token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0xf429f19b.
//
// Solidity: function getType(address _token) view returns(uint8)
func (_TokenManager *TokenManagerSession) GetType(_token common.Address) (uint8, error) {
	return _TokenManager.Contract.GetType(&_TokenManager.CallOpts, _token)
}

// GetType is a free data retrieval call binding the contract method 0xf429f19b.
//
// Solidity: function getType(address _token) view returns(uint8)
func (_TokenManager *TokenManagerCallerSession) GetType(_token common.Address) (uint8, error) {
	return _TokenManager.Contract.GetType(&_TokenManager.CallOpts, _token)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_TokenManager *TokenManagerCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_TokenManager *TokenManagerSession) SignerStorage() (common.Address, error) {
	return _TokenManager.Contract.SignerStorage(&_TokenManager.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_TokenManager *TokenManagerCallerSession) SignerStorage() (common.Address, error) {
	return _TokenManager.Contract.SignerStorage(&_TokenManager.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(uint8)
func (_TokenManager *TokenManagerCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(uint8)
func (_TokenManager *TokenManagerSession) SupportedTokens(arg0 common.Address) (uint8, error) {
	return _TokenManager.Contract.SupportedTokens(&_TokenManager.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(uint8)
func (_TokenManager *TokenManagerCallerSession) SupportedTokens(arg0 common.Address) (uint8, error) {
	return _TokenManager.Contract.SupportedTokens(&_TokenManager.CallOpts, arg0)
}

// DeployToken is a paid mutator transaction binding the contract method 0x1d12b938.
//
// Solidity: function deployToken(string name, string symbol) returns(address)
func (_TokenManager *TokenManagerTransactor) DeployToken(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "deployToken", name, symbol)
}

// DeployToken is a paid mutator transaction binding the contract method 0x1d12b938.
//
// Solidity: function deployToken(string name, string symbol) returns(address)
func (_TokenManager *TokenManagerSession) DeployToken(name string, symbol string) (*types.Transaction, error) {
	return _TokenManager.Contract.DeployToken(&_TokenManager.TransactOpts, name, symbol)
}

// DeployToken is a paid mutator transaction binding the contract method 0x1d12b938.
//
// Solidity: function deployToken(string name, string symbol) returns(address)
func (_TokenManager *TokenManagerTransactorSession) DeployToken(name string, symbol string) (*types.Transaction, error) {
	return _TokenManager.Contract.DeployToken(&_TokenManager.TransactOpts, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signerStorage) returns()
func (_TokenManager *TokenManagerTransactor) Initialize(opts *bind.TransactOpts, _signerStorage common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "initialize", _signerStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signerStorage) returns()
func (_TokenManager *TokenManagerSession) Initialize(_signerStorage common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Initialize(&_TokenManager.TransactOpts, _signerStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signerStorage) returns()
func (_TokenManager *TokenManagerTransactorSession) Initialize(_signerStorage common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Initialize(&_TokenManager.TransactOpts, _signerStorage)
}

// SetToken is a paid mutator transaction binding the contract method 0xc4767b95.
//
// Solidity: function setToken(address _token, uint8 _tokenType) returns()
func (_TokenManager *TokenManagerTransactor) SetToken(opts *bind.TransactOpts, _token common.Address, _tokenType uint8) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setToken", _token, _tokenType)
}

// SetToken is a paid mutator transaction binding the contract method 0xc4767b95.
//
// Solidity: function setToken(address _token, uint8 _tokenType) returns()
func (_TokenManager *TokenManagerSession) SetToken(_token common.Address, _tokenType uint8) (*types.Transaction, error) {
	return _TokenManager.Contract.SetToken(&_TokenManager.TransactOpts, _token, _tokenType)
}

// SetToken is a paid mutator transaction binding the contract method 0xc4767b95.
//
// Solidity: function setToken(address _token, uint8 _tokenType) returns()
func (_TokenManager *TokenManagerTransactorSession) SetToken(_token common.Address, _tokenType uint8) (*types.Transaction, error) {
	return _TokenManager.Contract.SetToken(&_TokenManager.TransactOpts, _token, _tokenType)
}

// TokenManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenManager contract.
type TokenManagerInitializedIterator struct {
	Event *TokenManagerInitialized // Event containing the contract specifics and raw log

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
func (it *TokenManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerInitialized)
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
		it.Event = new(TokenManagerInitialized)
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
func (it *TokenManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerInitialized represents a Initialized event raised by the TokenManager contract.
type TokenManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenManager *TokenManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenManagerInitializedIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenManagerInitializedIterator{contract: _TokenManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenManager *TokenManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerInitialized)
				if err := _TokenManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokenManager *TokenManagerFilterer) ParseInitialized(log types.Log) (*TokenManagerInitialized, error) {
	event := new(TokenManagerInitialized)
	if err := _TokenManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
