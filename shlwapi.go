// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libshlwapi uintptr

	// Functions
	assocCreate                uintptr
	assocGetPerceivedType      uintptr
	assocIsDangerous           uintptr
	assocQueryKey              uintptr
	assocQueryStringByKey      uintptr
	assocQueryString           uintptr
	chrCmpI                    uintptr
	colorAdjustLuma            uintptr
	colorHLSToRGB              uintptr
	colorRGBToHLS              uintptr
	connectToConnectionPoint   uintptr
	getAcceptLanguages         uintptr
	getMenuPosFromID           uintptr
	hashData                   uintptr
	iStream_Reset              uintptr
	iStream_Size               uintptr
	iUnknown_AtomicRelease     uintptr
	iUnknown_GetSite           uintptr
	iUnknown_GetWindow         uintptr
	iUnknown_QueryService      uintptr
	iUnknown_Set               uintptr
	iUnknown_SetSite           uintptr
	intlStrEqWorker            uintptr
	isCharSpace                uintptr
	isInternetESCEnabled       uintptr
	isOS                       uintptr
	mLFreeLibrary              uintptr
	mLLoadLibrary              uintptr
	parseURL                   uintptr
	pathAddBackslash           uintptr
	pathAddExtension           uintptr
	pathAppend                 uintptr
	pathBuildRoot              uintptr
	pathCanonicalize           uintptr
	pathCombine                uintptr
	pathCommonPrefix           uintptr
	pathCompactPathEx          uintptr
	pathCompactPath            uintptr
	pathCreateFromUrlAlloc     uintptr
	pathCreateFromUrl          uintptr
	pathFileExists             uintptr
	pathFindExtension          uintptr
	pathFindFileName           uintptr
	pathFindNextComponent      uintptr
	pathFindOnPath             uintptr
	pathFindSuffixArray        uintptr
	pathGetArgs                uintptr
	pathGetCharType            uintptr
	pathGetDriveNumber         uintptr
	pathIsContentType          uintptr
	pathIsDirectoryEmpty       uintptr
	pathIsDirectory            uintptr
	pathIsFileSpec             uintptr
	pathIsLFNFileSpec          uintptr
	pathIsNetworkPath          uintptr
	pathIsPrefix               uintptr
	pathIsRelative             uintptr
	pathIsRoot                 uintptr
	pathIsSameRoot             uintptr
	pathIsSystemFolder         uintptr
	pathIsUNCServerShare       uintptr
	pathIsUNCServer            uintptr
	pathIsUNC                  uintptr
	pathIsURL                  uintptr
	pathMakePretty             uintptr
	pathMakeSystemFolder       uintptr
	pathMatchSpec              uintptr
	pathParseIconLocation      uintptr
	pathQuoteSpaces            uintptr
	pathRelativePathTo         uintptr
	pathRemoveArgs             uintptr
	pathRemoveBackslash        uintptr
	pathRemoveBlanks           uintptr
	pathRemoveExtension        uintptr
	pathRemoveFileSpec         uintptr
	pathRenameExtension        uintptr
	pathSearchAndQualify       uintptr
	pathSetDlgItemPath         uintptr
	pathSkipRoot               uintptr
	pathStripPath              uintptr
	pathStripToRoot            uintptr
	pathUnExpandEnvStrings     uintptr
	pathUndecorate             uintptr
	pathUnmakeSystemFolder     uintptr
	pathUnquoteSpaces          uintptr
	qISearch                   uintptr
	sHAllocShared              uintptr
	sHAnsiToAnsi               uintptr
	sHAnsiToUnicode            uintptr
	sHAutoComplete             uintptr
	sHCopyKey                  uintptr
	sHCreateShellPalette       uintptr
	sHCreateStreamOnFileEx     uintptr
	sHCreateStreamOnFile       uintptr
	sHCreateStreamWrapper      uintptr
	sHCreateThread             uintptr
	sHCreateThreadRef          uintptr
	sHDeleteEmptyKey           uintptr
	sHDeleteKey                uintptr
	sHDeleteOrphanKey          uintptr
	sHDeleteValue              uintptr
	sHEnumKeyEx                uintptr
	sHEnumValue                uintptr
	sHFormatDateTime           uintptr
	sHFreeShared               uintptr
	sHGetInverseCMAP           uintptr
	sHGetThreadRef             uintptr
	sHGetValue                 uintptr
	sHGetViewStatePropertyBag  uintptr
	sHIsChildOrSelf            uintptr
	sHIsLowMemoryMachine       uintptr
	sHLoadIndirectString       uintptr
	sHLockShared               uintptr
	sHMessageBoxCheck          uintptr
	sHQueryInfoKey             uintptr
	sHQueryValueEx             uintptr
	sHRegCloseUSKey            uintptr
	sHRegCreateUSKey           uintptr
	sHRegDeleteEmptyUSKey      uintptr
	sHRegDeleteUSValue         uintptr
	sHRegDuplicateHKey         uintptr
	sHRegEnumUSKey             uintptr
	sHRegEnumUSValue           uintptr
	sHRegGetBoolUSValue        uintptr
	sHRegGetIntW               uintptr
	sHRegGetPath               uintptr
	sHRegGetUSValue            uintptr
	sHRegOpenUSKey             uintptr
	sHRegQueryInfoUSKey        uintptr
	sHRegQueryUSValue          uintptr
	sHRegSetPath               uintptr
	sHRegSetUSValue            uintptr
	sHRegWriteUSValue          uintptr
	sHRegisterValidateTemplate uintptr
	sHReleaseThreadRef         uintptr
	sHSendMessageBroadcast     uintptr
	sHSetThreadRef             uintptr
	sHSetValue                 uintptr
	sHSkipJunction             uintptr
	sHStrDup                   uintptr
	sHStripMneumonic           uintptr
	sHUnicodeToAnsi            uintptr
	sHUnicodeToUnicode         uintptr
	sHUnlockShared             uintptr
	strCSpnI                   uintptr
	strCSpn                    uintptr
	strCatBuff                 uintptr
	strCatChainW               uintptr
	strCatW                    uintptr
	strChrI                    uintptr
	strChrNW                   uintptr
	strChr                     uintptr
	strCmpC                    uintptr
	strCmpIC                   uintptr
	strCmpIW                   uintptr
	strCmpLogicalW             uintptr
	strCmpNC                   uintptr
	strCmpNIC                  uintptr
	strCmpNI                   uintptr
	strCmpN                    uintptr
	strCmpW                    uintptr
	strCpyNW                   uintptr
	strCpyW                    uintptr
	strDup                     uintptr
	strFormatByteSize64A       uintptr
	strFormatByteSize          uintptr
	strFormatKBSize            uintptr
	strFromTimeInterval        uintptr
	strIsIntlEqual             uintptr
	strNCat                    uintptr
	strPBrk                    uintptr
	strRChrI                   uintptr
	strRChr                    uintptr
	strRStrI                   uintptr
	strRetToBSTR               uintptr
	strRetToBuf                uintptr
	strRetToStr                uintptr
	strSpn                     uintptr
	strStrI                    uintptr
	strStrNIW                  uintptr
	strStrNW                   uintptr
	strStr                     uintptr
	strToInt64Ex               uintptr
	strToIntEx                 uintptr
	strToInt                   uintptr
	strTrim                    uintptr
	urlApplyScheme             uintptr
	urlCanonicalize            uintptr
	urlCombine                 uintptr
	urlCompare                 uintptr
	urlCreateFromPath          uintptr
	urlEscape                  uintptr
	urlFixupW                  uintptr
	urlGetLocation             uintptr
	urlGetPart                 uintptr
	urlHash                    uintptr
	urlIsNoHistory             uintptr
	urlIsOpaque                uintptr
	urlIs                      uintptr
	urlUnescape                uintptr
	whichPlatform              uintptr
)

