// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	liboleaut32 uintptr

	// Functions
	bSTR_UserFree                  uintptr
	bSTR_UserSize                  uintptr
	bstrFromVector                 uintptr
	clearCustData                  uintptr
	createDispTypeInfo             uintptr
	createErrorInfo                uintptr
	createStdDispatch              uintptr
	createTypeLib                  uintptr
	createTypeLib2                 uintptr
	dispCallFunc                   uintptr
	dispGetIDsOfNames              uintptr
	dispGetParam                   uintptr
	dispInvoke                     uintptr
	dosDateTimeToVariantTime       uintptr
	getErrorInfo                   uintptr
	getRecordInfoFromGuids         uintptr
	getRecordInfoFromTypeInfo      uintptr
	lHashValOfNameSysA             uintptr
	lPSAFEARRAY_UserFree           uintptr
	lPSAFEARRAY_UserSize           uintptr
	loadRegTypeLib                 uintptr
	loadTypeLib                    uintptr
	loadTypeLibEx                  uintptr
	oaBuildVersion                 uintptr
	oleCreateFontIndirect          uintptr
	oleCreatePictureIndirect       uintptr
	oleCreatePropertyFrame         uintptr
	oleCreatePropertyFrameIndirect uintptr
	oleIconToCursor                uintptr
	oleLoadPicture                 uintptr
	oleLoadPictureEx               uintptr
	oleLoadPictureFile             uintptr
	oleLoadPicturePath             uintptr
	oleSavePictureFile             uintptr
	oleTranslateColor              uintptr
	queryPathOfRegTypeLib          uintptr
	registerTypeLib                uintptr
	registerTypeLibForUser         uintptr
	safeArrayAccessData            uintptr
	safeArrayAllocData             uintptr
	safeArrayAllocDescriptor       uintptr
	safeArrayAllocDescriptorEx     uintptr
	safeArrayCopy                  uintptr
	safeArrayCopyData              uintptr
	safeArrayDestroy               uintptr
	safeArrayDestroyData           uintptr
	safeArrayDestroyDescriptor     uintptr
	safeArrayGetDim                uintptr
	safeArrayGetElement            uintptr
	safeArrayGetElemsize           uintptr
	safeArrayGetIID                uintptr
	safeArrayGetLBound             uintptr
	safeArrayGetRecordInfo         uintptr
	safeArrayGetUBound             uintptr
	safeArrayGetVartype            uintptr
	safeArrayLock                  uintptr
	safeArrayPtrOfIndex            uintptr
	safeArrayPutElement            uintptr
	safeArrayRedim                 uintptr
	safeArraySetIID                uintptr
	safeArraySetRecordInfo         uintptr
	safeArrayUnaccessData          uintptr
	safeArrayUnlock                uintptr
	setErrorInfo                   uintptr
	setOaNoCache                   uintptr
	sysAllocString                 uintptr
	sysAllocStringByteLen          uintptr
	sysAllocStringLen              uintptr
	sysFreeString                  uintptr
	sysReAllocString               uintptr
	sysReAllocStringLen            uintptr
	sysStringByteLen               uintptr
	sysStringLen                   uintptr
	systemTimeToVariantTime        uintptr
	unRegisterTypeLib              uintptr
	unRegisterTypeLibForUser       uintptr
	vARIANT_UserFree               uintptr
	vARIANT_UserSize               uintptr
	varAbs                         uintptr
	varAdd                         uintptr
	varAnd                         uintptr
	varBoolFromCy                  uintptr
	varBoolFromDate                uintptr
	varBoolFromDec                 uintptr
	varBoolFromDisp                uintptr
	varBoolFromI1                  uintptr
	varBoolFromI2                  uintptr
	varBoolFromI4                  uintptr
	varBoolFromI8                  uintptr
	varBoolFromR4                  uintptr
	varBoolFromR8                  uintptr
	varBoolFromStr                 uintptr
	varBoolFromUI1                 uintptr
	varBoolFromUI2                 uintptr
	varBoolFromUI4                 uintptr
	varBoolFromUI8                 uintptr
	varBstrCat                     uintptr
	varBstrCmp                     uintptr
	varBstrFromBool                uintptr
	varBstrFromCy                  uintptr
	varBstrFromDate                uintptr
	varBstrFromDec                 uintptr
	varBstrFromDisp                uintptr
	varBstrFromI1                  uintptr
	varBstrFromI2                  uintptr
	varBstrFromI4                  uintptr
	varBstrFromI8                  uintptr
	varBstrFromR4                  uintptr
	varBstrFromR8                  uintptr
	varBstrFromUI1                 uintptr
	varBstrFromUI2                 uintptr
	varBstrFromUI4                 uintptr
	varBstrFromUI8                 uintptr
	varCat                         uintptr
	varCmp                         uintptr
	varCyAbs                       uintptr
	varCyAdd                       uintptr
	varCyCmp                       uintptr
	varCyCmpR8                     uintptr
	varCyFix                       uintptr
	varCyFromBool                  uintptr
	varCyFromDate                  uintptr
	varCyFromDec                   uintptr
	varCyFromDisp                  uintptr
	varCyFromI1                    uintptr
	varCyFromI2                    uintptr
	varCyFromI4                    uintptr
	varCyFromI8                    uintptr
	varCyFromR4                    uintptr
	varCyFromR8                    uintptr
	varCyFromStr                   uintptr
	varCyFromUI1                   uintptr
	varCyFromUI2                   uintptr
	varCyFromUI4                   uintptr
	varCyFromUI8                   uintptr
	varCyInt                       uintptr
	varCyMul                       uintptr
	varCyMulI4                     uintptr
	varCyMulI8                     uintptr
	varCyNeg                       uintptr
	varCyRound                     uintptr
	varCySub                       uintptr
	varDateFromBool                uintptr
	varDateFromCy                  uintptr
	varDateFromDec                 uintptr
	varDateFromDisp                uintptr
	varDateFromI1                  uintptr
	varDateFromI2                  uintptr
	varDateFromI4                  uintptr
	varDateFromI8                  uintptr
	varDateFromR4                  uintptr
	varDateFromR8                  uintptr
	varDateFromStr                 uintptr
	varDateFromUI1                 uintptr
	varDateFromUI2                 uintptr
	varDateFromUI4                 uintptr
	varDateFromUI8                 uintptr
	varDateFromUdate               uintptr
	varDateFromUdateEx             uintptr
	varDecAbs                      uintptr
	varDecAdd                      uintptr
	varDecCmp                      uintptr
	varDecCmpR8                    uintptr
	varDecDiv                      uintptr
	varDecFix                      uintptr
	varDecFromBool                 uintptr
	varDecFromCy                   uintptr
	varDecFromDate                 uintptr
	varDecFromDisp                 uintptr
	varDecFromI1                   uintptr
	varDecFromI2                   uintptr
	varDecFromI4                   uintptr
	varDecFromI8                   uintptr
	varDecFromR4                   uintptr
	varDecFromR8                   uintptr
	varDecFromStr                  uintptr
	varDecFromUI1                  uintptr
	varDecFromUI2                  uintptr
	varDecFromUI4                  uintptr
	varDecFromUI8                  uintptr
	varDecInt                      uintptr
	varDecMul                      uintptr
	varDecNeg                      uintptr
	varDecRound                    uintptr
	varDecSub                      uintptr
	varDiv                         uintptr
	varEqv                         uintptr
	varFix                         uintptr
	varFormat                      uintptr
	varFormatCurrency              uintptr
	varFormatDateTime              uintptr
	varFormatFromTokens            uintptr
	varFormatNumber                uintptr
	varFormatPercent               uintptr
	varI1FromBool                  uintptr
	varI1FromCy                    uintptr
	varI1FromDate                  uintptr
	varI1FromDec                   uintptr
	varI1FromDisp                  uintptr
	varI1FromI2                    uintptr
	varI1FromI4                    uintptr
	varI1FromI8                    uintptr
	varI1FromR4                    uintptr
	varI1FromR8                    uintptr
	varI1FromStr                   uintptr
	varI1FromUI1                   uintptr
	varI1FromUI2                   uintptr
	varI1FromUI4                   uintptr
	varI1FromUI8                   uintptr
	varI2FromBool                  uintptr
	varI2FromCy                    uintptr
	varI2FromDate                  uintptr
	varI2FromDec                   uintptr
	varI2FromDisp                  uintptr
	varI2FromI1                    uintptr
	varI2FromI4                    uintptr
	varI2FromI8                    uintptr
	varI2FromR4                    uintptr
	varI2FromR8                    uintptr
	varI2FromStr                   uintptr
	varI2FromUI1                   uintptr
	varI2FromUI2                   uintptr
	varI2FromUI4                   uintptr
	varI2FromUI8                   uintptr
	varI4FromBool                  uintptr
	varI4FromCy                    uintptr
	varI4FromDate                  uintptr
	varI4FromDec                   uintptr
	varI4FromDisp                  uintptr
	varI4FromI1                    uintptr
	varI4FromI2                    uintptr
	varI4FromI8                    uintptr
	varI4FromR4                    uintptr
	varI4FromR8                    uintptr
	varI4FromStr                   uintptr
	varI4FromUI1                   uintptr
	varI4FromUI2                   uintptr
	varI4FromUI4                   uintptr
	varI4FromUI8                   uintptr
	varI8FromBool                  uintptr
	varI8FromCy                    uintptr
	varI8FromDate                  uintptr
	varI8FromDec                   uintptr
	varI8FromDisp                  uintptr
	varI8FromI1                    uintptr
	varI8FromI2                    uintptr
	varI8FromR4                    uintptr
	varI8FromR8                    uintptr
	varI8FromStr                   uintptr
	varI8FromUI1                   uintptr
	varI8FromUI2                   uintptr
	varI8FromUI4                   uintptr
	varI8FromUI8                   uintptr
	varIdiv                        uintptr
	varImp                         uintptr
	varInt                         uintptr
	varMod                         uintptr
	varMonthName                   uintptr
	varMul                         uintptr
	varNeg                         uintptr
	varNot                         uintptr
	varNumFromParseNum             uintptr
	varOr                          uintptr
	varParseNumFromStr             uintptr
	varPow                         uintptr
	varR4CmpR8                     uintptr
	varR4FromBool                  uintptr
	varR4FromCy                    uintptr
	varR4FromDate                  uintptr
	varR4FromDec                   uintptr
	varR4FromDisp                  uintptr
	varR4FromI1                    uintptr
	varR4FromI2                    uintptr
	varR4FromI4                    uintptr
	varR4FromI8                    uintptr
	varR4FromR8                    uintptr
	varR4FromStr                   uintptr
	varR4FromUI1                   uintptr
	varR4FromUI2                   uintptr
	varR4FromUI4                   uintptr
	varR4FromUI8                   uintptr
	varR8FromBool                  uintptr
	varR8FromCy                    uintptr
	varR8FromDate                  uintptr
	varR8FromDec                   uintptr
	varR8FromDisp                  uintptr
	varR8FromI1                    uintptr
	varR8FromI2                    uintptr
	varR8FromI4                    uintptr
	varR8FromI8                    uintptr
	varR8FromR4                    uintptr
	varR8FromStr                   uintptr
	varR8FromUI1                   uintptr
	varR8FromUI2                   uintptr
	varR8FromUI4                   uintptr
	varR8FromUI8                   uintptr
	varR8Pow                       uintptr
	varR8Round                     uintptr
	varRound                       uintptr
	varSub                         uintptr
	varTokenizeFormatString        uintptr
	varUI1FromBool                 uintptr
	varUI1FromCy                   uintptr
	varUI1FromDate                 uintptr
	varUI1FromDec                  uintptr
	varUI1FromDisp                 uintptr
	varUI1FromI1                   uintptr
	varUI1FromI2                   uintptr
	varUI1FromI4                   uintptr
	varUI1FromI8                   uintptr
	varUI1FromR4                   uintptr
	varUI1FromR8                   uintptr
	varUI1FromStr                  uintptr
	varUI1FromUI2                  uintptr
	varUI1FromUI4                  uintptr
	varUI1FromUI8                  uintptr
	varUI2FromBool                 uintptr
	varUI2FromCy                   uintptr
	varUI2FromDate                 uintptr
	varUI2FromDec                  uintptr
	varUI2FromDisp                 uintptr
	varUI2FromI1                   uintptr
	varUI2FromI2                   uintptr
	varUI2FromI4                   uintptr
	varUI2FromI8                   uintptr
	varUI2FromR4                   uintptr
	varUI2FromR8                   uintptr
	varUI2FromStr                  uintptr
	varUI2FromUI1                  uintptr
	varUI2FromUI4                  uintptr
	varUI2FromUI8                  uintptr
	varUI4FromBool                 uintptr
	varUI4FromCy                   uintptr
	varUI4FromDate                 uintptr
	varUI4FromDec                  uintptr
	varUI4FromDisp                 uintptr
	varUI4FromI1                   uintptr
	varUI4FromI2                   uintptr
	varUI4FromI4                   uintptr
	varUI4FromI8                   uintptr
	varUI4FromR4                   uintptr
	varUI4FromR8                   uintptr
	varUI4FromStr                  uintptr
	varUI4FromUI1                  uintptr
	varUI4FromUI2                  uintptr
	varUI4FromUI8                  uintptr
	varUI8FromBool                 uintptr
	varUI8FromCy                   uintptr
	varUI8FromDate                 uintptr
	varUI8FromDec                  uintptr
	varUI8FromDisp                 uintptr
	varUI8FromI1                   uintptr
	varUI8FromI2                   uintptr
	varUI8FromI8                   uintptr
	varUI8FromR4                   uintptr
	varUI8FromR8                   uintptr
	varUI8FromStr                  uintptr
	varUI8FromUI1                  uintptr
	varUI8FromUI2                  uintptr
	varUI8FromUI4                  uintptr
	varUdateFromDate               uintptr
	varWeekdayName                 uintptr
	varXor                         uintptr
	variantChangeType              uintptr
	variantChangeTypeEx            uintptr
	variantClear                   uintptr
	variantCopy                    uintptr
	variantCopyInd                 uintptr
	variantInit                    uintptr
	variantTimeToDosDateTime       uintptr
	variantTimeToSystemTime        uintptr
	vectorFromBstr                 uintptr
)

