// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	librpcrt4 uintptr

	// Functions
	cStdStubBuffer_AddRef                    uintptr
	cStdStubBuffer_Connect                   uintptr
	cStdStubBuffer_CountRefs                 uintptr
	cStdStubBuffer_DebugServerQueryInterface uintptr
	cStdStubBuffer_DebugServerRelease        uintptr
	cStdStubBuffer_Disconnect                uintptr
	cStdStubBuffer_QueryInterface            uintptr
	createStubFromTypeInfo                   uintptr
	iUnknown_AddRef_Proxy                    uintptr
	iUnknown_QueryInterface_Proxy            uintptr
	iUnknown_Release_Proxy                   uintptr
	i_RpcBindingInqLocalClientPID            uintptr
	i_RpcExceptionFilter                     uintptr
	i_RpcFree                                uintptr
	i_RpcFreeBuffer                          uintptr
	i_RpcGetBuffer                           uintptr
	i_RpcGetCurrentCallHandle                uintptr
	i_RpcMapWin32Status                      uintptr
	i_RpcNegotiateTransferSyntax             uintptr
	i_RpcReceive                             uintptr
	i_RpcSend                                uintptr
	i_RpcSendReceive                         uintptr
	nDRSContextMarshall                      uintptr
	nDRSContextMarshall2                     uintptr
	nDRSContextMarshallEx                    uintptr
	nDRSContextUnmarshall                    uintptr
	nDRSContextUnmarshall2                   uintptr
	nDRSContextUnmarshallEx                  uintptr
	ndrByteCountPointerBufferSize            uintptr
	ndrByteCountPointerFree                  uintptr
	ndrClearOutParameters                    uintptr
	ndrComplexArrayBufferSize                uintptr
	ndrComplexArrayFree                      uintptr
	ndrComplexArrayMemorySize                uintptr
	ndrComplexStructBufferSize               uintptr
	ndrComplexStructFree                     uintptr
	ndrComplexStructMemorySize               uintptr
	ndrConformantArrayBufferSize             uintptr
	ndrConformantArrayFree                   uintptr
	ndrConformantArrayMemorySize             uintptr
	ndrConformantStringBufferSize            uintptr
	ndrConformantStringMemorySize            uintptr
	ndrConformantStructBufferSize            uintptr
	ndrConformantStructFree                  uintptr
	ndrConformantStructMemorySize            uintptr
	ndrConformantVaryingArrayBufferSize      uintptr
	ndrConformantVaryingArrayFree            uintptr
	ndrConformantVaryingArrayMemorySize      uintptr
	ndrConformantVaryingStructBufferSize     uintptr
	ndrConformantVaryingStructFree           uintptr
	ndrConformantVaryingStructMemorySize     uintptr
	ndrContextHandleInitialize               uintptr
	ndrContextHandleSize                     uintptr
	ndrConvert                               uintptr
	ndrConvert2                              uintptr
	ndrCorrelationFree                       uintptr
	ndrCorrelationInitialize                 uintptr
	ndrCorrelationPass                       uintptr
	ndrEncapsulatedUnionBufferSize           uintptr
	ndrEncapsulatedUnionFree                 uintptr
	ndrEncapsulatedUnionMemorySize           uintptr
	ndrFixedArrayBufferSize                  uintptr
	ndrFixedArrayFree                        uintptr
	ndrFixedArrayMemorySize                  uintptr
	ndrFreeBuffer                            uintptr
	ndrInterfacePointerBufferSize            uintptr
	ndrInterfacePointerFree                  uintptr
	ndrInterfacePointerMemorySize            uintptr
	ndrNonConformantStringBufferSize         uintptr
	ndrNonConformantStringMemorySize         uintptr
	ndrNonEncapsulatedUnionBufferSize        uintptr
	ndrNonEncapsulatedUnionFree              uintptr
	ndrNonEncapsulatedUnionMemorySize        uintptr
	ndrOleFree                               uintptr
	ndrPointerBufferSize                     uintptr
	ndrPointerFree                           uintptr
	ndrPointerMemorySize                     uintptr
	ndrProxyErrorHandler                     uintptr
	ndrProxyFreeBuffer                       uintptr
	ndrProxyGetBuffer                        uintptr
	ndrProxySendReceive                      uintptr
	ndrRpcSmSetClientToOsf                   uintptr
	ndrServerCall2                           uintptr
	ndrServerContextMarshall                 uintptr
	ndrServerContextNewMarshall              uintptr
	ndrServerContextNewUnmarshall            uintptr
	ndrServerContextUnmarshall               uintptr
	ndrSimpleStructBufferSize                uintptr
	ndrSimpleStructFree                      uintptr
	ndrSimpleStructMemorySize                uintptr
	ndrUserMarshalBufferSize                 uintptr
	ndrUserMarshalFree                       uintptr
	ndrUserMarshalMemorySize                 uintptr
	ndrVaryingArrayBufferSize                uintptr
	ndrVaryingArrayFree                      uintptr
	ndrVaryingArrayMemorySize                uintptr
	ndrXmitOrRepAsBufferSize                 uintptr
	ndrXmitOrRepAsFree                       uintptr
	ndrXmitOrRepAsMemorySize                 uintptr
	rpcBindingFree                           uintptr
	rpcBindingSetOption                      uintptr
	rpcImpersonateClient                     uintptr
	rpcMgmtEnableIdleCleanup                 uintptr
	rpcMgmtIsServerListening                 uintptr
	rpcMgmtSetComTimeout                     uintptr
	rpcMgmtSetServerStackSize                uintptr
	rpcMgmtStopServerListening               uintptr
	rpcMgmtWaitServerListen                  uintptr
	rpcRevertToSelf                          uintptr
	rpcRevertToSelfEx                        uintptr
	rpcServerListen                          uintptr
	rpcSmDestroyClientContext                uintptr
	rpcSsDestroyClientContext                uintptr
	rpcSsDontSerializeContext                uintptr
)

