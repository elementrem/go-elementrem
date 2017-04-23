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

// Package ele implements the Elementrem protocol.
package ele

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/elementrem/elhash"
	"github.com/elementrem/go-elementrem/accounts"
	"github.com/elementrem/go-elementrem/common"
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/core/vm"
	"github.com/elementrem/go-elementrem/ele/downloader"
	"github.com/elementrem/go-elementrem/ele/filters"
	"github.com/elementrem/go-elementrem/ele/gasprice"
	"github.com/elementrem/go-elementrem/eledb"
	"github.com/elementrem/go-elementrem/event"
	"github.com/elementrem/go-elementrem/internal/eleapi"
	"github.com/elementrem/go-elementrem/logger"
	"github.com/elementrem/go-elementrem/logger/glog"
	"github.com/elementrem/go-elementrem/miner"
	"github.com/elementrem/go-elementrem/node"
	"github.com/elementrem/go-elementrem/p2p"
	"github.com/elementrem/go-elementrem/params"
	"github.com/elementrem/go-elementrem/pow"
	"github.com/elementrem/go-elementrem/rpc"
)

const (
	epochLength    = 30000
	elhashRevision = 23

	autoDAGcheckInterval = 10 * time.Hour
	autoDAGepochHeight   = epochLength / 2
)

var (
	datadirInUseErrnos = map[uint]bool{11: true, 32: true, 35: true}
	portInUseErrRE     = regexp.MustCompile("address already in use")
)

type Config struct {
	ChainConfig *params.ChainConfig // chain configuration

	NetworkId  int    // Network ID to use for selecting peers to connect to
	Genesis    string // Genesis JSON to seed the chain database with
	FastSync   bool   // Enables the state download based fast synchronisation algorithm
	LightMode  bool   // Running in light client mode
	LightServ  int    // Maximum percentage of time allowed for serving LES requests
	LightPeers int    // Maximum number of LES client peers
	MaxPeers   int    // Maximum number of global peers

	SkipBcVersionCheck bool // e.g. blockchain export
	DatabaseCache      int
	DatabaseHandles    int

	DocRoot   string
	AutoDAG   bool
	PowFake   bool
	PowTest   bool
	PowShared bool
	ExtraData []byte

	Elementbase    common.Address
	GasPrice     *big.Int
	MinerThreads int
	SolcPath     string

	GpoMinGasPrice          *big.Int
	GpoMaxGasPrice          *big.Int
	GpoFullBlockRatio       int
	GpobaseStepDown         int
	GpobaseStepUp           int
	GpobaseCorrectionFactor int

	EnablePreimageRecording bool

	TestGenesisBlock *types.Block   // Genesis block to seed the chain database with (testing only!)
	TestGenesisState eledb.Database // Genesis state to seed the database with (testing only!)
}

type LesServer interface {
	Start(srvr *p2p.Server)
	Stop()
	Protocols() []p2p.Protocol
}

// Elementrem implements the Elementrem full node service.
type Elementrem struct {
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan  chan bool // Channel for shutting down the elementrem
	stopDbUpgrade func()    // stop chain db sequential key upgrade
	// Handlers
	txPool          *core.TxPool
	txMu            sync.Mutex
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	lesServer       LesServer
	// DB interfaces
	chainDb eledb.Database // Block chain database

	eventMux       *event.TypeMux
	pow            pow.PoW
	accountManager *accounts.Manager

	ApiBackend *EleApiBackend

	miner        *miner.Miner
	Mining       bool
	MinerThreads int
	AutoDAG      bool
	autodagquit  chan bool
	elementbase    common.Address
	solcPath     string

	netVersionId  int
	netRPCService *eleapi.PublicNetAPI
}

func (s *Elementrem) AddLesServer(ls LesServer) {
	s.lesServer = ls
	s.protocolManager.lesServer = ls
}

