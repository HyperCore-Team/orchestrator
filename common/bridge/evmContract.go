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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unhaltDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"administratorDelayParam\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tssDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"initialGuardians\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdministrator\",\"type\":\"address\"}],\"name\":\"ChangedAdministrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChangedGuardians\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTssAddress\",\"type\":\"address\"}],\"name\":\"ChangedTssAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Halted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"PendingAdministrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PendingGuardians\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"PendingTssAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RegisteredRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"RevokedRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unhalted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unwrapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionsNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administratorChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administratorDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowKeyGen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"changeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"oldSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newSignature\",\"type\":\"bytes\"}],\"name\":\"changeTssAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"changeTssAddressDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmationsToFinality\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractDeploymentHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedBlockTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardianChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"guardiansVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"halt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isHalted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAdministratorDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTssAddressChangeDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minUnhaltDurationInBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newGuardians\",\"type\":\"address[]\"}],\"name\":\"nominateGuardians\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nominatedGuardians\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"proposeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeemsInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paramsHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestedAdministrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestedTssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"revokeRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAllowKeyGen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setChangeTssAddressDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"}],\"name\":\"setConfirmationsToFinality\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockTime\",\"type\":\"uint64\"}],\"name\":\"setEstimatedBlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemDelay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOwned\",\"type\":\"bool\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setUnhaltDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemDelayInBlocks\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"owned\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddressChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhalt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhaltDurationInBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unhaltedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AdministratorChangeBlock is a free data retrieval call binding the contract method 0x8a1a2375.
//
// Solidity: function administratorChangeBlock() view returns(uint256)
func (_Bridge *BridgeCaller) AdministratorChangeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "administratorChangeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdministratorChangeBlock is a free data retrieval call binding the contract method 0x8a1a2375.
//
// Solidity: function administratorChangeBlock() view returns(uint256)
func (_Bridge *BridgeSession) AdministratorChangeBlock() (*big.Int, error) {
	return _Bridge.Contract.AdministratorChangeBlock(&_Bridge.CallOpts)
}