func init() {
	// Library
	librpcrt4 = doLoadLibrary("rpcrt4.dll")

	// Functions
	cStdStubBuffer_AddRef = doGetProcAddress(librpcrt4, "CStdStubBuffer_AddRef")
	cStdStubBuffer_Connect = doGetProcAddress(librpcrt4, "CStdStubBuffer_Connect")
	cStdStubBuffer_CountRefs = doGetProcAddress(librpcrt4, "CStdStubBuffer_CountRefs")
	cStdStubBuffer_DebugServerQueryInterface = doGetProcAddress(librpcrt4, "CStdStubBuffer_DebugServerQueryInterface")
	cStdStubBuffer_DebugServerRelease = doGetProcAddress(librpcrt4, "CStdStubBuffer_DebugServerRelease")
	cStdStubBuffer_Disconnect = doGetProcAddress(librpcrt4, "CStdStubBuffer_Disconnect")
	cStdStubBuffer_QueryInterface = doGetProcAddress(librpcrt4, "CStdStubBuffer_QueryInterface")
	createStubFromTypeInfo = doGetProcAddress(librpcrt4, "CreateStubFromTypeInfo")
	iUnknown_AddRef_Proxy = doGetProcAddress(librpcrt4, "IUnknown_AddRef_Proxy")
	iUnknown_QueryInterface_Proxy = doGetProcAddress(librpcrt4, "IUnknown_QueryInterface_Proxy")
	iUnknown_Release_Proxy = doGetProcAddress(librpcrt4, "IUnknown_Release_Proxy")
	i_RpcBindingInqLocalClientPID = doGetProcAddress(librpcrt4, "I_RpcBindingInqLocalClientPID")
	i_RpcExceptionFilter = doGetProcAddress(librpcrt4, "I_RpcExceptionFilter")
	i_RpcFree = doGetProcAddress(librpcrt4, "I_RpcFree")
	i_RpcFreeBuffer = doGetProcAddress(librpcrt4, "I_RpcFreeBuffer")
	i_RpcGetBuffer = doGetProcAddress(librpcrt4, "I_RpcGetBuffer")
	i_RpcGetCurrentCallHandle = doGetProcAddress(librpcrt4, "I_RpcGetCurrentCallHandle")
	i_RpcMapWin32Status = doGetProcAddress(librpcrt4, "I_RpcMapWin32Status")
	i_RpcNegotiateTransferSyntax = doGetProcAddress(librpcrt4, "I_RpcNegotiateTransferSyntax")
	i_RpcReceive = doGetProcAddress(librpcrt4, "I_RpcReceive")
	i_RpcSend = doGetProcAddress(librpcrt4, "I_RpcSend")
	i_RpcSendReceive = doGetProcAddress(librpcrt4, "I_RpcSendReceive")
	nDRSContextMarshall = doGetProcAddress(librpcrt4, "NDRSContextMarshall")
	nDRSContextMarshall2 = doGetProcAddress(librpcrt4, "NDRSContextMarshall2")
	nDRSContextMarshallEx = doGetProcAddress(librpcrt4, "NDRSContextMarshallEx")
	nDRSContextUnmarshall = doGetProcAddress(librpcrt4, "NDRSContextUnmarshall")
	nDRSContextUnmarshall2 = doGetProcAddress(librpcrt4, "NDRSContextUnmarshall2")
	nDRSContextUnmarshallEx = doGetProcAddress(librpcrt4, "NDRSContextUnmarshallEx")
	ndrByteCountPointerBufferSize = doGetProcAddress(librpcrt4, "NdrByteCountPointerBufferSize")
	ndrByteCountPointerFree = doGetProcAddress(librpcrt4, "NdrByteCountPointerFree")
	ndrClearOutParameters = doGetProcAddress(librpcrt4, "NdrClearOutParameters")
	ndrComplexArrayBufferSize = doGetProcAddress(librpcrt4, "NdrComplexArrayBufferSize")
	ndrComplexArrayFree = doGetProcAddress(librpcrt4, "NdrComplexArrayFree")
	ndrComplexArrayMemorySize = doGetProcAddress(librpcrt4, "NdrComplexArrayMemorySize")
	ndrComplexStructBufferSize = doGetProcAddress(librpcrt4, "NdrComplexStructBufferSize")
	ndrComplexStructFree = doGetProcAddress(librpcrt4, "NdrComplexStructFree")
	ndrComplexStructMemorySize = doGetProcAddress(librpcrt4, "NdrComplexStructMemorySize")
	ndrConformantArrayBufferSize = doGetProcAddress(librpcrt4, "NdrConformantArrayBufferSize")
	ndrConformantArrayFree = doGetProcAddress(librpcrt4, "NdrConformantArrayFree")
	ndrConformantArrayMemorySize = doGetProcAddress(librpcrt4, "NdrConformantArrayMemorySize")
	ndrConformantStringBufferSize = doGetProcAddress(librpcrt4, "NdrConformantStringBufferSize")
	ndrConformantStringMemorySize = doGetProcAddress(librpcrt4, "NdrConformantStringMemorySize")
	ndrConformantStructBufferSize = doGetProcAddress(librpcrt4, "NdrConformantStructBufferSize")
	ndrConformantStructFree = doGetProcAddress(librpcrt4, "NdrConformantStructFree")
	ndrConformantStructMemorySize = doGetProcAddress(librpcrt4, "NdrConformantStructMemorySize")
	ndrConformantVaryingArrayBufferSize = doGetProcAddress(librpcrt4, "NdrConformantVaryingArrayBufferSize")
	ndrConformantVaryingArrayFree = doGetProcAddress(librpcrt4, "NdrConformantVaryingArrayFree")
	ndrConformantVaryingArrayMemorySize = doGetProcAddress(librpcrt4, "NdrConformantVaryingArrayMemorySize")
	ndrConformantVaryingStructBufferSize = doGetProcAddress(librpcrt4, "NdrConformantVaryingStructBufferSize")
	ndrConformantVaryingStructFree = doGetProcAddress(librpcrt4, "NdrConformantVaryingStructFree")
	ndrConformantVaryingStructMemorySize = doGetProcAddress(librpcrt4, "NdrConformantVaryingStructMemorySize")
	ndrContextHandleInitialize = doGetProcAddress(librpcrt4, "NdrContextHandleInitialize")
	ndrContextHandleSize = doGetProcAddress(librpcrt4, "NdrContextHandleSize")
	ndrConvert = doGetProcAddress(librpcrt4, "NdrConvert")
	ndrConvert2 = doGetProcAddress(librpcrt4, "NdrConvert2")
	ndrCorrelationFree = doGetProcAddress(librpcrt4, "NdrCorrelationFree")
	ndrCorrelationInitialize = doGetProcAddress(librpcrt4, "NdrCorrelationInitialize")
	ndrCorrelationPass = doGetProcAddress(librpcrt4, "NdrCorrelationPass")
	ndrEncapsulatedUnionBufferSize = doGetProcAddress(librpcrt4, "NdrEncapsulatedUnionBufferSize")
	ndrEncapsulatedUnionFree = doGetProcAddress(librpcrt4, "NdrEncapsulatedUnionFree")
	ndrEncapsulatedUnionMemorySize = doGetProcAddress(librpcrt4, "NdrEncapsulatedUnionMemorySize")
	ndrFixedArrayBufferSize = doGetProcAddress(librpcrt4, "NdrFixedArrayBufferSize")
	ndrFixedArrayFree = doGetProcAddress(librpcrt4, "NdrFixedArrayFree")
	ndrFixedArrayMemorySize = doGetProcAddress(librpcrt4, "NdrFixedArrayMemorySize")
	ndrFreeBuffer = doGetProcAddress(librpcrt4, "NdrFreeBuffer")
	ndrInterfacePointerBufferSize = doGetProcAddress(librpcrt4, "NdrInterfacePointerBufferSize")
	ndrInterfacePointerFree = doGetProcAddress(librpcrt4, "NdrInterfacePointerFree")
	ndrInterfacePointerMemorySize = doGetProcAddress(librpcrt4, "NdrInterfacePointerMemorySize")
	ndrNonConformantStringBufferSize = doGetProcAddress(librpcrt4, "NdrNonConformantStringBufferSize")
	ndrNonConformantStringMemorySize = doGetProcAddress(librpcrt4, "NdrNonConformantStringMemorySize")
	ndrNonEncapsulatedUnionBufferSize = doGetProcAddress(librpcrt4, "NdrNonEncapsulatedUnionBufferSize")
	ndrNonEncapsulatedUnionFree = doGetProcAddress(librpcrt4, "NdrNonEncapsulatedUnionFree")
	ndrNonEncapsulatedUnionMemorySize = doGetProcAddress(librpcrt4, "NdrNonEncapsulatedUnionMemorySize")
	ndrOleFree = doGetProcAddress(librpcrt4, "NdrOleFree")
	ndrPointerBufferSize = doGetProcAddress(librpcrt4, "NdrPointerBufferSize")
	ndrPointerFree = doGetProcAddress(librpcrt4, "NdrPointerFree")
	ndrPointerMemorySize = doGetProcAddress(librpcrt4, "NdrPointerMemorySize")
	ndrProxyErrorHandler = doGetProcAddress(librpcrt4, "NdrProxyErrorHandler")
	ndrProxyFreeBuffer = doGetProcAddress(librpcrt4, "NdrProxyFreeBuffer")
	ndrProxyGetBuffer = doGetProcAddress(librpcrt4, "NdrProxyGetBuffer")
	ndrProxySendReceive = doGetProcAddress(librpcrt4, "NdrProxySendReceive")
	ndrRpcSmSetClientToOsf = doGetProcAddress(librpcrt4, "NdrRpcSmSetClientToOsf")
	ndrServerCall2 = doGetProcAddress(librpcrt4, "NdrServerCall2")
	ndrServerContextMarshall = doGetProcAddress(librpcrt4, "NdrServerContextMarshall")
	ndrServerContextNewMarshall = doGetProcAddress(librpcrt4, "NdrServerContextNewMarshall")
	ndrServerContextNewUnmarshall = doGetProcAddress(librpcrt4, "NdrServerContextNewUnmarshall")
	ndrServerContextUnmarshall = doGetProcAddress(librpcrt4, "NdrServerContextUnmarshall")
	ndrSimpleStructBufferSize = doGetProcAddress(librpcrt4, "NdrSimpleStructBufferSize")
	ndrSimpleStructFree = doGetProcAddress(librpcrt4, "NdrSimpleStructFree")
	ndrSimpleStructMemorySize = doGetProcAddress(librpcrt4, "NdrSimpleStructMemorySize")
	ndrUserMarshalBufferSize = doGetProcAddress(librpcrt4, "NdrUserMarshalBufferSize")
	ndrUserMarshalFree = doGetProcAddress(librpcrt4, "NdrUserMarshalFree")
	ndrUserMarshalMemorySize = doGetProcAddress(librpcrt4, "NdrUserMarshalMemorySize")
	ndrVaryingArrayBufferSize = doGetProcAddress(librpcrt4, "NdrVaryingArrayBufferSize")
	ndrVaryingArrayFree = doGetProcAddress(librpcrt4, "NdrVaryingArrayFree")
	ndrVaryingArrayMemorySize = doGetProcAddress(librpcrt4, "NdrVaryingArrayMemorySize")
	ndrXmitOrRepAsBufferSize = doGetProcAddress(librpcrt4, "NdrXmitOrRepAsBufferSize")
	ndrXmitOrRepAsFree = doGetProcAddress(librpcrt4, "NdrXmitOrRepAsFree")
	ndrXmitOrRepAsMemorySize = doGetProcAddress(librpcrt4, "NdrXmitOrRepAsMemorySize")
	rpcBindingFree = doGetProcAddress(librpcrt4, "RpcBindingFree")
	rpcBindingSetOption = doGetProcAddress(librpcrt4, "RpcBindingSetOption")
	rpcImpersonateClient = doGetProcAddress(librpcrt4, "RpcImpersonateClient")
	rpcMgmtEnableIdleCleanup = doGetProcAddress(librpcrt4, "RpcMgmtEnableIdleCleanup")
	rpcMgmtIsServerListening = doGetProcAddress(librpcrt4, "RpcMgmtIsServerListening")
	rpcMgmtSetComTimeout = doGetProcAddress(librpcrt4, "RpcMgmtSetComTimeout")
	rpcMgmtSetServerStackSize = doGetProcAddress(librpcrt4, "RpcMgmtSetServerStackSize")
	rpcMgmtStopServerListening = doGetProcAddress(librpcrt4, "RpcMgmtStopServerListening")
	rpcMgmtWaitServerListen = doGetProcAddress(librpcrt4, "RpcMgmtWaitServerListen")
	rpcRevertToSelf = doGetProcAddress(librpcrt4, "RpcRevertToSelf")
	rpcRevertToSelfEx = doGetProcAddress(librpcrt4, "RpcRevertToSelfEx")
	rpcServerListen = doGetProcAddress(librpcrt4, "RpcServerListen")
	rpcSmDestroyClientContext = doGetProcAddress(librpcrt4, "RpcSmDestroyClientContext")
	rpcSsDestroyClientContext = doGetProcAddress(librpcrt4, "RpcSsDestroyClientContext")
	rpcSsDontSerializeContext = doGetProcAddress(librpcrt4, "RpcSsDontSerializeContext")
}

