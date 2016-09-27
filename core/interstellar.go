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

package core

import (
	"bytes"
	"math/big"

	"github.com/elementrem/go-elementrem/core/state"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/params"
)

// ValidateINTERSTELLARHeaderExtraData validates the extra-data field of a block header to
// ensure it conforms to INTERSTELLAR hyperz-leap rules.
//
// INTERSTELLAR hyperz-leap extension to the header validity:
//   a) if the node is no-fork, do not accept blocks in the [fork, fork+10) range
//      with the fork specific extra-data set
//   b) if the node is pro-fork, require blocks in the specific range to have the
//      unique extra-data set.
func ValidateINTERSTELLARHeaderExtraData(config *ChainConfig, header *types.Header) error {
	// Short circuit validation if the node doesn't care about the INTERSTELLAR fork
	if config.INTERSTELLARleapBlock == nil {
		return nil
	}
	// Make sure the block is within the fork's modified extra-data range
	limit := new(big.Int).Add(config.INTERSTELLARleapBlock, params.INTERSTELLARleapExtraRange)
	if header.Number.Cmp(config.INTERSTELLARleapBlock) < 0 || header.Number.Cmp(limit) >= 0 {
		return nil
	}
	// Depending whether we support or oppose the fork, validate the extra-data contents
	if config.INTERSTELLARleapSupport {
		if bytes.Compare(header.Extra, params.INTERSTELLARleapBlockExtra) != 0 {
			return ValidationError("INTERSTELLAR pro-fork bad block extra-data: 0x%x", header.Extra)
		}
	} else {
		if bytes.Compare(header.Extra, params.INTERSTELLARleapBlockExtra) == 0 {
			return ValidationError("INTERSTELLAR no-fork bad block extra-data: 0x%x", header.Extra)
		}
	}
	// All ok, header has the same extra-data we expect
	return nil
}

// ApplyINTERSTELLARHyperzLeap modifies the state database according to the INTERSTELLAR hyperz-leap
// rules, transferring all balances of a set of INTERSTELLAR accounts to a single refund
// contract.
func ApplyINTERSTELLARHyperzLeap(statedb *state.StateDB) {
	// Retrieve the contract to refund balances into
	refund := statedb.GetOrNewStateObject(params.INTERSTELLARRefundContract)

	// Move every INTERSTELLAR account and extra-balance account funds into the refund contract
	for _, addr := range params.INTERATELLARDrainList {
		if account := statedb.GetStateObject(addr); account != nil {
			refund.AddBalance(account.Balance())
			account.SetBalance(new(big.Int))
		}
	}
}
