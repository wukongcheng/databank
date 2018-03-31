package xcdata

//go:generate abigen --sol contract/xcdata.sol --pkg contract --out contract/xcdata.go

import (
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/contracts/xcdata/contract"
	"github.com/xcareteam/xci/core/types"
	"github.com/xcareteam/xci/accounts"
	"math/big"
)

var (
	MainNetAddress = common.HexToAddress("0x314159265dD8dbb310642f98f50C066173C1259b")
	TestNetAddress = common.HexToAddress("0xf0abf8eed4a5c3bea3ef735a72e7f1d53a0e3f55")
)

type XCData struct {
	*contract.XCDataSession
	contractBackend bind.ContractBackend
}

func NewXCData(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*XCData, error) {
	xcdata, err := contract.NewXCData(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &XCData{
		&contract.XCDataSession{
			Contract:     xcdata,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployXCData(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *XCData, error) {
	xcdataAddr, _, _, err := contract.DeployXCData(transactOpts, contractBackend)
	if err != nil {
		return xcdataAddr, nil, err
	}

	xcdata, err := NewXCData(transactOpts, xcdataAddr, contractBackend)
	if err != nil {
		return xcdataAddr, nil, err
	}

	return xcdataAddr, xcdata, nil
}

func GetXCData(accMng *accounts.Manager, backend bind.ContractBackend, address common.Address, passphrase string) (*XCData, error) {

	account := accounts.Account{Address: address}
	wallet, err := accMng.Find(account)
	if err != nil {
		return nil, err
	}

	transactOpts, err := wallet.NewKeyedTransactor(account, passphrase)
	if err != nil {
		return nil, err
	}

	contract, err := NewXCData(transactOpts, TestNetAddress, backend)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func GetXCDataReadOnly(backend bind.ContractBackend) (*XCData, error) {

	contract, err := NewXCData(&bind.TransactOpts{}, TestNetAddress, backend)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func (self *XCData) CommitData(did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return self.Contract.CommitData(&self.TransactOpts, did, datahash, encryptedAESKey)
}

func (self *XCData) CommitNewOwnerData(did string, datahash string, encryptedAESKey []byte) (*types.Transaction, error) {
	return self.Contract.CommitNewOwnerData(&self.TransactOpts, did, datahash ,encryptedAESKey);
}

func (self *XCData) DeletePreOwnerData(did string) (*types.Transaction, error) {
	return self.Contract.DeletePreOwnerData(&self.TransactOpts, did);
}

func (self *XCData) TransferDidOwner(did string, to common.Address) (*types.Transaction, error) {
	return self.Contract.TransferDidOwner(&self.TransactOpts, did, to);
}

func (self *XCData) AutherizeData(to common.Address, did string, index *big.Int, encryptedAESKey []byte) (*types.Transaction, error) {
	return self.Contract.AutherizeData(&self.TransactOpts, to, did, index ,encryptedAESKey);
}

func (self *XCData) GetDataLength(did string) (*big.Int, error) {
	return self.Contract.GetDataLength(&self.CallOpts, did);
}

func (self *XCData) GetData(did string, index *big.Int) (*big.Int, string, []byte, error) {
	return self.Contract.GetData(&self.CallOpts, did, index);
}

func (self *XCData) GetAutherizedDataLength(addr common.Address) (*big.Int, error) {
	return self.Contract.GetAutherizedDataLength(&self.CallOpts, addr);
}

func (self *XCData) GetAutherizedAESKeyByHash(addr common.Address, datahash string) ([]byte, error) {
	return self.Contract.GetAutherizedAESKeyByHash(&self.CallOpts, addr,datahash);
}