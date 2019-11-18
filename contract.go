// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openvasp

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

// OpenvaspABI is the input ABI used to generate the binding from.
const OpenvaspABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"addressLine\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buildingNo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"code\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"country\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"email\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"handshakeKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"identityClaims\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"postCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signingKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"street\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"town\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trustedPeers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"website\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addTrustedPeer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeTrustedPeer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Openvasp is an auto generated Go binding around an Ethereum contract.
type Openvasp struct {
	OpenvaspCaller     // Read-only binding to the contract
	OpenvaspTransactor // Write-only binding to the contract
	OpenvaspFilterer   // Log filterer for contract events
}

// OpenvaspCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenvaspCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenvaspTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenvaspTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenvaspFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenvaspFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenvaspSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenvaspSession struct {
	Contract     *Openvasp         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenvaspCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenvaspCallerSession struct {
	Contract *OpenvaspCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OpenvaspTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenvaspTransactorSession struct {
	Contract     *OpenvaspTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OpenvaspRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenvaspRaw struct {
	Contract *Openvasp // Generic contract binding to access the raw methods on
}

// OpenvaspCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenvaspCallerRaw struct {
	Contract *OpenvaspCaller // Generic read-only contract binding to access the raw methods on
}

// OpenvaspTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenvaspTransactorRaw struct {
	Contract *OpenvaspTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenvasp creates a new instance of Openvasp, bound to a specific deployed contract.
func NewOpenvasp(address common.Address, backend bind.ContractBackend) (*Openvasp, error) {
	contract, err := bindOpenvasp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Openvasp{OpenvaspCaller: OpenvaspCaller{contract: contract}, OpenvaspTransactor: OpenvaspTransactor{contract: contract}, OpenvaspFilterer: OpenvaspFilterer{contract: contract}}, nil
}

// NewOpenvaspCaller creates a new read-only instance of Openvasp, bound to a specific deployed contract.
func NewOpenvaspCaller(address common.Address, caller bind.ContractCaller) (*OpenvaspCaller, error) {
	contract, err := bindOpenvasp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenvaspCaller{contract: contract}, nil
}

// NewOpenvaspTransactor creates a new write-only instance of Openvasp, bound to a specific deployed contract.
func NewOpenvaspTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenvaspTransactor, error) {
	contract, err := bindOpenvasp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenvaspTransactor{contract: contract}, nil
}

// NewOpenvaspFilterer creates a new log filterer instance of Openvasp, bound to a specific deployed contract.
func NewOpenvaspFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenvaspFilterer, error) {
	contract, err := bindOpenvasp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenvaspFilterer{contract: contract}, nil
}

// bindOpenvasp binds a generic wrapper to an already deployed contract.
func bindOpenvasp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenvaspABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Openvasp *OpenvaspRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Openvasp.Contract.OpenvaspCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Openvasp *OpenvaspRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Openvasp.Contract.OpenvaspTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Openvasp *OpenvaspRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Openvasp.Contract.OpenvaspTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Openvasp *OpenvaspCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Openvasp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Openvasp *OpenvaspTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Openvasp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Openvasp *OpenvaspTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Openvasp.Contract.contract.Transact(opts, method, params...)
}

// AddressLine is a free data retrieval call binding the contract method 0x9cabdf66.
//
// Solidity: function addressLine() constant returns(string)
func (_Openvasp *OpenvaspCaller) AddressLine(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "addressLine")
	return *ret0, err
}

// AddressLine is a free data retrieval call binding the contract method 0x9cabdf66.
//
// Solidity: function addressLine() constant returns(string)
func (_Openvasp *OpenvaspSession) AddressLine() (string, error) {
	return _Openvasp.Contract.AddressLine(&_Openvasp.CallOpts)
}

