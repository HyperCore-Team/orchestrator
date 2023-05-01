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
	_ = abi.ConvertType
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unhaltDurationParam\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"administratorDelayParam\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"softDelayParam\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"initialGuardians\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Halted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"PendingAdministrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PendingGuardians\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PendingTokenInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTss\",\"type\":\"address\"}],\"name\":\"PendingTss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RegisteredRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"RevokedRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdministrator\",\"type\":\"address\"}],\"name\":\"SetAdministrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetAdministratorDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetAllowKeyGen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SetConfirmationsToFinality\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SetEstimatedBlockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SetGuardians\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetSoftDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SetTokenInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTss\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTss\",\"type\":\"address\"}],\"name\":\"SetTss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetUnhaltDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unhalted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unwrapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionsNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administratorDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowKeyGen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmationsToFinality\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractDeploymentHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedBlockTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"guardiansVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"halt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isHalted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAdministratorDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSoftDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minUnhaltDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newGuardians\",\"type\":\"address[]\"}],\"name\":\"nominateGuardians\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"proposeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeemsInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paramsHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"revokeRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"setAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setAdministratorDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAllowKeyGen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"}],\"name\":\"setConfirmationsToFinality\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockTime\",\"type\":\"uint64\"}],\"name\":\"setEstimatedBlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setSoftDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemDelay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bridgeable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"redeemable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOwned\",\"type\":\"bool\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTss\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"oldSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newSignature\",\"type\":\"bytes\"}],\"name\":\"setTss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setUnhaltDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"softDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"timeChallengesInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paramsHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemDelay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bridgeable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"redeemable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"owned\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tss\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhalt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhaltDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhaltedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_Bridge *BridgeCaller) ActionsNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "actionsNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_Bridge *BridgeSession) ActionsNonce() (*big.Int, error) {
	return _Bridge.Contract.ActionsNonce(&_Bridge.CallOpts)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_Bridge *BridgeCallerSession) ActionsNonce() (*big.Int, error) {
	return _Bridge.Contract.ActionsNonce(&_Bridge.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Bridge *BridgeCaller) Administrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "administrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Bridge *BridgeSession) Administrator() (common.Address, error) {
	return _Bridge.Contract.Administrator(&_Bridge.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Bridge *BridgeCallerSession) Administrator() (common.Address, error) {
	return _Bridge.Contract.Administrator(&_Bridge.CallOpts)
}

// AdministratorDelay is a free data retrieval call binding the contract method 0x7f34ad69.
//
// Solidity: function administratorDelay() view returns(uint256)
func (_Bridge *BridgeCaller) AdministratorDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "administratorDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdministratorDelay is a free data retrieval call binding the contract method 0x7f34ad69.
//
// Solidity: function administratorDelay() view returns(uint256)
func (_Bridge *BridgeSession) AdministratorDelay() (*big.Int, error) {
	return _Bridge.Contract.AdministratorDelay(&_Bridge.CallOpts)
}

// AdministratorDelay is a free data retrieval call binding the contract method 0x7f34ad69.
//
// Solidity: function administratorDelay() view returns(uint256)
func (_Bridge *BridgeCallerSession) AdministratorDelay() (*big.Int, error) {
	return _Bridge.Contract.AdministratorDelay(&_Bridge.CallOpts)
}

// AllowKeyGen is a free data retrieval call binding the contract method 0x25b90865.
//
// Solidity: function allowKeyGen() view returns(bool)
func (_Bridge *BridgeCaller) AllowKeyGen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowKeyGen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowKeyGen is a free data retrieval call binding the contract method 0x25b90865.
//
// Solidity: function allowKeyGen() view returns(bool)
func (_Bridge *BridgeSession) AllowKeyGen() (bool, error) {
	return _Bridge.Contract.AllowKeyGen(&_Bridge.CallOpts)
}

// AllowKeyGen is a free data retrieval call binding the contract method 0x25b90865.
//
// Solidity: function allowKeyGen() view returns(bool)
func (_Bridge *BridgeCallerSession) AllowKeyGen() (bool, error) {
	return _Bridge.Contract.AllowKeyGen(&_Bridge.CallOpts)
}

// ConfirmationsToFinality is a free data retrieval call binding the contract method 0x5e2b8902.
//
// Solidity: function confirmationsToFinality() view returns(uint64)
func (_Bridge *BridgeCaller) ConfirmationsToFinality(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "confirmationsToFinality")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConfirmationsToFinality is a free data retrieval call binding the contract method 0x5e2b8902.
//
// Solidity: function confirmationsToFinality() view returns(uint64)
func (_Bridge *BridgeSession) ConfirmationsToFinality() (uint64, error) {
	return _Bridge.Contract.ConfirmationsToFinality(&_Bridge.CallOpts)
}

// ConfirmationsToFinality is a free data retrieval call binding the contract method 0x5e2b8902.
//
// Solidity: function confirmationsToFinality() view returns(uint64)
func (_Bridge *BridgeCallerSession) ConfirmationsToFinality() (uint64, error) {
	return _Bridge.Contract.ConfirmationsToFinality(&_Bridge.CallOpts)
}

// ContractDeploymentHeight is a free data retrieval call binding the contract method 0xdeb005f7.
//
// Solidity: function contractDeploymentHeight() view returns(uint256)
func (_Bridge *BridgeCaller) ContractDeploymentHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "contractDeploymentHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractDeploymentHeight is a free data retrieval call binding the contract method 0xdeb005f7.
//
// Solidity: function contractDeploymentHeight() view returns(uint256)
func (_Bridge *BridgeSession) ContractDeploymentHeight() (*big.Int, error) {
	return _Bridge.Contract.ContractDeploymentHeight(&_Bridge.CallOpts)
}

// ContractDeploymentHeight is a free data retrieval call binding the contract method 0xdeb005f7.
//
// Solidity: function contractDeploymentHeight() view returns(uint256)
func (_Bridge *BridgeCallerSession) ContractDeploymentHeight() (*big.Int, error) {
	return _Bridge.Contract.ContractDeploymentHeight(&_Bridge.CallOpts)
}

// EstimatedBlockTime is a free data retrieval call binding the contract method 0xe4ebdd5e.
//
// Solidity: function estimatedBlockTime() view returns(uint64)
func (_Bridge *BridgeCaller) EstimatedBlockTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "estimatedBlockTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EstimatedBlockTime is a free data retrieval call binding the contract method 0xe4ebdd5e.
//
// Solidity: function estimatedBlockTime() view returns(uint64)
func (_Bridge *BridgeSession) EstimatedBlockTime() (uint64, error) {
	return _Bridge.Contract.EstimatedBlockTime(&_Bridge.CallOpts)
}

// EstimatedBlockTime is a free data retrieval call binding the contract method 0xe4ebdd5e.
//
// Solidity: function estimatedBlockTime() view returns(uint64)
func (_Bridge *BridgeCallerSession) EstimatedBlockTime() (uint64, error) {
	return _Bridge.Contract.EstimatedBlockTime(&_Bridge.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0xf560c734.
//
// Solidity: function guardians(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) Guardians(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "guardians", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0xf560c734.
//
// Solidity: function guardians(uint256 ) view returns(address)
func (_Bridge *BridgeSession) Guardians(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.Guardians(&_Bridge.CallOpts, arg0)
}

// Guardians is a free data retrieval call binding the contract method 0xf560c734.
//
// Solidity: function guardians(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) Guardians(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.Guardians(&_Bridge.CallOpts, arg0)
}

// GuardiansVotes is a free data retrieval call binding the contract method 0x2b4fdaf0.
//
// Solidity: function guardiansVotes(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) GuardiansVotes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "guardiansVotes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardiansVotes is a free data retrieval call binding the contract method 0x2b4fdaf0.
//
// Solidity: function guardiansVotes(uint256 ) view returns(address)
func (_Bridge *BridgeSession) GuardiansVotes(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.GuardiansVotes(&_Bridge.CallOpts, arg0)
}

// GuardiansVotes is a free data retrieval call binding the contract method 0x2b4fdaf0.
//
// Solidity: function guardiansVotes(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) GuardiansVotes(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.GuardiansVotes(&_Bridge.CallOpts, arg0)
}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Bridge *BridgeCaller) Halted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "halted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Bridge *BridgeSession) Halted() (bool, error) {
	return _Bridge.Contract.Halted(&_Bridge.CallOpts)
}

// Halted is a free data retrieval call binding the contract method 0xb9b8af0b.
//
// Solidity: function halted() view returns(bool)
func (_Bridge *BridgeCallerSession) Halted() (bool, error) {
	return _Bridge.Contract.Halted(&_Bridge.CallOpts)
}

// IsHalted is a free data retrieval call binding the contract method 0xc7ff1584.
//
// Solidity: function isHalted() view returns(bool)
func (_Bridge *BridgeCaller) IsHalted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isHalted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHalted is a free data retrieval call binding the contract method 0xc7ff1584.
//
// Solidity: function isHalted() view returns(bool)
func (_Bridge *BridgeSession) IsHalted() (bool, error) {
	return _Bridge.Contract.IsHalted(&_Bridge.CallOpts)
}

// IsHalted is a free data retrieval call binding the contract method 0xc7ff1584.
//
// Solidity: function isHalted() view returns(bool)
func (_Bridge *BridgeCallerSession) IsHalted() (bool, error) {
	return _Bridge.Contract.IsHalted(&_Bridge.CallOpts)
}

// MinAdministratorDelay is a free data retrieval call binding the contract method 0x4969bfac.
//
// Solidity: function minAdministratorDelay() view returns(uint256)
func (_Bridge *BridgeCaller) MinAdministratorDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minAdministratorDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAdministratorDelay is a free data retrieval call binding the contract method 0x4969bfac.
//
// Solidity: function minAdministratorDelay() view returns(uint256)
func (_Bridge *BridgeSession) MinAdministratorDelay() (*big.Int, error) {
	return _Bridge.Contract.MinAdministratorDelay(&_Bridge.CallOpts)
}

