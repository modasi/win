package win

import (
	"reflect"
	"unsafe"
)

var (
	types map[string]reflect.Type
)

func init() {
	types = make(map[string]reflect.Type)
	types["ABC"] = reflect.TypeOf((*ABC)(nil)).Elem()
	types["ABCFLOAT"] = reflect.TypeOf((*ABCFLOAT)(nil)).Elem()
	types["ABORTPROC"] = reflect.TypeOf((*ABORTPROC)(nil)).Elem()
	types["ACCEL"] = reflect.TypeOf((*ACCEL)(nil)).Elem()
	types["ACCESS_MASK"] = reflect.TypeOf((*ACCESS_MASK)(nil)).Elem()
	types["ACCESS_MODE"] = reflect.TypeOf((*ACCESS_MODE)(nil)).Elem()
	types["ACL"] = reflect.TypeOf((*ACL)(nil)).Elem()
	types["ADDRINFO"] = reflect.TypeOf((*ADDRINFO)(nil)).Elem()
	types["ADDRINFOEX"] = reflect.TypeOf((*ADDRINFOEX)(nil)).Elem()
	types["AFPROTOCOLS"] = reflect.TypeOf((*AFPROTOCOLS)(nil)).Elem()
	types["ALG_ID"] = reflect.TypeOf((*ALG_ID)(nil)).Elem()
	types["ALTTABINFO"] = reflect.TypeOf((*ALTTABINFO)(nil)).Elem()
	types["ARRAY_INFO"] = reflect.TypeOf((*ARRAY_INFO)(nil)).Elem()
	types["ASSOCF"] = reflect.TypeOf((*ASSOCF)(nil)).Elem()
	types["ASSOCKEY"] = reflect.TypeOf((*ASSOCKEY)(nil)).Elem()
	types["ASSOCSTR"] = reflect.TypeOf((*ASSOCSTR)(nil)).Elem()
	types["ATOM"] = reflect.TypeOf((*ATOM)(nil)).Elem()
	types["AUDIT_POLICY_INFORMATION"] = reflect.TypeOf((*AUDIT_POLICY_INFORMATION)(nil)).Elem()
	types["AUXCAPS"] = reflect.TypeOf((*AUXCAPS)(nil)).Elem()
	types["BITMAP"] = reflect.TypeOf((*BITMAP)(nil)).Elem()
	types["BITMAPINFO"] = reflect.TypeOf((*BITMAPINFO)(nil)).Elem()
	types["BITMAPINFOHEADER"] = reflect.TypeOf((*BITMAPINFOHEADER)(nil)).Elem()
	types["BLENDFUNCTION"] = reflect.TypeOf((*BLENDFUNCTION)(nil)).Elem()
	types["BLENDOBJ"] = reflect.TypeOf((*BLENDOBJ)(nil)).Elem()
	types["BOOL"] = reflect.TypeOf((*BOOL)(nil)).Elem()
	types["BOOLEAN"] = reflect.TypeOf((*BOOLEAN)(nil)).Elem()
	types["BP_ANIMATIONPARAMS"] = reflect.TypeOf((*BP_ANIMATIONPARAMS)(nil)).Elem()
	types["BP_BUFFERFORMAT"] = reflect.TypeOf((*BP_BUFFERFORMAT)(nil)).Elem()
	types["BP_PAINTPARAMS"] = reflect.TypeOf((*BP_PAINTPARAMS)(nil)).Elem()
	types["BRUSHOBJ"] = reflect.TypeOf((*BRUSHOBJ)(nil)).Elem()
	types["BSMINFO"] = reflect.TypeOf((*BSMINFO)(nil)).Elem()
	types["BSTR"] = reflect.TypeOf((*BSTR)(nil)).Elem()
	types["BYTE"] = reflect.TypeOf((*BYTE)(nil)).Elem()
	types["CALLCONV"] = reflect.TypeOf((*CALLCONV)(nil)).Elem()
	types["CANDIDATEFORM"] = reflect.TypeOf((*CANDIDATEFORM)(nil)).Elem()
	types["CANDIDATELIST"] = reflect.TypeOf((*CANDIDATELIST)(nil)).Elem()
	types["CCHOOKPROC"] = reflect.TypeOf((*CCHOOKPROC)(nil)).Elem()
	types["CERT_EXTENSION"] = reflect.TypeOf((*CERT_EXTENSION)(nil)).Elem()
	types["CERT_NAME_BLOB"] = reflect.TypeOf((*CERT_NAME_BLOB)(nil)).Elem()
	types["CERT_RDN_VALUE_BLOB"] = reflect.TypeOf((*CERT_RDN_VALUE_BLOB)(nil)).Elem()
	types["CHAR"] = reflect.TypeOf((*CHAR)(nil)).Elem()
	types["CHARSETINFO"] = reflect.TypeOf((*CHARSETINFO)(nil)).Elem()
	types["CHOOSECOLOR"] = reflect.TypeOf((*CHOOSECOLOR)(nil)).Elem()
	types["CHOOSEFONT"] = reflect.TypeOf((*CHOOSEFONT)(nil)).Elem()
	types["CIEXYZ"] = reflect.TypeOf((*CIEXYZ)(nil)).Elem()
	types["CIEXYZTRIPLE"] = reflect.TypeOf((*CIEXYZTRIPLE)(nil)).Elem()
	types["CLIPLINE"] = reflect.TypeOf((*CLIPLINE)(nil)).Elem()
	types["CLIPOBJ"] = reflect.TypeOf((*CLIPOBJ)(nil)).Elem()
	types["CLSID"] = reflect.TypeOf((*CLSID)(nil)).Elem()
	types["COLOR16"] = reflect.TypeOf((*COLOR16)(nil)).Elem()
	types["COLORADJUSTMENT"] = reflect.TypeOf((*COLORADJUSTMENT)(nil)).Elem()
	types["COLORMAP"] = reflect.TypeOf((*COLORMAP)(nil)).Elem()
	types["COLORREF"] = reflect.TypeOf((*COLORREF)(nil)).Elem()
	types["COMBOBOXINFO"] = reflect.TypeOf((*COMBOBOXINFO)(nil)).Elem()
	types["COMM_FAULT_OFFSETS"] = reflect.TypeOf((*COMM_FAULT_OFFSETS)(nil)).Elem()
	types["COMPOSITIONFORM"] = reflect.TypeOf((*COMPOSITIONFORM)(nil)).Elem()
	types["CONVCONTEXT"] = reflect.TypeOf((*CONVCONTEXT)(nil)).Elem()
	types["CONVINFO"] = reflect.TypeOf((*CONVINFO)(nil)).Elem()
	types["COORD"] = reflect.TypeOf((*COORD)(nil)).Elem()
	types["CREDENTIAL"] = reflect.TypeOf((*CREDENTIAL)(nil)).Elem()
	types["CREDENTIAL_ATTRIBUTE"] = reflect.TypeOf((*CREDENTIAL_ATTRIBUTE)(nil)).Elem()
	types["CRL_CONTEXT"] = reflect.TypeOf((*CRL_CONTEXT)(nil)).Elem()
	types["CRL_ENTRY"] = reflect.TypeOf((*CRL_ENTRY)(nil)).Elem()
	types["CRL_INFO"] = reflect.TypeOf((*CRL_INFO)(nil)).Elem()
	types["CRYPTOAPI_BLOB_"] = reflect.TypeOf((*CRYPTOAPI_BLOB_)(nil)).Elem()
	types["CRYPT_ALGORITHM_IDENTIFIER"] = reflect.TypeOf((*CRYPT_ALGORITHM_IDENTIFIER)(nil)).Elem()
	types["CRYPT_DATA_BLOB"] = reflect.TypeOf((*CRYPT_DATA_BLOB)(nil)).Elem()
	types["CRYPT_HASH_BLOB"] = reflect.TypeOf((*CRYPT_HASH_BLOB)(nil)).Elem()
	types["CRYPT_INTEGER_BLOB"] = reflect.TypeOf((*CRYPT_INTEGER_BLOB)(nil)).Elem()
	types["CRYPT_OBJID_BLOB"] = reflect.TypeOf((*CRYPT_OBJID_BLOB)(nil)).Elem()
	types["CSADDR_INFO"] = reflect.TypeOf((*CSADDR_INFO)(nil)).Elem()
	types["CS_TAG_GETTING_ROUTINE"] = reflect.TypeOf((*CS_TAG_GETTING_ROUTINE)(nil)).Elem()
	types["CS_TYPE_FROM_NETCS_ROUTINE"] = reflect.TypeOf((*CS_TYPE_FROM_NETCS_ROUTINE)(nil)).Elem()
	types["CS_TYPE_LOCAL_SIZE_ROUTINE"] = reflect.TypeOf((*CS_TYPE_LOCAL_SIZE_ROUTINE)(nil)).Elem()
	types["CS_TYPE_NET_SIZE_ROUTINE"] = reflect.TypeOf((*CS_TYPE_NET_SIZE_ROUTINE)(nil)).Elem()
	types["CS_TYPE_TO_NETCS_ROUTINE"] = reflect.TypeOf((*CS_TYPE_TO_NETCS_ROUTINE)(nil)).Elem()
	types["CURSORINFO"] = reflect.TypeOf((*CURSORINFO)(nil)).Elem()
	types["CUSTDATA"] = reflect.TypeOf((*CUSTDATA)(nil)).Elem()
	types["CUSTDATAITEM"] = reflect.TypeOf((*CUSTDATAITEM)(nil)).Elem()
	types["CY"] = reflect.TypeOf((*CY)(nil)).Elem()
	types["DACOMPARE"] = reflect.TypeOf((*DACOMPARE)(nil)).Elem()
	types["DAENUMCALLBACK"] = reflect.TypeOf((*DAENUMCALLBACK)(nil)).Elem()
	types["DATE"] = reflect.TypeOf((*DATE)(nil)).Elem()
	types["DATETIME"] = reflect.TypeOf((*DATETIME)(nil)).Elem()
	types["DECIMAL"] = reflect.TypeOf((*DECIMAL)(nil)).Elem()
	types["DESIGNVECTOR"] = reflect.TypeOf((*DESIGNVECTOR)(nil)).Elem()
	types["DESKTOPENUMPROC"] = reflect.TypeOf((*DESKTOPENUMPROC)(nil)).Elem()
	types["DEVMODE"] = reflect.TypeOf((*DEVMODE)(nil)).Elem()
	types["DHPDEV"] = reflect.TypeOf((*DHPDEV)(nil)).Elem()
	types["DHSURF"] = reflect.TypeOf((*DHSURF)(nil)).Elem()
	types["DISPID"] = reflect.TypeOf((*DISPID)(nil)).Elem()
	types["DISPLAY_DEVICE"] = reflect.TypeOf((*DISPLAY_DEVICE)(nil)).Elem()
	types["DISPPARAMS"] = reflect.TypeOf((*DISPPARAMS)(nil)).Elem()
	types["DLGPROC"] = reflect.TypeOf((*DLGPROC)(nil)).Elem()
	types["DLGTEMPLATE"] = reflect.TypeOf((*DLGTEMPLATE)(nil)).Elem()
	types["DOCINFO"] = reflect.TypeOf((*DOCINFO)(nil)).Elem()
	types["DOUBLE"] = reflect.TypeOf((*DOUBLE)(nil)).Elem()
	types["DRAWSTATEPROC"] = reflect.TypeOf((*DRAWSTATEPROC)(nil)).Elem()
	types["DRAWTEXTPARAMS"] = reflect.TypeOf((*DRAWTEXTPARAMS)(nil)).Elem()
	types["DTBGOPTS"] = reflect.TypeOf((*DTBGOPTS)(nil)).Elem()
	types["DTTOPTS"] = reflect.TypeOf((*DTTOPTS)(nil)).Elem()
	types["DTT_CALLBACK_PROC"] = reflect.TypeOf((*DTT_CALLBACK_PROC)(nil)).Elem()
	types["DWORD"] = reflect.TypeOf((*DWORD)(nil)).Elem()
	types["DWORD_PTR"] = reflect.TypeOf((*DWORD_PTR)(nil)).Elem()
	types["EFS_CERTIFICATE_BLOB"] = reflect.TypeOf((*EFS_CERTIFICATE_BLOB)(nil)).Elem()
	types["ENCRYPTION_CERTIFICATE"] = reflect.TypeOf((*ENCRYPTION_CERTIFICATE)(nil)).Elem()
	types["ENCRYPTION_CERTIFICATE_LIST"] = reflect.TypeOf((*ENCRYPTION_CERTIFICATE_LIST)(nil)).Elem()
	types["ENG_TIME_FIELDS"] = reflect.TypeOf((*ENG_TIME_FIELDS)(nil)).Elem()
	types["ENHMETAHEADER"] = reflect.TypeOf((*ENHMETAHEADER)(nil)).Elem()
	types["ENHMETARECORD"] = reflect.TypeOf((*ENHMETARECORD)(nil)).Elem()
	types["ENHMFENUMPROC"] = reflect.TypeOf((*ENHMFENUMPROC)(nil)).Elem()
	types["ENUMLOGFONTEX"] = reflect.TypeOf((*ENUMLOGFONTEX)(nil)).Elem()
	types["ENUMLOGFONTEXDV"] = reflect.TypeOf((*ENUMLOGFONTEXDV)(nil)).Elem()
	types["ENUMRESLANGPROC"] = reflect.TypeOf((*ENUMRESLANGPROC)(nil)).Elem()
	types["ENUM_PAGE_FILE_INFORMATION"] = reflect.TypeOf((*ENUM_PAGE_FILE_INFORMATION)(nil)).Elem()
	types["EXCEPINFO"] = reflect.TypeOf((*EXCEPINFO)(nil)).Elem()
	types["EXPLICIT_ACCESS"] = reflect.TypeOf((*EXPLICIT_ACCESS)(nil)).Elem()
	types["EXPR_EVAL"] = reflect.TypeOf((*EXPR_EVAL)(nil)).Elem()
	types["Error_status_t"] = reflect.TypeOf((*Error_status_t)(nil)).Elem()
	types["FARPROC"] = reflect.TypeOf((*FARPROC)(nil)).Elem()
	types["FD_GLYPHATTR"] = reflect.TypeOf((*FD_GLYPHATTR)(nil)).Elem()
	types["FILETIME"] = reflect.TypeOf((*FILETIME)(nil)).Elem()
	types["FINDREPLACE"] = reflect.TypeOf((*FINDREPLACE)(nil)).Elem()
	types["FIX"] = reflect.TypeOf((*FIX)(nil)).Elem()
	types["FIXED"] = reflect.TypeOf((*FIXED)(nil)).Elem()
	types["FIXED_INFO_W2KSP1"] = reflect.TypeOf((*FIXED_INFO_W2KSP1)(nil)).Elem()
	types["FLASHWINFO"] = reflect.TypeOf((*FLASHWINFO)(nil)).Elem()
	types["FLOAT"] = reflect.TypeOf((*FLOAT)(nil)).Elem()
	types["FLOATL"] = reflect.TypeOf((*FLOATL)(nil)).Elem()
	types["FLOAT_LONG"] = reflect.TypeOf((*FLOAT_LONG)(nil)).Elem()
	types["FLONG"] = reflect.TypeOf((*FLONG)(nil)).Elem()
	types["FLOWSPEC"] = reflect.TypeOf((*FLOWSPEC)(nil)).Elem()
	types["FONTDESC"] = reflect.TypeOf((*FONTDESC)(nil)).Elem()
	types["FONTENUMPROC"] = reflect.TypeOf((*FONTENUMPROC)(nil)).Elem()
	types["FONTINFO"] = reflect.TypeOf((*FONTINFO)(nil)).Elem()
	types["FONTOBJ"] = reflect.TypeOf((*FONTOBJ)(nil)).Elem()
	types["FONTSIGNATURE"] = reflect.TypeOf((*FONTSIGNATURE)(nil)).Elem()
	types["FULL_PTR_XLAT_TABLES"] = reflect.TypeOf((*FULL_PTR_XLAT_TABLES)(nil)).Elem()
	types["FXPT2DOT30"] = reflect.TypeOf((*FXPT2DOT30)(nil)).Elem()
	types["GCP_RESULTS"] = reflect.TypeOf((*GCP_RESULTS)(nil)).Elem()
	types["GENERIC_BINDING_INFO"] = reflect.TypeOf((*GENERIC_BINDING_INFO)(nil)).Elem()
	types["GENERIC_BINDING_ROUTINE"] = reflect.TypeOf((*GENERIC_BINDING_ROUTINE)(nil)).Elem()
	types["GENERIC_BINDING_ROUTINE_PAIR"] = reflect.TypeOf((*GENERIC_BINDING_ROUTINE_PAIR)(nil)).Elem()
	types["GENERIC_MAPPING"] = reflect.TypeOf((*GENERIC_MAPPING)(nil)).Elem()
	types["GENERIC_UNBIND_ROUTINE"] = reflect.TypeOf((*GENERIC_UNBIND_ROUTINE)(nil)).Elem()
	types["GEOID"] = reflect.TypeOf((*GEOID)(nil)).Elem()
	types["GESTURECONFIG"] = reflect.TypeOf((*GESTURECONFIG)(nil)).Elem()
	types["GESTUREINFO"] = reflect.TypeOf((*GESTUREINFO)(nil)).Elem()
	types["GLYPHBITS"] = reflect.TypeOf((*GLYPHBITS)(nil)).Elem()
	types["GLYPHDEF"] = reflect.TypeOf((*GLYPHDEF)(nil)).Elem()
	types["GLYPHMETRICS"] = reflect.TypeOf((*GLYPHMETRICS)(nil)).Elem()
	types["GLYPHMETRICSFLOAT"] = reflect.TypeOf((*GLYPHMETRICSFLOAT)(nil)).Elem()
	types["GLYPHPOS"] = reflect.TypeOf((*GLYPHPOS)(nil)).Elem()
	types["GLYPHSET"] = reflect.TypeOf((*GLYPHSET)(nil)).Elem()
	types["GLbitfield"] = reflect.TypeOf((*GLbitfield)(nil)).Elem()
	types["GLboolean"] = reflect.TypeOf((*GLboolean)(nil)).Elem()
	types["GLbyte"] = reflect.TypeOf((*GLbyte)(nil)).Elem()
	types["GLclampd"] = reflect.TypeOf((*GLclampd)(nil)).Elem()
	types["GLclampf"] = reflect.TypeOf((*GLclampf)(nil)).Elem()
	types["GLdouble"] = reflect.TypeOf((*GLdouble)(nil)).Elem()
	types["GLenum"] = reflect.TypeOf((*GLenum)(nil)).Elem()
	types["GLfloat"] = reflect.TypeOf((*GLfloat)(nil)).Elem()
	types["GLint"] = reflect.TypeOf((*GLint)(nil)).Elem()
	types["GLshort"] = reflect.TypeOf((*GLshort)(nil)).Elem()
	types["GLsizei"] = reflect.TypeOf((*GLsizei)(nil)).Elem()
	types["GLubyte"] = reflect.TypeOf((*GLubyte)(nil)).Elem()
	types["GLuint"] = reflect.TypeOf((*GLuint)(nil)).Elem()
	types["GLushort"] = reflect.TypeOf((*GLushort)(nil)).Elem()
	types["GOBJENUMPROC"] = reflect.TypeOf((*GOBJENUMPROC)(nil)).Elem()
	types["GRAYSTRINGPROC"] = reflect.TypeOf((*GRAYSTRINGPROC)(nil)).Elem()
	types["GROUP"] = reflect.TypeOf((*GROUP)(nil)).Elem()
	types["GUID"] = reflect.TypeOf((*GUID)(nil)).Elem()
	types["GUITHREADINFO"] = reflect.TypeOf((*GUITHREADINFO)(nil)).Elem()
	types["GdiplusStartupInput"] = reflect.TypeOf((*GdiplusStartupInput)(nil)).Elem()
	types["GdiplusStartupOutput"] = reflect.TypeOf((*GdiplusStartupOutput)(nil)).Elem()
	types["GpStatus"] = reflect.TypeOf((*GpStatus)(nil)).Elem()
	types["HACCEL"] = reflect.TypeOf((*HACCEL)(nil)).Elem()
	types["HANDLE"] = reflect.TypeOf((*HANDLE)(nil)).Elem()
	types["HANDLER_FUNCTION_EX"] = reflect.TypeOf((*HANDLER_FUNCTION_EX)(nil)).Elem()
	types["HANDLETABLE"] = reflect.TypeOf((*HANDLETABLE)(nil)).Elem()
	types["HANIMATIONBUFFER"] = reflect.TypeOf((*HANIMATIONBUFFER)(nil)).Elem()
	types["HARDWAREINPUT"] = reflect.TypeOf((*HARDWAREINPUT)(nil)).Elem()
	types["HBITMAP"] = reflect.TypeOf((*HBITMAP)(nil)).Elem()
	types["HBRUSH"] = reflect.TypeOf((*HBRUSH)(nil)).Elem()
	types["HCERTSTORE"] = reflect.TypeOf((*HCERTSTORE)(nil)).Elem()
	types["HCOLORSPACE"] = reflect.TypeOf((*HCOLORSPACE)(nil)).Elem()
	types["HCONV"] = reflect.TypeOf((*HCONV)(nil)).Elem()
	types["HCONVLIST"] = reflect.TypeOf((*HCONVLIST)(nil)).Elem()
	types["HCRYPTHASH"] = reflect.TypeOf((*HCRYPTHASH)(nil)).Elem()
	types["HCRYPTKEY"] = reflect.TypeOf((*HCRYPTKEY)(nil)).Elem()
	types["HCRYPTPROV"] = reflect.TypeOf((*HCRYPTPROV)(nil)).Elem()
	types["HCURSOR"] = reflect.TypeOf((*HCURSOR)(nil)).Elem()
	types["HDC"] = reflect.TypeOf((*HDC)(nil)).Elem()
	types["HDDEDATA"] = reflect.TypeOf((*HDDEDATA)(nil)).Elem()
	types["HDESK"] = reflect.TypeOf((*HDESK)(nil)).Elem()
	types["HDEV"] = reflect.TypeOf((*HDEV)(nil)).Elem()
	types["HDEVNOTIFY"] = reflect.TypeOf((*HDEVNOTIFY)(nil)).Elem()
	types["HDPA"] = reflect.TypeOf((*HDPA)(nil)).Elem()
	types["HDRVR"] = reflect.TypeOf((*HDRVR)(nil)).Elem()
	types["HDSA"] = reflect.TypeOf((*HDSA)(nil)).Elem()
	types["HDWP"] = reflect.TypeOf((*HDWP)(nil)).Elem()
	types["HELPINFO"] = reflect.TypeOf((*HELPINFO)(nil)).Elem()
	types["HENHMETAFILE"] = reflect.TypeOf((*HENHMETAFILE)(nil)).Elem()
	types["HFONT"] = reflect.TypeOf((*HFONT)(nil)).Elem()
	types["HGDIOBJ"] = reflect.TypeOf((*HGDIOBJ)(nil)).Elem()
	types["HGESTUREINFO"] = reflect.TypeOf((*HGESTUREINFO)(nil)).Elem()
	types["HGLOBAL"] = reflect.TypeOf((*HGLOBAL)(nil)).Elem()
	types["HGLRC"] = reflect.TypeOf((*HGLRC)(nil)).Elem()
	types["HGLYPH"] = reflect.TypeOf((*HGLYPH)(nil)).Elem()
	types["HHOOK"] = reflect.TypeOf((*HHOOK)(nil)).Elem()
	types["HICON"] = reflect.TypeOf((*HICON)(nil)).Elem()
	types["HIMAGELIST"] = reflect.TypeOf((*HIMAGELIST)(nil)).Elem()
	types["HIMC"] = reflect.TypeOf((*HIMC)(nil)).Elem()
	types["HIMCC"] = reflect.TypeOf((*HIMCC)(nil)).Elem()
	types["HINSTANCE"] = reflect.TypeOf((*HINSTANCE)(nil)).Elem()
	types["HKEY"] = reflect.TypeOf((*HKEY)(nil)).Elem()
	types["HKL"] = reflect.TypeOf((*HKL)(nil)).Elem()
	types["HLOCAL"] = reflect.TypeOf((*HLOCAL)(nil)).Elem()
	types["HMENU"] = reflect.TypeOf((*HMENU)(nil)).Elem()
	types["HMETAFILE"] = reflect.TypeOf((*HMETAFILE)(nil)).Elem()
	types["HMIDI"] = reflect.TypeOf((*HMIDI)(nil)).Elem()
	types["HMIDIIN"] = reflect.TypeOf((*HMIDIIN)(nil)).Elem()
	types["HMIDIOUT"] = reflect.TypeOf((*HMIDIOUT)(nil)).Elem()
	types["HMIDISTRM"] = reflect.TypeOf((*HMIDISTRM)(nil)).Elem()
	types["HMODULE"] = reflect.TypeOf((*HMODULE)(nil)).Elem()
	types["HMONITOR"] = reflect.TypeOf((*HMONITOR)(nil)).Elem()
	types["HOOKPROC"] = reflect.TypeOf((*HOOKPROC)(nil)).Elem()
	types["HPAINTBUFFER"] = reflect.TypeOf((*HPAINTBUFFER)(nil)).Elem()
	types["HPALETTE"] = reflect.TypeOf((*HPALETTE)(nil)).Elem()
	types["HPEN"] = reflect.TypeOf((*HPEN)(nil)).Elem()
	types["HPOWERNOTIFY"] = reflect.TypeOf((*HPOWERNOTIFY)(nil)).Elem()
	types["HPROPSHEETPAGE"] = reflect.TypeOf((*HPROPSHEETPAGE)(nil)).Elem()
	types["HPSXA"] = reflect.TypeOf((*HPSXA)(nil)).Elem()
	types["HRAWINPUT"] = reflect.TypeOf((*HRAWINPUT)(nil)).Elem()
	types["HRESULT"] = reflect.TypeOf((*HRESULT)(nil)).Elem()
	types["HRGN"] = reflect.TypeOf((*HRGN)(nil)).Elem()
	types["HRSRC"] = reflect.TypeOf((*HRSRC)(nil)).Elem()
	types["HSEMAPHORE"] = reflect.TypeOf((*HSEMAPHORE)(nil)).Elem()
	types["HSURF"] = reflect.TypeOf((*HSURF)(nil)).Elem()
	types["HSZ"] = reflect.TypeOf((*HSZ)(nil)).Elem()
	types["HTASK"] = reflect.TypeOf((*HTASK)(nil)).Elem()
	types["HTHEME"] = reflect.TypeOf((*HTHEME)(nil)).Elem()
	types["HTOUCHINPUT"] = reflect.TypeOf((*HTOUCHINPUT)(nil)).Elem()
	types["HUSKEY"] = reflect.TypeOf((*HUSKEY)(nil)).Elem()
	types["HWCT"] = reflect.TypeOf((*HWCT)(nil)).Elem()
	types["HWINEVENTHOOK"] = reflect.TypeOf((*HWINEVENTHOOK)(nil)).Elem()
	types["HWINSTA"] = reflect.TypeOf((*HWINSTA)(nil)).Elem()
	types["HWND"] = reflect.TypeOf((*HWND)(nil)).Elem()
	types["Handle_t"] = reflect.TypeOf((*Handle_t)(nil)).Elem()
	types["IBindCtx"] = reflect.TypeOf((*IBindCtx)(nil)).Elem()
	types["ICMENUMPROC"] = reflect.TypeOf((*ICMENUMPROC)(nil)).Elem()
	types["ICONINFO"] = reflect.TypeOf((*ICONINFO)(nil)).Elem()
	types["IConnectionPoint"] = reflect.TypeOf((*IConnectionPoint)(nil)).Elem()
	types["ICreateErrorInfo"] = reflect.TypeOf((*ICreateErrorInfo)(nil)).Elem()
	types["ICreateTypeLib"] = reflect.TypeOf((*ICreateTypeLib)(nil)).Elem()
	types["ICreateTypeLib2"] = reflect.TypeOf((*ICreateTypeLib2)(nil)).Elem()
	types["IDL_CS_CONVERT"] = reflect.TypeOf((*IDL_CS_CONVERT)(nil)).Elem()
	types["IDispatch"] = reflect.TypeOf((*IDispatch)(nil)).Elem()
	types["IErrorInfo"] = reflect.TypeOf((*IErrorInfo)(nil)).Elem()
	types["IFTYPE"] = reflect.TypeOf((*IFTYPE)(nil)).Elem()
	types["IF_INDEX"] = reflect.TypeOf((*IF_INDEX)(nil)).Elem()
	types["IF_LUID"] = reflect.TypeOf((*IF_LUID)(nil)).Elem()
	types["IID"] = reflect.TypeOf((*IID)(nil)).Elem()
	types["IMAGEINFO"] = reflect.TypeOf((*IMAGEINFO)(nil)).Elem()
	types["IMAGELISTDRAWPARAMS"] = reflect.TypeOf((*IMAGELISTDRAWPARAMS)(nil)).Elem()
	types["IMCENUMPROC"] = reflect.TypeOf((*IMCENUMPROC)(nil)).Elem()
	types["IMEMENUITEMINFO"] = reflect.TypeOf((*IMEMENUITEMINFO)(nil)).Elem()
	types["IMEPRO"] = reflect.TypeOf((*IMEPRO)(nil)).Elem()
	types["IMoniker"] = reflect.TypeOf((*IMoniker)(nil)).Elem()
	types["INITCOMMONCONTROLSEX"] = reflect.TypeOf((*INITCOMMONCONTROLSEX)(nil)).Elem()
	types["INPUT"] = reflect.TypeOf((*INPUT)(nil)).Elem()
	types["INPUTCONTEXT"] = reflect.TypeOf((*INPUTCONTEXT)(nil)).Elem()
	types["INT"] = reflect.TypeOf((*INT)(nil)).Elem()
	types["INTERFACEDATA"] = reflect.TypeOf((*INTERFACEDATA)(nil)).Elem()
	types["INTERFACE_HANDLE"] = reflect.TypeOf((*INTERFACE_HANDLE)(nil)).Elem()
	types["INTERNAL_IF_OPER_STATUS"] = reflect.TypeOf((*INTERNAL_IF_OPER_STATUS)(nil)).Elem()
	types["INTLIST"] = reflect.TypeOf((*INTLIST)(nil)).Elem()
	types["INT_PTR"] = reflect.TypeOf((*INT_PTR)(nil)).Elem()
	types["IPAddr"] = reflect.TypeOf((*IPAddr)(nil)).Elem()
	types["IPMask"] = reflect.TypeOf((*IPMask)(nil)).Elem()
	types["IP_ADAPTER_ADDRESSES_LH"] = reflect.TypeOf((*IP_ADAPTER_ADDRESSES_LH)(nil)).Elem()
	types["IP_ADAPTER_ANYCAST_ADDRESS_XP"] = reflect.TypeOf((*IP_ADAPTER_ANYCAST_ADDRESS_XP)(nil)).Elem()
	types["IP_ADAPTER_DNS_SERVER_ADDRESS_XP"] = reflect.TypeOf((*IP_ADAPTER_DNS_SERVER_ADDRESS_XP)(nil)).Elem()
	types["IP_ADAPTER_DNS_SUFFIX"] = reflect.TypeOf((*IP_ADAPTER_DNS_SUFFIX)(nil)).Elem()
	types["IP_ADAPTER_GATEWAY_ADDRESS_LH"] = reflect.TypeOf((*IP_ADAPTER_GATEWAY_ADDRESS_LH)(nil)).Elem()
	types["IP_ADAPTER_INDEX_MAP"] = reflect.TypeOf((*IP_ADAPTER_INDEX_MAP)(nil)).Elem()
	types["IP_ADAPTER_INFO"] = reflect.TypeOf((*IP_ADAPTER_INFO)(nil)).Elem()
	types["IP_ADAPTER_MULTICAST_ADDRESS_XP"] = reflect.TypeOf((*IP_ADAPTER_MULTICAST_ADDRESS_XP)(nil)).Elem()
	types["IP_ADAPTER_ORDER_MAP"] = reflect.TypeOf((*IP_ADAPTER_ORDER_MAP)(nil)).Elem()
	types["IP_ADAPTER_PREFIX_XP"] = reflect.TypeOf((*IP_ADAPTER_PREFIX_XP)(nil)).Elem()
	types["IP_ADAPTER_UNICAST_ADDRESS_LH"] = reflect.TypeOf((*IP_ADAPTER_UNICAST_ADDRESS_LH)(nil)).Elem()
	types["IP_ADAPTER_WINS_SERVER_ADDRESS_LH"] = reflect.TypeOf((*IP_ADAPTER_WINS_SERVER_ADDRESS_LH)(nil)).Elem()
	types["IP_ADDRESS_STRING"] = reflect.TypeOf((*IP_ADDRESS_STRING)(nil)).Elem()
	types["IP_ADDR_STRING"] = reflect.TypeOf((*IP_ADDR_STRING)(nil)).Elem()
	types["IP_DAD_STATE"] = reflect.TypeOf((*IP_DAD_STATE)(nil)).Elem()
	types["IP_INTERFACE_INFO"] = reflect.TypeOf((*IP_INTERFACE_INFO)(nil)).Elem()
	types["IP_MASK_STRING"] = reflect.TypeOf((*IP_MASK_STRING)(nil)).Elem()
	types["IP_PER_ADAPTER_INFO_W2KSP1"] = reflect.TypeOf((*IP_PER_ADAPTER_INFO_W2KSP1)(nil)).Elem()
	types["IP_PREFIX_ORIGIN"] = reflect.TypeOf((*IP_PREFIX_ORIGIN)(nil)).Elem()
	types["IP_STATUS"] = reflect.TypeOf((*IP_STATUS)(nil)).Elem()
	types["IP_SUFFIX_ORIGIN"] = reflect.TypeOf((*IP_SUFFIX_ORIGIN)(nil)).Elem()
	types["IRecordInfo"] = reflect.TypeOf((*IRecordInfo)(nil)).Elem()
	types["IRpcChannelBuffer"] = reflect.TypeOf((*IRpcChannelBuffer)(nil)).Elem()
	types["IRpcStubBuffer"] = reflect.TypeOf((*IRpcStubBuffer)(nil)).Elem()
	types["IStream"] = reflect.TypeOf((*IStream)(nil)).Elem()
	types["ITEMIDLIST"] = reflect.TypeOf((*ITEMIDLIST)(nil)).Elem()
	types["ITypeInfo"] = reflect.TypeOf((*ITypeInfo)(nil)).Elem()
	types["ITypeLib"] = reflect.TypeOf((*ITypeLib)(nil)).Elem()
	types["IUnknown"] = reflect.TypeOf((*IUnknown)(nil)).Elem()
	types["I_RPC_HANDLE"] = reflect.TypeOf((*I_RPC_HANDLE)(nil)).Elem()
	types["KERNINGPAIR"] = reflect.TypeOf((*KERNINGPAIR)(nil)).Elem()
	types["KEYBDINPUT"] = reflect.TypeOf((*KEYBDINPUT)(nil)).Elem()
	types["LANGID"] = reflect.TypeOf((*LANGID)(nil)).Elem()
	types["LARGE_INTEGER"] = reflect.TypeOf((*LARGE_INTEGER)(nil)).Elem()
	types["LASTINPUTINFO"] = reflect.TypeOf((*LASTINPUTINFO)(nil)).Elem()
	types["LAYERPLANEDESCRIPTOR"] = reflect.TypeOf((*LAYERPLANEDESCRIPTOR)(nil)).Elem()
	types["LCID"] = reflect.TypeOf((*LCID)(nil)).Elem()
	types["LCSCSTYPE"] = reflect.TypeOf((*LCSCSTYPE)(nil)).Elem()
	types["LCSGAMUTMATCH"] = reflect.TypeOf((*LCSGAMUTMATCH)(nil)).Elem()
	types["LCTYPE"] = reflect.TypeOf((*LCTYPE)(nil)).Elem()
	types["LINEATTRS"] = reflect.TypeOf((*LINEATTRS)(nil)).Elem()
	types["LINEDDAPROC"] = reflect.TypeOf((*LINEDDAPROC)(nil)).Elem()
	types["LOGBRUSH"] = reflect.TypeOf((*LOGBRUSH)(nil)).Elem()
	types["LOGCOLORSPACE"] = reflect.TypeOf((*LOGCOLORSPACE)(nil)).Elem()
	types["LOGFONT"] = reflect.TypeOf((*LOGFONT)(nil)).Elem()
	types["LOGPALETTE"] = reflect.TypeOf((*LOGPALETTE)(nil)).Elem()
	types["LOGPEN"] = reflect.TypeOf((*LOGPEN)(nil)).Elem()
	types["LONG"] = reflect.TypeOf((*LONG)(nil)).Elem()
	types["LONG64"] = reflect.TypeOf((*LONG64)(nil)).Elem()
	types["LONGLONG"] = reflect.TypeOf((*LONGLONG)(nil)).Elem()
	types["LONG_PTR"] = reflect.TypeOf((*LONG_PTR)(nil)).Elem()
	types["LPAFPROTOCOLS"] = reflect.TypeOf((*LPAFPROTOCOLS)(nil)).Elem()
	types["LPARAM"] = reflect.TypeOf((*LPARAM)(nil)).Elem()
	types["LPBLOB"] = reflect.TypeOf((*LPBLOB)(nil)).Elem()
	types["LPBYTE"] = reflect.TypeOf((*LPBYTE)(nil)).Elem()
	types["LPCFHOOKPROC"] = reflect.TypeOf((*LPCFHOOKPROC)(nil)).Elem()
	types["LPCHOOSEFONT"] = reflect.TypeOf((*LPCHOOSEFONT)(nil)).Elem()
	types["LPCITEMIDLIST"] = reflect.TypeOf((*LPCITEMIDLIST)(nil)).Elem()
	types["LPCOLESTR"] = reflect.TypeOf((*LPCOLESTR)(nil)).Elem()
	types["LPCONDITIONPROC"] = reflect.TypeOf((*LPCONDITIONPROC)(nil)).Elem()
	types["LPCSADDR_INFO"] = reflect.TypeOf((*LPCSADDR_INFO)(nil)).Elem()
	types["LPCSTR"] = reflect.TypeOf((*LPCSTR)(nil)).Elem()
	types["LPCWSTR"] = reflect.TypeOf((*LPCWSTR)(nil)).Elem()
	types["LPDEVMODE"] = reflect.TypeOf((*LPDEVMODE)(nil)).Elem()
	types["LPDISPATCH"] = reflect.TypeOf((*LPDISPATCH)(nil)).Elem()
	types["LPFINDREPLACE"] = reflect.TypeOf((*LPFINDREPLACE)(nil)).Elem()
	types["LPFRHOOKPROC"] = reflect.TypeOf((*LPFRHOOKPROC)(nil)).Elem()
	types["LPGCP_RESULTS"] = reflect.TypeOf((*LPGCP_RESULTS)(nil)).Elem()
	types["LPGUID"] = reflect.TypeOf((*LPGUID)(nil)).Elem()
	types["LPHELPINFO"] = reflect.TypeOf((*LPHELPINFO)(nil)).Elem()
	types["LPIMEMENUITEMINFO"] = reflect.TypeOf((*LPIMEMENUITEMINFO)(nil)).Elem()
	types["LPLOGCOLORSPACE"] = reflect.TypeOf((*LPLOGCOLORSPACE)(nil)).Elem()
	types["LPLOGFONT"] = reflect.TypeOf((*LPLOGFONT)(nil)).Elem()
	types["LPLOOKUPSERVICE_COMPLETION_ROUTINE"] = reflect.TypeOf((*LPLOOKUPSERVICE_COMPLETION_ROUTINE)(nil)).Elem()
	types["LPMONIKER"] = reflect.TypeOf((*LPMONIKER)(nil)).Elem()
	types["LPOFNHOOKPROC"] = reflect.TypeOf((*LPOFNHOOKPROC)(nil)).Elem()
	types["LPOLESTR"] = reflect.TypeOf((*LPOLESTR)(nil)).Elem()
	types["LPOPENFILENAME"] = reflect.TypeOf((*LPOPENFILENAME)(nil)).Elem()
	types["LPOUTLINETEXTMETRIC"] = reflect.TypeOf((*LPOUTLINETEXTMETRIC)(nil)).Elem()
	types["LPPAGEPAINTHOOK"] = reflect.TypeOf((*LPPAGEPAINTHOOK)(nil)).Elem()
	types["LPPAGESETUPDLG"] = reflect.TypeOf((*LPPAGESETUPDLG)(nil)).Elem()
	types["LPPAGESETUPHOOK"] = reflect.TypeOf((*LPPAGESETUPHOOK)(nil)).Elem()
	types["LPPRINTDLG"] = reflect.TypeOf((*LPPRINTDLG)(nil)).Elem()
	types["LPPRINTDLGEX"] = reflect.TypeOf((*LPPRINTDLGEX)(nil)).Elem()
	types["LPPRINTHOOKPROC"] = reflect.TypeOf((*LPPRINTHOOKPROC)(nil)).Elem()
	types["LPPRINTPAGERANGE"] = reflect.TypeOf((*LPPRINTPAGERANGE)(nil)).Elem()
	types["LPQOS"] = reflect.TypeOf((*LPQOS)(nil)).Elem()
	types["LPRASTERIZER_STATUS"] = reflect.TypeOf((*LPRASTERIZER_STATUS)(nil)).Elem()
	types["LPRECT"] = reflect.TypeOf((*LPRECT)(nil)).Elem()
	types["LPSAFEARRAY"] = reflect.TypeOf((*LPSAFEARRAY)(nil)).Elem()
	types["LPSETUPHOOKPROC"] = reflect.TypeOf((*LPSETUPHOOKPROC)(nil)).Elem()
	types["LPSOCKADDR"] = reflect.TypeOf((*LPSOCKADDR)(nil)).Elem()
	types["LPSTR"] = reflect.TypeOf((*LPSTR)(nil)).Elem()
	types["LPSTREAM"] = reflect.TypeOf((*LPSTREAM)(nil)).Elem()
	types["LPSTYLEBUF"] = reflect.TypeOf((*LPSTYLEBUF)(nil)).Elem()
	types["LPTEXTMETRIC"] = reflect.TypeOf((*LPTEXTMETRIC)(nil)).Elem()
	types["LPUNKNOWN"] = reflect.TypeOf((*LPUNKNOWN)(nil)).Elem()
	types["LPVOID"] = reflect.TypeOf((*LPVOID)(nil)).Elem()
	types["LPWPUPOSTMESSAGE"] = reflect.TypeOf((*LPWPUPOSTMESSAGE)(nil)).Elem()
	types["LPWSABUF"] = reflect.TypeOf((*LPWSABUF)(nil)).Elem()
	types["LPWSANAMESPACE_INFO"] = reflect.TypeOf((*LPWSANAMESPACE_INFO)(nil)).Elem()
	types["LPWSANSCLASSINFO"] = reflect.TypeOf((*LPWSANSCLASSINFO)(nil)).Elem()
	types["LPWSAOVERLAPPED"] = reflect.TypeOf((*LPWSAOVERLAPPED)(nil)).Elem()
	types["LPWSAOVERLAPPED_COMPLETION_ROUTINE"] = reflect.TypeOf((*LPWSAOVERLAPPED_COMPLETION_ROUTINE)(nil)).Elem()
	types["LPWSAPROTOCOL_INFO"] = reflect.TypeOf((*LPWSAPROTOCOL_INFO)(nil)).Elem()
	types["LPWSAQUERYSET"] = reflect.TypeOf((*LPWSAQUERYSET)(nil)).Elem()
	types["LPWSASERVICECLASSINFO"] = reflect.TypeOf((*LPWSASERVICECLASSINFO)(nil)).Elem()
	types["LPWSAVERSION"] = reflect.TypeOf((*LPWSAVERSION)(nil)).Elem()
	types["LPWSTR"] = reflect.TypeOf((*LPWSTR)(nil)).Elem()
	types["LRESULT"] = reflect.TypeOf((*LRESULT)(nil)).Elem()
	types["LUID"] = reflect.TypeOf((*LUID)(nil)).Elem()
	types["LUID_AND_ATTRIBUTES"] = reflect.TypeOf((*LUID_AND_ATTRIBUTES)(nil)).Elem()
	types["MALLOC_FREE_STRUCT"] = reflect.TypeOf((*MALLOC_FREE_STRUCT)(nil)).Elem()
	types["MARGINS"] = reflect.TypeOf((*MARGINS)(nil)).Elem()
	types["MAT2"] = reflect.TypeOf((*MAT2)(nil)).Elem()
	types["MCIDEVICEID"] = reflect.TypeOf((*MCIDEVICEID)(nil)).Elem()
	types["MCIERROR"] = reflect.TypeOf((*MCIERROR)(nil)).Elem()
	types["MENUBARINFO"] = reflect.TypeOf((*MENUBARINFO)(nil)).Elem()
	types["MENUINFO"] = reflect.TypeOf((*MENUINFO)(nil)).Elem()
	types["MENUITEMINFO"] = reflect.TypeOf((*MENUITEMINFO)(nil)).Elem()
	types["METAFILEPICT"] = reflect.TypeOf((*METAFILEPICT)(nil)).Elem()
	types["METARECORD"] = reflect.TypeOf((*METARECORD)(nil)).Elem()
	types["METHODDATA"] = reflect.TypeOf((*METHODDATA)(nil)).Elem()
	types["MFENUMPROC"] = reflect.TypeOf((*MFENUMPROC)(nil)).Elem()
	types["MIBICMPINFO"] = reflect.TypeOf((*MIBICMPINFO)(nil)).Elem()
	types["MIBICMPSTATS"] = reflect.TypeOf((*MIBICMPSTATS)(nil)).Elem()
	types["MIBICMPSTATS_EX"] = reflect.TypeOf((*MIBICMPSTATS_EX)(nil)).Elem()
	types["MIBICMPSTATS_EX_XPSP1"] = reflect.TypeOf((*MIBICMPSTATS_EX_XPSP1)(nil)).Elem()
	types["MIB_ICMP"] = reflect.TypeOf((*MIB_ICMP)(nil)).Elem()
	types["MIB_ICMP_EX_XPSP1"] = reflect.TypeOf((*MIB_ICMP_EX_XPSP1)(nil)).Elem()
	types["MIB_IFROW"] = reflect.TypeOf((*MIB_IFROW)(nil)).Elem()
	types["MIB_IFTABLE"] = reflect.TypeOf((*MIB_IFTABLE)(nil)).Elem()
	types["MIB_IPADDRROW"] = reflect.TypeOf((*MIB_IPADDRROW)(nil)).Elem()
	types["MIB_IPADDRROW_XP"] = reflect.TypeOf((*MIB_IPADDRROW_XP)(nil)).Elem()
	types["MIB_IPADDRTABLE"] = reflect.TypeOf((*MIB_IPADDRTABLE)(nil)).Elem()
	types["MIB_IPFORWARDROW"] = reflect.TypeOf((*MIB_IPFORWARDROW)(nil)).Elem()
	types["MIB_IPFORWARDTABLE"] = reflect.TypeOf((*MIB_IPFORWARDTABLE)(nil)).Elem()
	types["MIB_IPFORWARD_PROTO"] = reflect.TypeOf((*MIB_IPFORWARD_PROTO)(nil)).Elem()
	types["MIB_IPFORWARD_TYPE"] = reflect.TypeOf((*MIB_IPFORWARD_TYPE)(nil)).Elem()
	types["MIB_IPNETROW"] = reflect.TypeOf((*MIB_IPNETROW)(nil)).Elem()
	types["MIB_IPNETROW_LH"] = reflect.TypeOf((*MIB_IPNETROW_LH)(nil)).Elem()
	types["MIB_IPNETTABLE"] = reflect.TypeOf((*MIB_IPNETTABLE)(nil)).Elem()
	types["MIB_IPNET_TYPE"] = reflect.TypeOf((*MIB_IPNET_TYPE)(nil)).Elem()
	types["MIB_IPSTATS_FORWARDING"] = reflect.TypeOf((*MIB_IPSTATS_FORWARDING)(nil)).Elem()
	types["MIB_IPSTATS_LH"] = reflect.TypeOf((*MIB_IPSTATS_LH)(nil)).Elem()
	types["MIB_TCP6ROW_OWNER_MODULE"] = reflect.TypeOf((*MIB_TCP6ROW_OWNER_MODULE)(nil)).Elem()
	types["MIB_TCPROW_LH"] = reflect.TypeOf((*MIB_TCPROW_LH)(nil)).Elem()
	types["MIB_TCPROW_OWNER_MODULE"] = reflect.TypeOf((*MIB_TCPROW_OWNER_MODULE)(nil)).Elem()
	types["MIB_TCP_STATE"] = reflect.TypeOf((*MIB_TCP_STATE)(nil)).Elem()
	types["MIB_UDP6ROW_OWNER_MODULE"] = reflect.TypeOf((*MIB_UDP6ROW_OWNER_MODULE)(nil)).Elem()
	types["MIB_UDPROW_OWNER_MODULE"] = reflect.TypeOf((*MIB_UDPROW_OWNER_MODULE)(nil)).Elem()
	types["MIDIHDR"] = reflect.TypeOf((*MIDIHDR)(nil)).Elem()
	types["MIDIINCAPS"] = reflect.TypeOf((*MIDIINCAPS)(nil)).Elem()
	types["MIDL_STUB_DESC"] = reflect.TypeOf((*MIDL_STUB_DESC)(nil)).Elem()
	types["MIDL_STUB_MESSAGE"] = reflect.TypeOf((*MIDL_STUB_MESSAGE)(nil)).Elem()
	types["MIX"] = reflect.TypeOf((*MIX)(nil)).Elem()
	types["MMRESULT"] = reflect.TypeOf((*MMRESULT)(nil)).Elem()
	types["MMVERSION"] = reflect.TypeOf((*MMVERSION)(nil)).Elem()
	types["MODULEINFO"] = reflect.TypeOf((*MODULEINFO)(nil)).Elem()
	types["MONITORENUMPROC"] = reflect.TypeOf((*MONITORENUMPROC)(nil)).Elem()
	types["MONITORINFO"] = reflect.TypeOf((*MONITORINFO)(nil)).Elem()
	types["MOUSEINPUT"] = reflect.TypeOf((*MOUSEINPUT)(nil)).Elem()
	types["MOUSEMOVEPOINT"] = reflect.TypeOf((*MOUSEMOVEPOINT)(nil)).Elem()
	types["MRUCMPPROC"] = reflect.TypeOf((*MRUCMPPROC)(nil)).Elem()
	types["MRUINFO"] = reflect.TypeOf((*MRUINFO)(nil)).Elem()
	types["MSG"] = reflect.TypeOf((*MSG)(nil)).Elem()
	types["MSGBOXCALLBACK"] = reflect.TypeOf((*MSGBOXCALLBACK)(nil)).Elem()
	types["MSGBOXPARAMS"] = reflect.TypeOf((*MSGBOXPARAMS)(nil)).Elem()
	types["MULTIPLE_TRUSTEE_OPERATION"] = reflect.TypeOf((*MULTIPLE_TRUSTEE_OPERATION)(nil)).Elem()
	types["NDR_CS_ROUTINES"] = reflect.TypeOf((*NDR_CS_ROUTINES)(nil)).Elem()
	types["NDR_CS_SIZE_CONVERT_ROUTINES"] = reflect.TypeOf((*NDR_CS_SIZE_CONVERT_ROUTINES)(nil)).Elem()
	types["NDR_EXPR_DESC"] = reflect.TypeOf((*NDR_EXPR_DESC)(nil)).Elem()
	types["NDR_RUNDOWN"] = reflect.TypeOf((*NDR_RUNDOWN)(nil)).Elem()
	types["NDR_SCONTEXT"] = reflect.TypeOf((*NDR_SCONTEXT)(nil)).Elem()
	types["NDR_SCONTEXT_"] = reflect.TypeOf((*NDR_SCONTEXT_)(nil)).Elem()
	types["NET_IFINDEX"] = reflect.TypeOf((*NET_IFINDEX)(nil)).Elem()
	types["NET_IF_COMPARTMENT_ID"] = reflect.TypeOf((*NET_IF_COMPARTMENT_ID)(nil)).Elem()
	types["NET_IF_CONNECTION_TYPE"] = reflect.TypeOf((*NET_IF_CONNECTION_TYPE)(nil)).Elem()
	types["NET_IF_NETWORK_GUID"] = reflect.TypeOf((*NET_IF_NETWORK_GUID)(nil)).Elem()
	types["NET_LUID"] = reflect.TypeOf((*NET_LUID)(nil)).Elem()
	types["NET_LUID_LH"] = reflect.TypeOf((*NET_LUID_LH)(nil)).Elem()
	types["NL_DAD_STATE"] = reflect.TypeOf((*NL_DAD_STATE)(nil)).Elem()
	types["NL_PREFIX_ORIGIN"] = reflect.TypeOf((*NL_PREFIX_ORIGIN)(nil)).Elem()
	types["NL_ROUTE_PROTOCOL"] = reflect.TypeOf((*NL_ROUTE_PROTOCOL)(nil)).Elem()
	types["NL_SUFFIX_ORIGIN"] = reflect.TypeOf((*NL_SUFFIX_ORIGIN)(nil)).Elem()
	types["NUMPARSE"] = reflect.TypeOf((*NUMPARSE)(nil)).Elem()
	types["OBJECTS_AND_NAME"] = reflect.TypeOf((*OBJECTS_AND_NAME)(nil)).Elem()
	types["OBJECTS_AND_SID"] = reflect.TypeOf((*OBJECTS_AND_SID)(nil)).Elem()
	types["OBJECT_TYPE_LIST"] = reflect.TypeOf((*OBJECT_TYPE_LIST)(nil)).Elem()
	types["OCPFIPARAMS"] = reflect.TypeOf((*OCPFIPARAMS)(nil)).Elem()
	types["OLECHAR"] = reflect.TypeOf((*OLECHAR)(nil)).Elem()
	types["OLE_COLOR"] = reflect.TypeOf((*OLE_COLOR)(nil)).Elem()
	types["OPENFILENAME"] = reflect.TypeOf((*OPENFILENAME)(nil)).Elem()
	types["OUTLINETEXTMETRIC"] = reflect.TypeOf((*OUTLINETEXTMETRIC)(nil)).Elem()
	types["OVERLAPPED"] = reflect.TypeOf((*OVERLAPPED)(nil)).Elem()
	types["PADDRINFO"] = reflect.TypeOf((*PADDRINFO)(nil)).Elem()
	types["PAGESETUPDLG"] = reflect.TypeOf((*PAGESETUPDLG)(nil)).Elem()
	types["PAINTSTRUCT"] = reflect.TypeOf((*PAINTSTRUCT)(nil)).Elem()
	types["PALETTEENTRY"] = reflect.TypeOf((*PALETTEENTRY)(nil)).Elem()
	types["PANOSE"] = reflect.TypeOf((*PANOSE)(nil)).Elem()
	types["PARAMDATA"] = reflect.TypeOf((*PARAMDATA)(nil)).Elem()
	types["PARRAY_INFO"] = reflect.TypeOf((*PARRAY_INFO)(nil)).Elem()
	types["PARSEDURL"] = reflect.TypeOf((*PARSEDURL)(nil)).Elem()
	types["PATHDATA"] = reflect.TypeOf((*PATHDATA)(nil)).Elem()
	types["PATHOBJ"] = reflect.TypeOf((*PATHOBJ)(nil)).Elem()
	types["PAUDIT_POLICY_INFORMATION"] = reflect.TypeOf((*PAUDIT_POLICY_INFORMATION)(nil)).Elem()
	types["PBYTE"] = reflect.TypeOf((*PBYTE)(nil)).Elem()
	types["PCCRL_CONTEXT"] = reflect.TypeOf((*PCCRL_CONTEXT)(nil)).Elem()
	types["PCERT_EXTENSION"] = reflect.TypeOf((*PCERT_EXTENSION)(nil)).Elem()
	types["PCERT_NAME_BLOB"] = reflect.TypeOf((*PCERT_NAME_BLOB)(nil)).Elem()
	types["PCERT_RDN_VALUE_BLOB"] = reflect.TypeOf((*PCERT_RDN_VALUE_BLOB)(nil)).Elem()
	types["PCHAR"] = reflect.TypeOf((*PCHAR)(nil)).Elem()
	types["PCRL_ENTRY"] = reflect.TypeOf((*PCRL_ENTRY)(nil)).Elem()
	types["PCRL_INFO"] = reflect.TypeOf((*PCRL_INFO)(nil)).Elem()
	types["PCRYPT_INTEGER_BLOB"] = reflect.TypeOf((*PCRYPT_INTEGER_BLOB)(nil)).Elem()
	types["PCWSTR"] = reflect.TypeOf((*PCWSTR)(nil)).Elem()
	types["PDH_COUNTER_INFO"] = reflect.TypeOf((*PDH_COUNTER_INFO)(nil)).Elem()
	types["PDH_COUNTER_PATH_ELEMENTS"] = reflect.TypeOf((*PDH_COUNTER_PATH_ELEMENTS)(nil)).Elem()
	types["PDH_DATA_ITEM_PATH_ELEMENTS"] = reflect.TypeOf((*PDH_DATA_ITEM_PATH_ELEMENTS)(nil)).Elem()
	types["PDH_FMT_COUNTERVALUE"] = reflect.TypeOf((*PDH_FMT_COUNTERVALUE)(nil)).Elem()
	types["PDH_HCOUNTER"] = reflect.TypeOf((*PDH_HCOUNTER)(nil)).Elem()
	types["PDH_HLOG"] = reflect.TypeOf((*PDH_HLOG)(nil)).Elem()
	types["PDH_HQUERY"] = reflect.TypeOf((*PDH_HQUERY)(nil)).Elem()
	types["PDH_RAW_COUNTER"] = reflect.TypeOf((*PDH_RAW_COUNTER)(nil)).Elem()
	types["PDH_STATUS"] = reflect.TypeOf((*PDH_STATUS)(nil)).Elem()
	types["PENCRYPTION_CERTIFICATE"] = reflect.TypeOf((*PENCRYPTION_CERTIFICATE)(nil)).Elem()
	types["PENG_TIME_FIELDS"] = reflect.TypeOf((*PENG_TIME_FIELDS)(nil)).Elem()
	types["PENUM_PAGE_FILE_CALLBACK"] = reflect.TypeOf((*PENUM_PAGE_FILE_CALLBACK)(nil)).Elem()
	types["PENUM_PAGE_FILE_INFORMATION"] = reflect.TypeOf((*PENUM_PAGE_FILE_INFORMATION)(nil)).Elem()
	types["PERCEIVED"] = reflect.TypeOf((*PERCEIVED)(nil)).Elem()
	types["PERFORMANCE_INFORMATION"] = reflect.TypeOf((*PERFORMANCE_INFORMATION)(nil)).Elem()
	types["PERF_COUNTERSET_INSTANCE"] = reflect.TypeOf((*PERF_COUNTERSET_INSTANCE)(nil)).Elem()
	types["PFD_GLYPHATTR"] = reflect.TypeOf((*PFD_GLYPHATTR)(nil)).Elem()
	types["PFIXED_INFO"] = reflect.TypeOf((*PFIXED_INFO)(nil)).Elem()
	types["PFLOAT_LONG"] = reflect.TypeOf((*PFLOAT_LONG)(nil)).Elem()
	types["PFNCALLBACK"] = reflect.TypeOf((*PFNCALLBACK)(nil)).Elem()
	types["PFORMAT_STRING"] = reflect.TypeOf((*PFORMAT_STRING)(nil)).Elem()
	types["PGENERIC_BINDING_INFO"] = reflect.TypeOf((*PGENERIC_BINDING_INFO)(nil)).Elem()
	types["PGLYPHPOS"] = reflect.TypeOf((*PGLYPHPOS)(nil)).Elem()
	types["PHUSKEY"] = reflect.TypeOf((*PHUSKEY)(nil)).Elem()
	types["PICTDESC"] = reflect.TypeOf((*PICTDESC)(nil)).Elem()
	types["PINT_PTR"] = reflect.TypeOf((*PINT_PTR)(nil)).Elem()
	types["PIP_ADAPTER_ADDRESSES"] = reflect.TypeOf((*PIP_ADAPTER_ADDRESSES)(nil)).Elem()
	types["PIP_ADAPTER_ANYCAST_ADDRESS_XP"] = reflect.TypeOf((*PIP_ADAPTER_ANYCAST_ADDRESS_XP)(nil)).Elem()
	types["PIP_ADAPTER_DNS_SERVER_ADDRESS_XP"] = reflect.TypeOf((*PIP_ADAPTER_DNS_SERVER_ADDRESS_XP)(nil)).Elem()
	types["PIP_ADAPTER_DNS_SUFFIX"] = reflect.TypeOf((*PIP_ADAPTER_DNS_SUFFIX)(nil)).Elem()
	types["PIP_ADAPTER_GATEWAY_ADDRESS_LH"] = reflect.TypeOf((*PIP_ADAPTER_GATEWAY_ADDRESS_LH)(nil)).Elem()
	types["PIP_ADAPTER_INFO"] = reflect.TypeOf((*PIP_ADAPTER_INFO)(nil)).Elem()
	types["PIP_ADAPTER_MULTICAST_ADDRESS_XP"] = reflect.TypeOf((*PIP_ADAPTER_MULTICAST_ADDRESS_XP)(nil)).Elem()
	types["PIP_ADAPTER_ORDER_MAP"] = reflect.TypeOf((*PIP_ADAPTER_ORDER_MAP)(nil)).Elem()
	types["PIP_ADAPTER_PREFIX_XP"] = reflect.TypeOf((*PIP_ADAPTER_PREFIX_XP)(nil)).Elem()
	types["PIP_ADAPTER_UNICAST_ADDRESS_LH"] = reflect.TypeOf((*PIP_ADAPTER_UNICAST_ADDRESS_LH)(nil)).Elem()
	types["PIP_ADAPTER_WINS_SERVER_ADDRESS_LH"] = reflect.TypeOf((*PIP_ADAPTER_WINS_SERVER_ADDRESS_LH)(nil)).Elem()
	types["PIP_ADDR_STRING"] = reflect.TypeOf((*PIP_ADDR_STRING)(nil)).Elem()
	types["PIP_INTERFACE_INFO"] = reflect.TypeOf((*PIP_INTERFACE_INFO)(nil)).Elem()
	types["PIP_PER_ADAPTER_INFO"] = reflect.TypeOf((*PIP_PER_ADAPTER_INFO)(nil)).Elem()
	types["PIXELFORMATDESCRIPTOR"] = reflect.TypeOf((*PIXELFORMATDESCRIPTOR)(nil)).Elem()
	types["PMIB_ICMP"] = reflect.TypeOf((*PMIB_ICMP)(nil)).Elem()
	types["PMIB_ICMP_EX"] = reflect.TypeOf((*PMIB_ICMP_EX)(nil)).Elem()
	types["PMIB_IFROW"] = reflect.TypeOf((*PMIB_IFROW)(nil)).Elem()
	types["PMIB_IFTABLE"] = reflect.TypeOf((*PMIB_IFTABLE)(nil)).Elem()
	types["PMIB_IPADDRTABLE"] = reflect.TypeOf((*PMIB_IPADDRTABLE)(nil)).Elem()
	types["PMIB_IPFORWARDROW"] = reflect.TypeOf((*PMIB_IPFORWARDROW)(nil)).Elem()
	types["PMIB_IPFORWARDTABLE"] = reflect.TypeOf((*PMIB_IPFORWARDTABLE)(nil)).Elem()
	types["PMIB_IPNETROW"] = reflect.TypeOf((*PMIB_IPNETROW)(nil)).Elem()
	types["PMIB_IPNETTABLE"] = reflect.TypeOf((*PMIB_IPNETTABLE)(nil)).Elem()
	types["PMIB_IPSTATS"] = reflect.TypeOf((*PMIB_IPSTATS)(nil)).Elem()
	types["PMIB_TCP6ROW_OWNER_MODULE"] = reflect.TypeOf((*PMIB_TCP6ROW_OWNER_MODULE)(nil)).Elem()
	types["PMIB_TCPROW"] = reflect.TypeOf((*PMIB_TCPROW)(nil)).Elem()
	types["PMIB_TCPROW_OWNER_MODULE"] = reflect.TypeOf((*PMIB_TCPROW_OWNER_MODULE)(nil)).Elem()
	types["PMIB_UDP6ROW_OWNER_MODULE"] = reflect.TypeOf((*PMIB_UDP6ROW_OWNER_MODULE)(nil)).Elem()
	types["PMIB_UDPROW_OWNER_MODULE"] = reflect.TypeOf((*PMIB_UDPROW_OWNER_MODULE)(nil)).Elem()
	types["PMIDL_STUB_MESSAGE"] = reflect.TypeOf((*PMIDL_STUB_MESSAGE)(nil)).Elem()
	types["POINT"] = reflect.TypeOf((*POINT)(nil)).Elem()
	types["POINTFIX"] = reflect.TypeOf((*POINTFIX)(nil)).Elem()
	types["POINTFLOAT"] = reflect.TypeOf((*POINTFLOAT)(nil)).Elem()
	types["POINTL"] = reflect.TypeOf((*POINTL)(nil)).Elem()
	types["POINTQF"] = reflect.TypeOf((*POINTQF)(nil)).Elem()
	types["POINTS"] = reflect.TypeOf((*POINTS)(nil)).Elem()
	types["POLICY_AUDIT_EVENT_TYPE"] = reflect.TypeOf((*POLICY_AUDIT_EVENT_TYPE)(nil)).Elem()
	types["POLICY_AUDIT_SID_ARRAY"] = reflect.TypeOf((*POLICY_AUDIT_SID_ARRAY)(nil)).Elem()
	types["POLYTEXT"] = reflect.TypeOf((*POLYTEXT)(nil)).Elem()
	types["PPERFORMACE_INFORMATION"] = reflect.TypeOf((*PPERFORMACE_INFORMATION)(nil)).Elem()
	types["PPOLICY_AUDIT_EVENT_TYPE"] = reflect.TypeOf((*PPOLICY_AUDIT_EVENT_TYPE)(nil)).Elem()
	types["PPOLICY_AUDIT_SID_ARRAY"] = reflect.TypeOf((*PPOLICY_AUDIT_SID_ARRAY)(nil)).Elem()
	types["PPROCESS_MEMORY_COUNTERS"] = reflect.TypeOf((*PPROCESS_MEMORY_COUNTERS)(nil)).Elem()
	types["PPSAPI_WS_WATCH_INFORMATION"] = reflect.TypeOf((*PPSAPI_WS_WATCH_INFORMATION)(nil)).Elem()
	types["PPSAPI_WS_WATCH_INFORMATION_EX"] = reflect.TypeOf((*PPSAPI_WS_WATCH_INFORMATION_EX)(nil)).Elem()
	types["PRECTFX"] = reflect.TypeOf((*PRECTFX)(nil)).Elem()
	types["PRINTDLG"] = reflect.TypeOf((*PRINTDLG)(nil)).Elem()
	types["PRINTDLGEX"] = reflect.TypeOf((*PRINTDLGEX)(nil)).Elem()
	types["PRINTPAGERANGE"] = reflect.TypeOf((*PRINTPAGERANGE)(nil)).Elem()
	types["PRIVILEGE_SET"] = reflect.TypeOf((*PRIVILEGE_SET)(nil)).Elem()
	types["PROC"] = reflect.TypeOf((*PROC)(nil)).Elem()
	types["PROCESS_INFORMATION"] = reflect.TypeOf((*PROCESS_INFORMATION)(nil)).Elem()
	types["PROCESS_MEMORY_COUNTERS"] = reflect.TypeOf((*PROCESS_MEMORY_COUNTERS)(nil)).Elem()
	types["PROPENUMPROC"] = reflect.TypeOf((*PROPENUMPROC)(nil)).Elem()
	types["PROPENUMPROCEX"] = reflect.TypeOf((*PROPENUMPROCEX)(nil)).Elem()
	types["PROPERTYORIGIN"] = reflect.TypeOf((*PROPERTYORIGIN)(nil)).Elem()
	types["PROPSHEETCALLBACK"] = reflect.TypeOf((*PROPSHEETCALLBACK)(nil)).Elem()
	types["PROPSHEETHEADER"] = reflect.TypeOf((*PROPSHEETHEADER)(nil)).Elem()
	types["PROPSHEETHEADER_V2"] = reflect.TypeOf((*PROPSHEETHEADER_V2)(nil)).Elem()
	types["PROPSHEETPAGE"] = reflect.TypeOf((*PROPSHEETPAGE)(nil)).Elem()
	types["PROPSHEETPAGE_RESOURCE"] = reflect.TypeOf((*PROPSHEETPAGE_RESOURCE)(nil)).Elem()
	types["PROPSHEETPAGE_V4"] = reflect.TypeOf((*PROPSHEETPAGE_V4)(nil)).Elem()
	types["PRPC_MESSAGE"] = reflect.TypeOf((*PRPC_MESSAGE)(nil)).Elem()
	types["PRPC_SYNTAX_IDENTIFIER"] = reflect.TypeOf((*PRPC_SYNTAX_IDENTIFIER)(nil)).Elem()
	types["PSAPI_WS_WATCH_INFORMATION"] = reflect.TypeOf((*PSAPI_WS_WATCH_INFORMATION)(nil)).Elem()
	types["PSAPI_WS_WATCH_INFORMATION_EX"] = reflect.TypeOf((*PSAPI_WS_WATCH_INFORMATION_EX)(nil)).Elem()
	types["PSECURE_MEMORY_CACHE_CALLBACK"] = reflect.TypeOf((*PSECURE_MEMORY_CACHE_CALLBACK)(nil)).Elem()
	types["PSECURITY_DESCRIPTOR"] = reflect.TypeOf((*PSECURITY_DESCRIPTOR)(nil)).Elem()
	types["PSID"] = reflect.TypeOf((*PSID)(nil)).Elem()
	types["PSPCALLBACK"] = reflect.TypeOf((*PSPCALLBACK)(nil)).Elem()
	types["PSRWLOCK"] = reflect.TypeOf((*PSRWLOCK)(nil)).Elem()
	types["PSTR"] = reflect.TypeOf((*PSTR)(nil)).Elem()
	types["PTRIVERTEX"] = reflect.TypeOf((*PTRIVERTEX)(nil)).Elem()
	types["PULONG64"] = reflect.TypeOf((*PULONG64)(nil)).Elem()
	types["PUSHORT"] = reflect.TypeOf((*PUSHORT)(nil)).Elem()
	types["PVOID"] = reflect.TypeOf((*PVOID)(nil)).Elem()
	types["PWCHAR"] = reflect.TypeOf((*PWCHAR)(nil)).Elem()
	types["PWSTR"] = reflect.TypeOf((*PWSTR)(nil)).Elem()
	types["QITAB"] = reflect.TypeOf((*QITAB)(nil)).Elem()
	types["QOS"] = reflect.TypeOf((*QOS)(nil)).Elem()
	types["RASTERIZER_STATUS"] = reflect.TypeOf((*RASTERIZER_STATUS)(nil)).Elem()
	types["RAWHID"] = reflect.TypeOf((*RAWHID)(nil)).Elem()
	types["RAWINPUT"] = reflect.TypeOf((*RAWINPUT)(nil)).Elem()
	types["RAWINPUTDEVICE"] = reflect.TypeOf((*RAWINPUTDEVICE)(nil)).Elem()
	types["RAWINPUTDEVICELIST"] = reflect.TypeOf((*RAWINPUTDEVICELIST)(nil)).Elem()
	types["RAWINPUTHEADER"] = reflect.TypeOf((*RAWINPUTHEADER)(nil)).Elem()
	types["RAWKEYBOARD"] = reflect.TypeOf((*RAWKEYBOARD)(nil)).Elem()
	types["RAWMOUSE"] = reflect.TypeOf((*RAWMOUSE)(nil)).Elem()
	types["RECT"] = reflect.TypeOf((*RECT)(nil)).Elem()
	types["RECTFX"] = reflect.TypeOf((*RECTFX)(nil)).Elem()
	types["RECTL"] = reflect.TypeOf((*RECTL)(nil)).Elem()
	types["REFCLSID"] = reflect.TypeOf((*REFCLSID)(nil)).Elem()
	types["REFGUID"] = reflect.TypeOf((*REFGUID)(nil)).Elem()
	types["REFIID"] = reflect.TypeOf((*REFIID)(nil)).Elem()
	types["REGISTERWORDENUMPROC"] = reflect.TypeOf((*REGISTERWORDENUMPROC)(nil)).Elem()
	types["REGKIND"] = reflect.TypeOf((*REGKIND)(nil)).Elem()
	types["REGSAM"] = reflect.TypeOf((*REGSAM)(nil)).Elem()
	types["RGBQUAD"] = reflect.TypeOf((*RGBQUAD)(nil)).Elem()
	types["RGNDATA"] = reflect.TypeOf((*RGNDATA)(nil)).Elem()
	types["RGNDATAHEADER"] = reflect.TypeOf((*RGNDATAHEADER)(nil)).Elem()
	types["ROP4"] = reflect.TypeOf((*ROP4)(nil)).Elem()
	types["RPC_BINDING_HANDLE"] = reflect.TypeOf((*RPC_BINDING_HANDLE)(nil)).Elem()
	types["RPC_MESSAGE"] = reflect.TypeOf((*RPC_MESSAGE)(nil)).Elem()
	types["RPC_STATUS"] = reflect.TypeOf((*RPC_STATUS)(nil)).Elem()
	types["RPC_SYNTAX_IDENTIFIER"] = reflect.TypeOf((*RPC_SYNTAX_IDENTIFIER)(nil)).Elem()
	types["RPC_VERSION"] = reflect.TypeOf((*RPC_VERSION)(nil)).Elem()
	types["RTL_SRWLOCK"] = reflect.TypeOf((*RTL_SRWLOCK)(nil)).Elem()
	types["RUN"] = reflect.TypeOf((*RUN)(nil)).Elem()
	types["SAFEARRAY"] = reflect.TypeOf((*SAFEARRAY)(nil)).Elem()
	types["SAFEARRAYBOUND"] = reflect.TypeOf((*SAFEARRAYBOUND)(nil)).Elem()
	types["SAFER_LEVEL_HANDLE"] = reflect.TypeOf((*SAFER_LEVEL_HANDLE)(nil)).Elem()
	types["SCODE"] = reflect.TypeOf((*SCODE)(nil)).Elem()
	types["SCROLLBARINFO"] = reflect.TypeOf((*SCROLLBARINFO)(nil)).Elem()
	types["SCROLLINFO"] = reflect.TypeOf((*SCROLLINFO)(nil)).Elem()
	types["SC_HANDLE"] = reflect.TypeOf((*SC_HANDLE)(nil)).Elem()
	types["SC_LOCK"] = reflect.TypeOf((*SC_LOCK)(nil)).Elem()
	types["SECURITY_ATTRIBUTES"] = reflect.TypeOf((*SECURITY_ATTRIBUTES)(nil)).Elem()
	types["SECURITY_CONTEXT_TRACKING_MODE"] = reflect.TypeOf((*SECURITY_CONTEXT_TRACKING_MODE)(nil)).Elem()
	types["SECURITY_DESCRIPTOR"] = reflect.TypeOf((*SECURITY_DESCRIPTOR)(nil)).Elem()
	types["SECURITY_DESCRIPTOR_CONTROL"] = reflect.TypeOf((*SECURITY_DESCRIPTOR_CONTROL)(nil)).Elem()
	types["SECURITY_IMPERSONATION_LEVEL"] = reflect.TypeOf((*SECURITY_IMPERSONATION_LEVEL)(nil)).Elem()
	types["SECURITY_INFORMATION"] = reflect.TypeOf((*SECURITY_INFORMATION)(nil)).Elem()
	types["SECURITY_QUALITY_OF_SERVICE"] = reflect.TypeOf((*SECURITY_QUALITY_OF_SERVICE)(nil)).Elem()
	types["SENDASYNCPROC"] = reflect.TypeOf((*SENDASYNCPROC)(nil)).Elem()
	types["SERVICETYPE"] = reflect.TypeOf((*SERVICETYPE)(nil)).Elem()
	types["SERVICE_STATUS"] = reflect.TypeOf((*SERVICE_STATUS)(nil)).Elem()
	types["SERVICE_STATUS_HANDLE"] = reflect.TypeOf((*SERVICE_STATUS_HANDLE)(nil)).Elem()
	types["SE_OBJECT_TYPE"] = reflect.TypeOf((*SE_OBJECT_TYPE)(nil)).Elem()
	types["SHITEMID"] = reflect.TypeOf((*SHITEMID)(nil)).Elem()
	types["SHORT"] = reflect.TypeOf((*SHORT)(nil)).Elem()
	types["SHREGDEL_FLAGS"] = reflect.TypeOf((*SHREGDEL_FLAGS)(nil)).Elem()
	types["SHREGENUM_FLAGS"] = reflect.TypeOf((*SHREGENUM_FLAGS)(nil)).Elem()
	types["SID"] = reflect.TypeOf((*SID)(nil)).Elem()
	types["SID_AND_ATTRIBUTES"] = reflect.TypeOf((*SID_AND_ATTRIBUTES)(nil)).Elem()
	types["SID_IDENTIFIER_AUTHORITY"] = reflect.TypeOf((*SID_IDENTIFIER_AUTHORITY)(nil)).Elem()
	types["SIZE"] = reflect.TypeOf((*SIZE)(nil)).Elem()
	types["SIZEL"] = reflect.TypeOf((*SIZEL)(nil)).Elem()
	types["SIZE_T"] = reflect.TypeOf((*SIZE_T)(nil)).Elem()
	types["SOCKADDR"] = reflect.TypeOf((*SOCKADDR)(nil)).Elem()
	types["SOCKET"] = reflect.TypeOf((*SOCKET)(nil)).Elem()
	types["SOCKET_ADDRESS"] = reflect.TypeOf((*SOCKET_ADDRESS)(nil)).Elem()
	types["STARTUPINFO"] = reflect.TypeOf((*STARTUPINFO)(nil)).Elem()
	types["STROBJ"] = reflect.TypeOf((*STROBJ)(nil)).Elem()
	types["STRRET"] = reflect.TypeOf((*STRRET)(nil)).Elem()
	types["STYLEBUF"] = reflect.TypeOf((*STYLEBUF)(nil)).Elem()
	types["SUBCLASSPROC"] = reflect.TypeOf((*SUBCLASSPROC)(nil)).Elem()
	types["SURFOBJ"] = reflect.TypeOf((*SURFOBJ)(nil)).Elem()
	types["SYSKIND"] = reflect.TypeOf((*SYSKIND)(nil)).Elem()
	types["SYSTEMTIME"] = reflect.TypeOf((*SYSTEMTIME)(nil)).Elem()
	types["SYSTEM_INFO"] = reflect.TypeOf((*SYSTEM_INFO)(nil)).Elem()
	types["Sockaddr"] = reflect.TypeOf((*Sockaddr)(nil)).Elem()
	types["Socklen_t"] = reflect.TypeOf((*Socklen_t)(nil)).Elem()
	types["Status"] = reflect.TypeOf((*Status)(nil)).Elem()
	types["TASKDIALOGCALLBACK"] = reflect.TypeOf((*TASKDIALOGCALLBACK)(nil)).Elem()
	types["TASKDIALOGCONFIG"] = reflect.TypeOf((*TASKDIALOGCONFIG)(nil)).Elem()
	types["TASKDIALOG_BUTTON"] = reflect.TypeOf((*TASKDIALOG_BUTTON)(nil)).Elem()
	types["TASKDIALOG_COMMON_BUTTON_FLAGS"] = reflect.TypeOf((*TASKDIALOG_COMMON_BUTTON_FLAGS)(nil)).Elem()
	types["TASKDIALOG_FLAGS"] = reflect.TypeOf((*TASKDIALOG_FLAGS)(nil)).Elem()
	types["TBBUTTON"] = reflect.TypeOf((*TBBUTTON)(nil)).Elem()
	types["TCPIP_OWNER_MODULE_INFO_CLASS"] = reflect.TypeOf((*TCPIP_OWNER_MODULE_INFO_CLASS)(nil)).Elem()
	types["TCP_TABLE_CLASS"] = reflect.TypeOf((*TCP_TABLE_CLASS)(nil)).Elem()
	types["TEXTMETRIC"] = reflect.TypeOf((*TEXTMETRIC)(nil)).Elem()
	types["THEMESIZE"] = reflect.TypeOf((*THEMESIZE)(nil)).Elem()
	types["THREAD_START_ROUTINE"] = reflect.TypeOf((*THREAD_START_ROUTINE)(nil)).Elem()
	types["TIMERPROC"] = reflect.TypeOf((*TIMERPROC)(nil)).Elem()
	types["TITLEBARINFO"] = reflect.TypeOf((*TITLEBARINFO)(nil)).Elem()
	types["TOKEN_GROUPS"] = reflect.TypeOf((*TOKEN_GROUPS)(nil)).Elem()
	types["TOKEN_PRIVILEGES"] = reflect.TypeOf((*TOKEN_PRIVILEGES)(nil)).Elem()
	types["TOUCHINPUT"] = reflect.TypeOf((*TOUCHINPUT)(nil)).Elem()
	types["TPMPARAMS"] = reflect.TypeOf((*TPMPARAMS)(nil)).Elem()
	types["TRACKMOUSEEVENT"] = reflect.TypeOf((*TRACKMOUSEEVENT)(nil)).Elem()
	types["TRIVERTEX"] = reflect.TypeOf((*TRIVERTEX)(nil)).Elem()
	types["TRUSTEE"] = reflect.TypeOf((*TRUSTEE)(nil)).Elem()
	types["TRUSTEE_FORM"] = reflect.TypeOf((*TRUSTEE_FORM)(nil)).Elem()
	types["TRUSTEE_TYPE"] = reflect.TypeOf((*TRUSTEE_TYPE)(nil)).Elem()
	types["TUNNEL_TYPE"] = reflect.TypeOf((*TUNNEL_TYPE)(nil)).Elem()
	types["Time_t"] = reflect.TypeOf((*Time_t)(nil)).Elem()
	types["Timeval"] = reflect.TypeOf((*Timeval)(nil)).Elem()
	types["UCHAR"] = reflect.TypeOf((*UCHAR)(nil)).Elem()
	types["UDATE"] = reflect.TypeOf((*UDATE)(nil)).Elem()
	types["UDP_TABLE_CLASS"] = reflect.TypeOf((*UDP_TABLE_CLASS)(nil)).Elem()
	types["UINT"] = reflect.TypeOf((*UINT)(nil)).Elem()
	types["UINT8"] = reflect.TypeOf((*UINT8)(nil)).Elem()
	types["UINT_PTR"] = reflect.TypeOf((*UINT_PTR)(nil)).Elem()
	types["ULARGE_INTEGER"] = reflect.TypeOf((*ULARGE_INTEGER)(nil)).Elem()
	types["ULONG"] = reflect.TypeOf((*ULONG)(nil)).Elem()
	types["ULONG64"] = reflect.TypeOf((*ULONG64)(nil)).Elem()
	types["ULONGLONG"] = reflect.TypeOf((*ULONGLONG)(nil)).Elem()
	types["ULONG_PTR"] = reflect.TypeOf((*ULONG_PTR)(nil)).Elem()
	types["UPDATELAYEREDWINDOWINFO"] = reflect.TypeOf((*UPDATELAYEREDWINDOWINFO)(nil)).Elem()
	types["URLIS"] = reflect.TypeOf((*URLIS)(nil)).Elem()
	types["USER_MARSHAL_FREEING_ROUTINE"] = reflect.TypeOf((*USER_MARSHAL_FREEING_ROUTINE)(nil)).Elem()
	types["USER_MARSHAL_MARSHALLING_ROUTINE"] = reflect.TypeOf((*USER_MARSHAL_MARSHALLING_ROUTINE)(nil)).Elem()
	types["USER_MARSHAL_ROUTINE_QUADRUPLE"] = reflect.TypeOf((*USER_MARSHAL_ROUTINE_QUADRUPLE)(nil)).Elem()
	types["USER_MARSHAL_SIZING_ROUTINE"] = reflect.TypeOf((*USER_MARSHAL_SIZING_ROUTINE)(nil)).Elem()
	types["USER_MARSHAL_UNMARSHALLING_ROUTINE"] = reflect.TypeOf((*USER_MARSHAL_UNMARSHALLING_ROUTINE)(nil)).Elem()
	types["USHORT"] = reflect.TypeOf((*USHORT)(nil)).Elem()
	types["VARIANT"] = reflect.TypeOf((*VARIANT)(nil)).Elem()
	types["VARIANTARG"] = reflect.TypeOf((*VARIANTARG)(nil)).Elem()
	types["VARIANT_BOOL"] = reflect.TypeOf((*VARIANT_BOOL)(nil)).Elem()
	types["VARTYPE"] = reflect.TypeOf((*VARTYPE)(nil)).Elem()
	types["WCHAR"] = reflect.TypeOf((*WCHAR)(nil)).Elem()
	types["WCRANGE"] = reflect.TypeOf((*WCRANGE)(nil)).Elem()
	types["WGLSWAP"] = reflect.TypeOf((*WGLSWAP)(nil)).Elem()
	types["WINDOWINFO"] = reflect.TypeOf((*WINDOWINFO)(nil)).Elem()
	types["WINDOWPLACEMENT"] = reflect.TypeOf((*WINDOWPLACEMENT)(nil)).Elem()
	types["WINEVENTPROC"] = reflect.TypeOf((*WINEVENTPROC)(nil)).Elem()
	types["WINSTAENUMPROC"] = reflect.TypeOf((*WINSTAENUMPROC)(nil)).Elem()
	types["WNDCLASS"] = reflect.TypeOf((*WNDCLASS)(nil)).Elem()
	types["WNDCLASSEX"] = reflect.TypeOf((*WNDCLASSEX)(nil)).Elem()
	types["WNDENUMPROC"] = reflect.TypeOf((*WNDENUMPROC)(nil)).Elem()
	types["WNDPROC"] = reflect.TypeOf((*WNDPROC)(nil)).Elem()
	types["WORD"] = reflect.TypeOf((*WORD)(nil)).Elem()
	types["WPARAM"] = reflect.TypeOf((*WPARAM)(nil)).Elem()
	types["WSABUF"] = reflect.TypeOf((*WSABUF)(nil)).Elem()
	types["WSACOMPLETION"] = reflect.TypeOf((*WSACOMPLETION)(nil)).Elem()
	types["WSACOMPLETIONTYPE"] = reflect.TypeOf((*WSACOMPLETIONTYPE)(nil)).Elem()
	types["WSADATA"] = reflect.TypeOf((*WSADATA)(nil)).Elem()
	types["WSAECOMPARATOR"] = reflect.TypeOf((*WSAECOMPARATOR)(nil)).Elem()
	types["WSAESETSERVICEOP"] = reflect.TypeOf((*WSAESETSERVICEOP)(nil)).Elem()
	types["WSAEVENT"] = reflect.TypeOf((*WSAEVENT)(nil)).Elem()
	types["WSAMSG"] = reflect.TypeOf((*WSAMSG)(nil)).Elem()
	types["WSANAMESPACE_INFO"] = reflect.TypeOf((*WSANAMESPACE_INFO)(nil)).Elem()
	types["WSANETWORKEVENTS"] = reflect.TypeOf((*WSANETWORKEVENTS)(nil)).Elem()
	types["WSANSCLASSINFO"] = reflect.TypeOf((*WSANSCLASSINFO)(nil)).Elem()
	types["WSAPOLLFD"] = reflect.TypeOf((*WSAPOLLFD)(nil)).Elem()
	types["WSAPROTOCOLCHAIN"] = reflect.TypeOf((*WSAPROTOCOLCHAIN)(nil)).Elem()
	types["WSAPROTOCOL_INFO"] = reflect.TypeOf((*WSAPROTOCOL_INFO)(nil)).Elem()
	types["WSAQUERYSET"] = reflect.TypeOf((*WSAQUERYSET)(nil)).Elem()
	types["WSASERVICECLASSINFO"] = reflect.TypeOf((*WSASERVICECLASSINFO)(nil)).Elem()
	types["WSAVERSION"] = reflect.TypeOf((*WSAVERSION)(nil)).Elem()
	types["XFORM"] = reflect.TypeOf((*XFORM)(nil)).Elem()
	types["XFORML"] = reflect.TypeOf((*XFORML)(nil)).Elem()
	types["XFORMOBJ"] = reflect.TypeOf((*XFORMOBJ)(nil)).Elem()
	types["XLATEOBJ"] = reflect.TypeOf((*XLATEOBJ)(nil)).Elem()
	types["XLAT_SIDE"] = reflect.TypeOf((*XLAT_SIDE)(nil)).Elem()
	types["XMIT_HELPER_ROUTINE"] = reflect.TypeOf((*XMIT_HELPER_ROUTINE)(nil)).Elem()
	types["XMIT_ROUTINE_QUINTUPLE"] = reflect.TypeOf((*XMIT_ROUTINE_QUINTUPLE)(nil)).Elem()
	types["bool"] = reflect.TypeOf((*bool)(nil)).Elem()
	types["byte"] = reflect.TypeOf((*byte)(nil)).Elem()
	types["float32"] = reflect.TypeOf((*float32)(nil)).Elem()
	types["float64"] = reflect.TypeOf((*float64)(nil)).Elem()
	types["int"] = reflect.TypeOf((*int)(nil)).Elem()
	types["int16"] = reflect.TypeOf((*int16)(nil)).Elem()
	types["int32"] = reflect.TypeOf((*int32)(nil)).Elem()
	types["int8"] = reflect.TypeOf((*int8)(nil)).Elem()
	types["string"] = reflect.TypeOf((*string)(nil)).Elem()
	types["uint16"] = reflect.TypeOf((*uint16)(nil)).Elem()
	types["uint32"] = reflect.TypeOf((*uint32)(nil)).Elem()
	types["uintptr"] = reflect.TypeOf((*uintptr)(nil)).Elem()
}

