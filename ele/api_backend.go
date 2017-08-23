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

package ele

import (
	"math/big"

	"github.com/elementrem/go-elementrem/accounts"
	"github.com/elementrem/go-elementrem/common"
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/core/state"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/core/vm"
	"github.com/elementrem/go-elementrem/ele/downloader"
	"github.com/elementrem/go-elementrem/ele/gasprice"
	"github.com/elementrem/go-elementrem/eledb"
	"github.com/elementrem/go-elementrem/event"
	"github.com/elementrem/go-elementrem/internal/eleapi"
	"github.com/elementrem/go-elementrem/params"
	"github.com/elementrem/go-elementrem/rpc"
	"golang.org/x/net/context"
)

// EleApiBackend implements eleapi.Backend for full nodes
type EleApiBackend struct {
	ele *Elementrem
	gpo *gasprice.GasPriceOracle
}

func (b *EleApiBackend) ChainConfig() *params.ChainConfig {
	return b.ele.chainConfig
}

func (b *EleApiBackend) CurrentBlock() *types.Block {
	return b.ele.blockchain.CurrentBlock()
}

func (b *EleApiBackend) SetHead(number uint64) {
	b.ele.blockchain.SetHead(number)
}

func (b *EleApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.ele.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.ele.blockchain.CurrentBlock().Header(), nil
	}
	return b.ele.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *EleApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.ele.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.ele.blockchain.CurrentBlock(), nil
	}
	return b.ele.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *EleApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (eleapi.State, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.ele.miner.Pending()
		return EleApiState{state}, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.ele.BlockChain().StateAt(header.Root)
	return EleApiState{stateDb}, header, err
}

func (b *EleApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.ele.blockchain.GetBlockByHash(blockHash), nil
}

func (b *EleApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.ele.chainDb, blockHash, core.GetBlockNumber(b.ele.chainDb, blockHash)), nil
}

func (b *EleApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.ele.blockchain.GetTdByHash(blockHash)
}

func (b *EleApiBackend) GetVMEnv(ctx context.Context, msg core.Message, state eleapi.State, header *types.Header) (*vm.EVM, func() error, error) {
	statedb := state.(EleApiState).state
	from := statedb.GetOrNewStateObject(msg.From())
	from.SetBalance(common.MaxBig)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.ele.BlockChain())
	return vm.NewEVM(context, statedb, b.ele.chainConfig, vm.Config{}), vmError, nil
}

func (b *EleApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	b.ele.txPool.SetLocal(signedTx)
	return b.ele.txPool.Add(signedTx)
}

func (b *EleApiBackend) RemoveTx(txHash common.Hash) {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	b.ele.txPool.Remove(txHash)
}

func (b *EleApiBackend) GetPoolTransactions() (types.Transactions, error) {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	pending, err := b.ele.txPool.Pending()
	if err != nil {
		return nil, err
	}

	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *EleApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	return b.ele.txPool.Get(hash)
}

func (b *EleApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	return b.ele.txPool.State().GetNonce(addr), nil
}

func (b *EleApiBackend) Stats() (pending int, queued int) {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	return b.ele.txPool.Stats()
}

func (b *EleApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	b.ele.txMu.Lock()
	defer b.ele.txMu.Unlock()

	return b.ele.TxPool().Content()
}

func (b *EleApiBackend) Downloader() *downloader.Downloader {
	return b.ele.Downloader()
}

func (b *EleApiBackend) ProtocolVersion() int {
	return b.ele.EleVersion()
}

func (b *EleApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(), nil
}

func (b *EleApiBackend) ChainDb() eledb.Database {
	return b.ele.ChainDb()
}

func (b *EleApiBackend) EventMux() *event.TypeMux {
	return b.ele.EventMux()
}

func (b *EleApiBackend) AccountManager() *accounts.Manager {
	return b.ele.AccountManager()
}

type EleApiState struct {
	state *state.StateDB
}

func (s EleApiState) GetBalance(ctx context.Context, addr common.Address) (*big.Int, error) {
	return s.state.GetBalance(addr), nil
}

func (s EleApiState) GetCode(ctx context.Context, addr common.Address) ([]byte, error) {
	return s.state.GetCode(addr), nil
}

func (s EleApiState) GetState(ctx context.Context, a common.Address, b common.Hash) (common.Hash, error) {
	return s.state.GetState(a, b), nil
}

func (s EleApiState) GetNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return s.state.GetNonce(addr), nil
}
