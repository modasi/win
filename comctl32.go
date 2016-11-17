// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libcomctl32 uintptr

	// Functions
	createMappedBitmap           uintptr
	createPropertySheetPage      uintptr
	createStatusWindow           uintptr
	createToolbarEx              uintptr
	createUpDownControl          uintptr
	dPA_Create                   uintptr
	dPA_DeleteAllPtrs            uintptr
	dPA_DeletePtr                uintptr
	dPA_Destroy                  uintptr
	dPA_DestroyCallback          uintptr
	dPA_EnumCallback             uintptr
	dPA_GetPtr                   uintptr
	dPA_InsertPtr                uintptr
	dPA_Search                   uintptr
	dPA_SetPtr                   uintptr
	dPA_Sort                     uintptr
	dSA_Create                   uintptr
	dSA_DeleteAllItems           uintptr
	dSA_Destroy                  uintptr
	dSA_DestroyCallback          uintptr
	dSA_GetItemPtr               uintptr
	dSA_InsertItem               uintptr
	defSubclassProc              uintptr
	destroyPropertySheetPage     uintptr
	drawInsert                   uintptr
	drawStatusText               uintptr
	flatSB_EnableScrollBar       uintptr
	flatSB_GetScrollInfo         uintptr
	flatSB_GetScrollPos          uintptr
	flatSB_GetScrollProp         uintptr
	flatSB_GetScrollPropPtr      uintptr
	flatSB_GetScrollRange        uintptr
	flatSB_SetScrollInfo         uintptr
	flatSB_SetScrollPos          uintptr
	flatSB_SetScrollProp         uintptr
	flatSB_SetScrollRange        uintptr
	flatSB_ShowScrollBar         uintptr
	getEffectiveClientRect       uintptr
	getMUILanguage               uintptr
	imageList_Add                uintptr
	imageList_AddMasked          uintptr
	imageList_BeginDrag          uintptr
	imageList_Copy               uintptr
	imageList_Create             uintptr
	imageList_Destroy            uintptr
	imageList_DragEnter          uintptr
	imageList_DragLeave          uintptr
	imageList_DragMove           uintptr
	imageList_DragShowNolock     uintptr
	imageList_Draw               uintptr
	imageList_DrawEx             uintptr
	imageList_DrawIndirect       uintptr
	imageList_Duplicate          uintptr
	imageList_EndDrag            uintptr
	imageList_GetBkColor         uintptr
	imageList_GetDragImage       uintptr
	imageList_GetIcon            uintptr
	imageList_GetIconSize        uintptr
	imageList_GetImageCount      uintptr
	imageList_GetImageInfo       uintptr
	imageList_LoadImage          uintptr
	imageList_Merge              uintptr
	imageList_Read               uintptr
	imageList_Remove             uintptr
	imageList_Replace            uintptr
	imageList_ReplaceIcon        uintptr
	imageList_SetBkColor         uintptr
	imageList_SetDragCursorImage uintptr
	imageList_SetIconSize        uintptr
	imageList_SetImageCount      uintptr
	imageList_SetOverlayImage    uintptr
	imageList_Write              uintptr
	initCommonControls           uintptr
	initCommonControlsEx         uintptr
	initMUILanguage              uintptr
	initializeFlatSB             uintptr
	lBItemFromPt                 uintptr
	makeDragList                 uintptr
	menuHelp                     uintptr
	propertySheet                uintptr
	removeWindowSubclass         uintptr
	setWindowSubclass            uintptr
	showHideMenuCtl              uintptr
	str_SetPtrW                  uintptr
	taskDialog                   uintptr
	taskDialogIndirect           uintptr
	uninitializeFlatSB           uintptr
	addMRUStringW                uintptr
	createMRUListW               uintptr
	createToolbar                uintptr
	enumMRUListW                 uintptr
	freeMRUList                  uintptr
	imageList_AddIcon            uintptr
	imageList_GetFlags           uintptr
	imageList_GetImageRect       uintptr
	imageList_SetFilter          uintptr
	imageList_SetFlags           uintptr
)