type ABC struct {
	AbcA int32
	AbcB UINT
	AbcC int32
}
type ABCFLOAT struct {
	AbcfA FLOAT
	AbcfB FLOAT
	AbcfC FLOAT
}
type ABORTPROC func(unnamed0 HDC, unnamed1 int32) BOOL
type ACCEL struct {
	FVirt BYTE
	Key   WORD
	Cmd   WORD
}
type ACCESS_MASK uint32
type ACL struct {
	AclRevision byte
	Sbz1        byte
	AclSize     uint16
	AceCount    uint16
	Sbz2        uint16
}
type ADDRINFO struct {
	Ai_flags     int32
	Ai_family    int32
	Ai_socktype  int32
	Ai_protocol  int32
	Ai_addrlen   SIZE_T
	Ai_canonname PWSTR
	Ai_addr      uintptr // struct sockaddr *
	Ai_next      *ADDRINFO
}
type ADDRINFOEX struct {
	Ai_flags     int32
	Ai_family    int32
	Ai_socktype  int32
	Ai_protocol  int32
	Ai_addrlen   SIZE_T
	Ai_canonname PWSTR
	Ai_addr      uintptr // struct sockaddr*
	Ai_blob      LPVOID
	Ai_bloblen   SIZE_T
	Ai_provider  LPGUID
	Ai_next      *ADDRINFOEX
}
type AFPROTOCOLS struct {
	IAddressFamily INT
	IProtocol      INT
}
type ALG_ID uint32
type ALTTABINFO struct {
	CbSize    uint32
	CItems    int32
	CColumns  int32
	CRows     int32
	IColFocus int32
	IRowFocus int32
	CxItem    int32
	CyItem    int32
	PtStart   POINT
}
type ARRAY_INFO struct {
	Dimension             int32
	BufferConformanceMark *uint32
	BufferVarianceMark    *uint32
	MaxCountArray         *uint32
	OffsetArray           *uint32
	ActualCountArray      *uint32
}
type ATOM uint16
type AUDIT_POLICY_INFORMATION struct {
	AuditSubCategoryGuid GUID
	AuditingInformation  ULONG
	AuditCategoryGuid    GUID
}
type AUXCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	WTechnology    WORD
	WReserved1     WORD
	DwSupport      DWORD
}
type BITMAP struct {
	BmType       LONG
	BmWidth      LONG
	BmHeight     LONG
	BmWidthBytes LONG
	BmPlanes     WORD
	BmBitsPixel  WORD
	BmBits       LPVOID
}
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}
type BITMAPINFOHEADER struct {
	BiSize          DWORD
	BiWidth         LONG
	BiHeight        LONG
	BiPlanes        WORD
	BiBitCount      WORD
	BiCompression   DWORD
	BiSizeImage     DWORD
	BiXPelsPerMeter LONG
	BiYPelsPerMeter LONG
	BiClrUsed       DWORD
	BiClrImportant  DWORD
}
type BLENDOBJ struct {
	BlendFunction BLENDFUNCTION
}
type BOOLEAN byte
type BP_ANIMATIONPARAMS struct {
	CbSize     DWORD
	DwFlags    DWORD
	Style      BP_ANIMATIONSTYLE
	DwDuration DWORD
}
type BP_PAINTPARAMS struct {
	CbSize  DWORD
	DwFlags DWORD
	PrcExclude/*const*/ *RECT
	PBlendFunction/*const*/ *BLENDFUNCTION
}
type BRUSHOBJ struct {
	ISolidColor ULONG
	PvRbrush    PVOID
	FlColorType FLONG
}
type BSMINFO struct {
	CbSize uint32
	Hdesk  HDESK
	Hwnd   HWND
	Luid   LUID
}
type BSTR *OLECHAR
type BYTE byte
type CANDIDATEFORM struct {
	DwIndex      DWORD
	DwStyle      DWORD
	PtCurrentPos POINT
	RcArea       RECT
}
type CANDIDATELIST struct {
	DwSize      DWORD
	DwStyle     DWORD
	DwCount     DWORD
	DwSelection DWORD
	DwPageStart DWORD
	DwPageSize  DWORD
	DwOffset    [1]DWORD
}
type CCHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type CERT_EXTENSION struct {
	PszObjId  LPSTR
	FCritical BOOL
	Value     CRYPT_OBJID_BLOB
}
type CERT_NAME_BLOB CRYPTOAPI_BLOB_
type CERT_RDN_VALUE_BLOB CRYPTOAPI_BLOB_
type CHAR byte
type CHARSETINFO struct {
	CiCharset UINT
	CiACP     UINT
	Fs        FONTSIGNATURE
}
type CHOOSECOLOR struct {
	LStructSize    DWORD
	HwndOwner      HWND
	HInstance      HWND
	RgbResult      COLORREF
	LpCustColors   *COLORREF
	Flags          DWORD
	LCustData      LPARAM
	LpfnHook       uintptr // LPCCHOOKPROC
	LpTemplateName LPCWSTR
}
type CHOOSEFONT struct {
	LStructSize            DWORD
	HwndOwner              HWND
	HDC                    HDC
	LpLogFont              LPLOGFONT
	IPointSize             INT
	Flags                  DWORD
	RgbColors              COLORREF
	LCustData              LPARAM
	LpfnHook               uintptr // LPCFHOOKPROC
	LpTemplateName         LPCWSTR
	HInstance              HINSTANCE
	LpszStyle              LPWSTR
	NFontType              WORD
	___MISSING_ALIGNMENT__ WORD
	NSizeMin               INT
	NSizeMax               INT
}
type CIEXYZ struct {
	CiexyzX FXPT2DOT30
	CiexyzY FXPT2DOT30
	CiexyzZ FXPT2DOT30
}
type CIEXYZTRIPLE struct {
	CiexyzRed   CIEXYZ
	CiexyzGreen CIEXYZ
	CiexyzBlue  CIEXYZ
}
type CLIPLINE struct {
	PtfxA       POINTFIX
	PtfxB       POINTFIX
	LStyleState LONG
	C           ULONG
	Arun        [1]RUN
}
type CLIPOBJ struct {
	IUniq        ULONG
	RclBounds    RECTL
	IDComplexity BYTE
	IFComplexity BYTE
	IMode        BYTE
	FjOptions    BYTE
}
type CLSID GUID
type COLOR16 USHORT
type COLORADJUSTMENT struct {
	CaSize            WORD
	CaFlags           WORD
	CaIlluminantIndex WORD
	CaRedGamma        WORD
	CaGreenGamma      WORD
	CaBlueGamma       WORD
	CaReferenceBlack  WORD
	CaReferenceWhite  WORD
	CaContrast        SHORT
	CaBrightness      SHORT
	CaColorfulness    SHORT
	CaRedGreenTint    SHORT
}
type COLORMAP struct {
	From COLORREF
	To   COLORREF
}
type COLORREF uint32
type COMBOBOXINFO struct {
	CbSize      uint32
	RcItem      RECT
	RcButton    RECT
	StateButton uint32
	HwndCombo   HWND
	HwndItem    HWND
	HwndList    HWND
}
type COMM_FAULT_OFFSETS struct {
	CommOffset  int16
	FaultOffset int16
}
type COMPOSITIONFORM struct {
	DwStyle      DWORD
	PtCurrentPos POINT
	RcArea       RECT
}
type CONVCONTEXT struct {
	Cb         uint32
	WFlags     uint32
	WCountryID uint32
	ICodePage  int32
	DwLangID   uint32
	DwSecurity uint32
	Qos        SECURITY_QUALITY_OF_SERVICE
}
type CONVINFO struct {
	Cb            uint32
	HUser         *uint32 // DWORD_PTR
	HConvPartner  HCONV
	HszSvcPartner HSZ
	HszServiceReq HSZ
	HszTopic      HSZ
	HszItem       HSZ
	WFmt          uint32 // UINT
	WType         uint32 // UINT
	WStatus       uint32 // UINT
	WConvst       uint32 // UINT
	WLastError    uint32 // UINT
	HConvList     HCONVLIST
	ConvCtxt      CONVCONTEXT
	Hwnd          HWND
	HwndPartner   HWND
}
type CREDENTIAL struct {
	Flags              uint32
	Type               uint32
	TargetName         LPWSTR
	Comment            LPWSTR
	LastWritten        FILETIME
	CredentialBlobSize uint32
	CredentialBlob     *byte
	Persist            uint32
	AttributeCount     uint32
	Attributes         *CREDENTIAL_ATTRIBUTE
	TargetAlias        LPWSTR
	UserName           LPWSTR
}
type CREDENTIAL_ATTRIBUTE struct {
	Keyword   LPWSTR
	Flags     uint32
	ValueSize uint32
	Value     *byte
}
type CRL_CONTEXT struct {
	DwCertEncodingType DWORD
	PbCrlEncoded       *BYTE
	CbCrlEncoded       DWORD
	PCrlInfo           PCRL_INFO
	HCertStore         HCERTSTORE
}
type CRL_ENTRY struct {
	SerialNumber   CRYPT_INTEGER_BLOB
	RevocationDate FILETIME
	CExtension     DWORD
	RgExtension    PCERT_EXTENSION
}
type CRL_INFO struct {
	DwVersion          DWORD
	SignatureAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	Issuer             CERT_NAME_BLOB
	ThisUpdate         FILETIME
	NextUpdate         FILETIME
	CCRLEntry          DWORD
	RgCRLEntry         PCRL_ENTRY
	CExtension         DWORD
	RgExtension        PCERT_EXTENSION
}
type CRYPTOAPI_BLOB_ struct {
	CbData DWORD
	PbData *BYTE
}
type CRYPT_ALGORITHM_IDENTIFIER struct {
	PszObjId   LPSTR
	Parameters CRYPT_OBJID_BLOB
}
type CRYPT_DATA_BLOB CRYPTOAPI_BLOB_
type CRYPT_HASH_BLOB CRYPTOAPI_BLOB_
type CRYPT_INTEGER_BLOB CRYPTOAPI_BLOB_
type CRYPT_OBJID_BLOB CRYPTOAPI_BLOB_
type CSADDR_INFO struct {
	LocalAddr   SOCKET_ADDRESS
	RemoteAddr  SOCKET_ADDRESS
	ISocketType INT
	IProtocol   INT
}
type CS_TAG_GETTING_ROUTINE func(hBinding RPC_BINDING_HANDLE, fServerSide int32, pulSendingTag *uint32, pulDesiredReceivingTag *uint32, pulReceivingTag *uint32, pStatus *Error_status_t)
type CS_TYPE_FROM_NETCS_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, pNetworkData *byte, ulNetworkDataLength uint32, ulLocalBufferSize uint32, pLocalData uintptr, pulLocalDataLength *uint32, pStatus *Error_status_t)
type CS_TYPE_LOCAL_SIZE_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, ulNetworkBufferSize uint32, conversionType *IDL_CS_CONVERT, pulLocalBufferSize *uint32, pStatus *Error_status_t)
type CS_TYPE_NET_SIZE_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, ulLocalBufferSize uint32, conversionType *IDL_CS_CONVERT, pulNetworkBufferSize *uint32, pStatus *Error_status_t)
type CS_TYPE_TO_NETCS_ROUTINE func(hBinding RPC_BINDING_HANDLE, ulNetworkCodeSet uint32, pLocalData uintptr, ulLocalDataLength uint32, pNetworkData *byte, pulNetworkDataLength *uint32, pStatus *Error_status_t)
type CURSORINFO struct {
	CbSize      uint32
	Flags       uint32
	HCursor     HCURSOR
	PtScreenPos POINT
}
type CUSTDATA struct {
	CCustData   DWORD
	PrgCustData *CUSTDATAITEM
}
type CUSTDATAITEM struct {
	Guid     GUID
	VarValue VARIANTARG
}
type DACOMPARE func(p1 uintptr, p2 uintptr, lParam LPARAM) int32
type DAENUMCALLBACK func(p uintptr, pData uintptr) int32
type DATE float64
type DATETIME struct {
	Year  uint16 // WORD
	Month uint16 // WORD
	Day   uint16 // WORD
	Hour  uint16 // WORD
	Min   uint16 // WORD
	Sec   uint16 // WORD
}
type DECIMAL struct {
	WReserved USHORT
	union1    [2]byte
	Hi32      ULONG
	union2    [8]byte
}

