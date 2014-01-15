// +build windows

package ts

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// Retrieves information about the specified console screen buffer.
	// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms683171(v=vs.85).aspx
	screenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

//   Contains information about a console screen buffer.
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms682093(v=vs.85).aspx
type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         uint16
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

// Defines the coordinates of a character cell in a console screen buffer.
// The origin of the coordinate system (0,0) is at the top, left cell of the buffer.
// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms682119(v=vs.85).aspx
type COORD struct {
	X, Y uint16
}

// Defines the coordinates of the upper left and lower right corners of a rectangle.
// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms686311(v=vs.85).aspx
type SMALL_RECT struct {
	Left, Top, Right, Bottom uint16
}

func GetSize() (ws Size, err error) {
	var info CONSOLE_SCREEN_BUFFER_INFO
	rc, _, err := screenBufferInfo.Call(
		uintptr(syscall.Stdout),
		uintptr(unsafe.Pointer(&info)))

	if rc == 0 {
		return ws, err
	}

	ws = Size{info.SrWindow.Right,
		info.SrWindow.Bottom,
		info.DwCursorPosition.X,
		info.DwCursorPosition.Y}

	return ws, nil
}
