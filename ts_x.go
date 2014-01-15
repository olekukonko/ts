// +build !windows

package ts

import (
	"syscall"
	"unsafe"
)




func NewSize() Size {
	ws := Size{}
	if TIOCGWINSZ != 0 {
		syscall.Syscall(syscall.SYS_IOCTL,
			uintptr(0),
			uintptr(TIOCGWINSZ),
			uintptr(unsafe.Pointer(&ws)))

		return int(ws.ws_col)
	}
	return ws
}