func CStdStubBuffer_AddRef(this *IRpcStubBuffer) ULONG {
	ret1 := syscall3(cStdStubBuffer_AddRef, 1,
		uintptr(unsafe.Pointer(this)),
		0,
		0)
	return ULONG(ret1)
}

func CStdStubBuffer_Connect(this *IRpcStubBuffer, pUnkServer *IUnknown) HRESULT {
	ret1 := syscall3(cStdStubBuffer_Connect, 2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pUnkServer)),
		0)
	return HRESULT(ret1)
}

func CStdStubBuffer_CountRefs(this *IRpcStubBuffer) ULONG {
	ret1 := syscall3(cStdStubBuffer_CountRefs, 1,
		uintptr(unsafe.Pointer(this)),
		0,
		0)
	return ULONG(ret1)
}

func CStdStubBuffer_DebugServerQueryInterface(this *IRpcStubBuffer, ppv uintptr) HRESULT {
	ret1 := syscall3(cStdStubBuffer_DebugServerQueryInterface, 2,
		uintptr(unsafe.Pointer(this)),
		ppv,
		0)
	return HRESULT(ret1)
}

func CStdStubBuffer_DebugServerRelease(this *IRpcStubBuffer, pv uintptr) {
	syscall3(cStdStubBuffer_DebugServerRelease, 2,
		uintptr(unsafe.Pointer(this)),
		pv,
		0)
}

func CStdStubBuffer_Disconnect(this *IRpcStubBuffer) {
	syscall3(cStdStubBuffer_Disconnect, 1,
		uintptr(unsafe.Pointer(this)),
		0,
		0)
}

// TODO: Unknown type(s): RPCOLEMESSAGE *
// func CStdStubBuffer_Invoke(this *IRpcStubBuffer, pRpcMsg RPCOLEMESSAGE *, pRpcChannelBuffer *IRpcChannelBuffer) HRESULT

func CStdStubBuffer_QueryInterface(this *IRpcStubBuffer, riid REFIID, ppvObject uintptr) HRESULT {
	ret1 := syscall3(cStdStubBuffer_QueryInterface, 3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(riid)),
		ppvObject)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): IPSFactoryBuffer *
