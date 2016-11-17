package win

import (
	"unsafe"
)

type GESTUREINFO struct {
	CbSize       uint32 // UINT
	DwFlags      uint32
	DwID         uint32
	HwndTarget   HWND
	PtsLocation  POINTS
	DwInstanceID uint32
	DwSequenceID uint32
	UllArguments int64  // ULONGLONG
	CbExtraArgs  uint32 // UINT
}
type INPUT struct {
	Type uint32
	data [36]byte
}
func (this *INPUT) Mi() *MOUSEINPUT {
	return (*MOUSEINPUT)(unsafe.Pointer(&this.data[0]))
}
func (this *INPUT) Ki() *KEYBDINPUT {
	return (*KEYBDINPUT)(unsafe.Pointer(&this.data[0]))
}
func (this *INPUT) Hi() *HARDWAREINPUT {
	return (*HARDWAREINPUT)(unsafe.Pointer(&this.data[0]))
}
type INPUTCONTEXT struct {
	storage [352]byte
}
func (this *INPUTCONTEXT) HWnd() *HWND { // 8
	return (*HWND)(unsafe.Pointer(&this.storage[0]))
}
func (this *INPUTCONTEXT) FOpen() *BOOL { // 4
	return (*BOOL)(unsafe.Pointer(&this.storage[8]))
}
func (this *INPUTCONTEXT) PtStatusWndPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[12]))
}
func (this *INPUTCONTEXT) PtSoftKbdPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[20]))
}
func (this *INPUTCONTEXT) FdwConversion() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *INPUTCONTEXT) FdwSentence() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[32]))
}
func (this *INPUTCONTEXT) LfFont() *LOGFONT { // 92
	return (*LOGFONT)(unsafe.Pointer(&this.storage[36]))
}
func (this *INPUTCONTEXT) CfCompForm() *COMPOSITIONFORM { // 28
	return (*COMPOSITIONFORM)(unsafe.Pointer(&this.storage[128]))
}
func (this *INPUTCONTEXT) CfCandForm() *[4]CANDIDATEFORM { // 128
	return (*[4]CANDIDATEFORM)(unsafe.Pointer(&this.storage[156]))
}
func (this *INPUTCONTEXT) HCompStr() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[284]))
}
func (this *INPUTCONTEXT) HCandInfo() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[292]))
}
func (this *INPUTCONTEXT) HGuideLine() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[300]))
}
func (this *INPUTCONTEXT) HPrivate() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[308]))
}
func (this *INPUTCONTEXT) DwNumMsgBuf() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[316]))
}
func (this *INPUTCONTEXT) HMsgBuf() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[320]))
}
func (this *INPUTCONTEXT) FdwInit() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[328]))
}
func (this *INPUTCONTEXT) DwReserve() *[3]DWORD { // 12
	return (*[3]DWORD)(unsafe.Pointer(&this.storage[332]))
}
type IP_ADAPTER_ADDRESSES_LH struct {
	union1                 ULONGLONG
	Next                   *IP_ADAPTER_ADDRESSES_LH
	AdapterName            PCHAR
	FirstUnicastAddress    PIP_ADAPTER_UNICAST_ADDRESS_LH
	FirstAnycastAddress    PIP_ADAPTER_ANYCAST_ADDRESS_XP
	FirstMulticastAddress  PIP_ADAPTER_MULTICAST_ADDRESS_XP
	FirstDnsServerAddress  PIP_ADAPTER_DNS_SERVER_ADDRESS_XP
	DnsSuffix              PWCHAR
	Description            PWCHAR
	FriendlyName           PWCHAR
	PhysicalAddress        [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	PhysicalAddressLength  ULONG
	union2                 ULONG
	Mtu                    ULONG
	IfType                 IFTYPE
	OperStatus             IF_OPER_STATUS
	Ipv6IfIndex            IF_INDEX
	ZoneIndices            [16]ULONG
	FirstPrefix            PIP_ADAPTER_PREFIX_XP
	TransmitLinkSpeed      ULONG64
	ReceiveLinkSpeed       ULONG64
	FirstWinsServerAddress PIP_ADAPTER_WINS_SERVER_ADDRESS_LH
	FirstGatewayAddress    PIP_ADAPTER_GATEWAY_ADDRESS_LH
	Ipv4Metric             ULONG
	Ipv6Metric             ULONG
	Luid                   IF_LUID
	Dhcpv4Server           SOCKET_ADDRESS
	CompartmentId          NET_IF_COMPARTMENT_ID
	NetworkGuid            NET_IF_NETWORK_GUID
	ConnectionType         NET_IF_CONNECTION_TYPE
	TunnelType             TUNNEL_TYPE
	Dhcpv6Server           SOCKET_ADDRESS
	Dhcpv6ClientDuid       [MAX_DHCPV6_DUID_LENGTH]BYTE
	Dhcpv6ClientDuidLength ULONG
	Dhcpv6Iaid             ULONG
	FirstDnsSuffix         PIP_ADAPTER_DNS_SUFFIX
}
type IP_ADAPTER_ANYCAST_ADDRESS_XP struct {
	union1  ULONGLONG
	Next    *IP_ADAPTER_ANYCAST_ADDRESS_XP
	Address SOCKET_ADDRESS
}
type IP_ADAPTER_DNS_SERVER_ADDRESS_XP struct {
	union1  ULONGLONG
	Next    *IP_ADAPTER_DNS_SERVER_ADDRESS_XP
	Address SOCKET_ADDRESS
}
type IP_ADAPTER_GATEWAY_ADDRESS_LH struct {
	union1  ULONGLONG
	Next    *IP_ADAPTER_GATEWAY_ADDRESS_LH
	Address SOCKET_ADDRESS
}
type IP_ADAPTER_MULTICAST_ADDRESS_XP struct {
	union1  ULONGLONG
	Next    *IP_ADAPTER_MULTICAST_ADDRESS_XP
	Address SOCKET_ADDRESS
}
type IP_ADAPTER_WINS_SERVER_ADDRESS_LH struct {
	union1  ULONGLONG
	Next    *IP_ADAPTER_WINS_SERVER_ADDRESS_LH
	Address SOCKET_ADDRESS
}
type MIB_UDP6ROW_OWNER_MODULE struct {
	UcLocalAddr       [16]UCHAR
	DwLocalScopeId    DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
type MIB_UDPROW_OWNER_MODULE struct {
	DwLocalAddr       DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
type MIDIHDR struct {
	storage [112]byte
}
func (this *MIDIHDR) LpData() *LPSTR {
	return (*LPSTR)(unsafe.Pointer(&this.storage[0]))
}
func (this *MIDIHDR) DwBufferLength() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[8]))
}
func (this *MIDIHDR) DwBytesRecorded() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[12]))
}
func (this *MIDIHDR) DwUser() *DWORD_PTR {
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[16]))
}
func (this *MIDIHDR) DwFlags() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[24]))
}
func (this *MIDIHDR) LpNext() **MIDIHDR {
	return (**MIDIHDR)(unsafe.Pointer(&this.storage[28]))
}
func (this *MIDIHDR) Reserved() *DWORD_PTR {
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[36]))
}
func (this *MIDIHDR) DwOffset() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[44]))
}
func (this *MIDIHDR) DwReserved() *[8]DWORD_PTR {
	return (*[8]DWORD_PTR)(unsafe.Pointer(&this.storage[48]))
}
type PDH_RAW_COUNTER struct {
	CStatus     DWORD
	TimeStamp   FILETIME
	FirstValue  LONGLONG
	SecondValue LONGLONG
	MultiCount  DWORD
}
type PRINTDLG struct {
	storage [120]byte
}
func (this *PRINTDLG) LStructSize() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *PRINTDLG) HwndOwner() *HWND { // 8
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *PRINTDLG) HDevMode() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[12]))
}
func (this *PRINTDLG) HDevNames() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[20]))
}
func (this *PRINTDLG) HDC() *HDC { // 8
	return (*HDC)(unsafe.Pointer(&this.storage[28]))
}
func (this *PRINTDLG) Flags() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[36]))
}
func (this *PRINTDLG) NFromPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[40]))
}
func (this *PRINTDLG) NToPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[42]))
}
func (this *PRINTDLG) NMinPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[44]))
}
func (this *PRINTDLG) NMaxPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[46]))
}
func (this *PRINTDLG) NCopies() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[48]))
}
func (this *PRINTDLG) HInstance() *HINSTANCE { // 8
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[56]))
}
func (this *PRINTDLG) LCustData() *LPARAM { // 8
	return (*LPARAM)(unsafe.Pointer(&this.storage[64]))
}
func (this *PRINTDLG) LpfnPrintHook() *uintptr { // 8
	return (*uintptr)(unsafe.Pointer(&this.storage[72]))
}
func (this *PRINTDLG) LpfnSetupHook() *uintptr { // 8
	return (*uintptr)(unsafe.Pointer(&this.storage[80]))
}
func (this *PRINTDLG) LpPrintTemplateName() *LPCWSTR { // 8
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[88]))
}
func (this *PRINTDLG) LpSetupTemplateName() *LPCWSTR { // 8
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[96]))
}
func (this *PRINTDLG) HPrintTemplate() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[104]))
}
func (this *PRINTDLG) HSetupTemplate() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[112]))
}
type SIZE_T uint64
type STRRET struct {
	UType    UINT
	padding1 [4]byte
	cStr     [260]byte
	padding2 [4]byte
}
type TASKDIALOGCONFIG struct {
	storage [160]byte
}
func (this *TASKDIALOGCONFIG) CbSize() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[0]))
}
func (this *TASKDIALOGCONFIG) HwndParent() *HWND {
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *TASKDIALOGCONFIG) HInstance() *HINSTANCE {
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[12]))
}
func (this *TASKDIALOGCONFIG) DwFlags() *TASKDIALOG_FLAGS {
	return (*TASKDIALOG_FLAGS)(unsafe.Pointer(&this.storage[20]))
}
func (this *TASKDIALOGCONFIG) DwCommonButtons() *TASKDIALOG_COMMON_BUTTON_FLAGS {
	return (*TASKDIALOG_COMMON_BUTTON_FLAGS)(unsafe.Pointer(&this.storage[24]))
}
func (this *TASKDIALOGCONFIG) PszWindowTitle() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[28]))
}
func (this *TASKDIALOGCONFIG) HMainIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.storage[36]))
}
func (this *TASKDIALOGCONFIG) PszMainIcon() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[36]))
}
func (this *TASKDIALOGCONFIG) PszMainInstruction() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[44]))
}
func (this *TASKDIALOGCONFIG) PszContent() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[52]))
}
func (this *TASKDIALOGCONFIG) CButtons() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[60]))
}
func (this *TASKDIALOGCONFIG) PButtons() **TASKDIALOG_BUTTON {
	return (**TASKDIALOG_BUTTON)(unsafe.Pointer(&this.storage[64]))
}
func (this *TASKDIALOGCONFIG) NDefaultButton() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[72]))
}
func (this *TASKDIALOGCONFIG) CRadioButtons() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[76]))
}
func (this *TASKDIALOGCONFIG) PRadioButtons() **TASKDIALOG_BUTTON {
	return (**TASKDIALOG_BUTTON)(unsafe.Pointer(&this.storage[80]))
}
func (this *TASKDIALOGCONFIG) NDefaultRadioButton() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[88]))
}
func (this *TASKDIALOGCONFIG) PszVerificationText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[92]))
}
func (this *TASKDIALOGCONFIG) PszExpandedInformation() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[100]))
}
func (this *TASKDIALOGCONFIG) PszExpandedControlText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[108]))
}
func (this *TASKDIALOGCONFIG) PszCollapsedControlText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[116]))
}
func (this *TASKDIALOGCONFIG) HFooterIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.storage[124]))
}
func (this *TASKDIALOGCONFIG) PszFooterIcon() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[124]))
}
func (this *TASKDIALOGCONFIG) PszFooter() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[132]))
}
func (this *TASKDIALOGCONFIG) PfCallback() *uintptr {
	return (*uintptr)(unsafe.Pointer(&this.storage[140]))
}
func (this *TASKDIALOGCONFIG) LpCallbackData() *LONG_PTR {
	return (*LONG_PTR)(unsafe.Pointer(&this.storage[148]))
}
func (this *TASKDIALOGCONFIG) CxWidth() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[156]))
}
type TASKDIALOG_BUTTON struct {
	storage [12]byte
}
func (this *TASKDIALOG_BUTTON) NButtonID() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[0]))
}
func (this *TASKDIALOG_BUTTON) PszButtonText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[4]))
}
type TBBUTTON struct {
	IBitmap   int32
	IdCommand int32
	FsState   byte
	FsStyle   byte
	BReserved [6]byte
	DwData    *DWORD
	IString   uintptr
}
type VARIANT struct {
	union1 [24]byte
}
func (this *VARIANT) PRecInfo() *IRecordInfo {
	return (*IRecordInfo)(unsafe.Pointer(&this.union1[16]))
}
type WSADATA struct {
	WVersion       uint16
	WHighVersion   uint16
	IMaxSockets    uint16
	IMaxUdpDg      uint16
	LpVendorInfo   *byte
	SzDescription  [WSADESCRIPTION_LEN + 1]byte
	SzSystemStatus [WSASYS_STATUS_LEN + 1]byte
}
