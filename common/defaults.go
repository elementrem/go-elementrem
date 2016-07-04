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

package common

import (
	"path/filepath"
	"runtime"
)

const (
	DefaultIPCSocket = "gele.ipc"  // Default (relative) name of the IPC RPC socket
	DefaultHTTPHost  = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort  = 7075        // Default TCP port for the HTTP RPC server
	DefaultWSHost    = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort    = 7076        // Default TCP port for the websocket RPC server
)

// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := HomeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "Elementrem")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "Elementrem")
		} else {
			return filepath.Join(home, ".elementrem")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}
