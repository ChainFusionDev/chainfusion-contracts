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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"LiquidityPoolsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Reverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"TokenManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"ValidatorAddressUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAppAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAppAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_txHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPools\",\"outputs\":[{\"internalType\":\"contractLiquidityPools\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayBridge\",\"outputs\":[{\"internalType\":\"contractRelayBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"revertSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"setLiquidityPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractSignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"contractTokenManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_Bridge *BridgeCaller) BridgeAppAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeAppAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_Bridge *BridgeSession) BridgeAppAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeAppAddress(&_Bridge.CallOpts)
}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_Bridge *BridgeCallerSession) BridgeAppAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeAppAddress(&_Bridge.CallOpts)
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

// IsExecuted is a free data retrieval call binding the contract method 0x432f37e8.
//
// Solidity: function isExecuted(address _sender, bytes _txHash, address _token, address _receiver, uint256 _amount) view returns(bool)
func (_Bridge *BridgeCaller) IsExecuted(opts *bind.CallOpts, _sender common.Address, _txHash []byte, _token common.Address, _receiver common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isExecuted", _sender, _txHash, _token, _receiver, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecuted is a free data retrieval call binding the contract method 0x432f37e8.
//
// Solidity: function isExecuted(address _sender, bytes _txHash, address _token, address _receiver, uint256 _amount) view returns(bool)
func (_Bridge *BridgeSession) IsExecuted(_sender common.Address, _txHash []byte, _token common.Address, _receiver common.Address, _amount *big.Int) (bool, error) {
	return _Bridge.Contract.IsExecuted(&_Bridge.CallOpts, _sender, _txHash, _token, _receiver, _amount)
}

// IsExecuted is a free data retrieval call binding the contract method 0x432f37e8.
//
// Solidity: function isExecuted(address _sender, bytes _txHash, address _token, address _receiver, uint256 _amount) view returns(bool)
func (_Bridge *BridgeCallerSession) IsExecuted(_sender common.Address, _txHash []byte, _token common.Address, _receiver common.Address, _amount *big.Int) (bool, error) {
	return _Bridge.Contract.IsExecuted(&_Bridge.CallOpts, _sender, _txHash, _token, _receiver, _amount)
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

// RelayBridge is a free data retrieval call binding the contract method 0xa76aa0e2.
//
// Solidity: function relayBridge() view returns(address)
func (_Bridge *BridgeCaller) RelayBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "relayBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayBridge is a free data retrieval call binding the contract method 0xa76aa0e2.
//
// Solidity: function relayBridge() view returns(address)
func (_Bridge *BridgeSession) RelayBridge() (common.Address, error) {
	return _Bridge.Contract.RelayBridge(&_Bridge.CallOpts)
}

// RelayBridge is a free data retrieval call binding the contract method 0xa76aa0e2.
//
// Solidity: function relayBridge() view returns(address)
func (_Bridge *BridgeCallerSession) RelayBridge() (common.Address, error) {
	return _Bridge.Contract.RelayBridge(&_Bridge.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_Bridge *BridgeCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_Bridge *BridgeSession) SignerStorage() (common.Address, error) {
	return _Bridge.Contract.SignerStorage(&_Bridge.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_Bridge *BridgeCallerSession) SignerStorage() (common.Address, error) {
	return _Bridge.Contract.SignerStorage(&_Bridge.CallOpts)
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

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 , bytes data) returns()
func (_Bridge *BridgeTransactor) Execute(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "execute", arg0, data)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 , bytes data) returns()
func (_Bridge *BridgeSession) Execute(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Execute(&_Bridge.TransactOpts, arg0, data)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 , bytes data) returns()
func (_Bridge *BridgeTransactorSession) Execute(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Execute(&_Bridge.TransactOpts, arg0, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _relayBridgeAddress, address _signerStorage, address _tokenManager, address _liquidityPools, address _feeManager, address _bridgeAppAddress) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, _relayBridgeAddress common.Address, _signerStorage common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address, _bridgeAppAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", _relayBridgeAddress, _signerStorage, _tokenManager, _liquidityPools, _feeManager, _bridgeAppAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _relayBridgeAddress, address _signerStorage, address _tokenManager, address _liquidityPools, address _feeManager, address _bridgeAppAddress) returns()
func (_Bridge *BridgeSession) Initialize(_relayBridgeAddress common.Address, _signerStorage common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address, _bridgeAppAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, _relayBridgeAddress, _signerStorage, _tokenManager, _liquidityPools, _feeManager, _bridgeAppAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _relayBridgeAddress, address _signerStorage, address _tokenManager, address _liquidityPools, address _feeManager, address _bridgeAppAddress) returns()
func (_Bridge *BridgeTransactorSession) Initialize(_relayBridgeAddress common.Address, _signerStorage common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address, _bridgeAppAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, _relayBridgeAddress, _signerStorage, _tokenManager, _liquidityPools, _feeManager, _bridgeAppAddress)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 , bytes data) returns()
func (_Bridge *BridgeTransactor) RevertSend(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revertSend", arg0, data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 , bytes data) returns()
func (_Bridge *BridgeSession) RevertSend(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RevertSend(&_Bridge.TransactOpts, arg0, data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 , bytes data) returns()
func (_Bridge *BridgeTransactorSession) RevertSend(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.RevertSend(&_Bridge.TransactOpts, arg0, data)
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
	Sender             common.Address
	Token              common.Address
	DestinationToken   common.Address
	DestinationChainId *big.Int
	Receiver           common.Address
	Fee                *big.Int
	TransferAmount     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x38ae17e88c657bd2a0a842ef8220ce6d392351c85ca7115badee7e8a339ea01e.
//
// Solidity: event Deposited(address sender, address token, address destinationToken, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) FilterDeposited(opts *bind.FilterOpts) (*BridgeDepositedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedIterator{contract: _Bridge.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x38ae17e88c657bd2a0a842ef8220ce6d392351c85ca7115badee7e8a339ea01e.
//
// Solidity: event Deposited(address sender, address token, address destinationToken, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
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

// ParseDeposited is a log parse operation binding the contract event 0x38ae17e88c657bd2a0a842ef8220ce6d392351c85ca7115badee7e8a339ea01e.
//
// Solidity: event Deposited(address sender, address token, address destinationToken, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
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
	Sender             common.Address
	Token              common.Address
	DestinationChainId *big.Int
	Receiver           common.Address
	Fee                *big.Int
	TransferAmount     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0x594aceaea2948a8384c265adbb963c46e8d8b5c405d4d3b98622384fbbce00ac.
//
// Solidity: event DepositedNative(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_Bridge *BridgeFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*BridgeDepositedNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedNativeIterator{contract: _Bridge.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x594aceaea2948a8384c265adbb963c46e8d8b5c405d4d3b98622384fbbce00ac.
//
// Solidity: event DepositedNative(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
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

// ParseDepositedNative is a log parse operation binding the contract event 0x594aceaea2948a8384c265adbb963c46e8d8b5c405d4d3b98622384fbbce00ac.
//
// Solidity: event DepositedNative(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
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

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// BridgeRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the Bridge contract.
type BridgeRevertedIterator struct {
	Event *BridgeReverted // Event containing the contract specifics and raw log

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
func (it *BridgeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeReverted)
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
		it.Event = new(BridgeReverted)
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
func (it *BridgeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeReverted represents a Reverted event raised by the Bridge contract.
type BridgeReverted struct {
	Sender        common.Address
	Token         common.Address
	SourceChainId *big.Int
	Receiver      common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0x6b79def8f04467814352cf318743c5f80de8e5481a67b938b1ea3f9887683501.
//
// Solidity: event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_Bridge *BridgeFilterer) FilterReverted(opts *bind.FilterOpts) (*BridgeRevertedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Reverted")
	if err != nil {
		return nil, err
	}
	return &BridgeRevertedIterator{contract: _Bridge.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x6b79def8f04467814352cf318743c5f80de8e5481a67b938b1ea3f9887683501.
//
// Solidity: event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_Bridge *BridgeFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *BridgeReverted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Reverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeReverted)
				if err := _Bridge.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0x6b79def8f04467814352cf318743c5f80de8e5481a67b938b1ea3f9887683501.
//
// Solidity: event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_Bridge *BridgeFilterer) ParseReverted(log types.Log) (*BridgeReverted, error) {
	event := new(BridgeReverted)
	if err := _Bridge.contract.UnpackLog(event, "Reverted", log); err != nil {
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
	Sender        common.Address
	Token         common.Address
	SourceChainId *big.Int
	Receiver      common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0x621c87a009a8535d1a52333309173ce7117c61f0d228bb8737404185ea6764d2.
//
// Solidity: event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_Bridge *BridgeFilterer) FilterTransferred(opts *bind.FilterOpts) (*BridgeTransferredIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &BridgeTransferredIterator{contract: _Bridge.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0x621c87a009a8535d1a52333309173ce7117c61f0d228bb8737404185ea6764d2.
//
// Solidity: event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
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

// ParseTransferred is a log parse operation binding the contract event 0x621c87a009a8535d1a52333309173ce7117c61f0d228bb8737404185ea6764d2.
//
// Solidity: event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_Bridge *BridgeFilterer) ParseTransferred(log types.Log) (*BridgeTransferred, error) {
	event := new(BridgeTransferred)
	if err := _Bridge.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorAddressUpdatedIterator is returned from FilterValidatorAddressUpdated and is used to iterate over the raw logs and unpacked data for ValidatorAddressUpdated events raised by the Bridge contract.
type BridgeValidatorAddressUpdatedIterator struct {
	Event *BridgeValidatorAddressUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorAddressUpdated)
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
		it.Event = new(BridgeValidatorAddressUpdated)
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
func (it *BridgeValidatorAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorAddressUpdated represents a ValidatorAddressUpdated event raised by the Bridge contract.
type BridgeValidatorAddressUpdated struct {
	ValidatorAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorAddressUpdated is a free log retrieval operation binding the contract event 0xdf570ab3f5307e9adf3ff792b3dbfa0cd0d9f82e667552def683ae20be981f72.
//
// Solidity: event ValidatorAddressUpdated(address _validatorAddress)
func (_Bridge *BridgeFilterer) FilterValidatorAddressUpdated(opts *bind.FilterOpts) (*BridgeValidatorAddressUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ValidatorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorAddressUpdatedIterator{contract: _Bridge.contract, event: "ValidatorAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorAddressUpdated is a free log subscription operation binding the contract event 0xdf570ab3f5307e9adf3ff792b3dbfa0cd0d9f82e667552def683ae20be981f72.
//
// Solidity: event ValidatorAddressUpdated(address _validatorAddress)
func (_Bridge *BridgeFilterer) WatchValidatorAddressUpdated(opts *bind.WatchOpts, sink chan<- *BridgeValidatorAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ValidatorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorAddressUpdated)
				if err := _Bridge.contract.UnpackLog(event, "ValidatorAddressUpdated", log); err != nil {
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

// ParseValidatorAddressUpdated is a log parse operation binding the contract event 0xdf570ab3f5307e9adf3ff792b3dbfa0cd0d9f82e667552def683ae20be981f72.
//
// Solidity: event ValidatorAddressUpdated(address _validatorAddress)
func (_Bridge *BridgeFilterer) ParseValidatorAddressUpdated(log types.Log) (*BridgeValidatorAddressUpdated, error) {
	event := new(BridgeValidatorAddressUpdated)
	if err := _Bridge.contract.UnpackLog(event, "ValidatorAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
