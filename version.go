// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libversion uintptr

	// Functions
	getFileVersionInfoSize uintptr
	getFileVersionInfo     uintptr
	verFindFile            uintptr
	verInstallFile         uintptr
	verQueryValue          uintptr
)

func init() {
	// Library
	libversion = doLoadLibrary("version.dll")

	// Functions
	getFileVersionInfoSize = doGetProcAddress(libversion, "GetFileVersionInfoSizeW")
	getFileVersionInfo = doGetProcAddress(libversion, "GetFileVersionInfoW")
	verFindFile = doGetProcAddress(libversion, "VerFindFileW")
	verInstallFile = doGetProcAddress(libversion, "VerInstallFileW")
	verQueryValue = doGetProcAddress(libversion, "VerQueryValueW")
}

func GetFileVersionInfoSize(lptstrFilename string, lpdwHandle *uint32) uint32 {
	lptstrFilenameStr := unicode16FromString(lptstrFilename)
	ret1 := syscall3(getFileVersionInfoSize, 2,
		uintptr(unsafe.Pointer(&lptstrFilenameStr[0])),
		uintptr(unsafe.Pointer(lpdwHandle)),
		0)
	return uint32(ret1)
}

func GetFileVersionInfo(lptstrFilename string, dwHandle uint32, dwLen uint32, lpData LPVOID) bool {
	lptstrFilenameStr := unicode16FromString(lptstrFilename)
	ret1 := syscall6(getFileVersionInfo, 4,
		uintptr(unsafe.Pointer(&lptstrFilenameStr[0])),
		uintptr(dwHandle),
		uintptr(dwLen),
		uintptr(unsafe.Pointer(lpData)),
		0,
		0)
	return ret1 != 0
}

func VerFindFile(uFlags uint32, szFileName LPWSTR, szWinDir LPWSTR, szAppDir LPWSTR, szCurDir LPWSTR, lpuCurDirLen *uint32, szDestDir LPWSTR, lpuDestDirLen *uint32) uint32 {
	ret1 := syscall9(verFindFile, 8,
		uintptr(uFlags),
		uintptr(unsafe.Pointer(szFileName)),
		uintptr(unsafe.Pointer(szWinDir)),
		uintptr(unsafe.Pointer(szAppDir)),
		uintptr(unsafe.Pointer(szCurDir)),
		uintptr(unsafe.Pointer(lpuCurDirLen)),
		uintptr(unsafe.Pointer(szDestDir)),
		uintptr(unsafe.Pointer(lpuDestDirLen)),
		0)
	return uint32(ret1)
}

func VerInstallFile(uFlags uint32, szSrcFileName LPWSTR, szDestFileName LPWSTR, szSrcDir LPWSTR, szDestDir LPWSTR, szCurDir LPWSTR, szTmpFile LPWSTR, lpuTmpFileLen *uint32) uint32 {
	ret1 := syscall9(verInstallFile, 8,
		uintptr(uFlags),
		uintptr(unsafe.Pointer(szSrcFileName)),
		uintptr(unsafe.Pointer(szDestFileName)),
		uintptr(unsafe.Pointer(szSrcDir)),
		uintptr(unsafe.Pointer(szDestDir)),
		uintptr(unsafe.Pointer(szCurDir)),
		uintptr(unsafe.Pointer(szTmpFile)),
		uintptr(unsafe.Pointer(lpuTmpFileLen)),
		0)
	return uint32(ret1)
}

func VerQueryValue(pBlock /*const*/ LPVOID, lpSubBlock string, lplpBuffer *LPVOID, puLen *uint32) bool {
	lpSubBlockStr := unicode16FromString(lpSubBlock)
	ret1 := syscall6(verQueryValue, 4,
		uintptr(unsafe.Pointer(pBlock)),
		uintptr(unsafe.Pointer(&lpSubBlockStr[0])),
		uintptr(unsafe.Pointer(lplpBuffer)),
		uintptr(unsafe.Pointer(puLen)),
		0,
		0)
	return ret1 != 0
}
