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
const XCDataABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"datahash\",\"type\":\"string\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"commitNewOwnerData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"datahash\",\"type\":\"string\"}],\"name\":\"getAutherizedAESKeyByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deletePreOwnerData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"datahash\",\"type\":\"string\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"commitData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"encryptedAESKey\",\"type\":\"bytes\"}],\"name\":\"autherizeData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAutherizedDataLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferDidOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDataLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"datahash\",\"type\":\"string\"}],\"name\":\"Autherize\",\"type\":\"event\"}]"

// XCDataBin is the compiled bytecode used for deploying new contracts.
const XCDataBin = `0x6060604052341561000f57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506115c28061005e6000396000f300606060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631910e49d1461009e5780633e33f204146100f4578063402b6d2c1461021757806354f6165c146102dd57806394966c8f1461030b578063a29e81f314610361578063ac01603e146103cb578063df63bbd114610418578063ef0f9f4214610465575b600080fd5b34156100a957600080fd5b6100f260048080359060200190820180359060200191909192908035906020019082018035906020019190919290803590602001908201803590602001919091929050506104a7565b005b34156100ff57600080fd5b61012960048080359060200190820180359060200191909192908035906020019091905050610686565b604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015610173578082015181840152602081019050610158565b50505050905090810190601f1680156101a05780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101d95780820151818401526020810190506101be565b50505050905090810190601f1680156102065780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561022257600080fd5b610262600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190820180359060200191909192905050610925565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102a2578082015181840152602081019050610287565b50505050905090810190601f1680156102cf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e857600080fd5b61030960048080359060200190820180359060200191909192905050610a31565b005b341561031657600080fd5b61035f6004808035906020019082018035906020019190919290803590602001908201803590602001919091929080359060200190820180359060200191909192905050610aec565b005b341561036c57600080fd5b6103c9600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001919091929080359060200190919080359060200190820180359060200191909192905050610d2f565b005b34156103d657600080fd5b610402600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506110f3565b6040518082815260200191505060405180910390f35b341561042357600080fd5b6104636004808035906020019082018035906020019190919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061113c565b005b341561047057600080fd5b61049160048080359060200190820180359060200191909192905050611266565b6040518082815260200191505060405180910390f35b60006002878760405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561052d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561056757600080fd5b6001878760405180838380828437820191505092505050908152602001604051809103902060000180548060010182816105a1919061129c565b9160005260206000209060030201600060606040519081016040528042815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815250909190915060008201518160000155602082015181600101908051906020019061065c9291906112ce565b50604082015181600201908051906020019061067992919061134e565b5050505050505050505050565b60006106906113ce565b6106986113e2565b600060018787604051808383808284378201915050925050509081526020016040518091039020600001805490501115156106d257600080fd5b60018686604051808383808284378201915050925050509081526020016040518091039020600001805490508410151561070b57600080fd5b600186866040518083838082843782019150509250505090815260200160405180910390206000018481548110151561074057fe5b906000526020600020906003020160000154600187876040518083838082843782019150509250505090815260200160405180910390206000018581548110151561078757fe5b906000526020600020906003020160010160018888604051808383808284378201915050925050509081526020016040518091039020600001868154811015156107cd57fe5b9060005260206000209060030201600201818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108735780601f1061084857610100808354040283529160200191610873565b820191906000526020600020905b81548152906001019060200180831161085657829003601f168201915b50505050509150808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561090f5780601f106108e45761010080835404028352916020019161090f565b820191906000526020600020905b8154815290600101906020018083116108f257829003601f168201915b5050505050905092509250925093509350939050565b61092d6113e2565b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083836040518083838082843782019150509250505090815260200160405180910390208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a235780601f106109f857610100808354040283529160200191610a23565b820191906000526020600020905b815481529060010190602001808311610a0657829003601f168201915b505050505090509392505050565b3373ffffffffffffffffffffffffffffffffffffffff166002838360405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610ab057600080fd5b6001828260405180838380828437820191505092505050908152602001604051809103902060008082016000610ae691906113f6565b50505050565b60006002878760405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610bab578073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ba657600080fd5b610c10565b336002888860405180838380828437820191505092505050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600187876040518083838082843782019150509250505090815260200160405180910390206000018054806001018281610c4a919061129c565b9160005260206000209060030201600060606040519081016040528042815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508152509091909150600082015181600001556020820151816001019080519060200190610d059291906112ce565b506040820151816002019080519060200190610d2292919061134e565b5050505050505050505050565b6000806002878760405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610db657600080fd5b8773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610df157600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e2b57600080fd5b600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508383600360008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060018a8a60405180838380828437820191505092505050908152602001604051809103902060000188815481101515610ee357fe5b90600052602060002090600302016001016040518082805460018160011615610100020316600290048015610f4f5780601f10610f2d576101008083540402835291820191610f4f565b820191906000526020600020905b815481529060010190602001808311610f3b575b505091505090815260200160405180910390209190610f6f92919061141a565b5060018101600460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f1712dfe0223c657e473f7157be3a85c919d61050a672e35a3b30d382837e718988600189896040518083838082843782019150509250505090815260200160405180910390206000018781548110151561100e57fe5b9060005260206000209060030201600101604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156110da5780601f106110af576101008083540402835291602001916110da565b820191906000526020600020905b8154815290600101906020018083116110bd57829003601f168201915b5050935050505060405180910390a15050505050505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006002848460405180838380828437820191505092505050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156111c257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156111fc57600080fd5b816002858560405180838380828437820191505092505050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b60006001838360405180838380828437820191505092505050908152602001604051809103902060000180549050905092915050565b8154818355818115116112c9576003028160030283600052602060002091820191016112c8919061149a565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061130f57805160ff191683800117855561133d565b8280016001018555821561133d579182015b8281111561133c578251825591602001919060010190611321565b5b50905061134a91906114e1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061138f57805160ff19168380011785556113bd565b828001600101855582156113bd579182015b828111156113bc5782518255916020019190600101906113a1565b5b5090506113ca91906114e1565b5090565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b5080546000825560030290600052602060002090810190611417919061149a565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061145b57803560ff1916838001178555611489565b82800160010185558215611489579182015b8281111561148857823582559160200191906001019061146d565b5b50905061149691906114e1565b5090565b6114de91905b808211156114da576000808201600090556001820160006114c19190611506565b6002820160006114d1919061154e565b506003016114a0565b5090565b90565b61150391905b808211156114ff5760008160009055506001016114e7565b5090565b90565b50805460018160011615610100020316600290046000825580601f1061152c575061154b565b601f01602090049060005260206000209081019061154a91906114e1565b5b50565b50805460018160011615610100020316600290046000825580601f106115745750611593565b601f01602090049060005260206000209081019061159291906114e1565b5b505600a165627a7a723058208540063349cff9c56dc02740f84874ecebd755768b482777bcf9be99ddaac93f0029`

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

