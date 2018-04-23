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
const XCDataABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"datahash\",\"type\":\"string\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"commitNewOwnerData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAuthorizedDataLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"datahash\",\"type\":\"string\"}],\"name\":\"getAuthorizedAESKeyByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deletePreOwnerData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"datahash\",\"type\":\"string\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"commitData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"authorizeData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferDidOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDataLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"datahash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"NewCommitData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"datahash\",\"type\":\"string\"}],\"name\":\"Authorize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"did\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"TransferDid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DeleteDid\",\"type\":\"event\"}]"

// XCDataBin is the compiled bytecode used for deploying new contracts.
const XCDataBin = `0x6060604052341561000f57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118278061005e6000396000f300606060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631910e49d1461009e5780633498f326146100f45780633e33f204146101415780634d95ba291461026457806354f6165c1461032a57806394966c8f14610358578063ad10287d146103ae578063df63bbd114610418578063ef0f9f4214610465575b600080fd5b34156100a957600080fd5b6100f260048080359060200190820180359060200191909192908035906020019082018035906020019190919290803590602001908201803590602001919091929050506104a7565b005b34156100ff57600080fd5b61012b600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610748565b6040518082815260200191505060405180910390f35b341561014c57600080fd5b61017660048080359060200190820180359060200191909192908035906020019091905050610791565b604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156101c05780820151818401526020810190506101a5565b50505050905090810190601f1680156101ed5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561022657808201518184015260208101905061020b565b50505050905090810190601f1680156102535780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561026f57600080fd5b6102af600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190820180359060200191909192905050610a30565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102ef5780820151818401526020810190506102d4565b50505050905090810190601f16801561031c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561033557600080fd5b61035660048080359060200190820180359060200191909192905050610b3c565b005b341561036357600080fd5b6103ac6004808035906020019082018035906020019190919290803590602001908201803590602001919091929080359060200190820180359060200191909192905050610c5f565b005b34156103b957600080fd5b610416600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001919091929080359060200190919080359060200190820180359060200191909192905050610f64565b005b341561042357600080fd5b6104636004808035906020019082018035906020019190919290803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611322565b005b341561047057600080fd5b610491600480803590602001908201803590602001919091929050506114cb565b6040518082815260200191505060405180910390f35b6000806002888860405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561052e57600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561056857600080fd5b6001888860405180838380828437820191505092505050908152602001604051809103902060000180548060010182816105a29190611501565b916000526020600020906003020160006060604051908101604052804281526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815250909190915060008201518160000155602082015181600101908051906020019061065d929190611533565b50604082015181600201908051906020019061067a9291906115b3565b50505050600188886040518083838082843782019150509250505090815260200160405180910390206000018054905090503373ffffffffffffffffffffffffffffffffffffffff167fbe20497c5b443fdd6acb9daf8ccce79dd423acc869b205234fd5cb5670fd70f38989898986604051808060200180602001848152602001838103835288888281815260200192508082843782019150508381038252868682818152602001925080828437820191505097505050505050505060405180910390a25050505050505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600061079b611633565b6107a3611647565b600060018787604051808383808284378201915050925050509081526020016040518091039020600001805490501115156107dd57600080fd5b60018686604051808383808284378201915050925050509081526020016040518091039020600001805490508410151561081657600080fd5b600186866040518083838082843782019150509250505090815260200160405180910390206000018481548110151561084b57fe5b906000526020600020906003020160000154600187876040518083838082843782019150509250505090815260200160405180910390206000018581548110151561089257fe5b906000526020600020906003020160010160018888604051808383808284378201915050925050509081526020016040518091039020600001868154811015156108d857fe5b9060005260206000209060030201600201818054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561097e5780601f106109535761010080835404028352916020019161097e565b820191906000526020600020905b81548152906001019060200180831161096157829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a1a5780601f106109ef57610100808354040283529160200191610a1a565b820191906000526020600020905b8154815290600101906020018083116109fd57829003601f168201915b5050505050905092509250925093509350939050565b610a38611647565b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083836040518083838082843782019150509250505090815260200160405180910390208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b2e5780601f10610b0357610100808354040283529160200191610b2e565b820191906000526020600020905b815481529060010190602001808311610b1157829003601f168201915b505050505090509392505050565b3373ffffffffffffffffffffffffffffffffffffffff166002838360405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610bbb57600080fd5b6001828260405180838380828437820191505092505050908152602001604051809103902060008082016000610bf1919061165b565b50503373ffffffffffffffffffffffffffffffffffffffff167f13f9ad3f1bd92d9062334f3af088ffe250b1ec2816aac72fd4e60d8dfa651e2c8383604051808060200182810382528484828181526020019250808284378201915050935050505060405180910390a25050565b6000806002888860405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515610d1f578173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d1a57600080fd5b610d84565b336002898960405180838380828437820191505092505050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600188886040518083838082843782019150509250505090815260200160405180910390206000018054806001018281610dbe9190611501565b916000526020600020906003020160006060604051908101604052804281526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508152509091909150600082015181600001556020820151816001019080519060200190610e79929190611533565b506040820151816002019080519060200190610e969291906115b3565b50505050600188886040518083838082843782019150509250505090815260200160405180910390206000018054905090503373ffffffffffffffffffffffffffffffffffffffff167fbe20497c5b443fdd6acb9daf8ccce79dd423acc869b205234fd5cb5670fd70f38989898986604051808060200180602001848152602001838103835288888281815260200192508082843782019150508381038252868682818152602001925080828437820191505097505050505050505060405180910390a25050505050505050565b6000806002878760405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610feb57600080fd5b8773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561102657600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561106057600080fd5b600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508383600360008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060018a8a6040518083838082843782019150509250505090815260200160405180910390206000018881548110151561111857fe5b906000526020600020906003020160010160405180828054600181600116156101000203166002900480156111845780601f10611162576101008083540402835291820191611184565b820191906000526020600020905b815481529060010190602001808311611170575b5050915050908152602001604051809103902091906111a492919061167f565b5060018101600460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508773ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8a6fa22cbcd96428dae964201b351427af5a5a55c3239abdbe74b8771906cd9760018a8a6040518083838082843782019150509250505090815260200160405180910390206000018881548110151561127057fe5b9060005260206000209060030201600101604051808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561130a5780601f106112df5761010080835404028352916020019161130a565b820191906000526020600020905b8154815290600101906020018083116112ed57829003601f168201915b50509250505060405180910390a35050505050505050565b60006002848460405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156113a857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113e257600080fd5b816002858560405180838380828437820191505092505050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fa0f57e6ab8a104c5d04604dbe08125bafbdfad75e270a85687c9f87dda14dc548686604051808060200182810382528484828181526020019250808284378201915050935050505060405180910390a350505050565b60006001838360405180838380828437820191505092505050908152602001604051809103902060000180549050905092915050565b81548183558181151161152e5760030281600302836000526020600020918201910161152d91906116ff565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061157457805160ff19168380011785556115a2565b828001600101855582156115a2579182015b828111156115a1578251825591602001919060010190611586565b5b5090506115af9190611746565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115f457805160ff1916838001178555611622565b82800160010185558215611622579182015b82811115611621578251825591602001919060010190611606565b5b50905061162f9190611746565b5090565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b508054600082556003029060005260206000209081019061167c91906116ff565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116c057803560ff19168380011785556116ee565b828001600101855582156116ee579182015b828111156116ed5782358255916020019190600101906116d2565b5b5090506116fb9190611746565b5090565b61174391905b8082111561173f57600080820160009055600182016000611726919061176b565b60028201600061173691906117b3565b50600301611705565b5090565b90565b61176891905b8082111561176457600081600090555060010161174c565b5090565b90565b50805460018160011615610100020316600290046000825580601f1061179157506117b0565b601f0160209004906000526020600020908101906117af9190611746565b5b50565b50805460018160011615610100020316600290046000825580601f106117d957506117f8565b601f0160209004906000526020600020908101906117f79190611746565b5b505600a165627a7a72305820e6b10dc65470b3352d5890e23891445e98365616da8d423719e8683d6c6a3c180029`

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
	return address, tx, &XCData{XCDataCaller: XCDataCaller{contract: contract}, XCDataTransactor: XCDataTransactor{contract: contract}}, nil
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

