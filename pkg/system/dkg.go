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

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"RoundDataFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"RoundDataProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getRoundBroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRoundBroadcastData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"isRoundFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_rawData\",\"type\":\"bytes\"}],\"name\":\"roundBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DKGABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGMetaData.ABI instead.
var DKGABI = DKGMetaData.ABI

// DKG is an auto generated Go binding around an Ethereum contract.
type DKG struct {
	DKGCaller     // Read-only binding to the contract
	DKGTransactor // Write-only binding to the contract
	DKGFilterer   // Log filterer for contract events
}

// DKGCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGSession struct {
	Contract     *DKG              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGCallerSession struct {
	Contract *DKGCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DKGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGTransactorSession struct {
	Contract     *DKGTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGRaw struct {
	Contract *DKG // Generic contract binding to access the raw methods on
}

// DKGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGCallerRaw struct {
	Contract *DKGCaller // Generic read-only contract binding to access the raw methods on
}

// DKGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGTransactorRaw struct {
	Contract *DKGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKG creates a new instance of DKG, bound to a specific deployed contract.
func NewDKG(address common.Address, backend bind.ContractBackend) (*DKG, error) {
	contract, err := bindDKG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

// NewDKGCaller creates a new read-only instance of DKG, bound to a specific deployed contract.
func NewDKGCaller(address common.Address, caller bind.ContractCaller) (*DKGCaller, error) {
	contract, err := bindDKG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGCaller{contract: contract}, nil
}

// NewDKGTransactor creates a new write-only instance of DKG, bound to a specific deployed contract.
func NewDKGTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGTransactor, error) {
	contract, err := bindDKG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGTransactor{contract: contract}, nil
}

// NewDKGFilterer creates a new log filterer instance of DKG, bound to a specific deployed contract.
func NewDKGFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGFilterer, error) {
	contract, err := bindDKG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGFilterer{contract: contract}, nil
}

