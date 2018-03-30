// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"context"
	"math/big"

	"github.com/xcareteam/xci/accounts"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/common/math"
	"github.com/xcareteam/xci/core"
	"github.com/xcareteam/xci/core/bloombits"
	"github.com/xcareteam/xci/core/state"
	"github.com/xcareteam/xci/core/types"
	"github.com/xcareteam/xci/core/vm"
	"github.com/xcareteam/xci/eth/downloader"
	"github.com/xcareteam/xci/eth/gasprice"
	"github.com/xcareteam/xci/ethdb"
	"github.com/xcareteam/xci/event"
	"github.com/xcareteam/xci/params"
	"github.com/xcareteam/xci/rpc"
	"github.com/ipfs/go-ipfs-api"
	"bytes"
	"github.com/xcareteam/xci/contracts/xcdata"
	"fmt"
	"github.com/xcareteam/xci/common/randomGenerator"
	"github.com/xcareteam/xci/common/AES"
	"github.com/xcareteam/xci/crypto"
	"github.com/xcareteam/xci/crypto/ecies"
	"crypto/rand"
	"github.com/pkg/errors"
	"strings"
	"strconv"
	"github.com/ipfs/ipfs-cluster/api/rest/client"
	"github.com/ipfs/go-cid"
)

// EthApiBackend implements ethapi.Backend for full nodes
type EthApiBackend struct {
	eth *Ethereum
	gpo *gasprice.Oracle
}

func (b *EthApiBackend) ChainConfig() *params.ChainConfig {
	return b.eth.chainConfig
}

func (b *EthApiBackend) CurrentBlock() *types.Block {
	return b.eth.blockchain.CurrentBlock()
}

func (b *EthApiBackend) SetHead(number uint64) {
	b.eth.protocolManager.downloader.Cancel()
	b.eth.blockchain.SetHead(number)
}

func (b *EthApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.eth.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.eth.blockchain.CurrentBlock().Header(), nil
	}
	return b.eth.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *EthApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.eth.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.eth.blockchain.CurrentBlock(), nil
	}
	return b.eth.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *EthApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.eth.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.eth.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *EthApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.eth.blockchain.GetBlockByHash(blockHash), nil
}

func (b *EthApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.eth.chainDb, blockHash, core.GetBlockNumber(b.eth.chainDb, blockHash)), nil
}

func (b *EthApiBackend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	receipts := core.GetBlockReceipts(b.eth.chainDb, blockHash, core.GetBlockNumber(b.eth.chainDb, blockHash))
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *EthApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.eth.blockchain.GetTdByHash(blockHash)
}

func (b *EthApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.eth.BlockChain(), nil)
	return vm.NewEVM(context, state, b.eth.chainConfig, vmCfg), vmError, nil
}

func (b *EthApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *EthApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChainEvent(ch)
}

func (b *EthApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *EthApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *EthApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.eth.BlockChain().SubscribeLogsEvent(ch)
}

func (b *EthApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.eth.txPool.AddLocal(signedTx)
}

func (b *EthApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.eth.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *EthApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.eth.txPool.Get(hash)
}

func (b *EthApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.eth.txPool.State().GetNonce(addr), nil
}

func (b *EthApiBackend) Stats() (pending int, queued int) {
	return b.eth.txPool.Stats()
}

func (b *EthApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.eth.TxPool().Content()
}

func (b *EthApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.eth.TxPool().SubscribeTxPreEvent(ch)
}

