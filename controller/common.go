package controller

import (
	"github.com/lxn/win"
	"syscall"
)

func init() {
	RefreshWindowResolution()
}

func RefreshWindowResolution()  {
	xScreen = win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen = win.GetSystemMetrics(win.SM_CYSCREEN)
	if dpiScale == 0 {
		dpiScale = float64(win.GetDpiForWindow(win.GetForegroundWindow())) / 96.0
	}
}

func GetTaskbarHeight() int32 {
	ptr1, err := syscall.UTF16PtrFromString("Shell_TrayWnd")
	if err != nil {
		return 0
	}
	ptr2, err := syscall.UTF16PtrFromString("")
	if err != nil {
		return 0
	}
	var rect *win.RECT
	if !win.GetWindowRect(win.FindWindow(ptr1, ptr2), rect) {
		return 0
	}
	return rect.Bottom - rect.Top
}

func CalcDpiScaledSize(sizeW int32, sizeH int32) (int32, int32) {
	return int32(float64(sizeW) * dpiScale), int32(float64(sizeH) * dpiScale)
}

func CalcDpiCenterScaledSize(screenSize, size int32) int32 {
	return int32((float64(screenSize) / dpiScale) - float64(size)) / 2
}
