// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libuxtheme uintptr

	// Functions
	beginPanningFeedback                  uintptr
	endPanningFeedback                    uintptr
	updatePanningFeedback                 uintptr
	beginBufferedAnimation                uintptr
	beginBufferedPaint                    uintptr
	bufferedPaintClear                    uintptr
	bufferedPaintInit                     uintptr
	bufferedPaintRenderAnimation          uintptr
	bufferedPaintSetAlpha                 uintptr
	bufferedPaintStopAllAnimations        uintptr
	bufferedPaintUnInit                   uintptr
	closeThemeData                        uintptr
	drawThemeBackground                   uintptr
	drawThemeBackgroundEx                 uintptr
	drawThemeEdge                         uintptr
	drawThemeIcon                         uintptr
	drawThemeParentBackground             uintptr
	drawThemeText                         uintptr
	drawThemeTextEx                       uintptr
	enableThemeDialogTexture              uintptr
	enableTheming                         uintptr
	endBufferedAnimation                  uintptr
	endBufferedPaint                      uintptr
	getBufferedPaintBits                  uintptr
	getBufferedPaintDC                    uintptr
	getBufferedPaintTargetDC              uintptr
	getBufferedPaintTargetRect            uintptr
	getCurrentThemeName                   uintptr
	getThemeAppProperties                 uintptr
	getThemeBackgroundContentRect         uintptr
	getThemeBackgroundExtent              uintptr
	getThemeBackgroundRegion              uintptr
	getThemeBool                          uintptr
	getThemeColor                         uintptr
	getThemeDocumentationProperty         uintptr
	getThemeEnumValue                     uintptr
	getThemeFilename                      uintptr
	getThemeFont                          uintptr
	getThemeInt                           uintptr
	getThemeIntList                       uintptr
	getThemeMargins                       uintptr
	getThemeMetric                        uintptr
	getThemePartSize                      uintptr
	getThemePosition                      uintptr
	getThemePropertyOrigin                uintptr
	getThemeRect                          uintptr
	getThemeString                        uintptr
	getThemeSysBool                       uintptr
	getThemeSysColor                      uintptr
	getThemeSysColorBrush                 uintptr
	getThemeSysFont                       uintptr
	getThemeSysInt                        uintptr
	getThemeSysSize                       uintptr
	getThemeSysString                     uintptr
	getThemeTextExtent                    uintptr
	getThemeTextMetrics                   uintptr
	getThemeTransitionDuration            uintptr
	getWindowTheme                        uintptr
	hitTestThemeBackground                uintptr
	isAppThemed                           uintptr
	isThemeActive                         uintptr
	isThemeBackgroundPartiallyTransparent uintptr
	isThemeDialogTextureEnabled           uintptr
	isThemePartDefined                    uintptr
	openThemeData                         uintptr
	openThemeDataEx                       uintptr
	setThemeAppProperties                 uintptr
	setWindowTheme                        uintptr
)