func init() {
	// Library
	libshlwapi = doLoadLibrary("shlwapi.dll")

	// Functions
	assocCreate = doGetProcAddress(libshlwapi, "AssocCreate")
	assocGetPerceivedType = doGetProcAddress(libshlwapi, "AssocGetPerceivedType")
	assocIsDangerous = doGetProcAddress(libshlwapi, "AssocIsDangerous")
	assocQueryKey = doGetProcAddress(libshlwapi, "AssocQueryKeyW")
	assocQueryStringByKey = doGetProcAddress(libshlwapi, "AssocQueryStringByKeyW")
	assocQueryString = doGetProcAddress(libshlwapi, "AssocQueryStringW")
	chrCmpI = doGetProcAddress(libshlwapi, "ChrCmpIW")
	colorAdjustLuma = doGetProcAddress(libshlwapi, "ColorAdjustLuma")
	colorHLSToRGB = doGetProcAddress(libshlwapi, "ColorHLSToRGB")
	colorRGBToHLS = doGetProcAddress(libshlwapi, "ColorRGBToHLS")
	connectToConnectionPoint = doGetProcAddress(libshlwapi, "ConnectToConnectionPoint")
	getAcceptLanguages = doGetProcAddress(libshlwapi, "GetAcceptLanguagesW")
	getMenuPosFromID = doGetProcAddress(libshlwapi, "GetMenuPosFromID")
	hashData = doGetProcAddress(libshlwapi, "HashData")
	iStream_Reset = doGetProcAddress(libshlwapi, "IStream_Reset")
	iStream_Size = doGetProcAddress(libshlwapi, "IStream_Size")
	iUnknown_AtomicRelease = doGetProcAddress(libshlwapi, "IUnknown_AtomicRelease")
	iUnknown_GetSite = doGetProcAddress(libshlwapi, "IUnknown_GetSite")
	iUnknown_GetWindow = doGetProcAddress(libshlwapi, "IUnknown_GetWindow")
	iUnknown_QueryService = doGetProcAddress(libshlwapi, "IUnknown_QueryService")
	iUnknown_Set = doGetProcAddress(libshlwapi, "IUnknown_Set")
	iUnknown_SetSite = doGetProcAddress(libshlwapi, "IUnknown_SetSite")
	intlStrEqWorker = doGetProcAddress(libshlwapi, "IntlStrEqWorkerW")
	isCharSpace = doGetProcAddress(libshlwapi, "IsCharSpaceW")
	isInternetESCEnabled = doGetProcAddress(libshlwapi, "IsInternetESCEnabled")
	isOS = doGetProcAddress(libshlwapi, "IsOS")
	mLFreeLibrary = doGetProcAddress(libshlwapi, "MLFreeLibrary")
	mLLoadLibrary = doGetProcAddress(libshlwapi, "MLLoadLibraryW")
	parseURL = doGetProcAddress(libshlwapi, "ParseURLW")
	pathAddBackslash = doGetProcAddress(libshlwapi, "PathAddBackslashW")
	pathAddExtension = doGetProcAddress(libshlwapi, "PathAddExtensionW")
	pathAppend = doGetProcAddress(libshlwapi, "PathAppendW")
	pathBuildRoot = doGetProcAddress(libshlwapi, "PathBuildRootW")
	pathCanonicalize = doGetProcAddress(libshlwapi, "PathCanonicalizeW")
	pathCombine = doGetProcAddress(libshlwapi, "PathCombineW")
	pathCommonPrefix = doGetProcAddress(libshlwapi, "PathCommonPrefixW")
	pathCompactPathEx = doGetProcAddress(libshlwapi, "PathCompactPathExW")
	pathCompactPath = doGetProcAddress(libshlwapi, "PathCompactPathW")
	pathCreateFromUrlAlloc = doGetProcAddress(libshlwapi, "PathCreateFromUrlAlloc")
	pathCreateFromUrl = doGetProcAddress(libshlwapi, "PathCreateFromUrlW")
	pathFileExists = doGetProcAddress(libshlwapi, "PathFileExistsW")
	pathFindExtension = doGetProcAddress(libshlwapi, "PathFindExtensionW")
	pathFindFileName = doGetProcAddress(libshlwapi, "PathFindFileNameW")
	pathFindNextComponent = doGetProcAddress(libshlwapi, "PathFindNextComponentW")
	pathFindOnPath = doGetProcAddress(libshlwapi, "PathFindOnPathW")
	pathFindSuffixArray = doGetProcAddress(libshlwapi, "PathFindSuffixArrayW")
	pathGetArgs = doGetProcAddress(libshlwapi, "PathGetArgsW")
	pathGetCharType = doGetProcAddress(libshlwapi, "PathGetCharTypeW")
	pathGetDriveNumber = doGetProcAddress(libshlwapi, "PathGetDriveNumberW")
	pathIsContentType = doGetProcAddress(libshlwapi, "PathIsContentTypeW")
	pathIsDirectoryEmpty = doGetProcAddress(libshlwapi, "PathIsDirectoryEmptyW")
	pathIsDirectory = doGetProcAddress(libshlwapi, "PathIsDirectoryW")
	pathIsFileSpec = doGetProcAddress(libshlwapi, "PathIsFileSpecW")
	pathIsLFNFileSpec = doGetProcAddress(libshlwapi, "PathIsLFNFileSpecW")
	pathIsNetworkPath = doGetProcAddress(libshlwapi, "PathIsNetworkPathW")
	pathIsPrefix = doGetProcAddress(libshlwapi, "PathIsPrefixW")
	pathIsRelative = doGetProcAddress(libshlwapi, "PathIsRelativeW")
	pathIsRoot = doGetProcAddress(libshlwapi, "PathIsRootW")
	pathIsSameRoot = doGetProcAddress(libshlwapi, "PathIsSameRootW")
	pathIsSystemFolder = doGetProcAddress(libshlwapi, "PathIsSystemFolderW")
	pathIsUNCServerShare = doGetProcAddress(libshlwapi, "PathIsUNCServerShareW")
	pathIsUNCServer = doGetProcAddress(libshlwapi, "PathIsUNCServerW")
	pathIsUNC = doGetProcAddress(libshlwapi, "PathIsUNCW")
	pathIsURL = doGetProcAddress(libshlwapi, "PathIsURLW")
	pathMakePretty = doGetProcAddress(libshlwapi, "PathMakePrettyW")
	pathMakeSystemFolder = doGetProcAddress(libshlwapi, "PathMakeSystemFolderW")
	pathMatchSpec = doGetProcAddress(libshlwapi, "PathMatchSpecW")
	pathParseIconLocation = doGetProcAddress(libshlwapi, "PathParseIconLocationW")
	pathQuoteSpaces = doGetProcAddress(libshlwapi, "PathQuoteSpacesW")
	pathRelativePathTo = doGetProcAddress(libshlwapi, "PathRelativePathToW")
	pathRemoveArgs = doGetProcAddress(libshlwapi, "PathRemoveArgsW")
	pathRemoveBackslash = doGetProcAddress(libshlwapi, "PathRemoveBackslashW")
	pathRemoveBlanks = doGetProcAddress(libshlwapi, "PathRemoveBlanksW")
	pathRemoveExtension = doGetProcAddress(libshlwapi, "PathRemoveExtensionW")
	pathRemoveFileSpec = doGetProcAddress(libshlwapi, "PathRemoveFileSpecW")
	pathRenameExtension = doGetProcAddress(libshlwapi, "PathRenameExtensionW")
	pathSearchAndQualify = doGetProcAddress(libshlwapi, "PathSearchAndQualifyW")
	pathSetDlgItemPath = doGetProcAddress(libshlwapi, "PathSetDlgItemPathW")
	pathSkipRoot = doGetProcAddress(libshlwapi, "PathSkipRootW")
	pathStripPath = doGetProcAddress(libshlwapi, "PathStripPathW")
	pathStripToRoot = doGetProcAddress(libshlwapi, "PathStripToRootW")
	pathUnExpandEnvStrings = doGetProcAddress(libshlwapi, "PathUnExpandEnvStringsW")
	pathUndecorate = doGetProcAddress(libshlwapi, "PathUndecorateW")
	pathUnmakeSystemFolder = doGetProcAddress(libshlwapi, "PathUnmakeSystemFolderW")
	pathUnquoteSpaces = doGetProcAddress(libshlwapi, "PathUnquoteSpacesW")
	qISearch = doGetProcAddress(libshlwapi, "QISearch")
	sHAllocShared = doGetProcAddress(libshlwapi, "SHAllocShared")
	sHAnsiToAnsi = doGetProcAddress(libshlwapi, "SHAnsiToAnsi")
	sHAnsiToUnicode = doGetProcAddress(libshlwapi, "SHAnsiToUnicode")
	sHAutoComplete = doGetProcAddress(libshlwapi, "SHAutoComplete")
	sHCopyKey = doGetProcAddress(libshlwapi, "SHCopyKeyW")
	sHCreateShellPalette = doGetProcAddress(libshlwapi, "SHCreateShellPalette")
	sHCreateStreamOnFileEx = doGetProcAddress(libshlwapi, "SHCreateStreamOnFileEx")
	sHCreateStreamOnFile = doGetProcAddress(libshlwapi, "SHCreateStreamOnFileW")
	sHCreateStreamWrapper = doGetProcAddress(libshlwapi, "SHCreateStreamWrapper")
	sHCreateThread = doGetProcAddress(libshlwapi, "SHCreateThread")
	sHCreateThreadRef = doGetProcAddress(libshlwapi, "SHCreateThreadRef")
	sHDeleteEmptyKey = doGetProcAddress(libshlwapi, "SHDeleteEmptyKeyW")
	sHDeleteKey = doGetProcAddress(libshlwapi, "SHDeleteKeyW")
	sHDeleteOrphanKey = doGetProcAddress(libshlwapi, "SHDeleteOrphanKeyW")
	sHDeleteValue = doGetProcAddress(libshlwapi, "SHDeleteValueW")
	sHEnumKeyEx = doGetProcAddress(libshlwapi, "SHEnumKeyExW")
	sHEnumValue = doGetProcAddress(libshlwapi, "SHEnumValueW")
	sHFormatDateTime = doGetProcAddress(libshlwapi, "SHFormatDateTimeW")
	sHFreeShared = doGetProcAddress(libshlwapi, "SHFreeShared")
	sHGetInverseCMAP = doGetProcAddress(libshlwapi, "SHGetInverseCMAP")
	sHGetThreadRef = doGetProcAddress(libshlwapi, "SHGetThreadRef")
	sHGetValue = doGetProcAddress(libshlwapi, "SHGetValueW")
	sHGetViewStatePropertyBag = doGetProcAddress(libshlwapi, "SHGetViewStatePropertyBag")
	sHIsChildOrSelf = doGetProcAddress(libshlwapi, "SHIsChildOrSelf")
	sHIsLowMemoryMachine = doGetProcAddress(libshlwapi, "SHIsLowMemoryMachine")
	sHLoadIndirectString = doGetProcAddress(libshlwapi, "SHLoadIndirectString")
	sHLockShared = doGetProcAddress(libshlwapi, "SHLockShared")
	sHMessageBoxCheck = doGetProcAddress(libshlwapi, "SHMessageBoxCheckW")
	sHQueryInfoKey = doGetProcAddress(libshlwapi, "SHQueryInfoKeyW")
	sHQueryValueEx = doGetProcAddress(libshlwapi, "SHQueryValueExW")
	sHRegCloseUSKey = doGetProcAddress(libshlwapi, "SHRegCloseUSKey")
	sHRegCreateUSKey = doGetProcAddress(libshlwapi, "SHRegCreateUSKeyW")
	sHRegDeleteEmptyUSKey = doGetProcAddress(libshlwapi, "SHRegDeleteEmptyUSKeyW")
	sHRegDeleteUSValue = doGetProcAddress(libshlwapi, "SHRegDeleteUSValueW")
	sHRegDuplicateHKey = doGetProcAddress(libshlwapi, "SHRegDuplicateHKey")
	sHRegEnumUSKey = doGetProcAddress(libshlwapi, "SHRegEnumUSKeyW")
	sHRegEnumUSValue = doGetProcAddress(libshlwapi, "SHRegEnumUSValueW")
	sHRegGetBoolUSValue = doGetProcAddress(libshlwapi, "SHRegGetBoolUSValueW")
	sHRegGetIntW = doGetProcAddress(libshlwapi, "SHRegGetIntW")
	sHRegGetPath = doGetProcAddress(libshlwapi, "SHRegGetPathW")
	sHRegGetUSValue = doGetProcAddress(libshlwapi, "SHRegGetUSValueW")
	sHRegOpenUSKey = doGetProcAddress(libshlwapi, "SHRegOpenUSKeyW")
	sHRegQueryInfoUSKey = doGetProcAddress(libshlwapi, "SHRegQueryInfoUSKeyW")
	sHRegQueryUSValue = doGetProcAddress(libshlwapi, "SHRegQueryUSValueW")
	sHRegSetPath = doGetProcAddress(libshlwapi, "SHRegSetPathW")
	sHRegSetUSValue = doGetProcAddress(libshlwapi, "SHRegSetUSValueW")
	sHRegWriteUSValue = doGetProcAddress(libshlwapi, "SHRegWriteUSValueW")
	sHRegisterValidateTemplate = doGetProcAddress(libshlwapi, "SHRegisterValidateTemplate")
	sHReleaseThreadRef = doGetProcAddress(libshlwapi, "SHReleaseThreadRef")
	sHSendMessageBroadcast = doGetProcAddress(libshlwapi, "SHSendMessageBroadcastW")
	sHSetThreadRef = doGetProcAddress(libshlwapi, "SHSetThreadRef")
	sHSetValue = doGetProcAddress(libshlwapi, "SHSetValueW")
	sHSkipJunction = doGetProcAddress(libshlwapi, "SHSkipJunction")
	sHStrDup = doGetProcAddress(libshlwapi, "SHStrDupW")
	sHStripMneumonic = doGetProcAddress(libshlwapi, "SHStripMneumonicW")
	sHUnicodeToAnsi = doGetProcAddress(libshlwapi, "SHUnicodeToAnsi")
	sHUnicodeToUnicode = doGetProcAddress(libshlwapi, "SHUnicodeToUnicode")
	sHUnlockShared = doGetProcAddress(libshlwapi, "SHUnlockShared")
	strCSpnI = doGetProcAddress(libshlwapi, "StrCSpnIW")
	strCSpn = doGetProcAddress(libshlwapi, "StrCSpnW")
	strCatBuff = doGetProcAddress(libshlwapi, "StrCatBuffW")
	strCatChainW = doGetProcAddress(libshlwapi, "StrCatChainW")
	strCatW = doGetProcAddress(libshlwapi, "StrCatW")
	strChrI = doGetProcAddress(libshlwapi, "StrChrIW")
	strChrNW = doGetProcAddress(libshlwapi, "StrChrNW")
	strChr = doGetProcAddress(libshlwapi, "StrChrW")
	strCmpC = doGetProcAddress(libshlwapi, "StrCmpCW")
	strCmpIC = doGetProcAddress(libshlwapi, "StrCmpICW")
	strCmpIW = doGetProcAddress(libshlwapi, "StrCmpIW")
	strCmpLogicalW = doGetProcAddress(libshlwapi, "StrCmpLogicalW")
	strCmpNC = doGetProcAddress(libshlwapi, "StrCmpNCW")
	strCmpNIC = doGetProcAddress(libshlwapi, "StrCmpNICW")
	strCmpNI = doGetProcAddress(libshlwapi, "StrCmpNIW")
	strCmpN = doGetProcAddress(libshlwapi, "StrCmpNW")
	strCmpW = doGetProcAddress(libshlwapi, "StrCmpW")
	strCpyNW = doGetProcAddress(libshlwapi, "StrCpyNW")
	strCpyW = doGetProcAddress(libshlwapi, "StrCpyW")
	strDup = doGetProcAddress(libshlwapi, "StrDupW")
	strFormatByteSize64A = doGetProcAddress(libshlwapi, "StrFormatByteSize64A")
	strFormatByteSize = doGetProcAddress(libshlwapi, "StrFormatByteSizeW")
	strFormatKBSize = doGetProcAddress(libshlwapi, "StrFormatKBSizeW")
	strFromTimeInterval = doGetProcAddress(libshlwapi, "StrFromTimeIntervalW")
	strIsIntlEqual = doGetProcAddress(libshlwapi, "StrIsIntlEqualW")
	strNCat = doGetProcAddress(libshlwapi, "StrNCatW")
	strPBrk = doGetProcAddress(libshlwapi, "StrPBrkW")
	strRChrI = doGetProcAddress(libshlwapi, "StrRChrIW")
	strRChr = doGetProcAddress(libshlwapi, "StrRChrW")
	strRStrI = doGetProcAddress(libshlwapi, "StrRStrIW")
	strRetToBSTR = doGetProcAddress(libshlwapi, "StrRetToBSTR")
	strRetToBuf = doGetProcAddress(libshlwapi, "StrRetToBufW")
	strRetToStr = doGetProcAddress(libshlwapi, "StrRetToStrW")
	strSpn = doGetProcAddress(libshlwapi, "StrSpnW")
	strStrI = doGetProcAddress(libshlwapi, "StrStrIW")
	strStrNIW = doGetProcAddress(libshlwapi, "StrStrNIW")
	strStrNW = doGetProcAddress(libshlwapi, "StrStrNW")
	strStr = doGetProcAddress(libshlwapi, "StrStrW")
	strToInt64Ex = doGetProcAddress(libshlwapi, "StrToInt64ExW")
	strToIntEx = doGetProcAddress(libshlwapi, "StrToIntExW")
	strToInt = doGetProcAddress(libshlwapi, "StrToIntW")
	strTrim = doGetProcAddress(libshlwapi, "StrTrimW")
	urlApplyScheme = doGetProcAddress(libshlwapi, "UrlApplySchemeW")
	urlCanonicalize = doGetProcAddress(libshlwapi, "UrlCanonicalizeW")
	urlCombine = doGetProcAddress(libshlwapi, "UrlCombineW")
	urlCompare = doGetProcAddress(libshlwapi, "UrlCompareW")
	urlCreateFromPath = doGetProcAddress(libshlwapi, "UrlCreateFromPathW")
	urlEscape = doGetProcAddress(libshlwapi, "UrlEscapeW")
	urlFixupW = doGetProcAddress(libshlwapi, "UrlFixupW")
	urlGetLocation = doGetProcAddress(libshlwapi, "UrlGetLocationW")
	urlGetPart = doGetProcAddress(libshlwapi, "UrlGetPartW")
	urlHash = doGetProcAddress(libshlwapi, "UrlHashW")
	urlIsNoHistory = doGetProcAddress(libshlwapi, "UrlIsNoHistoryW")
	urlIsOpaque = doGetProcAddress(libshlwapi, "UrlIsOpaqueW")
	urlIs = doGetProcAddress(libshlwapi, "UrlIsW")
	urlUnescape = doGetProcAddress(libshlwapi, "UrlUnescapeW")
	whichPlatform = doGetProcAddress(libshlwapi, "WhichPlatform")
}