func init() {
	// Library
	liboleaut32 = doLoadLibrary("oleaut32.dll")

	// Functions
	bSTR_UserFree = doGetProcAddress(liboleaut32, "BSTR_UserFree")
	bSTR_UserSize = doGetProcAddress(liboleaut32, "BSTR_UserSize")
	bstrFromVector = doGetProcAddress(liboleaut32, "BstrFromVector")
	clearCustData = doGetProcAddress(liboleaut32, "ClearCustData")
	createDispTypeInfo = doGetProcAddress(liboleaut32, "CreateDispTypeInfo")
	createErrorInfo = doGetProcAddress(liboleaut32, "CreateErrorInfo")
	createStdDispatch = doGetProcAddress(liboleaut32, "CreateStdDispatch")
	createTypeLib = doGetProcAddress(liboleaut32, "CreateTypeLib")
	createTypeLib2 = doGetProcAddress(liboleaut32, "CreateTypeLib2")
	dispCallFunc = doGetProcAddress(liboleaut32, "DispCallFunc")
	dispGetIDsOfNames = doGetProcAddress(liboleaut32, "DispGetIDsOfNames")
	dispGetParam = doGetProcAddress(liboleaut32, "DispGetParam")
	dispInvoke = doGetProcAddress(liboleaut32, "DispInvoke")
	dosDateTimeToVariantTime = doGetProcAddress(liboleaut32, "DosDateTimeToVariantTime")
	getErrorInfo = doGetProcAddress(liboleaut32, "GetErrorInfo")
	getRecordInfoFromGuids = doGetProcAddress(liboleaut32, "GetRecordInfoFromGuids")
	getRecordInfoFromTypeInfo = doGetProcAddress(liboleaut32, "GetRecordInfoFromTypeInfo")
	lHashValOfNameSysA = doGetProcAddress(liboleaut32, "LHashValOfNameSysA")
	lPSAFEARRAY_UserFree = doGetProcAddress(liboleaut32, "LPSAFEARRAY_UserFree")
	lPSAFEARRAY_UserSize = doGetProcAddress(liboleaut32, "LPSAFEARRAY_UserSize")
	loadRegTypeLib = doGetProcAddress(liboleaut32, "LoadRegTypeLib")
	loadTypeLib = doGetProcAddress(liboleaut32, "LoadTypeLib")
	loadTypeLibEx = doGetProcAddress(liboleaut32, "LoadTypeLibEx")
	oaBuildVersion = doGetProcAddress(liboleaut32, "OaBuildVersion")
	oleCreateFontIndirect = doGetProcAddress(liboleaut32, "OleCreateFontIndirect")
	oleCreatePictureIndirect = doGetProcAddress(liboleaut32, "OleCreatePictureIndirect")
	oleCreatePropertyFrame = doGetProcAddress(liboleaut32, "OleCreatePropertyFrame")
	oleCreatePropertyFrameIndirect = doGetProcAddress(liboleaut32, "OleCreatePropertyFrameIndirect")
	oleIconToCursor = doGetProcAddress(liboleaut32, "OleIconToCursor")
	oleLoadPicture = doGetProcAddress(liboleaut32, "OleLoadPicture")
	oleLoadPictureEx = doGetProcAddress(liboleaut32, "OleLoadPictureEx")
	oleLoadPictureFile = doGetProcAddress(liboleaut32, "OleLoadPictureFile")
	oleLoadPicturePath = doGetProcAddress(liboleaut32, "OleLoadPicturePath")
	oleSavePictureFile = doGetProcAddress(liboleaut32, "OleSavePictureFile")
	oleTranslateColor = doGetProcAddress(liboleaut32, "OleTranslateColor")
	queryPathOfRegTypeLib = doGetProcAddress(liboleaut32, "QueryPathOfRegTypeLib")
	registerTypeLib = doGetProcAddress(liboleaut32, "RegisterTypeLib")
	registerTypeLibForUser = doGetProcAddress(liboleaut32, "RegisterTypeLibForUser")
	safeArrayAccessData = doGetProcAddress(liboleaut32, "SafeArrayAccessData")
	safeArrayAllocData = doGetProcAddress(liboleaut32, "SafeArrayAllocData")
	safeArrayAllocDescriptor = doGetProcAddress(liboleaut32, "SafeArrayAllocDescriptor")
	safeArrayAllocDescriptorEx = doGetProcAddress(liboleaut32, "SafeArrayAllocDescriptorEx")
	safeArrayCopy = doGetProcAddress(liboleaut32, "SafeArrayCopy")
	safeArrayCopyData = doGetProcAddress(liboleaut32, "SafeArrayCopyData")
	safeArrayDestroy = doGetProcAddress(liboleaut32, "SafeArrayDestroy")
	safeArrayDestroyData = doGetProcAddress(liboleaut32, "SafeArrayDestroyData")
	safeArrayDestroyDescriptor = doGetProcAddress(liboleaut32, "SafeArrayDestroyDescriptor")
	safeArrayGetDim = doGetProcAddress(liboleaut32, "SafeArrayGetDim")
	safeArrayGetElement = doGetProcAddress(liboleaut32, "SafeArrayGetElement")
	safeArrayGetElemsize = doGetProcAddress(liboleaut32, "SafeArrayGetElemsize")
	safeArrayGetIID = doGetProcAddress(liboleaut32, "SafeArrayGetIID")
	safeArrayGetLBound = doGetProcAddress(liboleaut32, "SafeArrayGetLBound")
	safeArrayGetRecordInfo = doGetProcAddress(liboleaut32, "SafeArrayGetRecordInfo")
	safeArrayGetUBound = doGetProcAddress(liboleaut32, "SafeArrayGetUBound")
	safeArrayGetVartype = doGetProcAddress(liboleaut32, "SafeArrayGetVartype")
	safeArrayLock = doGetProcAddress(liboleaut32, "SafeArrayLock")
	safeArrayPtrOfIndex = doGetProcAddress(liboleaut32, "SafeArrayPtrOfIndex")
	safeArrayPutElement = doGetProcAddress(liboleaut32, "SafeArrayPutElement")
	safeArrayRedim = doGetProcAddress(liboleaut32, "SafeArrayRedim")
	safeArraySetIID = doGetProcAddress(liboleaut32, "SafeArraySetIID")
	safeArraySetRecordInfo = doGetProcAddress(liboleaut32, "SafeArraySetRecordInfo")
	safeArrayUnaccessData = doGetProcAddress(liboleaut32, "SafeArrayUnaccessData")
	safeArrayUnlock = doGetProcAddress(liboleaut32, "SafeArrayUnlock")
	setErrorInfo = doGetProcAddress(liboleaut32, "SetErrorInfo")
	setOaNoCache = doGetProcAddress(liboleaut32, "SetOaNoCache")
	sysAllocString = doGetProcAddress(liboleaut32, "SysAllocString")
	sysAllocStringByteLen = doGetProcAddress(liboleaut32, "SysAllocStringByteLen")
	sysAllocStringLen = doGetProcAddress(liboleaut32, "SysAllocStringLen")
	sysFreeString = doGetProcAddress(liboleaut32, "SysFreeString")
	sysReAllocString = doGetProcAddress(liboleaut32, "SysReAllocString")
	sysReAllocStringLen = doGetProcAddress(liboleaut32, "SysReAllocStringLen")
	sysStringByteLen = doGetProcAddress(liboleaut32, "SysStringByteLen")
	sysStringLen = doGetProcAddress(liboleaut32, "SysStringLen")
	systemTimeToVariantTime = doGetProcAddress(liboleaut32, "SystemTimeToVariantTime")
	unRegisterTypeLib = doGetProcAddress(liboleaut32, "UnRegisterTypeLib")
	unRegisterTypeLibForUser = doGetProcAddress(liboleaut32, "UnRegisterTypeLibForUser")
	vARIANT_UserFree = doGetProcAddress(liboleaut32, "VARIANT_UserFree")
	vARIANT_UserSize = doGetProcAddress(liboleaut32, "VARIANT_UserSize")
	varAbs = doGetProcAddress(liboleaut32, "VarAbs")
	varAdd = doGetProcAddress(liboleaut32, "VarAdd")
	varAnd = doGetProcAddress(liboleaut32, "VarAnd")
	varBoolFromCy = doGetProcAddress(liboleaut32, "VarBoolFromCy")
	varBoolFromDate = doGetProcAddress(liboleaut32, "VarBoolFromDate")
	varBoolFromDec = doGetProcAddress(liboleaut32, "VarBoolFromDec")
	varBoolFromDisp = doGetProcAddress(liboleaut32, "VarBoolFromDisp")
	varBoolFromI1 = doGetProcAddress(liboleaut32, "VarBoolFromI1")
	varBoolFromI2 = doGetProcAddress(liboleaut32, "VarBoolFromI2")
	varBoolFromI4 = doGetProcAddress(liboleaut32, "VarBoolFromI4")
	varBoolFromI8 = doGetProcAddress(liboleaut32, "VarBoolFromI8")
	varBoolFromR4 = doGetProcAddress(liboleaut32, "VarBoolFromR4")
	varBoolFromR8 = doGetProcAddress(liboleaut32, "VarBoolFromR8")
	varBoolFromStr = doGetProcAddress(liboleaut32, "VarBoolFromStr")
	varBoolFromUI1 = doGetProcAddress(liboleaut32, "VarBoolFromUI1")
	varBoolFromUI2 = doGetProcAddress(liboleaut32, "VarBoolFromUI2")
	varBoolFromUI4 = doGetProcAddress(liboleaut32, "VarBoolFromUI4")
	varBoolFromUI8 = doGetProcAddress(liboleaut32, "VarBoolFromUI8")
	varBstrCat = doGetProcAddress(liboleaut32, "VarBstrCat")
	varBstrCmp = doGetProcAddress(liboleaut32, "VarBstrCmp")
	varBstrFromBool = doGetProcAddress(liboleaut32, "VarBstrFromBool")
	varBstrFromCy = doGetProcAddress(liboleaut32, "VarBstrFromCy")
	varBstrFromDate = doGetProcAddress(liboleaut32, "VarBstrFromDate")
	varBstrFromDec = doGetProcAddress(liboleaut32, "VarBstrFromDec")
	varBstrFromDisp = doGetProcAddress(liboleaut32, "VarBstrFromDisp")
	varBstrFromI1 = doGetProcAddress(liboleaut32, "VarBstrFromI1")
	varBstrFromI2 = doGetProcAddress(liboleaut32, "VarBstrFromI2")
	varBstrFromI4 = doGetProcAddress(liboleaut32, "VarBstrFromI4")
	varBstrFromI8 = doGetProcAddress(liboleaut32, "VarBstrFromI8")
	varBstrFromR4 = doGetProcAddress(liboleaut32, "VarBstrFromR4")
	varBstrFromR8 = doGetProcAddress(liboleaut32, "VarBstrFromR8")
	varBstrFromUI1 = doGetProcAddress(liboleaut32, "VarBstrFromUI1")
	varBstrFromUI2 = doGetProcAddress(liboleaut32, "VarBstrFromUI2")
	varBstrFromUI4 = doGetProcAddress(liboleaut32, "VarBstrFromUI4")
	varBstrFromUI8 = doGetProcAddress(liboleaut32, "VarBstrFromUI8")
	varCat = doGetProcAddress(liboleaut32, "VarCat")
	varCmp = doGetProcAddress(liboleaut32, "VarCmp")
	varCyAbs = doGetProcAddress(liboleaut32, "VarCyAbs")
	varCyAdd = doGetProcAddress(liboleaut32, "VarCyAdd")
	varCyCmp = doGetProcAddress(liboleaut32, "VarCyCmp")
	varCyCmpR8 = doGetProcAddress(liboleaut32, "VarCyCmpR8")
	varCyFix = doGetProcAddress(liboleaut32, "VarCyFix")
	varCyFromBool = doGetProcAddress(liboleaut32, "VarCyFromBool")
	varCyFromDate = doGetProcAddress(liboleaut32, "VarCyFromDate")
	varCyFromDec = doGetProcAddress(liboleaut32, "VarCyFromDec")
	varCyFromDisp = doGetProcAddress(liboleaut32, "VarCyFromDisp")
	varCyFromI1 = doGetProcAddress(liboleaut32, "VarCyFromI1")
	varCyFromI2 = doGetProcAddress(liboleaut32, "VarCyFromI2")
	varCyFromI4 = doGetProcAddress(liboleaut32, "VarCyFromI4")
	varCyFromI8 = doGetProcAddress(liboleaut32, "VarCyFromI8")
	varCyFromR4 = doGetProcAddress(liboleaut32, "VarCyFromR4")
	varCyFromR8 = doGetProcAddress(liboleaut32, "VarCyFromR8")
	varCyFromStr = doGetProcAddress(liboleaut32, "VarCyFromStr")
	varCyFromUI1 = doGetProcAddress(liboleaut32, "VarCyFromUI1")
	varCyFromUI2 = doGetProcAddress(liboleaut32, "VarCyFromUI2")
	varCyFromUI4 = doGetProcAddress(liboleaut32, "VarCyFromUI4")
	varCyFromUI8 = doGetProcAddress(liboleaut32, "VarCyFromUI8")
	varCyInt = doGetProcAddress(liboleaut32, "VarCyInt")
	varCyMul = doGetProcAddress(liboleaut32, "VarCyMul")
	varCyMulI4 = doGetProcAddress(liboleaut32, "VarCyMulI4")
	varCyMulI8 = doGetProcAddress(liboleaut32, "VarCyMulI8")
	varCyNeg = doGetProcAddress(liboleaut32, "VarCyNeg")
	varCyRound = doGetProcAddress(liboleaut32, "VarCyRound")
	varCySub = doGetProcAddress(liboleaut32, "VarCySub")
	varDateFromBool = doGetProcAddress(liboleaut32, "VarDateFromBool")
	varDateFromCy = doGetProcAddress(liboleaut32, "VarDateFromCy")
	varDateFromDec = doGetProcAddress(liboleaut32, "VarDateFromDec")
	varDateFromDisp = doGetProcAddress(liboleaut32, "VarDateFromDisp")
	varDateFromI1 = doGetProcAddress(liboleaut32, "VarDateFromI1")
	varDateFromI2 = doGetProcAddress(liboleaut32, "VarDateFromI2")
	varDateFromI4 = doGetProcAddress(liboleaut32, "VarDateFromI4")
	varDateFromI8 = doGetProcAddress(liboleaut32, "VarDateFromI8")
	varDateFromR4 = doGetProcAddress(liboleaut32, "VarDateFromR4")
	varDateFromR8 = doGetProcAddress(liboleaut32, "VarDateFromR8")
	varDateFromStr = doGetProcAddress(liboleaut32, "VarDateFromStr")
	varDateFromUI1 = doGetProcAddress(liboleaut32, "VarDateFromUI1")
	varDateFromUI2 = doGetProcAddress(liboleaut32, "VarDateFromUI2")
	varDateFromUI4 = doGetProcAddress(liboleaut32, "VarDateFromUI4")
	varDateFromUI8 = doGetProcAddress(liboleaut32, "VarDateFromUI8")
	varDateFromUdate = doGetProcAddress(liboleaut32, "VarDateFromUdate")
	varDateFromUdateEx = doGetProcAddress(liboleaut32, "VarDateFromUdateEx")
	varDecAbs = doGetProcAddress(liboleaut32, "VarDecAbs")
	varDecAdd = doGetProcAddress(liboleaut32, "VarDecAdd")
	varDecCmp = doGetProcAddress(liboleaut32, "VarDecCmp")
	varDecCmpR8 = doGetProcAddress(liboleaut32, "VarDecCmpR8")
	varDecDiv = doGetProcAddress(liboleaut32, "VarDecDiv")
	varDecFix = doGetProcAddress(liboleaut32, "VarDecFix")
	varDecFromBool = doGetProcAddress(liboleaut32, "VarDecFromBool")
	varDecFromCy = doGetProcAddress(liboleaut32, "VarDecFromCy")
	varDecFromDate = doGetProcAddress(liboleaut32, "VarDecFromDate")
	varDecFromDisp = doGetProcAddress(liboleaut32, "VarDecFromDisp")
	varDecFromI1 = doGetProcAddress(liboleaut32, "VarDecFromI1")
	varDecFromI2 = doGetProcAddress(liboleaut32, "VarDecFromI2")
	varDecFromI4 = doGetProcAddress(liboleaut32, "VarDecFromI4")
	varDecFromI8 = doGetProcAddress(liboleaut32, "VarDecFromI8")
	varDecFromR4 = doGetProcAddress(liboleaut32, "VarDecFromR4")
	varDecFromR8 = doGetProcAddress(liboleaut32, "VarDecFromR8")
	varDecFromStr = doGetProcAddress(liboleaut32, "VarDecFromStr")
	varDecFromUI1 = doGetProcAddress(liboleaut32, "VarDecFromUI1")
	varDecFromUI2 = doGetProcAddress(liboleaut32, "VarDecFromUI2")
	varDecFromUI4 = doGetProcAddress(liboleaut32, "VarDecFromUI4")
	varDecFromUI8 = doGetProcAddress(liboleaut32, "VarDecFromUI8")
	varDecInt = doGetProcAddress(liboleaut32, "VarDecInt")
	varDecMul = doGetProcAddress(liboleaut32, "VarDecMul")
	varDecNeg = doGetProcAddress(liboleaut32, "VarDecNeg")
	varDecRound = doGetProcAddress(liboleaut32, "VarDecRound")
	varDecSub = doGetProcAddress(liboleaut32, "VarDecSub")
	varDiv = doGetProcAddress(liboleaut32, "VarDiv")
	varEqv = doGetProcAddress(liboleaut32, "VarEqv")
	varFix = doGetProcAddress(liboleaut32, "VarFix")
	varFormat = doGetProcAddress(liboleaut32, "VarFormat")
	varFormatCurrency = doGetProcAddress(liboleaut32, "VarFormatCurrency")
	varFormatDateTime = doGetProcAddress(liboleaut32, "VarFormatDateTime")
	varFormatFromTokens = doGetProcAddress(liboleaut32, "VarFormatFromTokens")
	varFormatNumber = doGetProcAddress(liboleaut32, "VarFormatNumber")
	varFormatPercent = doGetProcAddress(liboleaut32, "VarFormatPercent")
	varI1FromBool = doGetProcAddress(liboleaut32, "VarI1FromBool")
	varI1FromCy = doGetProcAddress(liboleaut32, "VarI1FromCy")
	varI1FromDate = doGetProcAddress(liboleaut32, "VarI1FromDate")
	varI1FromDec = doGetProcAddress(liboleaut32, "VarI1FromDec")
	varI1FromDisp = doGetProcAddress(liboleaut32, "VarI1FromDisp")
	varI1FromI2 = doGetProcAddress(liboleaut32, "VarI1FromI2")
	varI1FromI4 = doGetProcAddress(liboleaut32, "VarI1FromI4")
	varI1FromI8 = doGetProcAddress(liboleaut32, "VarI1FromI8")
	varI1FromR4 = doGetProcAddress(liboleaut32, "VarI1FromR4")
	varI1FromR8 = doGetProcAddress(liboleaut32, "VarI1FromR8")
	varI1FromStr = doGetProcAddress(liboleaut32, "VarI1FromStr")
	varI1FromUI1 = doGetProcAddress(liboleaut32, "VarI1FromUI1")
	varI1FromUI2 = doGetProcAddress(liboleaut32, "VarI1FromUI2")
	varI1FromUI4 = doGetProcAddress(liboleaut32, "VarI1FromUI4")
	varI1FromUI8 = doGetProcAddress(liboleaut32, "VarI1FromUI8")
	varI2FromBool = doGetProcAddress(liboleaut32, "VarI2FromBool")
	varI2FromCy = doGetProcAddress(liboleaut32, "VarI2FromCy")
	varI2FromDate = doGetProcAddress(liboleaut32, "VarI2FromDate")
	varI2FromDec = doGetProcAddress(liboleaut32, "VarI2FromDec")
	varI2FromDisp = doGetProcAddress(liboleaut32, "VarI2FromDisp")
	varI2FromI1 = doGetProcAddress(liboleaut32, "VarI2FromI1")
	varI2FromI4 = doGetProcAddress(liboleaut32, "VarI2FromI4")
	varI2FromI8 = doGetProcAddress(liboleaut32, "VarI2FromI8")
	varI2FromR4 = doGetProcAddress(liboleaut32, "VarI2FromR4")
	varI2FromR8 = doGetProcAddress(liboleaut32, "VarI2FromR8")
	varI2FromStr = doGetProcAddress(liboleaut32, "VarI2FromStr")
	varI2FromUI1 = doGetProcAddress(liboleaut32, "VarI2FromUI1")
	varI2FromUI2 = doGetProcAddress(liboleaut32, "VarI2FromUI2")
	varI2FromUI4 = doGetProcAddress(liboleaut32, "VarI2FromUI4")
	varI2FromUI8 = doGetProcAddress(liboleaut32, "VarI2FromUI8")
	varI4FromBool = doGetProcAddress(liboleaut32, "VarI4FromBool")
	varI4FromCy = doGetProcAddress(liboleaut32, "VarI4FromCy")
	varI4FromDate = doGetProcAddress(liboleaut32, "VarI4FromDate")
	varI4FromDec = doGetProcAddress(liboleaut32, "VarI4FromDec")
	varI4FromDisp = doGetProcAddress(liboleaut32, "VarI4FromDisp")
	varI4FromI1 = doGetProcAddress(liboleaut32, "VarI4FromI1")
	varI4FromI2 = doGetProcAddress(liboleaut32, "VarI4FromI2")
	varI4FromI8 = doGetProcAddress(liboleaut32, "VarI4FromI8")
	varI4FromR4 = doGetProcAddress(liboleaut32, "VarI4FromR4")
	varI4FromR8 = doGetProcAddress(liboleaut32, "VarI4FromR8")
	varI4FromStr = doGetProcAddress(liboleaut32, "VarI4FromStr")
	varI4FromUI1 = doGetProcAddress(liboleaut32, "VarI4FromUI1")
	varI4FromUI2 = doGetProcAddress(liboleaut32, "VarI4FromUI2")
	varI4FromUI4 = doGetProcAddress(liboleaut32, "VarI4FromUI4")
	varI4FromUI8 = doGetProcAddress(liboleaut32, "VarI4FromUI8")
	varI8FromBool = doGetProcAddress(liboleaut32, "VarI8FromBool")
	varI8FromCy = doGetProcAddress(liboleaut32, "VarI8FromCy")
	varI8FromDate = doGetProcAddress(liboleaut32, "VarI8FromDate")
	varI8FromDec = doGetProcAddress(liboleaut32, "VarI8FromDec")
	varI8FromDisp = doGetProcAddress(liboleaut32, "VarI8FromDisp")
	varI8FromI1 = doGetProcAddress(liboleaut32, "VarI8FromI1")
	varI8FromI2 = doGetProcAddress(liboleaut32, "VarI8FromI2")
	varI8FromR4 = doGetProcAddress(liboleaut32, "VarI8FromR4")
	varI8FromR8 = doGetProcAddress(liboleaut32, "VarI8FromR8")
	varI8FromStr = doGetProcAddress(liboleaut32, "VarI8FromStr")
	varI8FromUI1 = doGetProcAddress(liboleaut32, "VarI8FromUI1")
	varI8FromUI2 = doGetProcAddress(liboleaut32, "VarI8FromUI2")
	varI8FromUI4 = doGetProcAddress(liboleaut32, "VarI8FromUI4")
	varI8FromUI8 = doGetProcAddress(liboleaut32, "VarI8FromUI8")
	varIdiv = doGetProcAddress(liboleaut32, "VarIdiv")
	varImp = doGetProcAddress(liboleaut32, "VarImp")
	varInt = doGetProcAddress(liboleaut32, "VarInt")
	varMod = doGetProcAddress(liboleaut32, "VarMod")
	varMonthName = doGetProcAddress(liboleaut32, "VarMonthName")
	varMul = doGetProcAddress(liboleaut32, "VarMul")
	varNeg = doGetProcAddress(liboleaut32, "VarNeg")
	varNot = doGetProcAddress(liboleaut32, "VarNot")
	varNumFromParseNum = doGetProcAddress(liboleaut32, "VarNumFromParseNum")
	varOr = doGetProcAddress(liboleaut32, "VarOr")
	varParseNumFromStr = doGetProcAddress(liboleaut32, "VarParseNumFromStr")
	varPow = doGetProcAddress(liboleaut32, "VarPow")
	varR4CmpR8 = doGetProcAddress(liboleaut32, "VarR4CmpR8")
	varR4FromBool = doGetProcAddress(liboleaut32, "VarR4FromBool")
	varR4FromCy = doGetProcAddress(liboleaut32, "VarR4FromCy")
	varR4FromDate = doGetProcAddress(liboleaut32, "VarR4FromDate")
	varR4FromDec = doGetProcAddress(liboleaut32, "VarR4FromDec")
	varR4FromDisp = doGetProcAddress(liboleaut32, "VarR4FromDisp")
	varR4FromI1 = doGetProcAddress(liboleaut32, "VarR4FromI1")
	varR4FromI2 = doGetProcAddress(liboleaut32, "VarR4FromI2")
	varR4FromI4 = doGetProcAddress(liboleaut32, "VarR4FromI4")
	varR4FromI8 = doGetProcAddress(liboleaut32, "VarR4FromI8")
	varR4FromR8 = doGetProcAddress(liboleaut32, "VarR4FromR8")
	varR4FromStr = doGetProcAddress(liboleaut32, "VarR4FromStr")
	varR4FromUI1 = doGetProcAddress(liboleaut32, "VarR4FromUI1")
	varR4FromUI2 = doGetProcAddress(liboleaut32, "VarR4FromUI2")
	varR4FromUI4 = doGetProcAddress(liboleaut32, "VarR4FromUI4")
	varR4FromUI8 = doGetProcAddress(liboleaut32, "VarR4FromUI8")
	varR8FromBool = doGetProcAddress(liboleaut32, "VarR8FromBool")
	varR8FromCy = doGetProcAddress(liboleaut32, "VarR8FromCy")
	varR8FromDate = doGetProcAddress(liboleaut32, "VarR8FromDate")
	varR8FromDec = doGetProcAddress(liboleaut32, "VarR8FromDec")
	varR8FromDisp = doGetProcAddress(liboleaut32, "VarR8FromDisp")
	varR8FromI1 = doGetProcAddress(liboleaut32, "VarR8FromI1")
	varR8FromI2 = doGetProcAddress(liboleaut32, "VarR8FromI2")
	varR8FromI4 = doGetProcAddress(liboleaut32, "VarR8FromI4")
	varR8FromI8 = doGetProcAddress(liboleaut32, "VarR8FromI8")
	varR8FromR4 = doGetProcAddress(liboleaut32, "VarR8FromR4")
	varR8FromStr = doGetProcAddress(liboleaut32, "VarR8FromStr")
	varR8FromUI1 = doGetProcAddress(liboleaut32, "VarR8FromUI1")
	varR8FromUI2 = doGetProcAddress(liboleaut32, "VarR8FromUI2")
	varR8FromUI4 = doGetProcAddress(liboleaut32, "VarR8FromUI4")
	varR8FromUI8 = doGetProcAddress(liboleaut32, "VarR8FromUI8")
	varR8Pow = doGetProcAddress(liboleaut32, "VarR8Pow")
	varR8Round = doGetProcAddress(liboleaut32, "VarR8Round")
	varRound = doGetProcAddress(liboleaut32, "VarRound")
	varSub = doGetProcAddress(liboleaut32, "VarSub")
	varTokenizeFormatString = doGetProcAddress(liboleaut32, "VarTokenizeFormatString")
	varUI1FromBool = doGetProcAddress(liboleaut32, "VarUI1FromBool")
	varUI1FromCy = doGetProcAddress(liboleaut32, "VarUI1FromCy")
	varUI1FromDate = doGetProcAddress(liboleaut32, "VarUI1FromDate")
	varUI1FromDec = doGetProcAddress(liboleaut32, "VarUI1FromDec")
	varUI1FromDisp = doGetProcAddress(liboleaut32, "VarUI1FromDisp")
	varUI1FromI1 = doGetProcAddress(liboleaut32, "VarUI1FromI1")
	varUI1FromI2 = doGetProcAddress(liboleaut32, "VarUI1FromI2")
	varUI1FromI4 = doGetProcAddress(liboleaut32, "VarUI1FromI4")
	varUI1FromI8 = doGetProcAddress(liboleaut32, "VarUI1FromI8")
	varUI1FromR4 = doGetProcAddress(liboleaut32, "VarUI1FromR4")
	varUI1FromR8 = doGetProcAddress(liboleaut32, "VarUI1FromR8")
	varUI1FromStr = doGetProcAddress(liboleaut32, "VarUI1FromStr")
	varUI1FromUI2 = doGetProcAddress(liboleaut32, "VarUI1FromUI2")
	varUI1FromUI4 = doGetProcAddress(liboleaut32, "VarUI1FromUI4")
	varUI1FromUI8 = doGetProcAddress(liboleaut32, "VarUI1FromUI8")
	varUI2FromBool = doGetProcAddress(liboleaut32, "VarUI2FromBool")
	varUI2FromCy = doGetProcAddress(liboleaut32, "VarUI2FromCy")
	varUI2FromDate = doGetProcAddress(liboleaut32, "VarUI2FromDate")
	varUI2FromDec = doGetProcAddress(liboleaut32, "VarUI2FromDec")
	varUI2FromDisp = doGetProcAddress(liboleaut32, "VarUI2FromDisp")
	varUI2FromI1 = doGetProcAddress(liboleaut32, "VarUI2FromI1")
	varUI2FromI2 = doGetProcAddress(liboleaut32, "VarUI2FromI2")
	varUI2FromI4 = doGetProcAddress(liboleaut32, "VarUI2FromI4")
	varUI2FromI8 = doGetProcAddress(liboleaut32, "VarUI2FromI8")
	varUI2FromR4 = doGetProcAddress(liboleaut32, "VarUI2FromR4")
	varUI2FromR8 = doGetProcAddress(liboleaut32, "VarUI2FromR8")
	varUI2FromStr = doGetProcAddress(liboleaut32, "VarUI2FromStr")
	varUI2FromUI1 = doGetProcAddress(liboleaut32, "VarUI2FromUI1")
	varUI2FromUI4 = doGetProcAddress(liboleaut32, "VarUI2FromUI4")
	varUI2FromUI8 = doGetProcAddress(liboleaut32, "VarUI2FromUI8")
	varUI4FromBool = doGetProcAddress(liboleaut32, "VarUI4FromBool")
	varUI4FromCy = doGetProcAddress(liboleaut32, "VarUI4FromCy")
	varUI4FromDate = doGetProcAddress(liboleaut32, "VarUI4FromDate")
	varUI4FromDec = doGetProcAddress(liboleaut32, "VarUI4FromDec")
	varUI4FromDisp = doGetProcAddress(liboleaut32, "VarUI4FromDisp")
	varUI4FromI1 = doGetProcAddress(liboleaut32, "VarUI4FromI1")
	varUI4FromI2 = doGetProcAddress(liboleaut32, "VarUI4FromI2")
	varUI4FromI4 = doGetProcAddress(liboleaut32, "VarUI4FromI4")
	varUI4FromI8 = doGetProcAddress(liboleaut32, "VarUI4FromI8")
	varUI4FromR4 = doGetProcAddress(liboleaut32, "VarUI4FromR4")
	varUI4FromR8 = doGetProcAddress(liboleaut32, "VarUI4FromR8")
	varUI4FromStr = doGetProcAddress(liboleaut32, "VarUI4FromStr")
	varUI4FromUI1 = doGetProcAddress(liboleaut32, "VarUI4FromUI1")
	varUI4FromUI2 = doGetProcAddress(liboleaut32, "VarUI4FromUI2")
	varUI4FromUI8 = doGetProcAddress(liboleaut32, "VarUI4FromUI8")
	varUI8FromBool = doGetProcAddress(liboleaut32, "VarUI8FromBool")
	varUI8FromCy = doGetProcAddress(liboleaut32, "VarUI8FromCy")
	varUI8FromDate = doGetProcAddress(liboleaut32, "VarUI8FromDate")
	varUI8FromDec = doGetProcAddress(liboleaut32, "VarUI8FromDec")
	varUI8FromDisp = doGetProcAddress(liboleaut32, "VarUI8FromDisp")
	varUI8FromI1 = doGetProcAddress(liboleaut32, "VarUI8FromI1")
	varUI8FromI2 = doGetProcAddress(liboleaut32, "VarUI8FromI2")
	varUI8FromI8 = doGetProcAddress(liboleaut32, "VarUI8FromI8")
	varUI8FromR4 = doGetProcAddress(liboleaut32, "VarUI8FromR4")
	varUI8FromR8 = doGetProcAddress(liboleaut32, "VarUI8FromR8")
	varUI8FromStr = doGetProcAddress(liboleaut32, "VarUI8FromStr")
	varUI8FromUI1 = doGetProcAddress(liboleaut32, "VarUI8FromUI1")
	varUI8FromUI2 = doGetProcAddress(liboleaut32, "VarUI8FromUI2")
	varUI8FromUI4 = doGetProcAddress(liboleaut32, "VarUI8FromUI4")
	varUdateFromDate = doGetProcAddress(liboleaut32, "VarUdateFromDate")
	varWeekdayName = doGetProcAddress(liboleaut32, "VarWeekdayName")
	varXor = doGetProcAddress(liboleaut32, "VarXor")
	variantChangeType = doGetProcAddress(liboleaut32, "VariantChangeType")
	variantChangeTypeEx = doGetProcAddress(liboleaut32, "VariantChangeTypeEx")
	variantClear = doGetProcAddress(liboleaut32, "VariantClear")
	variantCopy = doGetProcAddress(liboleaut32, "VariantCopy")
	variantCopyInd = doGetProcAddress(liboleaut32, "VariantCopyInd")
	variantInit = doGetProcAddress(liboleaut32, "VariantInit")
	variantTimeToDosDateTime = doGetProcAddress(liboleaut32, "VariantTimeToDosDateTime")
	variantTimeToSystemTime = doGetProcAddress(liboleaut32, "VariantTimeToSystemTime")
	vectorFromBstr = doGetProcAddress(liboleaut32, "VectorFromBstr")
}

