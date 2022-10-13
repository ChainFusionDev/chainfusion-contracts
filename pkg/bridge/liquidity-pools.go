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

// LiquidityPoolsMetaData contains all meta data concerning the LiquidityPools contract.
var LiquidityPoolsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Bridge\",\"type\":\"address\"}],\"name\":\"ERC20BridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercentage\",\"type\":\"uint256\"}],\"name\":\"FeePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenManager\",\"type\":\"address\"}],\"name\":\"TokenManagerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collectedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"distributeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Bridge\",\"outputs\":[{\"internalType\":\"contractERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20Bridge\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityPositions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"rewardsOwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Bridge\",\"type\":\"address\"}],\"name\":\"setERC20Bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"setFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractSignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"contractTokenManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_transferAmount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquidityPoolsABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityPoolsMetaData.ABI instead.
var LiquidityPoolsABI = LiquidityPoolsMetaData.ABI

// LiquidityPools is an auto generated Go binding around an Ethereum contract.
type LiquidityPools struct {
	LiquidityPoolsCaller     // Read-only binding to the contract
	LiquidityPoolsTransactor // Write-only binding to the contract
	LiquidityPoolsFilterer   // Log filterer for contract events
}

// LiquidityPoolsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityPoolsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityPoolsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityPoolsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityPoolsSession struct {
	Contract     *LiquidityPools   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityPoolsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityPoolsCallerSession struct {
	Contract *LiquidityPoolsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LiquidityPoolsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityPoolsTransactorSession struct {
	Contract     *LiquidityPoolsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LiquidityPoolsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityPoolsRaw struct {
	Contract *LiquidityPools // Generic contract binding to access the raw methods on
}

// LiquidityPoolsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityPoolsCallerRaw struct {
	Contract *LiquidityPoolsCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityPoolsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityPoolsTransactorRaw struct {
	Contract *LiquidityPoolsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityPools creates a new instance of LiquidityPools, bound to a specific deployed contract.
func NewLiquidityPools(address common.Address, backend bind.ContractBackend) (*LiquidityPools, error) {
	contract, err := bindLiquidityPools(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityPools{LiquidityPoolsCaller: LiquidityPoolsCaller{contract: contract}, LiquidityPoolsTransactor: LiquidityPoolsTransactor{contract: contract}, LiquidityPoolsFilterer: LiquidityPoolsFilterer{contract: contract}}, nil
}

// NewLiquidityPoolsCaller creates a new read-only instance of LiquidityPools, bound to a specific deployed contract.
func NewLiquidityPoolsCaller(address common.Address, caller bind.ContractCaller) (*LiquidityPoolsCaller, error) {
	contract, err := bindLiquidityPools(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsCaller{contract: contract}, nil
}

// NewLiquidityPoolsTransactor creates a new write-only instance of LiquidityPools, bound to a specific deployed contract.
func NewLiquidityPoolsTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityPoolsTransactor, error) {
	contract, err := bindLiquidityPools(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsTransactor{contract: contract}, nil
}

// NewLiquidityPoolsFilterer creates a new log filterer instance of LiquidityPools, bound to a specific deployed contract.
func NewLiquidityPoolsFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityPoolsFilterer, error) {
	contract, err := bindLiquidityPools(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsFilterer{contract: contract}, nil
}

// bindLiquidityPools binds a generic wrapper to an already deployed contract.
func bindLiquidityPools(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityPoolsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPools *LiquidityPoolsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPools.Contract.LiquidityPoolsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPools *LiquidityPoolsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPools.Contract.LiquidityPoolsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPools *LiquidityPoolsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPools.Contract.LiquidityPoolsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPools *LiquidityPoolsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPools.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPools *LiquidityPoolsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPools.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPools *LiquidityPoolsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPools.Contract.contract.Transact(opts, method, params...)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x181f37c8.
//
// Solidity: function availableLiquidity(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) AvailableLiquidity(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "availableLiquidity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x181f37c8.
//
// Solidity: function availableLiquidity(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) AvailableLiquidity(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.AvailableLiquidity(&_LiquidityPools.CallOpts, arg0)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x181f37c8.
//
// Solidity: function availableLiquidity(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) AvailableLiquidity(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.AvailableLiquidity(&_LiquidityPools.CallOpts, arg0)
}

// CollectedFees is a free data retrieval call binding the contract method 0x1cead9a7.
//
// Solidity: function collectedFees(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) CollectedFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "collectedFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectedFees is a free data retrieval call binding the contract method 0x1cead9a7.
//
// Solidity: function collectedFees(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) CollectedFees(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.CollectedFees(&_LiquidityPools.CallOpts, arg0)
}

// CollectedFees is a free data retrieval call binding the contract method 0x1cead9a7.
//
// Solidity: function collectedFees(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) CollectedFees(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.CollectedFees(&_LiquidityPools.CallOpts, arg0)
}

// Erc20Bridge is a free data retrieval call binding the contract method 0xf8ffea7e.
//
// Solidity: function erc20Bridge() view returns(address)
func (_LiquidityPools *LiquidityPoolsCaller) Erc20Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "erc20Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Bridge is a free data retrieval call binding the contract method 0xf8ffea7e.
//
// Solidity: function erc20Bridge() view returns(address)
func (_LiquidityPools *LiquidityPoolsSession) Erc20Bridge() (common.Address, error) {
	return _LiquidityPools.Contract.Erc20Bridge(&_LiquidityPools.CallOpts)
}

// Erc20Bridge is a free data retrieval call binding the contract method 0xf8ffea7e.
//
// Solidity: function erc20Bridge() view returns(address)
func (_LiquidityPools *LiquidityPoolsCallerSession) Erc20Bridge() (common.Address, error) {
	return _LiquidityPools.Contract.Erc20Bridge(&_LiquidityPools.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_LiquidityPools *LiquidityPoolsCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_LiquidityPools *LiquidityPoolsSession) FeeManager() (common.Address, error) {
	return _LiquidityPools.Contract.FeeManager(&_LiquidityPools.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_LiquidityPools *LiquidityPoolsCallerSession) FeeManager() (common.Address, error) {
	return _LiquidityPools.Contract.FeeManager(&_LiquidityPools.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) FeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) FeePercentage() (*big.Int, error) {
	return _LiquidityPools.Contract.FeePercentage(&_LiquidityPools.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) FeePercentage() (*big.Int, error) {
	return _LiquidityPools.Contract.FeePercentage(&_LiquidityPools.CallOpts)
}

// LiquidityPositions is a free data retrieval call binding the contract method 0x0d96e2be.
//
// Solidity: function liquidityPositions(address , address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_LiquidityPools *LiquidityPoolsCaller) LiquidityPositions(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "liquidityPositions", arg0, arg1)

	outstruct := new(struct {
		Balance          *big.Int
		LastRewardPoints *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardPoints = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LiquidityPositions is a free data retrieval call binding the contract method 0x0d96e2be.
//
// Solidity: function liquidityPositions(address , address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_LiquidityPools *LiquidityPoolsSession) LiquidityPositions(arg0 common.Address, arg1 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	return _LiquidityPools.Contract.LiquidityPositions(&_LiquidityPools.CallOpts, arg0, arg1)
}

// LiquidityPositions is a free data retrieval call binding the contract method 0x0d96e2be.
//
// Solidity: function liquidityPositions(address , address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_LiquidityPools *LiquidityPoolsCallerSession) LiquidityPositions(arg0 common.Address, arg1 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	return _LiquidityPools.Contract.LiquidityPositions(&_LiquidityPools.CallOpts, arg0, arg1)
}

// ProvidedLiquidity is a free data retrieval call binding the contract method 0xb046d729.
//
// Solidity: function providedLiquidity(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) ProvidedLiquidity(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "providedLiquidity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvidedLiquidity is a free data retrieval call binding the contract method 0xb046d729.
//
// Solidity: function providedLiquidity(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) ProvidedLiquidity(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.ProvidedLiquidity(&_LiquidityPools.CallOpts, arg0)
}

// ProvidedLiquidity is a free data retrieval call binding the contract method 0xb046d729.
//
// Solidity: function providedLiquidity(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) ProvidedLiquidity(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.ProvidedLiquidity(&_LiquidityPools.CallOpts, arg0)
}

// RewardsOwing is a free data retrieval call binding the contract method 0x97d5ede5.
//
// Solidity: function rewardsOwing(address _token) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) RewardsOwing(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "rewardsOwing", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsOwing is a free data retrieval call binding the contract method 0x97d5ede5.
//
// Solidity: function rewardsOwing(address _token) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) RewardsOwing(_token common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.RewardsOwing(&_LiquidityPools.CallOpts, _token)
}

// RewardsOwing is a free data retrieval call binding the contract method 0x97d5ede5.
//
// Solidity: function rewardsOwing(address _token) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) RewardsOwing(_token common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.RewardsOwing(&_LiquidityPools.CallOpts, _token)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_LiquidityPools *LiquidityPoolsCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_LiquidityPools *LiquidityPoolsSession) SignerStorage() (common.Address, error) {
	return _LiquidityPools.Contract.SignerStorage(&_LiquidityPools.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_LiquidityPools *LiquidityPoolsCallerSession) SignerStorage() (common.Address, error) {
	return _LiquidityPools.Contract.SignerStorage(&_LiquidityPools.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_LiquidityPools *LiquidityPoolsCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_LiquidityPools *LiquidityPoolsSession) TokenManager() (common.Address, error) {
	return _LiquidityPools.Contract.TokenManager(&_LiquidityPools.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_LiquidityPools *LiquidityPoolsCallerSession) TokenManager() (common.Address, error) {
	return _LiquidityPools.Contract.TokenManager(&_LiquidityPools.CallOpts)
}

// TotalRewardPoints is a free data retrieval call binding the contract method 0x306cc38d.
//
// Solidity: function totalRewardPoints(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) TotalRewardPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "totalRewardPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardPoints is a free data retrieval call binding the contract method 0x306cc38d.
//
// Solidity: function totalRewardPoints(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) TotalRewardPoints(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.TotalRewardPoints(&_LiquidityPools.CallOpts, arg0)
}

// TotalRewardPoints is a free data retrieval call binding the contract method 0x306cc38d.
//
// Solidity: function totalRewardPoints(address ) view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) TotalRewardPoints(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPools.Contract.TotalRewardPoints(&_LiquidityPools.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.AddLiquidity(&_LiquidityPools.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.AddLiquidity(&_LiquidityPools.TransactOpts, _token, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_LiquidityPools *LiquidityPoolsTransactor) AddNativeLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "addNativeLiquidity")
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_LiquidityPools *LiquidityPoolsSession) AddNativeLiquidity() (*types.Transaction, error) {
	return _LiquidityPools.Contract.AddNativeLiquidity(&_LiquidityPools.TransactOpts)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) AddNativeLiquidity() (*types.Transaction, error) {
	return _LiquidityPools.Contract.AddNativeLiquidity(&_LiquidityPools.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _token) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) ClaimRewards(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "claimRewards", _token)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _token) returns()
func (_LiquidityPools *LiquidityPoolsSession) ClaimRewards(_token common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.ClaimRewards(&_LiquidityPools.TransactOpts, _token)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _token) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) ClaimRewards(_token common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.ClaimRewards(&_LiquidityPools.TransactOpts, _token)
}

// DistributeFee is a paid mutator transaction binding the contract method 0x448e520f.
//
// Solidity: function distributeFee(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) DistributeFee(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "distributeFee", _token, _amount)
}