// GetAuthorizedAESKeyByHash is a free data retrieval call binding the contract method 0x402b6d2c.
//
// Solidity: function getAuthorizedAESKeyByHash(addr address, datahash string) constant returns(bytes)
func (_XCData *XCDataCaller) GetAuthorizedAESKeyByHash(opts *bind.CallOpts, addr common.Address, datahash string) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _XCData.contract.Call(opts, out, "getAuthorizedAESKeyByHash", addr, datahash)
	return *ret0, err
}

// GetAuthorizedAESKeyByHash is a free data retrieval call binding the contract method 0x402b6d2c.
//
// Solidity: function getAuthorizedAESKeyByHash(addr address, datahash string) constant returns(bytes)
func (_XCData *XCDataSession) GetAuthorizedAESKeyByHash(addr common.Address, datahash string) ([]byte, error) {
	return _XCData.Contract.GetAuthorizedAESKeyByHash(&_XCData.CallOpts, addr, datahash)
}

// GetAuthorizedAESKeyByHash is a free data retrieval call binding the contract method 0x402b6d2c.
//
// Solidity: function getAuthorizedAESKeyByHash(addr address, datahash string) constant returns(bytes)
func (_XCData *XCDataCallerSession) GetAuthorizedAESKeyByHash(addr common.Address, datahash string) ([]byte, error) {
	return _XCData.Contract.GetAuthorizedAESKeyByHash(&_XCData.CallOpts, addr, datahash)
}

