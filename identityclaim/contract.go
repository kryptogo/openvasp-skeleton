// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identityclaim

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IdentityclaimABI is the input ABI used to generate the binding from.
const IdentityclaimABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vasp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Identityclaim is an auto generated Go binding around an Ethereum contract.
type Identityclaim struct {
	IdentityclaimCaller     // Read-only binding to the contract
	IdentityclaimTransactor // Write-only binding to the contract
	IdentityclaimFilterer   // Log filterer for contract events
}

// IdentityclaimCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityclaimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityclaimTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityclaimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityclaimFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityclaimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityclaimSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityclaimSession struct {
	Contract     *Identityclaim    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityclaimCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityclaimCallerSession struct {
	Contract *IdentityclaimCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IdentityclaimTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityclaimTransactorSession struct {
	Contract     *IdentityclaimTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IdentityclaimRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityclaimRaw struct {
	Contract *Identityclaim // Generic contract binding to access the raw methods on
}

// IdentityclaimCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityclaimCallerRaw struct {
	Contract *IdentityclaimCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityclaimTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityclaimTransactorRaw struct {
	Contract *IdentityclaimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityclaim creates a new instance of Identityclaim, bound to a specific deployed contract.
func NewIdentityclaim(address common.Address, backend bind.ContractBackend) (*Identityclaim, error) {
	contract, err := bindIdentityclaim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identityclaim{IdentityclaimCaller: IdentityclaimCaller{contract: contract}, IdentityclaimTransactor: IdentityclaimTransactor{contract: contract}, IdentityclaimFilterer: IdentityclaimFilterer{contract: contract}}, nil
}

// NewIdentityclaimCaller creates a new read-only instance of Identityclaim, bound to a specific deployed contract.
func NewIdentityclaimCaller(address common.Address, caller bind.ContractCaller) (*IdentityclaimCaller, error) {
	contract, err := bindIdentityclaim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityclaimCaller{contract: contract}, nil
}

// NewIdentityclaimTransactor creates a new write-only instance of Identityclaim, bound to a specific deployed contract.
func NewIdentityclaimTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityclaimTransactor, error) {
	contract, err := bindIdentityclaim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityclaimTransactor{contract: contract}, nil
}

// NewIdentityclaimFilterer creates a new log filterer instance of Identityclaim, bound to a specific deployed contract.
func NewIdentityclaimFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityclaimFilterer, error) {
	contract, err := bindIdentityclaim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityclaimFilterer{contract: contract}, nil
}

// bindIdentityclaim binds a generic wrapper to an already deployed contract.
func bindIdentityclaim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityclaimABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identityclaim *IdentityclaimRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identityclaim.Contract.IdentityclaimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identityclaim *IdentityclaimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityclaim.Contract.IdentityclaimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identityclaim *IdentityclaimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identityclaim.Contract.IdentityclaimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identityclaim *IdentityclaimCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identityclaim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identityclaim *IdentityclaimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityclaim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identityclaim *IdentityclaimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identityclaim.Contract.contract.Transact(opts, method, params...)
}

// Claim is a free data retrieval call binding the contract method 0x4e71d92d.
//
// Solidity: function claim() constant returns(bytes32)
func (_Identityclaim *IdentityclaimCaller) Claim(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identityclaim.contract.Call(opts, out, "claim")
	return *ret0, err
}

// Claim is a free data retrieval call binding the contract method 0x4e71d92d.
//
// Solidity: function claim() constant returns(bytes32)
func (_Identityclaim *IdentityclaimSession) Claim() ([32]byte, error) {
	return _Identityclaim.Contract.Claim(&_Identityclaim.CallOpts)
}

// Claim is a free data retrieval call binding the contract method 0x4e71d92d.
//
// Solidity: function claim() constant returns(bytes32)
func (_Identityclaim *IdentityclaimCallerSession) Claim() ([32]byte, error) {
	return _Identityclaim.Contract.Claim(&_Identityclaim.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_Identityclaim *IdentityclaimCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identityclaim.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_Identityclaim *IdentityclaimSession) Issuer() (common.Address, error) {
	return _Identityclaim.Contract.Issuer(&_Identityclaim.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_Identityclaim *IdentityclaimCallerSession) Issuer() (common.Address, error) {
	return _Identityclaim.Contract.Issuer(&_Identityclaim.CallOpts)
}

// Vasp is a free data retrieval call binding the contract method 0x71cc58f9.
//
// Solidity: function vasp() constant returns(address)
func (_Identityclaim *IdentityclaimCaller) Vasp(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identityclaim.contract.Call(opts, out, "vasp")
	return *ret0, err
}

// Vasp is a free data retrieval call binding the contract method 0x71cc58f9.
//
// Solidity: function vasp() constant returns(address)
func (_Identityclaim *IdentityclaimSession) Vasp() (common.Address, error) {
	return _Identityclaim.Contract.Vasp(&_Identityclaim.CallOpts)
}

// Vasp is a free data retrieval call binding the contract method 0x71cc58f9.
//
// Solidity: function vasp() constant returns(address)
func (_Identityclaim *IdentityclaimCallerSession) Vasp() (common.Address, error) {
	return _Identityclaim.Contract.Vasp(&_Identityclaim.CallOpts)
}
