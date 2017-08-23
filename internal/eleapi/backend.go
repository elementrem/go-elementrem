// Copyright 2016-2017 The go-elementrem Authors
// This file is part of the go-elementrem library.
//
// The go-elementrem library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-elementrem library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-elementrem library. If not, see <http://www.gnu.org/licenses/>.

// Package eleapi implements the general Elementrem API functions.
package eleapi

import (
	"math/big"

	"github.com/elementrem/go-elementrem/accounts"
	"github.com/elementrem/go-elementrem/common"
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/core/vm"
	"github.com/elementrem/go-elementrem/ele/downloader"
	"github.com/elementrem/go-elementrem/eledb"
	"github.com/elementrem/go-elementrem/event"
	"github.com/elementrem/go-elementrem/params"
	"github.com/elementrem/go-elementrem/rpc"
	"golang.org/x/net/context"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// general Elementrem API
	Downloader() *downloader.Downloader
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() eledb.Database
	EventMux() *event.TypeMux
	AccountManager() *accounts.Manager
	// BlockChain API
	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (State, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetTd(blockHash common.Hash) *big.Int
	GetVMEnv(ctx context.Context, msg core.Message, state State, header *types.Header) (*vm.EVM, func() error, error)
	// TxPool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	RemoveTx(txHash common.Hash)
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)

	ChainConfig() *params.ChainConfig
	CurrentBlock() *types.Block
}

type State interface {
	GetBalance(ctx context.Context, addr common.Address) (*big.Int, error)
	GetCode(ctx context.Context, addr common.Address) ([]byte, error)
	GetState(ctx context.Context, a common.Address, b common.Hash) (common.Hash, error)
	GetNonce(ctx context.Context, addr common.Address) (uint64, error)
}

func GetAPIs(apiBackend Backend, solcPath string) []rpc.API {
	compiler := makeCompilerAPIs(solcPath)
	all := []rpc.API{
		{
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicElementremAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend),
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
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend),
			Public:    false,
		},
	}
	return append(compiler, all...)
}