func init() {
	// Library
	libcomctl32 = doLoadLibrary("comctl32.dll")

	// Functions
	createMappedBitmap = doGetProcAddress(libcomctl32, "CreateMappedBitmap")
	createPropertySheetPage = doGetProcAddress(libcomctl32, "CreatePropertySheetPageW")
	createStatusWindow = doGetProcAddress(libcomctl32, "CreateStatusWindowW")
	createToolbarEx = doGetProcAddress(libcomctl32, "CreateToolbarEx")
	createUpDownControl = doGetProcAddress(libcomctl32, "CreateUpDownControl")
	dPA_Create = doGetProcAddress(libcomctl32, "DPA_Create")
	dPA_DeleteAllPtrs = doGetProcAddress(libcomctl32, "DPA_DeleteAllPtrs")
	dPA_DeletePtr = doGetProcAddress(libcomctl32, "DPA_DeletePtr")
	dPA_Destroy = doGetProcAddress(libcomctl32, "DPA_Destroy")
	dPA_DestroyCallback = doGetProcAddress(libcomctl32, "DPA_DestroyCallback")
	dPA_EnumCallback = doGetProcAddress(libcomctl32, "DPA_EnumCallback")
	dPA_GetPtr = doGetProcAddress(libcomctl32, "DPA_GetPtr")
	dPA_InsertPtr = doGetProcAddress(libcomctl32, "DPA_InsertPtr")
	dPA_Search = doGetProcAddress(libcomctl32, "DPA_Search")
	dPA_SetPtr = doGetProcAddress(libcomctl32, "DPA_SetPtr")
	dPA_Sort = doGetProcAddress(libcomctl32, "DPA_Sort")
	dSA_Create = doGetProcAddress(libcomctl32, "DSA_Create")
	dSA_DeleteAllItems = doGetProcAddress(libcomctl32, "DSA_DeleteAllItems")
	dSA_Destroy = doGetProcAddress(libcomctl32, "DSA_Destroy")
	dSA_DestroyCallback = doGetProcAddress(libcomctl32, "DSA_DestroyCallback")
	dSA_GetItemPtr = doGetProcAddress(libcomctl32, "DSA_GetItemPtr")
	dSA_InsertItem = doGetProcAddress(libcomctl32, "DSA_InsertItem")
	defSubclassProc = doGetProcAddress(libcomctl32, "DefSubclassProc")
	destroyPropertySheetPage = doGetProcAddress(libcomctl32, "DestroyPropertySheetPage")
	drawInsert = doGetProcAddress(libcomctl32, "DrawInsert")
	drawStatusText = doGetProcAddress(libcomctl32, "DrawStatusTextW")
	flatSB_EnableScrollBar = doGetProcAddress(libcomctl32, "FlatSB_EnableScrollBar")
	flatSB_GetScrollInfo = doGetProcAddress(libcomctl32, "FlatSB_GetScrollInfo")
	flatSB_GetScrollPos = doGetProcAddress(libcomctl32, "FlatSB_GetScrollPos")
	flatSB_GetScrollProp = doGetProcAddress(libcomctl32, "FlatSB_GetScrollProp")
	flatSB_GetScrollPropPtr = doGetProcAddress(libcomctl32, "FlatSB_GetScrollPropPtr")
	flatSB_GetScrollRange = doGetProcAddress(libcomctl32, "FlatSB_GetScrollRange")
	flatSB_SetScrollInfo = doGetProcAddress(libcomctl32, "FlatSB_SetScrollInfo")
	flatSB_SetScrollPos = doGetProcAddress(libcomctl32, "FlatSB_SetScrollPos")
	flatSB_SetScrollProp = doGetProcAddress(libcomctl32, "FlatSB_SetScrollProp")
	flatSB_SetScrollRange = doGetProcAddress(libcomctl32, "FlatSB_SetScrollRange")
	flatSB_ShowScrollBar = doGetProcAddress(libcomctl32, "FlatSB_ShowScrollBar")
	getEffectiveClientRect = doGetProcAddress(libcomctl32, "GetEffectiveClientRect")
	getMUILanguage = doGetProcAddress(libcomctl32, "GetMUILanguage")
	imageList_Add = doGetProcAddress(libcomctl32, "ImageList_Add")
	imageList_AddMasked = doGetProcAddress(libcomctl32, "ImageList_AddMasked")
	imageList_BeginDrag = doGetProcAddress(libcomctl32, "ImageList_BeginDrag")
	imageList_Copy = doGetProcAddress(libcomctl32, "ImageList_Copy")
	imageList_Create = doGetProcAddress(libcomctl32, "ImageList_Create")
	imageList_Destroy = doGetProcAddress(libcomctl32, "ImageList_Destroy")
	imageList_DragEnter = doGetProcAddress(libcomctl32, "ImageList_DragEnter")
	imageList_DragLeave = doGetProcAddress(libcomctl32, "ImageList_DragLeave")
	imageList_DragMove = doGetProcAddress(libcomctl32, "ImageList_DragMove")
	imageList_DragShowNolock = doGetProcAddress(libcomctl32, "ImageList_DragShowNolock")
	imageList_Draw = doGetProcAddress(libcomctl32, "ImageList_Draw")
	imageList_DrawEx = doGetProcAddress(libcomctl32, "ImageList_DrawEx")
	imageList_DrawIndirect = doGetProcAddress(libcomctl32, "ImageList_DrawIndirect")
	imageList_Duplicate = doGetProcAddress(libcomctl32, "ImageList_Duplicate")
	imageList_EndDrag = doGetProcAddress(libcomctl32, "ImageList_EndDrag")
	imageList_GetBkColor = doGetProcAddress(libcomctl32, "ImageList_GetBkColor")
	imageList_GetDragImage = doGetProcAddress(libcomctl32, "ImageList_GetDragImage")
	imageList_GetIcon = doGetProcAddress(libcomctl32, "ImageList_GetIcon")
	imageList_GetIconSize = doGetProcAddress(libcomctl32, "ImageList_GetIconSize")
	imageList_GetImageCount = doGetProcAddress(libcomctl32, "ImageList_GetImageCount")
	imageList_GetImageInfo = doGetProcAddress(libcomctl32, "ImageList_GetImageInfo")
	imageList_LoadImage = doGetProcAddress(libcomctl32, "ImageList_LoadImageW")
	imageList_Merge = doGetProcAddress(libcomctl32, "ImageList_Merge")
	imageList_Read = doGetProcAddress(libcomctl32, "ImageList_Read")
	imageList_Remove = doGetProcAddress(libcomctl32, "ImageList_Remove")
	imageList_Replace = doGetProcAddress(libcomctl32, "ImageList_Replace")
	imageList_ReplaceIcon = doGetProcAddress(libcomctl32, "ImageList_ReplaceIcon")
	imageList_SetBkColor = doGetProcAddress(libcomctl32, "ImageList_SetBkColor")
	imageList_SetDragCursorImage = doGetProcAddress(libcomctl32, "ImageList_SetDragCursorImage")
	imageList_SetIconSize = doGetProcAddress(libcomctl32, "ImageList_SetIconSize")
	imageList_SetImageCount = doGetProcAddress(libcomctl32, "ImageList_SetImageCount")
	imageList_SetOverlayImage = doGetProcAddress(libcomctl32, "ImageList_SetOverlayImage")
	imageList_Write = doGetProcAddress(libcomctl32, "ImageList_Write")
	initCommonControls = doGetProcAddress(libcomctl32, "InitCommonControls")
	initCommonControlsEx = doGetProcAddress(libcomctl32, "InitCommonControlsEx")
	initMUILanguage = doGetProcAddress(libcomctl32, "InitMUILanguage")
	initializeFlatSB = doGetProcAddress(libcomctl32, "InitializeFlatSB")
	lBItemFromPt = doGetProcAddress(libcomctl32, "LBItemFromPt")
	makeDragList = doGetProcAddress(libcomctl32, "MakeDragList")
	menuHelp = doGetProcAddress(libcomctl32, "MenuHelp")
	propertySheet = doGetProcAddress(libcomctl32, "PropertySheetW")
	removeWindowSubclass = doGetProcAddress(libcomctl32, "RemoveWindowSubclass")
	setWindowSubclass = doGetProcAddress(libcomctl32, "SetWindowSubclass")
	showHideMenuCtl = doGetProcAddress(libcomctl32, "ShowHideMenuCtl")
	str_SetPtrW = doGetProcAddress(libcomctl32, "Str_SetPtrW")
	taskDialog = doGetProcAddress(libcomctl32, "TaskDialog")
	taskDialogIndirect = doGetProcAddress(libcomctl32, "TaskDialogIndirect")
	uninitializeFlatSB = doGetProcAddress(libcomctl32, "UninitializeFlatSB")
	addMRUStringW = doGetProcAddress(libcomctl32, "AddMRUStringW")
	createMRUListW = doGetProcAddress(libcomctl32, "CreateMRUListW")
	createToolbar = doGetProcAddress(libcomctl32, "CreateToolbar")
	enumMRUListW = doGetProcAddress(libcomctl32, "EnumMRUListW")
	freeMRUList = doGetProcAddress(libcomctl32, "FreeMRUList")
	imageList_AddIcon = doGetProcAddress(libcomctl32, "ImageList_AddIcon")
	imageList_GetFlags = doGetProcAddress(libcomctl32, "ImageList_GetFlags")
	imageList_GetImageRect = doGetProcAddress(libcomctl32, "ImageList_GetImageRect")
	imageList_SetFilter = doGetProcAddress(libcomctl32, "ImageList_SetFilter")
	imageList_SetFlags = doGetProcAddress(libcomctl32, "ImageList_SetFlags")
}