// func NdrCStdStubBuffer2_Release(this *IRpcStubBuffer, pPSF IPSFactoryBuffer *) ULONG

// TODO: Unknown type(s): IPSFactoryBuffer *
// func NdrCStdStubBuffer_Release(this *IRpcStubBuffer, pPSF IPSFactoryBuffer *) ULONG

// TODO: Unknown type(s): LPRPCSTUBBUFFER
// func CStdStubBuffer_IsIIDSupported(iface LPRPCSTUBBUFFER, riid REFIID) LPRPCSTUBBUFFER

// TODO: Unknown type(s): LPRPCPROXYBUFFER *, LPTYPEINFO
// func CreateProxyFromTypeInfo(pTypeInfo LPTYPEINFO, pUnkOuter LPUNKNOWN, riid REFIID, ppProxy LPRPCPROXYBUFFER *, ppv *LPVOID) HRESULT

func CreateStubFromTypeInfo(pTypeInfo *ITypeInfo, riid REFIID, pUnkServer *IUnknown, ppStub **IRpcStubBuffer) HRESULT {
	ret1 := syscall6(createStubFromTypeInfo, 4,
		uintptr(unsafe.Pointer(pTypeInfo)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pUnkServer)),
		uintptr(unsafe.Pointer(ppStub)),
		0,
		0)
	return HRESULT(ret1)
}

func IUnknown_AddRef_Proxy(iface LPUNKNOWN) ULONG {
	ret1 := syscall3(iUnknown_AddRef_Proxy, 1,
		uintptr(unsafe.Pointer(iface)),
		0,
		0)
	return ULONG(ret1)
}

func IUnknown_QueryInterface_Proxy(iface LPUNKNOWN, riid REFIID, ppvObj *LPVOID) HRESULT {
	ret1 := syscall3(iUnknown_QueryInterface_Proxy, 3,
		uintptr(unsafe.Pointer(iface)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObj)))
	return HRESULT(ret1)
}

func IUnknown_Release_Proxy(iface LPUNKNOWN) ULONG {
	ret1 := syscall3(iUnknown_Release_Proxy, 1,
		uintptr(unsafe.Pointer(iface)),
		0,
		0)
	return ULONG(ret1)
}

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func I_RpcAsyncAbortCall(pAsync PRPC_ASYNC_STATE, exceptionCode ULONG) RPC_STATUS

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func I_RpcAsyncSetHandle(pMsg PRPC_MESSAGE, pAsync PRPC_ASYNC_STATE) RPC_STATUS