// DistributeFee is a paid mutator transaction binding the contract method 0x448e520f.
//
// Solidity: function distributeFee(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsSession) DistributeFee(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.DistributeFee(&_LiquidityPools.TransactOpts, _token, _amount)
}

// DistributeFee is a paid mutator transaction binding the contract method 0x448e520f.
//
// Solidity: function distributeFee(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) DistributeFee(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.DistributeFee(&_LiquidityPools.TransactOpts, _token, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _signerStorage, address _tokenManager, address _erc20Bridge, address _feeManager, uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) Initialize(opts *bind.TransactOpts, _signerStorage common.Address, _tokenManager common.Address, _erc20Bridge common.Address, _feeManager common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "initialize", _signerStorage, _tokenManager, _erc20Bridge, _feeManager, _feePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _signerStorage, address _tokenManager, address _erc20Bridge, address _feeManager, uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsSession) Initialize(_signerStorage common.Address, _tokenManager common.Address, _erc20Bridge common.Address, _feeManager common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Initialize(&_LiquidityPools.TransactOpts, _signerStorage, _tokenManager, _erc20Bridge, _feeManager, _feePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _signerStorage, address _tokenManager, address _erc20Bridge, address _feeManager, uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) Initialize(_signerStorage common.Address, _tokenManager common.Address, _erc20Bridge common.Address, _feeManager common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Initialize(&_LiquidityPools.TransactOpts, _signerStorage, _tokenManager, _erc20Bridge, _feeManager, _feePercentage)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount) payable returns()
func (_LiquidityPools *LiquidityPoolsTransactor) RemoveLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "removeLiquidity", _token, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount) payable returns()
func (_LiquidityPools *LiquidityPoolsSession) RemoveLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.RemoveLiquidity(&_LiquidityPools.TransactOpts, _token, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount) payable returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) RemoveLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.RemoveLiquidity(&_LiquidityPools.TransactOpts, _token, _amount)
}

