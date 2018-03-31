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

// Package ethapi implements the general Ethereum API functions.
package ethapi

import (
	"context"
	"math/big"

	"github.com/xcareteam/xci/accounts"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/core"
	"github.com/xcareteam/xci/core/state"
	"github.com/xcareteam/xci/core/types"
	"github.com/xcareteam/xci/core/vm"
	"github.com/xcareteam/xci/eth/downloader"
	"github.com/xcareteam/xci/ethdb"
	"github.com/xcareteam/xci/event"
	"github.com/xcareteam/xci/params"
	"github.com/xcareteam/xci/rpc"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Ethereum API
	Downloader() *downloader.Downloader
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() ethdb.Database
	EventMux() *event.TypeMux
	AccountManager() *accounts.Manager

	// BlockChain API
	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetTd(blockHash common.Hash) *big.Int
	GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error)
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription

	// TxPool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeTxPreEvent(chan<- core.TxPreEvent) event.Subscription

	//XCARE API
	CommitXciData(address common.Address, passphrase string, ipfsEndpoint string, did string, data []byte) (common.Hash, error)
	CommitNewOwnerData(address common.Address, passphrase string, ipfsEndpoint string, did string, data []byte) (common.Hash, error)
	DeletePreOwnerData(address common.Address, passphrase string, did string) (common.Hash, error)
	TransferDidOwner(address common.Address, passphrase string, did string, to common.Address) (common.Hash, error)
	AuthorizeXcdata(address common.Address, passphrase string, publicKeyString string, did string, index *big.Int) (common.Hash, error)
	GetXciDataLength(did string) (*big.Int, error)
	GetXciData(address common.Address, passphrase string, ipfsEndpoint string, did string, index *big.Int) ([]byte, error)
	GetXciDataTimestampAndHash(did string, index *big.Int) (*big.Int, string, []byte, error)
	GetAutherizedDataLength(address common.Address) (*big.Int, error)
	GetAutherizedAESKeyByHash(address common.Address, hash string) ([]byte, error)
	GetAutherizedData(address common.Address, passphrase string, ipfsEndpoint string, ipfsHash string) ([]byte, error)

	ChainConfig() *params.ChainConfig
	CurrentBlock() *types.Block
}

func GetAPIs(apiBackend Backend) []rpc.API {
	nonceLock := new(AddrLocker)
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(apiBackend),
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend, nonceLock),
			Public:    false,
		}, {
			Namespace: "xcare",
			Version:   "1.0",
			Service:   NewPublicXcareAPI(apiBackend),
			Public:    true,
		},
	}
}