// New creates a new Elementrem object (including the
// initialisation of the common Elementrem object)
func New(ctx *node.ServiceContext, config *Config) (*Elementrem, error) {
	chainDb, err := CreateDB(ctx, config, "chaindata")
	if err != nil {
		return nil, err
	}
	stopDbUpgrade := upgradeSequentialKeys(chainDb)
	if err := SetupGenesisBlock(&chainDb, config); err != nil {
		return nil, err
	}
	pow, err := CreatePoW(config)
	if err != nil {
		return nil, err
	}

	ele := &Elementrem{
		chainDb:        chainDb,
		eventMux:       ctx.EventMux,
		accountManager: ctx.AccountManager,
		pow:            pow,
		shutdownChan:   make(chan bool),
		stopDbUpgrade:  stopDbUpgrade,
		netVersionId:   config.NetworkId,
		elementbase:      config.Elementbase,
		MinerThreads:   config.MinerThreads,
		AutoDAG:        config.AutoDAG,
		solcPath:       config.SolcPath,
	}

	if err := upgradeChainDatabase(chainDb); err != nil {
		return nil, err
	}
	if err := addMipmapBloomBins(chainDb); err != nil {
		return nil, err
	}

	glog.V(logger.Info).Infof("Protocol Versions: %v, Network Id: %v", ProtocolVersions, config.NetworkId)

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
		if bcVersion != core.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run gele upgradedb.\n", bcVersion, core.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
	}

	// load the genesis block or write a new one if no genesis
	// block is prenent in the database.
	genesis := core.GetBlock(chainDb, core.GetCanonicalHash(chainDb, 0), 0)
	if genesis == nil {
		genesis, err = core.WriteDefaultGenesisBlock(chainDb)
		if err != nil {
			return nil, err
		}
		glog.V(logger.Info).Infoln("Default genesis block is applied to The elementrem")
	}

	if config.ChainConfig == nil {
		return nil, errors.New("missing chain config")
	}
	core.WriteChainConfig(chainDb, genesis.Hash(), config.ChainConfig)

	ele.chainConfig = config.ChainConfig

	glog.V(logger.Info).Infoln("Chain config:", ele.chainConfig)

	ele.blockchain, err = core.NewBlockChain(chainDb, ele.chainConfig, ele.pow, ele.EventMux(), vm.Config{EnablePreimageRecording: config.EnablePreimageRecording})
	if err != nil {
		if err == core.ErrNoGenesis {
			return nil, fmt.Errorf(`No chain found. Please initialise a new chain using the "init" subcommand.`)
		}
		return nil, err
	}
	newPool := core.NewTxPool(ele.chainConfig, ele.EventMux(), ele.blockchain.State, ele.blockchain.GasLimit)
	ele.txPool = newPool

	maxPeers := config.MaxPeers
	if config.LightServ > 0 {
		// if we are running a light server, limit the number of ELE peers so that we reserve some space for incoming LES connections
		// temporary solution until the new peer connectivity API is finished
		halfPeers := maxPeers / 2
		maxPeers -= config.LightPeers
		if maxPeers < halfPeers {
			maxPeers = halfPeers
		}
	}

	if ele.protocolManager, err = NewProtocolManager(ele.chainConfig, config.FastSync, config.NetworkId, maxPeers, ele.eventMux, ele.txPool, ele.pow, ele.blockchain, chainDb); err != nil {
		return nil, err
	}
	ele.miner = miner.New(ele, ele.chainConfig, ele.EventMux(), ele.pow)
	ele.miner.SetGasPrice(config.GasPrice)
	ele.miner.SetExtra(config.ExtraData)

	gpoParams := &gasprice.GpoParams{
		GpoMinGasPrice:          config.GpoMinGasPrice,
		GpoMaxGasPrice:          config.GpoMaxGasPrice,
		GpoFullBlockRatio:       config.GpoFullBlockRatio,
		GpobaseStepDown:         config.GpobaseStepDown,
		GpobaseStepUp:           config.GpobaseStepUp,
		GpobaseCorrectionFactor: config.GpobaseCorrectionFactor,
	}
	gpo := gasprice.NewGasPriceOracle(ele.blockchain, chainDb, ele.eventMux, gpoParams)
	ele.ApiBackend = &EleApiBackend{ele, gpo}

	return ele, nil
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (eledb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if db, ok := db.(*eledb.LDBDatabase); ok {
		db.Meter("ele/db/chaindata/")
	}
	return db, err
}

