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

// Contains all the wrappers from the params package.

package gele

import (
	"github.com/elementrem/go-elementrem/core"
	"github.com/elementrem/go-elementrem/p2p/discv5"
	"github.com/elementrem/go-elementrem/params"
)

// MainnetChainConfig returns the chain configurations for the main Elementrem network.
func MainnetChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:        params.MainNetChainID.Int64(),
		HomesteadBlock: params.MainNetHomesteadBlock.Int64(),
		INTERSTELLARleapBlock:   params.MainNetINTERSTELLARleapBlock.Int64(),
		INTERSTELLARleapSupport: true,
		EIP150Block:    params.MainNetHomesteadGasRepriceBlock.Int64(),
		EIP150Hash:     Hash{params.MainNetHomesteadGasRepriceHash},
		EIP155Block:    params.MainNetSpuriousDragon.Int64(),
		EIP158Block:    params.MainNetSpuriousDragon.Int64(),
	}
}

// MainnetGenesis returns the JSON spec to use for the main Elementrem network. It
// is actually empty since that defaults to the hard coded binary genesis block.
func MainnetGenesis() string {
	return ""
}

// TestnetChainConfig returns the chain configurations for the Elementrem test network.
func TestnetChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:        params.TestNetChainID.Int64(),
		HomesteadBlock: params.TestNetHomesteadBlock.Int64(),
		INTERSTELLARleapBlock:   0,
		INTERSTELLARleapSupport: false,
		EIP150Block:    params.TestNetHomesteadGasRepriceBlock.Int64(),
		EIP150Hash:     Hash{params.TestNetHomesteadGasRepriceHash},
		EIP155Block:    params.TestNetSpuriousDragon.Int64(),
		EIP158Block:    params.TestNetSpuriousDragon.Int64(),
	}
}

// TestnetGenesis returns the JSON spec to use for the Elementrem test network.
func TestnetGenesis() string {
	return core.DefaultTestnetGenesisBlock()
}

// ChainConfig is the core config which determines the blockchain settings.
type ChainConfig struct {
	ChainID        int64 // Chain ID for replay protection
	HomesteadBlock int64 // Homestead switch block
	INTERSTELLARleapBlock   int64 // TheINTERSTELLAR hyperz-leap switch block
	INTERSTELLARleapSupport bool  // Whether the nodes supports or opposes the INTERSTELLAR hyperz-leap
	EIP150Block    int64 // Homestead gas reprice switch block
	EIP150Hash     Hash  // Homestead gas reprice switch block hash
	EIP155Block    int64 // Replay protection switch block
	EIP158Block    int64 // Empty account pruning switch block
}

// NewChainConfig creates a new chain configuration that transitions immediately
// to homestead and has no notion of the INTERSTELLAR leap (ideal for a private network).
func NewChainConfig() *ChainConfig {
	return new(ChainConfig)
}

// FoundationBootnodes returns the enode URLs of the P2P bootstrap nodes operated
// by the foundation running the V5 discovery protocol.
func FoundationBootnodes() *Enodes {
	nodes := &Enodes{nodes: make([]*discv5.Node, len(params.DiscoveryV5Bootnodes))}
	for i, url := range params.DiscoveryV5Bootnodes {
		nodes.nodes[i] = discv5.MustParseNode(url)
	}
	return nodes
}
