// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libuser32 uintptr

	// Functions
	activateKeyboardLayout             uintptr
	adjustWindowRect                   uintptr
	adjustWindowRectEx                 uintptr
	allowSetForegroundWindow           uintptr
	animateWindow                      uintptr
	anyPopup                           uintptr
	appendMenu                         uintptr
	arrangeIconicWindows               uintptr
	attachThreadInput                  uintptr
	beginDeferWindowPos                uintptr
	beginPaint                         uintptr
	blockInput                         uintptr
	bringWindowToTop                   uintptr
	broadcastSystemMessageEx           uintptr
	broadcastSystemMessage             uintptr
	callMsgFilter                      uintptr
	callNextHookEx                     uintptr
	callWindowProc                     uintptr
	cascadeWindows                     uintptr
	changeClipboardChain               uintptr
	changeDisplaySettingsEx            uintptr
	changeDisplaySettings              uintptr
	changeMenu                         uintptr
	charLowerBuff                      uintptr
	charLower                          uintptr
	charNextExA                        uintptr
	charNext                           uintptr
	charPrevExA                        uintptr
	charPrev                           uintptr
	charToOemBuff                      uintptr
	charToOem                          uintptr
	charUpperBuff                      uintptr
	charUpper                          uintptr
	checkDlgButton                     uintptr
	checkMenuItem                      uintptr
	checkMenuRadioItem                 uintptr
	checkRadioButton                   uintptr
	childWindowFromPoint               uintptr
	childWindowFromPointEx             uintptr
	clientToScreen                     uintptr
	clipCursor                         uintptr
	closeClipboard                     uintptr
	closeDesktop                       uintptr
	closeGestureInfoHandle             uintptr
	closeTouchInputHandle              uintptr
	closeWindow                        uintptr
	closeWindowStation                 uintptr
	copyAcceleratorTable               uintptr
	copyIcon                           uintptr
	copyImage                          uintptr
	copyRect                           uintptr
	countClipboardFormats              uintptr
	createAcceleratorTable             uintptr
	createCaret                        uintptr
	createCursor                       uintptr
	createDesktop                      uintptr
	createDialogIndirectParam          uintptr
	createDialogParam                  uintptr
	createIcon                         uintptr
	createIconFromResource             uintptr
	createIconFromResourceEx           uintptr
	createIconIndirect                 uintptr
	createMDIWindow                    uintptr
	createMenu                         uintptr
	createPopupMenu                    uintptr
	createWindowEx                     uintptr
	createWindowStation                uintptr
	ddeAbandonTransaction              uintptr
	ddeAccessData                      uintptr
	ddeAddData                         uintptr
	ddeClientTransaction               uintptr
	ddeCmpStringHandles                uintptr
	ddeConnect                         uintptr
	ddeConnectList                     uintptr
	ddeCreateDataHandle                uintptr
	ddeCreateStringHandle              uintptr
	ddeDisconnect                      uintptr
	ddeDisconnectList                  uintptr
	ddeEnableCallback                  uintptr
	ddeFreeDataHandle                  uintptr
	ddeFreeStringHandle                uintptr
	ddeGetData                         uintptr
	ddeGetLastError                    uintptr
	ddeImpersonateClient               uintptr
	ddeInitialize                      uintptr
	ddeKeepStringHandle                uintptr
	ddeNameService                     uintptr
	ddePostAdvise                      uintptr
	ddeQueryConvInfo                   uintptr
	ddeQueryNextServer                 uintptr
	ddeQueryString                     uintptr
	ddeReconnect                       uintptr
	ddeSetQualityOfService             uintptr
	ddeSetUserHandle                   uintptr
	ddeUnaccessData                    uintptr
	ddeUninitialize                    uintptr
	defDlgProc                         uintptr
	defFrameProc                       uintptr
	defMDIChildProc                    uintptr
	defRawInputProc                    uintptr
	defWindowProc                      uintptr
	deferWindowPos                     uintptr
	deleteMenu                         uintptr
	deregisterShellHookWindow          uintptr
	destroyAcceleratorTable            uintptr
	destroyCaret                       uintptr
	destroyCursor                      uintptr
	destroyIcon                        uintptr
	destroyMenu                        uintptr
	destroyWindow                      uintptr
	dialogBoxIndirectParam             uintptr
	dialogBoxParam                     uintptr
	disableProcessWindowsGhosting      uintptr
	dispatchMessage                    uintptr
	dlgDirListComboBox                 uintptr
	dlgDirList                         uintptr
	dlgDirSelectComboBoxEx             uintptr
	dlgDirSelectEx                     uintptr
	dragDetect                         uintptr
	dragObject                         uintptr
	drawAnimatedRects                  uintptr
	drawCaption                        uintptr
	drawEdge                           uintptr
	drawFocusRect                      uintptr
	drawFrameControl                   uintptr
	drawIcon                           uintptr
	drawIconEx                         uintptr
	drawMenuBar                        uintptr
	drawState                          uintptr
	drawTextEx                         uintptr
	drawText                           uintptr
	emptyClipboard                     uintptr
	enableMenuItem                     uintptr
	enableScrollBar                    uintptr
	enableWindow                       uintptr
	endDeferWindowPos                  uintptr
	endDialog                          uintptr
	endMenu                            uintptr
	endPaint                           uintptr
	endTask                            uintptr
	enumChildWindows                   uintptr
	enumClipboardFormats               uintptr
	enumDesktopWindows                 uintptr
	enumDesktops                       uintptr
	enumDisplayDevices                 uintptr
	enumDisplayMonitors                uintptr
	enumDisplaySettingsEx              uintptr
	enumDisplaySettings                uintptr
	enumPropsEx                        uintptr
	enumProps                          uintptr
	enumThreadWindows                  uintptr
	enumWindowStations                 uintptr
	enumWindows                        uintptr
	equalRect                          uintptr
	excludeUpdateRgn                   uintptr
	exitWindowsEx                      uintptr
	fillRect                           uintptr
	findWindowEx                       uintptr
	findWindow                         uintptr
	flashWindow                        uintptr
	flashWindowEx                      uintptr
	frameRect                          uintptr
	freeDDElParam                      uintptr
	getActiveWindow                    uintptr
	getAltTabInfo                      uintptr
	getAncestor                        uintptr
	getAsyncKeyState                   uintptr
	getCapture                         uintptr
	getCaretBlinkTime                  uintptr
	getCaretPos                        uintptr
	getClassInfoEx                     uintptr
	getClassInfo                       uintptr
	getClassLongPtr                    uintptr
	getClassLong                       uintptr
	getClassName                       uintptr
	getClassWord                       uintptr
	getClientRect                      uintptr
	getClipCursor                      uintptr
	getClipboardData                   uintptr
	getClipboardFormatName             uintptr
	getClipboardOwner                  uintptr
	getClipboardSequenceNumber         uintptr
	getClipboardViewer                 uintptr
	getComboBoxInfo                    uintptr
	getCursor                          uintptr
	getCursorInfo                      uintptr
	getCursorPos                       uintptr
	getDC                              uintptr
	getDCEx                            uintptr
	getDesktopWindow                   uintptr
	getDialogBaseUnits                 uintptr
	getDlgCtrlID                       uintptr
	getDlgItem                         uintptr
	getDlgItemInt                      uintptr
	getDlgItemText                     uintptr
	getDoubleClickTime                 uintptr
	getFocus                           uintptr
	getForegroundWindow                uintptr
	getGUIThreadInfo                   uintptr
	getGestureConfig                   uintptr
	getGestureExtraArgs                uintptr
	getGestureInfo                     uintptr
	getGuiResources                    uintptr
	getIconInfo                        uintptr
	getInputState                      uintptr
	getKBCodePage                      uintptr
	getKeyNameText                     uintptr
	getKeyState                        uintptr
	getKeyboardLayout                  uintptr
	getKeyboardLayoutList              uintptr
	getKeyboardLayoutName              uintptr
	getKeyboardState                   uintptr
	getKeyboardType                    uintptr
	getLastActivePopup                 uintptr
	getLastInputInfo                   uintptr
	getLayeredWindowAttributes         uintptr
	getListBoxInfo                     uintptr
	getMenu                            uintptr
	getMenuBarInfo                     uintptr
	getMenuCheckMarkDimensions         uintptr
	getMenuContextHelpId               uintptr
	getMenuDefaultItem                 uintptr
	getMenuInfo                        uintptr
	getMenuItemCount                   uintptr
	getMenuItemID                      uintptr
	getMenuItemInfo                    uintptr
	getMenuItemRect                    uintptr
	getMenuState                       uintptr
	getMenuString                      uintptr
	getMessageExtraInfo                uintptr
	getMessagePos                      uintptr
	getMessageTime                     uintptr
	getMessage                         uintptr
	getMonitorInfo                     uintptr
	getMouseMovePointsEx               uintptr
	getNextDlgGroupItem                uintptr
	getNextDlgTabItem                  uintptr
	getOpenClipboardWindow             uintptr
	getParent                          uintptr
	getPriorityClipboardFormat         uintptr
	getProcessDefaultLayout            uintptr
	getProcessWindowStation            uintptr
	getProp                            uintptr
	getQueueStatus                     uintptr
	getRawInputBuffer                  uintptr
	getRawInputData                    uintptr
	getRawInputDeviceInfo              uintptr
	getRawInputDeviceList              uintptr
	getRegisteredRawInputDevices       uintptr
	getScrollBarInfo                   uintptr
	getScrollInfo                      uintptr
	getScrollPos                       uintptr
	getScrollRange                     uintptr
	getShellWindow                     uintptr
	getSubMenu                         uintptr
	getSysColor                        uintptr
	getSysColorBrush                   uintptr
	getSystemMenu                      uintptr
	getSystemMetrics                   uintptr
	getTabbedTextExtent                uintptr
	getThreadDesktop                   uintptr
	getTitleBarInfo                    uintptr
	getTopWindow                       uintptr
	getTouchInputInfo                  uintptr
	getUpdateRect                      uintptr
	getUpdateRgn                       uintptr
	getUserObjectInformation           uintptr
	getUserObjectSecurity              uintptr
	getWindow                          uintptr
	getWindowContextHelpId             uintptr
	getWindowDC                        uintptr
	getWindowInfo                      uintptr
	getWindowLongPtr                   uintptr
	getWindowLong                      uintptr
	getWindowModuleFileName            uintptr
	getWindowPlacement                 uintptr
	getWindowRect                      uintptr
	getWindowRgn                       uintptr
	getWindowRgnBox                    uintptr
	getWindowTextLength                uintptr
	getWindowText                      uintptr
	getWindowThreadProcessId           uintptr
	getWindowWord                      uintptr
	grayString                         uintptr
	hideCaret                          uintptr
	hiliteMenuItem                     uintptr
	iMPGetIME                          uintptr
	iMPQueryIME                        uintptr
	iMPSetIME                          uintptr
	impersonateDdeClientWindow         uintptr
	inSendMessage                      uintptr
	inSendMessageEx                    uintptr
	inflateRect                        uintptr
	insertMenuItem                     uintptr
	insertMenu                         uintptr
	internalGetWindowText              uintptr
	intersectRect                      uintptr
	invalidateRect                     uintptr
	invalidateRgn                      uintptr
	invertRect                         uintptr
	isCharAlphaNumeric                 uintptr
	isCharAlpha                        uintptr
	isCharLower                        uintptr
	isCharUpper                        uintptr
	isChild                            uintptr
	isClipboardFormatAvailable         uintptr
	isDialogMessage                    uintptr
	isDlgButtonChecked                 uintptr
	isGUIThread                        uintptr
	isHungAppWindow                    uintptr
	isIconic                           uintptr
	isMenu                             uintptr
	isRectEmpty                        uintptr
	isTouchWindow                      uintptr
	isWinEventHookInstalled            uintptr
	isWindow                           uintptr
	isWindowEnabled                    uintptr
	isWindowUnicode                    uintptr
	isWindowVisible                    uintptr
	isWow64Message                     uintptr
	isZoomed                           uintptr
	killTimer                          uintptr
	loadAccelerators                   uintptr
	loadBitmap                         uintptr
	loadCursorFromFile                 uintptr
	loadCursor                         uintptr
	loadIcon                           uintptr
	loadImage                          uintptr
	loadKeyboardLayout                 uintptr
	loadMenuIndirect                   uintptr
	loadMenu                           uintptr
	loadString                         uintptr
	lockSetForegroundWindow            uintptr
	lockWindowUpdate                   uintptr
	lockWorkStation                    uintptr
	lookupIconIdFromDirectory          uintptr
	lookupIconIdFromDirectoryEx        uintptr
	mapDialogRect                      uintptr
	mapVirtualKeyEx                    uintptr
	mapVirtualKey                      uintptr
	mapWindowPoints                    uintptr
	menuItemFromPoint                  uintptr
	messageBeep                        uintptr
	messageBoxEx                       uintptr
	messageBoxIndirect                 uintptr
	messageBox                         uintptr
	modifyMenu                         uintptr
	monitorFromPoint                   uintptr
	monitorFromRect                    uintptr
	monitorFromWindow                  uintptr
	moveWindow                         uintptr
	msgWaitForMultipleObjects          uintptr
	msgWaitForMultipleObjectsEx        uintptr
	notifyWinEvent                     uintptr
	oemKeyScan                         uintptr
	oemToCharBuff                      uintptr
	oemToChar                          uintptr
	offsetRect                         uintptr
	openClipboard                      uintptr
	openDesktop                        uintptr
	openIcon                           uintptr
	openInputDesktop                   uintptr
	openWindowStation                  uintptr
	packDDElParam                      uintptr
	paintDesktop                       uintptr
	peekMessage                        uintptr
	postMessage                        uintptr
	postQuitMessage                    uintptr
	postThreadMessage                  uintptr
	printWindow                        uintptr
	privateExtractIcons                uintptr
	ptInRect                           uintptr
	realChildWindowFromPoint           uintptr
	realGetWindowClass                 uintptr
	redrawWindow                       uintptr
	registerClassEx                    uintptr
	registerClass                      uintptr
	registerClipboardFormat            uintptr
	registerDeviceNotification         uintptr
	registerHotKey                     uintptr
	registerPowerSettingNotification   uintptr
	registerRawInputDevices            uintptr
	registerShellHookWindow            uintptr
	registerTouchWindow                uintptr
	registerWindowMessage              uintptr
	releaseCapture                     uintptr
	releaseDC                          uintptr
	removeMenu                         uintptr
	removeProp                         uintptr
	replyMessage                       uintptr
	reuseDDElParam                     uintptr
	screenToClient                     uintptr
	scrollDC                           uintptr
	scrollWindow                       uintptr
	scrollWindowEx                     uintptr
	sendDlgItemMessage                 uintptr
	sendIMEMessageEx                   uintptr
	sendInput                          uintptr
	sendMessageCallback                uintptr
	sendMessageTimeout                 uintptr
	sendMessage                        uintptr
	sendNotifyMessage                  uintptr
	setActiveWindow                    uintptr
	setCapture                         uintptr
	setCaretBlinkTime                  uintptr
	setCaretPos                        uintptr
	setClassLongPtr                    uintptr
	setClassLong                       uintptr
	setClassWord                       uintptr
	setClipboardData                   uintptr
	setClipboardViewer                 uintptr
	setCursor                          uintptr
	setCursorPos                       uintptr
	setDebugErrorLevel                 uintptr
	setDlgItemInt                      uintptr
	setDlgItemText                     uintptr
	setDoubleClickTime                 uintptr
	setFocus                           uintptr
	setForegroundWindow                uintptr
	setGestureConfig                   uintptr
	setKeyboardState                   uintptr
	setLastErrorEx                     uintptr
	setLayeredWindowAttributes         uintptr
	setMenu                            uintptr
	setMenuContextHelpId               uintptr
	setMenuDefaultItem                 uintptr
	setMenuInfo                        uintptr
	setMenuItemBitmaps                 uintptr
	setMenuItemInfo                    uintptr
	setMessageExtraInfo                uintptr
	setMessageQueue                    uintptr
	setParent                          uintptr
	setProcessDefaultLayout            uintptr
	setProcessWindowStation            uintptr
	setProp                            uintptr
	setRect                            uintptr
	setRectEmpty                       uintptr
	setScrollInfo                      uintptr
	setScrollPos                       uintptr
	setScrollRange                     uintptr
	setSysColors                       uintptr
	setSystemCursor                    uintptr
	setThreadDesktop                   uintptr
	setTimer                           uintptr
	setUserObjectInformation           uintptr
	setUserObjectSecurity              uintptr
	setWinEventHook                    uintptr
	setWindowContextHelpId             uintptr
	setWindowLongPtr                   uintptr
	setWindowLong                      uintptr
	setWindowPlacement                 uintptr
	setWindowPos                       uintptr
	setWindowRgn                       uintptr
	setWindowText                      uintptr
	setWindowWord                      uintptr
	setWindowsHookEx                   uintptr
	setWindowsHook                     uintptr
	showCaret                          uintptr
	showCursor                         uintptr
	showOwnedPopups                    uintptr
	showScrollBar                      uintptr
	showWindow                         uintptr
	showWindowAsync                    uintptr
	subtractRect                       uintptr
	swapMouseButton                    uintptr
	switchDesktop                      uintptr
	switchToThisWindow                 uintptr
	systemParametersInfo               uintptr
	tabbedTextOut                      uintptr
	tileWindows                        uintptr
	toAscii                            uintptr
	toAsciiEx                          uintptr
	toUnicode                          uintptr
	toUnicodeEx                        uintptr
	trackMouseEvent                    uintptr
	trackPopupMenu                     uintptr
	trackPopupMenuEx                   uintptr
	translateAccelerator               uintptr
	translateMDISysAccel               uintptr
	translateMessage                   uintptr
	unhookWinEvent                     uintptr
	unhookWindowsHook                  uintptr
	unhookWindowsHookEx                uintptr
	unionRect                          uintptr
	unloadKeyboardLayout               uintptr
	unpackDDElParam                    uintptr
	unregisterClass                    uintptr
	unregisterDeviceNotification       uintptr
	unregisterHotKey                   uintptr
	unregisterPowerSettingNotification uintptr
	unregisterTouchWindow              uintptr
	updateLayeredWindow                uintptr
	updateLayeredWindowIndirect        uintptr
	updateWindow                       uintptr
	userHandleGrantAccess              uintptr
	validateRect                       uintptr
	validateRgn                        uintptr
	vkKeyScanEx                        uintptr
	vkKeyScan                          uintptr
	wINNLSEnableIME                    uintptr
	wINNLSGetEnableStatus              uintptr
	wINNLSGetIMEHotkey                 uintptr
	waitForInputIdle                   uintptr
	waitMessage                        uintptr
	winHelp                            uintptr
	windowFromDC                       uintptr
	windowFromPoint                    uintptr
	keybd_event                        uintptr
	mouse_event                        uintptr
	alignRects                         uintptr
	cascadeChildWindows                uintptr
	createDialogIndirectParamAorW      uintptr
	dialogBoxIndirectParamAorW         uintptr
	drawCaptionTemp                    uintptr
	drawMenuBarTemp                    uintptr
	getAppCompatFlags                  uintptr
	getAppCompatFlags2                 uintptr
	getCursorFrameInfo                 uintptr
	getInternalWindowPos               uintptr
	getProgmanWindow                   uintptr
	getTaskmanWindow                   uintptr
	killSystemTimer                    uintptr
	loadLocalFonts                     uintptr
	messageBoxTimeout                  uintptr
	privateExtractIconEx               uintptr
	registerLogonProcess               uintptr
	registerServicesProcess            uintptr
	registerSystemThread               uintptr
	registerTasklist                   uintptr
	scrollChildren                     uintptr
	setInternalWindowPos               uintptr
	setLogonNotifyWindow               uintptr
	setProgmanWindow                   uintptr
	setShellWindow                     uintptr
	setShellWindowEx                   uintptr
	setSysColorsTemp                   uintptr
	setSystemMenu                      uintptr
	setSystemTimer                     uintptr
	setTaskmanWindow                   uintptr
	setWindowStationUser               uintptr
	tileChildWindows                   uintptr
	user32InitializeImmEntryTable      uintptr
	userRealizePalette                 uintptr
	userRegisterWowHandlers            uintptr
)