// MinAdministratorDelay is a free data retrieval call binding the contract method 0x4969bfac.
//
// Solidity: function minAdministratorDelay() view returns(uint256)
func (_Bridge *BridgeCallerSession) MinAdministratorDelay() (*big.Int, error) {
	return _Bridge.Contract.MinAdministratorDelay(&_Bridge.CallOpts)
}

// MinSoftDelay is a free data retrieval call binding the contract method 0xb4edfaf3.
//
// Solidity: function minSoftDelay() view returns(uint256)
func (_Bridge *BridgeCaller) MinSoftDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minSoftDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSoftDelay is a free data retrieval call binding the contract method 0xb4edfaf3.
//
// Solidity: function minSoftDelay() view returns(uint256)
func (_Bridge *BridgeSession) MinSoftDelay() (*big.Int, error) {
	return _Bridge.Contract.MinSoftDelay(&_Bridge.CallOpts)
}

// MinSoftDelay is a free data retrieval call binding the contract method 0xb4edfaf3.
//
// Solidity: function minSoftDelay() view returns(uint256)
func (_Bridge *BridgeCallerSession) MinSoftDelay() (*big.Int, error) {
	return _Bridge.Contract.MinSoftDelay(&_Bridge.CallOpts)
}

// MinUnhaltDuration is a free data retrieval call binding the contract method 0x1ea9e496.
//
// Solidity: function minUnhaltDuration() view returns(uint256)
func (_Bridge *BridgeCaller) MinUnhaltDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minUnhaltDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinUnhaltDuration is a free data retrieval call binding the contract method 0x1ea9e496.
//
// Solidity: function minUnhaltDuration() view returns(uint256)
func (_Bridge *BridgeSession) MinUnhaltDuration() (*big.Int, error) {
	return _Bridge.Contract.MinUnhaltDuration(&_Bridge.CallOpts)
}

// MinUnhaltDuration is a free data retrieval call binding the contract method 0x1ea9e496.
//
// Solidity: function minUnhaltDuration() view returns(uint256)
func (_Bridge *BridgeCallerSession) MinUnhaltDuration() (*big.Int, error) {
	return _Bridge.Contract.MinUnhaltDuration(&_Bridge.CallOpts)
}

// RedeemsInfo is a free data retrieval call binding the contract method 0xa75f8096.
//
// Solidity: function redeemsInfo(uint256 ) view returns(uint256 blockNumber, bytes32 paramsHash)
func (_Bridge *BridgeCaller) RedeemsInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNumber *big.Int
	ParamsHash  [32]byte
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "redeemsInfo", arg0)

	outstruct := new(struct {
		BlockNumber *big.Int
		ParamsHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ParamsHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RedeemsInfo is a free data retrieval call binding the contract method 0xa75f8096.
//
// Solidity: function redeemsInfo(uint256 ) view returns(uint256 blockNumber, bytes32 paramsHash)
func (_Bridge *BridgeSession) RedeemsInfo(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	ParamsHash  [32]byte
}, error) {
	return _Bridge.Contract.RedeemsInfo(&_Bridge.CallOpts, arg0)
}

// RedeemsInfo is a free data retrieval call binding the contract method 0xa75f8096.
//
// Solidity: function redeemsInfo(uint256 ) view returns(uint256 blockNumber, bytes32 paramsHash)
func (_Bridge *BridgeCallerSession) RedeemsInfo(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	ParamsHash  [32]byte
}, error) {
	return _Bridge.Contract.RedeemsInfo(&_Bridge.CallOpts, arg0)
}

// SoftDelay is a free data retrieval call binding the contract method 0x08c3c888.
//
// Solidity: function softDelay() view returns(uint256)
func (_Bridge *BridgeCaller) SoftDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "softDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SoftDelay is a free data retrieval call binding the contract method 0x08c3c888.
//
// Solidity: function softDelay() view returns(uint256)
func (_Bridge *BridgeSession) SoftDelay() (*big.Int, error) {
	return _Bridge.Contract.SoftDelay(&_Bridge.CallOpts)
}

// SoftDelay is a free data retrieval call binding the contract method 0x08c3c888.
//
// Solidity: function softDelay() view returns(uint256)
func (_Bridge *BridgeCallerSession) SoftDelay() (*big.Int, error) {
	return _Bridge.Contract.SoftDelay(&_Bridge.CallOpts)
}

// TimeChallengesInfo is a free data retrieval call binding the contract method 0xa6fd410e.
//
// Solidity: function timeChallengesInfo(string ) view returns(uint256 blockNumber, bytes32 paramsHash)
func (_Bridge *BridgeCaller) TimeChallengesInfo(opts *bind.CallOpts, arg0 string) (struct {
	BlockNumber *big.Int
	ParamsHash  [32]byte
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "timeChallengesInfo", arg0)

	outstruct := new(struct {
		BlockNumber *big.Int
		ParamsHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ParamsHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// TimeChallengesInfo is a free data retrieval call binding the contract method 0xa6fd410e.
//
// Solidity: function timeChallengesInfo(string ) view returns(uint256 blockNumber, bytes32 paramsHash)
func (_Bridge *BridgeSession) TimeChallengesInfo(arg0 string) (struct {
	BlockNumber *big.Int
	ParamsHash  [32]byte
}, error) {
	return _Bridge.Contract.TimeChallengesInfo(&_Bridge.CallOpts, arg0)
}

// TimeChallengesInfo is a free data retrieval call binding the contract method 0xa6fd410e.
//
// Solidity: function timeChallengesInfo(string ) view returns(uint256 blockNumber, bytes32 paramsHash)
func (_Bridge *BridgeCallerSession) TimeChallengesInfo(arg0 string) (struct {
	BlockNumber *big.Int
	ParamsHash  [32]byte
}, error) {
	return _Bridge.Contract.TimeChallengesInfo(&_Bridge.CallOpts, arg0)
}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 minAmount, uint256 redeemDelay, bool bridgeable, bool redeemable, bool owned)
func (_Bridge *BridgeCaller) TokensInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinAmount   *big.Int
	RedeemDelay *big.Int
	Bridgeable  bool
	Redeemable  bool
	Owned       bool
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokensInfo", arg0)

	outstruct := new(struct {
		MinAmount   *big.Int
		RedeemDelay *big.Int
		Bridgeable  bool
		Redeemable  bool
		Owned       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RedeemDelay = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Bridgeable = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Redeemable = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Owned = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 minAmount, uint256 redeemDelay, bool bridgeable, bool redeemable, bool owned)
func (_Bridge *BridgeSession) TokensInfo(arg0 common.Address) (struct {
	MinAmount   *big.Int
	RedeemDelay *big.Int
	Bridgeable  bool
	Redeemable  bool
	Owned       bool
}, error) {
	return _Bridge.Contract.TokensInfo(&_Bridge.CallOpts, arg0)
}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 minAmount, uint256 redeemDelay, bool bridgeable, bool redeemable, bool owned)
func (_Bridge *BridgeCallerSession) TokensInfo(arg0 common.Address) (struct {
	MinAmount   *big.Int
	RedeemDelay *big.Int
	Bridgeable  bool
	Redeemable  bool
	Owned       bool
}, error) {
	return _Bridge.Contract.TokensInfo(&_Bridge.CallOpts, arg0)
}

// Tss is a free data retrieval call binding the contract method 0x6e6dbb51.
//
// Solidity: function tss() view returns(address)
func (_Bridge *BridgeCaller) Tss(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tss")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tss is a free data retrieval call binding the contract method 0x6e6dbb51.
//
// Solidity: function tss() view returns(address)
func (_Bridge *BridgeSession) Tss() (common.Address, error) {
	return _Bridge.Contract.Tss(&_Bridge.CallOpts)
}

// Tss is a free data retrieval call binding the contract method 0x6e6dbb51.
//
// Solidity: function tss() view returns(address)
func (_Bridge *BridgeCallerSession) Tss() (common.Address, error) {
	return _Bridge.Contract.Tss(&_Bridge.CallOpts)
}

// UnhaltDuration is a free data retrieval call binding the contract method 0xdc22c2ac.
//
// Solidity: function unhaltDuration() view returns(uint256)
func (_Bridge *BridgeCaller) UnhaltDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "unhaltDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnhaltDuration is a free data retrieval call binding the contract method 0xdc22c2ac.
//
// Solidity: function unhaltDuration() view returns(uint256)
func (_Bridge *BridgeSession) UnhaltDuration() (*big.Int, error) {
	return _Bridge.Contract.UnhaltDuration(&_Bridge.CallOpts)
}

// UnhaltDuration is a free data retrieval call binding the contract method 0xdc22c2ac.
//
// Solidity: function unhaltDuration() view returns(uint256)
func (_Bridge *BridgeCallerSession) UnhaltDuration() (*big.Int, error) {
	return _Bridge.Contract.UnhaltDuration(&_Bridge.CallOpts)
}

// UnhaltedAt is a free data retrieval call binding the contract method 0x69d11381.
//
// Solidity: function unhaltedAt() view returns(uint256)
func (_Bridge *BridgeCaller) UnhaltedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "unhaltedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnhaltedAt is a free data retrieval call binding the contract method 0x69d11381.
//
// Solidity: function unhaltedAt() view returns(uint256)
func (_Bridge *BridgeSession) UnhaltedAt() (*big.Int, error) {
	return _Bridge.Contract.UnhaltedAt(&_Bridge.CallOpts)
}

// UnhaltedAt is a free data retrieval call binding the contract method 0x69d11381.
//
// Solidity: function unhaltedAt() view returns(uint256)
func (_Bridge *BridgeCallerSession) UnhaltedAt() (*big.Int, error) {
	return _Bridge.Contract.UnhaltedAt(&_Bridge.CallOpts)
}

// VotesCount is a free data retrieval call binding the contract method 0xd209cd68.
//
// Solidity: function votesCount(address ) view returns(uint256)
func (_Bridge *BridgeCaller) VotesCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "votesCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesCount is a free data retrieval call binding the contract method 0xd209cd68.
//
// Solidity: function votesCount(address ) view returns(uint256)
func (_Bridge *BridgeSession) VotesCount(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.VotesCount(&_Bridge.CallOpts, arg0)
}

// VotesCount is a free data retrieval call binding the contract method 0xd209cd68.
//
// Solidity: function votesCount(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) VotesCount(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.VotesCount(&_Bridge.CallOpts, arg0)
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_Bridge *BridgeTransactor) Emergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "emergency")
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_Bridge *BridgeSession) Emergency() (*types.Transaction, error) {
	return _Bridge.Contract.Emergency(&_Bridge.TransactOpts)
}