func I_RpcBindingInqLocalClientPID(clientBinding RPC_BINDING_HANDLE, clientPID *ULONG) RPC_STATUS {
	ret1 := syscall3(i_RpcBindingInqLocalClientPID, 2,
		uintptr(clientBinding),
		uintptr(unsafe.Pointer(clientPID)),
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): unsigned int *
// func I_RpcBindingInqTransportType(binding RPC_BINDING_HANDLE, aType unsigned int *) RPC_STATUS

func I_RpcExceptionFilter(exceptionCode ULONG) int32 {
	ret1 := syscall3(i_RpcExceptionFilter, 1,
		uintptr(exceptionCode),
		0,
		0)
	return int32(ret1)
}

func I_RpcFree(object uintptr) {
	syscall3(i_RpcFree, 1,
		object,
		0,
		0)
}

func I_RpcFreeBuffer(pMsg PRPC_MESSAGE) RPC_STATUS {
	ret1 := syscall3(i_RpcFreeBuffer, 1,
		uintptr(unsafe.Pointer(pMsg)),
		0,
		0)
	return RPC_STATUS(ret1)
}

func I_RpcGetBuffer(pMsg PRPC_MESSAGE) RPC_STATUS {
	ret1 := syscall3(i_RpcGetBuffer, 1,
		uintptr(unsafe.Pointer(pMsg)),
		0,
		0)
	return RPC_STATUS(ret1)
}

func I_RpcGetCurrentCallHandle() RPC_BINDING_HANDLE {
	ret1 := syscall3(i_RpcGetCurrentCallHandle, 0,
		0,
		0,
		0)
	return RPC_BINDING_HANDLE(ret1)
}

func I_RpcMapWin32Status(status RPC_STATUS) LONG {
	ret1 := syscall3(i_RpcMapWin32Status, 1,
		uintptr(status),
		0,
		0)
	return LONG(ret1)
}

func I_RpcNegotiateTransferSyntax(pMsg PRPC_MESSAGE) RPC_STATUS {
	ret1 := syscall3(i_RpcNegotiateTransferSyntax, 1,
		uintptr(unsafe.Pointer(pMsg)),
		0,
		0)
	return RPC_STATUS(ret1)
}

func I_RpcReceive(pMsg PRPC_MESSAGE) RPC_STATUS {
	ret1 := syscall3(i_RpcReceive, 1,
		uintptr(unsafe.Pointer(pMsg)),
		0,
		0)
	return RPC_STATUS(ret1)
}

func I_RpcSend(pMsg PRPC_MESSAGE) RPC_STATUS {
	ret1 := syscall3(i_RpcSend, 1,
		uintptr(unsafe.Pointer(pMsg)),
		0,
		0)
	return RPC_STATUS(ret1)
}

func I_RpcSendReceive(pMsg PRPC_MESSAGE) RPC_STATUS {
	ret1 := syscall3(i_RpcSendReceive, 1,
		uintptr(unsafe.Pointer(pMsg)),
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): UUID *
// func I_UuidCreate(uuid UUID *) RPC_STATUS

// TODO: Unknown type(s): MIDL_ES_CODE, char * *, handle_t
// func MesBufferHandleReset(handle handle_t, handleStyle ULONG, operation MIDL_ES_CODE, buffer char * *, bufferSize ULONG, encodedSize *ULONG) RPC_STATUS

// TODO: Unknown type(s): MIDL_ES_READ, handle_t *
// func MesDecodeIncrementalHandleCreate(userState uintptr, readFn MIDL_ES_READ, pHandle handle_t *) RPC_STATUS

// TODO: Unknown type(s): MIDL_ES_ALLOC, MIDL_ES_WRITE, handle_t *
// func MesEncodeIncrementalHandleCreate(userState uintptr, allocFn MIDL_ES_ALLOC, writeFn MIDL_ES_WRITE, pHandle handle_t *) RPC_STATUS

// TODO: Unknown type(s): handle_t
// func MesHandleFree(handle handle_t) RPC_STATUS

// TODO: Unknown type(s): MIDL_ES_ALLOC, MIDL_ES_CODE, MIDL_ES_READ, MIDL_ES_WRITE, handle_t
// func MesIncrementalHandleReset(handle handle_t, userState uintptr, allocFn MIDL_ES_ALLOC, writeFn MIDL_ES_WRITE, readFn MIDL_ES_READ, operation MIDL_ES_CODE) RPC_STATUS

// TODO: Unknown type(s): NDR_CCONTEXT
// func NDRCContextBinding(cContext NDR_CCONTEXT) RPC_BINDING_HANDLE

// TODO: Unknown type(s): NDR_CCONTEXT
// func NDRCContextMarshall(cContext NDR_CCONTEXT, pBuff uintptr)

// TODO: Unknown type(s): NDR_CCONTEXT *
// func NDRCContextUnmarshall(cContext NDR_CCONTEXT *, hBinding RPC_BINDING_HANDLE, pBuff uintptr, dataRepresentation ULONG)

func NDRSContextMarshall(sContext NDR_SCONTEXT, pBuff uintptr, userRunDownIn NDR_RUNDOWN) {
	userRunDownInCallback := syscall.NewCallback(userRunDownIn)
	syscall3(nDRSContextMarshall, 3,
		uintptr(unsafe.Pointer(sContext)),
		pBuff,
		userRunDownInCallback)
}

func NDRSContextMarshall2(hBinding RPC_BINDING_HANDLE, sContext NDR_SCONTEXT, pBuff uintptr, userRunDownIn NDR_RUNDOWN, ctxGuard uintptr, flags ULONG) {
	userRunDownInCallback := syscall.NewCallback(userRunDownIn)
	syscall6(nDRSContextMarshall2, 6,
		uintptr(hBinding),
		uintptr(unsafe.Pointer(sContext)),
		pBuff,
		userRunDownInCallback,
		ctxGuard,
		uintptr(flags))
}

func NDRSContextMarshallEx(hBinding RPC_BINDING_HANDLE, sContext NDR_SCONTEXT, pBuff uintptr, userRunDownIn NDR_RUNDOWN) {
	userRunDownInCallback := syscall.NewCallback(userRunDownIn)
	syscall6(nDRSContextMarshallEx, 4,
		uintptr(hBinding),
		uintptr(unsafe.Pointer(sContext)),
		pBuff,
		userRunDownInCallback,
		0,
		0)
}

func NDRSContextUnmarshall(pBuff uintptr, dataRepresentation ULONG) NDR_SCONTEXT {
	ret1 := syscall3(nDRSContextUnmarshall, 2,
		pBuff,
		uintptr(dataRepresentation),
		0)
	return (NDR_SCONTEXT)(unsafe.Pointer(ret1))
}

func NDRSContextUnmarshall2(hBinding RPC_BINDING_HANDLE, pBuff uintptr, dataRepresentation ULONG, ctxGuard uintptr, flags ULONG) NDR_SCONTEXT {
	ret1 := syscall6(nDRSContextUnmarshall2, 5,
		uintptr(hBinding),
		pBuff,
		uintptr(dataRepresentation),
		ctxGuard,
		uintptr(flags),
		0)
	return (NDR_SCONTEXT)(unsafe.Pointer(ret1))
}

func NDRSContextUnmarshallEx(hBinding RPC_BINDING_HANDLE, pBuff uintptr, dataRepresentation ULONG) NDR_SCONTEXT {
	ret1 := syscall3(nDRSContextUnmarshallEx, 3,
		uintptr(hBinding),
		pBuff,
		uintptr(dataRepresentation))
	return (NDR_SCONTEXT)(unsafe.Pointer(ret1))
}

func NdrByteCountPointerBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrByteCountPointerBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrByteCountPointerFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrByteCountPointerFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrClearOutParameters(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING, argAddr uintptr) {
	syscall3(ndrClearOutParameters, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		argAddr)
}

// TODO: Unknown type(s): NDR_CCONTEXT
// func NdrClientContextMarshall(pStubMsg PMIDL_STUB_MESSAGE, contextHandle NDR_CCONTEXT, fCheck int32)

// TODO: Unknown type(s): NDR_CCONTEXT *
// func NdrClientContextUnmarshall(pStubMsg PMIDL_STUB_MESSAGE, pContextHandle NDR_CCONTEXT *, bindHandle RPC_BINDING_HANDLE)

// TODO: Unknown type(s): PMIDL_STUB_DESC
// func NdrClientInitializeNew(pRpcMessage PRPC_MESSAGE, pStubMsg PMIDL_STUB_MESSAGE, pStubDesc PMIDL_STUB_DESC, procNum uint32)

func NdrComplexArrayBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrComplexArrayBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrComplexArrayFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrComplexArrayFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrComplexArrayMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrComplexArrayMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrComplexStructBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrComplexStructBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrComplexStructFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrComplexStructFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrComplexStructMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrComplexStructMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrConformantArrayBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantArrayBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantArrayFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantArrayFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantArrayMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrConformantArrayMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrConformantStringBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantStringBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantStringMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrConformantStringMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrConformantStructBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantStructBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantStructFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantStructFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantStructMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrConformantStructMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrConformantVaryingArrayBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantVaryingArrayBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantVaryingArrayFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantVaryingArrayFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantVaryingArrayMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrConformantVaryingArrayMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrConformantVaryingStructBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantVaryingStructBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantVaryingStructFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConformantVaryingStructFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConformantVaryingStructMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrConformantVaryingStructMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrContextHandleInitialize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) NDR_SCONTEXT {
	ret1 := syscall3(ndrContextHandleInitialize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return (NDR_SCONTEXT)(unsafe.Pointer(ret1))
}

func NdrContextHandleSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrContextHandleSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrConvert(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrConvert, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
}

func NdrConvert2(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING, numberParams LONG) {
	syscall3(ndrConvert2, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		uintptr(numberParams))
}

func NdrCorrelationFree(pStubMsg PMIDL_STUB_MESSAGE) {
	syscall3(ndrCorrelationFree, 1,
		uintptr(unsafe.Pointer(pStubMsg)),
		0,
		0)
}

func NdrCorrelationInitialize(pStubMsg PMIDL_STUB_MESSAGE, pMemory uintptr, cacheSize ULONG, flags ULONG) {
	syscall6(ndrCorrelationInitialize, 4,
		uintptr(unsafe.Pointer(pStubMsg)),
		pMemory,
		uintptr(cacheSize),
		uintptr(flags),
		0,
		0)
}

func NdrCorrelationPass(pStubMsg PMIDL_STUB_MESSAGE) {
	syscall3(ndrCorrelationPass, 1,
		uintptr(unsafe.Pointer(pStubMsg)),
		0,
		0)
}

// TODO: Unknown type(s): CStdPSFactoryBuffer *
// func NdrDllCanUnloadNow(pPSFactoryBuffer CStdPSFactoryBuffer *) HRESULT

// TODO: Unknown type(s): CStdPSFactoryBuffer *, const ProxyFileInfo * *
// func NdrDllGetClassObject(rclsid /*const*/ REFCLSID, iid REFIID, ppv *LPVOID, pProxyFileList /*const*/ const ProxyFileInfo * *, pclsid /*const*/ *CLSID, pPSFactoryBuffer CStdPSFactoryBuffer *) HRESULT

// TODO: Unknown type(s): const ProxyFileInfo * *
// func NdrDllRegisterProxy(hDll HMODULE, pProxyFileList /*const*/ const ProxyFileInfo * *, pclsid /*const*/ *CLSID) HRESULT

// TODO: Unknown type(s): const ProxyFileInfo * *
// func NdrDllUnregisterProxy(hDll HMODULE, pProxyFileList /*const*/ const ProxyFileInfo * *, pclsid /*const*/ *CLSID) HRESULT

func NdrEncapsulatedUnionBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrEncapsulatedUnionBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrEncapsulatedUnionFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrEncapsulatedUnionFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrEncapsulatedUnionMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrEncapsulatedUnionMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrFixedArrayBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrFixedArrayBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrFixedArrayFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrFixedArrayFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrFixedArrayMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrFixedArrayMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrFreeBuffer(pStubMsg PMIDL_STUB_MESSAGE) {
	syscall3(ndrFreeBuffer, 1,
		uintptr(unsafe.Pointer(pStubMsg)),
		0,
		0)
}

// TODO: Unknown type(s): PFULL_PTR_XLAT_TABLES
// func NdrFullPointerFree(pXlatTables PFULL_PTR_XLAT_TABLES, pointer uintptr) int32

// TODO: Unknown type(s): PFULL_PTR_XLAT_TABLES
// func NdrFullPointerInsertRefId(pXlatTables PFULL_PTR_XLAT_TABLES, refId ULONG, pPointer uintptr)

// TODO: Unknown type(s): PFULL_PTR_XLAT_TABLES, unsigned char
// func NdrFullPointerQueryPointer(pXlatTables PFULL_PTR_XLAT_TABLES, pPointer uintptr, queryType unsigned char, pRefId *ULONG) int32

// TODO: Unknown type(s): PFULL_PTR_XLAT_TABLES, unsigned char
// func NdrFullPointerQueryRefId(pXlatTables PFULL_PTR_XLAT_TABLES, refId ULONG, queryType unsigned char, ppPointer uintptr) int32

// TODO: Unknown type(s): PFULL_PTR_XLAT_TABLES
// func NdrFullPointerXlatFree(pXlatTables PFULL_PTR_XLAT_TABLES)

// TODO: Unknown type(s): PFULL_PTR_XLAT_TABLES
// func NdrFullPointerXlatInit(numberOfPointers ULONG, xlatSide XLAT_SIDE) PFULL_PTR_XLAT_TABLES

func NdrInterfacePointerBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrInterfacePointerBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrInterfacePointerFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrInterfacePointerFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrInterfacePointerMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrInterfacePointerMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrNonConformantStringBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrNonConformantStringBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrNonConformantStringMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrNonConformantStringMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrNonEncapsulatedUnionBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrNonEncapsulatedUnionBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrNonEncapsulatedUnionFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrNonEncapsulatedUnionFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrNonEncapsulatedUnionMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrNonEncapsulatedUnionMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrOleFree(nodeToFree uintptr) {
	syscall3(ndrOleFree, 1,
		nodeToFree,
		0,
		0)
}

func NdrPointerBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrPointerBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrPointerFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrPointerFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrPointerMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrPointerMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrProxyErrorHandler(dwExceptionCode uint32) HRESULT {
	ret1 := syscall3(ndrProxyErrorHandler, 1,
		uintptr(dwExceptionCode),
		0,
		0)
	return HRESULT(ret1)
}

func NdrProxyFreeBuffer(this uintptr, pStubMsg PMIDL_STUB_MESSAGE) {
	syscall3(ndrProxyFreeBuffer, 2,
		this,
		uintptr(unsafe.Pointer(pStubMsg)),
		0)
}

func NdrProxyGetBuffer(this uintptr, pStubMsg PMIDL_STUB_MESSAGE) {
	syscall3(ndrProxyGetBuffer, 2,
		this,
		uintptr(unsafe.Pointer(pStubMsg)),
		0)
}

// TODO: Unknown type(s): PMIDL_STUB_DESC
// func NdrProxyInitialize(this uintptr, pRpcMsg PRPC_MESSAGE, pStubMsg PMIDL_STUB_MESSAGE, pStubDescriptor PMIDL_STUB_DESC, procNum uint32)

func NdrProxySendReceive(this uintptr, pStubMsg PMIDL_STUB_MESSAGE) {
	syscall3(ndrProxySendReceive, 2,
		this,
		uintptr(unsafe.Pointer(pStubMsg)),
		0)
}

func NdrRpcSmSetClientToOsf(pMessage PMIDL_STUB_MESSAGE) {
	syscall3(ndrRpcSmSetClientToOsf, 1,
		uintptr(unsafe.Pointer(pMessage)),
		0,
		0)
}

func NdrServerCall2(pRpcMsg PRPC_MESSAGE) {
	syscall3(ndrServerCall2, 1,
		uintptr(unsafe.Pointer(pRpcMsg)),
		0,
		0)
}

func NdrServerContextMarshall(pStubMsg PMIDL_STUB_MESSAGE, contextHandle NDR_SCONTEXT, rundownRoutine NDR_RUNDOWN) {
	rundownRoutineCallback := syscall.NewCallback(rundownRoutine)
	syscall3(ndrServerContextMarshall, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(contextHandle)),
		rundownRoutineCallback)
}

func NdrServerContextNewMarshall(pStubMsg PMIDL_STUB_MESSAGE, contextHandle NDR_SCONTEXT, rundownRoutine NDR_RUNDOWN, pFormat /*const*/ PFORMAT_STRING) {
	rundownRoutineCallback := syscall.NewCallback(rundownRoutine)
	syscall6(ndrServerContextNewMarshall, 4,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(contextHandle)),
		rundownRoutineCallback,
		uintptr(unsafe.Pointer(pFormat)),
		0,
		0)
}