func BSTR_UserFree(pFlags *ULONG, pstr *BSTR) {
	syscall3(bSTR_UserFree, 2,
		uintptr(unsafe.Pointer(pFlags)),
		uintptr(unsafe.Pointer(pstr)),
		0)
}

func BSTR_UserSize(pFlags *ULONG, start ULONG, pstr *BSTR) ULONG {
	ret1 := syscall3(bSTR_UserSize, 3,
		uintptr(unsafe.Pointer(pFlags)),
		uintptr(start),
		uintptr(unsafe.Pointer(pstr)))
	return ULONG(ret1)
}

func BstrFromVector(psa *SAFEARRAY, pbstr *BSTR) HRESULT {
	ret1 := syscall3(bstrFromVector, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(pbstr)),
		0)
	return HRESULT(ret1)
}

func ClearCustData(lpCust *CUSTDATA) {
	syscall3(clearCustData, 1,
		uintptr(unsafe.Pointer(lpCust)),
		0,
		0)
}

func CreateDispTypeInfo(pidata *INTERFACEDATA, lcid LCID, pptinfo **ITypeInfo) HRESULT {
	ret1 := syscall3(createDispTypeInfo, 3,
		uintptr(unsafe.Pointer(pidata)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pptinfo)))
	return HRESULT(ret1)
}