// Emergency is a paid mutator transaction binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() returns()
func (_Bridge *BridgeTransactorSession) Emergency() (*types.Transaction, error) {
	return _Bridge.Contract.Emergency(&_Bridge.TransactOpts)
}

// Halt is a paid mutator transaction binding the contract method 0x944e7cb1.
//
// Solidity: function halt(bytes signature) returns()
func (_Bridge *BridgeTransactor) Halt(opts *bind.TransactOpts, signature []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "halt", signature)
}

// Halt is a paid mutator transaction binding the contract method 0x944e7cb1.
//
// Solidity: function halt(bytes signature) returns()
func (_Bridge *BridgeSession) Halt(signature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Halt(&_Bridge.TransactOpts, signature)
}

// Halt is a paid mutator transaction binding the contract method 0x944e7cb1.
//
// Solidity: function halt(bytes signature) returns()
func (_Bridge *BridgeTransactorSession) Halt(signature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Halt(&_Bridge.TransactOpts, signature)
}

// NominateGuardians is a paid mutator transaction binding the contract method 0x661b2753.
//
// Solidity: function nominateGuardians(address[] newGuardians) returns()
func (_Bridge *BridgeTransactor) NominateGuardians(opts *bind.TransactOpts, newGuardians []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "nominateGuardians", newGuardians)
}

// NominateGuardians is a paid mutator transaction binding the contract method 0x661b2753.
//
// Solidity: function nominateGuardians(address[] newGuardians) returns()
func (_Bridge *BridgeSession) NominateGuardians(newGuardians []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.NominateGuardians(&_Bridge.TransactOpts, newGuardians)
}

// NominateGuardians is a paid mutator transaction binding the contract method 0x661b2753.
//
// Solidity: function nominateGuardians(address[] newGuardians) returns()
func (_Bridge *BridgeTransactorSession) NominateGuardians(newGuardians []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.NominateGuardians(&_Bridge.TransactOpts, newGuardians)
}

// ProposeAdministrator is a paid mutator transaction binding the contract method 0x10db6bb7.
//
// Solidity: function proposeAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeTransactor) ProposeAdministrator(opts *bind.TransactOpts, newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "proposeAdministrator", newAdministrator)
}

// ProposeAdministrator is a paid mutator transaction binding the contract method 0x10db6bb7.
//
// Solidity: function proposeAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeSession) ProposeAdministrator(newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ProposeAdministrator(&_Bridge.TransactOpts, newAdministrator)
}

// ProposeAdministrator is a paid mutator transaction binding the contract method 0x10db6bb7.
//
// Solidity: function proposeAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeTransactorSession) ProposeAdministrator(newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ProposeAdministrator(&_Bridge.TransactOpts, newAdministrator)
}

// Redeem is a paid mutator transaction binding the contract method 0xfb51c178.
//
// Solidity: function redeem(address to, address token, uint256 amount, uint256 nonce, bytes signature) returns()
func (_Bridge *BridgeTransactor) Redeem(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "redeem", to, token, amount, nonce, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0xfb51c178.
//
// Solidity: function redeem(address to, address token, uint256 amount, uint256 nonce, bytes signature) returns()
func (_Bridge *BridgeSession) Redeem(to common.Address, token common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Redeem(&_Bridge.TransactOpts, to, token, amount, nonce, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0xfb51c178.
//
// Solidity: function redeem(address to, address token, uint256 amount, uint256 nonce, bytes signature) returns()
func (_Bridge *BridgeTransactorSession) Redeem(to common.Address, token common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Redeem(&_Bridge.TransactOpts, to, token, amount, nonce, signature)
}

// RevokeRedeems is a paid mutator transaction binding the contract method 0x6717bcf9.
//
// Solidity: function revokeRedeems(uint256[] nonces) returns()
func (_Bridge *BridgeTransactor) RevokeRedeems(opts *bind.TransactOpts, nonces []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRedeems", nonces)
}

// RevokeRedeems is a paid mutator transaction binding the contract method 0x6717bcf9.
//
// Solidity: function revokeRedeems(uint256[] nonces) returns()
func (_Bridge *BridgeSession) RevokeRedeems(nonces []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRedeems(&_Bridge.TransactOpts, nonces)
}

// RevokeRedeems is a paid mutator transaction binding the contract method 0x6717bcf9.
//
// Solidity: function revokeRedeems(uint256[] nonces) returns()
func (_Bridge *BridgeTransactorSession) RevokeRedeems(nonces []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRedeems(&_Bridge.TransactOpts, nonces)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeTransactor) SetAdministrator(opts *bind.TransactOpts, newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAdministrator", newAdministrator)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeSession) SetAdministrator(newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdministrator(&_Bridge.TransactOpts, newAdministrator)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeTransactorSession) SetAdministrator(newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdministrator(&_Bridge.TransactOpts, newAdministrator)
}

// SetAdministratorDelay is a paid mutator transaction binding the contract method 0xb1d72e29.
//
// Solidity: function setAdministratorDelay(uint256 delay) returns()
func (_Bridge *BridgeTransactor) SetAdministratorDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAdministratorDelay", delay)
}

// SetAdministratorDelay is a paid mutator transaction binding the contract method 0xb1d72e29.
//
// Solidity: function setAdministratorDelay(uint256 delay) returns()
func (_Bridge *BridgeSession) SetAdministratorDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdministratorDelay(&_Bridge.TransactOpts, delay)
}

// SetAdministratorDelay is a paid mutator transaction binding the contract method 0xb1d72e29.
//
// Solidity: function setAdministratorDelay(uint256 delay) returns()
func (_Bridge *BridgeTransactorSession) SetAdministratorDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdministratorDelay(&_Bridge.TransactOpts, delay)
}

// SetAllowKeyGen is a paid mutator transaction binding the contract method 0x949ffc7b.
//
// Solidity: function setAllowKeyGen(bool value) returns()
func (_Bridge *BridgeTransactor) SetAllowKeyGen(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAllowKeyGen", value)
}

// SetAllowKeyGen is a paid mutator transaction binding the contract method 0x949ffc7b.
//
// Solidity: function setAllowKeyGen(bool value) returns()
func (_Bridge *BridgeSession) SetAllowKeyGen(value bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetAllowKeyGen(&_Bridge.TransactOpts, value)
}

// SetAllowKeyGen is a paid mutator transaction binding the contract method 0x949ffc7b.
//
// Solidity: function setAllowKeyGen(bool value) returns()
func (_Bridge *BridgeTransactorSession) SetAllowKeyGen(value bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetAllowKeyGen(&_Bridge.TransactOpts, value)
}

// SetConfirmationsToFinality is a paid mutator transaction binding the contract method 0x53a1dba9.
//
// Solidity: function setConfirmationsToFinality(uint64 confirmations) returns()
func (_Bridge *BridgeTransactor) SetConfirmationsToFinality(opts *bind.TransactOpts, confirmations uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setConfirmationsToFinality", confirmations)
}

// SetConfirmationsToFinality is a paid mutator transaction binding the contract method 0x53a1dba9.
//
// Solidity: function setConfirmationsToFinality(uint64 confirmations) returns()
func (_Bridge *BridgeSession) SetConfirmationsToFinality(confirmations uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetConfirmationsToFinality(&_Bridge.TransactOpts, confirmations)
}

// SetConfirmationsToFinality is a paid mutator transaction binding the contract method 0x53a1dba9.
//
// Solidity: function setConfirmationsToFinality(uint64 confirmations) returns()
func (_Bridge *BridgeTransactorSession) SetConfirmationsToFinality(confirmations uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetConfirmationsToFinality(&_Bridge.TransactOpts, confirmations)
}

// SetEstimatedBlockTime is a paid mutator transaction binding the contract method 0xb6f474e4.
//
// Solidity: function setEstimatedBlockTime(uint64 blockTime) returns()
func (_Bridge *BridgeTransactor) SetEstimatedBlockTime(opts *bind.TransactOpts, blockTime uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setEstimatedBlockTime", blockTime)
}

// SetEstimatedBlockTime is a paid mutator transaction binding the contract method 0xb6f474e4.
//
// Solidity: function setEstimatedBlockTime(uint64 blockTime) returns()
func (_Bridge *BridgeSession) SetEstimatedBlockTime(blockTime uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetEstimatedBlockTime(&_Bridge.TransactOpts, blockTime)
}

// SetEstimatedBlockTime is a paid mutator transaction binding the contract method 0xb6f474e4.
//
// Solidity: function setEstimatedBlockTime(uint64 blockTime) returns()
func (_Bridge *BridgeTransactorSession) SetEstimatedBlockTime(blockTime uint64) (*types.Transaction, error) {
	return _Bridge.Contract.SetEstimatedBlockTime(&_Bridge.TransactOpts, blockTime)
}

// SetSoftDelay is a paid mutator transaction binding the contract method 0x64d9e2e4.
//
// Solidity: function setSoftDelay(uint256 delay) returns()
func (_Bridge *BridgeTransactor) SetSoftDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setSoftDelay", delay)
}

// SetSoftDelay is a paid mutator transaction binding the contract method 0x64d9e2e4.
//
// Solidity: function setSoftDelay(uint256 delay) returns()
func (_Bridge *BridgeSession) SetSoftDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetSoftDelay(&_Bridge.TransactOpts, delay)
}

