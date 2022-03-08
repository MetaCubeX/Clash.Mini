package controller

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32            = windows.NewLazySystemDLL("user32")
	User32SendMessage = user32.NewProc("SendMessageW")

	shell32           = windows.NewLazySystemDLL("shell32")
	User32ExtractIcon = shell32.NewProc("ExtractIconW")
)

func ExtractIcon(exeFileName string, iconIndex int32) uintptr {
	e, _ := syscall.UTF16PtrFromString(exeFileName)
	ret, _, _ := User32ExtractIcon.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(e)),
		uintptr(iconIndex),
	)
	return ret
}

func SendMessage(hWnd unsafe.Pointer, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(User32SendMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg),
		wParam, lParam,
		0, 0)
	return ret
}
