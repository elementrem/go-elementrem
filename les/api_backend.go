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

package les

import (
	"math/big"

	"github.com/elementrem/go-elementrem/accounts"
	"github.com/elementrem/go-elementrem/common"
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/core/vm"
	"github.com/elementrem/go-elementrem/ele/downloader"
	"github.com/elementrem/go-elementrem/ele/gasprice"
	"github.com/elementrem/go-elementrem/eledb"
	"github.com/elementrem/go-elementrem/event"
	"github.com/elementrem/go-elementrem/internal/eleapi"
	"github.com/elementrem/go-elementrem/light"
	"github.com/elementrem/go-elementrem/params"
	"github.com/elementrem/go-elementrem/rpc"
	"golang.org/x/net/context"
)

type LesApiBackend struct {
	ele *LightElementrem
	gpo *gasprice.LightPriceOracle
}

func (b *LesApiBackend) ChainConfig() *params.ChainConfig {
	return b.ele.chainConfig
}

func (b *LesApiBackend) CurrentBlock() *types.Block {
	return types.NewBlockWithHeader(b.ele.BlockChain().CurrentHeader())
}

func (b *LesApiBackend) SetHead(number uint64) {
	b.ele.blockchain.SetHead(number)
}

func (b *LesApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	if blockNr == rpc.LatestBlockNumber || blockNr == rpc.PendingBlockNumber {
		return b.ele.blockchain.CurrentHeader(), nil
	}

	return b.ele.blockchain.GetHeaderByNumberOdr(ctx, uint64(blockNr))
}

func (b *LesApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, err
	}
	return b.GetBlock(ctx, header.Hash())
}

func (b *LesApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (eleapi.State, *types.Header, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	return light.NewLightState(light.StateTrieID(header), b.ele.odr), header, nil
}

func (b *LesApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.ele.blockchain.GetBlockByHash(ctx, blockHash)
}

func (b *LesApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return light.GetBlockReceipts(ctx, b.ele.odr, blockHash, core.GetBlockNumber(b.ele.chainDb, blockHash))
}

func (b *LesApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.ele.blockchain.GetTdByHash(blockHash)
}

func (b *LesApiBackend) GetVMEnv(ctx context.Context, msg core.Message, state eleapi.State, header *types.Header) (*vm.EVM, func() error, error) {
	stateDb := state.(*light.LightState).Copy()
	addr := msg.From()
	from, err := stateDb.GetOrNewStateObject(ctx, addr)
	if err != nil {
		return nil, nil, err
	}
	from.SetBalance(common.MaxBig)

	vmstate := light.NewVMState(ctx, stateDb)
	context := core.NewEVMContext(msg, header, b.ele.blockchain)
	return vm.NewEVM(context, vmstate, b.ele.chainConfig, vm.Config{}), vmstate.Error, nil
}

func (b *LesApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.ele.txPool.Add(ctx, signedTx)
}

func (b *LesApiBackend) RemoveTx(txHash common.Hash) {
	b.ele.txPool.RemoveTx(txHash)
}

func (b *LesApiBackend) GetPoolTransactions() (types.Transactions, error) {
	return b.ele.txPool.GetTransactions()
}

func (b *LesApiBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	return b.ele.txPool.GetTransaction(txHash)
}

func (b *LesApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.ele.txPool.GetNonce(ctx, addr)
}

func (b *LesApiBackend) Stats() (pending int, queued int) {
	return b.ele.txPool.Stats(), 0
}

func (b *LesApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.ele.txPool.Content()
}

func (b *LesApiBackend) Downloader() *downloader.Downloader {
	return b.ele.Downloader()
}

func (b *LesApiBackend) ProtocolVersion() int {
	return b.ele.LesVersion() + 10000
}

func (b *LesApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *LesApiBackend) ChainDb() eledb.Database {
	return b.ele.chainDb
}

func (b *LesApiBackend) EventMux() *event.TypeMux {
	return b.ele.eventMux
}

func (b *LesApiBackend) AccountManager() *accounts.Manager {
	return b.ele.accountManager
}