func AssocCreate(clsid CLSID, refiid REFIID, lpInterface uintptr) HRESULT {
	ret1 := syscall6(assocCreate, 6,
		uintptr(clsid.Data1),
		uintptr((uint32(clsid.Data2)<<16)|uint32(clsid.Data3)),
		uintptr((uint32(clsid.Data4[0])<<24)|(uint32(clsid.Data4[1]<<16))|(uint32(clsid.Data4[2]<<8))|uint32(clsid.Data4[3])),
		uintptr((uint32(clsid.Data4[4])<<24)|(uint32(clsid.Data4[5]<<16))|(uint32(clsid.Data4[6]<<8))|uint32(clsid.Data4[7])),
		uintptr(unsafe.Pointer(refiid)),
		lpInterface)
	return HRESULT(ret1)
}

func AssocGetPerceivedType(lpszExt string, lpType *PERCEIVED, lpFlag *int32, lppszType *LPWSTR) HRESULT {
	lpszExtStr := unicode16FromString(lpszExt)
	ret1 := syscall6(assocGetPerceivedType, 4,
		uintptr(unsafe.Pointer(&lpszExtStr[0])),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpFlag)),
		uintptr(unsafe.Pointer(lppszType)),
		0,
		0)
	return HRESULT(ret1)
}

func AssocIsDangerous(lpszAssoc string) bool {
	lpszAssocStr := unicode16FromString(lpszAssoc)
	ret1 := syscall3(assocIsDangerous, 1,
		uintptr(unsafe.Pointer(&lpszAssocStr[0])),
		0,
		0)
	return ret1 != 0
}

func AssocQueryKey(cfFlags ASSOCF, assockey ASSOCKEY, pszAssoc string, pszExtra string, phkeyOut *HKEY) HRESULT {
	pszAssocStr := unicode16FromString(pszAssoc)
	pszExtraStr := unicode16FromString(pszExtra)
	ret1 := syscall6(assocQueryKey, 5,
		uintptr(cfFlags),
		uintptr(assockey),
		uintptr(unsafe.Pointer(&pszAssocStr[0])),
		uintptr(unsafe.Pointer(&pszExtraStr[0])),
		uintptr(unsafe.Pointer(phkeyOut)),
		0)
	return HRESULT(ret1)
}

func AssocQueryStringByKey(cfFlags ASSOCF, str ASSOCSTR, hkAssoc HKEY, pszExtra string, pszOut LPWSTR, pcchOut *uint32) HRESULT {
	pszExtraStr := unicode16FromString(pszExtra)
	ret1 := syscall6(assocQueryStringByKey, 6,
		uintptr(cfFlags),
		uintptr(str),
		uintptr(hkAssoc),
		uintptr(unsafe.Pointer(&pszExtraStr[0])),
		uintptr(unsafe.Pointer(pszOut)),
		uintptr(unsafe.Pointer(pcchOut)))
	return HRESULT(ret1)
}

func AssocQueryString(cfFlags ASSOCF, str ASSOCSTR, pszAssoc string, pszExtra string, pszOut LPWSTR, pcchOut *uint32) HRESULT {
	pszAssocStr := unicode16FromString(pszAssoc)
	pszExtraStr := unicode16FromString(pszExtra)
	ret1 := syscall6(assocQueryString, 6,
		uintptr(cfFlags),
		uintptr(str),
		uintptr(unsafe.Pointer(&pszAssocStr[0])),
		uintptr(unsafe.Pointer(&pszExtraStr[0])),
		uintptr(unsafe.Pointer(pszOut)),
		uintptr(unsafe.Pointer(pcchOut)))
	return HRESULT(ret1)
}

func ChrCmpI(ch1 WCHAR, ch2 WCHAR) bool {
	ret1 := syscall3(chrCmpI, 2,
		uintptr(ch1),
		uintptr(ch2),
		0)
	return ret1 != 0
}

func ColorAdjustLuma(cRGB COLORREF, dwLuma int32, bUnknown bool) COLORREF {
	ret1 := syscall3(colorAdjustLuma, 3,
		uintptr(cRGB),
		uintptr(dwLuma),
		getUintptrFromBool(bUnknown))
	return COLORREF(ret1)
}

func ColorHLSToRGB(wHue uint16, wLuminosity uint16, wSaturation uint16) COLORREF {
	ret1 := syscall3(colorHLSToRGB, 3,
		uintptr(wHue),
		uintptr(wLuminosity),
		uintptr(wSaturation))
	return COLORREF(ret1)
}

func ColorRGBToHLS(cRGB COLORREF, pwHue *uint16, pwLuminance *uint16, pwSaturation *uint16) {
	syscall6(colorRGBToHLS, 4,
		uintptr(cRGB),
		uintptr(unsafe.Pointer(pwHue)),
		uintptr(unsafe.Pointer(pwLuminance)),
		uintptr(unsafe.Pointer(pwSaturation)),
		0,
		0)
}

func ConnectToConnectionPoint(lpUnkSink *IUnknown, riid REFIID, fConnect bool, lpUnknown *IUnknown, lpCookie *uint32, lppCP **IConnectionPoint) HRESULT {
	ret1 := syscall6(connectToConnectionPoint, 6,
		uintptr(unsafe.Pointer(lpUnkSink)),
		uintptr(unsafe.Pointer(riid)),
		getUintptrFromBool(fConnect),
		uintptr(unsafe.Pointer(lpUnknown)),
		uintptr(unsafe.Pointer(lpCookie)),
		uintptr(unsafe.Pointer(lppCP)))
	return HRESULT(ret1)
}

