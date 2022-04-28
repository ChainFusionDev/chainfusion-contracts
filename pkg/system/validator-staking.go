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

// ValidatorStakingMetaData contains all meta data concerning the ValidatorStaking contract.
var ValidatorStakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"}],\"name\":\"MinimalStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"WithdrawalPeriodUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"}],\"name\":\"setMinimalStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashingVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawalAnnouncements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ValidatorStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorStakingMetaData.ABI instead.
var ValidatorStakingABI = ValidatorStakingMetaData.ABI

// ValidatorStaking is an auto generated Go binding around an Ethereum contract.
type ValidatorStaking struct {
	ValidatorStakingCaller     // Read-only binding to the contract
	ValidatorStakingTransactor // Write-only binding to the contract
	ValidatorStakingFilterer   // Log filterer for contract events
}

// ValidatorStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorStakingSession struct {
	Contract     *ValidatorStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorStakingCallerSession struct {
	Contract *ValidatorStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorStakingTransactorSession struct {
	Contract     *ValidatorStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorStakingRaw struct {
	Contract *ValidatorStaking // Generic contract binding to access the raw methods on
}

// ValidatorStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorStakingCallerRaw struct {
	Contract *ValidatorStakingCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorStakingTransactorRaw struct {
	Contract *ValidatorStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorStaking creates a new instance of ValidatorStaking, bound to a specific deployed contract.
func NewValidatorStaking(address common.Address, backend bind.ContractBackend) (*ValidatorStaking, error) {
	contract, err := bindValidatorStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorStaking{ValidatorStakingCaller: ValidatorStakingCaller{contract: contract}, ValidatorStakingTransactor: ValidatorStakingTransactor{contract: contract}, ValidatorStakingFilterer: ValidatorStakingFilterer{contract: contract}}, nil
}

// NewValidatorStakingCaller creates a new read-only instance of ValidatorStaking, bound to a specific deployed contract.
func NewValidatorStakingCaller(address common.Address, caller bind.ContractCaller) (*ValidatorStakingCaller, error) {
	contract, err := bindValidatorStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingCaller{contract: contract}, nil
}

// NewValidatorStakingTransactor creates a new write-only instance of ValidatorStaking, bound to a specific deployed contract.
func NewValidatorStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorStakingTransactor, error) {
	contract, err := bindValidatorStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingTransactor{contract: contract}, nil
}

// NewValidatorStakingFilterer creates a new log filterer instance of ValidatorStaking, bound to a specific deployed contract.
func NewValidatorStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorStakingFilterer, error) {
	contract, err := bindValidatorStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingFilterer{contract: contract}, nil
}

// bindValidatorStaking binds a generic wrapper to an already deployed contract.
func bindValidatorStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorStaking *ValidatorStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorStaking.Contract.ValidatorStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorStaking *ValidatorStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.ValidatorStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorStaking *ValidatorStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.ValidatorStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorStaking *ValidatorStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorStaking *ValidatorStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorStaking *ValidatorStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.contract.Transact(opts, method, params...)
}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCaller) MinimalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "minimalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingSession) MinimalStake() (*big.Int, error) {
	return _ValidatorStaking.Contract.MinimalStake(&_ValidatorStaking.CallOpts)
}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCallerSession) MinimalStake() (*big.Int, error) {
	return _ValidatorStaking.Contract.MinimalStake(&_ValidatorStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorStaking *ValidatorStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorStaking *ValidatorStakingSession) Owner() (common.Address, error) {
	return _ValidatorStaking.Contract.Owner(&_ValidatorStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorStaking *ValidatorStakingCallerSession) Owner() (common.Address, error) {
	return _ValidatorStaking.Contract.Owner(&_ValidatorStaking.CallOpts)
}

// SlashingCount is a free data retrieval call binding the contract method 0xffe7ecc8.
//
// Solidity: function slashingCount(address ) view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCaller) SlashingCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "slashingCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingCount is a free data retrieval call binding the contract method 0xffe7ecc8.
//
// Solidity: function slashingCount(address ) view returns(uint256)
func (_ValidatorStaking *ValidatorStakingSession) SlashingCount(arg0 common.Address) (*big.Int, error) {
	return _ValidatorStaking.Contract.SlashingCount(&_ValidatorStaking.CallOpts, arg0)
}

// SlashingCount is a free data retrieval call binding the contract method 0xffe7ecc8.
//
// Solidity: function slashingCount(address ) view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCallerSession) SlashingCount(arg0 common.Address) (*big.Int, error) {
	return _ValidatorStaking.Contract.SlashingCount(&_ValidatorStaking.CallOpts, arg0)
}

// SlashingVotes is a free data retrieval call binding the contract method 0x752abe1a.
//
// Solidity: function slashingVotes(address , address ) view returns(bool)
func (_ValidatorStaking *ValidatorStakingCaller) SlashingVotes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "slashingVotes", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashingVotes is a free data retrieval call binding the contract method 0x752abe1a.
//
// Solidity: function slashingVotes(address , address ) view returns(bool)
func (_ValidatorStaking *ValidatorStakingSession) SlashingVotes(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _ValidatorStaking.Contract.SlashingVotes(&_ValidatorStaking.CallOpts, arg0, arg1)
}

// SlashingVotes is a free data retrieval call binding the contract method 0x752abe1a.
//
// Solidity: function slashingVotes(address , address ) view returns(bool)
func (_ValidatorStaking *ValidatorStakingCallerSession) SlashingVotes(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _ValidatorStaking.Contract.SlashingVotes(&_ValidatorStaking.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_ValidatorStaking *ValidatorStakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Validator common.Address
		Stake     *big.Int
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_ValidatorStaking *ValidatorStakingSession) Stakes(arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	return _ValidatorStaking.Contract.Stakes(&_ValidatorStaking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_ValidatorStaking *ValidatorStakingCallerSession) Stakes(arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	return _ValidatorStaking.Contract.Stakes(&_ValidatorStaking.CallOpts, arg0)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorStaking.Contract.ValidatorCount(&_ValidatorStaking.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCallerSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorStaking.Contract.ValidatorCount(&_ValidatorStaking.CallOpts)
}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_ValidatorStaking *ValidatorStakingCaller) WithdrawalAnnouncements(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "withdrawalAnnouncements", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Time   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_ValidatorStaking *ValidatorStakingSession) WithdrawalAnnouncements(arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _ValidatorStaking.Contract.WithdrawalAnnouncements(&_ValidatorStaking.CallOpts, arg0)
}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_ValidatorStaking *ValidatorStakingCallerSession) WithdrawalAnnouncements(arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _ValidatorStaking.Contract.WithdrawalAnnouncements(&_ValidatorStaking.CallOpts, arg0)
}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCaller) WithdrawalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorStaking.contract.Call(opts, &out, "withdrawalPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingSession) WithdrawalPeriod() (*big.Int, error) {
	return _ValidatorStaking.Contract.WithdrawalPeriod(&_ValidatorStaking.CallOpts)
}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_ValidatorStaking *ValidatorStakingCallerSession) WithdrawalPeriod() (*big.Int, error) {
	return _ValidatorStaking.Contract.WithdrawalPeriod(&_ValidatorStaking.CallOpts)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_ValidatorStaking *ValidatorStakingTransactor) AnnounceWithdrawal(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "announceWithdrawal", _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_ValidatorStaking *ValidatorStakingSession) AnnounceWithdrawal(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.AnnounceWithdrawal(&_ValidatorStaking.TransactOpts, _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) AnnounceWithdrawal(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.AnnounceWithdrawal(&_ValidatorStaking.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_ValidatorStaking *ValidatorStakingTransactor) Initialize(opts *bind.TransactOpts, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "initialize", _minimalStake, _withdrawalPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_ValidatorStaking *ValidatorStakingSession) Initialize(_minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Initialize(&_ValidatorStaking.TransactOpts, _minimalStake, _withdrawalPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) Initialize(_minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Initialize(&_ValidatorStaking.TransactOpts, _minimalStake, _withdrawalPeriod)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorStaking *ValidatorStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorStaking *ValidatorStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorStaking.Contract.RenounceOwnership(&_ValidatorStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorStaking.Contract.RenounceOwnership(&_ValidatorStaking.TransactOpts)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_ValidatorStaking *ValidatorStakingTransactor) SetMinimalStake(opts *bind.TransactOpts, _minimalStake *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "setMinimalStake", _minimalStake)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_ValidatorStaking *ValidatorStakingSession) SetMinimalStake(_minimalStake *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.SetMinimalStake(&_ValidatorStaking.TransactOpts, _minimalStake)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) SetMinimalStake(_minimalStake *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.SetMinimalStake(&_ValidatorStaking.TransactOpts, _minimalStake)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_ValidatorStaking *ValidatorStakingTransactor) SetWithdrawalPeriod(opts *bind.TransactOpts, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "setWithdrawalPeriod", _withdrawalPeriod)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_ValidatorStaking *ValidatorStakingSession) SetWithdrawalPeriod(_withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.SetWithdrawalPeriod(&_ValidatorStaking.TransactOpts, _withdrawalPeriod)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) SetWithdrawalPeriod(_withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.SetWithdrawalPeriod(&_ValidatorStaking.TransactOpts, _withdrawalPeriod)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_ValidatorStaking *ValidatorStakingTransactor) Slash(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "slash", _validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_ValidatorStaking *ValidatorStakingSession) Slash(_validator common.Address) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Slash(&_ValidatorStaking.TransactOpts, _validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) Slash(_validator common.Address) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Slash(&_ValidatorStaking.TransactOpts, _validator)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_ValidatorStaking *ValidatorStakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_ValidatorStaking *ValidatorStakingSession) Stake() (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Stake(&_ValidatorStaking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) Stake() (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Stake(&_ValidatorStaking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorStaking *ValidatorStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorStaking *ValidatorStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.TransferOwnership(&_ValidatorStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorStaking.Contract.TransferOwnership(&_ValidatorStaking.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorStaking *ValidatorStakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorStaking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorStaking *ValidatorStakingSession) Withdraw() (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Withdraw(&_ValidatorStaking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorStaking *ValidatorStakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ValidatorStaking.Contract.Withdraw(&_ValidatorStaking.TransactOpts)
}

// ValidatorStakingMinimalStakeUpdatedIterator is returned from FilterMinimalStakeUpdated and is used to iterate over the raw logs and unpacked data for MinimalStakeUpdated events raised by the ValidatorStaking contract.
type ValidatorStakingMinimalStakeUpdatedIterator struct {
	Event *ValidatorStakingMinimalStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingMinimalStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingMinimalStakeUpdated)
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
		it.Event = new(ValidatorStakingMinimalStakeUpdated)
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
func (it *ValidatorStakingMinimalStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingMinimalStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingMinimalStakeUpdated represents a MinimalStakeUpdated event raised by the ValidatorStaking contract.
type ValidatorStakingMinimalStakeUpdated struct {
	MinimalStake *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinimalStakeUpdated is a free log retrieval operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_ValidatorStaking *ValidatorStakingFilterer) FilterMinimalStakeUpdated(opts *bind.FilterOpts) (*ValidatorStakingMinimalStakeUpdatedIterator, error) {

	logs, sub, err := _ValidatorStaking.contract.FilterLogs(opts, "MinimalStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingMinimalStakeUpdatedIterator{contract: _ValidatorStaking.contract, event: "MinimalStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimalStakeUpdated is a free log subscription operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_ValidatorStaking *ValidatorStakingFilterer) WatchMinimalStakeUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorStakingMinimalStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _ValidatorStaking.contract.WatchLogs(opts, "MinimalStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingMinimalStakeUpdated)
				if err := _ValidatorStaking.contract.UnpackLog(event, "MinimalStakeUpdated", log); err != nil {
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

// ParseMinimalStakeUpdated is a log parse operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_ValidatorStaking *ValidatorStakingFilterer) ParseMinimalStakeUpdated(log types.Log) (*ValidatorStakingMinimalStakeUpdated, error) {
	event := new(ValidatorStakingMinimalStakeUpdated)
	if err := _ValidatorStaking.contract.UnpackLog(event, "MinimalStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorStaking contract.
type ValidatorStakingOwnershipTransferredIterator struct {
	Event *ValidatorStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingOwnershipTransferred)
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
		it.Event = new(ValidatorStakingOwnershipTransferred)
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
func (it *ValidatorStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorStaking contract.
type ValidatorStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorStaking *ValidatorStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingOwnershipTransferredIterator{contract: _ValidatorStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorStaking *ValidatorStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingOwnershipTransferred)
				if err := _ValidatorStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValidatorStaking *ValidatorStakingFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorStakingOwnershipTransferred, error) {
	event := new(ValidatorStakingOwnershipTransferred)
	if err := _ValidatorStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingWithdrawalPeriodUpdatedIterator is returned from FilterWithdrawalPeriodUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalPeriodUpdated events raised by the ValidatorStaking contract.
type ValidatorStakingWithdrawalPeriodUpdatedIterator struct {
	Event *ValidatorStakingWithdrawalPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingWithdrawalPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingWithdrawalPeriodUpdated)
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
		it.Event = new(ValidatorStakingWithdrawalPeriodUpdated)
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
func (it *ValidatorStakingWithdrawalPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingWithdrawalPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingWithdrawalPeriodUpdated represents a WithdrawalPeriodUpdated event raised by the ValidatorStaking contract.
type ValidatorStakingWithdrawalPeriodUpdated struct {
	WithdrawalPeriod *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalPeriodUpdated is a free log retrieval operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_ValidatorStaking *ValidatorStakingFilterer) FilterWithdrawalPeriodUpdated(opts *bind.FilterOpts) (*ValidatorStakingWithdrawalPeriodUpdatedIterator, error) {

	logs, sub, err := _ValidatorStaking.contract.FilterLogs(opts, "WithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingWithdrawalPeriodUpdatedIterator{contract: _ValidatorStaking.contract, event: "WithdrawalPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalPeriodUpdated is a free log subscription operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_ValidatorStaking *ValidatorStakingFilterer) WatchWithdrawalPeriodUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorStakingWithdrawalPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _ValidatorStaking.contract.WatchLogs(opts, "WithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingWithdrawalPeriodUpdated)
				if err := _ValidatorStaking.contract.UnpackLog(event, "WithdrawalPeriodUpdated", log); err != nil {
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

// ParseWithdrawalPeriodUpdated is a log parse operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_ValidatorStaking *ValidatorStakingFilterer) ParseWithdrawalPeriodUpdated(log types.Log) (*ValidatorStakingWithdrawalPeriodUpdated, error) {
	event := new(ValidatorStakingWithdrawalPeriodUpdated)
	if err := _ValidatorStaking.contract.UnpackLog(event, "WithdrawalPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
