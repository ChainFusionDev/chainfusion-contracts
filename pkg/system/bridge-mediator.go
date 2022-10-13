// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system

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

// BridgeMediatorMetaData contains all meta data concerning the BridgeMediator contract.
var BridgeMediatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mediatorAddress\",\"type\":\"address\"}],\"name\":\"MediatorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sourceData\",\"type\":\"bytes\"}],\"name\":\"mediate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeMediatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMediatorMetaData.ABI instead.
var BridgeMediatorABI = BridgeMediatorMetaData.ABI

// BridgeMediator is an auto generated Go binding around an Ethereum contract.
type BridgeMediator struct {
	BridgeMediatorCaller     // Read-only binding to the contract
	BridgeMediatorTransactor // Write-only binding to the contract
	BridgeMediatorFilterer   // Log filterer for contract events
}

// BridgeMediatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeMediatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeMediatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeMediatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeMediatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeMediatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeMediatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeMediatorSession struct {
	Contract     *BridgeMediator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeMediatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeMediatorCallerSession struct {
	Contract *BridgeMediatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeMediatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeMediatorTransactorSession struct {
	Contract     *BridgeMediatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeMediatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeMediatorRaw struct {
	Contract *BridgeMediator // Generic contract binding to access the raw methods on
}

// BridgeMediatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeMediatorCallerRaw struct {
	Contract *BridgeMediatorCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeMediatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeMediatorTransactorRaw struct {
	Contract *BridgeMediatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeMediator creates a new instance of BridgeMediator, bound to a specific deployed contract.
func NewBridgeMediator(address common.Address, backend bind.ContractBackend) (*BridgeMediator, error) {
	contract, err := bindBridgeMediator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeMediator{BridgeMediatorCaller: BridgeMediatorCaller{contract: contract}, BridgeMediatorTransactor: BridgeMediatorTransactor{contract: contract}, BridgeMediatorFilterer: BridgeMediatorFilterer{contract: contract}}, nil
}

// NewBridgeMediatorCaller creates a new read-only instance of BridgeMediator, bound to a specific deployed contract.
func NewBridgeMediatorCaller(address common.Address, caller bind.ContractCaller) (*BridgeMediatorCaller, error) {
	contract, err := bindBridgeMediator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorCaller{contract: contract}, nil
}

// NewBridgeMediatorTransactor creates a new write-only instance of BridgeMediator, bound to a specific deployed contract.
func NewBridgeMediatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeMediatorTransactor, error) {
	contract, err := bindBridgeMediator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorTransactor{contract: contract}, nil
}

// NewBridgeMediatorFilterer creates a new log filterer instance of BridgeMediator, bound to a specific deployed contract.
func NewBridgeMediatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeMediatorFilterer, error) {
	contract, err := bindBridgeMediator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorFilterer{contract: contract}, nil
}

// bindBridgeMediator binds a generic wrapper to an already deployed contract.
func bindBridgeMediator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeMediatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeMediator *BridgeMediatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeMediator.Contract.BridgeMediatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeMediator *BridgeMediatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeMediator.Contract.BridgeMediatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeMediator *BridgeMediatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeMediator.Contract.BridgeMediatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeMediator *BridgeMediatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeMediator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeMediator *BridgeMediatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeMediator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeMediator *BridgeMediatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeMediator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeMediator *BridgeMediatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeMediator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeMediator *BridgeMediatorSession) Owner() (common.Address, error) {
	return _BridgeMediator.Contract.Owner(&_BridgeMediator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeMediator *BridgeMediatorCallerSession) Owner() (common.Address, error) {
	return _BridgeMediator.Contract.Owner(&_BridgeMediator.CallOpts)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa20126ff.
//
// Solidity: function symbolToToken(uint256 , string ) view returns(address)
func (_BridgeMediator *BridgeMediatorCaller) SymbolToToken(opts *bind.CallOpts, arg0 *big.Int, arg1 string) (common.Address, error) {
	var out []interface{}
	err := _BridgeMediator.contract.Call(opts, &out, "symbolToToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa20126ff.
//
// Solidity: function symbolToToken(uint256 , string ) view returns(address)
func (_BridgeMediator *BridgeMediatorSession) SymbolToToken(arg0 *big.Int, arg1 string) (common.Address, error) {
	return _BridgeMediator.Contract.SymbolToToken(&_BridgeMediator.CallOpts, arg0, arg1)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa20126ff.
//
// Solidity: function symbolToToken(uint256 , string ) view returns(address)
func (_BridgeMediator *BridgeMediatorCallerSession) SymbolToToken(arg0 *big.Int, arg1 string) (common.Address, error) {
	return _BridgeMediator.Contract.SymbolToToken(&_BridgeMediator.CallOpts, arg0, arg1)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0xeff7bb86.
//
// Solidity: function tokenToSymbol(uint256 , address ) view returns(string)
func (_BridgeMediator *BridgeMediatorCaller) TokenToSymbol(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (string, error) {
	var out []interface{}
	err := _BridgeMediator.contract.Call(opts, &out, "tokenToSymbol", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0xeff7bb86.
//
// Solidity: function tokenToSymbol(uint256 , address ) view returns(string)
func (_BridgeMediator *BridgeMediatorSession) TokenToSymbol(arg0 *big.Int, arg1 common.Address) (string, error) {
	return _BridgeMediator.Contract.TokenToSymbol(&_BridgeMediator.CallOpts, arg0, arg1)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0xeff7bb86.
//
// Solidity: function tokenToSymbol(uint256 , address ) view returns(string)
func (_BridgeMediator *BridgeMediatorCallerSession) TokenToSymbol(arg0 *big.Int, arg1 common.Address) (string, error) {
	return _BridgeMediator.Contract.TokenToSymbol(&_BridgeMediator.CallOpts, arg0, arg1)
}

// AddToken is a paid mutator transaction binding the contract method 0x11457604.
//
// Solidity: function addToken(string symbol, uint256 chainId, address token) returns()
func (_BridgeMediator *BridgeMediatorTransactor) AddToken(opts *bind.TransactOpts, symbol string, chainId *big.Int, token common.Address) (*types.Transaction, error) {
	return _BridgeMediator.contract.Transact(opts, "addToken", symbol, chainId, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x11457604.
//
// Solidity: function addToken(string symbol, uint256 chainId, address token) returns()
func (_BridgeMediator *BridgeMediatorSession) AddToken(symbol string, chainId *big.Int, token common.Address) (*types.Transaction, error) {
	return _BridgeMediator.Contract.AddToken(&_BridgeMediator.TransactOpts, symbol, chainId, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x11457604.
//
// Solidity: function addToken(string symbol, uint256 chainId, address token) returns()
func (_BridgeMediator *BridgeMediatorTransactorSession) AddToken(symbol string, chainId *big.Int, token common.Address) (*types.Transaction, error) {
	return _BridgeMediator.Contract.AddToken(&_BridgeMediator.TransactOpts, symbol, chainId, token)
}

// Mediate is a paid mutator transaction binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 sourceChain, uint256 destinationChain, bytes sourceData) returns(bytes)
func (_BridgeMediator *BridgeMediatorTransactor) Mediate(opts *bind.TransactOpts, sourceChain *big.Int, destinationChain *big.Int, sourceData []byte) (*types.Transaction, error) {
	return _BridgeMediator.contract.Transact(opts, "mediate", sourceChain, destinationChain, sourceData)
}

// Mediate is a paid mutator transaction binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 sourceChain, uint256 destinationChain, bytes sourceData) returns(bytes)
func (_BridgeMediator *BridgeMediatorSession) Mediate(sourceChain *big.Int, destinationChain *big.Int, sourceData []byte) (*types.Transaction, error) {
	return _BridgeMediator.Contract.Mediate(&_BridgeMediator.TransactOpts, sourceChain, destinationChain, sourceData)
}

// Mediate is a paid mutator transaction binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 sourceChain, uint256 destinationChain, bytes sourceData) returns(bytes)
func (_BridgeMediator *BridgeMediatorTransactorSession) Mediate(sourceChain *big.Int, destinationChain *big.Int, sourceData []byte) (*types.Transaction, error) {
	return _BridgeMediator.Contract.Mediate(&_BridgeMediator.TransactOpts, sourceChain, destinationChain, sourceData)
}

// RemoveToken is a paid mutator transaction binding the contract method 0xe409ed98.
//
// Solidity: function removeToken(string symbol, uint256 chainId) returns()
func (_BridgeMediator *BridgeMediatorTransactor) RemoveToken(opts *bind.TransactOpts, symbol string, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeMediator.contract.Transact(opts, "removeToken", symbol, chainId)
}

// RemoveToken is a paid mutator transaction binding the contract method 0xe409ed98.
//
// Solidity: function removeToken(string symbol, uint256 chainId) returns()
func (_BridgeMediator *BridgeMediatorSession) RemoveToken(symbol string, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeMediator.Contract.RemoveToken(&_BridgeMediator.TransactOpts, symbol, chainId)
}

// RemoveToken is a paid mutator transaction binding the contract method 0xe409ed98.
//
// Solidity: function removeToken(string symbol, uint256 chainId) returns()
func (_BridgeMediator *BridgeMediatorTransactorSession) RemoveToken(symbol string, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeMediator.Contract.RemoveToken(&_BridgeMediator.TransactOpts, symbol, chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeMediator *BridgeMediatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeMediator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeMediator *BridgeMediatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeMediator.Contract.RenounceOwnership(&_BridgeMediator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeMediator *BridgeMediatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeMediator.Contract.RenounceOwnership(&_BridgeMediator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeMediator *BridgeMediatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeMediator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeMediator *BridgeMediatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeMediator.Contract.TransferOwnership(&_BridgeMediator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeMediator *BridgeMediatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeMediator.Contract.TransferOwnership(&_BridgeMediator.TransactOpts, newOwner)
}

// BridgeMediatorAddedTokenIterator is returned from FilterAddedToken and is used to iterate over the raw logs and unpacked data for AddedToken events raised by the BridgeMediator contract.
type BridgeMediatorAddedTokenIterator struct {
	Event *BridgeMediatorAddedToken // Event containing the contract specifics and raw log

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
func (it *BridgeMediatorAddedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMediatorAddedToken)
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
		it.Event = new(BridgeMediatorAddedToken)
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
func (it *BridgeMediatorAddedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMediatorAddedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMediatorAddedToken represents a AddedToken event raised by the BridgeMediator contract.
type BridgeMediatorAddedToken struct {
	Symbol  string
	ChainId *big.Int
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0x1eea762d9514744d01160b932a7a6c1cbee2788c286607b3cc00999f6a530ea9.
//
// Solidity: event AddedToken(string symbol, uint256 chainId, address token)
func (_BridgeMediator *BridgeMediatorFilterer) FilterAddedToken(opts *bind.FilterOpts) (*BridgeMediatorAddedTokenIterator, error) {

	logs, sub, err := _BridgeMediator.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorAddedTokenIterator{contract: _BridgeMediator.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0x1eea762d9514744d01160b932a7a6c1cbee2788c286607b3cc00999f6a530ea9.
//
// Solidity: event AddedToken(string symbol, uint256 chainId, address token)
func (_BridgeMediator *BridgeMediatorFilterer) WatchAddedToken(opts *bind.WatchOpts, sink chan<- *BridgeMediatorAddedToken) (event.Subscription, error) {

	logs, sub, err := _BridgeMediator.contract.WatchLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMediatorAddedToken)
				if err := _BridgeMediator.contract.UnpackLog(event, "AddedToken", log); err != nil {
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

// ParseAddedToken is a log parse operation binding the contract event 0x1eea762d9514744d01160b932a7a6c1cbee2788c286607b3cc00999f6a530ea9.
//
// Solidity: event AddedToken(string symbol, uint256 chainId, address token)
func (_BridgeMediator *BridgeMediatorFilterer) ParseAddedToken(log types.Log) (*BridgeMediatorAddedToken, error) {
	event := new(BridgeMediatorAddedToken)
	if err := _BridgeMediator.contract.UnpackLog(event, "AddedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMediatorMediatorAddressIterator is returned from FilterMediatorAddress and is used to iterate over the raw logs and unpacked data for MediatorAddress events raised by the BridgeMediator contract.
type BridgeMediatorMediatorAddressIterator struct {
	Event *BridgeMediatorMediatorAddress // Event containing the contract specifics and raw log

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
func (it *BridgeMediatorMediatorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMediatorMediatorAddress)
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
		it.Event = new(BridgeMediatorMediatorAddress)
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
func (it *BridgeMediatorMediatorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMediatorMediatorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMediatorMediatorAddress represents a MediatorAddress event raised by the BridgeMediator contract.
type BridgeMediatorMediatorAddress struct {
	MediatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMediatorAddress is a free log retrieval operation binding the contract event 0xbb7882191c73e87a43b5ce5731ce01aceb048e6cafc123d2122cc3d365ad7c31.
//
// Solidity: event MediatorAddress(address mediatorAddress)
func (_BridgeMediator *BridgeMediatorFilterer) FilterMediatorAddress(opts *bind.FilterOpts) (*BridgeMediatorMediatorAddressIterator, error) {

	logs, sub, err := _BridgeMediator.contract.FilterLogs(opts, "MediatorAddress")
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorMediatorAddressIterator{contract: _BridgeMediator.contract, event: "MediatorAddress", logs: logs, sub: sub}, nil
}

// WatchMediatorAddress is a free log subscription operation binding the contract event 0xbb7882191c73e87a43b5ce5731ce01aceb048e6cafc123d2122cc3d365ad7c31.
//
// Solidity: event MediatorAddress(address mediatorAddress)
func (_BridgeMediator *BridgeMediatorFilterer) WatchMediatorAddress(opts *bind.WatchOpts, sink chan<- *BridgeMediatorMediatorAddress) (event.Subscription, error) {

	logs, sub, err := _BridgeMediator.contract.WatchLogs(opts, "MediatorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMediatorMediatorAddress)
				if err := _BridgeMediator.contract.UnpackLog(event, "MediatorAddress", log); err != nil {
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

// ParseMediatorAddress is a log parse operation binding the contract event 0xbb7882191c73e87a43b5ce5731ce01aceb048e6cafc123d2122cc3d365ad7c31.
//
// Solidity: event MediatorAddress(address mediatorAddress)
func (_BridgeMediator *BridgeMediatorFilterer) ParseMediatorAddress(log types.Log) (*BridgeMediatorMediatorAddress, error) {
	event := new(BridgeMediatorMediatorAddress)
	if err := _BridgeMediator.contract.UnpackLog(event, "MediatorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMediatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeMediator contract.
type BridgeMediatorOwnershipTransferredIterator struct {
	Event *BridgeMediatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeMediatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMediatorOwnershipTransferred)
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
		it.Event = new(BridgeMediatorOwnershipTransferred)
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
func (it *BridgeMediatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMediatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMediatorOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeMediator contract.
type BridgeMediatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeMediator *BridgeMediatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeMediatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeMediator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorOwnershipTransferredIterator{contract: _BridgeMediator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeMediator *BridgeMediatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeMediatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeMediator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMediatorOwnershipTransferred)
				if err := _BridgeMediator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeMediator *BridgeMediatorFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeMediatorOwnershipTransferred, error) {
	event := new(BridgeMediatorOwnershipTransferred)
	if err := _BridgeMediator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMediatorRemovedTokenIterator is returned from FilterRemovedToken and is used to iterate over the raw logs and unpacked data for RemovedToken events raised by the BridgeMediator contract.
type BridgeMediatorRemovedTokenIterator struct {
	Event *BridgeMediatorRemovedToken // Event containing the contract specifics and raw log

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
func (it *BridgeMediatorRemovedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMediatorRemovedToken)
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
		it.Event = new(BridgeMediatorRemovedToken)
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
func (it *BridgeMediatorRemovedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMediatorRemovedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMediatorRemovedToken represents a RemovedToken event raised by the BridgeMediator contract.
type BridgeMediatorRemovedToken struct {
	Symbol  string
	ChainId *big.Int
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedToken is a free log retrieval operation binding the contract event 0x747510f5d47eee15c8194e6944a1339805da89ef768309e849b2aacde53f3b44.
//
// Solidity: event RemovedToken(string symbol, uint256 chainId, address token)
func (_BridgeMediator *BridgeMediatorFilterer) FilterRemovedToken(opts *bind.FilterOpts) (*BridgeMediatorRemovedTokenIterator, error) {

	logs, sub, err := _BridgeMediator.contract.FilterLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return &BridgeMediatorRemovedTokenIterator{contract: _BridgeMediator.contract, event: "RemovedToken", logs: logs, sub: sub}, nil
}

// WatchRemovedToken is a free log subscription operation binding the contract event 0x747510f5d47eee15c8194e6944a1339805da89ef768309e849b2aacde53f3b44.
//
// Solidity: event RemovedToken(string symbol, uint256 chainId, address token)
func (_BridgeMediator *BridgeMediatorFilterer) WatchRemovedToken(opts *bind.WatchOpts, sink chan<- *BridgeMediatorRemovedToken) (event.Subscription, error) {

	logs, sub, err := _BridgeMediator.contract.WatchLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMediatorRemovedToken)
				if err := _BridgeMediator.contract.UnpackLog(event, "RemovedToken", log); err != nil {
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

// ParseRemovedToken is a log parse operation binding the contract event 0x747510f5d47eee15c8194e6944a1339805da89ef768309e849b2aacde53f3b44.
//
// Solidity: event RemovedToken(string symbol, uint256 chainId, address token)
func (_BridgeMediator *BridgeMediatorFilterer) ParseRemovedToken(log types.Log) (*BridgeMediatorRemovedToken, error) {
	event := new(BridgeMediatorRemovedToken)
	if err := _BridgeMediator.contract.UnpackLog(event, "RemovedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
