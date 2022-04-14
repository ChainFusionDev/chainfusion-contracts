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

// BroadcastData is an auto generated low-level Go binding around an user-defined struct.
type BroadcastData struct {
	Provided    bool
	PublicData  [][]byte
	PrivateData [][]byte
}

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Round1Filled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Round1Provided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Round2Filled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Round2Provided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Round3Filled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Round3Provided\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRound1BroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRound1BroadcastData\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"provided\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"publicData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"privateData\",\"type\":\"bytes[]\"}],\"internalType\":\"structBroadcastData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRound2BroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRound2BroadcastData\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"provided\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"publicData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"privateData\",\"type\":\"bytes[]\"}],\"internalType\":\"structBroadcastData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRound3BroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRound3BroadcastData\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"provided\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"publicData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"privateData\",\"type\":\"bytes[]\"}],\"internalType\":\"structBroadcastData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_public\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_private\",\"type\":\"bytes[]\"}],\"name\":\"round1Broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_public\",\"type\":\"bytes[]\"}],\"name\":\"round2Broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_public\",\"type\":\"bytes[]\"}],\"name\":\"round3Broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetRound1BroadcastCount is a free data retrieval call binding the contract method 0x2c1075c6.
//
// Solidity: function getRound1BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGCaller) GetRound1BroadcastCount(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRound1BroadcastCount", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound1BroadcastCount is a free data retrieval call binding the contract method 0x2c1075c6.
//
// Solidity: function getRound1BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGSession) GetRound1BroadcastCount(_id [32]byte) (*big.Int, error) {
	return _DKG.Contract.GetRound1BroadcastCount(&_DKG.CallOpts, _id)
}

// GetRound1BroadcastCount is a free data retrieval call binding the contract method 0x2c1075c6.
//
// Solidity: function getRound1BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGCallerSession) GetRound1BroadcastCount(_id [32]byte) (*big.Int, error) {
	return _DKG.Contract.GetRound1BroadcastCount(&_DKG.CallOpts, _id)
}

// GetRound1BroadcastData is a free data retrieval call binding the contract method 0xe0635b22.
//
// Solidity: function getRound1BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGCaller) GetRound1BroadcastData(opts *bind.CallOpts, _id [32]byte, _validator common.Address) (BroadcastData, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRound1BroadcastData", _id, _validator)

	if err != nil {
		return *new(BroadcastData), err
	}

	out0 := *abi.ConvertType(out[0], new(BroadcastData)).(*BroadcastData)

	return out0, err

}

// GetRound1BroadcastData is a free data retrieval call binding the contract method 0xe0635b22.
//
// Solidity: function getRound1BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGSession) GetRound1BroadcastData(_id [32]byte, _validator common.Address) (BroadcastData, error) {
	return _DKG.Contract.GetRound1BroadcastData(&_DKG.CallOpts, _id, _validator)
}

// GetRound1BroadcastData is a free data retrieval call binding the contract method 0xe0635b22.
//
// Solidity: function getRound1BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGCallerSession) GetRound1BroadcastData(_id [32]byte, _validator common.Address) (BroadcastData, error) {
	return _DKG.Contract.GetRound1BroadcastData(&_DKG.CallOpts, _id, _validator)
}

// GetRound2BroadcastCount is a free data retrieval call binding the contract method 0x5af41e3a.
//
// Solidity: function getRound2BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGCaller) GetRound2BroadcastCount(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRound2BroadcastCount", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound2BroadcastCount is a free data retrieval call binding the contract method 0x5af41e3a.
//
// Solidity: function getRound2BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGSession) GetRound2BroadcastCount(_id [32]byte) (*big.Int, error) {
	return _DKG.Contract.GetRound2BroadcastCount(&_DKG.CallOpts, _id)
}

// GetRound2BroadcastCount is a free data retrieval call binding the contract method 0x5af41e3a.
//
// Solidity: function getRound2BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGCallerSession) GetRound2BroadcastCount(_id [32]byte) (*big.Int, error) {
	return _DKG.Contract.GetRound2BroadcastCount(&_DKG.CallOpts, _id)
}

