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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationToken\",\"type\":\"address\"}],\"name\":\"setDestinationToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isMintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractSignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// IsTokenEnabled is a free data retrieval call binding the contract method 0x748538d9.
//
// Solidity: function isTokenEnabled(address _token) view returns(bool)
func (_TokenManager *TokenManagerCaller) IsTokenEnabled(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "isTokenEnabled", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenEnabled is a free data retrieval call binding the contract method 0x748538d9.
//
// Solidity: function isTokenEnabled(address _token) view returns(bool)
func (_TokenManager *TokenManagerSession) IsTokenEnabled(_token common.Address) (bool, error) {
	return _TokenManager.Contract.IsTokenEnabled(&_TokenManager.CallOpts, _token)
}

// IsTokenEnabled is a free data retrieval call binding the contract method 0x748538d9.
//
// Solidity: function isTokenEnabled(address _token) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsTokenEnabled(_token common.Address) (bool, error) {
	return _TokenManager.Contract.IsTokenEnabled(&_TokenManager.CallOpts, _token)
}

// IsTokenMintable is a free data retrieval call binding the contract method 0x5da619a9.
//
// Solidity: function isTokenMintable(address _token) view returns(bool)
func (_TokenManager *TokenManagerCaller) IsTokenMintable(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "isTokenMintable", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenMintable is a free data retrieval call binding the contract method 0x5da619a9.
//
// Solidity: function isTokenMintable(address _token) view returns(bool)
func (_TokenManager *TokenManagerSession) IsTokenMintable(_token common.Address) (bool, error) {
	return _TokenManager.Contract.IsTokenMintable(&_TokenManager.CallOpts, _token)
}

// IsTokenMintable is a free data retrieval call binding the contract method 0x5da619a9.
//
// Solidity: function isTokenMintable(address _token) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsTokenMintable(_token common.Address) (bool, error) {
	return _TokenManager.Contract.IsTokenMintable(&_TokenManager.CallOpts, _token)
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
// Solidity: function supportedTokens(address ) view returns(bool isEnabled, bool isMintable)
func (_TokenManager *TokenManagerCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsEnabled  bool
	IsMintable bool
}, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "supportedTokens", arg0)

	outstruct := new(struct {
		IsEnabled  bool
		IsMintable bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsMintable = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool isEnabled, bool isMintable)
func (_TokenManager *TokenManagerSession) SupportedTokens(arg0 common.Address) (struct {
	IsEnabled  bool
	IsMintable bool
}, error) {
	return _TokenManager.Contract.SupportedTokens(&_TokenManager.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool isEnabled, bool isMintable)
func (_TokenManager *TokenManagerCallerSession) SupportedTokens(arg0 common.Address) (struct {
	IsEnabled  bool
	IsMintable bool
}, error) {
	return _TokenManager.Contract.SupportedTokens(&_TokenManager.CallOpts, arg0)
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

// SetDestinationToken is a paid mutator transaction binding the contract method 0xff1b14be.
//
// Solidity: function setDestinationToken(uint256 _chainId, address _token, address _destinationToken) returns()
func (_TokenManager *TokenManagerTransactor) SetDestinationToken(opts *bind.TransactOpts, _chainId *big.Int, _token common.Address, _destinationToken common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setDestinationToken", _chainId, _token, _destinationToken)
}

// SetDestinationToken is a paid mutator transaction binding the contract method 0xff1b14be.
//
// Solidity: function setDestinationToken(uint256 _chainId, address _token, address _destinationToken) returns()
func (_TokenManager *TokenManagerSession) SetDestinationToken(_chainId *big.Int, _token common.Address, _destinationToken common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.SetDestinationToken(&_TokenManager.TransactOpts, _chainId, _token, _destinationToken)
}

// SetDestinationToken is a paid mutator transaction binding the contract method 0xff1b14be.
//
// Solidity: function setDestinationToken(uint256 _chainId, address _token, address _destinationToken) returns()
func (_TokenManager *TokenManagerTransactorSession) SetDestinationToken(_chainId *big.Int, _token common.Address, _destinationToken common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.SetDestinationToken(&_TokenManager.TransactOpts, _chainId, _token, _destinationToken)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x17fe72bb.
//
// Solidity: function setEnabled(address _token, bool _isEnabled) returns()
func (_TokenManager *TokenManagerTransactor) SetEnabled(opts *bind.TransactOpts, _token common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setEnabled", _token, _isEnabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x17fe72bb.
//
// Solidity: function setEnabled(address _token, bool _isEnabled) returns()
func (_TokenManager *TokenManagerSession) SetEnabled(_token common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _TokenManager.Contract.SetEnabled(&_TokenManager.TransactOpts, _token, _isEnabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x17fe72bb.
//
// Solidity: function setEnabled(address _token, bool _isEnabled) returns()
func (_TokenManager *TokenManagerTransactorSession) SetEnabled(_token common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _TokenManager.Contract.SetEnabled(&_TokenManager.TransactOpts, _token, _isEnabled)
}

// SetMintable is a paid mutator transaction binding the contract method 0xf7eb06c4.
//
// Solidity: function setMintable(address _token, bool _isMintable) returns()
func (_TokenManager *TokenManagerTransactor) SetMintable(opts *bind.TransactOpts, _token common.Address, _isMintable bool) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setMintable", _token, _isMintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0xf7eb06c4.
//
// Solidity: function setMintable(address _token, bool _isMintable) returns()
func (_TokenManager *TokenManagerSession) SetMintable(_token common.Address, _isMintable bool) (*types.Transaction, error) {
	return _TokenManager.Contract.SetMintable(&_TokenManager.TransactOpts, _token, _isMintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0xf7eb06c4.
//
// Solidity: function setMintable(address _token, bool _isMintable) returns()
func (_TokenManager *TokenManagerTransactorSession) SetMintable(_token common.Address, _isMintable bool) (*types.Transaction, error) {
	return _TokenManager.Contract.SetMintable(&_TokenManager.TransactOpts, _token, _isMintable)
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