// SetERC20Bridge is a paid mutator transaction binding the contract method 0xe88866a5.
//
// Solidity: function setERC20Bridge(address _erc20Bridge) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) SetERC20Bridge(opts *bind.TransactOpts, _erc20Bridge common.Address) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "setERC20Bridge", _erc20Bridge)
}

// SetERC20Bridge is a paid mutator transaction binding the contract method 0xe88866a5.
//
// Solidity: function setERC20Bridge(address _erc20Bridge) returns()
func (_LiquidityPools *LiquidityPoolsSession) SetERC20Bridge(_erc20Bridge common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetERC20Bridge(&_LiquidityPools.TransactOpts, _erc20Bridge)
}

// SetERC20Bridge is a paid mutator transaction binding the contract method 0xe88866a5.
//
// Solidity: function setERC20Bridge(address _erc20Bridge) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) SetERC20Bridge(_erc20Bridge common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetERC20Bridge(&_LiquidityPools.TransactOpts, _erc20Bridge)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) SetFeeManager(opts *bind.TransactOpts, _feeManager common.Address) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "setFeeManager", _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_LiquidityPools *LiquidityPoolsSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetFeeManager(&_LiquidityPools.TransactOpts, _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetFeeManager(&_LiquidityPools.TransactOpts, _feeManager)
}

