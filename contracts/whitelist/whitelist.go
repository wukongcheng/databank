package whitelist

//go:generate abigen --sol contract/whitelist.sol --pkg contract --out contract/whitelist.go

import (
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/common"
	//"github.com/xcareteam/xci/eth"
	"github.com/xcareteam/xci/contracts/whitelist/contract"
	"github.com/xcareteam/xci/core/types"
	//"github.com/xcareteam/xci/node"
	//"github.com/xcareteam/xci/internal/ethapi"
	//"github.com/xcareteam/xci/les"
	"github.com/xcareteam/xci/accounts"
)

var (
	MainNetAddress = common.HexToAddress("0x314159265dD8dbb310642f98f50C066173C1259b")
	TestNetAddress = common.HexToAddress("0x112234455c3a32fd11230c42e7bccd4a84e02010")
)

type WhiteList struct {
	*contract.WhitelistSession
	contractBackend bind.ContractBackend
}

func NewWhiteList(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*WhiteList, error) {
	whitelist, err := contract.NewWhitelist(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &WhiteList{
		&contract.WhitelistSession{
			Contract:     whitelist,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func NewWhiteListCaller(contractAddr common.Address, contractBackend bind.ContractBackend) (*contract.WhitelistCaller, error) {
	caller, err := contract.NewWhitelistCaller(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return caller, nil
}

func DeployWhiteList(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *WhiteList, error) {
	whitelistAddr, _, _, err := contract.DeployWhitelist(transactOpts, contractBackend)
	if err != nil {
		return whitelistAddr, nil, err
	}

	whitelist, err := NewWhiteList(transactOpts, whitelistAddr, contractBackend)
	if err != nil {
		return whitelistAddr, nil, err
	}

	return whitelistAddr, whitelist, nil
}

func (self *WhiteList) AddNewNode(enode string, DIDJson string) (*types.Transaction, error) {
	return self.Contract.AddNewNode(&self.TransactOpts, enode, DIDJson)
}

func (self *WhiteList) GetDID(enode string) (string, error) {
	return self.Contract.GetDID(&self.CallOpts, enode);
}

//func getEthereumBackend(node *node.Node) (*eth.Ethereum, ethapi.Backend, error) {
//	var apiBackend ethapi.Backend
//	var ethereum *eth.Ethereum
//	if err := node.Service(&ethereum); err == nil {
//		apiBackend = ethereum.ApiBackend
//	} else {
//		var ethereum *les.LightEthereum
//		if err := node.Service(&ethereum); err == nil {
//			apiBackend = ethereum.ApiBackend
//		} else {
//			return nil, nil, err
//		}
//	}
//
//	return ethereum, apiBackend, nil
//}

func GetNewWhiteList(accMng *accounts.Manager, backend bind.ContractBackend, address common.Address, passphrase string) (*WhiteList, error) {

	account := accounts.Account{Address: address}
	wallet, err := accMng.Find(account)
	if err != nil {
		return nil, err
	}

	transactOpts, err := wallet.NewKeyedTransactor(account, passphrase)
	if err != nil {
		return nil, err
	}

	contract, err := NewWhiteList(transactOpts, TestNetAddress, backend)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

//func GetNewWhiteListCaller(node *node.Node) (*contract.WhitelistCaller, error) {
//	_, apiBackend, err := getEthereumBackend(node)
//	if err != nil {
//		return nil, err
//	}
//
//	contractBackend2 := eth.NewContractBackend(apiBackend)
//	var contractBackend bind.ContractBackend
//	contractBackend = contractBackend2
//
//	return NewWhiteListCaller(TestNetAddress, contractBackend)
//}
