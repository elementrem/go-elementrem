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

package params

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/elementrem/go-elementrem/common"
)

// INTERSTELLAR leap will jump to the entire block as the new strong network.
// the Elementrem test network is deprecated.
var TestNetINTERSTELLARleapBlock *big.Int

// MainNetINTERSTELLARleapBlock is the block number where the INTERSTELLAR hyperz-leap commences on
// the Elementrem main network.
var MainNetINTERSTELLARleapBlock = big.NewInt(0)

// INTERSTELLARleapBlockExtra is the block header extra-data field to set for the INTERSTELLAR fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
var INTERSTELLARleapBlockExtra = common.FromHex("0xaa57ee17a7257988a36b7b418cd716d315028b19")

// INTERSTELLARleapExtraRange is the number of consecutive blocks from the INTERSTELLAR fork point
// to override the extra-data in to prevent no-fork attacks.
var INTERSTELLARleapExtraRange = big.NewInt(1)

// INTERSTELLARRefundContract is the address of the none.
var INTERSTELLARRefundContract = common.HexToAddress("0x")

// INTERATELLARDrainList is deprecated.
var INTERATELLARDrainList []common.Address

func init() {
	// Parse the list of INTERSTELLAR accounts to drain
	var list []map[string]string
	if err := json.Unmarshal([]byte(interstellarDrainListJSON), &list); err != nil {
		panic(fmt.Errorf("Failed to parse INTERSTELLAR drain list: %v", err))
	}
	// Collect all the accounts that need draining
	for _, interstellar := range list {
		INTERATELLARDrainList = append(INTERATELLARDrainList, common.HexToAddress(interstellar["address"]))
		INTERATELLARDrainList = append(INTERATELLARDrainList, common.HexToAddress(interstellar["extraBalanceAccount"]))
	}
}

const interstellarDrainListJSON = `
[
   {
      "address":"",
      "balance":"",
      "extraBalance":"",
      "extraBalanceAccount":""
   }
]
`