// GetRound2BroadcastData is a free data retrieval call binding the contract method 0x6954add5.
//
// Solidity: function getRound2BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGCaller) GetRound2BroadcastData(opts *bind.CallOpts, _id [32]byte, _validator common.Address) (BroadcastData, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRound2BroadcastData", _id, _validator)

	if err != nil {
		return *new(BroadcastData), err
	}

	out0 := *abi.ConvertType(out[0], new(BroadcastData)).(*BroadcastData)

	return out0, err

}

// GetRound2BroadcastData is a free data retrieval call binding the contract method 0x6954add5.
//
// Solidity: function getRound2BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGSession) GetRound2BroadcastData(_id [32]byte, _validator common.Address) (BroadcastData, error) {
	return _DKG.Contract.GetRound2BroadcastData(&_DKG.CallOpts, _id, _validator)
}

// GetRound2BroadcastData is a free data retrieval call binding the contract method 0x6954add5.
//
// Solidity: function getRound2BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGCallerSession) GetRound2BroadcastData(_id [32]byte, _validator common.Address) (BroadcastData, error) {
	return _DKG.Contract.GetRound2BroadcastData(&_DKG.CallOpts, _id, _validator)
}

// GetRound3BroadcastCount is a free data retrieval call binding the contract method 0xf0a7e833.
//
// Solidity: function getRound3BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGCaller) GetRound3BroadcastCount(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRound3BroadcastCount", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound3BroadcastCount is a free data retrieval call binding the contract method 0xf0a7e833.
//
// Solidity: function getRound3BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGSession) GetRound3BroadcastCount(_id [32]byte) (*big.Int, error) {
	return _DKG.Contract.GetRound3BroadcastCount(&_DKG.CallOpts, _id)
}

// GetRound3BroadcastCount is a free data retrieval call binding the contract method 0xf0a7e833.
//
// Solidity: function getRound3BroadcastCount(bytes32 _id) view returns(uint256)
func (_DKG *DKGCallerSession) GetRound3BroadcastCount(_id [32]byte) (*big.Int, error) {
	return _DKG.Contract.GetRound3BroadcastCount(&_DKG.CallOpts, _id)
}

// GetRound3BroadcastData is a free data retrieval call binding the contract method 0x560e6a08.
//
// Solidity: function getRound3BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGCaller) GetRound3BroadcastData(opts *bind.CallOpts, _id [32]byte, _validator common.Address) (BroadcastData, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRound3BroadcastData", _id, _validator)

	if err != nil {
		return *new(BroadcastData), err
	}

	out0 := *abi.ConvertType(out[0], new(BroadcastData)).(*BroadcastData)

	return out0, err

}

// GetRound3BroadcastData is a free data retrieval call binding the contract method 0x560e6a08.
//
// Solidity: function getRound3BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGSession) GetRound3BroadcastData(_id [32]byte, _validator common.Address) (BroadcastData, error) {
	return _DKG.Contract.GetRound3BroadcastData(&_DKG.CallOpts, _id, _validator)
}

