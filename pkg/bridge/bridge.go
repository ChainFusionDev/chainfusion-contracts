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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"LiquidityPoolsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"TokenManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"ValidatorManagerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvalsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_txHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approveTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_txHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPools\",\"outputs\":[{\"internalType\":\"contractLiquidityPools\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"setLiquidityPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"setValidatorManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"contractTokenManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorManager\",\"outputs\":[{\"internalType\":\"contractValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Approvals is a free data retrieval call binding the contract method 0x24f6c776.
//
// Solidity: function approvals(bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCaller) Approvals(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "approvals", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approvals is a free data retrieval call binding the contract method 0x24f6c776.
//
// Solidity: function approvals(bytes32 , address ) view returns(bool)
func (_Bridge *BridgeSession) Approvals(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Bridge.Contract.Approvals(&_Bridge.CallOpts, arg0, arg1)
}

// Approvals is a free data retrieval call binding the contract method 0x24f6c776.
//
// Solidity: function approvals(bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Approvals(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Bridge.Contract.Approvals(&_Bridge.CallOpts, arg0, arg1)
}

// ApprovalsCount is a free data retrieval call binding the contract method 0xdf52d05d.
//
// Solidity: function approvalsCount(bytes32 ) view returns(uint256)
func (_Bridge *BridgeCaller) ApprovalsCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "approvalsCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovalsCount is a free data retrieval call binding the contract method 0xdf52d05d.
//
// Solidity: function approvalsCount(bytes32 ) view returns(uint256)
func (_Bridge *BridgeSession) ApprovalsCount(arg0 [32]byte) (*big.Int, error) {
	return _Bridge.Contract.ApprovalsCount(&_Bridge.CallOpts, arg0)
}

// ApprovalsCount is a free data retrieval call binding the contract method 0xdf52d05d.
//
// Solidity: function approvalsCount(bytes32 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) ApprovalsCount(arg0 [32]byte) (*big.Int, error) {
	return _Bridge.Contract.ApprovalsCount(&_Bridge.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Executed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "executed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Executed(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Executed(&_Bridge.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Executed(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Executed(&_Bridge.CallOpts, arg0)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Bridge *BridgeCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Bridge *BridgeSession) FeeManager() (common.Address, error) {
	return _Bridge.Contract.FeeManager(&_Bridge.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Bridge *BridgeCallerSession) FeeManager() (common.Address, error) {
	return _Bridge.Contract.FeeManager(&_Bridge.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x32498425.
//
// Solidity: function isApproved(bytes _txHash, address _token, address _receiver, uint256 _amount) view returns(bool)
func (_Bridge *BridgeCaller) IsApproved(opts *bind.CallOpts, _txHash []byte, _token common.Address, _receiver common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isApproved", _txHash, _token, _receiver, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x32498425.
//
// Solidity: function isApproved(bytes _txHash, address _token, address _receiver, uint256 _amount) view returns(bool)
func (_Bridge *BridgeSession) IsApproved(_txHash []byte, _token common.Address, _receiver common.Address, _amount *big.Int) (bool, error) {
	return _Bridge.Contract.IsApproved(&_Bridge.CallOpts, _txHash, _token, _receiver, _amount)
}

// IsApproved is a free data retrieval call binding the contract method 0x32498425.
//
// Solidity: function isApproved(bytes _txHash, address _token, address _receiver, uint256 _amount) view returns(bool)
func (_Bridge *BridgeCallerSession) IsApproved(_txHash []byte, _token common.Address, _receiver common.Address, _amount *big.Int) (bool, error) {
	return _Bridge.Contract.IsApproved(&_Bridge.CallOpts, _txHash, _token, _receiver, _amount)
}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_Bridge *BridgeCaller) LiquidityPools(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "liquidityPools")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_Bridge *BridgeSession) LiquidityPools() (common.Address, error) {
	return _Bridge.Contract.LiquidityPools(&_Bridge.CallOpts)
}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_Bridge *BridgeCallerSession) LiquidityPools() (common.Address, error) {
	return _Bridge.Contract.LiquidityPools(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_Bridge *BridgeCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_Bridge *BridgeSession) TokenManager() (common.Address, error) {
	return _Bridge.Contract.TokenManager(&_Bridge.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_Bridge *BridgeCallerSession) TokenManager() (common.Address, error) {
	return _Bridge.Contract.TokenManager(&_Bridge.CallOpts)
}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Bridge *BridgeCaller) ValidatorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "validatorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Bridge *BridgeSession) ValidatorManager() (common.Address, error) {
	return _Bridge.Contract.ValidatorManager(&_Bridge.CallOpts)
}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Bridge *BridgeCallerSession) ValidatorManager() (common.Address, error) {
	return _Bridge.Contract.ValidatorManager(&_Bridge.CallOpts)
}

