// +build !windows

// Copyright 2014 Oleku Konko All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This module is a Terminal  API for the Go Programming Language.
// The protocols were written in pure Go and works on windows and unix systems

package ts

import (
	"syscall"
	"unsafe"
)

// Get Windows Size
func GetSize() (ws Size, err error) {
	_, _, ec := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))

	err = getError(ec)

	if TIOCGWINSZ == 0 && err != nil {
		ws = Size{80, 25, 0, 0}
	}
	return ws, err
}

func getError(ec interface{}) (err error) {
	switch v := ec.(type) {
		// Some implementation return syscall.Errno number
	case syscall.Errno:
		if v != 0 {
			err = syscall.Errno(v)
		}
		// Some implementation return error
	case error:
		err = ec.(error)
	default:
		err = nil
	}
	return
}
