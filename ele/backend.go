// Copyright 2016 The go-elementrem Authors.
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
	"bytes"
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
	"github.com/elementrem/go-elementrem/common/compiler"
	"github.com/elementrem/go-elementrem/common/httpclient"
	"github.com/elementrem/go-elementrem/common/registrar/elereg"
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/core/vm"
	"github.com/elementrem/go-elementrem/ele/downloader"
	"github.com/elementrem/go-elementrem/ele/filters"
	"github.com/elementrem/go-elementrem/eledb"
	"github.com/elementrem/go-elementrem/event"
	"github.com/elementrem/go-elementrem/logger"
	"github.com/elementrem/go-elementrem/logger/glog"
	"github.com/elementrem/go-elementrem/miner"
	"github.com/elementrem/go-elementrem/node"
	"github.com/elementrem/go-elementrem/p2p"
	"github.com/elementrem/go-elementrem/rlp"
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
	ChainConfig *core.ChainConfig // chain configuration

	NetworkId int    // Network ID to use for selecting peers to connect to
	Genesis   string // Genesis JSON to seed the chain database with
	FastSync  bool   // Enables the state download based fast synchronisation algorithm

	BlockChainVersion  int
	SkipBcVersionCheck bool // e.g. blockchain export
	DatabaseCache      int
	DatabaseHandles    int

	NatSpec   bool
	DocRoot   string
	AutoDAG   bool
	PowTest   bool
	PowShared bool
	ExtraData []byte

	AccountManager *accounts.Manager
	Elementbase      common.Address
	GasPrice       *big.Int
	MinerThreads   int
	SolcPath       string

	GpoMinGasPrice          *big.Int
	GpoMaxGasPrice          *big.Int
	GpoFullBlockRatio       int
	GpobaseStepDown         int
	GpobaseStepUp           int
	GpobaseCorrectionFactor int

	EnableJit bool
	ForceJit  bool

	TestGenesisBlock *types.Block   // Genesis block to seed the chain database with (testing only!)
	TestGenesisState eledb.Database // Genesis state to seed the database with (testing only!)
}

type Elementrem struct {
	chainConfig *core.ChainConfig
	// Channel for shutting down the elementrem
	shutdownChan chan bool

	// DB interfaces
	chainDb eledb.Database // Block chain database
	dappDb  eledb.Database // Dapp database

	// Handlers
	txPool          *core.TxPool
	txMu            sync.Mutex
	blockchain      *core.BlockChain
	accountManager  *accounts.Manager
	pow             *elhash.Elhash
	protocolManager *ProtocolManager
	SolcPath        string
	solc            *compiler.Solidity
	gpo             *GasPriceOracle

	GpoMinGasPrice          *big.Int
	GpoMaxGasPrice          *big.Int
	GpoFullBlockRatio       int
	GpobaseStepDown         int
	GpobaseStepUp           int
	GpobaseCorrectionFactor int

	httpclient *httpclient.HTTPClient

	eventMux *event.TypeMux
	miner    *miner.Miner

	Mining        bool
	MinerThreads  int
	NatSpec       bool
	AutoDAG       bool
	PowTest       bool
	autodagquit   chan bool
	elementbase     common.Address
	netVersionId  int
	netRPCService *PublicNetAPI
}