func init() {
	// Library
	libuxtheme = doLoadLibrary("uxtheme.dll")

	// Functions
	beginPanningFeedback = doGetProcAddress(libuxtheme, "BeginPanningFeedback")
	endPanningFeedback = doGetProcAddress(libuxtheme, "EndPanningFeedback")
	updatePanningFeedback = doGetProcAddress(libuxtheme, "UpdatePanningFeedback")
	beginBufferedAnimation = doGetProcAddress(libuxtheme, "BeginBufferedAnimation")
	beginBufferedPaint = doGetProcAddress(libuxtheme, "BeginBufferedPaint")
	bufferedPaintClear = doGetProcAddress(libuxtheme, "BufferedPaintClear")
	bufferedPaintInit = doGetProcAddress(libuxtheme, "BufferedPaintInit")
	bufferedPaintRenderAnimation = doGetProcAddress(libuxtheme, "BufferedPaintRenderAnimation")
	bufferedPaintSetAlpha = doGetProcAddress(libuxtheme, "BufferedPaintSetAlpha")
	bufferedPaintStopAllAnimations = doGetProcAddress(libuxtheme, "BufferedPaintStopAllAnimations")
	bufferedPaintUnInit = doGetProcAddress(libuxtheme, "BufferedPaintUnInit")
	closeThemeData = doGetProcAddress(libuxtheme, "CloseThemeData")
	drawThemeBackground = doGetProcAddress(libuxtheme, "DrawThemeBackground")
	drawThemeBackgroundEx = doGetProcAddress(libuxtheme, "DrawThemeBackgroundEx")
	drawThemeEdge = doGetProcAddress(libuxtheme, "DrawThemeEdge")
	drawThemeIcon = doGetProcAddress(libuxtheme, "DrawThemeIcon")
	drawThemeParentBackground = doGetProcAddress(libuxtheme, "DrawThemeParentBackground")
	drawThemeText = doGetProcAddress(libuxtheme, "DrawThemeText")
	drawThemeTextEx = doGetProcAddress(libuxtheme, "DrawThemeTextEx")
	enableThemeDialogTexture = doGetProcAddress(libuxtheme, "EnableThemeDialogTexture")
	enableTheming = doGetProcAddress(libuxtheme, "EnableTheming")
	endBufferedAnimation = doGetProcAddress(libuxtheme, "EndBufferedAnimation")
	endBufferedPaint = doGetProcAddress(libuxtheme, "EndBufferedPaint")
	getBufferedPaintBits = doGetProcAddress(libuxtheme, "GetBufferedPaintBits")
	getBufferedPaintDC = doGetProcAddress(libuxtheme, "GetBufferedPaintDC")
	getBufferedPaintTargetDC = doGetProcAddress(libuxtheme, "GetBufferedPaintTargetDC")
	getBufferedPaintTargetRect = doGetProcAddress(libuxtheme, "GetBufferedPaintTargetRect")
	getCurrentThemeName = doGetProcAddress(libuxtheme, "GetCurrentThemeName")
	getThemeAppProperties = doGetProcAddress(libuxtheme, "GetThemeAppProperties")
	getThemeBackgroundContentRect = doGetProcAddress(libuxtheme, "GetThemeBackgroundContentRect")
	getThemeBackgroundExtent = doGetProcAddress(libuxtheme, "GetThemeBackgroundExtent")
	getThemeBackgroundRegion = doGetProcAddress(libuxtheme, "GetThemeBackgroundRegion")
	getThemeBool = doGetProcAddress(libuxtheme, "GetThemeBool")
	getThemeColor = doGetProcAddress(libuxtheme, "GetThemeColor")
	getThemeDocumentationProperty = doGetProcAddress(libuxtheme, "GetThemeDocumentationProperty")
	getThemeEnumValue = doGetProcAddress(libuxtheme, "GetThemeEnumValue")
	getThemeFilename = doGetProcAddress(libuxtheme, "GetThemeFilename")
	getThemeFont = doGetProcAddress(libuxtheme, "GetThemeFont")
	getThemeInt = doGetProcAddress(libuxtheme, "GetThemeInt")
	getThemeIntList = doGetProcAddress(libuxtheme, "GetThemeIntList")
	getThemeMargins = doGetProcAddress(libuxtheme, "GetThemeMargins")
	getThemeMetric = doGetProcAddress(libuxtheme, "GetThemeMetric")
	getThemePartSize = doGetProcAddress(libuxtheme, "GetThemePartSize")
	getThemePosition = doGetProcAddress(libuxtheme, "GetThemePosition")
	getThemePropertyOrigin = doGetProcAddress(libuxtheme, "GetThemePropertyOrigin")
	getThemeRect = doGetProcAddress(libuxtheme, "GetThemeRect")
	getThemeString = doGetProcAddress(libuxtheme, "GetThemeString")
	getThemeSysBool = doGetProcAddress(libuxtheme, "GetThemeSysBool")
	getThemeSysColor = doGetProcAddress(libuxtheme, "GetThemeSysColor")
	getThemeSysColorBrush = doGetProcAddress(libuxtheme, "GetThemeSysColorBrush")
	getThemeSysFont = doGetProcAddress(libuxtheme, "GetThemeSysFont")
	getThemeSysInt = doGetProcAddress(libuxtheme, "GetThemeSysInt")
	getThemeSysSize = doGetProcAddress(libuxtheme, "GetThemeSysSize")
	getThemeSysString = doGetProcAddress(libuxtheme, "GetThemeSysString")
	getThemeTextExtent = doGetProcAddress(libuxtheme, "GetThemeTextExtent")
	getThemeTextMetrics = doGetProcAddress(libuxtheme, "GetThemeTextMetrics")
	getThemeTransitionDuration = doGetProcAddress(libuxtheme, "GetThemeTransitionDuration")
	getWindowTheme = doGetProcAddress(libuxtheme, "GetWindowTheme")
	hitTestThemeBackground = doGetProcAddress(libuxtheme, "HitTestThemeBackground")
	isAppThemed = doGetProcAddress(libuxtheme, "IsAppThemed")
	isThemeActive = doGetProcAddress(libuxtheme, "IsThemeActive")
	isThemeBackgroundPartiallyTransparent = doGetProcAddress(libuxtheme, "IsThemeBackgroundPartiallyTransparent")
	isThemeDialogTextureEnabled = doGetProcAddress(libuxtheme, "IsThemeDialogTextureEnabled")
	isThemePartDefined = doGetProcAddress(libuxtheme, "IsThemePartDefined")
	openThemeData = doGetProcAddress(libuxtheme, "OpenThemeData")
	openThemeDataEx = doGetProcAddress(libuxtheme, "OpenThemeDataEx")
	setThemeAppProperties = doGetProcAddress(libuxtheme, "SetThemeAppProperties")
	setWindowTheme = doGetProcAddress(libuxtheme, "SetWindowTheme")
}

