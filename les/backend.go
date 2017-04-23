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

// Package les implements the Light Elementrem Subprotocol.
package les

import (
	"errors"
	"fmt"
	"time"

	"github.com/elementrem/go-elementrem/accounts"
	"github.com/elementrem/go-elementrem/common"
	"github.com/elementrem/go-elementrem/common/compiler"
	"github.com/elementrem/go-elementrem/common/hexutil"
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/ele"
	"github.com/elementrem/go-elementrem/ele/downloader"
	"github.com/elementrem/go-elementrem/ele/filters"
	"github.com/elementrem/go-elementrem/ele/gasprice"
	"github.com/elementrem/go-elementrem/eledb"
	"github.com/elementrem/go-elementrem/event"
	"github.com/elementrem/go-elementrem/internal/eleapi"
	"github.com/elementrem/go-elementrem/light"
	"github.com/elementrem/go-elementrem/logger"
	"github.com/elementrem/go-elementrem/logger/glog"
	"github.com/elementrem/go-elementrem/node"
	"github.com/elementrem/go-elementrem/p2p"
	"github.com/elementrem/go-elementrem/params"
	"github.com/elementrem/go-elementrem/pow"
	rpc "github.com/elementrem/go-elementrem/rpc"
)

type LightElementrem struct {
	odr         *LesOdr
	relay       *LesTxRelay
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan chan bool
	// Handlers
	txPool          *light.TxPool
	blockchain      *light.LightChain
	protocolManager *ProtocolManager
	// DB interfaces
	chainDb eledb.Database // Block chain database

	ApiBackend *LesApiBackend

	eventMux       *event.TypeMux
	pow            pow.PoW
	accountManager *accounts.Manager
	solcPath       string
	solc           *compiler.Solidity

	netVersionId  int
	netRPCService *eleapi.PublicNetAPI
}

func New(ctx *node.ServiceContext, config *ele.Config) (*LightElementrem, error) {
	chainDb, err := ele.CreateDB(ctx, config, "lightchaindata")
	if err != nil {
		return nil, err
	}
	if err := ele.SetupGenesisBlock(&chainDb, config); err != nil {
		return nil, err
	}
	pow, err := ele.CreatePoW(config)
	if err != nil {
		return nil, err
	}

	odr := NewLesOdr(chainDb)
	relay := NewLesTxRelay()
	ele := &LightElementrem{
		odr:            odr,
		relay:          relay,
		chainDb:        chainDb,
		eventMux:       ctx.EventMux,
		accountManager: ctx.AccountManager,
		pow:            pow,
		shutdownChan:   make(chan bool),
		netVersionId:   config.NetworkId,
		solcPath:       config.SolcPath,
	}

	if config.ChainConfig == nil {
		return nil, errors.New("missing chain config")
	}
	ele.chainConfig = config.ChainConfig
	ele.blockchain, err = light.NewLightChain(odr, ele.chainConfig, ele.pow, ele.eventMux)
	if err != nil {
		if err == core.ErrNoGenesis {
			return nil, fmt.Errorf(`Genesis block not found. Please supply a genesis block with the "--genesis /path/to/file" argument`)
		}
		return nil, err
	}

	ele.txPool = light.NewTxPool(ele.chainConfig, ele.eventMux, ele.blockchain, ele.relay)
	if ele.protocolManager, err = NewProtocolManager(ele.chainConfig, config.LightMode, config.NetworkId, ele.eventMux, ele.pow, ele.blockchain, nil, chainDb, odr, relay); err != nil {
		return nil, err
	}

	ele.ApiBackend = &LesApiBackend{ele, nil}
	ele.ApiBackend.gpo = gasprice.NewLightPriceOracle(ele.ApiBackend)
	return ele, nil
}

type LightDummyAPI struct{}

// Elementbase is the address that mining rewards will be send to
func (s *LightDummyAPI) Elementbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Coinbase is the address that mining rewards will be send to (alias for Elementbase)
func (s *LightDummyAPI) Coinbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Hashrate returns the POW hashrate
func (s *LightDummyAPI) Hashrate() hexutil.Uint {
	return 0
}

// Mining returns an indication if this node is currently mining.
func (s *LightDummyAPI) Mining() bool {
	return false
}

// APIs returns the collection of RPC services the elementrem package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *LightElementrem) APIs() []rpc.API {
	return append(eleapi.GetAPIs(s.ApiBackend, s.solcPath), []rpc.API{
		{
			Namespace: "ele",
			Version:   "1.0",
			Service:   &LightDummyAPI{},
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, true),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *LightElementrem) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *LightElementrem) BlockChain() *light.LightChain      { return s.blockchain }
func (s *LightElementrem) TxPool() *light.TxPool              { return s.txPool }
func (s *LightElementrem) LesVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *LightElementrem) Downloader() *downloader.Downloader { return s.protocolManager.downloader }
func (s *LightElementrem) EventMux() *event.TypeMux           { return s.eventMux }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *LightElementrem) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

// Start implements node.Service, starting all internal goroutines needed by the
// Elementrem protocol implementation.
func (s *LightElementrem) Start(srvr *p2p.Server) error {
	glog.V(logger.Info).Infof("WARNING: light client mode is an experimental feature")
	s.netRPCService = eleapi.NewPublicNetAPI(srvr, s.netVersionId)
	s.protocolManager.Start(srvr)
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Elementrem protocol.
func (s *LightElementrem) Stop() error {
	s.odr.Stop()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()

	s.eventMux.Stop()

	time.Sleep(time.Millisecond * 200)
	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
