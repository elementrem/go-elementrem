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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/elementrem/go-elementrem/metrics"
)

var (
	hashInMeter      = metrics.NewMeter("ele/downloader/hashes/in")
	hashReqTimer     = metrics.NewTimer("ele/downloader/hashes/req")
	hashDropMeter    = metrics.NewMeter("ele/downloader/hashes/drop")
	hashTimeoutMeter = metrics.NewMeter("ele/downloader/hashes/timeout")

	blockInMeter      = metrics.NewMeter("ele/downloader/blocks/in")
	blockReqTimer     = metrics.NewTimer("ele/downloader/blocks/req")
	blockDropMeter    = metrics.NewMeter("ele/downloader/blocks/drop")
	blockTimeoutMeter = metrics.NewMeter("ele/downloader/blocks/timeout")

	headerInMeter      = metrics.NewMeter("ele/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("ele/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("ele/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("ele/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("ele/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("ele/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("ele/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("ele/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("ele/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("ele/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("ele/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("ele/downloader/receipts/timeout")

	stateInMeter      = metrics.NewMeter("ele/downloader/states/in")
	stateReqTimer     = metrics.NewTimer("ele/downloader/states/req")
	stateDropMeter    = metrics.NewMeter("ele/downloader/states/drop")
	stateTimeoutMeter = metrics.NewMeter("ele/downloader/states/timeout")
)
