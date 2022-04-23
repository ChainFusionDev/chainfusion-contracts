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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collectedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityPositions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"rewardsOwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"setFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"contractTokenManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// BASEDIVISOR is a free data retrieval call binding the contract method 0x3ed77114.
//
// Solidity: function BASE_DIVISOR() view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCaller) BASEDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "BASE_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEDIVISOR is a free data retrieval call binding the contract method 0x3ed77114.
//
// Solidity: function BASE_DIVISOR() view returns(uint256)
func (_LiquidityPools *LiquidityPoolsSession) BASEDIVISOR() (*big.Int, error) {
	return _LiquidityPools.Contract.BASEDIVISOR(&_LiquidityPools.CallOpts)
}

// BASEDIVISOR is a free data retrieval call binding the contract method 0x3ed77114.
//
// Solidity: function BASE_DIVISOR() view returns(uint256)
func (_LiquidityPools *LiquidityPoolsCallerSession) BASEDIVISOR() (*big.Int, error) {
	return _LiquidityPools.Contract.BASEDIVISOR(&_LiquidityPools.CallOpts)
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

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LiquidityPools *LiquidityPoolsCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LiquidityPools *LiquidityPoolsSession) Bridge() (common.Address, error) {
	return _LiquidityPools.Contract.Bridge(&_LiquidityPools.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LiquidityPools *LiquidityPoolsCallerSession) Bridge() (common.Address, error) {
	return _LiquidityPools.Contract.Bridge(&_LiquidityPools.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPools *LiquidityPoolsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPools.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPools *LiquidityPoolsSession) Owner() (common.Address, error) {
	return _LiquidityPools.Contract.Owner(&_LiquidityPools.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPools *LiquidityPoolsCallerSession) Owner() (common.Address, error) {
	return _LiquidityPools.Contract.Owner(&_LiquidityPools.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _tokenManager, address _bridge, uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) Initialize(opts *bind.TransactOpts, _tokenManager common.Address, _bridge common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "initialize", _tokenManager, _bridge, _feePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _tokenManager, address _bridge, uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsSession) Initialize(_tokenManager common.Address, _bridge common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Initialize(&_LiquidityPools.TransactOpts, _tokenManager, _bridge, _feePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _tokenManager, address _bridge, uint256 _feePercentage) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) Initialize(_tokenManager common.Address, _bridge common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Initialize(&_LiquidityPools.TransactOpts, _tokenManager, _bridge, _feePercentage)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) RemoveLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "removeLiquidity", _token, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsSession) RemoveLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.RemoveLiquidity(&_LiquidityPools.TransactOpts, _token, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) RemoveLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.RemoveLiquidity(&_LiquidityPools.TransactOpts, _token, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPools *LiquidityPoolsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPools *LiquidityPoolsSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPools.Contract.RenounceOwnership(&_LiquidityPools.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPools.Contract.RenounceOwnership(&_LiquidityPools.TransactOpts)
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
// Solidity: function transfer(address _token, address _receiver, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) Transfer(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "transfer", _token, _receiver, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _receiver, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsSession) Transfer(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Transfer(&_LiquidityPools.TransactOpts, _token, _receiver, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _receiver, uint256 _amount) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) Transfer(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPools.Contract.Transfer(&_LiquidityPools.TransactOpts, _token, _receiver, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPools *LiquidityPoolsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPools.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPools *LiquidityPoolsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.TransferOwnership(&_LiquidityPools.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPools *LiquidityPoolsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPools.Contract.TransferOwnership(&_LiquidityPools.TransactOpts, newOwner)
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

// LiquidityPoolsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidityPools contract.
type LiquidityPoolsOwnershipTransferredIterator struct {
	Event *LiquidityPoolsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolsOwnershipTransferred)
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
		it.Event = new(LiquidityPoolsOwnershipTransferred)
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
func (it *LiquidityPoolsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolsOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidityPools contract.
type LiquidityPoolsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPools *LiquidityPoolsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidityPoolsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPools.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolsOwnershipTransferredIterator{contract: _LiquidityPools.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPools *LiquidityPoolsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidityPoolsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPools.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolsOwnershipTransferred)
				if err := _LiquidityPools.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LiquidityPools *LiquidityPoolsFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidityPoolsOwnershipTransferred, error) {
	event := new(LiquidityPoolsOwnershipTransferred)
	if err := _LiquidityPools.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