func BeginPanningFeedback(hwnd HWND) bool {
	ret1 := syscall3(beginPanningFeedback, 1,
		uintptr(hwnd),
		0,
		0)
	return ret1 != 0
}

func EndPanningFeedback(hwnd HWND, fAnimateBack bool) bool {
	ret1 := syscall3(endPanningFeedback, 2,
		uintptr(hwnd),
		getUintptrFromBool(fAnimateBack),
		0)
	return ret1 != 0
}

func UpdatePanningFeedback(hwnd HWND, lTotalOverpanOffsetX LONG, lTotalOverpanOffsetY LONG, fInInertia bool) bool {
	ret1 := syscall6(updatePanningFeedback, 4,
		uintptr(hwnd),
		uintptr(lTotalOverpanOffsetX),
		uintptr(lTotalOverpanOffsetY),
		getUintptrFromBool(fInInertia),
		0,
		0)
	return ret1 != 0
}

func BeginBufferedAnimation(hwnd HWND, hdcTarget HDC, rcTarget /*const*/ *RECT, dwFormat BP_BUFFERFORMAT, pPaintParams *BP_PAINTPARAMS, pAnimationParams *BP_ANIMATIONPARAMS, phdcFrom *HDC, phdcTo *HDC) HANIMATIONBUFFER {
	ret1 := syscall9(beginBufferedAnimation, 8,
		uintptr(hwnd),
		uintptr(hdcTarget),
		uintptr(unsafe.Pointer(rcTarget)),
		uintptr(dwFormat),
		uintptr(unsafe.Pointer(pPaintParams)),
		uintptr(unsafe.Pointer(pAnimationParams)),
		uintptr(unsafe.Pointer(phdcFrom)),
		uintptr(unsafe.Pointer(phdcTo)),
		0)
	return HANIMATIONBUFFER(ret1)
}

func BeginBufferedPaint(hdcTarget HDC, prcTarget /*const*/ *RECT, dwFormat BP_BUFFERFORMAT, pPaintParams *BP_PAINTPARAMS, phdc *HDC) HPAINTBUFFER {
	ret1 := syscall6(beginBufferedPaint, 5,
		uintptr(hdcTarget),
		uintptr(unsafe.Pointer(prcTarget)),
		uintptr(dwFormat),
		uintptr(unsafe.Pointer(pPaintParams)),
		uintptr(unsafe.Pointer(phdc)),
		0)
	return HPAINTBUFFER(ret1)
}

func BufferedPaintClear(hBufferedPaint HPAINTBUFFER, prc /*const*/ *RECT) HRESULT {
	ret1 := syscall3(bufferedPaintClear, 2,
		uintptr(hBufferedPaint),
		uintptr(unsafe.Pointer(prc)),
		0)
	return HRESULT(ret1)
}