func NdrServerContextNewUnmarshall(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) NDR_SCONTEXT {
	ret1 := syscall3(ndrServerContextNewUnmarshall, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return (NDR_SCONTEXT)(unsafe.Pointer(ret1))
}

func NdrServerContextUnmarshall(pStubMsg PMIDL_STUB_MESSAGE) NDR_SCONTEXT {
	ret1 := syscall3(ndrServerContextUnmarshall, 1,
		uintptr(unsafe.Pointer(pStubMsg)),
		0,
		0)
	return (NDR_SCONTEXT)(unsafe.Pointer(ret1))
}

func NdrSimpleStructBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrSimpleStructBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrSimpleStructFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrSimpleStructFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrSimpleStructMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrSimpleStructMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

// TODO: Unknown type(s): unsigned char
// func NdrSimpleTypeMarshall(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, formatChar unsigned char)

// TODO: Unknown type(s): unsigned char
// func NdrSimpleTypeUnmarshall(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, formatChar unsigned char)

// TODO: Unknown type(s): struct IRpcChannelBuffer *, struct IRpcStubBuffer *
// func NdrStubCall2(pThis struct IRpcStubBuffer *, pChannel struct IRpcChannelBuffer *, pRpcMsg PRPC_MESSAGE, pdwStubPhase *uint32) LONG

// TODO: Unknown type(s): LPRPCCHANNELBUFFER, LPRPCSTUBBUFFER
// func NdrStubGetBuffer(iface LPRPCSTUBBUFFER, pRpcChannelBuffer LPRPCCHANNELBUFFER, pStubMsg PMIDL_STUB_MESSAGE)

// TODO: Unknown type(s): LPRPCCHANNELBUFFER, PMIDL_STUB_DESC
// func NdrStubInitialize(pRpcMsg PRPC_MESSAGE, pStubMsg PMIDL_STUB_MESSAGE, pStubDescriptor PMIDL_STUB_DESC, pRpcChannelBuffer LPRPCCHANNELBUFFER)

func NdrUserMarshalBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrUserMarshalBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrUserMarshalFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrUserMarshalFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrUserMarshalMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrUserMarshalMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrVaryingArrayBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrVaryingArrayBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrVaryingArrayFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrVaryingArrayFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrVaryingArrayMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrVaryingArrayMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

func NdrXmitOrRepAsBufferSize(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrXmitOrRepAsBufferSize, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrXmitOrRepAsFree(pStubMsg PMIDL_STUB_MESSAGE, pMemory *byte, pFormat /*const*/ PFORMAT_STRING) {
	syscall3(ndrXmitOrRepAsFree, 3,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pMemory)),
		uintptr(unsafe.Pointer(pFormat)))
}