// SetupGenesisBlock initializes the genesis block for an Elementrem service
func SetupGenesisBlock(chainDb *eledb.Database, config *Config) error {
	// Load up any custom genesis block if requested
	if len(config.Genesis) > 0 {
		block, err := core.WriteGenesisBlock(*chainDb, strings.NewReader(config.Genesis))
		if err != nil {
			return err
		}
		glog.V(logger.Info).Infof("Successfully wrote custom genesis block: %x", block.Hash())
	}
	// Load up a test setup if directly injected
	if config.TestGenesisState != nil {
		*chainDb = config.TestGenesisState
	}
	if config.TestGenesisBlock != nil {
		core.WriteTd(*chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.NumberU64(), config.TestGenesisBlock.Difficulty())
		core.WriteBlock(*chainDb, config.TestGenesisBlock)
		core.WriteCanonicalHash(*chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.NumberU64())
		core.WriteHeadBlockHash(*chainDb, config.TestGenesisBlock.Hash())
	}
	return nil
}

// CreatePoW creates the required type of PoW instance for an Elementrem service
func CreatePoW(config *Config) (pow.PoW, error) {
	switch {
	case config.PowFake:
		glog.V(logger.Info).Infof("elhash used in fake mode")
		return pow.PoW(core.FakePow{}), nil
	case config.PowTest:
		glog.V(logger.Info).Infof("elhash used in test mode")
		return elhash.NewForTesting()
	case config.PowShared:
		glog.V(logger.Info).Infof("elhash used in shared mode")
		return elhash.NewShared(), nil
	default:
		return elhash.New(), nil
	}
}

// APIs returns the collection of RPC services the elementrem package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Elementrem) APIs() []rpc.API {
	return append(eleapi.GetAPIs(s.ApiBackend, s.solcPath), []rpc.API{
		{
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicElementremAPI(s),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicMinerAPI(s),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(s),
			Public:    false,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(s.chainConfig, s),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *Elementrem) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Elementrem) Elementbase() (eb common.Address, err error) {
	if s.elementbase != (common.Address{}) {
		return s.elementbase, nil
	}
	if wallets := s.AccountManager().Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			return accounts[0].Address, nil
		}
	}
	return common.Address{}, fmt.Errorf("elementbase address must be explicitly specified")
}

// set in js console via admin interface or wrapper from cli flags
func (self *Elementrem) SetElementbase(elementbase common.Address) {
	self.elementbase = elementbase
	self.miner.SetElementbase(elementbase)
}

func (s *Elementrem) StartMining(threads int) error {
	eb, err := s.Elementbase()
	if err != nil {
		err = fmt.Errorf("Cannot start mining without elementbase address: %v", err)
		glog.V(logger.Error).Infoln(err)
		return err
	}
	go s.miner.Start(eb, threads)
	return nil
}

func (s *Elementrem) StopMining()         { s.miner.Stop() }
func (s *Elementrem) IsMining() bool      { return s.miner.Mining() }
func (s *Elementrem) Miner() *miner.Miner { return s.miner }

func (s *Elementrem) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Elementrem) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Elementrem) TxPool() *core.TxPool               { return s.txPool }
func (s *Elementrem) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Elementrem) Pow() pow.PoW                       { return s.pow }
func (s *Elementrem) ChainDb() eledb.Database            { return s.chainDb }
func (s *Elementrem) IsListening() bool                  { return true } // Always listening
func (s *Elementrem) EleVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Elementrem) NetVersion() int                    { return s.netVersionId }
func (s *Elementrem) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *Elementrem) Protocols() []p2p.Protocol {
	if s.lesServer == nil {
		return s.protocolManager.SubProtocols
	} else {
		return append(s.protocolManager.SubProtocols, s.lesServer.Protocols()...)
	}
}

