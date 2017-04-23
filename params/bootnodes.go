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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Elementrem network.
var MainnetBootnodes = []string{
	// ELE/DEV Go Bootnodes
	"enode://fb28247da1355b8ba8d0c81efc93aba6a4a52865e8918d8a683ce387c4e22c104553ebc3c8bb0e2222a211b7c7ce47626a84216fb64ea6cc2fff355383ecc12e@52.51.84.216:30707",
	"enode://95948e41d02f524d5b44672a6e8bb9c59b734694247b4d8f4123234605704b72a7ffc238d385d0643530c62d11bf5bc8b5fa27050e017df1291932c26c252c4f@50.17.164.173:30707",
	"enode://f408ea82752eed6d141e85fa7e12fcc7fb0c17656c46a2ebae1b2c88545e68313766d488fd6ccfc14feaa1fa69bf4e8915e1e6481524ae7c51a1565c44739e27@52.160.107.78:30707",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestnetBootnodes = []string{
	// ELE/DEV Go Bootnodes
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
