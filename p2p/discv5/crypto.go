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

package discv5

import (
	//"github.com/btcsuite/btcd/btcec"
	"github.com/elementrem/go-elementrem/crypto/secp256k1"
)

func S256() *secp256k1.BitCurve {
	return secp256k1.S256()
}

// This version should be used for NaCl compilation
/*func S256() *btcec.KoblitzCurve {
	return S256()
}*/