func BufferedPaintInit() HRESULT {
	ret1 := syscall3(bufferedPaintInit, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func BufferedPaintRenderAnimation(hwnd HWND, hdcTarget HDC) bool {
	ret1 := syscall3(bufferedPaintRenderAnimation, 2,
		uintptr(hwnd),
		uintptr(hdcTarget),
		0)
	return ret1 != 0
}

func BufferedPaintSetAlpha(hBufferedPaint HPAINTBUFFER, prc /*const*/ *RECT, alpha byte) HRESULT {
	ret1 := syscall3(bufferedPaintSetAlpha, 3,
		uintptr(hBufferedPaint),
		uintptr(unsafe.Pointer(prc)),
		uintptr(alpha))
	return HRESULT(ret1)
}

func BufferedPaintStopAllAnimations(hwnd HWND) HRESULT {
	ret1 := syscall3(bufferedPaintStopAllAnimations, 1,
		uintptr(hwnd),
		0,
		0)
	return HRESULT(ret1)
}

func BufferedPaintUnInit() HRESULT {
	ret1 := syscall3(bufferedPaintUnInit, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func CloseThemeData(hTheme HTHEME) HRESULT {
	ret1 := syscall3(closeThemeData, 1,
		uintptr(hTheme),
		0,
		0)
	return HRESULT(ret1)
}

func DrawThemeBackground(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pRect /*const*/ *RECT, pClipRect /*const*/ *RECT) HRESULT {
	ret1 := syscall6(drawThemeBackground, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pClipRect)))
	return HRESULT(ret1)
}

func DrawThemeBackgroundEx(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pRect /*const*/ *RECT, pOptions /*const*/ *DTBGOPTS) HRESULT {
	ret1 := syscall6(drawThemeBackgroundEx, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pOptions)))
	return HRESULT(ret1)
}

func DrawThemeEdge(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pDestRect /*const*/ *RECT, uEdge uint32, uFlags uint32, pContentRect *RECT) HRESULT {
	ret1 := syscall9(drawThemeEdge, 8,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pDestRect)),
		uintptr(uEdge),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(pContentRect)),
		0)
	return HRESULT(ret1)
}

func DrawThemeIcon(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pRect /*const*/ *RECT, himl HIMAGELIST, iImageIndex int32) HRESULT {
	ret1 := syscall9(drawThemeIcon, 7,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(himl),
		uintptr(iImageIndex),
		0,
		0)
	return HRESULT(ret1)
}

func DrawThemeParentBackground(hwnd HWND, hdc HDC, prc *RECT) HRESULT {
	ret1 := syscall3(drawThemeParentBackground, 3,
		uintptr(hwnd),
		uintptr(hdc),
		uintptr(unsafe.Pointer(prc)))
	return HRESULT(ret1)
}

func DrawThemeText(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pszText string, iCharCount int32, flags uint32, flags2 uint32, pRect /*const*/ *RECT) HRESULT {
	pszTextStr := unicode16FromString(pszText)
	ret1 := syscall9(drawThemeText, 9,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(&pszTextStr[0])),
		uintptr(iCharCount),
		uintptr(flags),
		uintptr(flags2),
		uintptr(unsafe.Pointer(pRect)))
	return HRESULT(ret1)
}

func DrawThemeTextEx(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pszText string, iCharCount int32, flags uint32, rect *RECT, options /*const*/ *DTTOPTS) HRESULT {
	pszTextStr := unicode16FromString(pszText)
	ret1 := syscall9(drawThemeTextEx, 9,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(&pszTextStr[0])),
		uintptr(iCharCount),
		uintptr(flags),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(options)))
	return HRESULT(ret1)
}

func EnableThemeDialogTexture(hwnd HWND, dwFlags uint32) HRESULT {
	ret1 := syscall3(enableThemeDialogTexture, 2,
		uintptr(hwnd),
		uintptr(dwFlags),
		0)
	return HRESULT(ret1)
}

func EnableTheming(fEnable bool) HRESULT {
	ret1 := syscall3(enableTheming, 1,
		getUintptrFromBool(fEnable),
		0,
		0)
	return HRESULT(ret1)
}

func EndBufferedAnimation(hbpAnimation HANIMATIONBUFFER, fUpdateTarget bool) HRESULT {
	ret1 := syscall3(endBufferedAnimation, 2,
		uintptr(hbpAnimation),
		getUintptrFromBool(fUpdateTarget),
		0)
	return HRESULT(ret1)
}

