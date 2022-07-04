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

// ValidatorManagerMetaData contains all meta data concerning the ValidatorManager contract.
var ValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredApprovals\",\"type\":\"uint256\"}],\"name\":\"RequiredApprovalsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredApprovals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requiredApprovals\",\"type\":\"uint256\"}],\"name\":\"setRequiredApprovals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorManagerMetaData.ABI instead.
var ValidatorManagerABI = ValidatorManagerMetaData.ABI

// ValidatorManager is an auto generated Go binding around an Ethereum contract.
type ValidatorManager struct {
	ValidatorManagerCaller     // Read-only binding to the contract
	ValidatorManagerTransactor // Write-only binding to the contract
	ValidatorManagerFilterer   // Log filterer for contract events
}

// ValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorManagerSession struct {
	Contract     *ValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorManagerCallerSession struct {
	Contract *ValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorManagerTransactorSession struct {
	Contract     *ValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorManagerRaw struct {
	Contract *ValidatorManager // Generic contract binding to access the raw methods on
}

// ValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorManagerCallerRaw struct {
	Contract *ValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactorRaw struct {
	Contract *ValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorManager creates a new instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManager(address common.Address, backend bind.ContractBackend) (*ValidatorManager, error) {
	contract, err := bindValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}, ValidatorManagerFilterer: ValidatorManagerFilterer{contract: contract}}, nil
}

// NewValidatorManagerCaller creates a new read-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*ValidatorManagerCaller, error) {
	contract, err := bindValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerCaller{contract: contract}, nil
}

// NewValidatorManagerTransactor creates a new write-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorManagerTransactor, error) {
	contract, err := bindValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerTransactor{contract: contract}, nil
}

// NewValidatorManagerFilterer creates a new log filterer instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorManagerFilterer, error) {
	contract, err := bindValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerFilterer{contract: contract}, nil
}

// bindValidatorManager binds a generic wrapper to an already deployed contract.
func bindValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.ValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_ValidatorManager *ValidatorManagerCaller) IsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "isValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_ValidatorManager *ValidatorManagerSession) IsValidator(arg0 common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsValidator(&_ValidatorManager.CallOpts, arg0)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_ValidatorManager *ValidatorManagerCallerSession) IsValidator(arg0 common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsValidator(&_ValidatorManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) Owner() (common.Address, error) {
	return _ValidatorManager.Contract.Owner(&_ValidatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) Owner() (common.Address, error) {
	return _ValidatorManager.Contract.Owner(&_ValidatorManager.CallOpts)
}

// RequiredApprovals is a free data retrieval call binding the contract method 0x99c1aadc.
//
// Solidity: function requiredApprovals() view returns(uint256)
func (_ValidatorManager *ValidatorManagerCaller) RequiredApprovals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "requiredApprovals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredApprovals is a free data retrieval call binding the contract method 0x99c1aadc.
//
// Solidity: function requiredApprovals() view returns(uint256)
func (_ValidatorManager *ValidatorManagerSession) RequiredApprovals() (*big.Int, error) {
	return _ValidatorManager.Contract.RequiredApprovals(&_ValidatorManager.CallOpts)
}

