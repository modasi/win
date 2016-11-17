// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libimm32 uintptr

	// Functions
	immAssociateContext        uintptr
	immAssociateContextEx      uintptr
	immConfigureIME            uintptr
	immCreateContext           uintptr
	immDestroyContext          uintptr
	immDisableIME              uintptr
	immDisableTextFrameService uintptr
	immEnumInputContext        uintptr
	immEnumRegisterWord        uintptr
	immEscape                  uintptr
	immGetCandidateListCount   uintptr
	immGetCandidateList        uintptr
	immGetCandidateWindow      uintptr
	immGetCompositionFont      uintptr
	immGetCompositionString    uintptr
	immGetCompositionWindow    uintptr
	immGetContext              uintptr
	immGetConversionList       uintptr
	immGetConversionStatus     uintptr
	immGetDefaultIMEWnd        uintptr
	immGetDescription          uintptr
	immGetGuideLine            uintptr
	immGetIMEFileName          uintptr
	immGetImeMenuItems         uintptr
	immGetOpenStatus           uintptr
	immGetProperty             uintptr
	immGetRegisterWordStyle    uintptr
	immGetStatusWindowPos      uintptr
	immGetVirtualKey           uintptr
	immInstallIME              uintptr
	immIsIME                   uintptr
	immIsUIMessage             uintptr
	immNotifyIME               uintptr
	immRegisterWord            uintptr
	immReleaseContext          uintptr
	immSetCandidateWindow      uintptr
	immSetCompositionFont      uintptr
	immSetCompositionString    uintptr
	immSetCompositionWindow    uintptr
	immSetConversionStatus     uintptr
	immSetOpenStatus           uintptr
	immSetStatusWindowPos      uintptr
	immSimulateHotKey          uintptr
	immUnregisterWord          uintptr
	immCreateIMCC              uintptr
	immCreateSoftKeyboard      uintptr
	immDestroyIMCC             uintptr
	immDestroySoftKeyboard     uintptr
	immGenerateMessage         uintptr
	immGetHotKey               uintptr
	immGetIMCCLockCount        uintptr
	immGetIMCCSize             uintptr
	immGetIMCLockCount         uintptr
	immLockIMC                 uintptr
	immLockIMCC                uintptr
	immProcessKey              uintptr
	immReSizeIMCC              uintptr
	immRequestMessage          uintptr
	immShowSoftKeyboard        uintptr
	immTranslateMessage        uintptr
	immUnlockIMC               uintptr
	immUnlockIMCC              uintptr
)