func init() {
	// Library
	libuser32 = doLoadLibrary("user32.dll")

	// Functions
	activateKeyboardLayout = doGetProcAddress(libuser32, "ActivateKeyboardLayout")
	adjustWindowRect = doGetProcAddress(libuser32, "AdjustWindowRect")
	adjustWindowRectEx = doGetProcAddress(libuser32, "AdjustWindowRectEx")
	allowSetForegroundWindow = doGetProcAddress(libuser32, "AllowSetForegroundWindow")
	animateWindow = doGetProcAddress(libuser32, "AnimateWindow")
	anyPopup = doGetProcAddress(libuser32, "AnyPopup")
	appendMenu = doGetProcAddress(libuser32, "AppendMenuW")
	arrangeIconicWindows = doGetProcAddress(libuser32, "ArrangeIconicWindows")
	attachThreadInput = doGetProcAddress(libuser32, "AttachThreadInput")
	beginDeferWindowPos = doGetProcAddress(libuser32, "BeginDeferWindowPos")
	beginPaint = doGetProcAddress(libuser32, "BeginPaint")
	blockInput = doGetProcAddress(libuser32, "BlockInput")
	bringWindowToTop = doGetProcAddress(libuser32, "BringWindowToTop")
	broadcastSystemMessageEx = doGetProcAddress(libuser32, "BroadcastSystemMessageExW")
	broadcastSystemMessage = doGetProcAddress(libuser32, "BroadcastSystemMessageW")
	callMsgFilter = doGetProcAddress(libuser32, "CallMsgFilterW")
	callNextHookEx = doGetProcAddress(libuser32, "CallNextHookEx")
	callWindowProc = doGetProcAddress(libuser32, "CallWindowProcW")
	cascadeWindows = doGetProcAddress(libuser32, "CascadeWindows")
	changeClipboardChain = doGetProcAddress(libuser32, "ChangeClipboardChain")
	changeDisplaySettingsEx = doGetProcAddress(libuser32, "ChangeDisplaySettingsExW")
	changeDisplaySettings = doGetProcAddress(libuser32, "ChangeDisplaySettingsW")
	changeMenu = doGetProcAddress(libuser32, "ChangeMenuW")
	charLowerBuff = doGetProcAddress(libuser32, "CharLowerBuffW")
	charLower = doGetProcAddress(libuser32, "CharLowerW")
	charNextExA = doGetProcAddress(libuser32, "CharNextExA")
	charNext = doGetProcAddress(libuser32, "CharNextW")
	charPrevExA = doGetProcAddress(libuser32, "CharPrevExA")
	charPrev = doGetProcAddress(libuser32, "CharPrevW")
	charToOemBuff = doGetProcAddress(libuser32, "CharToOemBuffW")
	charToOem = doGetProcAddress(libuser32, "CharToOemW")
	charUpperBuff = doGetProcAddress(libuser32, "CharUpperBuffW")
	charUpper = doGetProcAddress(libuser32, "CharUpperW")
	checkDlgButton = doGetProcAddress(libuser32, "CheckDlgButton")
	checkMenuItem = doGetProcAddress(libuser32, "CheckMenuItem")
	checkMenuRadioItem = doGetProcAddress(libuser32, "CheckMenuRadioItem")
	checkRadioButton = doGetProcAddress(libuser32, "CheckRadioButton")
	childWindowFromPoint = doGetProcAddress(libuser32, "ChildWindowFromPoint")
	childWindowFromPointEx = doGetProcAddress(libuser32, "ChildWindowFromPointEx")
	clientToScreen = doGetProcAddress(libuser32, "ClientToScreen")
	clipCursor = doGetProcAddress(libuser32, "ClipCursor")
	closeClipboard = doGetProcAddress(libuser32, "CloseClipboard")
	closeDesktop = doGetProcAddress(libuser32, "CloseDesktop")
	closeGestureInfoHandle = doGetProcAddress(libuser32, "CloseGestureInfoHandle")
	closeTouchInputHandle = doGetProcAddress(libuser32, "CloseTouchInputHandle")
	closeWindow = doGetProcAddress(libuser32, "CloseWindow")
	closeWindowStation = doGetProcAddress(libuser32, "CloseWindowStation")
	copyAcceleratorTable = doGetProcAddress(libuser32, "CopyAcceleratorTableW")
	copyIcon = doGetProcAddress(libuser32, "CopyIcon")
	copyImage = doGetProcAddress(libuser32, "CopyImage")
	copyRect = doGetProcAddress(libuser32, "CopyRect")
	countClipboardFormats = doGetProcAddress(libuser32, "CountClipboardFormats")
	createAcceleratorTable = doGetProcAddress(libuser32, "CreateAcceleratorTableW")
	createCaret = doGetProcAddress(libuser32, "CreateCaret")
	createCursor = doGetProcAddress(libuser32, "CreateCursor")
	createDesktop = doGetProcAddress(libuser32, "CreateDesktopW")
	createDialogIndirectParam = doGetProcAddress(libuser32, "CreateDialogIndirectParamW")
	createDialogParam = doGetProcAddress(libuser32, "CreateDialogParamW")
	createIcon = doGetProcAddress(libuser32, "CreateIcon")
	createIconFromResource = doGetProcAddress(libuser32, "CreateIconFromResource")
	createIconFromResourceEx = doGetProcAddress(libuser32, "CreateIconFromResourceEx")
	createIconIndirect = doGetProcAddress(libuser32, "CreateIconIndirect")
	createMDIWindow = doGetProcAddress(libuser32, "CreateMDIWindowW")
	createMenu = doGetProcAddress(libuser32, "CreateMenu")
	createPopupMenu = doGetProcAddress(libuser32, "CreatePopupMenu")
	createWindowEx = doGetProcAddress(libuser32, "CreateWindowExW")
	createWindowStation = doGetProcAddress(libuser32, "CreateWindowStationW")
	ddeAbandonTransaction = doGetProcAddress(libuser32, "DdeAbandonTransaction")
	ddeAccessData = doGetProcAddress(libuser32, "DdeAccessData")
	ddeAddData = doGetProcAddress(libuser32, "DdeAddData")
	ddeClientTransaction = doGetProcAddress(libuser32, "DdeClientTransaction")
	ddeCmpStringHandles = doGetProcAddress(libuser32, "DdeCmpStringHandles")
	ddeConnect = doGetProcAddress(libuser32, "DdeConnect")
	ddeConnectList = doGetProcAddress(libuser32, "DdeConnectList")
	ddeCreateDataHandle = doGetProcAddress(libuser32, "DdeCreateDataHandle")
	ddeCreateStringHandle = doGetProcAddress(libuser32, "DdeCreateStringHandleW")
	ddeDisconnect = doGetProcAddress(libuser32, "DdeDisconnect")
	ddeDisconnectList = doGetProcAddress(libuser32, "DdeDisconnectList")
	ddeEnableCallback = doGetProcAddress(libuser32, "DdeEnableCallback")
	ddeFreeDataHandle = doGetProcAddress(libuser32, "DdeFreeDataHandle")
	ddeFreeStringHandle = doGetProcAddress(libuser32, "DdeFreeStringHandle")
	ddeGetData = doGetProcAddress(libuser32, "DdeGetData")
	ddeGetLastError = doGetProcAddress(libuser32, "DdeGetLastError")
	ddeImpersonateClient = doGetProcAddress(libuser32, "DdeImpersonateClient")
	ddeInitialize = doGetProcAddress(libuser32, "DdeInitializeW")
	ddeKeepStringHandle = doGetProcAddress(libuser32, "DdeKeepStringHandle")
	ddeNameService = doGetProcAddress(libuser32, "DdeNameService")
	ddePostAdvise = doGetProcAddress(libuser32, "DdePostAdvise")
	ddeQueryConvInfo = doGetProcAddress(libuser32, "DdeQueryConvInfo")
	ddeQueryNextServer = doGetProcAddress(libuser32, "DdeQueryNextServer")
	ddeQueryString = doGetProcAddress(libuser32, "DdeQueryStringW")
	ddeReconnect = doGetProcAddress(libuser32, "DdeReconnect")
	ddeSetQualityOfService = doGetProcAddress(libuser32, "DdeSetQualityOfService")
	ddeSetUserHandle = doGetProcAddress(libuser32, "DdeSetUserHandle")
	ddeUnaccessData = doGetProcAddress(libuser32, "DdeUnaccessData")
	ddeUninitialize = doGetProcAddress(libuser32, "DdeUninitialize")
	defDlgProc = doGetProcAddress(libuser32, "DefDlgProcW")
	defFrameProc = doGetProcAddress(libuser32, "DefFrameProcW")
	defMDIChildProc = doGetProcAddress(libuser32, "DefMDIChildProcW")
	defRawInputProc = doGetProcAddress(libuser32, "DefRawInputProc")
	defWindowProc = doGetProcAddress(libuser32, "DefWindowProcW")
	deferWindowPos = doGetProcAddress(libuser32, "DeferWindowPos")
	deleteMenu = doGetProcAddress(libuser32, "DeleteMenu")
	deregisterShellHookWindow = doGetProcAddress(libuser32, "DeregisterShellHookWindow")
	destroyAcceleratorTable = doGetProcAddress(libuser32, "DestroyAcceleratorTable")
	destroyCaret = doGetProcAddress(libuser32, "DestroyCaret")
	destroyCursor = doGetProcAddress(libuser32, "DestroyCursor")
	destroyIcon = doGetProcAddress(libuser32, "DestroyIcon")
	destroyMenu = doGetProcAddress(libuser32, "DestroyMenu")
	destroyWindow = doGetProcAddress(libuser32, "DestroyWindow")
	dialogBoxIndirectParam = doGetProcAddress(libuser32, "DialogBoxIndirectParamW")
	dialogBoxParam = doGetProcAddress(libuser32, "DialogBoxParamW")
	disableProcessWindowsGhosting = doGetProcAddress(libuser32, "DisableProcessWindowsGhosting")
	dispatchMessage = doGetProcAddress(libuser32, "DispatchMessageW")
	dlgDirListComboBox = doGetProcAddress(libuser32, "DlgDirListComboBoxW")
	dlgDirList = doGetProcAddress(libuser32, "DlgDirListW")
	dlgDirSelectComboBoxEx = doGetProcAddress(libuser32, "DlgDirSelectComboBoxExW")
	dlgDirSelectEx = doGetProcAddress(libuser32, "DlgDirSelectExW")
	dragDetect = doGetProcAddress(libuser32, "DragDetect")
	dragObject = doGetProcAddress(libuser32, "DragObject")
	drawAnimatedRects = doGetProcAddress(libuser32, "DrawAnimatedRects")
	drawCaption = doGetProcAddress(libuser32, "DrawCaption")
	drawEdge = doGetProcAddress(libuser32, "DrawEdge")
	drawFocusRect = doGetProcAddress(libuser32, "DrawFocusRect")
	drawFrameControl = doGetProcAddress(libuser32, "DrawFrameControl")
	drawIcon = doGetProcAddress(libuser32, "DrawIcon")
	drawIconEx = doGetProcAddress(libuser32, "DrawIconEx")
	drawMenuBar = doGetProcAddress(libuser32, "DrawMenuBar")
	drawState = doGetProcAddress(libuser32, "DrawStateW")
	drawTextEx = doGetProcAddress(libuser32, "DrawTextExW")
	drawText = doGetProcAddress(libuser32, "DrawTextW")
	emptyClipboard = doGetProcAddress(libuser32, "EmptyClipboard")
	enableMenuItem = doGetProcAddress(libuser32, "EnableMenuItem")
	enableScrollBar = doGetProcAddress(libuser32, "EnableScrollBar")
	enableWindow = doGetProcAddress(libuser32, "EnableWindow")
	endDeferWindowPos = doGetProcAddress(libuser32, "EndDeferWindowPos")
	endDialog = doGetProcAddress(libuser32, "EndDialog")
	endMenu = doGetProcAddress(libuser32, "EndMenu")
	endPaint = doGetProcAddress(libuser32, "EndPaint")
	endTask = doGetProcAddress(libuser32, "EndTask")
	enumChildWindows = doGetProcAddress(libuser32, "EnumChildWindows")
	enumClipboardFormats = doGetProcAddress(libuser32, "EnumClipboardFormats")
	enumDesktopWindows = doGetProcAddress(libuser32, "EnumDesktopWindows")
	enumDesktops = doGetProcAddress(libuser32, "EnumDesktopsW")
	enumDisplayDevices = doGetProcAddress(libuser32, "EnumDisplayDevicesW")
	enumDisplayMonitors = doGetProcAddress(libuser32, "EnumDisplayMonitors")
	enumDisplaySettingsEx = doGetProcAddress(libuser32, "EnumDisplaySettingsExW")
	enumDisplaySettings = doGetProcAddress(libuser32, "EnumDisplaySettingsW")
	enumPropsEx = doGetProcAddress(libuser32, "EnumPropsExW")
	enumProps = doGetProcAddress(libuser32, "EnumPropsW")
	enumThreadWindows = doGetProcAddress(libuser32, "EnumThreadWindows")
	enumWindowStations = doGetProcAddress(libuser32, "EnumWindowStationsW")
	enumWindows = doGetProcAddress(libuser32, "EnumWindows")
	equalRect = doGetProcAddress(libuser32, "EqualRect")
	excludeUpdateRgn = doGetProcAddress(libuser32, "ExcludeUpdateRgn")
	exitWindowsEx = doGetProcAddress(libuser32, "ExitWindowsEx")
	fillRect = doGetProcAddress(libuser32, "FillRect")
	findWindowEx = doGetProcAddress(libuser32, "FindWindowExW")
	findWindow = doGetProcAddress(libuser32, "FindWindowW")
	flashWindow = doGetProcAddress(libuser32, "FlashWindow")
	flashWindowEx = doGetProcAddress(libuser32, "FlashWindowEx")
	frameRect = doGetProcAddress(libuser32, "FrameRect")
	freeDDElParam = doGetProcAddress(libuser32, "FreeDDElParam")
	getActiveWindow = doGetProcAddress(libuser32, "GetActiveWindow")
	getAltTabInfo = doGetProcAddress(libuser32, "GetAltTabInfoW")
	getAncestor = doGetProcAddress(libuser32, "GetAncestor")
	getAsyncKeyState = doGetProcAddress(libuser32, "GetAsyncKeyState")
	getCapture = doGetProcAddress(libuser32, "GetCapture")
	getCaretBlinkTime = doGetProcAddress(libuser32, "GetCaretBlinkTime")
	getCaretPos = doGetProcAddress(libuser32, "GetCaretPos")
	getClassInfoEx = doGetProcAddress(libuser32, "GetClassInfoExW")
	getClassInfo = doGetProcAddress(libuser32, "GetClassInfoW")
	getClassLongPtr = doGetProcAddress(libuser32, "GetClassLongPtrW")
	getClassLong = doGetProcAddress(libuser32, "GetClassLongW")
	getClassName = doGetProcAddress(libuser32, "GetClassNameW")
	getClassWord = doGetProcAddress(libuser32, "GetClassWord")
	getClientRect = doGetProcAddress(libuser32, "GetClientRect")
	getClipCursor = doGetProcAddress(libuser32, "GetClipCursor")
	getClipboardData = doGetProcAddress(libuser32, "GetClipboardData")
	getClipboardFormatName = doGetProcAddress(libuser32, "GetClipboardFormatNameW")
	getClipboardOwner = doGetProcAddress(libuser32, "GetClipboardOwner")
	getClipboardSequenceNumber = doGetProcAddress(libuser32, "GetClipboardSequenceNumber")
	getClipboardViewer = doGetProcAddress(libuser32, "GetClipboardViewer")
	getComboBoxInfo = doGetProcAddress(libuser32, "GetComboBoxInfo")
	getCursor = doGetProcAddress(libuser32, "GetCursor")
	getCursorInfo = doGetProcAddress(libuser32, "GetCursorInfo")
	getCursorPos = doGetProcAddress(libuser32, "GetCursorPos")
	getDC = doGetProcAddress(libuser32, "GetDC")
	getDCEx = doGetProcAddress(libuser32, "GetDCEx")
	getDesktopWindow = doGetProcAddress(libuser32, "GetDesktopWindow")
	getDialogBaseUnits = doGetProcAddress(libuser32, "GetDialogBaseUnits")
	getDlgCtrlID = doGetProcAddress(libuser32, "GetDlgCtrlID")
	getDlgItem = doGetProcAddress(libuser32, "GetDlgItem")
	getDlgItemInt = doGetProcAddress(libuser32, "GetDlgItemInt")
	getDlgItemText = doGetProcAddress(libuser32, "GetDlgItemTextW")
	getDoubleClickTime = doGetProcAddress(libuser32, "GetDoubleClickTime")
	getFocus = doGetProcAddress(libuser32, "GetFocus")
	getForegroundWindow = doGetProcAddress(libuser32, "GetForegroundWindow")
	getGUIThreadInfo = doGetProcAddress(libuser32, "GetGUIThreadInfo")
	getGestureConfig = doGetProcAddress(libuser32, "GetGestureConfig")
	getGestureExtraArgs = doGetProcAddress(libuser32, "GetGestureExtraArgs")
	getGestureInfo = doGetProcAddress(libuser32, "GetGestureInfo")
	getGuiResources = doGetProcAddress(libuser32, "GetGuiResources")
	getIconInfo = doGetProcAddress(libuser32, "GetIconInfo")
	getInputState = doGetProcAddress(libuser32, "GetInputState")
	getKBCodePage = doGetProcAddress(libuser32, "GetKBCodePage")
	getKeyNameText = doGetProcAddress(libuser32, "GetKeyNameTextW")
	getKeyState = doGetProcAddress(libuser32, "GetKeyState")
	getKeyboardLayout = doGetProcAddress(libuser32, "GetKeyboardLayout")
	getKeyboardLayoutList = doGetProcAddress(libuser32, "GetKeyboardLayoutList")
	getKeyboardLayoutName = doGetProcAddress(libuser32, "GetKeyboardLayoutNameW")
	getKeyboardState = doGetProcAddress(libuser32, "GetKeyboardState")
	getKeyboardType = doGetProcAddress(libuser32, "GetKeyboardType")
	getLastActivePopup = doGetProcAddress(libuser32, "GetLastActivePopup")
	getLastInputInfo = doGetProcAddress(libuser32, "GetLastInputInfo")
	getLayeredWindowAttributes = doGetProcAddress(libuser32, "GetLayeredWindowAttributes")
	getListBoxInfo = doGetProcAddress(libuser32, "GetListBoxInfo")
	getMenu = doGetProcAddress(libuser32, "GetMenu")
	getMenuBarInfo = doGetProcAddress(libuser32, "GetMenuBarInfo")
	getMenuCheckMarkDimensions = doGetProcAddress(libuser32, "GetMenuCheckMarkDimensions")
	getMenuContextHelpId = doGetProcAddress(libuser32, "GetMenuContextHelpId")
	getMenuDefaultItem = doGetProcAddress(libuser32, "GetMenuDefaultItem")
	getMenuInfo = doGetProcAddress(libuser32, "GetMenuInfo")
	getMenuItemCount = doGetProcAddress(libuser32, "GetMenuItemCount")
	getMenuItemID = doGetProcAddress(libuser32, "GetMenuItemID")
	getMenuItemInfo = doGetProcAddress(libuser32, "GetMenuItemInfoW")
	getMenuItemRect = doGetProcAddress(libuser32, "GetMenuItemRect")
	getMenuState = doGetProcAddress(libuser32, "GetMenuState")
	getMenuString = doGetProcAddress(libuser32, "GetMenuStringW")
	getMessageExtraInfo = doGetProcAddress(libuser32, "GetMessageExtraInfo")
	getMessagePos = doGetProcAddress(libuser32, "GetMessagePos")
	getMessageTime = doGetProcAddress(libuser32, "GetMessageTime")
	getMessage = doGetProcAddress(libuser32, "GetMessageW")
	getMonitorInfo = doGetProcAddress(libuser32, "GetMonitorInfoW")
	getMouseMovePointsEx = doGetProcAddress(libuser32, "GetMouseMovePointsEx")
	getNextDlgGroupItem = doGetProcAddress(libuser32, "GetNextDlgGroupItem")
	getNextDlgTabItem = doGetProcAddress(libuser32, "GetNextDlgTabItem")
	getOpenClipboardWindow = doGetProcAddress(libuser32, "GetOpenClipboardWindow")
	getParent = doGetProcAddress(libuser32, "GetParent")
	getPriorityClipboardFormat = doGetProcAddress(libuser32, "GetPriorityClipboardFormat")
	getProcessDefaultLayout = doGetProcAddress(libuser32, "GetProcessDefaultLayout")
	getProcessWindowStation = doGetProcAddress(libuser32, "GetProcessWindowStation")
	getProp = doGetProcAddress(libuser32, "GetPropW")
	getQueueStatus = doGetProcAddress(libuser32, "GetQueueStatus")
	getRawInputBuffer = doGetProcAddress(libuser32, "GetRawInputBuffer")
	getRawInputData = doGetProcAddress(libuser32, "GetRawInputData")
	getRawInputDeviceInfo = doGetProcAddress(libuser32, "GetRawInputDeviceInfoW")
	getRawInputDeviceList = doGetProcAddress(libuser32, "GetRawInputDeviceList")
	getRegisteredRawInputDevices = doGetProcAddress(libuser32, "GetRegisteredRawInputDevices")
	getScrollBarInfo = doGetProcAddress(libuser32, "GetScrollBarInfo")
	getScrollInfo = doGetProcAddress(libuser32, "GetScrollInfo")
	getScrollPos = doGetProcAddress(libuser32, "GetScrollPos")
	getScrollRange = doGetProcAddress(libuser32, "GetScrollRange")
	getShellWindow = doGetProcAddress(libuser32, "GetShellWindow")
	getSubMenu = doGetProcAddress(libuser32, "GetSubMenu")
	getSysColor = doGetProcAddress(libuser32, "GetSysColor")
	getSysColorBrush = doGetProcAddress(libuser32, "GetSysColorBrush")
	getSystemMenu = doGetProcAddress(libuser32, "GetSystemMenu")
	getSystemMetrics = doGetProcAddress(libuser32, "GetSystemMetrics")
	getTabbedTextExtent = doGetProcAddress(libuser32, "GetTabbedTextExtentW")
	getThreadDesktop = doGetProcAddress(libuser32, "GetThreadDesktop")
	getTitleBarInfo = doGetProcAddress(libuser32, "GetTitleBarInfo")
	getTopWindow = doGetProcAddress(libuser32, "GetTopWindow")
	getTouchInputInfo = doGetProcAddress(libuser32, "GetTouchInputInfo")
	getUpdateRect = doGetProcAddress(libuser32, "GetUpdateRect")
	getUpdateRgn = doGetProcAddress(libuser32, "GetUpdateRgn")
	getUserObjectInformation = doGetProcAddress(libuser32, "GetUserObjectInformationW")
	getUserObjectSecurity = doGetProcAddress(libuser32, "GetUserObjectSecurity")
	getWindow = doGetProcAddress(libuser32, "GetWindow")
	getWindowContextHelpId = doGetProcAddress(libuser32, "GetWindowContextHelpId")
	getWindowDC = doGetProcAddress(libuser32, "GetWindowDC")
	getWindowInfo = doGetProcAddress(libuser32, "GetWindowInfo")
	getWindowLongPtr = doGetProcAddress(libuser32, "GetWindowLongPtrW")
	getWindowLong = doGetProcAddress(libuser32, "GetWindowLongW")
	getWindowModuleFileName = doGetProcAddress(libuser32, "GetWindowModuleFileNameW")
	getWindowPlacement = doGetProcAddress(libuser32, "GetWindowPlacement")
	getWindowRect = doGetProcAddress(libuser32, "GetWindowRect")
	getWindowRgn = doGetProcAddress(libuser32, "GetWindowRgn")
	getWindowRgnBox = doGetProcAddress(libuser32, "GetWindowRgnBox")
	getWindowTextLength = doGetProcAddress(libuser32, "GetWindowTextLengthW")
	getWindowText = doGetProcAddress(libuser32, "GetWindowTextW")
	getWindowThreadProcessId = doGetProcAddress(libuser32, "GetWindowThreadProcessId")
	getWindowWord = doGetProcAddress(libuser32, "GetWindowWord")
	grayString = doGetProcAddress(libuser32, "GrayStringW")
	hideCaret = doGetProcAddress(libuser32, "HideCaret")
	hiliteMenuItem = doGetProcAddress(libuser32, "HiliteMenuItem")
	iMPGetIME = doGetProcAddress(libuser32, "IMPGetIMEW")
	iMPQueryIME = doGetProcAddress(libuser32, "IMPQueryIMEW")
	iMPSetIME = doGetProcAddress(libuser32, "IMPSetIMEW")
	impersonateDdeClientWindow = doGetProcAddress(libuser32, "ImpersonateDdeClientWindow")
	inSendMessage = doGetProcAddress(libuser32, "InSendMessage")
	inSendMessageEx = doGetProcAddress(libuser32, "InSendMessageEx")
	inflateRect = doGetProcAddress(libuser32, "InflateRect")
	insertMenuItem = doGetProcAddress(libuser32, "InsertMenuItemW")
	insertMenu = doGetProcAddress(libuser32, "InsertMenuW")
	internalGetWindowText = doGetProcAddress(libuser32, "InternalGetWindowText")
	intersectRect = doGetProcAddress(libuser32, "IntersectRect")
	invalidateRect = doGetProcAddress(libuser32, "InvalidateRect")
	invalidateRgn = doGetProcAddress(libuser32, "InvalidateRgn")
	invertRect = doGetProcAddress(libuser32, "InvertRect")
	isCharAlphaNumeric = doGetProcAddress(libuser32, "IsCharAlphaNumericW")
	isCharAlpha = doGetProcAddress(libuser32, "IsCharAlphaW")
	isCharLower = doGetProcAddress(libuser32, "IsCharLowerW")
	isCharUpper = doGetProcAddress(libuser32, "IsCharUpperW")
	isChild = doGetProcAddress(libuser32, "IsChild")
	isClipboardFormatAvailable = doGetProcAddress(libuser32, "IsClipboardFormatAvailable")
	isDialogMessage = doGetProcAddress(libuser32, "IsDialogMessageW")
	isDlgButtonChecked = doGetProcAddress(libuser32, "IsDlgButtonChecked")
	isGUIThread = doGetProcAddress(libuser32, "IsGUIThread")
	isHungAppWindow = doGetProcAddress(libuser32, "IsHungAppWindow")
	isIconic = doGetProcAddress(libuser32, "IsIconic")
	isMenu = doGetProcAddress(libuser32, "IsMenu")
	isRectEmpty = doGetProcAddress(libuser32, "IsRectEmpty")
	isTouchWindow = doGetProcAddress(libuser32, "IsTouchWindow")
	isWinEventHookInstalled = doGetProcAddress(libuser32, "IsWinEventHookInstalled")
	isWindow = doGetProcAddress(libuser32, "IsWindow")
	isWindowEnabled = doGetProcAddress(libuser32, "IsWindowEnabled")
	isWindowUnicode = doGetProcAddress(libuser32, "IsWindowUnicode")
	isWindowVisible = doGetProcAddress(libuser32, "IsWindowVisible")
	isWow64Message = doGetProcAddress(libuser32, "IsWow64Message")
	isZoomed = doGetProcAddress(libuser32, "IsZoomed")
	killTimer = doGetProcAddress(libuser32, "KillTimer")
	loadAccelerators = doGetProcAddress(libuser32, "LoadAcceleratorsW")
	loadBitmap = doGetProcAddress(libuser32, "LoadBitmapW")
	loadCursorFromFile = doGetProcAddress(libuser32, "LoadCursorFromFileW")
	loadCursor = doGetProcAddress(libuser32, "LoadCursorW")
	loadIcon = doGetProcAddress(libuser32, "LoadIconW")
	loadImage = doGetProcAddress(libuser32, "LoadImageW")
	loadKeyboardLayout = doGetProcAddress(libuser32, "LoadKeyboardLayoutW")
	loadMenuIndirect = doGetProcAddress(libuser32, "LoadMenuIndirectW")
	loadMenu = doGetProcAddress(libuser32, "LoadMenuW")
	loadString = doGetProcAddress(libuser32, "LoadStringW")
	lockSetForegroundWindow = doGetProcAddress(libuser32, "LockSetForegroundWindow")
	lockWindowUpdate = doGetProcAddress(libuser32, "LockWindowUpdate")
	lockWorkStation = doGetProcAddress(libuser32, "LockWorkStation")
	lookupIconIdFromDirectory = doGetProcAddress(libuser32, "LookupIconIdFromDirectory")
	lookupIconIdFromDirectoryEx = doGetProcAddress(libuser32, "LookupIconIdFromDirectoryEx")
	mapDialogRect = doGetProcAddress(libuser32, "MapDialogRect")
	mapVirtualKeyEx = doGetProcAddress(libuser32, "MapVirtualKeyExW")
	mapVirtualKey = doGetProcAddress(libuser32, "MapVirtualKeyW")
	mapWindowPoints = doGetProcAddress(libuser32, "MapWindowPoints")
	menuItemFromPoint = doGetProcAddress(libuser32, "MenuItemFromPoint")
	messageBeep = doGetProcAddress(libuser32, "MessageBeep")
	messageBoxEx = doGetProcAddress(libuser32, "MessageBoxExW")
	messageBoxIndirect = doGetProcAddress(libuser32, "MessageBoxIndirectW")
	messageBox = doGetProcAddress(libuser32, "MessageBoxW")
	modifyMenu = doGetProcAddress(libuser32, "ModifyMenuW")
	monitorFromPoint = doGetProcAddress(libuser32, "MonitorFromPoint")
	monitorFromRect = doGetProcAddress(libuser32, "MonitorFromRect")
	monitorFromWindow = doGetProcAddress(libuser32, "MonitorFromWindow")
	moveWindow = doGetProcAddress(libuser32, "MoveWindow")
	msgWaitForMultipleObjects = doGetProcAddress(libuser32, "MsgWaitForMultipleObjects")
	msgWaitForMultipleObjectsEx = doGetProcAddress(libuser32, "MsgWaitForMultipleObjectsEx")
	notifyWinEvent = doGetProcAddress(libuser32, "NotifyWinEvent")
	oemKeyScan = doGetProcAddress(libuser32, "OemKeyScan")
	oemToCharBuff = doGetProcAddress(libuser32, "OemToCharBuffW")
	oemToChar = doGetProcAddress(libuser32, "OemToCharW")
	offsetRect = doGetProcAddress(libuser32, "OffsetRect")
	openClipboard = doGetProcAddress(libuser32, "OpenClipboard")
	openDesktop = doGetProcAddress(libuser32, "OpenDesktopW")
	openIcon = doGetProcAddress(libuser32, "OpenIcon")
	openInputDesktop = doGetProcAddress(libuser32, "OpenInputDesktop")
	openWindowStation = doGetProcAddress(libuser32, "OpenWindowStationW")
	packDDElParam = doGetProcAddress(libuser32, "PackDDElParam")
	paintDesktop = doGetProcAddress(libuser32, "PaintDesktop")
	peekMessage = doGetProcAddress(libuser32, "PeekMessageW")
	postMessage = doGetProcAddress(libuser32, "PostMessageW")
	postQuitMessage = doGetProcAddress(libuser32, "PostQuitMessage")
	postThreadMessage = doGetProcAddress(libuser32, "PostThreadMessageW")
	printWindow = doGetProcAddress(libuser32, "PrintWindow")
	privateExtractIcons = doGetProcAddress(libuser32, "PrivateExtractIconsW")
	ptInRect = doGetProcAddress(libuser32, "PtInRect")
	realChildWindowFromPoint = doGetProcAddress(libuser32, "RealChildWindowFromPoint")
	realGetWindowClass = doGetProcAddress(libuser32, "RealGetWindowClassW")
	redrawWindow = doGetProcAddress(libuser32, "RedrawWindow")
	registerClassEx = doGetProcAddress(libuser32, "RegisterClassExW")
	registerClass = doGetProcAddress(libuser32, "RegisterClassW")
	registerClipboardFormat = doGetProcAddress(libuser32, "RegisterClipboardFormatW")
	registerDeviceNotification = doGetProcAddress(libuser32, "RegisterDeviceNotificationW")
	registerHotKey = doGetProcAddress(libuser32, "RegisterHotKey")
	registerPowerSettingNotification = doGetProcAddress(libuser32, "RegisterPowerSettingNotification")
	registerRawInputDevices = doGetProcAddress(libuser32, "RegisterRawInputDevices")
	registerShellHookWindow = doGetProcAddress(libuser32, "RegisterShellHookWindow")
	registerTouchWindow = doGetProcAddress(libuser32, "RegisterTouchWindow")
	registerWindowMessage = doGetProcAddress(libuser32, "RegisterWindowMessageW")
	releaseCapture = doGetProcAddress(libuser32, "ReleaseCapture")
	releaseDC = doGetProcAddress(libuser32, "ReleaseDC")
	removeMenu = doGetProcAddress(libuser32, "RemoveMenu")
	removeProp = doGetProcAddress(libuser32, "RemovePropW")
	replyMessage = doGetProcAddress(libuser32, "ReplyMessage")
	reuseDDElParam = doGetProcAddress(libuser32, "ReuseDDElParam")
	screenToClient = doGetProcAddress(libuser32, "ScreenToClient")
	scrollDC = doGetProcAddress(libuser32, "ScrollDC")
	scrollWindow = doGetProcAddress(libuser32, "ScrollWindow")
	scrollWindowEx = doGetProcAddress(libuser32, "ScrollWindowEx")
	sendDlgItemMessage = doGetProcAddress(libuser32, "SendDlgItemMessageW")
	sendIMEMessageEx = doGetProcAddress(libuser32, "SendIMEMessageExW")
	sendInput = doGetProcAddress(libuser32, "SendInput")
	sendMessageCallback = doGetProcAddress(libuser32, "SendMessageCallbackW")
	sendMessageTimeout = doGetProcAddress(libuser32, "SendMessageTimeoutW")
	sendMessage = doGetProcAddress(libuser32, "SendMessageW")
	sendNotifyMessage = doGetProcAddress(libuser32, "SendNotifyMessageW")
	setActiveWindow = doGetProcAddress(libuser32, "SetActiveWindow")
	setCapture = doGetProcAddress(libuser32, "SetCapture")
	setCaretBlinkTime = doGetProcAddress(libuser32, "SetCaretBlinkTime")
	setCaretPos = doGetProcAddress(libuser32, "SetCaretPos")
	setClassLongPtr = doGetProcAddress(libuser32, "SetClassLongPtrW")
	setClassLong = doGetProcAddress(libuser32, "SetClassLongW")
	setClassWord = doGetProcAddress(libuser32, "SetClassWord")
	setClipboardData = doGetProcAddress(libuser32, "SetClipboardData")
	setClipboardViewer = doGetProcAddress(libuser32, "SetClipboardViewer")
	setCursor = doGetProcAddress(libuser32, "SetCursor")
	setCursorPos = doGetProcAddress(libuser32, "SetCursorPos")
	setDebugErrorLevel = doGetProcAddress(libuser32, "SetDebugErrorLevel")
	setDlgItemInt = doGetProcAddress(libuser32, "SetDlgItemInt")
	setDlgItemText = doGetProcAddress(libuser32, "SetDlgItemTextW")
	setDoubleClickTime = doGetProcAddress(libuser32, "SetDoubleClickTime")
	setFocus = doGetProcAddress(libuser32, "SetFocus")
	setForegroundWindow = doGetProcAddress(libuser32, "SetForegroundWindow")
	setGestureConfig = doGetProcAddress(libuser32, "SetGestureConfig")
	setKeyboardState = doGetProcAddress(libuser32, "SetKeyboardState")
	setLastErrorEx = doGetProcAddress(libuser32, "SetLastErrorEx")
	setLayeredWindowAttributes = doGetProcAddress(libuser32, "SetLayeredWindowAttributes")
	setMenu = doGetProcAddress(libuser32, "SetMenu")
	setMenuContextHelpId = doGetProcAddress(libuser32, "SetMenuContextHelpId")
	setMenuDefaultItem = doGetProcAddress(libuser32, "SetMenuDefaultItem")
	setMenuInfo = doGetProcAddress(libuser32, "SetMenuInfo")
	setMenuItemBitmaps = doGetProcAddress(libuser32, "SetMenuItemBitmaps")
	setMenuItemInfo = doGetProcAddress(libuser32, "SetMenuItemInfoW")
	setMessageExtraInfo = doGetProcAddress(libuser32, "SetMessageExtraInfo")
	setMessageQueue = doGetProcAddress(libuser32, "SetMessageQueue")
	setParent = doGetProcAddress(libuser32, "SetParent")
	setProcessDefaultLayout = doGetProcAddress(libuser32, "SetProcessDefaultLayout")
	setProcessWindowStation = doGetProcAddress(libuser32, "SetProcessWindowStation")
	setProp = doGetProcAddress(libuser32, "SetPropW")
	setRect = doGetProcAddress(libuser32, "SetRect")
	setRectEmpty = doGetProcAddress(libuser32, "SetRectEmpty")
	setScrollInfo = doGetProcAddress(libuser32, "SetScrollInfo")
	setScrollPos = doGetProcAddress(libuser32, "SetScrollPos")
	setScrollRange = doGetProcAddress(libuser32, "SetScrollRange")
	setSysColors = doGetProcAddress(libuser32, "SetSysColors")
	setSystemCursor = doGetProcAddress(libuser32, "SetSystemCursor")
	setThreadDesktop = doGetProcAddress(libuser32, "SetThreadDesktop")
	setTimer = doGetProcAddress(libuser32, "SetTimer")
	setUserObjectInformation = doGetProcAddress(libuser32, "SetUserObjectInformationW")
	setUserObjectSecurity = doGetProcAddress(libuser32, "SetUserObjectSecurity")
	setWinEventHook = doGetProcAddress(libuser32, "SetWinEventHook")
	setWindowContextHelpId = doGetProcAddress(libuser32, "SetWindowContextHelpId")
	setWindowLongPtr = doGetProcAddress(libuser32, "SetWindowLongPtrW")
	setWindowLong = doGetProcAddress(libuser32, "SetWindowLongW")
	setWindowPlacement = doGetProcAddress(libuser32, "SetWindowPlacement")
	setWindowPos = doGetProcAddress(libuser32, "SetWindowPos")
	setWindowRgn = doGetProcAddress(libuser32, "SetWindowRgn")
	setWindowText = doGetProcAddress(libuser32, "SetWindowTextW")
	setWindowWord = doGetProcAddress(libuser32, "SetWindowWord")
	setWindowsHookEx = doGetProcAddress(libuser32, "SetWindowsHookExW")
	setWindowsHook = doGetProcAddress(libuser32, "SetWindowsHookW")
	showCaret = doGetProcAddress(libuser32, "ShowCaret")
	showCursor = doGetProcAddress(libuser32, "ShowCursor")
	showOwnedPopups = doGetProcAddress(libuser32, "ShowOwnedPopups")
	showScrollBar = doGetProcAddress(libuser32, "ShowScrollBar")
	showWindow = doGetProcAddress(libuser32, "ShowWindow")
	showWindowAsync = doGetProcAddress(libuser32, "ShowWindowAsync")
	subtractRect = doGetProcAddress(libuser32, "SubtractRect")
	swapMouseButton = doGetProcAddress(libuser32, "SwapMouseButton")
	switchDesktop = doGetProcAddress(libuser32, "SwitchDesktop")
	switchToThisWindow = doGetProcAddress(libuser32, "SwitchToThisWindow")
	systemParametersInfo = doGetProcAddress(libuser32, "SystemParametersInfoW")
	tabbedTextOut = doGetProcAddress(libuser32, "TabbedTextOutW")
	tileWindows = doGetProcAddress(libuser32, "TileWindows")
	toAscii = doGetProcAddress(libuser32, "ToAscii")
	toAsciiEx = doGetProcAddress(libuser32, "ToAsciiEx")
	toUnicode = doGetProcAddress(libuser32, "ToUnicode")
	toUnicodeEx = doGetProcAddress(libuser32, "ToUnicodeEx")
	trackMouseEvent = doGetProcAddress(libuser32, "TrackMouseEvent")
	trackPopupMenu = doGetProcAddress(libuser32, "TrackPopupMenu")
	trackPopupMenuEx = doGetProcAddress(libuser32, "TrackPopupMenuEx")
	translateAccelerator = doGetProcAddress(libuser32, "TranslateAcceleratorW")
	translateMDISysAccel = doGetProcAddress(libuser32, "TranslateMDISysAccel")
	translateMessage = doGetProcAddress(libuser32, "TranslateMessage")
	unhookWinEvent = doGetProcAddress(libuser32, "UnhookWinEvent")
	unhookWindowsHook = doGetProcAddress(libuser32, "UnhookWindowsHook")
	unhookWindowsHookEx = doGetProcAddress(libuser32, "UnhookWindowsHookEx")
	unionRect = doGetProcAddress(libuser32, "UnionRect")
	unloadKeyboardLayout = doGetProcAddress(libuser32, "UnloadKeyboardLayout")
	unpackDDElParam = doGetProcAddress(libuser32, "UnpackDDElParam")
	unregisterClass = doGetProcAddress(libuser32, "UnregisterClassW")
	unregisterDeviceNotification = doGetProcAddress(libuser32, "UnregisterDeviceNotification")
	unregisterHotKey = doGetProcAddress(libuser32, "UnregisterHotKey")
	unregisterPowerSettingNotification = doGetProcAddress(libuser32, "UnregisterPowerSettingNotification")
	unregisterTouchWindow = doGetProcAddress(libuser32, "UnregisterTouchWindow")
	updateLayeredWindow = doGetProcAddress(libuser32, "UpdateLayeredWindow")
	updateLayeredWindowIndirect = doGetProcAddress(libuser32, "UpdateLayeredWindowIndirect")
	updateWindow = doGetProcAddress(libuser32, "UpdateWindow")
	userHandleGrantAccess = doGetProcAddress(libuser32, "UserHandleGrantAccess")
	validateRect = doGetProcAddress(libuser32, "ValidateRect")
	validateRgn = doGetProcAddress(libuser32, "ValidateRgn")
	vkKeyScanEx = doGetProcAddress(libuser32, "VkKeyScanExW")
	vkKeyScan = doGetProcAddress(libuser32, "VkKeyScanW")
	wINNLSEnableIME = doGetProcAddress(libuser32, "WINNLSEnableIME")
	wINNLSGetEnableStatus = doGetProcAddress(libuser32, "WINNLSGetEnableStatus")
	wINNLSGetIMEHotkey = doGetProcAddress(libuser32, "WINNLSGetIMEHotkey")
	waitForInputIdle = doGetProcAddress(libuser32, "WaitForInputIdle")
	waitMessage = doGetProcAddress(libuser32, "WaitMessage")
	winHelp = doGetProcAddress(libuser32, "WinHelpW")
	windowFromDC = doGetProcAddress(libuser32, "WindowFromDC")
	windowFromPoint = doGetProcAddress(libuser32, "WindowFromPoint")
	keybd_event = doGetProcAddress(libuser32, "keybd_event")
	mouse_event = doGetProcAddress(libuser32, "mouse_event")
	alignRects = doGetProcAddress(libuser32, "AlignRects")
	cascadeChildWindows = doGetProcAddress(libuser32, "CascadeChildWindows")
	createDialogIndirectParamAorW = doGetProcAddress(libuser32, "CreateDialogIndirectParamAorW")
	dialogBoxIndirectParamAorW = doGetProcAddress(libuser32, "DialogBoxIndirectParamAorW")
	drawCaptionTemp = doGetProcAddress(libuser32, "DrawCaptionTempW")
	drawMenuBarTemp = doGetProcAddress(libuser32, "DrawMenuBarTemp")
	getAppCompatFlags = doGetProcAddress(libuser32, "GetAppCompatFlags")
	getAppCompatFlags2 = doGetProcAddress(libuser32, "GetAppCompatFlags2")
	getCursorFrameInfo = doGetProcAddress(libuser32, "GetCursorFrameInfo")
	getInternalWindowPos = doGetProcAddress(libuser32, "GetInternalWindowPos")
	getProgmanWindow = doGetProcAddress(libuser32, "GetProgmanWindow")
	getTaskmanWindow = doGetProcAddress(libuser32, "GetTaskmanWindow")
	killSystemTimer = doGetProcAddress(libuser32, "KillSystemTimer")
	loadLocalFonts = doGetProcAddress(libuser32, "LoadLocalFonts")
	messageBoxTimeout = doGetProcAddress(libuser32, "MessageBoxTimeoutW")
	privateExtractIconEx = doGetProcAddress(libuser32, "PrivateExtractIconExW")
	registerLogonProcess = doGetProcAddress(libuser32, "RegisterLogonProcess")
	registerServicesProcess = doGetProcAddress(libuser32, "RegisterServicesProcess")
	registerSystemThread = doGetProcAddress(libuser32, "RegisterSystemThread")
	registerTasklist = doGetProcAddress(libuser32, "RegisterTasklist")
	scrollChildren = doGetProcAddress(libuser32, "ScrollChildren")
	setInternalWindowPos = doGetProcAddress(libuser32, "SetInternalWindowPos")
	setLogonNotifyWindow = doGetProcAddress(libuser32, "SetLogonNotifyWindow")
	setProgmanWindow = doGetProcAddress(libuser32, "SetProgmanWindow")
	setShellWindow = doGetProcAddress(libuser32, "SetShellWindow")
	setShellWindowEx = doGetProcAddress(libuser32, "SetShellWindowEx")
	setSysColorsTemp = doGetProcAddress(libuser32, "SetSysColorsTemp")
	setSystemMenu = doGetProcAddress(libuser32, "SetSystemMenu")
	setSystemTimer = doGetProcAddress(libuser32, "SetSystemTimer")
	setTaskmanWindow = doGetProcAddress(libuser32, "SetTaskmanWindow")
	setWindowStationUser = doGetProcAddress(libuser32, "SetWindowStationUser")
	tileChildWindows = doGetProcAddress(libuser32, "TileChildWindows")
	user32InitializeImmEntryTable = doGetProcAddress(libuser32, "User32InitializeImmEntryTable")
	userRealizePalette = doGetProcAddress(libuser32, "UserRealizePalette")
	userRegisterWowHandlers = doGetProcAddress(libuser32, "UserRegisterWowHandlers")
}

