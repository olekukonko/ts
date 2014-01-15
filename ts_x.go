// +build !windows

package ts

import (
	"syscall"
	"unsafe"
)

// Get Windows Size
func GetSize() (ws Size, err error) {

	rc, _, ec := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))

	if rc == 0 {
		// Set Default size if OS is unknown
		if TIOCGWINSZ == 0 {
			ws = Size{80, 25, 0, 0}
		}
		err = syscall.Errno(ec)
		return ws, err
	}

	return ws, err
}