func (this *DECIMAL) Scale() *BYTE {
	return (*BYTE)(unsafe.Pointer(&this.union1[0]))
}
func (this *DECIMAL) Sign() *BYTE {
	return (*BYTE)(unsafe.Pointer(&this.union1[1]))
}
func (this *DECIMAL) Signscale() *USHORT {
	return (*USHORT)(unsafe.Pointer(&this.union1[0]))
}
func (this *DECIMAL) Lo32() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union2[0]))
}
func (this *DECIMAL) Mid32() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union2[4]))
}
func (this *DECIMAL) Lo64() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union2[0]))
}

type DESIGNVECTOR struct {
	DvReserved DWORD
	DvNumAxes  DWORD
	DvValues   [MM_MAX_NUMAXES]LONG
}
type DESKTOPENUMPROC func(lpszDesktop LPWSTR, lParam LPARAM) BOOL
type DEVMODE struct {
	DmDeviceName       [CCHDEVICENAME]uint16
	DmSpecVersion      uint16
	DmDriverVersion    uint16
	DmSize             uint16
	DmDriverExtra      uint16
	DmFields           uint32
	DmOrientation      int16
	DmPaperSize        int16
	DmPaperLength      int16
	DmPaperWidth       int16
	DmScale            int16
	DmCopies           int16
	DmDefaultSource    int16
	DmPrintQuality     int16
	DmColor            int16
	DmDuplex           int16
	DmYResolution      int16
	DmTTOption         int16
	DmCollate          int16
	DmFormName         [CCHFORMNAME]uint16
	DmLogPixels        uint16
	DmBitsPerPel       uint32
	DmPelsWidth        uint32
	DmPelsHeight       uint32
	DmDisplayFlags     uint32
	DmDisplayFrequency uint32
	DmICMMethod        uint32
	DmICMIntent        uint32
	DmMediaType        uint32
	DmDitherType       uint32
	DmReserved1        uint32
	DmReserved2        uint32
	DmPanningWidth     uint32
	DmPanningHeight    uint32
}
type DHPDEV HANDLE
type DHSURF HANDLE
type DISPID LONG
type DISPLAY_DEVICE struct {
	Cb           uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	StateFlags   uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}