func CreateErrorInfo(pperrinfo **ICreateErrorInfo) HRESULT {
	ret1 := syscall3(createErrorInfo, 1,
		uintptr(unsafe.Pointer(pperrinfo)),
		0,
		0)
	return HRESULT(ret1)
}

func CreateStdDispatch(punkOuter *IUnknown, pvThis uintptr, ptinfo *ITypeInfo, stddisp **IUnknown) HRESULT {
	ret1 := syscall6(createStdDispatch, 4,
		uintptr(unsafe.Pointer(punkOuter)),
		pvThis,
		uintptr(unsafe.Pointer(ptinfo)),
		uintptr(unsafe.Pointer(stddisp)),
		0,
		0)
	return HRESULT(ret1)
}

func CreateTypeLib(syskind SYSKIND, szFile /*const*/ LPCOLESTR, ppctlib **ICreateTypeLib) HRESULT {
	ret1 := syscall3(createTypeLib, 3,
		uintptr(syskind),
		uintptr(unsafe.Pointer(szFile)),
		uintptr(unsafe.Pointer(ppctlib)))
	return HRESULT(ret1)
}

func CreateTypeLib2(syskind SYSKIND, szFile /*const*/ LPCOLESTR, ppctlib **ICreateTypeLib2) HRESULT {
	ret1 := syscall3(createTypeLib2, 3,
		uintptr(syskind),
		uintptr(unsafe.Pointer(szFile)),
		uintptr(unsafe.Pointer(ppctlib)))
	return HRESULT(ret1)
}