func NdrXmitOrRepAsMemorySize(pStubMsg PMIDL_STUB_MESSAGE, pFormat /*const*/ PFORMAT_STRING) ULONG {
	ret1 := syscall3(ndrXmitOrRepAsMemorySize, 2,
		uintptr(unsafe.Pointer(pStubMsg)),
		uintptr(unsafe.Pointer(pFormat)),
		0)
	return ULONG(ret1)
}

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func RpcAsyncAbortCall(pAsync PRPC_ASYNC_STATE, exceptionCode ULONG) RPC_STATUS

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func RpcAsyncCancelCall(pAsync PRPC_ASYNC_STATE, fAbortCall bool) RPC_STATUS

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func RpcAsyncCompleteCall(pAsync PRPC_ASYNC_STATE, reply uintptr) RPC_STATUS

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func RpcAsyncGetCallStatus(pAsync PRPC_ASYNC_STATE) RPC_STATUS

// TODO: Unknown type(s): PRPC_ASYNC_STATE
// func RpcAsyncInitializeHandle(pAsync PRPC_ASYNC_STATE, size uint32) RPC_STATUS

func RpcBindingFree(binding *RPC_BINDING_HANDLE) RPC_STATUS {
	ret1 := syscall3(rpcBindingFree, 1,
		uintptr(unsafe.Pointer(binding)),
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_WSTR
// func RpcBindingFromStringBinding(stringBinding RPC_WSTR, binding *RPC_BINDING_HANDLE) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func RpcBindingInqObject(binding RPC_BINDING_HANDLE, objectUuid UUID *) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func RpcBindingSetObject(binding RPC_BINDING_HANDLE, objectUuid UUID *) RPC_STATUS

func RpcBindingSetOption(bindingHandle RPC_BINDING_HANDLE, option ULONG, optionValue *uint32) RPC_STATUS {
	ret1 := syscall3(rpcBindingSetOption, 3,
		uintptr(bindingHandle),
		uintptr(option),
		uintptr(unsafe.Pointer(optionValue)))
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_WSTR *
// func RpcBindingToStringBinding(binding RPC_BINDING_HANDLE, stringBinding RPC_WSTR *) RPC_STATUS

// TODO: Unknown type(s): RPC_BINDING_VECTOR * *
// func RpcBindingVectorFree(bindingVector RPC_BINDING_VECTOR * *) RPC_STATUS

// TODO: Unknown type(s): RPC_BINDING_VECTOR *, RPC_IF_HANDLE, RPC_WSTR, UUID_VECTOR *
// func RpcEpRegisterNoReplace(ifSpec RPC_IF_HANDLE, bindingVector RPC_BINDING_VECTOR *, uuidVector UUID_VECTOR *, annotation RPC_WSTR) RPC_STATUS

// TODO: Unknown type(s): RPC_BINDING_VECTOR *, RPC_IF_HANDLE, RPC_WSTR, UUID_VECTOR *
// func RpcEpRegister(ifSpec RPC_IF_HANDLE, bindingVector RPC_BINDING_VECTOR *, uuidVector UUID_VECTOR *, annotation RPC_WSTR) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_HANDLE
// func RpcEpResolveBinding(binding RPC_BINDING_HANDLE, ifSpec RPC_IF_HANDLE) RPC_STATUS

// TODO: Unknown type(s): RPC_BINDING_VECTOR *, RPC_IF_HANDLE, UUID_VECTOR *
// func RpcEpUnregister(ifSpec RPC_IF_HANDLE, bindingVector RPC_BINDING_VECTOR *, uuidVector UUID_VECTOR *) RPC_STATUS

func RpcImpersonateClient(bindingHandle RPC_BINDING_HANDLE) RPC_STATUS {
	ret1 := syscall3(rpcImpersonateClient, 1,
		uintptr(bindingHandle),
		0,
		0)
	return RPC_STATUS(ret1)
}

func RpcMgmtEnableIdleCleanup() RPC_STATUS {
	ret1 := syscall3(rpcMgmtEnableIdleCleanup, 0,
		0,
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_EP_INQ_HANDLE *, RPC_IF_ID *, UUID *
// func RpcMgmtEpEltInqBegin(binding RPC_BINDING_HANDLE, inquiryType ULONG, ifId RPC_IF_ID *, versOption ULONG, objectUuid UUID *, inquiryContext RPC_EP_INQ_HANDLE *) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_ID_VECTOR * *
// func RpcMgmtInqIfIds(binding RPC_BINDING_HANDLE, ifIdVector RPC_IF_ID_VECTOR * *) RPC_STATUS

// TODO: Unknown type(s): RPC_STATS_VECTOR * *
// func RpcMgmtInqStats(binding RPC_BINDING_HANDLE, statistics RPC_STATS_VECTOR * *) RPC_STATUS

func RpcMgmtIsServerListening(binding RPC_BINDING_HANDLE) RPC_STATUS {
	ret1 := syscall3(rpcMgmtIsServerListening, 1,
		uintptr(binding),
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_MGMT_AUTHORIZATION_FN
// func RpcMgmtSetAuthorizationFn(fn RPC_MGMT_AUTHORIZATION_FN) RPC_STATUS

func RpcMgmtSetComTimeout(bindingHandle RPC_BINDING_HANDLE, timeout uint32) RPC_STATUS {
	ret1 := syscall3(rpcMgmtSetComTimeout, 2,
		uintptr(bindingHandle),
		uintptr(timeout),
		0)
	return RPC_STATUS(ret1)
}

func RpcMgmtSetServerStackSize(threadStackSize ULONG) RPC_STATUS {
	ret1 := syscall3(rpcMgmtSetServerStackSize, 1,
		uintptr(threadStackSize),
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_STATS_VECTOR * *
// func RpcMgmtStatsVectorFree(statsVector RPC_STATS_VECTOR * *) RPC_STATUS

func RpcMgmtStopServerListening(binding RPC_BINDING_HANDLE) RPC_STATUS {
	ret1 := syscall3(rpcMgmtStopServerListening, 1,
		uintptr(binding),
		0,
		0)
	return RPC_STATUS(ret1)
}

func RpcMgmtWaitServerListen() RPC_STATUS {
	ret1 := syscall3(rpcMgmtWaitServerListen, 0,
		0,
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_PROTSEQ_VECTORW * *
// func RpcNetworkInqProtseqs(protseqs RPC_PROTSEQ_VECTORW * *) RPC_STATUS

// TODO: Unknown type(s): RPC_WSTR
// func RpcNetworkIsProtseqValid(protseq RPC_WSTR) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func RpcObjectSetType(objUuid UUID *, typeUuid UUID *) RPC_STATUS

// TODO: Unknown type(s): RPC_PROTSEQ_VECTORW * *
// func RpcProtseqVectorFree(protseqs RPC_PROTSEQ_VECTORW * *) RPC_STATUS

// TODO: Unknown type(s): DECLSPEC_NORETURN
// func RpcRaiseException(exception RPC_STATUS) DECLSPEC_NORETURN

func RpcRevertToSelf() RPC_STATUS {
	ret1 := syscall3(rpcRevertToSelf, 0,
		0,
		0,
		0)
	return RPC_STATUS(ret1)
}

func RpcRevertToSelfEx(bindingHandle RPC_BINDING_HANDLE) RPC_STATUS {
	ret1 := syscall3(rpcRevertToSelfEx, 1,
		uintptr(bindingHandle),
		0,
		0)
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_BINDING_VECTOR * *
// func RpcServerInqBindings(bindingVector RPC_BINDING_VECTOR * *) RPC_STATUS

func RpcServerListen(minimumCallThreads uint32, maxCalls uint32, dontWait uint32) RPC_STATUS {
	ret1 := syscall3(rpcServerListen, 3,
		uintptr(minimumCallThreads),
		uintptr(maxCalls),
		uintptr(dontWait))
	return RPC_STATUS(ret1)
}

// TODO: Unknown type(s): RPC_AUTH_KEY_RETRIEVAL_FN, RPC_WSTR
// func RpcServerRegisterAuthInfo(serverPrincName RPC_WSTR, authnSvc ULONG, getKeyFn RPC_AUTH_KEY_RETRIEVAL_FN, arg LPVOID) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_HANDLE, RPC_MGR_EPV *, UUID *
// func RpcServerRegisterIf(ifSpec RPC_IF_HANDLE, mgrTypeUuid UUID *, mgrEpv RPC_MGR_EPV *) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_CALLBACK_FN *, RPC_IF_HANDLE, RPC_MGR_EPV *, UUID *
// func RpcServerRegisterIf2(ifSpec RPC_IF_HANDLE, mgrTypeUuid UUID *, mgrEpv RPC_MGR_EPV *, flags uint32, maxCalls uint32, maxRpcSize uint32, ifCallbackFn RPC_IF_CALLBACK_FN *) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_CALLBACK_FN *, RPC_IF_HANDLE, RPC_MGR_EPV *, UUID *
// func RpcServerRegisterIfEx(ifSpec RPC_IF_HANDLE, mgrTypeUuid UUID *, mgrEpv RPC_MGR_EPV *, flags uint32, maxCalls uint32, ifCallbackFn RPC_IF_CALLBACK_FN *) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_HANDLE, UUID *
// func RpcServerUnregisterIf(ifSpec RPC_IF_HANDLE, mgrTypeUuid UUID *, waitForCallsToComplete uint32) RPC_STATUS

// TODO: Unknown type(s): RPC_IF_HANDLE, UUID *
// func RpcServerUnregisterIfEx(ifSpec RPC_IF_HANDLE, mgrTypeUuid UUID *, rundownContextHandles int32) RPC_STATUS

// TODO: Unknown type(s): PRPC_POLICY, RPC_WSTR
// func RpcServerUseProtseqEpEx(protseq RPC_WSTR, maxCalls uint32, endpoint RPC_WSTR, securityDescriptor LPVOID, lpPolicy PRPC_POLICY) RPC_STATUS

// TODO: Unknown type(s): RPC_WSTR
// func RpcServerUseProtseqEp(protseq RPC_WSTR, maxCalls uint32, endpoint RPC_WSTR, securityDescriptor LPVOID) RPC_STATUS

// TODO: Unknown type(s): RPC_WSTR
// func RpcServerUseProtseq(protseq RPC_WSTR, maxCalls uint32, securityDescriptor uintptr) RPC_STATUS

func RpcSmDestroyClientContext(contextHandle uintptr) RPC_STATUS {
	ret1 := syscall3(rpcSmDestroyClientContext, 1,
		contextHandle,
		0,
		0)
	return RPC_STATUS(ret1)
}

func RpcSsDestroyClientContext(contextHandle uintptr) {
	syscall3(rpcSsDestroyClientContext, 1,
		contextHandle,
		0,
		0)
}

func RpcSsDontSerializeContext() {
	syscall3(rpcSsDontSerializeContext, 0,
		0,
		0,
		0)
}

// TODO: Unknown type(s): RPC_WSTR, RPC_WSTR *
// func RpcStringBindingCompose(objUuid RPC_WSTR, protseq RPC_WSTR, networkAddr RPC_WSTR, endpoint RPC_WSTR, options RPC_WSTR, stringBinding RPC_WSTR *) RPC_STATUS

// TODO: Unknown type(s): RPC_WSTR, RPC_WSTR *
// func RpcStringBindingParse(stringBinding RPC_WSTR, objUuid RPC_WSTR *, protseq RPC_WSTR *, networkAddr RPC_WSTR *, endpoint RPC_WSTR *, options RPC_WSTR *) RPC_STATUS

// TODO: Unknown type(s): RPC_WSTR *
// func RpcStringFree(string RPC_WSTR *) RPC_STATUS

// TODO: Unknown type(s): twr_t * *
// func TowerConstruct(object /*const*/ *RPC_SYNTAX_IDENTIFIER, syntax /*const*/ *RPC_SYNTAX_IDENTIFIER, protseq /*const*/ *CHAR, endpoint /*const*/ *CHAR, address /*const*/ *CHAR, tower twr_t * *) RPC_STATUS

// TODO: Unknown type(s): char * *, const twr_t *
// func TowerExplode(tower /*const*/ const twr_t *, object PRPC_SYNTAX_IDENTIFIER, syntax PRPC_SYNTAX_IDENTIFIER, protseq char * *, endpoint char * *, address char * *) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func UuidCompare(uuid1 UUID *, uuid2 UUID *, status *RPC_STATUS) int32

// TODO: Unknown type(s): UUID *
// func UuidCreate(uuid UUID *) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func UuidCreateNil(uuid UUID *) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func UuidCreateSequential(uuid UUID *) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func UuidEqual(uuid1 UUID *, uuid2 UUID *, status *RPC_STATUS) int32

// TODO: Unknown type(s): RPC_WSTR, UUID *
// func UuidFromString(s RPC_WSTR, uuid UUID *) RPC_STATUS

// TODO: Unknown type(s): UUID *
// func UuidHash(uuid UUID *, status *RPC_STATUS) int16

// TODO: Unknown type(s): UUID *
// func UuidIsNil(uuid UUID *, status *RPC_STATUS) int32

// TODO: Unknown type(s): RPC_WSTR *, UUID *
// func UuidToString(uuid UUID *, stringUuid RPC_WSTR *) RPC_STATUS