// GetRound3BroadcastData is a free data retrieval call binding the contract method 0x560e6a08.
//
// Solidity: function getRound3BroadcastData(bytes32 _id, address _validator) view returns((bool,bytes[],bytes[]))
func (_DKG *DKGCallerSession) GetRound3BroadcastData(_id [32]byte, _validator common.Address) (BroadcastData, error) {
	return _DKG.Contract.GetRound3BroadcastData(&_DKG.CallOpts, _id, _validator)
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

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_DKG *DKGCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_DKG *DKGSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _DKG.Contract.Validators(&_DKG.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_DKG *DKGCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _DKG.Contract.Validators(&_DKG.CallOpts, arg0)
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

// Round1Broadcast is a paid mutator transaction binding the contract method 0x6f540e43.
//
// Solidity: function round1Broadcast(bytes32 _id, bytes[] _public, bytes[] _private) returns()
func (_DKG *DKGTransactor) Round1Broadcast(opts *bind.TransactOpts, _id [32]byte, _public [][]byte, _private [][]byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "round1Broadcast", _id, _public, _private)
}

// Round1Broadcast is a paid mutator transaction binding the contract method 0x6f540e43.
//
// Solidity: function round1Broadcast(bytes32 _id, bytes[] _public, bytes[] _private) returns()
func (_DKG *DKGSession) Round1Broadcast(_id [32]byte, _public [][]byte, _private [][]byte) (*types.Transaction, error) {
	return _DKG.Contract.Round1Broadcast(&_DKG.TransactOpts, _id, _public, _private)
}

// Round1Broadcast is a paid mutator transaction binding the contract method 0x6f540e43.
//
// Solidity: function round1Broadcast(bytes32 _id, bytes[] _public, bytes[] _private) returns()
func (_DKG *DKGTransactorSession) Round1Broadcast(_id [32]byte, _public [][]byte, _private [][]byte) (*types.Transaction, error) {
	return _DKG.Contract.Round1Broadcast(&_DKG.TransactOpts, _id, _public, _private)
}

// Round2Broadcast is a paid mutator transaction binding the contract method 0x4ee2c5ae.
//
// Solidity: function round2Broadcast(bytes32 _id, bytes[] _public) returns()
func (_DKG *DKGTransactor) Round2Broadcast(opts *bind.TransactOpts, _id [32]byte, _public [][]byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "round2Broadcast", _id, _public)
}

// Round2Broadcast is a paid mutator transaction binding the contract method 0x4ee2c5ae.
//
// Solidity: function round2Broadcast(bytes32 _id, bytes[] _public) returns()
func (_DKG *DKGSession) Round2Broadcast(_id [32]byte, _public [][]byte) (*types.Transaction, error) {
	return _DKG.Contract.Round2Broadcast(&_DKG.TransactOpts, _id, _public)
}

// Round2Broadcast is a paid mutator transaction binding the contract method 0x4ee2c5ae.
//
// Solidity: function round2Broadcast(bytes32 _id, bytes[] _public) returns()
func (_DKG *DKGTransactorSession) Round2Broadcast(_id [32]byte, _public [][]byte) (*types.Transaction, error) {
	return _DKG.Contract.Round2Broadcast(&_DKG.TransactOpts, _id, _public)
}

// Round3Broadcast is a paid mutator transaction binding the contract method 0x37f5d622.
//
// Solidity: function round3Broadcast(bytes32 _id, bytes[] _public) returns()
func (_DKG *DKGTransactor) Round3Broadcast(opts *bind.TransactOpts, _id [32]byte, _public [][]byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "round3Broadcast", _id, _public)
}

// Round3Broadcast is a paid mutator transaction binding the contract method 0x37f5d622.
//
// Solidity: function round3Broadcast(bytes32 _id, bytes[] _public) returns()
func (_DKG *DKGSession) Round3Broadcast(_id [32]byte, _public [][]byte) (*types.Transaction, error) {
	return _DKG.Contract.Round3Broadcast(&_DKG.TransactOpts, _id, _public)
}