// RequiredApprovals is a free data retrieval call binding the contract method 0x99c1aadc.
//
// Solidity: function requiredApprovals() view returns(uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) RequiredApprovals() (*big.Int, error) {
	return _ValidatorManager.Contract.RequiredApprovals(&_ValidatorManager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorManager *ValidatorManagerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorManager.Contract.Validators(&_ValidatorManager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorManager.Contract.Validators(&_ValidatorManager.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorManager *ValidatorManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorManager *ValidatorManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorManager.Contract.RenounceOwnership(&_ValidatorManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorManager.Contract.RenounceOwnership(&_ValidatorManager.TransactOpts)
}

// SetRequiredApprovals is a paid mutator transaction binding the contract method 0x222a242e.
//
// Solidity: function setRequiredApprovals(uint256 _requiredApprovals) returns()
func (_ValidatorManager *ValidatorManagerTransactor) SetRequiredApprovals(opts *bind.TransactOpts, _requiredApprovals *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "setRequiredApprovals", _requiredApprovals)
}

// SetRequiredApprovals is a paid mutator transaction binding the contract method 0x222a242e.
//
// Solidity: function setRequiredApprovals(uint256 _requiredApprovals) returns()
func (_ValidatorManager *ValidatorManagerSession) SetRequiredApprovals(_requiredApprovals *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetRequiredApprovals(&_ValidatorManager.TransactOpts, _requiredApprovals)
}

// SetRequiredApprovals is a paid mutator transaction binding the contract method 0x222a242e.
//
// Solidity: function setRequiredApprovals(uint256 _requiredApprovals) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) SetRequiredApprovals(_requiredApprovals *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetRequiredApprovals(&_ValidatorManager.TransactOpts, _requiredApprovals)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_ValidatorManager *ValidatorManagerTransactor) SetValidators(opts *bind.TransactOpts, _validators []common.Address) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "setValidators", _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_ValidatorManager *ValidatorManagerSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetValidators(&_ValidatorManager.TransactOpts, _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetValidators(&_ValidatorManager.TransactOpts, _validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorManager *ValidatorManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorManager *ValidatorManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TransferOwnership(&_ValidatorManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TransferOwnership(&_ValidatorManager.TransactOpts, newOwner)
}

// ValidatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorManager contract.
type ValidatorManagerInitializedIterator struct {
	Event *ValidatorManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerInitialized)
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
		it.Event = new(ValidatorManagerInitialized)
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
func (it *ValidatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerInitialized represents a Initialized event raised by the ValidatorManager contract.
type ValidatorManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorManager *ValidatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorManagerInitializedIterator, error) {

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerInitializedIterator{contract: _ValidatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorManager *ValidatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerInitialized)
				if err := _ValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ValidatorManager *ValidatorManagerFilterer) ParseInitialized(log types.Log) (*ValidatorManagerInitialized, error) {
	event := new(ValidatorManagerInitialized)
	if err := _ValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorManager contract.
type ValidatorManagerOwnershipTransferredIterator struct {
	Event *ValidatorManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerOwnershipTransferred)
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
		it.Event = new(ValidatorManagerOwnershipTransferred)
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
func (it *ValidatorManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorManager contract.
type ValidatorManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorManager *ValidatorManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerOwnershipTransferredIterator{contract: _ValidatorManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorManager *ValidatorManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerOwnershipTransferred)
				if err := _ValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValidatorManager *ValidatorManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorManagerOwnershipTransferred, error) {
	event := new(ValidatorManagerOwnershipTransferred)
	if err := _ValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerRequiredApprovalsUpdatedIterator is returned from FilterRequiredApprovalsUpdated and is used to iterate over the raw logs and unpacked data for RequiredApprovalsUpdated events raised by the ValidatorManager contract.
type ValidatorManagerRequiredApprovalsUpdatedIterator struct {
	Event *ValidatorManagerRequiredApprovalsUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerRequiredApprovalsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerRequiredApprovalsUpdated)
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
		it.Event = new(ValidatorManagerRequiredApprovalsUpdated)
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
func (it *ValidatorManagerRequiredApprovalsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerRequiredApprovalsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerRequiredApprovalsUpdated represents a RequiredApprovalsUpdated event raised by the ValidatorManager contract.
type ValidatorManagerRequiredApprovalsUpdated struct {
	RequiredApprovals *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRequiredApprovalsUpdated is a free log retrieval operation binding the contract event 0xdb445dd03a560dac1b5553e3d3a8d6396138df6ab80174d52bfb898a57f42936.
//
// Solidity: event RequiredApprovalsUpdated(uint256 requiredApprovals)
func (_ValidatorManager *ValidatorManagerFilterer) FilterRequiredApprovalsUpdated(opts *bind.FilterOpts) (*ValidatorManagerRequiredApprovalsUpdatedIterator, error) {

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "RequiredApprovalsUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerRequiredApprovalsUpdatedIterator{contract: _ValidatorManager.contract, event: "RequiredApprovalsUpdated", logs: logs, sub: sub}, nil
}

// WatchRequiredApprovalsUpdated is a free log subscription operation binding the contract event 0xdb445dd03a560dac1b5553e3d3a8d6396138df6ab80174d52bfb898a57f42936.
//
// Solidity: event RequiredApprovalsUpdated(uint256 requiredApprovals)
func (_ValidatorManager *ValidatorManagerFilterer) WatchRequiredApprovalsUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerRequiredApprovalsUpdated) (event.Subscription, error) {

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "RequiredApprovalsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerRequiredApprovalsUpdated)
				if err := _ValidatorManager.contract.UnpackLog(event, "RequiredApprovalsUpdated", log); err != nil {
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

// ParseRequiredApprovalsUpdated is a log parse operation binding the contract event 0xdb445dd03a560dac1b5553e3d3a8d6396138df6ab80174d52bfb898a57f42936.
//
// Solidity: event RequiredApprovalsUpdated(uint256 requiredApprovals)
func (_ValidatorManager *ValidatorManagerFilterer) ParseRequiredApprovalsUpdated(log types.Log) (*ValidatorManagerRequiredApprovalsUpdated, error) {
	event := new(ValidatorManagerRequiredApprovalsUpdated)
	if err := _ValidatorManager.contract.UnpackLog(event, "RequiredApprovalsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the ValidatorManager contract.
type ValidatorManagerValidatorsUpdatedIterator struct {
	Event *ValidatorManagerValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorsUpdated)
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
		it.Event = new(ValidatorManagerValidatorsUpdated)
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
func (it *ValidatorManagerValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorsUpdated represents a ValidatorsUpdated event raised by the ValidatorManager contract.
type ValidatorManagerValidatorsUpdated struct {
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0x095f1141a8102fc81879f247c8ea3c0aa8a60f19127b978c5f1397ecc245b40d.
//
// Solidity: event ValidatorsUpdated(address[] validators)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*ValidatorManagerValidatorsUpdatedIterator, error) {

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorsUpdatedIterator{contract: _ValidatorManager.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0x095f1141a8102fc81879f247c8ea3c0aa8a60f19127b978c5f1397ecc245b40d.
//
// Solidity: event ValidatorsUpdated(address[] validators)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorsUpdated)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
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

// ParseValidatorsUpdated is a log parse operation binding the contract event 0x095f1141a8102fc81879f247c8ea3c0aa8a60f19127b978c5f1397ecc245b40d.
//
// Solidity: event ValidatorsUpdated(address[] validators)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorsUpdated(log types.Log) (*ValidatorManagerValidatorsUpdated, error) {
	event := new(ValidatorManagerValidatorsUpdated)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