type DISPPARAMS struct {
	Rgvarg            *VARIANTARG
	RgdispidNamedArgs *DISPID
	CArgs             UINT
	CNamedArgs        UINT
}
type DLGPROC func(hwndDlg HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) int32
type DLGTEMPLATE struct {
	storage [18]byte
}

func (this *DLGTEMPLATE) Style() *uint32 {
	return (*uint32)(unsafe.Pointer(&this.storage[0]))
}
func (this *DLGTEMPLATE) DwExtendedStyle() *uint32 {
	return (*uint32)(unsafe.Pointer(&this.storage[4]))
}
func (this *DLGTEMPLATE) Cdit() *uint16 {
	return (*uint16)(unsafe.Pointer(&this.storage[8]))
}
func (this *DLGTEMPLATE) X() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[10]))
}
func (this *DLGTEMPLATE) Y() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[12]))
}
func (this *DLGTEMPLATE) Cx() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[14]))
}
func (this *DLGTEMPLATE) Cy() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[16]))
}

type DOCINFO struct {
	CbSize       int32
	LpszDocName  LPCWSTR
	LpszOutput   LPCWSTR
	LpszDatatype LPCWSTR
	FwType       DWORD
}
type DOUBLE float64
type DRAWSTATEPROC func(hdc HDC, lData uintptr, wData uintptr, cx int32, cy int32) BOOL
type DRAWTEXTPARAMS struct {
	CbSize        uint32
	ITabLength    int32
	ILeftMargin   int32
	IRightMargin  int32
	UiLengthDrawn uint32
}
type DTBGOPTS struct {
	DwSize  DWORD
	DwFlags DWORD
	RcClip  RECT
}
type DTTOPTS struct {
	DwSize              DWORD
	DwFlags             DWORD
	CrText              COLORREF
	CrBorder            COLORREF
	CrShadow            COLORREF
	ITextShadowType     int32
	PtShadowOffset      POINT
	IBorderSize         int32
	IFontPropId         int32
	IColorPropId        int32
	IStateId            int32
	FApplyOverlay       BOOL
	IGlowSize           int32
	PfnDrawTextCallback uintptr // DTT_CALLBACK_PROC
	LParam              LPARAM
}
type DTT_CALLBACK_PROC func(hdc HDC, pszText LPWSTR, cchText int32, prc LPRECT, dwFlags UINT, lParam LPARAM) int32
type DWORD_PTR *DWORD
type EFS_CERTIFICATE_BLOB struct {
	DwCertEncodingType DWORD
	CbData             DWORD
	PbData             PBYTE
}
type ENCRYPTION_CERTIFICATE struct {
	CbTotalLength DWORD
	PUserSid      *SID
	PCertBlob     *EFS_CERTIFICATE_BLOB
}
type ENCRYPTION_CERTIFICATE_LIST struct {
	NUsers DWORD
	PUsers *PENCRYPTION_CERTIFICATE
}
type ENG_TIME_FIELDS struct {
	UsYear         USHORT
	UsMonth        USHORT
	UsDay          USHORT
	UsHour         USHORT
	UsMinute       USHORT
	UsSecond       USHORT
	UsMilliseconds USHORT
	UsWeekday      USHORT
}
type ENHMETAHEADER struct {
	IType          DWORD
	NSize          DWORD
	RclBounds      RECTL
	RclFrame       RECTL
	DSignature     DWORD
	NVersion       DWORD
	NBytes         DWORD
	NRecords       DWORD
	NHandles       WORD
	SReserved      WORD
	NDescription   DWORD
	OffDescription DWORD
	NPalEntries    DWORD
	SzlDevice      SIZEL
	SzlMillimeters SIZEL
	CbPixelFormat  DWORD
	OffPixelFormat DWORD
	BOpenGL        DWORD
	SzlMicrometers SIZEL
}
type ENHMETARECORD struct {
	IType DWORD
	NSize DWORD
	DParm [1]DWORD
}
type ENHMFENUMPROC func(hdc HDC, lpht *HANDLETABLE, lpmr /*const*/ *ENHMETARECORD, nHandles int32, data LPARAM) int32
type ENUMLOGFONTEX struct {
	ElfLogFont  LOGFONT
	ElfFullName [LF_FULLFACESIZE]WCHAR
	ElfStyle    [LF_FACESIZE]WCHAR
	ElfScript   [LF_FACESIZE]WCHAR
}
type ENUMLOGFONTEXDV struct {
	ElfEnumLogfontEx ENUMLOGFONTEX
	ElfDesignVector  DESIGNVECTOR
}
type ENUMRESLANGPROC func(hModule HMODULE, lpType string, lpName string, wLanguage WORD, lParam LONG_PTR) BOOL
type ENUM_PAGE_FILE_INFORMATION struct {
	Cb         DWORD
	Reserved   DWORD
	TotalSize  SIZE_T
	TotalInUse SIZE_T
	PeakUsage  SIZE_T
}
type EXCEPINFO struct {
	WCode             WORD
	WReserved         WORD
	BstrSource        BSTR
	BstrDescription   BSTR
	BstrHelpFile      BSTR
	DwHelpContext     DWORD
	PvReserved        PVOID
	PfnDeferredFillIn uintptr
	Scode             SCODE
}