// SetFeePercentage is a paid mutator transaction binding the contract method 0xae06c1b7.
//
// Solidity: function setFeePercentage(uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) SetFeePercentage(opts *bind.TransactOpts, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "setFeePercentage", _feePercentage)
}

// SetFeePercentage is a paid mutator transaction binding the contract method 0xae06c1b7.
//
// Solidity: function setFeePercentage(uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsSession) SetFeePercentage(_feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetFeePercentage(&_LiquidityPools.TransactOpts, _feePercentage)
}

// SetFeePercentage is a paid mutator transaction binding the contract method 0xae06c1b7.
//
// Solidity: function setFeePercentage(uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) SetFeePercentage(_feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetFeePercentage(&_LiquidityPools.TransactOpts, _feePercentage)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) SetTokenManager(opts *bind.TransactOpts, _tokenManager common.Address) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "setTokenManager", _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_LiquidityPools *LiquidityPoolsSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetTokenManager(&_LiquidityPools.TransactOpts, _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.SetTokenManager(&_LiquidityPools.TransactOpts, _tokenManager)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _receiver, uint256 _transferAmount) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) Transfer(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _transferAmount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "transfer", _token, _receiver, _transferAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _receiver, uint256 _transferAmount) returns()
func (_LiquidityPools *LiquidityPoolsSession) Transfer(_token common.Address, _receiver common.Address, _transferAmount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Transfer(&_LiquidityPools.TransactOpts, _token, _receiver, _transferAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _receiver, uint256 _transferAmount) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) Transfer(_token common.Address, _receiver common.Address, _transferAmount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Transfer(&_LiquidityPools.TransactOpts, _token, _receiver, _transferAmount)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address _receiver, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) TransferNative(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "transferNative", _receiver, _amount)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address _receiver, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsSession) TransferNative(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.TransferNative(&_LiquidityPools.TransactOpts, _receiver, _amount)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address _receiver, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) TransferNative(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.TransferNative(&_LiquidityPools.TransactOpts, _receiver, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPools *LiquidityPoolsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPools.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPools *LiquidityPoolsSession) Receive() (*types.Transaction, error) {
	return _LiquidityPools.Contract.Receive(&_LiquidityPools.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidityPools.Contract.Receive(&_LiquidityPools.TransactOpts)
}

// LiquidityPoolsERC20BridgeUpdatedIterator is returned from FilterERC20BridgeUpdated and is used to iterate over the raw logs and unpacked data for ERC20BridgeUpdated events raised by the LiquidityPools contract.
type LiquidityPoolsERC20BridgeUpdatedIterator struct {
	Event *LiquidityPoolsERC20BridgeUpdated // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsERC20BridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsERC20BridgeUpdated)
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
		it.Event = new(LiquidityPoolsERC20BridgeUpdated)
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
func (it *LiquidityPoolsERC20BridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsERC20BridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsERC20BridgeUpdated represents a ERC20BridgeUpdated event raised by the LiquidityPools contract.
type LiquidityPoolsERC20BridgeUpdated struct {
	Erc20Bridge common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeUpdated is a free log retrieval operation binding the contract event 0x4dcff21652c7d0d1d66be34beea4eeea84364a02f61ed804cfd084772324e63b.
//
// Solidity: event ERC20BridgeUpdated(address erc20Bridge)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterERC20BridgeUpdated(opts *bind.FilterOpts) (*LiquidityPoolsERC20BridgeUpdatedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "ERC20BridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsERC20BridgeUpdatedIterator{contract: _LiquidityPools.contract, event: "ERC20BridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeUpdated is a free log subscription operation binding the contract event 0x4dcff21652c7d0d1d66be34beea4eeea84364a02f61ed804cfd084772324e63b.
//
// Solidity: event ERC20BridgeUpdated(address erc20Bridge)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchERC20BridgeUpdated(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsERC20BridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "ERC20BridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsERC20BridgeUpdated)
				if err := _LiquidityPools.contract.UnpackLog(event, "ERC20BridgeUpdated", log); err != nil {
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

// ParseERC20BridgeUpdated is a log parse operation binding the contract event 0x4dcff21652c7d0d1d66be34beea4eeea84364a02f61ed804cfd084772324e63b.
//
// Solidity: event ERC20BridgeUpdated(address erc20Bridge)
func (_LiquidityPools *LiquidityPoolsFilterer) ParseERC20BridgeUpdated(log types.Log) (*LiquidityPoolsERC20BridgeUpdated, error) {
	event := new(LiquidityPoolsERC20BridgeUpdated)
	if err := _LiquidityPools.contract.UnpackLog(event, "ERC20BridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolsFeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the LiquidityPools contract.
type LiquidityPoolsFeeManagerUpdatedIterator struct {
	Event *LiquidityPoolsFeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsFeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsFeeManagerUpdated)
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
		it.Event = new(LiquidityPoolsFeeManagerUpdated)
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
func (it *LiquidityPoolsFeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsFeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsFeeManagerUpdated represents a FeeManagerUpdated event raised by the LiquidityPools contract.
type LiquidityPoolsFeeManagerUpdated struct {
	FeeManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address feeManager)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterFeeManagerUpdated(opts *bind.FilterOpts) (*LiquidityPoolsFeeManagerUpdatedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsFeeManagerUpdatedIterator{contract: _LiquidityPools.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address feeManager)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsFeeManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsFeeManagerUpdated)
				if err := _LiquidityPools.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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
// Solidity: event FeeManagerUpdated(address feeManager)
func (_LiquidityPools *LiquidityPoolsFilterer) ParseFeeManagerUpdated(log types.Log) (*LiquidityPoolsFeeManagerUpdated, error) {
	event := new(LiquidityPoolsFeeManagerUpdated)
	if err := _LiquidityPools.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolsFeePercentageUpdatedIterator is returned from FilterFeePercentageUpdated and is used to iterate over the raw logs and unpacked data for FeePercentageUpdated events raised by the LiquidityPools contract.
type LiquidityPoolsFeePercentageUpdatedIterator struct {
	Event *LiquidityPoolsFeePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsFeePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsFeePercentageUpdated)
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
		it.Event = new(LiquidityPoolsFeePercentageUpdated)
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
func (it *LiquidityPoolsFeePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsFeePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsFeePercentageUpdated represents a FeePercentageUpdated event raised by the LiquidityPools contract.
type LiquidityPoolsFeePercentageUpdated struct {
	FeePercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeePercentageUpdated is a free log retrieval operation binding the contract event 0x74516f05eb4bd2461d57aa1e935ee553f86a3e02bfed7759f2f772915de3d9be.
//
// Solidity: event FeePercentageUpdated(uint256 feePercentage)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterFeePercentageUpdated(opts *bind.FilterOpts) (*LiquidityPoolsFeePercentageUpdatedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "FeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsFeePercentageUpdatedIterator{contract: _LiquidityPools.contract, event: "FeePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePercentageUpdated is a free log subscription operation binding the contract event 0x74516f05eb4bd2461d57aa1e935ee553f86a3e02bfed7759f2f772915de3d9be.
//
// Solidity: event FeePercentageUpdated(uint256 feePercentage)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchFeePercentageUpdated(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsFeePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "FeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsFeePercentageUpdated)
				if err := _LiquidityPools.contract.UnpackLog(event, "FeePercentageUpdated", log); err != nil {
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

// ParseFeePercentageUpdated is a log parse operation binding the contract event 0x74516f05eb4bd2461d57aa1e935ee553f86a3e02bfed7759f2f772915de3d9be.
//
// Solidity: event FeePercentageUpdated(uint256 feePercentage)
func (_LiquidityPools *LiquidityPoolsFilterer) ParseFeePercentageUpdated(log types.Log) (*LiquidityPoolsFeePercentageUpdated, error) {
	event := new(LiquidityPoolsFeePercentageUpdated)
	if err := _LiquidityPools.contract.UnpackLog(event, "FeePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LiquidityPools contract.
type LiquidityPoolsInitializedIterator struct {
	Event *LiquidityPoolsInitialized // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsInitialized)
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
		it.Event = new(LiquidityPoolsInitialized)
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
func (it *LiquidityPoolsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsInitialized represents a Initialized event raised by the LiquidityPools contract.
type LiquidityPoolsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterInitialized(opts *bind.FilterOpts) (*LiquidityPoolsInitializedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsInitializedIterator{contract: _LiquidityPools.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsInitialized) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsInitialized)
				if err := _LiquidityPools.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LiquidityPools *LiquidityPoolsFilterer) ParseInitialized(log types.Log) (*LiquidityPoolsInitialized, error) {
	event := new(LiquidityPoolsInitialized)
	if err := _LiquidityPools.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolsLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the LiquidityPools contract.
type LiquidityPoolsLiquidityAddedIterator struct {
	Event *LiquidityPoolsLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsLiquidityAdded)
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
		it.Event = new(LiquidityPoolsLiquidityAdded)
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
func (it *LiquidityPoolsLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsLiquidityAdded represents a LiquidityAdded event raised by the LiquidityPools contract.
type LiquidityPoolsLiquidityAdded struct {
	Token   common.Address
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x9d278c56ba6dc86a12eefe6b43112bd6e06648eb4ec0b950ee2d783d40e2acb4.
//
// Solidity: event LiquidityAdded(address token, address account, uint256 amount)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*LiquidityPoolsLiquidityAddedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsLiquidityAddedIterator{contract: _LiquidityPools.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x9d278c56ba6dc86a12eefe6b43112bd6e06648eb4ec0b950ee2d783d40e2acb4.
//
// Solidity: event LiquidityAdded(address token, address account, uint256 amount)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsLiquidityAdded)
				if err := _LiquidityPools.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x9d278c56ba6dc86a12eefe6b43112bd6e06648eb4ec0b950ee2d783d40e2acb4.
//
// Solidity: event LiquidityAdded(address token, address account, uint256 amount)
func (_LiquidityPools *LiquidityPoolsFilterer) ParseLiquidityAdded(log types.Log) (*LiquidityPoolsLiquidityAdded, error) {
	event := new(LiquidityPoolsLiquidityAdded)
	if err := _LiquidityPools.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolsLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the LiquidityPools contract.
type LiquidityPoolsLiquidityRemovedIterator struct {
	Event *LiquidityPoolsLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsLiquidityRemoved)
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
		it.Event = new(LiquidityPoolsLiquidityRemoved)
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
func (it *LiquidityPoolsLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsLiquidityRemoved represents a LiquidityRemoved event raised by the LiquidityPools contract.
type LiquidityPoolsLiquidityRemoved struct {
	Token   common.Address
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0x983e86fda8e7b1e2eae380201830eaf1cac55772e8e39583da349865e8178863.
//
// Solidity: event LiquidityRemoved(address token, address account, uint256 amount)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts) (*LiquidityPoolsLiquidityRemovedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "LiquidityRemoved")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsLiquidityRemovedIterator{contract: _LiquidityPools.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0x983e86fda8e7b1e2eae380201830eaf1cac55772e8e39583da349865e8178863.
//
// Solidity: event LiquidityRemoved(address token, address account, uint256 amount)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsLiquidityRemoved) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "LiquidityRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsLiquidityRemoved)
				if err := _LiquidityPools.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0x983e86fda8e7b1e2eae380201830eaf1cac55772e8e39583da349865e8178863.
//
// Solidity: event LiquidityRemoved(address token, address account, uint256 amount)
func (_LiquidityPools *LiquidityPoolsFilterer) ParseLiquidityRemoved(log types.Log) (*LiquidityPoolsLiquidityRemoved, error) {
	event := new(LiquidityPoolsLiquidityRemoved)
	if err := _LiquidityPools.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolsTokenManagerUpdatedIterator is returned from FilterTokenManagerUpdated and is used to iterate over the raw logs and unpacked data for TokenManagerUpdated events raised by the LiquidityPools contract.
type LiquidityPoolsTokenManagerUpdatedIterator struct {
	Event *LiquidityPoolsTokenManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsTokenManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsTokenManagerUpdated)
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
		it.Event = new(LiquidityPoolsTokenManagerUpdated)
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
func (it *LiquidityPoolsTokenManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsTokenManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsTokenManagerUpdated represents a TokenManagerUpdated event raised by the LiquidityPools contract.
type LiquidityPoolsTokenManagerUpdated struct {
	TokenManager common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenManagerUpdated is a free log retrieval operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address tokenManager)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterTokenManagerUpdated(opts *bind.FilterOpts) (*LiquidityPoolsTokenManagerUpdatedIterator, error) {

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "TokenManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsTokenManagerUpdatedIterator{contract: _LiquidityPools.contract, event: "TokenManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenManagerUpdated is a free log subscription operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address tokenManager)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchTokenManagerUpdated(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsTokenManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "TokenManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsTokenManagerUpdated)
				if err := _LiquidityPools.contract.UnpackLog(event, "TokenManagerUpdated", log); err != nil {
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
// Solidity: event TokenManagerUpdated(address tokenManager)
func (_LiquidityPools *LiquidityPoolsFilterer) ParseTokenManagerUpdated(log types.Log) (*LiquidityPoolsTokenManagerUpdated, error) {
	event := new(LiquidityPoolsTokenManagerUpdated)
	if err := _LiquidityPools.contract.UnpackLog(event, "TokenManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