func init() {
	// Library
	libimm32 = doLoadLibrary("imm32.dll")

	// Functions
	immAssociateContext = doGetProcAddress(libimm32, "ImmAssociateContext")
	immAssociateContextEx = doGetProcAddress(libimm32, "ImmAssociateContextEx")
	immConfigureIME = doGetProcAddress(libimm32, "ImmConfigureIMEW")
	immCreateContext = doGetProcAddress(libimm32, "ImmCreateContext")
	immDestroyContext = doGetProcAddress(libimm32, "ImmDestroyContext")
	immDisableIME = doGetProcAddress(libimm32, "ImmDisableIME")
	immDisableTextFrameService = doGetProcAddress(libimm32, "ImmDisableTextFrameService")
	immEnumInputContext = doGetProcAddress(libimm32, "ImmEnumInputContext")
	immEnumRegisterWord = doGetProcAddress(libimm32, "ImmEnumRegisterWordW")
	immEscape = doGetProcAddress(libimm32, "ImmEscapeW")
	immGetCandidateListCount = doGetProcAddress(libimm32, "ImmGetCandidateListCountW")
	immGetCandidateList = doGetProcAddress(libimm32, "ImmGetCandidateListW")
	immGetCandidateWindow = doGetProcAddress(libimm32, "ImmGetCandidateWindow")
	immGetCompositionFont = doGetProcAddress(libimm32, "ImmGetCompositionFontW")
	immGetCompositionString = doGetProcAddress(libimm32, "ImmGetCompositionStringW")
	immGetCompositionWindow = doGetProcAddress(libimm32, "ImmGetCompositionWindow")
	immGetContext = doGetProcAddress(libimm32, "ImmGetContext")
	immGetConversionList = doGetProcAddress(libimm32, "ImmGetConversionListW")
	immGetConversionStatus = doGetProcAddress(libimm32, "ImmGetConversionStatus")
	immGetDefaultIMEWnd = doGetProcAddress(libimm32, "ImmGetDefaultIMEWnd")
	immGetDescription = doGetProcAddress(libimm32, "ImmGetDescriptionW")
	immGetGuideLine = doGetProcAddress(libimm32, "ImmGetGuideLineW")
	immGetIMEFileName = doGetProcAddress(libimm32, "ImmGetIMEFileNameW")
	immGetImeMenuItems = doGetProcAddress(libimm32, "ImmGetImeMenuItemsW")
	immGetOpenStatus = doGetProcAddress(libimm32, "ImmGetOpenStatus")
	immGetProperty = doGetProcAddress(libimm32, "ImmGetProperty")
	immGetRegisterWordStyle = doGetProcAddress(libimm32, "ImmGetRegisterWordStyleW")
	immGetStatusWindowPos = doGetProcAddress(libimm32, "ImmGetStatusWindowPos")
	immGetVirtualKey = doGetProcAddress(libimm32, "ImmGetVirtualKey")
	immInstallIME = doGetProcAddress(libimm32, "ImmInstallIMEW")
	immIsIME = doGetProcAddress(libimm32, "ImmIsIME")
	immIsUIMessage = doGetProcAddress(libimm32, "ImmIsUIMessageW")
	immNotifyIME = doGetProcAddress(libimm32, "ImmNotifyIME")
	immRegisterWord = doGetProcAddress(libimm32, "ImmRegisterWordW")
	immReleaseContext = doGetProcAddress(libimm32, "ImmReleaseContext")
	immSetCandidateWindow = doGetProcAddress(libimm32, "ImmSetCandidateWindow")
	immSetCompositionFont = doGetProcAddress(libimm32, "ImmSetCompositionFontW")
	immSetCompositionString = doGetProcAddress(libimm32, "ImmSetCompositionStringW")
	immSetCompositionWindow = doGetProcAddress(libimm32, "ImmSetCompositionWindow")
	immSetConversionStatus = doGetProcAddress(libimm32, "ImmSetConversionStatus")
	immSetOpenStatus = doGetProcAddress(libimm32, "ImmSetOpenStatus")
	immSetStatusWindowPos = doGetProcAddress(libimm32, "ImmSetStatusWindowPos")
	immSimulateHotKey = doGetProcAddress(libimm32, "ImmSimulateHotKey")
	immUnregisterWord = doGetProcAddress(libimm32, "ImmUnregisterWordW")
	immCreateIMCC = doGetProcAddress(libimm32, "ImmCreateIMCC")
	immCreateSoftKeyboard = doGetProcAddress(libimm32, "ImmCreateSoftKeyboard")
	immDestroyIMCC = doGetProcAddress(libimm32, "ImmDestroyIMCC")
	immDestroySoftKeyboard = doGetProcAddress(libimm32, "ImmDestroySoftKeyboard")
	immGenerateMessage = doGetProcAddress(libimm32, "ImmGenerateMessage")
	immGetHotKey = doGetProcAddress(libimm32, "ImmGetHotKey")
	immGetIMCCLockCount = doGetProcAddress(libimm32, "ImmGetIMCCLockCount")
	immGetIMCCSize = doGetProcAddress(libimm32, "ImmGetIMCCSize")
	immGetIMCLockCount = doGetProcAddress(libimm32, "ImmGetIMCLockCount")
	immLockIMC = doGetProcAddress(libimm32, "ImmLockIMC")
	immLockIMCC = doGetProcAddress(libimm32, "ImmLockIMCC")
	immProcessKey = doGetProcAddress(libimm32, "ImmProcessKey")
	immReSizeIMCC = doGetProcAddress(libimm32, "ImmReSizeIMCC")
	immRequestMessage = doGetProcAddress(libimm32, "ImmRequestMessageW")
	immShowSoftKeyboard = doGetProcAddress(libimm32, "ImmShowSoftKeyboard")
	immTranslateMessage = doGetProcAddress(libimm32, "ImmTranslateMessage")
	immUnlockIMC = doGetProcAddress(libimm32, "ImmUnlockIMC")
	immUnlockIMCC = doGetProcAddress(libimm32, "ImmUnlockIMCC")
}