// SetSoftDelay is a paid mutator transaction binding the contract method 0x64d9e2e4.
//
// Solidity: function setSoftDelay(uint256 delay) returns()
func (_Bridge *BridgeTransactorSession) SetSoftDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetSoftDelay(&_Bridge.TransactOpts, delay)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x39cbc5b9.
//
// Solidity: function setTokenInfo(address token, uint256 minAmount, uint256 redeemDelay, bool bridgeable, bool redeemable, bool isOwned) returns()
func (_Bridge *BridgeTransactor) SetTokenInfo(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, redeemDelay *big.Int, bridgeable bool, redeemable bool, isOwned bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTokenInfo", token, minAmount, redeemDelay, bridgeable, redeemable, isOwned)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x39cbc5b9.
//
// Solidity: function setTokenInfo(address token, uint256 minAmount, uint256 redeemDelay, bool bridgeable, bool redeemable, bool isOwned) returns()
func (_Bridge *BridgeSession) SetTokenInfo(token common.Address, minAmount *big.Int, redeemDelay *big.Int, bridgeable bool, redeemable bool, isOwned bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenInfo(&_Bridge.TransactOpts, token, minAmount, redeemDelay, bridgeable, redeemable, isOwned)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x39cbc5b9.
//
// Solidity: function setTokenInfo(address token, uint256 minAmount, uint256 redeemDelay, bool bridgeable, bool redeemable, bool isOwned) returns()
func (_Bridge *BridgeTransactorSession) SetTokenInfo(token common.Address, minAmount *big.Int, redeemDelay *big.Int, bridgeable bool, redeemable bool, isOwned bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenInfo(&_Bridge.TransactOpts, token, minAmount, redeemDelay, bridgeable, redeemable, isOwned)
}

// SetTss is a paid mutator transaction binding the contract method 0x72461d7b.
//
// Solidity: function setTss(address newTss, bytes oldSignature, bytes newSignature) returns()
func (_Bridge *BridgeTransactor) SetTss(opts *bind.TransactOpts, newTss common.Address, oldSignature []byte, newSignature []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTss", newTss, oldSignature, newSignature)
}

// SetTss is a paid mutator transaction binding the contract method 0x72461d7b.
//
// Solidity: function setTss(address newTss, bytes oldSignature, bytes newSignature) returns()
func (_Bridge *BridgeSession) SetTss(newTss common.Address, oldSignature []byte, newSignature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.SetTss(&_Bridge.TransactOpts, newTss, oldSignature, newSignature)
}

// SetTss is a paid mutator transaction binding the contract method 0x72461d7b.
//
// Solidity: function setTss(address newTss, bytes oldSignature, bytes newSignature) returns()
func (_Bridge *BridgeTransactorSession) SetTss(newTss common.Address, oldSignature []byte, newSignature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.SetTss(&_Bridge.TransactOpts, newTss, oldSignature, newSignature)
}

// SetUnhaltDuration is a paid mutator transaction binding the contract method 0xbca9525d.
//
// Solidity: function setUnhaltDuration(uint256 duration) returns()
func (_Bridge *BridgeTransactor) SetUnhaltDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setUnhaltDuration", duration)
}

// SetUnhaltDuration is a paid mutator transaction binding the contract method 0xbca9525d.
//
// Solidity: function setUnhaltDuration(uint256 duration) returns()
func (_Bridge *BridgeSession) SetUnhaltDuration(duration *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetUnhaltDuration(&_Bridge.TransactOpts, duration)
}

// SetUnhaltDuration is a paid mutator transaction binding the contract method 0xbca9525d.
//
// Solidity: function setUnhaltDuration(uint256 duration) returns()
func (_Bridge *BridgeTransactorSession) SetUnhaltDuration(duration *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetUnhaltDuration(&_Bridge.TransactOpts, duration)
}

// Unhalt is a paid mutator transaction binding the contract method 0xcb3e64fd.
//
// Solidity: function unhalt() returns()
func (_Bridge *BridgeTransactor) Unhalt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unhalt")
}

// Unhalt is a paid mutator transaction binding the contract method 0xcb3e64fd.
//
// Solidity: function unhalt() returns()
func (_Bridge *BridgeSession) Unhalt() (*types.Transaction, error) {
	return _Bridge.Contract.Unhalt(&_Bridge.TransactOpts)
}

// Unhalt is a paid mutator transaction binding the contract method 0xcb3e64fd.
//
// Solidity: function unhalt() returns()
func (_Bridge *BridgeTransactorSession) Unhalt() (*types.Transaction, error) {
	return _Bridge.Contract.Unhalt(&_Bridge.TransactOpts)
}

// Unwrap is a paid mutator transaction binding the contract method 0x7a07b345.
//
// Solidity: function unwrap(address token, uint256 amount, string to) returns()
func (_Bridge *BridgeTransactor) Unwrap(opts *bind.TransactOpts, token common.Address, amount *big.Int, to string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unwrap", token, amount, to)
}

// Unwrap is a paid mutator transaction binding the contract method 0x7a07b345.
//
// Solidity: function unwrap(address token, uint256 amount, string to) returns()
func (_Bridge *BridgeSession) Unwrap(token common.Address, amount *big.Int, to string) (*types.Transaction, error) {
	return _Bridge.Contract.Unwrap(&_Bridge.TransactOpts, token, amount, to)
}

// Unwrap is a paid mutator transaction binding the contract method 0x7a07b345.
//
// Solidity: function unwrap(address token, uint256 amount, string to) returns()
func (_Bridge *BridgeTransactorSession) Unwrap(token common.Address, amount *big.Int, to string) (*types.Transaction, error) {
	return _Bridge.Contract.Unwrap(&_Bridge.TransactOpts, token, amount, to)
}