func EndBufferedPaint(hPaintBuffer HPAINTBUFFER, fUpdateTarget bool) HRESULT {
	ret1 := syscall3(endBufferedPaint, 2,
		uintptr(hPaintBuffer),
		getUintptrFromBool(fUpdateTarget),
		0)
	return HRESULT(ret1)
}

func GetBufferedPaintBits(hBufferedPaint HPAINTBUFFER, ppbBuffer **RGBQUAD, pcxRow *int) HRESULT {
	ret1 := syscall3(getBufferedPaintBits, 3,
		uintptr(hBufferedPaint),
		uintptr(unsafe.Pointer(ppbBuffer)),
		uintptr(unsafe.Pointer(pcxRow)))
	return HRESULT(ret1)
}

func GetBufferedPaintDC(hBufferedPaint HPAINTBUFFER) HDC {
	ret1 := syscall3(getBufferedPaintDC, 1,
		uintptr(hBufferedPaint),
		0,
		0)
	return HDC(ret1)
}

func GetBufferedPaintTargetDC(hBufferedPaint HPAINTBUFFER) HDC {
	ret1 := syscall3(getBufferedPaintTargetDC, 1,
		uintptr(hBufferedPaint),
		0,
		0)
	return HDC(ret1)
}

func GetBufferedPaintTargetRect(hBufferedPaint HPAINTBUFFER, prc *RECT) HRESULT {
	ret1 := syscall3(getBufferedPaintTargetRect, 2,
		uintptr(hBufferedPaint),
		uintptr(unsafe.Pointer(prc)),
		0)
	return HRESULT(ret1)
}

func GetCurrentThemeName(pszThemeFileName LPWSTR, dwMaxNameChars int32, pszColorBuff LPWSTR, cchMaxColorChars int32, pszSizeBuff LPWSTR, cchMaxSizeChars int32) HRESULT {
	ret1 := syscall6(getCurrentThemeName, 6,
		uintptr(unsafe.Pointer(pszThemeFileName)),
		uintptr(dwMaxNameChars),
		uintptr(unsafe.Pointer(pszColorBuff)),
		uintptr(cchMaxColorChars),
		uintptr(unsafe.Pointer(pszSizeBuff)),
		uintptr(cchMaxSizeChars))
	return HRESULT(ret1)
}

func GetThemeAppProperties() uint32 {
	ret1 := syscall3(getThemeAppProperties, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetThemeBackgroundContentRect(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pBoundingRect /*const*/ *RECT, pContentRect *RECT) HRESULT {
	ret1 := syscall6(getThemeBackgroundContentRect, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pBoundingRect)),
		uintptr(unsafe.Pointer(pContentRect)))
	return HRESULT(ret1)
}

func GetThemeBackgroundExtent(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pContentRect /*const*/ *RECT, pExtentRect *RECT) HRESULT {
	ret1 := syscall6(getThemeBackgroundExtent, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pContentRect)),
		uintptr(unsafe.Pointer(pExtentRect)))
	return HRESULT(ret1)
}

func GetThemeBackgroundRegion(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pRect /*const*/ *RECT, pRegion *HRGN) HRESULT {
	ret1 := syscall6(getThemeBackgroundRegion, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pRegion)))
	return HRESULT(ret1)
}

func GetThemeBool(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pfVal *BOOL) HRESULT {
	ret1 := syscall6(getThemeBool, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pfVal)),
		0)
	return HRESULT(ret1)
}

func GetThemeColor(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pColor *COLORREF) HRESULT {
	ret1 := syscall6(getThemeColor, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pColor)),
		0)
	return HRESULT(ret1)
}

func GetThemeDocumentationProperty(pszThemeName string, pszPropertyName string, pszValueBuff LPWSTR, cchMaxValChars int32) HRESULT {
	pszThemeNameStr := unicode16FromString(pszThemeName)
	pszPropertyNameStr := unicode16FromString(pszPropertyName)
	ret1 := syscall6(getThemeDocumentationProperty, 4,
		uintptr(unsafe.Pointer(&pszThemeNameStr[0])),
		uintptr(unsafe.Pointer(&pszPropertyNameStr[0])),
		uintptr(unsafe.Pointer(pszValueBuff)),
		uintptr(cchMaxValChars),
		0,
		0)
	return HRESULT(ret1)
}