func ImmAssociateContext(unnamed0 HWND, unnamed1 HIMC) HIMC {
	ret1 := syscall3(immAssociateContext, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return HIMC(ret1)
}

func ImmAssociateContextEx(unnamed0 HWND, unnamed1 HIMC, unnamed2 uint32) bool {
	ret1 := syscall3(immAssociateContextEx, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2))
	return ret1 != 0
}

func ImmConfigureIME(unnamed0 HKL, unnamed1 HWND, unnamed2 uint32, unnamed3 LPVOID) bool {
	ret1 := syscall6(immConfigureIME, 4,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unsafe.Pointer(unnamed3)),
		0,
		0)
	return ret1 != 0
}

func ImmCreateContext() HIMC {
	ret1 := syscall3(immCreateContext, 0,
		0,
		0,
		0)
	return HIMC(ret1)
}

func ImmDestroyContext(unnamed0 HIMC) bool {
	ret1 := syscall3(immDestroyContext, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func ImmDisableIME(unnamed0 uint32) bool {
	ret1 := syscall3(immDisableIME, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func ImmDisableTextFrameService(idThread uint32) bool {
	ret1 := syscall3(immDisableTextFrameService, 1,
		uintptr(idThread),
		0,
		0)
	return ret1 != 0
}

func ImmEnumInputContext(idThread uint32, lpfn IMCENUMPROC, lParam LPARAM) bool {
	lpfnCallback := syscall.NewCallback(func(unnamed0RawArg HIMC, unnamed1RawArg LPARAM) uintptr {
		ret := lpfn(unnamed0RawArg, unnamed1RawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(immEnumInputContext, 3,
		uintptr(idThread),
		lpfnCallback,
		uintptr(lParam))
	return ret1 != 0
}

func ImmEnumRegisterWord(unnamed0 HKL, unnamed1 REGISTERWORDENUMPROC, lpszReading string, unnamed3 uint32, lpszRegister string, unnamed5 LPVOID) uint32 {
	lpszReadingStr := unicode16FromString(lpszReading)
	lpszRegisterStr := unicode16FromString(lpszRegister)
	unnamed1Callback := syscall.NewCallback(func(lpszReadingRawArg /*const*/ *uint16, unnamed1RawArg DWORD, lpszStringRawArg /*const*/ *uint16, unnamed3RawArg LPVOID) uintptr {
		lpszReading := stringFromUnicode16(lpszReadingRawArg)
		lpszString := stringFromUnicode16(lpszStringRawArg)
		ret := unnamed1(lpszReading, unnamed1RawArg, lpszString, unnamed3RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(immEnumRegisterWord, 6,
		uintptr(unnamed0),
		unnamed1Callback,
		uintptr(unsafe.Pointer(&lpszReadingStr[0])),
		uintptr(unnamed3),
		uintptr(unsafe.Pointer(&lpszRegisterStr[0])),
		uintptr(unsafe.Pointer(unnamed5)))
	return uint32(ret1)
}

func ImmEscape(unnamed0 HKL, unnamed1 HIMC, unnamed2 uint32, unnamed3 LPVOID) LRESULT {
	ret1 := syscall6(immEscape, 4,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unsafe.Pointer(unnamed3)),
		0,
		0)
	return LRESULT(ret1)
}

func ImmGetCandidateListCount(unnamed0 HIMC, lpdwListCount *uint32) uint32 {
	ret1 := syscall3(immGetCandidateListCount, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(lpdwListCount)),
		0)
	return uint32(ret1)
}

func ImmGetCandidateList(unnamed0 HIMC, deIndex uint32, unnamed2 *CANDIDATELIST, dwBufLen uint32) uint32 {
	ret1 := syscall6(immGetCandidateList, 4,
		uintptr(unnamed0),
		uintptr(deIndex),
		uintptr(unsafe.Pointer(unnamed2)),
		uintptr(dwBufLen),
		0,
		0)
	return uint32(ret1)
}

func ImmGetCandidateWindow(unnamed0 HIMC, unnamed1 uint32, unnamed2 *CANDIDATEFORM) bool {
	ret1 := syscall3(immGetCandidateWindow, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unsafe.Pointer(unnamed2)))
	return ret1 != 0
}

func ImmGetCompositionFont(unnamed0 HIMC, unnamed1 LPLOGFONT) bool {
	ret1 := syscall3(immGetCompositionFont, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmGetCompositionString(unnamed0 HIMC, unnamed1 uint32, unnamed2 LPVOID, unnamed3 uint32) LONG {
	ret1 := syscall6(immGetCompositionString, 4,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unsafe.Pointer(unnamed2)),
		uintptr(unnamed3),
		0,
		0)
	return LONG(ret1)
}

func ImmGetCompositionWindow(unnamed0 HIMC, unnamed1 *COMPOSITIONFORM) bool {
	ret1 := syscall3(immGetCompositionWindow, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmGetContext(unnamed0 HWND) HIMC {
	ret1 := syscall3(immGetContext, 1,
		uintptr(unnamed0),
		0,
		0)
	return HIMC(ret1)
}

func ImmGetConversionList(unnamed0 HKL, unnamed1 HIMC, unnamed2 string, unnamed3 *CANDIDATELIST, dwBufLen uint32, uFlag uint32) uint32 {
	unnamed2Str := unicode16FromString(unnamed2)
	ret1 := syscall6(immGetConversionList, 6,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unsafe.Pointer(&unnamed2Str[0])),
		uintptr(unsafe.Pointer(unnamed3)),
		uintptr(dwBufLen),
		uintptr(uFlag))
	return uint32(ret1)
}

func ImmGetConversionStatus(unnamed0 HIMC, unnamed1 *uint32, unnamed2 *uint32) bool {
	ret1 := syscall3(immGetConversionStatus, 3,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(unsafe.Pointer(unnamed2)))
	return ret1 != 0
}

func ImmGetDefaultIMEWnd(unnamed0 HWND) HWND {
	ret1 := syscall3(immGetDefaultIMEWnd, 1,
		uintptr(unnamed0),
		0,
		0)
	return HWND(ret1)
}

func ImmGetDescription(unnamed0 HKL, unnamed1 LPWSTR, uBufLen uint32) uint32 {
	ret1 := syscall3(immGetDescription, 3,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(uBufLen))
	return uint32(ret1)
}

func ImmGetGuideLine(unnamed0 HIMC, dwIndex uint32, unnamed2 LPWSTR, dwBufLen uint32) uint32 {
	ret1 := syscall6(immGetGuideLine, 4,
		uintptr(unnamed0),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(unnamed2)),
		uintptr(dwBufLen),
		0,
		0)
	return uint32(ret1)
}

func ImmGetIMEFileName(unnamed0 HKL, unnamed1 LPWSTR, uBufLen uint32) uint32 {
	ret1 := syscall3(immGetIMEFileName, 3,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(uBufLen))
	return uint32(ret1)
}

func ImmGetImeMenuItems(unnamed0 HIMC, unnamed1 uint32, unnamed2 uint32, unnamed3 LPIMEMENUITEMINFO, unnamed4 LPIMEMENUITEMINFO, unnamed5 uint32) uint32 {
	ret1 := syscall6(immGetImeMenuItems, 6,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unsafe.Pointer(unnamed3)),
		uintptr(unsafe.Pointer(unnamed4)),
		uintptr(unnamed5))
	return uint32(ret1)
}

func ImmGetOpenStatus(unnamed0 HIMC) bool {
	ret1 := syscall3(immGetOpenStatus, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func ImmGetProperty(unnamed0 HKL, unnamed1 uint32) uint32 {
	ret1 := syscall3(immGetProperty, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return uint32(ret1)
}

func ImmGetRegisterWordStyle(unnamed0 HKL, nItem uint32, unnamed2 LPSTYLEBUF) uint32 {
	ret1 := syscall3(immGetRegisterWordStyle, 3,
		uintptr(unnamed0),
		uintptr(nItem),
		uintptr(unsafe.Pointer(unnamed2)))
	return uint32(ret1)
}

func ImmGetStatusWindowPos(unnamed0 HIMC, unnamed1 *POINT) bool {
	ret1 := syscall3(immGetStatusWindowPos, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmGetVirtualKey(unnamed0 HWND) uint32 {
	ret1 := syscall3(immGetVirtualKey, 1,
		uintptr(unnamed0),
		0,
		0)
	return uint32(ret1)
}

func ImmInstallIME(lpszIMEFileName string, lpszLayoutText string) HKL {
	lpszIMEFileNameStr := unicode16FromString(lpszIMEFileName)
	lpszLayoutTextStr := unicode16FromString(lpszLayoutText)
	ret1 := syscall3(immInstallIME, 2,
		uintptr(unsafe.Pointer(&lpszIMEFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpszLayoutTextStr[0])),
		0)
	return HKL(ret1)
}

func ImmIsIME(unnamed0 HKL) bool {
	ret1 := syscall3(immIsIME, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func ImmIsUIMessage(unnamed0 HWND, unnamed1 uint32, unnamed2 WPARAM, unnamed3 LPARAM) bool {
	ret1 := syscall6(immIsUIMessage, 4,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		0,
		0)
	return ret1 != 0
}

func ImmNotifyIME(unnamed0 HIMC, dwAction uint32, dwIndex uint32, dwValue uint32) bool {
	ret1 := syscall6(immNotifyIME, 4,
		uintptr(unnamed0),
		uintptr(dwAction),
		uintptr(dwIndex),
		uintptr(dwValue),
		0,
		0)
	return ret1 != 0
}

func ImmRegisterWord(unnamed0 HKL, lpszReading string, unnamed2 uint32, lpszRegister string) bool {
	lpszReadingStr := unicode16FromString(lpszReading)
	lpszRegisterStr := unicode16FromString(lpszRegister)
	ret1 := syscall6(immRegisterWord, 4,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(&lpszReadingStr[0])),
		uintptr(unnamed2),
		uintptr(unsafe.Pointer(&lpszRegisterStr[0])),
		0,
		0)
	return ret1 != 0
}

func ImmReleaseContext(unnamed0 HWND, unnamed1 HIMC) bool {
	ret1 := syscall3(immReleaseContext, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func ImmSetCandidateWindow(unnamed0 HIMC, unnamed1 *CANDIDATEFORM) bool {
	ret1 := syscall3(immSetCandidateWindow, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmSetCompositionFont(unnamed0 HIMC, unnamed1 LPLOGFONT) bool {
	ret1 := syscall3(immSetCompositionFont, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmSetCompositionString(unnamed0 HIMC, dwIndex uint32, lpComp LPVOID, unnamed3 uint32, lpRead LPVOID, unnamed5 uint32) bool {
	ret1 := syscall6(immSetCompositionString, 6,
		uintptr(unnamed0),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(lpComp)),
		uintptr(unnamed3),
		uintptr(unsafe.Pointer(lpRead)),
		uintptr(unnamed5))
	return ret1 != 0
}

func ImmSetCompositionWindow(unnamed0 HIMC, unnamed1 *COMPOSITIONFORM) bool {
	ret1 := syscall3(immSetCompositionWindow, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmSetConversionStatus(unnamed0 HIMC, unnamed1 uint32, unnamed2 uint32) bool {
	ret1 := syscall3(immSetConversionStatus, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2))
	return ret1 != 0
}

func ImmSetOpenStatus(unnamed0 HIMC, unnamed1 bool) bool {
	ret1 := syscall3(immSetOpenStatus, 2,
		uintptr(unnamed0),
		getUintptrFromBool(unnamed1),
		0)
	return ret1 != 0
}

func ImmSetStatusWindowPos(unnamed0 HIMC, unnamed1 *POINT) bool {
	ret1 := syscall3(immSetStatusWindowPos, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImmSimulateHotKey(unnamed0 HWND, unnamed1 uint32) bool {
	ret1 := syscall3(immSimulateHotKey, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func ImmUnregisterWord(unnamed0 HKL, lpszReading string, unnamed2 uint32, lpszUnregister string) bool {
	lpszReadingStr := unicode16FromString(lpszReading)
	lpszUnregisterStr := unicode16FromString(lpszUnregister)
	ret1 := syscall6(immUnregisterWord, 4,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(&lpszReadingStr[0])),
		uintptr(unnamed2),
		uintptr(unsafe.Pointer(&lpszUnregisterStr[0])),
		0,
		0)
	return ret1 != 0
}

func ImmCreateIMCC(size uint32) HIMCC {
	ret1 := syscall3(immCreateIMCC, 1,
		uintptr(size),
		0,
		0)
	return HIMCC(ret1)
}

func ImmCreateSoftKeyboard(uType uint32, hOwner uint32, x int32, y int32) HWND {
	ret1 := syscall6(immCreateSoftKeyboard, 4,
		uintptr(uType),
		uintptr(hOwner),
		uintptr(x),
		uintptr(y),
		0,
		0)
	return HWND(ret1)
}

func ImmDestroyIMCC(block HIMCC) HIMCC {
	ret1 := syscall3(immDestroyIMCC, 1,
		uintptr(block),
		0,
		0)
	return HIMCC(ret1)
}

func ImmDestroySoftKeyboard(hSoftWnd HWND) bool {
	ret1 := syscall3(immDestroySoftKeyboard, 1,
		uintptr(hSoftWnd),
		0,
		0)
	return ret1 != 0
}

func ImmGenerateMessage(hIMC HIMC) bool {
	ret1 := syscall3(immGenerateMessage, 1,
		uintptr(hIMC),
		0,
		0)
	return ret1 != 0
}

func ImmGetHotKey(hotkey uint32, modifiers *UINT, key *UINT, hkl HKL) bool {
	ret1 := syscall6(immGetHotKey, 4,
		uintptr(hotkey),
		uintptr(unsafe.Pointer(modifiers)),
		uintptr(unsafe.Pointer(key)),
		uintptr(hkl),
		0,
		0)
	return ret1 != 0
}

func ImmGetIMCCLockCount(imcc HIMCC) uint32 {
	ret1 := syscall3(immGetIMCCLockCount, 1,
		uintptr(imcc),
		0,
		0)
	return uint32(ret1)
}

func ImmGetIMCCSize(imcc HIMCC) uint32 {
	ret1 := syscall3(immGetIMCCSize, 1,
		uintptr(imcc),
		0,
		0)
	return uint32(ret1)
}

func ImmGetIMCLockCount(hIMC HIMC) uint32 {
	ret1 := syscall3(immGetIMCLockCount, 1,
		uintptr(hIMC),
		0,
		0)
	return uint32(ret1)
}

func ImmLockIMC(hIMC HIMC) *INPUTCONTEXT {
	ret1 := syscall3(immLockIMC, 1,
		uintptr(hIMC),
		0,
		0)
	return (*INPUTCONTEXT)(unsafe.Pointer(ret1))
}

func ImmLockIMCC(imcc HIMCC) LPVOID {
	ret1 := syscall3(immLockIMCC, 1,
		uintptr(imcc),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func ImmProcessKey(hwnd HWND, hKL HKL, vKey uint32, lKeyData LPARAM, unknown uint32) bool {
	ret1 := syscall6(immProcessKey, 5,
		uintptr(hwnd),
		uintptr(hKL),
		uintptr(vKey),
		uintptr(lKeyData),
		uintptr(unknown),
		0)
	return ret1 != 0
}

func ImmReSizeIMCC(imcc HIMCC, size uint32) HIMCC {
	ret1 := syscall3(immReSizeIMCC, 2,
		uintptr(imcc),
		uintptr(size),
		0)
	return HIMCC(ret1)
}

func ImmRequestMessage(hIMC HIMC, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall3(immRequestMessage, 3,
		uintptr(hIMC),
		uintptr(wParam),
		uintptr(lParam))
	return LRESULT(ret1)
}

func ImmShowSoftKeyboard(hSoftWnd HWND, nCmdShow int32) bool {
	ret1 := syscall3(immShowSoftKeyboard, 2,
		uintptr(hSoftWnd),
		uintptr(nCmdShow),
		0)
	return ret1 != 0
}

func ImmTranslateMessage(hwnd HWND, msg uint32, wParam WPARAM, lKeyData LPARAM) bool {
	ret1 := syscall6(immTranslateMessage, 4,
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lKeyData),
		0,
		0)
	return ret1 != 0
}

func ImmUnlockIMC(hIMC HIMC) bool {
	ret1 := syscall3(immUnlockIMC, 1,
		uintptr(hIMC),
		0,
		0)
	return ret1 != 0
}

func ImmUnlockIMCC(imcc HIMCC) bool {
	ret1 := syscall3(immUnlockIMCC, 1,
		uintptr(imcc),
		0,
		0)
	return ret1 != 0
}