func GetAcceptLanguages(langbuf LPWSTR, buflen *uint32) HRESULT {
	ret1 := syscall3(getAcceptLanguages, 2,
		uintptr(unsafe.Pointer(langbuf)),
		uintptr(unsafe.Pointer(buflen)),
		0)
	return HRESULT(ret1)
}

func GetMenuPosFromID(hMenu HMENU, wID uint32) int32 {
	ret1 := syscall3(getMenuPosFromID, 2,
		uintptr(hMenu),
		uintptr(wID),
		0)
	return int32(ret1)
}

func HashData(lpSrc /*const*/ *byte, nSrcLen uint32, lpDest *byte, nDestLen uint32) HRESULT {
	ret1 := syscall6(hashData, 4,
		uintptr(unsafe.Pointer(lpSrc)),
		uintptr(nSrcLen),
		uintptr(unsafe.Pointer(lpDest)),
		uintptr(nDestLen),
		0,
		0)
	return HRESULT(ret1)
}

func IStream_Reset(lpStream *IStream) HRESULT {
	ret1 := syscall3(iStream_Reset, 1,
		uintptr(unsafe.Pointer(lpStream)),
		0,
		0)
	return HRESULT(ret1)
}

func IStream_Size(lpStream *IStream, lpulSize *ULARGE_INTEGER) HRESULT {
	ret1 := syscall3(iStream_Size, 2,
		uintptr(unsafe.Pointer(lpStream)),
		uintptr(unsafe.Pointer(lpulSize)),
		0)
	return HRESULT(ret1)
}

func IUnknown_AtomicRelease(lpUnknown **IUnknown) {
	syscall3(iUnknown_AtomicRelease, 1,
		uintptr(unsafe.Pointer(lpUnknown)),
		0,
		0)
}

func IUnknown_GetSite(lpUnknown LPUNKNOWN, iid REFIID, lppSite *PVOID) HRESULT {
	ret1 := syscall3(iUnknown_GetSite, 3,
		uintptr(unsafe.Pointer(lpUnknown)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(lppSite)))
	return HRESULT(ret1)
}

func IUnknown_GetWindow(lpUnknown *IUnknown, lphWnd *HWND) HRESULT {
	ret1 := syscall3(iUnknown_GetWindow, 2,
		uintptr(unsafe.Pointer(lpUnknown)),
		uintptr(unsafe.Pointer(lphWnd)),
		0)
	return HRESULT(ret1)
}

func IUnknown_QueryService(unnamed0 *IUnknown, unnamed1 REFGUID, unnamed2 REFIID, unnamed3 *LPVOID) HRESULT {
	ret1 := syscall6(iUnknown_QueryService, 4,
		uintptr(unsafe.Pointer(unnamed0)),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(unsafe.Pointer(unnamed2)),
		uintptr(unsafe.Pointer(unnamed3)),
		0,
		0)
	return HRESULT(ret1)
}

func IUnknown_Set(lppDest **IUnknown, lpUnknown *IUnknown) {
	syscall3(iUnknown_Set, 2,
		uintptr(unsafe.Pointer(lppDest)),
		uintptr(unsafe.Pointer(lpUnknown)),
		0)
}

func IUnknown_SetSite(obj *IUnknown, site *IUnknown) HRESULT {
	ret1 := syscall3(iUnknown_SetSite, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(site)),
		0)
	return HRESULT(ret1)
}

func IntlStrEqWorker(bCase bool, lpszStr string, lpszComp string, iLen int32) bool {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall6(intlStrEqWorker, 4,
		getUintptrFromBool(bCase),
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		uintptr(iLen),
		0,
		0)
	return ret1 != 0
}

func IsCharSpace(wc WCHAR) bool {
	ret1 := syscall3(isCharSpace, 1,
		uintptr(wc),
		0,
		0)
	return ret1 != 0
}