func GetThemeEnumValue(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, piVal *int) HRESULT {
	ret1 := syscall6(getThemeEnumValue, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(piVal)),
		0)
	return HRESULT(ret1)
}

func GetThemeFilename(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pszThemeFilename LPWSTR, cchMaxBuffChars int32) HRESULT {
	ret1 := syscall6(getThemeFilename, 6,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pszThemeFilename)),
		uintptr(cchMaxBuffChars))
	return HRESULT(ret1)
}

func GetThemeFont(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, iPropId int32, pFont *LOGFONT) HRESULT {
	ret1 := syscall6(getThemeFont, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pFont)))
	return HRESULT(ret1)
}

func GetThemeInt(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, piVal *int) HRESULT {
	ret1 := syscall6(getThemeInt, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(piVal)),
		0)
	return HRESULT(ret1)
}

func GetThemeIntList(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pIntList *INTLIST) HRESULT {
	ret1 := syscall6(getThemeIntList, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pIntList)),
		0)
	return HRESULT(ret1)
}

func GetThemeMargins(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, iPropId int32, prc *RECT, pMargins *MARGINS) HRESULT {
	ret1 := syscall9(getThemeMargins, 7,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(prc)),
		uintptr(unsafe.Pointer(pMargins)),
		0,
		0)
	return HRESULT(ret1)
}

func GetThemeMetric(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, iPropId int32, piVal *int) HRESULT {
	ret1 := syscall6(getThemeMetric, 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(piVal)))
	return HRESULT(ret1)
}

func GetThemePartSize(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, prc *RECT, eSize THEMESIZE, psz *SIZE) HRESULT {
	ret1 := syscall9(getThemePartSize, 7,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(prc)),
		uintptr(eSize),
		uintptr(unsafe.Pointer(psz)),
		0,
		0)
	return HRESULT(ret1)
}

func GetThemePosition(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pPoint *POINT) HRESULT {
	ret1 := syscall6(getThemePosition, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pPoint)),
		0)
	return HRESULT(ret1)
}

func GetThemePropertyOrigin(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pOrigin *PROPERTYORIGIN) HRESULT {
	ret1 := syscall6(getThemePropertyOrigin, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pOrigin)),
		0)
	return HRESULT(ret1)
}

func GetThemeRect(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pRect *RECT) HRESULT {
	ret1 := syscall6(getThemeRect, 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pRect)),
		0)
	return HRESULT(ret1)
}

func GetThemeString(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pszBuff LPWSTR, cchMaxBuffChars int32) HRESULT {
	ret1 := syscall6(getThemeString, 6,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pszBuff)),
		uintptr(cchMaxBuffChars))
	return HRESULT(ret1)
}

func GetThemeSysBool(hTheme HTHEME, iBoolID int32) bool {
	ret1 := syscall3(getThemeSysBool, 2,
		uintptr(hTheme),
		uintptr(iBoolID),
		0)
	return ret1 != 0
}

func GetThemeSysColor(hTheme HTHEME, iColorID int32) COLORREF {
	ret1 := syscall3(getThemeSysColor, 2,
		uintptr(hTheme),
		uintptr(iColorID),
		0)
	return COLORREF(ret1)
}

func GetThemeSysColorBrush(hTheme HTHEME, iColorID int32) HBRUSH {
	ret1 := syscall3(getThemeSysColorBrush, 2,
		uintptr(hTheme),
		uintptr(iColorID),
		0)
	return HBRUSH(ret1)
}

func GetThemeSysFont(hTheme HTHEME, iFontID int32, plf *LOGFONT) HRESULT {
	ret1 := syscall3(getThemeSysFont, 3,
		uintptr(hTheme),
		uintptr(iFontID),
		uintptr(unsafe.Pointer(plf)))
	return HRESULT(ret1)
}

func GetThemeSysInt(hTheme HTHEME, iIntID int32, piValue *int) HRESULT {
	ret1 := syscall3(getThemeSysInt, 3,
		uintptr(hTheme),
		uintptr(iIntID),
		uintptr(unsafe.Pointer(piValue)))
	return HRESULT(ret1)
}

func GetThemeSysSize(hTheme HTHEME, iSizeID int32) int32 {
	ret1 := syscall3(getThemeSysSize, 2,
		uintptr(hTheme),
		uintptr(iSizeID),
		0)
	return int32(ret1)
}

