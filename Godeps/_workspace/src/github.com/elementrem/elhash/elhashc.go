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

package elhash

/*
 -mno-stack-arg-probe disables stack probing which avoids the function
 __chkstk_ms being linked. this avoids a clash of this symbol as we also
 separately link the secp256k1 lib which ends up defining this symbol

 1. https://gcc.gnu.org/onlinedocs/gccint/Stack-Checking.html
 2. https://groups.google.com/forum/#!msg/golang-dev/v1bziURSQ4k/88fXuJ24e-gJ
 3. https://groups.google.com/forum/#!topic/golang-nuts/VNP6Mwz_B6o

*/

/*
#cgo CFLAGS: -std=gnu99 -Wall
#cgo windows CFLAGS: -mno-stack-arg-probe
#cgo LDFLAGS: -lm

#include "src/libelhash/internal.c"
#include "src/libelhash/sha3.c"
#include "src/libelhash/io.c"

#ifdef _WIN32
#	include "src/libelhash/io_win32.c"
#	include "src/libelhash/mmap_win32.c"
#else
#	include "src/libelhash/io_posix.c"
#endif

// 'gateway function' for calling back into go.
extern int elhashGoCallback(unsigned);
int elhashGoCallback_cgo(unsigned percent) { return elhashGoCallback(percent); }

*/
import "C"