// AddressLine is a free data retrieval call binding the contract method 0x9cabdf66.
//
// Solidity: function addressLine() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) AddressLine() (string, error) {
	return _Openvasp.Contract.AddressLine(&_Openvasp.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() constant returns(address)
func (_Openvasp *OpenvaspCaller) Administrator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "administrator")
	return *ret0, err
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() constant returns(address)
func (_Openvasp *OpenvaspSession) Administrator() (common.Address, error) {
	return _Openvasp.Contract.Administrator(&_Openvasp.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() constant returns(address)
func (_Openvasp *OpenvaspCallerSession) Administrator() (common.Address, error) {
	return _Openvasp.Contract.Administrator(&_Openvasp.CallOpts)
}

// BuildingNo is a free data retrieval call binding the contract method 0x0222fd40.
//
// Solidity: function buildingNo() constant returns(bytes32)
func (_Openvasp *OpenvaspCaller) BuildingNo(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "buildingNo")
	return *ret0, err
}

// BuildingNo is a free data retrieval call binding the contract method 0x0222fd40.
//
// Solidity: function buildingNo() constant returns(bytes32)
func (_Openvasp *OpenvaspSession) BuildingNo() ([32]byte, error) {
	return _Openvasp.Contract.BuildingNo(&_Openvasp.CallOpts)
}

// BuildingNo is a free data retrieval call binding the contract method 0x0222fd40.
//
// Solidity: function buildingNo() constant returns(bytes32)
func (_Openvasp *OpenvaspCallerSession) BuildingNo() ([32]byte, error) {
	return _Openvasp.Contract.BuildingNo(&_Openvasp.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(uint8)
func (_Openvasp *OpenvaspCaller) Channels(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "channels", arg0)
	return *ret0, err
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(uint8)
func (_Openvasp *OpenvaspSession) Channels(arg0 *big.Int) (uint8, error) {
	return _Openvasp.Contract.Channels(&_Openvasp.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(uint8)
func (_Openvasp *OpenvaspCallerSession) Channels(arg0 *big.Int) (uint8, error) {
	return _Openvasp.Contract.Channels(&_Openvasp.CallOpts, arg0)
}

// Code is a free data retrieval call binding the contract method 0x24c12bf6.
//
// Solidity: function code() constant returns(bytes8)
func (_Openvasp *OpenvaspCaller) Code(opts *bind.CallOpts) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "code")
	return *ret0, err
}

// Code is a free data retrieval call binding the contract method 0x24c12bf6.
//
// Solidity: function code() constant returns(bytes8)
func (_Openvasp *OpenvaspSession) Code() ([8]byte, error) {
	return _Openvasp.Contract.Code(&_Openvasp.CallOpts)
}

// Code is a free data retrieval call binding the contract method 0x24c12bf6.
//
// Solidity: function code() constant returns(bytes8)
func (_Openvasp *OpenvaspCallerSession) Code() ([8]byte, error) {
	return _Openvasp.Contract.Code(&_Openvasp.CallOpts)
}

// Country is a free data retrieval call binding the contract method 0xd8b0b499.
//
// Solidity: function country() constant returns(bytes2)
func (_Openvasp *OpenvaspCaller) Country(opts *bind.CallOpts) ([2]byte, error) {
	var (
		ret0 = new([2]byte)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "country")
	return *ret0, err
}

// Country is a free data retrieval call binding the contract method 0xd8b0b499.
//
// Solidity: function country() constant returns(bytes2)
func (_Openvasp *OpenvaspSession) Country() ([2]byte, error) {
	return _Openvasp.Contract.Country(&_Openvasp.CallOpts)
}

// Country is a free data retrieval call binding the contract method 0xd8b0b499.
//
// Solidity: function country() constant returns(bytes2)
func (_Openvasp *OpenvaspCallerSession) Country() ([2]byte, error) {
	return _Openvasp.Contract.Country(&_Openvasp.CallOpts)
}

// Email is a free data retrieval call binding the contract method 0x820e93f5.
//
// Solidity: function email() constant returns(string)
func (_Openvasp *OpenvaspCaller) Email(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "email")
	return *ret0, err
}

// Email is a free data retrieval call binding the contract method 0x820e93f5.
//
// Solidity: function email() constant returns(string)
func (_Openvasp *OpenvaspSession) Email() (string, error) {
	return _Openvasp.Contract.Email(&_Openvasp.CallOpts)
}

// Email is a free data retrieval call binding the contract method 0x820e93f5.
//
// Solidity: function email() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) Email() (string, error) {
	return _Openvasp.Contract.Email(&_Openvasp.CallOpts)
}

// HandshakeKey is a free data retrieval call binding the contract method 0x7cc1c587.
//
// Solidity: function handshakeKey() constant returns(string)
func (_Openvasp *OpenvaspCaller) HandshakeKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "handshakeKey")
	return *ret0, err
}

// HandshakeKey is a free data retrieval call binding the contract method 0x7cc1c587.
//
// Solidity: function handshakeKey() constant returns(string)
func (_Openvasp *OpenvaspSession) HandshakeKey() (string, error) {
	return _Openvasp.Contract.HandshakeKey(&_Openvasp.CallOpts)
}

// HandshakeKey is a free data retrieval call binding the contract method 0x7cc1c587.
//
// Solidity: function handshakeKey() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) HandshakeKey() (string, error) {
	return _Openvasp.Contract.HandshakeKey(&_Openvasp.CallOpts)
}

// IdentityClaims is a free data retrieval call binding the contract method 0x8edea32c.
//
// Solidity: function identityClaims(uint256 ) constant returns(address)
func (_Openvasp *OpenvaspCaller) IdentityClaims(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "identityClaims", arg0)
	return *ret0, err
}