func IsInternetESCEnabled() bool {
	ret1 := syscall3(isInternetESCEnabled, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func IsOS(feature uint32) bool {
	ret1 := syscall3(isOS, 1,
		uintptr(feature),
		0,
		0)
	return ret1 != 0
}

func MLFreeLibrary(hModule HMODULE) bool {
	ret1 := syscall3(mLFreeLibrary, 1,
		uintptr(hModule),
		0,
		0)
	return ret1 != 0
}

func MLLoadLibrary(new_mod string, inst_hwnd HMODULE, dwCrossCodePage uint32) HMODULE {
	new_modStr := unicode16FromString(new_mod)
	ret1 := syscall3(mLLoadLibrary, 3,
		uintptr(unsafe.Pointer(&new_modStr[0])),
		uintptr(inst_hwnd),
		uintptr(dwCrossCodePage))
	return HMODULE(ret1)
}

func ParseURL(x string, y *PARSEDURL) HRESULT {
	xStr := unicode16FromString(x)
	ret1 := syscall3(parseURL, 2,
		uintptr(unsafe.Pointer(&xStr[0])),
		uintptr(unsafe.Pointer(y)),
		0)
	return HRESULT(ret1)
}

func PathAddBackslash(lpszPath LPWSTR) LPWSTR {
	ret1 := syscall3(pathAddBackslash, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathAddExtension(lpszPath LPWSTR, lpszExtension string) bool {
	lpszExtensionStr := unicode16FromString(lpszExtension)
	ret1 := syscall3(pathAddExtension, 2,
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(unsafe.Pointer(&lpszExtensionStr[0])),
		0)
	return ret1 != 0
}

func PathAppend(lpszPath LPWSTR, lpszAppend string) bool {
	lpszAppendStr := unicode16FromString(lpszAppend)
	ret1 := syscall3(pathAppend, 2,
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(unsafe.Pointer(&lpszAppendStr[0])),
		0)
	return ret1 != 0
}

func PathBuildRoot(lpszPath LPWSTR, drive int32) LPWSTR {
	ret1 := syscall3(pathBuildRoot, 2,
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(drive),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathCanonicalize(lpszBuf LPWSTR, lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathCanonicalize, 2,
		uintptr(unsafe.Pointer(lpszBuf)),
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0)
	return ret1 != 0
}

func PathCombine(lpszDest LPWSTR, lpszDir string, lpszFile string) LPWSTR {
	lpszDirStr := unicode16FromString(lpszDir)
	lpszFileStr := unicode16FromString(lpszFile)
	ret1 := syscall3(pathCombine, 3,
		uintptr(unsafe.Pointer(lpszDest)),
		uintptr(unsafe.Pointer(&lpszDirStr[0])),
		uintptr(unsafe.Pointer(&lpszFileStr[0])))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathCommonPrefix(lpszFile1 string, lpszFile2 string, achPath LPWSTR) int32 {
	lpszFile1Str := unicode16FromString(lpszFile1)
	lpszFile2Str := unicode16FromString(lpszFile2)
	ret1 := syscall3(pathCommonPrefix, 3,
		uintptr(unsafe.Pointer(&lpszFile1Str[0])),
		uintptr(unsafe.Pointer(&lpszFile2Str[0])),
		uintptr(unsafe.Pointer(achPath)))
	return int32(ret1)
}

func PathCompactPathEx(lpszDest LPWSTR, lpszPath string, cchMax uint32, dwFlags uint32) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall6(pathCompactPathEx, 4,
		uintptr(unsafe.Pointer(lpszDest)),
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(cchMax),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func PathCompactPath(hDC HDC, lpszPath LPWSTR, dx uint32) bool {
	ret1 := syscall3(pathCompactPath, 3,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(dx))
	return ret1 != 0
}

func PathCreateFromUrlAlloc(pszUrl string, pszPath *LPWSTR, dwReserved uint32) HRESULT {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall3(pathCreateFromUrlAlloc, 3,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		uintptr(unsafe.Pointer(pszPath)),
		uintptr(dwReserved))
	return HRESULT(ret1)
}

func PathCreateFromUrl(pszUrl string, pszPath LPWSTR, pcchPath *uint32, dwReserved uint32) HRESULT {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall6(pathCreateFromUrl, 4,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		uintptr(unsafe.Pointer(pszPath)),
		uintptr(unsafe.Pointer(pcchPath)),
		uintptr(dwReserved),
		0,
		0)
	return HRESULT(ret1)
}

func PathFileExists(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathFileExists, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathFindExtension(lpszPath string) LPWSTR {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathFindExtension, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathFindFileName(lpszPath string) LPWSTR {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathFindFileName, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathFindNextComponent(lpszPath string) LPWSTR {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathFindNextComponent, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathFindOnPath(lpszFile LPWSTR, lppszOtherDirs *LPCWSTR) bool {
	ret1 := syscall3(pathFindOnPath, 2,
		uintptr(unsafe.Pointer(lpszFile)),
		uintptr(unsafe.Pointer(lppszOtherDirs)),
		0)
	return ret1 != 0
}

func PathFindSuffixArray(lpszSuffix string, lppszArray *LPCWSTR, dwCount int32) string {
	lpszSuffixStr := unicode16FromString(lpszSuffix)
	ret1 := syscall3(pathFindSuffixArray, 3,
		uintptr(unsafe.Pointer(&lpszSuffixStr[0])),
		uintptr(unsafe.Pointer(lppszArray)),
		uintptr(dwCount))
	return stringFromUnicode16((*uint16)(unsafe.Pointer(ret1)))
}

func PathGetArgs(lpszPath string) LPWSTR {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathGetArgs, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathGetCharType(ch WCHAR) uint32 {
	ret1 := syscall3(pathGetCharType, 1,
		uintptr(ch),
		0,
		0)
	return uint32(ret1)
}

func PathGetDriveNumber(lpszPath string) int32 {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathGetDriveNumber, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return int32(ret1)
}

func PathIsContentType(lpszPath string, lpszContentType string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	lpszContentTypeStr := unicode16FromString(lpszContentType)
	ret1 := syscall3(pathIsContentType, 2,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(unsafe.Pointer(&lpszContentTypeStr[0])),
		0)
	return ret1 != 0
}

func PathIsDirectoryEmpty(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsDirectoryEmpty, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsDirectory(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsDirectory, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsFileSpec(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsFileSpec, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsLFNFileSpec(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsLFNFileSpec, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsNetworkPath(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsNetworkPath, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsPrefix(lpszPrefix string, lpszPath string) bool {
	lpszPrefixStr := unicode16FromString(lpszPrefix)
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsPrefix, 2,
		uintptr(unsafe.Pointer(&lpszPrefixStr[0])),
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0)
	return ret1 != 0
}

func PathIsRelative(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsRelative, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsRoot(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsRoot, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsSameRoot(lpszPath1 string, lpszPath2 string) bool {
	lpszPath1Str := unicode16FromString(lpszPath1)
	lpszPath2Str := unicode16FromString(lpszPath2)
	ret1 := syscall3(pathIsSameRoot, 2,
		uintptr(unsafe.Pointer(&lpszPath1Str[0])),
		uintptr(unsafe.Pointer(&lpszPath2Str[0])),
		0)
	return ret1 != 0
}

func PathIsSystemFolder(lpszPath string, dwAttrib uint32) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsSystemFolder, 2,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(dwAttrib),
		0)
	return ret1 != 0
}

func PathIsUNCServerShare(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsUNCServerShare, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsUNCServer(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsUNCServer, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsUNC(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathIsUNC, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathIsURL(lpstrPath string) bool {
	lpstrPathStr := unicode16FromString(lpstrPath)
	ret1 := syscall3(pathIsURL, 1,
		uintptr(unsafe.Pointer(&lpstrPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathMakePretty(lpszPath LPWSTR) bool {
	ret1 := syscall3(pathMakePretty, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
	return ret1 != 0
}

func PathMakeSystemFolder(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathMakeSystemFolder, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathMatchSpec(lpszPath string, lpszMask string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	lpszMaskStr := unicode16FromString(lpszMask)
	ret1 := syscall3(pathMatchSpec, 2,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(unsafe.Pointer(&lpszMaskStr[0])),
		0)
	return ret1 != 0
}

func PathParseIconLocation(lpszPath LPWSTR) int32 {
	ret1 := syscall3(pathParseIconLocation, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
	return int32(ret1)
}

func PathQuoteSpaces(lpszPath LPWSTR) {
	syscall3(pathQuoteSpaces, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func PathRelativePathTo(lpszPath LPWSTR, lpszFrom string, dwAttrFrom uint32, lpszTo string, dwAttrTo uint32) bool {
	lpszFromStr := unicode16FromString(lpszFrom)
	lpszToStr := unicode16FromString(lpszTo)
	ret1 := syscall6(pathRelativePathTo, 5,
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(unsafe.Pointer(&lpszFromStr[0])),
		uintptr(dwAttrFrom),
		uintptr(unsafe.Pointer(&lpszToStr[0])),
		uintptr(dwAttrTo),
		0)
	return ret1 != 0
}

func PathRemoveArgs(lpszPath LPWSTR) {
	syscall3(pathRemoveArgs, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func PathRemoveBackslash(lpszPath LPWSTR) LPWSTR {
	ret1 := syscall3(pathRemoveBackslash, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathRemoveBlanks(lpszPath LPWSTR) {
	syscall3(pathRemoveBlanks, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func PathRemoveExtension(lpszPath LPWSTR) {
	syscall3(pathRemoveExtension, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func PathRemoveFileSpec(lpszPath LPWSTR) bool {
	ret1 := syscall3(pathRemoveFileSpec, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
	return ret1 != 0
}

func PathRenameExtension(lpszPath LPWSTR, lpszExt string) bool {
	lpszExtStr := unicode16FromString(lpszExt)
	ret1 := syscall3(pathRenameExtension, 2,
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(unsafe.Pointer(&lpszExtStr[0])),
		0)
	return ret1 != 0
}

func PathSearchAndQualify(lpszPath string, lpszBuf LPWSTR, cchBuf uint32) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathSearchAndQualify, 3,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(unsafe.Pointer(lpszBuf)),
		uintptr(cchBuf))
	return ret1 != 0
}

func PathSetDlgItemPath(hDlg HWND, id int32, lpszPath string) {
	lpszPathStr := unicode16FromString(lpszPath)
	syscall3(pathSetDlgItemPath, 3,
		uintptr(hDlg),
		uintptr(id),
		uintptr(unsafe.Pointer(&lpszPathStr[0])))
}

func PathSkipRoot(lpszPath string) LPWSTR {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathSkipRoot, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func PathStripPath(lpszPath LPWSTR) {
	syscall3(pathStripPath, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func PathStripToRoot(lpszPath LPWSTR) bool {
	ret1 := syscall3(pathStripToRoot, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
	return ret1 != 0
}

func PathUnExpandEnvStrings(path string, buffer LPWSTR, buf_len uint32) bool {
	pathStr := unicode16FromString(path)
	ret1 := syscall3(pathUnExpandEnvStrings, 3,
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(buf_len))
	return ret1 != 0
}

func PathUndecorate(lpszPath LPWSTR) {
	syscall3(pathUndecorate, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func PathUnmakeSystemFolder(lpszPath string) bool {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(pathUnmakeSystemFolder, 1,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func PathUnquoteSpaces(lpszPath LPWSTR) {
	syscall3(pathUnquoteSpaces, 1,
		uintptr(unsafe.Pointer(lpszPath)),
		0,
		0)
}

func QISearch(base uintptr, table /*const*/ *QITAB, riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall6(qISearch, 4,
		base,
		uintptr(unsafe.Pointer(table)),
		uintptr(unsafe.Pointer(riid)),
		ppv,
		0,
		0)
	return HRESULT(ret1)
}

func SHAllocShared(lpvData LPVOID, dwSize uint32, dwProcId uint32) HANDLE {
	ret1 := syscall3(sHAllocShared, 3,
		uintptr(unsafe.Pointer(lpvData)),
		uintptr(dwSize),
		uintptr(dwProcId))
	return HANDLE(ret1)
}

func SHAnsiToAnsi(lpszSrc /*const*/ LPCSTR, lpszDst LPSTR, iLen int32) uint32 {
	ret1 := syscall3(sHAnsiToAnsi, 3,
		uintptr(unsafe.Pointer(lpszSrc)),
		uintptr(unsafe.Pointer(lpszDst)),
		uintptr(iLen))
	return uint32(ret1)
}

func SHAnsiToUnicode(lpSrcStr /*const*/ LPCSTR, lpDstStr LPWSTR, iLen int32) uint32 {
	ret1 := syscall3(sHAnsiToUnicode, 3,
		uintptr(unsafe.Pointer(lpSrcStr)),
		uintptr(unsafe.Pointer(lpDstStr)),
		uintptr(iLen))
	return uint32(ret1)
}

func SHAutoComplete(hwndEdit HWND, dwFlags uint32) HRESULT {
	ret1 := syscall3(sHAutoComplete, 2,
		uintptr(hwndEdit),
		uintptr(dwFlags),
		0)
	return HRESULT(ret1)
}

func SHCopyKey(hKeySrc HKEY, lpszSrcSubKey string, hKeyDst HKEY, dwReserved uint32) uint32 {
	lpszSrcSubKeyStr := unicode16FromString(lpszSrcSubKey)
	ret1 := syscall6(sHCopyKey, 4,
		uintptr(hKeySrc),
		uintptr(unsafe.Pointer(&lpszSrcSubKeyStr[0])),
		uintptr(hKeyDst),
		uintptr(dwReserved),
		0,
		0)
	return uint32(ret1)
}

func SHCreateShellPalette(hdc HDC) HPALETTE {
	ret1 := syscall3(sHCreateShellPalette, 1,
		uintptr(hdc),
		0,
		0)
	return HPALETTE(ret1)
}

func SHCreateStreamOnFileEx(lpszPath string, dwMode uint32, dwAttributes uint32, bCreate bool, lpTemplate *IStream, lppStream **IStream) HRESULT {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall6(sHCreateStreamOnFileEx, 6,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(dwMode),
		uintptr(dwAttributes),
		getUintptrFromBool(bCreate),
		uintptr(unsafe.Pointer(lpTemplate)),
		uintptr(unsafe.Pointer(lppStream)))
	return HRESULT(ret1)
}

func SHCreateStreamOnFile(lpszPath string, dwMode uint32, lppStream **IStream) HRESULT {
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall3(sHCreateStreamOnFile, 3,
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(dwMode),
		uintptr(unsafe.Pointer(lppStream)))
	return HRESULT(ret1)
}

func SHCreateStreamWrapper(lpbData *byte, dwDataLen uint32, dwReserved uint32, lppStream **IStream) HRESULT {
	ret1 := syscall6(sHCreateStreamWrapper, 4,
		uintptr(unsafe.Pointer(lpbData)),
		uintptr(dwDataLen),
		uintptr(dwReserved),
		uintptr(unsafe.Pointer(lppStream)),
		0,
		0)
	return HRESULT(ret1)
}

func SHCreateThread(pfnThreadProc THREAD_START_ROUTINE, pData uintptr, dwFlags uint32, pfnCallback THREAD_START_ROUTINE) bool {
	pfnThreadProcCallback := syscall.NewCallback(func(lpThreadParameterRawArg LPVOID) uintptr {
		ret := pfnThreadProc(lpThreadParameterRawArg)
		return uintptr(ret)
	})
	pfnCallbackCallback := syscall.NewCallback(func(lpThreadParameterRawArg LPVOID) uintptr {
		ret := pfnCallback(lpThreadParameterRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(sHCreateThread, 4,
		pfnThreadProcCallback,
		pData,
		uintptr(dwFlags),
		pfnCallbackCallback,
		0,
		0)
	return ret1 != 0
}

func SHCreateThreadRef(lprefcount *LONG, lppUnknown **IUnknown) HRESULT {
	ret1 := syscall3(sHCreateThreadRef, 2,
		uintptr(unsafe.Pointer(lprefcount)),
		uintptr(unsafe.Pointer(lppUnknown)),
		0)
	return HRESULT(ret1)
}

func SHDeleteEmptyKey(hKey HKEY, lpszSubKey string) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	ret1 := syscall3(sHDeleteEmptyKey, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		0)
	return uint32(ret1)
}

func SHDeleteKey(hKey HKEY, lpszSubKey string) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	ret1 := syscall3(sHDeleteKey, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		0)
	return uint32(ret1)
}

func SHDeleteOrphanKey(hKey HKEY, lpszSubKey string) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	ret1 := syscall3(sHDeleteOrphanKey, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		0)
	return uint32(ret1)
}

func SHDeleteValue(hKey HKEY, lpszSubKey string, lpszValue string) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	lpszValueStr := unicode16FromString(lpszValue)
	ret1 := syscall3(sHDeleteValue, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpszValueStr[0])))
	return uint32(ret1)
}

func SHEnumKeyEx(hKey HKEY, dwIndex uint32, lpszSubKey LPWSTR, pwLen *uint32) LONG {
	ret1 := syscall6(sHEnumKeyEx, 4,
		uintptr(hKey),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(lpszSubKey)),
		uintptr(unsafe.Pointer(pwLen)),
		0,
		0)
	return LONG(ret1)
}

func SHEnumValue(hKey HKEY, dwIndex uint32, lpszValue LPWSTR, pwLen *uint32, pwType *uint32, pvData LPVOID, pcbData *uint32) LONG {
	ret1 := syscall9(sHEnumValue, 7,
		uintptr(hKey),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(lpszValue)),
		uintptr(unsafe.Pointer(pwLen)),
		uintptr(unsafe.Pointer(pwType)),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(unsafe.Pointer(pcbData)),
		0,
		0)
	return LONG(ret1)
}

func SHFormatDateTime(fileTime /*const*/ *FILETIME, flags *uint32, buf LPWSTR, size uint32) int32 {
	ret1 := syscall6(sHFormatDateTime, 4,
		uintptr(unsafe.Pointer(fileTime)),
		uintptr(unsafe.Pointer(flags)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(size),
		0,
		0)
	return int32(ret1)
}

func SHFreeShared(hShared HANDLE, dwProcId uint32) bool {
	ret1 := syscall3(sHFreeShared, 2,
		uintptr(hShared),
		uintptr(dwProcId),
		0)
	return ret1 != 0
}

func SHGetInverseCMAP(dest *uint32, dwSize uint32) HRESULT {
	ret1 := syscall3(sHGetInverseCMAP, 2,
		uintptr(unsafe.Pointer(dest)),
		uintptr(dwSize),
		0)
	return HRESULT(ret1)
}

func SHGetThreadRef(lppUnknown **IUnknown) HRESULT {
	ret1 := syscall3(sHGetThreadRef, 1,
		uintptr(unsafe.Pointer(lppUnknown)),
		0,
		0)
	return HRESULT(ret1)
}

func SHGetValue(hKey HKEY, lpszSubKey string, lpszValue string, pwType *uint32, pvData LPVOID, pcbData *uint32) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	lpszValueStr := unicode16FromString(lpszValue)
	ret1 := syscall6(sHGetValue, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpszValueStr[0])),
		uintptr(unsafe.Pointer(pwType)),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(unsafe.Pointer(pcbData)))
	return uint32(ret1)
}

func SHGetViewStatePropertyBag(pidl /*const*/ LPCITEMIDLIST, bag_name LPWSTR, flags uint32, riid REFIID, ppv uintptr) HRESULT {
	ret1 := syscall6(sHGetViewStatePropertyBag, 5,
		uintptr(unsafe.Pointer(pidl)),
		uintptr(unsafe.Pointer(bag_name)),
		uintptr(flags),
		uintptr(unsafe.Pointer(riid)),
		ppv,
		0)
	return HRESULT(ret1)
}

func SHIsChildOrSelf(hParent HWND, hChild HWND) bool {
	ret1 := syscall3(sHIsChildOrSelf, 2,
		uintptr(hParent),
		uintptr(hChild),
		0)
	return ret1 != 0
}

func SHIsLowMemoryMachine(x uint32) bool {
	ret1 := syscall3(sHIsLowMemoryMachine, 1,
		uintptr(x),
		0,
		0)
	return ret1 != 0
}

func SHLoadIndirectString(src string, dst LPWSTR, dst_len uint32, reserved uintptr) HRESULT {
	srcStr := unicode16FromString(src)
	ret1 := syscall6(sHLoadIndirectString, 4,
		uintptr(unsafe.Pointer(&srcStr[0])),
		uintptr(unsafe.Pointer(dst)),
		uintptr(dst_len),
		reserved,
		0,
		0)
	return HRESULT(ret1)
}

func SHLockShared(hShared HANDLE, dwProcId uint32) LPVOID {
	ret1 := syscall3(sHLockShared, 2,
		uintptr(hShared),
		uintptr(dwProcId),
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func SHMessageBoxCheck(unnamed0 HWND, unnamed1 string, unnamed2 string, unnamed3 uint32, unnamed4 INT_PTR, unnamed5 string) INT_PTR {
	unnamed1Str := unicode16FromString(unnamed1)
	unnamed2Str := unicode16FromString(unnamed2)
	unnamed5Str := unicode16FromString(unnamed5)
	ret1 := syscall6(sHMessageBoxCheck, 6,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(&unnamed1Str[0])),
		uintptr(unsafe.Pointer(&unnamed2Str[0])),
		uintptr(unnamed3),
		uintptr(unsafe.Pointer(unnamed4)),
		uintptr(unsafe.Pointer(&unnamed5Str[0])))
	return (INT_PTR)(unsafe.Pointer(ret1))
}

func SHQueryInfoKey(hKey HKEY, pwSubKeys *uint32, pwSubKeyMax *uint32, pwValues *uint32, pwValueMax *uint32) LONG {
	ret1 := syscall6(sHQueryInfoKey, 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(pwSubKeys)),
		uintptr(unsafe.Pointer(pwSubKeyMax)),
		uintptr(unsafe.Pointer(pwValues)),
		uintptr(unsafe.Pointer(pwValueMax)),
		0)
	return LONG(ret1)
}

func SHQueryValueEx(hKey HKEY, lpszValue string, lpReserved *uint32, pwType *uint32, pvData LPVOID, pcbData *uint32) uint32 {
	lpszValueStr := unicode16FromString(lpszValue)
	ret1 := syscall6(sHQueryValueEx, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszValueStr[0])),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(pwType)),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(unsafe.Pointer(pcbData)))
	return uint32(ret1)
}

func SHRegCloseUSKey(hUSKey HUSKEY) LONG {
	ret1 := syscall3(sHRegCloseUSKey, 1,
		uintptr(hUSKey),
		0,
		0)
	return LONG(ret1)
}

func SHRegCreateUSKey(path string, samDesired REGSAM, relative_key HUSKEY, new_uskey PHUSKEY, flags uint32) LONG {
	pathStr := unicode16FromString(path)
	ret1 := syscall6(sHRegCreateUSKey, 5,
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(samDesired),
		uintptr(relative_key),
		uintptr(unsafe.Pointer(new_uskey)),
		uintptr(flags),
		0)
	return LONG(ret1)
}

func SHRegDeleteEmptyUSKey(hUSKey HUSKEY, pszValue string, delRegFlags SHREGDEL_FLAGS) LONG {
	pszValueStr := unicode16FromString(pszValue)
	ret1 := syscall3(sHRegDeleteEmptyUSKey, 3,
		uintptr(hUSKey),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		uintptr(delRegFlags))
	return LONG(ret1)
}

func SHRegDeleteUSValue(hUSKey HUSKEY, pszValue string, delRegFlags SHREGDEL_FLAGS) LONG {
	pszValueStr := unicode16FromString(pszValue)
	ret1 := syscall3(sHRegDeleteUSValue, 3,
		uintptr(hUSKey),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		uintptr(delRegFlags))
	return LONG(ret1)
}

func SHRegDuplicateHKey(hKey HKEY) HKEY {
	ret1 := syscall3(sHRegDuplicateHKey, 1,
		uintptr(hKey),
		0,
		0)
	return HKEY(ret1)
}

func SHRegEnumUSKey(hUSKey HUSKEY, dwIndex uint32, pszName LPWSTR, pcchValueNameLen *uint32, enumRegFlags SHREGENUM_FLAGS) LONG {
	ret1 := syscall6(sHRegEnumUSKey, 5,
		uintptr(hUSKey),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(pszName)),
		uintptr(unsafe.Pointer(pcchValueNameLen)),
		uintptr(enumRegFlags),
		0)
	return LONG(ret1)
}

func SHRegEnumUSValue(hUSKey HUSKEY, dwIndex uint32, pszValueName LPWSTR, pcchValueNameLen *uint32, pdwType *uint32, pvData LPVOID, pcbData *uint32, enumRegFlags SHREGENUM_FLAGS) LONG {
	ret1 := syscall9(sHRegEnumUSValue, 8,
		uintptr(hUSKey),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(pszValueName)),
		uintptr(unsafe.Pointer(pcchValueNameLen)),
		uintptr(unsafe.Pointer(pdwType)),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(unsafe.Pointer(pcbData)),
		uintptr(enumRegFlags),
		0)
	return LONG(ret1)
}

func SHRegGetBoolUSValue(pszSubKey string, pszValue string, fIgnoreHKCU bool, fDefault bool) bool {
	pszSubKeyStr := unicode16FromString(pszSubKey)
	pszValueStr := unicode16FromString(pszValue)
	ret1 := syscall6(sHRegGetBoolUSValue, 4,
		uintptr(unsafe.Pointer(&pszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		getUintptrFromBool(fIgnoreHKCU),
		getUintptrFromBool(fDefault),
		0,
		0)
	return ret1 != 0
}

func SHRegGetIntW(hKey HKEY, lpszValue string, iDefault int32) int32 {
	lpszValueStr := unicode16FromString(lpszValue)
	ret1 := syscall3(sHRegGetIntW, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszValueStr[0])),
		uintptr(iDefault))
	return int32(ret1)
}

func SHRegGetPath(hKey HKEY, lpszSubKey string, lpszValue string, lpszPath LPWSTR, dwFlags uint32) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	lpszValueStr := unicode16FromString(lpszValue)
	ret1 := syscall6(sHRegGetPath, 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpszValueStr[0])),
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(dwFlags),
		0)
	return uint32(ret1)
}

func SHRegGetUSValue(pSubKey string, pValue string, pwType *uint32, pvData LPVOID, pcbData *uint32, flagIgnoreHKCU bool, pDefaultData LPVOID, wDefaultDataSize uint32) LONG {
	pSubKeyStr := unicode16FromString(pSubKey)
	pValueStr := unicode16FromString(pValue)
	ret1 := syscall9(sHRegGetUSValue, 8,
		uintptr(unsafe.Pointer(&pSubKeyStr[0])),
		uintptr(unsafe.Pointer(&pValueStr[0])),
		uintptr(unsafe.Pointer(pwType)),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(unsafe.Pointer(pcbData)),
		getUintptrFromBool(flagIgnoreHKCU),
		uintptr(unsafe.Pointer(pDefaultData)),
		uintptr(wDefaultDataSize),
		0)
	return LONG(ret1)
}

func SHRegOpenUSKey(path string, accessType REGSAM, hRelativeUSKey HUSKEY, phNewUSKey PHUSKEY, fIgnoreHKCU bool) LONG {
	pathStr := unicode16FromString(path)
	ret1 := syscall6(sHRegOpenUSKey, 5,
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(accessType),
		uintptr(hRelativeUSKey),
		uintptr(unsafe.Pointer(phNewUSKey)),
		getUintptrFromBool(fIgnoreHKCU),
		0)
	return LONG(ret1)
}

func SHRegQueryInfoUSKey(hUSKey HUSKEY, pcSubKeys *uint32, pcchMaxSubKeyLen *uint32, pcValues *uint32, pcchMaxValueNameLen *uint32, enumRegFlags SHREGENUM_FLAGS) LONG {
	ret1 := syscall6(sHRegQueryInfoUSKey, 6,
		uintptr(hUSKey),
		uintptr(unsafe.Pointer(pcSubKeys)),
		uintptr(unsafe.Pointer(pcchMaxSubKeyLen)),
		uintptr(unsafe.Pointer(pcValues)),
		uintptr(unsafe.Pointer(pcchMaxValueNameLen)),
		uintptr(enumRegFlags))
	return LONG(ret1)
}

func SHRegQueryUSValue(hUSKey HUSKEY, pszValue string, pdwType *uint32, pvData LPVOID, pcbData *uint32, fIgnoreHKCU bool, pvDefaultData LPVOID, dwDefaultDataSize uint32) LONG {
	pszValueStr := unicode16FromString(pszValue)
	ret1 := syscall9(sHRegQueryUSValue, 8,
		uintptr(hUSKey),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		uintptr(unsafe.Pointer(pdwType)),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(unsafe.Pointer(pcbData)),
		getUintptrFromBool(fIgnoreHKCU),
		uintptr(unsafe.Pointer(pvDefaultData)),
		uintptr(dwDefaultDataSize),
		0)
	return LONG(ret1)
}

func SHRegSetPath(hKey HKEY, lpszSubKey string, lpszValue string, lpszPath string, dwFlags uint32) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	lpszValueStr := unicode16FromString(lpszValue)
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall6(sHRegSetPath, 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpszValueStr[0])),
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		uintptr(dwFlags),
		0)
	return uint32(ret1)
}

func SHRegSetUSValue(pszSubKey string, pszValue string, dwType uint32, pvData LPVOID, cbData uint32, dwFlags uint32) LONG {
	pszSubKeyStr := unicode16FromString(pszSubKey)
	pszValueStr := unicode16FromString(pszValue)
	ret1 := syscall6(sHRegSetUSValue, 6,
		uintptr(unsafe.Pointer(&pszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		uintptr(dwType),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(cbData),
		uintptr(dwFlags))
	return LONG(ret1)
}

func SHRegWriteUSValue(hUSKey HUSKEY, pszValue string, dwType uint32, pvData LPVOID, cbData uint32, dwFlags uint32) LONG {
	pszValueStr := unicode16FromString(pszValue)
	ret1 := syscall6(sHRegWriteUSValue, 6,
		uintptr(hUSKey),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		uintptr(dwType),
		uintptr(unsafe.Pointer(pvData)),
		uintptr(cbData),
		uintptr(dwFlags))
	return LONG(ret1)
}

func SHRegisterValidateTemplate(filename string, unknown bool) HRESULT {
	filenameStr := unicode16FromString(filename)
	ret1 := syscall3(sHRegisterValidateTemplate, 2,
		uintptr(unsafe.Pointer(&filenameStr[0])),
		getUintptrFromBool(unknown),
		0)
	return HRESULT(ret1)
}

func SHReleaseThreadRef() HRESULT {
	ret1 := syscall3(sHReleaseThreadRef, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func SHSendMessageBroadcast(uMsg uint32, wParam WPARAM, lParam LPARAM) uint32 {
	ret1 := syscall3(sHSendMessageBroadcast, 3,
		uintptr(uMsg),
		uintptr(wParam),
		uintptr(lParam))
	return uint32(ret1)
}

func SHSetThreadRef(lpUnknown *IUnknown) HRESULT {
	ret1 := syscall3(sHSetThreadRef, 1,
		uintptr(unsafe.Pointer(lpUnknown)),
		0,
		0)
	return HRESULT(ret1)
}

func SHSetValue(hKey HKEY, lpszSubKey string, lpszValue string, dwType uint32, pvData /*const*/ uintptr, cbData uint32) uint32 {
	lpszSubKeyStr := unicode16FromString(lpszSubKey)
	lpszValueStr := unicode16FromString(lpszValue)
	ret1 := syscall6(sHSetValue, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpszSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpszValueStr[0])),
		uintptr(dwType),
		pvData,
		uintptr(cbData))
	return uint32(ret1)
}

func SHSkipJunction(pbc *IBindCtx, pclsid /*const*/ *CLSID) bool {
	ret1 := syscall3(sHSkipJunction, 2,
		uintptr(unsafe.Pointer(pbc)),
		uintptr(unsafe.Pointer(pclsid)),
		0)
	return ret1 != 0
}

func SHStrDup(src string, dest *LPWSTR) HRESULT {
	srcStr := unicode16FromString(src)
	ret1 := syscall3(sHStrDup, 2,
		uintptr(unsafe.Pointer(&srcStr[0])),
		uintptr(unsafe.Pointer(dest)),
		0)
	return HRESULT(ret1)
}

func SHStripMneumonic(lpszStr string) WCHAR {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(sHStripMneumonic, 1,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		0,
		0)
	return WCHAR(ret1)
}

func SHUnicodeToAnsi(lpSrcStr string, lpDstStr LPSTR, iLen int32) int32 {
	lpSrcStrStr := unicode16FromString(lpSrcStr)
	ret1 := syscall3(sHUnicodeToAnsi, 3,
		uintptr(unsafe.Pointer(&lpSrcStrStr[0])),
		uintptr(unsafe.Pointer(lpDstStr)),
		uintptr(iLen))
	return int32(ret1)
}

func SHUnicodeToUnicode(lpszSrc string, lpszDst LPWSTR, iLen int32) uint32 {
	lpszSrcStr := unicode16FromString(lpszSrc)
	ret1 := syscall3(sHUnicodeToUnicode, 3,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(lpszDst)),
		uintptr(iLen))
	return uint32(ret1)
}

func SHUnlockShared(lpView LPVOID) bool {
	ret1 := syscall3(sHUnlockShared, 1,
		uintptr(unsafe.Pointer(lpView)),
		0,
		0)
	return ret1 != 0
}

func StrCSpnI(lpszStr string, lpszMatch string) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszMatchStr := unicode16FromString(lpszMatch)
	ret1 := syscall3(strCSpnI, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszMatchStr[0])),
		0)
	return int32(ret1)
}

func StrCSpn(lpszStr string, lpszMatch string) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszMatchStr := unicode16FromString(lpszMatch)
	ret1 := syscall3(strCSpn, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszMatchStr[0])),
		0)
	return int32(ret1)
}

func StrCatBuff(lpszStr LPWSTR, lpszCat string, cchMax int32) LPWSTR {
	lpszCatStr := unicode16FromString(lpszCat)
	ret1 := syscall3(strCatBuff, 3,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(unsafe.Pointer(&lpszCatStr[0])),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrCatChainW(lpszStr LPWSTR, cchMax uint32, ichAt uint32, lpszCat string) uint32 {
	lpszCatStr := unicode16FromString(lpszCat)
	ret1 := syscall6(strCatChainW, 4,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(cchMax),
		uintptr(ichAt),
		uintptr(unsafe.Pointer(&lpszCatStr[0])),
		0,
		0)
	return uint32(ret1)
}

func StrCatW(lpszStr LPWSTR, lpszSrc string) LPWSTR {
	lpszSrcStr := unicode16FromString(lpszSrc)
	ret1 := syscall3(strCatW, 2,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrChrI(lpszStr string, ch WCHAR) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(strChrI, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(ch),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrChrNW(lpszStr string, ch WCHAR, cchMax uint32) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(strChrNW, 3,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(ch),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrChr(lpszStr string, ch WCHAR) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(strChr, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(ch),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrCmpC(lpszSrc string, lpszCmp string) uint32 {
	lpszSrcStr := unicode16FromString(lpszSrc)
	lpszCmpStr := unicode16FromString(lpszCmp)
	ret1 := syscall3(strCmpC, 2,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(&lpszCmpStr[0])),
		0)
	return uint32(ret1)
}

func StrCmpIC(lpszSrc string, lpszCmp string) uint32 {
	lpszSrcStr := unicode16FromString(lpszSrc)
	lpszCmpStr := unicode16FromString(lpszCmp)
	ret1 := syscall3(strCmpIC, 2,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(&lpszCmpStr[0])),
		0)
	return uint32(ret1)
}

func StrCmpIW(lpszStr string, lpszComp string) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall3(strCmpIW, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		0)
	return int32(ret1)
}

func StrCmpLogicalW(lpszStr string, lpszComp string) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall3(strCmpLogicalW, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		0)
	return int32(ret1)
}

func StrCmpNC(lpszSrc string, lpszCmp string, aLen int32) uint32 {
	lpszSrcStr := unicode16FromString(lpszSrc)
	lpszCmpStr := unicode16FromString(lpszCmp)
	ret1 := syscall3(strCmpNC, 3,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(&lpszCmpStr[0])),
		uintptr(aLen))
	return uint32(ret1)
}

func StrCmpNIC(lpszSrc string, lpszCmp string, aLen uint32) uint32 {
	lpszSrcStr := unicode16FromString(lpszSrc)
	lpszCmpStr := unicode16FromString(lpszCmp)
	ret1 := syscall3(strCmpNIC, 3,
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		uintptr(unsafe.Pointer(&lpszCmpStr[0])),
		uintptr(aLen))
	return uint32(ret1)
}

func StrCmpNI(lpszStr string, lpszComp string, iLen int32) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall3(strCmpNI, 3,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		uintptr(iLen))
	return int32(ret1)
}

func StrCmpN(lpszStr string, lpszComp string, iLen int32) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall3(strCmpN, 3,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		uintptr(iLen))
	return int32(ret1)
}

func StrCmpW(lpszStr string, lpszComp string) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall3(strCmpW, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		0)
	return int32(ret1)
}

func StrCpyNW(dst LPWSTR, src string, count int32) LPWSTR {
	srcStr := unicode16FromString(src)
	ret1 := syscall3(strCpyNW, 3,
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(&srcStr[0])),
		uintptr(count))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrCpyW(lpszStr LPWSTR, lpszSrc string) LPWSTR {
	lpszSrcStr := unicode16FromString(lpszSrc)
	ret1 := syscall3(strCpyW, 2,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(unsafe.Pointer(&lpszSrcStr[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrDup(lpszStr string) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(strDup, 1,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrFormatByteSize64A(llBytes LONGLONG, lpszDest LPSTR, cchMax uint32) LPSTR {
	ret1 := syscall3(strFormatByteSize64A, 3,
		uintptr(llBytes),
		uintptr(unsafe.Pointer(lpszDest)),
		uintptr(cchMax))
	return (LPSTR)(unsafe.Pointer(ret1))
}

func StrFormatByteSize(llBytes LONGLONG, lpszDest LPWSTR, cchMax uint32) LPWSTR {
	ret1 := syscall3(strFormatByteSize, 3,
		uintptr(llBytes),
		uintptr(unsafe.Pointer(lpszDest)),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrFormatKBSize(llBytes LONGLONG, lpszDest LPWSTR, cchMax uint32) LPWSTR {
	ret1 := syscall3(strFormatKBSize, 3,
		uintptr(llBytes),
		uintptr(unsafe.Pointer(lpszDest)),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrFromTimeInterval(lpszStr LPWSTR, cchMax uint32, dwMS uint32, iDigits int32) int32 {
	ret1 := syscall6(strFromTimeInterval, 4,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(cchMax),
		uintptr(dwMS),
		uintptr(iDigits),
		0,
		0)
	return int32(ret1)
}

func StrIsIntlEqual(bCase bool, lpszStr string, lpszComp string, iLen int32) bool {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszCompStr := unicode16FromString(lpszComp)
	ret1 := syscall6(strIsIntlEqual, 4,
		getUintptrFromBool(bCase),
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszCompStr[0])),
		uintptr(iLen),
		0,
		0)
	return ret1 != 0
}

func StrNCat(lpszStr LPWSTR, lpszCat string, cchMax int32) LPWSTR {
	lpszCatStr := unicode16FromString(lpszCat)
	ret1 := syscall3(strNCat, 3,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(unsafe.Pointer(&lpszCatStr[0])),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrPBrk(lpszStr string, lpszMatch string) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszMatchStr := unicode16FromString(lpszMatch)
	ret1 := syscall3(strPBrk, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszMatchStr[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrRChrI(str string, end string, ch uint16) LPWSTR {
	strStr := unicode16FromString(str)
	endStr := unicode16FromString(end)
	ret1 := syscall3(strRChrI, 3,
		uintptr(unsafe.Pointer(&strStr[0])),
		uintptr(unsafe.Pointer(&endStr[0])),
		uintptr(ch))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrRChr(str string, end string, ch uint16) LPWSTR {
	strStr := unicode16FromString(str)
	endStr := unicode16FromString(end)
	ret1 := syscall3(strRChr, 3,
		uintptr(unsafe.Pointer(&strStr[0])),
		uintptr(unsafe.Pointer(&endStr[0])),
		uintptr(ch))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrRStrI(lpszStr string, lpszEnd string, lpszSearch string) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszEndStr := unicode16FromString(lpszEnd)
	lpszSearchStr := unicode16FromString(lpszSearch)
	ret1 := syscall3(strRStrI, 3,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszEndStr[0])),
		uintptr(unsafe.Pointer(&lpszSearchStr[0])))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrRetToBSTR(lpStrRet *STRRET, pidl /*const*/ LPCITEMIDLIST, pBstrOut *BSTR) HRESULT {
	ret1 := syscall3(strRetToBSTR, 3,
		uintptr(unsafe.Pointer(lpStrRet)),
		uintptr(unsafe.Pointer(pidl)),
		uintptr(unsafe.Pointer(pBstrOut)))
	return HRESULT(ret1)
}

func StrRetToBuf(src *STRRET, pidl /*const*/ *ITEMIDLIST, dest LPWSTR, aLen uint32) HRESULT {
	ret1 := syscall6(strRetToBuf, 4,
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(pidl)),
		uintptr(unsafe.Pointer(dest)),
		uintptr(aLen),
		0,
		0)
	return HRESULT(ret1)
}

func StrRetToStr(lpStrRet *STRRET, pidl /*const*/ *ITEMIDLIST, ppszName *LPWSTR) HRESULT {
	ret1 := syscall3(strRetToStr, 3,
		uintptr(unsafe.Pointer(lpStrRet)),
		uintptr(unsafe.Pointer(pidl)),
		uintptr(unsafe.Pointer(ppszName)))
	return HRESULT(ret1)
}

func StrSpn(lpszStr string, lpszMatch string) int32 {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszMatchStr := unicode16FromString(lpszMatch)
	ret1 := syscall3(strSpn, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszMatchStr[0])),
		0)
	return int32(ret1)
}

func StrStrI(lpszStr string, lpszSearch string) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszSearchStr := unicode16FromString(lpszSearch)
	ret1 := syscall3(strStrI, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszSearchStr[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrStrNIW(lpFirst string, lpSrch string, cchMax uint32) LPWSTR {
	lpFirstStr := unicode16FromString(lpFirst)
	lpSrchStr := unicode16FromString(lpSrch)
	ret1 := syscall3(strStrNIW, 3,
		uintptr(unsafe.Pointer(&lpFirstStr[0])),
		uintptr(unsafe.Pointer(&lpSrchStr[0])),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrStrNW(lpFirst string, lpSrch string, cchMax uint32) LPWSTR {
	lpFirstStr := unicode16FromString(lpFirst)
	lpSrchStr := unicode16FromString(lpSrch)
	ret1 := syscall3(strStrNW, 3,
		uintptr(unsafe.Pointer(&lpFirstStr[0])),
		uintptr(unsafe.Pointer(&lpSrchStr[0])),
		uintptr(cchMax))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrStr(lpszStr string, lpszSearch string) LPWSTR {
	lpszStrStr := unicode16FromString(lpszStr)
	lpszSearchStr := unicode16FromString(lpszSearch)
	ret1 := syscall3(strStr, 2,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(unsafe.Pointer(&lpszSearchStr[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func StrToInt64Ex(lpszStr string, dwFlags uint32, lpiRet *LONGLONG) bool {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(strToInt64Ex, 3,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpiRet)))
	return ret1 != 0
}

func StrToIntEx(lpszStr string, dwFlags uint32, lpiRet *int32) bool {
	lpszStrStr := unicode16FromString(lpszStr)
	ret1 := syscall3(strToIntEx, 3,
		uintptr(unsafe.Pointer(&lpszStrStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpiRet)))
	return ret1 != 0
}

func StrToInt(lpString string) int32 {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(strToInt, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return int32(ret1)
}

func StrTrim(lpszStr LPWSTR, lpszTrim string) bool {
	lpszTrimStr := unicode16FromString(lpszTrim)
	ret1 := syscall3(strTrim, 2,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(unsafe.Pointer(&lpszTrimStr[0])),
		0)
	return ret1 != 0
}

func UrlApplyScheme(pszIn string, pszOut LPWSTR, pcchOut *uint32, dwFlags uint32) HRESULT {
	pszInStr := unicode16FromString(pszIn)
	ret1 := syscall6(urlApplyScheme, 4,
		uintptr(unsafe.Pointer(&pszInStr[0])),
		uintptr(unsafe.Pointer(pszOut)),
		uintptr(unsafe.Pointer(pcchOut)),
		uintptr(dwFlags),
		0,
		0)
	return HRESULT(ret1)
}

func UrlCanonicalize(pszUrl string, pszCanonicalized LPWSTR, pcchCanonicalized *uint32, dwFlags uint32) HRESULT {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall6(urlCanonicalize, 4,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		uintptr(unsafe.Pointer(pszCanonicalized)),
		uintptr(unsafe.Pointer(pcchCanonicalized)),
		uintptr(dwFlags),
		0,
		0)
	return HRESULT(ret1)
}

func UrlCombine(pszBase string, pszRelative string, pszCombined LPWSTR, pcchCombined *uint32, dwFlags uint32) HRESULT {
	pszBaseStr := unicode16FromString(pszBase)
	pszRelativeStr := unicode16FromString(pszRelative)
	ret1 := syscall6(urlCombine, 5,
		uintptr(unsafe.Pointer(&pszBaseStr[0])),
		uintptr(unsafe.Pointer(&pszRelativeStr[0])),
		uintptr(unsafe.Pointer(pszCombined)),
		uintptr(unsafe.Pointer(pcchCombined)),
		uintptr(dwFlags),
		0)
	return HRESULT(ret1)
}

func UrlCompare(pszUrl1 string, pszUrl2 string, fIgnoreSlash bool) int32 {
	pszUrl1Str := unicode16FromString(pszUrl1)
	pszUrl2Str := unicode16FromString(pszUrl2)
	ret1 := syscall3(urlCompare, 3,
		uintptr(unsafe.Pointer(&pszUrl1Str[0])),
		uintptr(unsafe.Pointer(&pszUrl2Str[0])),
		getUintptrFromBool(fIgnoreSlash))
	return int32(ret1)
}

func UrlCreateFromPath(pszPath string, pszUrl LPWSTR, pcchUrl *uint32, dwReserved uint32) HRESULT {
	pszPathStr := unicode16FromString(pszPath)
	ret1 := syscall6(urlCreateFromPath, 4,
		uintptr(unsafe.Pointer(&pszPathStr[0])),
		uintptr(unsafe.Pointer(pszUrl)),
		uintptr(unsafe.Pointer(pcchUrl)),
		uintptr(dwReserved),
		0,
		0)
	return HRESULT(ret1)
}

func UrlEscape(pszUrl string, pszEscaped LPWSTR, pcchEscaped *uint32, dwFlags uint32) HRESULT {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall6(urlEscape, 4,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		uintptr(unsafe.Pointer(pszEscaped)),
		uintptr(unsafe.Pointer(pcchEscaped)),
		uintptr(dwFlags),
		0,
		0)
	return HRESULT(ret1)
}

func UrlFixupW(url string, translatedUrl LPWSTR, maxChars uint32) HRESULT {
	urlStr := unicode16FromString(url)
	ret1 := syscall3(urlFixupW, 3,
		uintptr(unsafe.Pointer(&urlStr[0])),
		uintptr(unsafe.Pointer(translatedUrl)),
		uintptr(maxChars))
	return HRESULT(ret1)
}

func UrlGetLocation(pszUrl string) string {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall3(urlGetLocation, 1,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		0,
		0)
	return stringFromUnicode16((*uint16)(unsafe.Pointer(ret1)))
}

func UrlGetPart(pszIn string, pszOut LPWSTR, pcchOut *uint32, dwPart uint32, dwFlags uint32) HRESULT {
	pszInStr := unicode16FromString(pszIn)
	ret1 := syscall6(urlGetPart, 5,
		uintptr(unsafe.Pointer(&pszInStr[0])),
		uintptr(unsafe.Pointer(pszOut)),
		uintptr(unsafe.Pointer(pcchOut)),
		uintptr(dwPart),
		uintptr(dwFlags),
		0)
	return HRESULT(ret1)
}

func UrlHash(pszUrl string, lpDest *byte, nDestLen uint32) HRESULT {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall3(urlHash, 3,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		uintptr(unsafe.Pointer(lpDest)),
		uintptr(nDestLen))
	return HRESULT(ret1)
}

func UrlIsNoHistory(pszUrl string) bool {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall3(urlIsNoHistory, 1,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		0,
		0)
	return ret1 != 0
}

func UrlIsOpaque(pszUrl string) bool {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall3(urlIsOpaque, 1,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		0,
		0)
	return ret1 != 0
}

func UrlIs(pszUrl string, urlis URLIS) bool {
	pszUrlStr := unicode16FromString(pszUrl)
	ret1 := syscall3(urlIs, 2,
		uintptr(unsafe.Pointer(&pszUrlStr[0])),
		uintptr(urlis),
		0)
	return ret1 != 0
}

func UrlUnescape(pszUrl LPWSTR, pszUnescaped LPWSTR, pcchUnescaped *uint32, dwFlags uint32) HRESULT {
	ret1 := syscall6(urlUnescape, 4,
		uintptr(unsafe.Pointer(pszUrl)),
		uintptr(unsafe.Pointer(pszUnescaped)),
		uintptr(unsafe.Pointer(pcchUnescaped)),
		uintptr(dwFlags),
		0,
		0)
	return HRESULT(ret1)
}

func WhichPlatform() uint32 {
	ret1 := syscall3(whichPlatform, 0,
		0,
		0,
		0)
	return uint32(ret1)
}