func CreateMappedBitmap(hInstance HINSTANCE, idBitmap INT_PTR, wFlags uint32, lpColorMap *COLORMAP, iNumMaps int32) HBITMAP {
	ret1 := syscall6(createMappedBitmap, 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(idBitmap)),
		uintptr(wFlags),
		uintptr(unsafe.Pointer(lpColorMap)),
		uintptr(iNumMaps),
		0)
	return HBITMAP(ret1)
}

func CreatePropertySheetPage(constPropSheetPagePointer /*const*/ *PROPSHEETPAGE) HPROPSHEETPAGE {
	ret1 := syscall3(createPropertySheetPage, 1,
		uintptr(unsafe.Pointer(constPropSheetPagePointer)),
		0,
		0)
	return HPROPSHEETPAGE(ret1)
}

func CreateStatusWindow(style LONG, lpszText string, hwndParent HWND, wID uint32) HWND {
	lpszTextStr := unicode16FromString(lpszText)
	ret1 := syscall6(createStatusWindow, 4,
		uintptr(style),
		uintptr(unsafe.Pointer(&lpszTextStr[0])),
		uintptr(hwndParent),
		uintptr(wID),
		0,
		0)
	return HWND(ret1)
}

func CreateToolbarEx(hwnd HWND, ws uint32, wID uint32, nBitmaps int32, hBMInst HINSTANCE, wBMID *uint32, lpButtons /*const*/ *TBBUTTON, iNumButtons int32, dxButton int32, dyButton int32, dxBitmap int32, dyBitmap int32, uStructSize uint32) HWND {
	ret1 := syscall15(createToolbarEx, 13,
		uintptr(hwnd),
		uintptr(ws),
		uintptr(wID),
		uintptr(nBitmaps),
		uintptr(hBMInst),
		uintptr(unsafe.Pointer(wBMID)),
		uintptr(unsafe.Pointer(lpButtons)),
		uintptr(iNumButtons),
		uintptr(dxButton),
		uintptr(dyButton),
		uintptr(dxBitmap),
		uintptr(dyBitmap),
		uintptr(uStructSize),
		0,
		0)
	return HWND(ret1)
}

func CreateUpDownControl(dwStyle uint32, x int32, y int32, cx int32, cy int32, hParent HWND, nID int32, hInst HINSTANCE, hBuddy HWND, nUpper int32, nLower int32, nPos int32) HWND {
	ret1 := syscall12(createUpDownControl, 12,
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent),
		uintptr(nID),
		uintptr(hInst),
		uintptr(hBuddy),
		uintptr(nUpper),
		uintptr(nLower),
		uintptr(nPos))
	return HWND(ret1)
}

func DPA_Create(cItemGrow int32) HDPA {
	ret1 := syscall3(dPA_Create, 1,
		uintptr(cItemGrow),
		0,
		0)
	return HDPA(ret1)
}

func DPA_DeleteAllPtrs(hdpa HDPA) bool {
	ret1 := syscall3(dPA_DeleteAllPtrs, 1,
		uintptr(hdpa),
		0,
		0)
	return ret1 != 0
}

