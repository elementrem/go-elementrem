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

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/elementrem/go-elementrem/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("ele/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("ele/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("ele/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("ele/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("ele/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("ele/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("ele/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("ele/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("ele/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("ele/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("ele/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("ele/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("ele/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("ele/fetcher/filter/bodies/out")
)
