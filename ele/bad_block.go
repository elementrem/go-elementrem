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

package ele

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/elementrem/go-elementrem/common"
	"github.com/elementrem/go-elementrem/core/types"
	"github.com/elementrem/go-elementrem/logger"
	"github.com/elementrem/go-elementrem/logger/glog"
	"github.com/elementrem/go-elementrem/rlp"
)

const (
	// The Elementrem main network genesis block.
	defaultGenesisHash = "0xc7130c992fb03c628c86439f8e5d1facdbf48f3f3c115ffd99a0e3abefa1016f"
	badBlocksURL       = ""
)

var EnableBadBlockReporting = false

func sendBadBlockReport(block *types.Block, err error) {
	if !EnableBadBlockReporting {
		return
	}

	var (
		blockRLP, _ = rlp.EncodeToBytes(block)
		params      = map[string]interface{}{
			"block":     common.Bytes2Hex(blockRLP),
			"blockHash": block.Hash().Hex(),
			"errortype": err.Error(),
			"client":    "go",
		}
	)
	if !block.ReceivedAt.IsZero() {
		params["receivedAt"] = block.ReceivedAt.UTC().String()
	}
	if p, ok := block.ReceivedFrom.(*peer); ok {
		params["receivedFrom"] = map[string]interface{}{
			"enode":           fmt.Sprintf("enode://%x@%v", p.ID(), p.RemoteAddr()),
			"name":            p.Name(),
			"protocolVersion": p.version,
		}
	}
	jsonStr, _ := json.Marshal(map[string]interface{}{"method": "ele_badBlock", "id": "1", "jsonrpc": "2.0", "params": []interface{}{params}})
	client := http.Client{Timeout: 8 * time.Second}
	resp, err := client.Post(badBlocksURL, "application/json", bytes.NewReader(jsonStr))
	if err != nil {
		glog.V(logger.Debug).Infoln(err)
		return
	}
	glog.V(logger.Debug).Infof("Bad Block Report posted (%d)", resp.StatusCode)
	resp.Body.Close()
}