// IdentityClaims is a free data retrieval call binding the contract method 0x8edea32c.
//
// Solidity: function identityClaims(uint256 ) constant returns(address)
func (_Openvasp *OpenvaspSession) IdentityClaims(arg0 *big.Int) (common.Address, error) {
	return _Openvasp.Contract.IdentityClaims(&_Openvasp.CallOpts, arg0)
}

// IdentityClaims is a free data retrieval call binding the contract method 0x8edea32c.
//
// Solidity: function identityClaims(uint256 ) constant returns(address)
func (_Openvasp *OpenvaspCallerSession) IdentityClaims(arg0 *big.Int) (common.Address, error) {
	return _Openvasp.Contract.IdentityClaims(&_Openvasp.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Openvasp *OpenvaspCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Openvasp *OpenvaspSession) Name() (string, error) {
	return _Openvasp.Contract.Name(&_Openvasp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) Name() (string, error) {
	return _Openvasp.Contract.Name(&_Openvasp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Openvasp *OpenvaspCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Openvasp *OpenvaspSession) Owner() (common.Address, error) {
	return _Openvasp.Contract.Owner(&_Openvasp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Openvasp *OpenvaspCallerSession) Owner() (common.Address, error) {
	return _Openvasp.Contract.Owner(&_Openvasp.CallOpts)
}

// PostCode is a free data retrieval call binding the contract method 0xa69ae55b.
//
// Solidity: function postCode() constant returns(bytes32)
func (_Openvasp *OpenvaspCaller) PostCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "postCode")
	return *ret0, err
}

// PostCode is a free data retrieval call binding the contract method 0xa69ae55b.
//
// Solidity: function postCode() constant returns(bytes32)
func (_Openvasp *OpenvaspSession) PostCode() ([32]byte, error) {
	return _Openvasp.Contract.PostCode(&_Openvasp.CallOpts)
}

// PostCode is a free data retrieval call binding the contract method 0xa69ae55b.
//
// Solidity: function postCode() constant returns(bytes32)
func (_Openvasp *OpenvaspCallerSession) PostCode() ([32]byte, error) {
	return _Openvasp.Contract.PostCode(&_Openvasp.CallOpts)
}

// SigningKey is a free data retrieval call binding the contract method 0x66f27fd3.
//
// Solidity: function signingKey() constant returns(string)
func (_Openvasp *OpenvaspCaller) SigningKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "signingKey")
	return *ret0, err
}

// SigningKey is a free data retrieval call binding the contract method 0x66f27fd3.
//
// Solidity: function signingKey() constant returns(string)
func (_Openvasp *OpenvaspSession) SigningKey() (string, error) {
	return _Openvasp.Contract.SigningKey(&_Openvasp.CallOpts)
}

// SigningKey is a free data retrieval call binding the contract method 0x66f27fd3.
//
// Solidity: function signingKey() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) SigningKey() (string, error) {
	return _Openvasp.Contract.SigningKey(&_Openvasp.CallOpts)
}

// Street is a free data retrieval call binding the contract method 0x74196626.
//
// Solidity: function street() constant returns(string)
func (_Openvasp *OpenvaspCaller) Street(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "street")
	return *ret0, err
}

// Street is a free data retrieval call binding the contract method 0x74196626.
//
// Solidity: function street() constant returns(string)
func (_Openvasp *OpenvaspSession) Street() (string, error) {
	return _Openvasp.Contract.Street(&_Openvasp.CallOpts)
}

// Street is a free data retrieval call binding the contract method 0x74196626.
//
// Solidity: function street() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) Street() (string, error) {
	return _Openvasp.Contract.Street(&_Openvasp.CallOpts)
}

// Town is a free data retrieval call binding the contract method 0xf5f3468f.
//
// Solidity: function town() constant returns(string)
func (_Openvasp *OpenvaspCaller) Town(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "town")
	return *ret0, err
}

// Town is a free data retrieval call binding the contract method 0xf5f3468f.
//
// Solidity: function town() constant returns(string)
func (_Openvasp *OpenvaspSession) Town() (string, error) {
	return _Openvasp.Contract.Town(&_Openvasp.CallOpts)
}

// Town is a free data retrieval call binding the contract method 0xf5f3468f.
//
// Solidity: function town() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) Town() (string, error) {
	return _Openvasp.Contract.Town(&_Openvasp.CallOpts)
}

// TrustedPeers is a free data retrieval call binding the contract method 0x4e612e79.
//
// Solidity: function trustedPeers(uint256 ) constant returns(address)
func (_Openvasp *OpenvaspCaller) TrustedPeers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "trustedPeers", arg0)
	return *ret0, err
}

// TrustedPeers is a free data retrieval call binding the contract method 0x4e612e79.
//
// Solidity: function trustedPeers(uint256 ) constant returns(address)
func (_Openvasp *OpenvaspSession) TrustedPeers(arg0 *big.Int) (common.Address, error) {
	return _Openvasp.Contract.TrustedPeers(&_Openvasp.CallOpts, arg0)
}

