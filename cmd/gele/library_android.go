// Copyright 2016 The go-elementrem Authors.
// This file is part of go-elementrem.
//
// go-elementrem is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-elementrem is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-elementrem. If not, see <http://www.gnu.org/licenses/>.

// Contains specialized code for running Gele on Android.

package main

// #include <android/log.h>
// #cgo LDFLAGS: -llog
import "C"
import (
	"bufio"
	"os"
)

func init() {
	// Redirect the standard output and error to logcat
	oldStdout, oldStderr := os.Stdout, os.Stderr

	outRead, outWrite, _ := os.Pipe()
	errRead, errWrite, _ := os.Pipe()

	os.Stdout = outWrite
	os.Stderr = errWrite

	go func() {
		scanner := bufio.NewScanner(outRead)
		for scanner.Scan() {
			line := scanner.Text()
			C.__android_log_write(C.ANDROID_LOG_INFO, C.CString("Stdout"), C.CString(line))
			oldStdout.WriteString(line + "\n")
		}
	}()

	go func() {
		scanner := bufio.NewScanner(errRead)
		for scanner.Scan() {
			line := scanner.Text()
			C.__android_log_write(C.ANDROID_LOG_INFO, C.CString("Stderr"), C.CString(line))
			oldStderr.WriteString(line + "\n")
		}
	}()
}