func ActivateKeyboardLayout(hkl HKL, flags uint32) HKL {
	ret1 := syscall3(activateKeyboardLayout, 2,
		uintptr(hkl),
		uintptr(flags),
		0)
	return HKL(ret1)
}

func AdjustWindowRect(lpRect *RECT, dwStyle uint32, bMenu bool) bool {
	ret1 := syscall3(adjustWindowRect, 3,
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(dwStyle),
		getUintptrFromBool(bMenu))
	return ret1 != 0
}

func AdjustWindowRectEx(lpRect *RECT, dwStyle uint32, bMenu bool, dwExStyle uint32) bool {
	ret1 := syscall6(adjustWindowRectEx, 4,
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(dwStyle),
		getUintptrFromBool(bMenu),
		uintptr(dwExStyle),
		0,
		0)
	return ret1 != 0
}

func AllowSetForegroundWindow(dwProcessId uint32) bool {
	ret1 := syscall3(allowSetForegroundWindow, 1,
		uintptr(dwProcessId),
		0,
		0)
	return ret1 != 0
}

func AnimateWindow(hWnd HWND, dwTime uint32, dwFlags uint32) bool {
	ret1 := syscall3(animateWindow, 3,
		uintptr(hWnd),
		uintptr(dwTime),
		uintptr(dwFlags))
	return ret1 != 0
}

func AnyPopup() bool {
	ret1 := syscall3(anyPopup, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func AppendMenu(hMenu HMENU, uFlags uint32, uIDNewItem *uint32, lpNewItem string) bool {
	lpNewItemStr := unicode16FromString(lpNewItem)
	ret1 := syscall6(appendMenu, 4,
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(uIDNewItem)),
		uintptr(unsafe.Pointer(&lpNewItemStr[0])),
		0,
		0)
	return ret1 != 0
}

func ArrangeIconicWindows(hWnd HWND) uint32 {
	ret1 := syscall3(arrangeIconicWindows, 1,
		uintptr(hWnd),
		0,
		0)
	return uint32(ret1)
}

func AttachThreadInput(idAttach uint32, idAttachTo uint32, fAttach bool) bool {
	ret1 := syscall3(attachThreadInput, 3,
		uintptr(idAttach),
		uintptr(idAttachTo),
		getUintptrFromBool(fAttach))
	return ret1 != 0
}

func BeginDeferWindowPos(nNumWindows int32) HDWP {
	ret1 := syscall3(beginDeferWindowPos, 1,
		uintptr(nNumWindows),
		0,
		0)
	return HDWP(ret1)
}

func BeginPaint(hWnd HWND, lpPaint *PAINTSTRUCT) HDC {
	ret1 := syscall3(beginPaint, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPaint)),
		0)
	return HDC(ret1)
}

func BlockInput(fBlockIt bool) bool {
	ret1 := syscall3(blockInput, 1,
		getUintptrFromBool(fBlockIt),
		0,
		0)
	return ret1 != 0
}