// ApproveTransfer is a paid mutator transaction binding the contract method 0x7871e4e0.
//
// Solidity: function approveTransfer(bytes _txHash, address _token, uint256 _sourceChainId, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) ApproveTransfer(opts *bind.TransactOpts, _txHash []byte, _token common.Address, _sourceChainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "approveTransfer", _txHash, _token, _sourceChainId, _receiver, _amount)
}

// ApproveTransfer is a paid mutator transaction binding the contract method 0x7871e4e0.
//
// Solidity: function approveTransfer(bytes _txHash, address _token, uint256 _sourceChainId, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeSession) ApproveTransfer(_txHash []byte, _token common.Address, _sourceChainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ApproveTransfer(&_Bridge.TransactOpts, _txHash, _token, _sourceChainId, _receiver, _amount)
}

// ApproveTransfer is a paid mutator transaction binding the contract method 0x7871e4e0.
//
// Solidity: function approveTransfer(bytes _txHash, address _token, uint256 _sourceChainId, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) ApproveTransfer(_txHash []byte, _token common.Address, _sourceChainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ApproveTransfer(&_Bridge.TransactOpts, _txHash, _token, _sourceChainId, _receiver, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address _token, uint256 _chainId, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _chainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", _token, _chainId, _receiver, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address _token, uint256 _chainId, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeSession) Deposit(_token common.Address, _chainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, _token, _chainId, _receiver, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address _token, uint256 _chainId, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) Deposit(_token common.Address, _chainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, _token, _chainId, _receiver, _amount)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 _chainId, address _receiver) payable returns()
func (_Bridge *BridgeTransactor) DepositNative(opts *bind.TransactOpts, _chainId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositNative", _chainId, _receiver)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 _chainId, address _receiver) payable returns()
func (_Bridge *BridgeSession) DepositNative(_chainId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, _chainId, _receiver)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 _chainId, address _receiver) payable returns()
func (_Bridge *BridgeTransactorSession) DepositNative(_chainId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, _chainId, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _validatorManager, address _tokenManager, address _liquidityPools, address _feeManager) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _validatorManager common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", _owner, _validatorManager, _tokenManager, _liquidityPools, _feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _validatorManager, address _tokenManager, address _liquidityPools, address _feeManager) returns()
func (_Bridge *BridgeSession) Initialize(_owner common.Address, _validatorManager common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, _owner, _validatorManager, _tokenManager, _liquidityPools, _feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _validatorManager, address _tokenManager, address _liquidityPools, address _feeManager) returns()
func (_Bridge *BridgeTransactorSession) Initialize(_owner common.Address, _validatorManager common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, _owner, _validatorManager, _tokenManager, _liquidityPools, _feeManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_Bridge *BridgeTransactor) SetFeeManager(opts *bind.TransactOpts, _feeManager common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFeeManager", _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_Bridge *BridgeSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeManager(&_Bridge.TransactOpts, _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_Bridge *BridgeTransactorSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeManager(&_Bridge.TransactOpts, _feeManager)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_Bridge *BridgeTransactor) SetLiquidityPools(opts *bind.TransactOpts, _liquidityPools common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setLiquidityPools", _liquidityPools)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_Bridge *BridgeSession) SetLiquidityPools(_liquidityPools common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetLiquidityPools(&_Bridge.TransactOpts, _liquidityPools)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_Bridge *BridgeTransactorSession) SetLiquidityPools(_liquidityPools common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetLiquidityPools(&_Bridge.TransactOpts, _liquidityPools)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_Bridge *BridgeTransactor) SetTokenManager(opts *bind.TransactOpts, _tokenManager common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTokenManager", _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_Bridge *BridgeSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenManager(&_Bridge.TransactOpts, _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_Bridge *BridgeTransactorSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenManager(&_Bridge.TransactOpts, _tokenManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Bridge *BridgeTransactor) SetValidatorManager(opts *bind.TransactOpts, _validatorManager common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setValidatorManager", _validatorManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Bridge *BridgeSession) SetValidatorManager(_validatorManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetValidatorManager(&_Bridge.TransactOpts, _validatorManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Bridge *BridgeTransactorSession) SetValidatorManager(_validatorManager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetValidatorManager(&_Bridge.TransactOpts, _validatorManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the Bridge contract.
type BridgeApprovedIterator struct {
	Event *BridgeApproved // Event containing the contract specifics and raw log

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
func (it *BridgeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeApproved)
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
		it.Event = new(BridgeApproved)
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
func (it *BridgeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeApproved represents a Approved event raised by the Bridge contract.
type BridgeApproved struct {
	Id        [32]byte
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x90b535106f433f3869de4725076776361dc9946b091db11206fa1b526a85ecd4.
//
// Solidity: event Approved(bytes32 id, address validator)
func (_Bridge *BridgeFilterer) FilterApproved(opts *bind.FilterOpts) (*BridgeApprovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Approved")
	if err != nil {
		return nil, err
	}
	return &BridgeApprovedIterator{contract: _Bridge.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x90b535106f433f3869de4725076776361dc9946b091db11206fa1b526a85ecd4.
//
// Solidity: event Approved(bytes32 id, address validator)
func (_Bridge *BridgeFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *BridgeApproved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Approved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeApproved)
				if err := _Bridge.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x90b535106f433f3869de4725076776361dc9946b091db11206fa1b526a85ecd4.
//
// Solidity: event Approved(bytes32 id, address validator)
func (_Bridge *BridgeFilterer) ParseApproved(log types.Log) (*BridgeApproved, error) {
	event := new(BridgeApproved)
	if err := _Bridge.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Bridge contract.
type BridgeDepositedIterator struct {
	Event *BridgeDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposited)
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
		it.Event = new(BridgeDeposited)
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
func (it *BridgeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposited represents a Deposited event raised by the Bridge contract.
type BridgeDeposited struct {
	Token              common.Address
	DestinationToken   common.Address
	DestinationChainId *big.Int
	Receiver           common.Address
	Fee                *big.Int
	TransferAmount     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xb64b8d2a723c97c9a2594ccb3778465edeb4e8b73151aee791183eeda118efff.
//
// Solidity: event Deposited(address token, address destinationToken, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) FilterDeposited(opts *bind.FilterOpts) (*BridgeDepositedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedIterator{contract: _Bridge.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xb64b8d2a723c97c9a2594ccb3778465edeb4e8b73151aee791183eeda118efff.
//
// Solidity: event Deposited(address token, address destinationToken, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *BridgeDeposited) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposited)
				if err := _Bridge.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xb64b8d2a723c97c9a2594ccb3778465edeb4e8b73151aee791183eeda118efff.
//
// Solidity: event Deposited(address token, address destinationToken, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) ParseDeposited(log types.Log) (*BridgeDeposited, error) {
	event := new(BridgeDeposited)
	if err := _Bridge.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the Bridge contract.
type BridgeDepositedNativeIterator struct {
	Event *BridgeDepositedNative // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedNative)
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
		it.Event = new(BridgeDepositedNative)
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
func (it *BridgeDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedNative represents a DepositedNative event raised by the Bridge contract.
type BridgeDepositedNative struct {
	Token              common.Address
	DestinationChainId *big.Int
	Receiver           common.Address
	Fee                *big.Int
	TransferAmount     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0xe342e0b03345f9d35e9e811082f2c4d8b50a4236c3d3bb140babfdd7ca18b294.
//
// Solidity: event DepositedNative(address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*BridgeDepositedNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedNativeIterator{contract: _Bridge.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0xe342e0b03345f9d35e9e811082f2c4d8b50a4236c3d3bb140babfdd7ca18b294.
//
// Solidity: event DepositedNative(address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *BridgeDepositedNative) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedNative)
				if err := _Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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

// ParseDepositedNative is a log parse operation binding the contract event 0xe342e0b03345f9d35e9e811082f2c4d8b50a4236c3d3bb140babfdd7ca18b294.
//
// Solidity: event DepositedNative(address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) ParseDepositedNative(log types.Log) (*BridgeDepositedNative, error) {
	event := new(BridgeDepositedNative)
	if err := _Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the Bridge contract.
type BridgeFeeManagerUpdatedIterator struct {
	Event *BridgeFeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerUpdated)
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
		it.Event = new(BridgeFeeManagerUpdated)
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
func (it *BridgeFeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerUpdated represents a FeeManagerUpdated event raised by the Bridge contract.
type BridgeFeeManagerUpdated struct {
	FeeManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address _feeManager)
func (_Bridge *BridgeFilterer) FilterFeeManagerUpdated(opts *bind.FilterOpts) (*BridgeFeeManagerUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerUpdatedIterator{contract: _Bridge.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address _feeManager)
func (_Bridge *BridgeFilterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerUpdated)
				if err := _Bridge.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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

// ParseFeeManagerUpdated is a log parse operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address _feeManager)
func (_Bridge *BridgeFilterer) ParseFeeManagerUpdated(log types.Log) (*BridgeFeeManagerUpdated, error) {
	event := new(BridgeFeeManagerUpdated)
	if err := _Bridge.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLiquidityPoolsUpdatedIterator is returned from FilterLiquidityPoolsUpdated and is used to iterate over the raw logs and unpacked data for LiquidityPoolsUpdated events raised by the Bridge contract.
type BridgeLiquidityPoolsUpdatedIterator struct {
	Event *BridgeLiquidityPoolsUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeLiquidityPoolsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLiquidityPoolsUpdated)
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
		it.Event = new(BridgeLiquidityPoolsUpdated)
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
func (it *BridgeLiquidityPoolsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLiquidityPoolsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLiquidityPoolsUpdated represents a LiquidityPoolsUpdated event raised by the Bridge contract.
type BridgeLiquidityPoolsUpdated struct {
	LiquidityPools common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidityPoolsUpdated is a free log retrieval operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_Bridge *BridgeFilterer) FilterLiquidityPoolsUpdated(opts *bind.FilterOpts) (*BridgeLiquidityPoolsUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LiquidityPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeLiquidityPoolsUpdatedIterator{contract: _Bridge.contract, event: "LiquidityPoolsUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityPoolsUpdated is a free log subscription operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_Bridge *BridgeFilterer) WatchLiquidityPoolsUpdated(opts *bind.WatchOpts, sink chan<- *BridgeLiquidityPoolsUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LiquidityPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLiquidityPoolsUpdated)
				if err := _Bridge.contract.UnpackLog(event, "LiquidityPoolsUpdated", log); err != nil {
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

// ParseLiquidityPoolsUpdated is a log parse operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_Bridge *BridgeFilterer) ParseLiquidityPoolsUpdated(log types.Log) (*BridgeLiquidityPoolsUpdated, error) {
	event := new(BridgeLiquidityPoolsUpdated)
	if err := _Bridge.contract.UnpackLog(event, "LiquidityPoolsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenManagerUpdatedIterator is returned from FilterTokenManagerUpdated and is used to iterate over the raw logs and unpacked data for TokenManagerUpdated events raised by the Bridge contract.
type BridgeTokenManagerUpdatedIterator struct {
	Event *BridgeTokenManagerUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTokenManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenManagerUpdated)
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
		it.Event = new(BridgeTokenManagerUpdated)
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
func (it *BridgeTokenManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenManagerUpdated represents a TokenManagerUpdated event raised by the Bridge contract.
type BridgeTokenManagerUpdated struct {
	TokenManager common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenManagerUpdated is a free log retrieval operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address _tokenManager)
func (_Bridge *BridgeFilterer) FilterTokenManagerUpdated(opts *bind.FilterOpts) (*BridgeTokenManagerUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTokenManagerUpdatedIterator{contract: _Bridge.contract, event: "TokenManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenManagerUpdated is a free log subscription operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address _tokenManager)
func (_Bridge *BridgeFilterer) WatchTokenManagerUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTokenManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenManagerUpdated)
				if err := _Bridge.contract.UnpackLog(event, "TokenManagerUpdated", log); err != nil {
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

// ParseTokenManagerUpdated is a log parse operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address _tokenManager)
func (_Bridge *BridgeFilterer) ParseTokenManagerUpdated(log types.Log) (*BridgeTokenManagerUpdated, error) {
	event := new(BridgeTokenManagerUpdated)
	if err := _Bridge.contract.UnpackLog(event, "TokenManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the Bridge contract.
type BridgeTransferredIterator struct {
	Event *BridgeTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTransferred)
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
		it.Event = new(BridgeTransferred)
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
func (it *BridgeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTransferred represents a Transferred event raised by the Bridge contract.
type BridgeTransferred struct {
	Token         common.Address
	SourceChainId *big.Int
	Receiver      common.Address
	Amount        *big.Int
	Validator     common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0x11937e05f51288f2b167a3d7cef81d81f3662b2765f9e042a91089619aec438a.
//
// Solidity: event Transferred(address token, uint256 sourceChainId, address receiver, uint256 amount, address validator)
func (_Bridge *BridgeFilterer) FilterTransferred(opts *bind.FilterOpts) (*BridgeTransferredIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &BridgeTransferredIterator{contract: _Bridge.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0x11937e05f51288f2b167a3d7cef81d81f3662b2765f9e042a91089619aec438a.
//
// Solidity: event Transferred(address token, uint256 sourceChainId, address receiver, uint256 amount, address validator)
func (_Bridge *BridgeFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *BridgeTransferred) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTransferred)
				if err := _Bridge.contract.UnpackLog(event, "Transferred", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0x11937e05f51288f2b167a3d7cef81d81f3662b2765f9e042a91089619aec438a.
//
// Solidity: event Transferred(address token, uint256 sourceChainId, address receiver, uint256 amount, address validator)
func (_Bridge *BridgeFilterer) ParseTransferred(log types.Log) (*BridgeTransferred, error) {
	event := new(BridgeTransferred)
	if err := _Bridge.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorManagerUpdatedIterator is returned from FilterValidatorManagerUpdated and is used to iterate over the raw logs and unpacked data for ValidatorManagerUpdated events raised by the Bridge contract.
type BridgeValidatorManagerUpdatedIterator struct {
	Event *BridgeValidatorManagerUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorManagerUpdated)
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
		it.Event = new(BridgeValidatorManagerUpdated)
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
func (it *BridgeValidatorManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorManagerUpdated represents a ValidatorManagerUpdated event raised by the Bridge contract.
type BridgeValidatorManagerUpdated struct {
	ValidatorManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorManagerUpdated is a free log retrieval operation binding the contract event 0xcc1d286dbc29083de2c32a167d46163f88f978e5b16fea673f47cc4e413c9ef5.
//
// Solidity: event ValidatorManagerUpdated(address _validatorManager)
func (_Bridge *BridgeFilterer) FilterValidatorManagerUpdated(opts *bind.FilterOpts) (*BridgeValidatorManagerUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ValidatorManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorManagerUpdatedIterator{contract: _Bridge.contract, event: "ValidatorManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorManagerUpdated is a free log subscription operation binding the contract event 0xcc1d286dbc29083de2c32a167d46163f88f978e5b16fea673f47cc4e413c9ef5.
//
// Solidity: event ValidatorManagerUpdated(address _validatorManager)
func (_Bridge *BridgeFilterer) WatchValidatorManagerUpdated(opts *bind.WatchOpts, sink chan<- *BridgeValidatorManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ValidatorManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorManagerUpdated)
				if err := _Bridge.contract.UnpackLog(event, "ValidatorManagerUpdated", log); err != nil {
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

// ParseValidatorManagerUpdated is a log parse operation binding the contract event 0xcc1d286dbc29083de2c32a167d46163f88f978e5b16fea673f47cc4e413c9ef5.
//
// Solidity: event ValidatorManagerUpdated(address _validatorManager)
func (_Bridge *BridgeFilterer) ParseValidatorManagerUpdated(log types.Log) (*BridgeValidatorManagerUpdated, error) {
	event := new(BridgeValidatorManagerUpdated)
	if err := _Bridge.contract.UnpackLog(event, "ValidatorManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