func DPA_DeletePtr(hdpa HDPA, i int32) uintptr {
	ret1 := syscall3(dPA_DeletePtr, 2,
		uintptr(hdpa),
		uintptr(i),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func DPA_Destroy(hdpa HDPA) bool {
	ret1 := syscall3(dPA_Destroy, 1,
		uintptr(hdpa),
		0,
		0)
	return ret1 != 0
}

func DPA_DestroyCallback(hdpa HDPA, pfnCB DAENUMCALLBACK, pData uintptr) {
	pfnCBCallback := syscall.NewCallback(func(pRawArg uintptr, pDataRawArg uintptr) uintptr {
		ret := pfnCB(pRawArg, pDataRawArg)
		return uintptr(ret)
	})
	syscall3(dPA_DestroyCallback, 3,
		uintptr(hdpa),
		pfnCBCallback,
		pData)
}

func DPA_EnumCallback(hdpa HDPA, pfnCB DAENUMCALLBACK, pData uintptr) {
	pfnCBCallback := syscall.NewCallback(func(pRawArg uintptr, pDataRawArg uintptr) uintptr {
		ret := pfnCB(pRawArg, pDataRawArg)
		return uintptr(ret)
	})
	syscall3(dPA_EnumCallback, 3,
		uintptr(hdpa),
		pfnCBCallback,
		pData)
}

func DPA_GetPtr(hdpa HDPA, i INT_PTR) uintptr {
	ret1 := syscall3(dPA_GetPtr, 2,
		uintptr(hdpa),
		uintptr(unsafe.Pointer(i)),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func DPA_InsertPtr(hdpa HDPA, i int32, p uintptr) int32 {
	ret1 := syscall3(dPA_InsertPtr, 3,
		uintptr(hdpa),
		uintptr(i),
		p)
	return int32(ret1)
}

func DPA_Search(hdpa HDPA, pFind uintptr, iStart int32, pfnCompare DACOMPARE, lParam LPARAM, options uint32) int32 {
	pfnCompareCallback := syscall.NewCallback(func(p1RawArg uintptr, p2RawArg uintptr, lParamRawArg LPARAM) uintptr {
		ret := pfnCompare(p1RawArg, p2RawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(dPA_Search, 6,
		uintptr(hdpa),
		pFind,
		uintptr(iStart),
		pfnCompareCallback,
		uintptr(lParam),
		uintptr(options))
	return int32(ret1)
}

func DPA_SetPtr(hdpa HDPA, i int32, p uintptr) bool {
	ret1 := syscall3(dPA_SetPtr, 3,
		uintptr(hdpa),
		uintptr(i),
		p)
	return ret1 != 0
}

func DPA_Sort(hdpa HDPA, pfnCompare DACOMPARE, lParam LPARAM) bool {
	pfnCompareCallback := syscall.NewCallback(func(p1RawArg uintptr, p2RawArg uintptr, lParamRawArg LPARAM) uintptr {
		ret := pfnCompare(p1RawArg, p2RawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(dPA_Sort, 3,
		uintptr(hdpa),
		pfnCompareCallback,
		uintptr(lParam))
	return ret1 != 0
}

func DSA_Create(cbItem int32, cItemGrow int32) HDSA {
	ret1 := syscall3(dSA_Create, 2,
		uintptr(cbItem),
		uintptr(cItemGrow),
		0)
	return HDSA(ret1)
}

func DSA_DeleteAllItems(hdsa HDSA) bool {
	ret1 := syscall3(dSA_DeleteAllItems, 1,
		uintptr(hdsa),
		0,
		0)
	return ret1 != 0
}

func DSA_Destroy(hdsa HDSA) bool {
	ret1 := syscall3(dSA_Destroy, 1,
		uintptr(hdsa),
		0,
		0)
	return ret1 != 0
}

func DSA_DestroyCallback(hdsa HDSA, pfnCB DAENUMCALLBACK, pData uintptr) {
	pfnCBCallback := syscall.NewCallback(func(pRawArg uintptr, pDataRawArg uintptr) uintptr {
		ret := pfnCB(pRawArg, pDataRawArg)
		return uintptr(ret)
	})
	syscall3(dSA_DestroyCallback, 3,
		uintptr(hdsa),
		pfnCBCallback,
		pData)
}

func DSA_GetItemPtr(hdsa HDSA, i int32) uintptr {
	ret1 := syscall3(dSA_GetItemPtr, 2,
		uintptr(hdsa),
		uintptr(i),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func DSA_InsertItem(hdsa HDSA, i int32, pitem /*const*/ uintptr) int32 {
	ret1 := syscall3(dSA_InsertItem, 3,
		uintptr(hdsa),
		uintptr(i),
		pitem)
	return int32(ret1)
}

func DefSubclassProc(hWnd HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(defSubclassProc, 4,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return LRESULT(ret1)
}

func DestroyPropertySheetPage(unnamed0 HPROPSHEETPAGE) bool {
	ret1 := syscall3(destroyPropertySheetPage, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func DrawInsert(handParent HWND, hLB HWND, nItem int32) {
	syscall3(drawInsert, 3,
		uintptr(handParent),
		uintptr(hLB),
		uintptr(nItem))
}

func DrawStatusText(hDC HDC, lprc /*const*/ *RECT, pszText string, uFlags uint32) {
	pszTextStr := unicode16FromString(pszText)
	syscall6(drawStatusText, 4,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(unsafe.Pointer(&pszTextStr[0])),
		uintptr(uFlags),
		0,
		0)
}

func FlatSB_EnableScrollBar(unnamed0 HWND, unnamed1 int32, unnamed2 uint32) bool {
	ret1 := syscall3(flatSB_EnableScrollBar, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2))
	return ret1 != 0
}

func FlatSB_GetScrollInfo(unnamed0 HWND, code int32, unnamed2 *SCROLLINFO) bool {
	ret1 := syscall3(flatSB_GetScrollInfo, 3,
		uintptr(unnamed0),
		uintptr(code),
		uintptr(unsafe.Pointer(unnamed2)))
	return ret1 != 0
}

func FlatSB_GetScrollPos(unnamed0 HWND, code int32) int32 {
	ret1 := syscall3(flatSB_GetScrollPos, 2,
		uintptr(unnamed0),
		uintptr(code),
		0)
	return int32(ret1)
}

func FlatSB_GetScrollProp(unnamed0 HWND, propIndex int32, unnamed2 *int32) bool {
	ret1 := syscall3(flatSB_GetScrollProp, 3,
		uintptr(unnamed0),
		uintptr(propIndex),
		uintptr(unsafe.Pointer(unnamed2)))
	return ret1 != 0
}

func FlatSB_GetScrollPropPtr(unnamed0 HWND, propIndex int32, unnamed2 PINT_PTR) bool {
	ret1 := syscall3(flatSB_GetScrollPropPtr, 3,
		uintptr(unnamed0),
		uintptr(propIndex),
		uintptr(unsafe.Pointer(unnamed2)))
	return ret1 != 0
}

func FlatSB_GetScrollRange(unnamed0 HWND, code int32, unnamed2 *int32, unnamed3 *int32) bool {
	ret1 := syscall6(flatSB_GetScrollRange, 4,
		uintptr(unnamed0),
		uintptr(code),
		uintptr(unsafe.Pointer(unnamed2)),
		uintptr(unsafe.Pointer(unnamed3)),
		0,
		0)
	return ret1 != 0
}

func FlatSB_SetScrollInfo(unnamed0 HWND, code int32, unnamed2 *SCROLLINFO, fRedraw bool) int32 {
	ret1 := syscall6(flatSB_SetScrollInfo, 4,
		uintptr(unnamed0),
		uintptr(code),
		uintptr(unsafe.Pointer(unnamed2)),
		getUintptrFromBool(fRedraw),
		0,
		0)
	return int32(ret1)
}

func FlatSB_SetScrollPos(unnamed0 HWND, code int32, pos int32, fRedraw bool) int32 {
	ret1 := syscall6(flatSB_SetScrollPos, 4,
		uintptr(unnamed0),
		uintptr(code),
		uintptr(pos),
		getUintptrFromBool(fRedraw),
		0,
		0)
	return int32(ret1)
}

func FlatSB_SetScrollProp(unnamed0 HWND, index uint32, newValue INT_PTR, unnamed3 bool) bool {
	ret1 := syscall6(flatSB_SetScrollProp, 4,
		uintptr(unnamed0),
		uintptr(index),
		uintptr(unsafe.Pointer(newValue)),
		getUintptrFromBool(unnamed3),
		0,
		0)
	return ret1 != 0
}

func FlatSB_SetScrollRange(unnamed0 HWND, code int32, min int32, max int32, fRedraw bool) int32 {
	ret1 := syscall6(flatSB_SetScrollRange, 5,
		uintptr(unnamed0),
		uintptr(code),
		uintptr(min),
		uintptr(max),
		getUintptrFromBool(fRedraw),
		0)
	return int32(ret1)
}

func FlatSB_ShowScrollBar(unnamed0 HWND, code int32, unnamed2 bool) bool {
	ret1 := syscall3(flatSB_ShowScrollBar, 3,
		uintptr(unnamed0),
		uintptr(code),
		getUintptrFromBool(unnamed2))
	return ret1 != 0
}

func GetEffectiveClientRect(hWnd HWND, lprc *RECT, lpInfo /*const*/ *int32) {
	syscall3(getEffectiveClientRect, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(unsafe.Pointer(lpInfo)))
}

func GetMUILanguage() LANGID {
	ret1 := syscall3(getMUILanguage, 0,
		0,
		0,
		0)
	return LANGID(ret1)
}

func ImageList_Add(himl HIMAGELIST, hbmImage HBITMAP, hbmMask HBITMAP) int32 {
	ret1 := syscall3(imageList_Add, 3,
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(hbmMask))
	return int32(ret1)
}

func ImageList_AddMasked(himl HIMAGELIST, hbmImage HBITMAP, crMask COLORREF) int32 {
	ret1 := syscall3(imageList_AddMasked, 3,
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(crMask))
	return int32(ret1)
}

func ImageList_BeginDrag(himlTrack HIMAGELIST, iTrack int32, dxHotspot int32, dyHotspot int32) bool {
	ret1 := syscall6(imageList_BeginDrag, 4,
		uintptr(himlTrack),
		uintptr(iTrack),
		uintptr(dxHotspot),
		uintptr(dyHotspot),
		0,
		0)
	return ret1 != 0
}

func ImageList_Copy(himlDst HIMAGELIST, iDst int32, himlSrc HIMAGELIST, iSrc int32, uFlags uint32) bool {
	ret1 := syscall6(imageList_Copy, 5,
		uintptr(himlDst),
		uintptr(iDst),
		uintptr(himlSrc),
		uintptr(iSrc),
		uintptr(uFlags),
		0)
	return ret1 != 0
}

func ImageList_Create(cx int32, cy int32, flags uint32, cInitial int32, cGrow int32) HIMAGELIST {
	ret1 := syscall6(imageList_Create, 5,
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
		uintptr(cInitial),
		uintptr(cGrow),
		0)
	return HIMAGELIST(ret1)
}

func ImageList_Destroy(himl HIMAGELIST) bool {
	ret1 := syscall3(imageList_Destroy, 1,
		uintptr(himl),
		0,
		0)
	return ret1 != 0
}

func ImageList_DragEnter(hwndLock HWND, x int32, y int32) bool {
	ret1 := syscall3(imageList_DragEnter, 3,
		uintptr(hwndLock),
		uintptr(x),
		uintptr(y))
	return ret1 != 0
}

func ImageList_DragLeave(hwndLock HWND) bool {
	ret1 := syscall3(imageList_DragLeave, 1,
		uintptr(hwndLock),
		0,
		0)
	return ret1 != 0
}

func ImageList_DragMove(x int32, y int32) bool {
	ret1 := syscall3(imageList_DragMove, 2,
		uintptr(x),
		uintptr(y),
		0)
	return ret1 != 0
}

func ImageList_DragShowNolock(fShow bool) bool {
	ret1 := syscall3(imageList_DragShowNolock, 1,
		getUintptrFromBool(fShow),
		0,
		0)
	return ret1 != 0
}

func ImageList_Draw(himl HIMAGELIST, i int32, hdcDst HDC, x int32, y int32, fStyle uint32) bool {
	ret1 := syscall6(imageList_Draw, 6,
		uintptr(himl),
		uintptr(i),
		uintptr(hdcDst),
		uintptr(x),
		uintptr(y),
		uintptr(fStyle))
	return ret1 != 0
}

func ImageList_DrawEx(himl HIMAGELIST, i int32, hdcDst HDC, x int32, y int32, dx int32, dy int32, rgbBk COLORREF, rgbFg COLORREF, fStyle uint32) bool {
	ret1 := syscall12(imageList_DrawEx, 10,
		uintptr(himl),
		uintptr(i),
		uintptr(hdcDst),
		uintptr(x),
		uintptr(y),
		uintptr(dx),
		uintptr(dy),
		uintptr(rgbBk),
		uintptr(rgbFg),
		uintptr(fStyle),
		0,
		0)
	return ret1 != 0
}

func ImageList_DrawIndirect(pimldp *IMAGELISTDRAWPARAMS) bool {
	ret1 := syscall3(imageList_DrawIndirect, 1,
		uintptr(unsafe.Pointer(pimldp)),
		0,
		0)
	return ret1 != 0
}

func ImageList_Duplicate(himl HIMAGELIST) HIMAGELIST {
	ret1 := syscall3(imageList_Duplicate, 1,
		uintptr(himl),
		0,
		0)
	return HIMAGELIST(ret1)
}

func ImageList_EndDrag() {
	syscall3(imageList_EndDrag, 0,
		0,
		0,
		0)
}

func ImageList_GetBkColor(himl HIMAGELIST) COLORREF {
	ret1 := syscall3(imageList_GetBkColor, 1,
		uintptr(himl),
		0,
		0)
	return COLORREF(ret1)
}

func ImageList_GetDragImage(ppt *POINT, pptHotspot *POINT) HIMAGELIST {
	ret1 := syscall3(imageList_GetDragImage, 2,
		uintptr(unsafe.Pointer(ppt)),
		uintptr(unsafe.Pointer(pptHotspot)),
		0)
	return HIMAGELIST(ret1)
}

func ImageList_GetIcon(himl HIMAGELIST, i int32, flags uint32) HICON {
	ret1 := syscall3(imageList_GetIcon, 3,
		uintptr(himl),
		uintptr(i),
		uintptr(flags))
	return HICON(ret1)
}

func ImageList_GetIconSize(himl HIMAGELIST, cx *int, cy *int) bool {
	ret1 := syscall3(imageList_GetIconSize, 3,
		uintptr(himl),
		uintptr(unsafe.Pointer(cx)),
		uintptr(unsafe.Pointer(cy)))
	return ret1 != 0
}

func ImageList_GetImageCount(himl HIMAGELIST) int32 {
	ret1 := syscall3(imageList_GetImageCount, 1,
		uintptr(himl),
		0,
		0)
	return int32(ret1)
}

func ImageList_GetImageInfo(himl HIMAGELIST, i int32, pImageInfo *IMAGEINFO) bool {
	ret1 := syscall3(imageList_GetImageInfo, 3,
		uintptr(himl),
		uintptr(i),
		uintptr(unsafe.Pointer(pImageInfo)))
	return ret1 != 0
}

func ImageList_LoadImage(hi HINSTANCE, lpbmp string, cx int32, cGrow int32, crMask COLORREF, uType uint32, uFlags uint32) HIMAGELIST {
	lpbmpStr := unicode16FromString(lpbmp)
	ret1 := syscall9(imageList_LoadImage, 7,
		uintptr(hi),
		uintptr(unsafe.Pointer(&lpbmpStr[0])),
		uintptr(cx),
		uintptr(cGrow),
		uintptr(crMask),
		uintptr(uType),
		uintptr(uFlags),
		0,
		0)
	return HIMAGELIST(ret1)
}

func ImageList_Merge(himl1 HIMAGELIST, i1 int32, himl2 HIMAGELIST, i2 int32, dx int32, dy int32) HIMAGELIST {
	ret1 := syscall6(imageList_Merge, 6,
		uintptr(himl1),
		uintptr(i1),
		uintptr(himl2),
		uintptr(i2),
		uintptr(dx),
		uintptr(dy))
	return HIMAGELIST(ret1)
}

func ImageList_Read(pstm LPSTREAM) HIMAGELIST {
	ret1 := syscall3(imageList_Read, 1,
		uintptr(unsafe.Pointer(pstm)),
		0,
		0)
	return HIMAGELIST(ret1)
}

func ImageList_Remove(himl HIMAGELIST, i int32) bool {
	ret1 := syscall3(imageList_Remove, 2,
		uintptr(himl),
		uintptr(i),
		0)
	return ret1 != 0
}

func ImageList_Replace(himl HIMAGELIST, i int32, hbmImage HBITMAP, hbmMask HBITMAP) bool {
	ret1 := syscall6(imageList_Replace, 4,
		uintptr(himl),
		uintptr(i),
		uintptr(hbmImage),
		uintptr(hbmMask),
		0,
		0)
	return ret1 != 0
}

func ImageList_ReplaceIcon(himl HIMAGELIST, i int32, hicon HICON) int32 {
	ret1 := syscall3(imageList_ReplaceIcon, 3,
		uintptr(himl),
		uintptr(i),
		uintptr(hicon))
	return int32(ret1)
}

func ImageList_SetBkColor(himl HIMAGELIST, clrBk COLORREF) COLORREF {
	ret1 := syscall3(imageList_SetBkColor, 2,
		uintptr(himl),
		uintptr(clrBk),
		0)
	return COLORREF(ret1)
}

func ImageList_SetDragCursorImage(himlDrag HIMAGELIST, iDrag int32, dxHotspot int32, dyHotspot int32) bool {
	ret1 := syscall6(imageList_SetDragCursorImage, 4,
		uintptr(himlDrag),
		uintptr(iDrag),
		uintptr(dxHotspot),
		uintptr(dyHotspot),
		0,
		0)
	return ret1 != 0
}

func ImageList_SetIconSize(himl HIMAGELIST, cx int32, cy int32) bool {
	ret1 := syscall3(imageList_SetIconSize, 3,
		uintptr(himl),
		uintptr(cx),
		uintptr(cy))
	return ret1 != 0
}

func ImageList_SetImageCount(himl HIMAGELIST, uNewCount uint32) bool {
	ret1 := syscall3(imageList_SetImageCount, 2,
		uintptr(himl),
		uintptr(uNewCount),
		0)
	return ret1 != 0
}

func ImageList_SetOverlayImage(himl HIMAGELIST, iImage int32, iOverlay int32) bool {
	ret1 := syscall3(imageList_SetOverlayImage, 3,
		uintptr(himl),
		uintptr(iImage),
		uintptr(iOverlay))
	return ret1 != 0
}

func ImageList_Write(himl HIMAGELIST, pstm LPSTREAM) bool {
	ret1 := syscall3(imageList_Write, 2,
		uintptr(himl),
		uintptr(unsafe.Pointer(pstm)),
		0)
	return ret1 != 0
}

func InitCommonControls() {
	syscall3(initCommonControls, 0,
		0,
		0,
		0)
}

func InitCommonControlsEx(unnamed0 /*const*/ *INITCOMMONCONTROLSEX) bool {
	ret1 := syscall3(initCommonControlsEx, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func InitMUILanguage(uiLang LANGID) {
	syscall3(initMUILanguage, 1,
		uintptr(uiLang),
		0,
		0)
}

func InitializeFlatSB(unnamed0 HWND) bool {
	ret1 := syscall3(initializeFlatSB, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func LBItemFromPt(hLB HWND, pt POINT, bAutoScroll bool) int32 {
	ret1 := syscall6(lBItemFromPt, 4,
		uintptr(hLB),
		uintptr(pt.X),
		uintptr(pt.Y),
		getUintptrFromBool(bAutoScroll),
		0,
		0)
	return int32(ret1)
}

func MakeDragList(hLB HWND) bool {
	ret1 := syscall3(makeDragList, 1,
		uintptr(hLB),
		0,
		0)
	return ret1 != 0
}

func MenuHelp(uMsg uint32, wParam WPARAM, lParam LPARAM, hMainMenu HMENU, hInst HINSTANCE, hwndStatus HWND, lpwIDs *UINT) {
	syscall9(menuHelp, 7,
		uintptr(uMsg),
		uintptr(wParam),
		uintptr(lParam),
		uintptr(hMainMenu),
		uintptr(hInst),
		uintptr(hwndStatus),
		uintptr(unsafe.Pointer(lpwIDs)),
		0,
		0)
}

func PropertySheet(unnamed0 /*const*/ *PROPSHEETHEADER) INT_PTR {
	ret1 := syscall3(propertySheet, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return (INT_PTR)(unsafe.Pointer(ret1))
}

func RemoveWindowSubclass(hWnd HWND, pfnSubclass SUBCLASSPROC, uIdSubclass *uint32) bool {
	pfnSubclassCallback := syscall.NewCallback(func(hWndRawArg HWND, uMsgRawArg UINT, wParamRawArg WPARAM, lParamRawArg LPARAM, uIdSubclassRawArg UINT_PTR, dwRefDataRawArg DWORD_PTR) uintptr {
		ret := pfnSubclass(hWndRawArg, uMsgRawArg, wParamRawArg, lParamRawArg, uIdSubclassRawArg, dwRefDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(removeWindowSubclass, 3,
		uintptr(hWnd),
		pfnSubclassCallback,
		uintptr(unsafe.Pointer(uIdSubclass)))
	return ret1 != 0
}

func SetWindowSubclass(hWnd HWND, pfnSubclass SUBCLASSPROC, uIdSubclass *uint32, dwRefData *uint32) bool {
	pfnSubclassCallback := syscall.NewCallback(func(hWndRawArg HWND, uMsgRawArg UINT, wParamRawArg WPARAM, lParamRawArg LPARAM, uIdSubclassRawArg UINT_PTR, dwRefDataRawArg DWORD_PTR) uintptr {
		ret := pfnSubclass(hWndRawArg, uMsgRawArg, wParamRawArg, lParamRawArg, uIdSubclassRawArg, dwRefDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(setWindowSubclass, 4,
		uintptr(hWnd),
		pfnSubclassCallback,
		uintptr(unsafe.Pointer(uIdSubclass)),
		uintptr(unsafe.Pointer(dwRefData)),
		0,
		0)
	return ret1 != 0
}

func ShowHideMenuCtl(hWnd HWND, uFlags *uint32, lpInfo *int32) bool {
	ret1 := syscall3(showHideMenuCtl, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(uFlags)),
		uintptr(unsafe.Pointer(lpInfo)))
	return ret1 != 0
}

func Str_SetPtrW(ppsz *LPWSTR, psz string) bool {
	pszStr := unicode16FromString(psz)
	ret1 := syscall3(str_SetPtrW, 2,
		uintptr(unsafe.Pointer(ppsz)),
		uintptr(unsafe.Pointer(&pszStr[0])),
		0)
	return ret1 != 0
}

func TaskDialog(hwndOwner HWND, hInstance HINSTANCE, pszWindowTitle string, pszMainInstruction string, pszContent string, dwCommonButtons TASKDIALOG_COMMON_BUTTON_FLAGS, pszIcon string, pnButton *int) HRESULT {
	pszWindowTitleStr := unicode16FromString(pszWindowTitle)
	pszMainInstructionStr := unicode16FromString(pszMainInstruction)
	pszContentStr := unicode16FromString(pszContent)
	pszIconStr := unicode16FromString(pszIcon)
	ret1 := syscall9(taskDialog, 8,
		uintptr(hwndOwner),
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&pszWindowTitleStr[0])),
		uintptr(unsafe.Pointer(&pszMainInstructionStr[0])),
		uintptr(unsafe.Pointer(&pszContentStr[0])),
		uintptr(dwCommonButtons),
		uintptr(unsafe.Pointer(&pszIconStr[0])),
		uintptr(unsafe.Pointer(pnButton)),
		0)
	return HRESULT(ret1)
}

func TaskDialogIndirect(pTaskConfig /*const*/ *TASKDIALOGCONFIG, pnButton *int, pnRadioButton *int, pfVerificationFlagChecked *BOOL) HRESULT {
	ret1 := syscall6(taskDialogIndirect, 4,
		uintptr(unsafe.Pointer(pTaskConfig)),
		uintptr(unsafe.Pointer(pnButton)),
		uintptr(unsafe.Pointer(pnRadioButton)),
		uintptr(unsafe.Pointer(pfVerificationFlagChecked)),
		0,
		0)
	return HRESULT(ret1)
}

func UninitializeFlatSB(unnamed0 HWND) HRESULT {
	ret1 := syscall3(uninitializeFlatSB, 1,
		uintptr(unnamed0),
		0,
		0)
	return HRESULT(ret1)
}

func AddMRUStringW(hList HANDLE, lpszString string) int32 {
	lpszStringStr := unicode16FromString(lpszString)
	ret1 := syscall3(addMRUStringW, 2,
		uintptr(hList),
		uintptr(unsafe.Pointer(&lpszStringStr[0])),
		0)
	return int32(ret1)
}

func CreateMRUListW(infoW /*const*/ *MRUINFO) HANDLE {
	ret1 := syscall3(createMRUListW, 1,
		uintptr(unsafe.Pointer(infoW)),
		0,
		0)
	return HANDLE(ret1)
}

func CreateToolbar(hwnd HWND, style uint32, wID uint32, nBitmaps int32, hBMInst HINSTANCE, wBMID uint32, lpButtons /*const*/ *TBBUTTON, iNumButtons int32) HWND {
	ret1 := syscall9(createToolbar, 8,
		uintptr(hwnd),
		uintptr(style),
		uintptr(wID),
		uintptr(nBitmaps),
		uintptr(hBMInst),
		uintptr(wBMID),
		uintptr(unsafe.Pointer(lpButtons)),
		uintptr(iNumButtons),
		0)
	return HWND(ret1)
}

func EnumMRUListW(hList HANDLE, nItemPos int32, lpBuffer LPVOID, nBufferSize uint32) int32 {
	ret1 := syscall6(enumMRUListW, 4,
		uintptr(hList),
		uintptr(nItemPos),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nBufferSize),
		0,
		0)
	return int32(ret1)
}

func FreeMRUList(hMRUList HANDLE) {
	syscall3(freeMRUList, 1,
		uintptr(hMRUList),
		0,
		0)
}

func ImageList_AddIcon(himl HIMAGELIST, hIcon HICON) int32 {
	ret1 := syscall3(imageList_AddIcon, 2,
		uintptr(himl),
		uintptr(hIcon),
		0)
	return int32(ret1)
}

func ImageList_GetFlags(himl HIMAGELIST) uint32 {
	ret1 := syscall3(imageList_GetFlags, 1,
		uintptr(himl),
		0,
		0)
	return uint32(ret1)
}

func ImageList_GetImageRect(himl HIMAGELIST, i int32, lpRect *RECT) bool {
	ret1 := syscall3(imageList_GetImageRect, 3,
		uintptr(himl),
		uintptr(i),
		uintptr(unsafe.Pointer(lpRect)))
	return ret1 != 0
}

func ImageList_SetFilter(himl HIMAGELIST, i int32, dwFilter uint32) bool {
	ret1 := syscall3(imageList_SetFilter, 3,
		uintptr(himl),
		uintptr(i),
		uintptr(dwFilter))
	return ret1 != 0
}

func ImageList_SetFlags(himl HIMAGELIST, flags uint32) uint32 {
	ret1 := syscall3(imageList_SetFlags, 2,
		uintptr(himl),
		uintptr(flags),
		0)
	return uint32(ret1)
}