func (this *EXCEPINFO) DeferredFillIn() func(unnamed0 *EXCEPINFO) HRESULT {
	return func(unnamed0 *EXCEPINFO) HRESULT {
		ret := syscall3(this.PfnDeferredFillIn, 1,
			uintptr(unsafe.Pointer(unnamed0)),
			0,
			0)
		return HRESULT(ret)
	}
}

type EXPLICIT_ACCESS struct {
	GrfAccessPermissions DWORD
	GrfAccessMode        ACCESS_MODE
	GrfInheritance       DWORD
	Trustee              TRUSTEE
}
type EXPR_EVAL func(unnamed0 *MIDL_STUB_MESSAGE)
type Error_status_t uint32
type FARPROC func() INT_PTR
type FD_GLYPHATTR struct {
	CjThis     ULONG
	CGlyphs    ULONG
	IMode      ULONG
	AGlyphAttr [1]BYTE
}
type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}
type FINDREPLACE struct {
	LStructSize      DWORD
	HwndOwner        HWND
	HInstance        HINSTANCE
	Flags            DWORD
	LpstrFindWhat    LPWSTR
	LpstrReplaceWith LPWSTR
	WFindWhatLen     WORD
	WReplaceWithLen  WORD
	LCustData        LPARAM
	LpfnHook         uintptr // LPFRHOOKPROC
	LpTemplateName   LPCWSTR
}
type FIX LONG
type FIXED struct {
	Fract WORD
	Value int16
}
type FIXED_INFO_W2KSP1 struct {
	HostName         [MAX_HOSTNAME_LEN + 4]CHAR
	DomainName       [MAX_DOMAIN_NAME_LEN + 4]CHAR
	CurrentDnsServer PIP_ADDR_STRING
	DnsServerList    IP_ADDR_STRING
	NodeType         UINT
	ScopeId          [MAX_SCOPE_ID_LEN + 4]CHAR
	EnableRouting    UINT
	EnableProxy      UINT
	EnableDns        UINT
}
type FLASHWINFO struct {
	CbSize    uint32 // UINT
	Hwnd      HWND
	DwFlags   uint32
	UCount    uint32 // UINT
	DwTimeout uint32
}
type FLOAT float32
type FLOATL FLOAT
type FLOAT_LONG struct {
	storage [4]byte
}

func (this *FLOAT_LONG) E() *FLOATL {
	return (*FLOATL)(unsafe.Pointer(&this.storage[0]))
}
func (this *FLOAT_LONG) L() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[0]))
}

type FLONG uint32
type FLOWSPEC struct {
	TokenRate          ULONG
	TokenBucketSize    ULONG
	PeakBandwidth      ULONG
	Latency            ULONG
	DelayVariation     ULONG
	ServiceType        SERVICETYPE
	MaxSduSize         ULONG
	MinimumPolicedSize ULONG
}
type FONTDESC struct {
	CbSizeofstruct UINT
	LpstrName      LPOLESTR
	CySize         CY
	SWeight        SHORT
	SCharset       SHORT
	FItalic        BOOL
	FUnderline     BOOL
	FStrikethrough BOOL
}
type FONTENUMPROC func(unnamed0 /*const*/ *LOGFONT, unnamed1 /*const*/ *TEXTMETRIC, unnamed2 DWORD, unnamed3 LPARAM) int32
type FONTINFO struct {
	CjThis           ULONG
	FlCaps           FLONG
	CGlyphsSupported ULONG
	CjMaxGlyph1      ULONG
	CjMaxGlyph4      ULONG
	CjMaxGlyph8      ULONG
	CjMaxGlyph32     ULONG
}
type FONTOBJ struct {
	IUniq        ULONG
	IFace        ULONG
	CxMax        ULONG
	FlFontType   FLONG
	ITTUniq      ULONG_PTR
	IFile        ULONG_PTR
	SizLogResPpi SIZE
	UlStyleSize  ULONG
	PvConsumer   PVOID
	PvProducer   PVOID
}
type FONTSIGNATURE struct {
	FsUsb [4]DWORD
	FsCsb [2]DWORD
}
type FULL_PTR_XLAT_TABLES struct {
	RefIdToPointer uintptr
	PointerToRefId uintptr
	NextRefId      uint32
	XlatSide       XLAT_SIDE
}
type FXPT2DOT30 int32
type GCP_RESULTS struct {
	LStructSize DWORD
	LpOutString LPWSTR
	LpOrder     *UINT
	LpDx        *int32
	LpCaretPos  *int32
	LpClass     LPSTR
	LpGlyphs    LPWSTR
	NGlyphs     UINT
	NMaxFit     int32
}
type GENERIC_BINDING_INFO struct {
	PObj      uintptr
	Size      uint32
	PfnBind   uintptr // GENERIC_BINDING_ROUTINE
	PfnUnbind uintptr // GENERIC_UNBIND_ROUTINE
}
type GENERIC_BINDING_ROUTINE func(unnamed0 uintptr) uintptr
type GENERIC_BINDING_ROUTINE_PAIR struct {
	PfnBind   uintptr // GENERIC_BINDING_ROUTINE
	PfnUnbind uintptr // GENERIC_UNBIND_ROUTINE
}
type GENERIC_MAPPING struct {
	GenericRead    ACCESS_MASK
	GenericWrite   ACCESS_MASK
	GenericExecute ACCESS_MASK
	GenericAll     ACCESS_MASK
}
type GENERIC_UNBIND_ROUTINE func(unnamed0 uintptr, unnamed1 *byte)
type GEOID LONG
type GESTURECONFIG struct {
	DwID    uint32
	DwWant  uint32
	DwBlock uint32
}
type GLYPHBITS struct {
	PtlOrigin  POINTL
	SizlBitmap SIZEL
	Aj         [1]BYTE
}
type GLYPHDEF struct {
	p uintptr
}

func (this *GLYPHDEF) Pgb() *GLYPHBITS {
	return (*GLYPHBITS)(unsafe.Pointer(this.p))
}
func (this *GLYPHDEF) Ppo() *PATHOBJ {
	return (*PATHOBJ)(unsafe.Pointer(this.p))
}

type GLYPHMETRICS struct {
	GmBlackBoxX     UINT
	GmBlackBoxY     UINT
	GmptGlyphOrigin POINT
	GmCellIncX      int16
	GmCellIncY      int16
}
type GLYPHMETRICSFLOAT struct {
	GmfBlackBoxX     float32
	GmfBlackBoxY     float32
	GmfptGlyphOrigin POINTFLOAT
	GmfCellIncX      float32
	GmfCellIncY      float32
}
type GLYPHPOS struct {
	Hg   HGLYPH
	Pgdf *GLYPHDEF
	Ptl  POINTL
}
type GLYPHSET struct {
	CbThis           DWORD
	FlAccel          DWORD
	CGlyphsSupported DWORD
	CRanges          DWORD
	Ranges           [1]WCRANGE
}
type GLbitfield uint32
type GLboolean byte
type GLbyte int8
type GLclampd float64
type GLclampf float32
type GLdouble float64
type GLenum uint32
type GLfloat float32
type GLint int32
type GLshort int16
type GLsizei uint32
type GLubyte uint8
type GLuint uint32
type GLushort uint16
type GOBJENUMPROC func(unnamed0 LPVOID, unnamed1 LPARAM) int32
type GRAYSTRINGPROC func(hdc HDC, lParam LPARAM, cchData int) BOOL
type GROUP uint32
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}
type GUITHREADINFO struct {
	CbSize        uint32
	Flags         uint32
	HwndActive    HWND
	HwndFocus     HWND
	HwndCapture   HWND
	HwndMenuOwner HWND
	HwndMoveSize  HWND
	HwndCaret     HWND
	RcCaret       RECT
}
type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr // DebugEventProc
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}
type GdiplusStartupOutput struct {
	storage [2]uintptr
}
type HACCEL HANDLE
type HANDLE uintptr
type HANDLER_FUNCTION_EX func(dwControl uint32, dwEventType uint32, lpEventData uintptr, lpContext uintptr) uint32
type HANDLETABLE struct {
	ObjectHandle [1]HGDIOBJ
}
type HANIMATIONBUFFER HANDLE
type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
}
type HBITMAP HGDIOBJ
type HBRUSH HGDIOBJ
type HCERTSTORE uintptr
type HCOLORSPACE HANDLE
type HCONV HANDLE
type HCONVLIST HANDLE
type HCRYPTHASH uintptr
type HCRYPTKEY uintptr
type HCRYPTPROV uintptr
type HCURSOR HANDLE
type HDC HANDLE
type HDDEDATA HANDLE
type HDESK HANDLE
type HDEV HANDLE
type HDEVNOTIFY uintptr
type HDPA uintptr
type HDRVR HANDLE
type HDSA uintptr
type HDWP HANDLE
type HELPINFO struct {
	CbSize       UINT
	IContextType int32
	ICtrlId      int32
	HItemHandle  HANDLE
	DwContextId  DWORD_PTR
	MousePos     POINT
}
type HENHMETAFILE HANDLE
type HFONT HGDIOBJ
type HGDIOBJ HANDLE
type HGESTUREINFO HANDLE
type HGLOBAL HANDLE
type HGLRC HANDLE
type HGLYPH ULONG
type HHOOK HANDLE
type HICON HANDLE
type HIMAGELIST uintptr
type HIMC HANDLE
type HIMCC HANDLE
type HINSTANCE HANDLE
type HKEY HANDLE
type HKL HANDLE
type HLOCAL HANDLE
type HMENU HANDLE
type HMETAFILE HANDLE
type HMIDI HANDLE
type HMIDIIN HANDLE
type HMIDIOUT HANDLE
type HMIDISTRM HANDLE
type HMODULE uintptr
type HMONITOR HANDLE
type HOOKPROC func(code int32, wParam WPARAM, lParam LPARAM) LRESULT
type HPAINTBUFFER HANDLE
type HPALETTE HGDIOBJ
type HPEN HGDIOBJ
type HPOWERNOTIFY uintptr
type HPROPSHEETPAGE uintptr
type HPSXA HANDLE
type HRAWINPUT HANDLE
type HRESULT int32
type HRGN HANDLE
type HRSRC HANDLE
type HSEMAPHORE HANDLE
type HSURF HANDLE
type HSZ HANDLE
type HTASK HANDLE
type HTHEME HANDLE
type HTOUCHINPUT HANDLE
type HUSKEY HANDLE
type HWCT uintptr
type HWINEVENTHOOK HANDLE
type HWINSTA HANDLE
type HWND HANDLE
type Handle_t RPC_BINDING_HANDLE
type IBindCtx struct {
	lpVtbl uintptr
}
type ICMENUMPROC func(unnamed0 LPWSTR, unnamed1 LPARAM) int32
type ICONINFO struct {
	FIcon    BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}
type IConnectionPoint struct {
	lpVtbl uintptr
}
type ICreateErrorInfo struct {
	lpVtbl uintptr
}
type ICreateTypeLib struct {
	lpVtbl uintptr
}
type ICreateTypeLib2 struct {
	lpVtbl uintptr
}
type IDispatch struct {
	lpVtbl uintptr
}
type IErrorInfo struct {
	lpVtbl uintptr
}
type IFTYPE ULONG
type IF_INDEX NET_IFINDEX
type IF_LUID NET_LUID
type IID GUID
type IMAGEINFO struct {
	HbmImage HBITMAP
	HbmMask  HBITMAP
	Unused1  int32
	Unused2  int32
	RcImage  RECT
}
type IMAGELISTDRAWPARAMS struct {
	CbSize   DWORD
	Himl     HIMAGELIST
	I        int32
	HdcDst   HDC
	X        int32
	Y        int32
	Cx       int32
	Cy       int32
	XBitmap  int32
	YBitmap  int32
	RgbBk    COLORREF
	RgbFg    COLORREF
	FStyle   UINT
	DwRop    DWORD
	FState   DWORD
	Frame    DWORD
	CrEffect COLORREF
}
type IMCENUMPROC func(unnamed0 HIMC, unnamed1 LPARAM) BOOL
type IMEMENUITEMINFO struct {
	CbSize        UINT
	FType         UINT
	FState        UINT
	WID           UINT
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    DWORD
	SzString      [IMEMENUITEM_STRING_SIZE]WCHAR
	HbmpItem      HBITMAP
}
type IMEPRO struct {
	HWnd          HWND
	InstDate      DATETIME
	WVersion      uint32 // UINT
	SzDescription [50]uint16
	SzName        [80]uint16
	SzOptions     [30]uint16
}
type IMoniker struct {
	lpVtbl uintptr
}
type INITCOMMONCONTROLSEX struct {
	DwSize DWORD
	DwICC  DWORD
}
type INT int32
type INTERFACEDATA struct {
	Pmethdata *METHODDATA
	CMembers  UINT
}
type INTERFACE_HANDLE PVOID
type INTLIST struct {
	IValueCount int32
	IValues     [MAX_INTLIST_COUNT]int32
}
type INT_PTR *int32
type IPAddr ULONG
type IPMask ULONG

func (this *IP_ADAPTER_ADDRESSES_LH) Alignment() *ULONGLONG {
	return &this.union1
}
func (this *IP_ADAPTER_ADDRESSES_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_ADDRESSES_LH) IfIndex() *IF_INDEX {
	return (*IF_INDEX)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
func (this *IP_ADAPTER_ADDRESSES_LH) Flags() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union2))
}
func (this *IP_ADAPTER_ADDRESSES_LH) DdnsEnabled() bool {
	return this.union2 == 0x80000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) RegisterAdapterSuffix() bool {
	return this.union2 == 0x20000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Dhcpv4Enabled() bool {
	return this.union2 == 0x10000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) ReceiveOnly() bool {
	return this.union2 == 0x8000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) NoMulticast() bool {
	return this.union2 == 0x2000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv6OtherStatefulConfig() bool {
	return this.union2 == 0x1000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) NetbiosOverTcpipEnabled() bool {
	return this.union2 == 0x800000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv4Enabled() bool {
	return this.union2 == 0x200000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv6Enabled() bool {
	return this.union2 == 0x100000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv6ManagedAddressConfigurationSupported() bool {
	return this.union2 == 0x80000
}
func (this *IP_ADAPTER_ANYCAST_ADDRESS_XP) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_ANYCAST_ADDRESS_XP) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_ANYCAST_ADDRESS_XP) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
func (this *IP_ADAPTER_DNS_SERVER_ADDRESS_XP) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_DNS_SERVER_ADDRESS_XP) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_DNS_SERVER_ADDRESS_XP) Reserved() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

type IP_ADAPTER_DNS_SUFFIX struct {
	Next   *IP_ADAPTER_DNS_SUFFIX
	String [MAX_DNS_SUFFIX_STRING_LENGTH]WCHAR
}

func (this *IP_ADAPTER_GATEWAY_ADDRESS_LH) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_GATEWAY_ADDRESS_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_GATEWAY_ADDRESS_LH) Reserved() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