func New(ctx *node.ServiceContext, config *Config) (*Elementrem, error) {
	// Open the chain database and perform any upgrades needed
	chainDb, err := ctx.OpenDatabase("chaindata", config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := chainDb.(*eledb.LDBDatabase); ok {
		db.Meter("ele/db/chaindata/")
	}
	if err := upgradeChainDatabase(chainDb); err != nil {
		return nil, err
	}
	if err := addMipmapBloomBins(chainDb); err != nil {
		return nil, err
	}

	dappDb, err := ctx.OpenDatabase("dapp", config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := dappDb.(*eledb.LDBDatabase); ok {
		db.Meter("ele/db/dapp/")
	}
	glog.V(logger.Info).Infof("Protocol Versions: %v, Network Id: %v", ProtocolVersions, config.NetworkId)

	// Load up any custom genesis block if requested
	if len(config.Genesis) > 0 {
		block, err := core.WriteGenesisBlock(chainDb, strings.NewReader(config.Genesis))
		if err != nil {
			return nil, err
		}
		glog.V(logger.Info).Infof("Successfully wrote custom genesis block: %x", block.Hash())
	}

	// Load up a test setup if directly injected
	if config.TestGenesisState != nil {
		chainDb = config.TestGenesisState
	}
	if config.TestGenesisBlock != nil {
		core.WriteTd(chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.Difficulty())
		core.WriteBlock(chainDb, config.TestGenesisBlock)
		core.WriteCanonicalHash(chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.NumberU64())
		core.WriteHeadBlockHash(chainDb, config.TestGenesisBlock.Hash())
	}

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
		if bcVersion != config.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run gele upgradedb.\n", bcVersion, config.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, config.BlockChainVersion)
	}
	glog.V(logger.Info).Infof("Blockchain DB Version: %d", config.BlockChainVersion)

	ele := &Elementrem{
		shutdownChan:            make(chan bool),
		chainDb:                 chainDb,
		dappDb:                  dappDb,
		eventMux:                ctx.EventMux,
		accountManager:          config.AccountManager,
		elementbase:               config.Elementbase,
		netVersionId:            config.NetworkId,
		NatSpec:                 config.NatSpec,
		MinerThreads:            config.MinerThreads,
		SolcPath:                config.SolcPath,
		AutoDAG:                 config.AutoDAG,
		PowTest:                 config.PowTest,
		GpoMinGasPrice:          config.GpoMinGasPrice,
		GpoMaxGasPrice:          config.GpoMaxGasPrice,
		GpoFullBlockRatio:       config.GpoFullBlockRatio,
		GpobaseStepDown:         config.GpobaseStepDown,
		GpobaseStepUp:           config.GpobaseStepUp,
		GpobaseCorrectionFactor: config.GpobaseCorrectionFactor,
		httpclient:              httpclient.New(config.DocRoot),
	}
	switch {
	case config.PowTest:
		glog.V(logger.Info).Infof("elhash used in test mode")
		ele.pow, err = elhash.NewForTesting()
		if err != nil {
			return nil, err
		}
	case config.PowShared:
		glog.V(logger.Info).Infof("elhash used in shared mode")
		ele.pow = elhash.NewShared()

	default:
		ele.pow = elhash.New()
	}

	// load the genesis block or write a new one if no genesis
	// block is prenent in the database.
	genesis := core.GetBlock(chainDb, core.GetCanonicalHash(chainDb, 0))
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
	ele.chainConfig.VmConfig = vm.Config{
		EnableJit: config.EnableJit,
		ForceJit:  config.ForceJit,
	}

	ele.blockchain, err = core.NewBlockChain(chainDb, ele.chainConfig, ele.pow, ele.EventMux())
	if err != nil {
		if err == core.ErrNoGenesis {
			return nil, fmt.Errorf(`No chain found. Please initialise a new chain using the "init" subcommand.`)
		}
		return nil, err
	}
	ele.gpo = NewGasPriceOracle(ele)

	newPool := core.NewTxPool(ele.chainConfig, ele.EventMux(), ele.blockchain.State, ele.blockchain.GasLimit)
	ele.txPool = newPool

	if ele.protocolManager, err = NewProtocolManager(ele.chainConfig, config.FastSync, config.NetworkId, ele.eventMux, ele.txPool, ele.pow, ele.blockchain, chainDb); err != nil {
		return nil, err
	}
	ele.miner = miner.New(ele, ele.chainConfig, ele.EventMux(), ele.pow)
	ele.miner.SetGasPrice(config.GasPrice)
	ele.miner.SetExtra(config.ExtraData)

	return ele, nil
}

// APIs returns the collection of RPC services the elementrem package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Elementrem) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicElementremAPI(s),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(s.accountManager),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(s),
			Public:    false,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(s.chainConfig, s.blockchain, s.miner, s.chainDb, s.gpo, s.eventMux, s.accountManager),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(s),
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
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(s),
			Public:    true,
		}, {
			Namespace: "ele",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.chainDb, s.eventMux),
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
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   elereg.NewPrivateRegistarAPI(s.chainConfig, s.blockchain, s.chainDb, s.txPool, s.accountManager),
		},
	}
}

func (s *Elementrem) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Elementrem) Elementbase() (eb common.Address, err error) {
	eb = s.elementbase
	if (eb == common.Address{}) {
		firstAccount, err := s.AccountManager().AccountByIndex(0)
		eb = firstAccount.Address
		if err != nil {
			return eb, fmt.Errorf("elementbase address must be explicitly specified")
		}
	}
	return eb, nil
}

// set in js console via admin interface or wrapper from cli flags
func (self *Elementrem) SetElementbase(elementbase common.Address) {
	self.elementbase = elementbase
	self.miner.SetElementbase(elementbase)
}

func (s *Elementrem) StopMining()         { s.miner.Stop() }
func (s *Elementrem) IsMining() bool      { return s.miner.Mining() }
func (s *Elementrem) Miner() *miner.Miner { return s.miner }

func (s *Elementrem) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Elementrem) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Elementrem) TxPool() *core.TxPool               { return s.txPool }
func (s *Elementrem) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Elementrem) ChainDb() eledb.Database            { return s.chainDb }
func (s *Elementrem) DappDb() eledb.Database             { return s.dappDb }
func (s *Elementrem) IsListening() bool                  { return true } // Always listening
func (s *Elementrem) EleVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Elementrem) NetVersion() int                    { return s.netVersionId }
func (s *Elementrem) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *Elementrem) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