func DispCallFunc(pvInstance uintptr, oVft *uint32, cc CALLCONV, vtReturn VARTYPE, cActuals uint32, prgvt *VARTYPE, prgpvarg **VARIANTARG, pvargResult *VARIANT) HRESULT {
	ret1 := syscall9(dispCallFunc, 8,
		pvInstance,
		uintptr(unsafe.Pointer(oVft)),
		uintptr(cc),
		uintptr(vtReturn),
		uintptr(cActuals),
		uintptr(unsafe.Pointer(prgvt)),
		uintptr(unsafe.Pointer(prgpvarg)),
		uintptr(unsafe.Pointer(pvargResult)),
		0)
	return HRESULT(ret1)
}

func DispGetIDsOfNames(ptinfo *ITypeInfo, rgszNames **OLECHAR, cNames uint32, rgdispid *DISPID) HRESULT {
	ret1 := syscall6(dispGetIDsOfNames, 4,
		uintptr(unsafe.Pointer(ptinfo)),
		uintptr(unsafe.Pointer(rgszNames)),
		uintptr(cNames),
		uintptr(unsafe.Pointer(rgdispid)),
		0,
		0)
	return HRESULT(ret1)
}

func DispGetParam(pdispparams *DISPPARAMS, position uint32, vtTarg VARTYPE, pvarResult *VARIANT, puArgErr *UINT) HRESULT {
	ret1 := syscall6(dispGetParam, 5,
		uintptr(unsafe.Pointer(pdispparams)),
		uintptr(position),
		uintptr(vtTarg),
		uintptr(unsafe.Pointer(pvarResult)),
		uintptr(unsafe.Pointer(puArgErr)),
		0)
	return HRESULT(ret1)
}

func DispInvoke(_this uintptr, ptinfo *ITypeInfo, dispidMember DISPID, wFlags USHORT, pparams *DISPPARAMS, pvarResult *VARIANT, pexcepinfo *EXCEPINFO, puArgErr *UINT) HRESULT {
	ret1 := syscall9(dispInvoke, 8,
		_this,
		uintptr(unsafe.Pointer(ptinfo)),
		uintptr(dispidMember),
		uintptr(wFlags),
		uintptr(unsafe.Pointer(pparams)),
		uintptr(unsafe.Pointer(pvarResult)),
		uintptr(unsafe.Pointer(pexcepinfo)),
		uintptr(unsafe.Pointer(puArgErr)),
		0)
	return HRESULT(ret1)
}

func DosDateTimeToVariantTime(wDosDate USHORT, wDosTime USHORT, pDateOut *float64) int32 {
	ret1 := syscall3(dosDateTimeToVariantTime, 3,
		uintptr(wDosDate),
		uintptr(wDosTime),
		uintptr(unsafe.Pointer(pDateOut)))
	return int32(ret1)
}

func GetErrorInfo(dwReserved ULONG, pperrinfo **IErrorInfo) HRESULT {
	ret1 := syscall3(getErrorInfo, 2,
		uintptr(dwReserved),
		uintptr(unsafe.Pointer(pperrinfo)),
		0)
	return HRESULT(ret1)
}

func GetRecordInfoFromGuids(rGuidTypeLib REFGUID, uVerMajor ULONG, uVerMinor ULONG, lcid LCID, rGuidTypeInfo REFGUID, ppRecInfo **IRecordInfo) HRESULT {
	ret1 := syscall6(getRecordInfoFromGuids, 6,
		uintptr(unsafe.Pointer(rGuidTypeLib)),
		uintptr(uVerMajor),
		uintptr(uVerMinor),
		uintptr(lcid),
		uintptr(unsafe.Pointer(rGuidTypeInfo)),
		uintptr(unsafe.Pointer(ppRecInfo)))
	return HRESULT(ret1)
}

func GetRecordInfoFromTypeInfo(pTI *ITypeInfo, ppRecInfo **IRecordInfo) HRESULT {
	ret1 := syscall3(getRecordInfoFromTypeInfo, 2,
		uintptr(unsafe.Pointer(pTI)),
		uintptr(unsafe.Pointer(ppRecInfo)),
		0)
	return HRESULT(ret1)
}

func LHashValOfNameSysA(skind SYSKIND, lcid LCID, lpStr /*const*/ LPCSTR) ULONG {
	ret1 := syscall3(lHashValOfNameSysA, 3,
		uintptr(skind),
		uintptr(lcid),
		uintptr(unsafe.Pointer(lpStr)))
	return ULONG(ret1)
}

func LPSAFEARRAY_UserFree(pFlags *ULONG, ppsa *LPSAFEARRAY) {
	syscall3(lPSAFEARRAY_UserFree, 2,
		uintptr(unsafe.Pointer(pFlags)),
		uintptr(unsafe.Pointer(ppsa)),
		0)
}

func LPSAFEARRAY_UserSize(pFlags *ULONG, startingSize ULONG, ppsa *LPSAFEARRAY) ULONG {
	ret1 := syscall3(lPSAFEARRAY_UserSize, 3,
		uintptr(unsafe.Pointer(pFlags)),
		uintptr(startingSize),
		uintptr(unsafe.Pointer(ppsa)))
	return ULONG(ret1)
}

func LoadRegTypeLib(rguid REFGUID, wVerMajor uint16, wVerMinor uint16, lcid LCID, ppTLib **ITypeLib) HRESULT {
	ret1 := syscall6(loadRegTypeLib, 5,
		uintptr(unsafe.Pointer(rguid)),
		uintptr(wVerMajor),
		uintptr(wVerMinor),
		uintptr(lcid),
		uintptr(unsafe.Pointer(ppTLib)),
		0)
	return HRESULT(ret1)
}

func LoadTypeLib(szFile /*const*/ *OLECHAR, pptLib **ITypeLib) HRESULT {
	ret1 := syscall3(loadTypeLib, 2,
		uintptr(unsafe.Pointer(szFile)),
		uintptr(unsafe.Pointer(pptLib)),
		0)
	return HRESULT(ret1)
}

func LoadTypeLibEx(szFile /*const*/ LPCOLESTR, regkind REGKIND, pptLib **ITypeLib) HRESULT {
	ret1 := syscall3(loadTypeLibEx, 3,
		uintptr(unsafe.Pointer(szFile)),
		uintptr(regkind),
		uintptr(unsafe.Pointer(pptLib)))
	return HRESULT(ret1)
}

func OaBuildVersion() ULONG {
	ret1 := syscall3(oaBuildVersion, 0,
		0,
		0,
		0)
	return ULONG(ret1)
}

func OleCreateFontIndirect(lpFontDesc *FONTDESC, riid REFIID, ppvObj *LPVOID) HRESULT {
	ret1 := syscall3(oleCreateFontIndirect, 3,
		uintptr(unsafe.Pointer(lpFontDesc)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObj)))
	return HRESULT(ret1)
}

func OleCreatePictureIndirect(lpPictDesc *PICTDESC, riid REFIID, own bool, ppvObj uintptr) HRESULT {
	ret1 := syscall6(oleCreatePictureIndirect, 4,
		uintptr(unsafe.Pointer(lpPictDesc)),
		uintptr(unsafe.Pointer(riid)),
		getUintptrFromBool(own),
		ppvObj,
		0,
		0)
	return HRESULT(ret1)
}

func OleCreatePropertyFrame(hwndOwner HWND, x uint32, y uint32, lpszCaption /*const*/ LPCOLESTR, cObjects ULONG, ppUnk *LPUNKNOWN, cPages ULONG, pPageClsID *CLSID, lcid LCID, dwReserved uint32, pvReserved LPVOID) HRESULT {
	ret1 := syscall12(oleCreatePropertyFrame, 11,
		uintptr(hwndOwner),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpszCaption)),
		uintptr(cObjects),
		uintptr(unsafe.Pointer(ppUnk)),
		uintptr(cPages),
		uintptr(unsafe.Pointer(pPageClsID)),
		uintptr(lcid),
		uintptr(dwReserved),
		uintptr(unsafe.Pointer(pvReserved)),
		0)
	return HRESULT(ret1)
}

func OleCreatePropertyFrameIndirect(lpParams *OCPFIPARAMS) HRESULT {
	ret1 := syscall3(oleCreatePropertyFrameIndirect, 1,
		uintptr(unsafe.Pointer(lpParams)),
		0,
		0)
	return HRESULT(ret1)
}

func OleIconToCursor(hinstExe HINSTANCE, hIcon HICON) HCURSOR {
	ret1 := syscall3(oleIconToCursor, 2,
		uintptr(hinstExe),
		uintptr(hIcon),
		0)
	return HCURSOR(ret1)
}

func OleLoadPicture(lpstream LPSTREAM, lSize LONG, fRunmode bool, riid REFIID, ppvObj *LPVOID) HRESULT {
	ret1 := syscall6(oleLoadPicture, 5,
		uintptr(unsafe.Pointer(lpstream)),
		uintptr(lSize),
		getUintptrFromBool(fRunmode),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObj)),
		0)
	return HRESULT(ret1)
}

func OleLoadPictureEx(lpstream LPSTREAM, lSize LONG, fRunmode bool, riid REFIID, xsiz uint32, ysiz uint32, flags uint32, ppvObj *LPVOID) HRESULT {
	ret1 := syscall9(oleLoadPictureEx, 8,
		uintptr(unsafe.Pointer(lpstream)),
		uintptr(lSize),
		getUintptrFromBool(fRunmode),
		uintptr(unsafe.Pointer(riid)),
		uintptr(xsiz),
		uintptr(ysiz),
		uintptr(flags),
		uintptr(unsafe.Pointer(ppvObj)),
		0)
	return HRESULT(ret1)
}

func OleLoadPictureFile(file VARIANT, picture *LPDISPATCH) HRESULT {
	args := unpackVARIANT(file)
	args = append(args, uintptr(unsafe.Pointer(picture)))
	ret := syscallN(oleLoadPictureFile, args)
	return HRESULT(ret)
}

func OleLoadPicturePath(szURLorPath LPOLESTR, punkCaller LPUNKNOWN, dwReserved uint32, clrReserved OLE_COLOR, riid REFIID, ppvRet *LPVOID) HRESULT {
	ret1 := syscall6(oleLoadPicturePath, 6,
		uintptr(unsafe.Pointer(szURLorPath)),
		uintptr(unsafe.Pointer(punkCaller)),
		uintptr(dwReserved),
		uintptr(clrReserved),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvRet)))
	return HRESULT(ret1)
}

func OleSavePictureFile(picture *IDispatch, filename BSTR) HRESULT {
	ret1 := syscall3(oleSavePictureFile, 2,
		uintptr(unsafe.Pointer(picture)),
		uintptr(unsafe.Pointer(filename)),
		0)
	return HRESULT(ret1)
}

func OleTranslateColor(clr OLE_COLOR, hpal HPALETTE, pColorRef *COLORREF) HRESULT {
	ret1 := syscall3(oleTranslateColor, 3,
		uintptr(clr),
		uintptr(hpal),
		uintptr(unsafe.Pointer(pColorRef)))
	return HRESULT(ret1)
}

func QueryPathOfRegTypeLib(guid REFGUID, wMaj uint16, wMin uint16, lcid LCID, path *BSTR) HRESULT {
	ret1 := syscall6(queryPathOfRegTypeLib, 5,
		uintptr(unsafe.Pointer(guid)),
		uintptr(wMaj),
		uintptr(wMin),
		uintptr(lcid),
		uintptr(unsafe.Pointer(path)),
		0)
	return HRESULT(ret1)
}

func RegisterTypeLib(ptlib *ITypeLib, szFullPath *OLECHAR, szHelpDir *OLECHAR) HRESULT {
	ret1 := syscall3(registerTypeLib, 3,
		uintptr(unsafe.Pointer(ptlib)),
		uintptr(unsafe.Pointer(szFullPath)),
		uintptr(unsafe.Pointer(szHelpDir)))
	return HRESULT(ret1)
}

func RegisterTypeLibForUser(ptlib *ITypeLib, szFullPath *OLECHAR, szHelpDir *OLECHAR) HRESULT {
	ret1 := syscall3(registerTypeLibForUser, 3,
		uintptr(unsafe.Pointer(ptlib)),
		uintptr(unsafe.Pointer(szFullPath)),
		uintptr(unsafe.Pointer(szHelpDir)))
	return HRESULT(ret1)
}

func SafeArrayAccessData(psa *SAFEARRAY, ppvData uintptr) HRESULT {
	ret1 := syscall3(safeArrayAccessData, 2,
		uintptr(unsafe.Pointer(psa)),
		ppvData,
		0)
	return HRESULT(ret1)
}

func SafeArrayAllocData(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayAllocData, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SafeArrayAllocDescriptor(cDims uint32, ppsaOut **SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayAllocDescriptor, 2,
		uintptr(cDims),
		uintptr(unsafe.Pointer(ppsaOut)),
		0)
	return HRESULT(ret1)
}

func SafeArrayAllocDescriptorEx(vt VARTYPE, cDims uint32, ppsaOut **SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayAllocDescriptorEx, 3,
		uintptr(vt),
		uintptr(cDims),
		uintptr(unsafe.Pointer(ppsaOut)))
	return HRESULT(ret1)
}

