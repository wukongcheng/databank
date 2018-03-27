// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/xcareteam/xci/accounts/abi"
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/core/types"
)

// XCDataABI is the input ABI used to generate the binding from.
const XCDataABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"datahash\",\"type\":\"string\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"commitNewOwnerData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deletePreOwnerData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAutherizeDataLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"datahash\",\"type\":\"string\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"commitData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"autherizeData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferDidOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAutherizeData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDataLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// XCDataBin is the compiled bytecode used for deploying new contracts.
const XCDataBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610f3a8061003b6000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631910e49d811461009d5780633e33f204146100d557806354f6165c146101df578063688a3ddc146101fd57806394966c8f1461022e578063a29e81f314610264578063df63bbd11461029f578063e7430c0a146102cb578063ef0f9f42146102ed575b600080fd5b34156100a857600080fd5b6100d3602460048035828101929082013591813580830192908201359160443591820191013561030b565b005b34156100e057600080fd5b6100f7602460048035828101929101359035610478565b604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561013e578082015183820152602001610126565b50505050905090810190601f16801561016b5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101a1578082015183820152602001610189565b50505050905090810190601f1680156101ce5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34156101ea57600080fd5b6100d36004803560248101910135610700565b341561020857600080fd5b61021c600160a060020a036004351661077c565b60405190815260200160405180910390f35b341561023957600080fd5b6100d36024600480358281019290820135918135808301929082013591604435918201910135610797565b341561026f57600080fd5b6100d360048035600160a060020a031690602480358082019290810135916044359160643591820191013561087e565b34156102aa57600080fd5b6100d36024600480358281019291013590600160a060020a03903516610aa2565b34156102d657600080fd5b6100f7600160a060020a0360043516602435610b5a565b34156102f857600080fd5b61021c6004803560248101910135610d61565b6000600387876040518083838082843782019150509250505090815260200160405190819003902054600160a060020a0316905080151561034b57600080fd5b80600160a060020a031633600160a060020a031614151561036b57600080fd5b60018787604051808383808284378201915050925050509081526020016040519081900390208054600181016103a18382610d91565b9160005260206000209060030201600060606040519081016040528042815260200189898080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050815260200187878080601f016020809104026020016040519081016040528181529291906020840183838082843750505092909352509193925083915050518155602082015181600101908051610450929160200190610dc2565b5060408201518160020190805161046b929160200190610dc2565b5050505050505050505050565b6000610482610e40565b61048a610e40565b6000600187876040518083838082843782019150509250505090815260200160405190819003902054116104bd57600080fd5b60018686604051808383808284378201915050925050509081526020016040519081900390205484106104ef57600080fd5b600186866040518083838082843782019150509250505090815260200160405190819003902080548590811061052157fe5b906000526020600020906003020160000154600187876040518083838082843782019150509250505090815260200160405190819003902080548690811061056557fe5b906000526020600020906003020160010160018888604051808383808284378201915050925050509081526020016040519081900390208054879081106105a857fe5b9060005260206000209060030201600201818054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561064e5780601f106106235761010080835404028352916020019161064e565b820191906000526020600020905b81548152906001019060200180831161063157829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106ea5780601f106106bf576101008083540402835291602001916106ea565b820191906000526020600020905b8154815290600101906020018083116106cd57829003601f168201915b5050505050905092509250925093509350939050565b33600160a060020a0316600383836040518083838082843782019150509250505090815260200160405190819003902054600160a060020a03161461074457600080fd5b600182826040518083838082843782019150509250505090815260200160405190819003902060006107768282610e52565b50505050565b600160a060020a031660009081526002602052604090205490565b6000600387876040518083838082843782019150509250505090815260200160405190819003902054600160a060020a0316905080156107f65780600160a060020a031633600160a060020a03161415156107f157600080fd5b61036b565b336003888860405180838380828437820191505092505050908152602001604051908190039020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560018787604051808383808284378201915050925050509081526020016040519081900390208054600181016103a18382610d91565b6000600386866040518083838082843782019150509250505090815260200160405190819003902054600160a060020a031690508015156108be57600080fd5b80600160a060020a031633600160a060020a03161415156108de57600080fd5b600160a060020a03871660009081526002602052604090208054600181016109068382610d91565b9160005260206000209060030201600060606040519081016040528060018b8b6040518083838082843782019150509250505090815260200160405190819003902080548a90811061095457fe5b906000526020600020906003020160000154815260200160018b8b6040518083838082843782019150509250505090815260200160405190819003902080548a90811061099d57fe5b90600052602060002090600302016001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a425780601f10610a1757610100808354040283529160200191610a42565b820191906000526020600020905b815481529060010190602001808311610a2557829003601f168201915b5050505050815260200187878080601f016020809104026020016040519081016040528181529291906020840183838082843750505092909352509193925083915050518155602082015181600101908051610450929160200190610dc2565b6000600384846040518083838082843782019150509250505090815260200160405190819003902054600160a060020a03169050801515610ae257600080fd5b80600160a060020a031633600160a060020a0316141515610b0257600080fd5b816003858560405180838380828437820191505092505050908152602001604051908190039020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905550505050565b6000610b64610e40565b610b6c610e40565b600160a060020a0385166000908152600260205260409020805485908110610b9057fe5b60009182526020808320600390920290910154600160a060020a038816835260029091526040909120805486908110610bc557fe5b90600052602060002090600302016001016002600088600160a060020a0316600160a060020a0316815260200190815260200160002060000186815481101515610c0b57fe5b9060005260206000209060030201600201818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cb15780601f10610c8657610100808354040283529160200191610cb1565b820191906000526020600020905b815481529060010190602001808311610c9457829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d4d5780601f10610d2257610100808354040283529160200191610d4d565b820191906000526020600020905b815481529060010190602001808311610d3057829003601f168201915b505050505090509250925092509250925092565b60006001838360405180838380828437820191505092505050908152602001604051908190039020549392505050565b815481835581811511610dbd57600302816003028360005260206000209182019101610dbd9190610e76565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e0357805160ff1916838001178555610e30565b82800160010185558215610e30579182015b82811115610e30578251825591602001919060010190610e15565b50610e3c929150610eb0565b5090565b60206040519081016040526000815290565b5080546000825560030290600052602060002090810190610e739190610e76565b50565b610ead91905b80821115610e3c576000808255610e966001830182610eca565b610ea4600283016000610eca565b50600301610e7c565b90565b610ead91905b80821115610e3c5760008155600101610eb6565b50805460018160011615610100020316600290046000825580601f10610ef05750610e73565b601f016020900490600052602060002090810190610e739190610eb05600a165627a7a723058204d451149848389c3971650b2a8ec7d8e5d43c8213ec7b7629f3e6040e2a8c2e20029`

// DeployXCData deploys a new Ethereum contract, binding an instance of XCData to it.
func DeployXCData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XCData, error) {
	parsed, err := abi.JSON(strings.NewReader(XCDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XCDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XCData{XCDataCaller: XCDataCaller{contract: contract}, XCDataTransactor: XCDataTransactor{contract: contract}, XCDataFilterer: XCDataFilterer{contract: contract}}, nil
}

// XCData is an auto generated Go binding around an Ethereum contract.
type XCData struct {
	XCDataCaller     // Read-only binding to the contract
	XCDataTransactor // Write-only binding to the contract
	XCDataFilterer   // Log filterer for contract events
}

// XCDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type XCDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XCDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XCDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XCDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XCDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XCDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XCDataSession struct {
	Contract     *XCData           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XCDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XCDataCallerSession struct {
	Contract *XCDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XCDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XCDataTransactorSession struct {
	Contract     *XCDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XCDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type XCDataRaw struct {
	Contract *XCData // Generic contract binding to access the raw methods on
}

// XCDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XCDataCallerRaw struct {
	Contract *XCDataCaller // Generic read-only contract binding to access the raw methods on
}

// XCDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XCDataTransactorRaw struct {
	Contract *XCDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXCData creates a new instance of XCData, bound to a specific deployed contract.
func NewXCData(address common.Address, backend bind.ContractBackend) (*XCData, error) {
	contract, err := bindXCData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XCData{XCDataCaller: XCDataCaller{contract: contract}, XCDataTransactor: XCDataTransactor{contract: contract}, XCDataFilterer: XCDataFilterer{contract: contract}}, nil
}

// NewXCDataCaller creates a new read-only instance of XCData, bound to a specific deployed contract.
func NewXCDataCaller(address common.Address, caller bind.ContractCaller) (*XCDataCaller, error) {
	contract, err := bindXCData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XCDataCaller{contract: contract}, nil
}

// NewXCDataTransactor creates a new write-only instance of XCData, bound to a specific deployed contract.
func NewXCDataTransactor(address common.Address, transactor bind.ContractTransactor) (*XCDataTransactor, error) {
	contract, err := bindXCData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XCDataTransactor{contract: contract}, nil
}

// NewXCDataFilterer creates a new log filterer instance of XCData, bound to a specific deployed contract.
func NewXCDataFilterer(address common.Address, filterer bind.ContractFilterer) (*XCDataFilterer, error) {
	contract, err := bindXCData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XCDataFilterer{contract: contract}, nil
}

// bindXCData binds a generic wrapper to an already deployed contract.
func bindXCData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XCDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XCData *XCDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XCData.Contract.XCDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XCData *XCDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XCData.Contract.XCDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XCData *XCDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XCData.Contract.XCDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XCData *XCDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XCData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XCData *XCDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XCData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XCData *XCDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XCData.Contract.contract.Transact(opts, method, params...)
}

// GetAutherizeData is a free data retrieval call binding the contract method 0xe7430c0a.
//
// Solidity: function getAutherizeData(addr address, index uint256) constant returns(uint256, string, bytes)
func (_XCData *XCDataCaller) GetAutherizeData(opts *bind.CallOpts, addr common.Address, index *big.Int) (*big.Int, string, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
		ret2 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _XCData.contract.Call(opts, out, "getAutherizeData", addr, index)
	return *ret0, *ret1, *ret2, err
}

// GetAutherizeData is a free data retrieval call binding the contract method 0xe7430c0a.
//
// Solidity: function getAutherizeData(addr address, index uint256) constant returns(uint256, string, bytes)
func (_XCData *XCDataSession) GetAutherizeData(addr common.Address, index *big.Int) (*big.Int, string, []byte, error) {
	return _XCData.Contract.GetAutherizeData(&_XCData.CallOpts, addr, index)
}

// GetAutherizeData is a free data retrieval call binding the contract method 0xe7430c0a.
//
// Solidity: function getAutherizeData(addr address, index uint256) constant returns(uint256, string, bytes)
func (_XCData *XCDataCallerSession) GetAutherizeData(addr common.Address, index *big.Int) (*big.Int, string, []byte, error) {
	return _XCData.Contract.GetAutherizeData(&_XCData.CallOpts, addr, index)
}

// GetAutherizeDataLength is a free data retrieval call binding the contract method 0x688a3ddc.
//
// Solidity: function getAutherizeDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataCaller) GetAutherizeDataLength(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XCData.contract.Call(opts, out, "getAutherizeDataLength", addr)
	return *ret0, err
}

// GetAutherizeDataLength is a free data retrieval call binding the contract method 0x688a3ddc.
//
// Solidity: function getAutherizeDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataSession) GetAutherizeDataLength(addr common.Address) (*big.Int, error) {
	return _XCData.Contract.GetAutherizeDataLength(&_XCData.CallOpts, addr)
}

// GetAutherizeDataLength is a free data retrieval call binding the contract method 0x688a3ddc.
//
// Solidity: function getAutherizeDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataCallerSession) GetAutherizeDataLength(addr common.Address) (*big.Int, error) {
	return _XCData.Contract.GetAutherizeDataLength(&_XCData.CallOpts, addr)
}

// GetData is a free data retrieval call binding the contract method 0x3e33f204.
//
// Solidity: function getData(did string, index uint256) constant returns(uint256, string, bytes)
func (_XCData *XCDataCaller) GetData(opts *bind.CallOpts, did string, index *big.Int) (*big.Int, string, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
		ret2 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _XCData.contract.Call(opts, out, "getData", did, index)
	return *ret0, *ret1, *ret2, err
}

// GetData is a free data retrieval call binding the contract method 0x3e33f204.
//
// Solidity: function getData(did string, index uint256) constant returns(uint256, string, bytes)
func (_XCData *XCDataSession) GetData(did string, index *big.Int) (*big.Int, string, []byte, error) {
	return _XCData.Contract.GetData(&_XCData.CallOpts, did, index)
}

// GetData is a free data retrieval call binding the contract method 0x3e33f204.
//
// Solidity: function getData(did string, index uint256) constant returns(uint256, string, bytes)
func (_XCData *XCDataCallerSession) GetData(did string, index *big.Int) (*big.Int, string, []byte, error) {
	return _XCData.Contract.GetData(&_XCData.CallOpts, did, index)
}

// GetDataLength is a free data retrieval call binding the contract method 0xef0f9f42.
//
// Solidity: function getDataLength(did string) constant returns(uint256)
func (_XCData *XCDataCaller) GetDataLength(opts *bind.CallOpts, did string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XCData.contract.Call(opts, out, "getDataLength", did)
	return *ret0, err
}

// GetDataLength is a free data retrieval call binding the contract method 0xef0f9f42.
//
// Solidity: function getDataLength(did string) constant returns(uint256)
func (_XCData *XCDataSession) GetDataLength(did string) (*big.Int, error) {
	return _XCData.Contract.GetDataLength(&_XCData.CallOpts, did)
}

// GetDataLength is a free data retrieval call binding the contract method 0xef0f9f42.
//
// Solidity: function getDataLength(did string) constant returns(uint256)
func (_XCData *XCDataCallerSession) GetDataLength(did string) (*big.Int, error) {
	return _XCData.Contract.GetDataLength(&_XCData.CallOpts, did)
}

// AutherizeData is a paid mutator transaction binding the contract method 0xa29e81f3.
//
// Solidity: function autherizeData(to address, did string, index uint256, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactor) AutherizeData(opts *bind.TransactOpts, to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.contract.Transact(opts, "autherizeData", to, did, index, encryptedAESKey)
}

// AutherizeData is a paid mutator transaction binding the contract method 0xa29e81f3.
//
// Solidity: function autherizeData(to address, did string, index uint256, encryptedAESKey bytes) returns()
func (_XCData *XCDataSession) AutherizeData(to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.AutherizeData(&_XCData.TransactOpts, to, did, index, encryptedAESKey)
}

// AutherizeData is a paid mutator transaction binding the contract method 0xa29e81f3.
//
// Solidity: function autherizeData(to address, did string, index uint256, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactorSession) AutherizeData(to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.AutherizeData(&_XCData.TransactOpts, to, did, index, encryptedAESKey)
}

// CommitData is a paid mutator transaction binding the contract method 0x94966c8f.
//
// Solidity: function commitData(did string, datahash string, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactor) CommitData(opts *bind.TransactOpts, did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.contract.Transact(opts, "commitData", did, datahash, encryptedAESKey)
}

// CommitData is a paid mutator transaction binding the contract method 0x94966c8f.
//
// Solidity: function commitData(did string, datahash string, encryptedAESKey bytes) returns()
func (_XCData *XCDataSession) CommitData(did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.CommitData(&_XCData.TransactOpts, did, datahash, encryptedAESKey)
}

// CommitData is a paid mutator transaction binding the contract method 0x94966c8f.
//
// Solidity: function commitData(did string, datahash string, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactorSession) CommitData(did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.CommitData(&_XCData.TransactOpts, did, datahash, encryptedAESKey)
}

// CommitNewOwnerData is a paid mutator transaction binding the contract method 0x1910e49d.
//
// Solidity: function commitNewOwnerData(did string, datahash string, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactor) CommitNewOwnerData(opts *bind.TransactOpts, did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.contract.Transact(opts, "commitNewOwnerData", did, datahash, encryptedAESKey)
}

// CommitNewOwnerData is a paid mutator transaction binding the contract method 0x1910e49d.
//
// Solidity: function commitNewOwnerData(did string, datahash string, encryptedAESKey bytes) returns()
func (_XCData *XCDataSession) CommitNewOwnerData(did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.CommitNewOwnerData(&_XCData.TransactOpts, did, datahash, encryptedAESKey)
}

// CommitNewOwnerData is a paid mutator transaction binding the contract method 0x1910e49d.
//
// Solidity: function commitNewOwnerData(did string, datahash string, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactorSession) CommitNewOwnerData(did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.CommitNewOwnerData(&_XCData.TransactOpts, did, datahash, encryptedAESKey)
}

// DeletePreOwnerData is a paid mutator transaction binding the contract method 0x54f6165c.
//
// Solidity: function deletePreOwnerData(did string) returns()
func (_XCData *XCDataTransactor) DeletePreOwnerData(opts *bind.TransactOpts, did string) (*types.Transaction, error) {
	return _XCData.contract.Transact(opts, "deletePreOwnerData", did)
}

// DeletePreOwnerData is a paid mutator transaction binding the contract method 0x54f6165c.
//
// Solidity: function deletePreOwnerData(did string) returns()
func (_XCData *XCDataSession) DeletePreOwnerData(did string) (*types.Transaction, error) {
	return _XCData.Contract.DeletePreOwnerData(&_XCData.TransactOpts, did)
}

// DeletePreOwnerData is a paid mutator transaction binding the contract method 0x54f6165c.
//
// Solidity: function deletePreOwnerData(did string) returns()
func (_XCData *XCDataTransactorSession) DeletePreOwnerData(did string) (*types.Transaction, error) {
	return _XCData.Contract.DeletePreOwnerData(&_XCData.TransactOpts, did)
}

// TransferDidOwner is a paid mutator transaction binding the contract method 0xdf63bbd1.
//
// Solidity: function transferDidOwner(did string, to address) returns()
func (_XCData *XCDataTransactor) TransferDidOwner(opts *bind.TransactOpts, did string, to common.Address) (*types.Transaction, error) {
	return _XCData.contract.Transact(opts, "transferDidOwner", did, to)
}

// TransferDidOwner is a paid mutator transaction binding the contract method 0xdf63bbd1.
//
// Solidity: function transferDidOwner(did string, to address) returns()
func (_XCData *XCDataSession) TransferDidOwner(did string, to common.Address) (*types.Transaction, error) {
	return _XCData.Contract.TransferDidOwner(&_XCData.TransactOpts, did, to)
}

// TransferDidOwner is a paid mutator transaction binding the contract method 0xdf63bbd1.
//
// Solidity: function transferDidOwner(did string, to address) returns()
func (_XCData *XCDataTransactorSession) TransferDidOwner(did string, to common.Address) (*types.Transaction, error) {
	return _XCData.Contract.TransferDidOwner(&_XCData.TransactOpts, did, to)
}