// Start implements node.Service, starting all internal goroutines needed by the
// Elementrem protocol implementation.
func (s *Elementrem) Start(srvr *p2p.Server) error {
	s.netRPCService = eleapi.NewPublicNetAPI(srvr, s.NetVersion())
	if s.AutoDAG {
		s.StartAutoDAG()
	}
	s.protocolManager.Start()
	if s.lesServer != nil {
		s.lesServer.Start(srvr)
	}
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Elementrem protocol.
func (s *Elementrem) Stop() error {
	if s.stopDbUpgrade != nil {
		s.stopDbUpgrade()
	}
	s.blockchain.Stop()
	s.protocolManager.Stop()
	if s.lesServer != nil {
		s.lesServer.Stop()
	}
	s.txPool.Stop()
	s.miner.Stop()
	s.eventMux.Stop()

	s.StopAutoDAG()

	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}

// This function will wait for a shutdown and resumes main thread execution
func (s *Elementrem) WaitForShutdown() {
	<-s.shutdownChan
}

// StartAutoDAG() spawns a go routine that checks the DAG every autoDAGcheckInterval
// by default that is 10 times per epoch
// in epoch n, if we past autoDAGepochHeight within-epoch blocks,
// it calls elhash.MakeDAG  to pregenerate the DAG for the next epoch n+1
// if it does not exist yet as well as remove the DAG for epoch n-1
// the loop quits if autodagquit channel is closed, it can safely restart and
// stop any number of times.
// For any more sophisticated pattern of DAG generation, use CLI subcommand
// makedag
func (self *Elementrem) StartAutoDAG() {
	if self.autodagquit != nil {
		return // already started
	}
	go func() {
		glog.V(logger.Info).Infof("Automatic pregeneration of elhash DAG ON (elhash dir: %s)", elhash.DefaultDir)
		var nextEpoch uint64
		timer := time.After(0)
		self.autodagquit = make(chan bool)
		for {
			select {
			case <-timer:
				glog.V(logger.Info).Infof("checking DAG (elhash dir: %s)", elhash.DefaultDir)
				currentBlock := self.BlockChain().CurrentBlock().NumberU64()
				thisEpoch := currentBlock / epochLength
				if nextEpoch <= thisEpoch {
					if currentBlock%epochLength > autoDAGepochHeight {
						if thisEpoch > 0 {
							previousDag, previousDagFull := dagFiles(thisEpoch - 1)
							os.Remove(filepath.Join(elhash.DefaultDir, previousDag))
							os.Remove(filepath.Join(elhash.DefaultDir, previousDagFull))
							glog.V(logger.Info).Infof("removed DAG for epoch %d (%s)", thisEpoch-1, previousDag)
						}
						nextEpoch = thisEpoch + 1
						dag, _ := dagFiles(nextEpoch)
						if _, err := os.Stat(dag); os.IsNotExist(err) {
							glog.V(logger.Info).Infof("Pregenerating DAG for epoch %d (%s)", nextEpoch, dag)
							err := elhash.MakeDAG(nextEpoch*epochLength, "") // "" -> elhash.DefaultDir
							if err != nil {
								glog.V(logger.Error).Infof("Error generating DAG for epoch %d (%s)", nextEpoch, dag)
								return
							}
						} else {
							glog.V(logger.Error).Infof("DAG for epoch %d (%s)", nextEpoch, dag)
						}
					}
				}
				timer = time.After(autoDAGcheckInterval)
			case <-self.autodagquit:
				return
			}
		}
	}()
}

// stopAutoDAG stops automatic DAG pregeneration by quitting the loop
func (self *Elementrem) StopAutoDAG() {
	if self.autodagquit != nil {
		close(self.autodagquit)
		self.autodagquit = nil
	}
	glog.V(logger.Info).Infof("Automatic pregeneration of elhash DAG OFF (elhash dir: %s)", elhash.DefaultDir)
}

// dagFiles(epoch) returns the two alternative DAG filenames (not a path)
// 1) <revision>-<hex(seedhash[8])> 2) full-R<revision>-<hex(seedhash[8])>
func dagFiles(epoch uint64) (string, string) {
	seedHash, _ := elhash.GetSeedHash(epoch * epochLength)
	dag := fmt.Sprintf("full-R%d-%x", elhashRevision, seedHash[:8])
	return dag, "full-R" + dag
}