func SafeArrayCopy(psa *SAFEARRAY, ppsaOut **SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayCopy, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(ppsaOut)),
		0)
	return HRESULT(ret1)
}

func SafeArrayCopyData(psaSource *SAFEARRAY, psaTarget *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayCopyData, 2,
		uintptr(unsafe.Pointer(psaSource)),
		uintptr(unsafe.Pointer(psaTarget)),
		0)
	return HRESULT(ret1)
}

func SafeArrayDestroy(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayDestroy, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SafeArrayDestroyData(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayDestroyData, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SafeArrayDestroyDescriptor(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayDestroyDescriptor, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SafeArrayGetDim(psa *SAFEARRAY) uint32 {
	ret1 := syscall3(safeArrayGetDim, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return uint32(ret1)
}

func SafeArrayGetElement(psa *SAFEARRAY, rgIndices *LONG, pvData uintptr) HRESULT {
	ret1 := syscall3(safeArrayGetElement, 3,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(rgIndices)),
		pvData)
	return HRESULT(ret1)
}

func SafeArrayGetElemsize(psa *SAFEARRAY) uint32 {
	ret1 := syscall3(safeArrayGetElemsize, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return uint32(ret1)
}

func SafeArrayGetIID(psa *SAFEARRAY, pGuid *GUID) HRESULT {
	ret1 := syscall3(safeArrayGetIID, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(pGuid)),
		0)
	return HRESULT(ret1)
}

func SafeArrayGetLBound(psa *SAFEARRAY, nDim uint32, plLbound *LONG) HRESULT {
	ret1 := syscall3(safeArrayGetLBound, 3,
		uintptr(unsafe.Pointer(psa)),
		uintptr(nDim),
		uintptr(unsafe.Pointer(plLbound)))
	return HRESULT(ret1)
}

func SafeArrayGetRecordInfo(psa *SAFEARRAY, pRinfo **IRecordInfo) HRESULT {
	ret1 := syscall3(safeArrayGetRecordInfo, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(pRinfo)),
		0)
	return HRESULT(ret1)
}

func SafeArrayGetUBound(psa *SAFEARRAY, nDim uint32, plUbound *LONG) HRESULT {
	ret1 := syscall3(safeArrayGetUBound, 3,
		uintptr(unsafe.Pointer(psa)),
		uintptr(nDim),
		uintptr(unsafe.Pointer(plUbound)))
	return HRESULT(ret1)
}

func SafeArrayGetVartype(psa *SAFEARRAY, pvt *VARTYPE) HRESULT {
	ret1 := syscall3(safeArrayGetVartype, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(pvt)),
		0)
	return HRESULT(ret1)
}

func SafeArrayLock(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayLock, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SafeArrayPtrOfIndex(psa *SAFEARRAY, rgIndices *LONG, ppvData uintptr) HRESULT {
	ret1 := syscall3(safeArrayPtrOfIndex, 3,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(rgIndices)),
		ppvData)
	return HRESULT(ret1)
}

func SafeArrayPutElement(psa *SAFEARRAY, rgIndices *LONG, pvData uintptr) HRESULT {
	ret1 := syscall3(safeArrayPutElement, 3,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(rgIndices)),
		pvData)
	return HRESULT(ret1)
}

func SafeArrayRedim(psa *SAFEARRAY, psabound *SAFEARRAYBOUND) HRESULT {
	ret1 := syscall3(safeArrayRedim, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(psabound)),
		0)
	return HRESULT(ret1)
}

func SafeArraySetIID(psa *SAFEARRAY, guid REFGUID) HRESULT {
	ret1 := syscall3(safeArraySetIID, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(guid)),
		0)
	return HRESULT(ret1)
}

func SafeArraySetRecordInfo(psa *SAFEARRAY, pRinfo *IRecordInfo) HRESULT {
	ret1 := syscall3(safeArraySetRecordInfo, 2,
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(pRinfo)),
		0)
	return HRESULT(ret1)
}

func SafeArrayUnaccessData(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayUnaccessData, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SafeArrayUnlock(psa *SAFEARRAY) HRESULT {
	ret1 := syscall3(safeArrayUnlock, 1,
		uintptr(unsafe.Pointer(psa)),
		0,
		0)
	return HRESULT(ret1)
}

func SetErrorInfo(dwReserved ULONG, perrinfo *IErrorInfo) HRESULT {
	ret1 := syscall3(setErrorInfo, 2,
		uintptr(dwReserved),
		uintptr(unsafe.Pointer(perrinfo)),
		0)
	return HRESULT(ret1)
}

func SetOaNoCache() {
	syscall3(setOaNoCache, 0,
		0,
		0,
		0)
}

func SysAllocString(str /*const*/ LPCOLESTR) BSTR {
	ret1 := syscall3(sysAllocString, 1,
		uintptr(unsafe.Pointer(str)),
		0,
		0)
	return (BSTR)(unsafe.Pointer(ret1))
}

func SysAllocStringByteLen(str /*const*/ LPCSTR, aLen uint32) BSTR {
	ret1 := syscall3(sysAllocStringByteLen, 2,
		uintptr(unsafe.Pointer(str)),
		uintptr(aLen),
		0)
	return (BSTR)(unsafe.Pointer(ret1))
}

func SysAllocStringLen(str /*const*/ *OLECHAR, aLen uint32) BSTR {
	ret1 := syscall3(sysAllocStringLen, 2,
		uintptr(unsafe.Pointer(str)),
		uintptr(aLen),
		0)
	return (BSTR)(unsafe.Pointer(ret1))
}

func SysFreeString(str BSTR) {
	syscall3(sysFreeString, 1,
		uintptr(unsafe.Pointer(str)),
		0,
		0)
}

func SysReAllocString(old *BSTR, str /*const*/ LPCOLESTR) int32 {
	ret1 := syscall3(sysReAllocString, 2,
		uintptr(unsafe.Pointer(old)),
		uintptr(unsafe.Pointer(str)),
		0)
	return int32(ret1)
}

func SysReAllocStringLen(old *BSTR, str /*const*/ *OLECHAR, aLen uint32) int32 {
	ret1 := syscall3(sysReAllocStringLen, 3,
		uintptr(unsafe.Pointer(old)),
		uintptr(unsafe.Pointer(str)),
		uintptr(aLen))
	return int32(ret1)
}

func SysStringByteLen(str BSTR) uint32 {
	ret1 := syscall3(sysStringByteLen, 1,
		uintptr(unsafe.Pointer(str)),
		0,
		0)
	return uint32(ret1)
}

func SysStringLen(str BSTR) uint32 {
	ret1 := syscall3(sysStringLen, 1,
		uintptr(unsafe.Pointer(str)),
		0,
		0)
	return uint32(ret1)
}

func SystemTimeToVariantTime(lpSt *SYSTEMTIME, pDateOut *float64) int32 {
	ret1 := syscall3(systemTimeToVariantTime, 2,
		uintptr(unsafe.Pointer(lpSt)),
		uintptr(unsafe.Pointer(pDateOut)),
		0)
	return int32(ret1)
}

func UnRegisterTypeLib(libid REFGUID, wVerMajor uint16, wVerMinor uint16, lcid LCID, syskind SYSKIND) HRESULT {
	ret1 := syscall6(unRegisterTypeLib, 5,
		uintptr(unsafe.Pointer(libid)),
		uintptr(wVerMajor),
		uintptr(wVerMinor),
		uintptr(lcid),
		uintptr(syskind),
		0)
	return HRESULT(ret1)
}

func UnRegisterTypeLibForUser(libid REFGUID, wVerMajor uint16, wVerMinor uint16, lcid LCID, syskind SYSKIND) HRESULT {
	ret1 := syscall6(unRegisterTypeLibForUser, 5,
		uintptr(unsafe.Pointer(libid)),
		uintptr(wVerMajor),
		uintptr(wVerMinor),
		uintptr(lcid),
		uintptr(syskind),
		0)
	return HRESULT(ret1)
}

func VARIANT_UserFree(pFlags *ULONG, pvar *VARIANT) {
	syscall3(vARIANT_UserFree, 2,
		uintptr(unsafe.Pointer(pFlags)),
		uintptr(unsafe.Pointer(pvar)),
		0)
}

func VARIANT_UserSize(pFlags *ULONG, start ULONG, pvar *VARIANT) ULONG {
	ret1 := syscall3(vARIANT_UserSize, 3,
		uintptr(unsafe.Pointer(pFlags)),
		uintptr(start),
		uintptr(unsafe.Pointer(pvar)))
	return ULONG(ret1)
}

func VarAbs(pVarIn *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varAbs, 2,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(pVarOut)),
		0)
	return HRESULT(ret1)
}

func VarAdd(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varAdd, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarAnd(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varAnd, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarBoolFromCy(cyIn CY, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pBoolOut)))
	return HRESULT(ret1)
}

func VarBoolFromDate(dateIn DATE, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromDec(pDecIn *DECIMAL, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromDec, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromDisp(pdispIn *IDispatch, lcid LCID, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pBoolOut)))
	return HRESULT(ret1)
}