func GetThemeSysString(hTheme HTHEME, iStringID int32, pszStringBuff LPWSTR, cchMaxStringChars int32) HRESULT {
	ret1 := syscall6(getThemeSysString, 4,
		uintptr(hTheme),
		uintptr(iStringID),
		uintptr(unsafe.Pointer(pszStringBuff)),
		uintptr(cchMaxStringChars),
		0,
		0)
	return HRESULT(ret1)
}

func GetThemeTextExtent(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pszText string, iCharCount int32, dwTextFlags uint32, pBoundingRect /*const*/ *RECT, pExtentRect *RECT) HRESULT {
	pszTextStr := unicode16FromString(pszText)
	ret1 := syscall9(getThemeTextExtent, 9,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(&pszTextStr[0])),
		uintptr(iCharCount),
		uintptr(dwTextFlags),
		uintptr(unsafe.Pointer(pBoundingRect)),
		uintptr(unsafe.Pointer(pExtentRect)))
	return HRESULT(ret1)
}

func GetThemeTextMetrics(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, ptm *TEXTMETRIC) HRESULT {
	ret1 := syscall6(getThemeTextMetrics, 5,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(ptm)),
		0)
	return HRESULT(ret1)
}

func GetThemeTransitionDuration(hTheme HTHEME, iPartId int32, iStateIdFrom int32, iStateIdTo int32, iPropId int32, pdwDuration *uint32) HRESULT {
	ret1 := syscall6(getThemeTransitionDuration, 6,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateIdFrom),
		uintptr(iStateIdTo),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pdwDuration)))
	return HRESULT(ret1)
}

func GetWindowTheme(hwnd HWND) HTHEME {
	ret1 := syscall3(getWindowTheme, 1,
		uintptr(hwnd),
		0,
		0)
	return HTHEME(ret1)
}

func HitTestThemeBackground(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, dwOptions uint32, pRect /*const*/ *RECT, hrgn HRGN, ptTest POINT, pwHitTestCode *WORD) HRESULT {
	ret1 := syscall12(hitTestThemeBackground, 10,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(dwOptions),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(hrgn),
		uintptr(ptTest.X),
		uintptr(ptTest.Y),
		uintptr(unsafe.Pointer(pwHitTestCode)),
		0,
		0)
	return HRESULT(ret1)
}

func IsAppThemed() bool {
	ret1 := syscall3(isAppThemed, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func IsThemeActive() bool {
	ret1 := syscall3(isThemeActive, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func IsThemeBackgroundPartiallyTransparent(hTheme HTHEME, iPartId int32, iStateId int32) bool {
	ret1 := syscall3(isThemeBackgroundPartiallyTransparent, 3,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId))
	return ret1 != 0
}

func IsThemeDialogTextureEnabled(hwnd HWND) bool {
	ret1 := syscall3(isThemeDialogTextureEnabled, 1,
		uintptr(hwnd),
		0,
		0)
	return ret1 != 0
}

func IsThemePartDefined(hTheme HTHEME, iPartId int32, iStateId int32) bool {
	ret1 := syscall3(isThemePartDefined, 3,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId))
	return ret1 != 0
}

func OpenThemeData(hwnd HWND, classlist string) HTHEME {
	classlistStr := unicode16FromString(classlist)
	ret1 := syscall3(openThemeData, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&classlistStr[0])),
		0)
	return HTHEME(ret1)
}

func OpenThemeDataEx(hwnd HWND, pszClassList string, flags uint32) HTHEME {
	pszClassListStr := unicode16FromString(pszClassList)
	ret1 := syscall3(openThemeDataEx, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pszClassListStr[0])),
		uintptr(flags))
	return HTHEME(ret1)
}

func SetThemeAppProperties(dwFlags uint32) {
	syscall3(setThemeAppProperties, 1,
		uintptr(dwFlags),
		0,
		0)
}

func SetWindowTheme(hwnd HWND, pszSubAppName string, pszSubIdList string) HRESULT {
	pszSubAppNameStr := unicode16FromString(pszSubAppName)
	pszSubIdListStr := unicode16FromString(pszSubIdList)
	ret1 := syscall3(setWindowTheme, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pszSubAppNameStr[0])),
		uintptr(unsafe.Pointer(&pszSubIdListStr[0])))
	return HRESULT(ret1)
}