// TrustedPeers is a free data retrieval call binding the contract method 0x4e612e79.
//
// Solidity: function trustedPeers(uint256 ) constant returns(address)
func (_Openvasp *OpenvaspCallerSession) TrustedPeers(arg0 *big.Int) (common.Address, error) {
	return _Openvasp.Contract.TrustedPeers(&_Openvasp.CallOpts, arg0)
}

// Website is a free data retrieval call binding the contract method 0xbeb0a416.
//
// Solidity: function website() constant returns(string)
func (_Openvasp *OpenvaspCaller) Website(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Openvasp.contract.Call(opts, out, "website")
	return *ret0, err
}

// Website is a free data retrieval call binding the contract method 0xbeb0a416.
//
// Solidity: function website() constant returns(string)
func (_Openvasp *OpenvaspSession) Website() (string, error) {
	return _Openvasp.Contract.Website(&_Openvasp.CallOpts)
}

// Website is a free data retrieval call binding the contract method 0xbeb0a416.
//
// Solidity: function website() constant returns(string)
func (_Openvasp *OpenvaspCallerSession) Website() (string, error) {
	return _Openvasp.Contract.Website(&_Openvasp.CallOpts)
}

// AddClaim is a paid mutator transaction binding the contract method 0x69f46779.
//
// Solidity: function addClaim(address ) returns()
func (_Openvasp *OpenvaspTransactor) AddClaim(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.contract.Transact(opts, "addClaim", arg0)
}

// AddClaim is a paid mutator transaction binding the contract method 0x69f46779.
//
// Solidity: function addClaim(address ) returns()
func (_Openvasp *OpenvaspSession) AddClaim(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.AddClaim(&_Openvasp.TransactOpts, arg0)
}

// AddClaim is a paid mutator transaction binding the contract method 0x69f46779.
//
// Solidity: function addClaim(address ) returns()
func (_Openvasp *OpenvaspTransactorSession) AddClaim(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.AddClaim(&_Openvasp.TransactOpts, arg0)
}

// AddTrustedPeer is a paid mutator transaction binding the contract method 0x5e165c26.
//
// Solidity: function addTrustedPeer(address ) returns()
func (_Openvasp *OpenvaspTransactor) AddTrustedPeer(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.contract.Transact(opts, "addTrustedPeer", arg0)
}

// AddTrustedPeer is a paid mutator transaction binding the contract method 0x5e165c26.
//
// Solidity: function addTrustedPeer(address ) returns()
func (_Openvasp *OpenvaspSession) AddTrustedPeer(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.AddTrustedPeer(&_Openvasp.TransactOpts, arg0)
}

// AddTrustedPeer is a paid mutator transaction binding the contract method 0x5e165c26.
//
// Solidity: function addTrustedPeer(address ) returns()
func (_Openvasp *OpenvaspTransactorSession) AddTrustedPeer(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.AddTrustedPeer(&_Openvasp.TransactOpts, arg0)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4401f86a.
//
// Solidity: function removeClaim(address ) returns()
func (_Openvasp *OpenvaspTransactor) RemoveClaim(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.contract.Transact(opts, "removeClaim", arg0)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4401f86a.
//
// Solidity: function removeClaim(address ) returns()
func (_Openvasp *OpenvaspSession) RemoveClaim(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.RemoveClaim(&_Openvasp.TransactOpts, arg0)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4401f86a.
//
// Solidity: function removeClaim(address ) returns()
func (_Openvasp *OpenvaspTransactorSession) RemoveClaim(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.RemoveClaim(&_Openvasp.TransactOpts, arg0)
}

// RemoveTrustedPeer is a paid mutator transaction binding the contract method 0xe64a2376.
//
// Solidity: function removeTrustedPeer(address ) returns()
func (_Openvasp *OpenvaspTransactor) RemoveTrustedPeer(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.contract.Transact(opts, "removeTrustedPeer", arg0)
}

// RemoveTrustedPeer is a paid mutator transaction binding the contract method 0xe64a2376.
//
// Solidity: function removeTrustedPeer(address ) returns()
func (_Openvasp *OpenvaspSession) RemoveTrustedPeer(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.RemoveTrustedPeer(&_Openvasp.TransactOpts, arg0)
}

// RemoveTrustedPeer is a paid mutator transaction binding the contract method 0xe64a2376.
//
// Solidity: function removeTrustedPeer(address ) returns()
func (_Openvasp *OpenvaspTransactorSession) RemoveTrustedPeer(arg0 common.Address) (*types.Transaction, error) {
	return _Openvasp.Contract.RemoveTrustedPeer(&_Openvasp.TransactOpts, arg0)
}