func BringWindowToTop(hWnd HWND) bool {
	ret1 := syscall3(bringWindowToTop, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func BroadcastSystemMessageEx(flags uint32, lpInfo *uint32, msg uint32, wParam WPARAM, lParam LPARAM, pbsmInfo *BSMINFO) int32 {
	ret1 := syscall6(broadcastSystemMessageEx, 6,
		uintptr(flags),
		uintptr(unsafe.Pointer(lpInfo)),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		uintptr(unsafe.Pointer(pbsmInfo)))
	return int32(ret1)
}

func BroadcastSystemMessage(flags uint32, lpInfo *uint32, msg uint32, wParam WPARAM, lParam LPARAM) int32 {
	ret1 := syscall6(broadcastSystemMessage, 5,
		uintptr(flags),
		uintptr(unsafe.Pointer(lpInfo)),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0)
	return int32(ret1)
}

func CallMsgFilter(lpMsg *MSG, nCode int32) bool {
	ret1 := syscall3(callMsgFilter, 2,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(nCode),
		0)
	return ret1 != 0
}

func CallNextHookEx(hhk HHOOK, nCode int32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(callNextHookEx, 4,
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return LRESULT(ret1)
}

func CallWindowProc(lpPrevWndFunc WNDPROC, hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	lpPrevWndFuncCallback := syscall.NewCallback(func(unnamed0RawArg HWND, unnamed1RawArg UINT, unnamed2RawArg WPARAM, unnamed3RawArg LPARAM) uintptr {
		ret := lpPrevWndFunc(unnamed0RawArg, unnamed1RawArg, unnamed2RawArg, unnamed3RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(callWindowProc, 5,
		lpPrevWndFuncCallback,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0)
	return LRESULT(ret1)
}

func CascadeWindows(hwndParent HWND, wHow uint32, lpRect /*const*/ *RECT, cKids uint32, lpKids /*const*/ *HWND) uint16 {
	ret1 := syscall6(cascadeWindows, 5,
		uintptr(hwndParent),
		uintptr(wHow),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(cKids),
		uintptr(unsafe.Pointer(lpKids)),
		0)
	return uint16(ret1)
}

func ChangeClipboardChain(hWndRemove HWND, hWndNewNext HWND) bool {
	ret1 := syscall3(changeClipboardChain, 2,
		uintptr(hWndRemove),
		uintptr(hWndNewNext),
		0)
	return ret1 != 0
}

func ChangeDisplaySettingsEx(lpszDeviceName string, lpDevMode LPDEVMODE, hwnd HWND, dwflags uint32, lParam LPVOID) LONG {
	lpszDeviceNameStr := unicode16FromString(lpszDeviceName)
	ret1 := syscall6(changeDisplaySettingsEx, 5,
		uintptr(unsafe.Pointer(&lpszDeviceNameStr[0])),
		uintptr(unsafe.Pointer(lpDevMode)),
		uintptr(hwnd),
		uintptr(dwflags),
		uintptr(unsafe.Pointer(lParam)),
		0)
	return LONG(ret1)
}

func ChangeDisplaySettings(lpDevMode LPDEVMODE, dwFlags uint32) LONG {
	ret1 := syscall3(changeDisplaySettings, 2,
		uintptr(unsafe.Pointer(lpDevMode)),
		uintptr(dwFlags),
		0)
	return LONG(ret1)
}

func ChangeMenu(hMenu HMENU, cmd uint32, lpszNewItem string, cmdInsert uint32, flags uint32) bool {
	lpszNewItemStr := unicode16FromString(lpszNewItem)
	ret1 := syscall6(changeMenu, 5,
		uintptr(hMenu),
		uintptr(cmd),
		uintptr(unsafe.Pointer(&lpszNewItemStr[0])),
		uintptr(cmdInsert),
		uintptr(flags),
		0)
	return ret1 != 0
}

func CharLowerBuff(lpsz LPWSTR, cchLength uint32) uint32 {
	ret1 := syscall3(charLowerBuff, 2,
		uintptr(unsafe.Pointer(lpsz)),
		uintptr(cchLength),
		0)
	return uint32(ret1)
}

func CharLower(lpsz LPWSTR) LPWSTR {
	ret1 := syscall3(charLower, 1,
		uintptr(unsafe.Pointer(lpsz)),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func CharNextExA(codePage uint16, lpCurrentChar /*const*/ LPCSTR, dwFlags uint32) LPSTR {
	ret1 := syscall3(charNextExA, 3,
		uintptr(codePage),
		uintptr(unsafe.Pointer(lpCurrentChar)),
		uintptr(dwFlags))
	return (LPSTR)(unsafe.Pointer(ret1))
}

func CharNext(lpsz string) LPWSTR {
	lpszStr := unicode16FromString(lpsz)
	ret1 := syscall3(charNext, 1,
		uintptr(unsafe.Pointer(&lpszStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func CharPrevExA(codePage uint16, lpStart /*const*/ LPCSTR, lpCurrentChar /*const*/ LPCSTR, dwFlags uint32) LPSTR {
	ret1 := syscall6(charPrevExA, 4,
		uintptr(codePage),
		uintptr(unsafe.Pointer(lpStart)),
		uintptr(unsafe.Pointer(lpCurrentChar)),
		uintptr(dwFlags),
		0,
		0)
	return (LPSTR)(unsafe.Pointer(ret1))
}

func CharPrev(lpszStart string, lpszCurrent string) LPWSTR {
	lpszStartStr := unicode16FromString(lpszStart)
	lpszCurrentStr := unicode16FromString(lpszCurrent)
	ret1 := syscall3(charPrev, 2,
		uintptr(unsafe.Pointer(&lpszStartStr[0])),
		uintptr(unsafe.Pointer(&lpszCurrentStr[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func CharToOemBuff(lpszSrc string, lpszDst LPSTR, cchDstLength uint32) bool {
	lpszSrcStr := unicode16FromString(lpszSrc)
	ret1 := syscall3(charToOemBuff, 3,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(lpszDst)),
		uintptr(cchDstLength))
	return ret1 != 0
}

func CharToOem(lpszSrc string, lpszDst LPSTR) bool {
	lpszSrcStr := unicode16FromString(lpszSrc)
	ret1 := syscall3(charToOem, 2,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(lpszDst)),
		0)
	return ret1 != 0
}

func CharUpperBuff(lpsz LPWSTR, cchLength uint32) uint32 {
	ret1 := syscall3(charUpperBuff, 2,
		uintptr(unsafe.Pointer(lpsz)),
		uintptr(cchLength),
		0)
	return uint32(ret1)
}

func CharUpper(lpsz LPWSTR) LPWSTR {
	ret1 := syscall3(charUpper, 1,
		uintptr(unsafe.Pointer(lpsz)),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func CheckDlgButton(hDlg HWND, nIDButton int32, uCheck uint32) bool {
	ret1 := syscall3(checkDlgButton, 3,
		uintptr(hDlg),
		uintptr(nIDButton),
		uintptr(uCheck))
	return ret1 != 0
}

func CheckMenuItem(hMenu HMENU, uIDCheckItem uint32, uCheck uint32) uint32 {
	ret1 := syscall3(checkMenuItem, 3,
		uintptr(hMenu),
		uintptr(uIDCheckItem),
		uintptr(uCheck))
	return uint32(ret1)
}

func CheckMenuRadioItem(hmenu HMENU, first uint32, last uint32, check uint32, flags uint32) bool {
	ret1 := syscall6(checkMenuRadioItem, 5,
		uintptr(hmenu),
		uintptr(first),
		uintptr(last),
		uintptr(check),
		uintptr(flags),
		0)
	return ret1 != 0
}

func CheckRadioButton(hDlg HWND, nIDFirstButton int32, nIDLastButton int32, nIDCheckButton int32) bool {
	ret1 := syscall6(checkRadioButton, 4,
		uintptr(hDlg),
		uintptr(nIDFirstButton),
		uintptr(nIDLastButton),
		uintptr(nIDCheckButton),
		0,
		0)
	return ret1 != 0
}

func ChildWindowFromPoint(hWndParent HWND, point POINT) HWND {
	ret1 := syscall3(childWindowFromPoint, 3,
		uintptr(hWndParent),
		uintptr(point.X),
		uintptr(point.Y))
	return HWND(ret1)
}

func ChildWindowFromPointEx(hwnd HWND, pt POINT, flags uint32) HWND {
	ret1 := syscall6(childWindowFromPointEx, 4,
		uintptr(hwnd),
		uintptr(pt.X),
		uintptr(pt.Y),
		uintptr(flags),
		0,
		0)
	return HWND(ret1)
}

func ClientToScreen(hWnd HWND, lpPoint *POINT) bool {
	ret1 := syscall3(clientToScreen, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPoint)),
		0)
	return ret1 != 0
}

func ClipCursor(lpRect /*const*/ *RECT) bool {
	ret1 := syscall3(clipCursor, 1,
		uintptr(unsafe.Pointer(lpRect)),
		0,
		0)
	return ret1 != 0
}

func CloseClipboard() bool {
	ret1 := syscall3(closeClipboard, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func CloseDesktop(hDesktop HDESK) bool {
	ret1 := syscall3(closeDesktop, 1,
		uintptr(hDesktop),
		0,
		0)
	return ret1 != 0
}

func CloseGestureInfoHandle(hGestureInfo HGESTUREINFO) bool {
	ret1 := syscall3(closeGestureInfoHandle, 1,
		uintptr(hGestureInfo),
		0,
		0)
	return ret1 != 0
}

func CloseTouchInputHandle(hTouchInput HTOUCHINPUT) bool {
	ret1 := syscall3(closeTouchInputHandle, 1,
		uintptr(hTouchInput),
		0,
		0)
	return ret1 != 0
}

func CloseWindow(hWnd HWND) bool {
	ret1 := syscall3(closeWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func CloseWindowStation(hWinSta HWINSTA) bool {
	ret1 := syscall3(closeWindowStation, 1,
		uintptr(hWinSta),
		0,
		0)
	return ret1 != 0
}

func CopyAcceleratorTable(hAccelSrc HACCEL, lpAccelDst *ACCEL, cAccelEntries int32) int32 {
	ret1 := syscall3(copyAcceleratorTable, 3,
		uintptr(hAccelSrc),
		uintptr(unsafe.Pointer(lpAccelDst)),
		uintptr(cAccelEntries))
	return int32(ret1)
}

func CopyIcon(hIcon HICON) HICON {
	ret1 := syscall3(copyIcon, 1,
		uintptr(hIcon),
		0,
		0)
	return HICON(ret1)
}

func CopyImage(h HANDLE, aType uint32, cx int32, cy int32, flags uint32) HANDLE {
	ret1 := syscall6(copyImage, 5,
		uintptr(h),
		uintptr(aType),
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
		0)
	return HANDLE(ret1)
}

func CopyRect(lprcDst *RECT, lprcSrc /*const*/ *RECT) bool {
	ret1 := syscall3(copyRect, 2,
		uintptr(unsafe.Pointer(lprcDst)),
		uintptr(unsafe.Pointer(lprcSrc)),
		0)
	return ret1 != 0
}

func CountClipboardFormats() int32 {
	ret1 := syscall3(countClipboardFormats, 0,
		0,
		0,
		0)
	return int32(ret1)
}

func CreateAcceleratorTable(paccel *ACCEL, cAccel int32) HACCEL {
	ret1 := syscall3(createAcceleratorTable, 2,
		uintptr(unsafe.Pointer(paccel)),
		uintptr(cAccel),
		0)
	return HACCEL(ret1)
}

func CreateCaret(hWnd HWND, hBitmap HBITMAP, nWidth int32, nHeight int32) bool {
	ret1 := syscall6(createCaret, 4,
		uintptr(hWnd),
		uintptr(hBitmap),
		uintptr(nWidth),
		uintptr(nHeight),
		0,
		0)
	return ret1 != 0
}

func CreateCursor(hInst HINSTANCE, xHotSpot int32, yHotSpot int32, nWidth int32, nHeight int32, pvANDPlane /*const*/ uintptr, pvXORPlane /*const*/ uintptr) HCURSOR {
	ret1 := syscall9(createCursor, 7,
		uintptr(hInst),
		uintptr(xHotSpot),
		uintptr(yHotSpot),
		uintptr(nWidth),
		uintptr(nHeight),
		pvANDPlane,
		pvXORPlane,
		0,
		0)
	return HCURSOR(ret1)
}

func CreateDesktop(lpszDesktop string, lpszDevice string, pDevmode LPDEVMODE, dwFlags uint32, dwDesiredAccess ACCESS_MASK, lpsa *SECURITY_ATTRIBUTES) HDESK {
	lpszDesktopStr := unicode16FromString(lpszDesktop)
	lpszDeviceStr := unicode16FromString(lpszDevice)
	ret1 := syscall6(createDesktop, 6,
		uintptr(unsafe.Pointer(&lpszDesktopStr[0])),
		uintptr(unsafe.Pointer(&lpszDeviceStr[0])),
		uintptr(unsafe.Pointer(pDevmode)),
		uintptr(dwFlags),
		uintptr(dwDesiredAccess),
		uintptr(unsafe.Pointer(lpsa)))
	return HDESK(ret1)
}

func CreateDialogIndirectParam(hInstance HINSTANCE, lpTemplate /*const*/ *DLGTEMPLATE, hWndParent HWND, lpDialogFunc DLGPROC, dwInitParam LPARAM) HWND {
	lpDialogFuncCallback := syscall.NewCallback(func(hwndDlgRawArg HWND, uMsgRawArg uint32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := lpDialogFunc(hwndDlgRawArg, uMsgRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(createDialogIndirectParam, 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplate)),
		uintptr(hWndParent),
		lpDialogFuncCallback,
		uintptr(dwInitParam),
		0)
	return HWND(ret1)
}

func CreateDialogParam(hInstance HINSTANCE, lpTemplateName string, hWndParent HWND, lpDialogFunc DLGPROC, dwInitParam LPARAM) HWND {
	lpTemplateNameStr := unicode16FromString(lpTemplateName)
	lpDialogFuncCallback := syscall.NewCallback(func(hwndDlgRawArg HWND, uMsgRawArg uint32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := lpDialogFunc(hwndDlgRawArg, uMsgRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(createDialogParam, 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpTemplateNameStr[0])),
		uintptr(hWndParent),
		lpDialogFuncCallback,
		uintptr(dwInitParam),
		0)
	return HWND(ret1)
}

func CreateIcon(hInstance HINSTANCE, nWidth int32, nHeight int32, cPlanes byte, cBitsPixel byte, lpbANDbits /*const*/ *byte, lpbXORbits /*const*/ *byte) HICON {
	ret1 := syscall9(createIcon, 7,
		uintptr(hInstance),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(cPlanes),
		uintptr(cBitsPixel),
		uintptr(unsafe.Pointer(lpbANDbits)),
		uintptr(unsafe.Pointer(lpbXORbits)),
		0,
		0)
	return HICON(ret1)
}

func CreateIconFromResource(presbits *byte, dwResSize uint32, fIcon bool, dwVer uint32) HICON {
	ret1 := syscall6(createIconFromResource, 4,
		uintptr(unsafe.Pointer(presbits)),
		uintptr(dwResSize),
		getUintptrFromBool(fIcon),
		uintptr(dwVer),
		0,
		0)
	return HICON(ret1)
}

func CreateIconFromResourceEx(presbits *byte, dwResSize uint32, fIcon bool, dwVer uint32, cxDesired int32, cyDesired int32, flags uint32) HICON {
	ret1 := syscall9(createIconFromResourceEx, 7,
		uintptr(unsafe.Pointer(presbits)),
		uintptr(dwResSize),
		getUintptrFromBool(fIcon),
		uintptr(dwVer),
		uintptr(cxDesired),
		uintptr(cyDesired),
		uintptr(flags),
		0,
		0)
	return HICON(ret1)
}

func CreateIconIndirect(piconinfo *ICONINFO) HICON {
	ret1 := syscall3(createIconIndirect, 1,
		uintptr(unsafe.Pointer(piconinfo)),
		0,
		0)
	return HICON(ret1)
}

func CreateMDIWindow(lpClassName string, lpWindowName string, dwStyle uint32, x int32, y int32, nWidth int32, nHeight int32, hWndParent HWND, hInstance HINSTANCE, lParam LPARAM) HWND {
	lpClassNameStr := unicode16FromString(lpClassName)
	lpWindowNameStr := unicode16FromString(lpWindowName)
	ret1 := syscall12(createMDIWindow, 10,
		uintptr(unsafe.Pointer(&lpClassNameStr[0])),
		uintptr(unsafe.Pointer(&lpWindowNameStr[0])),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hInstance),
		uintptr(lParam),
		0,
		0)
	return HWND(ret1)
}

func CreateMenu() HMENU {
	ret1 := syscall3(createMenu, 0,
		0,
		0,
		0)
	return HMENU(ret1)
}

func CreatePopupMenu() HMENU {
	ret1 := syscall3(createPopupMenu, 0,
		0,
		0,
		0)
	return HMENU(ret1)
}

func CreateWindowEx(dwExStyle uint32, lpClassName string, lpWindowName string, dwStyle uint32, x int32, y int32, nWidth int32, nHeight int32, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) HWND {
	lpClassNameStr := unicode16FromString(lpClassName)
	lpWindowNameStr := unicode16FromString(lpWindowName)
	ret1 := syscall12(createWindowEx, 12,
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(&lpClassNameStr[0])),
		uintptr(unsafe.Pointer(&lpWindowNameStr[0])),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpParam)))
	return HWND(ret1)
}

func CreateWindowStation(lpwinsta string, dwFlags uint32, dwDesiredAccess ACCESS_MASK, lpsa *SECURITY_ATTRIBUTES) HWINSTA {
	lpwinstaStr := unicode16FromString(lpwinsta)
	ret1 := syscall6(createWindowStation, 4,
		uintptr(unsafe.Pointer(&lpwinstaStr[0])),
		uintptr(dwFlags),
		uintptr(dwDesiredAccess),
		uintptr(unsafe.Pointer(lpsa)),
		0,
		0)
	return HWINSTA(ret1)
}

func DdeAbandonTransaction(idInst uint32, hConv HCONV, idTransaction uint32) bool {
	ret1 := syscall3(ddeAbandonTransaction, 3,
		uintptr(idInst),
		uintptr(hConv),
		uintptr(idTransaction))
	return ret1 != 0
}

func DdeAccessData(hData HDDEDATA, pcbDataSize *uint32) *byte {
	ret1 := syscall3(ddeAccessData, 2,
		uintptr(hData),
		uintptr(unsafe.Pointer(pcbDataSize)),
		0)
	return (*byte)(unsafe.Pointer(ret1))
}

func DdeAddData(hData HDDEDATA, pSrc *byte, cb uint32, cbOff uint32) HDDEDATA {
	ret1 := syscall6(ddeAddData, 4,
		uintptr(hData),
		uintptr(unsafe.Pointer(pSrc)),
		uintptr(cb),
		uintptr(cbOff),
		0,
		0)
	return HDDEDATA(ret1)
}

func DdeClientTransaction(pData *byte, cbData uint32, hConv HCONV, hszItem HSZ, wFmt uint32, wType uint32, dwTimeout uint32, pdwResult *uint32) HDDEDATA {
	ret1 := syscall9(ddeClientTransaction, 8,
		uintptr(unsafe.Pointer(pData)),
		uintptr(cbData),
		uintptr(hConv),
		uintptr(hszItem),
		uintptr(wFmt),
		uintptr(wType),
		uintptr(dwTimeout),
		uintptr(unsafe.Pointer(pdwResult)),
		0)
	return HDDEDATA(ret1)
}

func DdeCmpStringHandles(hsz1 HSZ, hsz2 HSZ) int32 {
	ret1 := syscall3(ddeCmpStringHandles, 2,
		uintptr(hsz1),
		uintptr(hsz2),
		0)
	return int32(ret1)
}

func DdeConnect(idInst uint32, hszService HSZ, hszTopic HSZ, pCC *CONVCONTEXT) HCONV {
	ret1 := syscall6(ddeConnect, 4,
		uintptr(idInst),
		uintptr(hszService),
		uintptr(hszTopic),
		uintptr(unsafe.Pointer(pCC)),
		0,
		0)
	return HCONV(ret1)
}

func DdeConnectList(idInst uint32, hszService HSZ, hszTopic HSZ, hConvList HCONVLIST, pCC *CONVCONTEXT) HCONVLIST {
	ret1 := syscall6(ddeConnectList, 5,
		uintptr(idInst),
		uintptr(hszService),
		uintptr(hszTopic),
		uintptr(hConvList),
		uintptr(unsafe.Pointer(pCC)),
		0)
	return HCONVLIST(ret1)
}

func DdeCreateDataHandle(idInst uint32, pSrc *byte, cb uint32, cbOff uint32, hszItem HSZ, wFmt uint32, afCmd uint32) HDDEDATA {
	ret1 := syscall9(ddeCreateDataHandle, 7,
		uintptr(idInst),
		uintptr(unsafe.Pointer(pSrc)),
		uintptr(cb),
		uintptr(cbOff),
		uintptr(hszItem),
		uintptr(wFmt),
		uintptr(afCmd),
		0,
		0)
	return HDDEDATA(ret1)
}

func DdeCreateStringHandle(idInst uint32, psz string, iCodePage int32) HSZ {
	pszStr := unicode16FromString(psz)
	ret1 := syscall3(ddeCreateStringHandle, 3,
		uintptr(idInst),
		uintptr(unsafe.Pointer(&pszStr[0])),
		uintptr(iCodePage))
	return HSZ(ret1)
}

func DdeDisconnect(hConv HCONV) bool {
	ret1 := syscall3(ddeDisconnect, 1,
		uintptr(hConv),
		0,
		0)
	return ret1 != 0
}

func DdeDisconnectList(hConvList HCONVLIST) bool {
	ret1 := syscall3(ddeDisconnectList, 1,
		uintptr(hConvList),
		0,
		0)
	return ret1 != 0
}

func DdeEnableCallback(idInst uint32, hConv HCONV, wCmd uint32) bool {
	ret1 := syscall3(ddeEnableCallback, 3,
		uintptr(idInst),
		uintptr(hConv),
		uintptr(wCmd))
	return ret1 != 0
}

func DdeFreeDataHandle(hData HDDEDATA) bool {
	ret1 := syscall3(ddeFreeDataHandle, 1,
		uintptr(hData),
		0,
		0)
	return ret1 != 0
}

func DdeFreeStringHandle(idInst uint32, hsz HSZ) bool {
	ret1 := syscall3(ddeFreeStringHandle, 2,
		uintptr(idInst),
		uintptr(hsz),
		0)
	return ret1 != 0
}

func DdeGetData(hData HDDEDATA, pDst *byte, cbMax uint32, cbOff uint32) uint32 {
	ret1 := syscall6(ddeGetData, 4,
		uintptr(hData),
		uintptr(unsafe.Pointer(pDst)),
		uintptr(cbMax),
		uintptr(cbOff),
		0,
		0)
	return uint32(ret1)
}

func DdeGetLastError(idInst uint32) uint32 {
	ret1 := syscall3(ddeGetLastError, 1,
		uintptr(idInst),
		0,
		0)
	return uint32(ret1)
}

func DdeImpersonateClient(hConv HCONV) bool {
	ret1 := syscall3(ddeImpersonateClient, 1,
		uintptr(hConv),
		0,
		0)
	return ret1 != 0
}

func DdeInitialize(pidInst *uint32, pfnCallback PFNCALLBACK, afCmd uint32, ulRes uint32) uint32 {
	pfnCallbackCallback := syscall.NewCallback(func(wTypeRawArg uint32, wFmtRawArg uint32, hConvRawArg HCONV, hsz1RawArg HSZ, hsz2RawArg HSZ, hDataRawArg HDDEDATA, dwData1RawArg uintptr, dwData2RawArg uintptr) uintptr {
		ret := pfnCallback(wTypeRawArg, wFmtRawArg, hConvRawArg, hsz1RawArg, hsz2RawArg, hDataRawArg, dwData1RawArg, dwData2RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(ddeInitialize, 4,
		uintptr(unsafe.Pointer(pidInst)),
		pfnCallbackCallback,
		uintptr(afCmd),
		uintptr(ulRes),
		0,
		0)
	return uint32(ret1)
}

func DdeKeepStringHandle(idInst uint32, hsz HSZ) bool {
	ret1 := syscall3(ddeKeepStringHandle, 2,
		uintptr(idInst),
		uintptr(hsz),
		0)
	return ret1 != 0
}

func DdeNameService(idInst uint32, hsz1 HSZ, hsz2 HSZ, afCmd uint32) HDDEDATA {
	ret1 := syscall6(ddeNameService, 4,
		uintptr(idInst),
		uintptr(hsz1),
		uintptr(hsz2),
		uintptr(afCmd),
		0,
		0)
	return HDDEDATA(ret1)
}

func DdePostAdvise(idInst uint32, hszTopic HSZ, hszItem HSZ) bool {
	ret1 := syscall3(ddePostAdvise, 3,
		uintptr(idInst),
		uintptr(hszTopic),
		uintptr(hszItem))
	return ret1 != 0
}

func DdeQueryConvInfo(hConv HCONV, idTransaction uint32, pConvInfo *CONVINFO) uint32 {
	ret1 := syscall3(ddeQueryConvInfo, 3,
		uintptr(hConv),
		uintptr(idTransaction),
		uintptr(unsafe.Pointer(pConvInfo)))
	return uint32(ret1)
}

func DdeQueryNextServer(hConvList HCONVLIST, hConvPrev HCONV) HCONV {
	ret1 := syscall3(ddeQueryNextServer, 2,
		uintptr(hConvList),
		uintptr(hConvPrev),
		0)
	return HCONV(ret1)
}

func DdeQueryString(idInst uint32, hsz HSZ, psz LPWSTR, cchMax uint32, iCodePage int32) uint32 {
	ret1 := syscall6(ddeQueryString, 5,
		uintptr(idInst),
		uintptr(hsz),
		uintptr(unsafe.Pointer(psz)),
		uintptr(cchMax),
		uintptr(iCodePage),
		0)
	return uint32(ret1)
}

func DdeReconnect(hConv HCONV) HCONV {
	ret1 := syscall3(ddeReconnect, 1,
		uintptr(hConv),
		0,
		0)
	return HCONV(ret1)
}

func DdeSetQualityOfService(hwndClient HWND, pqosNew /*const*/ *SECURITY_QUALITY_OF_SERVICE, pqosPrev *SECURITY_QUALITY_OF_SERVICE) bool {
	ret1 := syscall3(ddeSetQualityOfService, 3,
		uintptr(hwndClient),
		uintptr(unsafe.Pointer(pqosNew)),
		uintptr(unsafe.Pointer(pqosPrev)))
	return ret1 != 0
}

func DdeSetUserHandle(hConv HCONV, id uint32, hUser *uint32) bool {
	ret1 := syscall3(ddeSetUserHandle, 3,
		uintptr(hConv),
		uintptr(id),
		uintptr(unsafe.Pointer(hUser)))
	return ret1 != 0
}

func DdeUnaccessData(hData HDDEDATA) bool {
	ret1 := syscall3(ddeUnaccessData, 1,
		uintptr(hData),
		0,
		0)
	return ret1 != 0
}

func DdeUninitialize(idInst uint32) bool {
	ret1 := syscall3(ddeUninitialize, 1,
		uintptr(idInst),
		0,
		0)
	return ret1 != 0
}

func DefDlgProc(hDlg HWND, msg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(defDlgProc, 4,
		uintptr(hDlg),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return LRESULT(ret1)
}

func DefFrameProc(hWnd HWND, hWndMDIClient HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(defFrameProc, 5,
		uintptr(hWnd),
		uintptr(hWndMDIClient),
		uintptr(uMsg),
		uintptr(wParam),
		uintptr(lParam),
		0)
	return LRESULT(ret1)
}

func DefMDIChildProc(hWnd HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(defMDIChildProc, 4,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return LRESULT(ret1)
}

func DefRawInputProc(paRawInput uintptr, nInput int32, cbSizeHeader uint32) LRESULT {
	ret1 := syscall3(defRawInputProc, 3,
		paRawInput,
		uintptr(nInput),
		uintptr(cbSizeHeader))
	return LRESULT(ret1)
}

func DefWindowProc(hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(defWindowProc, 4,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return LRESULT(ret1)
}

func DeferWindowPos(hWinPosInfo HDWP, hWnd HWND, hWndInsertAfter HWND, x int32, y int32, cx int32, cy int32, uFlags uint32) HDWP {
	ret1 := syscall9(deferWindowPos, 8,
		uintptr(hWinPosInfo),
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags),
		0)
	return HDWP(ret1)
}

func DeleteMenu(hMenu HMENU, uPosition uint32, uFlags uint32) bool {
	ret1 := syscall3(deleteMenu, 3,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags))
	return ret1 != 0
}

func DeregisterShellHookWindow(hwnd HWND) bool {
	ret1 := syscall3(deregisterShellHookWindow, 1,
		uintptr(hwnd),
		0,
		0)
	return ret1 != 0
}

func DestroyAcceleratorTable(hAccel HACCEL) bool {
	ret1 := syscall3(destroyAcceleratorTable, 1,
		uintptr(hAccel),
		0,
		0)
	return ret1 != 0
}

func DestroyCaret() bool {
	ret1 := syscall3(destroyCaret, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func DestroyCursor(hCursor HCURSOR) bool {
	ret1 := syscall3(destroyCursor, 1,
		uintptr(hCursor),
		0,
		0)
	return ret1 != 0
}

func DestroyIcon(hIcon HICON) bool {
	ret1 := syscall3(destroyIcon, 1,
		uintptr(hIcon),
		0,
		0)
	return ret1 != 0
}

func DestroyMenu(hMenu HMENU) bool {
	ret1 := syscall3(destroyMenu, 1,
		uintptr(hMenu),
		0,
		0)
	return ret1 != 0
}

func DestroyWindow(hWnd HWND) bool {
	ret1 := syscall3(destroyWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func DialogBoxIndirectParam(hInstance HINSTANCE, hDialogTemplate /*const*/ *DLGTEMPLATE, hWndParent HWND, lpDialogFunc DLGPROC, dwInitParam LPARAM) INT_PTR {
	lpDialogFuncCallback := syscall.NewCallback(func(hwndDlgRawArg HWND, uMsgRawArg uint32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := lpDialogFunc(hwndDlgRawArg, uMsgRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(dialogBoxIndirectParam, 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(hDialogTemplate)),
		uintptr(hWndParent),
		lpDialogFuncCallback,
		uintptr(dwInitParam),
		0)
	return (INT_PTR)(unsafe.Pointer(ret1))
}

func DialogBoxParam(hInstance HINSTANCE, lpTemplateName string, hWndParent HWND, lpDialogFunc DLGPROC, dwInitParam LPARAM) INT_PTR {
	lpTemplateNameStr := unicode16FromString(lpTemplateName)
	lpDialogFuncCallback := syscall.NewCallback(func(hwndDlgRawArg HWND, uMsgRawArg uint32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := lpDialogFunc(hwndDlgRawArg, uMsgRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(dialogBoxParam, 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpTemplateNameStr[0])),
		uintptr(hWndParent),
		lpDialogFuncCallback,
		uintptr(dwInitParam),
		0)
	return (INT_PTR)(unsafe.Pointer(ret1))
}

func DisableProcessWindowsGhosting() {
	syscall3(disableProcessWindowsGhosting, 0,
		0,
		0,
		0)
}

func DispatchMessage(lpMsg /*const*/ *MSG) LRESULT {
	ret1 := syscall3(dispatchMessage, 1,
		uintptr(unsafe.Pointer(lpMsg)),
		0,
		0)
	return LRESULT(ret1)
}

func DlgDirListComboBox(hDlg HWND, lpPathSpec LPWSTR, nIDComboBox int32, nIDStaticPath int32, uFiletype uint32) int32 {
	ret1 := syscall6(dlgDirListComboBox, 5,
		uintptr(hDlg),
		uintptr(unsafe.Pointer(lpPathSpec)),
		uintptr(nIDComboBox),
		uintptr(nIDStaticPath),
		uintptr(uFiletype),
		0)
	return int32(ret1)
}

func DlgDirList(hDlg HWND, lpPathSpec LPWSTR, nIDListBox int32, nIDStaticPath int32, uFileType uint32) int32 {
	ret1 := syscall6(dlgDirList, 5,
		uintptr(hDlg),
		uintptr(unsafe.Pointer(lpPathSpec)),
		uintptr(nIDListBox),
		uintptr(nIDStaticPath),
		uintptr(uFileType),
		0)
	return int32(ret1)
}

func DlgDirSelectComboBoxEx(hwndDlg HWND, lpString LPWSTR, cchOut int32, idComboBox int32) bool {
	ret1 := syscall6(dlgDirSelectComboBoxEx, 4,
		uintptr(hwndDlg),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(cchOut),
		uintptr(idComboBox),
		0,
		0)
	return ret1 != 0
}

func DlgDirSelectEx(hwndDlg HWND, lpString LPWSTR, chCount int32, idListBox int32) bool {
	ret1 := syscall6(dlgDirSelectEx, 4,
		uintptr(hwndDlg),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(chCount),
		uintptr(idListBox),
		0,
		0)
	return ret1 != 0
}

func DragDetect(hwnd HWND, pt POINT) bool {
	ret1 := syscall3(dragDetect, 3,
		uintptr(hwnd),
		uintptr(pt.X),
		uintptr(pt.Y))
	return ret1 != 0
}

func DragObject(hwndParent HWND, hwndFrom HWND, fmt uint32, data *uint32, hcur HCURSOR) uint32 {
	ret1 := syscall6(dragObject, 5,
		uintptr(hwndParent),
		uintptr(hwndFrom),
		uintptr(fmt),
		uintptr(unsafe.Pointer(data)),
		uintptr(hcur),
		0)
	return uint32(ret1)
}

func DrawAnimatedRects(hwnd HWND, idAni int32, lprcFrom /*const*/ *RECT, lprcTo /*const*/ *RECT) bool {
	ret1 := syscall6(drawAnimatedRects, 4,
		uintptr(hwnd),
		uintptr(idAni),
		uintptr(unsafe.Pointer(lprcFrom)),
		uintptr(unsafe.Pointer(lprcTo)),
		0,
		0)
	return ret1 != 0
}

func DrawCaption(hwnd HWND, hdc HDC, lprect /*const*/ *RECT, flags uint32) bool {
	ret1 := syscall6(drawCaption, 4,
		uintptr(hwnd),
		uintptr(hdc),
		uintptr(unsafe.Pointer(lprect)),
		uintptr(flags),
		0,
		0)
	return ret1 != 0
}

func DrawEdge(hdc HDC, qrc *RECT, edge uint32, grfFlags uint32) bool {
	ret1 := syscall6(drawEdge, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(qrc)),
		uintptr(edge),
		uintptr(grfFlags),
		0,
		0)
	return ret1 != 0
}

func DrawFocusRect(hDC HDC, lprc /*const*/ *RECT) bool {
	ret1 := syscall3(drawFocusRect, 2,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		0)
	return ret1 != 0
}

func DrawFrameControl(unnamed0 HDC, unnamed1 *RECT, unnamed2 uint32, unnamed3 uint32) bool {
	ret1 := syscall6(drawFrameControl, 4,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(unnamed2),
		uintptr(unnamed3),
		0,
		0)
	return ret1 != 0
}

func DrawIcon(hDC HDC, x int32, y int32, hIcon HICON) bool {
	ret1 := syscall6(drawIcon, 4,
		uintptr(hDC),
		uintptr(x),
		uintptr(y),
		uintptr(hIcon),
		0,
		0)
	return ret1 != 0
}

func DrawIconEx(hdc HDC, xLeft int32, yTop int32, hIcon HICON, cxWidth int32, cyWidth int32, istepIfAniCur uint32, hbrFlickerFreeDraw HBRUSH, diFlags uint32) bool {
	ret1 := syscall9(drawIconEx, 9,
		uintptr(hdc),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(hIcon),
		uintptr(cxWidth),
		uintptr(cyWidth),
		uintptr(istepIfAniCur),
		uintptr(hbrFlickerFreeDraw),
		uintptr(diFlags))
	return ret1 != 0
}

func DrawMenuBar(hWnd HWND) bool {
	ret1 := syscall3(drawMenuBar, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func DrawState(hdc HDC, hbrFore HBRUSH, qfnCallBack DRAWSTATEPROC, lData LPARAM, wData WPARAM, x int32, y int32, cx int32, cy int32, uFlags uint32) bool {
	qfnCallBackCallback := syscall.NewCallback(func(hdcRawArg HDC, lDataRawArg uintptr, wDataRawArg uintptr, cxRawArg int32, cyRawArg int32) uintptr {
		ret := qfnCallBack(hdcRawArg, lDataRawArg, wDataRawArg, cxRawArg, cyRawArg)
		return uintptr(ret)
	})
	ret1 := syscall12(drawState, 10,
		uintptr(hdc),
		uintptr(hbrFore),
		qfnCallBackCallback,
		uintptr(lData),
		uintptr(wData),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags),
		0,
		0)
	return ret1 != 0
}

func DrawTextEx(hdc HDC, lpchText LPWSTR, cchText int32, lprc *RECT, format uint32, lpdtp *DRAWTEXTPARAMS) int32 {
	ret1 := syscall6(drawTextEx, 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpchText)),
		uintptr(cchText),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(format),
		uintptr(unsafe.Pointer(lpdtp)))
	return int32(ret1)
}

func DrawText(hdc HDC, lpchText string, cchText int32, lprc *RECT, format uint32) int32 {
	lpchTextStr := unicode16FromString(lpchText)
	ret1 := syscall6(drawText, 5,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpchTextStr[0])),
		uintptr(cchText),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(format),
		0)
	return int32(ret1)
}

func EmptyClipboard() bool {
	ret1 := syscall3(emptyClipboard, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func EnableMenuItem(hMenu HMENU, uIDEnableItem uint32, uEnable uint32) bool {
	ret1 := syscall3(enableMenuItem, 3,
		uintptr(hMenu),
		uintptr(uIDEnableItem),
		uintptr(uEnable))
	return ret1 != 0
}

func EnableScrollBar(hWnd HWND, wSBflags uint32, wArrows uint32) bool {
	ret1 := syscall3(enableScrollBar, 3,
		uintptr(hWnd),
		uintptr(wSBflags),
		uintptr(wArrows))
	return ret1 != 0
}

func EnableWindow(hWnd HWND, bEnable bool) bool {
	ret1 := syscall3(enableWindow, 2,
		uintptr(hWnd),
		getUintptrFromBool(bEnable),
		0)
	return ret1 != 0
}

func EndDeferWindowPos(hWinPosInfo HDWP) bool {
	ret1 := syscall3(endDeferWindowPos, 1,
		uintptr(hWinPosInfo),
		0,
		0)
	return ret1 != 0
}

func EndDialog(hDlg HWND, nResult INT_PTR) bool {
	ret1 := syscall3(endDialog, 2,
		uintptr(hDlg),
		uintptr(unsafe.Pointer(nResult)),
		0)
	return ret1 != 0
}

func EndMenu() bool {
	ret1 := syscall3(endMenu, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func EndPaint(hWnd HWND, lpPaint /*const*/ *PAINTSTRUCT) bool {
	ret1 := syscall3(endPaint, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPaint)),
		0)
	return ret1 != 0
}

func EndTask(hWnd HWND, fShutDown bool, fForce bool) bool {
	ret1 := syscall3(endTask, 3,
		uintptr(hWnd),
		getUintptrFromBool(fShutDown),
		getUintptrFromBool(fForce))
	return ret1 != 0
}

func EnumChildWindows(hWndParent HWND, lpEnumFunc WNDENUMPROC, lParam LPARAM) bool {
	lpEnumFuncCallback := syscall.NewCallback(func(hWndRawArg HWND, lParamRawArg LPARAM) uintptr {
		ret := lpEnumFunc(hWndRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumChildWindows, 3,
		uintptr(hWndParent),
		lpEnumFuncCallback,
		uintptr(lParam))
	return ret1 != 0
}

func EnumClipboardFormats(format uint32) uint32 {
	ret1 := syscall3(enumClipboardFormats, 1,
		uintptr(format),
		0,
		0)
	return uint32(ret1)
}

func EnumDesktopWindows(hDesktop HDESK, lpfn WNDENUMPROC, lParam LPARAM) bool {
	lpfnCallback := syscall.NewCallback(func(hWndRawArg HWND, lParamRawArg LPARAM) uintptr {
		ret := lpfn(hWndRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumDesktopWindows, 3,
		uintptr(hDesktop),
		lpfnCallback,
		uintptr(lParam))
	return ret1 != 0
}

func EnumDesktops(hwinsta HWINSTA, lpEnumFunc DESKTOPENUMPROC, lParam LPARAM) bool {
	lpEnumFuncCallback := syscall.NewCallback(func(lpszDesktopRawArg LPWSTR, lParamRawArg LPARAM) uintptr {
		ret := lpEnumFunc(lpszDesktopRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumDesktops, 3,
		uintptr(hwinsta),
		lpEnumFuncCallback,
		uintptr(lParam))
	return ret1 != 0
}

func EnumDisplayDevices(lpDevice string, iDevNum uint32, lpDisplayDevice *DISPLAY_DEVICE, dwFlags uint32) bool {
	lpDeviceStr := unicode16FromString(lpDevice)
	ret1 := syscall6(enumDisplayDevices, 4,
		uintptr(unsafe.Pointer(&lpDeviceStr[0])),
		uintptr(iDevNum),
		uintptr(unsafe.Pointer(lpDisplayDevice)),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func EnumDisplayMonitors(hdc HDC, lprcClip /*const*/ *RECT, lpfnEnum MONITORENUMPROC, dwData LPARAM) bool {
	lpfnEnumCallback := syscall.NewCallback(func(hMonitorRawArg HMONITOR, hdcMonitorRawArg HDC, lprcMonitorRawArg *RECT, dwDataRawArg uintptr) uintptr {
		ret := lpfnEnum(hMonitorRawArg, hdcMonitorRawArg, lprcMonitorRawArg, dwDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumDisplayMonitors, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lprcClip)),
		lpfnEnumCallback,
		uintptr(dwData),
		0,
		0)
	return ret1 != 0
}

func EnumDisplaySettingsEx(lpszDeviceName string, iModeNum uint32, lpDevMode LPDEVMODE, dwFlags uint32) bool {
	lpszDeviceNameStr := unicode16FromString(lpszDeviceName)
	ret1 := syscall6(enumDisplaySettingsEx, 4,
		uintptr(unsafe.Pointer(&lpszDeviceNameStr[0])),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(lpDevMode)),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func EnumDisplaySettings(lpszDeviceName string, iModeNum uint32, lpDevMode LPDEVMODE) bool {
	lpszDeviceNameStr := unicode16FromString(lpszDeviceName)
	ret1 := syscall3(enumDisplaySettings, 3,
		uintptr(unsafe.Pointer(&lpszDeviceNameStr[0])),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(lpDevMode)))
	return ret1 != 0
}

func EnumPropsEx(hWnd HWND, lpEnumFunc PROPENUMPROCEX, lParam LPARAM) int32 {
	lpEnumFuncCallback := syscall.NewCallback(func(hwndRawArg HWND, lpszStringRawArg LPWSTR, hDataRawArg HANDLE, dwDataRawArg uintptr) uintptr {
		ret := lpEnumFunc(hwndRawArg, lpszStringRawArg, hDataRawArg, dwDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumPropsEx, 3,
		uintptr(hWnd),
		lpEnumFuncCallback,
		uintptr(lParam))
	return int32(ret1)
}

func EnumProps(hWnd HWND, lpEnumFunc PROPENUMPROC) int32 {
	lpEnumFuncCallback := syscall.NewCallback(func(hWndRawArg HWND, lpszStringRawArg /*const*/ *uint16, hDataRawArg HANDLE) uintptr {
		lpszString := stringFromUnicode16(lpszStringRawArg)
		ret := lpEnumFunc(hWndRawArg, lpszString, hDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumProps, 2,
		uintptr(hWnd),
		lpEnumFuncCallback,
		0)
	return int32(ret1)
}

func EnumThreadWindows(dwThreadId uint32, lpfn WNDENUMPROC, lParam LPARAM) bool {
	lpfnCallback := syscall.NewCallback(func(hWndRawArg HWND, lParamRawArg LPARAM) uintptr {
		ret := lpfn(hWndRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumThreadWindows, 3,
		uintptr(dwThreadId),
		lpfnCallback,
		uintptr(lParam))
	return ret1 != 0
}

func EnumWindowStations(lpEnumFunc WINSTAENUMPROC, lParam LPARAM) bool {
	lpEnumFuncCallback := syscall.NewCallback(func(lpszWindowStationRawArg LPWSTR, lParamRawArg LPARAM) uintptr {
		ret := lpEnumFunc(lpszWindowStationRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumWindowStations, 2,
		lpEnumFuncCallback,
		uintptr(lParam),
		0)
	return ret1 != 0
}

func EnumWindows(lpEnumFunc WNDENUMPROC, lParam LPARAM) bool {
	lpEnumFuncCallback := syscall.NewCallback(func(hWndRawArg HWND, lParamRawArg LPARAM) uintptr {
		ret := lpEnumFunc(hWndRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumWindows, 2,
		lpEnumFuncCallback,
		uintptr(lParam),
		0)
	return ret1 != 0
}

func EqualRect(lprc1 /*const*/ *RECT, lprc2 /*const*/ *RECT) bool {
	ret1 := syscall3(equalRect, 2,
		uintptr(unsafe.Pointer(lprc1)),
		uintptr(unsafe.Pointer(lprc2)),
		0)
	return ret1 != 0
}

func ExcludeUpdateRgn(hDC HDC, hWnd HWND) int32 {
	ret1 := syscall3(excludeUpdateRgn, 2,
		uintptr(hDC),
		uintptr(hWnd),
		0)
	return int32(ret1)
}

func ExitWindowsEx(uFlags uint32, dwReason uint32) bool {
	ret1 := syscall3(exitWindowsEx, 2,
		uintptr(uFlags),
		uintptr(dwReason),
		0)
	return ret1 != 0
}

func FillRect(hDC HDC, lprc /*const*/ *RECT, hbr HBRUSH) int32 {
	ret1 := syscall3(fillRect, 3,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))
	return int32(ret1)
}

func FindWindowEx(hWndParent HWND, hWndChildAfter HWND, lpszClass string, lpszWindow string) HWND {
	lpszClassStr := unicode16FromString(lpszClass)
	lpszWindowStr := unicode16FromString(lpszWindow)
	ret1 := syscall6(findWindowEx, 4,
		uintptr(hWndParent),
		uintptr(hWndChildAfter),
		uintptr(unsafe.Pointer(&lpszClassStr[0])),
		uintptr(unsafe.Pointer(&lpszWindowStr[0])),
		0,
		0)
	return HWND(ret1)
}

func FindWindow(lpClassName string, lpWindowName string) HWND {
	lpClassNameStr := unicode16FromString(lpClassName)
	lpWindowNameStr := unicode16FromString(lpWindowName)
	ret1 := syscall3(findWindow, 2,
		uintptr(unsafe.Pointer(&lpClassNameStr[0])),
		uintptr(unsafe.Pointer(&lpWindowNameStr[0])),
		0)
	return HWND(ret1)
}

func FlashWindow(hWnd HWND, bInvert bool) bool {
	ret1 := syscall3(flashWindow, 2,
		uintptr(hWnd),
		getUintptrFromBool(bInvert),
		0)
	return ret1 != 0
}

func FlashWindowEx(pfwi *FLASHWINFO) bool {
	ret1 := syscall3(flashWindowEx, 1,
		uintptr(unsafe.Pointer(pfwi)),
		0,
		0)
	return ret1 != 0
}

func FrameRect(hDC HDC, lprc /*const*/ *RECT, hbr HBRUSH) int32 {
	ret1 := syscall3(frameRect, 3,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))
	return int32(ret1)
}

func FreeDDElParam(msg uint32, lParam LPARAM) bool {
	ret1 := syscall3(freeDDElParam, 2,
		uintptr(msg),
		uintptr(lParam),
		0)
	return ret1 != 0
}

func GetActiveWindow() HWND {
	ret1 := syscall3(getActiveWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetAltTabInfo(hwnd HWND, iItem int32, pati *ALTTABINFO, pszItemText LPWSTR, cchItemText uint32) bool {
	ret1 := syscall6(getAltTabInfo, 5,
		uintptr(hwnd),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pati)),
		uintptr(unsafe.Pointer(pszItemText)),
		uintptr(cchItemText),
		0)
	return ret1 != 0
}

func GetAncestor(hwnd HWND, gaFlags uint32) HWND {
	ret1 := syscall3(getAncestor, 2,
		uintptr(hwnd),
		uintptr(gaFlags),
		0)
	return HWND(ret1)
}

func GetAsyncKeyState(vKey int32) int16 {
	ret1 := syscall3(getAsyncKeyState, 1,
		uintptr(vKey),
		0,
		0)
	return int16(ret1)
}

func GetCapture() HWND {
	ret1 := syscall3(getCapture, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetCaretBlinkTime() uint32 {
	ret1 := syscall3(getCaretBlinkTime, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetCaretPos(lpPoint *POINT) bool {
	ret1 := syscall3(getCaretPos, 1,
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)
	return ret1 != 0
}

func GetClassInfoEx(hInstance HINSTANCE, lpszClass string, lpwcx *WNDCLASSEX) bool {
	lpszClassStr := unicode16FromString(lpszClass)
	ret1 := syscall3(getClassInfoEx, 3,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpszClassStr[0])),
		uintptr(unsafe.Pointer(lpwcx)))
	return ret1 != 0
}

func GetClassInfo(hInstance HINSTANCE, lpClassName string, lpWndClass *WNDCLASS) bool {
	lpClassNameStr := unicode16FromString(lpClassName)
	ret1 := syscall3(getClassInfo, 3,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpClassNameStr[0])),
		uintptr(unsafe.Pointer(lpWndClass)))
	return ret1 != 0
}

func GetClassLongPtr(hWnd HWND, nIndex int32) *uint32 {
	ret1 := syscall3(getClassLongPtr, 2,
		uintptr(hWnd),
		uintptr(nIndex),
		0)
	return (*uint32)(unsafe.Pointer(ret1))
}

func GetClassLong(hWnd HWND, nIndex int32) uint32 {
	ret1 := syscall3(getClassLong, 2,
		uintptr(hWnd),
		uintptr(nIndex),
		0)
	return uint32(ret1)
}

func GetClassName(hWnd HWND, lpClassName LPWSTR, nMaxCount int32) int32 {
	ret1 := syscall3(getClassName, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(nMaxCount))
	return int32(ret1)
}

func GetClassWord(hWnd HWND, nIndex int32) uint16 {
	ret1 := syscall3(getClassWord, 2,
		uintptr(hWnd),
		uintptr(nIndex),
		0)
	return uint16(ret1)
}

func GetClientRect(hWnd HWND, lpRect *RECT) bool {
	ret1 := syscall3(getClientRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		0)
	return ret1 != 0
}

func GetClipCursor(lpRect *RECT) bool {
	ret1 := syscall3(getClipCursor, 1,
		uintptr(unsafe.Pointer(lpRect)),
		0,
		0)
	return ret1 != 0
}

func GetClipboardData(uFormat uint32) HANDLE {
	ret1 := syscall3(getClipboardData, 1,
		uintptr(uFormat),
		0,
		0)
	return HANDLE(ret1)
}

func GetClipboardFormatName(format uint32, lpszFormatName LPWSTR, cchMaxCount int32) int32 {
	ret1 := syscall3(getClipboardFormatName, 3,
		uintptr(format),
		uintptr(unsafe.Pointer(lpszFormatName)),
		uintptr(cchMaxCount))
	return int32(ret1)
}

func GetClipboardOwner() HWND {
	ret1 := syscall3(getClipboardOwner, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetClipboardSequenceNumber() uint32 {
	ret1 := syscall3(getClipboardSequenceNumber, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetClipboardViewer() HWND {
	ret1 := syscall3(getClipboardViewer, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetComboBoxInfo(hwndCombo HWND, pcbi *COMBOBOXINFO) bool {
	ret1 := syscall3(getComboBoxInfo, 2,
		uintptr(hwndCombo),
		uintptr(unsafe.Pointer(pcbi)),
		0)
	return ret1 != 0
}

func GetCursor() HCURSOR {
	ret1 := syscall3(getCursor, 0,
		0,
		0,
		0)
	return HCURSOR(ret1)
}

func GetCursorInfo(pci *CURSORINFO) bool {
	ret1 := syscall3(getCursorInfo, 1,
		uintptr(unsafe.Pointer(pci)),
		0,
		0)
	return ret1 != 0
}

func GetCursorPos(lpPoint *POINT) bool {
	ret1 := syscall3(getCursorPos, 1,
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)
	return ret1 != 0
}

func GetDC(hWnd HWND) HDC {
	ret1 := syscall3(getDC, 1,
		uintptr(hWnd),
		0,
		0)
	return HDC(ret1)
}

func GetDCEx(hWnd HWND, hrgnClip HRGN, flags uint32) HDC {
	ret1 := syscall3(getDCEx, 3,
		uintptr(hWnd),
		uintptr(hrgnClip),
		uintptr(flags))
	return HDC(ret1)
}

func GetDesktopWindow() HWND {
	ret1 := syscall3(getDesktopWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetDialogBaseUnits() int32 {
	ret1 := syscall3(getDialogBaseUnits, 0,
		0,
		0,
		0)
	return int32(ret1)
}

func GetDlgCtrlID(hWnd HWND) int32 {
	ret1 := syscall3(getDlgCtrlID, 1,
		uintptr(hWnd),
		0,
		0)
	return int32(ret1)
}

func GetDlgItem(hDlg HWND, nIDDlgItem int32) HWND {
	ret1 := syscall3(getDlgItem, 2,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		0)
	return HWND(ret1)
}

func GetDlgItemInt(hDlg HWND, nIDDlgItem int32, lpTranslated *BOOL, bSigned bool) uint32 {
	ret1 := syscall6(getDlgItemInt, 4,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		uintptr(unsafe.Pointer(lpTranslated)),
		getUintptrFromBool(bSigned),
		0,
		0)
	return uint32(ret1)
}

func GetDlgItemText(hDlg HWND, nIDDlgItem int32, lpString LPWSTR, cchMax int32) uint32 {
	ret1 := syscall6(getDlgItemText, 4,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(cchMax),
		0,
		0)
	return uint32(ret1)
}

func GetDoubleClickTime() uint32 {
	ret1 := syscall3(getDoubleClickTime, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetFocus() HWND {
	ret1 := syscall3(getFocus, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetForegroundWindow() HWND {
	ret1 := syscall3(getForegroundWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetGUIThreadInfo(idThread uint32, pgui *GUITHREADINFO) bool {
	ret1 := syscall3(getGUIThreadInfo, 2,
		uintptr(idThread),
		uintptr(unsafe.Pointer(pgui)),
		0)
	return ret1 != 0
}

func GetGestureConfig(hwnd HWND, dwReserved uint32, dwFlags uint32, pcIDs *uint32, pGestureConfig *GESTURECONFIG, cbSize uint32) bool {
	ret1 := syscall6(getGestureConfig, 6,
		uintptr(hwnd),
		uintptr(dwReserved),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pcIDs)),
		uintptr(unsafe.Pointer(pGestureConfig)),
		uintptr(cbSize))
	return ret1 != 0
}

func GetGestureExtraArgs(hGestureInfo HGESTUREINFO, cbExtraArgs uint32, pExtraArgs *byte) bool {
	ret1 := syscall3(getGestureExtraArgs, 3,
		uintptr(hGestureInfo),
		uintptr(cbExtraArgs),
		uintptr(unsafe.Pointer(pExtraArgs)))
	return ret1 != 0
}

func GetGestureInfo(hGestureInfo HGESTUREINFO, pGestureInfo *GESTUREINFO) bool {
	ret1 := syscall3(getGestureInfo, 2,
		uintptr(hGestureInfo),
		uintptr(unsafe.Pointer(pGestureInfo)),
		0)
	return ret1 != 0
}

func GetGuiResources(hProcess HANDLE, uiFlags uint32) uint32 {
	ret1 := syscall3(getGuiResources, 2,
		uintptr(hProcess),
		uintptr(uiFlags),
		0)
	return uint32(ret1)
}

func GetIconInfo(hIcon HICON, piconinfo *ICONINFO) bool {
	ret1 := syscall3(getIconInfo, 2,
		uintptr(hIcon),
		uintptr(unsafe.Pointer(piconinfo)),
		0)
	return ret1 != 0
}

func GetInputState() bool {
	ret1 := syscall3(getInputState, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func GetKBCodePage() uint32 {
	ret1 := syscall3(getKBCodePage, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetKeyNameText(lParam LONG, lpString LPWSTR, cchSize int32) int32 {
	ret1 := syscall3(getKeyNameText, 3,
		uintptr(lParam),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(cchSize))
	return int32(ret1)
}

func GetKeyState(nVirtKey int32) int16 {
	ret1 := syscall3(getKeyState, 1,
		uintptr(nVirtKey),
		0,
		0)
	return int16(ret1)
}

func GetKeyboardLayout(idThread uint32) HKL {
	ret1 := syscall3(getKeyboardLayout, 1,
		uintptr(idThread),
		0,
		0)
	return HKL(ret1)
}

func GetKeyboardLayoutList(nBuff int32, lpList *HKL) int32 {
	ret1 := syscall3(getKeyboardLayoutList, 2,
		uintptr(nBuff),
		uintptr(unsafe.Pointer(lpList)),
		0)
	return int32(ret1)
}

func GetKeyboardLayoutName(pwszKLID LPWSTR) bool {
	ret1 := syscall3(getKeyboardLayoutName, 1,
		uintptr(unsafe.Pointer(pwszKLID)),
		0,
		0)
	return ret1 != 0
}

func GetKeyboardState(lpKeyState *byte) bool {
	ret1 := syscall3(getKeyboardState, 1,
		uintptr(unsafe.Pointer(lpKeyState)),
		0,
		0)
	return ret1 != 0
}

func GetKeyboardType(nTypeFlag int32) int32 {
	ret1 := syscall3(getKeyboardType, 1,
		uintptr(nTypeFlag),
		0,
		0)
	return int32(ret1)
}

func GetLastActivePopup(hWnd HWND) HWND {
	ret1 := syscall3(getLastActivePopup, 1,
		uintptr(hWnd),
		0,
		0)
	return HWND(ret1)
}

func GetLastInputInfo(plii *LASTINPUTINFO) bool {
	ret1 := syscall3(getLastInputInfo, 1,
		uintptr(unsafe.Pointer(plii)),
		0,
		0)
	return ret1 != 0
}

func GetLayeredWindowAttributes(hwnd HWND, pcrKey *COLORREF, pbAlpha *byte, pdwFlags *uint32) bool {
	ret1 := syscall6(getLayeredWindowAttributes, 4,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pcrKey)),
		uintptr(unsafe.Pointer(pbAlpha)),
		uintptr(unsafe.Pointer(pdwFlags)),
		0,
		0)
	return ret1 != 0
}

func GetListBoxInfo(hwnd HWND) uint32 {
	ret1 := syscall3(getListBoxInfo, 1,
		uintptr(hwnd),
		0,
		0)
	return uint32(ret1)
}

func GetMenu(hWnd HWND) HMENU {
	ret1 := syscall3(getMenu, 1,
		uintptr(hWnd),
		0,
		0)
	return HMENU(ret1)
}

func GetMenuBarInfo(hwnd HWND, idObject LONG, idItem LONG, pmbi *MENUBARINFO) bool {
	ret1 := syscall6(getMenuBarInfo, 4,
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idItem),
		uintptr(unsafe.Pointer(pmbi)),
		0,
		0)
	return ret1 != 0
}

func GetMenuCheckMarkDimensions() LONG {
	ret1 := syscall3(getMenuCheckMarkDimensions, 0,
		0,
		0,
		0)
	return LONG(ret1)
}

func GetMenuContextHelpId(unnamed0 HMENU) uint32 {
	ret1 := syscall3(getMenuContextHelpId, 1,
		uintptr(unnamed0),
		0,
		0)
	return uint32(ret1)
}

func GetMenuDefaultItem(hMenu HMENU, fByPos uint32, gmdiFlags uint32) uint32 {
	ret1 := syscall3(getMenuDefaultItem, 3,
		uintptr(hMenu),
		uintptr(fByPos),
		uintptr(gmdiFlags))
	return uint32(ret1)
}

func GetMenuInfo(unnamed0 HMENU, unnamed1 *MENUINFO) bool {
	ret1 := syscall3(getMenuInfo, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func GetMenuItemCount(hMenu HMENU) int32 {
	ret1 := syscall3(getMenuItemCount, 1,
		uintptr(hMenu),
		0,
		0)
	return int32(ret1)
}

func GetMenuItemID(hMenu HMENU, nPos int32) uint32 {
	ret1 := syscall3(getMenuItemID, 2,
		uintptr(hMenu),
		uintptr(nPos),
		0)
	return uint32(ret1)
}

func GetMenuItemInfo(hmenu HMENU, item uint32, fByPosition bool, lpmii *MENUITEMINFO) bool {
	ret1 := syscall6(getMenuItemInfo, 4,
		uintptr(hmenu),
		uintptr(item),
		getUintptrFromBool(fByPosition),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)
	return ret1 != 0
}

func GetMenuItemRect(hWnd HWND, hMenu HMENU, uItem uint32, lprcItem *RECT) bool {
	ret1 := syscall6(getMenuItemRect, 4,
		uintptr(hWnd),
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(unsafe.Pointer(lprcItem)),
		0,
		0)
	return ret1 != 0
}

func GetMenuState(hMenu HMENU, uId uint32, uFlags uint32) uint32 {
	ret1 := syscall3(getMenuState, 3,
		uintptr(hMenu),
		uintptr(uId),
		uintptr(uFlags))
	return uint32(ret1)
}

func GetMenuString(hMenu HMENU, uIDItem uint32, lpString LPWSTR, cchMax int32, flags uint32) int32 {
	ret1 := syscall6(getMenuString, 5,
		uintptr(hMenu),
		uintptr(uIDItem),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(cchMax),
		uintptr(flags),
		0)
	return int32(ret1)
}

func GetMessageExtraInfo() LPARAM {
	ret1 := syscall3(getMessageExtraInfo, 0,
		0,
		0,
		0)
	return LPARAM(ret1)
}

func GetMessagePos() uint32 {
	ret1 := syscall3(getMessagePos, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetMessageTime() LONG {
	ret1 := syscall3(getMessageTime, 0,
		0,
		0,
		0)
	return LONG(ret1)
}

func GetMessage(lpMsg *MSG, hWnd HWND, wMsgFilterMin uint32, wMsgFilterMax uint32) bool {
	ret1 := syscall6(getMessage, 4,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		0,
		0)
	return ret1 != 0
}

func GetMonitorInfo(hMonitor HMONITOR, lpmi *MONITORINFO) bool {
	ret1 := syscall3(getMonitorInfo, 2,
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lpmi)),
		0)
	return ret1 != 0
}

func GetMouseMovePointsEx(cbSize uint32, lppt *MOUSEMOVEPOINT, lpptBuf *MOUSEMOVEPOINT, nBufPoints int32, resolution uint32) int32 {
	ret1 := syscall6(getMouseMovePointsEx, 5,
		uintptr(cbSize),
		uintptr(unsafe.Pointer(lppt)),
		uintptr(unsafe.Pointer(lpptBuf)),
		uintptr(nBufPoints),
		uintptr(resolution),
		0)
	return int32(ret1)
}

func GetNextDlgGroupItem(hDlg HWND, hCtl HWND, bPrevious bool) HWND {
	ret1 := syscall3(getNextDlgGroupItem, 3,
		uintptr(hDlg),
		uintptr(hCtl),
		getUintptrFromBool(bPrevious))
	return HWND(ret1)
}

func GetNextDlgTabItem(hDlg HWND, hCtl HWND, bPrevious bool) HWND {
	ret1 := syscall3(getNextDlgTabItem, 3,
		uintptr(hDlg),
		uintptr(hCtl),
		getUintptrFromBool(bPrevious))
	return HWND(ret1)
}

func GetOpenClipboardWindow() HWND {
	ret1 := syscall3(getOpenClipboardWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetParent(hWnd HWND) HWND {
	ret1 := syscall3(getParent, 1,
		uintptr(hWnd),
		0,
		0)
	return HWND(ret1)
}

func GetPriorityClipboardFormat(paFormatPriorityList *UINT, cFormats int32) int32 {
	ret1 := syscall3(getPriorityClipboardFormat, 2,
		uintptr(unsafe.Pointer(paFormatPriorityList)),
		uintptr(cFormats),
		0)
	return int32(ret1)
}

func GetProcessDefaultLayout(pdwDefaultLayout *uint32) bool {
	ret1 := syscall3(getProcessDefaultLayout, 1,
		uintptr(unsafe.Pointer(pdwDefaultLayout)),
		0,
		0)
	return ret1 != 0
}

func GetProcessWindowStation() HWINSTA {
	ret1 := syscall3(getProcessWindowStation, 0,
		0,
		0,
		0)
	return HWINSTA(ret1)
}

func GetProp(hWnd HWND, lpString string) HANDLE {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(getProp, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0)
	return HANDLE(ret1)
}

func GetQueueStatus(flags uint32) uint32 {
	ret1 := syscall3(getQueueStatus, 1,
		uintptr(flags),
		0,
		0)
	return uint32(ret1)
}

func GetRawInputBuffer(pData *RAWINPUT, pcbSize *uint32, cbSizeHeader uint32) uint32 {
	ret1 := syscall3(getRawInputBuffer, 3,
		uintptr(unsafe.Pointer(pData)),
		uintptr(unsafe.Pointer(pcbSize)),
		uintptr(cbSizeHeader))
	return uint32(ret1)
}

func GetRawInputData(hRawInput HRAWINPUT, uiCommand uint32, pData LPVOID, pcbSize *uint32, cbSizeHeader uint32) uint32 {
	ret1 := syscall6(getRawInputData, 5,
		uintptr(hRawInput),
		uintptr(uiCommand),
		uintptr(unsafe.Pointer(pData)),
		uintptr(unsafe.Pointer(pcbSize)),
		uintptr(cbSizeHeader),
		0)
	return uint32(ret1)
}

func GetRawInputDeviceInfo(hDevice HANDLE, uiCommand uint32, pData LPVOID, pcbSize *uint32) uint32 {
	ret1 := syscall6(getRawInputDeviceInfo, 4,
		uintptr(hDevice),
		uintptr(uiCommand),
		uintptr(unsafe.Pointer(pData)),
		uintptr(unsafe.Pointer(pcbSize)),
		0,
		0)
	return uint32(ret1)
}

func GetRawInputDeviceList(pRawInputDeviceList *RAWINPUTDEVICELIST, puiNumDevices *uint32, cbSize uint32) uint32 {
	ret1 := syscall3(getRawInputDeviceList, 3,
		uintptr(unsafe.Pointer(pRawInputDeviceList)),
		uintptr(unsafe.Pointer(puiNumDevices)),
		uintptr(cbSize))
	return uint32(ret1)
}

func GetRegisteredRawInputDevices(pRawInputDevices *RAWINPUTDEVICE, puiNumDevices *uint32, cbSize uint32) uint32 {
	ret1 := syscall3(getRegisteredRawInputDevices, 3,
		uintptr(unsafe.Pointer(pRawInputDevices)),
		uintptr(unsafe.Pointer(puiNumDevices)),
		uintptr(cbSize))
	return uint32(ret1)
}

func GetScrollBarInfo(hwnd HWND, idObject LONG, psbi *SCROLLBARINFO) bool {
	ret1 := syscall3(getScrollBarInfo, 3,
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(unsafe.Pointer(psbi)))
	return ret1 != 0
}

func GetScrollInfo(hwnd HWND, nBar int32, lpsi *SCROLLINFO) bool {
	ret1 := syscall3(getScrollInfo, 3,
		uintptr(hwnd),
		uintptr(nBar),
		uintptr(unsafe.Pointer(lpsi)))
	return ret1 != 0
}

func GetScrollPos(hWnd HWND, nBar int32) int32 {
	ret1 := syscall3(getScrollPos, 2,
		uintptr(hWnd),
		uintptr(nBar),
		0)
	return int32(ret1)
}

func GetScrollRange(hWnd HWND, nBar int32, lpMinPos *int32, lpMaxPos *int32) bool {
	ret1 := syscall6(getScrollRange, 4,
		uintptr(hWnd),
		uintptr(nBar),
		uintptr(unsafe.Pointer(lpMinPos)),
		uintptr(unsafe.Pointer(lpMaxPos)),
		0,
		0)
	return ret1 != 0
}

func GetShellWindow() HWND {
	ret1 := syscall3(getShellWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetSubMenu(hMenu HMENU, nPos int32) HMENU {
	ret1 := syscall3(getSubMenu, 2,
		uintptr(hMenu),
		uintptr(nPos),
		0)
	return HMENU(ret1)
}

func GetSysColor(nIndex int32) uint32 {
	ret1 := syscall3(getSysColor, 1,
		uintptr(nIndex),
		0,
		0)
	return uint32(ret1)
}

func GetSysColorBrush(nIndex int32) HBRUSH {
	ret1 := syscall3(getSysColorBrush, 1,
		uintptr(nIndex),
		0,
		0)
	return HBRUSH(ret1)
}

func GetSystemMenu(hWnd HWND, bRevert bool) HMENU {
	ret1 := syscall3(getSystemMenu, 2,
		uintptr(hWnd),
		getUintptrFromBool(bRevert),
		0)
	return HMENU(ret1)
}

func GetSystemMetrics(nIndex int32) int32 {
	ret1 := syscall3(getSystemMetrics, 1,
		uintptr(nIndex),
		0,
		0)
	return int32(ret1)
}

func GetTabbedTextExtent(hdc HDC, lpString string, chCount int32, nTabPositions int32, lpnTabStopPositions /*const*/ *int32) uint32 {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall6(getTabbedTextExtent, 5,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(chCount),
		uintptr(nTabPositions),
		uintptr(unsafe.Pointer(lpnTabStopPositions)),
		0)
	return uint32(ret1)
}

func GetThreadDesktop(dwThreadId uint32) HDESK {
	ret1 := syscall3(getThreadDesktop, 1,
		uintptr(dwThreadId),
		0,
		0)
	return HDESK(ret1)
}

func GetTitleBarInfo(hwnd HWND, pti *TITLEBARINFO) bool {
	ret1 := syscall3(getTitleBarInfo, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pti)),
		0)
	return ret1 != 0
}

func GetTopWindow(hWnd HWND) HWND {
	ret1 := syscall3(getTopWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return HWND(ret1)
}

func GetTouchInputInfo(hTouchInput HTOUCHINPUT, cInputs uint32, pInputs *TOUCHINPUT, cbSize int32) bool {
	ret1 := syscall6(getTouchInputInfo, 4,
		uintptr(hTouchInput),
		uintptr(cInputs),
		uintptr(unsafe.Pointer(pInputs)),
		uintptr(cbSize),
		0,
		0)
	return ret1 != 0
}

func GetUpdateRect(hWnd HWND, lpRect *RECT, bErase bool) bool {
	ret1 := syscall3(getUpdateRect, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		getUintptrFromBool(bErase))
	return ret1 != 0
}

func GetUpdateRgn(hWnd HWND, hRgn HRGN, bErase bool) int32 {
	ret1 := syscall3(getUpdateRgn, 3,
		uintptr(hWnd),
		uintptr(hRgn),
		getUintptrFromBool(bErase))
	return int32(ret1)
}

func GetUserObjectInformation(hObj HANDLE, nIndex int32, pvInfo uintptr, nLength uint32, lpnLengthNeeded *uint32) bool {
	ret1 := syscall6(getUserObjectInformation, 5,
		uintptr(hObj),
		uintptr(nIndex),
		pvInfo,
		uintptr(nLength),
		uintptr(unsafe.Pointer(lpnLengthNeeded)),
		0)
	return ret1 != 0
}

func GetUserObjectSecurity(hObj HANDLE, pSIRequested *SECURITY_INFORMATION, pSID PSECURITY_DESCRIPTOR, nLength uint32, lpnLengthNeeded *uint32) bool {
	ret1 := syscall6(getUserObjectSecurity, 5,
		uintptr(hObj),
		uintptr(unsafe.Pointer(pSIRequested)),
		uintptr(unsafe.Pointer(pSID)),
		uintptr(nLength),
		uintptr(unsafe.Pointer(lpnLengthNeeded)),
		0)
	return ret1 != 0
}

func GetWindow(hWnd HWND, uCmd uint32) HWND {
	ret1 := syscall3(getWindow, 2,
		uintptr(hWnd),
		uintptr(uCmd),
		0)
	return HWND(ret1)
}

func GetWindowContextHelpId(unnamed0 HWND) uint32 {
	ret1 := syscall3(getWindowContextHelpId, 1,
		uintptr(unnamed0),
		0,
		0)
	return uint32(ret1)
}

func GetWindowDC(hWnd HWND) HDC {
	ret1 := syscall3(getWindowDC, 1,
		uintptr(hWnd),
		0,
		0)
	return HDC(ret1)
}

func GetWindowInfo(hwnd HWND, pwi *WINDOWINFO) bool {
	ret1 := syscall3(getWindowInfo, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pwi)),
		0)
	return ret1 != 0
}

func GetWindowLongPtr(hWnd HWND, nIndex int32) uintptr {
	ret1 := syscall3(getWindowLongPtr, 2,
		uintptr(hWnd),
		uintptr(nIndex),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func GetWindowLong(hWnd HWND, nIndex int32) LONG {
	ret1 := syscall3(getWindowLong, 2,
		uintptr(hWnd),
		uintptr(nIndex),
		0)
	return LONG(ret1)
}

func GetWindowModuleFileName(hwnd HWND, pszFileName LPWSTR, cchFileNameMax uint32) uint32 {
	ret1 := syscall3(getWindowModuleFileName, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pszFileName)),
		uintptr(cchFileNameMax))
	return uint32(ret1)
}

func GetWindowPlacement(hWnd HWND, lpwndpl *WINDOWPLACEMENT) bool {
	ret1 := syscall3(getWindowPlacement, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpwndpl)),
		0)
	return ret1 != 0
}

func GetWindowRect(hWnd HWND, lpRect *RECT) bool {
	ret1 := syscall3(getWindowRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		0)
	return ret1 != 0
}

func GetWindowRgn(hWnd HWND, hRgn HRGN) int32 {
	ret1 := syscall3(getWindowRgn, 2,
		uintptr(hWnd),
		uintptr(hRgn),
		0)
	return int32(ret1)
}

func GetWindowRgnBox(hWnd HWND, lprc *RECT) int32 {
	ret1 := syscall3(getWindowRgnBox, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lprc)),
		0)
	return int32(ret1)
}

func GetWindowTextLength(hWnd HWND) int32 {
	ret1 := syscall3(getWindowTextLength, 1,
		uintptr(hWnd),
		0,
		0)
	return int32(ret1)
}

func GetWindowText(hWnd HWND, lpString LPWSTR, nMaxCount int32) int32 {
	ret1 := syscall3(getWindowText, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(nMaxCount))
	return int32(ret1)
}

func GetWindowThreadProcessId(hWnd HWND, lpdwProcessId *uint32) uint32 {
	ret1 := syscall3(getWindowThreadProcessId, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpdwProcessId)),
		0)
	return uint32(ret1)
}

func GetWindowWord(hWnd HWND, nIndex int32) uint16 {
	ret1 := syscall3(getWindowWord, 2,
		uintptr(hWnd),
		uintptr(nIndex),
		0)
	return uint16(ret1)
}

func GrayString(hDC HDC, hBrush HBRUSH, lpOutputFunc GRAYSTRINGPROC, lpData LPARAM, nCount int32, x int32, y int32, nWidth int32, nHeight int32) bool {
	lpOutputFuncCallback := syscall.NewCallback(func(hdcRawArg HDC, lParamRawArg LPARAM, cchDataRawArg int) uintptr {
		ret := lpOutputFunc(hdcRawArg, lParamRawArg, cchDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall9(grayString, 9,
		uintptr(hDC),
		uintptr(hBrush),
		lpOutputFuncCallback,
		uintptr(lpData),
		uintptr(nCount),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight))
	return ret1 != 0
}

func HideCaret(hWnd HWND) bool {
	ret1 := syscall3(hideCaret, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func HiliteMenuItem(hWnd HWND, hMenu HMENU, uIDHiliteItem uint32, uHilite uint32) bool {
	ret1 := syscall6(hiliteMenuItem, 4,
		uintptr(hWnd),
		uintptr(hMenu),
		uintptr(uIDHiliteItem),
		uintptr(uHilite),
		0,
		0)
	return ret1 != 0
}

func IMPGetIME(unnamed0 HWND, unnamed1 *IMEPRO) bool {
	ret1 := syscall3(iMPGetIME, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func IMPQueryIME(unnamed0 *IMEPRO) bool {
	ret1 := syscall3(iMPQueryIME, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func IMPSetIME(unnamed0 HWND, unnamed1 *IMEPRO) bool {
	ret1 := syscall3(iMPSetIME, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func ImpersonateDdeClientWindow(hWndClient HWND, hWndServer HWND) bool {
	ret1 := syscall3(impersonateDdeClientWindow, 2,
		uintptr(hWndClient),
		uintptr(hWndServer),
		0)
	return ret1 != 0
}

func InSendMessage() bool {
	ret1 := syscall3(inSendMessage, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func InSendMessageEx(lpReserved LPVOID) uint32 {
	ret1 := syscall3(inSendMessageEx, 1,
		uintptr(unsafe.Pointer(lpReserved)),
		0,
		0)
	return uint32(ret1)
}

func InflateRect(lprc *RECT, dx int32, dy int32) bool {
	ret1 := syscall3(inflateRect, 3,
		uintptr(unsafe.Pointer(lprc)),
		uintptr(dx),
		uintptr(dy))
	return ret1 != 0
}

func InsertMenuItem(hmenu HMENU, item uint32, fByPosition bool, lpmi /*const*/ *MENUITEMINFO) bool {
	ret1 := syscall6(insertMenuItem, 4,
		uintptr(hmenu),
		uintptr(item),
		getUintptrFromBool(fByPosition),
		uintptr(unsafe.Pointer(lpmi)),
		0,
		0)
	return ret1 != 0
}

func InsertMenu(hMenu HMENU, uPosition uint32, uFlags uint32, uIDNewItem *uint32, lpNewItem string) bool {
	lpNewItemStr := unicode16FromString(lpNewItem)
	ret1 := syscall6(insertMenu, 5,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(uIDNewItem)),
		uintptr(unsafe.Pointer(&lpNewItemStr[0])),
		0)
	return ret1 != 0
}

func InternalGetWindowText(hWnd HWND, pString LPWSTR, cchMaxCount int32) int32 {
	ret1 := syscall3(internalGetWindowText, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pString)),
		uintptr(cchMaxCount))
	return int32(ret1)
}

func IntersectRect(lprcDst *RECT, lprcSrc1 /*const*/ *RECT, lprcSrc2 /*const*/ *RECT) bool {
	ret1 := syscall3(intersectRect, 3,
		uintptr(unsafe.Pointer(lprcDst)),
		uintptr(unsafe.Pointer(lprcSrc1)),
		uintptr(unsafe.Pointer(lprcSrc2)))
	return ret1 != 0
}

func InvalidateRect(hWnd HWND, lpRect /*const*/ *RECT, bErase bool) bool {
	ret1 := syscall3(invalidateRect, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		getUintptrFromBool(bErase))
	return ret1 != 0
}

func InvalidateRgn(hWnd HWND, hRgn HRGN, bErase bool) bool {
	ret1 := syscall3(invalidateRgn, 3,
		uintptr(hWnd),
		uintptr(hRgn),
		getUintptrFromBool(bErase))
	return ret1 != 0
}

func InvertRect(hDC HDC, lprc /*const*/ *RECT) bool {
	ret1 := syscall3(invertRect, 2,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		0)
	return ret1 != 0
}

func IsCharAlphaNumeric(ch WCHAR) bool {
	ret1 := syscall3(isCharAlphaNumeric, 1,
		uintptr(ch),
		0,
		0)
	return ret1 != 0
}

func IsCharAlpha(ch WCHAR) bool {
	ret1 := syscall3(isCharAlpha, 1,
		uintptr(ch),
		0,
		0)
	return ret1 != 0
}

func IsCharLower(ch WCHAR) bool {
	ret1 := syscall3(isCharLower, 1,
		uintptr(ch),
		0,
		0)
	return ret1 != 0
}

func IsCharUpper(ch WCHAR) bool {
	ret1 := syscall3(isCharUpper, 1,
		uintptr(ch),
		0,
		0)
	return ret1 != 0
}

func IsChild(hWndParent HWND, hWnd HWND) bool {
	ret1 := syscall3(isChild, 2,
		uintptr(hWndParent),
		uintptr(hWnd),
		0)
	return ret1 != 0
}

func IsClipboardFormatAvailable(format uint32) bool {
	ret1 := syscall3(isClipboardFormatAvailable, 1,
		uintptr(format),
		0,
		0)
	return ret1 != 0
}

func IsDialogMessage(hDlg HWND, lpMsg *MSG) bool {
	ret1 := syscall3(isDialogMessage, 2,
		uintptr(hDlg),
		uintptr(unsafe.Pointer(lpMsg)),
		0)
	return ret1 != 0
}

func IsDlgButtonChecked(hDlg HWND, nIDButton int32) uint32 {
	ret1 := syscall3(isDlgButtonChecked, 2,
		uintptr(hDlg),
		uintptr(nIDButton),
		0)
	return uint32(ret1)
}

func IsGUIThread(bConvert bool) bool {
	ret1 := syscall3(isGUIThread, 1,
		getUintptrFromBool(bConvert),
		0,
		0)
	return ret1 != 0
}

func IsHungAppWindow(hwnd HWND) bool {
	ret1 := syscall3(isHungAppWindow, 1,
		uintptr(hwnd),
		0,
		0)
	return ret1 != 0
}

func IsIconic(hWnd HWND) bool {
	ret1 := syscall3(isIconic, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func IsMenu(hMenu HMENU) bool {
	ret1 := syscall3(isMenu, 1,
		uintptr(hMenu),
		0,
		0)
	return ret1 != 0
}

func IsRectEmpty(lprc /*const*/ *RECT) bool {
	ret1 := syscall3(isRectEmpty, 1,
		uintptr(unsafe.Pointer(lprc)),
		0,
		0)
	return ret1 != 0
}

func IsTouchWindow(hwnd HWND, pulFlags *uint32) bool {
	ret1 := syscall3(isTouchWindow, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pulFlags)),
		0)
	return ret1 != 0
}

func IsWinEventHookInstalled(event uint32) bool {
	ret1 := syscall3(isWinEventHookInstalled, 1,
		uintptr(event),
		0,
		0)
	return ret1 != 0
}

func IsWindow(hWnd HWND) bool {
	ret1 := syscall3(isWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func IsWindowEnabled(hWnd HWND) bool {
	ret1 := syscall3(isWindowEnabled, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func IsWindowUnicode(hWnd HWND) bool {
	ret1 := syscall3(isWindowUnicode, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func IsWindowVisible(hWnd HWND) bool {
	ret1 := syscall3(isWindowVisible, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func IsWow64Message() bool {
	ret1 := syscall3(isWow64Message, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func IsZoomed(hWnd HWND) bool {
	ret1 := syscall3(isZoomed, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func KillTimer(hWnd HWND, uIDEvent *uint32) bool {
	ret1 := syscall3(killTimer, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(uIDEvent)),
		0)
	return ret1 != 0
}

func LoadAccelerators(hInstance HINSTANCE, lpTableName string) HACCEL {
	lpTableNameStr := unicode16FromString(lpTableName)
	ret1 := syscall3(loadAccelerators, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpTableNameStr[0])),
		0)
	return HACCEL(ret1)
}

func LoadBitmap(hInstance HINSTANCE, lpBitmapName string) HBITMAP {
	lpBitmapNameStr := unicode16FromString(lpBitmapName)
	ret1 := syscall3(loadBitmap, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpBitmapNameStr[0])),
		0)
	return HBITMAP(ret1)
}

func LoadCursorFromFile(lpFileName string) HCURSOR {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(loadCursorFromFile, 1,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return HCURSOR(ret1)
}

func LoadCursor(hInstance HINSTANCE, lpCursorName string) HCURSOR {
	lpCursorNameStr := unicode16FromString(lpCursorName)
	ret1 := syscall3(loadCursor, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpCursorNameStr[0])),
		0)
	return HCURSOR(ret1)
}

func LoadIcon(hInstance HINSTANCE, lpIconName string) HICON {
	lpIconNameStr := unicode16FromString(lpIconName)
	ret1 := syscall3(loadIcon, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpIconNameStr[0])),
		0)
	return HICON(ret1)
}

func LoadImage(hInst HINSTANCE, name string, aType uint32, cx int32, cy int32, fuLoad uint32) HANDLE {
	nameStr := unicode16FromString(name)
	ret1 := syscall6(loadImage, 6,
		uintptr(hInst),
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(aType),
		uintptr(cx),
		uintptr(cy),
		uintptr(fuLoad))
	return HANDLE(ret1)
}

func LoadKeyboardLayout(pwszKLID string, flags uint32) HKL {
	pwszKLIDStr := unicode16FromString(pwszKLID)
	ret1 := syscall3(loadKeyboardLayout, 2,
		uintptr(unsafe.Pointer(&pwszKLIDStr[0])),
		uintptr(flags),
		0)
	return HKL(ret1)
}

func LoadMenuIndirect(lpMenuTemplate /*const*/ uintptr) HMENU {
	ret1 := syscall3(loadMenuIndirect, 1,
		lpMenuTemplate,
		0,
		0)
	return HMENU(ret1)
}

func LoadMenu(hInstance HINSTANCE, lpMenuName string) HMENU {
	lpMenuNameStr := unicode16FromString(lpMenuName)
	ret1 := syscall3(loadMenu, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(&lpMenuNameStr[0])),
		0)
	return HMENU(ret1)
}

func LoadString(hInstance HINSTANCE, uID uint32, lpBuffer LPWSTR, cchBufferMax int32) int32 {
	ret1 := syscall6(loadString, 4,
		uintptr(hInstance),
		uintptr(uID),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cchBufferMax),
		0,
		0)
	return int32(ret1)
}

func LockSetForegroundWindow(uLockCode uint32) bool {
	ret1 := syscall3(lockSetForegroundWindow, 1,
		uintptr(uLockCode),
		0,
		0)
	return ret1 != 0
}

func LockWindowUpdate(hWndLock HWND) bool {
	ret1 := syscall3(lockWindowUpdate, 1,
		uintptr(hWndLock),
		0,
		0)
	return ret1 != 0
}

func LockWorkStation() bool {
	ret1 := syscall3(lockWorkStation, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func LookupIconIdFromDirectory(presbits *byte, fIcon bool) int32 {
	ret1 := syscall3(lookupIconIdFromDirectory, 2,
		uintptr(unsafe.Pointer(presbits)),
		getUintptrFromBool(fIcon),
		0)
	return int32(ret1)
}

func LookupIconIdFromDirectoryEx(presbits *byte, fIcon bool, cxDesired int32, cyDesired int32, flags uint32) int32 {
	ret1 := syscall6(lookupIconIdFromDirectoryEx, 5,
		uintptr(unsafe.Pointer(presbits)),
		getUintptrFromBool(fIcon),
		uintptr(cxDesired),
		uintptr(cyDesired),
		uintptr(flags),
		0)
	return int32(ret1)
}

func MapDialogRect(hDlg HWND, lpRect *RECT) bool {
	ret1 := syscall3(mapDialogRect, 2,
		uintptr(hDlg),
		uintptr(unsafe.Pointer(lpRect)),
		0)
	return ret1 != 0
}

func MapVirtualKeyEx(uCode uint32, uMapType uint32, dwhkl HKL) uint32 {
	ret1 := syscall3(mapVirtualKeyEx, 3,
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl))
	return uint32(ret1)
}

func MapVirtualKey(uCode uint32, uMapType uint32) uint32 {
	ret1 := syscall3(mapVirtualKey, 2,
		uintptr(uCode),
		uintptr(uMapType),
		0)
	return uint32(ret1)
}

func MapWindowPoints(hWndFrom HWND, hWndTo HWND, lpPoints *POINT, cPoints uint32) int32 {
	ret1 := syscall6(mapWindowPoints, 4,
		uintptr(hWndFrom),
		uintptr(hWndTo),
		uintptr(unsafe.Pointer(lpPoints)),
		uintptr(cPoints),
		0,
		0)
	return int32(ret1)
}

func MenuItemFromPoint(hWnd HWND, hMenu HMENU, ptScreen POINT) int32 {
	ret1 := syscall6(menuItemFromPoint, 4,
		uintptr(hWnd),
		uintptr(hMenu),
		uintptr(ptScreen.X),
		uintptr(ptScreen.Y),
		0,
		0)
	return int32(ret1)
}

func MessageBeep(uType uint32) bool {
	ret1 := syscall3(messageBeep, 1,
		uintptr(uType),
		0,
		0)
	return ret1 != 0
}

func MessageBoxEx(hWnd HWND, lpText string, lpCaption string, uType uint32, wLanguageId uint16) int32 {
	lpTextStr := unicode16FromString(lpText)
	lpCaptionStr := unicode16FromString(lpCaption)
	ret1 := syscall6(messageBoxEx, 5,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&lpTextStr[0])),
		uintptr(unsafe.Pointer(&lpCaptionStr[0])),
		uintptr(uType),
		uintptr(wLanguageId),
		0)
	return int32(ret1)
}

func MessageBoxIndirect(lpmbp /*const*/ *MSGBOXPARAMS) int32 {
	ret1 := syscall3(messageBoxIndirect, 1,
		uintptr(unsafe.Pointer(lpmbp)),
		0,
		0)
	return int32(ret1)
}

func MessageBox(hWnd HWND, lpText string, lpCaption string, uType uint32) int32 {
	lpTextStr := unicode16FromString(lpText)
	lpCaptionStr := unicode16FromString(lpCaption)
	ret1 := syscall6(messageBox, 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&lpTextStr[0])),
		uintptr(unsafe.Pointer(&lpCaptionStr[0])),
		uintptr(uType),
		0,
		0)
	return int32(ret1)
}

func ModifyMenu(hMnu HMENU, uPosition uint32, uFlags uint32, uIDNewItem *uint32, lpNewItem string) bool {
	lpNewItemStr := unicode16FromString(lpNewItem)
	ret1 := syscall6(modifyMenu, 5,
		uintptr(hMnu),
		uintptr(uPosition),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(uIDNewItem)),
		uintptr(unsafe.Pointer(&lpNewItemStr[0])),
		0)
	return ret1 != 0
}

func MonitorFromPoint(pt POINT, dwFlags uint32) HMONITOR {
	ret1 := syscall3(monitorFromPoint, 3,
		uintptr(pt.X),
		uintptr(pt.Y),
		uintptr(dwFlags))
	return HMONITOR(ret1)
}

func MonitorFromRect(lprc /*const*/ *RECT, dwFlags uint32) HMONITOR {
	ret1 := syscall3(monitorFromRect, 2,
		uintptr(unsafe.Pointer(lprc)),
		uintptr(dwFlags),
		0)
	return HMONITOR(ret1)
}

func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret1 := syscall3(monitorFromWindow, 2,
		uintptr(hwnd),
		uintptr(dwFlags),
		0)
	return HMONITOR(ret1)
}

func MoveWindow(hWnd HWND, x int32, y int32, nWidth int32, nHeight int32, bRepaint bool) bool {
	ret1 := syscall6(moveWindow, 6,
		uintptr(hWnd),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		getUintptrFromBool(bRepaint))
	return ret1 != 0
}

func MsgWaitForMultipleObjects(nCount uint32, pHandles /*const*/ *HANDLE, fWaitAll bool, dwMilliseconds uint32, dwWakeMask uint32) uint32 {
	ret1 := syscall6(msgWaitForMultipleObjects, 5,
		uintptr(nCount),
		uintptr(unsafe.Pointer(pHandles)),
		getUintptrFromBool(fWaitAll),
		uintptr(dwMilliseconds),
		uintptr(dwWakeMask),
		0)
	return uint32(ret1)
}

func MsgWaitForMultipleObjectsEx(nCount uint32, pHandles /*const*/ *HANDLE, dwMilliseconds uint32, dwWakeMask uint32, dwFlags uint32) uint32 {
	ret1 := syscall6(msgWaitForMultipleObjectsEx, 5,
		uintptr(nCount),
		uintptr(unsafe.Pointer(pHandles)),
		uintptr(dwMilliseconds),
		uintptr(dwWakeMask),
		uintptr(dwFlags),
		0)
	return uint32(ret1)
}

func NotifyWinEvent(event uint32, hwnd HWND, idObject LONG, idChild LONG) {
	syscall6(notifyWinEvent, 4,
		uintptr(event),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		0,
		0)
}

func OemKeyScan(wOemChar uint16) uint32 {
	ret1 := syscall3(oemKeyScan, 1,
		uintptr(wOemChar),
		0,
		0)
	return uint32(ret1)
}

func OemToCharBuff(lpszSrc /*const*/ LPCSTR, lpszDst LPWSTR, cchDstLength uint32) bool {
	ret1 := syscall3(oemToCharBuff, 3,
		uintptr(unsafe.Pointer(lpszSrc)),
		uintptr(unsafe.Pointer(lpszDst)),
		uintptr(cchDstLength))
	return ret1 != 0
}

func OemToChar(lpszSrc /*const*/ LPCSTR, lpszDst LPWSTR) bool {
	ret1 := syscall3(oemToChar, 2,
		uintptr(unsafe.Pointer(lpszSrc)),
		uintptr(unsafe.Pointer(lpszDst)),
		0)
	return ret1 != 0
}

func OffsetRect(lprc *RECT, dx int32, dy int32) bool {
	ret1 := syscall3(offsetRect, 3,
		uintptr(unsafe.Pointer(lprc)),
		uintptr(dx),
		uintptr(dy))
	return ret1 != 0
}

func OpenClipboard(hWndNewOwner HWND) bool {
	ret1 := syscall3(openClipboard, 1,
		uintptr(hWndNewOwner),
		0,
		0)
	return ret1 != 0
}

func OpenDesktop(lpszDesktop string, dwFlags uint32, fInherit bool, dwDesiredAccess ACCESS_MASK) HDESK {
	lpszDesktopStr := unicode16FromString(lpszDesktop)
	ret1 := syscall6(openDesktop, 4,
		uintptr(unsafe.Pointer(&lpszDesktopStr[0])),
		uintptr(dwFlags),
		getUintptrFromBool(fInherit),
		uintptr(dwDesiredAccess),
		0,
		0)
	return HDESK(ret1)
}

func OpenIcon(hWnd HWND) bool {
	ret1 := syscall3(openIcon, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func OpenInputDesktop(dwFlags uint32, fInherit bool, dwDesiredAccess ACCESS_MASK) HDESK {
	ret1 := syscall3(openInputDesktop, 3,
		uintptr(dwFlags),
		getUintptrFromBool(fInherit),
		uintptr(dwDesiredAccess))
	return HDESK(ret1)
}

func OpenWindowStation(lpszWinSta string, fInherit bool, dwDesiredAccess ACCESS_MASK) HWINSTA {
	lpszWinStaStr := unicode16FromString(lpszWinSta)
	ret1 := syscall3(openWindowStation, 3,
		uintptr(unsafe.Pointer(&lpszWinStaStr[0])),
		getUintptrFromBool(fInherit),
		uintptr(dwDesiredAccess))
	return HWINSTA(ret1)
}

func PackDDElParam(msg uint32, uiLo *uint32, uiHi *uint32) LPARAM {
	ret1 := syscall3(packDDElParam, 3,
		uintptr(msg),
		uintptr(unsafe.Pointer(uiLo)),
		uintptr(unsafe.Pointer(uiHi)))
	return LPARAM(ret1)
}

func PaintDesktop(hdc HDC) bool {
	ret1 := syscall3(paintDesktop, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func PeekMessage(lpMsg *MSG, hWnd HWND, wMsgFilterMin uint32, wMsgFilterMax uint32, wRemoveMsg uint32) bool {
	ret1 := syscall6(peekMessage, 5,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
		0)
	return ret1 != 0
}

func PostMessage(hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM) bool {
	ret1 := syscall6(postMessage, 4,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return ret1 != 0
}

func PostQuitMessage(nExitCode int32) {
	syscall3(postQuitMessage, 1,
		uintptr(nExitCode),
		0,
		0)
}

func PostThreadMessage(idThread uint32, msg uint32, wParam WPARAM, lParam LPARAM) bool {
	ret1 := syscall6(postThreadMessage, 4,
		uintptr(idThread),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return ret1 != 0
}

func PrintWindow(hwnd HWND, hdcBlt HDC, nFlags uint32) bool {
	ret1 := syscall3(printWindow, 3,
		uintptr(hwnd),
		uintptr(hdcBlt),
		uintptr(nFlags))
	return ret1 != 0
}

func PrivateExtractIcons(szFileName string, nIconIndex int32, cxIcon int32, cyIcon int32, phicon *HICON, piconid *UINT, nIcons uint32, flags uint32) uint32 {
	szFileNameStr := unicode16FromString(szFileName)
	ret1 := syscall9(privateExtractIcons, 8,
		uintptr(unsafe.Pointer(&szFileNameStr[0])),
		uintptr(nIconIndex),
		uintptr(cxIcon),
		uintptr(cyIcon),
		uintptr(unsafe.Pointer(phicon)),
		uintptr(unsafe.Pointer(piconid)),
		uintptr(nIcons),
		uintptr(flags),
		0)
	return uint32(ret1)
}

func PtInRect(lprc /*const*/ *RECT, pt POINT) bool {
	ret1 := syscall3(ptInRect, 3,
		uintptr(unsafe.Pointer(lprc)),
		uintptr(pt.X),
		uintptr(pt.Y))
	return ret1 != 0
}

func RealChildWindowFromPoint(hwndParent HWND, ptParentClientCoords POINT) HWND {
	ret1 := syscall3(realChildWindowFromPoint, 3,
		uintptr(hwndParent),
		uintptr(ptParentClientCoords.X),
		uintptr(ptParentClientCoords.Y))
	return HWND(ret1)
}

func RealGetWindowClass(hwnd HWND, ptszClassName LPWSTR, cchClassNameMax uint32) uint32 {
	ret1 := syscall3(realGetWindowClass, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(ptszClassName)),
		uintptr(cchClassNameMax))
	return uint32(ret1)
}

func RedrawWindow(hWnd HWND, lprcUpdate /*const*/ *RECT, hrgnUpdate HRGN, flags uint32) bool {
	ret1 := syscall6(redrawWindow, 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lprcUpdate)),
		uintptr(hrgnUpdate),
		uintptr(flags),
		0,
		0)
	return ret1 != 0
}

func RegisterClassEx(unnamed0 /*const*/ *WNDCLASSEX) ATOM {
	ret1 := syscall3(registerClassEx, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ATOM(ret1)
}

func RegisterClass(lpWndClass /*const*/ *WNDCLASS) ATOM {
	ret1 := syscall3(registerClass, 1,
		uintptr(unsafe.Pointer(lpWndClass)),
		0,
		0)
	return ATOM(ret1)
}

func RegisterClipboardFormat(lpszFormat string) uint32 {
	lpszFormatStr := unicode16FromString(lpszFormat)
	ret1 := syscall3(registerClipboardFormat, 1,
		uintptr(unsafe.Pointer(&lpszFormatStr[0])),
		0,
		0)
	return uint32(ret1)
}

func RegisterDeviceNotification(hRecipient HANDLE, notificationFilter LPVOID, flags uint32) HDEVNOTIFY {
	ret1 := syscall3(registerDeviceNotification, 3,
		uintptr(hRecipient),
		uintptr(unsafe.Pointer(notificationFilter)),
		uintptr(flags))
	return HDEVNOTIFY(ret1)
}

func RegisterHotKey(hWnd HWND, id int32, fsModifiers uint32, vk uint32) bool {
	ret1 := syscall6(registerHotKey, 4,
		uintptr(hWnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vk),
		0,
		0)
	return ret1 != 0
}

func RegisterPowerSettingNotification(hRecipient HANDLE, powerSettingGuid /*const*/ *GUID, flags uint32) HPOWERNOTIFY {
	ret1 := syscall3(registerPowerSettingNotification, 3,
		uintptr(hRecipient),
		uintptr(unsafe.Pointer(powerSettingGuid)),
		uintptr(flags))
	return HPOWERNOTIFY(ret1)
}

func RegisterRawInputDevices(pRawInputDevices /*const*/ *RAWINPUTDEVICE, uiNumDevices uint32, cbSize uint32) bool {
	ret1 := syscall3(registerRawInputDevices, 3,
		uintptr(unsafe.Pointer(pRawInputDevices)),
		uintptr(uiNumDevices),
		uintptr(cbSize))
	return ret1 != 0
}

func RegisterShellHookWindow(hwnd HWND) bool {
	ret1 := syscall3(registerShellHookWindow, 1,
		uintptr(hwnd),
		0,
		0)
	return ret1 != 0
}

func RegisterTouchWindow(hwnd HWND, ulFlags ULONG) bool {
	ret1 := syscall3(registerTouchWindow, 2,
		uintptr(hwnd),
		uintptr(ulFlags),
		0)
	return ret1 != 0
}

func RegisterWindowMessage(lpString string) uint32 {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(registerWindowMessage, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return uint32(ret1)
}

func ReleaseCapture() bool {
	ret1 := syscall3(releaseCapture, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func ReleaseDC(hWnd HWND, hDC HDC) int32 {
	ret1 := syscall3(releaseDC, 2,
		uintptr(hWnd),
		uintptr(hDC),
		0)
	return int32(ret1)
}

func RemoveMenu(hMenu HMENU, uPosition uint32, uFlags uint32) bool {
	ret1 := syscall3(removeMenu, 3,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags))
	return ret1 != 0
}

func RemoveProp(hWnd HWND, lpString string) HANDLE {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(removeProp, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0)
	return HANDLE(ret1)
}

func ReplyMessage(lResult LRESULT) bool {
	ret1 := syscall3(replyMessage, 1,
		uintptr(lResult),
		0,
		0)
	return ret1 != 0
}

func ReuseDDElParam(lParam LPARAM, msgIn uint32, msgOut uint32, uiLo *uint32, uiHi *uint32) LPARAM {
	ret1 := syscall6(reuseDDElParam, 5,
		uintptr(lParam),
		uintptr(msgIn),
		uintptr(msgOut),
		uintptr(unsafe.Pointer(uiLo)),
		uintptr(unsafe.Pointer(uiHi)),
		0)
	return LPARAM(ret1)
}

func ScreenToClient(hWnd HWND, lpPoint *POINT) bool {
	ret1 := syscall3(screenToClient, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPoint)),
		0)
	return ret1 != 0
}

func ScrollDC(hDC HDC, dx int32, dy int32, lprcScroll /*const*/ *RECT, lprcClip /*const*/ *RECT, hrgnUpdate HRGN, lprcUpdate *RECT) bool {
	ret1 := syscall9(scrollDC, 7,
		uintptr(hDC),
		uintptr(dx),
		uintptr(dy),
		uintptr(unsafe.Pointer(lprcScroll)),
		uintptr(unsafe.Pointer(lprcClip)),
		uintptr(hrgnUpdate),
		uintptr(unsafe.Pointer(lprcUpdate)),
		0,
		0)
	return ret1 != 0
}

func ScrollWindow(hWnd HWND, xAmount int32, yAmount int32, lpRect /*const*/ *RECT, lpClipRect /*const*/ *RECT) bool {
	ret1 := syscall6(scrollWindow, 5,
		uintptr(hWnd),
		uintptr(xAmount),
		uintptr(yAmount),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(unsafe.Pointer(lpClipRect)),
		0)
	return ret1 != 0
}

func ScrollWindowEx(hWnd HWND, dx int32, dy int32, prcScroll /*const*/ *RECT, prcClip /*const*/ *RECT, hrgnUpdate HRGN, prcUpdate *RECT, flags uint32) int32 {
	ret1 := syscall9(scrollWindowEx, 8,
		uintptr(hWnd),
		uintptr(dx),
		uintptr(dy),
		uintptr(unsafe.Pointer(prcScroll)),
		uintptr(unsafe.Pointer(prcClip)),
		uintptr(hrgnUpdate),
		uintptr(unsafe.Pointer(prcUpdate)),
		uintptr(flags),
		0)
	return int32(ret1)
}

func SendDlgItemMessage(hDlg HWND, nIDDlgItem int32, msg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(sendDlgItemMessage, 5,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0)
	return LRESULT(ret1)
}

func SendIMEMessageEx(unnamed0 HWND, unnamed1 LPARAM) LRESULT {
	ret1 := syscall3(sendIMEMessageEx, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return LRESULT(ret1)
}

func SendInput(cInputs uint32, pInputs *INPUT, cbSize int32) uint32 {
	ret1 := syscall3(sendInput, 3,
		uintptr(cInputs),
		uintptr(unsafe.Pointer(pInputs)),
		uintptr(cbSize))
	return uint32(ret1)
}

func SendMessageCallback(hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM, lpResultCallBack SENDASYNCPROC, dwData *uint32) bool {
	lpResultCallBackCallback := syscall.NewCallback(lpResultCallBack)
	ret1 := syscall6(sendMessageCallback, 6,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		lpResultCallBackCallback,
		uintptr(unsafe.Pointer(dwData)))
	return ret1 != 0
}

func SendMessageTimeout(hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM, fuFlags uint32, uTimeout uint32, lpdwResult *uintptr) LRESULT {
	ret1 := syscall9(sendMessageTimeout, 7,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		uintptr(fuFlags),
		uintptr(uTimeout),
		uintptr(unsafe.Pointer(lpdwResult)),
		0,
		0)
	return LRESULT(ret1)
}

func SendMessage(hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret1 := syscall6(sendMessage, 4,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return LRESULT(ret1)
}

func SendNotifyMessage(hWnd HWND, msg uint32, wParam WPARAM, lParam LPARAM) bool {
	ret1 := syscall6(sendNotifyMessage, 4,
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
	return ret1 != 0
}

func SetActiveWindow(hWnd HWND) HWND {
	ret1 := syscall3(setActiveWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return HWND(ret1)
}

func SetCapture(hWnd HWND) HWND {
	ret1 := syscall3(setCapture, 1,
		uintptr(hWnd),
		0,
		0)
	return HWND(ret1)
}

func SetCaretBlinkTime(uMSeconds uint32) bool {
	ret1 := syscall3(setCaretBlinkTime, 1,
		uintptr(uMSeconds),
		0,
		0)
	return ret1 != 0
}

func SetCaretPos(x int32, y int32) bool {
	ret1 := syscall3(setCaretPos, 2,
		uintptr(x),
		uintptr(y),
		0)
	return ret1 != 0
}

func SetClassLongPtr(hWnd HWND, nIndex int32, dwNewLong uintptr) *uint32 {
	ret1 := syscall3(setClassLongPtr, 3,
		uintptr(hWnd),
		uintptr(nIndex),
		dwNewLong)
	return (*uint32)(unsafe.Pointer(ret1))
}

func SetClassLong(hWnd HWND, nIndex int32, dwNewLong LONG) uint32 {
	ret1 := syscall3(setClassLong, 3,
		uintptr(hWnd),
		uintptr(nIndex),
		uintptr(dwNewLong))
	return uint32(ret1)
}

func SetClassWord(hWnd HWND, nIndex int32, wNewWord uint16) uint16 {
	ret1 := syscall3(setClassWord, 3,
		uintptr(hWnd),
		uintptr(nIndex),
		uintptr(wNewWord))
	return uint16(ret1)
}

func SetClipboardData(uFormat uint32, hMem HANDLE) HANDLE {
	ret1 := syscall3(setClipboardData, 2,
		uintptr(uFormat),
		uintptr(hMem),
		0)
	return HANDLE(ret1)
}

func SetClipboardViewer(hWndNewViewer HWND) HWND {
	ret1 := syscall3(setClipboardViewer, 1,
		uintptr(hWndNewViewer),
		0,
		0)
	return HWND(ret1)
}

func SetCursor(hCursor HCURSOR) HCURSOR {
	ret1 := syscall3(setCursor, 1,
		uintptr(hCursor),
		0,
		0)
	return HCURSOR(ret1)
}

func SetCursorPos(x int32, y int32) bool {
	ret1 := syscall3(setCursorPos, 2,
		uintptr(x),
		uintptr(y),
		0)
	return ret1 != 0
}

func SetDebugErrorLevel(dwLevel uint32) {
	syscall3(setDebugErrorLevel, 1,
		uintptr(dwLevel),
		0,
		0)
}

func SetDlgItemInt(hDlg HWND, nIDDlgItem int32, uValue uint32, bSigned bool) bool {
	ret1 := syscall6(setDlgItemInt, 4,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		uintptr(uValue),
		getUintptrFromBool(bSigned),
		0,
		0)
	return ret1 != 0
}

func SetDlgItemText(hDlg HWND, nIDDlgItem int32, lpString string) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(setDlgItemText, 3,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		uintptr(unsafe.Pointer(&lpStringStr[0])))
	return ret1 != 0
}

func SetDoubleClickTime(unnamed0 uint32) bool {
	ret1 := syscall3(setDoubleClickTime, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func SetFocus(hWnd HWND) HWND {
	ret1 := syscall3(setFocus, 1,
		uintptr(hWnd),
		0,
		0)
	return HWND(ret1)
}

func SetForegroundWindow(hWnd HWND) bool {
	ret1 := syscall3(setForegroundWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func SetGestureConfig(hwnd HWND, dwReserved uint32, cIDs uint32, pGestureConfig *GESTURECONFIG, cbSize uint32) bool {
	ret1 := syscall6(setGestureConfig, 5,
		uintptr(hwnd),
		uintptr(dwReserved),
		uintptr(cIDs),
		uintptr(unsafe.Pointer(pGestureConfig)),
		uintptr(cbSize),
		0)
	return ret1 != 0
}

func SetKeyboardState(lpKeyState *byte) bool {
	ret1 := syscall3(setKeyboardState, 1,
		uintptr(unsafe.Pointer(lpKeyState)),
		0,
		0)
	return ret1 != 0
}

func SetLastErrorEx(dwErrCode uint32, dwType uint32) {
	syscall3(setLastErrorEx, 2,
		uintptr(dwErrCode),
		uintptr(dwType),
		0)
}

func SetLayeredWindowAttributes(hwnd HWND, crKey COLORREF, bAlpha byte, dwFlags uint32) bool {
	ret1 := syscall6(setLayeredWindowAttributes, 4,
		uintptr(hwnd),
		uintptr(crKey),
		uintptr(bAlpha),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func SetMenu(hWnd HWND, hMenu HMENU) bool {
	ret1 := syscall3(setMenu, 2,
		uintptr(hWnd),
		uintptr(hMenu),
		0)
	return ret1 != 0
}

func SetMenuContextHelpId(unnamed0 HMENU, unnamed1 uint32) bool {
	ret1 := syscall3(setMenuContextHelpId, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func SetMenuDefaultItem(hMenu HMENU, uItem uint32, fByPos uint32) bool {
	ret1 := syscall3(setMenuDefaultItem, 3,
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(fByPos))
	return ret1 != 0
}

func SetMenuInfo(unnamed0 HMENU, unnamed1 /*const*/ *MENUINFO) bool {
	ret1 := syscall3(setMenuInfo, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return ret1 != 0
}

func SetMenuItemBitmaps(hMenu HMENU, uPosition uint32, uFlags uint32, hBitmapUnchecked HBITMAP, hBitmapChecked HBITMAP) bool {
	ret1 := syscall6(setMenuItemBitmaps, 5,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags),
		uintptr(hBitmapUnchecked),
		uintptr(hBitmapChecked),
		0)
	return ret1 != 0
}

func SetMenuItemInfo(hmenu HMENU, item uint32, fByPositon bool, lpmii /*const*/ *MENUITEMINFO) bool {
	ret1 := syscall6(setMenuItemInfo, 4,
		uintptr(hmenu),
		uintptr(item),
		getUintptrFromBool(fByPositon),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)
	return ret1 != 0
}

func SetMessageExtraInfo(lParam LPARAM) LPARAM {
	ret1 := syscall3(setMessageExtraInfo, 1,
		uintptr(lParam),
		0,
		0)
	return LPARAM(ret1)
}

func SetMessageQueue(cMessagesMax int32) bool {
	ret1 := syscall3(setMessageQueue, 1,
		uintptr(cMessagesMax),
		0,
		0)
	return ret1 != 0
}

func SetParent(hWndChild HWND, hWndNewParent HWND) HWND {
	ret1 := syscall3(setParent, 2,
		uintptr(hWndChild),
		uintptr(hWndNewParent),
		0)
	return HWND(ret1)
}

func SetProcessDefaultLayout(dwDefaultLayout uint32) bool {
	ret1 := syscall3(setProcessDefaultLayout, 1,
		uintptr(dwDefaultLayout),
		0,
		0)
	return ret1 != 0
}

func SetProcessWindowStation(hWinSta HWINSTA) bool {
	ret1 := syscall3(setProcessWindowStation, 1,
		uintptr(hWinSta),
		0,
		0)
	return ret1 != 0
}

func SetProp(hWnd HWND, lpString string, hData HANDLE) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(setProp, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(hData))
	return ret1 != 0
}

func SetRect(lprc *RECT, xLeft int32, yTop int32, xRight int32, yBottom int32) bool {
	ret1 := syscall6(setRect, 5,
		uintptr(unsafe.Pointer(lprc)),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(xRight),
		uintptr(yBottom),
		0)
	return ret1 != 0
}

func SetRectEmpty(lprc *RECT) bool {
	ret1 := syscall3(setRectEmpty, 1,
		uintptr(unsafe.Pointer(lprc)),
		0,
		0)
	return ret1 != 0
}

func SetScrollInfo(hwnd HWND, nBar int32, lpsi /*const*/ *SCROLLINFO, redraw bool) int32 {
	ret1 := syscall6(setScrollInfo, 4,
		uintptr(hwnd),
		uintptr(nBar),
		uintptr(unsafe.Pointer(lpsi)),
		getUintptrFromBool(redraw),
		0,
		0)
	return int32(ret1)
}

func SetScrollPos(hWnd HWND, nBar int32, nPos int32, bRedraw bool) int32 {
	ret1 := syscall6(setScrollPos, 4,
		uintptr(hWnd),
		uintptr(nBar),
		uintptr(nPos),
		getUintptrFromBool(bRedraw),
		0,
		0)
	return int32(ret1)
}

func SetScrollRange(hWnd HWND, nBar int32, nMinPos int32, nMaxPos int32, bRedraw bool) bool {
	ret1 := syscall6(setScrollRange, 5,
		uintptr(hWnd),
		uintptr(nBar),
		uintptr(nMinPos),
		uintptr(nMaxPos),
		getUintptrFromBool(bRedraw),
		0)
	return ret1 != 0
}

func SetSysColors(cElements int32, lpaElements /*const*/ *int32, lpaRgbValues /*const*/ *COLORREF) bool {
	ret1 := syscall3(setSysColors, 3,
		uintptr(cElements),
		uintptr(unsafe.Pointer(lpaElements)),
		uintptr(unsafe.Pointer(lpaRgbValues)))
	return ret1 != 0
}

func SetSystemCursor(hcur HCURSOR, id uint32) bool {
	ret1 := syscall3(setSystemCursor, 2,
		uintptr(hcur),
		uintptr(id),
		0)
	return ret1 != 0
}

func SetThreadDesktop(hDesktop HDESK) bool {
	ret1 := syscall3(setThreadDesktop, 1,
		uintptr(hDesktop),
		0,
		0)
	return ret1 != 0
}

func SetTimer(hWnd HWND, nIDEvent *uint32, uElapse uint32, lpTimerFunc TIMERPROC) *uint32 {
	lpTimerFuncCallback := syscall.NewCallback(lpTimerFunc)
	ret1 := syscall6(setTimer, 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(nIDEvent)),
		uintptr(uElapse),
		lpTimerFuncCallback,
		0,
		0)
	return (*uint32)(unsafe.Pointer(ret1))
}

func SetUserObjectInformation(hObj HANDLE, nIndex int32, pvInfo uintptr, nLength uint32) bool {
	ret1 := syscall6(setUserObjectInformation, 4,
		uintptr(hObj),
		uintptr(nIndex),
		pvInfo,
		uintptr(nLength),
		0,
		0)
	return ret1 != 0
}

func SetUserObjectSecurity(hObj HANDLE, pSIRequested *SECURITY_INFORMATION, pSID PSECURITY_DESCRIPTOR) bool {
	ret1 := syscall3(setUserObjectSecurity, 3,
		uintptr(hObj),
		uintptr(unsafe.Pointer(pSIRequested)),
		uintptr(unsafe.Pointer(pSID)))
	return ret1 != 0
}

func SetWinEventHook(eventMin uint32, eventMax uint32, hmodWinEventProc HMODULE, pfnWinEventProc WINEVENTPROC, idProcess uint32, idThread uint32, dwFlags uint32) HWINEVENTHOOK {
	pfnWinEventProcCallback := syscall.NewCallback(pfnWinEventProc)
	ret1 := syscall9(setWinEventHook, 7,
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		pfnWinEventProcCallback,
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
		0,
		0)
	return HWINEVENTHOOK(ret1)
}

func SetWindowContextHelpId(unnamed0 HWND, unnamed1 uint32) bool {
	ret1 := syscall3(setWindowContextHelpId, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func SetWindowLongPtr(hWnd HWND, nIndex int32, dwNewLong uintptr) uintptr {
	ret1 := syscall3(setWindowLongPtr, 3,
		uintptr(hWnd),
		uintptr(nIndex),
		dwNewLong)
	return (uintptr)(unsafe.Pointer(ret1))
}

func SetWindowLong(hWnd HWND, nIndex int32, dwNewLong LONG) LONG {
	ret1 := syscall3(setWindowLong, 3,
		uintptr(hWnd),
		uintptr(nIndex),
		uintptr(dwNewLong))
	return LONG(ret1)
}

func SetWindowPlacement(hWnd HWND, lpwndpl /*const*/ *WINDOWPLACEMENT) bool {
	ret1 := syscall3(setWindowPlacement, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpwndpl)),
		0)
	return ret1 != 0
}

func SetWindowPos(hWnd HWND, hWndInsertAfter HWND, x int32, y int32, cx int32, cy int32, uFlags uint32) bool {
	ret1 := syscall9(setWindowPos, 7,
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags),
		0,
		0)
	return ret1 != 0
}

func SetWindowRgn(hWnd HWND, hRgn HRGN, bRedraw bool) int32 {
	ret1 := syscall3(setWindowRgn, 3,
		uintptr(hWnd),
		uintptr(hRgn),
		getUintptrFromBool(bRedraw))
	return int32(ret1)
}

func SetWindowText(hWnd HWND, lpString string) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(setWindowText, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0)
	return ret1 != 0
}

func SetWindowWord(hWnd HWND, nIndex int32, wNewWord uint16) uint16 {
	ret1 := syscall3(setWindowWord, 3,
		uintptr(hWnd),
		uintptr(nIndex),
		uintptr(wNewWord))
	return uint16(ret1)
}

func SetWindowsHookEx(idHook int32, lpfn HOOKPROC, hmod HINSTANCE, dwThreadId uint32) HHOOK {
	lpfnCallback := syscall.NewCallback(func(codeRawArg int32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := lpfn(codeRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(setWindowsHookEx, 4,
		uintptr(idHook),
		lpfnCallback,
		uintptr(hmod),
		uintptr(dwThreadId),
		0,
		0)
	return HHOOK(ret1)
}

func SetWindowsHook(nFilterType int32, pfnFilterProc HOOKPROC) HHOOK {
	pfnFilterProcCallback := syscall.NewCallback(func(codeRawArg int32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := pfnFilterProc(codeRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(setWindowsHook, 2,
		uintptr(nFilterType),
		pfnFilterProcCallback,
		0)
	return HHOOK(ret1)
}

func ShowCaret(hWnd HWND) bool {
	ret1 := syscall3(showCaret, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func ShowCursor(bShow bool) int32 {
	ret1 := syscall3(showCursor, 1,
		getUintptrFromBool(bShow),
		0,
		0)
	return int32(ret1)
}

func ShowOwnedPopups(hWnd HWND, fShow bool) bool {
	ret1 := syscall3(showOwnedPopups, 2,
		uintptr(hWnd),
		getUintptrFromBool(fShow),
		0)
	return ret1 != 0
}

func ShowScrollBar(hWnd HWND, wBar int32, bShow bool) bool {
	ret1 := syscall3(showScrollBar, 3,
		uintptr(hWnd),
		uintptr(wBar),
		getUintptrFromBool(bShow))
	return ret1 != 0
}

func ShowWindow(hWnd HWND, nCmdShow int32) bool {
	ret1 := syscall3(showWindow, 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)
	return ret1 != 0
}

func ShowWindowAsync(hWnd HWND, nCmdShow int32) bool {
	ret1 := syscall3(showWindowAsync, 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)
	return ret1 != 0
}

func SubtractRect(lprcDst *RECT, lprcSrc1 /*const*/ *RECT, lprcSrc2 /*const*/ *RECT) bool {
	ret1 := syscall3(subtractRect, 3,
		uintptr(unsafe.Pointer(lprcDst)),
		uintptr(unsafe.Pointer(lprcSrc1)),
		uintptr(unsafe.Pointer(lprcSrc2)))
	return ret1 != 0
}

func SwapMouseButton(fSwap bool) bool {
	ret1 := syscall3(swapMouseButton, 1,
		getUintptrFromBool(fSwap),
		0,
		0)
	return ret1 != 0
}

func SwitchDesktop(hDesktop HDESK) bool {
	ret1 := syscall3(switchDesktop, 1,
		uintptr(hDesktop),
		0,
		0)
	return ret1 != 0
}

func SwitchToThisWindow(hwnd HWND, fUnknown bool) {
	syscall3(switchToThisWindow, 2,
		uintptr(hwnd),
		getUintptrFromBool(fUnknown),
		0)
}

func SystemParametersInfo(uiAction uint32, uiParam uint32, pvParam uintptr, fWinIni uint32) bool {
	ret1 := syscall6(systemParametersInfo, 4,
		uintptr(uiAction),
		uintptr(uiParam),
		pvParam,
		uintptr(fWinIni),
		0,
		0)
	return ret1 != 0
}

func TabbedTextOut(hdc HDC, x int32, y int32, lpString string, chCount int32, nTabPositions int32, lpnTabStopPositions /*const*/ *int32, nTabOrigin int32) LONG {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall9(tabbedTextOut, 8,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(chCount),
		uintptr(nTabPositions),
		uintptr(unsafe.Pointer(lpnTabStopPositions)),
		uintptr(nTabOrigin),
		0)
	return LONG(ret1)
}

func TileWindows(hwndParent HWND, wHow uint32, lpRect /*const*/ *RECT, cKids uint32, lpKids /*const*/ *HWND) uint16 {
	ret1 := syscall6(tileWindows, 5,
		uintptr(hwndParent),
		uintptr(wHow),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(cKids),
		uintptr(unsafe.Pointer(lpKids)),
		0)
	return uint16(ret1)
}

func ToAscii(uVirtKey uint32, uScanCode uint32, lpKeyState /*const*/ *byte, lpChar *uint16, uFlags uint32) int32 {
	ret1 := syscall6(toAscii, 5,
		uintptr(uVirtKey),
		uintptr(uScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(lpChar)),
		uintptr(uFlags),
		0)
	return int32(ret1)
}

func ToAsciiEx(uVirtKey uint32, uScanCode uint32, lpKeyState /*const*/ *byte, lpChar *uint16, uFlags uint32, dwhkl HKL) int32 {
	ret1 := syscall6(toAsciiEx, 6,
		uintptr(uVirtKey),
		uintptr(uScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(lpChar)),
		uintptr(uFlags),
		uintptr(dwhkl))
	return int32(ret1)
}

func ToUnicode(wVirtKey uint32, wScanCode uint32, lpKeyState /*const*/ *byte, pwszBuff LPWSTR, cchBuff int32, wFlags uint32) int32 {
	ret1 := syscall6(toUnicode, 6,
		uintptr(wVirtKey),
		uintptr(wScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(pwszBuff)),
		uintptr(cchBuff),
		uintptr(wFlags))
	return int32(ret1)
}

func ToUnicodeEx(wVirtKey uint32, wScanCode uint32, lpKeyState /*const*/ *byte, pwszBuff LPWSTR, cchBuff int32, wFlags uint32, dwhkl HKL) int32 {
	ret1 := syscall9(toUnicodeEx, 7,
		uintptr(wVirtKey),
		uintptr(wScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(pwszBuff)),
		uintptr(cchBuff),
		uintptr(wFlags),
		uintptr(dwhkl),
		0,
		0)
	return int32(ret1)
}

func TrackMouseEvent(lpEventTrack *TRACKMOUSEEVENT) bool {
	ret1 := syscall3(trackMouseEvent, 1,
		uintptr(unsafe.Pointer(lpEventTrack)),
		0,
		0)
	return ret1 != 0
}

func TrackPopupMenu(hMenu HMENU, uFlags uint32, x int32, y int32, nReserved int32, hWnd HWND, prcRect /*const*/ *RECT) bool {
	ret1 := syscall9(trackPopupMenu, 7,
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(x),
		uintptr(y),
		uintptr(nReserved),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(prcRect)),
		0,
		0)
	return ret1 != 0
}

func TrackPopupMenuEx(unnamed0 HMENU, unnamed1 uint32, unnamed2 int32, unnamed3 int32, unnamed4 HWND, unnamed5 *TPMPARAMS) bool {
	ret1 := syscall6(trackPopupMenuEx, 6,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		uintptr(unnamed4),
		uintptr(unsafe.Pointer(unnamed5)))
	return ret1 != 0
}

func TranslateAccelerator(hWnd HWND, hAccTable HACCEL, lpMsg *MSG) int32 {
	ret1 := syscall3(translateAccelerator, 3,
		uintptr(hWnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)))
	return int32(ret1)
}

func TranslateMDISysAccel(hWndClient HWND, lpMsg *MSG) bool {
	ret1 := syscall3(translateMDISysAccel, 2,
		uintptr(hWndClient),
		uintptr(unsafe.Pointer(lpMsg)),
		0)
	return ret1 != 0
}

func TranslateMessage(lpMsg /*const*/ *MSG) bool {
	ret1 := syscall3(translateMessage, 1,
		uintptr(unsafe.Pointer(lpMsg)),
		0,
		0)
	return ret1 != 0
}

func UnhookWinEvent(hWinEventHook HWINEVENTHOOK) bool {
	ret1 := syscall3(unhookWinEvent, 1,
		uintptr(hWinEventHook),
		0,
		0)
	return ret1 != 0
}

func UnhookWindowsHook(nCode int32, pfnFilterProc HOOKPROC) bool {
	pfnFilterProcCallback := syscall.NewCallback(func(codeRawArg int32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := pfnFilterProc(codeRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(unhookWindowsHook, 2,
		uintptr(nCode),
		pfnFilterProcCallback,
		0)
	return ret1 != 0
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret1 := syscall3(unhookWindowsHookEx, 1,
		uintptr(hhk),
		0,
		0)
	return ret1 != 0
}

func UnionRect(lprcDst *RECT, lprcSrc1 /*const*/ *RECT, lprcSrc2 /*const*/ *RECT) bool {
	ret1 := syscall3(unionRect, 3,
		uintptr(unsafe.Pointer(lprcDst)),
		uintptr(unsafe.Pointer(lprcSrc1)),
		uintptr(unsafe.Pointer(lprcSrc2)))
	return ret1 != 0
}

func UnloadKeyboardLayout(hkl HKL) bool {
	ret1 := syscall3(unloadKeyboardLayout, 1,
		uintptr(hkl),
		0,
		0)
	return ret1 != 0
}

func UnpackDDElParam(msg uint32, lParam LPARAM, puiLo *uintptr, puiHi *uintptr) bool {
	ret1 := syscall6(unpackDDElParam, 4,
		uintptr(msg),
		uintptr(lParam),
		uintptr(unsafe.Pointer(puiLo)),
		uintptr(unsafe.Pointer(puiHi)),
		0,
		0)
	return ret1 != 0
}

func UnregisterClass(lpClassName string, hInstance HINSTANCE) bool {
	lpClassNameStr := unicode16FromString(lpClassName)
	ret1 := syscall3(unregisterClass, 2,
		uintptr(unsafe.Pointer(&lpClassNameStr[0])),
		uintptr(hInstance),
		0)
	return ret1 != 0
}

func UnregisterDeviceNotification(handle HDEVNOTIFY) bool {
	ret1 := syscall3(unregisterDeviceNotification, 1,
		uintptr(handle),
		0,
		0)
	return ret1 != 0
}

func UnregisterHotKey(hWnd HWND, id int32) bool {
	ret1 := syscall3(unregisterHotKey, 2,
		uintptr(hWnd),
		uintptr(id),
		0)
	return ret1 != 0
}

func UnregisterPowerSettingNotification(handle HPOWERNOTIFY) bool {
	ret1 := syscall3(unregisterPowerSettingNotification, 1,
		uintptr(handle),
		0,
		0)
	return ret1 != 0
}

func UnregisterTouchWindow(hwnd HWND) bool {
	ret1 := syscall3(unregisterTouchWindow, 1,
		uintptr(hwnd),
		0,
		0)
	return ret1 != 0
}

func UpdateLayeredWindow(hWnd HWND, hdcDst HDC, pptDst *POINT, psize *SIZE, hdcSrc HDC, pptSrc *POINT, crKey COLORREF, pblend *BLENDFUNCTION, dwFlags uint32) bool {
	ret1 := syscall9(updateLayeredWindow, 9,
		uintptr(hWnd),
		uintptr(hdcDst),
		uintptr(unsafe.Pointer(pptDst)),
		uintptr(unsafe.Pointer(psize)),
		uintptr(hdcSrc),
		uintptr(unsafe.Pointer(pptSrc)),
		uintptr(crKey),
		uintptr(unsafe.Pointer(pblend)),
		uintptr(dwFlags))
	return ret1 != 0
}

func UpdateLayeredWindowIndirect(hWnd HWND, pULWInfo /*const*/ *UPDATELAYEREDWINDOWINFO) bool {
	ret1 := syscall3(updateLayeredWindowIndirect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pULWInfo)),
		0)
	return ret1 != 0
}

func UpdateWindow(hWnd HWND) bool {
	ret1 := syscall3(updateWindow, 1,
		uintptr(hWnd),
		0,
		0)
	return ret1 != 0
}

func UserHandleGrantAccess(hUserHandle HANDLE, hJob HANDLE, bGrant bool) bool {
	ret1 := syscall3(userHandleGrantAccess, 3,
		uintptr(hUserHandle),
		uintptr(hJob),
		getUintptrFromBool(bGrant))
	return ret1 != 0
}

func ValidateRect(hWnd HWND, lpRect /*const*/ *RECT) bool {
	ret1 := syscall3(validateRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		0)
	return ret1 != 0
}

func ValidateRgn(hWnd HWND, hRgn HRGN) bool {
	ret1 := syscall3(validateRgn, 2,
		uintptr(hWnd),
		uintptr(hRgn),
		0)
	return ret1 != 0
}

func VkKeyScanEx(ch WCHAR, dwhkl HKL) int16 {
	ret1 := syscall3(vkKeyScanEx, 2,
		uintptr(ch),
		uintptr(dwhkl),
		0)
	return int16(ret1)
}

func VkKeyScan(ch WCHAR) int16 {
	ret1 := syscall3(vkKeyScan, 1,
		uintptr(ch),
		0,
		0)
	return int16(ret1)
}

func WINNLSEnableIME(unnamed0 HWND, unnamed1 bool) bool {
	ret1 := syscall3(wINNLSEnableIME, 2,
		uintptr(unnamed0),
		getUintptrFromBool(unnamed1),
		0)
	return ret1 != 0
}

func WINNLSGetEnableStatus(unnamed0 HWND) bool {
	ret1 := syscall3(wINNLSGetEnableStatus, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func WINNLSGetIMEHotkey(unnamed0 HWND) uint32 {
	ret1 := syscall3(wINNLSGetIMEHotkey, 1,
		uintptr(unnamed0),
		0,
		0)
	return uint32(ret1)
}

func WaitForInputIdle(hProcess HANDLE, dwMilliseconds uint32) uint32 {
	ret1 := syscall3(waitForInputIdle, 2,
		uintptr(hProcess),
		uintptr(dwMilliseconds),
		0)
	return uint32(ret1)
}

func WaitMessage() bool {
	ret1 := syscall3(waitMessage, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func WinHelp(hWndMain HWND, lpszHelp string, uCommand uint32, dwData *uint32) bool {
	lpszHelpStr := unicode16FromString(lpszHelp)
	ret1 := syscall6(winHelp, 4,
		uintptr(hWndMain),
		uintptr(unsafe.Pointer(&lpszHelpStr[0])),
		uintptr(uCommand),
		uintptr(unsafe.Pointer(dwData)),
		0,
		0)
	return ret1 != 0
}

func WindowFromDC(hDC HDC) HWND {
	ret1 := syscall3(windowFromDC, 1,
		uintptr(hDC),
		0,
		0)
	return HWND(ret1)
}

func WindowFromPoint(point POINT) HWND {
	ret1 := syscall3(windowFromPoint, 2,
		uintptr(point.X),
		uintptr(point.Y),
		0)
	return HWND(ret1)
}

func Keybd_event(bVk byte, bScan byte, dwFlags uint32, dwExtraInfo *uint32) {
	syscall6(keybd_event, 4,
		uintptr(bVk),
		uintptr(bScan),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(dwExtraInfo)),
		0,
		0)
}

func Mouse_event(dwFlags uint32, dx uint32, dy uint32, dwData uint32, dwExtraInfo *uint32) {
	syscall6(mouse_event, 5,
		uintptr(dwFlags),
		uintptr(dx),
		uintptr(dy),
		uintptr(dwData),
		uintptr(unsafe.Pointer(dwExtraInfo)),
		0)
}

func AlignRects(rect *RECT, b uint32, c uint32, d uint32) bool {
	ret1 := syscall6(alignRects, 4,
		uintptr(unsafe.Pointer(rect)),
		uintptr(b),
		uintptr(c),
		uintptr(d),
		0,
		0)
	return ret1 != 0
}

func CascadeChildWindows(parent HWND, flags uint32) uint16 {
	ret1 := syscall3(cascadeChildWindows, 2,
		uintptr(parent),
		uintptr(flags),
		0)
	return uint16(ret1)
}

func CreateDialogIndirectParamAorW(hInst HINSTANCE, dlgTemplate /*const*/ uintptr, owner HWND, dlgProc DLGPROC, param LPARAM, flags uint32) HWND {
	dlgProcCallback := syscall.NewCallback(func(hwndDlgRawArg HWND, uMsgRawArg uint32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := dlgProc(hwndDlgRawArg, uMsgRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(createDialogIndirectParamAorW, 6,
		uintptr(hInst),
		dlgTemplate,
		uintptr(owner),
		dlgProcCallback,
		uintptr(param),
		uintptr(flags))
	return HWND(ret1)
}

func DialogBoxIndirectParamAorW(hInstance HINSTANCE, template /*const*/ uintptr, owner HWND, dlgProc DLGPROC, param LPARAM, flags uint32) INT_PTR {
	dlgProcCallback := syscall.NewCallback(func(hwndDlgRawArg HWND, uMsgRawArg uint32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := dlgProc(hwndDlgRawArg, uMsgRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(dialogBoxIndirectParamAorW, 6,
		uintptr(hInstance),
		template,
		uintptr(owner),
		dlgProcCallback,
		uintptr(param),
		uintptr(flags))
	return (INT_PTR)(unsafe.Pointer(ret1))
}

func DrawCaptionTemp(hwnd HWND, hdc HDC, rect /*const*/ *RECT, hFont HFONT, hIcon HICON, str string, uFlags uint32) bool {
	strStr := unicode16FromString(str)
	ret1 := syscall9(drawCaptionTemp, 7,
		uintptr(hwnd),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rect)),
		uintptr(hFont),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(&strStr[0])),
		uintptr(uFlags),
		0,
		0)
	return ret1 != 0
}

func DrawMenuBarTemp(hwnd HWND, hDC HDC, lprect *RECT, hMenu HMENU, hFont HFONT) uint32 {
	ret1 := syscall6(drawMenuBarTemp, 5,
		uintptr(hwnd),
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprect)),
		uintptr(hMenu),
		uintptr(hFont),
		0)
	return uint32(ret1)
}

func GetAppCompatFlags(hTask HTASK) uint32 {
	ret1 := syscall3(getAppCompatFlags, 1,
		uintptr(hTask),
		0,
		0)
	return uint32(ret1)
}

func GetAppCompatFlags2(hTask HTASK) uint32 {
	ret1 := syscall3(getAppCompatFlags2, 1,
		uintptr(hTask),
		0,
		0)
	return uint32(ret1)
}

func GetCursorFrameInfo(hCursor HCURSOR, reserved uint32, istep uint32, rate_jiffies *uint32, num_steps *uint32) HCURSOR {
	ret1 := syscall6(getCursorFrameInfo, 5,
		uintptr(hCursor),
		uintptr(reserved),
		uintptr(istep),
		uintptr(unsafe.Pointer(rate_jiffies)),
		uintptr(unsafe.Pointer(num_steps)),
		0)
	return HCURSOR(ret1)
}

func GetInternalWindowPos(hwnd HWND, rectWnd *RECT, ptIcon *POINT) uint32 {
	ret1 := syscall3(getInternalWindowPos, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rectWnd)),
		uintptr(unsafe.Pointer(ptIcon)))
	return uint32(ret1)
}

func GetProgmanWindow() HWND {
	ret1 := syscall3(getProgmanWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func GetTaskmanWindow() HWND {
	ret1 := syscall3(getTaskmanWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

func KillSystemTimer(hwnd HWND, id *uint32) bool {
	ret1 := syscall3(killSystemTimer, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(id)),
		0)
	return ret1 != 0
}

func LoadLocalFonts() {
	syscall3(loadLocalFonts, 0,
		0,
		0,
		0)
}

func MessageBoxTimeout(hWnd HWND, text string, title string, aType uint32, langid uint16, timeout uint32) int32 {
	textStr := unicode16FromString(text)
	titleStr := unicode16FromString(title)
	ret1 := syscall6(messageBoxTimeout, 6,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&textStr[0])),
		uintptr(unsafe.Pointer(&titleStr[0])),
		uintptr(aType),
		uintptr(langid),
		uintptr(timeout))
	return int32(ret1)
}

func PrivateExtractIconEx(lpwstrFile string, nIndex int32, phIconLarge *HICON, phIconSmall *HICON, nIcons uint32) uint32 {
	lpwstrFileStr := unicode16FromString(lpwstrFile)
	ret1 := syscall6(privateExtractIconEx, 5,
		uintptr(unsafe.Pointer(&lpwstrFileStr[0])),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(phIconLarge)),
		uintptr(unsafe.Pointer(phIconSmall)),
		uintptr(nIcons),
		0)
	return uint32(ret1)
}

func RegisterLogonProcess(hprocess HANDLE, x bool) uint32 {
	ret1 := syscall3(registerLogonProcess, 2,
		uintptr(hprocess),
		getUintptrFromBool(x),
		0)
	return uint32(ret1)
}

func RegisterServicesProcess(servicesProcessId uint32) int32 {
	ret1 := syscall3(registerServicesProcess, 1,
		uintptr(servicesProcessId),
		0,
		0)
	return int32(ret1)
}

func RegisterSystemThread(flags uint32, reserved uint32) {
	syscall3(registerSystemThread, 2,
		uintptr(flags),
		uintptr(reserved),
		0)
}

func RegisterTasklist(x uint32) uint32 {
	ret1 := syscall3(registerTasklist, 1,
		uintptr(x),
		0,
		0)
	return uint32(ret1)
}

func ScrollChildren(hWnd HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) {
	syscall6(scrollChildren, 4,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(wParam),
		uintptr(lParam),
		0,
		0)
}

func SetInternalWindowPos(hwnd HWND, showCmd uint32, rect *RECT, pt *POINT) {
	syscall6(setInternalWindowPos, 4,
		uintptr(hwnd),
		uintptr(showCmd),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(pt)),
		0,
		0)
}

func SetLogonNotifyWindow(hwinsta HWINSTA, hwnd HWND) uint32 {
	ret1 := syscall3(setLogonNotifyWindow, 2,
		uintptr(hwinsta),
		uintptr(hwnd),
		0)
	return uint32(ret1)
}

func SetProgmanWindow(hwnd HWND) HWND {
	ret1 := syscall3(setProgmanWindow, 1,
		uintptr(hwnd),
		0,
		0)
	return HWND(ret1)
}

func SetShellWindow(hwndShell HWND) bool {
	ret1 := syscall3(setShellWindow, 1,
		uintptr(hwndShell),
		0,
		0)
	return ret1 != 0
}

func SetShellWindowEx(hwndShell HWND, hwndListView HWND) bool {
	ret1 := syscall3(setShellWindowEx, 2,
		uintptr(hwndShell),
		uintptr(hwndListView),
		0)
	return ret1 != 0
}

func SetSysColorsTemp(pPens /*const*/ *COLORREF, pBrushes /*const*/ *HBRUSH, n *uint32) *uint32 {
	ret1 := syscall3(setSysColorsTemp, 3,
		uintptr(unsafe.Pointer(pPens)),
		uintptr(unsafe.Pointer(pBrushes)),
		uintptr(unsafe.Pointer(n)))
	return (*uint32)(unsafe.Pointer(ret1))
}

func SetSystemMenu(hwnd HWND, hMenu HMENU) bool {
	ret1 := syscall3(setSystemMenu, 2,
		uintptr(hwnd),
		uintptr(hMenu),
		0)
	return ret1 != 0
}

func SetSystemTimer(hwnd HWND, id *uint32, timeout uint32, proc TIMERPROC) *uint32 {
	procCallback := syscall.NewCallback(proc)
	ret1 := syscall6(setSystemTimer, 4,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(id)),
		uintptr(timeout),
		procCallback,
		0,
		0)
	return (*uint32)(unsafe.Pointer(ret1))
}

func SetTaskmanWindow(hwnd HWND) HWND {
	ret1 := syscall3(setTaskmanWindow, 1,
		uintptr(hwnd),
		0,
		0)
	return HWND(ret1)
}

func SetWindowStationUser(x1 uint32, x2 uint32) uint32 {
	ret1 := syscall3(setWindowStationUser, 2,
		uintptr(x1),
		uintptr(x2),
		0)
	return uint32(ret1)
}

func TileChildWindows(parent HWND, flags uint32) uint16 {
	ret1 := syscall3(tileChildWindows, 2,
		uintptr(parent),
		uintptr(flags),
		0)
	return uint16(ret1)
}

func User32InitializeImmEntryTable(unnamed0 uint32) bool {
	ret1 := syscall3(user32InitializeImmEntryTable, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func UserRealizePalette(hDC HDC) uint32 {
	ret1 := syscall3(userRealizePalette, 1,
		uintptr(hDC),
		0,
		0)
	return uint32(ret1)
}

func UserRegisterWowHandlers(aNew /*const*/ uintptr, orig uintptr) {
	syscall3(userRegisterWowHandlers, 2,
		aNew,
		orig,
		0)
}
