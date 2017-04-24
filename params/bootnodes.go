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
	"enode://4c20c27d0331c2961faa7f678d74a03344dc6f6562ee62fff3891e85297f1ce48d1a667900d0da92b8085ed26883d7521b4a4ed2ff542b7ad2fa3cc4f1554e40@52.51.84.216:30707",
	"enode://a24a2c48f6b42c52521c2a0e7dff64c1f4ca98226aadbba499d08a59eb19b2b0802aa35f77b668ff1eb39d7f2ef53445e8624f02a7b42930355b43bc59cb4c93@50.17.164.173:30707",
	"enode://ef13e446f647dd3d86bace64bd0e704b4782ae9eb9930154eae0e18975f7bb13b6a0150da32b033de47423a425f205e76f2099d5f647fc32c61c073c40a3ae41@52.160.107.78:30707",
	"enode://79bab1a6035da93c343a8f27f70d19aa588f13a0c03cbed8f61cae8cb486816fe2b2cf3b2e1540100d43f5d0b69fff71050aa9acfca3802f03fb73493472a055@34.195.98.164:30707",
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