type IP_ADAPTER_INDEX_MAP struct {
	Index ULONG
	Name  [MAX_ADAPTER_NAME]WCHAR
}
type IP_ADAPTER_INFO struct {
	Next                *IP_ADAPTER_INFO
	ComboIndex          DWORD
	AdapterName         [MAX_ADAPTER_NAME_LENGTH + 4]CHAR
	Description         [MAX_ADAPTER_DESCRIPTION_LENGTH + 4]CHAR
	AddressLength       UINT
	Address             [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	Index               DWORD
	Type                UINT
	DhcpEnabled         UINT
	CurrentIpAddress    PIP_ADDR_STRING
	IpAddressList       IP_ADDR_STRING
	GatewayList         IP_ADDR_STRING
	DhcpServer          IP_ADDR_STRING
	HaveWins            BOOL
	PrimaryWinsServer   IP_ADDR_STRING
	SecondaryWinsServer IP_ADDR_STRING
	LeaseObtained       Time_t
	LeaseExpires        Time_t
}

func (this *IP_ADAPTER_MULTICAST_ADDRESS_XP) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_MULTICAST_ADDRESS_XP) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_MULTICAST_ADDRESS_XP) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

type IP_ADAPTER_ORDER_MAP struct {
	NumAdapters  ULONG
	AdapterOrder [1]ULONG
}
type IP_ADAPTER_PREFIX_XP struct {
	union1       ULONGLONG
	Next         *IP_ADAPTER_PREFIX_XP
	Address      SOCKET_ADDRESS
	PrefixLength ULONG
}

func (this *IP_ADAPTER_PREFIX_XP) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_PREFIX_XP) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_PREFIX_XP) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

type IP_ADAPTER_UNICAST_ADDRESS_LH struct {
	union1             ULONGLONG
	Next               *IP_ADAPTER_UNICAST_ADDRESS_LH
	Address            SOCKET_ADDRESS
	PrefixOrigin       IP_PREFIX_ORIGIN
	SuffixOrigin       IP_SUFFIX_ORIGIN
	DadState           IP_DAD_STATE
	ValidLifetime      ULONG
	PreferredLifetime  ULONG
	LeaseLifetime      ULONG
	OnLinkPrefixLength UINT8
}

func (this *IP_ADAPTER_UNICAST_ADDRESS_LH) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_UNICAST_ADDRESS_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_UNICAST_ADDRESS_LH) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
func (this *IP_ADAPTER_WINS_SERVER_ADDRESS_LH) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_WINS_SERVER_ADDRESS_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_WINS_SERVER_ADDRESS_LH) Reserved() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

type IP_ADDRESS_STRING struct {
	String [4 * 4]CHAR
}
type IP_ADDR_STRING struct {
	Next      *IP_ADDR_STRING
	IpAddress IP_ADDRESS_STRING
	IpMask    IP_MASK_STRING
	Context   DWORD
}
type IP_DAD_STATE NL_DAD_STATE
type IP_INTERFACE_INFO struct {
	NumAdapters LONG
	Adapter     [1]IP_ADAPTER_INDEX_MAP
}
type IP_MASK_STRING struct {
	String [4 * 4]CHAR
}
type IP_PER_ADAPTER_INFO_W2KSP1 struct {
	AutoconfigEnabled UINT
	AutoconfigActive  UINT
	CurrentDnsServer  PIP_ADDR_STRING
	DnsServerList     IP_ADDR_STRING
}
type IP_PREFIX_ORIGIN NL_PREFIX_ORIGIN
type IP_STATUS ULONG
type IP_SUFFIX_ORIGIN NL_SUFFIX_ORIGIN
type IRecordInfo struct {
	lpVtbl uintptr
}
type IRpcChannelBuffer struct {
	lpVtbl uintptr
}
type IRpcStubBuffer struct {
	lpVtbl uintptr
}
type IStream struct {
	lpVtbl uintptr
}
type ITEMIDLIST struct {
	Mkid SHITEMID
}
type ITypeInfo struct {
	lpVtbl uintptr
}
type ITypeLib struct {
	lpVtbl uintptr
}
type IUnknown struct {
	lpVtbl uintptr
}
type I_RPC_HANDLE uintptr
type KERNINGPAIR struct {
	WFirst      WORD
	WSecond     WORD
	IKernAmount int32
}
type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr // ULONG_PTR
}
type LANGID uint16
type LARGE_INTEGER struct {
	QuadPart int64
}

func (l *LARGE_INTEGER) LowPart() *uint32 {
	return (*uint32)(unsafe.Pointer(&l.QuadPart))
}
func (l *LARGE_INTEGER) HighPart() *int32 {
	return (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&l.QuadPart)) + uintptr(4)))
}

type LASTINPUTINFO struct {
	CbSize uint32 // UINT
	DwTime uint32
}
type LAYERPLANEDESCRIPTOR struct {
	NSize           uint16
	NVersion        uint16
	DwFlags         uint32
	IPixelType      byte
	CColorBits      byte
	CRedBits        byte
	CRedShift       byte
	CGreenBits      byte
	CGreenShift     byte
	CBlueBits       byte
	CBlueShift      byte
	CAlphaBits      byte
	CAlphaShift     byte
	CAccumBits      byte
	CAccumRedBits   byte
	CAccumGreenBits byte
	CAccumBlueBits  byte
	CAccumAlphaBits byte
	CDepthBits      byte
	CStencilBits    byte
	CAuxBuffers     byte
	ILayerPlane     byte
	BReserved       byte
	CrTransparent   COLORREF
}
type LCID uint32
type LCSCSTYPE LONG
type LCSGAMUTMATCH LONG
type LCTYPE uint32
type LINEATTRS struct {
	Fl           FLONG
	IJoin        ULONG
	IEndCap      ULONG
	ElWidth      FLOAT_LONG
	EMiterLimit  FLOATL
	Cstyle       ULONG
	Pstyle       PFLOAT_LONG
	ElStyleState FLOAT_LONG
}
type LINEDDAPROC func(unnamed0 int32, unnamed1 int32, unnamed LPARAM)
type LOGBRUSH struct {
	LbStyle UINT
	LbColor COLORREF
	LbHatch ULONG_PTR
}
type LOGCOLORSPACE struct {
	LcsSignature  DWORD
	LcsVersion    DWORD
	LcsSize       DWORD
	LcsCSType     LCSCSTYPE
	LcsIntent     LCSGAMUTMATCH
	LcsEndpoints  CIEXYZTRIPLE
	LcsGammaRed   DWORD
	LcsGammaGreen DWORD
	LcsGammaBlue  DWORD
	LcsFilename   [MAX_PATH]WCHAR
}
type LOGFONT struct {
	LfHeight         LONG
	LfWidth          LONG
	LfEscapement     LONG
	LfOrientation    LONG
	LfWeight         LONG
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       [LF_FACESIZE]WCHAR
}
type LOGPALETTE struct {
	PalVersion    WORD
	PalNumEntries WORD
	PalPalEntry   [1]PALETTEENTRY
}
type LOGPEN struct {
	LopnStyle UINT
	LopnWidth POINT
	LopnColor COLORREF
}
type LONG int32
type LONG64 int64
type LONGLONG int64
type LONG_PTR *int32
type LPAFPROTOCOLS *AFPROTOCOLS
type LPARAM uintptr
type tagBLOB struct {
	CbSize    ULONG
	PBlobData *BYTE
}
type LPBLOB *tagBLOB
type LPBYTE *byte
type LPCFHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPCHOOSEFONT *CHOOSEFONT
type LPCITEMIDLIST *ITEMIDLIST
type LPCOLESTR *OLECHAR
type LPCONDITIONPROC func(lpCallerId LPWSABUF, lpCallerData LPWSABUF, lpSQOS LPQOS, lpGQOS LPQOS, lpCalleeId LPWSABUF, lpCalleeData LPWSABUF, g *GROUP, dwCallbackData DWORD_PTR) int32
type LPCSADDR_INFO *CSADDR_INFO
type LPCSTR *byte
type LPCWSTR *uint16
type LPDEVMODE *DEVMODE
type LPDISPATCH *IDispatch
type LPFINDREPLACE *FINDREPLACE
type LPFRHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPGCP_RESULTS *GCP_RESULTS
type LPGUID *GUID
type LPHELPINFO *HELPINFO
type LPIMEMENUITEMINFO *IMEMENUITEMINFO
type LPLOGCOLORSPACE *LOGCOLORSPACE
type LPLOGFONT *LOGFONT
type LPLOOKUPSERVICE_COMPLETION_ROUTINE func(dwError DWORD, dwBytes DWORD, lpOverlapped LPWSAOVERLAPPED)
type LPMONIKER *IMoniker
type LPOFNHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPOLESTR *OLECHAR
type LPOPENFILENAME *OPENFILENAME
type LPOUTLINETEXTMETRIC *OUTLINETEXTMETRIC
type LPPAGEPAINTHOOK func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPPAGESETUPDLG *PAGESETUPDLG
type LPPAGESETUPHOOK func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPPRINTDLG *PRINTDLG
type LPPRINTDLGEX *PRINTDLGEX
type LPPRINTHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPPRINTPAGERANGE *PRINTPAGERANGE
type LPQOS *QOS
type LPRASTERIZER_STATUS *RASTERIZER_STATUS
type LPRECT *RECT
type LPSAFEARRAY *SAFEARRAY
type LPSETUPHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
type LPSOCKADDR *SOCKADDR
type LPSTR *CHAR
type LPSTREAM *IStream
type LPSTYLEBUF *STYLEBUF
type LPTEXTMETRIC *TEXTMETRIC
type LPUNKNOWN *IUnknown
type LPVOID uintptr
type LPWPUPOSTMESSAGE func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) BOOL
type LPWSABUF *WSABUF
type LPWSANAMESPACE_INFO *WSANAMESPACE_INFO
type LPWSANSCLASSINFO *WSANSCLASSINFO
type LPWSAOVERLAPPED *OVERLAPPED
type LPWSAOVERLAPPED_COMPLETION_ROUTINE func(dwError DWORD, cbTransferred DWORD, lpOverlapped LPWSAOVERLAPPED, dwFlags DWORD)
type LPWSAPROTOCOL_INFO *WSAPROTOCOL_INFO
type LPWSAQUERYSET *WSAQUERYSET
type LPWSASERVICECLASSINFO *WSASERVICECLASSINFO
type LPWSAVERSION *WSAVERSION
type LPWSTR *uint16
type LRESULT uintptr
type LUID struct {
	LowPart  uint32
	HighPart int32
}
type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes ULONG
}
type MALLOC_FREE_STRUCT struct {
	PfnAllocate uintptr // void* (__RPC_USER *pfnAllocate)(size_t)
	PfnFree     uintptr // void (__RPC_USER *pfnFree)(void *)
}
type MARGINS struct {
	CxLeftWidth    int32
	CxRightWidth   int32
	CyTopHeight    int32
	CyBottomHeight int32
}
type MAT2 struct {
	EM11 FIXED
	EM12 FIXED
	EM21 FIXED
	EM22 FIXED
}
type MCIDEVICEID UINT
type MCIERROR DWORD
type MENUBARINFO struct {
	CbSize          uint32
	RcBar           RECT
	HMenu           HMENU
	HwndMenu        HWND
	bitfieldedFlags uint32
	// BOOL fBarFocused:1;
	// BOOL fFocused:1;
}

func (i *MENUBARINFO) FBarFocused() bool {
	return (i.bitfieldedFlags & 1) == 1
}
func (i *MENUBARINFO) FFocused() bool {
	return (i.bitfieldedFlags & 2) == 2
}

type MENUINFO struct {
	CbSize          uint32
	FMask           uint32
	DwStyle         uint32
	CyMax           uint32
	HbrBack         HBRUSH
	DwContextHelpID uint32
	DwMenuData      uintptr
}
type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32
	FType         uint32
	FState        uint32
	WID           uint32
	HSubMenu      HMENU
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    uintptr
	DwTypeData    *uint16
	Cch           uint32
	HbmpItem      HBITMAP
}
type METAFILEPICT struct {
	Mm   LONG
	XExt LONG
	YExt LONG
	HMF  HMETAFILE
}
type METARECORD struct {
	RdSize     DWORD
	RdFunction WORD
	RdParm     [1]WORD
}
type METHODDATA struct {
	SzName   *OLECHAR
	Ppdata   *PARAMDATA
	Dispid   DISPID
	IMeth    UINT
	Cc       CALLCONV
	CArgs    UINT
	WFlags   WORD
	VtReturn VARTYPE
}
type MFENUMPROC func(hdc HDC, lpht *HANDLETABLE, lpMR *METARECORD, nObj int32, param LPARAM) int32
type MIBICMPINFO struct {
	IcmpInStats  MIBICMPSTATS
	IcmpOutStats MIBICMPSTATS
}
type MIBICMPSTATS struct {
	DwMsgs          DWORD
	DwErrors        DWORD
	DwDestUnreachs  DWORD
	DwTimeExcds     DWORD
	DwParmProbs     DWORD
	DwSrcQuenchs    DWORD
	DwRedirects     DWORD
	DwEchos         DWORD
	DwEchoReps      DWORD
	DwTimestamps    DWORD
	DwTimestampReps DWORD
	DwAddrMasks     DWORD
	DwAddrMaskReps  DWORD
}
type MIBICMPSTATS_EX MIBICMPSTATS_EX_XPSP1
type MIBICMPSTATS_EX_XPSP1 struct {
	DwMsgs        DWORD
	DwErrors      DWORD
	RgdwTypeCount [256]DWORD
}
type MIB_ICMP struct {
	Stats MIBICMPINFO
}
type MIB_ICMP_EX_XPSP1 struct {
	IcmpInStats  MIBICMPSTATS_EX
	IcmpOutStats MIBICMPSTATS_EX
}
type MIB_IFROW struct {
	WszName           [MAX_INTERFACE_NAME_LEN]WCHAR
	DwIndex           IF_INDEX
	DwType            IFTYPE
	DwMtu             DWORD
	DwSpeed           DWORD
	DwPhysAddrLen     DWORD
	BPhysAddr         [MAXLEN_PHYSADDR]UCHAR
	DwAdminStatus     DWORD
	DwOperStatus      INTERNAL_IF_OPER_STATUS
	DwLastChange      DWORD
	DwInOctets        DWORD
	DwInUcastPkts     DWORD
	DwInNUcastPkts    DWORD
	DwInDiscards      DWORD
	DwInErrors        DWORD
	DwInUnknownProtos DWORD
	DwOutOctets       DWORD
	DwOutUcastPkts    DWORD
	DwOutNUcastPkts   DWORD
	DwOutDiscards     DWORD
	DwOutErrors       DWORD
	DwOutQLen         DWORD
	DwDescrLen        DWORD
	BDescr            [MAXLEN_IFDESCR]UCHAR
}
type MIB_IFTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IFROW
}
type MIB_IPADDRROW MIB_IPADDRROW_XP
type MIB_IPADDRROW_XP struct {
	DwAddr      DWORD
	DwIndex     IF_INDEX
	DwMask      DWORD
	DwBCastAddr DWORD
	DwReasmSize DWORD
	Unused1     uint16
	WType       uint16
}
type MIB_IPADDRTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IPADDRROW
}
type MIB_IPFORWARDROW struct {
	DwForwardDest      DWORD
	DwForwardMask      DWORD
	DwForwardPolicy    DWORD
	DwForwardNextHop   DWORD
	DwForwardIfIndex   IF_INDEX
	ForwardType        MIB_IPFORWARD_TYPE
	ForwardProto       MIB_IPFORWARD_PROTO
	DwForwardAge       DWORD
	DwForwardNextHopAS DWORD
	DwForwardMetric1   DWORD
	DwForwardMetric2   DWORD
	DwForwardMetric3   DWORD
	DwForwardMetric4   DWORD
	DwForwardMetric5   DWORD
}
type MIB_IPFORWARDTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IPFORWARDROW
}
type MIB_IPFORWARD_PROTO NL_ROUTE_PROTOCOL
type MIB_IPNETROW MIB_IPNETROW_LH
type MIB_IPNETROW_LH struct {
	DwIndex       IF_INDEX
	DwPhysAddrLen DWORD
	BPhysAddr     [MAXLEN_PHYSADDR]UCHAR
	DwAddr        DWORD
	Type          MIB_IPNET_TYPE
}
type MIB_IPNETTABLE struct {
	DwNumEntries DWORD
	Table        [ANY_SIZE]MIB_IPNETROW
}
type MIB_IPSTATS_LH struct {
	Forwarding        MIB_IPSTATS_FORWARDING
	DwDefaultTTL      DWORD
	DwInReceives      DWORD
	DwInHdrErrors     DWORD
	DwInAddrErrors    DWORD
	DwForwDatagrams   DWORD
	DwInUnknownProtos DWORD
	DwInDiscards      DWORD
	DwInDelivers      DWORD
	DwOutRequests     DWORD
	DwRoutingDiscards DWORD
	DwOutDiscards     DWORD
	DwOutNoRoutes     DWORD
	DwReasmTimeout    DWORD
	DwReasmReqds      DWORD
	DwReasmOks        DWORD
	DwReasmFails      DWORD
	DwFragOks         DWORD
	DwFragFails       DWORD
	DwFragCreates     DWORD
	DwNumIf           DWORD
	DwNumAddr         DWORD
	DwNumRoutes       DWORD
}
type MIB_TCP6ROW_OWNER_MODULE struct {
	UcLocalAddr       [16]UCHAR
	DwLocalScopeId    DWORD
	DwLocalPort       DWORD
	UcRemoteAddr      [16]UCHAR
	DwRemoteScopeId   DWORD
	DwRemotePort      DWORD
	DwState           DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
type MIB_TCPROW_LH struct {
	State        MIB_TCP_STATE
	DwLocalAddr  DWORD
	DwLocalPort  DWORD
	DwRemoteAddr DWORD
	DwRemotePort DWORD
}
type MIB_TCPROW_OWNER_MODULE struct {
	DwState           DWORD
	DwLocalAddr       DWORD
	DwLocalPort       DWORD
	DwRemoteAddr      DWORD
	DwRemotePort      DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}

func (this *MIB_UDP6ROW_OWNER_MODULE) DwFlags() *int32 {
	return &this.dwFlags
}
func (this *MIB_UDP6ROW_OWNER_MODULE) SpecificPortBind() int32 {
	return this.dwFlags & 0x1
}
func (this *MIB_UDPROW_OWNER_MODULE) DwFlags() *int32 {
	return &this.dwFlags
}
func (this *MIB_UDPROW_OWNER_MODULE) SpecificPortBind() int32 {
	return this.dwFlags & 0x1
}

type MIDIINCAPS struct {
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
	DwSupport      DWORD
}
type MIDL_STUB_DESC struct {
	RpcInterfaceInformation     uintptr
	PfnAllocate                 uintptr // void* (__RPC_API *pfnAllocate)(size_t)
	PfnFree                     uintptr // void (__RPC_API *pfnFree)(void *)
	IMPLICIT_HANDLE_INFO        uintptr
	ApfnNdrRundownRoutines      uintptr // const NDR_RUNDOWN*
	AGenericBindingRoutinePairs uintptr // const GENERIC_BINDING_ROUTINE_PAIR*
	ApfnExprEval                uintptr // const EXPR_EVAL*
	AXmitQuintuple              uintptr // const XMIT_ROUTINE_QUINTUPLE*
	PFormatTypes/*const*/ *byte
	FCheckBounds      int32
	Version           uint32
	PMallocFreeStruct uintptr // MALLOC_FREE_STRUCT*
	MIDLVersion       int32
	CommFaultOffsets/*const*/ *COMM_FAULT_OFFSETS
	AUserMarshalQuadruple uintptr // const USER_MARSHAL_ROUTINE_QUADRUPLE*
	NotifyRoutineTable    uintptr // const NDR_NOTIFY_ROUTINE*
	MFlags                ULONG_PTR
	CsRoutineTables/*const*/ *NDR_CS_ROUTINES
	ProxyServerInfo uintptr
	PExprInfo/*const*/ *NDR_EXPR_DESC
}

func (this MIDL_STUB_DESC) PAutoHandle() *Handle_t {
	return (*Handle_t)(unsafe.Pointer(&this.IMPLICIT_HANDLE_INFO))
}
func (this MIDL_STUB_DESC) PPrimitiveHandle() *Handle_t {
	return (*Handle_t)(unsafe.Pointer(&this.IMPLICIT_HANDLE_INFO))
}
func (this MIDL_STUB_DESC) PGenericBindingInfo() PGENERIC_BINDING_INFO {
	return (PGENERIC_BINDING_INFO)(unsafe.Pointer(&this.IMPLICIT_HANDLE_INFO))
}

type MIDL_STUB_MESSAGE struct {
	RpcMsg                 PRPC_MESSAGE
	Buffer                 *byte
	BufferStart            *byte
	BufferEnd              *byte
	BufferMark             *byte
	BufferLength           uint32
	MemorySize             uint32
	Memory                 *byte
	IsClient               byte
	Pad                    byte
	UFlags2                uint16
	ReuseBuffer            int32
	PAllocAllNodesContext  uintptr // struct NDR_ALLOC_ALL_NODES_CONTEXT*
	PPointerQueueState     uintptr // struct NDR_POINTER_QUEUE_STATE*
	IgnoreEmbeddedPointers int32
	PointerBufferMark      *byte
	CorrDespIncrement      byte
	uFlags                 byte
	UniquePtrCount         uint16
	MaxCount               ULONG_PTR
	Offset                 uint32
	ActualCount            uint32
	PfnAllocate            uintptr // void*(__RPC_API *pfnAllocate)(size_t)
	PfnFree                uintptr // void(__RPC_API *pfnFree)(void*)
	StackTop               *byte
	PPresentedType         *byte
	PTransmitType          *byte
	SavedHandle            Handle_t
	StubDesc/*const*/ *MIDL_STUB_DESC
	FullPtrXlatTables *FULL_PTR_XLAT_TABLES
	FullPtrRefId      uint32
	PointerLength     uint32
	fBitField32       uint32
	/*
	   int                             fInDontFree       :1;
	   int                             fDontCallFreeInst :1;
	   int                             fInOnlyParam      :1;
	   int                             fHasReturn        :1;
	   int                             fHasExtensions    :1;
	   int                             fHasNewCorrDesc   :1;
	   int                             fIsIn             :1;
	   int                             fIsOut            :1;
	   int                             fIsOicf           :1;
	   int                             fBufferValid      :1;
	   int                             fHasMemoryValidateCallback: 1;
	   int                             fInFree             :1;
	   int                             fNeedMCCP         :1;
	   int                             fUnused           :3;
	   int                             fUnused2          :16;
	*/
	DwDestContext       uint32
	PvDestContext       uintptr
	SavedContextHandles *NDR_SCONTEXT
	ParamNumber         int32
	PRpcChannelBuffer   *IRpcChannelBuffer
	PArrayInfo          PARRAY_INFO
	SizePtrCountArray   *uint32
	SizePtrOffsetArray  *uint32
	SizePtrLengthArray  *uint32
	PArgQueue           uintptr
	DwStubPhase         uint32
	LowStackMark        uintptr
	PAsyncMsg           uintptr // PNDR_ASYNC_MESSAGE
	PCorrInfo           uintptr // PNDR_CORRELATION_INFO
	PCorrMemory         *byte
	PMemoryList         uintptr
	PCSInfo             INT_PTR
	ConformanceMark     *byte
	VarianceMark        *byte
	Unused              INT_PTR
	PContext            uintptr // struct _NDR_PROC_CONTEXT*
	ContextHandleHash   uintptr
	PUserMarshalList    uintptr
	Reserved51_3        INT_PTR
	Reserved51_4        INT_PTR
	Reserved51_5        INT_PTR
}
type MIX ULONG
type MMRESULT uint32
type MMVERSION UINT
type MODULEINFO struct {
	LpBaseOfDll LPVOID
	SizeOfImage DWORD
	EntryPoint  LPVOID
}
type MONITORENUMPROC func(hMonitor HMONITOR, hdcMonitor HDC, lprcMonitor *RECT, dwData uintptr) BOOL
type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}
type MOUSEINPUT struct {
	Dx          int32 // LONG
	Dy          int32 // LONG
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr // ULONG_PTR
}
type MOUSEMOVEPOINT struct {
	X           int32
	Y           int32
	Time        uint32
	DwExtraInfo uintptr // ULONG_PTR
}
type MRUCMPPROC func(pString1 string, pString2 string) int32
type MRUINFO struct {
	CbSize      DWORD
	UMax        UINT
	FFlags      UINT
	HKey        HKEY
	LpszSubKey  LPCWSTR
	LpfnCompare uintptr // MRUCMPPROC
}
type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}
type MSGBOXCALLBACK func(lpHelpInfo LPHELPINFO)
type MSGBOXPARAMS struct {
	CbSize             uint32
	HwndOwner          HWND
	HInstance          HINSTANCE
	LpszText           *uint16 // LPCWSTR
	LpszCaption        *uint16 // LPCWSTR
	DwStyle            uint32
	LpszIcon           *uint16 // LPCWSTR
	DwContextHelpId    *uint32 // DWORD_PTR
	LpfnMsgBoxCallback uintptr // MSGBOXCALLBACK
	DwLanguageId       uint32
}
type NDR_CS_ROUTINES struct {
	PSizeConvertRoutines *NDR_CS_SIZE_CONVERT_ROUTINES // NDR_CS_SIZE_CONVERT_ROUTINES
	PTagGettingRoutines  uintptr                       // CS_TAG_GETTING_ROUTINE
}
type NDR_CS_SIZE_CONVERT_ROUTINES struct {
	PfnNetSize   uintptr // CS_TYPE_NET_SIZE_ROUTINE
	PfnToNetCs   uintptr // CS_TYPE_TO_NETCS_ROUTINE
	PfnLocalSize uintptr // CS_TYPE_LOCAL_SIZE_ROUTINE
	PfnFromNetCs uintptr // CS_TYPE_FROM_NETCS_ROUTINE
}
type NDR_EXPR_DESC struct {
	POffset/*const*/ *uint16
	PFormatExpr PFORMAT_STRING
}
type NDR_RUNDOWN func(context uintptr)
type NDR_SCONTEXT *NDR_SCONTEXT_
type NDR_SCONTEXT_ struct {
	Pad         [2]uintptr
	UserContext uintptr
}
type NET_IFINDEX ULONG
type NET_IF_NETWORK_GUID GUID
type NET_LUID NET_LUID_LH
type NET_LUID_LH struct {
	Value ULONG64
}