// bindDKG binds a generic wrapper to an already deployed contract.
func bindDKG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DKGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.DKGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGCaller) GetCurrentValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getCurrentValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGSession) GetCurrentValidators() ([]common.Address, error) {
	return _DKG.Contract.GetCurrentValidators(&_DKG.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGCallerSession) GetCurrentValidators() ([]common.Address, error) {
	return _DKG.Contract.GetCurrentValidators(&_DKG.CallOpts)
}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _id, uint256 _round) view returns(uint256)
func (_DKG *DKGCaller) GetRoundBroadcastCount(opts *bind.CallOpts, _id *big.Int, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRoundBroadcastCount", _id, _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _id, uint256 _round) view returns(uint256)
func (_DKG *DKGSession) GetRoundBroadcastCount(_id *big.Int, _round *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetRoundBroadcastCount(&_DKG.CallOpts, _id, _round)
}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _id, uint256 _round) view returns(uint256)
func (_DKG *DKGCallerSession) GetRoundBroadcastCount(_id *big.Int, _round *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetRoundBroadcastCount(&_DKG.CallOpts, _id, _round)
}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _id, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGCaller) GetRoundBroadcastData(opts *bind.CallOpts, _id *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRoundBroadcastData", _id, _round, _validator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _id, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGSession) GetRoundBroadcastData(_id *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	return _DKG.Contract.GetRoundBroadcastData(&_DKG.CallOpts, _id, _round, _validator)
}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _id, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGCallerSession) GetRoundBroadcastData(_id *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	return _DKG.Contract.GetRoundBroadcastData(&_DKG.CallOpts, _id, _round, _validator)
}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _id) view returns(address[])
func (_DKG *DKGCaller) GetValidators(opts *bind.CallOpts, _id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidators", _id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _id) view returns(address[])
func (_DKG *DKGSession) GetValidators(_id *big.Int) ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts, _id)
}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _id) view returns(address[])
func (_DKG *DKGCallerSession) GetValidators(_id *big.Int) ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts, _id)
}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _id, uint256 _round) view returns(bool)
func (_DKG *DKGCaller) IsRoundFilled(opts *bind.CallOpts, _id *big.Int, _round *big.Int) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isRoundFilled", _id, _round)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _id, uint256 _round) view returns(bool)
func (_DKG *DKGSession) IsRoundFilled(_id *big.Int, _round *big.Int) (bool, error) {
	return _DKG.Contract.IsRoundFilled(&_DKG.CallOpts, _id, _round)
}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _id, uint256 _round) view returns(bool)
func (_DKG *DKGCallerSession) IsRoundFilled(_id *big.Int, _round *big.Int) (bool, error) {
	return _DKG.Contract.IsRoundFilled(&_DKG.CallOpts, _id, _round)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_DKG *DKGCaller) IsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_DKG *DKGSession) IsValidator(arg0 common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, arg0)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_DKG *DKGCallerSession) IsValidator(arg0 common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DKG *DKGCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DKG *DKGSession) Owner() (common.Address, error) {
	return _DKG.Contract.Owner(&_DKG.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DKG *DKGCallerSession) Owner() (common.Address, error) {
	return _DKG.Contract.Owner(&_DKG.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xdcf2793a.
//
// Solidity: function validators(uint256 , uint256 ) view returns(address)
func (_DKG *DKGCaller) Validators(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "validators", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xdcf2793a.
//
// Solidity: function validators(uint256 , uint256 ) view returns(address)
func (_DKG *DKGSession) Validators(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _DKG.Contract.Validators(&_DKG.CallOpts, arg0, arg1)
}

// Validators is a free data retrieval call binding the contract method 0xdcf2793a.
//
// Solidity: function validators(uint256 , uint256 ) view returns(address)
func (_DKG *DKGCallerSession) Validators(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _DKG.Contract.Validators(&_DKG.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _validators) returns()
func (_DKG *DKGTransactor) Initialize(opts *bind.TransactOpts, _validators []common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "initialize", _validators)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _validators) returns()
func (_DKG *DKGSession) Initialize(_validators []common.Address) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _validators)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _validators) returns()
func (_DKG *DKGTransactorSession) Initialize(_validators []common.Address) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DKG *DKGTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DKG *DKGSession) RenounceOwnership() (*types.Transaction, error) {
	return _DKG.Contract.RenounceOwnership(&_DKG.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DKG *DKGTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DKG.Contract.RenounceOwnership(&_DKG.TransactOpts)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _id, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGTransactor) RoundBroadcast(opts *bind.TransactOpts, _id *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "roundBroadcast", _id, _round, _rawData)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _id, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGSession) RoundBroadcast(_id *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.Contract.RoundBroadcast(&_DKG.TransactOpts, _id, _round, _rawData)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _id, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGTransactorSession) RoundBroadcast(_id *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.Contract.RoundBroadcast(&_DKG.TransactOpts, _id, _round, _rawData)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_DKG *DKGTransactor) SetValidators(opts *bind.TransactOpts, _validators []common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "setValidators", _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_DKG *DKGSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _DKG.Contract.SetValidators(&_DKG.TransactOpts, _validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] _validators) returns()
func (_DKG *DKGTransactorSession) SetValidators(_validators []common.Address) (*types.Transaction, error) {
	return _DKG.Contract.SetValidators(&_DKG.TransactOpts, _validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DKG *DKGTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DKG *DKGSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DKG.Contract.TransferOwnership(&_DKG.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DKG *DKGTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DKG.Contract.TransferOwnership(&_DKG.TransactOpts, newOwner)
}

// DKGOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DKG contract.
type DKGOwnershipTransferredIterator struct {
	Event *DKGOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DKGOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGOwnershipTransferred)
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
		it.Event = new(DKGOwnershipTransferred)
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
func (it *DKGOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGOwnershipTransferred represents a OwnershipTransferred event raised by the DKG contract.
type DKGOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DKG *DKGFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DKGOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DKGOwnershipTransferredIterator{contract: _DKG.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DKG *DKGFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DKGOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGOwnershipTransferred)
				if err := _DKG.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DKG *DKGFilterer) ParseOwnershipTransferred(log types.Log) (*DKGOwnershipTransferred, error) {
	event := new(DKGOwnershipTransferred)
	if err := _DKG.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRoundDataFilledIterator is returned from FilterRoundDataFilled and is used to iterate over the raw logs and unpacked data for RoundDataFilled events raised by the DKG contract.
type DKGRoundDataFilledIterator struct {
	Event *DKGRoundDataFilled // Event containing the contract specifics and raw log

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
func (it *DKGRoundDataFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRoundDataFilled)
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
		it.Event = new(DKGRoundDataFilled)
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
func (it *DKGRoundDataFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRoundDataFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRoundDataFilled represents a RoundDataFilled event raised by the DKG contract.
type DKGRoundDataFilled struct {
	Id    *big.Int
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRoundDataFilled is a free log retrieval operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 id, uint256 round)
func (_DKG *DKGFilterer) FilterRoundDataFilled(opts *bind.FilterOpts) (*DKGRoundDataFilledIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "RoundDataFilled")
	if err != nil {
		return nil, err
	}
	return &DKGRoundDataFilledIterator{contract: _DKG.contract, event: "RoundDataFilled", logs: logs, sub: sub}, nil
}

// WatchRoundDataFilled is a free log subscription operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 id, uint256 round)
func (_DKG *DKGFilterer) WatchRoundDataFilled(opts *bind.WatchOpts, sink chan<- *DKGRoundDataFilled) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "RoundDataFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRoundDataFilled)
				if err := _DKG.contract.UnpackLog(event, "RoundDataFilled", log); err != nil {
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

// ParseRoundDataFilled is a log parse operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 id, uint256 round)
func (_DKG *DKGFilterer) ParseRoundDataFilled(log types.Log) (*DKGRoundDataFilled, error) {
	event := new(DKGRoundDataFilled)
	if err := _DKG.contract.UnpackLog(event, "RoundDataFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRoundDataProvidedIterator is returned from FilterRoundDataProvided and is used to iterate over the raw logs and unpacked data for RoundDataProvided events raised by the DKG contract.
type DKGRoundDataProvidedIterator struct {
	Event *DKGRoundDataProvided // Event containing the contract specifics and raw log

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
func (it *DKGRoundDataProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRoundDataProvided)
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
		it.Event = new(DKGRoundDataProvided)
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
func (it *DKGRoundDataProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRoundDataProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRoundDataProvided represents a RoundDataProvided event raised by the DKG contract.
type DKGRoundDataProvided struct {
	Id        *big.Int
	Round     *big.Int
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRoundDataProvided is a free log retrieval operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 id, uint256 round, address validator)
func (_DKG *DKGFilterer) FilterRoundDataProvided(opts *bind.FilterOpts) (*DKGRoundDataProvidedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "RoundDataProvided")
	if err != nil {
		return nil, err
	}
	return &DKGRoundDataProvidedIterator{contract: _DKG.contract, event: "RoundDataProvided", logs: logs, sub: sub}, nil
}

// WatchRoundDataProvided is a free log subscription operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 id, uint256 round, address validator)
func (_DKG *DKGFilterer) WatchRoundDataProvided(opts *bind.WatchOpts, sink chan<- *DKGRoundDataProvided) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "RoundDataProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRoundDataProvided)
				if err := _DKG.contract.UnpackLog(event, "RoundDataProvided", log); err != nil {
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

// ParseRoundDataProvided is a log parse operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 id, uint256 round, address validator)
func (_DKG *DKGFilterer) ParseRoundDataProvided(log types.Log) (*DKGRoundDataProvided, error) {
	event := new(DKGRoundDataProvided)
	if err := _DKG.contract.UnpackLog(event, "RoundDataProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the DKG contract.
type DKGValidatorsUpdatedIterator struct {
	Event *DKGValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *DKGValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGValidatorsUpdated)
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
		it.Event = new(DKGValidatorsUpdated)
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
func (it *DKGValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGValidatorsUpdated represents a ValidatorsUpdated event raised by the DKG contract.
type DKGValidatorsUpdated struct {
	Id         *big.Int
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 id, address[] validators)
func (_DKG *DKGFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*DKGValidatorsUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGValidatorsUpdatedIterator{contract: _DKG.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 id, address[] validators)
func (_DKG *DKGFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *DKGValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGValidatorsUpdated)
				if err := _DKG.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
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

// ParseValidatorsUpdated is a log parse operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 id, address[] validators)
func (_DKG *DKGFilterer) ParseValidatorsUpdated(log types.Log) (*DKGValidatorsUpdated, error) {
	event := new(DKGValidatorsUpdated)
	if err := _DKG.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