// GetAutherizedAESKeyByHash is a free data retrieval call binding the contract method 0x402b6d2c.
//
// Solidity: function getAutherizedAESKeyByHash(addr address, datahash string) constant returns(bytes)
func (_XCData *XCDataCaller) GetAutherizedAESKeyByHash(opts *bind.CallOpts, addr common.Address, datahash string) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _XCData.contract.Call(opts, out, "getAutherizedAESKeyByHash", addr, datahash)
	return *ret0, err
}

// GetAutherizedAESKeyByHash is a free data retrieval call binding the contract method 0x402b6d2c.
//
// Solidity: function getAutherizedAESKeyByHash(addr address, datahash string) constant returns(bytes)
func (_XCData *XCDataSession) GetAutherizedAESKeyByHash(addr common.Address, datahash string) ([]byte, error) {
	return _XCData.Contract.GetAutherizedAESKeyByHash(&_XCData.CallOpts, addr, datahash)
}

// GetAutherizedAESKeyByHash is a free data retrieval call binding the contract method 0x402b6d2c.
//
// Solidity: function getAutherizedAESKeyByHash(addr address, datahash string) constant returns(bytes)
func (_XCData *XCDataCallerSession) GetAutherizedAESKeyByHash(addr common.Address, datahash string) ([]byte, error) {
	return _XCData.Contract.GetAutherizedAESKeyByHash(&_XCData.CallOpts, addr, datahash)
}

// GetAutherizedDataLength is a free data retrieval call binding the contract method 0xac01603e.
//
// Solidity: function getAutherizedDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataCaller) GetAutherizedDataLength(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XCData.contract.Call(opts, out, "getAutherizedDataLength", addr)
	return *ret0, err
}

// GetAutherizedDataLength is a free data retrieval call binding the contract method 0xac01603e.
//
// Solidity: function getAutherizedDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataSession) GetAutherizedDataLength(addr common.Address) (*big.Int, error) {
	return _XCData.Contract.GetAutherizedDataLength(&_XCData.CallOpts, addr)
}

// GetAutherizedDataLength is a free data retrieval call binding the contract method 0xac01603e.
//
// Solidity: function getAutherizedDataLength(addr address) constant returns(uint256)
func (_XCData *XCDataCallerSession) GetAutherizedDataLength(addr common.Address) (*big.Int, error) {
	return _XCData.Contract.GetAutherizedDataLength(&_XCData.CallOpts, addr)
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