// AdministratorChangeBlock is a free data retrieval call binding the contract method 0x8a1a2375.
//
// Solidity: function administratorChangeBlock() view returns(uint256)
func (_Bridge *BridgeCallerSession) AdministratorChangeBlock() (*big.Int, error) {
	return _Bridge.Contract.AdministratorChangeBlock(&_Bridge.CallOpts)
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

// ChangeTssAddressDelay is a free data retrieval call binding the contract method 0x6c2e7be4.
//
// Solidity: function changeTssAddressDelay() view returns(uint256)
func (_Bridge *BridgeCaller) ChangeTssAddressDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "changeTssAddressDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChangeTssAddressDelay is a free data retrieval call binding the contract method 0x6c2e7be4.
//
// Solidity: function changeTssAddressDelay() view returns(uint256)
func (_Bridge *BridgeSession) ChangeTssAddressDelay() (*big.Int, error) {
	return _Bridge.Contract.ChangeTssAddressDelay(&_Bridge.CallOpts)
}

// ChangeTssAddressDelay is a free data retrieval call binding the contract method 0x6c2e7be4.
//
// Solidity: function changeTssAddressDelay() view returns(uint256)
func (_Bridge *BridgeCallerSession) ChangeTssAddressDelay() (*big.Int, error) {
	return _Bridge.Contract.ChangeTssAddressDelay(&_Bridge.CallOpts)
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

// GuardianChangeBlock is a free data retrieval call binding the contract method 0x92d669a9.
//
// Solidity: function guardianChangeBlock() view returns(uint256)
func (_Bridge *BridgeCaller) GuardianChangeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "guardianChangeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardianChangeBlock is a free data retrieval call binding the contract method 0x92d669a9.
//
// Solidity: function guardianChangeBlock() view returns(uint256)
func (_Bridge *BridgeSession) GuardianChangeBlock() (*big.Int, error) {
	return _Bridge.Contract.GuardianChangeBlock(&_Bridge.CallOpts)
}

// GuardianChangeBlock is a free data retrieval call binding the contract method 0x92d669a9.
//
// Solidity: function guardianChangeBlock() view returns(uint256)
func (_Bridge *BridgeCallerSession) GuardianChangeBlock() (*big.Int, error) {
	return _Bridge.Contract.GuardianChangeBlock(&_Bridge.CallOpts)
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

// MinTssAddressChangeDelay is a free data retrieval call binding the contract method 0xf547e0c8.
//
// Solidity: function minTssAddressChangeDelay() view returns(uint256)
func (_Bridge *BridgeCaller) MinTssAddressChangeDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minTssAddressChangeDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTssAddressChangeDelay is a free data retrieval call binding the contract method 0xf547e0c8.
//
// Solidity: function minTssAddressChangeDelay() view returns(uint256)
func (_Bridge *BridgeSession) MinTssAddressChangeDelay() (*big.Int, error) {
	return _Bridge.Contract.MinTssAddressChangeDelay(&_Bridge.CallOpts)
}

// MinTssAddressChangeDelay is a free data retrieval call binding the contract method 0xf547e0c8.
//
// Solidity: function minTssAddressChangeDelay() view returns(uint256)
func (_Bridge *BridgeCallerSession) MinTssAddressChangeDelay() (*big.Int, error) {
	return _Bridge.Contract.MinTssAddressChangeDelay(&_Bridge.CallOpts)
}

// MinUnhaltDurationInBlocks is a free data retrieval call binding the contract method 0x57161706.
//
// Solidity: function minUnhaltDurationInBlocks() view returns(uint256)
func (_Bridge *BridgeCaller) MinUnhaltDurationInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minUnhaltDurationInBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinUnhaltDurationInBlocks is a free data retrieval call binding the contract method 0x57161706.
//
// Solidity: function minUnhaltDurationInBlocks() view returns(uint256)
func (_Bridge *BridgeSession) MinUnhaltDurationInBlocks() (*big.Int, error) {
	return _Bridge.Contract.MinUnhaltDurationInBlocks(&_Bridge.CallOpts)
}

// MinUnhaltDurationInBlocks is a free data retrieval call binding the contract method 0x57161706.
//
// Solidity: function minUnhaltDurationInBlocks() view returns(uint256)
func (_Bridge *BridgeCallerSession) MinUnhaltDurationInBlocks() (*big.Int, error) {
	return _Bridge.Contract.MinUnhaltDurationInBlocks(&_Bridge.CallOpts)
}

// NominatedGuardians is a free data retrieval call binding the contract method 0x8029992f.
//
// Solidity: function nominatedGuardians(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) NominatedGuardians(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nominatedGuardians", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedGuardians is a free data retrieval call binding the contract method 0x8029992f.
//
// Solidity: function nominatedGuardians(uint256 ) view returns(address)
func (_Bridge *BridgeSession) NominatedGuardians(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.NominatedGuardians(&_Bridge.CallOpts, arg0)
}

// NominatedGuardians is a free data retrieval call binding the contract method 0x8029992f.
//
// Solidity: function nominatedGuardians(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) NominatedGuardians(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.NominatedGuardians(&_Bridge.CallOpts, arg0)
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

// RequestedAdministrator is a free data retrieval call binding the contract method 0xeb9f8edd.
//
// Solidity: function requestedAdministrator() view returns(address)
func (_Bridge *BridgeCaller) RequestedAdministrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "requestedAdministrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequestedAdministrator is a free data retrieval call binding the contract method 0xeb9f8edd.
//
// Solidity: function requestedAdministrator() view returns(address)
func (_Bridge *BridgeSession) RequestedAdministrator() (common.Address, error) {
	return _Bridge.Contract.RequestedAdministrator(&_Bridge.CallOpts)
}

// RequestedAdministrator is a free data retrieval call binding the contract method 0xeb9f8edd.
//
// Solidity: function requestedAdministrator() view returns(address)
func (_Bridge *BridgeCallerSession) RequestedAdministrator() (common.Address, error) {
	return _Bridge.Contract.RequestedAdministrator(&_Bridge.CallOpts)
}

// RequestedTssAddress is a free data retrieval call binding the contract method 0x9cecc98e.
//
// Solidity: function requestedTssAddress() view returns(address)
func (_Bridge *BridgeCaller) RequestedTssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "requestedTssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequestedTssAddress is a free data retrieval call binding the contract method 0x9cecc98e.
//
// Solidity: function requestedTssAddress() view returns(address)
func (_Bridge *BridgeSession) RequestedTssAddress() (common.Address, error) {
	return _Bridge.Contract.RequestedTssAddress(&_Bridge.CallOpts)
}

// RequestedTssAddress is a free data retrieval call binding the contract method 0x9cecc98e.
//
// Solidity: function requestedTssAddress() view returns(address)
func (_Bridge *BridgeCallerSession) RequestedTssAddress() (common.Address, error) {
	return _Bridge.Contract.RequestedTssAddress(&_Bridge.CallOpts)
}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 minAmount, uint256 redeemDelayInBlocks, bool allowed, bool owned)
func (_Bridge *BridgeCaller) TokensInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinAmount           *big.Int
	RedeemDelayInBlocks *big.Int
	Allowed             bool
	Owned               bool
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokensInfo", arg0)

	outstruct := new(struct {
		MinAmount           *big.Int
		RedeemDelayInBlocks *big.Int
		Allowed             bool
		Owned               bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RedeemDelayInBlocks = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Allowed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Owned = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 minAmount, uint256 redeemDelayInBlocks, bool allowed, bool owned)
func (_Bridge *BridgeSession) TokensInfo(arg0 common.Address) (struct {
	MinAmount           *big.Int
	RedeemDelayInBlocks *big.Int
	Allowed             bool
	Owned               bool
}, error) {
	return _Bridge.Contract.TokensInfo(&_Bridge.CallOpts, arg0)
}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 minAmount, uint256 redeemDelayInBlocks, bool allowed, bool owned)
func (_Bridge *BridgeCallerSession) TokensInfo(arg0 common.Address) (struct {
	MinAmount           *big.Int
	RedeemDelayInBlocks *big.Int
	Allowed             bool
	Owned               bool
}, error) {
	return _Bridge.Contract.TokensInfo(&_Bridge.CallOpts, arg0)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_Bridge *BridgeCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_Bridge *BridgeSession) TssAddress() (common.Address, error) {
	return _Bridge.Contract.TssAddress(&_Bridge.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_Bridge *BridgeCallerSession) TssAddress() (common.Address, error) {
	return _Bridge.Contract.TssAddress(&_Bridge.CallOpts)
}

// TssAddressChangeBlock is a free data retrieval call binding the contract method 0x4d78ec6d.
//
// Solidity: function tssAddressChangeBlock() view returns(uint256)
func (_Bridge *BridgeCaller) TssAddressChangeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tssAddressChangeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TssAddressChangeBlock is a free data retrieval call binding the contract method 0x4d78ec6d.
//
// Solidity: function tssAddressChangeBlock() view returns(uint256)
func (_Bridge *BridgeSession) TssAddressChangeBlock() (*big.Int, error) {
	return _Bridge.Contract.TssAddressChangeBlock(&_Bridge.CallOpts)
}

// TssAddressChangeBlock is a free data retrieval call binding the contract method 0x4d78ec6d.
//
// Solidity: function tssAddressChangeBlock() view returns(uint256)
func (_Bridge *BridgeCallerSession) TssAddressChangeBlock() (*big.Int, error) {
	return _Bridge.Contract.TssAddressChangeBlock(&_Bridge.CallOpts)
}

// UnhaltDurationInBlocks is a free data retrieval call binding the contract method 0x35e96ed8.
//
// Solidity: function unhaltDurationInBlocks() view returns(uint256)
func (_Bridge *BridgeCaller) UnhaltDurationInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "unhaltDurationInBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnhaltDurationInBlocks is a free data retrieval call binding the contract method 0x35e96ed8.
//
// Solidity: function unhaltDurationInBlocks() view returns(uint256)
func (_Bridge *BridgeSession) UnhaltDurationInBlocks() (*big.Int, error) {
	return _Bridge.Contract.UnhaltDurationInBlocks(&_Bridge.CallOpts)
}

// UnhaltDurationInBlocks is a free data retrieval call binding the contract method 0x35e96ed8.
//
// Solidity: function unhaltDurationInBlocks() view returns(uint256)
func (_Bridge *BridgeCallerSession) UnhaltDurationInBlocks() (*big.Int, error) {
	return _Bridge.Contract.UnhaltDurationInBlocks(&_Bridge.CallOpts)
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

// ChangeAdministrator is a paid mutator transaction binding the contract method 0x988a9727.
//
// Solidity: function changeAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeTransactor) ChangeAdministrator(opts *bind.TransactOpts, newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeAdministrator", newAdministrator)
}

// ChangeAdministrator is a paid mutator transaction binding the contract method 0x988a9727.
//
// Solidity: function changeAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeSession) ChangeAdministrator(newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeAdministrator(&_Bridge.TransactOpts, newAdministrator)
}

// ChangeAdministrator is a paid mutator transaction binding the contract method 0x988a9727.
//
// Solidity: function changeAdministrator(address newAdministrator) returns()
func (_Bridge *BridgeTransactorSession) ChangeAdministrator(newAdministrator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeAdministrator(&_Bridge.TransactOpts, newAdministrator)
}

// ChangeTssAddress is a paid mutator transaction binding the contract method 0x9ee657c6.
//
// Solidity: function changeTssAddress(address newTssAddress, bytes oldSignature, bytes newSignature) returns()
func (_Bridge *BridgeTransactor) ChangeTssAddress(opts *bind.TransactOpts, newTssAddress common.Address, oldSignature []byte, newSignature []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeTssAddress", newTssAddress, oldSignature, newSignature)
}

// ChangeTssAddress is a paid mutator transaction binding the contract method 0x9ee657c6.
//
// Solidity: function changeTssAddress(address newTssAddress, bytes oldSignature, bytes newSignature) returns()
func (_Bridge *BridgeSession) ChangeTssAddress(newTssAddress common.Address, oldSignature []byte, newSignature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeTssAddress(&_Bridge.TransactOpts, newTssAddress, oldSignature, newSignature)
}

// ChangeTssAddress is a paid mutator transaction binding the contract method 0x9ee657c6.
//
// Solidity: function changeTssAddress(address newTssAddress, bytes oldSignature, bytes newSignature) returns()
func (_Bridge *BridgeTransactorSession) ChangeTssAddress(newTssAddress common.Address, oldSignature []byte, newSignature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeTssAddress(&_Bridge.TransactOpts, newTssAddress, oldSignature, newSignature)
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

// SetChangeTssAddressDelay is a paid mutator transaction binding the contract method 0xa876530a.
//
// Solidity: function setChangeTssAddressDelay(uint256 delay) returns()
func (_Bridge *BridgeTransactor) SetChangeTssAddressDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setChangeTssAddressDelay", delay)
}

// SetChangeTssAddressDelay is a paid mutator transaction binding the contract method 0xa876530a.
//
// Solidity: function setChangeTssAddressDelay(uint256 delay) returns()
func (_Bridge *BridgeSession) SetChangeTssAddressDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetChangeTssAddressDelay(&_Bridge.TransactOpts, delay)
}

// SetChangeTssAddressDelay is a paid mutator transaction binding the contract method 0xa876530a.
//
// Solidity: function setChangeTssAddressDelay(uint256 delay) returns()
func (_Bridge *BridgeTransactorSession) SetChangeTssAddressDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetChangeTssAddressDelay(&_Bridge.TransactOpts, delay)
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

// SetTokenInfo is a paid mutator transaction binding the contract method 0x51557449.
//
// Solidity: function setTokenInfo(address token, uint256 minAmount, uint256 redeemDelay, bool allowed, bool isOwned) returns()
func (_Bridge *BridgeTransactor) SetTokenInfo(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, redeemDelay *big.Int, allowed bool, isOwned bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTokenInfo", token, minAmount, redeemDelay, allowed, isOwned)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x51557449.
//
// Solidity: function setTokenInfo(address token, uint256 minAmount, uint256 redeemDelay, bool allowed, bool isOwned) returns()
func (_Bridge *BridgeSession) SetTokenInfo(token common.Address, minAmount *big.Int, redeemDelay *big.Int, allowed bool, isOwned bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenInfo(&_Bridge.TransactOpts, token, minAmount, redeemDelay, allowed, isOwned)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x51557449.
//
// Solidity: function setTokenInfo(address token, uint256 minAmount, uint256 redeemDelay, bool allowed, bool isOwned) returns()
func (_Bridge *BridgeTransactorSession) SetTokenInfo(token common.Address, minAmount *big.Int, redeemDelay *big.Int, allowed bool, isOwned bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenInfo(&_Bridge.TransactOpts, token, minAmount, redeemDelay, allowed, isOwned)
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

// BridgeChangedAdministratorIterator is returned from FilterChangedAdministrator and is used to iterate over the raw logs and unpacked data for ChangedAdministrator events raised by the Bridge contract.
type BridgeChangedAdministratorIterator struct {
	Event *BridgeChangedAdministrator // Event containing the contract specifics and raw log

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
func (it *BridgeChangedAdministratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChangedAdministrator)
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
		it.Event = new(BridgeChangedAdministrator)
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
func (it *BridgeChangedAdministratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChangedAdministratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChangedAdministrator represents a ChangedAdministrator event raised by the Bridge contract.
type BridgeChangedAdministrator struct {
	NewAdministrator common.Address
	OldAdministrator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChangedAdministrator is a free log retrieval operation binding the contract event 0xd26c57e646f4fdb3f746240d8408ad897af56bb48d9ec17aa31859be0bccd795.
//
// Solidity: event ChangedAdministrator(address indexed newAdministrator, address oldAdministrator)
func (_Bridge *BridgeFilterer) FilterChangedAdministrator(opts *bind.FilterOpts, newAdministrator []common.Address) (*BridgeChangedAdministratorIterator, error) {

	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChangedAdministrator", newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeChangedAdministratorIterator{contract: _Bridge.contract, event: "ChangedAdministrator", logs: logs, sub: sub}, nil
}

// WatchChangedAdministrator is a free log subscription operation binding the contract event 0xd26c57e646f4fdb3f746240d8408ad897af56bb48d9ec17aa31859be0bccd795.
//
// Solidity: event ChangedAdministrator(address indexed newAdministrator, address oldAdministrator)
func (_Bridge *BridgeFilterer) WatchChangedAdministrator(opts *bind.WatchOpts, sink chan<- *BridgeChangedAdministrator, newAdministrator []common.Address) (event.Subscription, error) {

	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChangedAdministrator", newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChangedAdministrator)
				if err := _Bridge.contract.UnpackLog(event, "ChangedAdministrator", log); err != nil {
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

// ParseChangedAdministrator is a log parse operation binding the contract event 0xd26c57e646f4fdb3f746240d8408ad897af56bb48d9ec17aa31859be0bccd795.
//
// Solidity: event ChangedAdministrator(address indexed newAdministrator, address oldAdministrator)
func (_Bridge *BridgeFilterer) ParseChangedAdministrator(log types.Log) (*BridgeChangedAdministrator, error) {
	event := new(BridgeChangedAdministrator)
	if err := _Bridge.contract.UnpackLog(event, "ChangedAdministrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeChangedGuardiansIterator is returned from FilterChangedGuardians and is used to iterate over the raw logs and unpacked data for ChangedGuardians events raised by the Bridge contract.
type BridgeChangedGuardiansIterator struct {
	Event *BridgeChangedGuardians // Event containing the contract specifics and raw log

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
func (it *BridgeChangedGuardiansIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChangedGuardians)
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
		it.Event = new(BridgeChangedGuardians)
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
func (it *BridgeChangedGuardiansIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChangedGuardiansIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChangedGuardians represents a ChangedGuardians event raised by the Bridge contract.
type BridgeChangedGuardians struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangedGuardians is a free log retrieval operation binding the contract event 0x5f25b7a8dc5f9542e3cf70285ca7075f66d4688be0e75e950d816c6cfbd7b336.
//
// Solidity: event ChangedGuardians()
func (_Bridge *BridgeFilterer) FilterChangedGuardians(opts *bind.FilterOpts) (*BridgeChangedGuardiansIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChangedGuardians")
	if err != nil {
		return nil, err
	}
	return &BridgeChangedGuardiansIterator{contract: _Bridge.contract, event: "ChangedGuardians", logs: logs, sub: sub}, nil
}

// WatchChangedGuardians is a free log subscription operation binding the contract event 0x5f25b7a8dc5f9542e3cf70285ca7075f66d4688be0e75e950d816c6cfbd7b336.
//
// Solidity: event ChangedGuardians()
func (_Bridge *BridgeFilterer) WatchChangedGuardians(opts *bind.WatchOpts, sink chan<- *BridgeChangedGuardians) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChangedGuardians")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChangedGuardians)
				if err := _Bridge.contract.UnpackLog(event, "ChangedGuardians", log); err != nil {
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

// ParseChangedGuardians is a log parse operation binding the contract event 0x5f25b7a8dc5f9542e3cf70285ca7075f66d4688be0e75e950d816c6cfbd7b336.
//
// Solidity: event ChangedGuardians()
func (_Bridge *BridgeFilterer) ParseChangedGuardians(log types.Log) (*BridgeChangedGuardians, error) {
	event := new(BridgeChangedGuardians)
	if err := _Bridge.contract.UnpackLog(event, "ChangedGuardians", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeChangedTssAddressIterator is returned from FilterChangedTssAddress and is used to iterate over the raw logs and unpacked data for ChangedTssAddress events raised by the Bridge contract.
type BridgeChangedTssAddressIterator struct {
	Event *BridgeChangedTssAddress // Event containing the contract specifics and raw log

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
func (it *BridgeChangedTssAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChangedTssAddress)
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
		it.Event = new(BridgeChangedTssAddress)
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
func (it *BridgeChangedTssAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChangedTssAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChangedTssAddress represents a ChangedTssAddress event raised by the Bridge contract.
type BridgeChangedTssAddress struct {
	NewTssAddress common.Address
	OldTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangedTssAddress is a free log retrieval operation binding the contract event 0xda91bc0c8df0b1b4750875013dc731906e54b411afea3459c0e923f2c424a432.
//
// Solidity: event ChangedTssAddress(address indexed newTssAddress, address oldTssAddress)
func (_Bridge *BridgeFilterer) FilterChangedTssAddress(opts *bind.FilterOpts, newTssAddress []common.Address) (*BridgeChangedTssAddressIterator, error) {

	var newTssAddressRule []interface{}
	for _, newTssAddressItem := range newTssAddress {
		newTssAddressRule = append(newTssAddressRule, newTssAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChangedTssAddress", newTssAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeChangedTssAddressIterator{contract: _Bridge.contract, event: "ChangedTssAddress", logs: logs, sub: sub}, nil
}

// WatchChangedTssAddress is a free log subscription operation binding the contract event 0xda91bc0c8df0b1b4750875013dc731906e54b411afea3459c0e923f2c424a432.
//
// Solidity: event ChangedTssAddress(address indexed newTssAddress, address oldTssAddress)
func (_Bridge *BridgeFilterer) WatchChangedTssAddress(opts *bind.WatchOpts, sink chan<- *BridgeChangedTssAddress, newTssAddress []common.Address) (event.Subscription, error) {

	var newTssAddressRule []interface{}
	for _, newTssAddressItem := range newTssAddress {
		newTssAddressRule = append(newTssAddressRule, newTssAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChangedTssAddress", newTssAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChangedTssAddress)
				if err := _Bridge.contract.UnpackLog(event, "ChangedTssAddress", log); err != nil {
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

// ParseChangedTssAddress is a log parse operation binding the contract event 0xda91bc0c8df0b1b4750875013dc731906e54b411afea3459c0e923f2c424a432.
//
// Solidity: event ChangedTssAddress(address indexed newTssAddress, address oldTssAddress)
func (_Bridge *BridgeFilterer) ParseChangedTssAddress(log types.Log) (*BridgeChangedTssAddress, error) {
	event := new(BridgeChangedTssAddress)
	if err := _Bridge.contract.UnpackLog(event, "ChangedTssAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// BridgePendingTssAddressIterator is returned from FilterPendingTssAddress and is used to iterate over the raw logs and unpacked data for PendingTssAddress events raised by the Bridge contract.
type BridgePendingTssAddressIterator struct {
	Event *BridgePendingTssAddress // Event containing the contract specifics and raw log

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
func (it *BridgePendingTssAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePendingTssAddress)
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
		it.Event = new(BridgePendingTssAddress)
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
func (it *BridgePendingTssAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePendingTssAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePendingTssAddress represents a PendingTssAddress event raised by the Bridge contract.
type BridgePendingTssAddress struct {
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingTssAddress is a free log retrieval operation binding the contract event 0xdb0a7ee6019c8bdae9668159d0a941d7463cde5e39e6658226e34fa02a706680.
//
// Solidity: event PendingTssAddress(address indexed newTssAddress)
func (_Bridge *BridgeFilterer) FilterPendingTssAddress(opts *bind.FilterOpts, newTssAddress []common.Address) (*BridgePendingTssAddressIterator, error) {

	var newTssAddressRule []interface{}
	for _, newTssAddressItem := range newTssAddress {
		newTssAddressRule = append(newTssAddressRule, newTssAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PendingTssAddress", newTssAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgePendingTssAddressIterator{contract: _Bridge.contract, event: "PendingTssAddress", logs: logs, sub: sub}, nil
}

// WatchPendingTssAddress is a free log subscription operation binding the contract event 0xdb0a7ee6019c8bdae9668159d0a941d7463cde5e39e6658226e34fa02a706680.
//
// Solidity: event PendingTssAddress(address indexed newTssAddress)
func (_Bridge *BridgeFilterer) WatchPendingTssAddress(opts *bind.WatchOpts, sink chan<- *BridgePendingTssAddress, newTssAddress []common.Address) (event.Subscription, error) {

	var newTssAddressRule []interface{}
	for _, newTssAddressItem := range newTssAddress {
		newTssAddressRule = append(newTssAddressRule, newTssAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PendingTssAddress", newTssAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePendingTssAddress)
				if err := _Bridge.contract.UnpackLog(event, "PendingTssAddress", log); err != nil {
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

// ParsePendingTssAddress is a log parse operation binding the contract event 0xdb0a7ee6019c8bdae9668159d0a941d7463cde5e39e6658226e34fa02a706680.
//
// Solidity: event PendingTssAddress(address indexed newTssAddress)
func (_Bridge *BridgeFilterer) ParsePendingTssAddress(log types.Log) (*BridgePendingTssAddress, error) {
	event := new(BridgePendingTssAddress)
	if err := _Bridge.contract.UnpackLog(event, "PendingTssAddress", log); err != nil {
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