// Round3Broadcast is a paid mutator transaction binding the contract method 0x37f5d622.
//
// Solidity: function round3Broadcast(bytes32 _id, bytes[] _public) returns()
func (_DKG *DKGTransactorSession) Round3Broadcast(_id [32]byte, _public [][]byte) (*types.Transaction, error) {
	return _DKG.Contract.Round3Broadcast(&_DKG.TransactOpts, _id, _public)
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

// DKGRound1FilledIterator is returned from FilterRound1Filled and is used to iterate over the raw logs and unpacked data for Round1Filled events raised by the DKG contract.
type DKGRound1FilledIterator struct {
	Event *DKGRound1Filled // Event containing the contract specifics and raw log

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
func (it *DKGRound1FilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRound1Filled)
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
		it.Event = new(DKGRound1Filled)
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
func (it *DKGRound1FilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRound1FilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRound1Filled represents a Round1Filled event raised by the DKG contract.
type DKGRound1Filled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRound1Filled is a free log retrieval operation binding the contract event 0x826f5f7890a0cc448c30fd82aa3cb1e024b28aaf8177d7bfead21c7c96479f90.
//
// Solidity: event Round1Filled(bytes32 id)
func (_DKG *DKGFilterer) FilterRound1Filled(opts *bind.FilterOpts) (*DKGRound1FilledIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Round1Filled")
	if err != nil {
		return nil, err
	}
	return &DKGRound1FilledIterator{contract: _DKG.contract, event: "Round1Filled", logs: logs, sub: sub}, nil
}

// WatchRound1Filled is a free log subscription operation binding the contract event 0x826f5f7890a0cc448c30fd82aa3cb1e024b28aaf8177d7bfead21c7c96479f90.
//
// Solidity: event Round1Filled(bytes32 id)
func (_DKG *DKGFilterer) WatchRound1Filled(opts *bind.WatchOpts, sink chan<- *DKGRound1Filled) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Round1Filled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRound1Filled)
				if err := _DKG.contract.UnpackLog(event, "Round1Filled", log); err != nil {
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

// ParseRound1Filled is a log parse operation binding the contract event 0x826f5f7890a0cc448c30fd82aa3cb1e024b28aaf8177d7bfead21c7c96479f90.
//
// Solidity: event Round1Filled(bytes32 id)
func (_DKG *DKGFilterer) ParseRound1Filled(log types.Log) (*DKGRound1Filled, error) {
	event := new(DKGRound1Filled)
	if err := _DKG.contract.UnpackLog(event, "Round1Filled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRound1ProvidedIterator is returned from FilterRound1Provided and is used to iterate over the raw logs and unpacked data for Round1Provided events raised by the DKG contract.
type DKGRound1ProvidedIterator struct {
	Event *DKGRound1Provided // Event containing the contract specifics and raw log

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
func (it *DKGRound1ProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRound1Provided)
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
		it.Event = new(DKGRound1Provided)
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
func (it *DKGRound1ProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRound1ProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRound1Provided represents a Round1Provided event raised by the DKG contract.
type DKGRound1Provided struct {
	Id        [32]byte
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRound1Provided is a free log retrieval operation binding the contract event 0xff0bbf01029d35ad1de4979157c2b093013823420a431f12d6702e97f9ad4d80.
//
// Solidity: event Round1Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) FilterRound1Provided(opts *bind.FilterOpts) (*DKGRound1ProvidedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Round1Provided")
	if err != nil {
		return nil, err
	}
	return &DKGRound1ProvidedIterator{contract: _DKG.contract, event: "Round1Provided", logs: logs, sub: sub}, nil
}

// WatchRound1Provided is a free log subscription operation binding the contract event 0xff0bbf01029d35ad1de4979157c2b093013823420a431f12d6702e97f9ad4d80.
//
// Solidity: event Round1Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) WatchRound1Provided(opts *bind.WatchOpts, sink chan<- *DKGRound1Provided) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Round1Provided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRound1Provided)
				if err := _DKG.contract.UnpackLog(event, "Round1Provided", log); err != nil {
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

// ParseRound1Provided is a log parse operation binding the contract event 0xff0bbf01029d35ad1de4979157c2b093013823420a431f12d6702e97f9ad4d80.
//
// Solidity: event Round1Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) ParseRound1Provided(log types.Log) (*DKGRound1Provided, error) {
	event := new(DKGRound1Provided)
	if err := _DKG.contract.UnpackLog(event, "Round1Provided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRound2FilledIterator is returned from FilterRound2Filled and is used to iterate over the raw logs and unpacked data for Round2Filled events raised by the DKG contract.
type DKGRound2FilledIterator struct {
	Event *DKGRound2Filled // Event containing the contract specifics and raw log

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
func (it *DKGRound2FilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRound2Filled)
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
		it.Event = new(DKGRound2Filled)
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
func (it *DKGRound2FilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRound2FilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRound2Filled represents a Round2Filled event raised by the DKG contract.
type DKGRound2Filled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRound2Filled is a free log retrieval operation binding the contract event 0x54a928510d9ccbdf7d88149e28187e49f1a15466afc130831cd089ca4169b45a.
//
// Solidity: event Round2Filled(bytes32 id)
func (_DKG *DKGFilterer) FilterRound2Filled(opts *bind.FilterOpts) (*DKGRound2FilledIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Round2Filled")
	if err != nil {
		return nil, err
	}
	return &DKGRound2FilledIterator{contract: _DKG.contract, event: "Round2Filled", logs: logs, sub: sub}, nil
}

// WatchRound2Filled is a free log subscription operation binding the contract event 0x54a928510d9ccbdf7d88149e28187e49f1a15466afc130831cd089ca4169b45a.
//
// Solidity: event Round2Filled(bytes32 id)
func (_DKG *DKGFilterer) WatchRound2Filled(opts *bind.WatchOpts, sink chan<- *DKGRound2Filled) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Round2Filled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRound2Filled)
				if err := _DKG.contract.UnpackLog(event, "Round2Filled", log); err != nil {
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

// ParseRound2Filled is a log parse operation binding the contract event 0x54a928510d9ccbdf7d88149e28187e49f1a15466afc130831cd089ca4169b45a.
//
// Solidity: event Round2Filled(bytes32 id)
func (_DKG *DKGFilterer) ParseRound2Filled(log types.Log) (*DKGRound2Filled, error) {
	event := new(DKGRound2Filled)
	if err := _DKG.contract.UnpackLog(event, "Round2Filled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRound2ProvidedIterator is returned from FilterRound2Provided and is used to iterate over the raw logs and unpacked data for Round2Provided events raised by the DKG contract.
type DKGRound2ProvidedIterator struct {
	Event *DKGRound2Provided // Event containing the contract specifics and raw log

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
func (it *DKGRound2ProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRound2Provided)
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
		it.Event = new(DKGRound2Provided)
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
func (it *DKGRound2ProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRound2ProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRound2Provided represents a Round2Provided event raised by the DKG contract.
type DKGRound2Provided struct {
	Id        [32]byte
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRound2Provided is a free log retrieval operation binding the contract event 0x3558ea64e4bcb51d449be257bf9463a6ae305caed00eb7b758c81e175962da0a.
//
// Solidity: event Round2Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) FilterRound2Provided(opts *bind.FilterOpts) (*DKGRound2ProvidedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Round2Provided")
	if err != nil {
		return nil, err
	}
	return &DKGRound2ProvidedIterator{contract: _DKG.contract, event: "Round2Provided", logs: logs, sub: sub}, nil
}

// WatchRound2Provided is a free log subscription operation binding the contract event 0x3558ea64e4bcb51d449be257bf9463a6ae305caed00eb7b758c81e175962da0a.
//
// Solidity: event Round2Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) WatchRound2Provided(opts *bind.WatchOpts, sink chan<- *DKGRound2Provided) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Round2Provided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRound2Provided)
				if err := _DKG.contract.UnpackLog(event, "Round2Provided", log); err != nil {
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

// ParseRound2Provided is a log parse operation binding the contract event 0x3558ea64e4bcb51d449be257bf9463a6ae305caed00eb7b758c81e175962da0a.
//
// Solidity: event Round2Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) ParseRound2Provided(log types.Log) (*DKGRound2Provided, error) {
	event := new(DKGRound2Provided)
	if err := _DKG.contract.UnpackLog(event, "Round2Provided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRound3FilledIterator is returned from FilterRound3Filled and is used to iterate over the raw logs and unpacked data for Round3Filled events raised by the DKG contract.
type DKGRound3FilledIterator struct {
	Event *DKGRound3Filled // Event containing the contract specifics and raw log

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
func (it *DKGRound3FilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRound3Filled)
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
		it.Event = new(DKGRound3Filled)
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
func (it *DKGRound3FilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRound3FilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRound3Filled represents a Round3Filled event raised by the DKG contract.
type DKGRound3Filled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRound3Filled is a free log retrieval operation binding the contract event 0x5bdd744a1c2812910c75c3d1cc27e109d0397cb4c58542c313b9e530a8dc57db.
//
// Solidity: event Round3Filled(bytes32 id)
func (_DKG *DKGFilterer) FilterRound3Filled(opts *bind.FilterOpts) (*DKGRound3FilledIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Round3Filled")
	if err != nil {
		return nil, err
	}
	return &DKGRound3FilledIterator{contract: _DKG.contract, event: "Round3Filled", logs: logs, sub: sub}, nil
}

// WatchRound3Filled is a free log subscription operation binding the contract event 0x5bdd744a1c2812910c75c3d1cc27e109d0397cb4c58542c313b9e530a8dc57db.
//
// Solidity: event Round3Filled(bytes32 id)
func (_DKG *DKGFilterer) WatchRound3Filled(opts *bind.WatchOpts, sink chan<- *DKGRound3Filled) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Round3Filled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRound3Filled)
				if err := _DKG.contract.UnpackLog(event, "Round3Filled", log); err != nil {
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

// ParseRound3Filled is a log parse operation binding the contract event 0x5bdd744a1c2812910c75c3d1cc27e109d0397cb4c58542c313b9e530a8dc57db.
//
// Solidity: event Round3Filled(bytes32 id)
func (_DKG *DKGFilterer) ParseRound3Filled(log types.Log) (*DKGRound3Filled, error) {
	event := new(DKGRound3Filled)
	if err := _DKG.contract.UnpackLog(event, "Round3Filled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRound3ProvidedIterator is returned from FilterRound3Provided and is used to iterate over the raw logs and unpacked data for Round3Provided events raised by the DKG contract.
type DKGRound3ProvidedIterator struct {
	Event *DKGRound3Provided // Event containing the contract specifics and raw log

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
func (it *DKGRound3ProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRound3Provided)
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
		it.Event = new(DKGRound3Provided)
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
func (it *DKGRound3ProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRound3ProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRound3Provided represents a Round3Provided event raised by the DKG contract.
type DKGRound3Provided struct {
	Id        [32]byte
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRound3Provided is a free log retrieval operation binding the contract event 0x9266d80263563030832e287e62a2b307cd4428d1df44fdc13be65f8b0104dbc9.
//
// Solidity: event Round3Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) FilterRound3Provided(opts *bind.FilterOpts) (*DKGRound3ProvidedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Round3Provided")
	if err != nil {
		return nil, err
	}
	return &DKGRound3ProvidedIterator{contract: _DKG.contract, event: "Round3Provided", logs: logs, sub: sub}, nil
}

// WatchRound3Provided is a free log subscription operation binding the contract event 0x9266d80263563030832e287e62a2b307cd4428d1df44fdc13be65f8b0104dbc9.
//
// Solidity: event Round3Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) WatchRound3Provided(opts *bind.WatchOpts, sink chan<- *DKGRound3Provided) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Round3Provided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRound3Provided)
				if err := _DKG.contract.UnpackLog(event, "Round3Provided", log); err != nil {
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

// ParseRound3Provided is a log parse operation binding the contract event 0x9266d80263563030832e287e62a2b307cd4428d1df44fdc13be65f8b0104dbc9.
//
// Solidity: event Round3Provided(bytes32 id, address validator)
func (_DKG *DKGFilterer) ParseRound3Provided(log types.Log) (*DKGRound3Provided, error) {
	event := new(DKGRound3Provided)
	if err := _DKG.contract.UnpackLog(event, "Round3Provided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