// GetAuthorizedDataLength is a free data retrieval call binding the contract method 0xac01603e.
//
// Solidity: function getAuthorizedDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataCaller) GetAuthorizedDataLength(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XCData.contract.Call(opts, out, "getAuthorizedDataLength", addr)
	return *ret0, err
}

// GetAuthorizedDataLength is a free data retrieval call binding the contract method 0xac01603e.
//
// Solidity: function getAuthorizedDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataSession) GetAuthorizedDataLength(addr common.Address) (*big.Int, error) {
	return _XCData.Contract.GetAuthorizedDataLength(&_XCData.CallOpts, addr)
}

// GetAuthorizedDataLength is a free data retrieval call binding the contract method 0xac01603e.
//
// Solidity: function getAuthorizedDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataCallerSession) GetAuthorizedDataLength(addr common.Address) (*big.Int, error) {
	return _XCData.Contract.GetAuthorizedDataLength(&_XCData.CallOpts, addr)
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

// AuthorizeData is a paid mutator transaction binding the contract method 0xa29e81f3.
//
// Solidity: function authorizeData(to address, did string, index uint256, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactor) AuthorizeData(opts *bind.TransactOpts, to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.contract.Transact(opts, "authorizeData", to, did, index, encryptedAESKey)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xa29e81f3.
//
// Solidity: function authorizeData(to address, did string, index uint256, encryptedAESKey bytes) returns()
func (_XCData *XCDataSession) AuthorizeData(to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.AuthorizeData(&_XCData.TransactOpts, to, did, index, encryptedAESKey)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xa29e81f3.
//
// Solidity: function authorizeData(to address, did string, index uint256, encryptedAESKey bytes) returns()
func (_XCData *XCDataTransactorSession) AuthorizeData(to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return _XCData.Contract.AuthorizeData(&_XCData.TransactOpts, to, did, index, encryptedAESKey)
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