// BridgeHaltedIterator is returned from FilterHalted and is used to iterate over the raw logs and unpacked data for Halted events raised by the Bridge contract.
type BridgeHaltedIterator struct {
	Event *BridgeHalted // Event containing the contract specifics and raw log

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
func (it *BridgeHaltedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHalted)
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
		it.Event = new(BridgeHalted)
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
func (it *BridgeHaltedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHaltedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHalted represents a Halted event raised by the Bridge contract.
type BridgeHalted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHalted is a free log retrieval operation binding the contract event 0x1ee9080f6b55ca44ce58681c8162e6c1ac1c47e1da791a4a1c1ec6186d8af1f3.
//
// Solidity: event Halted()
func (_Bridge *BridgeFilterer) FilterHalted(opts *bind.FilterOpts) (*BridgeHaltedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Halted")
	if err != nil {
		return nil, err
	}
	return &BridgeHaltedIterator{contract: _Bridge.contract, event: "Halted", logs: logs, sub: sub}, nil
}

// WatchHalted is a free log subscription operation binding the contract event 0x1ee9080f6b55ca44ce58681c8162e6c1ac1c47e1da791a4a1c1ec6186d8af1f3.
//
// Solidity: event Halted()
func (_Bridge *BridgeFilterer) WatchHalted(opts *bind.WatchOpts, sink chan<- *BridgeHalted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Halted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHalted)
				if err := _Bridge.contract.UnpackLog(event, "Halted", log); err != nil {
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

// ParseHalted is a log parse operation binding the contract event 0x1ee9080f6b55ca44ce58681c8162e6c1ac1c47e1da791a4a1c1ec6186d8af1f3.
//
// Solidity: event Halted()
func (_Bridge *BridgeFilterer) ParseHalted(log types.Log) (*BridgeHalted, error) {
	event := new(BridgeHalted)
	if err := _Bridge.contract.UnpackLog(event, "Halted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePendingAdministratorIterator is returned from FilterPendingAdministrator and is used to iterate over the raw logs and unpacked data for PendingAdministrator events raised by the Bridge contract.
type BridgePendingAdministratorIterator struct {
	Event *BridgePendingAdministrator // Event containing the contract specifics and raw log

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
func (it *BridgePendingAdministratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePendingAdministrator)
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
		it.Event = new(BridgePendingAdministrator)
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
func (it *BridgePendingAdministratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePendingAdministratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePendingAdministrator represents a PendingAdministrator event raised by the Bridge contract.
type BridgePendingAdministrator struct {
	NewAdministrator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPendingAdministrator is a free log retrieval operation binding the contract event 0x2eb4029903a72f19daef97dcf670179630438575b5ffcd45a5537b7e5a5dafc9.
//
// Solidity: event PendingAdministrator(address indexed newAdministrator)
func (_Bridge *BridgeFilterer) FilterPendingAdministrator(opts *bind.FilterOpts, newAdministrator []common.Address) (*BridgePendingAdministratorIterator, error) {

	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PendingAdministrator", newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return &BridgePendingAdministratorIterator{contract: _Bridge.contract, event: "PendingAdministrator", logs: logs, sub: sub}, nil
}

// WatchPendingAdministrator is a free log subscription operation binding the contract event 0x2eb4029903a72f19daef97dcf670179630438575b5ffcd45a5537b7e5a5dafc9.
//
// Solidity: event PendingAdministrator(address indexed newAdministrator)
func (_Bridge *BridgeFilterer) WatchPendingAdministrator(opts *bind.WatchOpts, sink chan<- *BridgePendingAdministrator, newAdministrator []common.Address) (event.Subscription, error) {

	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PendingAdministrator", newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePendingAdministrator)
				if err := _Bridge.contract.UnpackLog(event, "PendingAdministrator", log); err != nil {
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

// ParsePendingAdministrator is a log parse operation binding the contract event 0x2eb4029903a72f19daef97dcf670179630438575b5ffcd45a5537b7e5a5dafc9.
//
// Solidity: event PendingAdministrator(address indexed newAdministrator)
func (_Bridge *BridgeFilterer) ParsePendingAdministrator(log types.Log) (*BridgePendingAdministrator, error) {
	event := new(BridgePendingAdministrator)
	if err := _Bridge.contract.UnpackLog(event, "PendingAdministrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePendingGuardiansIterator is returned from FilterPendingGuardians and is used to iterate over the raw logs and unpacked data for PendingGuardians events raised by the Bridge contract.
type BridgePendingGuardiansIterator struct {
	Event *BridgePendingGuardians // Event containing the contract specifics and raw log

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
func (it *BridgePendingGuardiansIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePendingGuardians)
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
		it.Event = new(BridgePendingGuardians)
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
func (it *BridgePendingGuardiansIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePendingGuardiansIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePendingGuardians represents a PendingGuardians event raised by the Bridge contract.
type BridgePendingGuardians struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPendingGuardians is a free log retrieval operation binding the contract event 0xc9dcac3c9c0c7b1e477f2e4a899bd3feb9b15f0903484a9f120bb44a3b87efd8.
//
// Solidity: event PendingGuardians()
func (_Bridge *BridgeFilterer) FilterPendingGuardians(opts *bind.FilterOpts) (*BridgePendingGuardiansIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PendingGuardians")
	if err != nil {
		return nil, err
	}
	return &BridgePendingGuardiansIterator{contract: _Bridge.contract, event: "PendingGuardians", logs: logs, sub: sub}, nil
}

// WatchPendingGuardians is a free log subscription operation binding the contract event 0xc9dcac3c9c0c7b1e477f2e4a899bd3feb9b15f0903484a9f120bb44a3b87efd8.
//
// Solidity: event PendingGuardians()
func (_Bridge *BridgeFilterer) WatchPendingGuardians(opts *bind.WatchOpts, sink chan<- *BridgePendingGuardians) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PendingGuardians")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePendingGuardians)
				if err := _Bridge.contract.UnpackLog(event, "PendingGuardians", log); err != nil {
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

// ParsePendingGuardians is a log parse operation binding the contract event 0xc9dcac3c9c0c7b1e477f2e4a899bd3feb9b15f0903484a9f120bb44a3b87efd8.
//
// Solidity: event PendingGuardians()
func (_Bridge *BridgeFilterer) ParsePendingGuardians(log types.Log) (*BridgePendingGuardians, error) {
	event := new(BridgePendingGuardians)
	if err := _Bridge.contract.UnpackLog(event, "PendingGuardians", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePendingTokenInfoIterator is returned from FilterPendingTokenInfo and is used to iterate over the raw logs and unpacked data for PendingTokenInfo events raised by the Bridge contract.
type BridgePendingTokenInfoIterator struct {
	Event *BridgePendingTokenInfo // Event containing the contract specifics and raw log

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
func (it *BridgePendingTokenInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePendingTokenInfo)
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
		it.Event = new(BridgePendingTokenInfo)
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
func (it *BridgePendingTokenInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePendingTokenInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePendingTokenInfo represents a PendingTokenInfo event raised by the Bridge contract.
type BridgePendingTokenInfo struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPendingTokenInfo is a free log retrieval operation binding the contract event 0xb29cc088f961a7dbe14c460642a2d3afbbccc7af072a5e38e1f9227c5e5e7222.
//
// Solidity: event PendingTokenInfo(address indexed token)
func (_Bridge *BridgeFilterer) FilterPendingTokenInfo(opts *bind.FilterOpts, token []common.Address) (*BridgePendingTokenInfoIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PendingTokenInfo", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgePendingTokenInfoIterator{contract: _Bridge.contract, event: "PendingTokenInfo", logs: logs, sub: sub}, nil
}

// WatchPendingTokenInfo is a free log subscription operation binding the contract event 0xb29cc088f961a7dbe14c460642a2d3afbbccc7af072a5e38e1f9227c5e5e7222.
//
// Solidity: event PendingTokenInfo(address indexed token)
func (_Bridge *BridgeFilterer) WatchPendingTokenInfo(opts *bind.WatchOpts, sink chan<- *BridgePendingTokenInfo, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PendingTokenInfo", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePendingTokenInfo)
				if err := _Bridge.contract.UnpackLog(event, "PendingTokenInfo", log); err != nil {
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

// ParsePendingTokenInfo is a log parse operation binding the contract event 0xb29cc088f961a7dbe14c460642a2d3afbbccc7af072a5e38e1f9227c5e5e7222.
//
// Solidity: event PendingTokenInfo(address indexed token)
func (_Bridge *BridgeFilterer) ParsePendingTokenInfo(log types.Log) (*BridgePendingTokenInfo, error) {
	event := new(BridgePendingTokenInfo)
	if err := _Bridge.contract.UnpackLog(event, "PendingTokenInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePendingTssIterator is returned from FilterPendingTss and is used to iterate over the raw logs and unpacked data for PendingTss events raised by the Bridge contract.
type BridgePendingTssIterator struct {
	Event *BridgePendingTss // Event containing the contract specifics and raw log

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
func (it *BridgePendingTssIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePendingTss)
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
		it.Event = new(BridgePendingTss)
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
func (it *BridgePendingTssIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePendingTssIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePendingTss represents a PendingTss event raised by the Bridge contract.
type BridgePendingTss struct {
	NewTss common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPendingTss is a free log retrieval operation binding the contract event 0xad1c846290c418eec330f68f49d58cf79ff618f063a767abf9ee1c895f9d0f77.
//
// Solidity: event PendingTss(address indexed newTss)
func (_Bridge *BridgeFilterer) FilterPendingTss(opts *bind.FilterOpts, newTss []common.Address) (*BridgePendingTssIterator, error) {

	var newTssRule []interface{}
	for _, newTssItem := range newTss {
		newTssRule = append(newTssRule, newTssItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PendingTss", newTssRule)
	if err != nil {
		return nil, err
	}
	return &BridgePendingTssIterator{contract: _Bridge.contract, event: "PendingTss", logs: logs, sub: sub}, nil
}

// WatchPendingTss is a free log subscription operation binding the contract event 0xad1c846290c418eec330f68f49d58cf79ff618f063a767abf9ee1c895f9d0f77.
//
// Solidity: event PendingTss(address indexed newTss)
func (_Bridge *BridgeFilterer) WatchPendingTss(opts *bind.WatchOpts, sink chan<- *BridgePendingTss, newTss []common.Address) (event.Subscription, error) {

	var newTssRule []interface{}
	for _, newTssItem := range newTss {
		newTssRule = append(newTssRule, newTssItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PendingTss", newTssRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePendingTss)
				if err := _Bridge.contract.UnpackLog(event, "PendingTss", log); err != nil {
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

// ParsePendingTss is a log parse operation binding the contract event 0xad1c846290c418eec330f68f49d58cf79ff618f063a767abf9ee1c895f9d0f77.
//
// Solidity: event PendingTss(address indexed newTss)
func (_Bridge *BridgeFilterer) ParsePendingTss(log types.Log) (*BridgePendingTss, error) {
	event := new(BridgePendingTss)
	if err := _Bridge.contract.UnpackLog(event, "PendingTss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the Bridge contract.
type BridgeRedeemedIterator struct {
	Event *BridgeRedeemed // Event containing the contract specifics and raw log

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
func (it *BridgeRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRedeemed)
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
		it.Event = new(BridgeRedeemed)
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
func (it *BridgeRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRedeemed represents a Redeemed event raised by the Bridge contract.
type BridgeRedeemed struct {
	Nonce  *big.Int
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x18737e07ba2aac9c230bdd7119bde1c2d51cef17a2910224f55819e4b0651ea1.
//
// Solidity: event Redeemed(uint256 indexed nonce, address indexed to, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterRedeemed(opts *bind.FilterOpts, nonce []*big.Int, to []common.Address, token []common.Address) (*BridgeRedeemedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Redeemed", nonceRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRedeemedIterator{contract: _Bridge.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x18737e07ba2aac9c230bdd7119bde1c2d51cef17a2910224f55819e4b0651ea1.
//
// Solidity: event Redeemed(uint256 indexed nonce, address indexed to, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *BridgeRedeemed, nonce []*big.Int, to []common.Address, token []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Redeemed", nonceRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRedeemed)
				if err := _Bridge.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x18737e07ba2aac9c230bdd7119bde1c2d51cef17a2910224f55819e4b0651ea1.
//
// Solidity: event Redeemed(uint256 indexed nonce, address indexed to, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseRedeemed(log types.Log) (*BridgeRedeemed, error) {
	event := new(BridgeRedeemed)
	if err := _Bridge.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRegisteredRedeemIterator is returned from FilterRegisteredRedeem and is used to iterate over the raw logs and unpacked data for RegisteredRedeem events raised by the Bridge contract.
type BridgeRegisteredRedeemIterator struct {
	Event *BridgeRegisteredRedeem // Event containing the contract specifics and raw log

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
func (it *BridgeRegisteredRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegisteredRedeem)
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
		it.Event = new(BridgeRegisteredRedeem)
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
func (it *BridgeRegisteredRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegisteredRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegisteredRedeem represents a RegisteredRedeem event raised by the Bridge contract.
type BridgeRegisteredRedeem struct {
	Nonce  *big.Int
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredRedeem is a free log retrieval operation binding the contract event 0x2bbb1da11bfccd157ef0abea11bc0c5b42f9816aedf585fd1603b11f45b887be.
//
// Solidity: event RegisteredRedeem(uint256 indexed nonce, address indexed to, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterRegisteredRedeem(opts *bind.FilterOpts, nonce []*big.Int, to []common.Address, token []common.Address) (*BridgeRegisteredRedeemIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RegisteredRedeem", nonceRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRegisteredRedeemIterator{contract: _Bridge.contract, event: "RegisteredRedeem", logs: logs, sub: sub}, nil
}

// WatchRegisteredRedeem is a free log subscription operation binding the contract event 0x2bbb1da11bfccd157ef0abea11bc0c5b42f9816aedf585fd1603b11f45b887be.
//
// Solidity: event RegisteredRedeem(uint256 indexed nonce, address indexed to, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchRegisteredRedeem(opts *bind.WatchOpts, sink chan<- *BridgeRegisteredRedeem, nonce []*big.Int, to []common.Address, token []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RegisteredRedeem", nonceRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegisteredRedeem)
				if err := _Bridge.contract.UnpackLog(event, "RegisteredRedeem", log); err != nil {
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

// ParseRegisteredRedeem is a log parse operation binding the contract event 0x2bbb1da11bfccd157ef0abea11bc0c5b42f9816aedf585fd1603b11f45b887be.
//
// Solidity: event RegisteredRedeem(uint256 indexed nonce, address indexed to, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseRegisteredRedeem(log types.Log) (*BridgeRegisteredRedeem, error) {
	event := new(BridgeRegisteredRedeem)
	if err := _Bridge.contract.UnpackLog(event, "RegisteredRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRevokedRedeemIterator is returned from FilterRevokedRedeem and is used to iterate over the raw logs and unpacked data for RevokedRedeem events raised by the Bridge contract.
type BridgeRevokedRedeemIterator struct {
	Event *BridgeRevokedRedeem // Event containing the contract specifics and raw log

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
func (it *BridgeRevokedRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRevokedRedeem)
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
		it.Event = new(BridgeRevokedRedeem)
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
func (it *BridgeRevokedRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRevokedRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRevokedRedeem represents a RevokedRedeem event raised by the Bridge contract.
type BridgeRevokedRedeem struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevokedRedeem is a free log retrieval operation binding the contract event 0xda456efa0805c22cd348af0a9a79bc8ab7c618fd5d9b761d72b327a65e98b3c1.
//
// Solidity: event RevokedRedeem(uint256 indexed nonce)
func (_Bridge *BridgeFilterer) FilterRevokedRedeem(opts *bind.FilterOpts, nonce []*big.Int) (*BridgeRevokedRedeemIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RevokedRedeem", nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRevokedRedeemIterator{contract: _Bridge.contract, event: "RevokedRedeem", logs: logs, sub: sub}, nil
}

// WatchRevokedRedeem is a free log subscription operation binding the contract event 0xda456efa0805c22cd348af0a9a79bc8ab7c618fd5d9b761d72b327a65e98b3c1.
//
// Solidity: event RevokedRedeem(uint256 indexed nonce)
func (_Bridge *BridgeFilterer) WatchRevokedRedeem(opts *bind.WatchOpts, sink chan<- *BridgeRevokedRedeem, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RevokedRedeem", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRevokedRedeem)
				if err := _Bridge.contract.UnpackLog(event, "RevokedRedeem", log); err != nil {
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

// ParseRevokedRedeem is a log parse operation binding the contract event 0xda456efa0805c22cd348af0a9a79bc8ab7c618fd5d9b761d72b327a65e98b3c1.
//
// Solidity: event RevokedRedeem(uint256 indexed nonce)
func (_Bridge *BridgeFilterer) ParseRevokedRedeem(log types.Log) (*BridgeRevokedRedeem, error) {
	event := new(BridgeRevokedRedeem)
	if err := _Bridge.contract.UnpackLog(event, "RevokedRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetAdministratorIterator is returned from FilterSetAdministrator and is used to iterate over the raw logs and unpacked data for SetAdministrator events raised by the Bridge contract.
type BridgeSetAdministratorIterator struct {
	Event *BridgeSetAdministrator // Event containing the contract specifics and raw log

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
func (it *BridgeSetAdministratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetAdministrator)
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
		it.Event = new(BridgeSetAdministrator)
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
func (it *BridgeSetAdministratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetAdministratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetAdministrator represents a SetAdministrator event raised by the Bridge contract.
type BridgeSetAdministrator struct {
	NewAdministrator common.Address
	OldAdministrator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetAdministrator is a free log retrieval operation binding the contract event 0x525117b0c654ea4d9950e296eca92b77aa3c6fdeba6217801c6f32a0ed7a0922.
//
// Solidity: event SetAdministrator(address indexed newAdministrator, address oldAdministrator)
func (_Bridge *BridgeFilterer) FilterSetAdministrator(opts *bind.FilterOpts, newAdministrator []common.Address) (*BridgeSetAdministratorIterator, error) {

	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetAdministrator", newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSetAdministratorIterator{contract: _Bridge.contract, event: "SetAdministrator", logs: logs, sub: sub}, nil
}

// WatchSetAdministrator is a free log subscription operation binding the contract event 0x525117b0c654ea4d9950e296eca92b77aa3c6fdeba6217801c6f32a0ed7a0922.
//
// Solidity: event SetAdministrator(address indexed newAdministrator, address oldAdministrator)
func (_Bridge *BridgeFilterer) WatchSetAdministrator(opts *bind.WatchOpts, sink chan<- *BridgeSetAdministrator, newAdministrator []common.Address) (event.Subscription, error) {

	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetAdministrator", newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetAdministrator)
				if err := _Bridge.contract.UnpackLog(event, "SetAdministrator", log); err != nil {
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

// ParseSetAdministrator is a log parse operation binding the contract event 0x525117b0c654ea4d9950e296eca92b77aa3c6fdeba6217801c6f32a0ed7a0922.
//
// Solidity: event SetAdministrator(address indexed newAdministrator, address oldAdministrator)
func (_Bridge *BridgeFilterer) ParseSetAdministrator(log types.Log) (*BridgeSetAdministrator, error) {
	event := new(BridgeSetAdministrator)
	if err := _Bridge.contract.UnpackLog(event, "SetAdministrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetAdministratorDelayIterator is returned from FilterSetAdministratorDelay and is used to iterate over the raw logs and unpacked data for SetAdministratorDelay events raised by the Bridge contract.
type BridgeSetAdministratorDelayIterator struct {
	Event *BridgeSetAdministratorDelay // Event containing the contract specifics and raw log

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
func (it *BridgeSetAdministratorDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetAdministratorDelay)
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
		it.Event = new(BridgeSetAdministratorDelay)
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
func (it *BridgeSetAdministratorDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetAdministratorDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetAdministratorDelay represents a SetAdministratorDelay event raised by the Bridge contract.
type BridgeSetAdministratorDelay struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetAdministratorDelay is a free log retrieval operation binding the contract event 0x0cfb9bdc3485531feee3d5fc58e83d442a0f5de6b44a0cb275ece047cc3d691c.
//
// Solidity: event SetAdministratorDelay(uint256 arg0)
func (_Bridge *BridgeFilterer) FilterSetAdministratorDelay(opts *bind.FilterOpts) (*BridgeSetAdministratorDelayIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetAdministratorDelay")
	if err != nil {
		return nil, err
	}
	return &BridgeSetAdministratorDelayIterator{contract: _Bridge.contract, event: "SetAdministratorDelay", logs: logs, sub: sub}, nil
}

// WatchSetAdministratorDelay is a free log subscription operation binding the contract event 0x0cfb9bdc3485531feee3d5fc58e83d442a0f5de6b44a0cb275ece047cc3d691c.
//
// Solidity: event SetAdministratorDelay(uint256 arg0)
func (_Bridge *BridgeFilterer) WatchSetAdministratorDelay(opts *bind.WatchOpts, sink chan<- *BridgeSetAdministratorDelay) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetAdministratorDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetAdministratorDelay)
				if err := _Bridge.contract.UnpackLog(event, "SetAdministratorDelay", log); err != nil {
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

// ParseSetAdministratorDelay is a log parse operation binding the contract event 0x0cfb9bdc3485531feee3d5fc58e83d442a0f5de6b44a0cb275ece047cc3d691c.
//
// Solidity: event SetAdministratorDelay(uint256 arg0)
func (_Bridge *BridgeFilterer) ParseSetAdministratorDelay(log types.Log) (*BridgeSetAdministratorDelay, error) {
	event := new(BridgeSetAdministratorDelay)
	if err := _Bridge.contract.UnpackLog(event, "SetAdministratorDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetAllowKeyGenIterator is returned from FilterSetAllowKeyGen and is used to iterate over the raw logs and unpacked data for SetAllowKeyGen events raised by the Bridge contract.
type BridgeSetAllowKeyGenIterator struct {
	Event *BridgeSetAllowKeyGen // Event containing the contract specifics and raw log

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
func (it *BridgeSetAllowKeyGenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetAllowKeyGen)
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
		it.Event = new(BridgeSetAllowKeyGen)
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
func (it *BridgeSetAllowKeyGenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetAllowKeyGenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetAllowKeyGen represents a SetAllowKeyGen event raised by the Bridge contract.
type BridgeSetAllowKeyGen struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetAllowKeyGen is a free log retrieval operation binding the contract event 0xffa5f6b5dcc0b57eaa767458e1fba30bf9e2498b9b85114f874d7d5913c7d860.
//
// Solidity: event SetAllowKeyGen(bool arg0)
func (_Bridge *BridgeFilterer) FilterSetAllowKeyGen(opts *bind.FilterOpts) (*BridgeSetAllowKeyGenIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetAllowKeyGen")
	if err != nil {
		return nil, err
	}
	return &BridgeSetAllowKeyGenIterator{contract: _Bridge.contract, event: "SetAllowKeyGen", logs: logs, sub: sub}, nil
}

// WatchSetAllowKeyGen is a free log subscription operation binding the contract event 0xffa5f6b5dcc0b57eaa767458e1fba30bf9e2498b9b85114f874d7d5913c7d860.
//
// Solidity: event SetAllowKeyGen(bool arg0)
func (_Bridge *BridgeFilterer) WatchSetAllowKeyGen(opts *bind.WatchOpts, sink chan<- *BridgeSetAllowKeyGen) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetAllowKeyGen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetAllowKeyGen)
				if err := _Bridge.contract.UnpackLog(event, "SetAllowKeyGen", log); err != nil {
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

// ParseSetAllowKeyGen is a log parse operation binding the contract event 0xffa5f6b5dcc0b57eaa767458e1fba30bf9e2498b9b85114f874d7d5913c7d860.
//
// Solidity: event SetAllowKeyGen(bool arg0)
func (_Bridge *BridgeFilterer) ParseSetAllowKeyGen(log types.Log) (*BridgeSetAllowKeyGen, error) {
	event := new(BridgeSetAllowKeyGen)
	if err := _Bridge.contract.UnpackLog(event, "SetAllowKeyGen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetConfirmationsToFinalityIterator is returned from FilterSetConfirmationsToFinality and is used to iterate over the raw logs and unpacked data for SetConfirmationsToFinality events raised by the Bridge contract.
type BridgeSetConfirmationsToFinalityIterator struct {
	Event *BridgeSetConfirmationsToFinality // Event containing the contract specifics and raw log

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
func (it *BridgeSetConfirmationsToFinalityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetConfirmationsToFinality)
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
		it.Event = new(BridgeSetConfirmationsToFinality)
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
func (it *BridgeSetConfirmationsToFinalityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetConfirmationsToFinalityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetConfirmationsToFinality represents a SetConfirmationsToFinality event raised by the Bridge contract.
type BridgeSetConfirmationsToFinality struct {
	Arg0 uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetConfirmationsToFinality is a free log retrieval operation binding the contract event 0x7a2bd81684c5f899a41d138f583876506d31e8f3e88f7a8d124ed169cdd6bfe0.
//
// Solidity: event SetConfirmationsToFinality(uint64 arg0)
func (_Bridge *BridgeFilterer) FilterSetConfirmationsToFinality(opts *bind.FilterOpts) (*BridgeSetConfirmationsToFinalityIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetConfirmationsToFinality")
	if err != nil {
		return nil, err
	}
	return &BridgeSetConfirmationsToFinalityIterator{contract: _Bridge.contract, event: "SetConfirmationsToFinality", logs: logs, sub: sub}, nil
}

// WatchSetConfirmationsToFinality is a free log subscription operation binding the contract event 0x7a2bd81684c5f899a41d138f583876506d31e8f3e88f7a8d124ed169cdd6bfe0.
//
// Solidity: event SetConfirmationsToFinality(uint64 arg0)
func (_Bridge *BridgeFilterer) WatchSetConfirmationsToFinality(opts *bind.WatchOpts, sink chan<- *BridgeSetConfirmationsToFinality) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetConfirmationsToFinality")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetConfirmationsToFinality)
				if err := _Bridge.contract.UnpackLog(event, "SetConfirmationsToFinality", log); err != nil {
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

// ParseSetConfirmationsToFinality is a log parse operation binding the contract event 0x7a2bd81684c5f899a41d138f583876506d31e8f3e88f7a8d124ed169cdd6bfe0.
//
// Solidity: event SetConfirmationsToFinality(uint64 arg0)
func (_Bridge *BridgeFilterer) ParseSetConfirmationsToFinality(log types.Log) (*BridgeSetConfirmationsToFinality, error) {
	event := new(BridgeSetConfirmationsToFinality)
	if err := _Bridge.contract.UnpackLog(event, "SetConfirmationsToFinality", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetEstimatedBlockTimeIterator is returned from FilterSetEstimatedBlockTime and is used to iterate over the raw logs and unpacked data for SetEstimatedBlockTime events raised by the Bridge contract.
type BridgeSetEstimatedBlockTimeIterator struct {
	Event *BridgeSetEstimatedBlockTime // Event containing the contract specifics and raw log

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
func (it *BridgeSetEstimatedBlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetEstimatedBlockTime)
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
		it.Event = new(BridgeSetEstimatedBlockTime)
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
func (it *BridgeSetEstimatedBlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetEstimatedBlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetEstimatedBlockTime represents a SetEstimatedBlockTime event raised by the Bridge contract.
type BridgeSetEstimatedBlockTime struct {
	Arg0 uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetEstimatedBlockTime is a free log retrieval operation binding the contract event 0xf1ab3333803d2c9783b7929c6c99bd8b15ede77294566a9c8d97471dec21b90a.
//
// Solidity: event SetEstimatedBlockTime(uint64 arg0)
func (_Bridge *BridgeFilterer) FilterSetEstimatedBlockTime(opts *bind.FilterOpts) (*BridgeSetEstimatedBlockTimeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetEstimatedBlockTime")
	if err != nil {
		return nil, err
	}
	return &BridgeSetEstimatedBlockTimeIterator{contract: _Bridge.contract, event: "SetEstimatedBlockTime", logs: logs, sub: sub}, nil
}

// WatchSetEstimatedBlockTime is a free log subscription operation binding the contract event 0xf1ab3333803d2c9783b7929c6c99bd8b15ede77294566a9c8d97471dec21b90a.
//
// Solidity: event SetEstimatedBlockTime(uint64 arg0)
func (_Bridge *BridgeFilterer) WatchSetEstimatedBlockTime(opts *bind.WatchOpts, sink chan<- *BridgeSetEstimatedBlockTime) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetEstimatedBlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetEstimatedBlockTime)
				if err := _Bridge.contract.UnpackLog(event, "SetEstimatedBlockTime", log); err != nil {
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

// ParseSetEstimatedBlockTime is a log parse operation binding the contract event 0xf1ab3333803d2c9783b7929c6c99bd8b15ede77294566a9c8d97471dec21b90a.
//
// Solidity: event SetEstimatedBlockTime(uint64 arg0)
func (_Bridge *BridgeFilterer) ParseSetEstimatedBlockTime(log types.Log) (*BridgeSetEstimatedBlockTime, error) {
	event := new(BridgeSetEstimatedBlockTime)
	if err := _Bridge.contract.UnpackLog(event, "SetEstimatedBlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetGuardiansIterator is returned from FilterSetGuardians and is used to iterate over the raw logs and unpacked data for SetGuardians events raised by the Bridge contract.
type BridgeSetGuardiansIterator struct {
	Event *BridgeSetGuardians // Event containing the contract specifics and raw log

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
func (it *BridgeSetGuardiansIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetGuardians)
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
		it.Event = new(BridgeSetGuardians)
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
func (it *BridgeSetGuardiansIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetGuardiansIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetGuardians represents a SetGuardians event raised by the Bridge contract.
type BridgeSetGuardians struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetGuardians is a free log retrieval operation binding the contract event 0xade9cb9e0d9bd2fc419ba2eb751a8adb44beb8a13c106e78a18b2afba085aaa5.
//
// Solidity: event SetGuardians()
func (_Bridge *BridgeFilterer) FilterSetGuardians(opts *bind.FilterOpts) (*BridgeSetGuardiansIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetGuardians")
	if err != nil {
		return nil, err
	}
	return &BridgeSetGuardiansIterator{contract: _Bridge.contract, event: "SetGuardians", logs: logs, sub: sub}, nil
}

// WatchSetGuardians is a free log subscription operation binding the contract event 0xade9cb9e0d9bd2fc419ba2eb751a8adb44beb8a13c106e78a18b2afba085aaa5.
//
// Solidity: event SetGuardians()
func (_Bridge *BridgeFilterer) WatchSetGuardians(opts *bind.WatchOpts, sink chan<- *BridgeSetGuardians) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetGuardians")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetGuardians)
				if err := _Bridge.contract.UnpackLog(event, "SetGuardians", log); err != nil {
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

// ParseSetGuardians is a log parse operation binding the contract event 0xade9cb9e0d9bd2fc419ba2eb751a8adb44beb8a13c106e78a18b2afba085aaa5.
//
// Solidity: event SetGuardians()
func (_Bridge *BridgeFilterer) ParseSetGuardians(log types.Log) (*BridgeSetGuardians, error) {
	event := new(BridgeSetGuardians)
	if err := _Bridge.contract.UnpackLog(event, "SetGuardians", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetSoftDelayIterator is returned from FilterSetSoftDelay and is used to iterate over the raw logs and unpacked data for SetSoftDelay events raised by the Bridge contract.
type BridgeSetSoftDelayIterator struct {
	Event *BridgeSetSoftDelay // Event containing the contract specifics and raw log

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
func (it *BridgeSetSoftDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetSoftDelay)
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
		it.Event = new(BridgeSetSoftDelay)
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
func (it *BridgeSetSoftDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetSoftDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetSoftDelay represents a SetSoftDelay event raised by the Bridge contract.
type BridgeSetSoftDelay struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetSoftDelay is a free log retrieval operation binding the contract event 0xbccdc80e0b9b1548ca651855b773882ed962ad44b1af12d3021605254b7160f1.
//
// Solidity: event SetSoftDelay(uint256 arg0)
func (_Bridge *BridgeFilterer) FilterSetSoftDelay(opts *bind.FilterOpts) (*BridgeSetSoftDelayIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetSoftDelay")
	if err != nil {
		return nil, err
	}
	return &BridgeSetSoftDelayIterator{contract: _Bridge.contract, event: "SetSoftDelay", logs: logs, sub: sub}, nil
}

// WatchSetSoftDelay is a free log subscription operation binding the contract event 0xbccdc80e0b9b1548ca651855b773882ed962ad44b1af12d3021605254b7160f1.
//
// Solidity: event SetSoftDelay(uint256 arg0)
func (_Bridge *BridgeFilterer) WatchSetSoftDelay(opts *bind.WatchOpts, sink chan<- *BridgeSetSoftDelay) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetSoftDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetSoftDelay)
				if err := _Bridge.contract.UnpackLog(event, "SetSoftDelay", log); err != nil {
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

// ParseSetSoftDelay is a log parse operation binding the contract event 0xbccdc80e0b9b1548ca651855b773882ed962ad44b1af12d3021605254b7160f1.
//
// Solidity: event SetSoftDelay(uint256 arg0)
func (_Bridge *BridgeFilterer) ParseSetSoftDelay(log types.Log) (*BridgeSetSoftDelay, error) {
	event := new(BridgeSetSoftDelay)
	if err := _Bridge.contract.UnpackLog(event, "SetSoftDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetTokenInfoIterator is returned from FilterSetTokenInfo and is used to iterate over the raw logs and unpacked data for SetTokenInfo events raised by the Bridge contract.
type BridgeSetTokenInfoIterator struct {
	Event *BridgeSetTokenInfo // Event containing the contract specifics and raw log

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
func (it *BridgeSetTokenInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetTokenInfo)
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
		it.Event = new(BridgeSetTokenInfo)
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
func (it *BridgeSetTokenInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetTokenInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetTokenInfo represents a SetTokenInfo event raised by the Bridge contract.
type BridgeSetTokenInfo struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetTokenInfo is a free log retrieval operation binding the contract event 0x629fbdbc83c9e7e03a2983bce587b9da5a1c320cddc61f714f9dc81c10a66233.
//
// Solidity: event SetTokenInfo(address indexed token)
func (_Bridge *BridgeFilterer) FilterSetTokenInfo(opts *bind.FilterOpts, token []common.Address) (*BridgeSetTokenInfoIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetTokenInfo", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSetTokenInfoIterator{contract: _Bridge.contract, event: "SetTokenInfo", logs: logs, sub: sub}, nil
}

// WatchSetTokenInfo is a free log subscription operation binding the contract event 0x629fbdbc83c9e7e03a2983bce587b9da5a1c320cddc61f714f9dc81c10a66233.
//
// Solidity: event SetTokenInfo(address indexed token)
func (_Bridge *BridgeFilterer) WatchSetTokenInfo(opts *bind.WatchOpts, sink chan<- *BridgeSetTokenInfo, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetTokenInfo", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetTokenInfo)
				if err := _Bridge.contract.UnpackLog(event, "SetTokenInfo", log); err != nil {
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

// ParseSetTokenInfo is a log parse operation binding the contract event 0x629fbdbc83c9e7e03a2983bce587b9da5a1c320cddc61f714f9dc81c10a66233.
//
// Solidity: event SetTokenInfo(address indexed token)
func (_Bridge *BridgeFilterer) ParseSetTokenInfo(log types.Log) (*BridgeSetTokenInfo, error) {
	event := new(BridgeSetTokenInfo)
	if err := _Bridge.contract.UnpackLog(event, "SetTokenInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetTssIterator is returned from FilterSetTss and is used to iterate over the raw logs and unpacked data for SetTss events raised by the Bridge contract.
type BridgeSetTssIterator struct {
	Event *BridgeSetTss // Event containing the contract specifics and raw log

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
func (it *BridgeSetTssIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetTss)
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
		it.Event = new(BridgeSetTss)
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
func (it *BridgeSetTssIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetTssIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetTss represents a SetTss event raised by the Bridge contract.
type BridgeSetTss struct {
	NewTss common.Address
	OldTss common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTss is a free log retrieval operation binding the contract event 0xa37f87a74fe0e5d7a2c17e4d76a9b3467a6f29260495db13ebd6a486cf7d3cef.
//
// Solidity: event SetTss(address indexed newTss, address oldTss)
func (_Bridge *BridgeFilterer) FilterSetTss(opts *bind.FilterOpts, newTss []common.Address) (*BridgeSetTssIterator, error) {

	var newTssRule []interface{}
	for _, newTssItem := range newTss {
		newTssRule = append(newTssRule, newTssItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetTss", newTssRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSetTssIterator{contract: _Bridge.contract, event: "SetTss", logs: logs, sub: sub}, nil
}

// WatchSetTss is a free log subscription operation binding the contract event 0xa37f87a74fe0e5d7a2c17e4d76a9b3467a6f29260495db13ebd6a486cf7d3cef.
//
// Solidity: event SetTss(address indexed newTss, address oldTss)
func (_Bridge *BridgeFilterer) WatchSetTss(opts *bind.WatchOpts, sink chan<- *BridgeSetTss, newTss []common.Address) (event.Subscription, error) {

	var newTssRule []interface{}
	for _, newTssItem := range newTss {
		newTssRule = append(newTssRule, newTssItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetTss", newTssRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetTss)
				if err := _Bridge.contract.UnpackLog(event, "SetTss", log); err != nil {
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

// ParseSetTss is a log parse operation binding the contract event 0xa37f87a74fe0e5d7a2c17e4d76a9b3467a6f29260495db13ebd6a486cf7d3cef.
//
// Solidity: event SetTss(address indexed newTss, address oldTss)
func (_Bridge *BridgeFilterer) ParseSetTss(log types.Log) (*BridgeSetTss, error) {
	event := new(BridgeSetTss)
	if err := _Bridge.contract.UnpackLog(event, "SetTss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetUnhaltDurationIterator is returned from FilterSetUnhaltDuration and is used to iterate over the raw logs and unpacked data for SetUnhaltDuration events raised by the Bridge contract.
type BridgeSetUnhaltDurationIterator struct {
	Event *BridgeSetUnhaltDuration // Event containing the contract specifics and raw log

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
func (it *BridgeSetUnhaltDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetUnhaltDuration)
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
		it.Event = new(BridgeSetUnhaltDuration)
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
func (it *BridgeSetUnhaltDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetUnhaltDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetUnhaltDuration represents a SetUnhaltDuration event raised by the Bridge contract.
type BridgeSetUnhaltDuration struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetUnhaltDuration is a free log retrieval operation binding the contract event 0x2af629b5e397edd1e130c83e04770e3f5d9dd22424ecdbc204359406bbd3c7fa.
//
// Solidity: event SetUnhaltDuration(uint256 arg0)
func (_Bridge *BridgeFilterer) FilterSetUnhaltDuration(opts *bind.FilterOpts) (*BridgeSetUnhaltDurationIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetUnhaltDuration")
	if err != nil {
		return nil, err
	}
	return &BridgeSetUnhaltDurationIterator{contract: _Bridge.contract, event: "SetUnhaltDuration", logs: logs, sub: sub}, nil
}

// WatchSetUnhaltDuration is a free log subscription operation binding the contract event 0x2af629b5e397edd1e130c83e04770e3f5d9dd22424ecdbc204359406bbd3c7fa.
//
// Solidity: event SetUnhaltDuration(uint256 arg0)
func (_Bridge *BridgeFilterer) WatchSetUnhaltDuration(opts *bind.WatchOpts, sink chan<- *BridgeSetUnhaltDuration) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetUnhaltDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetUnhaltDuration)
				if err := _Bridge.contract.UnpackLog(event, "SetUnhaltDuration", log); err != nil {
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

// ParseSetUnhaltDuration is a log parse operation binding the contract event 0x2af629b5e397edd1e130c83e04770e3f5d9dd22424ecdbc204359406bbd3c7fa.
//
// Solidity: event SetUnhaltDuration(uint256 arg0)
func (_Bridge *BridgeFilterer) ParseSetUnhaltDuration(log types.Log) (*BridgeSetUnhaltDuration, error) {
	event := new(BridgeSetUnhaltDuration)
	if err := _Bridge.contract.UnpackLog(event, "SetUnhaltDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnhaltedIterator is returned from FilterUnhalted and is used to iterate over the raw logs and unpacked data for Unhalted events raised by the Bridge contract.
type BridgeUnhaltedIterator struct {
	Event *BridgeUnhalted // Event containing the contract specifics and raw log

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
func (it *BridgeUnhaltedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnhalted)
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
		it.Event = new(BridgeUnhalted)
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
func (it *BridgeUnhaltedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnhaltedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnhalted represents a Unhalted event raised by the Bridge contract.
type BridgeUnhalted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnhalted is a free log retrieval operation binding the contract event 0x7c46a5e7a10434913e987d799d659758880ce8e790692e13e66ddfae4cc9afca.
//
// Solidity: event Unhalted()
func (_Bridge *BridgeFilterer) FilterUnhalted(opts *bind.FilterOpts) (*BridgeUnhaltedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unhalted")
	if err != nil {
		return nil, err
	}
	return &BridgeUnhaltedIterator{contract: _Bridge.contract, event: "Unhalted", logs: logs, sub: sub}, nil
}

// WatchUnhalted is a free log subscription operation binding the contract event 0x7c46a5e7a10434913e987d799d659758880ce8e790692e13e66ddfae4cc9afca.
//
// Solidity: event Unhalted()
func (_Bridge *BridgeFilterer) WatchUnhalted(opts *bind.WatchOpts, sink chan<- *BridgeUnhalted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unhalted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnhalted)
				if err := _Bridge.contract.UnpackLog(event, "Unhalted", log); err != nil {
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

// ParseUnhalted is a log parse operation binding the contract event 0x7c46a5e7a10434913e987d799d659758880ce8e790692e13e66ddfae4cc9afca.
//
// Solidity: event Unhalted()
func (_Bridge *BridgeFilterer) ParseUnhalted(log types.Log) (*BridgeUnhalted, error) {
	event := new(BridgeUnhalted)
	if err := _Bridge.contract.UnpackLog(event, "Unhalted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnwrappedIterator is returned from FilterUnwrapped and is used to iterate over the raw logs and unpacked data for Unwrapped events raised by the Bridge contract.
type BridgeUnwrappedIterator struct {
	Event *BridgeUnwrapped // Event containing the contract specifics and raw log

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
func (it *BridgeUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnwrapped)
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
		it.Event = new(BridgeUnwrapped)
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
func (it *BridgeUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnwrapped represents a Unwrapped event raised by the Bridge contract.
type BridgeUnwrapped struct {
	From   common.Address
	Token  common.Address
	To     string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnwrapped is a free log retrieval operation binding the contract event 0x43a02321f21cdbe5a1c0c27634bfb1afed9b77259b534e83f247ad155c425a01.
//
// Solidity: event Unwrapped(address indexed from, address indexed token, string to, uint256 amount)
func (_Bridge *BridgeFilterer) FilterUnwrapped(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*BridgeUnwrappedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unwrapped", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnwrappedIterator{contract: _Bridge.contract, event: "Unwrapped", logs: logs, sub: sub}, nil
}

// WatchUnwrapped is a free log subscription operation binding the contract event 0x43a02321f21cdbe5a1c0c27634bfb1afed9b77259b534e83f247ad155c425a01.
//
// Solidity: event Unwrapped(address indexed from, address indexed token, string to, uint256 amount)
func (_Bridge *BridgeFilterer) WatchUnwrapped(opts *bind.WatchOpts, sink chan<- *BridgeUnwrapped, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unwrapped", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnwrapped)
				if err := _Bridge.contract.UnpackLog(event, "Unwrapped", log); err != nil {
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

// ParseUnwrapped is a log parse operation binding the contract event 0x43a02321f21cdbe5a1c0c27634bfb1afed9b77259b534e83f247ad155c425a01.
//
// Solidity: event Unwrapped(address indexed from, address indexed token, string to, uint256 amount)
func (_Bridge *BridgeFilterer) ParseUnwrapped(log types.Log) (*BridgeUnwrapped, error) {
	event := new(BridgeUnwrapped)
	if err := _Bridge.contract.UnpackLog(event, "Unwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
