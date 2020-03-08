// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnsrecord

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

// DnsrecordABI is the input ABI used to generate the binding from.
const DnsrecordABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"recType\",\"type\":\"int8\"},{\"internalType\":\"string\",\"name\":\"recValue\",\"type\":\"string\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"recType\",\"type\":\"int8\"}],\"name\":\"getRecord\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DnsrecordBin is the compiled bytecode used for deploying new contracts.
var DnsrecordBin = "0x608060405234801561001057600080fd5b5061056d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630b36777f1461003b578063a303bcac1461019a575b600080fd5b6101986004803603606081101561005157600080fd5b810190808035906020019064010000000081111561006e57600080fd5b82018360208201111561008057600080fd5b803590602001918460018302840111640100000000831117156100a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560000b90602001909291908035906020019064010000000081111561011257600080fd5b82018360208201111561012457600080fd5b8035906020019184600183028401116401000000008311171561014657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506102db565b005b610260600480360360408110156101b057600080fd5b81019080803590602001906401000000008111156101cd57600080fd5b8201836020820111156101df57600080fd5b8035906020019184600183028401116401000000008311171561020157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560000b9060200190929190505050610372565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102a0578082015181840152602081019050610285565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b806000846040518082805190602001908083835b6020831061031257805182526020820191506020810190506020830392506102ef565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008460000b8152602001908152602001600020908051906020019061036c929190610492565b50505050565b60606000836040518082805190602001908083835b602083106103aa5780518252602082019150602081019050602083039250610387565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008360000b81526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104855780601f1061045a57610100808354040283529160200191610485565b820191906000526020600020905b81548152906001019060200180831161046857829003601f168201915b5050505050905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104d357805160ff1916838001178555610501565b82800160010185558215610501579182015b828111156105005782518255916020019190600101906104e5565b5b50905061050e9190610512565b5090565b61053491905b80821115610530576000816000905550600101610518565b5090565b9056fea264697066735822122008c65502163aab889eaf47f374c45c3bc1e26325faabfabd2b1b616f49e4e4ad64736f6c63430006030033"

// DeployDnsrecord deploys a new Ethereum contract, binding an instance of Dnsrecord to it.
func DeployDnsrecord(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dnsrecord, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsrecordABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DnsrecordBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dnsrecord{DnsrecordCaller: DnsrecordCaller{contract: contract}, DnsrecordTransactor: DnsrecordTransactor{contract: contract}, DnsrecordFilterer: DnsrecordFilterer{contract: contract}}, nil
}

// Dnsrecord is an auto generated Go binding around an Ethereum contract.
type Dnsrecord struct {
	DnsrecordCaller     // Read-only binding to the contract
	DnsrecordTransactor // Write-only binding to the contract
	DnsrecordFilterer   // Log filterer for contract events
}

// DnsrecordCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsrecordCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsrecordTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsrecordTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsrecordFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsrecordFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsrecordSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsrecordSession struct {
	Contract     *Dnsrecord        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsrecordCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsrecordCallerSession struct {
	Contract *DnsrecordCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DnsrecordTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsrecordTransactorSession struct {
	Contract     *DnsrecordTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DnsrecordRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsrecordRaw struct {
	Contract *Dnsrecord // Generic contract binding to access the raw methods on
}

// DnsrecordCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsrecordCallerRaw struct {
	Contract *DnsrecordCaller // Generic read-only contract binding to access the raw methods on
}

// DnsrecordTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsrecordTransactorRaw struct {
	Contract *DnsrecordTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsrecord creates a new instance of Dnsrecord, bound to a specific deployed contract.
func NewDnsrecord(address common.Address, backend bind.ContractBackend) (*Dnsrecord, error) {
	contract, err := bindDnsrecord(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dnsrecord{DnsrecordCaller: DnsrecordCaller{contract: contract}, DnsrecordTransactor: DnsrecordTransactor{contract: contract}, DnsrecordFilterer: DnsrecordFilterer{contract: contract}}, nil
}

// NewDnsrecordCaller creates a new read-only instance of Dnsrecord, bound to a specific deployed contract.
func NewDnsrecordCaller(address common.Address, caller bind.ContractCaller) (*DnsrecordCaller, error) {
	contract, err := bindDnsrecord(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsrecordCaller{contract: contract}, nil
}

// NewDnsrecordTransactor creates a new write-only instance of Dnsrecord, bound to a specific deployed contract.
func NewDnsrecordTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsrecordTransactor, error) {
	contract, err := bindDnsrecord(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsrecordTransactor{contract: contract}, nil
}

// NewDnsrecordFilterer creates a new log filterer instance of Dnsrecord, bound to a specific deployed contract.
func NewDnsrecordFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsrecordFilterer, error) {
	contract, err := bindDnsrecord(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsrecordFilterer{contract: contract}, nil
}

// bindDnsrecord binds a generic wrapper to an already deployed contract.
func bindDnsrecord(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsrecordABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dnsrecord *DnsrecordRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dnsrecord.Contract.DnsrecordCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dnsrecord *DnsrecordRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dnsrecord.Contract.DnsrecordTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dnsrecord *DnsrecordRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dnsrecord.Contract.DnsrecordTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dnsrecord *DnsrecordCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dnsrecord.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dnsrecord *DnsrecordTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dnsrecord.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dnsrecord *DnsrecordTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dnsrecord.Contract.contract.Transact(opts, method, params...)
}

// GetRecord is a free data retrieval call binding the contract method 0xa303bcac.
//
// Solidity: function getRecord(string key, int8 recType) constant returns(string)
func (_Dnsrecord *DnsrecordCaller) GetRecord(opts *bind.CallOpts, key string, recType int8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dnsrecord.contract.Call(opts, out, "getRecord", key, recType)
	return *ret0, err
}

// GetRecord is a free data retrieval call binding the contract method 0xa303bcac.
//
// Solidity: function getRecord(string key, int8 recType) constant returns(string)
func (_Dnsrecord *DnsrecordSession) GetRecord(key string, recType int8) (string, error) {
	return _Dnsrecord.Contract.GetRecord(&_Dnsrecord.CallOpts, key, recType)
}

// GetRecord is a free data retrieval call binding the contract method 0xa303bcac.
//
// Solidity: function getRecord(string key, int8 recType) constant returns(string)
func (_Dnsrecord *DnsrecordCallerSession) GetRecord(key string, recType int8) (string, error) {
	return _Dnsrecord.Contract.GetRecord(&_Dnsrecord.CallOpts, key, recType)
}

// AddRecord is a paid mutator transaction binding the contract method 0x0b36777f.
//
// Solidity: function addRecord(string key, int8 recType, string recValue) returns()
func (_Dnsrecord *DnsrecordTransactor) AddRecord(opts *bind.TransactOpts, key string, recType int8, recValue string) (*types.Transaction, error) {
	return _Dnsrecord.contract.Transact(opts, "addRecord", key, recType, recValue)
}

// AddRecord is a paid mutator transaction binding the contract method 0x0b36777f.
//
// Solidity: function addRecord(string key, int8 recType, string recValue) returns()
func (_Dnsrecord *DnsrecordSession) AddRecord(key string, recType int8, recValue string) (*types.Transaction, error) {
	return _Dnsrecord.Contract.AddRecord(&_Dnsrecord.TransactOpts, key, recType, recValue)
}

// AddRecord is a paid mutator transaction binding the contract method 0x0b36777f.
//
// Solidity: function addRecord(string key, int8 recType, string recValue) returns()
func (_Dnsrecord *DnsrecordTransactorSession) AddRecord(key string, recType int8, recValue string) (*types.Transaction, error) {
	return _Dnsrecord.Contract.AddRecord(&_Dnsrecord.TransactOpts, key, recType, recValue)
}