func (b *EthApiBackend) CommitXciData(address common.Address, passphrase string, ipfsEndpoint string, did string, data []byte) (common.Hash, error) {

	account := accounts.Account{Address: address}

	wallet,err :=b.eth.accountManager.Find(account)
	if err != nil {
		return common.Hash{}, err
	}

	//AES key generate
	AESKey,err := randomGenerator.GenerateRandomString(32)
	if err != nil {
		return common.Hash{}, err
	}

	//AES encryption
	encryptedData,err := AES.Encrypt([]byte(AESKey),data)
	if err != nil {
		return common.Hash{}, err
	}

	ipfsInfo := strings.Split(ipfsEndpoint, ":")
	if len(ipfsInfo)!=5 {
		return common.Hash{}, errors.New("ipfsEndpoint syntax error, expected syntax: localhost:5001:9094:2:3, Here we suppose ipfs and ipfs-cluter are on the same host, 2 is replicationFactorMin and 3 is replicationFactorMax")
	}
	ipfsHost := ipfsInfo[0]
	ipfsPort := ipfsInfo[1]
	ipfsClusterPort := ipfsInfo[2]

	replicationFactorMin,err := strconv.Atoi(ipfsInfo[3])
	if err != nil {
		return common.Hash{}, err
	}
	replicationFactorMax,err := strconv.Atoi(ipfsInfo[4])
	if err != nil {
		return common.Hash{}, err
	}
	//Open ipfs shell
	ipfsShell := shell.NewShell(ipfsHost+":"+ipfsPort)
	//Save AES encrypted data to IPFS
	mhash, err := ipfsShell.Add(bytes.NewReader([]byte(encryptedData)))
	if err != nil{
		return common.Hash{}, err
	}

	cfg := &client.Config{
		Host:              ipfsHost,
		Port:              ipfsClusterPort,
		DisableKeepAlives: false,
	}

	ipfsClusterClient, err := client.NewClient(cfg)
	if err != nil {
		return common.Hash{}, err
	}

	ci,err := cid.Decode(mhash)
	if err != nil {
		return common.Hash{}, err
	}

	err = ipfsClusterClient.Pin(ci,replicationFactorMin,replicationFactorMax,"")
	if err != nil {
		return common.Hash{}, err
	}

	encryptiedAESKey,err := wallet.EncryptDataWithPublicKey(account,passphrase,[]byte(AESKey))

	if err != nil {
		return common.Hash{}, err
	}

	xcData,err := xcdata.GetXCData(b.eth.accountManager, NewContractBackend(b.eth.ApiBackend), address, passphrase)

	if err != nil {
		return common.Hash{}, err
	}
	//Save ipfs hash and AES to blockchain
	tx, err := xcData.CommitData(did, mhash,encryptiedAESKey)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (b *EthApiBackend) CommitNewOwnerData(address common.Address, passphrase string, ipfsEndpoint string, did string, data []byte) (common.Hash, error) {

	account := accounts.Account{Address: address}

	wallet,err :=b.eth.accountManager.Find(account)
	if err != nil {
		return common.Hash{}, err
	}

	//AES key generate
	AESKey,err := randomGenerator.GenerateRandomString(32)
	if err != nil {
		return common.Hash{}, err
	}

	//AES encryption
	encryptedData,err := AES.Encrypt([]byte(AESKey),data)
	if err != nil {
		return common.Hash{}, err
	}

	ipfsInfo := strings.Split(ipfsEndpoint, ":")
	if len(ipfsInfo)!=5 {
		return common.Hash{}, errors.New("ipfsEndpoint syntax error, expected syntax: localhost:5001:9094:2:3, Here we suppose ipfs and ipfs-cluter are on the same host, 2 is replicationFactorMin and 3 is replicationFactorMax")
	}
	ipfsHost := ipfsInfo[0]
	ipfsPort := ipfsInfo[1]
	ipfsClusterPort := ipfsInfo[2]

	replicationFactorMin,err := strconv.Atoi(ipfsInfo[3])
	if err != nil {
		return common.Hash{}, err
	}
	replicationFactorMax,err := strconv.Atoi(ipfsInfo[4])
	if err != nil {
		return common.Hash{}, err
	}
	//Open ipfs shell
	ipfsShell := shell.NewShell(ipfsHost+":"+ipfsPort)
	//Save AES encrypted data to IPFS
	mhash, err := ipfsShell.Add(bytes.NewReader([]byte(encryptedData)))
	if err != nil{
		return common.Hash{}, err
	}

	cfg := &client.Config{
		Host:              ipfsHost,
		Port:              ipfsClusterPort,
		DisableKeepAlives: false,
	}

	ipfsClusterClient, err := client.NewClient(cfg)
	if err != nil {
		return common.Hash{}, err
	}

	ci,err := cid.Decode(mhash)
	if err != nil {
		return common.Hash{}, err
	}

	err = ipfsClusterClient.Pin(ci,replicationFactorMin,replicationFactorMax,"")
	if err != nil {
		return common.Hash{}, err
	}

	encryptiedAESKey,err := wallet.EncryptDataWithPublicKey(account,passphrase,[]byte(AESKey))
	if err != nil {
		return common.Hash{}, err
	}

	xcData,err := xcdata.GetXCData(b.eth.accountManager, NewContractBackend(b.eth.ApiBackend), address, passphrase)
	if err != nil {
		return common.Hash{}, err
	}

	//Save ipfs hash and AES to blockchain
	tx, err := xcData.CommitNewOwnerData(did, mhash,encryptiedAESKey)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (b *EthApiBackend) DeletePreOwnerData(address common.Address, passphrase string, did string) (common.Hash, error) {

	xcData,err := xcdata.GetXCData(b.eth.accountManager, NewContractBackend(b.eth.ApiBackend), address, passphrase)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := xcData.DeletePreOwnerData(did)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (b *EthApiBackend) TransferDidOwner(address common.Address, passphrase string, did string, to common.Address) (common.Hash, error) {

	xcData,err := xcdata.GetXCData(b.eth.accountManager, NewContractBackend(b.eth.ApiBackend), address, passphrase)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := xcData.TransferDidOwner(did,to)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (b *EthApiBackend) AuthorizeXcdata(address common.Address, passphrase string, publicKeyString string, did string, index *big.Int) (common.Hash, error) {

	publicKeyByte := common.FromHex(publicKeyString)
	if publicKeyByte == nil {
		return common.Hash{}, errors.New("Error: publicKey string is not hexadecimal string")
	}

	publicKey := crypto.ToECDSAPub(publicKeyByte)

	if publicKey.X == nil || publicKey.Y == nil {
		return common.Hash{}, errors.New("Error: invalid public key")
	}

	eciesPublic := ecies.ImportECDSAPublic(publicKey)

	toAddress := crypto.PubkeyToAddress(*publicKey)

	xcData,err := xcdata.GetXCData(b.eth.accountManager, NewContractBackend(b.eth.ApiBackend), address, passphrase)
	if err != nil {
		return common.Hash{}, err
	}

	_, _, encryptedAESKey, err := xcData.GetData(did,index)
	if err != nil {
		return common.Hash{}, err
	}

	account := accounts.Account{Address: address}

	wallet,err :=b.eth.accountManager.Find(account)
	if err != nil {
		return common.Hash{}, err
	}
	decryptedAESKey,err := wallet.DecryptDataWithPrivateKey(account,passphrase,encryptedAESKey)
	if err != nil {
		return common.Hash{}, err
	}

	reencryptedAESKey,err := ecies.Encrypt(rand.Reader, eciesPublic, decryptedAESKey, nil, nil)
	if err != nil {
		return common.Hash{}, err
	}

	//Save ipfs hash and AES to blockchain
	tx, err := xcData.AutherizeData(toAddress, did, index, reencryptedAESKey)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(),nil
}

func (b *EthApiBackend) GetXciDataLength(did string) (*big.Int, error) {

	xcData,err := xcdata.GetXCDataReadOnly(NewContractBackend(b.eth.ApiBackend))
	if err != nil {
		return nil, err
	}

	length, err := xcData.GetDataLength(did)
	if err != nil {
		return nil, err
	}

	return length, nil
}

func (b *EthApiBackend) GetXciData(address common.Address, passphrase string, ipfsEndpoint string, did string, index *big.Int) ([]byte, error) {

	_, ipfsHash, encryptedAESKey, err := b.GetXciDataTimestampAndHash(did,index)
	if err != nil{
		return nil, err
	}

	url := fmt.Sprintf("/ipfs/%s",ipfsHash)
	ipfsShell := shell.NewShell(ipfsEndpoint)

	rc, err := ipfsShell.Cat(url)
	if err != nil{
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)

	account := accounts.Account{Address: address}

	wallet,err :=b.eth.accountManager.Find(account)
	if err != nil {
		return nil, err
	}

	AESKey,err := wallet.DecryptDataWithPrivateKey(account,passphrase,encryptedAESKey)
	if err != nil {
		return nil, err
	}

	//AES encryption
	decryptedData,err := AES.Decrypt(AESKey,buf.Bytes())
	if err != nil {
		return nil, err
	}

	return decryptedData,nil
}

func (b *EthApiBackend) GetXciDataTimestampAndHash(did string, index *big.Int) (*big.Int, string, []byte, error) {

	xcData,err := xcdata.GetXCDataReadOnly(NewContractBackend(b.eth.ApiBackend))
	if err != nil {
		return nil, "", nil, err
	}

	timestamp, ipfsHash, encryptedAESKey, err := xcData.GetData(did,index)
	if err != nil {
		return nil, "", nil, err
	}

	return timestamp, ipfsHash, encryptedAESKey, nil
}

func (b *EthApiBackend) GetAutherizedDataLength(address common.Address) (*big.Int, error) {

	xcData,err := xcdata.GetXCDataReadOnly(NewContractBackend(b.eth.ApiBackend))
	if err != nil {
		return nil, err
	}

	length, err := xcData.GetAutherizedDataLength(address)
	if err != nil {
		return nil, err
	}

	return length, nil
}

func (b *EthApiBackend) GetAutherizedAESKeyByHash(address common.Address, hash string) ([]byte, error) {

	xcData,err := xcdata.GetXCDataReadOnly(NewContractBackend(b.eth.ApiBackend))
	if err != nil {
		return nil, err
	}

	autherizedAESKey, err := xcData.GetAutherizedAESKeyByHash(address,hash)
	if err != nil {
		return nil, err
	}

	return autherizedAESKey, nil
}

func (b *EthApiBackend) GetAutherizedData(address common.Address, passphrase string, ipfsEndpoint string, ipfsHash string) ([]byte, error) {

	xcData,err := xcdata.GetXCDataReadOnly(NewContractBackend(b.eth.ApiBackend))
	if err != nil {
		return nil, err
	}

	autherizedAESKey, err := xcData.GetAutherizedAESKeyByHash(address,ipfsHash)
	if err != nil {
		return nil, err
	}

	account := accounts.Account{Address: address}
	wallet,err :=b.eth.accountManager.Find(account)
	if err != nil{
		return nil, err
	}

	AESKey,err := wallet.DecryptDataWithPrivateKey(account,passphrase,autherizedAESKey)
	if err != nil{
		return nil, err
	}

	url := fmt.Sprintf("/ipfs/%s",ipfsHash)
	ipfsShell := shell.NewShell(ipfsEndpoint)

	rc, err := ipfsShell.Cat(url)
	if err != nil{
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)

	autherizedData,err := AES.Decrypt(AESKey,buf.Bytes())
	if err != nil{
		return nil, err
	}

	return autherizedData, nil
}

func (b *EthApiBackend) Downloader() *downloader.Downloader {
	return b.eth.Downloader()
}

func (b *EthApiBackend) ProtocolVersion() int {
	return b.eth.EthVersion()
}

func (b *EthApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *EthApiBackend) ChainDb() ethdb.Database {
	return b.eth.ChainDb()
}

func (b *EthApiBackend) EventMux() *event.TypeMux {
	return b.eth.EventMux()
}

func (b *EthApiBackend) AccountManager() *accounts.Manager {
	return b.eth.AccountManager()
}

func (b *EthApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.eth.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *EthApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.eth.bloomRequests)
	}
}