func (this *NET_LUID_LH) Reserved() ULONG64 {
	v := this.Value
	return (v & 0xFFFFFF0000000000) >> (64 - 24)
}
func (this *NET_LUID_LH) NetLuidIndex() ULONG64 {
	v := this.Value
	return (v & 0xFFFFFF0000) >> 16
}
func (this *NET_LUID_LH) IfType() ULONG64 {
	v := this.Value
	return v & 0xFFFF
}

type NUMPARSE struct {
	CDig       INT
	DwInFlags  ULONG
	DwOutFlags ULONG
	CchUsed    INT
	NBaseShift INT
	NPwr10     INT
}
type OBJECTS_AND_NAME struct {
	ObjectsPresent          DWORD
	ObjectType              SE_OBJECT_TYPE
	ObjectTypeName          LPWSTR
	InheritedObjectTypeName LPWSTR
	PtstrName               LPWSTR
}
type OBJECTS_AND_SID struct {
	ObjectsPresent          DWORD
	ObjectTypeGuid          GUID
	InheritedObjectTypeGuid GUID
	PSid                    *SID
}
type OBJECT_TYPE_LIST struct {
	Level      USHORT
	Sbz        USHORT
	ObjectType *GUID
}
type OCPFIPARAMS struct {
	CbStructSize          ULONG
	HWndOwner             HWND
	X                     int32
	Y                     int32
	LpszCaption           LPCOLESTR
	CObjects              ULONG
	LplpUnk               *LPUNKNOWN
	CPages                ULONG
	LpPages               *CLSID
	Lcid                  LCID
	DispidInitialProperty DISPID
}
type OLECHAR WCHAR
type OLE_COLOR DWORD
type OPENFILENAME struct {
	LStructSize       DWORD
	HwndOwner         HWND
	HInstance         HINSTANCE
	LpstrFilter       LPCWSTR
	LpstrCustomFilter LPWSTR
	NMaxCustFilter    DWORD
	NFilterIndex      DWORD
	LpstrFile         LPWSTR
	NMaxFile          DWORD
	LpstrFileTitle    LPWSTR
	NMaxFileTitle     DWORD
	LpstrInitialDir   LPCWSTR
	LpstrTitle        LPCWSTR
	Flags             DWORD
	NFileOffset       WORD
	NFileExtension    WORD
	LpstrDefExt       LPCWSTR
	LCustData         LPARAM
	LpfnHook          LPOFNHOOKPROC
	LpTemplateName    LPCWSTR
	PvReserved        uintptr
	DwReserved        DWORD
	FlagsEx           DWORD
}
type OUTLINETEXTMETRIC struct {
	OtmSize                UINT
	OtmTextMetrics         TEXTMETRIC
	OtmFiller              BYTE
	OtmPanoseNumber        PANOSE
	OtmfsSelection         UINT
	OtmfsType              UINT
	OtmsCharSlopeRise      int32
	OtmsCharSlopeRun       int32
	OtmItalicAngle         int32
	OtmEMSquare            UINT
	OtmAscent              int32
	OtmDescent             int32
	OtmLineGap             UINT
	OtmsCapEmHeight        UINT
	OtmsXHeight            UINT
	OtmrcFontBox           RECT
	OtmMacAscent           int32
	OtmMacDescent          int32
	OtmMacLineGap          UINT
	OtmusMinimumPPEM       UINT
	OtmptSubscriptSize     POINT
	OtmptSubscriptOffset   POINT
	OtmptSuperscriptSize   POINT
	OtmptSuperscriptOffset POINT
	OtmsStrikeoutSize      UINT
	OtmsStrikeoutPosition  int32
	OtmsUnderscoreSize     int32
	OtmsUnderscorePosition int32
	OtmpFamilyName         PSTR
	OtmpFaceName           PSTR
	OtmpStyleName          PSTR
	OtmpFullName           PSTR
}
type OVERLAPPED struct {
	Internal     ULONG_PTR
	InternalHigh ULONG_PTR
	union1       [8]byte
	HEvent       HANDLE
}

func (this *OVERLAPPED) Offset() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.union1[0]))
}
func (this *OVERLAPPED) OffsetHigh() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.union1[4]))
}
func (this *OVERLAPPED) Pointer() *PVOID {
	return (*PVOID)(unsafe.Pointer(&this.union1[0]))
}

type PADDRINFO *ADDRINFO
type PAGESETUPDLG struct {
	LStructSize             DWORD
	HwndOwner               HWND
	HDevMode                HGLOBAL
	HDevNames               HGLOBAL
	Flags                   DWORD
	PtPaperSize             POINT
	RtMinMargin             RECT
	RtMargin                RECT
	HInstance               HINSTANCE
	LCustData               LPARAM
	LpfnPageSetupHook       LPPAGESETUPHOOK
	LpfnPagePaintHook       LPPAGEPAINTHOOK
	LpPageSetupTemplateName LPCWSTR
	HPageSetupTemplate      HGLOBAL
}
type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}
type PALETTEENTRY struct {
	PeRed   byte
	PeGreen byte
	PeBlue  byte
	PeFlags byte
}
type PANOSE struct {
	BFamilyType      BYTE
	BSerifStyle      BYTE
	BWeight          BYTE
	BProportion      BYTE
	BContrast        BYTE
	BStrokeVariation BYTE
	BArmStyle        BYTE
	BLetterform      BYTE
	BMidline         BYTE
	BXHeight         BYTE
}
type PARAMDATA struct {
	SzName *OLECHAR
	Vt     VARTYPE
}
type PARRAY_INFO *ARRAY_INFO
type PARSEDURL struct {
	CbSize      DWORD
	PszProtocol LPCWSTR
	CchProtocol UINT
	PszSuffix   LPCWSTR
	CchSuffix   UINT
	NScheme     UINT
}
type PATHDATA struct {
	Flags FLONG
	Count ULONG
	Pptfx *POINTFIX
}
type PATHOBJ struct {
	Fl      FLONG
	CCurves ULONG
}
type PAUDIT_POLICY_INFORMATION *AUDIT_POLICY_INFORMATION
type PBYTE *byte
type PCCRL_CONTEXT /*const*/ *CRL_CONTEXT
type PCERT_EXTENSION *CERT_EXTENSION
type PCERT_NAME_BLOB *CERT_NAME_BLOB
type PCERT_RDN_VALUE_BLOB *CERT_RDN_VALUE_BLOB
type PCHAR *CHAR
type PCRL_ENTRY *CRL_ENTRY
type PCRL_INFO *CRL_INFO
type PCRYPT_INTEGER_BLOB *CRYPT_INTEGER_BLOB
type PCWSTR *uint16
type PDH_COUNTER_INFO struct {
	DwLength        DWORD
	DwType          DWORD
	CVersion        DWORD
	CStatus         DWORD
	LScale          LONG
	LDefaultScale   LONG
	DwUserData      DWORD_PTR
	DwQueryUserData DWORD_PTR
	SzFullPath      LPWSTR
	union1          [44]byte
	SzExplainText   LPWSTR
	DataBuffer      [1]DWORD
}

func (this *PDH_COUNTER_INFO) DataItemPath() *PDH_DATA_ITEM_PATH_ELEMENTS {
	return (*PDH_DATA_ITEM_PATH_ELEMENTS)(unsafe.Pointer(&this.union1[0]))
}
func (this *PDH_COUNTER_INFO) CounterPath() *PDH_COUNTER_PATH_ELEMENTS {
	return (*PDH_COUNTER_PATH_ELEMENTS)(unsafe.Pointer(&this.union1[0]))
}
func (this *PDH_COUNTER_INFO) SzMachineName() *LPWSTR {
	return (*LPWSTR)(unsafe.Pointer(&this.union1[0]))
}
func (this *PDH_COUNTER_INFO) SzObjectName() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)))
}
func (this *PDH_COUNTER_INFO) SzInstanceName() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*2))
}
func (this *PDH_COUNTER_INFO) SzParentInstance() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*3))
}
func (this *PDH_COUNTER_INFO) DwInstanceIndex() *DWORD {
	var ptr LPWSTR
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*4))
}
func (this *PDH_COUNTER_INFO) SzCounterName() *LPWSTR {
	var ptr LPWSTR
	return (*LPWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + unsafe.Sizeof(ptr)*4 + uintptr(4)))
}

type PDH_COUNTER_PATH_ELEMENTS struct {
	SzMachineName    LPWSTR
	SzObjectName     LPWSTR
	SzInstanceName   LPWSTR
	SzParentInstance LPWSTR
	DwInstanceIndex  DWORD
	SzCounterName    LPWSTR
}
type PDH_DATA_ITEM_PATH_ELEMENTS struct {
	SzMachineName  LPWSTR
	ObjectGUID     GUID
	wItemId        DWORD
	SzInstanceName LPWSTR
}
type PDH_FMT_COUNTERVALUE struct {
	storage [16]byte
}

func (this *PDH_FMT_COUNTERVALUE) CStatus() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *PDH_FMT_COUNTERVALUE) LongValue() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) DoubleValue() *float64 {
	return (*float64)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) LargeValue() *LONGLONG {
	return (*LONGLONG)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) AnsiStringValue() *LPCSTR {
	return (*LPCSTR)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) WideStringValue() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[4]))
}

type PDH_HCOUNTER HANDLE
type PDH_HLOG HANDLE
type PDH_HQUERY HANDLE
type PDH_STATUS LONG
type PENCRYPTION_CERTIFICATE *ENCRYPTION_CERTIFICATE
type PENG_TIME_FIELDS *ENG_TIME_FIELDS
type PENUM_PAGE_FILE_CALLBACK func(pContext LPVOID, pPageFileInfo PENUM_PAGE_FILE_INFORMATION, lpFilename string) BOOL
type PENUM_PAGE_FILE_INFORMATION *ENUM_PAGE_FILE_INFORMATION
type PERFORMANCE_INFORMATION struct {
	Cb                DWORD
	CommitTotal       SIZE_T
	CommitLimit       SIZE_T
	CommitPeak        SIZE_T
	PhysicalTotal     SIZE_T
	PhysicalAvailable SIZE_T
	SystemCache       SIZE_T
	KernelTotal       SIZE_T
	KernelPaged       SIZE_T
	KernelNonpaged    SIZE_T
	PageSize          SIZE_T
	HandleCount       DWORD
	ProcessCount      DWORD
	ThreadCount       DWORD
}
type PERF_COUNTERSET_INSTANCE struct {
	CounterSetGuid     GUID
	DwSize             ULONG
	InstanceId         ULONG
	InstanceNameOffset ULONG
	InstanceNameSize   ULONG
}
type PFD_GLYPHATTR *FD_GLYPHATTR
type PFIXED_INFO *FIXED_INFO_W2KSP1
type PFLOAT_LONG *FLOAT_LONG
type PFNCALLBACK func(wType uint32, wFmt uint32, hConv HCONV, hsz1 HSZ, hsz2 HSZ, hData HDDEDATA, dwData1 uintptr, dwData2 uintptr) HDDEDATA
type PFORMAT_STRING *byte
type PGENERIC_BINDING_INFO *GENERIC_BINDING_INFO
type PGLYPHPOS *GLYPHPOS
type PHUSKEY *HUSKEY
type PICTDESC struct {
	CbSizeofstruct UINT
	PicType        UINT
	union1         uintptr
	union2         int32
	union3         int32
}

func (this *PICTDESC) Hbitmap() HBITMAP {
	return HBITMAP(this.union1)
}
func (this *PICTDESC) Hpal() HPALETTE {
	var ptr uintptr
	if is64 {
		*(*int32)(unsafe.Pointer(&ptr)) = this.union2
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + 4)) = this.union3
	} else {
		*(*int32)(unsafe.Pointer(&ptr)) = this.union2
	}
	return HPALETTE(ptr)
}
func (this *PICTDESC) Hmeta() HMETAFILE {
	return HMETAFILE(this.union1)
}
func (this *PICTDESC) XExt() int32 {
	return this.union2
}
func (this *PICTDESC) YExt() int32 {
	return this.union3
}
func (this *PICTDESC) Hicon() HICON {
	return HICON(this.union1)
}
func (this *PICTDESC) Hemf() HENHMETAFILE {
	return HENHMETAFILE(this.union1)
}

type PINT_PTR *INT_PTR
type PIP_ADAPTER_ADDRESSES *IP_ADAPTER_ADDRESSES_LH
type PIP_ADAPTER_ANYCAST_ADDRESS_XP *IP_ADAPTER_ANYCAST_ADDRESS_XP
type PIP_ADAPTER_DNS_SERVER_ADDRESS_XP *IP_ADAPTER_DNS_SERVER_ADDRESS_XP
type PIP_ADAPTER_DNS_SUFFIX *IP_ADAPTER_DNS_SUFFIX
type PIP_ADAPTER_GATEWAY_ADDRESS_LH *IP_ADAPTER_GATEWAY_ADDRESS_LH
type PIP_ADAPTER_INFO *IP_ADAPTER_INFO
type PIP_ADAPTER_MULTICAST_ADDRESS_XP *IP_ADAPTER_MULTICAST_ADDRESS_XP
type PIP_ADAPTER_ORDER_MAP *IP_ADAPTER_ORDER_MAP
type PIP_ADAPTER_PREFIX_XP *IP_ADAPTER_PREFIX_XP
type PIP_ADAPTER_UNICAST_ADDRESS_LH *IP_ADAPTER_UNICAST_ADDRESS_LH
type PIP_ADAPTER_WINS_SERVER_ADDRESS_LH *IP_ADAPTER_WINS_SERVER_ADDRESS_LH
type PIP_ADDR_STRING *IP_ADDR_STRING
type PIP_INTERFACE_INFO *IP_INTERFACE_INFO
type PIP_PER_ADAPTER_INFO *IP_PER_ADAPTER_INFO_W2KSP1
type PIXELFORMATDESCRIPTOR struct {
	NSize           uint16
	NVersion        uint16
	DwFlags         uint32
	IPixelType      byte
	CColorBits      byte
	CRedBits        byte
	CRedShift       byte
	CGreenBits      byte
	CGreenShift     byte
	CBlueBits       byte
	CBlueShift      byte
	CAlphaBits      byte
	CAlphaShift     byte
	CAccumBits      byte
	CAccumRedBits   byte
	CAccumGreenBits byte
	CAccumBlueBits  byte
	CAccumAlphaBits byte
	CDepthBits      byte
	CStencilBits    byte
	CAuxBuffers     byte
	ILayerType      byte
	BReserved       byte
	DwLayerMask     uint32
	DwVisibleMask   uint32
	DwDamageMask    uint32
}
type PMIB_ICMP *MIB_ICMP
type PMIB_ICMP_EX *MIB_ICMP_EX_XPSP1
type PMIB_IFROW *MIB_IFROW
type PMIB_IFTABLE *MIB_IFTABLE
type PMIB_IPADDRTABLE *MIB_IPADDRTABLE
type PMIB_IPFORWARDROW *MIB_IPFORWARDROW
type PMIB_IPFORWARDTABLE *MIB_IPFORWARDTABLE
type PMIB_IPNETROW *MIB_IPNETROW_LH
type PMIB_IPNETTABLE *MIB_IPNETTABLE
type PMIB_IPSTATS *MIB_IPSTATS_LH
type PMIB_TCP6ROW_OWNER_MODULE *MIB_TCP6ROW_OWNER_MODULE
type PMIB_TCPROW *MIB_TCPROW_LH
type PMIB_TCPROW_OWNER_MODULE *MIB_TCPROW_OWNER_MODULE
type PMIB_UDP6ROW_OWNER_MODULE *MIB_UDP6ROW_OWNER_MODULE
type PMIB_UDPROW_OWNER_MODULE *MIB_UDPROW_OWNER_MODULE
type PMIDL_STUB_MESSAGE *MIDL_STUB_MESSAGE
type POINT struct {
	X, Y int32
}
type POINTFIX struct {
	X FIX
	Y FIX
}
type POINTFLOAT struct {
	X float32
	Y float32
}
type POINTL struct {
	X LONG
	Y LONG
}
type POINTQF struct {
	X LARGE_INTEGER
	Y LARGE_INTEGER
}
type POINTS struct {
	X int16 // SHORT
	Y int16 // SHORT
}
type POLICY_AUDIT_SID_ARRAY struct {
	UsersCount   ULONG
	UserSidArray *PSID
}
type POLYTEXT struct {
	X       int32
	Y       int32
	N       UINT
	Lpstr   LPCWSTR
	UiFlags UINT
	Rcl     RECT
	Pdx     *int32
}
type PPERFORMACE_INFORMATION *PERFORMANCE_INFORMATION
type PPOLICY_AUDIT_EVENT_TYPE *POLICY_AUDIT_EVENT_TYPE
type PPOLICY_AUDIT_SID_ARRAY *POLICY_AUDIT_SID_ARRAY
type PPROCESS_MEMORY_COUNTERS *PROCESS_MEMORY_COUNTERS
type PPSAPI_WS_WATCH_INFORMATION *PSAPI_WS_WATCH_INFORMATION
type PPSAPI_WS_WATCH_INFORMATION_EX *PSAPI_WS_WATCH_INFORMATION_EX
type PRECTFX *RECTFX
type PRINTDLGEX struct {
	LStructSize         DWORD
	HwndOwner           HWND
	HDevMode            HGLOBAL
	HDevNames           HGLOBAL
	HDC                 HDC
	Flags               DWORD
	Flags2              DWORD
	ExclusionFlags      DWORD
	NPageRanges         DWORD
	NMaxPageRanges      DWORD
	LpPageRanges        LPPRINTPAGERANGE
	NMinPage            DWORD
	NMaxPage            DWORD
	NCopies             DWORD
	HInstance           HINSTANCE
	LpPrintTemplateName LPCWSTR
	LpCallback          LPUNKNOWN
	NPropertyPages      DWORD
	LphPropertyPages    *HPROPSHEETPAGE
	NStartPage          DWORD
	DwResultAction      DWORD
}
type PRINTPAGERANGE struct {
	NFromPage DWORD
	NToPage   DWORD
}
type PRIVILEGE_SET struct {
	PrivilegeCount ULONG
	Control        ULONG
	Privilege      [ANYSIZE_ARRAY]LUID_AND_ATTRIBUTES
}
type PROC uintptr
type PROCESS_INFORMATION struct {
	HProcess    HANDLE
	HThread     HANDLE
	DwProcessId DWORD
	DwThreadId  DWORD
}
type PROCESS_MEMORY_COUNTERS struct {
	Cb                         DWORD
	PageFaultCount             DWORD
	PeakWorkingSetSize         SIZE_T
	WorkingSetSize             SIZE_T
	QuotaPeakPagedPoolUsage    SIZE_T
	QuotaPagedPoolUsage        SIZE_T
	QuotaPeakNonPagedPoolUsage SIZE_T
	QuotaNonPagedPoolUsage     SIZE_T
	PagefileUsage              SIZE_T
	PeakPagefileUsage          SIZE_T
}
type PROPENUMPROC func(hWnd HWND, lpszString string, hData HANDLE) BOOL
type PROPENUMPROCEX func(hwnd HWND, lpszString LPWSTR, hData HANDLE, dwData uintptr) BOOL
type PROPSHEETCALLBACK func(unnamed0 HWND, unnamed1 UINT, unnamed2 LPARAM) int32
type PROPSHEETHEADER PROPSHEETHEADER_V2
type PROPSHEETHEADER_V2 struct {
	dwSize       DWORD
	dwFlags      DWORD
	hwndParent   HWND
	hInstance    HINSTANCE
	union1       uintptr
	PszCaption   LPCWSTR
	NPages       UINT
	union2       uintptr
	union3       uintptr
	PfnCallback  uintptr // PFNPROPSHEETCALLBACK
	union4       uintptr
	HplWatermark HPALETTE
	union5       uintptr
}

func (this *PROPSHEETHEADER_V2) HIcon() HICON {
	return HICON(this.union1)
}
func (this *PROPSHEETHEADER_V2) PszIcon() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union1)))
}
func (this *PROPSHEETHEADER_V2) NStartPage() UINT {
	return *(*UINT)(unsafe.Pointer(&this.union2))
}
func (this *PROPSHEETHEADER_V2) PStartPage() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union2)))
}
func (this *PROPSHEETHEADER_V2) Ppsp() /*const*/ *PROPSHEETPAGE {
	return (*PROPSHEETPAGE)(unsafe.Pointer(this.union3))
}
func (this *PROPSHEETHEADER_V2) Phpage() HPROPSHEETPAGE {
	return HPROPSHEETPAGE(this.union3)
}
func (this *PROPSHEETHEADER_V2) HbmWatermark() HBITMAP {
	return HBITMAP(this.union4)
}
func (this *PROPSHEETHEADER_V2) PszbmWatermark() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union4)))
}
func (this *PROPSHEETHEADER_V2) HbmHeader() HBITMAP {
	return HBITMAP(this.union5)
}
func (this *PROPSHEETHEADER_V2) PszbmHeader() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union5)))
}

type PROPSHEETPAGE PROPSHEETPAGE_V4
type PROPSHEETPAGE_RESOURCE *DLGTEMPLATE // const DLGTEMPLATE*
type PROPSHEETPAGE_V4 struct {
	DwSize            DWORD
	DwFlags           DWORD
	HInstance         HINSTANCE
	union1            uintptr
	union2            uintptr
	PszTitle          LPCWSTR
	PfnDlgProc        uintptr // DLGPROC
	LParam            uintptr
	PfnCallback       uintptr // PSPCALLBACK
	PcRefParent       *UINT
	PszHeaderTitle    LPCWSTR
	PszHeaderSubTitle LPCWSTR
	HActCtx           HANDLE
	union3            uintptr
}

func (this *PROPSHEETPAGE_V4) PszTemplate() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.union1))
}
func (this *PROPSHEETPAGE_V4) PResource() *PROPSHEETPAGE_RESOURCE {
	return (*PROPSHEETPAGE_RESOURCE)(unsafe.Pointer(&this.union1))
}
func (this *PROPSHEETPAGE_V4) HIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.union2))
}
func (this *PROPSHEETPAGE_V4) PszIcon() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.union2))
}
func (this *PROPSHEETPAGE_V4) HbmHeader() *HBITMAP {
	return (*HBITMAP)(unsafe.Pointer(&this.union3))
}
func (this *PROPSHEETPAGE_V4) PszbmHeader() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.union3))
}

type PRPC_MESSAGE *RPC_MESSAGE
type PRPC_SYNTAX_IDENTIFIER *RPC_SYNTAX_IDENTIFIER
type PSAPI_WS_WATCH_INFORMATION struct {
	FaultingPc LPVOID
	FaultingVa LPVOID
}
type PSAPI_WS_WATCH_INFORMATION_EX struct {
	BasicInfo        PSAPI_WS_WATCH_INFORMATION
	FaultingThreadId ULONG_PTR
	Flags            ULONG_PTR
}
type PSECURE_MEMORY_CACHE_CALLBACK func(Addr PVOID, Range SIZE_T) BOOLEAN
type PSECURITY_DESCRIPTOR *SECURITY_DESCRIPTOR
type PSID uintptr
type PSPCALLBACK func(hwnd HWND, uMsg UINT, ppsp *PROPSHEETPAGE) UINT
type PSRWLOCK *RTL_SRWLOCK
type PSTR *CHAR
type PTRIVERTEX *TRIVERTEX
type PULONG64 *ULONG64
type PUSHORT *USHORT
type PVOID uintptr
type PWCHAR *WCHAR
type PWSTR *WCHAR
type QITAB struct {
	Piid/*const*/ *IID
	DwOffset int32
}
type QOS struct {
	SendingFlowspec   FLOWSPEC
	ReceivingFlowspec FLOWSPEC
	ProviderSpecific  WSABUF
}
type RASTERIZER_STATUS struct {
	NSize       int16
	WFlags      int16
	NLanguageID int16
}
type RAWHID struct {
	DwSizeHid uint32
	DwCount   uint32
	BRawData  [1]byte
}
type RAWINPUT struct {
	Header RAWINPUTHEADER
	data   [24]byte
}

