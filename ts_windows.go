// +build windows

package ts

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	screenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

type COORD struct {
	X,        Y int16
}

type SMALL_RECT struct {
	Left,          Top,          Right,          Bottom int16
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         uint16
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

func Getize() Size {
	var csbi CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := screenBufferInfo.Call(
		uintptr(syscall.Stdout),
		uintptr(unsafe.Pointer(&csbi)))

	fmt.Println(ret, csbi)

	size := Size{csbi.SrWindow.Right,
		csbi.SrWindow.Bottom,
		csbi.DwCursorPosition.X,
		csbi.DwCursorPosition.Y}

	return size
}
