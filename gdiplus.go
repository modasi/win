// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libgdiplus uintptr

	// Functions
	gdiplusNotificationHook   uintptr
	gdiplusNotificationUnhook uintptr
	gdiplusStartup            uintptr
)

func init() {
	// Library
	libgdiplus = doLoadLibrary("gdiplus.dll")

	// Functions
	gdiplusNotificationHook = doGetProcAddress(libgdiplus, "GdiplusNotificationHook")
	gdiplusNotificationUnhook = doGetProcAddress(libgdiplus, "GdiplusNotificationUnhook")
	gdiplusStartup = doGetProcAddress(libgdiplus, "GdiplusStartup")
}

func GdiplusNotificationHook(token *ULONG_PTR) GpStatus {
	ret1 := syscall3(gdiplusNotificationHook, 1,
		uintptr(unsafe.Pointer(token)),
		0,
		0)
	return GpStatus(ret1)
}

func GdiplusNotificationUnhook(token *uint32) {
	syscall3(gdiplusNotificationUnhook, 1,
		uintptr(unsafe.Pointer(token)),
		0,
		0)
}

func GdiplusStartup(token *ULONG_PTR, input /*const*/ *GdiplusStartupInput, output *GdiplusStartupOutput) Status {
	ret1 := syscall3(gdiplusStartup, 3,
		uintptr(unsafe.Pointer(token)),
		uintptr(unsafe.Pointer(input)),
		uintptr(unsafe.Pointer(output)))
	return Status(ret1)
}