func VarBoolFromI1(cIn int8, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromI2(sIn int16, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromI4(lIn LONG, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromI4, 2,
		uintptr(lIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromI8(llIn LONG64, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromR4(fltIn float32, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromR8(dblIn float64, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall6(varBoolFromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pBoolOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBoolFromUI1(bIn byte, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromUI2(usIn USHORT, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromUI4(ulIn ULONG, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBoolFromUI8(ullIn ULONG64, pBoolOut *VARIANT_BOOL) HRESULT {
	ret1 := syscall3(varBoolFromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pBoolOut)),
		0)
	return HRESULT(ret1)
}

func VarBstrCat(pbstrLeft BSTR, pbstrRight BSTR, pbstrOut *BSTR) HRESULT {
	ret1 := syscall3(varBstrCat, 3,
		uintptr(unsafe.Pointer(pbstrLeft)),
		uintptr(unsafe.Pointer(pbstrRight)),
		uintptr(unsafe.Pointer(pbstrOut)))
	return HRESULT(ret1)
}

func VarBstrCmp(pbstrLeft BSTR, pbstrRight BSTR, lcid LCID, dwFlags uint32) HRESULT {
	ret1 := syscall6(varBstrCmp, 4,
		uintptr(unsafe.Pointer(pbstrLeft)),
		uintptr(unsafe.Pointer(pbstrRight)),
		uintptr(lcid),
		uintptr(dwFlags),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromBool(boolIn VARIANT_BOOL, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromBool, 4,
		uintptr(boolIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromCy(cyIn CY, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromCy, 5,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0)
	return HRESULT(ret1)
}

func VarBstrFromDate(dateIn DATE, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromDate, 4,
		uintptr(dateIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromDec(pDecIn *DECIMAL, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromDec, 4,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromDisp(pdispIn *IDispatch, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromDisp, 4,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromI1(cIn int8, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromI1, 4,
		uintptr(cIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromI2(sIn int16, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromI2, 4,
		uintptr(sIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromI4(lIn LONG, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromI4, 4,
		uintptr(lIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromI8(llIn LONG64, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromI8, 4,
		uintptr(llIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromR4(fltIn float32, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromR4, 4,
		uintptr(fltIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromR8(dblIn float64, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromR8, 4,
		uintptr(dblIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromUI1(bIn byte, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromUI1, 4,
		uintptr(bIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromUI2(usIn USHORT, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromUI2, 4,
		uintptr(usIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromUI4(ulIn ULONG, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromUI4, 4,
		uintptr(ulIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarBstrFromUI8(ullIn ULONG64, lcid LCID, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varBstrFromUI8, 4,
		uintptr(ullIn),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarCat(left *VARIANT, right *VARIANT, out *VARIANT) HRESULT {
	ret1 := syscall3(varCat, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(out)))
	return HRESULT(ret1)
}

func VarCmp(left *VARIANT, right *VARIANT, lcid LCID, flags uint32) HRESULT {
	ret1 := syscall6(varCmp, 4,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(lcid),
		uintptr(flags),
		0,
		0)
	return HRESULT(ret1)
}

func VarCyAbs(cyIn /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyAbs, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)))
	return HRESULT(ret1)
}

func VarCyAdd(cyLeft /*const*/ CY, cyRight /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCyAdd, 5,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(*(*uint32)(unsafe.Pointer(&cyRight))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyRight)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyCmp(cyLeft /*const*/ CY, cyRight /*const*/ CY) HRESULT {
	ret1 := syscall6(varCyCmp, 4,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(*(*uint32)(unsafe.Pointer(&cyRight))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyRight)) + uintptr(4)))),
		0,
		0)
	return HRESULT(ret1)
}

func VarCyCmpR8(cyLeft /*const*/ CY, dblRight float64) HRESULT {
	ret1 := syscall3(varCyCmpR8, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(dblRight))
	return HRESULT(ret1)
}

func VarCyFix(cyIn /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFix, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)))
	return HRESULT(ret1)
}

func VarCyFromBool(boolIn VARIANT_BOOL, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromDate(dateIn DATE, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromDec(pdecIn *DECIMAL, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromDisp(pdispIn *IDispatch, lcid LCID, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pCyOut)))
	return HRESULT(ret1)
}

func VarCyFromI1(cIn int8, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromI2(sIn int16, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromI4(lIn LONG, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromI4, 2,
		uintptr(lIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromI8(llIn LONG64, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromR4(fltIn float32, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromR8(dblIn float64, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCyFromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pCyOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarCyFromUI1(bIn byte, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromUI2(usIn USHORT, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromUI4(ulIn ULONG, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyFromUI8(ullIn ULONG64, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyFromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyInt(cyIn /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyInt, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)))
	return HRESULT(ret1)
}

func VarCyMul(cyLeft /*const*/ CY, cyRight /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCyMul, 5,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(*(*uint32)(unsafe.Pointer(&cyRight))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyRight)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarCyMulI4(cyLeft /*const*/ CY, lRight LONG, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCyMulI4, 4,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(lRight),
		uintptr(unsafe.Pointer(pCyOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarCyMulI8(cyLeft /*const*/ CY, llRight LONG64, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCyMulI8, 4,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(llRight),
		uintptr(unsafe.Pointer(pCyOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarCyNeg(cyIn /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall3(varCyNeg, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)))
	return HRESULT(ret1)
}

func VarCyRound(cyIn /*const*/ CY, cDecimals int32, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCyRound, 4,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(cDecimals),
		uintptr(unsafe.Pointer(pCyOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarCySub(cyLeft /*const*/ CY, cyRight /*const*/ CY, pCyOut *CY) HRESULT {
	ret1 := syscall6(varCySub, 5,
		uintptr(*(*uint32)(unsafe.Pointer(&cyLeft))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyLeft)) + uintptr(4)))),
		uintptr(*(*uint32)(unsafe.Pointer(&cyRight))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyRight)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pCyOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromBool(boolIn VARIANT_BOOL, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromCy(cyIn CY, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pdateOut)))
	return HRESULT(ret1)
}

func VarDateFromDec(pdecIn *DECIMAL, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromDisp(pdispIn *IDispatch, lcid LCID, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pdateOut)))
	return HRESULT(ret1)
}

func VarDateFromI1(cIn int8, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromI2(sIn int16, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromI4(lIn LONG, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromI4, 2,
		uintptr(lIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromI8(llIn LONG64, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromR4(fltIn float32, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromR8(dblIn float64, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pdateOut *DATE) HRESULT {
	ret1 := syscall6(varDateFromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pdateOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarDateFromUI1(bIn byte, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromUI2(uiIn USHORT, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromUI2, 2,
		uintptr(uiIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromUI4(ulIn ULONG, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromUI8(ullIn ULONG64, pdateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pdateOut)),
		0)
	return HRESULT(ret1)
}

func VarDateFromUdate(pUdateIn *UDATE, dwFlags ULONG, pDateOut *DATE) HRESULT {
	ret1 := syscall3(varDateFromUdate, 3,
		uintptr(unsafe.Pointer(pUdateIn)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pDateOut)))
	return HRESULT(ret1)
}

func VarDateFromUdateEx(pUdateIn *UDATE, lcid LCID, dwFlags ULONG, pDateOut *DATE) HRESULT {
	ret1 := syscall6(varDateFromUdateEx, 4,
		uintptr(unsafe.Pointer(pUdateIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pDateOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarDecAbs(pDecIn /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecAbs, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecAdd(pDecLeft /*const*/ *DECIMAL, pDecRight /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecAdd, 3,
		uintptr(unsafe.Pointer(pDecLeft)),
		uintptr(unsafe.Pointer(pDecRight)),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDecCmp(pDecLeft /*const*/ *DECIMAL, pDecRight /*const*/ *DECIMAL) HRESULT {
	ret1 := syscall3(varDecCmp, 2,
		uintptr(unsafe.Pointer(pDecLeft)),
		uintptr(unsafe.Pointer(pDecRight)),
		0)
	return HRESULT(ret1)
}

func VarDecCmpR8(pDecLeft /*const*/ *DECIMAL, dblRight float64) HRESULT {
	ret1 := syscall3(varDecCmpR8, 2,
		uintptr(unsafe.Pointer(pDecLeft)),
		uintptr(dblRight),
		0)
	return HRESULT(ret1)
}

func VarDecDiv(pDecLeft /*const*/ *DECIMAL, pDecRight /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecDiv, 3,
		uintptr(unsafe.Pointer(pDecLeft)),
		uintptr(unsafe.Pointer(pDecRight)),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDecFix(pDecIn /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFix, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromBool(bIn VARIANT_BOOL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromBool, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromCy(cyIn CY, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDecFromDate(dateIn DATE, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromDisp(pdispIn *IDispatch, lcid LCID, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDecFromI1(cIn int8, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromI2(sIn int16, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromI4(lIn LONG, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromI4, 2,
		uintptr(lIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromI8(llIn LONG64, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromR4(fltIn float32, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromR8(dblIn float64, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall6(varDecFromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pDecOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarDecFromUI1(bIn byte, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromUI2(usIn USHORT, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromUI4(ulIn ULONG, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecFromUI8(ullIn ULONG64, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecFromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecInt(pDecIn /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecInt, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecMul(pDecLeft /*const*/ *DECIMAL, pDecRight /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecMul, 3,
		uintptr(unsafe.Pointer(pDecLeft)),
		uintptr(unsafe.Pointer(pDecRight)),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDecNeg(pDecIn /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecNeg, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pDecOut)),
		0)
	return HRESULT(ret1)
}

func VarDecRound(pDecIn /*const*/ *DECIMAL, cDecimals int32, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecRound, 3,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(cDecimals),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDecSub(pDecLeft /*const*/ *DECIMAL, pDecRight /*const*/ *DECIMAL, pDecOut *DECIMAL) HRESULT {
	ret1 := syscall3(varDecSub, 3,
		uintptr(unsafe.Pointer(pDecLeft)),
		uintptr(unsafe.Pointer(pDecRight)),
		uintptr(unsafe.Pointer(pDecOut)))
	return HRESULT(ret1)
}

func VarDiv(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varDiv, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarEqv(pVarLeft *VARIANT, pVarRight *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varEqv, 3,
		uintptr(unsafe.Pointer(pVarLeft)),
		uintptr(unsafe.Pointer(pVarRight)),
		uintptr(unsafe.Pointer(pVarOut)))
	return HRESULT(ret1)
}

func VarFix(pVarIn *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varFix, 2,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(pVarOut)),
		0)
	return HRESULT(ret1)
}

func VarFormat(pVarIn *VARIANT, lpszFormat LPOLESTR, nFirstDay int32, nFirstWeek int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varFormat, 6,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(lpszFormat)),
		uintptr(nFirstDay),
		uintptr(nFirstWeek),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)))
	return HRESULT(ret1)
}

func VarFormatCurrency(pVarIn *VARIANT, nDigits int32, nLeading int32, nParens int32, nGrouping int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall9(varFormatCurrency, 7,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(nDigits),
		uintptr(nLeading),
		uintptr(nParens),
		uintptr(nGrouping),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarFormatDateTime(pVarIn *VARIANT, nFormat int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varFormatDateTime, 4,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(nFormat),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarFormatFromTokens(pVarIn *VARIANT, lpszFormat LPOLESTR, rgbTok *byte, dwFlags ULONG, pbstrOut *BSTR, lcid LCID) HRESULT {
	ret1 := syscall6(varFormatFromTokens, 6,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(lpszFormat)),
		uintptr(unsafe.Pointer(rgbTok)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		uintptr(lcid))
	return HRESULT(ret1)
}

func VarFormatNumber(pVarIn *VARIANT, nDigits int32, nLeading int32, nParens int32, nGrouping int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall9(varFormatNumber, 7,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(nDigits),
		uintptr(nLeading),
		uintptr(nParens),
		uintptr(nGrouping),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarFormatPercent(pVarIn *VARIANT, nDigits int32, nLeading int32, nParens int32, nGrouping int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall9(varFormatPercent, 7,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(nDigits),
		uintptr(nLeading),
		uintptr(nParens),
		uintptr(nGrouping),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarI1FromBool(boolIn VARIANT_BOOL, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromCy(cyIn CY, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pcOut)))
	return HRESULT(ret1)
}

func VarI1FromDate(dateIn DATE, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromDec(pdecIn *DECIMAL, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromDisp(pdispIn *IDispatch, lcid LCID, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pcOut)))
	return HRESULT(ret1)
}

func VarI1FromI2(sIn int16, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromI4(iIn LONG, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromI4, 2,
		uintptr(iIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromI8(llIn LONG64, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromR4(fltIn float32, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromR8(dblIn float64, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pcOut *int8) HRESULT {
	ret1 := syscall6(varI1FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pcOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarI1FromUI1(bIn byte, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromUI2(usIn USHORT, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromUI4(ulIn ULONG, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI1FromUI8(ullIn ULONG64, pcOut *int8) HRESULT {
	ret1 := syscall3(varI1FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pcOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromBool(boolIn VARIANT_BOOL, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromCy(cyIn CY, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(psOut)))
	return HRESULT(ret1)
}

func VarI2FromDate(dateIn DATE, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromDec(pdecIn *DECIMAL, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromDisp(pdispIn *IDispatch, lcid LCID, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(psOut)))
	return HRESULT(ret1)
}

func VarI2FromI1(cIn int8, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromI4(iIn LONG, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromI4, 2,
		uintptr(iIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromI8(llIn LONG64, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromR4(fltIn float32, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromR8(dblIn float64, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, psOut *SHORT) HRESULT {
	ret1 := syscall6(varI2FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(psOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarI2FromUI1(bIn byte, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromUI2(usIn USHORT, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromUI4(ulIn ULONG, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI2FromUI8(ullIn ULONG64, psOut *SHORT) HRESULT {
	ret1 := syscall3(varI2FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(psOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromBool(boolIn VARIANT_BOOL, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromCy(cyIn CY, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(piOut)))
	return HRESULT(ret1)
}

func VarI4FromDate(dateIn DATE, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromDec(pdecIn *DECIMAL, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromDisp(pdispIn *IDispatch, lcid LCID, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(piOut)))
	return HRESULT(ret1)
}

func VarI4FromI1(cIn int8, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromI2(sIn int16, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromI8(llIn LONG64, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromR4(fltIn float32, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromR8(dblIn float64, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, piOut *LONG) HRESULT {
	ret1 := syscall6(varI4FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(piOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarI4FromUI1(bIn byte, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromUI2(usIn USHORT, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromUI4(ulIn ULONG, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI4FromUI8(ullIn ULONG64, piOut *LONG) HRESULT {
	ret1 := syscall3(varI4FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(piOut)),
		0)
	return HRESULT(ret1)
}

func VarI8FromBool(boolIn VARIANT_BOOL, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromCy(cyIn CY, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pi64Out)))
	return HRESULT(ret1)
}

func VarI8FromDate(dateIn DATE, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromDec(pdecIn *DECIMAL, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromDisp(pdispIn *IDispatch, lcid LCID, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pi64Out)))
	return HRESULT(ret1)
}

func VarI8FromI1(cIn int8, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromI2(sIn int16, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromR4(fltIn float32, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromR8(dblIn float64, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pi64Out *LONG64) HRESULT {
	ret1 := syscall6(varI8FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pi64Out)),
		0,
		0)
	return HRESULT(ret1)
}

func VarI8FromUI1(bIn byte, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromUI2(usIn USHORT, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromUI4(ulIn ULONG, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarI8FromUI8(ullIn ULONG64, pi64Out *LONG64) HRESULT {
	ret1 := syscall3(varI8FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pi64Out)),
		0)
	return HRESULT(ret1)
}

func VarIdiv(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varIdiv, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarImp(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varImp, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarInt(pVarIn *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varInt, 2,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(pVarOut)),
		0)
	return HRESULT(ret1)
}

func VarMod(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varMod, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarMonthName(iMonth int32, fAbbrev int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varMonthName, 4,
		uintptr(iMonth),
		uintptr(fAbbrev),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarMul(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varMul, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarNeg(pVarIn *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varNeg, 2,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(pVarOut)),
		0)
	return HRESULT(ret1)
}

func VarNot(pVarIn *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varNot, 2,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(unsafe.Pointer(pVarOut)),
		0)
	return HRESULT(ret1)
}

func VarNumFromParseNum(pNumprs *NUMPARSE, rgbDig *byte, dwVtBits ULONG, pVarDst *VARIANT) HRESULT {
	ret1 := syscall6(varNumFromParseNum, 4,
		uintptr(unsafe.Pointer(pNumprs)),
		uintptr(unsafe.Pointer(rgbDig)),
		uintptr(dwVtBits),
		uintptr(unsafe.Pointer(pVarDst)),
		0,
		0)
	return HRESULT(ret1)
}

func VarOr(pVarLeft *VARIANT, pVarRight *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varOr, 3,
		uintptr(unsafe.Pointer(pVarLeft)),
		uintptr(unsafe.Pointer(pVarRight)),
		uintptr(unsafe.Pointer(pVarOut)))
	return HRESULT(ret1)
}

func VarParseNumFromStr(lpszStr *OLECHAR, lcid LCID, dwFlags ULONG, pNumprs *NUMPARSE, rgbDig *byte) HRESULT {
	ret1 := syscall6(varParseNumFromStr, 5,
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pNumprs)),
		uintptr(unsafe.Pointer(rgbDig)),
		0)
	return HRESULT(ret1)
}

func VarPow(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varPow, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarR4CmpR8(fltLeft float32, dblRight float64) HRESULT {
	ret1 := syscall3(varR4CmpR8, 2,
		uintptr(fltLeft),
		uintptr(dblRight),
		0)
	return HRESULT(ret1)
}

func VarR4FromBool(boolIn VARIANT_BOOL, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromCy(cyIn CY, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pFltOut)))
	return HRESULT(ret1)
}

func VarR4FromDate(dateIn DATE, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromDec(pDecIn *DECIMAL, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromDec, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromDisp(pdispIn *IDispatch, lcid LCID, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pFltOut)))
	return HRESULT(ret1)
}

func VarR4FromI1(cIn int8, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromI2(sIn int16, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromI4(lIn LONG, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromI4, 2,
		uintptr(lIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromI8(llIn LONG64, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromR8(dblIn float64, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pFltOut *float32) HRESULT {
	ret1 := syscall6(varR4FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pFltOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarR4FromUI1(bIn byte, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromUI2(usIn USHORT, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromUI4(ulIn ULONG, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR4FromUI8(ullIn ULONG64, pFltOut *float32) HRESULT {
	ret1 := syscall3(varR4FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pFltOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromBool(boolIn VARIANT_BOOL, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromCy(cyIn CY, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pDblOut)))
	return HRESULT(ret1)
}

func VarR8FromDate(dateIn DATE, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromDec(pDecIn /*const*/ *DECIMAL, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromDec, 2,
		uintptr(unsafe.Pointer(pDecIn)),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromDisp(pdispIn *IDispatch, lcid LCID, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pDblOut)))
	return HRESULT(ret1)
}

func VarR8FromI1(cIn int8, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromI2(sIn int16, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromI4(lIn LONG, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromI4, 2,
		uintptr(lIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromI8(llIn LONG64, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromR4(fltIn float32, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pDblOut *float64) HRESULT {
	ret1 := syscall6(varR8FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pDblOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarR8FromUI1(bIn byte, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromUI2(usIn USHORT, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromUI4(ulIn ULONG, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8FromUI8(ullIn ULONG64, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pDblOut)),
		0)
	return HRESULT(ret1)
}

func VarR8Pow(dblLeft float64, dblPow float64, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8Pow, 3,
		uintptr(dblLeft),
		uintptr(dblPow),
		uintptr(unsafe.Pointer(pDblOut)))
	return HRESULT(ret1)
}

func VarR8Round(dblIn float64, nDig int32, pDblOut *float64) HRESULT {
	ret1 := syscall3(varR8Round, 3,
		uintptr(dblIn),
		uintptr(nDig),
		uintptr(unsafe.Pointer(pDblOut)))
	return HRESULT(ret1)
}

func VarRound(pVarIn *VARIANT, deci int32, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varRound, 3,
		uintptr(unsafe.Pointer(pVarIn)),
		uintptr(deci),
		uintptr(unsafe.Pointer(pVarOut)))
	return HRESULT(ret1)
}

func VarSub(left *VARIANT, right *VARIANT, result *VARIANT) HRESULT {
	ret1 := syscall3(varSub, 3,
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(result)))
	return HRESULT(ret1)
}

func VarTokenizeFormatString(lpszFormat LPOLESTR, rgbTok *byte, cbTok int32, nFirstDay int32, nFirstWeek int32, lcid LCID, pcbActual *int) HRESULT {
	ret1 := syscall9(varTokenizeFormatString, 7,
		uintptr(unsafe.Pointer(lpszFormat)),
		uintptr(unsafe.Pointer(rgbTok)),
		uintptr(cbTok),
		uintptr(nFirstDay),
		uintptr(nFirstWeek),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pcbActual)),
		0,
		0)
	return HRESULT(ret1)
}

func VarUI1FromBool(boolIn VARIANT_BOOL, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromCy(cyIn CY, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pbOut)))
	return HRESULT(ret1)
}

func VarUI1FromDate(dateIn DATE, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromDec(pdecIn *DECIMAL, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromDisp(pdispIn *IDispatch, lcid LCID, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pbOut)))
	return HRESULT(ret1)
}

func VarUI1FromI1(cIn int8, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromI2(sIn int16, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromI4(iIn LONG, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromI4, 2,
		uintptr(iIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromI8(llIn LONG64, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromR4(fltIn float32, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromR8(dblIn float64, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pbOut *byte) HRESULT {
	ret1 := syscall6(varUI1FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarUI1FromUI2(usIn USHORT, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromUI4(ulIn ULONG, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI1FromUI8(ullIn ULONG64, pbOut *byte) HRESULT {
	ret1 := syscall3(varUI1FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pbOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromBool(boolIn VARIANT_BOOL, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromCy(cyIn CY, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pusOut)))
	return HRESULT(ret1)
}

func VarUI2FromDate(dateIn DATE, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromDec(pdecIn *DECIMAL, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromDisp(pdispIn *IDispatch, lcid LCID, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pusOut)))
	return HRESULT(ret1)
}

func VarUI2FromI1(cIn int8, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromI2(sIn int16, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromI4(iIn LONG, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromI4, 2,
		uintptr(iIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromI8(llIn LONG64, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromR4(fltIn float32, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromR8(dblIn float64, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pusOut *USHORT) HRESULT {
	ret1 := syscall6(varUI2FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pusOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarUI2FromUI1(bIn byte, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromUI4(ulIn ULONG, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI2FromUI8(ullIn ULONG64, pusOut *USHORT) HRESULT {
	ret1 := syscall3(varUI2FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pusOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromBool(boolIn VARIANT_BOOL, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromCy(cyIn CY, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pulOut)))
	return HRESULT(ret1)
}

func VarUI4FromDate(dateIn DATE, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromDec(pdecIn *DECIMAL, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromDisp(pdispIn *IDispatch, lcid LCID, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pulOut)))
	return HRESULT(ret1)
}

func VarUI4FromI1(cIn int8, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromI2(sIn int16, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromI4(iIn LONG, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromI4, 2,
		uintptr(iIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromI8(llIn LONG64, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromR4(fltIn float32, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromR8(dblIn float64, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pulOut *ULONG) HRESULT {
	ret1 := syscall6(varUI4FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pulOut)),
		0,
		0)
	return HRESULT(ret1)
}

func VarUI4FromUI1(bIn byte, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromUI2(usIn USHORT, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI4FromUI8(ullIn ULONG64, pulOut *ULONG) HRESULT {
	ret1 := syscall3(varUI4FromUI8, 2,
		uintptr(ullIn),
		uintptr(unsafe.Pointer(pulOut)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromBool(boolIn VARIANT_BOOL, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromBool, 2,
		uintptr(boolIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromCy(cyIn CY, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromCy, 3,
		uintptr(*(*uint32)(unsafe.Pointer(&cyIn))),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cyIn)) + uintptr(4)))),
		uintptr(unsafe.Pointer(pui64Out)))
	return HRESULT(ret1)
}

func VarUI8FromDate(dateIn DATE, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromDate, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromDec(pdecIn *DECIMAL, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromDec, 2,
		uintptr(unsafe.Pointer(pdecIn)),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromDisp(pdispIn *IDispatch, lcid LCID, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromDisp, 3,
		uintptr(unsafe.Pointer(pdispIn)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(pui64Out)))
	return HRESULT(ret1)
}

func VarUI8FromI1(cIn int8, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromI1, 2,
		uintptr(cIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromI2(sIn int16, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromI2, 2,
		uintptr(sIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromI8(llIn LONG64, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromI8, 2,
		uintptr(llIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromR4(fltIn float32, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromR4, 2,
		uintptr(fltIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromR8(dblIn float64, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromR8, 2,
		uintptr(dblIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromStr(strIn *OLECHAR, lcid LCID, dwFlags ULONG, pui64Out *ULONG64) HRESULT {
	ret1 := syscall6(varUI8FromStr, 4,
		uintptr(unsafe.Pointer(strIn)),
		uintptr(lcid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pui64Out)),
		0,
		0)
	return HRESULT(ret1)
}

func VarUI8FromUI1(bIn byte, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromUI1, 2,
		uintptr(bIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromUI2(usIn USHORT, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromUI2, 2,
		uintptr(usIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUI8FromUI4(ulIn ULONG, pui64Out *ULONG64) HRESULT {
	ret1 := syscall3(varUI8FromUI4, 2,
		uintptr(ulIn),
		uintptr(unsafe.Pointer(pui64Out)),
		0)
	return HRESULT(ret1)
}

func VarUdateFromDate(dateIn DATE, dwFlags ULONG, lpUdate *UDATE) HRESULT {
	ret1 := syscall3(varUdateFromDate, 3,
		uintptr(dateIn),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpUdate)))
	return HRESULT(ret1)
}

func VarWeekdayName(iWeekday int32, fAbbrev int32, iFirstDay int32, dwFlags ULONG, pbstrOut *BSTR) HRESULT {
	ret1 := syscall6(varWeekdayName, 5,
		uintptr(iWeekday),
		uintptr(fAbbrev),
		uintptr(iFirstDay),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbstrOut)),
		0)
	return HRESULT(ret1)
}

func VarXor(pVarLeft *VARIANT, pVarRight *VARIANT, pVarOut *VARIANT) HRESULT {
	ret1 := syscall3(varXor, 3,
		uintptr(unsafe.Pointer(pVarLeft)),
		uintptr(unsafe.Pointer(pVarRight)),
		uintptr(unsafe.Pointer(pVarOut)))
	return HRESULT(ret1)
}

func VariantChangeType(pvargDest *VARIANTARG, pvargSrc *VARIANTARG, wFlags USHORT, vt VARTYPE) HRESULT {
	ret1 := syscall6(variantChangeType, 4,
		uintptr(unsafe.Pointer(pvargDest)),
		uintptr(unsafe.Pointer(pvargSrc)),
		uintptr(wFlags),
		uintptr(vt),
		0,
		0)
	return HRESULT(ret1)
}

func VariantChangeTypeEx(pvargDest *VARIANTARG, pvargSrc *VARIANTARG, lcid LCID, wFlags USHORT, vt VARTYPE) HRESULT {
	ret1 := syscall6(variantChangeTypeEx, 5,
		uintptr(unsafe.Pointer(pvargDest)),
		uintptr(unsafe.Pointer(pvargSrc)),
		uintptr(lcid),
		uintptr(wFlags),
		uintptr(vt),
		0)
	return HRESULT(ret1)
}

func VariantClear(pVarg *VARIANTARG) HRESULT {
	ret1 := syscall3(variantClear, 1,
		uintptr(unsafe.Pointer(pVarg)),
		0,
		0)
	return HRESULT(ret1)
}

func VariantCopy(pvargDest *VARIANTARG, pvargSrc *VARIANTARG) HRESULT {
	ret1 := syscall3(variantCopy, 2,
		uintptr(unsafe.Pointer(pvargDest)),
		uintptr(unsafe.Pointer(pvargSrc)),
		0)
	return HRESULT(ret1)
}

func VariantCopyInd(pvargDest *VARIANT, pvargSrc *VARIANTARG) HRESULT {
	ret1 := syscall3(variantCopyInd, 2,
		uintptr(unsafe.Pointer(pvargDest)),
		uintptr(unsafe.Pointer(pvargSrc)),
		0)
	return HRESULT(ret1)
}

func VariantInit(pVarg *VARIANTARG) {
	syscall3(variantInit, 1,
		uintptr(unsafe.Pointer(pVarg)),
		0,
		0)
}

func VariantTimeToDosDateTime(dateIn float64, pwDosDate *USHORT, pwDosTime *USHORT) int32 {
	ret1 := syscall3(variantTimeToDosDateTime, 3,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(pwDosDate)),
		uintptr(unsafe.Pointer(pwDosTime)))
	return int32(ret1)
}

func VariantTimeToSystemTime(dateIn float64, lpSt *SYSTEMTIME) int32 {
	ret1 := syscall3(variantTimeToSystemTime, 2,
		uintptr(dateIn),
		uintptr(unsafe.Pointer(lpSt)),
		0)
	return int32(ret1)
}

func VectorFromBstr(bstr BSTR, ppsa **SAFEARRAY) HRESULT {
	ret1 := syscall3(vectorFromBstr, 2,
		uintptr(unsafe.Pointer(bstr)),
		uintptr(unsafe.Pointer(ppsa)),
		0)
	return HRESULT(ret1)
}