func (r *RAWINPUT) Mouse() *RAWMOUSE {
	return (*RAWMOUSE)(unsafe.Pointer(&r.data[0]))
}
func (r *RAWINPUT) Keyboard() *RAWKEYBOARD {
	return (*RAWKEYBOARD)(unsafe.Pointer(&r.data[0]))
}
func (r *RAWINPUT) HID() *RAWHID {
	return (*RAWHID)(unsafe.Pointer(&r.data[0]))
}

type RAWINPUTDEVICE struct {
	UsUsagePage uint16
	UsUsage     uint16
	DwFlags     uint32
	HwndTarget  HWND
}
type RAWINPUTDEVICELIST struct {
	HDevice HANDLE
	DwType  uint32
}
type RAWINPUTHEADER struct {
	DwType  uint32
	DwSize  uint32
	HDevice HANDLE
	WParam  uintptr
}
type RAWKEYBOARD struct {
	MakeCode         uint16
	Flags            uint16
	Reserved         int16
	VKey             uint16
	Message          uint32
	ExtraInformation uint32
}
type RAWMOUSE struct {
	UsFlags            uint16
	padding            [2]byte
	UsButtonFlags      uint16
	UsButtonData       uint16
	UlRawButtons       uint32
	LLastX             int32
	LLastY             int32
	UlExtraInformation uint32
}
type RECT struct {
	Left, Top, Right, Bottom int32
}
type RECTFX struct {
	XLeft   FIX
	YTop    FIX
	XRight  FIX
	YBottom FIX
}
type RECTL struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}
type REFCLSID /*const*/ *IID
type REFGUID /*const*/ *GUID
type REFIID /*const*/ *IID
type REGISTERWORDENUMPROC func(lpszReading string, unnamed1 DWORD, lpszString string, unnamed3 LPVOID) int32
type REGSAM uint32
type RGBQUAD struct {
	RgbBlue     BYTE
	RgbGreen    BYTE
	RgbRed      BYTE
	RgbReserved BYTE
}
type RGNDATA struct {
	Rdh    RGNDATAHEADER
	Buffer [1]byte
}
type RGNDATAHEADER struct {
	DwSize   DWORD
	IType    DWORD
	NCount   DWORD
	NRgnSize DWORD
	RcBound  RECT
}
type ROP4 ULONG
type RPC_BINDING_HANDLE I_RPC_HANDLE
type RPC_MESSAGE struct {
	Handle                  RPC_BINDING_HANDLE
	DataRepresentation      uint32
	Buffer                  uintptr
	BufferLength            uint32
	ProcNum                 uint32
	TransferSyntax          PRPC_SYNTAX_IDENTIFIER
	RpcInterfaceInformation uintptr
	ReservedForRuntime      uintptr
	ManagerEpv              uintptr
	ImportContext           uintptr
	RpcFlags                uint32
}
type RPC_STATUS int32
type RPC_SYNTAX_IDENTIFIER struct {
	SyntaxGUID    GUID
	SyntaxVersion RPC_VERSION
}
type RPC_VERSION struct {
	MajorVersion uint16
	MinorVersion uint16
}
type RTL_SRWLOCK struct {
	Ptr PVOID
}
type RUN struct {
	IStart LONG
	IStop  LONG
}
type SAFEARRAY struct {
	CDims      USHORT
	FFeatures  USHORT
	CbElements ULONG
	CLocks     ULONG
	PvData     PVOID
	Rgsabound  [1]SAFEARRAYBOUND
}
type SAFEARRAYBOUND struct {
	CElements ULONG
	LLbound   LONG
}
type SAFER_LEVEL_HANDLE HANDLE
type SCODE LONG
type SCROLLBARINFO struct {
	CbSize        uint32
	RcScrollBar   RECT
	DxyLineButton int32
	XyThumbTop    int32
	XyThumbBottom int32
	Reserved      int32
	Rgstate       [CCHILDREN_SCROLLBAR + 1]uint32
}
type SCROLLINFO struct {
	CbSize    uint32
	FMask     uint32
	NMin      int32
	NMax      int32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}
type SC_HANDLE HANDLE
type SC_LOCK uintptr
type SECURITY_ATTRIBUTES struct {
	NLength              uint32
	LpSecurityDescriptor uintptr
	BInheritHandle       BOOL
}
type SECURITY_CONTEXT_TRACKING_MODE BOOLEAN
type SECURITY_DESCRIPTOR struct {
	Revision byte
	Sbz1     byte
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    PSID
	Group    PSID
	Sacl     *ACL
	Dacl     *ACL
}
type SECURITY_DESCRIPTOR_CONTROL uint16
type SECURITY_IMPERSONATION_LEVEL int32
type SECURITY_INFORMATION ULONG
type SECURITY_QUALITY_OF_SERVICE struct {
	storage [12]byte
}

func (this *SECURITY_QUALITY_OF_SERVICE) Length() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *SECURITY_QUALITY_OF_SERVICE) ImpersonationLevel() *SECURITY_IMPERSONATION_LEVEL {
	return (*SECURITY_IMPERSONATION_LEVEL)(unsafe.Pointer(&this.storage[4]))
}
func (this *SECURITY_QUALITY_OF_SERVICE) ContextTrackingMode() *SECURITY_CONTEXT_TRACKING_MODE {
	return (*SECURITY_CONTEXT_TRACKING_MODE)(unsafe.Pointer(&this.storage[8]))
}
func (this *SECURITY_QUALITY_OF_SERVICE) EffectiveOnly() *BOOLEAN {
	return (*BOOLEAN)(unsafe.Pointer(&this.storage[9]))
}

type SENDASYNCPROC func(hwnd HWND, uMsg uint32, dwData uintptr, lResult LRESULT)
type SERVICE_STATUS struct {
	DwServiceType             DWORD
	DwCurrentState            DWORD
	DwControlsAccepted        DWORD
	DwWin32ExitCode           DWORD
	DwServiceSpecificExitCode DWORD
	DwCheckPoint              DWORD
	DwWaitHint                DWORD
}
type SERVICE_STATUS_HANDLE HANDLE
type SHITEMID struct {
	storage [3]byte
}

func (this *SHITEMID) Cb() *USHORT {
	return (*USHORT)(unsafe.Pointer(&this.storage[0]))
}
func (this *SHITEMID) AbID() *BYTE {
	return (*BYTE)(unsafe.Pointer(&this.storage[2]))
}

type SHORT int16
type SID struct {
	Revision            UCHAR
	SubAuthorityCount   UCHAR
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        [ANYSIZE_ARRAY]ULONG
}
type SID_AND_ATTRIBUTES struct {
	Sid        PSID
	Attributes ULONG
}
type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]UCHAR
}
type SIZE struct {
	Cx, Cy int32
}
type SIZEL SIZE
type SOCKADDR struct {
	Sa_family ADDRESS_FAMILY
	Sa_data   [14]CHAR
}
type SOCKET UINT_PTR
type SOCKET_ADDRESS struct {
	LpSockaddr      LPSOCKADDR
	ISockaddrLength INT
}
type STARTUPINFO struct {
	Cb              DWORD
	LpReserved      LPWSTR
	LpDesktop       LPWSTR
	LpTitle         LPWSTR
	DwX             DWORD
	DwY             DWORD
	DwXSize         DWORD
	DwYSize         DWORD
	DwXCountChars   DWORD
	DwYCountChars   DWORD
	DwFillAttribute DWORD
	DwFlags         DWORD
	WShowWindow     WORD
	CbReserved2     WORD
	LpReserved2     LPBYTE
	HStdInput       HANDLE
	HStdOutput      HANDLE
	HStdError       HANDLE
}
type STROBJ struct {
	CGlyphs     ULONG
	FlAccel     FLONG
	UlCharInc   ULONG
	RclBkGround RECTL
	Pgp         *GLYPHPOS
	PwszOrg     LPWSTR
}

func (this *STRRET) POleStr() *LPWSTR {
	return (*LPWSTR)(unsafe.Pointer(&this.cStr[0]))
}
func (this *STRRET) UOffset() *UINT {
	return (*UINT)(unsafe.Pointer(&this.cStr[0]))
}
func (this *STRRET) CStr() **byte {
	return (**byte)(unsafe.Pointer(&this.cStr[0]))
}

type STYLEBUF struct {
	DwStyle       DWORD
	SzDescription [STYLE_DESCRIPTION_SIZE]WCHAR
}
type SUBCLASSPROC func(hWnd HWND, uMsg UINT, wParam WPARAM, lParam LPARAM, uIdSubclass UINT_PTR, dwRefData DWORD_PTR) LRESULT
type SURFOBJ struct {
	Dhsurf        DHSURF
	Hsurf         HSURF
	Dhpdev        DHPDEV
	Hdev          HDEV
	SizlBitmap    SIZEL
	CjBits        ULONG
	PvBits        PVOID
	PvScan0       PVOID
	LDelta        LONG
	IUniq         ULONG
	IBitmapFormat ULONG
	IType         USHORT
	FjBitmap      USHORT
}
type SYSTEMTIME struct {
	WYear         WORD
	WMonth        WORD
	WDayOfWeek    WORD
	WDay          WORD
	WHour         WORD
	WMinute       WORD
	WSecond       WORD
	WMilliseconds WORD
}
type SYSTEM_INFO struct {
	WProcessorArchitecture      uint16
	WReserved                   uint16
	DwPageSize                  uint32
	LpMinimumApplicationAddress uintptr
	LpMaximumApplicationAddress uintptr
	DwActiveProcessorMask       uintptr
	DwNumberOfProcessors        uint32
	DwProcessorType             uint32
	DwAllocationGranularity     uint32
	WProcessorLevel             uint16
	WProcessorRevision          uint16
}
type Sockaddr struct {
	Sa_family uint16
	Sa_data   [14]CHAR
}
type Socklen_t int32
type Status GpStatus
type TASKDIALOGCALLBACK func(hwnd HWND, msg UINT, wParam WPARAM, lParam LPARAM, lpRefData LONG_PTR) HRESULT
type TEXTMETRIC struct {
	TmHeight           LONG
	TmAscent           LONG
	TmDescent          LONG
	TmInternalLeading  LONG
	TmExternalLeading  LONG
	TmAveCharWidth     LONG
	TmMaxCharWidth     LONG
	TmWeight           LONG
	TmOverhang         LONG
	TmDigitizedAspectX LONG
	TmDigitizedAspectY LONG
	TmFirstChar        WCHAR
	TmLastChar         WCHAR
	TmDefaultChar      WCHAR
	TmBreakChar        WCHAR
	TmItalic           BYTE
	TmUnderlined       BYTE
	TmStruckOut        BYTE
	TmPitchAndFamily   BYTE
	TmCharSet          BYTE
}
type THREAD_START_ROUTINE func(lpThreadParameter LPVOID) DWORD
type TIMERPROC func(hwnd HWND, uMsg uint32, idEvent uintptr, dwTime uint32)
type TITLEBARINFO struct {
	CbSize     uint32
	RcTitleBar RECT
	Rgstate    [CCHILDREN_TITLEBAR + 1]uint32
}
type TOKEN_GROUPS struct {
	GroupCount ULONG
	Groups     [ANYSIZE_ARRAY]SID_AND_ATTRIBUTES
}
type TOKEN_PRIVILEGES struct {
	PrivilegeCount ULONG
	Privileges     [ANYSIZE_ARRAY]LUID_AND_ATTRIBUTES
}
type TOUCHINPUT struct {
	X           int32 // LONG
	Y           int32 // LONG
	HSource     HANDLE
	DwID        uint32
	DwFlags     uint32
	DwMask      uint32
	DwTime      uint32
	DwExtraInfo uintptr // ULONG_PTR
	CxContact   uint32
	CyContact   uint32
}
type TPMPARAMS struct {
	CbSize    uint32
	RcExclude RECT
}
type TRACKMOUSEEVENT struct {
	CbSize      uint32
	DwFlags     uint32
	HwndTrack   HWND
	DwHoverTime uint32
}
type TRIVERTEX struct {
	X     LONG
	Y     LONG
	Red   COLOR16
	Green COLOR16
	Blue  COLOR16
	Alpha COLOR16
}
type TRUSTEE struct {
	PMultipleTrustee         *TRUSTEE
	MultipleTrusteeOperation MULTIPLE_TRUSTEE_OPERATION
	TrusteeForm              TRUSTEE_FORM
	TrusteeType              TRUSTEE_TYPE
	PtstrName                LPWSTR
}
type Time_t int64
type Timeval struct {
	Tv_sec  int32
	Tv_usec int32
}
type UCHAR byte
type UDATE struct {
	St         SYSTEMTIME
	WDayOfYear USHORT
}
type UINT uint32
type UINT8 uint8
type UINT_PTR uintptr
type ULARGE_INTEGER struct {
	QuadPart ULONGLONG
}
type ULONG uint32
type ULONG64 uint64
type ULONGLONG uint64
type ULONG_PTR *uint32
type UPDATELAYEREDWINDOWINFO struct {
	CbSize   uint32
	HdcDst   HDC
	PptDst   *POINT // const POINT*
	Psize    *SIZE  // const SIZE*
	HdcSrc   HDC
	PptSrc   *POINT // const POINT*
	CrKey    COLORREF
	Pblend   uintptr // const BLENDFUNCTION*
	DwFlags  uint32
	PrcDirty *RECT // const RECT*
}
type USER_MARSHAL_FREEING_ROUTINE func(unnamed0 *uint32, unnamed1 uintptr)
type USER_MARSHAL_MARSHALLING_ROUTINE func(unnamed0 *uint32, unnamed1 *byte, unnamed2 uintptr) *byte
type USER_MARSHAL_ROUTINE_QUADRUPLE struct {
	PfnBufferSize uintptr // USER_MARSHAL_SIZING_ROUTINE
	PfnMarshall   uintptr // USER_MARSHAL_MARSHALLING_ROUTINE
	PfnUnmarshall uintptr // USER_MARSHAL_UNMARSHALLING_ROUTINE
	PfnFree       uintptr // USER_MARSHAL_FREEING_ROUTINE
}
type USER_MARSHAL_SIZING_ROUTINE func(unnamed0 *uint32, unnamed1 uint32, unnamed2 uintptr) uint32
type USER_MARSHAL_UNMARSHALLING_ROUTINE func(unnamed0 *uint32, unnamed1 *byte, unnamed2 uintptr) *byte
type USHORT uint16

func (this *VARIANT) Vt() *VARTYPE {
	return (*VARTYPE)(unsafe.Pointer(&this.union1[0]))
}
func (this *VARIANT) WReserved1() *WORD {
	return (*WORD)(unsafe.Pointer(&this.union1[2]))
}
func (this *VARIANT) WReserved2() *WORD {
	return (*WORD)(unsafe.Pointer(&this.union1[4]))
}
func (this *VARIANT) WReserved3() *WORD {
	return (*WORD)(unsafe.Pointer(&this.union1[6]))
}
func (this *VARIANT) LlVal() *LONGLONG        { return (*LONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) LVal() *LONG             { return (*LONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) BVal() *BYTE             { return (*BYTE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) IVal() *SHORT            { return (*SHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) FltVal() *FLOAT          { return (*FLOAT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) DblVal() *DOUBLE         { return (*DOUBLE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) BoolVal() *VARIANT_BOOL  { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Bool() *VARIANT_BOOL     { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Scode() *SCODE           { return (*SCODE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) CyVal() *CY              { return (*CY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Date() *DATE             { return (*DATE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) BstrVal() *BSTR          { return (*BSTR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PunkVal() *IUnknown      { return (*IUnknown)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PdispVal() *IDispatch    { return (*IDispatch)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Parray() *SAFEARRAY      { return (*SAFEARRAY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PbVal() *BYTE            { return (*BYTE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PiVal() *SHORT           { return (*SHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PlVal() *LONG            { return (*LONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PllVal() *LONGLONG       { return (*LONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PfltVal() *FLOAT         { return (*FLOAT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PdblVal() *DOUBLE        { return (*DOUBLE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PboolVal() *VARIANT_BOOL { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pbool() *VARIANT_BOOL    { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pscode() *SCODE          { return (*SCODE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PcyVal() *CY             { return (*CY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pdate() *DATE            { return (*DATE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PbstrVal() *BSTR         { return (*BSTR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PpunkVal() **IUnknown    { return (**IUnknown)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PpdispVal() **IDispatch  { return (**IDispatch)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pparray() **SAFEARRAY    { return (**SAFEARRAY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PvarVal() *VARIANT       { return (*VARIANT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Byref() PVOID            { return (PVOID)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) CVal() CHAR              { return *(*CHAR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UiVal() USHORT           { return *(*USHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UlVal() ULONG            { return *(*ULONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UllVal() ULONGLONG       { return *(*ULONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) IntVal() INT             { return *(*INT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UintVal() UINT           { return *(*UINT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PdecVal() *DECIMAL       { return (*DECIMAL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PcVal() *CHAR            { return (*CHAR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PuiVal() *USHORT         { return (*USHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PulVal() *ULONG          { return (*ULONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PullVal() *ULONGLONG     { return (*ULONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PintVal() *INT           { return (*INT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PuintVal() *UINT         { return (*UINT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PvRecord() PVOID {
	return (PVOID)(unsafe.Pointer(&this.union1[8]))
}
func (this *VARIANT) DecVal() DECIMAL {
	return *(*DECIMAL)(unsafe.Pointer(&this.union1[0]))
}
func unpackVARIANT(v VARIANT) []uintptr {
	size := int(unsafe.Sizeof(v))
	size += size % 4
	step := 4
	n := size / step
	ret := []uintptr{}
	ptr := uintptr(unsafe.Pointer(&v))
	for i := 0; i < n; i++ {
		ret = append(ret, *(*uintptr)(unsafe.Pointer(ptr + uintptr(step*i))))
	}
	return ret
}

type VARIANTARG VARIANT
type VARIANT_BOOL int16
type VARTYPE uint16
type WCHAR uint16
type WCRANGE struct {
	WcLow   WCHAR
	CGlyphs USHORT
}
type WGLSWAP struct {
	Hdc     HDC
	UiFlags uint32 // UINT
}
type WINDOWINFO struct {
	CbSize          uint32 // DWORD
	RcWindow        RECT
	RcClient        RECT
	DwStyle         uint32 // DWORD
	DwExStyle       uint32 // DWORD
	DwWindowStatus  uint32 // DWORD
	CxWindowBorders uint32 // UINT
	CyWindowBorders uint32 // UINT
	AtomWindowType  ATOM
	WCreatorVersion uint16 // WORD
}
type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    POINT
	PtMaxPosition    POINT
	RcNormalPosition RECT
}
type WINEVENTPROC func(hWinEventHook HWINEVENTHOOK, event uint32, hwnd HWND, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32)
type WINSTAENUMPROC func(lpszWindowStation LPWSTR, lParam LPARAM) BOOL
type WNDCLASS struct {
	Style         uint32
	LpfnWndProc   uintptr // WNDPROC
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16 // LPCWSTR
	LpszClassName *uint16 // LPCWSTR
}
type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}
type WNDENUMPROC func(hWnd HWND, lParam LPARAM) BOOL
type WNDPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) LRESULT
type WORD uint16
type WPARAM uintptr
type WSABUF struct {
	Len ULONG
	Buf *CHAR
}
type WSACOMPLETION struct {
	storage [4]uintptr
}
type WSACOMPLETION_WindowMessage struct {
	HWnd    HWND
	UMsg    UINT
	Context WPARAM
}
type WSACOMPLETION_Event struct {
	LpOverlapped LPWSAOVERLAPPED
}
type WSACOMPLETION_Apc struct {
	LpOverlapped       LPWSAOVERLAPPED
	LpfnCompletionProc uintptr // LPWSAOVERLAPPED_COMPLETION_ROUTINE
}
type WSACOMPLETION_Port struct {
	LpOverlapped LPWSAOVERLAPPED
	HPort        HANDLE
	Key          ULONG_PTR
}

func (this *WSACOMPLETION) Type() *WSACOMPLETIONTYPE {
	return (*WSACOMPLETIONTYPE)(unsafe.Pointer(&this))
}
func (this *WSACOMPLETION) WindowMessage() *WSACOMPLETION_WindowMessage {
	return (*WSACOMPLETION_WindowMessage)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}
func (this *WSACOMPLETION) Event() *WSACOMPLETION_Event {
	return (*WSACOMPLETION_Event)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}
func (this *WSACOMPLETION) Apc() *WSACOMPLETION_Apc {
	return (*WSACOMPLETION_Apc)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}
func (this *WSACOMPLETION) Port() *WSACOMPLETION_Port {
	return (*WSACOMPLETION_Port)(unsafe.Pointer(uintptr(unsafe.Pointer(&this)) + uintptr(4)))
}

type WSAEVENT HANDLE
type WSAMSG struct {
	Name          LPSOCKADDR
	Namelen       INT
	LpBuffers     LPWSABUF
	DwBufferCount ULONG
	Control       WSABUF
	DwFlags       ULONG
}
type WSANAMESPACE_INFO struct {
	NSProviderId   GUID
	DwNameSpace    DWORD
	FActive        BOOL
	DwVersion      DWORD
	LpszIdentifier LPWSTR
}
type WSANETWORKEVENTS struct {
	LNetworkEvents int32
	IErrorCode     [FD_MAX_EVENTS]int32
}
type WSANSCLASSINFO struct {
	LpszName    LPWSTR
	DwNameSpace DWORD
	DwValueType DWORD
	DwValueSize DWORD
	LpValue     LPVOID
}
type WSAPOLLFD struct {
	Fd      SOCKET
	Events  SHORT
	Revents SHORT
}
type WSAPROTOCOLCHAIN struct {
	ChainLen     int32
	ChainEntries [MAX_PROTOCOL_CHAIN]DWORD
}
type WSAPROTOCOL_INFO struct {
	DwServiceFlags1    DWORD
	DwServiceFlags2    DWORD
	DwServiceFlags3    DWORD
	DwServiceFlags4    DWORD
	DwProviderFlags    DWORD
	ProviderId         GUID
	DwCatalogEntryId   DWORD
	ProtocolChain      WSAPROTOCOLCHAIN
	IVersion           int32
	IAddressFamily     int32
	IMaxSockAddr       int32
	IMinSockAddr       int32
	ISocketType        int32
	IProtocol          int32
	IProtocolMaxOffset int32
	INetworkByteOrder  int32
	ISecurityScheme    int32
	DwMessageSize      DWORD
	DwProviderReserved DWORD
	SzProtocol         [WSAPROTOCOL_LEN + 1]WCHAR
}
type WSAQUERYSET struct {
	DwSize                  DWORD
	LpszServiceInstanceName LPWSTR
	LpServiceClassId        LPGUID
	LpVersion               LPWSAVERSION
	LpszComment             LPWSTR
	DwNameSpace             DWORD
	LpNSProviderId          LPGUID
	LpszContext             LPWSTR
	DwNumberOfProtocols     DWORD
	LpafpProtocols          LPAFPROTOCOLS
	LpszQueryString         LPWSTR
	DwNumberOfCsAddrs       DWORD
	LpcsaBuffer             LPCSADDR_INFO
	DwOutputFlags           DWORD
	LpBlob                  LPBLOB
}
type WSASERVICECLASSINFO struct {
	LpServiceClassId     LPGUID
	LpszServiceClassName LPWSTR
	DwCount              DWORD
	LpClassInfos         LPWSANSCLASSINFO
}
type WSAVERSION struct {
	DwVersion DWORD
	EcHow     WSAECOMPARATOR
}
type XFORM struct {
	EM11 FLOAT
	EM12 FLOAT
	EM21 FLOAT
	EM22 FLOAT
	EDx  FLOAT
	EDy  FLOAT
}
type XFORML struct {
	EM11 FLOATL
	EM12 FLOATL
	EM21 FLOATL
	EM22 FLOATL
	EDx  FLOATL
	EDy  FLOATL
}
type XFORMOBJ struct {
	UlReserved ULONG
}
type XLATEOBJ struct {
	IUniq    ULONG
	FlXlate  FLONG
	ISrcType USHORT
	IDstType USHORT
	CEntries ULONG
	PulXlate *ULONG
}
type XMIT_HELPER_ROUTINE func(unnamed PMIDL_STUB_MESSAGE)
type XMIT_ROUTINE_QUINTUPLE struct {
	PfnTranslateToXmit   uintptr // XMIT_HELPER_ROUTINE
	PfnTranslateFromXmit uintptr // XMIT_HELPER_ROUTINE
	PfnFreeXmit          uintptr // XMIT_HELPER_ROUTINE
	PfnFreeInst          uintptr // XMIT_HELPER_ROUTINE
}

func Typeof(name string) reflect.Type {
	return types[name]
}

func TypeNames() []string {
	ret := make([]string, len(types))
	i := 0
	for name, _ := range types {
		ret[i] = name
		i++
	}
	return ret
}