// Start implements node.Service, starting all internal goroutines needed by the
// Elementrem protocol implementation.
func (s *Elementrem) Start(srvr *p2p.Server) error {
	if s.AutoDAG {
		s.StartAutoDAG()
	}
	s.protocolManager.Start()
	s.netRPCService = NewPublicNetAPI(srvr, s.NetVersion())
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Elementrem protocol.
func (s *Elementrem) Stop() error {
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.miner.Stop()
	s.eventMux.Stop()

	s.StopAutoDAG()

	s.chainDb.Close()
	s.dappDb.Close()
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

// HTTPClient returns the light http client used for fetching offchain docs
// (natspec, source for verification)
func (self *Elementrem) HTTPClient() *httpclient.HTTPClient {
	return self.httpclient
}

func (self *Elementrem) Solc() (*compiler.Solidity, error) {
	var err error
	if self.solc == nil {
		self.solc, err = compiler.New(self.SolcPath)
	}
	return self.solc, err
}

// set in js console via admin interface or wrapper from cli flags
func (self *Elementrem) SetSolc(solcPath string) (*compiler.Solidity, error) {
	self.SolcPath = solcPath
	self.solc = nil
	return self.Solc()
}

// dagFiles(epoch) returns the two alternative DAG filenames (not a path)
// 1) <revision>-<hex(seedhash[8])> 2) full-R<revision>-<hex(seedhash[8])>
func dagFiles(epoch uint64) (string, string) {
	seedHash, _ := elhash.GetSeedHash(epoch * epochLength)
	dag := fmt.Sprintf("full-R%d-%x", elhashRevision, seedHash[:8])
	return dag, "full-R" + dag
}

// upgradeChainDatabase ensures that the chain database stores block split into
// separate header and body entries.
func upgradeChainDatabase(db eledb.Database) error {
	// Short circuit if the head block is stored already as separate header and body
	data, err := db.Get([]byte("LastBlock"))
	if err != nil {
		return nil
	}
	head := common.BytesToHash(data)

	if block := core.GetBlockByHashOld(db, head); block == nil {
		return nil
	}
	// At least some of the database is still the old format, upgrade (skip the head block!)
	glog.V(logger.Info).Info("Old database detected, upgrading...")

	if db, ok := db.(*eledb.LDBDatabase); ok {
		blockPrefix := []byte("block-hash-")
		for it := db.NewIterator(); it.Next(); {
			// Skip anything other than a combined block
			if !bytes.HasPrefix(it.Key(), blockPrefix) {
				continue
			}
			// Skip the head block (merge last to signal upgrade completion)
			if bytes.HasSuffix(it.Key(), head.Bytes()) {
				continue
			}
			// Load the block, split and serialize (order!)
			block := core.GetBlockByHashOld(db, common.BytesToHash(bytes.TrimPrefix(it.Key(), blockPrefix)))

			if err := core.WriteTd(db, block.Hash(), block.DeprecatedTd()); err != nil {
				return err
			}
			if err := core.WriteBody(db, block.Hash(), block.Body()); err != nil {
				return err
			}
			if err := core.WriteHeader(db, block.Header()); err != nil {
				return err
			}
			if err := db.Delete(it.Key()); err != nil {
				return err
			}
		}
		// Lastly, upgrade the head block, disabling the upgrade mechanism
		current := core.GetBlockByHashOld(db, head)

		if err := core.WriteTd(db, current.Hash(), current.DeprecatedTd()); err != nil {
			return err
		}
		if err := core.WriteBody(db, current.Hash(), current.Body()); err != nil {
			return err
		}
		if err := core.WriteHeader(db, current.Header()); err != nil {
			return err
		}
	}
	return nil
}

func addMipmapBloomBins(db eledb.Database) (err error) {
	const mipmapVersion uint = 2

	// check if the version is set. We ignore data for now since there's
	// only one version so we can easily ignore it for now
	var data []byte
	data, _ = db.Get([]byte("setting-mipmap-version"))
	if len(data) > 0 {
		var version uint
		if err := rlp.DecodeBytes(data, &version); err == nil && version == mipmapVersion {
			return nil
		}
	}

	defer func() {
		if err == nil {
			var val []byte
			val, err = rlp.EncodeToBytes(mipmapVersion)
			if err == nil {
				err = db.Put([]byte("setting-mipmap-version"), val)
			}
			return
		}
	}()
	latestBlock := core.GetBlock(db, core.GetHeadBlockHash(db))
	if latestBlock == nil { // clean database
		return
	}

	tstart := time.Now()
	glog.V(logger.Info).Infoln("upgrading db log bloom bins")
	for i := uint64(0); i <= latestBlock.NumberU64(); i++ {
		hash := core.GetCanonicalHash(db, i)
		if (hash == common.Hash{}) {
			return fmt.Errorf("chain db corrupted. Could not find block %d.", i)
		}
		core.WriteMipmapBloom(db, i, core.GetBlockReceipts(db, hash))
	}
	glog.V(logger.Info).Infoln("upgrade completed in", time.Since(tstart))
	return nil
}
