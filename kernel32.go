// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libkernel32 uintptr

	// Functions
	acquireSRWLockExclusive               uintptr
	acquireSRWLockShared                  uintptr
	activateActCtx                        uintptr
	addAtom                               uintptr
	addConsoleAlias                       uintptr
	addIntegrityLabelToBoundaryDescriptor uintptr
	addRefActCtx                          uintptr
	addSIDToBoundaryDescriptor            uintptr
	addSecureMemoryCacheCallback          uintptr
	allocConsole                          uintptr
	applicationRecoveryFinished           uintptr
	applicationRecoveryInProgress         uintptr
	areFileApisANSI                       uintptr
	assignProcessToJobObject              uintptr
	attachConsole                         uintptr
	backupRead                            uintptr
	backupSeek                            uintptr
	backupWrite                           uintptr
	beep                                  uintptr
	beginUpdateResource                   uintptr
	callNamedPipe                         uintptr
	cancelDeviceWakeupRequest             uintptr
	cancelIo                              uintptr
	cancelIoEx                            uintptr
	cancelSynchronousIo                   uintptr
	cancelTimerQueueTimer                 uintptr
	cancelWaitableTimer                   uintptr
	changeTimerQueueTimer                 uintptr
	checkNameLegalDOS8Dot3                uintptr
	checkRemoteDebuggerPresent            uintptr
	clearCommBreak                        uintptr
	closeHandle                           uintptr
	closePrivateNamespace                 uintptr
	compareFileTime                       uintptr
	connectNamedPipe                      uintptr
	continueDebugEvent                    uintptr
	convertDefaultLocale                  uintptr
	convertFiberToThread                  uintptr
	convertThreadToFiber                  uintptr
	convertThreadToFiberEx                uintptr
	copyFile                              uintptr
	copyLZFile                            uintptr
	createBoundaryDescriptor              uintptr
	createConsoleScreenBuffer             uintptr
	createDirectoryEx                     uintptr
	createDirectoryTransacted             uintptr
	createDirectory                       uintptr
	createEventEx                         uintptr
	createEvent                           uintptr
	createFileMappingNuma                 uintptr
	createFileMapping                     uintptr
	createFileTransacted                  uintptr
	createFile                            uintptr
	createHardLinkTransacted              uintptr
	createHardLink                        uintptr
	createIoCompletionPort                uintptr
	createJobObject                       uintptr
	createMailslot                        uintptr
	createMutexEx                         uintptr
	createMutex                           uintptr
	createNamedPipe                       uintptr
	createPipe                            uintptr
	createPrivateNamespace                uintptr
	createProcess                         uintptr
	createRemoteThread                    uintptr
	createSemaphoreEx                     uintptr
	createSemaphore                       uintptr
	createSymbolicLinkTransacted          uintptr
	createSymbolicLink                    uintptr
	createTapePartition                   uintptr
	createThread                          uintptr
	createTimerQueue                      uintptr
	createToolhelp32Snapshot              uintptr
	createWaitableTimerEx                 uintptr
	createWaitableTimer                   uintptr
	deactivateActCtx                      uintptr
	debugActiveProcess                    uintptr
	debugActiveProcessStop                uintptr
	debugBreak                            uintptr
	debugBreakProcess                     uintptr
	debugSetProcessKillOnExit             uintptr
	decodePointer                         uintptr
	decodeSystemPointer                   uintptr
	defineDosDevice                       uintptr
	deleteAtom                            uintptr
	deleteBoundaryDescriptor              uintptr
	deleteFiber                           uintptr
	deleteFileTransacted                  uintptr
	deleteFile                            uintptr
	deleteTimerQueue                      uintptr
	deleteTimerQueueEx                    uintptr
	deleteTimerQueueTimer                 uintptr
	deleteVolumeMountPoint                uintptr
	deviceIoControl                       uintptr
	disableThreadLibraryCalls             uintptr
	disableThreadProfiling                uintptr
	discardVirtualMemory                  uintptr
	disconnectNamedPipe                   uintptr
	dnsHostnameToComputerName             uintptr
	dosDateTimeToFileTime                 uintptr
	duplicateHandle                       uintptr
	encodePointer                         uintptr
	encodeSystemPointer                   uintptr
	endUpdateResource                     uintptr
	enumResourceLanguagesEx               uintptr
	enumResourceLanguages                 uintptr
	enumSystemFirmwareTables              uintptr
	eraseTape                             uintptr
	escapeCommFunction                    uintptr
	exitProcess                           uintptr
	exitThread                            uintptr
	expandEnvironmentStrings              uintptr
	fatalAppExit                          uintptr
	fatalExit                             uintptr
	fileTimeToDosDateTime                 uintptr
	fileTimeToLocalFileTime               uintptr
	fileTimeToSystemTime                  uintptr
	fillConsoleOutputAttribute            uintptr
	fillConsoleOutputCharacter            uintptr
	findAtom                              uintptr
	findClose                             uintptr
	findCloseChangeNotification           uintptr
	findFirstChangeNotification           uintptr
	findFirstFileNameTransactedW          uintptr
	findFirstFileNameW                    uintptr
	findFirstVolumeMountPoint             uintptr
	findFirstVolume                       uintptr
	findNLSString                         uintptr
	findNextChangeNotification            uintptr
	findNextFileNameW                     uintptr
	findNextStreamW                       uintptr
	findNextVolumeMountPoint              uintptr
	findNextVolume                        uintptr
	findResourceEx                        uintptr
	findResource                          uintptr
	findStringOrdinal                     uintptr
	findVolumeClose                       uintptr
	findVolumeMountPointClose             uintptr
	flsFree                               uintptr
	flsGetValue                           uintptr
	flsSetValue                           uintptr
	flushConsoleInputBuffer               uintptr
	flushFileBuffers                      uintptr
	flushInstructionCache                 uintptr
	flushProcessWriteBuffers              uintptr
	flushViewOfFile                       uintptr
	freeConsole                           uintptr
	freeLibrary                           uintptr
	freeLibraryAndExitThread              uintptr
	freeResource                          uintptr
	generateConsoleCtrlEvent              uintptr
	getACP                                uintptr
	getActiveProcessorCount               uintptr
	getActiveProcessorGroupCount          uintptr
	getApplicationRestartSettings         uintptr
	getAtomName                           uintptr
	getBinaryType                         uintptr
	getCommMask                           uintptr
	getCommModemStatus                    uintptr
	getCommandLine                        uintptr
	getCompressedFileSizeTransacted       uintptr
	getCompressedFileSize                 uintptr
	getComputerName                       uintptr
	getConsoleAliasExesLength             uintptr
	getConsoleAliasExes                   uintptr
	getConsoleAlias                       uintptr
	getConsoleAliasesLength               uintptr
	getConsoleAliases                     uintptr
	getConsoleCP                          uintptr
	getConsoleDisplayMode                 uintptr
	getConsoleFontSize                    uintptr
	getConsoleMode                        uintptr
	getConsoleOriginalTitle               uintptr
	getConsoleOutputCP                    uintptr
	getConsoleProcessList                 uintptr
	getConsoleTitle                       uintptr
	getConsoleWindow                      uintptr
	getCurrentActCtx                      uintptr
	getCurrentDirectory                   uintptr
	getCurrentProcess                     uintptr
	getCurrentProcessId                   uintptr
	getCurrentProcessorNumber             uintptr
	getCurrentThread                      uintptr
	getCurrentThreadId                    uintptr
	getDateFormatEx                       uintptr
	getDateFormat                         uintptr
	getDevicePowerState                   uintptr
	getDiskFreeSpace                      uintptr
	getDllDirectory                       uintptr
	getDriveType                          uintptr
	getDurationFormat                     uintptr
	getDurationFormatEx                   uintptr
	getEnvironmentVariable                uintptr
	getErrorMode                          uintptr
	getExitCodeProcess                    uintptr
	getExitCodeThread                     uintptr
	getExpandedName                       uintptr
	getFileAttributes                     uintptr
	getFileBandwidthReservation           uintptr
	getFileSize                           uintptr
	getFileTime                           uintptr
	getFileType                           uintptr
	getFinalPathNameByHandle              uintptr
	getFirmwareEnvironmentVariable        uintptr
	getFullPathNameTransacted             uintptr
	getFullPathName                       uintptr
	getHandleInformation                  uintptr
	getLargePageMinimum                   uintptr
	getLargestConsoleWindowSize           uintptr
	getLastError                          uintptr
	getLocalTime                          uintptr
	getLocaleInfoEx                       uintptr
	getLocaleInfo                         uintptr
	getLogicalDriveStrings                uintptr
	getLogicalDrives                      uintptr
	getLongPathNameTransacted             uintptr
	getLongPathName                       uintptr
	getMailslotInfo                       uintptr
	getMaximumProcessorCount              uintptr
	getMaximumProcessorGroupCount         uintptr
	getModuleFileName                     uintptr
	getModuleHandleEx                     uintptr
	getModuleHandle                       uintptr
	getNamedPipeClientComputerName        uintptr
	getNamedPipeClientProcessId           uintptr
	getNamedPipeClientSessionId           uintptr
	getNamedPipeHandleState               uintptr
	getNamedPipeInfo                      uintptr
	getNamedPipeServerProcessId           uintptr
	getNamedPipeServerSessionId           uintptr
	getNativeSystemInfo                   uintptr
	getNumaHighestNodeNumber              uintptr
	getNumaNodeNumberFromHandle           uintptr
	getNumaProximityNodeEx                uintptr
	getNumberOfConsoleInputEvents         uintptr
	getNumberOfConsoleMouseButtons        uintptr
	getOEMCP                              uintptr
	getOverlappedResult                   uintptr
	getPriorityClass                      uintptr
	getPrivateProfileInt                  uintptr
	getPrivateProfileSectionNames         uintptr
	getPrivateProfileSection              uintptr
	getPrivateProfileString               uintptr
	getPrivateProfileStruct               uintptr
	getProcAddress                        uintptr
	getProcessAffinityMask                uintptr
	getProcessDEPPolicy                   uintptr
	getProcessGroupAffinity               uintptr
	getProcessHandleCount                 uintptr
	getProcessHeap                        uintptr
	getProcessHeaps                       uintptr
	getProcessId                          uintptr
	getProcessIdOfThread                  uintptr
	getProcessPriorityBoost               uintptr
	getProcessShutdownParameters          uintptr
	getProcessTimes                       uintptr
	getProcessVersion                     uintptr
	getProductInfo                        uintptr
	getProfileInt                         uintptr
	getProfileSection                     uintptr
	getProfileString                      uintptr
	getShortPathName                      uintptr
	getStartupInfo                        uintptr
	getStdHandle                          uintptr
	getStringScripts                      uintptr
	getSystemDefaultLCID                  uintptr
	getSystemDefaultLangID                uintptr
	getSystemDefaultLocaleName            uintptr
	getSystemDefaultUILanguage            uintptr
	getSystemDirectory                    uintptr
	getSystemFirmwareTable                uintptr
	getSystemInfo                         uintptr
	getSystemRegistryQuota                uintptr
	getSystemTime                         uintptr
	getSystemTimeAdjustment               uintptr
	getSystemTimeAsFileTime               uintptr
	getSystemTimePreciseAsFileTime        uintptr
	getSystemTimes                        uintptr
	getSystemWindowsDirectory             uintptr
	getSystemWow64Directory               uintptr
	getTapeParameters                     uintptr
	getTapePosition                       uintptr
	getTapeStatus                         uintptr
	getTempFileName                       uintptr
	getTempPath                           uintptr
	getThreadErrorMode                    uintptr
	getThreadIOPendingFlag                uintptr
	getThreadId                           uintptr
	getThreadLocale                       uintptr
	getThreadPriority                     uintptr
	getThreadPriorityBoost                uintptr
	getThreadTimes                        uintptr
	getThreadUILanguage                   uintptr
	getTickCount                          uintptr
	getTickCount64                        uintptr
	getTimeFormatEx                       uintptr
	getTimeFormat                         uintptr
	getUserDefaultLCID                    uintptr
	getUserDefaultLangID                  uintptr
	getUserDefaultLocaleName              uintptr
	getUserDefaultUILanguage              uintptr
	getVersion                            uintptr
	getVolumeInformationByHandleW         uintptr
	getVolumeInformation                  uintptr
	getVolumeNameForVolumeMountPoint      uintptr
	getVolumePathName                     uintptr
	getWindowsDirectory                   uintptr
	getWriteWatch                         uintptr
	globalAddAtom                         uintptr
	globalAlloc                           uintptr
	globalCompact                         uintptr
	globalDeleteAtom                      uintptr
	globalFindAtom                        uintptr
	globalFix                             uintptr
	globalFlags                           uintptr
	globalFree                            uintptr
	globalGetAtomName                     uintptr
	globalHandle                          uintptr
	globalLock                            uintptr
	globalReAlloc                         uintptr
	globalSize                            uintptr
	globalUnWire                          uintptr
	globalUnfix                           uintptr
	globalUnlock                          uintptr
	globalWire                            uintptr
	heapAlloc                             uintptr
	heapCompact                           uintptr
	heapCreate                            uintptr
	heapDestroy                           uintptr
	heapFree                              uintptr
	heapLock                              uintptr
	heapReAlloc                           uintptr
	heapSize                              uintptr
	heapUnlock                            uintptr
	heapValidate                          uintptr
	idnToAscii                            uintptr
	idnToNameprepUnicode                  uintptr
	idnToUnicode                          uintptr
	initAtomTable                         uintptr
	initializeSRWLock                     uintptr
	isBadCodePtr                          uintptr
	isBadHugeReadPtr                      uintptr
	isBadHugeWritePtr                     uintptr
	isBadReadPtr                          uintptr
	isBadStringPtr                        uintptr
	isBadWritePtr                         uintptr
	isDBCSLeadByte                        uintptr
	isDBCSLeadByteEx                      uintptr
	isDebuggerPresent                     uintptr
	isProcessInJob                        uintptr
	isProcessorFeaturePresent             uintptr
	isSystemResumeAutomatic               uintptr
	isThreadAFiber                        uintptr
	isValidCodePage                       uintptr
	isValidLocale                         uintptr
	isValidLocaleName                     uintptr
	isWow64Process                        uintptr
	lCIDToLocaleName                      uintptr
	lCMapString                           uintptr
	lZClose                               uintptr
	lZCopy                                uintptr
	lZDone                                uintptr
	lZInit                                uintptr
	lZRead                                uintptr
	lZSeek                                uintptr
	lZStart                               uintptr
	loadLibraryEx                         uintptr
	loadLibrary                           uintptr
	loadModule                            uintptr
	loadPackagedLibrary                   uintptr
	loadResource                          uintptr
	localAlloc                            uintptr
	localCompact                          uintptr
	localFileTimeToFileTime               uintptr
	localFlags                            uintptr
	localFree                             uintptr
	localHandle                           uintptr
	localLock                             uintptr
	localReAlloc                          uintptr
	localShrink                           uintptr
	localSize                             uintptr
	localUnlock                           uintptr
	localeNameToLCID                      uintptr
	lockFile                              uintptr
	lockFileEx                            uintptr
	lockResource                          uintptr
	mapViewOfFile                         uintptr
	mapViewOfFileEx                       uintptr
	mapViewOfFileExNuma                   uintptr
	moveFileEx                            uintptr
	moveFile                              uintptr
	mulDiv                                uintptr
	needCurrentDirectoryForExePath        uintptr
	notifyUILanguageChange                uintptr
	openEvent                             uintptr
	openFileMapping                       uintptr
	openJobObject                         uintptr
	openMutex                             uintptr
	openPrivateNamespace                  uintptr
	openProcess                           uintptr
	openSemaphore                         uintptr
	openThread                            uintptr
	openWaitableTimer                     uintptr
	outputDebugString                     uintptr
	peekNamedPipe                         uintptr
	postQueuedCompletionStatus            uintptr
	prepareTape                           uintptr
	processIdToSessionId                  uintptr
	pulseEvent                            uintptr
	purgeComm                             uintptr
	queryActCtxSettingsW                  uintptr
	queryActCtxW                          uintptr
	queryDosDevice                        uintptr
	queryFullProcessImageName             uintptr
	queryIdleProcessorCycleTime           uintptr
	queryIdleProcessorCycleTimeEx         uintptr
	queryMemoryResourceNotification       uintptr
	queryPerformanceCounter               uintptr
	queryPerformanceFrequency             uintptr
	queryProcessAffinityUpdateMode        uintptr
	queryProcessCycleTime                 uintptr
	queryThreadCycleTime                  uintptr
	queueUserWorkItem                     uintptr
	raiseException                        uintptr
	reOpenFile                            uintptr
	readConsoleOutputAttribute            uintptr
	readConsoleOutputCharacter            uintptr
	readConsole                           uintptr
	readFile                              uintptr
	readProcessMemory                     uintptr
	reclaimVirtualMemory                  uintptr
	registerApplicationRestart            uintptr
	releaseActCtx                         uintptr
	releaseMutex                          uintptr
	releaseSRWLockExclusive               uintptr
	releaseSRWLockShared                  uintptr
	releaseSemaphore                      uintptr
	removeDirectoryTransacted             uintptr
	removeDirectory                       uintptr
	removeSecureMemoryCacheCallback       uintptr
	removeVectoredContinueHandler         uintptr
	removeVectoredExceptionHandler        uintptr
	replaceFile                           uintptr
	replacePartitionUnit                  uintptr
	requestDeviceWakeup                   uintptr
	resetEvent                            uintptr
	resetWriteWatch                       uintptr
	resolveLocaleName                     uintptr
	restoreLastError                      uintptr
	resumeThread                          uintptr
	searchPath                            uintptr
	setCommBreak                          uintptr
	setCommMask                           uintptr
	setComputerName                       uintptr
	setConsoleActiveScreenBuffer          uintptr
	setConsoleCP                          uintptr
	setConsoleCursorPosition              uintptr
	setConsoleMode                        uintptr
	setConsoleOutputCP                    uintptr
	setConsoleScreenBufferSize            uintptr
	setConsoleTextAttribute               uintptr
	setConsoleTitle                       uintptr
	setCurrentDirectory                   uintptr
	setDllDirectory                       uintptr
	setEndOfFile                          uintptr
	setEnvironmentVariable                uintptr
	setErrorMode                          uintptr
	setEvent                              uintptr
	setFileApisToANSI                     uintptr
	setFileApisToOEM                      uintptr
	setFileAttributesTransacted           uintptr
	setFileAttributes                     uintptr
	setFileBandwidthReservation           uintptr
	setFileCompletionNotificationModes    uintptr
	setFilePointer                        uintptr
	setFileShortName                      uintptr
	setFileTime                           uintptr
	setFileValidData                      uintptr
	setFirmwareEnvironmentVariable        uintptr
	setHandleCount                        uintptr
	setHandleInformation                  uintptr
	setLastError                          uintptr
	setLocalTime                          uintptr
	setLocaleInfo                         uintptr
	setMailslotInfo                       uintptr
	setMessageWaitingIndicator            uintptr
	setNamedPipeHandleState               uintptr
	setPriorityClass                      uintptr
	setProcessAffinityMask                uintptr
	setProcessAffinityUpdateMode          uintptr
	setProcessDEPPolicy                   uintptr
	setProcessPriorityBoost               uintptr
	setProcessShutdownParameters          uintptr
	setProcessWorkingSetSize              uintptr
	setProcessWorkingSetSizeEx            uintptr
	setSearchPathMode                     uintptr
	setStdHandle                          uintptr
	setStdHandleEx                        uintptr
	setSystemFileCacheSize                uintptr
	setSystemPowerState                   uintptr
	setSystemTime                         uintptr
	setSystemTimeAdjustment               uintptr
	setTapeParameters                     uintptr
	setTapePosition                       uintptr
	setThreadAffinityMask                 uintptr
	setThreadErrorMode                    uintptr
	setThreadIdealProcessor               uintptr
	setThreadLocale                       uintptr
	setThreadPriority                     uintptr
	setThreadPriorityBoost                uintptr
	setThreadStackGuarantee               uintptr
	setThreadUILanguage                   uintptr
	setUserGeoID                          uintptr
	setVolumeLabel                        uintptr
	setVolumeMountPoint                   uintptr
	setupComm                             uintptr
	signalObjectAndWait                   uintptr
	sizeofResource                        uintptr
	sleep                                 uintptr
	sleepEx                               uintptr
	suspendThread                         uintptr
	switchToFiber                         uintptr
	switchToThread                        uintptr
	systemTimeToFileTime                  uintptr
	terminateJobObject                    uintptr
	terminateProcess                      uintptr
	terminateThread                       uintptr
	tlsAlloc                              uintptr
	tlsFree                               uintptr
	tlsGetValue                           uintptr
	tlsSetValue                           uintptr
	toolhelp32ReadProcessMemory           uintptr
	transactNamedPipe                     uintptr
	transmitCommChar                      uintptr
	tryAcquireSRWLockExclusive            uintptr
	tryAcquireSRWLockShared               uintptr
	unlockFile                            uintptr
	unlockFileEx                          uintptr
	unmapViewOfFile                       uintptr
	unregisterApplicationRecoveryCallback uintptr
	unregisterApplicationRestart          uintptr
	unregisterWait                        uintptr
	unregisterWaitEx                      uintptr
	updateResource                        uintptr
	verLanguageName                       uintptr
	verifyScripts                         uintptr
	virtualAlloc                          uintptr
	virtualAllocEx                        uintptr
	virtualAllocExNuma                    uintptr
	virtualFree                           uintptr
	virtualFreeEx                         uintptr
	virtualLock                           uintptr
	virtualProtect                        uintptr
	virtualProtectEx                      uintptr
	virtualUnlock                         uintptr
	wTSGetActiveConsoleSessionId          uintptr
	waitCommEvent                         uintptr
	waitForMultipleObjects                uintptr
	waitForMultipleObjectsEx              uintptr
	waitForSingleObject                   uintptr
	waitForSingleObjectEx                 uintptr
	waitNamedPipe                         uintptr
	werGetFlags                           uintptr
	werRegisterMemoryBlock                uintptr
	werRegisterRuntimeExceptionModule     uintptr
	werSetFlags                           uintptr
	werUnregisterFile                     uintptr
	werUnregisterMemoryBlock              uintptr
	werUnregisterRuntimeExceptionModule   uintptr
	winExec                               uintptr
	wow64DisableWow64FsRedirection        uintptr
	wow64EnableWow64FsRedirection         uintptr
	wow64RevertWow64FsRedirection         uintptr
	wow64SuspendThread                    uintptr
	writeConsoleOutputAttribute           uintptr
	writeConsoleOutputCharacter           uintptr
	writeConsole                          uintptr
	writeFile                             uintptr
	writePrivateProfileSection            uintptr
	writePrivateProfileString             uintptr
	writePrivateProfileStruct             uintptr
	writeProcessMemory                    uintptr
	writeProfileSection                   uintptr
	writeProfileString                    uintptr
	writeTapemark                         uintptr
	zombifyActCtx                         uintptr
	lstrcat                               uintptr
	lstrcmp                               uintptr
	lstrcmpi                              uintptr
	lstrcpy                               uintptr
	lstrcpyn                              uintptr
	lstrlen                               uintptr
	closeConsoleHandle                    uintptr
	closeProfileUserMapping               uintptr
	cmdBatNotification                    uintptr
	delayLoadFailureHook                  uintptr
	duplicateConsoleHandle                uintptr
	expungeConsoleCommandHistory          uintptr
	getConsoleCommandHistoryLength        uintptr
	getConsoleCommandHistory              uintptr
	getConsoleInputExeName                uintptr
	getConsoleInputWaitHandle             uintptr
	getConsoleKeyboardLayoutName          uintptr
	getNumberOfConsoleFonts               uintptr
	k32EmptyWorkingSet                    uintptr
	k32EnumDeviceDrivers                  uintptr
	k32EnumPageFiles                      uintptr
	k32EnumProcessModules                 uintptr
	k32EnumProcessModulesEx               uintptr
	k32EnumProcesses                      uintptr
	k32GetDeviceDriverBaseName            uintptr
	k32GetDeviceDriverFileName            uintptr
	k32GetMappedFileName                  uintptr
	k32GetModuleBaseName                  uintptr
	k32GetModuleFileNameEx                uintptr
	k32GetModuleInformation               uintptr
	k32GetProcessImageFileName            uintptr
	k32GetProcessMemoryInfo               uintptr
	k32GetWsChanges                       uintptr
	k32InitializeProcessForWsWatch        uintptr
	k32QueryWorkingSet                    uintptr
	k32QueryWorkingSetEx                  uintptr
	openConsoleW                          uintptr
	openProfileUserMapping                uintptr
	rtlCaptureStackBackTrace              uintptr
	rtlCompareMemory                      uintptr
	rtlCopyMemory                         uintptr
	rtlFillMemory                         uintptr
	rtlMoveMemory                         uintptr
	rtlPcToFileHeader                     uintptr
	rtlZeroMemory                         uintptr
	setCPGlobal                           uintptr
	setConsoleFont                        uintptr
	setConsoleIcon                        uintptr
	setConsoleInputExeName                uintptr
	setConsoleKeyShortcuts                uintptr
	setTermsrvAppInstallMode              uintptr
	termsrvAppInstallMode                 uintptr
	uTRegister                            uintptr
	uTUnRegister                          uintptr
	verSetConditionMask                   uintptr
	verifyConsoleIoHandle                 uintptr
)

func init() {
	// Library
	libkernel32 = doLoadLibrary("kernel32.dll")

	// Functions
	acquireSRWLockExclusive = doGetProcAddress(libkernel32, "AcquireSRWLockExclusive")
	acquireSRWLockShared = doGetProcAddress(libkernel32, "AcquireSRWLockShared")
	activateActCtx = doGetProcAddress(libkernel32, "ActivateActCtx")
	addAtom = doGetProcAddress(libkernel32, "AddAtomW")
	addConsoleAlias = doGetProcAddress(libkernel32, "AddConsoleAliasW")
	addIntegrityLabelToBoundaryDescriptor = doGetProcAddress(libkernel32, "AddIntegrityLabelToBoundaryDescriptor")
	addRefActCtx = doGetProcAddress(libkernel32, "AddRefActCtx")
	addSIDToBoundaryDescriptor = doGetProcAddress(libkernel32, "AddSIDToBoundaryDescriptor")
	addSecureMemoryCacheCallback = doGetProcAddress(libkernel32, "AddSecureMemoryCacheCallback")
	allocConsole = doGetProcAddress(libkernel32, "AllocConsole")
	applicationRecoveryFinished = doGetProcAddress(libkernel32, "ApplicationRecoveryFinished")
	applicationRecoveryInProgress = doGetProcAddress(libkernel32, "ApplicationRecoveryInProgress")
	areFileApisANSI = doGetProcAddress(libkernel32, "AreFileApisANSI")
	assignProcessToJobObject = doGetProcAddress(libkernel32, "AssignProcessToJobObject")
	attachConsole = doGetProcAddress(libkernel32, "AttachConsole")
	backupRead = doGetProcAddress(libkernel32, "BackupRead")
	backupSeek = doGetProcAddress(libkernel32, "BackupSeek")
	backupWrite = doGetProcAddress(libkernel32, "BackupWrite")
	beep = doGetProcAddress(libkernel32, "Beep")
	beginUpdateResource = doGetProcAddress(libkernel32, "BeginUpdateResourceW")
	callNamedPipe = doGetProcAddress(libkernel32, "CallNamedPipeW")
	cancelDeviceWakeupRequest = doGetProcAddress(libkernel32, "CancelDeviceWakeupRequest")
	cancelIo = doGetProcAddress(libkernel32, "CancelIo")
	cancelIoEx = doGetProcAddress(libkernel32, "CancelIoEx")
	cancelSynchronousIo = doGetProcAddress(libkernel32, "CancelSynchronousIo")
	cancelTimerQueueTimer = doGetProcAddress(libkernel32, "CancelTimerQueueTimer")
	cancelWaitableTimer = doGetProcAddress(libkernel32, "CancelWaitableTimer")
	changeTimerQueueTimer = doGetProcAddress(libkernel32, "ChangeTimerQueueTimer")
	checkNameLegalDOS8Dot3 = doGetProcAddress(libkernel32, "CheckNameLegalDOS8Dot3W")
	checkRemoteDebuggerPresent = doGetProcAddress(libkernel32, "CheckRemoteDebuggerPresent")
	clearCommBreak = doGetProcAddress(libkernel32, "ClearCommBreak")
	closeHandle = doGetProcAddress(libkernel32, "CloseHandle")
	closePrivateNamespace = doGetProcAddress(libkernel32, "ClosePrivateNamespace")
	compareFileTime = doGetProcAddress(libkernel32, "CompareFileTime")
	connectNamedPipe = doGetProcAddress(libkernel32, "ConnectNamedPipe")
	continueDebugEvent = doGetProcAddress(libkernel32, "ContinueDebugEvent")
	convertDefaultLocale = doGetProcAddress(libkernel32, "ConvertDefaultLocale")
	convertFiberToThread = doGetProcAddress(libkernel32, "ConvertFiberToThread")
	convertThreadToFiber = doGetProcAddress(libkernel32, "ConvertThreadToFiber")
	convertThreadToFiberEx = doGetProcAddress(libkernel32, "ConvertThreadToFiberEx")
	copyFile = doGetProcAddress(libkernel32, "CopyFileW")
	copyLZFile = doGetProcAddress(libkernel32, "CopyLZFile")
	createBoundaryDescriptor = doGetProcAddress(libkernel32, "CreateBoundaryDescriptorW")
	createConsoleScreenBuffer = doGetProcAddress(libkernel32, "CreateConsoleScreenBuffer")
	createDirectoryEx = doGetProcAddress(libkernel32, "CreateDirectoryExW")
	createDirectoryTransacted = doGetProcAddress(libkernel32, "CreateDirectoryTransactedW")
	createDirectory = doGetProcAddress(libkernel32, "CreateDirectoryW")
	createEventEx = doGetProcAddress(libkernel32, "CreateEventExW")
	createEvent = doGetProcAddress(libkernel32, "CreateEventW")
	createFileMappingNuma = doGetProcAddress(libkernel32, "CreateFileMappingNumaW")
	createFileMapping = doGetProcAddress(libkernel32, "CreateFileMappingW")
	createFileTransacted = doGetProcAddress(libkernel32, "CreateFileTransactedW")
	createFile = doGetProcAddress(libkernel32, "CreateFileW")
	createHardLinkTransacted = doGetProcAddress(libkernel32, "CreateHardLinkTransactedW")
	createHardLink = doGetProcAddress(libkernel32, "CreateHardLinkW")
	createIoCompletionPort = doGetProcAddress(libkernel32, "CreateIoCompletionPort")
	createJobObject = doGetProcAddress(libkernel32, "CreateJobObjectW")
	createMailslot = doGetProcAddress(libkernel32, "CreateMailslotW")
	createMutexEx = doGetProcAddress(libkernel32, "CreateMutexExW")
	createMutex = doGetProcAddress(libkernel32, "CreateMutexW")
	createNamedPipe = doGetProcAddress(libkernel32, "CreateNamedPipeW")
	createPipe = doGetProcAddress(libkernel32, "CreatePipe")
	createPrivateNamespace = doGetProcAddress(libkernel32, "CreatePrivateNamespaceW")
	createProcess = doGetProcAddress(libkernel32, "CreateProcessW")
	createRemoteThread = doGetProcAddress(libkernel32, "CreateRemoteThread")
	createSemaphoreEx = doGetProcAddress(libkernel32, "CreateSemaphoreExW")
	createSemaphore = doGetProcAddress(libkernel32, "CreateSemaphoreW")
	createSymbolicLinkTransacted = doGetProcAddress(libkernel32, "CreateSymbolicLinkTransactedW")
	createSymbolicLink = doGetProcAddress(libkernel32, "CreateSymbolicLinkW")
	createTapePartition = doGetProcAddress(libkernel32, "CreateTapePartition")
	createThread = doGetProcAddress(libkernel32, "CreateThread")
	createTimerQueue = doGetProcAddress(libkernel32, "CreateTimerQueue")
	createToolhelp32Snapshot = doGetProcAddress(libkernel32, "CreateToolhelp32Snapshot")
	createWaitableTimerEx = doGetProcAddress(libkernel32, "CreateWaitableTimerExW")
	createWaitableTimer = doGetProcAddress(libkernel32, "CreateWaitableTimerW")
	deactivateActCtx = doGetProcAddress(libkernel32, "DeactivateActCtx")
	debugActiveProcess = doGetProcAddress(libkernel32, "DebugActiveProcess")
	debugActiveProcessStop = doGetProcAddress(libkernel32, "DebugActiveProcessStop")
	debugBreak = doGetProcAddress(libkernel32, "DebugBreak")
	debugBreakProcess = doGetProcAddress(libkernel32, "DebugBreakProcess")
	debugSetProcessKillOnExit = doGetProcAddress(libkernel32, "DebugSetProcessKillOnExit")
	decodePointer = doGetProcAddress(libkernel32, "DecodePointer")
	decodeSystemPointer = doGetProcAddress(libkernel32, "DecodeSystemPointer")
	defineDosDevice = doGetProcAddress(libkernel32, "DefineDosDeviceW")
	deleteAtom = doGetProcAddress(libkernel32, "DeleteAtom")
	deleteBoundaryDescriptor = doGetProcAddress(libkernel32, "DeleteBoundaryDescriptor")
	deleteFiber = doGetProcAddress(libkernel32, "DeleteFiber")
	deleteFileTransacted = doGetProcAddress(libkernel32, "DeleteFileTransactedW")
	deleteFile = doGetProcAddress(libkernel32, "DeleteFileW")
	deleteTimerQueue = doGetProcAddress(libkernel32, "DeleteTimerQueue")
	deleteTimerQueueEx = doGetProcAddress(libkernel32, "DeleteTimerQueueEx")
	deleteTimerQueueTimer = doGetProcAddress(libkernel32, "DeleteTimerQueueTimer")
	deleteVolumeMountPoint = doGetProcAddress(libkernel32, "DeleteVolumeMountPointW")
	deviceIoControl = doGetProcAddress(libkernel32, "DeviceIoControl")
	disableThreadLibraryCalls = doGetProcAddress(libkernel32, "DisableThreadLibraryCalls")
	disableThreadProfiling = doGetProcAddress(libkernel32, "DisableThreadProfiling")
	discardVirtualMemory = doGetProcAddress(libkernel32, "DiscardVirtualMemory")
	disconnectNamedPipe = doGetProcAddress(libkernel32, "DisconnectNamedPipe")
	dnsHostnameToComputerName = doGetProcAddress(libkernel32, "DnsHostnameToComputerNameW")
	dosDateTimeToFileTime = doGetProcAddress(libkernel32, "DosDateTimeToFileTime")
	duplicateHandle = doGetProcAddress(libkernel32, "DuplicateHandle")
	encodePointer = doGetProcAddress(libkernel32, "EncodePointer")
	encodeSystemPointer = doGetProcAddress(libkernel32, "EncodeSystemPointer")
	endUpdateResource = doGetProcAddress(libkernel32, "EndUpdateResourceW")
	enumResourceLanguagesEx = doGetProcAddress(libkernel32, "EnumResourceLanguagesExW")
	enumResourceLanguages = doGetProcAddress(libkernel32, "EnumResourceLanguagesW")
	enumSystemFirmwareTables = doGetProcAddress(libkernel32, "EnumSystemFirmwareTables")
	eraseTape = doGetProcAddress(libkernel32, "EraseTape")
	escapeCommFunction = doGetProcAddress(libkernel32, "EscapeCommFunction")
	exitProcess = doGetProcAddress(libkernel32, "ExitProcess")
	exitThread = doGetProcAddress(libkernel32, "ExitThread")
	expandEnvironmentStrings = doGetProcAddress(libkernel32, "ExpandEnvironmentStringsW")
	fatalAppExit = doGetProcAddress(libkernel32, "FatalAppExitW")
	fatalExit = doGetProcAddress(libkernel32, "FatalExit")
	fileTimeToDosDateTime = doGetProcAddress(libkernel32, "FileTimeToDosDateTime")
	fileTimeToLocalFileTime = doGetProcAddress(libkernel32, "FileTimeToLocalFileTime")
	fileTimeToSystemTime = doGetProcAddress(libkernel32, "FileTimeToSystemTime")
	fillConsoleOutputAttribute = doGetProcAddress(libkernel32, "FillConsoleOutputAttribute")
	fillConsoleOutputCharacter = doGetProcAddress(libkernel32, "FillConsoleOutputCharacterW")
	findAtom = doGetProcAddress(libkernel32, "FindAtomW")
	findClose = doGetProcAddress(libkernel32, "FindClose")
	findCloseChangeNotification = doGetProcAddress(libkernel32, "FindCloseChangeNotification")
	findFirstChangeNotification = doGetProcAddress(libkernel32, "FindFirstChangeNotificationW")
	findFirstFileNameTransactedW = doGetProcAddress(libkernel32, "FindFirstFileNameTransactedW")
	findFirstFileNameW = doGetProcAddress(libkernel32, "FindFirstFileNameW")
	findFirstVolumeMountPoint = doGetProcAddress(libkernel32, "FindFirstVolumeMountPointW")
	findFirstVolume = doGetProcAddress(libkernel32, "FindFirstVolumeW")
	findNLSString = doGetProcAddress(libkernel32, "FindNLSString")
	findNextChangeNotification = doGetProcAddress(libkernel32, "FindNextChangeNotification")
	findNextFileNameW = doGetProcAddress(libkernel32, "FindNextFileNameW")
	findNextStreamW = doGetProcAddress(libkernel32, "FindNextStreamW")
	findNextVolumeMountPoint = doGetProcAddress(libkernel32, "FindNextVolumeMountPointW")
	findNextVolume = doGetProcAddress(libkernel32, "FindNextVolumeW")
	findResourceEx = doGetProcAddress(libkernel32, "FindResourceExW")
	findResource = doGetProcAddress(libkernel32, "FindResourceW")
	findStringOrdinal = doGetProcAddress(libkernel32, "FindStringOrdinal")
	findVolumeClose = doGetProcAddress(libkernel32, "FindVolumeClose")
	findVolumeMountPointClose = doGetProcAddress(libkernel32, "FindVolumeMountPointClose")
	flsFree = doGetProcAddress(libkernel32, "FlsFree")
	flsGetValue = doGetProcAddress(libkernel32, "FlsGetValue")
	flsSetValue = doGetProcAddress(libkernel32, "FlsSetValue")
	flushConsoleInputBuffer = doGetProcAddress(libkernel32, "FlushConsoleInputBuffer")
	flushFileBuffers = doGetProcAddress(libkernel32, "FlushFileBuffers")
	flushInstructionCache = doGetProcAddress(libkernel32, "FlushInstructionCache")
	flushProcessWriteBuffers = doGetProcAddress(libkernel32, "FlushProcessWriteBuffers")
	flushViewOfFile = doGetProcAddress(libkernel32, "FlushViewOfFile")
	freeConsole = doGetProcAddress(libkernel32, "FreeConsole")
	freeLibrary = doGetProcAddress(libkernel32, "FreeLibrary")
	freeLibraryAndExitThread = doGetProcAddress(libkernel32, "FreeLibraryAndExitThread")
	freeResource = doGetProcAddress(libkernel32, "FreeResource")
	generateConsoleCtrlEvent = doGetProcAddress(libkernel32, "GenerateConsoleCtrlEvent")
	getACP = doGetProcAddress(libkernel32, "GetACP")
	getActiveProcessorCount = doGetProcAddress(libkernel32, "GetActiveProcessorCount")
	getActiveProcessorGroupCount = doGetProcAddress(libkernel32, "GetActiveProcessorGroupCount")
	getApplicationRestartSettings = doGetProcAddress(libkernel32, "GetApplicationRestartSettings")
	getAtomName = doGetProcAddress(libkernel32, "GetAtomNameW")
	getBinaryType = doGetProcAddress(libkernel32, "GetBinaryTypeW")
	getCommMask = doGetProcAddress(libkernel32, "GetCommMask")
	getCommModemStatus = doGetProcAddress(libkernel32, "GetCommModemStatus")
	getCommandLine = doGetProcAddress(libkernel32, "GetCommandLineW")
	getCompressedFileSizeTransacted = doGetProcAddress(libkernel32, "GetCompressedFileSizeTransactedW")
	getCompressedFileSize = doGetProcAddress(libkernel32, "GetCompressedFileSizeW")
	getComputerName = doGetProcAddress(libkernel32, "GetComputerNameW")
	getConsoleAliasExesLength = doGetProcAddress(libkernel32, "GetConsoleAliasExesLengthW")
	getConsoleAliasExes = doGetProcAddress(libkernel32, "GetConsoleAliasExesW")
	getConsoleAlias = doGetProcAddress(libkernel32, "GetConsoleAliasW")
	getConsoleAliasesLength = doGetProcAddress(libkernel32, "GetConsoleAliasesLengthW")
	getConsoleAliases = doGetProcAddress(libkernel32, "GetConsoleAliasesW")
	getConsoleCP = doGetProcAddress(libkernel32, "GetConsoleCP")
	getConsoleDisplayMode = doGetProcAddress(libkernel32, "GetConsoleDisplayMode")
	getConsoleFontSize = doGetProcAddress(libkernel32, "GetConsoleFontSize")
	getConsoleMode = doGetProcAddress(libkernel32, "GetConsoleMode")
	getConsoleOriginalTitle = doGetProcAddress(libkernel32, "GetConsoleOriginalTitleW")
	getConsoleOutputCP = doGetProcAddress(libkernel32, "GetConsoleOutputCP")
	getConsoleProcessList = doGetProcAddress(libkernel32, "GetConsoleProcessList")
	getConsoleTitle = doGetProcAddress(libkernel32, "GetConsoleTitleW")
	getConsoleWindow = doGetProcAddress(libkernel32, "GetConsoleWindow")
	getCurrentActCtx = doGetProcAddress(libkernel32, "GetCurrentActCtx")
	getCurrentDirectory = doGetProcAddress(libkernel32, "GetCurrentDirectoryW")
	getCurrentProcess = doGetProcAddress(libkernel32, "GetCurrentProcess")
	getCurrentProcessId = doGetProcAddress(libkernel32, "GetCurrentProcessId")
	getCurrentProcessorNumber = doGetProcAddress(libkernel32, "GetCurrentProcessorNumber")
	getCurrentThread = doGetProcAddress(libkernel32, "GetCurrentThread")
	getCurrentThreadId = doGetProcAddress(libkernel32, "GetCurrentThreadId")
	getDateFormatEx = doGetProcAddress(libkernel32, "GetDateFormatEx")
	getDateFormat = doGetProcAddress(libkernel32, "GetDateFormatW")
	getDevicePowerState = doGetProcAddress(libkernel32, "GetDevicePowerState")
	getDiskFreeSpace = doGetProcAddress(libkernel32, "GetDiskFreeSpaceW")
	getDllDirectory = doGetProcAddress(libkernel32, "GetDllDirectoryW")
	getDriveType = doGetProcAddress(libkernel32, "GetDriveTypeW")
	getDurationFormat = doGetProcAddress(libkernel32, "GetDurationFormat")
	getDurationFormatEx = doGetProcAddress(libkernel32, "GetDurationFormatEx")
	getEnvironmentVariable = doGetProcAddress(libkernel32, "GetEnvironmentVariableW")
	getErrorMode = doGetProcAddress(libkernel32, "GetErrorMode")
	getExitCodeProcess = doGetProcAddress(libkernel32, "GetExitCodeProcess")
	getExitCodeThread = doGetProcAddress(libkernel32, "GetExitCodeThread")
	getExpandedName = doGetProcAddress(libkernel32, "GetExpandedNameW")
	getFileAttributes = doGetProcAddress(libkernel32, "GetFileAttributesW")
	getFileBandwidthReservation = doGetProcAddress(libkernel32, "GetFileBandwidthReservation")
	getFileSize = doGetProcAddress(libkernel32, "GetFileSize")
	getFileTime = doGetProcAddress(libkernel32, "GetFileTime")
	getFileType = doGetProcAddress(libkernel32, "GetFileType")
	getFinalPathNameByHandle = doGetProcAddress(libkernel32, "GetFinalPathNameByHandleW")
	getFirmwareEnvironmentVariable = doGetProcAddress(libkernel32, "GetFirmwareEnvironmentVariableW")
	getFullPathNameTransacted = doGetProcAddress(libkernel32, "GetFullPathNameTransactedW")
	getFullPathName = doGetProcAddress(libkernel32, "GetFullPathNameW")
	getHandleInformation = doGetProcAddress(libkernel32, "GetHandleInformation")
	getLargePageMinimum = doGetProcAddress(libkernel32, "GetLargePageMinimum")
	getLargestConsoleWindowSize = doGetProcAddress(libkernel32, "GetLargestConsoleWindowSize")
	getLastError = doGetProcAddress(libkernel32, "GetLastError")
	getLocalTime = doGetProcAddress(libkernel32, "GetLocalTime")
	getLocaleInfoEx = doGetProcAddress(libkernel32, "GetLocaleInfoEx")
	getLocaleInfo = doGetProcAddress(libkernel32, "GetLocaleInfoW")
	getLogicalDriveStrings = doGetProcAddress(libkernel32, "GetLogicalDriveStringsW")
	getLogicalDrives = doGetProcAddress(libkernel32, "GetLogicalDrives")
	getLongPathNameTransacted = doGetProcAddress(libkernel32, "GetLongPathNameTransactedW")
	getLongPathName = doGetProcAddress(libkernel32, "GetLongPathNameW")
	getMailslotInfo = doGetProcAddress(libkernel32, "GetMailslotInfo")
	getMaximumProcessorCount = doGetProcAddress(libkernel32, "GetMaximumProcessorCount")
	getMaximumProcessorGroupCount = doGetProcAddress(libkernel32, "GetMaximumProcessorGroupCount")
	getModuleFileName = doGetProcAddress(libkernel32, "GetModuleFileNameW")
	getModuleHandleEx = doGetProcAddress(libkernel32, "GetModuleHandleExW")
	getModuleHandle = doGetProcAddress(libkernel32, "GetModuleHandleW")
	getNamedPipeClientComputerName = doGetProcAddress(libkernel32, "GetNamedPipeClientComputerNameW")
	getNamedPipeClientProcessId = doGetProcAddress(libkernel32, "GetNamedPipeClientProcessId")
	getNamedPipeClientSessionId = doGetProcAddress(libkernel32, "GetNamedPipeClientSessionId")
	getNamedPipeHandleState = doGetProcAddress(libkernel32, "GetNamedPipeHandleStateW")
	getNamedPipeInfo = doGetProcAddress(libkernel32, "GetNamedPipeInfo")
	getNamedPipeServerProcessId = doGetProcAddress(libkernel32, "GetNamedPipeServerProcessId")
	getNamedPipeServerSessionId = doGetProcAddress(libkernel32, "GetNamedPipeServerSessionId")
	getNativeSystemInfo = doGetProcAddress(libkernel32, "GetNativeSystemInfo")
	getNumaHighestNodeNumber = doGetProcAddress(libkernel32, "GetNumaHighestNodeNumber")
	getNumaNodeNumberFromHandle = doGetProcAddress(libkernel32, "GetNumaNodeNumberFromHandle")
	getNumaProximityNodeEx = doGetProcAddress(libkernel32, "GetNumaProximityNodeEx")
	getNumberOfConsoleInputEvents = doGetProcAddress(libkernel32, "GetNumberOfConsoleInputEvents")
	getNumberOfConsoleMouseButtons = doGetProcAddress(libkernel32, "GetNumberOfConsoleMouseButtons")
	getOEMCP = doGetProcAddress(libkernel32, "GetOEMCP")
	getOverlappedResult = doGetProcAddress(libkernel32, "GetOverlappedResult")
	getPriorityClass = doGetProcAddress(libkernel32, "GetPriorityClass")
	getPrivateProfileInt = doGetProcAddress(libkernel32, "GetPrivateProfileIntW")
	getPrivateProfileSectionNames = doGetProcAddress(libkernel32, "GetPrivateProfileSectionNamesW")
	getPrivateProfileSection = doGetProcAddress(libkernel32, "GetPrivateProfileSectionW")
	getPrivateProfileString = doGetProcAddress(libkernel32, "GetPrivateProfileStringW")
	getPrivateProfileStruct = doGetProcAddress(libkernel32, "GetPrivateProfileStructW")
	getProcAddress = doGetProcAddress(libkernel32, "GetProcAddress")
	getProcessAffinityMask = doGetProcAddress(libkernel32, "GetProcessAffinityMask")
	getProcessDEPPolicy = doGetProcAddress(libkernel32, "GetProcessDEPPolicy")
	getProcessGroupAffinity = doGetProcAddress(libkernel32, "GetProcessGroupAffinity")
	getProcessHandleCount = doGetProcAddress(libkernel32, "GetProcessHandleCount")
	getProcessHeap = doGetProcAddress(libkernel32, "GetProcessHeap")
	getProcessHeaps = doGetProcAddress(libkernel32, "GetProcessHeaps")
	getProcessId = doGetProcAddress(libkernel32, "GetProcessId")
	getProcessIdOfThread = doGetProcAddress(libkernel32, "GetProcessIdOfThread")
	getProcessPriorityBoost = doGetProcAddress(libkernel32, "GetProcessPriorityBoost")
	getProcessShutdownParameters = doGetProcAddress(libkernel32, "GetProcessShutdownParameters")
	getProcessTimes = doGetProcAddress(libkernel32, "GetProcessTimes")
	getProcessVersion = doGetProcAddress(libkernel32, "GetProcessVersion")
	getProductInfo = doGetProcAddress(libkernel32, "GetProductInfo")
	getProfileInt = doGetProcAddress(libkernel32, "GetProfileIntW")
	getProfileSection = doGetProcAddress(libkernel32, "GetProfileSectionW")
	getProfileString = doGetProcAddress(libkernel32, "GetProfileStringW")
	getShortPathName = doGetProcAddress(libkernel32, "GetShortPathNameW")
	getStartupInfo = doGetProcAddress(libkernel32, "GetStartupInfoW")
	getStdHandle = doGetProcAddress(libkernel32, "GetStdHandle")
	getStringScripts = doGetProcAddress(libkernel32, "GetStringScripts")
	getSystemDefaultLCID = doGetProcAddress(libkernel32, "GetSystemDefaultLCID")
	getSystemDefaultLangID = doGetProcAddress(libkernel32, "GetSystemDefaultLangID")
	getSystemDefaultLocaleName = doGetProcAddress(libkernel32, "GetSystemDefaultLocaleName")
	getSystemDefaultUILanguage = doGetProcAddress(libkernel32, "GetSystemDefaultUILanguage")
	getSystemDirectory = doGetProcAddress(libkernel32, "GetSystemDirectoryW")
	getSystemFirmwareTable = doGetProcAddress(libkernel32, "GetSystemFirmwareTable")
	getSystemInfo = doGetProcAddress(libkernel32, "GetSystemInfo")
	getSystemRegistryQuota = doGetProcAddress(libkernel32, "GetSystemRegistryQuota")
	getSystemTime = doGetProcAddress(libkernel32, "GetSystemTime")
	getSystemTimeAdjustment = doGetProcAddress(libkernel32, "GetSystemTimeAdjustment")
	getSystemTimeAsFileTime = doGetProcAddress(libkernel32, "GetSystemTimeAsFileTime")
	getSystemTimePreciseAsFileTime = doGetProcAddress(libkernel32, "GetSystemTimePreciseAsFileTime")
	getSystemTimes = doGetProcAddress(libkernel32, "GetSystemTimes")
	getSystemWindowsDirectory = doGetProcAddress(libkernel32, "GetSystemWindowsDirectoryW")
	getSystemWow64Directory = doGetProcAddress(libkernel32, "GetSystemWow64DirectoryW")
	getTapeParameters = doGetProcAddress(libkernel32, "GetTapeParameters")
	getTapePosition = doGetProcAddress(libkernel32, "GetTapePosition")
	getTapeStatus = doGetProcAddress(libkernel32, "GetTapeStatus")
	getTempFileName = doGetProcAddress(libkernel32, "GetTempFileNameW")
	getTempPath = doGetProcAddress(libkernel32, "GetTempPathW")
	getThreadErrorMode = doGetProcAddress(libkernel32, "GetThreadErrorMode")
	getThreadIOPendingFlag = doGetProcAddress(libkernel32, "GetThreadIOPendingFlag")
	getThreadId = doGetProcAddress(libkernel32, "GetThreadId")
	getThreadLocale = doGetProcAddress(libkernel32, "GetThreadLocale")
	getThreadPriority = doGetProcAddress(libkernel32, "GetThreadPriority")
	getThreadPriorityBoost = doGetProcAddress(libkernel32, "GetThreadPriorityBoost")
	getThreadTimes = doGetProcAddress(libkernel32, "GetThreadTimes")
	getThreadUILanguage = doGetProcAddress(libkernel32, "GetThreadUILanguage")
	getTickCount = doGetProcAddress(libkernel32, "GetTickCount")
	getTickCount64 = doGetProcAddress(libkernel32, "GetTickCount64")
	getTimeFormatEx = doGetProcAddress(libkernel32, "GetTimeFormatEx")
	getTimeFormat = doGetProcAddress(libkernel32, "GetTimeFormatW")
	getUserDefaultLCID = doGetProcAddress(libkernel32, "GetUserDefaultLCID")
	getUserDefaultLangID = doGetProcAddress(libkernel32, "GetUserDefaultLangID")
	getUserDefaultLocaleName = doGetProcAddress(libkernel32, "GetUserDefaultLocaleName")
	getUserDefaultUILanguage = doGetProcAddress(libkernel32, "GetUserDefaultUILanguage")
	getVersion = doGetProcAddress(libkernel32, "GetVersion")
	getVolumeInformationByHandleW = doGetProcAddress(libkernel32, "GetVolumeInformationByHandleW")
	getVolumeInformation = doGetProcAddress(libkernel32, "GetVolumeInformationW")
	getVolumeNameForVolumeMountPoint = doGetProcAddress(libkernel32, "GetVolumeNameForVolumeMountPointW")
	getVolumePathName = doGetProcAddress(libkernel32, "GetVolumePathNameW")
	getWindowsDirectory = doGetProcAddress(libkernel32, "GetWindowsDirectoryW")
	getWriteWatch = doGetProcAddress(libkernel32, "GetWriteWatch")
	globalAddAtom = doGetProcAddress(libkernel32, "GlobalAddAtomW")
	globalAlloc = doGetProcAddress(libkernel32, "GlobalAlloc")
	globalCompact = doGetProcAddress(libkernel32, "GlobalCompact")
	globalDeleteAtom = doGetProcAddress(libkernel32, "GlobalDeleteAtom")
	globalFindAtom = doGetProcAddress(libkernel32, "GlobalFindAtomW")
	globalFix = doGetProcAddress(libkernel32, "GlobalFix")
	globalFlags = doGetProcAddress(libkernel32, "GlobalFlags")
	globalFree = doGetProcAddress(libkernel32, "GlobalFree")
	globalGetAtomName = doGetProcAddress(libkernel32, "GlobalGetAtomNameW")
	globalHandle = doGetProcAddress(libkernel32, "GlobalHandle")
	globalLock = doGetProcAddress(libkernel32, "GlobalLock")
	globalReAlloc = doGetProcAddress(libkernel32, "GlobalReAlloc")
	globalSize = doGetProcAddress(libkernel32, "GlobalSize")
	globalUnWire = doGetProcAddress(libkernel32, "GlobalUnWire")
	globalUnfix = doGetProcAddress(libkernel32, "GlobalUnfix")
	globalUnlock = doGetProcAddress(libkernel32, "GlobalUnlock")
	globalWire = doGetProcAddress(libkernel32, "GlobalWire")
	heapAlloc = doGetProcAddress(libkernel32, "HeapAlloc")
	heapCompact = doGetProcAddress(libkernel32, "HeapCompact")
	heapCreate = doGetProcAddress(libkernel32, "HeapCreate")
	heapDestroy = doGetProcAddress(libkernel32, "HeapDestroy")
	heapFree = doGetProcAddress(libkernel32, "HeapFree")
	heapLock = doGetProcAddress(libkernel32, "HeapLock")
	heapReAlloc = doGetProcAddress(libkernel32, "HeapReAlloc")
	heapSize = doGetProcAddress(libkernel32, "HeapSize")
	heapUnlock = doGetProcAddress(libkernel32, "HeapUnlock")
	heapValidate = doGetProcAddress(libkernel32, "HeapValidate")
	idnToAscii = doGetProcAddress(libkernel32, "IdnToAscii")
	idnToNameprepUnicode = doGetProcAddress(libkernel32, "IdnToNameprepUnicode")
	idnToUnicode = doGetProcAddress(libkernel32, "IdnToUnicode")
	initAtomTable = doGetProcAddress(libkernel32, "InitAtomTable")
	initializeSRWLock = doGetProcAddress(libkernel32, "InitializeSRWLock")
	isBadCodePtr = doGetProcAddress(libkernel32, "IsBadCodePtr")
	isBadHugeReadPtr = doGetProcAddress(libkernel32, "IsBadHugeReadPtr")
	isBadHugeWritePtr = doGetProcAddress(libkernel32, "IsBadHugeWritePtr")
	isBadReadPtr = doGetProcAddress(libkernel32, "IsBadReadPtr")
	isBadStringPtr = doGetProcAddress(libkernel32, "IsBadStringPtrW")
	isBadWritePtr = doGetProcAddress(libkernel32, "IsBadWritePtr")
	isDBCSLeadByte = doGetProcAddress(libkernel32, "IsDBCSLeadByte")
	isDBCSLeadByteEx = doGetProcAddress(libkernel32, "IsDBCSLeadByteEx")
	isDebuggerPresent = doGetProcAddress(libkernel32, "IsDebuggerPresent")
	isProcessInJob = doGetProcAddress(libkernel32, "IsProcessInJob")
	isProcessorFeaturePresent = doGetProcAddress(libkernel32, "IsProcessorFeaturePresent")
	isSystemResumeAutomatic = doGetProcAddress(libkernel32, "IsSystemResumeAutomatic")
	isThreadAFiber = doGetProcAddress(libkernel32, "IsThreadAFiber")
	isValidCodePage = doGetProcAddress(libkernel32, "IsValidCodePage")
	isValidLocale = doGetProcAddress(libkernel32, "IsValidLocale")
	isValidLocaleName = doGetProcAddress(libkernel32, "IsValidLocaleName")
	isWow64Process = doGetProcAddress(libkernel32, "IsWow64Process")
	lCIDToLocaleName = doGetProcAddress(libkernel32, "LCIDToLocaleName")
	lCMapString = doGetProcAddress(libkernel32, "LCMapStringW")
	lZClose = doGetProcAddress(libkernel32, "LZClose")
	lZCopy = doGetProcAddress(libkernel32, "LZCopy")
	lZDone = doGetProcAddress(libkernel32, "LZDone")
	lZInit = doGetProcAddress(libkernel32, "LZInit")
	lZRead = doGetProcAddress(libkernel32, "LZRead")
	lZSeek = doGetProcAddress(libkernel32, "LZSeek")
	lZStart = doGetProcAddress(libkernel32, "LZStart")
	loadLibraryEx = doGetProcAddress(libkernel32, "LoadLibraryExW")
	loadLibrary = doGetProcAddress(libkernel32, "LoadLibraryW")
	loadModule = doGetProcAddress(libkernel32, "LoadModule")
	loadPackagedLibrary = doGetProcAddress(libkernel32, "LoadPackagedLibrary")
	loadResource = doGetProcAddress(libkernel32, "LoadResource")
	localAlloc = doGetProcAddress(libkernel32, "LocalAlloc")
	localCompact = doGetProcAddress(libkernel32, "LocalCompact")
	localFileTimeToFileTime = doGetProcAddress(libkernel32, "LocalFileTimeToFileTime")
	localFlags = doGetProcAddress(libkernel32, "LocalFlags")
	localFree = doGetProcAddress(libkernel32, "LocalFree")
	localHandle = doGetProcAddress(libkernel32, "LocalHandle")
	localLock = doGetProcAddress(libkernel32, "LocalLock")
	localReAlloc = doGetProcAddress(libkernel32, "LocalReAlloc")
	localShrink = doGetProcAddress(libkernel32, "LocalShrink")
	localSize = doGetProcAddress(libkernel32, "LocalSize")
	localUnlock = doGetProcAddress(libkernel32, "LocalUnlock")
	localeNameToLCID = doGetProcAddress(libkernel32, "LocaleNameToLCID")
	lockFile = doGetProcAddress(libkernel32, "LockFile")
	lockFileEx = doGetProcAddress(libkernel32, "LockFileEx")
	lockResource = doGetProcAddress(libkernel32, "LockResource")
	mapViewOfFile = doGetProcAddress(libkernel32, "MapViewOfFile")
	mapViewOfFileEx = doGetProcAddress(libkernel32, "MapViewOfFileEx")
	mapViewOfFileExNuma = doGetProcAddress(libkernel32, "MapViewOfFileExNuma")
	moveFileEx = doGetProcAddress(libkernel32, "MoveFileExW")
	moveFile = doGetProcAddress(libkernel32, "MoveFileW")
	mulDiv = doGetProcAddress(libkernel32, "MulDiv")
	needCurrentDirectoryForExePath = doGetProcAddress(libkernel32, "NeedCurrentDirectoryForExePathW")
	notifyUILanguageChange = doGetProcAddress(libkernel32, "NotifyUILanguageChange")
	openEvent = doGetProcAddress(libkernel32, "OpenEventW")
	openFileMapping = doGetProcAddress(libkernel32, "OpenFileMappingW")
	openJobObject = doGetProcAddress(libkernel32, "OpenJobObjectW")
	openMutex = doGetProcAddress(libkernel32, "OpenMutexW")
	openPrivateNamespace = doGetProcAddress(libkernel32, "OpenPrivateNamespaceW")
	openProcess = doGetProcAddress(libkernel32, "OpenProcess")
	openSemaphore = doGetProcAddress(libkernel32, "OpenSemaphoreW")
	openThread = doGetProcAddress(libkernel32, "OpenThread")
	openWaitableTimer = doGetProcAddress(libkernel32, "OpenWaitableTimerW")
	outputDebugString = doGetProcAddress(libkernel32, "OutputDebugStringW")
	peekNamedPipe = doGetProcAddress(libkernel32, "PeekNamedPipe")
	postQueuedCompletionStatus = doGetProcAddress(libkernel32, "PostQueuedCompletionStatus")
	prepareTape = doGetProcAddress(libkernel32, "PrepareTape")
	processIdToSessionId = doGetProcAddress(libkernel32, "ProcessIdToSessionId")
	pulseEvent = doGetProcAddress(libkernel32, "PulseEvent")
	purgeComm = doGetProcAddress(libkernel32, "PurgeComm")
	queryActCtxSettingsW = doGetProcAddress(libkernel32, "QueryActCtxSettingsW")
	queryActCtxW = doGetProcAddress(libkernel32, "QueryActCtxW")
	queryDosDevice = doGetProcAddress(libkernel32, "QueryDosDeviceW")
	queryFullProcessImageName = doGetProcAddress(libkernel32, "QueryFullProcessImageNameW")
	queryIdleProcessorCycleTime = doGetProcAddress(libkernel32, "QueryIdleProcessorCycleTime")
	queryIdleProcessorCycleTimeEx = doGetProcAddress(libkernel32, "QueryIdleProcessorCycleTimeEx")
	queryMemoryResourceNotification = doGetProcAddress(libkernel32, "QueryMemoryResourceNotification")
	queryPerformanceCounter = doGetProcAddress(libkernel32, "QueryPerformanceCounter")
	queryPerformanceFrequency = doGetProcAddress(libkernel32, "QueryPerformanceFrequency")
	queryProcessAffinityUpdateMode = doGetProcAddress(libkernel32, "QueryProcessAffinityUpdateMode")
	queryProcessCycleTime = doGetProcAddress(libkernel32, "QueryProcessCycleTime")
	queryThreadCycleTime = doGetProcAddress(libkernel32, "QueryThreadCycleTime")
	queueUserWorkItem = doGetProcAddress(libkernel32, "QueueUserWorkItem")
	raiseException = doGetProcAddress(libkernel32, "RaiseException")
	reOpenFile = doGetProcAddress(libkernel32, "ReOpenFile")
	readConsoleOutputAttribute = doGetProcAddress(libkernel32, "ReadConsoleOutputAttribute")
	readConsoleOutputCharacter = doGetProcAddress(libkernel32, "ReadConsoleOutputCharacterW")
	readConsole = doGetProcAddress(libkernel32, "ReadConsoleW")
	readFile = doGetProcAddress(libkernel32, "ReadFile")
	readProcessMemory = doGetProcAddress(libkernel32, "ReadProcessMemory")
	reclaimVirtualMemory = doGetProcAddress(libkernel32, "ReclaimVirtualMemory")
	registerApplicationRestart = doGetProcAddress(libkernel32, "RegisterApplicationRestart")
	releaseActCtx = doGetProcAddress(libkernel32, "ReleaseActCtx")
	releaseMutex = doGetProcAddress(libkernel32, "ReleaseMutex")
	releaseSRWLockExclusive = doGetProcAddress(libkernel32, "ReleaseSRWLockExclusive")
	releaseSRWLockShared = doGetProcAddress(libkernel32, "ReleaseSRWLockShared")
	releaseSemaphore = doGetProcAddress(libkernel32, "ReleaseSemaphore")
	removeDirectoryTransacted = doGetProcAddress(libkernel32, "RemoveDirectoryTransactedW")
	removeDirectory = doGetProcAddress(libkernel32, "RemoveDirectoryW")
	removeSecureMemoryCacheCallback = doGetProcAddress(libkernel32, "RemoveSecureMemoryCacheCallback")
	removeVectoredContinueHandler = doGetProcAddress(libkernel32, "RemoveVectoredContinueHandler")
	removeVectoredExceptionHandler = doGetProcAddress(libkernel32, "RemoveVectoredExceptionHandler")
	replaceFile = doGetProcAddress(libkernel32, "ReplaceFileW")
	replacePartitionUnit = doGetProcAddress(libkernel32, "ReplacePartitionUnit")
	requestDeviceWakeup = doGetProcAddress(libkernel32, "RequestDeviceWakeup")
	resetEvent = doGetProcAddress(libkernel32, "ResetEvent")
	resetWriteWatch = doGetProcAddress(libkernel32, "ResetWriteWatch")
	resolveLocaleName = doGetProcAddress(libkernel32, "ResolveLocaleName")
	restoreLastError = doGetProcAddress(libkernel32, "RestoreLastError")
	resumeThread = doGetProcAddress(libkernel32, "ResumeThread")
	searchPath = doGetProcAddress(libkernel32, "SearchPathW")
	setCommBreak = doGetProcAddress(libkernel32, "SetCommBreak")
	setCommMask = doGetProcAddress(libkernel32, "SetCommMask")
	setComputerName = doGetProcAddress(libkernel32, "SetComputerNameW")
	setConsoleActiveScreenBuffer = doGetProcAddress(libkernel32, "SetConsoleActiveScreenBuffer")
	setConsoleCP = doGetProcAddress(libkernel32, "SetConsoleCP")
	setConsoleCursorPosition = doGetProcAddress(libkernel32, "SetConsoleCursorPosition")
	setConsoleMode = doGetProcAddress(libkernel32, "SetConsoleMode")
	setConsoleOutputCP = doGetProcAddress(libkernel32, "SetConsoleOutputCP")
	setConsoleScreenBufferSize = doGetProcAddress(libkernel32, "SetConsoleScreenBufferSize")
	setConsoleTextAttribute = doGetProcAddress(libkernel32, "SetConsoleTextAttribute")
	setConsoleTitle = doGetProcAddress(libkernel32, "SetConsoleTitleW")
	setCurrentDirectory = doGetProcAddress(libkernel32, "SetCurrentDirectoryW")
	setDllDirectory = doGetProcAddress(libkernel32, "SetDllDirectoryW")
	setEndOfFile = doGetProcAddress(libkernel32, "SetEndOfFile")
	setEnvironmentVariable = doGetProcAddress(libkernel32, "SetEnvironmentVariableW")
	setErrorMode = doGetProcAddress(libkernel32, "SetErrorMode")
	setEvent = doGetProcAddress(libkernel32, "SetEvent")
	setFileApisToANSI = doGetProcAddress(libkernel32, "SetFileApisToANSI")
	setFileApisToOEM = doGetProcAddress(libkernel32, "SetFileApisToOEM")
	setFileAttributesTransacted = doGetProcAddress(libkernel32, "SetFileAttributesTransactedW")
	setFileAttributes = doGetProcAddress(libkernel32, "SetFileAttributesW")
	setFileBandwidthReservation = doGetProcAddress(libkernel32, "SetFileBandwidthReservation")
	setFileCompletionNotificationModes = doGetProcAddress(libkernel32, "SetFileCompletionNotificationModes")
	setFilePointer = doGetProcAddress(libkernel32, "SetFilePointer")
	setFileShortName = doGetProcAddress(libkernel32, "SetFileShortNameW")
	setFileTime = doGetProcAddress(libkernel32, "SetFileTime")
	setFileValidData = doGetProcAddress(libkernel32, "SetFileValidData")
	setFirmwareEnvironmentVariable = doGetProcAddress(libkernel32, "SetFirmwareEnvironmentVariableW")
	setHandleCount = doGetProcAddress(libkernel32, "SetHandleCount")
	setHandleInformation = doGetProcAddress(libkernel32, "SetHandleInformation")
	setLastError = doGetProcAddress(libkernel32, "SetLastError")
	setLocalTime = doGetProcAddress(libkernel32, "SetLocalTime")
	setLocaleInfo = doGetProcAddress(libkernel32, "SetLocaleInfoW")
	setMailslotInfo = doGetProcAddress(libkernel32, "SetMailslotInfo")
	setMessageWaitingIndicator = doGetProcAddress(libkernel32, "SetMessageWaitingIndicator")
	setNamedPipeHandleState = doGetProcAddress(libkernel32, "SetNamedPipeHandleState")
	setPriorityClass = doGetProcAddress(libkernel32, "SetPriorityClass")
	setProcessAffinityMask = doGetProcAddress(libkernel32, "SetProcessAffinityMask")
	setProcessAffinityUpdateMode = doGetProcAddress(libkernel32, "SetProcessAffinityUpdateMode")
	setProcessDEPPolicy = doGetProcAddress(libkernel32, "SetProcessDEPPolicy")
	setProcessPriorityBoost = doGetProcAddress(libkernel32, "SetProcessPriorityBoost")
	setProcessShutdownParameters = doGetProcAddress(libkernel32, "SetProcessShutdownParameters")
	setProcessWorkingSetSize = doGetProcAddress(libkernel32, "SetProcessWorkingSetSize")
	setProcessWorkingSetSizeEx = doGetProcAddress(libkernel32, "SetProcessWorkingSetSizeEx")
	setSearchPathMode = doGetProcAddress(libkernel32, "SetSearchPathMode")
	setStdHandle = doGetProcAddress(libkernel32, "SetStdHandle")
	setStdHandleEx = doGetProcAddress(libkernel32, "SetStdHandleEx")
	setSystemFileCacheSize = doGetProcAddress(libkernel32, "SetSystemFileCacheSize")
	setSystemPowerState = doGetProcAddress(libkernel32, "SetSystemPowerState")
	setSystemTime = doGetProcAddress(libkernel32, "SetSystemTime")
	setSystemTimeAdjustment = doGetProcAddress(libkernel32, "SetSystemTimeAdjustment")
	setTapeParameters = doGetProcAddress(libkernel32, "SetTapeParameters")
	setTapePosition = doGetProcAddress(libkernel32, "SetTapePosition")
	setThreadAffinityMask = doGetProcAddress(libkernel32, "SetThreadAffinityMask")
	setThreadErrorMode = doGetProcAddress(libkernel32, "SetThreadErrorMode")
	setThreadIdealProcessor = doGetProcAddress(libkernel32, "SetThreadIdealProcessor")
	setThreadLocale = doGetProcAddress(libkernel32, "SetThreadLocale")
	setThreadPriority = doGetProcAddress(libkernel32, "SetThreadPriority")
	setThreadPriorityBoost = doGetProcAddress(libkernel32, "SetThreadPriorityBoost")
	setThreadStackGuarantee = doGetProcAddress(libkernel32, "SetThreadStackGuarantee")
	setThreadUILanguage = doGetProcAddress(libkernel32, "SetThreadUILanguage")
	setUserGeoID = doGetProcAddress(libkernel32, "SetUserGeoID")
	setVolumeLabel = doGetProcAddress(libkernel32, "SetVolumeLabelW")
	setVolumeMountPoint = doGetProcAddress(libkernel32, "SetVolumeMountPointW")
	setupComm = doGetProcAddress(libkernel32, "SetupComm")
	signalObjectAndWait = doGetProcAddress(libkernel32, "SignalObjectAndWait")
	sizeofResource = doGetProcAddress(libkernel32, "SizeofResource")
	sleep = doGetProcAddress(libkernel32, "Sleep")
	sleepEx = doGetProcAddress(libkernel32, "SleepEx")
	suspendThread = doGetProcAddress(libkernel32, "SuspendThread")
	switchToFiber = doGetProcAddress(libkernel32, "SwitchToFiber")
	switchToThread = doGetProcAddress(libkernel32, "SwitchToThread")
	systemTimeToFileTime = doGetProcAddress(libkernel32, "SystemTimeToFileTime")
	terminateJobObject = doGetProcAddress(libkernel32, "TerminateJobObject")
	terminateProcess = doGetProcAddress(libkernel32, "TerminateProcess")
	terminateThread = doGetProcAddress(libkernel32, "TerminateThread")
	tlsAlloc = doGetProcAddress(libkernel32, "TlsAlloc")
	tlsFree = doGetProcAddress(libkernel32, "TlsFree")
	tlsGetValue = doGetProcAddress(libkernel32, "TlsGetValue")
	tlsSetValue = doGetProcAddress(libkernel32, "TlsSetValue")
	toolhelp32ReadProcessMemory = doGetProcAddress(libkernel32, "Toolhelp32ReadProcessMemory")
	transactNamedPipe = doGetProcAddress(libkernel32, "TransactNamedPipe")
	transmitCommChar = doGetProcAddress(libkernel32, "TransmitCommChar")
	tryAcquireSRWLockExclusive = doGetProcAddress(libkernel32, "TryAcquireSRWLockExclusive")
	tryAcquireSRWLockShared = doGetProcAddress(libkernel32, "TryAcquireSRWLockShared")
	unlockFile = doGetProcAddress(libkernel32, "UnlockFile")
	unlockFileEx = doGetProcAddress(libkernel32, "UnlockFileEx")
	unmapViewOfFile = doGetProcAddress(libkernel32, "UnmapViewOfFile")
	unregisterApplicationRecoveryCallback = doGetProcAddress(libkernel32, "UnregisterApplicationRecoveryCallback")
	unregisterApplicationRestart = doGetProcAddress(libkernel32, "UnregisterApplicationRestart")
	unregisterWait = doGetProcAddress(libkernel32, "UnregisterWait")
	unregisterWaitEx = doGetProcAddress(libkernel32, "UnregisterWaitEx")
	updateResource = doGetProcAddress(libkernel32, "UpdateResourceW")
	verLanguageName = doGetProcAddress(libkernel32, "VerLanguageNameW")
	verifyScripts = doGetProcAddress(libkernel32, "VerifyScripts")
	virtualAlloc = doGetProcAddress(libkernel32, "VirtualAlloc")
	virtualAllocEx = doGetProcAddress(libkernel32, "VirtualAllocEx")
	virtualAllocExNuma = doGetProcAddress(libkernel32, "VirtualAllocExNuma")
	virtualFree = doGetProcAddress(libkernel32, "VirtualFree")
	virtualFreeEx = doGetProcAddress(libkernel32, "VirtualFreeEx")
	virtualLock = doGetProcAddress(libkernel32, "VirtualLock")
	virtualProtect = doGetProcAddress(libkernel32, "VirtualProtect")
	virtualProtectEx = doGetProcAddress(libkernel32, "VirtualProtectEx")
	virtualUnlock = doGetProcAddress(libkernel32, "VirtualUnlock")
	wTSGetActiveConsoleSessionId = doGetProcAddress(libkernel32, "WTSGetActiveConsoleSessionId")
	waitCommEvent = doGetProcAddress(libkernel32, "WaitCommEvent")
	waitForMultipleObjects = doGetProcAddress(libkernel32, "WaitForMultipleObjects")
	waitForMultipleObjectsEx = doGetProcAddress(libkernel32, "WaitForMultipleObjectsEx")
	waitForSingleObject = doGetProcAddress(libkernel32, "WaitForSingleObject")
	waitForSingleObjectEx = doGetProcAddress(libkernel32, "WaitForSingleObjectEx")
	waitNamedPipe = doGetProcAddress(libkernel32, "WaitNamedPipeW")
	werGetFlags = doGetProcAddress(libkernel32, "WerGetFlags")
	werRegisterMemoryBlock = doGetProcAddress(libkernel32, "WerRegisterMemoryBlock")
	werRegisterRuntimeExceptionModule = doGetProcAddress(libkernel32, "WerRegisterRuntimeExceptionModule")
	werSetFlags = doGetProcAddress(libkernel32, "WerSetFlags")
	werUnregisterFile = doGetProcAddress(libkernel32, "WerUnregisterFile")
	werUnregisterMemoryBlock = doGetProcAddress(libkernel32, "WerUnregisterMemoryBlock")
	werUnregisterRuntimeExceptionModule = doGetProcAddress(libkernel32, "WerUnregisterRuntimeExceptionModule")
	winExec = doGetProcAddress(libkernel32, "WinExec")
	wow64DisableWow64FsRedirection = doGetProcAddress(libkernel32, "Wow64DisableWow64FsRedirection")
	wow64EnableWow64FsRedirection = doGetProcAddress(libkernel32, "Wow64EnableWow64FsRedirection")
	wow64RevertWow64FsRedirection = doGetProcAddress(libkernel32, "Wow64RevertWow64FsRedirection")
	wow64SuspendThread = doGetProcAddress(libkernel32, "Wow64SuspendThread")
	writeConsoleOutputAttribute = doGetProcAddress(libkernel32, "WriteConsoleOutputAttribute")
	writeConsoleOutputCharacter = doGetProcAddress(libkernel32, "WriteConsoleOutputCharacterW")
	writeConsole = doGetProcAddress(libkernel32, "WriteConsoleW")
	writeFile = doGetProcAddress(libkernel32, "WriteFile")
	writePrivateProfileSection = doGetProcAddress(libkernel32, "WritePrivateProfileSectionW")
	writePrivateProfileString = doGetProcAddress(libkernel32, "WritePrivateProfileStringW")
	writePrivateProfileStruct = doGetProcAddress(libkernel32, "WritePrivateProfileStructW")
	writeProcessMemory = doGetProcAddress(libkernel32, "WriteProcessMemory")
	writeProfileSection = doGetProcAddress(libkernel32, "WriteProfileSectionW")
	writeProfileString = doGetProcAddress(libkernel32, "WriteProfileStringW")
	writeTapemark = doGetProcAddress(libkernel32, "WriteTapemark")
	zombifyActCtx = doGetProcAddress(libkernel32, "ZombifyActCtx")
	lstrcat = doGetProcAddress(libkernel32, "lstrcatW")
	lstrcmp = doGetProcAddress(libkernel32, "lstrcmpW")
	lstrcmpi = doGetProcAddress(libkernel32, "lstrcmpiW")
	lstrcpy = doGetProcAddress(libkernel32, "lstrcpyW")
	lstrcpyn = doGetProcAddress(libkernel32, "lstrcpynW")
	lstrlen = doGetProcAddress(libkernel32, "lstrlenW")
	closeConsoleHandle = doGetProcAddress(libkernel32, "CloseConsoleHandle")
	closeProfileUserMapping = doGetProcAddress(libkernel32, "CloseProfileUserMapping")
	cmdBatNotification = doGetProcAddress(libkernel32, "CmdBatNotification")
	delayLoadFailureHook = doGetProcAddress(libkernel32, "DelayLoadFailureHook")
	duplicateConsoleHandle = doGetProcAddress(libkernel32, "DuplicateConsoleHandle")
	expungeConsoleCommandHistory = doGetProcAddress(libkernel32, "ExpungeConsoleCommandHistoryW")
	getConsoleCommandHistoryLength = doGetProcAddress(libkernel32, "GetConsoleCommandHistoryLengthW")
	getConsoleCommandHistory = doGetProcAddress(libkernel32, "GetConsoleCommandHistoryW")
	getConsoleInputExeName = doGetProcAddress(libkernel32, "GetConsoleInputExeNameW")
	getConsoleInputWaitHandle = doGetProcAddress(libkernel32, "GetConsoleInputWaitHandle")
	getConsoleKeyboardLayoutName = doGetProcAddress(libkernel32, "GetConsoleKeyboardLayoutNameW")
	getNumberOfConsoleFonts = doGetProcAddress(libkernel32, "GetNumberOfConsoleFonts")
	k32EmptyWorkingSet = doGetProcAddress(libkernel32, "K32EmptyWorkingSet")
	k32EnumDeviceDrivers = doGetProcAddress(libkernel32, "K32EnumDeviceDrivers")
	k32EnumPageFiles = doGetProcAddress(libkernel32, "K32EnumPageFilesW")
	k32EnumProcessModules = doGetProcAddress(libkernel32, "K32EnumProcessModules")
	k32EnumProcessModulesEx = doGetProcAddress(libkernel32, "K32EnumProcessModulesEx")
	k32EnumProcesses = doGetProcAddress(libkernel32, "K32EnumProcesses")
	k32GetDeviceDriverBaseName = doGetProcAddress(libkernel32, "K32GetDeviceDriverBaseNameW")
	k32GetDeviceDriverFileName = doGetProcAddress(libkernel32, "K32GetDeviceDriverFileNameW")
	k32GetMappedFileName = doGetProcAddress(libkernel32, "K32GetMappedFileNameW")
	k32GetModuleBaseName = doGetProcAddress(libkernel32, "K32GetModuleBaseNameW")
	k32GetModuleFileNameEx = doGetProcAddress(libkernel32, "K32GetModuleFileNameExW")
	k32GetModuleInformation = doGetProcAddress(libkernel32, "K32GetModuleInformation")
	k32GetProcessImageFileName = doGetProcAddress(libkernel32, "K32GetProcessImageFileNameW")
	k32GetProcessMemoryInfo = doGetProcAddress(libkernel32, "K32GetProcessMemoryInfo")
	k32GetWsChanges = doGetProcAddress(libkernel32, "K32GetWsChanges")
	k32InitializeProcessForWsWatch = doGetProcAddress(libkernel32, "K32InitializeProcessForWsWatch")
	k32QueryWorkingSet = doGetProcAddress(libkernel32, "K32QueryWorkingSet")
	k32QueryWorkingSetEx = doGetProcAddress(libkernel32, "K32QueryWorkingSetEx")
	openConsoleW = doGetProcAddress(libkernel32, "OpenConsoleW")
	openProfileUserMapping = doGetProcAddress(libkernel32, "OpenProfileUserMapping")
	rtlCaptureStackBackTrace = doGetProcAddress(libkernel32, "RtlCaptureStackBackTrace")
	rtlCompareMemory = doGetProcAddress(libkernel32, "RtlCompareMemory")
	rtlCopyMemory = doGetProcAddress(libkernel32, "RtlCopyMemory")
	rtlFillMemory = doGetProcAddress(libkernel32, "RtlFillMemory")
	rtlMoveMemory = doGetProcAddress(libkernel32, "RtlMoveMemory")
	rtlPcToFileHeader = doGetProcAddress(libkernel32, "RtlPcToFileHeader")
	rtlZeroMemory = doGetProcAddress(libkernel32, "RtlZeroMemory")
	setCPGlobal = doGetProcAddress(libkernel32, "SetCPGlobal")
	setConsoleFont = doGetProcAddress(libkernel32, "SetConsoleFont")
	setConsoleIcon = doGetProcAddress(libkernel32, "SetConsoleIcon")
	setConsoleInputExeName = doGetProcAddress(libkernel32, "SetConsoleInputExeNameW")
	setConsoleKeyShortcuts = doGetProcAddress(libkernel32, "SetConsoleKeyShortcuts")
	setTermsrvAppInstallMode = doGetProcAddress(libkernel32, "SetTermsrvAppInstallMode")
	termsrvAppInstallMode = doGetProcAddress(libkernel32, "TermsrvAppInstallMode")
	uTRegister = doGetProcAddress(libkernel32, "UTRegister")
	uTUnRegister = doGetProcAddress(libkernel32, "UTUnRegister")
	verSetConditionMask = doGetProcAddress(libkernel32, "VerSetConditionMask")
	verifyConsoleIoHandle = doGetProcAddress(libkernel32, "VerifyConsoleIoHandle")
}

func AcquireSRWLockExclusive(sRWLock PSRWLOCK) {
	syscall3(acquireSRWLockExclusive, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
}

func AcquireSRWLockShared(sRWLock PSRWLOCK) {
	syscall3(acquireSRWLockShared, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
}

func ActivateActCtx(hActCtx HANDLE, lpCookie *ULONG_PTR) bool {
	ret1 := syscall3(activateActCtx, 2,
		uintptr(hActCtx),
		uintptr(unsafe.Pointer(lpCookie)),
		0)
	return ret1 != 0
}

func AddAtom(lpString string) ATOM {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(addAtom, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return ATOM(ret1)
}

func AddConsoleAlias(source LPWSTR, target LPWSTR, exeName LPWSTR) bool {
	ret1 := syscall3(addConsoleAlias, 3,
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(target)),
		uintptr(unsafe.Pointer(exeName)))
	return ret1 != 0
}

func AddIntegrityLabelToBoundaryDescriptor(boundaryDescriptor *HANDLE, integrityLabel PSID) bool {
	ret1 := syscall3(addIntegrityLabelToBoundaryDescriptor, 2,
		uintptr(unsafe.Pointer(boundaryDescriptor)),
		uintptr(integrityLabel),
		0)
	return ret1 != 0
}

func AddRefActCtx(hActCtx HANDLE) {
	syscall3(addRefActCtx, 1,
		uintptr(hActCtx),
		0,
		0)
}

func AddSIDToBoundaryDescriptor(boundaryDescriptor *HANDLE, requiredSid PSID) bool {
	ret1 := syscall3(addSIDToBoundaryDescriptor, 2,
		uintptr(unsafe.Pointer(boundaryDescriptor)),
		uintptr(requiredSid),
		0)
	return ret1 != 0
}

func AddSecureMemoryCacheCallback(pfnCallBack PSECURE_MEMORY_CACHE_CALLBACK) bool {
	pfnCallBackCallback := syscall.NewCallback(func(AddrRawArg PVOID, RangeRawArg SIZE_T) uintptr {
		ret := pfnCallBack(AddrRawArg, RangeRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(addSecureMemoryCacheCallback, 1,
		pfnCallBackCallback,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PVECTORED_EXCEPTION_HANDLER
// func AddVectoredContinueHandler(first ULONG, handler PVECTORED_EXCEPTION_HANDLER) uintptr

// TODO: Unknown type(s): PVECTORED_EXCEPTION_HANDLER
// func AddVectoredExceptionHandler(first ULONG, handler PVECTORED_EXCEPTION_HANDLER) uintptr

func AllocConsole() bool {
	ret1 := syscall3(allocConsole, 0,
		0,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PULONG_PTR
// func AllocateUserPhysicalPages(hProcess HANDLE, numberOfPages PULONG_PTR, pageArray PULONG_PTR) bool

// TODO: Unknown type(s): PULONG_PTR
// func AllocateUserPhysicalPagesNuma(hProcess HANDLE, numberOfPages PULONG_PTR, pageArray PULONG_PTR, nndPreferred uint32) bool

func ApplicationRecoveryFinished(bSuccess bool) {
	syscall3(applicationRecoveryFinished, 1,
		getUintptrFromBool(bSuccess),
		0,
		0)
}

func ApplicationRecoveryInProgress(pbCancelled *BOOL) HRESULT {
	ret1 := syscall3(applicationRecoveryInProgress, 1,
		uintptr(unsafe.Pointer(pbCancelled)),
		0,
		0)
	return HRESULT(ret1)
}

func AreFileApisANSI() bool {
	ret1 := syscall3(areFileApisANSI, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func AssignProcessToJobObject(hJob HANDLE, hProcess HANDLE) bool {
	ret1 := syscall3(assignProcessToJobObject, 2,
		uintptr(hJob),
		uintptr(hProcess),
		0)
	return ret1 != 0
}

func AttachConsole(dwProcessId uint32) bool {
	ret1 := syscall3(attachConsole, 1,
		uintptr(dwProcessId),
		0,
		0)
	return ret1 != 0
}

func BackupRead(hFile HANDLE, lpBuffer *byte, nNumberOfBytesToRead uint32, lpNumberOfBytesRead *uint32, bAbort bool, bProcessSecurity bool, lpContext *LPVOID) bool {
	ret1 := syscall9(backupRead, 7,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nNumberOfBytesToRead),
		uintptr(unsafe.Pointer(lpNumberOfBytesRead)),
		getUintptrFromBool(bAbort),
		getUintptrFromBool(bProcessSecurity),
		uintptr(unsafe.Pointer(lpContext)),
		0,
		0)
	return ret1 != 0
}

func BackupSeek(hFile HANDLE, dwLowBytesToSeek uint32, dwHighBytesToSeek uint32, lpdwLowByteSeeked *uint32, lpdwHighByteSeeked *uint32, lpContext *LPVOID) bool {
	ret1 := syscall6(backupSeek, 6,
		uintptr(hFile),
		uintptr(dwLowBytesToSeek),
		uintptr(dwHighBytesToSeek),
		uintptr(unsafe.Pointer(lpdwLowByteSeeked)),
		uintptr(unsafe.Pointer(lpdwHighByteSeeked)),
		uintptr(unsafe.Pointer(lpContext)))
	return ret1 != 0
}

func BackupWrite(hFile HANDLE, lpBuffer *byte, nNumberOfBytesToWrite uint32, lpNumberOfBytesWritten *uint32, bAbort bool, bProcessSecurity bool, lpContext *LPVOID) bool {
	ret1 := syscall9(backupWrite, 7,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nNumberOfBytesToWrite),
		uintptr(unsafe.Pointer(lpNumberOfBytesWritten)),
		getUintptrFromBool(bAbort),
		getUintptrFromBool(bProcessSecurity),
		uintptr(unsafe.Pointer(lpContext)),
		0,
		0)
	return ret1 != 0
}

func Beep(dwFreq uint32, dwDuration uint32) bool {
	ret1 := syscall3(beep, 2,
		uintptr(dwFreq),
		uintptr(dwDuration),
		0)
	return ret1 != 0
}

func BeginUpdateResource(pFileName string, bDeleteExistingResources bool) HANDLE {
	pFileNameStr := unicode16FromString(pFileName)
	ret1 := syscall3(beginUpdateResource, 2,
		uintptr(unsafe.Pointer(&pFileNameStr[0])),
		getUintptrFromBool(bDeleteExistingResources),
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): LPOVERLAPPED_COMPLETION_ROUTINE
// func BindIoCompletionCallback(fileHandle HANDLE, function LPOVERLAPPED_COMPLETION_ROUTINE, flags ULONG) bool

// TODO: Unknown type(s): LPCOMMTIMEOUTS, LPDCB
// func BuildCommDCBAndTimeouts(lpDef string, lpDCB LPDCB, lpCommTimeouts LPCOMMTIMEOUTS) bool

// TODO: Unknown type(s): LPDCB
// func BuildCommDCB(lpDef string, lpDCB LPDCB) bool

func CallNamedPipe(lpNamedPipeName string, lpInBuffer LPVOID, nInBufferSize uint32, lpOutBuffer LPVOID, nOutBufferSize uint32, lpBytesRead *uint32, nTimeOut uint32) bool {
	lpNamedPipeNameStr := unicode16FromString(lpNamedPipeName)
	ret1 := syscall9(callNamedPipe, 7,
		uintptr(unsafe.Pointer(&lpNamedPipeNameStr[0])),
		uintptr(unsafe.Pointer(lpInBuffer)),
		uintptr(nInBufferSize),
		uintptr(unsafe.Pointer(lpOutBuffer)),
		uintptr(nOutBufferSize),
		uintptr(unsafe.Pointer(lpBytesRead)),
		uintptr(nTimeOut),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PTP_CALLBACK_INSTANCE
// func CallbackMayRunLong(pci PTP_CALLBACK_INSTANCE) bool

func CancelDeviceWakeupRequest(hDevice HANDLE) bool {
	ret1 := syscall3(cancelDeviceWakeupRequest, 1,
		uintptr(hDevice),
		0,
		0)
	return ret1 != 0
}

func CancelIo(hFile HANDLE) bool {
	ret1 := syscall3(cancelIo, 1,
		uintptr(hFile),
		0,
		0)
	return ret1 != 0
}

func CancelIoEx(hFile HANDLE, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall3(cancelIoEx, 2,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0)
	return ret1 != 0
}

func CancelSynchronousIo(hThread HANDLE) bool {
	ret1 := syscall3(cancelSynchronousIo, 1,
		uintptr(hThread),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PTP_IO
// func CancelThreadpoolIo(pio PTP_IO)

func CancelTimerQueueTimer(timerQueue HANDLE, timer HANDLE) bool {
	ret1 := syscall3(cancelTimerQueueTimer, 2,
		uintptr(timerQueue),
		uintptr(timer),
		0)
	return ret1 != 0
}

func CancelWaitableTimer(hTimer HANDLE) bool {
	ret1 := syscall3(cancelWaitableTimer, 1,
		uintptr(hTimer),
		0,
		0)
	return ret1 != 0
}

func ChangeTimerQueueTimer(timerQueue HANDLE, timer HANDLE, dueTime ULONG, period ULONG) bool {
	ret1 := syscall6(changeTimerQueueTimer, 4,
		uintptr(timerQueue),
		uintptr(timer),
		uintptr(dueTime),
		uintptr(period),
		0,
		0)
	return ret1 != 0
}

func CheckNameLegalDOS8Dot3(lpName string, lpOemName LPSTR, oemNameSize uint32, pbNameContainsSpaces *BOOL, pbNameLegal *BOOL) bool {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(checkNameLegalDOS8Dot3, 5,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(lpOemName)),
		uintptr(oemNameSize),
		uintptr(unsafe.Pointer(pbNameContainsSpaces)),
		uintptr(unsafe.Pointer(pbNameLegal)),
		0)
	return ret1 != 0
}

func CheckRemoteDebuggerPresent(hProcess HANDLE, pbDebuggerPresent *BOOL) bool {
	ret1 := syscall3(checkRemoteDebuggerPresent, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(pbDebuggerPresent)),
		0)
	return ret1 != 0
}

func ClearCommBreak(hFile HANDLE) bool {
	ret1 := syscall3(clearCommBreak, 1,
		uintptr(hFile),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCOMSTAT
// func ClearCommError(hFile HANDLE, lpErrors *uint32, lpStat LPCOMSTAT) bool

func CloseHandle(hObject HANDLE) bool {
	ret1 := syscall3(closeHandle, 1,
		uintptr(hObject),
		0,
		0)
	return ret1 != 0
}

func ClosePrivateNamespace(handle HANDLE, flags ULONG) BOOLEAN {
	ret1 := syscall3(closePrivateNamespace, 2,
		uintptr(handle),
		uintptr(flags),
		0)
	return BOOLEAN(ret1)
}

// TODO: Unknown type(s): PTP_POOL
// func CloseThreadpool(ptpp PTP_POOL)

// TODO: Unknown type(s): PTP_CLEANUP_GROUP
// func CloseThreadpoolCleanupGroup(ptpcg PTP_CLEANUP_GROUP)

// TODO: Unknown type(s): PTP_CLEANUP_GROUP
// func CloseThreadpoolCleanupGroupMembers(ptpcg PTP_CLEANUP_GROUP, fCancelPendingCallbacks bool, pvCleanupContext uintptr)

// TODO: Unknown type(s): PTP_IO
// func CloseThreadpoolIo(pio PTP_IO)

// TODO: Unknown type(s): PTP_TIMER
// func CloseThreadpoolTimer(pti PTP_TIMER)

// TODO: Unknown type(s): PTP_WAIT
// func CloseThreadpoolWait(pwa PTP_WAIT)

// TODO: Unknown type(s): PTP_WORK
// func CloseThreadpoolWork(pwk PTP_WORK)

// TODO: Unknown type(s): LPCOMMCONFIG
// func CommConfigDialog(lpszName string, hWnd HWND, lpCC LPCOMMCONFIG) bool

func CompareFileTime(lpFileTime1 /*const*/ *FILETIME, lpFileTime2 /*const*/ *FILETIME) LONG {
	ret1 := syscall3(compareFileTime, 2,
		uintptr(unsafe.Pointer(lpFileTime1)),
		uintptr(unsafe.Pointer(lpFileTime2)),
		0)
	return LONG(ret1)
}

// TODO: Unknown type(s): LPCWCH, LPNLSVERSIONINFO
// func CompareStringEx(lpLocaleName string, dwCmpFlags uint32, lpString1 LPCWCH, cchCount1 int32, lpString2 LPCWCH, cchCount2 int32, lpVersionInformation LPNLSVERSIONINFO, lpReserved LPVOID, lParam LPARAM) int32

// TODO: Unknown type(s): LPCWCH
// func CompareStringOrdinal(lpString1 LPCWCH, cchCount1 int32, lpString2 LPCWCH, cchCount2 int32, bIgnoreCase bool) int32

// TODO: Unknown type(s): PCNZWCH
// func CompareString(locale LCID, dwCmpFlags uint32, lpString1 PCNZWCH, cchCount1 int32, lpString2 PCNZWCH, cchCount2 int32) int32

func ConnectNamedPipe(hNamedPipe HANDLE, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall3(connectNamedPipe, 2,
		uintptr(hNamedPipe),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0)
	return ret1 != 0
}

func ContinueDebugEvent(dwProcessId uint32, dwThreadId uint32, dwContinueStatus uint32) bool {
	ret1 := syscall3(continueDebugEvent, 3,
		uintptr(dwProcessId),
		uintptr(dwThreadId),
		uintptr(dwContinueStatus))
	return ret1 != 0
}

func ConvertDefaultLocale(locale LCID) LCID {
	ret1 := syscall3(convertDefaultLocale, 1,
		uintptr(locale),
		0,
		0)
	return LCID(ret1)
}

func ConvertFiberToThread() bool {
	ret1 := syscall3(convertFiberToThread, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func ConvertThreadToFiber(lpParameter LPVOID) LPVOID {
	ret1 := syscall3(convertThreadToFiber, 1,
		uintptr(unsafe.Pointer(lpParameter)),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func ConvertThreadToFiberEx(lpParameter LPVOID, dwFlags uint32) LPVOID {
	ret1 := syscall3(convertThreadToFiberEx, 2,
		uintptr(unsafe.Pointer(lpParameter)),
		uintptr(dwFlags),
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): LPPROGRESS_ROUTINE
// func CopyFileEx(lpExistingFileName string, lpNewFileName string, lpProgressRoutine LPPROGRESS_ROUTINE, lpData LPVOID, pbCancel *BOOL, dwCopyFlags uint32) bool

// TODO: Unknown type(s): LPPROGRESS_ROUTINE
// func CopyFileTransacted(lpExistingFileName string, lpNewFileName string, lpProgressRoutine LPPROGRESS_ROUTINE, lpData LPVOID, pbCancel *BOOL, dwCopyFlags uint32, hTransaction HANDLE) bool

func CopyFile(lpExistingFileName string, lpNewFileName string, bFailIfExists bool) bool {
	lpExistingFileNameStr := unicode16FromString(lpExistingFileName)
	lpNewFileNameStr := unicode16FromString(lpNewFileName)
	ret1 := syscall3(copyFile, 3,
		uintptr(unsafe.Pointer(&lpExistingFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpNewFileNameStr[0])),
		getUintptrFromBool(bFailIfExists))
	return ret1 != 0
}

func CopyLZFile(unnamed0 int32, unnamed1 int32) LONG {
	ret1 := syscall3(copyLZFile, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return LONG(ret1)
}

// TODO: Unknown type(s): PCACTCTXW
// func CreateActCtx(pActCtx PCACTCTXW) HANDLE

func CreateBoundaryDescriptor(name string, flags ULONG) HANDLE {
	nameStr := unicode16FromString(name)
	ret1 := syscall3(createBoundaryDescriptor, 2,
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(flags),
		0)
	return HANDLE(ret1)
}

func CreateConsoleScreenBuffer(dwDesiredAccess uint32, dwShareMode uint32, lpSecurityAttributes /*const*/ *SECURITY_ATTRIBUTES, dwFlags uint32, lpScreenBufferData LPVOID) HANDLE {
	ret1 := syscall6(createConsoleScreenBuffer, 5,
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpScreenBufferData)),
		0)
	return HANDLE(ret1)
}

func CreateDirectoryEx(lpTemplateDirectory string, lpNewDirectory string, lpSecurityAttributes *SECURITY_ATTRIBUTES) bool {
	lpTemplateDirectoryStr := unicode16FromString(lpTemplateDirectory)
	lpNewDirectoryStr := unicode16FromString(lpNewDirectory)
	ret1 := syscall3(createDirectoryEx, 3,
		uintptr(unsafe.Pointer(&lpTemplateDirectoryStr[0])),
		uintptr(unsafe.Pointer(&lpNewDirectoryStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)))
	return ret1 != 0
}

func CreateDirectoryTransacted(lpTemplateDirectory string, lpNewDirectory string, lpSecurityAttributes *SECURITY_ATTRIBUTES, hTransaction HANDLE) bool {
	lpTemplateDirectoryStr := unicode16FromString(lpTemplateDirectory)
	lpNewDirectoryStr := unicode16FromString(lpNewDirectory)
	ret1 := syscall6(createDirectoryTransacted, 4,
		uintptr(unsafe.Pointer(&lpTemplateDirectoryStr[0])),
		uintptr(unsafe.Pointer(&lpNewDirectoryStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(hTransaction),
		0,
		0)
	return ret1 != 0
}

func CreateDirectory(lpPathName string, lpSecurityAttributes *SECURITY_ATTRIBUTES) bool {
	lpPathNameStr := unicode16FromString(lpPathName)
	ret1 := syscall3(createDirectory, 2,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		0)
	return ret1 != 0
}

func CreateEventEx(lpEventAttributes *SECURITY_ATTRIBUTES, lpName string, dwFlags uint32, dwDesiredAccess uint32) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createEventEx, 4,
		uintptr(unsafe.Pointer(lpEventAttributes)),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(dwFlags),
		uintptr(dwDesiredAccess),
		0,
		0)
	return HANDLE(ret1)
}

func CreateEvent(lpEventAttributes *SECURITY_ATTRIBUTES, bManualReset bool, bInitialState bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createEvent, 4,
		uintptr(unsafe.Pointer(lpEventAttributes)),
		getUintptrFromBool(bManualReset),
		getUintptrFromBool(bInitialState),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		0,
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): LPFIBER_START_ROUTINE
// func CreateFiber(dwStackSize SIZE_T, lpStartAddress LPFIBER_START_ROUTINE, lpParameter LPVOID) LPVOID

// TODO: Unknown type(s): LPFIBER_START_ROUTINE
// func CreateFiberEx(dwStackCommitSize SIZE_T, dwStackReserveSize SIZE_T, dwFlags uint32, lpStartAddress LPFIBER_START_ROUTINE, lpParameter LPVOID) LPVOID

// TODO: Unknown type(s): LPCREATEFILE2_EXTENDED_PARAMETERS
// func CreateFile2(lpFileName string, dwDesiredAccess uint32, dwShareMode uint32, dwCreationDisposition uint32, pCreateExParams LPCREATEFILE2_EXTENDED_PARAMETERS) HANDLE

func CreateFileMappingNuma(hFile HANDLE, lpFileMappingAttributes *SECURITY_ATTRIBUTES, flProtect uint32, dwMaximumSizeHigh uint32, dwMaximumSizeLow uint32, lpName string, nndPreferred uint32) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall9(createFileMappingNuma, 7,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpFileMappingAttributes)),
		uintptr(flProtect),
		uintptr(dwMaximumSizeHigh),
		uintptr(dwMaximumSizeLow),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(nndPreferred),
		0,
		0)
	return HANDLE(ret1)
}

func CreateFileMapping(hFile HANDLE, lpFileMappingAttributes *SECURITY_ATTRIBUTES, flProtect uint32, dwMaximumSizeHigh uint32, dwMaximumSizeLow uint32, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createFileMapping, 6,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpFileMappingAttributes)),
		uintptr(flProtect),
		uintptr(dwMaximumSizeHigh),
		uintptr(dwMaximumSizeLow),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

func CreateFileTransacted(lpFileName string, dwDesiredAccess uint32, dwShareMode uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition uint32, dwFlagsAndAttributes uint32, hTemplateFile HANDLE, hTransaction HANDLE, pusMiniVersion PUSHORT, lpExtendedParameter uintptr) HANDLE {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall12(createFileTransacted, 10,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(dwCreationDisposition),
		uintptr(dwFlagsAndAttributes),
		uintptr(hTemplateFile),
		uintptr(hTransaction),
		uintptr(unsafe.Pointer(pusMiniVersion)),
		lpExtendedParameter,
		0,
		0)
	return HANDLE(ret1)
}

func CreateFile(lpFileName string, dwDesiredAccess uint32, dwShareMode uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition uint32, dwFlagsAndAttributes uint32, hTemplateFile HANDLE) HANDLE {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall9(createFile, 7,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(dwCreationDisposition),
		uintptr(dwFlagsAndAttributes),
		uintptr(hTemplateFile),
		0,
		0)
	return HANDLE(ret1)
}

func CreateHardLinkTransacted(lpFileName string, lpExistingFileName string, lpSecurityAttributes *SECURITY_ATTRIBUTES, hTransaction HANDLE) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	lpExistingFileNameStr := unicode16FromString(lpExistingFileName)
	ret1 := syscall6(createHardLinkTransacted, 4,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpExistingFileNameStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(hTransaction),
		0,
		0)
	return ret1 != 0
}

func CreateHardLink(lpFileName string, lpExistingFileName string, lpSecurityAttributes *SECURITY_ATTRIBUTES) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	lpExistingFileNameStr := unicode16FromString(lpExistingFileName)
	ret1 := syscall3(createHardLink, 3,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpExistingFileNameStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)))
	return ret1 != 0
}

func CreateIoCompletionPort(fileHandle HANDLE, existingCompletionPort HANDLE, completionKey *uint32, numberOfConcurrentThreads uint32) HANDLE {
	ret1 := syscall6(createIoCompletionPort, 4,
		uintptr(fileHandle),
		uintptr(existingCompletionPort),
		uintptr(unsafe.Pointer(completionKey)),
		uintptr(numberOfConcurrentThreads),
		0,
		0)
	return HANDLE(ret1)
}

func CreateJobObject(lpJobAttributes *SECURITY_ATTRIBUTES, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(createJobObject, 2,
		uintptr(unsafe.Pointer(lpJobAttributes)),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): PJOB_SET_ARRAY
// func CreateJobSet(numJob ULONG, userJobSet PJOB_SET_ARRAY, flags ULONG) bool

func CreateMailslot(lpName string, nMaxMessageSize uint32, lReadTimeout uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createMailslot, 4,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(nMaxMessageSize),
		uintptr(lReadTimeout),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		0,
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): MEMORY_RESOURCE_NOTIFICATION_TYPE
// func CreateMemoryResourceNotification(notificationType MEMORY_RESOURCE_NOTIFICATION_TYPE) HANDLE

func CreateMutexEx(lpMutexAttributes *SECURITY_ATTRIBUTES, lpName string, dwFlags uint32, dwDesiredAccess uint32) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createMutexEx, 4,
		uintptr(unsafe.Pointer(lpMutexAttributes)),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(dwFlags),
		uintptr(dwDesiredAccess),
		0,
		0)
	return HANDLE(ret1)
}

func CreateMutex(lpMutexAttributes *SECURITY_ATTRIBUTES, bInitialOwner bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(createMutex, 3,
		uintptr(unsafe.Pointer(lpMutexAttributes)),
		getUintptrFromBool(bInitialOwner),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

func CreateNamedPipe(lpName string, dwOpenMode uint32, dwPipeMode uint32, nMaxInstances uint32, nOutBufferSize uint32, nInBufferSize uint32, nDefaultTimeOut uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall9(createNamedPipe, 8,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(dwOpenMode),
		uintptr(dwPipeMode),
		uintptr(nMaxInstances),
		uintptr(nOutBufferSize),
		uintptr(nInBufferSize),
		uintptr(nDefaultTimeOut),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		0)
	return HANDLE(ret1)
}

func CreatePipe(hReadPipe *HANDLE, hWritePipe *HANDLE, lpPipeAttributes *SECURITY_ATTRIBUTES, nSize uint32) bool {
	ret1 := syscall6(createPipe, 4,
		uintptr(unsafe.Pointer(hReadPipe)),
		uintptr(unsafe.Pointer(hWritePipe)),
		uintptr(unsafe.Pointer(lpPipeAttributes)),
		uintptr(nSize),
		0,
		0)
	return ret1 != 0
}

func CreatePrivateNamespace(lpPrivateNamespaceAttributes *SECURITY_ATTRIBUTES, lpBoundaryDescriptor LPVOID, lpAliasPrefix string) HANDLE {
	lpAliasPrefixStr := unicode16FromString(lpAliasPrefix)
	ret1 := syscall3(createPrivateNamespace, 3,
		uintptr(unsafe.Pointer(lpPrivateNamespaceAttributes)),
		uintptr(unsafe.Pointer(lpBoundaryDescriptor)),
		uintptr(unsafe.Pointer(&lpAliasPrefixStr[0])))
	return HANDLE(ret1)
}

func CreateProcess(lpApplicationName string, lpCommandLine LPWSTR, lpProcessAttributes *SECURITY_ATTRIBUTES, lpThreadAttributes *SECURITY_ATTRIBUTES, bInheritHandles bool, dwCreationFlags uint32, lpEnvironment LPVOID, lpCurrentDirectory string, lpStartupInfo *STARTUPINFO, lpProcessInformation *PROCESS_INFORMATION) bool {
	lpApplicationNameStr := unicode16FromString(lpApplicationName)
	lpCurrentDirectoryStr := unicode16FromString(lpCurrentDirectory)
	ret1 := syscall12(createProcess, 10,
		uintptr(unsafe.Pointer(&lpApplicationNameStr[0])),
		uintptr(unsafe.Pointer(lpCommandLine)),
		uintptr(unsafe.Pointer(lpProcessAttributes)),
		uintptr(unsafe.Pointer(lpThreadAttributes)),
		getUintptrFromBool(bInheritHandles),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpEnvironment)),
		uintptr(unsafe.Pointer(&lpCurrentDirectoryStr[0])),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)),
		0,
		0)
	return ret1 != 0
}

func CreateRemoteThread(hProcess HANDLE, lpThreadAttributes *SECURITY_ATTRIBUTES, dwStackSize SIZE_T, lpStartAddress THREAD_START_ROUTINE, lpParameter LPVOID, dwCreationFlags uint32, lpThreadId *uint32) HANDLE {
	lpStartAddressCallback := syscall.NewCallback(func(lpThreadParameterRawArg LPVOID) uintptr {
		ret := lpStartAddress(lpThreadParameterRawArg)
		return uintptr(ret)
	})
	ret1 := syscall9(createRemoteThread, 7,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpThreadAttributes)),
		uintptr(dwStackSize),
		lpStartAddressCallback,
		uintptr(unsafe.Pointer(lpParameter)),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpThreadId)),
		0,
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): LPPROC_THREAD_ATTRIBUTE_LIST
// func CreateRemoteThreadEx(hProcess HANDLE, lpThreadAttributes *SECURITY_ATTRIBUTES, dwStackSize SIZE_T, lpStartAddress THREAD_START_ROUTINE, lpParameter LPVOID, dwCreationFlags uint32, lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST, lpThreadId *uint32) HANDLE

func CreateSemaphoreEx(lpSemaphoreAttributes *SECURITY_ATTRIBUTES, lInitialCount LONG, lMaximumCount LONG, lpName string, dwFlags uint32, dwDesiredAccess uint32) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createSemaphoreEx, 6,
		uintptr(unsafe.Pointer(lpSemaphoreAttributes)),
		uintptr(lInitialCount),
		uintptr(lMaximumCount),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(dwFlags),
		uintptr(dwDesiredAccess))
	return HANDLE(ret1)
}

func CreateSemaphore(lpSemaphoreAttributes *SECURITY_ATTRIBUTES, lInitialCount LONG, lMaximumCount LONG, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(createSemaphore, 4,
		uintptr(unsafe.Pointer(lpSemaphoreAttributes)),
		uintptr(lInitialCount),
		uintptr(lMaximumCount),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		0,
		0)
	return HANDLE(ret1)
}

func CreateSymbolicLinkTransacted(lpSymlinkFileName string, lpTargetFileName string, dwFlags uint32, hTransaction HANDLE) BOOLEAN {
	lpSymlinkFileNameStr := unicode16FromString(lpSymlinkFileName)
	lpTargetFileNameStr := unicode16FromString(lpTargetFileName)
	ret1 := syscall6(createSymbolicLinkTransacted, 4,
		uintptr(unsafe.Pointer(&lpSymlinkFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpTargetFileNameStr[0])),
		uintptr(dwFlags),
		uintptr(hTransaction),
		0,
		0)
	return BOOLEAN(ret1)
}

func CreateSymbolicLink(lpSymlinkFileName string, lpTargetFileName string, dwFlags uint32) BOOLEAN {
	lpSymlinkFileNameStr := unicode16FromString(lpSymlinkFileName)
	lpTargetFileNameStr := unicode16FromString(lpTargetFileName)
	ret1 := syscall3(createSymbolicLink, 3,
		uintptr(unsafe.Pointer(&lpSymlinkFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpTargetFileNameStr[0])),
		uintptr(dwFlags))
	return BOOLEAN(ret1)
}

func CreateTapePartition(hDevice HANDLE, dwPartitionMethod uint32, dwCount uint32, dwSize uint32) uint32 {
	ret1 := syscall6(createTapePartition, 4,
		uintptr(hDevice),
		uintptr(dwPartitionMethod),
		uintptr(dwCount),
		uintptr(dwSize),
		0,
		0)
	return uint32(ret1)
}

func CreateThread(lpThreadAttributes *SECURITY_ATTRIBUTES, dwStackSize SIZE_T, lpStartAddress THREAD_START_ROUTINE, lpParameter LPVOID, dwCreationFlags uint32, lpThreadId *uint32) HANDLE {
	lpStartAddressCallback := syscall.NewCallback(func(lpThreadParameterRawArg LPVOID) uintptr {
		ret := lpStartAddress(lpThreadParameterRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(createThread, 6,
		uintptr(unsafe.Pointer(lpThreadAttributes)),
		uintptr(dwStackSize),
		lpStartAddressCallback,
		uintptr(unsafe.Pointer(lpParameter)),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpThreadId)))
	return HANDLE(ret1)
}

// TODO: Unknown type(s): PTP_POOL
// func CreateThreadpool(reserved uintptr) PTP_POOL

// TODO: Unknown type(s): PTP_CLEANUP_GROUP
// func CreateThreadpoolCleanupGroup() PTP_CLEANUP_GROUP

// TODO: Unknown type(s): PTP_CALLBACK_ENVIRON, PTP_IO, PTP_WIN32_IO_CALLBACK
// func CreateThreadpoolIo(fl HANDLE, pfnio PTP_WIN32_IO_CALLBACK, pv uintptr, pcbe PTP_CALLBACK_ENVIRON) PTP_IO

// TODO: Unknown type(s): PTP_CALLBACK_ENVIRON, PTP_TIMER, PTP_TIMER_CALLBACK
// func CreateThreadpoolTimer(pfnti PTP_TIMER_CALLBACK, pv uintptr, pcbe PTP_CALLBACK_ENVIRON) PTP_TIMER

// TODO: Unknown type(s): PTP_CALLBACK_ENVIRON, PTP_WAIT, PTP_WAIT_CALLBACK
// func CreateThreadpoolWait(pfnwa PTP_WAIT_CALLBACK, pv uintptr, pcbe PTP_CALLBACK_ENVIRON) PTP_WAIT

// TODO: Unknown type(s): PTP_CALLBACK_ENVIRON, PTP_WORK, PTP_WORK_CALLBACK
// func CreateThreadpoolWork(pfnwk PTP_WORK_CALLBACK, pv uintptr, pcbe PTP_CALLBACK_ENVIRON) PTP_WORK

func CreateTimerQueue() HANDLE {
	ret1 := syscall3(createTimerQueue, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): WAITORTIMERCALLBACK
// func CreateTimerQueueTimer(phNewTimer *HANDLE, timerQueue HANDLE, callback WAITORTIMERCALLBACK, parameter uintptr, dueTime uint32, period uint32, flags ULONG) bool

func CreateToolhelp32Snapshot(dwFlags uint32, th32ProcessID uint32) HANDLE {
	ret1 := syscall3(createToolhelp32Snapshot, 2,
		uintptr(dwFlags),
		uintptr(th32ProcessID),
		0)
	return HANDLE(ret1)
}

func CreateWaitableTimerEx(lpTimerAttributes *SECURITY_ATTRIBUTES, lpTimerName string, dwFlags uint32, dwDesiredAccess uint32) HANDLE {
	lpTimerNameStr := unicode16FromString(lpTimerName)
	ret1 := syscall6(createWaitableTimerEx, 4,
		uintptr(unsafe.Pointer(lpTimerAttributes)),
		uintptr(unsafe.Pointer(&lpTimerNameStr[0])),
		uintptr(dwFlags),
		uintptr(dwDesiredAccess),
		0,
		0)
	return HANDLE(ret1)
}

func CreateWaitableTimer(lpTimerAttributes *SECURITY_ATTRIBUTES, bManualReset bool, lpTimerName string) HANDLE {
	lpTimerNameStr := unicode16FromString(lpTimerName)
	ret1 := syscall3(createWaitableTimer, 3,
		uintptr(unsafe.Pointer(lpTimerAttributes)),
		getUintptrFromBool(bManualReset),
		uintptr(unsafe.Pointer(&lpTimerNameStr[0])))
	return HANDLE(ret1)
}

func DeactivateActCtx(dwFlags uint32, ulCookie *uint32) bool {
	ret1 := syscall3(deactivateActCtx, 2,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(ulCookie)),
		0)
	return ret1 != 0
}

func DebugActiveProcess(dwProcessId uint32) bool {
	ret1 := syscall3(debugActiveProcess, 1,
		uintptr(dwProcessId),
		0,
		0)
	return ret1 != 0
}

func DebugActiveProcessStop(dwProcessId uint32) bool {
	ret1 := syscall3(debugActiveProcessStop, 1,
		uintptr(dwProcessId),
		0,
		0)
	return ret1 != 0
}

func DebugBreak() {
	syscall3(debugBreak, 0,
		0,
		0,
		0)
}

func DebugBreakProcess(process HANDLE) bool {
	ret1 := syscall3(debugBreakProcess, 1,
		uintptr(process),
		0,
		0)
	return ret1 != 0
}

func DebugSetProcessKillOnExit(killOnExit bool) bool {
	ret1 := syscall3(debugSetProcessKillOnExit, 1,
		getUintptrFromBool(killOnExit),
		0,
		0)
	return ret1 != 0
}

func DecodePointer(ptr uintptr) uintptr {
	ret1 := syscall3(decodePointer, 1,
		ptr,
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func DecodeSystemPointer(ptr uintptr) uintptr {
	ret1 := syscall3(decodeSystemPointer, 1,
		ptr,
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func DefineDosDevice(dwFlags uint32, lpDeviceName string, lpTargetPath string) bool {
	lpDeviceNameStr := unicode16FromString(lpDeviceName)
	lpTargetPathStr := unicode16FromString(lpTargetPath)
	ret1 := syscall3(defineDosDevice, 3,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpDeviceNameStr[0])),
		uintptr(unsafe.Pointer(&lpTargetPathStr[0])))
	return ret1 != 0
}

func DeleteAtom(nAtom ATOM) ATOM {
	ret1 := syscall3(deleteAtom, 1,
		uintptr(nAtom),
		0,
		0)
	return ATOM(ret1)
}

func DeleteBoundaryDescriptor(boundaryDescriptor HANDLE) {
	syscall3(deleteBoundaryDescriptor, 1,
		uintptr(boundaryDescriptor),
		0,
		0)
}

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func DeleteCriticalSection(lpCriticalSection LPCRITICAL_SECTION)

func DeleteFiber(lpFiber LPVOID) {
	syscall3(deleteFiber, 1,
		uintptr(unsafe.Pointer(lpFiber)),
		0,
		0)
}

func DeleteFileTransacted(lpFileName string, hTransaction HANDLE) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(deleteFileTransacted, 2,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(hTransaction),
		0)
	return ret1 != 0
}

func DeleteFile(lpFileName string) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(deleteFile, 1,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPPROC_THREAD_ATTRIBUTE_LIST
// func DeleteProcThreadAttributeList(lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST)

func DeleteTimerQueue(timerQueue HANDLE) bool {
	ret1 := syscall3(deleteTimerQueue, 1,
		uintptr(timerQueue),
		0,
		0)
	return ret1 != 0
}

func DeleteTimerQueueEx(timerQueue HANDLE, completionEvent HANDLE) bool {
	ret1 := syscall3(deleteTimerQueueEx, 2,
		uintptr(timerQueue),
		uintptr(completionEvent),
		0)
	return ret1 != 0
}

func DeleteTimerQueueTimer(timerQueue HANDLE, timer HANDLE, completionEvent HANDLE) bool {
	ret1 := syscall3(deleteTimerQueueTimer, 3,
		uintptr(timerQueue),
		uintptr(timer),
		uintptr(completionEvent))
	return ret1 != 0
}

func DeleteVolumeMountPoint(lpszVolumeMountPoint string) bool {
	lpszVolumeMountPointStr := unicode16FromString(lpszVolumeMountPoint)
	ret1 := syscall3(deleteVolumeMountPoint, 1,
		uintptr(unsafe.Pointer(&lpszVolumeMountPointStr[0])),
		0,
		0)
	return ret1 != 0
}

func DeviceIoControl(hDevice HANDLE, dwIoControlCode uint32, lpInBuffer LPVOID, nInBufferSize uint32, lpOutBuffer LPVOID, nOutBufferSize uint32, lpBytesReturned *uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall9(deviceIoControl, 8,
		uintptr(hDevice),
		uintptr(dwIoControlCode),
		uintptr(unsafe.Pointer(lpInBuffer)),
		uintptr(nInBufferSize),
		uintptr(unsafe.Pointer(lpOutBuffer)),
		uintptr(nOutBufferSize),
		uintptr(unsafe.Pointer(lpBytesReturned)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0)
	return ret1 != 0
}

func DisableThreadLibraryCalls(hLibModule HMODULE) bool {
	ret1 := syscall3(disableThreadLibraryCalls, 1,
		uintptr(hLibModule),
		0,
		0)
	return ret1 != 0
}

func DisableThreadProfiling(performanceDataHandle HANDLE) uint32 {
	ret1 := syscall3(disableThreadProfiling, 1,
		uintptr(performanceDataHandle),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PTP_CALLBACK_INSTANCE
// func DisassociateCurrentThreadFromCallback(pci PTP_CALLBACK_INSTANCE)

func DiscardVirtualMemory(virtualAddress uintptr, size SIZE_T) uint32 {
	ret1 := syscall3(discardVirtualMemory, 2,
		virtualAddress,
		uintptr(size),
		0)
	return uint32(ret1)
}

func DisconnectNamedPipe(hNamedPipe HANDLE) bool {
	ret1 := syscall3(disconnectNamedPipe, 1,
		uintptr(hNamedPipe),
		0,
		0)
	return ret1 != 0
}

func DnsHostnameToComputerName(hostname string, computerName LPWSTR, nSize *uint32) bool {
	hostnameStr := unicode16FromString(hostname)
	ret1 := syscall3(dnsHostnameToComputerName, 3,
		uintptr(unsafe.Pointer(&hostnameStr[0])),
		uintptr(unsafe.Pointer(computerName)),
		uintptr(unsafe.Pointer(nSize)))
	return ret1 != 0
}

func DosDateTimeToFileTime(wFatDate uint16, wFatTime uint16, lpFileTime *FILETIME) bool {
	ret1 := syscall3(dosDateTimeToFileTime, 3,
		uintptr(wFatDate),
		uintptr(wFatTime),
		uintptr(unsafe.Pointer(lpFileTime)))
	return ret1 != 0
}

func DuplicateHandle(hSourceProcessHandle HANDLE, hSourceHandle HANDLE, hTargetProcessHandle HANDLE, lpTargetHandle *HANDLE, dwDesiredAccess uint32, bInheritHandle bool, dwOptions uint32) bool {
	ret1 := syscall9(duplicateHandle, 7,
		uintptr(hSourceProcessHandle),
		uintptr(hSourceHandle),
		uintptr(hTargetProcessHandle),
		uintptr(unsafe.Pointer(lpTargetHandle)),
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(dwOptions),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): DWORD64
// func EnableThreadProfiling(threadHandle HANDLE, flags uint32, hardwareCounters DWORD64, performanceDataHandle *HANDLE) uint32

func EncodePointer(ptr uintptr) uintptr {
	ret1 := syscall3(encodePointer, 1,
		ptr,
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func EncodeSystemPointer(ptr uintptr) uintptr {
	ret1 := syscall3(encodeSystemPointer, 1,
		ptr,
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func EndUpdateResource(hUpdate HANDLE, fDiscard bool) bool {
	ret1 := syscall3(endUpdateResource, 2,
		uintptr(hUpdate),
		getUintptrFromBool(fDiscard),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func EnterCriticalSection(lpCriticalSection LPCRITICAL_SECTION)

// TODO: Unknown type(s): CALINFO_ENUMPROCEXEX, CALTYPE
// func EnumCalendarInfoExEx(pCalInfoEnumProcExEx CALINFO_ENUMPROCEXEX, lpLocaleName string, calendar CALID, lpReserved string, calType CALTYPE, lParam LPARAM) bool

// TODO: Unknown type(s): CALINFO_ENUMPROCEXW, CALTYPE
// func EnumCalendarInfoEx(lpCalInfoEnumProcEx CALINFO_ENUMPROCEXW, locale LCID, calendar CALID, calType CALTYPE) bool

// TODO: Unknown type(s): CALINFO_ENUMPROCW, CALTYPE
// func EnumCalendarInfo(lpCalInfoEnumProc CALINFO_ENUMPROCW, locale LCID, calendar CALID, calType CALTYPE) bool

// TODO: Unknown type(s): DATEFMT_ENUMPROCEXEX
// func EnumDateFormatsExEx(lpDateFmtEnumProcExEx DATEFMT_ENUMPROCEXEX, lpLocaleName string, dwFlags uint32, lParam LPARAM) bool

// TODO: Unknown type(s): DATEFMT_ENUMPROCEXW
// func EnumDateFormatsEx(lpDateFmtEnumProcEx DATEFMT_ENUMPROCEXW, locale LCID, dwFlags uint32) bool

// TODO: Unknown type(s): DATEFMT_ENUMPROCW
// func EnumDateFormats(lpDateFmtEnumProc DATEFMT_ENUMPROCW, locale LCID, dwFlags uint32) bool

// TODO: Unknown type(s): LANGGROUPLOCALE_ENUMPROCW, LGRPID
// func EnumLanguageGroupLocales(lpLangGroupLocaleEnumProc LANGGROUPLOCALE_ENUMPROCW, languageGroup LGRPID, dwFlags uint32, lParam uintptr) bool

func EnumResourceLanguagesEx(hModule HMODULE, lpType string, lpName string, lpEnumFunc ENUMRESLANGPROC, lParam uintptr, dwFlags uint32, langId LANGID) bool {
	lpTypeStr := unicode16FromString(lpType)
	lpNameStr := unicode16FromString(lpName)
	lpEnumFuncCallback := syscall.NewCallback(func(hModuleRawArg HMODULE, lpTypeRawArg /*const*/ *uint16, lpNameRawArg /*const*/ *uint16, wLanguageRawArg WORD, lParamRawArg LONG_PTR) uintptr {
		lpType := stringFromUnicode16(lpTypeRawArg)
		lpName := stringFromUnicode16(lpNameRawArg)
		ret := lpEnumFunc(hModuleRawArg, lpType, lpName, wLanguageRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall9(enumResourceLanguagesEx, 7,
		uintptr(hModule),
		uintptr(unsafe.Pointer(&lpTypeStr[0])),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		lpEnumFuncCallback,
		lParam,
		uintptr(dwFlags),
		uintptr(langId),
		0,
		0)
	return ret1 != 0
}

func EnumResourceLanguages(hModule HMODULE, lpType string, lpName string, lpEnumFunc ENUMRESLANGPROC, lParam uintptr) bool {
	lpTypeStr := unicode16FromString(lpType)
	lpNameStr := unicode16FromString(lpName)
	lpEnumFuncCallback := syscall.NewCallback(func(hModuleRawArg HMODULE, lpTypeRawArg /*const*/ *uint16, lpNameRawArg /*const*/ *uint16, wLanguageRawArg WORD, lParamRawArg LONG_PTR) uintptr {
		lpType := stringFromUnicode16(lpTypeRawArg)
		lpName := stringFromUnicode16(lpNameRawArg)
		ret := lpEnumFunc(hModuleRawArg, lpType, lpName, wLanguageRawArg, lParamRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumResourceLanguages, 5,
		uintptr(hModule),
		uintptr(unsafe.Pointer(&lpTypeStr[0])),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		lpEnumFuncCallback,
		lParam,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): ENUMRESNAMEPROCW
// func EnumResourceNamesEx(hModule HMODULE, lpType string, lpEnumFunc ENUMRESNAMEPROCW, lParam uintptr, dwFlags uint32, langId LANGID) bool

// TODO: Unknown type(s): ENUMRESNAMEPROCW
// func EnumResourceNames(hModule HMODULE, lpType string, lpEnumFunc ENUMRESNAMEPROCW, lParam uintptr) bool

// TODO: Unknown type(s): ENUMRESTYPEPROCW
// func EnumResourceTypesEx(hModule HMODULE, lpEnumFunc ENUMRESTYPEPROCW, lParam uintptr, dwFlags uint32, langId LANGID) bool

// TODO: Unknown type(s): ENUMRESTYPEPROCW
// func EnumResourceTypes(hModule HMODULE, lpEnumFunc ENUMRESTYPEPROCW, lParam uintptr) bool

// TODO: Unknown type(s): CODEPAGE_ENUMPROCW
// func EnumSystemCodePages(lpCodePageEnumProc CODEPAGE_ENUMPROCW, dwFlags uint32) bool

func EnumSystemFirmwareTables(firmwareTableProviderSignature uint32, pFirmwareTableEnumBuffer uintptr, bufferSize uint32) uint32 {
	ret1 := syscall3(enumSystemFirmwareTables, 3,
		uintptr(firmwareTableProviderSignature),
		pFirmwareTableEnumBuffer,
		uintptr(bufferSize))
	return uint32(ret1)
}

// TODO: Unknown type(s): GEOCLASS, GEO_ENUMPROC
// func EnumSystemGeoID(geoClass GEOCLASS, parentGeoId GEOID, lpGeoEnumProc GEO_ENUMPROC) bool

// TODO: Unknown type(s): LANGUAGEGROUP_ENUMPROCW
// func EnumSystemLanguageGroups(lpLanguageGroupEnumProc LANGUAGEGROUP_ENUMPROCW, dwFlags uint32, lParam uintptr) bool

// TODO: Unknown type(s): LOCALE_ENUMPROCEX
// func EnumSystemLocalesEx(lpLocaleEnumProcEx LOCALE_ENUMPROCEX, dwFlags uint32, lParam LPARAM, lpReserved LPVOID) bool

// TODO: Unknown type(s): LOCALE_ENUMPROCW
// func EnumSystemLocales(lpLocaleEnumProc LOCALE_ENUMPROCW, dwFlags uint32) bool

// TODO: Unknown type(s): TIMEFMT_ENUMPROCEX
// func EnumTimeFormatsEx(lpTimeFmtEnumProcEx TIMEFMT_ENUMPROCEX, lpLocaleName string, dwFlags uint32, lParam LPARAM) bool

// TODO: Unknown type(s): TIMEFMT_ENUMPROCW
// func EnumTimeFormats(lpTimeFmtEnumProc TIMEFMT_ENUMPROCW, locale LCID, dwFlags uint32) bool

// TODO: Unknown type(s): UILANGUAGE_ENUMPROCW
// func EnumUILanguages(lpUILanguageEnumProc UILANGUAGE_ENUMPROCW, dwFlags uint32, lParam uintptr) bool

func EraseTape(hDevice HANDLE, dwEraseType uint32, bImmediate bool) uint32 {
	ret1 := syscall3(eraseTape, 3,
		uintptr(hDevice),
		uintptr(dwEraseType),
		getUintptrFromBool(bImmediate))
	return uint32(ret1)
}

func EscapeCommFunction(hFile HANDLE, dwFunc uint32) bool {
	ret1 := syscall3(escapeCommFunction, 2,
		uintptr(hFile),
		uintptr(dwFunc),
		0)
	return ret1 != 0
}

func ExitProcess(uExitCode uint32) {
	syscall3(exitProcess, 1,
		uintptr(uExitCode),
		0,
		0)
}

func ExitThread(dwExitCode uint32) {
	syscall3(exitThread, 1,
		uintptr(dwExitCode),
		0,
		0)
}

func ExpandEnvironmentStrings(lpSrc string, lpDst LPWSTR, nSize uint32) uint32 {
	lpSrcStr := unicode16FromString(lpSrc)
	ret1 := syscall3(expandEnvironmentStrings, 3,
		uintptr(unsafe.Pointer(&lpSrcStr[0])),
		uintptr(unsafe.Pointer(lpDst)),
		uintptr(nSize))
	return uint32(ret1)
}

func FatalAppExit(uAction uint32, lpMessageText string) {
	lpMessageTextStr := unicode16FromString(lpMessageText)
	syscall3(fatalAppExit, 2,
		uintptr(uAction),
		uintptr(unsafe.Pointer(&lpMessageTextStr[0])),
		0)
}

func FatalExit(exitCode int32) {
	syscall3(fatalExit, 1,
		uintptr(exitCode),
		0,
		0)
}

func FileTimeToDosDateTime(lpFileTime /*const*/ *FILETIME, lpFatDate *uint16, lpFatTime *uint16) bool {
	ret1 := syscall3(fileTimeToDosDateTime, 3,
		uintptr(unsafe.Pointer(lpFileTime)),
		uintptr(unsafe.Pointer(lpFatDate)),
		uintptr(unsafe.Pointer(lpFatTime)))
	return ret1 != 0
}

func FileTimeToLocalFileTime(lpFileTime /*const*/ *FILETIME, lpLocalFileTime *FILETIME) bool {
	ret1 := syscall3(fileTimeToLocalFileTime, 2,
		uintptr(unsafe.Pointer(lpFileTime)),
		uintptr(unsafe.Pointer(lpLocalFileTime)),
		0)
	return ret1 != 0
}

func FileTimeToSystemTime(lpFileTime /*const*/ *FILETIME, lpSystemTime *SYSTEMTIME) bool {
	ret1 := syscall3(fileTimeToSystemTime, 2,
		uintptr(unsafe.Pointer(lpFileTime)),
		uintptr(unsafe.Pointer(lpSystemTime)),
		0)
	return ret1 != 0
}

func FillConsoleOutputAttribute(hConsoleOutput HANDLE, wAttribute uint16, nLength uint32, dwWriteCoord COORD, lpNumberOfAttrsWritten *uint32) bool {
	ret1 := syscall6(fillConsoleOutputAttribute, 5,
		uintptr(hConsoleOutput),
		uintptr(wAttribute),
		uintptr(nLength),
		getUintptrFromCOORD(dwWriteCoord),
		uintptr(unsafe.Pointer(lpNumberOfAttrsWritten)),
		0)
	return ret1 != 0
}

func FillConsoleOutputCharacter(hConsoleOutput HANDLE, cCharacter WCHAR, nLength uint32, dwWriteCoord COORD, lpNumberOfCharsWritten *uint32) bool {
	ret1 := syscall6(fillConsoleOutputCharacter, 5,
		uintptr(hConsoleOutput),
		uintptr(cCharacter),
		uintptr(nLength),
		getUintptrFromCOORD(dwWriteCoord),
		uintptr(unsafe.Pointer(lpNumberOfCharsWritten)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PACTCTX_SECTION_KEYED_DATA
// func FindActCtxSectionGuid(dwFlags uint32, lpExtensionGuid /*const*/ *GUID, ulSectionId ULONG, lpGuidToFind /*const*/ *GUID, returnedData PACTCTX_SECTION_KEYED_DATA) bool

// TODO: Unknown type(s): PACTCTX_SECTION_KEYED_DATA
// func FindActCtxSectionString(dwFlags uint32, lpExtensionGuid /*const*/ *GUID, ulSectionId ULONG, lpStringToFind string, returnedData PACTCTX_SECTION_KEYED_DATA) bool

func FindAtom(lpString string) ATOM {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(findAtom, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return ATOM(ret1)
}

func FindClose(hFindFile HANDLE) bool {
	ret1 := syscall3(findClose, 1,
		uintptr(hFindFile),
		0,
		0)
	return ret1 != 0
}

func FindCloseChangeNotification(hChangeHandle HANDLE) bool {
	ret1 := syscall3(findCloseChangeNotification, 1,
		uintptr(hChangeHandle),
		0,
		0)
	return ret1 != 0
}

func FindFirstChangeNotification(lpPathName string, bWatchSubtree bool, dwNotifyFilter uint32) HANDLE {
	lpPathNameStr := unicode16FromString(lpPathName)
	ret1 := syscall3(findFirstChangeNotification, 3,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		getUintptrFromBool(bWatchSubtree),
		uintptr(dwNotifyFilter))
	return HANDLE(ret1)
}

// TODO: Unknown type(s): FINDEX_INFO_LEVELS, FINDEX_SEARCH_OPS
// func FindFirstFileEx(lpFileName string, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData LPVOID, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter LPVOID, dwAdditionalFlags uint32) HANDLE

func FindFirstFileNameTransactedW(lpFileName string, dwFlags uint32, stringLength *uint32, linkName PWSTR, hTransaction HANDLE) HANDLE {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(findFirstFileNameTransactedW, 5,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(stringLength)),
		uintptr(unsafe.Pointer(linkName)),
		uintptr(hTransaction),
		0)
	return HANDLE(ret1)
}

func FindFirstFileNameW(lpFileName string, dwFlags uint32, stringLength *uint32, linkName PWSTR) HANDLE {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(findFirstFileNameW, 4,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(stringLength)),
		uintptr(unsafe.Pointer(linkName)),
		0,
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): FINDEX_INFO_LEVELS, FINDEX_SEARCH_OPS
// func FindFirstFileTransacted(lpFileName string, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData LPVOID, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter LPVOID, dwAdditionalFlags uint32, hTransaction HANDLE) HANDLE

// TODO: Unknown type(s): LPWIN32_FIND_DATAW
// func FindFirstFile(lpFileName string, lpFindFileData LPWIN32_FIND_DATAW) HANDLE

// TODO: Unknown type(s): STREAM_INFO_LEVELS
// func FindFirstStreamTransactedW(lpFileName string, infoLevel STREAM_INFO_LEVELS, lpFindStreamData LPVOID, dwFlags uint32, hTransaction HANDLE) HANDLE

// TODO: Unknown type(s): STREAM_INFO_LEVELS
// func FindFirstStreamW(lpFileName string, infoLevel STREAM_INFO_LEVELS, lpFindStreamData LPVOID, dwFlags uint32) HANDLE

func FindFirstVolumeMountPoint(lpszRootPathName string, lpszVolumeMountPoint LPWSTR, cchBufferLength uint32) HANDLE {
	lpszRootPathNameStr := unicode16FromString(lpszRootPathName)
	ret1 := syscall3(findFirstVolumeMountPoint, 3,
		uintptr(unsafe.Pointer(&lpszRootPathNameStr[0])),
		uintptr(unsafe.Pointer(lpszVolumeMountPoint)),
		uintptr(cchBufferLength))
	return HANDLE(ret1)
}

func FindFirstVolume(lpszVolumeName LPWSTR, cchBufferLength uint32) HANDLE {
	ret1 := syscall3(findFirstVolume, 2,
		uintptr(unsafe.Pointer(lpszVolumeName)),
		uintptr(cchBufferLength),
		0)
	return HANDLE(ret1)
}

func FindNLSString(locale LCID, dwFindNLSStringFlags uint32, lpStringSource string, cchSource int32, lpStringValue string, cchValue int32, pcchFound *int32) int32 {
	lpStringSourceStr := unicode16FromString(lpStringSource)
	lpStringValueStr := unicode16FromString(lpStringValue)
	ret1 := syscall9(findNLSString, 7,
		uintptr(locale),
		uintptr(dwFindNLSStringFlags),
		uintptr(unsafe.Pointer(&lpStringSourceStr[0])),
		uintptr(cchSource),
		uintptr(unsafe.Pointer(&lpStringValueStr[0])),
		uintptr(cchValue),
		uintptr(unsafe.Pointer(pcchFound)),
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): LPNLSVERSIONINFO
// func FindNLSStringEx(lpLocaleName string, dwFindNLSStringFlags uint32, lpStringSource string, cchSource int32, lpStringValue string, cchValue int32, pcchFound *int32, lpVersionInformation LPNLSVERSIONINFO, lpReserved LPVOID, sortHandle LPARAM) int32

func FindNextChangeNotification(hChangeHandle HANDLE) bool {
	ret1 := syscall3(findNextChangeNotification, 1,
		uintptr(hChangeHandle),
		0,
		0)
	return ret1 != 0
}

func FindNextFileNameW(hFindStream HANDLE, stringLength *uint32, linkName PWSTR) bool {
	ret1 := syscall3(findNextFileNameW, 3,
		uintptr(hFindStream),
		uintptr(unsafe.Pointer(stringLength)),
		uintptr(unsafe.Pointer(linkName)))
	return ret1 != 0
}

// TODO: Unknown type(s): LPWIN32_FIND_DATAW
// func FindNextFile(hFindFile HANDLE, lpFindFileData LPWIN32_FIND_DATAW) bool

func FindNextStreamW(hFindStream HANDLE, lpFindStreamData LPVOID) bool {
	ret1 := syscall3(findNextStreamW, 2,
		uintptr(hFindStream),
		uintptr(unsafe.Pointer(lpFindStreamData)),
		0)
	return ret1 != 0
}

func FindNextVolumeMountPoint(hFindVolumeMountPoint HANDLE, lpszVolumeMountPoint LPWSTR, cchBufferLength uint32) bool {
	ret1 := syscall3(findNextVolumeMountPoint, 3,
		uintptr(hFindVolumeMountPoint),
		uintptr(unsafe.Pointer(lpszVolumeMountPoint)),
		uintptr(cchBufferLength))
	return ret1 != 0
}

func FindNextVolume(hFindVolume HANDLE, lpszVolumeName LPWSTR, cchBufferLength uint32) bool {
	ret1 := syscall3(findNextVolume, 3,
		uintptr(hFindVolume),
		uintptr(unsafe.Pointer(lpszVolumeName)),
		uintptr(cchBufferLength))
	return ret1 != 0
}

func FindResourceEx(hModule HMODULE, lpType string, lpName string, wLanguage uint16) HRSRC {
	lpTypeStr := unicode16FromString(lpType)
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(findResourceEx, 4,
		uintptr(hModule),
		uintptr(unsafe.Pointer(&lpTypeStr[0])),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(wLanguage),
		0,
		0)
	return HRSRC(ret1)
}

func FindResource(hModule HMODULE, lpName string, lpType string) HRSRC {
	lpNameStr := unicode16FromString(lpName)
	lpTypeStr := unicode16FromString(lpType)
	ret1 := syscall3(findResource, 3,
		uintptr(hModule),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(&lpTypeStr[0])))
	return HRSRC(ret1)
}

func FindStringOrdinal(dwFindStringOrdinalFlags uint32, lpStringSource string, cchSource int32, lpStringValue string, cchValue int32, bIgnoreCase bool) int32 {
	lpStringSourceStr := unicode16FromString(lpStringSource)
	lpStringValueStr := unicode16FromString(lpStringValue)
	ret1 := syscall6(findStringOrdinal, 6,
		uintptr(dwFindStringOrdinalFlags),
		uintptr(unsafe.Pointer(&lpStringSourceStr[0])),
		uintptr(cchSource),
		uintptr(unsafe.Pointer(&lpStringValueStr[0])),
		uintptr(cchValue),
		getUintptrFromBool(bIgnoreCase))
	return int32(ret1)
}

func FindVolumeClose(hFindVolume HANDLE) bool {
	ret1 := syscall3(findVolumeClose, 1,
		uintptr(hFindVolume),
		0,
		0)
	return ret1 != 0
}

func FindVolumeMountPointClose(hFindVolumeMountPoint HANDLE) bool {
	ret1 := syscall3(findVolumeMountPointClose, 1,
		uintptr(hFindVolumeMountPoint),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PFLS_CALLBACK_FUNCTION
// func FlsAlloc(lpCallback PFLS_CALLBACK_FUNCTION) uint32

func FlsFree(dwFlsIndex uint32) bool {
	ret1 := syscall3(flsFree, 1,
		uintptr(dwFlsIndex),
		0,
		0)
	return ret1 != 0
}

func FlsGetValue(dwFlsIndex uint32) uintptr {
	ret1 := syscall3(flsGetValue, 1,
		uintptr(dwFlsIndex),
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func FlsSetValue(dwFlsIndex uint32, lpFlsData uintptr) bool {
	ret1 := syscall3(flsSetValue, 2,
		uintptr(dwFlsIndex),
		lpFlsData,
		0)
	return ret1 != 0
}

func FlushConsoleInputBuffer(hConsoleInput HANDLE) bool {
	ret1 := syscall3(flushConsoleInputBuffer, 1,
		uintptr(hConsoleInput),
		0,
		0)
	return ret1 != 0
}

func FlushFileBuffers(hFile HANDLE) bool {
	ret1 := syscall3(flushFileBuffers, 1,
		uintptr(hFile),
		0,
		0)
	return ret1 != 0
}

func FlushInstructionCache(hProcess HANDLE, lpBaseAddress /*const*/ uintptr, dwSize SIZE_T) bool {
	ret1 := syscall3(flushInstructionCache, 3,
		uintptr(hProcess),
		lpBaseAddress,
		uintptr(dwSize))
	return ret1 != 0
}

func FlushProcessWriteBuffers() {
	syscall3(flushProcessWriteBuffers, 0,
		0,
		0,
		0)
}

func FlushViewOfFile(lpBaseAddress /*const*/ uintptr, dwNumberOfBytesToFlush SIZE_T) bool {
	ret1 := syscall3(flushViewOfFile, 2,
		lpBaseAddress,
		uintptr(dwNumberOfBytesToFlush),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCWCH
// func FoldString(dwMapFlags uint32, lpSrcStr LPCWCH, cchSrc int32, lpDestStr LPWSTR, cchDest int32) int32

func FreeConsole() bool {
	ret1 := syscall3(freeConsole, 0,
		0,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPWCH
// func FreeEnvironmentStrings(penv LPWCH) bool

func FreeLibrary(hLibModule HMODULE) bool {
	ret1 := syscall3(freeLibrary, 1,
		uintptr(hLibModule),
		0,
		0)
	return ret1 != 0
}

func FreeLibraryAndExitThread(hLibModule HMODULE, dwExitCode uint32) {
	syscall3(freeLibraryAndExitThread, 2,
		uintptr(hLibModule),
		uintptr(dwExitCode),
		0)
}

// TODO: Unknown type(s): PTP_CALLBACK_INSTANCE
// func FreeLibraryWhenCallbackReturns(pci PTP_CALLBACK_INSTANCE, mod HMODULE)

func FreeResource(hResData HGLOBAL) bool {
	ret1 := syscall3(freeResource, 1,
		uintptr(hResData),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PULONG_PTR
// func FreeUserPhysicalPages(hProcess HANDLE, numberOfPages PULONG_PTR, pageArray PULONG_PTR) bool

func GenerateConsoleCtrlEvent(dwCtrlEvent uint32, dwProcessGroupId uint32) bool {
	ret1 := syscall3(generateConsoleCtrlEvent, 2,
		uintptr(dwCtrlEvent),
		uintptr(dwProcessGroupId),
		0)
	return ret1 != 0
}

func GetACP() uint32 {
	ret1 := syscall3(getACP, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetActiveProcessorCount(groupNumber uint16) uint32 {
	ret1 := syscall3(getActiveProcessorCount, 1,
		uintptr(groupNumber),
		0,
		0)
	return uint32(ret1)
}

func GetActiveProcessorGroupCount() uint16 {
	ret1 := syscall3(getActiveProcessorGroupCount, 0,
		0,
		0,
		0)
	return uint16(ret1)
}

// TODO: Unknown type(s): APPLICATION_RECOVERY_CALLBACK *
// func GetApplicationRecoveryCallback(hProcess HANDLE, pRecoveryCallback APPLICATION_RECOVERY_CALLBACK *, ppvParameter *PVOID, pdwPingInterval *DWORD, pdwFlags *DWORD) HRESULT

func GetApplicationRestartSettings(hProcess HANDLE, pwzCommandline PWSTR, pcchSize *DWORD, pdwFlags *DWORD) HRESULT {
	ret1 := syscall6(getApplicationRestartSettings, 4,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(pwzCommandline)),
		uintptr(unsafe.Pointer(pcchSize)),
		uintptr(unsafe.Pointer(pdwFlags)),
		0,
		0)
	return HRESULT(ret1)
}

func GetAtomName(nAtom ATOM, lpBuffer LPWSTR, nSize int32) uint32 {
	ret1 := syscall3(getAtomName, 3,
		uintptr(nAtom),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetBinaryType(lpApplicationName string, lpBinaryType *uint32) bool {
	lpApplicationNameStr := unicode16FromString(lpApplicationName)
	ret1 := syscall3(getBinaryType, 2,
		uintptr(unsafe.Pointer(&lpApplicationNameStr[0])),
		uintptr(unsafe.Pointer(lpBinaryType)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCPINFO
// func GetCPInfo(codePage uint32, lpCPInfo LPCPINFO) bool

// TODO: Unknown type(s): LPCPINFOEXW
// func GetCPInfoEx(codePage uint32, dwFlags uint32, lpCPInfoEx LPCPINFOEXW) bool

// TODO: Unknown type(s): CALTYPE
// func GetCalendarInfoEx(lpLocaleName string, calendar CALID, lpReserved string, calType CALTYPE, lpCalData LPWSTR, cchData int32, lpValue *uint32) int32

// TODO: Unknown type(s): CALTYPE
// func GetCalendarInfo(locale LCID, calendar CALID, calType CALTYPE, lpCalData LPWSTR, cchData int32, lpValue *uint32) int32

// TODO: Unknown type(s): LPCOMMCONFIG
// func GetCommConfig(hCommDev HANDLE, lpCC LPCOMMCONFIG, lpdwSize *uint32) bool

func GetCommMask(hFile HANDLE, lpEvtMask *uint32) bool {
	ret1 := syscall3(getCommMask, 2,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpEvtMask)),
		0)
	return ret1 != 0
}

func GetCommModemStatus(hFile HANDLE, lpModemStat *uint32) bool {
	ret1 := syscall3(getCommModemStatus, 2,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpModemStat)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCOMMPROP
// func GetCommProperties(hFile HANDLE, lpCommProp LPCOMMPROP) bool

// TODO: Unknown type(s): LPDCB
// func GetCommState(hFile HANDLE, lpDCB LPDCB) bool

// TODO: Unknown type(s): LPCOMMTIMEOUTS
// func GetCommTimeouts(hFile HANDLE, lpCommTimeouts LPCOMMTIMEOUTS) bool

func GetCommandLine() LPWSTR {
	ret1 := syscall3(getCommandLine, 0,
		0,
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func GetCompressedFileSizeTransacted(lpFileName string, lpFileSizeHigh *uint32, hTransaction HANDLE) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(getCompressedFileSizeTransacted, 3,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(lpFileSizeHigh)),
		uintptr(hTransaction))
	return uint32(ret1)
}

func GetCompressedFileSize(lpFileName string, lpFileSizeHigh *uint32) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(getCompressedFileSize, 2,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(lpFileSizeHigh)),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): COMPUTER_NAME_FORMAT
// func GetComputerNameEx(nameType COMPUTER_NAME_FORMAT, lpBuffer LPWSTR, nSize *uint32) bool

func GetComputerName(lpBuffer LPWSTR, nSize *uint32) bool {
	ret1 := syscall3(getComputerName, 2,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(unsafe.Pointer(nSize)),
		0)
	return ret1 != 0
}

func GetConsoleAliasExesLength() uint32 {
	ret1 := syscall3(getConsoleAliasExesLength, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetConsoleAliasExes(exeNameBuffer LPWSTR, exeNameBufferLength uint32) uint32 {
	ret1 := syscall3(getConsoleAliasExes, 2,
		uintptr(unsafe.Pointer(exeNameBuffer)),
		uintptr(exeNameBufferLength),
		0)
	return uint32(ret1)
}

func GetConsoleAlias(source LPWSTR, targetBuffer LPWSTR, targetBufferLength uint32, exeName LPWSTR) uint32 {
	ret1 := syscall6(getConsoleAlias, 4,
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(targetBuffer)),
		uintptr(targetBufferLength),
		uintptr(unsafe.Pointer(exeName)),
		0,
		0)
	return uint32(ret1)
}

func GetConsoleAliasesLength(exeName LPWSTR) uint32 {
	ret1 := syscall3(getConsoleAliasesLength, 1,
		uintptr(unsafe.Pointer(exeName)),
		0,
		0)
	return uint32(ret1)
}

func GetConsoleAliases(aliasBuffer LPWSTR, aliasBufferLength uint32, exeName LPWSTR) uint32 {
	ret1 := syscall3(getConsoleAliases, 3,
		uintptr(unsafe.Pointer(aliasBuffer)),
		uintptr(aliasBufferLength),
		uintptr(unsafe.Pointer(exeName)))
	return uint32(ret1)
}

func GetConsoleCP() uint32 {
	ret1 := syscall3(getConsoleCP, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PCONSOLE_CURSOR_INFO
// func GetConsoleCursorInfo(hConsoleOutput HANDLE, lpConsoleCursorInfo PCONSOLE_CURSOR_INFO) bool

func GetConsoleDisplayMode(lpModeFlags *uint32) bool {
	ret1 := syscall3(getConsoleDisplayMode, 1,
		uintptr(unsafe.Pointer(lpModeFlags)),
		0,
		0)
	return ret1 != 0
}

func GetConsoleFontSize(hConsoleOutput HANDLE, nFont uint32) COORD {
	ret1 := syscall3(getConsoleFontSize, 2,
		uintptr(hConsoleOutput),
		uintptr(nFont),
		0)
	return getCOORDFromUintptr(ret1)
}

// TODO: Unknown type(s): PCONSOLE_HISTORY_INFO
// func GetConsoleHistoryInfo(lpConsoleHistoryInfo PCONSOLE_HISTORY_INFO) bool

func GetConsoleMode(hConsoleHandle HANDLE, lpMode *uint32) bool {
	ret1 := syscall3(getConsoleMode, 2,
		uintptr(hConsoleHandle),
		uintptr(unsafe.Pointer(lpMode)),
		0)
	return ret1 != 0
}

func GetConsoleOriginalTitle(lpConsoleTitle LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(getConsoleOriginalTitle, 2,
		uintptr(unsafe.Pointer(lpConsoleTitle)),
		uintptr(nSize),
		0)
	return uint32(ret1)
}

func GetConsoleOutputCP() uint32 {
	ret1 := syscall3(getConsoleOutputCP, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetConsoleProcessList(lpdwProcessList *uint32, dwProcessCount uint32) uint32 {
	ret1 := syscall3(getConsoleProcessList, 2,
		uintptr(unsafe.Pointer(lpdwProcessList)),
		uintptr(dwProcessCount),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PCONSOLE_SCREEN_BUFFER_INFO
// func GetConsoleScreenBufferInfo(hConsoleOutput HANDLE, lpConsoleScreenBufferInfo PCONSOLE_SCREEN_BUFFER_INFO) bool

// TODO: Unknown type(s): PCONSOLE_SCREEN_BUFFER_INFOEX
// func GetConsoleScreenBufferInfoEx(hConsoleOutput HANDLE, lpConsoleScreenBufferInfoEx PCONSOLE_SCREEN_BUFFER_INFOEX) bool

// TODO: Unknown type(s): PCONSOLE_SELECTION_INFO
// func GetConsoleSelectionInfo(lpConsoleSelectionInfo PCONSOLE_SELECTION_INFO) bool

func GetConsoleTitle(lpConsoleTitle LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(getConsoleTitle, 2,
		uintptr(unsafe.Pointer(lpConsoleTitle)),
		uintptr(nSize),
		0)
	return uint32(ret1)
}

func GetConsoleWindow() HWND {
	ret1 := syscall3(getConsoleWindow, 0,
		0,
		0,
		0)
	return HWND(ret1)
}

// TODO: Unknown type(s): CONST CURRENCYFMTW *
// func GetCurrencyFormatEx(lpLocaleName string, dwFlags uint32, lpValue string, lpFormat /*const*/ CONST CURRENCYFMTW *, lpCurrencyStr LPWSTR, cchCurrency int32) int32

// TODO: Unknown type(s): CONST CURRENCYFMTW *
// func GetCurrencyFormat(locale LCID, dwFlags uint32, lpValue string, lpFormat /*const*/ CONST CURRENCYFMTW *, lpCurrencyStr LPWSTR, cchCurrency int32) int32

func GetCurrentActCtx(lphActCtx *HANDLE) bool {
	ret1 := syscall3(getCurrentActCtx, 1,
		uintptr(unsafe.Pointer(lphActCtx)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCONSOLE_FONT_INFO
// func GetCurrentConsoleFont(hConsoleOutput HANDLE, bMaximumWindow bool, lpConsoleCurrentFont PCONSOLE_FONT_INFO) bool

// TODO: Unknown type(s): PCONSOLE_FONT_INFOEX
// func GetCurrentConsoleFontEx(hConsoleOutput HANDLE, bMaximumWindow bool, lpConsoleCurrentFontEx PCONSOLE_FONT_INFOEX) bool

func GetCurrentDirectory(nBufferLength uint32, lpBuffer LPWSTR) uint32 {
	ret1 := syscall3(getCurrentDirectory, 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)
	return uint32(ret1)
}

func GetCurrentProcess() HANDLE {
	ret1 := syscall3(getCurrentProcess, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

func GetCurrentProcessId() uint32 {
	ret1 := syscall3(getCurrentProcessId, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetCurrentProcessorNumber() uint32 {
	ret1 := syscall3(getCurrentProcessorNumber, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PPROCESSOR_NUMBER
// func GetCurrentProcessorNumberEx(procNumber PPROCESSOR_NUMBER)

func GetCurrentThread() HANDLE {
	ret1 := syscall3(getCurrentThread, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

func GetCurrentThreadId() uint32 {
	ret1 := syscall3(getCurrentThreadId, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetDateFormatEx(lpLocaleName string, dwFlags uint32, lpDate /*const*/ *SYSTEMTIME, lpFormat string, lpDateStr LPWSTR, cchDate int32, lpCalendar string) int32 {
	lpLocaleNameStr := unicode16FromString(lpLocaleName)
	lpFormatStr := unicode16FromString(lpFormat)
	lpCalendarStr := unicode16FromString(lpCalendar)
	ret1 := syscall9(getDateFormatEx, 7,
		uintptr(unsafe.Pointer(&lpLocaleNameStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpDate)),
		uintptr(unsafe.Pointer(&lpFormatStr[0])),
		uintptr(unsafe.Pointer(lpDateStr)),
		uintptr(cchDate),
		uintptr(unsafe.Pointer(&lpCalendarStr[0])),
		0,
		0)
	return int32(ret1)
}

func GetDateFormat(locale LCID, dwFlags uint32, lpDate /*const*/ *SYSTEMTIME, lpFormat string, lpDateStr LPWSTR, cchDate int32) int32 {
	lpFormatStr := unicode16FromString(lpFormat)
	ret1 := syscall6(getDateFormat, 6,
		uintptr(locale),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpDate)),
		uintptr(unsafe.Pointer(&lpFormatStr[0])),
		uintptr(unsafe.Pointer(lpDateStr)),
		uintptr(cchDate))
	return int32(ret1)
}

// TODO: Unknown type(s): LPCOMMCONFIG
// func GetDefaultCommConfig(lpszName string, lpCC LPCOMMCONFIG, lpdwSize *uint32) bool

func GetDevicePowerState(hDevice HANDLE, pfOn *BOOL) bool {
	ret1 := syscall3(getDevicePowerState, 2,
		uintptr(hDevice),
		uintptr(unsafe.Pointer(pfOn)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PULARGE_INTEGER
// func GetDiskFreeSpaceEx(lpDirectoryName string, lpFreeBytesAvailableToCaller PULARGE_INTEGER, lpTotalNumberOfBytes PULARGE_INTEGER, lpTotalNumberOfFreeBytes PULARGE_INTEGER) bool

func GetDiskFreeSpace(lpRootPathName string, lpSectorsPerCluster *uint32, lpBytesPerSector *uint32, lpNumberOfFreeClusters *uint32, lpTotalNumberOfClusters *uint32) bool {
	lpRootPathNameStr := unicode16FromString(lpRootPathName)
	ret1 := syscall6(getDiskFreeSpace, 5,
		uintptr(unsafe.Pointer(&lpRootPathNameStr[0])),
		uintptr(unsafe.Pointer(lpSectorsPerCluster)),
		uintptr(unsafe.Pointer(lpBytesPerSector)),
		uintptr(unsafe.Pointer(lpNumberOfFreeClusters)),
		uintptr(unsafe.Pointer(lpTotalNumberOfClusters)),
		0)
	return ret1 != 0
}

func GetDllDirectory(nBufferLength uint32, lpBuffer LPWSTR) uint32 {
	ret1 := syscall3(getDllDirectory, 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)
	return uint32(ret1)
}

func GetDriveType(lpRootPathName string) uint32 {
	lpRootPathNameStr := unicode16FromString(lpRootPathName)
	ret1 := syscall3(getDriveType, 1,
		uintptr(unsafe.Pointer(&lpRootPathNameStr[0])),
		0,
		0)
	return uint32(ret1)
}

func GetDurationFormat(locale LCID, dwFlags uint32, lpDuration /*const*/ *SYSTEMTIME, ullDuration ULONGLONG, lpFormat string, lpDurationStr LPWSTR, cchDuration int32) int32 {
	lpFormatStr := unicode16FromString(lpFormat)
	ret1 := syscall9(getDurationFormat, 7,
		uintptr(locale),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpDuration)),
		uintptr(ullDuration),
		uintptr(unsafe.Pointer(&lpFormatStr[0])),
		uintptr(unsafe.Pointer(lpDurationStr)),
		uintptr(cchDuration),
		0,
		0)
	return int32(ret1)
}

func GetDurationFormatEx(lpLocaleName string, dwFlags uint32, lpDuration /*const*/ *SYSTEMTIME, ullDuration ULONGLONG, lpFormat string, lpDurationStr LPWSTR, cchDuration int32) int32 {
	lpLocaleNameStr := unicode16FromString(lpLocaleName)
	lpFormatStr := unicode16FromString(lpFormat)
	ret1 := syscall9(getDurationFormatEx, 7,
		uintptr(unsafe.Pointer(&lpLocaleNameStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpDuration)),
		uintptr(ullDuration),
		uintptr(unsafe.Pointer(&lpFormatStr[0])),
		uintptr(unsafe.Pointer(lpDurationStr)),
		uintptr(cchDuration),
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): PDYNAMIC_TIME_ZONE_INFORMATION
// func GetDynamicTimeZoneInformation(pTimeZoneInformation PDYNAMIC_TIME_ZONE_INFORMATION) uint32

// TODO: Unknown type(s): LPWCH
// func GetEnvironmentStrings() LPWCH

func GetEnvironmentVariable(lpName string, lpBuffer LPWSTR, nSize uint32) uint32 {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(getEnvironmentVariable, 3,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetErrorMode() uint32 {
	ret1 := syscall3(getErrorMode, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetExitCodeProcess(hProcess HANDLE, lpExitCode *uint32) bool {
	ret1 := syscall3(getExitCodeProcess, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpExitCode)),
		0)
	return ret1 != 0
}

func GetExitCodeThread(hThread HANDLE, lpExitCode *uint32) bool {
	ret1 := syscall3(getExitCodeThread, 2,
		uintptr(hThread),
		uintptr(unsafe.Pointer(lpExitCode)),
		0)
	return ret1 != 0
}

func GetExpandedName(unnamed0 LPWSTR, unnamed1 LPWSTR) int32 {
	ret1 := syscall3(getExpandedName, 2,
		uintptr(unsafe.Pointer(unnamed0)),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): GET_FILEEX_INFO_LEVELS
// func GetFileAttributesEx(lpFileName string, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation LPVOID) bool

// TODO: Unknown type(s): GET_FILEEX_INFO_LEVELS
// func GetFileAttributesTransacted(lpFileName string, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation LPVOID, hTransaction HANDLE) bool

func GetFileAttributes(lpFileName string) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(getFileAttributes, 1,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return uint32(ret1)
}

func GetFileBandwidthReservation(hFile HANDLE, lpPeriodMilliseconds *uint32, lpBytesPerPeriod *uint32, pDiscardable *BOOL, lpTransferSize *uint32, lpNumOutstandingRequests *uint32) bool {
	ret1 := syscall6(getFileBandwidthReservation, 6,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpPeriodMilliseconds)),
		uintptr(unsafe.Pointer(lpBytesPerPeriod)),
		uintptr(unsafe.Pointer(pDiscardable)),
		uintptr(unsafe.Pointer(lpTransferSize)),
		uintptr(unsafe.Pointer(lpNumOutstandingRequests)))
	return ret1 != 0
}

// TODO: Unknown type(s): LPBY_HANDLE_FILE_INFORMATION
// func GetFileInformationByHandle(hFile HANDLE, lpFileInformation LPBY_HANDLE_FILE_INFORMATION) bool

// TODO: Unknown type(s): FILE_INFO_BY_HANDLE_CLASS
// func GetFileInformationByHandleEx(hFile HANDLE, fileInformationClass FILE_INFO_BY_HANDLE_CLASS, lpFileInformation LPVOID, dwBufferSize uint32) bool

// TODO: Unknown type(s): PFILEMUIINFO
// func GetFileMUIInfo(dwFlags uint32, pcwszFilePath string, pFileMUIInfo PFILEMUIINFO, pcbFileMUIInfo *uint32) bool

// TODO: Unknown type(s): PULONGLONG
// func GetFileMUIPath(dwFlags uint32, pcwszFilePath string, pwszLanguage PWSTR, pcchLanguage *uint32, pwszFileMUIPath PWSTR, pcchFileMUIPath *uint32, pululEnumerator PULONGLONG) bool

func GetFileSize(hFile HANDLE, lpFileSizeHigh *uint32) uint32 {
	ret1 := syscall3(getFileSize, 2,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpFileSizeHigh)),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PLARGE_INTEGER
// func GetFileSizeEx(hFile HANDLE, lpFileSize PLARGE_INTEGER) bool

func GetFileTime(hFile HANDLE, lpCreationTime *FILETIME, lpLastAccessTime *FILETIME, lpLastWriteTime *FILETIME) bool {
	ret1 := syscall6(getFileTime, 4,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpLastAccessTime)),
		uintptr(unsafe.Pointer(lpLastWriteTime)),
		0,
		0)
	return ret1 != 0
}

func GetFileType(hFile HANDLE) uint32 {
	ret1 := syscall3(getFileType, 1,
		uintptr(hFile),
		0,
		0)
	return uint32(ret1)
}

func GetFinalPathNameByHandle(hFile HANDLE, lpszFilePath LPWSTR, cchFilePath uint32, dwFlags uint32) uint32 {
	ret1 := syscall6(getFinalPathNameByHandle, 4,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpszFilePath)),
		uintptr(cchFilePath),
		uintptr(dwFlags),
		0,
		0)
	return uint32(ret1)
}

func GetFirmwareEnvironmentVariable(lpName string, lpGuid string, pBuffer uintptr, nSize uint32) uint32 {
	lpNameStr := unicode16FromString(lpName)
	lpGuidStr := unicode16FromString(lpGuid)
	ret1 := syscall6(getFirmwareEnvironmentVariable, 4,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(&lpGuidStr[0])),
		pBuffer,
		uintptr(nSize),
		0,
		0)
	return uint32(ret1)
}

func GetFullPathNameTransacted(lpFileName string, nBufferLength uint32, lpBuffer LPWSTR, lpFilePart *LPWSTR, hTransaction HANDLE) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(getFullPathNameTransacted, 5,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(unsafe.Pointer(lpFilePart)),
		uintptr(hTransaction),
		0)
	return uint32(ret1)
}

func GetFullPathName(lpFileName string, nBufferLength uint32, lpBuffer LPWSTR, lpFilePart *LPWSTR) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(getFullPathName, 4,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(unsafe.Pointer(lpFilePart)),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): GEOTYPE
// func GetGeoInfo(location GEOID, geoType GEOTYPE, lpGeoData LPWSTR, cchData int32, langId LANGID) int32

func GetHandleInformation(hObject HANDLE, lpdwFlags *uint32) bool {
	ret1 := syscall3(getHandleInformation, 2,
		uintptr(hObject),
		uintptr(unsafe.Pointer(lpdwFlags)),
		0)
	return ret1 != 0
}

func GetLargePageMinimum() SIZE_T {
	ret1 := syscall3(getLargePageMinimum, 0,
		0,
		0,
		0)
	return SIZE_T(ret1)
}

func GetLargestConsoleWindowSize(hConsoleOutput HANDLE) COORD {
	ret1 := syscall3(getLargestConsoleWindowSize, 1,
		uintptr(hConsoleOutput),
		0,
		0)
	return getCOORDFromUintptr(ret1)
}

func GetLastError() uint32 {
	ret1 := syscall3(getLastError, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetLocalTime(lpSystemTime *SYSTEMTIME) {
	syscall3(getLocalTime, 1,
		uintptr(unsafe.Pointer(lpSystemTime)),
		0,
		0)
}

func GetLocaleInfoEx(lpLocaleName string, lCType LCTYPE, lpLCData LPWSTR, cchData int32) int32 {
	lpLocaleNameStr := unicode16FromString(lpLocaleName)
	ret1 := syscall6(getLocaleInfoEx, 4,
		uintptr(unsafe.Pointer(&lpLocaleNameStr[0])),
		uintptr(lCType),
		uintptr(unsafe.Pointer(lpLCData)),
		uintptr(cchData),
		0,
		0)
	return int32(ret1)
}

func GetLocaleInfo(locale LCID, lCType LCTYPE, lpLCData LPWSTR, cchData int32) int32 {
	ret1 := syscall6(getLocaleInfo, 4,
		uintptr(locale),
		uintptr(lCType),
		uintptr(unsafe.Pointer(lpLCData)),
		uintptr(cchData),
		0,
		0)
	return int32(ret1)
}

func GetLogicalDriveStrings(nBufferLength uint32, lpBuffer LPWSTR) uint32 {
	ret1 := syscall3(getLogicalDriveStrings, 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)
	return uint32(ret1)
}

func GetLogicalDrives() uint32 {
	ret1 := syscall3(getLogicalDrives, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PSYSTEM_LOGICAL_PROCESSOR_INFORMATION
// func GetLogicalProcessorInformation(buffer PSYSTEM_LOGICAL_PROCESSOR_INFORMATION, returnedLength *DWORD) bool

// TODO: Unknown type(s): LOGICAL_PROCESSOR_RELATIONSHIP, PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX
// func GetLogicalProcessorInformationEx(relationshipType LOGICAL_PROCESSOR_RELATIONSHIP, buffer PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX, returnedLength *DWORD) bool

func GetLongPathNameTransacted(lpszShortPath string, lpszLongPath LPWSTR, cchBuffer uint32, hTransaction HANDLE) uint32 {
	lpszShortPathStr := unicode16FromString(lpszShortPath)
	ret1 := syscall6(getLongPathNameTransacted, 4,
		uintptr(unsafe.Pointer(&lpszShortPathStr[0])),
		uintptr(unsafe.Pointer(lpszLongPath)),
		uintptr(cchBuffer),
		uintptr(hTransaction),
		0,
		0)
	return uint32(ret1)
}

func GetLongPathName(lpszShortPath string, lpszLongPath LPWSTR, cchBuffer uint32) uint32 {
	lpszShortPathStr := unicode16FromString(lpszShortPath)
	ret1 := syscall3(getLongPathName, 3,
		uintptr(unsafe.Pointer(&lpszShortPathStr[0])),
		uintptr(unsafe.Pointer(lpszLongPath)),
		uintptr(cchBuffer))
	return uint32(ret1)
}

func GetMailslotInfo(hMailslot HANDLE, lpMaxMessageSize *uint32, lpNextSize *uint32, lpMessageCount *uint32, lpReadTimeout *uint32) bool {
	ret1 := syscall6(getMailslotInfo, 5,
		uintptr(hMailslot),
		uintptr(unsafe.Pointer(lpMaxMessageSize)),
		uintptr(unsafe.Pointer(lpNextSize)),
		uintptr(unsafe.Pointer(lpMessageCount)),
		uintptr(unsafe.Pointer(lpReadTimeout)),
		0)
	return ret1 != 0
}

func GetMaximumProcessorCount(groupNumber uint16) uint32 {
	ret1 := syscall3(getMaximumProcessorCount, 1,
		uintptr(groupNumber),
		0,
		0)
	return uint32(ret1)
}

func GetMaximumProcessorGroupCount() uint16 {
	ret1 := syscall3(getMaximumProcessorGroupCount, 0,
		0,
		0,
		0)
	return uint16(ret1)
}

func GetModuleFileName(hModule HMODULE, lpFilename LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(getModuleFileName, 3,
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpFilename)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetModuleHandleEx(dwFlags uint32, lpModuleName string, phModule *HMODULE) bool {
	lpModuleNameStr := unicode16FromString(lpModuleName)
	ret1 := syscall3(getModuleHandleEx, 3,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpModuleNameStr[0])),
		uintptr(unsafe.Pointer(phModule)))
	return ret1 != 0
}

func GetModuleHandle(lpModuleName string) HMODULE {
	lpModuleNameStr := unicode16FromString(lpModuleName)
	ret1 := syscall3(getModuleHandle, 1,
		uintptr(unsafe.Pointer(&lpModuleNameStr[0])),
		0,
		0)
	return HMODULE(ret1)
}

// TODO: Unknown type(s): LPNLSVERSIONINFO, NLS_FUNCTION
// func GetNLSVersion(function NLS_FUNCTION, locale LCID, lpVersionInformation LPNLSVERSIONINFO) bool

// TODO: Unknown type(s): LPNLSVERSIONINFOEX, NLS_FUNCTION
// func GetNLSVersionEx(function NLS_FUNCTION, lpLocaleName string, lpVersionInformation LPNLSVERSIONINFOEX) bool

func GetNamedPipeClientComputerName(pipe HANDLE, clientComputerName LPWSTR, clientComputerNameLength ULONG) bool {
	ret1 := syscall3(getNamedPipeClientComputerName, 3,
		uintptr(pipe),
		uintptr(unsafe.Pointer(clientComputerName)),
		uintptr(clientComputerNameLength))
	return ret1 != 0
}

func GetNamedPipeClientProcessId(pipe HANDLE, clientProcessId *uint32) bool {
	ret1 := syscall3(getNamedPipeClientProcessId, 2,
		uintptr(pipe),
		uintptr(unsafe.Pointer(clientProcessId)),
		0)
	return ret1 != 0
}

func GetNamedPipeClientSessionId(pipe HANDLE, clientSessionId *uint32) bool {
	ret1 := syscall3(getNamedPipeClientSessionId, 2,
		uintptr(pipe),
		uintptr(unsafe.Pointer(clientSessionId)),
		0)
	return ret1 != 0
}

func GetNamedPipeHandleState(hNamedPipe HANDLE, lpState *uint32, lpCurInstances *uint32, lpMaxCollectionCount *uint32, lpCollectDataTimeout *uint32, lpUserName LPWSTR, nMaxUserNameSize uint32) bool {
	ret1 := syscall9(getNamedPipeHandleState, 7,
		uintptr(hNamedPipe),
		uintptr(unsafe.Pointer(lpState)),
		uintptr(unsafe.Pointer(lpCurInstances)),
		uintptr(unsafe.Pointer(lpMaxCollectionCount)),
		uintptr(unsafe.Pointer(lpCollectDataTimeout)),
		uintptr(unsafe.Pointer(lpUserName)),
		uintptr(nMaxUserNameSize),
		0,
		0)
	return ret1 != 0
}

func GetNamedPipeInfo(hNamedPipe HANDLE, lpFlags *uint32, lpOutBufferSize *uint32, lpInBufferSize *uint32, lpMaxInstances *uint32) bool {
	ret1 := syscall6(getNamedPipeInfo, 5,
		uintptr(hNamedPipe),
		uintptr(unsafe.Pointer(lpFlags)),
		uintptr(unsafe.Pointer(lpOutBufferSize)),
		uintptr(unsafe.Pointer(lpInBufferSize)),
		uintptr(unsafe.Pointer(lpMaxInstances)),
		0)
	return ret1 != 0
}

func GetNamedPipeServerProcessId(pipe HANDLE, serverProcessId *uint32) bool {
	ret1 := syscall3(getNamedPipeServerProcessId, 2,
		uintptr(pipe),
		uintptr(unsafe.Pointer(serverProcessId)),
		0)
	return ret1 != 0
}

func GetNamedPipeServerSessionId(pipe HANDLE, serverSessionId *uint32) bool {
	ret1 := syscall3(getNamedPipeServerSessionId, 2,
		uintptr(pipe),
		uintptr(unsafe.Pointer(serverSessionId)),
		0)
	return ret1 != 0
}

func GetNativeSystemInfo(lpSystemInfo *SYSTEM_INFO) {
	syscall3(getNativeSystemInfo, 1,
		uintptr(unsafe.Pointer(lpSystemInfo)),
		0,
		0)
}

// TODO: Unknown type(s): PULONGLONG
// func GetNumaAvailableMemoryNode(node UCHAR, availableBytes PULONGLONG) bool

// TODO: Unknown type(s): PULONGLONG
// func GetNumaAvailableMemoryNodeEx(node USHORT, availableBytes PULONGLONG) bool

func GetNumaHighestNodeNumber(highestNodeNumber *uint32) bool {
	ret1 := syscall3(getNumaHighestNodeNumber, 1,
		uintptr(unsafe.Pointer(highestNodeNumber)),
		0,
		0)
	return ret1 != 0
}

func GetNumaNodeNumberFromHandle(hFile HANDLE, nodeNumber PUSHORT) bool {
	ret1 := syscall3(getNumaNodeNumberFromHandle, 2,
		uintptr(hFile),
		uintptr(unsafe.Pointer(nodeNumber)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PULONGLONG
// func GetNumaNodeProcessorMask(node UCHAR, processorMask PULONGLONG) bool

// TODO: Unknown type(s): PGROUP_AFFINITY
// func GetNumaNodeProcessorMaskEx(node USHORT, processorMask PGROUP_AFFINITY) bool

// TODO: Unknown type(s): PUCHAR
// func GetNumaProcessorNode(processor UCHAR, nodeNumber PUCHAR) bool

// TODO: Unknown type(s): PPROCESSOR_NUMBER
// func GetNumaProcessorNodeEx(processor PPROCESSOR_NUMBER, nodeNumber PUSHORT) bool

// TODO: Unknown type(s): PUCHAR
// func GetNumaProximityNode(proximityId ULONG, nodeNumber PUCHAR) bool

func GetNumaProximityNodeEx(proximityId ULONG, nodeNumber PUSHORT) bool {
	ret1 := syscall3(getNumaProximityNodeEx, 2,
		uintptr(proximityId),
		uintptr(unsafe.Pointer(nodeNumber)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CONST NUMBERFMTW *
// func GetNumberFormatEx(lpLocaleName string, dwFlags uint32, lpValue string, lpFormat /*const*/ CONST NUMBERFMTW *, lpNumberStr LPWSTR, cchNumber int32) int32

// TODO: Unknown type(s): CONST NUMBERFMTW *
// func GetNumberFormat(locale LCID, dwFlags uint32, lpValue string, lpFormat /*const*/ CONST NUMBERFMTW *, lpNumberStr LPWSTR, cchNumber int32) int32

func GetNumberOfConsoleInputEvents(hConsoleInput HANDLE, lpNumberOfEvents *uint32) bool {
	ret1 := syscall3(getNumberOfConsoleInputEvents, 2,
		uintptr(hConsoleInput),
		uintptr(unsafe.Pointer(lpNumberOfEvents)),
		0)
	return ret1 != 0
}

func GetNumberOfConsoleMouseButtons(lpNumberOfMouseButtons *uint32) bool {
	ret1 := syscall3(getNumberOfConsoleMouseButtons, 1,
		uintptr(unsafe.Pointer(lpNumberOfMouseButtons)),
		0,
		0)
	return ret1 != 0
}

func GetOEMCP() uint32 {
	ret1 := syscall3(getOEMCP, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetOverlappedResult(hFile HANDLE, lpOverlapped *OVERLAPPED, lpNumberOfBytesTransferred *uint32, bWait bool) bool {
	ret1 := syscall6(getOverlappedResult, 4,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpOverlapped)),
		uintptr(unsafe.Pointer(lpNumberOfBytesTransferred)),
		getUintptrFromBool(bWait),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PULONGLONG
// func GetPhysicallyInstalledSystemMemory(totalMemoryInKilobytes PULONGLONG) bool

func GetPriorityClass(hProcess HANDLE) uint32 {
	ret1 := syscall3(getPriorityClass, 1,
		uintptr(hProcess),
		0,
		0)
	return uint32(ret1)
}

func GetPrivateProfileInt(lpAppName string, lpKeyName string, nDefault int32, lpFileName string) uint32 {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpKeyNameStr := unicode16FromString(lpKeyName)
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(getPrivateProfileInt, 4,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpKeyNameStr[0])),
		uintptr(nDefault),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return uint32(ret1)
}

func GetPrivateProfileSectionNames(lpszReturnBuffer LPWSTR, nSize uint32, lpFileName string) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(getPrivateProfileSectionNames, 3,
		uintptr(unsafe.Pointer(lpszReturnBuffer)),
		uintptr(nSize),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])))
	return uint32(ret1)
}

func GetPrivateProfileSection(lpAppName string, lpReturnedString LPWSTR, nSize uint32, lpFileName string) uint32 {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(getPrivateProfileSection, 4,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(lpReturnedString)),
		uintptr(nSize),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return uint32(ret1)
}

func GetPrivateProfileString(lpAppName string, lpKeyName string, lpDefault string, lpReturnedString LPWSTR, nSize uint32, lpFileName string) uint32 {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpKeyNameStr := unicode16FromString(lpKeyName)
	lpDefaultStr := unicode16FromString(lpDefault)
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(getPrivateProfileString, 6,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpKeyNameStr[0])),
		uintptr(unsafe.Pointer(&lpDefaultStr[0])),
		uintptr(unsafe.Pointer(lpReturnedString)),
		uintptr(nSize),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])))
	return uint32(ret1)
}

func GetPrivateProfileStruct(lpszSection string, lpszKey string, lpStruct LPVOID, uSizeStruct uint32, szFile string) bool {
	lpszSectionStr := unicode16FromString(lpszSection)
	lpszKeyStr := unicode16FromString(lpszKey)
	szFileStr := unicode16FromString(szFile)
	ret1 := syscall6(getPrivateProfileStruct, 5,
		uintptr(unsafe.Pointer(&lpszSectionStr[0])),
		uintptr(unsafe.Pointer(&lpszKeyStr[0])),
		uintptr(unsafe.Pointer(lpStruct)),
		uintptr(uSizeStruct),
		uintptr(unsafe.Pointer(&szFileStr[0])),
		0)
	return ret1 != 0
}

func GetProcAddress(hModule HMODULE, lpProcName /*const*/ LPCSTR) FARPROC {
	ret1 := syscall3(getProcAddress, 2,
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpProcName)),
		0)
	return func() INT_PTR {
		ret2 := syscall3(ret1, 0,
			0,
			0,
			0)
		return (INT_PTR)(unsafe.Pointer(ret2))
	}
}

func GetProcessAffinityMask(hProcess HANDLE, lpProcessAffinityMask *uintptr, lpSystemAffinityMask *uintptr) bool {
	ret1 := syscall3(getProcessAffinityMask, 3,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpProcessAffinityMask)),
		uintptr(unsafe.Pointer(lpSystemAffinityMask)))
	return ret1 != 0
}

func GetProcessDEPPolicy(hProcess HANDLE, lpFlags *uint32, lpPermanent *BOOL) bool {
	ret1 := syscall3(getProcessDEPPolicy, 3,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpFlags)),
		uintptr(unsafe.Pointer(lpPermanent)))
	return ret1 != 0
}

func GetProcessGroupAffinity(hProcess HANDLE, groupCount PUSHORT, groupArray PUSHORT) bool {
	ret1 := syscall3(getProcessGroupAffinity, 3,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(groupCount)),
		uintptr(unsafe.Pointer(groupArray)))
	return ret1 != 0
}

func GetProcessHandleCount(hProcess HANDLE, pdwHandleCount *DWORD) bool {
	ret1 := syscall3(getProcessHandleCount, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(pdwHandleCount)),
		0)
	return ret1 != 0
}

func GetProcessHeap() HANDLE {
	ret1 := syscall3(getProcessHeap, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

func GetProcessHeaps(numberOfHeaps uint32, processHeaps *HANDLE) uint32 {
	ret1 := syscall3(getProcessHeaps, 2,
		uintptr(numberOfHeaps),
		uintptr(unsafe.Pointer(processHeaps)),
		0)
	return uint32(ret1)
}

func GetProcessId(process HANDLE) uint32 {
	ret1 := syscall3(getProcessId, 1,
		uintptr(process),
		0,
		0)
	return uint32(ret1)
}

func GetProcessIdOfThread(thread HANDLE) uint32 {
	ret1 := syscall3(getProcessIdOfThread, 1,
		uintptr(thread),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PIO_COUNTERS
// func GetProcessIoCounters(hProcess HANDLE, lpIoCounters PIO_COUNTERS) bool

// TODO: Unknown type(s): PZZWSTR
// func GetProcessPreferredUILanguages(dwFlags uint32, pulNumLanguages *uint32, pwszLanguagesBuffer PZZWSTR, pcchLanguagesBuffer *uint32) bool

func GetProcessPriorityBoost(hProcess HANDLE, pDisablePriorityBoost *BOOL) bool {
	ret1 := syscall3(getProcessPriorityBoost, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(pDisablePriorityBoost)),
		0)
	return ret1 != 0
}

func GetProcessShutdownParameters(lpdwLevel *uint32, lpdwFlags *uint32) bool {
	ret1 := syscall3(getProcessShutdownParameters, 2,
		uintptr(unsafe.Pointer(lpdwLevel)),
		uintptr(unsafe.Pointer(lpdwFlags)),
		0)
	return ret1 != 0
}

func GetProcessTimes(hProcess HANDLE, lpCreationTime *FILETIME, lpExitTime *FILETIME, lpKernelTime *FILETIME, lpUserTime *FILETIME) bool {
	ret1 := syscall6(getProcessTimes, 5,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpExitTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)),
		0)
	return ret1 != 0
}

func GetProcessVersion(processId uint32) uint32 {
	ret1 := syscall3(getProcessVersion, 1,
		uintptr(processId),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PSIZE_T
// func GetProcessWorkingSetSize(hProcess HANDLE, lpMinimumWorkingSetSize PSIZE_T, lpMaximumWorkingSetSize PSIZE_T) bool

// TODO: Unknown type(s): PSIZE_T
// func GetProcessWorkingSetSizeEx(hProcess HANDLE, lpMinimumWorkingSetSize PSIZE_T, lpMaximumWorkingSetSize PSIZE_T, flags *DWORD) bool

// TODO: Unknown type(s): PSYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION
// func GetProcessorSystemCycleTime(group USHORT, buffer PSYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION, returnedLength *DWORD) bool

func GetProductInfo(dwOSMajorVersion uint32, dwOSMinorVersion uint32, dwSpMajorVersion uint32, dwSpMinorVersion uint32, pdwReturnedProductType *DWORD) bool {
	ret1 := syscall6(getProductInfo, 5,
		uintptr(dwOSMajorVersion),
		uintptr(dwOSMinorVersion),
		uintptr(dwSpMajorVersion),
		uintptr(dwSpMinorVersion),
		uintptr(unsafe.Pointer(pdwReturnedProductType)),
		0)
	return ret1 != 0
}

func GetProfileInt(lpAppName string, lpKeyName string, nDefault int32) uint32 {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpKeyNameStr := unicode16FromString(lpKeyName)
	ret1 := syscall3(getProfileInt, 3,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpKeyNameStr[0])),
		uintptr(nDefault))
	return uint32(ret1)
}

func GetProfileSection(lpAppName string, lpReturnedString LPWSTR, nSize uint32) uint32 {
	lpAppNameStr := unicode16FromString(lpAppName)
	ret1 := syscall3(getProfileSection, 3,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(lpReturnedString)),
		uintptr(nSize))
	return uint32(ret1)
}

func GetProfileString(lpAppName string, lpKeyName string, lpDefault string, lpReturnedString LPWSTR, nSize uint32) uint32 {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpKeyNameStr := unicode16FromString(lpKeyName)
	lpDefaultStr := unicode16FromString(lpDefault)
	ret1 := syscall6(getProfileString, 5,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpKeyNameStr[0])),
		uintptr(unsafe.Pointer(&lpDefaultStr[0])),
		uintptr(unsafe.Pointer(lpReturnedString)),
		uintptr(nSize),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PULONG_PTR
// func GetQueuedCompletionStatus(completionPort HANDLE, lpNumberOfBytesTransferred *uint32, lpCompletionKey PULONG_PTR, lpOverlapped *LPOVERLAPPED, dwMilliseconds uint32) bool

// TODO: Unknown type(s): LPOVERLAPPED_ENTRY
// func GetQueuedCompletionStatusEx(completionPort HANDLE, lpCompletionPortEntries LPOVERLAPPED_ENTRY, ulCount ULONG, ulNumEntriesRemoved *uint32, dwMilliseconds uint32, fAlertable bool) bool

func GetShortPathName(lpszLongPath string, lpszShortPath LPWSTR, cchBuffer uint32) uint32 {
	lpszLongPathStr := unicode16FromString(lpszLongPath)
	ret1 := syscall3(getShortPathName, 3,
		uintptr(unsafe.Pointer(&lpszLongPathStr[0])),
		uintptr(unsafe.Pointer(lpszShortPath)),
		uintptr(cchBuffer))
	return uint32(ret1)
}

func GetStartupInfo(lpStartupInfo *STARTUPINFO) {
	syscall3(getStartupInfo, 1,
		uintptr(unsafe.Pointer(lpStartupInfo)),
		0,
		0)
}

func GetStdHandle(nStdHandle uint32) HANDLE {
	ret1 := syscall3(getStdHandle, 1,
		uintptr(nStdHandle),
		0,
		0)
	return HANDLE(ret1)
}

func GetStringScripts(dwFlags uint32, lpString string, cchString int32, lpScripts LPWSTR, cchScripts int32) int32 {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall6(getStringScripts, 5,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(cchString),
		uintptr(unsafe.Pointer(lpScripts)),
		uintptr(cchScripts),
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): LPCWCH
// func GetStringTypeEx(locale LCID, dwInfoType uint32, lpSrcStr LPCWCH, cchSrc int32, lpCharType *uint16) bool

// TODO: Unknown type(s): LPCWCH
// func GetStringType(dwInfoType uint32, lpSrcStr LPCWCH, cchSrc int32, lpCharType *uint16) bool

// TODO: Unknown type(s): DEP_SYSTEM_POLICY_TYPE
// func GetSystemDEPPolicy() DEP_SYSTEM_POLICY_TYPE

func GetSystemDefaultLCID() LCID {
	ret1 := syscall3(getSystemDefaultLCID, 0,
		0,
		0,
		0)
	return LCID(ret1)
}

func GetSystemDefaultLangID() LANGID {
	ret1 := syscall3(getSystemDefaultLangID, 0,
		0,
		0,
		0)
	return LANGID(ret1)
}

func GetSystemDefaultLocaleName(lpLocaleName LPWSTR, cchLocaleName int32) int32 {
	ret1 := syscall3(getSystemDefaultLocaleName, 2,
		uintptr(unsafe.Pointer(lpLocaleName)),
		uintptr(cchLocaleName),
		0)
	return int32(ret1)
}

func GetSystemDefaultUILanguage() LANGID {
	ret1 := syscall3(getSystemDefaultUILanguage, 0,
		0,
		0,
		0)
	return LANGID(ret1)
}

func GetSystemDirectory(lpBuffer LPWSTR, uSize uint32) uint32 {
	ret1 := syscall3(getSystemDirectory, 2,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(uSize),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PSIZE_T
// func GetSystemFileCacheSize(lpMinimumFileCacheSize PSIZE_T, lpMaximumFileCacheSize PSIZE_T, lpFlags *DWORD) bool

func GetSystemFirmwareTable(firmwareTableProviderSignature uint32, firmwareTableID uint32, pFirmwareTableBuffer uintptr, bufferSize uint32) uint32 {
	ret1 := syscall6(getSystemFirmwareTable, 4,
		uintptr(firmwareTableProviderSignature),
		uintptr(firmwareTableID),
		pFirmwareTableBuffer,
		uintptr(bufferSize),
		0,
		0)
	return uint32(ret1)
}

func GetSystemInfo(lpSystemInfo *SYSTEM_INFO) {
	syscall3(getSystemInfo, 1,
		uintptr(unsafe.Pointer(lpSystemInfo)),
		0,
		0)
}

// TODO: Unknown type(s): LPSYSTEM_POWER_STATUS
// func GetSystemPowerStatus(lpSystemPowerStatus LPSYSTEM_POWER_STATUS) bool

// TODO: Unknown type(s): PZZWSTR
// func GetSystemPreferredUILanguages(dwFlags uint32, pulNumLanguages *uint32, pwszLanguagesBuffer PZZWSTR, pcchLanguagesBuffer *uint32) bool

func GetSystemRegistryQuota(pdwQuotaAllowed *DWORD, pdwQuotaUsed *DWORD) bool {
	ret1 := syscall3(getSystemRegistryQuota, 2,
		uintptr(unsafe.Pointer(pdwQuotaAllowed)),
		uintptr(unsafe.Pointer(pdwQuotaUsed)),
		0)
	return ret1 != 0
}

func GetSystemTime(lpSystemTime *SYSTEMTIME) {
	syscall3(getSystemTime, 1,
		uintptr(unsafe.Pointer(lpSystemTime)),
		0,
		0)
}

func GetSystemTimeAdjustment(lpTimeAdjustment *DWORD, lpTimeIncrement *DWORD, lpTimeAdjustmentDisabled *BOOL) bool {
	ret1 := syscall3(getSystemTimeAdjustment, 3,
		uintptr(unsafe.Pointer(lpTimeAdjustment)),
		uintptr(unsafe.Pointer(lpTimeIncrement)),
		uintptr(unsafe.Pointer(lpTimeAdjustmentDisabled)))
	return ret1 != 0
}

func GetSystemTimeAsFileTime(lpSystemTimeAsFileTime *FILETIME) {
	syscall3(getSystemTimeAsFileTime, 1,
		uintptr(unsafe.Pointer(lpSystemTimeAsFileTime)),
		0,
		0)
}

func GetSystemTimePreciseAsFileTime(lpSystemTimeAsFileTime *FILETIME) {
	syscall3(getSystemTimePreciseAsFileTime, 1,
		uintptr(unsafe.Pointer(lpSystemTimeAsFileTime)),
		0,
		0)
}

func GetSystemTimes(lpIdleTime *FILETIME, lpKernelTime *FILETIME, lpUserTime *FILETIME) bool {
	ret1 := syscall3(getSystemTimes, 3,
		uintptr(unsafe.Pointer(lpIdleTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)))
	return ret1 != 0
}

func GetSystemWindowsDirectory(lpBuffer LPWSTR, uSize uint32) uint32 {
	ret1 := syscall3(getSystemWindowsDirectory, 2,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(uSize),
		0)
	return uint32(ret1)
}

func GetSystemWow64Directory(lpBuffer LPWSTR, uSize uint32) uint32 {
	ret1 := syscall3(getSystemWow64Directory, 2,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(uSize),
		0)
	return uint32(ret1)
}

func GetTapeParameters(hDevice HANDLE, dwOperation uint32, lpdwSize *uint32, lpTapeInformation LPVOID) uint32 {
	ret1 := syscall6(getTapeParameters, 4,
		uintptr(hDevice),
		uintptr(dwOperation),
		uintptr(unsafe.Pointer(lpdwSize)),
		uintptr(unsafe.Pointer(lpTapeInformation)),
		0,
		0)
	return uint32(ret1)
}

func GetTapePosition(hDevice HANDLE, dwPositionType uint32, lpdwPartition *uint32, lpdwOffsetLow *uint32, lpdwOffsetHigh *uint32) uint32 {
	ret1 := syscall6(getTapePosition, 5,
		uintptr(hDevice),
		uintptr(dwPositionType),
		uintptr(unsafe.Pointer(lpdwPartition)),
		uintptr(unsafe.Pointer(lpdwOffsetLow)),
		uintptr(unsafe.Pointer(lpdwOffsetHigh)),
		0)
	return uint32(ret1)
}

func GetTapeStatus(hDevice HANDLE) uint32 {
	ret1 := syscall3(getTapeStatus, 1,
		uintptr(hDevice),
		0,
		0)
	return uint32(ret1)
}

func GetTempFileName(lpPathName string, lpPrefixString string, uUnique uint32, lpTempFileName LPWSTR) uint32 {
	lpPathNameStr := unicode16FromString(lpPathName)
	lpPrefixStringStr := unicode16FromString(lpPrefixString)
	ret1 := syscall6(getTempFileName, 4,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		uintptr(unsafe.Pointer(&lpPrefixStringStr[0])),
		uintptr(uUnique),
		uintptr(unsafe.Pointer(lpTempFileName)),
		0,
		0)
	return uint32(ret1)
}

func GetTempPath(nBufferLength uint32, lpBuffer LPWSTR) uint32 {
	ret1 := syscall3(getTempPath, 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): LPCONTEXT
// func GetThreadContext(hThread HANDLE, lpContext LPCONTEXT) bool

func GetThreadErrorMode() uint32 {
	ret1 := syscall3(getThreadErrorMode, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PGROUP_AFFINITY
// func GetThreadGroupAffinity(hThread HANDLE, groupAffinity PGROUP_AFFINITY) bool

func GetThreadIOPendingFlag(hThread HANDLE, lpIOIsPending *BOOL) bool {
	ret1 := syscall3(getThreadIOPendingFlag, 2,
		uintptr(hThread),
		uintptr(unsafe.Pointer(lpIOIsPending)),
		0)
	return ret1 != 0
}

func GetThreadId(thread HANDLE) uint32 {
	ret1 := syscall3(getThreadId, 1,
		uintptr(thread),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PPROCESSOR_NUMBER
// func GetThreadIdealProcessorEx(hThread HANDLE, lpIdealProcessor PPROCESSOR_NUMBER) bool

func GetThreadLocale() LCID {
	ret1 := syscall3(getThreadLocale, 0,
		0,
		0,
		0)
	return LCID(ret1)
}

// TODO: Unknown type(s): PZZWSTR
// func GetThreadPreferredUILanguages(dwFlags uint32, pulNumLanguages *uint32, pwszLanguagesBuffer PZZWSTR, pcchLanguagesBuffer *uint32) bool

func GetThreadPriority(hThread HANDLE) int32 {
	ret1 := syscall3(getThreadPriority, 1,
		uintptr(hThread),
		0,
		0)
	return int32(ret1)
}

func GetThreadPriorityBoost(hThread HANDLE, pDisablePriorityBoost *BOOL) bool {
	ret1 := syscall3(getThreadPriorityBoost, 2,
		uintptr(hThread),
		uintptr(unsafe.Pointer(pDisablePriorityBoost)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPLDT_ENTRY
// func GetThreadSelectorEntry(hThread HANDLE, dwSelector uint32, lpSelectorEntry LPLDT_ENTRY) bool

func GetThreadTimes(hThread HANDLE, lpCreationTime *FILETIME, lpExitTime *FILETIME, lpKernelTime *FILETIME, lpUserTime *FILETIME) bool {
	ret1 := syscall6(getThreadTimes, 5,
		uintptr(hThread),
		uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpExitTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)),
		0)
	return ret1 != 0
}

func GetThreadUILanguage() LANGID {
	ret1 := syscall3(getThreadUILanguage, 0,
		0,
		0,
		0)
	return LANGID(ret1)
}

func GetTickCount() uint32 {
	ret1 := syscall3(getTickCount, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GetTickCount64() ULONGLONG {
	ret1 := syscall3(getTickCount64, 0,
		0,
		0,
		0)
	return ULONGLONG(ret1)
}

func GetTimeFormatEx(lpLocaleName string, dwFlags uint32, lpTime /*const*/ *SYSTEMTIME, lpFormat string, lpTimeStr LPWSTR, cchTime int32) int32 {
	lpLocaleNameStr := unicode16FromString(lpLocaleName)
	lpFormatStr := unicode16FromString(lpFormat)
	ret1 := syscall6(getTimeFormatEx, 6,
		uintptr(unsafe.Pointer(&lpLocaleNameStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpTime)),
		uintptr(unsafe.Pointer(&lpFormatStr[0])),
		uintptr(unsafe.Pointer(lpTimeStr)),
		uintptr(cchTime))
	return int32(ret1)
}

func GetTimeFormat(locale LCID, dwFlags uint32, lpTime /*const*/ *SYSTEMTIME, lpFormat string, lpTimeStr LPWSTR, cchTime int32) int32 {
	lpFormatStr := unicode16FromString(lpFormat)
	ret1 := syscall6(getTimeFormat, 6,
		uintptr(locale),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpTime)),
		uintptr(unsafe.Pointer(&lpFormatStr[0])),
		uintptr(unsafe.Pointer(lpTimeStr)),
		uintptr(cchTime))
	return int32(ret1)
}

// TODO: Unknown type(s): LPTIME_ZONE_INFORMATION
// func GetTimeZoneInformation(lpTimeZoneInformation LPTIME_ZONE_INFORMATION) uint32

// TODO: Unknown type(s): LPTIME_ZONE_INFORMATION, PDYNAMIC_TIME_ZONE_INFORMATION
// func GetTimeZoneInformationForYear(wYear USHORT, pdtzi PDYNAMIC_TIME_ZONE_INFORMATION, ptzi LPTIME_ZONE_INFORMATION) bool

// TODO: Unknown type(s): PCZZWSTR, PZZWSTR
// func GetUILanguageInfo(dwFlags uint32, pwmszLanguage PCZZWSTR, pwszFallbackLanguages PZZWSTR, pcchFallbackLanguages *DWORD, pAttributes *DWORD) bool

func GetUserDefaultLCID() LCID {
	ret1 := syscall3(getUserDefaultLCID, 0,
		0,
		0,
		0)
	return LCID(ret1)
}

func GetUserDefaultLangID() LANGID {
	ret1 := syscall3(getUserDefaultLangID, 0,
		0,
		0,
		0)
	return LANGID(ret1)
}

func GetUserDefaultLocaleName(lpLocaleName LPWSTR, cchLocaleName int32) int32 {
	ret1 := syscall3(getUserDefaultLocaleName, 2,
		uintptr(unsafe.Pointer(lpLocaleName)),
		uintptr(cchLocaleName),
		0)
	return int32(ret1)
}

func GetUserDefaultUILanguage() LANGID {
	ret1 := syscall3(getUserDefaultUILanguage, 0,
		0,
		0,
		0)
	return LANGID(ret1)
}

// TODO: Unknown type(s): GEOCLASS
// func GetUserGeoID(geoClass GEOCLASS) GEOID

// TODO: Unknown type(s): PZZWSTR
// func GetUserPreferredUILanguages(dwFlags uint32, pulNumLanguages *uint32, pwszLanguagesBuffer PZZWSTR, pcchLanguagesBuffer *uint32) bool

func GetVersion() uint32 {
	ret1 := syscall3(getVersion, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): LPOSVERSIONINFOW
// func GetVersionEx(lpVersionInformation LPOSVERSIONINFOW) bool

func GetVolumeInformationByHandleW(hFile HANDLE, lpVolumeNameBuffer LPWSTR, nVolumeNameSize uint32, lpVolumeSerialNumber *uint32, lpMaximumComponentLength *uint32, lpFileSystemFlags *uint32, lpFileSystemNameBuffer LPWSTR, nFileSystemNameSize uint32) bool {
	ret1 := syscall9(getVolumeInformationByHandleW, 8,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpVolumeNameBuffer)),
		uintptr(nVolumeNameSize),
		uintptr(unsafe.Pointer(lpVolumeSerialNumber)),
		uintptr(unsafe.Pointer(lpMaximumComponentLength)),
		uintptr(unsafe.Pointer(lpFileSystemFlags)),
		uintptr(unsafe.Pointer(lpFileSystemNameBuffer)),
		uintptr(nFileSystemNameSize),
		0)
	return ret1 != 0
}

func GetVolumeInformation(lpRootPathName string, lpVolumeNameBuffer LPWSTR, nVolumeNameSize uint32, lpVolumeSerialNumber *uint32, lpMaximumComponentLength *uint32, lpFileSystemFlags *uint32, lpFileSystemNameBuffer LPWSTR, nFileSystemNameSize uint32) bool {
	lpRootPathNameStr := unicode16FromString(lpRootPathName)
	ret1 := syscall9(getVolumeInformation, 8,
		uintptr(unsafe.Pointer(&lpRootPathNameStr[0])),
		uintptr(unsafe.Pointer(lpVolumeNameBuffer)),
		uintptr(nVolumeNameSize),
		uintptr(unsafe.Pointer(lpVolumeSerialNumber)),
		uintptr(unsafe.Pointer(lpMaximumComponentLength)),
		uintptr(unsafe.Pointer(lpFileSystemFlags)),
		uintptr(unsafe.Pointer(lpFileSystemNameBuffer)),
		uintptr(nFileSystemNameSize),
		0)
	return ret1 != 0
}

func GetVolumeNameForVolumeMountPoint(lpszVolumeMountPoint string, lpszVolumeName LPWSTR, cchBufferLength uint32) bool {
	lpszVolumeMountPointStr := unicode16FromString(lpszVolumeMountPoint)
	ret1 := syscall3(getVolumeNameForVolumeMountPoint, 3,
		uintptr(unsafe.Pointer(&lpszVolumeMountPointStr[0])),
		uintptr(unsafe.Pointer(lpszVolumeName)),
		uintptr(cchBufferLength))
	return ret1 != 0
}

func GetVolumePathName(lpszFileName string, lpszVolumePathName LPWSTR, cchBufferLength uint32) bool {
	lpszFileNameStr := unicode16FromString(lpszFileName)
	ret1 := syscall3(getVolumePathName, 3,
		uintptr(unsafe.Pointer(&lpszFileNameStr[0])),
		uintptr(unsafe.Pointer(lpszVolumePathName)),
		uintptr(cchBufferLength))
	return ret1 != 0
}

// TODO: Unknown type(s): LPWCH
// func GetVolumePathNamesForVolumeName(lpszVolumeName string, lpszVolumePathNames LPWCH, cchBufferLength uint32, lpcchReturnLength *DWORD) bool

func GetWindowsDirectory(lpBuffer LPWSTR, uSize uint32) uint32 {
	ret1 := syscall3(getWindowsDirectory, 2,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(uSize),
		0)
	return uint32(ret1)
}

func GetWriteWatch(dwFlags uint32, lpBaseAddress uintptr, dwRegionSize SIZE_T, lpAddresses *PVOID, lpdwCount *ULONG_PTR, lpdwGranularity *uint32) uint32 {
	ret1 := syscall6(getWriteWatch, 6,
		uintptr(dwFlags),
		lpBaseAddress,
		uintptr(dwRegionSize),
		uintptr(unsafe.Pointer(lpAddresses)),
		uintptr(unsafe.Pointer(lpdwCount)),
		uintptr(unsafe.Pointer(lpdwGranularity)))
	return uint32(ret1)
}

func GlobalAddAtom(lpString string) ATOM {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(globalAddAtom, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return ATOM(ret1)
}

func GlobalAlloc(uFlags uint32, dwBytes SIZE_T) HGLOBAL {
	ret1 := syscall3(globalAlloc, 2,
		uintptr(uFlags),
		uintptr(dwBytes),
		0)
	return HGLOBAL(ret1)
}

func GlobalCompact(dwMinFree uint32) SIZE_T {
	ret1 := syscall3(globalCompact, 1,
		uintptr(dwMinFree),
		0,
		0)
	return SIZE_T(ret1)
}

func GlobalDeleteAtom(nAtom ATOM) ATOM {
	ret1 := syscall3(globalDeleteAtom, 1,
		uintptr(nAtom),
		0,
		0)
	return ATOM(ret1)
}

func GlobalFindAtom(lpString string) ATOM {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(globalFindAtom, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return ATOM(ret1)
}

func GlobalFix(hMem HGLOBAL) {
	syscall3(globalFix, 1,
		uintptr(hMem),
		0,
		0)
}

func GlobalFlags(hMem HGLOBAL) uint32 {
	ret1 := syscall3(globalFlags, 1,
		uintptr(hMem),
		0,
		0)
	return uint32(ret1)
}

func GlobalFree(hMem HGLOBAL) HGLOBAL {
	ret1 := syscall3(globalFree, 1,
		uintptr(hMem),
		0,
		0)
	return HGLOBAL(ret1)
}

func GlobalGetAtomName(nAtom ATOM, lpBuffer LPWSTR, nSize int32) uint32 {
	ret1 := syscall3(globalGetAtomName, 3,
		uintptr(nAtom),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize))
	return uint32(ret1)
}

func GlobalHandle(pMem /*const*/ uintptr) HGLOBAL {
	ret1 := syscall3(globalHandle, 1,
		pMem,
		0,
		0)
	return HGLOBAL(ret1)
}

func GlobalLock(hMem HGLOBAL) LPVOID {
	ret1 := syscall3(globalLock, 1,
		uintptr(hMem),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): LPMEMORYSTATUS
// func GlobalMemoryStatus(lpBuffer LPMEMORYSTATUS)

// TODO: Unknown type(s): LPMEMORYSTATUSEX
// func GlobalMemoryStatusEx(lpBuffer LPMEMORYSTATUSEX) bool

func GlobalReAlloc(hMem HGLOBAL, dwBytes SIZE_T, uFlags uint32) HGLOBAL {
	ret1 := syscall3(globalReAlloc, 3,
		uintptr(hMem),
		uintptr(dwBytes),
		uintptr(uFlags))
	return HGLOBAL(ret1)
}

func GlobalSize(hMem HGLOBAL) SIZE_T {
	ret1 := syscall3(globalSize, 1,
		uintptr(hMem),
		0,
		0)
	return SIZE_T(ret1)
}

func GlobalUnWire(hMem HGLOBAL) bool {
	ret1 := syscall3(globalUnWire, 1,
		uintptr(hMem),
		0,
		0)
	return ret1 != 0
}

func GlobalUnfix(hMem HGLOBAL) {
	syscall3(globalUnfix, 1,
		uintptr(hMem),
		0,
		0)
}

func GlobalUnlock(hMem HGLOBAL) bool {
	ret1 := syscall3(globalUnlock, 1,
		uintptr(hMem),
		0,
		0)
	return ret1 != 0
}

func GlobalWire(hMem HGLOBAL) LPVOID {
	ret1 := syscall3(globalWire, 1,
		uintptr(hMem),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): LPHEAPENTRY32
// func Heap32First(lphe LPHEAPENTRY32, th32ProcessID uint32, th32HeapID *uint32) bool

// TODO: Unknown type(s): LPHEAPLIST32
// func Heap32ListFirst(hSnapshot HANDLE, lphl LPHEAPLIST32) bool

// TODO: Unknown type(s): LPHEAPLIST32
// func Heap32ListNext(hSnapshot HANDLE, lphl LPHEAPLIST32) bool

// TODO: Unknown type(s): LPHEAPENTRY32
// func Heap32Next(lphe LPHEAPENTRY32) bool

func HeapAlloc(hHeap HANDLE, dwFlags uint32, dwBytes SIZE_T) LPVOID {
	ret1 := syscall3(heapAlloc, 3,
		uintptr(hHeap),
		uintptr(dwFlags),
		uintptr(dwBytes))
	return (LPVOID)(unsafe.Pointer(ret1))
}

func HeapCompact(hHeap HANDLE, dwFlags uint32) SIZE_T {
	ret1 := syscall3(heapCompact, 2,
		uintptr(hHeap),
		uintptr(dwFlags),
		0)
	return SIZE_T(ret1)
}

func HeapCreate(flOptions uint32, dwInitialSize SIZE_T, dwMaximumSize SIZE_T) HANDLE {
	ret1 := syscall3(heapCreate, 3,
		uintptr(flOptions),
		uintptr(dwInitialSize),
		uintptr(dwMaximumSize))
	return HANDLE(ret1)
}

func HeapDestroy(hHeap HANDLE) bool {
	ret1 := syscall3(heapDestroy, 1,
		uintptr(hHeap),
		0,
		0)
	return ret1 != 0
}

func HeapFree(hHeap HANDLE, dwFlags uint32, lpMem LPVOID) bool {
	ret1 := syscall3(heapFree, 3,
		uintptr(hHeap),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpMem)))
	return ret1 != 0
}

func HeapLock(hHeap HANDLE) bool {
	ret1 := syscall3(heapLock, 1,
		uintptr(hHeap),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): HEAP_INFORMATION_CLASS, PSIZE_T
// func HeapQueryInformation(heapHandle HANDLE, heapInformationClass HEAP_INFORMATION_CLASS, heapInformation uintptr, heapInformationLength SIZE_T, returnLength PSIZE_T) bool

func HeapReAlloc(hHeap HANDLE, dwFlags uint32, lpMem LPVOID, dwBytes SIZE_T) LPVOID {
	ret1 := syscall6(heapReAlloc, 4,
		uintptr(hHeap),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpMem)),
		uintptr(dwBytes),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): HEAP_INFORMATION_CLASS
// func HeapSetInformation(heapHandle HANDLE, heapInformationClass HEAP_INFORMATION_CLASS, heapInformation uintptr, heapInformationLength SIZE_T) bool

func HeapSize(hHeap HANDLE, dwFlags uint32, lpMem /*const*/ uintptr) SIZE_T {
	ret1 := syscall3(heapSize, 3,
		uintptr(hHeap),
		uintptr(dwFlags),
		lpMem)
	return SIZE_T(ret1)
}

// TODO: Unknown type(s): LPHEAP_SUMMARY
// func HeapSummary(hHeap HANDLE, dwFlags uint32, lpSummary LPHEAP_SUMMARY) bool

func HeapUnlock(hHeap HANDLE) bool {
	ret1 := syscall3(heapUnlock, 1,
		uintptr(hHeap),
		0,
		0)
	return ret1 != 0
}

func HeapValidate(hHeap HANDLE, dwFlags uint32, lpMem /*const*/ uintptr) bool {
	ret1 := syscall3(heapValidate, 3,
		uintptr(hHeap),
		uintptr(dwFlags),
		lpMem)
	return ret1 != 0
}

// TODO: Unknown type(s): LPPROCESS_HEAP_ENTRY
// func HeapWalk(hHeap HANDLE, lpEntry LPPROCESS_HEAP_ENTRY) bool

func IdnToAscii(dwFlags uint32, lpUnicodeCharStr string, cchUnicodeChar int32, lpASCIICharStr LPWSTR, cchASCIIChar int32) int32 {
	lpUnicodeCharStrStr := unicode16FromString(lpUnicodeCharStr)
	ret1 := syscall6(idnToAscii, 5,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpUnicodeCharStrStr[0])),
		uintptr(cchUnicodeChar),
		uintptr(unsafe.Pointer(lpASCIICharStr)),
		uintptr(cchASCIIChar),
		0)
	return int32(ret1)
}

func IdnToNameprepUnicode(dwFlags uint32, lpUnicodeCharStr string, cchUnicodeChar int32, lpNameprepCharStr LPWSTR, cchNameprepChar int32) int32 {
	lpUnicodeCharStrStr := unicode16FromString(lpUnicodeCharStr)
	ret1 := syscall6(idnToNameprepUnicode, 5,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpUnicodeCharStrStr[0])),
		uintptr(cchUnicodeChar),
		uintptr(unsafe.Pointer(lpNameprepCharStr)),
		uintptr(cchNameprepChar),
		0)
	return int32(ret1)
}

func IdnToUnicode(dwFlags uint32, lpASCIICharStr string, cchASCIIChar int32, lpUnicodeCharStr LPWSTR, cchUnicodeChar int32) int32 {
	lpASCIICharStrStr := unicode16FromString(lpASCIICharStr)
	ret1 := syscall6(idnToUnicode, 5,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpASCIICharStrStr[0])),
		uintptr(cchASCIIChar),
		uintptr(unsafe.Pointer(lpUnicodeCharStr)),
		uintptr(cchUnicodeChar),
		0)
	return int32(ret1)
}

func InitAtomTable(nSize uint32) bool {
	ret1 := syscall3(initAtomTable, 1,
		uintptr(nSize),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPINIT_ONCE
// func InitOnceBeginInitialize(lpInitOnce LPINIT_ONCE, dwFlags uint32, fPending *BOOL, lpContext *LPVOID) bool

// TODO: Unknown type(s): LPINIT_ONCE
// func InitOnceComplete(lpInitOnce LPINIT_ONCE, dwFlags uint32, lpContext LPVOID) bool

// TODO: Unknown type(s): PINIT_ONCE, PINIT_ONCE_FN
// func InitOnceExecuteOnce(initOnce PINIT_ONCE, initFn PINIT_ONCE_FN, parameter uintptr, context *LPVOID) bool

// TODO: Unknown type(s): PINIT_ONCE
// func InitOnceInitialize(initOnce PINIT_ONCE)

// TODO: Unknown type(s): PCONDITION_VARIABLE
// func InitializeConditionVariable(conditionVariable PCONDITION_VARIABLE)

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func InitializeCriticalSection(lpCriticalSection LPCRITICAL_SECTION)

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func InitializeCriticalSectionAndSpinCount(lpCriticalSection LPCRITICAL_SECTION, dwSpinCount uint32) bool

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func InitializeCriticalSectionEx(lpCriticalSection LPCRITICAL_SECTION, dwSpinCount uint32, flags uint32) bool

// TODO: Unknown type(s): LPPROC_THREAD_ATTRIBUTE_LIST, PSIZE_T
// func InitializeProcThreadAttributeList(lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST, dwAttributeCount uint32, dwFlags uint32, lpSize PSIZE_T) bool

// TODO: Unknown type(s): PSLIST_HEADER
// func InitializeSListHead(listHead PSLIST_HEADER)

func InitializeSRWLock(sRWLock PSRWLOCK) {
	syscall3(initializeSRWLock, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
}

// TODO: Unknown type(s): PSLIST_ENTRY, PSLIST_HEADER
// func InterlockedFlushSList(listHead PSLIST_HEADER) PSLIST_ENTRY

// TODO: Unknown type(s): PSLIST_ENTRY, PSLIST_HEADER
// func InterlockedPopEntrySList(listHead PSLIST_HEADER) PSLIST_ENTRY

// TODO: Unknown type(s): PSLIST_ENTRY, PSLIST_HEADER
// func InterlockedPushEntrySList(listHead PSLIST_HEADER, listEntry PSLIST_ENTRY) PSLIST_ENTRY

func IsBadCodePtr(lpfn FARPROC) bool {
	lpfnCallback := syscall.NewCallback(func() uintptr {
		ret := lpfn()
		return uintptr(unsafe.Pointer(ret))
	})
	ret1 := syscall3(isBadCodePtr, 1,
		lpfnCallback,
		0,
		0)
	return ret1 != 0
}

func IsBadHugeReadPtr(lp /*const*/ uintptr, ucb *uint32) bool {
	ret1 := syscall3(isBadHugeReadPtr, 2,
		lp,
		uintptr(unsafe.Pointer(ucb)),
		0)
	return ret1 != 0
}

func IsBadHugeWritePtr(lp LPVOID, ucb *uint32) bool {
	ret1 := syscall3(isBadHugeWritePtr, 2,
		uintptr(unsafe.Pointer(lp)),
		uintptr(unsafe.Pointer(ucb)),
		0)
	return ret1 != 0
}

func IsBadReadPtr(lp /*const*/ uintptr, ucb *uint32) bool {
	ret1 := syscall3(isBadReadPtr, 2,
		lp,
		uintptr(unsafe.Pointer(ucb)),
		0)
	return ret1 != 0
}

func IsBadStringPtr(lpsz string, ucchMax *uint32) bool {
	lpszStr := unicode16FromString(lpsz)
	ret1 := syscall3(isBadStringPtr, 2,
		uintptr(unsafe.Pointer(&lpszStr[0])),
		uintptr(unsafe.Pointer(ucchMax)),
		0)
	return ret1 != 0
}

func IsBadWritePtr(lp LPVOID, ucb *uint32) bool {
	ret1 := syscall3(isBadWritePtr, 2,
		uintptr(unsafe.Pointer(lp)),
		uintptr(unsafe.Pointer(ucb)),
		0)
	return ret1 != 0
}

func IsDBCSLeadByte(testChar byte) bool {
	ret1 := syscall3(isDBCSLeadByte, 1,
		uintptr(testChar),
		0,
		0)
	return ret1 != 0
}

func IsDBCSLeadByteEx(codePage uint32, testChar byte) bool {
	ret1 := syscall3(isDBCSLeadByteEx, 2,
		uintptr(codePage),
		uintptr(testChar),
		0)
	return ret1 != 0
}

func IsDebuggerPresent() bool {
	ret1 := syscall3(isDebuggerPresent, 0,
		0,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPNLSVERSIONINFO, NLS_FUNCTION
// func IsNLSDefinedString(function NLS_FUNCTION, dwFlags uint32, lpVersionInformation LPNLSVERSIONINFO, lpString string, cchStr int32) bool

// TODO: Unknown type(s): NORM_FORM
// func IsNormalizedString(normForm NORM_FORM, lpString string, cwLength int32) bool

func IsProcessInJob(processHandle HANDLE, jobHandle HANDLE, result *BOOL) bool {
	ret1 := syscall3(isProcessInJob, 3,
		uintptr(processHandle),
		uintptr(jobHandle),
		uintptr(unsafe.Pointer(result)))
	return ret1 != 0
}

func IsProcessorFeaturePresent(processorFeature uint32) bool {
	ret1 := syscall3(isProcessorFeaturePresent, 1,
		uintptr(processorFeature),
		0,
		0)
	return ret1 != 0
}

func IsSystemResumeAutomatic() bool {
	ret1 := syscall3(isSystemResumeAutomatic, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func IsThreadAFiber() bool {
	ret1 := syscall3(isThreadAFiber, 0,
		0,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PTP_TIMER
// func IsThreadpoolTimerSet(pti PTP_TIMER) bool

func IsValidCodePage(codePage uint32) bool {
	ret1 := syscall3(isValidCodePage, 1,
		uintptr(codePage),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LGRPID
// func IsValidLanguageGroup(languageGroup LGRPID, dwFlags uint32) bool

func IsValidLocale(locale LCID, dwFlags uint32) bool {
	ret1 := syscall3(isValidLocale, 2,
		uintptr(locale),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func IsValidLocaleName(lpLocaleName string) bool {
	lpLocaleNameStr := unicode16FromString(lpLocaleName)
	ret1 := syscall3(isValidLocaleName, 1,
		uintptr(unsafe.Pointer(&lpLocaleNameStr[0])),
		0,
		0)
	return ret1 != 0
}

func IsWow64Process(hProcess HANDLE, wow64Process *BOOL) bool {
	ret1 := syscall3(isWow64Process, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(wow64Process)),
		0)
	return ret1 != 0
}

func LCIDToLocaleName(locale LCID, lpName LPWSTR, cchName int32, dwFlags uint32) int32 {
	ret1 := syscall6(lCIDToLocaleName, 4,
		uintptr(locale),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(cchName),
		uintptr(dwFlags),
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): LPNLSVERSIONINFO
// func LCMapStringEx(lpLocaleName string, dwMapFlags uint32, lpSrcStr string, cchSrc int32, lpDestStr LPWSTR, cchDest int32, lpVersionInformation LPNLSVERSIONINFO, lpReserved LPVOID, sortHandle LPARAM) int32

func LCMapString(locale LCID, dwMapFlags uint32, lpSrcStr string, cchSrc int32, lpDestStr LPWSTR, cchDest int32) int32 {
	lpSrcStrStr := unicode16FromString(lpSrcStr)
	ret1 := syscall6(lCMapString, 6,
		uintptr(locale),
		uintptr(dwMapFlags),
		uintptr(unsafe.Pointer(&lpSrcStrStr[0])),
		uintptr(cchSrc),
		uintptr(unsafe.Pointer(lpDestStr)),
		uintptr(cchDest))
	return int32(ret1)
}

func LZClose(unnamed0 int32) {
	syscall3(lZClose, 1,
		uintptr(unnamed0),
		0,
		0)
}

func LZCopy(unnamed0 int32, unnamed1 int32) LONG {
	ret1 := syscall3(lZCopy, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return LONG(ret1)
}

func LZDone() {
	syscall3(lZDone, 0,
		0,
		0,
		0)
}

func LZInit(unnamed0 int32) int32 {
	ret1 := syscall3(lZInit, 1,
		uintptr(unnamed0),
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): LPOFSTRUCT
// func LZOpenFile(unnamed0 LPWSTR, unnamed1 LPOFSTRUCT, unnamed2 uint16) int32

func LZRead(unnamed0 int32, unnamed1 LPSTR, unnamed2 int32) int32 {
	ret1 := syscall3(lZRead, 3,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(unnamed2))
	return int32(ret1)
}

func LZSeek(unnamed0 int32, unnamed1 LONG, unnamed2 int32) LONG {
	ret1 := syscall3(lZSeek, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2))
	return LONG(ret1)
}

func LZStart() int32 {
	ret1 := syscall3(lZStart, 0,
		0,
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func LeaveCriticalSection(lpCriticalSection LPCRITICAL_SECTION)

// TODO: Unknown type(s): PCRITICAL_SECTION, PTP_CALLBACK_INSTANCE
// func LeaveCriticalSectionWhenCallbackReturns(pci PTP_CALLBACK_INSTANCE, pcs PCRITICAL_SECTION)

func LoadLibraryEx(lpLibFileName string, hFile HANDLE, dwFlags uint32) HMODULE {
	lpLibFileNameStr := unicode16FromString(lpLibFileName)
	ret1 := syscall3(loadLibraryEx, 3,
		uintptr(unsafe.Pointer(&lpLibFileNameStr[0])),
		uintptr(hFile),
		uintptr(dwFlags))
	return HMODULE(ret1)
}

func LoadLibrary(lpLibFileName string) HMODULE {
	lpLibFileNameStr := unicode16FromString(lpLibFileName)
	ret1 := syscall3(loadLibrary, 1,
		uintptr(unsafe.Pointer(&lpLibFileNameStr[0])),
		0,
		0)
	return HMODULE(ret1)
}

func LoadModule(lpModuleName /*const*/ LPCSTR, lpParameterBlock LPVOID) uint32 {
	ret1 := syscall3(loadModule, 2,
		uintptr(unsafe.Pointer(lpModuleName)),
		uintptr(unsafe.Pointer(lpParameterBlock)),
		0)
	return uint32(ret1)
}

func LoadPackagedLibrary(lpwLibFileName string, reserved uint32) HMODULE {
	lpwLibFileNameStr := unicode16FromString(lpwLibFileName)
	ret1 := syscall3(loadPackagedLibrary, 2,
		uintptr(unsafe.Pointer(&lpwLibFileNameStr[0])),
		uintptr(reserved),
		0)
	return HMODULE(ret1)
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
	ret1 := syscall3(loadResource, 2,
		uintptr(hModule),
		uintptr(hResInfo),
		0)
	return HGLOBAL(ret1)
}

func LocalAlloc(uFlags uint32, uBytes SIZE_T) HLOCAL {
	ret1 := syscall3(localAlloc, 2,
		uintptr(uFlags),
		uintptr(uBytes),
		0)
	return HLOCAL(ret1)
}

func LocalCompact(uMinFree uint32) SIZE_T {
	ret1 := syscall3(localCompact, 1,
		uintptr(uMinFree),
		0,
		0)
	return SIZE_T(ret1)
}

func LocalFileTimeToFileTime(lpLocalFileTime /*const*/ *FILETIME, lpFileTime *FILETIME) bool {
	ret1 := syscall3(localFileTimeToFileTime, 2,
		uintptr(unsafe.Pointer(lpLocalFileTime)),
		uintptr(unsafe.Pointer(lpFileTime)),
		0)
	return ret1 != 0
}

func LocalFlags(hMem HLOCAL) uint32 {
	ret1 := syscall3(localFlags, 1,
		uintptr(hMem),
		0,
		0)
	return uint32(ret1)
}

func LocalFree(hMem HLOCAL) HLOCAL {
	ret1 := syscall3(localFree, 1,
		uintptr(hMem),
		0,
		0)
	return HLOCAL(ret1)
}

func LocalHandle(pMem /*const*/ uintptr) HLOCAL {
	ret1 := syscall3(localHandle, 1,
		pMem,
		0,
		0)
	return HLOCAL(ret1)
}

func LocalLock(hMem HLOCAL) LPVOID {
	ret1 := syscall3(localLock, 1,
		uintptr(hMem),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func LocalReAlloc(hMem HLOCAL, uBytes SIZE_T, uFlags uint32) HLOCAL {
	ret1 := syscall3(localReAlloc, 3,
		uintptr(hMem),
		uintptr(uBytes),
		uintptr(uFlags))
	return HLOCAL(ret1)
}

func LocalShrink(hMem HLOCAL, cbNewSize uint32) SIZE_T {
	ret1 := syscall3(localShrink, 2,
		uintptr(hMem),
		uintptr(cbNewSize),
		0)
	return SIZE_T(ret1)
}

func LocalSize(hMem HLOCAL) SIZE_T {
	ret1 := syscall3(localSize, 1,
		uintptr(hMem),
		0,
		0)
	return SIZE_T(ret1)
}

func LocalUnlock(hMem HLOCAL) bool {
	ret1 := syscall3(localUnlock, 1,
		uintptr(hMem),
		0,
		0)
	return ret1 != 0
}

func LocaleNameToLCID(lpName string, dwFlags uint32) LCID {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(localeNameToLCID, 2,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(dwFlags),
		0)
	return LCID(ret1)
}

func LockFile(hFile HANDLE, dwFileOffsetLow uint32, dwFileOffsetHigh uint32, nNumberOfBytesToLockLow uint32, nNumberOfBytesToLockHigh uint32) bool {
	ret1 := syscall6(lockFile, 5,
		uintptr(hFile),
		uintptr(dwFileOffsetLow),
		uintptr(dwFileOffsetHigh),
		uintptr(nNumberOfBytesToLockLow),
		uintptr(nNumberOfBytesToLockHigh),
		0)
	return ret1 != 0
}

func LockFileEx(hFile HANDLE, dwFlags uint32, dwReserved uint32, nNumberOfBytesToLockLow uint32, nNumberOfBytesToLockHigh uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall6(lockFileEx, 6,
		uintptr(hFile),
		uintptr(dwFlags),
		uintptr(dwReserved),
		uintptr(nNumberOfBytesToLockLow),
		uintptr(nNumberOfBytesToLockHigh),
		uintptr(unsafe.Pointer(lpOverlapped)))
	return ret1 != 0
}

func LockResource(hResData HGLOBAL) LPVOID {
	ret1 := syscall3(lockResource, 1,
		uintptr(hResData),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PULONG_PTR
// func MapUserPhysicalPages(virtualAddress uintptr, numberOfPages *uint32, pageArray PULONG_PTR) bool

// TODO: Unknown type(s): PULONG_PTR
// func MapUserPhysicalPagesScatter(virtualAddresses *PVOID, numberOfPages *uint32, pageArray PULONG_PTR) bool

func MapViewOfFile(hFileMappingObject HANDLE, dwDesiredAccess uint32, dwFileOffsetHigh uint32, dwFileOffsetLow uint32, dwNumberOfBytesToMap SIZE_T) LPVOID {
	ret1 := syscall6(mapViewOfFile, 5,
		uintptr(hFileMappingObject),
		uintptr(dwDesiredAccess),
		uintptr(dwFileOffsetHigh),
		uintptr(dwFileOffsetLow),
		uintptr(dwNumberOfBytesToMap),
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func MapViewOfFileEx(hFileMappingObject HANDLE, dwDesiredAccess uint32, dwFileOffsetHigh uint32, dwFileOffsetLow uint32, dwNumberOfBytesToMap SIZE_T, lpBaseAddress LPVOID) LPVOID {
	ret1 := syscall6(mapViewOfFileEx, 6,
		uintptr(hFileMappingObject),
		uintptr(dwDesiredAccess),
		uintptr(dwFileOffsetHigh),
		uintptr(dwFileOffsetLow),
		uintptr(dwNumberOfBytesToMap),
		uintptr(unsafe.Pointer(lpBaseAddress)))
	return (LPVOID)(unsafe.Pointer(ret1))
}

func MapViewOfFileExNuma(hFileMappingObject HANDLE, dwDesiredAccess uint32, dwFileOffsetHigh uint32, dwFileOffsetLow uint32, dwNumberOfBytesToMap SIZE_T, lpBaseAddress LPVOID, nndPreferred uint32) LPVOID {
	ret1 := syscall9(mapViewOfFileExNuma, 7,
		uintptr(hFileMappingObject),
		uintptr(dwDesiredAccess),
		uintptr(dwFileOffsetHigh),
		uintptr(dwFileOffsetLow),
		uintptr(dwNumberOfBytesToMap),
		uintptr(unsafe.Pointer(lpBaseAddress)),
		uintptr(nndPreferred),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): LPMODULEENTRY32W
// func Module32FirstW(hSnapshot HANDLE, lpme LPMODULEENTRY32W) bool

// TODO: Unknown type(s): LPMODULEENTRY32W
// func Module32NextW(hSnapshot HANDLE, lpme LPMODULEENTRY32W) bool

func MoveFileEx(lpExistingFileName string, lpNewFileName string, dwFlags uint32) bool {
	lpExistingFileNameStr := unicode16FromString(lpExistingFileName)
	lpNewFileNameStr := unicode16FromString(lpNewFileName)
	ret1 := syscall3(moveFileEx, 3,
		uintptr(unsafe.Pointer(&lpExistingFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpNewFileNameStr[0])),
		uintptr(dwFlags))
	return ret1 != 0
}

// TODO: Unknown type(s): LPPROGRESS_ROUTINE
// func MoveFileTransacted(lpExistingFileName string, lpNewFileName string, lpProgressRoutine LPPROGRESS_ROUTINE, lpData LPVOID, dwFlags uint32, hTransaction HANDLE) bool

func MoveFile(lpExistingFileName string, lpNewFileName string) bool {
	lpExistingFileNameStr := unicode16FromString(lpExistingFileName)
	lpNewFileNameStr := unicode16FromString(lpNewFileName)
	ret1 := syscall3(moveFile, 2,
		uintptr(unsafe.Pointer(&lpExistingFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpNewFileNameStr[0])),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPPROGRESS_ROUTINE
// func MoveFileWithProgress(lpExistingFileName string, lpNewFileName string, lpProgressRoutine LPPROGRESS_ROUTINE, lpData LPVOID, dwFlags uint32) bool

func MulDiv(nNumber int32, nNumerator int32, nDenominator int32) int32 {
	ret1 := syscall3(mulDiv, 3,
		uintptr(nNumber),
		uintptr(nNumerator),
		uintptr(nDenominator))
	return int32(ret1)
}

// TODO: Unknown type(s): LPCCH
// func MultiByteToWideChar(codePage uint32, dwFlags uint32, lpMultiByteStr LPCCH, cbMultiByte int32, lpWideCharStr LPWSTR, cchWideChar int32) int32

func NeedCurrentDirectoryForExePath(exeName string) bool {
	exeNameStr := unicode16FromString(exeName)
	ret1 := syscall3(needCurrentDirectoryForExePath, 1,
		uintptr(unsafe.Pointer(&exeNameStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): NORM_FORM
// func NormalizeString(normForm NORM_FORM, lpSrcString string, cwSrcLength int32, lpDstString LPWSTR, cwDstLength int32) int32

func NotifyUILanguageChange(dwFlags uint32, pcwstrNewLanguage string, pcwstrPreviousLanguage string, dwReserved uint32, pdwStatusRtrn *DWORD) bool {
	pcwstrNewLanguageStr := unicode16FromString(pcwstrNewLanguage)
	pcwstrPreviousLanguageStr := unicode16FromString(pcwstrPreviousLanguage)
	ret1 := syscall6(notifyUILanguageChange, 5,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&pcwstrNewLanguageStr[0])),
		uintptr(unsafe.Pointer(&pcwstrPreviousLanguageStr[0])),
		uintptr(dwReserved),
		uintptr(unsafe.Pointer(pdwStatusRtrn)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): OFFER_PRIORITY
// func OfferVirtualMemory(virtualAddress uintptr, size SIZE_T, priority OFFER_PRIORITY) uint32

func OpenEvent(dwDesiredAccess uint32, bInheritHandle bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(openEvent, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

// TODO: Unknown type(s): LPOFSTRUCT
// func OpenFile(lpFileName /*const*/ LPCSTR, lpReOpenBuff LPOFSTRUCT, uStyle uint32) HFILE

// TODO: Unknown type(s): LPFILE_ID_DESCRIPTOR
// func OpenFileById(hVolumeHint HANDLE, lpFileId LPFILE_ID_DESCRIPTOR, dwDesiredAccess uint32, dwShareMode uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwFlagsAndAttributes uint32) HANDLE

func OpenFileMapping(dwDesiredAccess uint32, bInheritHandle bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(openFileMapping, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

func OpenJobObject(dwDesiredAccess uint32, bInheritHandle bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(openJobObject, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

func OpenMutex(dwDesiredAccess uint32, bInheritHandle bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(openMutex, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

func OpenPrivateNamespace(lpBoundaryDescriptor LPVOID, lpAliasPrefix string) HANDLE {
	lpAliasPrefixStr := unicode16FromString(lpAliasPrefix)
	ret1 := syscall3(openPrivateNamespace, 2,
		uintptr(unsafe.Pointer(lpBoundaryDescriptor)),
		uintptr(unsafe.Pointer(&lpAliasPrefixStr[0])),
		0)
	return HANDLE(ret1)
}

func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, dwProcessId uint32) HANDLE {
	ret1 := syscall3(openProcess, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(dwProcessId))
	return HANDLE(ret1)
}

func OpenSemaphore(dwDesiredAccess uint32, bInheritHandle bool, lpName string) HANDLE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(openSemaphore, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(unsafe.Pointer(&lpNameStr[0])))
	return HANDLE(ret1)
}

func OpenThread(dwDesiredAccess uint32, bInheritHandle bool, dwThreadId uint32) HANDLE {
	ret1 := syscall3(openThread, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(dwThreadId))
	return HANDLE(ret1)
}

func OpenWaitableTimer(dwDesiredAccess uint32, bInheritHandle bool, lpTimerName string) HANDLE {
	lpTimerNameStr := unicode16FromString(lpTimerName)
	ret1 := syscall3(openWaitableTimer, 3,
		uintptr(dwDesiredAccess),
		getUintptrFromBool(bInheritHandle),
		uintptr(unsafe.Pointer(&lpTimerNameStr[0])))
	return HANDLE(ret1)
}

func OutputDebugString(lpOutputString string) {
	lpOutputStringStr := unicode16FromString(lpOutputString)
	syscall3(outputDebugString, 1,
		uintptr(unsafe.Pointer(&lpOutputStringStr[0])),
		0,
		0)
}

// TODO: Unknown type(s): PINPUT_RECORD
// func PeekConsoleInput(hConsoleInput HANDLE, lpBuffer PINPUT_RECORD, nLength uint32, lpNumberOfEventsRead *uint32) bool

func PeekNamedPipe(hNamedPipe HANDLE, lpBuffer LPVOID, nBufferSize uint32, lpBytesRead *uint32, lpTotalBytesAvail *uint32, lpBytesLeftThisMessage *uint32) bool {
	ret1 := syscall6(peekNamedPipe, 6,
		uintptr(hNamedPipe),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nBufferSize),
		uintptr(unsafe.Pointer(lpBytesRead)),
		uintptr(unsafe.Pointer(lpTotalBytesAvail)),
		uintptr(unsafe.Pointer(lpBytesLeftThisMessage)))
	return ret1 != 0
}

func PostQueuedCompletionStatus(completionPort HANDLE, dwNumberOfBytesTransferred uint32, dwCompletionKey *uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall6(postQueuedCompletionStatus, 4,
		uintptr(completionPort),
		uintptr(dwNumberOfBytesTransferred),
		uintptr(unsafe.Pointer(dwCompletionKey)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): POWER_REQUEST_TYPE
// func PowerClearRequest(powerRequest HANDLE, requestType POWER_REQUEST_TYPE) bool

// TODO: Unknown type(s): PREASON_CONTEXT
// func PowerCreateRequest(context PREASON_CONTEXT) HANDLE

// TODO: Unknown type(s): POWER_REQUEST_TYPE
// func PowerSetRequest(powerRequest HANDLE, requestType POWER_REQUEST_TYPE) bool

// TODO: Unknown type(s): PWIN32_MEMORY_RANGE_ENTRY
// func PrefetchVirtualMemory(hProcess HANDLE, numberOfEntries *uint32, virtualAddresses PWIN32_MEMORY_RANGE_ENTRY, flags ULONG) bool

func PrepareTape(hDevice HANDLE, dwOperation uint32, bImmediate bool) uint32 {
	ret1 := syscall3(prepareTape, 3,
		uintptr(hDevice),
		uintptr(dwOperation),
		getUintptrFromBool(bImmediate))
	return uint32(ret1)
}

// TODO: Unknown type(s): LPPROCESSENTRY32W
// func Process32FirstW(hSnapshot HANDLE, lppe LPPROCESSENTRY32W) bool

// TODO: Unknown type(s): LPPROCESSENTRY32W
// func Process32NextW(hSnapshot HANDLE, lppe LPPROCESSENTRY32W) bool

func ProcessIdToSessionId(dwProcessId uint32, pSessionId *uint32) bool {
	ret1 := syscall3(processIdToSessionId, 2,
		uintptr(dwProcessId),
		uintptr(unsafe.Pointer(pSessionId)),
		0)
	return ret1 != 0
}

func PulseEvent(hEvent HANDLE) bool {
	ret1 := syscall3(pulseEvent, 1,
		uintptr(hEvent),
		0,
		0)
	return ret1 != 0
}

func PurgeComm(hFile HANDLE, dwFlags uint32) bool {
	ret1 := syscall3(purgeComm, 2,
		uintptr(hFile),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func QueryActCtxSettingsW(dwFlags uint32, hActCtx HANDLE, settingsNameSpace string, settingName string, pvBuffer PWSTR, dwBuffer SIZE_T, pdwWrittenOrRequired *SIZE_T) bool {
	settingsNameSpaceStr := unicode16FromString(settingsNameSpace)
	settingNameStr := unicode16FromString(settingName)
	ret1 := syscall9(queryActCtxSettingsW, 7,
		uintptr(dwFlags),
		uintptr(hActCtx),
		uintptr(unsafe.Pointer(&settingsNameSpaceStr[0])),
		uintptr(unsafe.Pointer(&settingNameStr[0])),
		uintptr(unsafe.Pointer(pvBuffer)),
		uintptr(dwBuffer),
		uintptr(unsafe.Pointer(pdwWrittenOrRequired)),
		0,
		0)
	return ret1 != 0
}

func QueryActCtxW(dwFlags uint32, hActCtx HANDLE, pvSubInstance uintptr, ulInfoClass ULONG, pvBuffer uintptr, cbBuffer SIZE_T, pcbWrittenOrRequired *SIZE_T) bool {
	ret1 := syscall9(queryActCtxW, 7,
		uintptr(dwFlags),
		uintptr(hActCtx),
		pvSubInstance,
		uintptr(ulInfoClass),
		pvBuffer,
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(pcbWrittenOrRequired)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PSLIST_HEADER
// func QueryDepthSList(listHead PSLIST_HEADER) USHORT

func QueryDosDevice(lpDeviceName string, lpTargetPath LPWSTR, ucchMax uint32) uint32 {
	lpDeviceNameStr := unicode16FromString(lpDeviceName)
	ret1 := syscall3(queryDosDevice, 3,
		uintptr(unsafe.Pointer(&lpDeviceNameStr[0])),
		uintptr(unsafe.Pointer(lpTargetPath)),
		uintptr(ucchMax))
	return uint32(ret1)
}

func QueryFullProcessImageName(hProcess HANDLE, dwFlags uint32, lpExeName LPWSTR, lpdwSize *DWORD) bool {
	ret1 := syscall6(queryFullProcessImageName, 4,
		uintptr(hProcess),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpExeName)),
		uintptr(unsafe.Pointer(lpdwSize)),
		0,
		0)
	return ret1 != 0
}

func QueryIdleProcessorCycleTime(bufferLength *uint32, processorIdleCycleTime PULONG64) bool {
	ret1 := syscall3(queryIdleProcessorCycleTime, 2,
		uintptr(unsafe.Pointer(bufferLength)),
		uintptr(unsafe.Pointer(processorIdleCycleTime)),
		0)
	return ret1 != 0
}

func QueryIdleProcessorCycleTimeEx(group USHORT, bufferLength *uint32, processorIdleCycleTime PULONG64) bool {
	ret1 := syscall3(queryIdleProcessorCycleTimeEx, 3,
		uintptr(group),
		uintptr(unsafe.Pointer(bufferLength)),
		uintptr(unsafe.Pointer(processorIdleCycleTime)))
	return ret1 != 0
}

// TODO: Unknown type(s): JOBOBJECTINFOCLASS
// func QueryInformationJobObject(hJob HANDLE, jobObjectInformationClass JOBOBJECTINFOCLASS, lpJobObjectInformation LPVOID, cbJobObjectInformationLength uint32, lpReturnLength *uint32) bool

func QueryMemoryResourceNotification(resourceNotificationHandle HANDLE, resourceState *BOOL) bool {
	ret1 := syscall3(queryMemoryResourceNotification, 2,
		uintptr(resourceNotificationHandle),
		uintptr(unsafe.Pointer(resourceState)),
		0)
	return ret1 != 0
}

func QueryPerformanceCounter(lpPerformanceCount *LARGE_INTEGER) bool {
	ret1 := syscall3(queryPerformanceCounter, 1,
		uintptr(unsafe.Pointer(lpPerformanceCount)),
		0,
		0)
	return ret1 != 0
}

func QueryPerformanceFrequency(lpFrequency *LARGE_INTEGER) bool {
	ret1 := syscall3(queryPerformanceFrequency, 1,
		uintptr(unsafe.Pointer(lpFrequency)),
		0,
		0)
	return ret1 != 0
}

func QueryProcessAffinityUpdateMode(hProcess HANDLE, lpdwFlags *uint32) bool {
	ret1 := syscall3(queryProcessAffinityUpdateMode, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpdwFlags)),
		0)
	return ret1 != 0
}

func QueryProcessCycleTime(processHandle HANDLE, cycleTime PULONG64) bool {
	ret1 := syscall3(queryProcessCycleTime, 2,
		uintptr(processHandle),
		uintptr(unsafe.Pointer(cycleTime)),
		0)
	return ret1 != 0
}

func QueryThreadCycleTime(threadHandle HANDLE, cycleTime PULONG64) bool {
	ret1 := syscall3(queryThreadCycleTime, 2,
		uintptr(threadHandle),
		uintptr(unsafe.Pointer(cycleTime)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PBOOLEAN
// func QueryThreadProfiling(threadHandle HANDLE, enabled PBOOLEAN) uint32

// TODO: Unknown type(s): PTP_POOL, PTP_POOL_STACK_INFORMATION
// func QueryThreadpoolStackInformation(ptpp PTP_POOL, ptpsi PTP_POOL_STACK_INFORMATION) bool

// TODO: Unknown type(s): PULONGLONG
// func QueryUnbiasedInterruptTime(unbiasedTime PULONGLONG) bool

// TODO: Unknown type(s): PAPCFUNC
// func QueueUserAPC(pfnAPC PAPCFUNC, hThread HANDLE, dwData *uint32) uint32

func QueueUserWorkItem(function THREAD_START_ROUTINE, context uintptr, flags ULONG) bool {
	functionCallback := syscall.NewCallback(func(lpThreadParameterRawArg LPVOID) uintptr {
		ret := function(lpThreadParameterRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(queueUserWorkItem, 3,
		functionCallback,
		context,
		uintptr(flags))
	return ret1 != 0
}

func RaiseException(dwExceptionCode uint32, dwExceptionFlags uint32, nNumberOfArguments uint32, lpArguments /*const*/ *ULONG_PTR) {
	syscall6(raiseException, 4,
		uintptr(dwExceptionCode),
		uintptr(dwExceptionFlags),
		uintptr(nNumberOfArguments),
		uintptr(unsafe.Pointer(lpArguments)),
		0,
		0)
}

// TODO: Unknown type(s): PCONTEXT, PEXCEPTION_RECORD
// func RaiseFailFastException(pExceptionRecord PEXCEPTION_RECORD, pContextRecord PCONTEXT, dwFlags uint32)

func ReOpenFile(hOriginalFile HANDLE, dwDesiredAccess uint32, dwShareMode uint32, dwFlagsAndAttributes uint32) HANDLE {
	ret1 := syscall6(reOpenFile, 4,
		uintptr(hOriginalFile),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		uintptr(dwFlagsAndAttributes),
		0,
		0)
	return HANDLE(ret1)
}

// TODO: Unknown type(s): PINPUT_RECORD
// func ReadConsoleInput(hConsoleInput HANDLE, lpBuffer PINPUT_RECORD, nLength uint32, lpNumberOfEventsRead *uint32) bool

func ReadConsoleOutputAttribute(hConsoleOutput HANDLE, lpAttribute *uint16, nLength uint32, dwReadCoord COORD, lpNumberOfAttrsRead *uint32) bool {
	ret1 := syscall6(readConsoleOutputAttribute, 5,
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(lpAttribute)),
		uintptr(nLength),
		getUintptrFromCOORD(dwReadCoord),
		uintptr(unsafe.Pointer(lpNumberOfAttrsRead)),
		0)
	return ret1 != 0
}

func ReadConsoleOutputCharacter(hConsoleOutput HANDLE, lpCharacter LPWSTR, nLength uint32, dwReadCoord COORD, lpNumberOfCharsRead *uint32) bool {
	ret1 := syscall6(readConsoleOutputCharacter, 5,
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(lpCharacter)),
		uintptr(nLength),
		getUintptrFromCOORD(dwReadCoord),
		uintptr(unsafe.Pointer(lpNumberOfCharsRead)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCHAR_INFO, PSMALL_RECT
// func ReadConsoleOutput(hConsoleOutput HANDLE, lpBuffer PCHAR_INFO, dwBufferSize COORD, dwBufferCoord COORD, lpReadRegion PSMALL_RECT) bool

func ReadConsole(hConsoleInput HANDLE, lpBuffer LPVOID, nNumberOfCharsToRead uint32, lpNumberOfCharsRead *uint32, lpReserved LPVOID) bool {
	ret1 := syscall6(readConsole, 5,
		uintptr(hConsoleInput),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nNumberOfCharsToRead),
		uintptr(unsafe.Pointer(lpNumberOfCharsRead)),
		uintptr(unsafe.Pointer(lpReserved)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPOVERLAPPED_COMPLETION_ROUTINE
// func ReadDirectoryChangesW(hDirectory HANDLE, lpBuffer LPVOID, nBufferLength uint32, bWatchSubtree bool, dwNotifyFilter uint32, lpBytesReturned *uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) bool

func ReadFile(hFile HANDLE, lpBuffer LPVOID, nNumberOfBytesToRead uint32, lpNumberOfBytesRead *uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall6(readFile, 5,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nNumberOfBytesToRead),
		uintptr(unsafe.Pointer(lpNumberOfBytesRead)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPOVERLAPPED_COMPLETION_ROUTINE
// func ReadFileEx(hFile HANDLE, lpBuffer LPVOID, nNumberOfBytesToRead uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) bool

// TODO: Unknown type(s): FILE_SEGMENT_ELEMENT*
// func ReadFileScatter(hFile HANDLE, aSegmentArray FILE_SEGMENT_ELEMENT*, nNumberOfBytesToRead uint32, lpReserved *uint32, lpOverlapped *OVERLAPPED) bool

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress /*const*/ uintptr, lpBuffer LPVOID, nSize SIZE_T, lpNumberOfBytesRead *SIZE_T) bool {
	ret1 := syscall6(readProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpNumberOfBytesRead)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PPERFORMANCE_DATA
// func ReadThreadProfilingData(performanceDataHandle HANDLE, flags uint32, performanceData PPERFORMANCE_DATA) uint32

func ReclaimVirtualMemory(virtualAddress uintptr, size SIZE_T) uint32 {
	ret1 := syscall3(reclaimVirtualMemory, 2,
		virtualAddress,
		uintptr(size),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): APPLICATION_RECOVERY_CALLBACK
// func RegisterApplicationRecoveryCallback(pRecoveyCallback APPLICATION_RECOVERY_CALLBACK, pvParameter uintptr, dwPingInterval uint32, dwFlags uint32) HRESULT

func RegisterApplicationRestart(pwzCommandline string, dwFlags uint32) HRESULT {
	pwzCommandlineStr := unicode16FromString(pwzCommandline)
	ret1 := syscall3(registerApplicationRestart, 2,
		uintptr(unsafe.Pointer(&pwzCommandlineStr[0])),
		uintptr(dwFlags),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): WAITORTIMERCALLBACK
// func RegisterWaitForSingleObject(phNewWaitObject *HANDLE, hObject HANDLE, callback WAITORTIMERCALLBACK, context uintptr, dwMilliseconds ULONG, dwFlags ULONG) bool

func ReleaseActCtx(hActCtx HANDLE) {
	syscall3(releaseActCtx, 1,
		uintptr(hActCtx),
		0,
		0)
}

func ReleaseMutex(hMutex HANDLE) bool {
	ret1 := syscall3(releaseMutex, 1,
		uintptr(hMutex),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PTP_CALLBACK_INSTANCE
// func ReleaseMutexWhenCallbackReturns(pci PTP_CALLBACK_INSTANCE, mut HANDLE)

func ReleaseSRWLockExclusive(sRWLock PSRWLOCK) {
	syscall3(releaseSRWLockExclusive, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
}

func ReleaseSRWLockShared(sRWLock PSRWLOCK) {
	syscall3(releaseSRWLockShared, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
}

func ReleaseSemaphore(hSemaphore HANDLE, lReleaseCount LONG, lpPreviousCount *LONG) bool {
	ret1 := syscall3(releaseSemaphore, 3,
		uintptr(hSemaphore),
		uintptr(lReleaseCount),
		uintptr(unsafe.Pointer(lpPreviousCount)))
	return ret1 != 0
}

// TODO: Unknown type(s): PTP_CALLBACK_INSTANCE
// func ReleaseSemaphoreWhenCallbackReturns(pci PTP_CALLBACK_INSTANCE, sem HANDLE, crel uint32)

func RemoveDirectoryTransacted(lpPathName string, hTransaction HANDLE) bool {
	lpPathNameStr := unicode16FromString(lpPathName)
	ret1 := syscall3(removeDirectoryTransacted, 2,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		uintptr(hTransaction),
		0)
	return ret1 != 0
}

func RemoveDirectory(lpPathName string) bool {
	lpPathNameStr := unicode16FromString(lpPathName)
	ret1 := syscall3(removeDirectory, 1,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		0,
		0)
	return ret1 != 0
}

func RemoveSecureMemoryCacheCallback(pfnCallBack PSECURE_MEMORY_CACHE_CALLBACK) bool {
	pfnCallBackCallback := syscall.NewCallback(func(AddrRawArg PVOID, RangeRawArg SIZE_T) uintptr {
		ret := pfnCallBack(AddrRawArg, RangeRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(removeSecureMemoryCacheCallback, 1,
		pfnCallBackCallback,
		0,
		0)
	return ret1 != 0
}

func RemoveVectoredContinueHandler(handle uintptr) ULONG {
	ret1 := syscall3(removeVectoredContinueHandler, 1,
		handle,
		0,
		0)
	return ULONG(ret1)
}

func RemoveVectoredExceptionHandler(handle uintptr) ULONG {
	ret1 := syscall3(removeVectoredExceptionHandler, 1,
		handle,
		0,
		0)
	return ULONG(ret1)
}

func ReplaceFile(lpReplacedFileName string, lpReplacementFileName string, lpBackupFileName string, dwReplaceFlags uint32, lpExclude LPVOID, lpReserved LPVOID) bool {
	lpReplacedFileNameStr := unicode16FromString(lpReplacedFileName)
	lpReplacementFileNameStr := unicode16FromString(lpReplacementFileName)
	lpBackupFileNameStr := unicode16FromString(lpBackupFileName)
	ret1 := syscall6(replaceFile, 6,
		uintptr(unsafe.Pointer(&lpReplacedFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpReplacementFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpBackupFileNameStr[0])),
		uintptr(dwReplaceFlags),
		uintptr(unsafe.Pointer(lpExclude)),
		uintptr(unsafe.Pointer(lpReserved)))
	return ret1 != 0
}

func ReplacePartitionUnit(targetPartition PWSTR, sparePartition PWSTR, flags ULONG) bool {
	ret1 := syscall3(replacePartitionUnit, 3,
		uintptr(unsafe.Pointer(targetPartition)),
		uintptr(unsafe.Pointer(sparePartition)),
		uintptr(flags))
	return ret1 != 0
}

func RequestDeviceWakeup(hDevice HANDLE) bool {
	ret1 := syscall3(requestDeviceWakeup, 1,
		uintptr(hDevice),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LATENCY_TIME
// func RequestWakeupLatency(latency LATENCY_TIME) bool

func ResetEvent(hEvent HANDLE) bool {
	ret1 := syscall3(resetEvent, 1,
		uintptr(hEvent),
		0,
		0)
	return ret1 != 0
}

func ResetWriteWatch(lpBaseAddress LPVOID, dwRegionSize SIZE_T) uint32 {
	ret1 := syscall3(resetWriteWatch, 2,
		uintptr(unsafe.Pointer(lpBaseAddress)),
		uintptr(dwRegionSize),
		0)
	return uint32(ret1)
}

func ResolveLocaleName(lpNameToResolve string, lpLocaleName LPWSTR, cchLocaleName int32) int32 {
	lpNameToResolveStr := unicode16FromString(lpNameToResolve)
	ret1 := syscall3(resolveLocaleName, 3,
		uintptr(unsafe.Pointer(&lpNameToResolveStr[0])),
		uintptr(unsafe.Pointer(lpLocaleName)),
		uintptr(cchLocaleName))
	return int32(ret1)
}

func RestoreLastError(dwErrCode uint32) {
	syscall3(restoreLastError, 1,
		uintptr(dwErrCode),
		0,
		0)
}

func ResumeThread(hThread HANDLE) uint32 {
	ret1 := syscall3(resumeThread, 1,
		uintptr(hThread),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): CONST CHAR_INFO *, CONST SMALL_RECT *
// func ScrollConsoleScreenBuffer(hConsoleOutput HANDLE, lpScrollRectangle /*const*/ CONST SMALL_RECT *, lpClipRectangle /*const*/ CONST SMALL_RECT *, dwDestinationOrigin COORD, lpFill /*const*/ CONST CHAR_INFO *) bool

func SearchPath(lpPath string, lpFileName string, lpExtension string, nBufferLength uint32, lpBuffer LPWSTR, lpFilePart *LPWSTR) uint32 {
	lpPathStr := unicode16FromString(lpPath)
	lpFileNameStr := unicode16FromString(lpFileName)
	lpExtensionStr := unicode16FromString(lpExtension)
	ret1 := syscall6(searchPath, 6,
		uintptr(unsafe.Pointer(&lpPathStr[0])),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(&lpExtensionStr[0])),
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(unsafe.Pointer(lpFilePart)))
	return uint32(ret1)
}

// TODO: Unknown type(s): CALTYPE
// func SetCalendarInfo(locale LCID, calendar CALID, calType CALTYPE, lpCalData string) bool

func SetCommBreak(hFile HANDLE) bool {
	ret1 := syscall3(setCommBreak, 1,
		uintptr(hFile),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCOMMCONFIG
// func SetCommConfig(hCommDev HANDLE, lpCC LPCOMMCONFIG, dwSize uint32) bool

func SetCommMask(hFile HANDLE, dwEvtMask uint32) bool {
	ret1 := syscall3(setCommMask, 2,
		uintptr(hFile),
		uintptr(dwEvtMask),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPDCB
// func SetCommState(hFile HANDLE, lpDCB LPDCB) bool

// TODO: Unknown type(s): LPCOMMTIMEOUTS
// func SetCommTimeouts(hFile HANDLE, lpCommTimeouts LPCOMMTIMEOUTS) bool

// TODO: Unknown type(s): COMPUTER_NAME_FORMAT
// func SetComputerNameEx(nameType COMPUTER_NAME_FORMAT, lpBuffer string) bool

func SetComputerName(lpComputerName string) bool {
	lpComputerNameStr := unicode16FromString(lpComputerName)
	ret1 := syscall3(setComputerName, 1,
		uintptr(unsafe.Pointer(&lpComputerNameStr[0])),
		0,
		0)
	return ret1 != 0
}

func SetConsoleActiveScreenBuffer(hConsoleOutput HANDLE) bool {
	ret1 := syscall3(setConsoleActiveScreenBuffer, 1,
		uintptr(hConsoleOutput),
		0,
		0)
	return ret1 != 0
}

func SetConsoleCP(wCodePageID uint32) bool {
	ret1 := syscall3(setConsoleCP, 1,
		uintptr(wCodePageID),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PHANDLER_ROUTINE
// func SetConsoleCtrlHandler(handlerRoutine PHANDLER_ROUTINE, add bool) bool

// TODO: Unknown type(s): CONST CONSOLE_CURSOR_INFO *
// func SetConsoleCursorInfo(hConsoleOutput HANDLE, lpConsoleCursorInfo /*const*/ CONST CONSOLE_CURSOR_INFO *) bool

func SetConsoleCursorPosition(hConsoleOutput HANDLE, dwCursorPosition COORD) bool {
	ret1 := syscall3(setConsoleCursorPosition, 2,
		uintptr(hConsoleOutput),
		getUintptrFromCOORD(dwCursorPosition),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCOORD
// func SetConsoleDisplayMode(hConsoleOutput HANDLE, dwFlags uint32, lpNewScreenBufferDimensions PCOORD) bool

// TODO: Unknown type(s): PCONSOLE_HISTORY_INFO
// func SetConsoleHistoryInfo(lpConsoleHistoryInfo PCONSOLE_HISTORY_INFO) bool

func SetConsoleMode(hConsoleHandle HANDLE, dwMode uint32) bool {
	ret1 := syscall3(setConsoleMode, 2,
		uintptr(hConsoleHandle),
		uintptr(dwMode),
		0)
	return ret1 != 0
}

func SetConsoleOutputCP(wCodePageID uint32) bool {
	ret1 := syscall3(setConsoleOutputCP, 1,
		uintptr(wCodePageID),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCONSOLE_SCREEN_BUFFER_INFOEX
// func SetConsoleScreenBufferInfoEx(hConsoleOutput HANDLE, lpConsoleScreenBufferInfoEx PCONSOLE_SCREEN_BUFFER_INFOEX) bool

func SetConsoleScreenBufferSize(hConsoleOutput HANDLE, dwSize COORD) bool {
	ret1 := syscall3(setConsoleScreenBufferSize, 2,
		uintptr(hConsoleOutput),
		getUintptrFromCOORD(dwSize),
		0)
	return ret1 != 0
}

func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes uint16) bool {
	ret1 := syscall3(setConsoleTextAttribute, 2,
		uintptr(hConsoleOutput),
		uintptr(wAttributes),
		0)
	return ret1 != 0
}

func SetConsoleTitle(lpConsoleTitle string) bool {
	lpConsoleTitleStr := unicode16FromString(lpConsoleTitle)
	ret1 := syscall3(setConsoleTitle, 1,
		uintptr(unsafe.Pointer(&lpConsoleTitleStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CONST SMALL_RECT *
// func SetConsoleWindowInfo(hConsoleOutput HANDLE, bAbsolute bool, lpConsoleWindow /*const*/ CONST SMALL_RECT *) bool

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func SetCriticalSectionSpinCount(lpCriticalSection LPCRITICAL_SECTION, dwSpinCount uint32) uint32

// TODO: Unknown type(s): PCONSOLE_FONT_INFOEX
// func SetCurrentConsoleFontEx(hConsoleOutput HANDLE, bMaximumWindow bool, lpConsoleCurrentFontEx PCONSOLE_FONT_INFOEX) bool

func SetCurrentDirectory(lpPathName string) bool {
	lpPathNameStr := unicode16FromString(lpPathName)
	ret1 := syscall3(setCurrentDirectory, 1,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPCOMMCONFIG
// func SetDefaultCommConfig(lpszName string, lpCC LPCOMMCONFIG, dwSize uint32) bool

func SetDllDirectory(lpPathName string) bool {
	lpPathNameStr := unicode16FromString(lpPathName)
	ret1 := syscall3(setDllDirectory, 1,
		uintptr(unsafe.Pointer(&lpPathNameStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CONST DYNAMIC_TIME_ZONE_INFORMATION *
// func SetDynamicTimeZoneInformation(lpTimeZoneInformation /*const*/ CONST DYNAMIC_TIME_ZONE_INFORMATION *) bool

func SetEndOfFile(hFile HANDLE) bool {
	ret1 := syscall3(setEndOfFile, 1,
		uintptr(hFile),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPWCH
// func SetEnvironmentStrings(newEnvironment LPWCH) bool

func SetEnvironmentVariable(lpName string, lpValue string) bool {
	lpNameStr := unicode16FromString(lpName)
	lpValueStr := unicode16FromString(lpValue)
	ret1 := syscall3(setEnvironmentVariable, 2,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(&lpValueStr[0])),
		0)
	return ret1 != 0
}

func SetErrorMode(uMode uint32) uint32 {
	ret1 := syscall3(setErrorMode, 1,
		uintptr(uMode),
		0,
		0)
	return uint32(ret1)
}

func SetEvent(hEvent HANDLE) bool {
	ret1 := syscall3(setEvent, 1,
		uintptr(hEvent),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PTP_CALLBACK_INSTANCE
// func SetEventWhenCallbackReturns(pci PTP_CALLBACK_INSTANCE, evt HANDLE)

func SetFileApisToANSI() {
	syscall3(setFileApisToANSI, 0,
		0,
		0,
		0)
}

func SetFileApisToOEM() {
	syscall3(setFileApisToOEM, 0,
		0,
		0,
		0)
}

func SetFileAttributesTransacted(lpFileName string, dwFileAttributes uint32, hTransaction HANDLE) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(setFileAttributesTransacted, 3,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwFileAttributes),
		uintptr(hTransaction))
	return ret1 != 0
}

func SetFileAttributes(lpFileName string, dwFileAttributes uint32) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(setFileAttributes, 2,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwFileAttributes),
		0)
	return ret1 != 0
}

func SetFileBandwidthReservation(hFile HANDLE, nPeriodMilliseconds uint32, nBytesPerPeriod uint32, bDiscardable bool, lpTransferSize *uint32, lpNumOutstandingRequests *uint32) bool {
	ret1 := syscall6(setFileBandwidthReservation, 6,
		uintptr(hFile),
		uintptr(nPeriodMilliseconds),
		uintptr(nBytesPerPeriod),
		getUintptrFromBool(bDiscardable),
		uintptr(unsafe.Pointer(lpTransferSize)),
		uintptr(unsafe.Pointer(lpNumOutstandingRequests)))
	return ret1 != 0
}

func SetFileCompletionNotificationModes(fileHandle HANDLE, flags UCHAR) bool {
	ret1 := syscall3(setFileCompletionNotificationModes, 2,
		uintptr(fileHandle),
		uintptr(flags),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): FILE_INFO_BY_HANDLE_CLASS
// func SetFileInformationByHandle(hFile HANDLE, fileInformationClass FILE_INFO_BY_HANDLE_CLASS, lpFileInformation LPVOID, dwBufferSize uint32) bool

// TODO: Unknown type(s): PUCHAR
// func SetFileIoOverlappedRange(fileHandle HANDLE, overlappedRangeStart PUCHAR, length ULONG) bool

func SetFilePointer(hFile HANDLE, lDistanceToMove LONG, lpDistanceToMoveHigh *int32, dwMoveMethod uint32) uint32 {
	ret1 := syscall6(setFilePointer, 4,
		uintptr(hFile),
		uintptr(lDistanceToMove),
		uintptr(unsafe.Pointer(lpDistanceToMoveHigh)),
		uintptr(dwMoveMethod),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PLARGE_INTEGER
// func SetFilePointerEx(hFile HANDLE, liDistanceToMove LARGE_INTEGER, lpNewFilePointer PLARGE_INTEGER, dwMoveMethod uint32) bool

func SetFileShortName(hFile HANDLE, lpShortName string) bool {
	lpShortNameStr := unicode16FromString(lpShortName)
	ret1 := syscall3(setFileShortName, 2,
		uintptr(hFile),
		uintptr(unsafe.Pointer(&lpShortNameStr[0])),
		0)
	return ret1 != 0
}

func SetFileTime(hFile HANDLE, lpCreationTime /*const*/ *FILETIME, lpLastAccessTime /*const*/ *FILETIME, lpLastWriteTime /*const*/ *FILETIME) bool {
	ret1 := syscall6(setFileTime, 4,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpLastAccessTime)),
		uintptr(unsafe.Pointer(lpLastWriteTime)),
		0,
		0)
	return ret1 != 0
}

func SetFileValidData(hFile HANDLE, validDataLength LONGLONG) bool {
	ret1 := syscall3(setFileValidData, 2,
		uintptr(hFile),
		uintptr(validDataLength),
		0)
	return ret1 != 0
}

func SetFirmwareEnvironmentVariable(lpName string, lpGuid string, pValue uintptr, nSize uint32) bool {
	lpNameStr := unicode16FromString(lpName)
	lpGuidStr := unicode16FromString(lpGuid)
	ret1 := syscall6(setFirmwareEnvironmentVariable, 4,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(&lpGuidStr[0])),
		pValue,
		uintptr(nSize),
		0,
		0)
	return ret1 != 0
}

func SetHandleCount(uNumber uint32) uint32 {
	ret1 := syscall3(setHandleCount, 1,
		uintptr(uNumber),
		0,
		0)
	return uint32(ret1)
}

func SetHandleInformation(hObject HANDLE, dwMask uint32, dwFlags uint32) bool {
	ret1 := syscall3(setHandleInformation, 3,
		uintptr(hObject),
		uintptr(dwMask),
		uintptr(dwFlags))
	return ret1 != 0
}

// TODO: Unknown type(s): JOBOBJECTINFOCLASS
// func SetInformationJobObject(hJob HANDLE, jobObjectInformationClass JOBOBJECTINFOCLASS, lpJobObjectInformation LPVOID, cbJobObjectInformationLength uint32) bool

func SetLastError(dwErrCode uint32) {
	syscall3(setLastError, 1,
		uintptr(dwErrCode),
		0,
		0)
}

func SetLocalTime(lpSystemTime /*const*/ *SYSTEMTIME) bool {
	ret1 := syscall3(setLocalTime, 1,
		uintptr(unsafe.Pointer(lpSystemTime)),
		0,
		0)
	return ret1 != 0
}

func SetLocaleInfo(locale LCID, lCType LCTYPE, lpLCData string) bool {
	lpLCDataStr := unicode16FromString(lpLCData)
	ret1 := syscall3(setLocaleInfo, 3,
		uintptr(locale),
		uintptr(lCType),
		uintptr(unsafe.Pointer(&lpLCDataStr[0])))
	return ret1 != 0
}

func SetMailslotInfo(hMailslot HANDLE, lReadTimeout uint32) bool {
	ret1 := syscall3(setMailslotInfo, 2,
		uintptr(hMailslot),
		uintptr(lReadTimeout),
		0)
	return ret1 != 0
}

func SetMessageWaitingIndicator(hMsgIndicator HANDLE, ulMsgCount ULONG) bool {
	ret1 := syscall3(setMessageWaitingIndicator, 2,
		uintptr(hMsgIndicator),
		uintptr(ulMsgCount),
		0)
	return ret1 != 0
}

func SetNamedPipeHandleState(hNamedPipe HANDLE, lpMode *uint32, lpMaxCollectionCount *uint32, lpCollectDataTimeout *uint32) bool {
	ret1 := syscall6(setNamedPipeHandleState, 4,
		uintptr(hNamedPipe),
		uintptr(unsafe.Pointer(lpMode)),
		uintptr(unsafe.Pointer(lpMaxCollectionCount)),
		uintptr(unsafe.Pointer(lpCollectDataTimeout)),
		0,
		0)
	return ret1 != 0
}

func SetPriorityClass(hProcess HANDLE, dwPriorityClass uint32) bool {
	ret1 := syscall3(setPriorityClass, 2,
		uintptr(hProcess),
		uintptr(dwPriorityClass),
		0)
	return ret1 != 0
}

func SetProcessAffinityMask(hProcess HANDLE, dwProcessAffinityMask *uint32) bool {
	ret1 := syscall3(setProcessAffinityMask, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(dwProcessAffinityMask)),
		0)
	return ret1 != 0
}

func SetProcessAffinityUpdateMode(hProcess HANDLE, dwFlags uint32) bool {
	ret1 := syscall3(setProcessAffinityUpdateMode, 2,
		uintptr(hProcess),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func SetProcessDEPPolicy(dwFlags uint32) bool {
	ret1 := syscall3(setProcessDEPPolicy, 1,
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCZZWSTR
// func SetProcessPreferredUILanguages(dwFlags uint32, pwszLanguagesBuffer PCZZWSTR, pulNumLanguages *uint32) bool

func SetProcessPriorityBoost(hProcess HANDLE, bDisablePriorityBoost bool) bool {
	ret1 := syscall3(setProcessPriorityBoost, 2,
		uintptr(hProcess),
		getUintptrFromBool(bDisablePriorityBoost),
		0)
	return ret1 != 0
}

func SetProcessShutdownParameters(dwLevel uint32, dwFlags uint32) bool {
	ret1 := syscall3(setProcessShutdownParameters, 2,
		uintptr(dwLevel),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func SetProcessWorkingSetSize(hProcess HANDLE, dwMinimumWorkingSetSize SIZE_T, dwMaximumWorkingSetSize SIZE_T) bool {
	ret1 := syscall3(setProcessWorkingSetSize, 3,
		uintptr(hProcess),
		uintptr(dwMinimumWorkingSetSize),
		uintptr(dwMaximumWorkingSetSize))
	return ret1 != 0
}

func SetProcessWorkingSetSizeEx(hProcess HANDLE, dwMinimumWorkingSetSize SIZE_T, dwMaximumWorkingSetSize SIZE_T, flags uint32) bool {
	ret1 := syscall6(setProcessWorkingSetSizeEx, 4,
		uintptr(hProcess),
		uintptr(dwMinimumWorkingSetSize),
		uintptr(dwMaximumWorkingSetSize),
		uintptr(flags),
		0,
		0)
	return ret1 != 0
}

func SetSearchPathMode(flags uint32) bool {
	ret1 := syscall3(setSearchPathMode, 1,
		uintptr(flags),
		0,
		0)
	return ret1 != 0
}

func SetStdHandle(nStdHandle uint32, hHandle HANDLE) bool {
	ret1 := syscall3(setStdHandle, 2,
		uintptr(nStdHandle),
		uintptr(hHandle),
		0)
	return ret1 != 0
}

func SetStdHandleEx(nStdHandle uint32, hHandle HANDLE, phPrevValue *HANDLE) bool {
	ret1 := syscall3(setStdHandleEx, 3,
		uintptr(nStdHandle),
		uintptr(hHandle),
		uintptr(unsafe.Pointer(phPrevValue)))
	return ret1 != 0
}

func SetSystemFileCacheSize(minimumFileCacheSize SIZE_T, maximumFileCacheSize SIZE_T, flags uint32) bool {
	ret1 := syscall3(setSystemFileCacheSize, 3,
		uintptr(minimumFileCacheSize),
		uintptr(maximumFileCacheSize),
		uintptr(flags))
	return ret1 != 0
}

func SetSystemPowerState(fSuspend bool, fForce bool) bool {
	ret1 := syscall3(setSystemPowerState, 2,
		getUintptrFromBool(fSuspend),
		getUintptrFromBool(fForce),
		0)
	return ret1 != 0
}

func SetSystemTime(lpSystemTime /*const*/ *SYSTEMTIME) bool {
	ret1 := syscall3(setSystemTime, 1,
		uintptr(unsafe.Pointer(lpSystemTime)),
		0,
		0)
	return ret1 != 0
}

func SetSystemTimeAdjustment(dwTimeAdjustment uint32, bTimeAdjustmentDisabled bool) bool {
	ret1 := syscall3(setSystemTimeAdjustment, 2,
		uintptr(dwTimeAdjustment),
		getUintptrFromBool(bTimeAdjustmentDisabled),
		0)
	return ret1 != 0
}

func SetTapeParameters(hDevice HANDLE, dwOperation uint32, lpTapeInformation LPVOID) uint32 {
	ret1 := syscall3(setTapeParameters, 3,
		uintptr(hDevice),
		uintptr(dwOperation),
		uintptr(unsafe.Pointer(lpTapeInformation)))
	return uint32(ret1)
}

func SetTapePosition(hDevice HANDLE, dwPositionMethod uint32, dwPartition uint32, dwOffsetLow uint32, dwOffsetHigh uint32, bImmediate bool) uint32 {
	ret1 := syscall6(setTapePosition, 6,
		uintptr(hDevice),
		uintptr(dwPositionMethod),
		uintptr(dwPartition),
		uintptr(dwOffsetLow),
		uintptr(dwOffsetHigh),
		getUintptrFromBool(bImmediate))
	return uint32(ret1)
}

func SetThreadAffinityMask(hThread HANDLE, dwThreadAffinityMask *uint32) *uint32 {
	ret1 := syscall3(setThreadAffinityMask, 2,
		uintptr(hThread),
		uintptr(unsafe.Pointer(dwThreadAffinityMask)),
		0)
	return (*uint32)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): CONST CONTEXT *
// func SetThreadContext(hThread HANDLE, lpContext /*const*/ CONST CONTEXT *) bool

func SetThreadErrorMode(dwNewMode uint32, lpOldMode *uint32) bool {
	ret1 := syscall3(setThreadErrorMode, 2,
		uintptr(dwNewMode),
		uintptr(unsafe.Pointer(lpOldMode)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): EXECUTION_STATE
// func SetThreadExecutionState(esFlags EXECUTION_STATE) EXECUTION_STATE

// TODO: Unknown type(s): CONST GROUP_AFFINITY *, PGROUP_AFFINITY
// func SetThreadGroupAffinity(hThread HANDLE, groupAffinity /*const*/ CONST GROUP_AFFINITY *, previousGroupAffinity PGROUP_AFFINITY) bool

func SetThreadIdealProcessor(hThread HANDLE, dwIdealProcessor uint32) uint32 {
	ret1 := syscall3(setThreadIdealProcessor, 2,
		uintptr(hThread),
		uintptr(dwIdealProcessor),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PPROCESSOR_NUMBER
// func SetThreadIdealProcessorEx(hThread HANDLE, lpIdealProcessor PPROCESSOR_NUMBER, lpPreviousIdealProcessor PPROCESSOR_NUMBER) bool

func SetThreadLocale(locale LCID) bool {
	ret1 := syscall3(setThreadLocale, 1,
		uintptr(locale),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCZZWSTR
// func SetThreadPreferredUILanguages(dwFlags uint32, pwszLanguagesBuffer PCZZWSTR, pulNumLanguages *uint32) bool

func SetThreadPriority(hThread HANDLE, nPriority int32) bool {
	ret1 := syscall3(setThreadPriority, 2,
		uintptr(hThread),
		uintptr(nPriority),
		0)
	return ret1 != 0
}

func SetThreadPriorityBoost(hThread HANDLE, bDisablePriorityBoost bool) bool {
	ret1 := syscall3(setThreadPriorityBoost, 2,
		uintptr(hThread),
		getUintptrFromBool(bDisablePriorityBoost),
		0)
	return ret1 != 0
}

func SetThreadStackGuarantee(stackSizeInBytes *uint32) bool {
	ret1 := syscall3(setThreadStackGuarantee, 1,
		uintptr(unsafe.Pointer(stackSizeInBytes)),
		0,
		0)
	return ret1 != 0
}

func SetThreadUILanguage(langId LANGID) LANGID {
	ret1 := syscall3(setThreadUILanguage, 1,
		uintptr(langId),
		0,
		0)
	return LANGID(ret1)
}

// TODO: Unknown type(s): PTP_POOL, PTP_POOL_STACK_INFORMATION
// func SetThreadpoolStackInformation(ptpp PTP_POOL, ptpsi PTP_POOL_STACK_INFORMATION) bool

// TODO: Unknown type(s): PTP_POOL
// func SetThreadpoolThreadMaximum(ptpp PTP_POOL, cthrdMost uint32)

// TODO: Unknown type(s): PTP_POOL
// func SetThreadpoolThreadMinimum(ptpp PTP_POOL, cthrdMic uint32) bool

// TODO: Unknown type(s): PFILETIME, PTP_TIMER
// func SetThreadpoolTimer(pti PTP_TIMER, pftDueTime PFILETIME, msPeriod uint32, msWindowLength uint32)

// TODO: Unknown type(s): PFILETIME, PTP_WAIT
// func SetThreadpoolWait(pwa PTP_WAIT, h HANDLE, pftTimeout PFILETIME)

// TODO: Unknown type(s): CONST TIME_ZONE_INFORMATION *
// func SetTimeZoneInformation(lpTimeZoneInformation /*const*/ CONST TIME_ZONE_INFORMATION *) bool

// TODO: Unknown type(s): WAITORTIMERCALLBACK
// func SetTimerQueueTimer(timerQueue HANDLE, callback WAITORTIMERCALLBACK, parameter uintptr, dueTime uint32, period uint32, preferIo bool) HANDLE

// TODO: Unknown type(s): LPTOP_LEVEL_EXCEPTION_FILTER
// func SetUnhandledExceptionFilter(lpTopLevelExceptionFilter LPTOP_LEVEL_EXCEPTION_FILTER) LPTOP_LEVEL_EXCEPTION_FILTER

func SetUserGeoID(geoId GEOID) bool {
	ret1 := syscall3(setUserGeoID, 1,
		uintptr(geoId),
		0,
		0)
	return ret1 != 0
}

func SetVolumeLabel(lpRootPathName string, lpVolumeName string) bool {
	lpRootPathNameStr := unicode16FromString(lpRootPathName)
	lpVolumeNameStr := unicode16FromString(lpVolumeName)
	ret1 := syscall3(setVolumeLabel, 2,
		uintptr(unsafe.Pointer(&lpRootPathNameStr[0])),
		uintptr(unsafe.Pointer(&lpVolumeNameStr[0])),
		0)
	return ret1 != 0
}

func SetVolumeMountPoint(lpszVolumeMountPoint string, lpszVolumeName string) bool {
	lpszVolumeMountPointStr := unicode16FromString(lpszVolumeMountPoint)
	lpszVolumeNameStr := unicode16FromString(lpszVolumeName)
	ret1 := syscall3(setVolumeMountPoint, 2,
		uintptr(unsafe.Pointer(&lpszVolumeMountPointStr[0])),
		uintptr(unsafe.Pointer(&lpszVolumeNameStr[0])),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PTIMERAPCROUTINE
// func SetWaitableTimer(hTimer HANDLE, lpDueTime /*const*/ *LARGE_INTEGER, lPeriod LONG, pfnCompletionRoutine PTIMERAPCROUTINE, lpArgToCompletionRoutine LPVOID, fResume bool) bool

// TODO: Unknown type(s): PREASON_CONTEXT, PTIMERAPCROUTINE
// func SetWaitableTimerEx(hTimer HANDLE, lpDueTime /*const*/ *LARGE_INTEGER, lPeriod LONG, pfnCompletionRoutine PTIMERAPCROUTINE, lpArgToCompletionRoutine LPVOID, wakeContext PREASON_CONTEXT, tolerableDelay ULONG) bool

func SetupComm(hFile HANDLE, dwInQueue uint32, dwOutQueue uint32) bool {
	ret1 := syscall3(setupComm, 3,
		uintptr(hFile),
		uintptr(dwInQueue),
		uintptr(dwOutQueue))
	return ret1 != 0
}

func SignalObjectAndWait(hObjectToSignal HANDLE, hObjectToWaitOn HANDLE, dwMilliseconds uint32, bAlertable bool) uint32 {
	ret1 := syscall6(signalObjectAndWait, 4,
		uintptr(hObjectToSignal),
		uintptr(hObjectToWaitOn),
		uintptr(dwMilliseconds),
		getUintptrFromBool(bAlertable),
		0,
		0)
	return uint32(ret1)
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
	ret1 := syscall3(sizeofResource, 2,
		uintptr(hModule),
		uintptr(hResInfo),
		0)
	return uint32(ret1)
}

func Sleep(dwMilliseconds uint32) {
	syscall3(sleep, 1,
		uintptr(dwMilliseconds),
		0,
		0)
}

// TODO: Unknown type(s): PCONDITION_VARIABLE, PCRITICAL_SECTION
// func SleepConditionVariableCS(conditionVariable PCONDITION_VARIABLE, criticalSection PCRITICAL_SECTION, dwMilliseconds uint32) bool

// TODO: Unknown type(s): PCONDITION_VARIABLE
// func SleepConditionVariableSRW(conditionVariable PCONDITION_VARIABLE, sRWLock PSRWLOCK, dwMilliseconds uint32, flags ULONG) bool

func SleepEx(dwMilliseconds uint32, bAlertable bool) uint32 {
	ret1 := syscall3(sleepEx, 2,
		uintptr(dwMilliseconds),
		getUintptrFromBool(bAlertable),
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): PTP_IO
// func StartThreadpoolIo(pio PTP_IO)

// TODO: Unknown type(s): PTP_WORK
// func SubmitThreadpoolWork(pwk PTP_WORK)

func SuspendThread(hThread HANDLE) uint32 {
	ret1 := syscall3(suspendThread, 1,
		uintptr(hThread),
		0,
		0)
	return uint32(ret1)
}

func SwitchToFiber(lpFiber LPVOID) {
	syscall3(switchToFiber, 1,
		uintptr(unsafe.Pointer(lpFiber)),
		0,
		0)
}

func SwitchToThread() bool {
	ret1 := syscall3(switchToThread, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func SystemTimeToFileTime(lpSystemTime /*const*/ *SYSTEMTIME, lpFileTime *FILETIME) bool {
	ret1 := syscall3(systemTimeToFileTime, 2,
		uintptr(unsafe.Pointer(lpSystemTime)),
		uintptr(unsafe.Pointer(lpFileTime)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CONST TIME_ZONE_INFORMATION *
// func SystemTimeToTzSpecificLocalTime(lpTimeZoneInformation /*const*/ CONST TIME_ZONE_INFORMATION *, lpUniversalTime /*const*/ *SYSTEMTIME, lpLocalTime *SYSTEMTIME) bool

func TerminateJobObject(hJob HANDLE, uExitCode uint32) bool {
	ret1 := syscall3(terminateJobObject, 2,
		uintptr(hJob),
		uintptr(uExitCode),
		0)
	return ret1 != 0
}

func TerminateProcess(hProcess HANDLE, uExitCode uint32) bool {
	ret1 := syscall3(terminateProcess, 2,
		uintptr(hProcess),
		uintptr(uExitCode),
		0)
	return ret1 != 0
}

func TerminateThread(hThread HANDLE, dwExitCode uint32) bool {
	ret1 := syscall3(terminateThread, 2,
		uintptr(hThread),
		uintptr(dwExitCode),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPTHREADENTRY32
// func Thread32First(hSnapshot HANDLE, lpte LPTHREADENTRY32) bool

// TODO: Unknown type(s): LPTHREADENTRY32
// func Thread32Next(hSnapshot HANDLE, lpte LPTHREADENTRY32) bool

func TlsAlloc() uint32 {
	ret1 := syscall3(tlsAlloc, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func TlsFree(dwTlsIndex uint32) bool {
	ret1 := syscall3(tlsFree, 1,
		uintptr(dwTlsIndex),
		0,
		0)
	return ret1 != 0
}

func TlsGetValue(dwTlsIndex uint32) LPVOID {
	ret1 := syscall3(tlsGetValue, 1,
		uintptr(dwTlsIndex),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func TlsSetValue(dwTlsIndex uint32, lpTlsValue LPVOID) bool {
	ret1 := syscall3(tlsSetValue, 2,
		uintptr(dwTlsIndex),
		uintptr(unsafe.Pointer(lpTlsValue)),
		0)
	return ret1 != 0
}

func Toolhelp32ReadProcessMemory(th32ProcessID uint32, lpBaseAddress /*const*/ uintptr, lpBuffer LPVOID, cbRead SIZE_T, lpNumberOfBytesRead *SIZE_T) bool {
	ret1 := syscall6(toolhelp32ReadProcessMemory, 5,
		uintptr(th32ProcessID),
		lpBaseAddress,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cbRead),
		uintptr(unsafe.Pointer(lpNumberOfBytesRead)),
		0)
	return ret1 != 0
}

func TransactNamedPipe(hNamedPipe HANDLE, lpInBuffer LPVOID, nInBufferSize uint32, lpOutBuffer LPVOID, nOutBufferSize uint32, lpBytesRead *uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall9(transactNamedPipe, 7,
		uintptr(hNamedPipe),
		uintptr(unsafe.Pointer(lpInBuffer)),
		uintptr(nInBufferSize),
		uintptr(unsafe.Pointer(lpOutBuffer)),
		uintptr(nOutBufferSize),
		uintptr(unsafe.Pointer(lpBytesRead)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0,
		0)
	return ret1 != 0
}

func TransmitCommChar(hFile HANDLE, cChar byte) bool {
	ret1 := syscall3(transmitCommChar, 2,
		uintptr(hFile),
		uintptr(cChar),
		0)
	return ret1 != 0
}

func TryAcquireSRWLockExclusive(sRWLock PSRWLOCK) BOOLEAN {
	ret1 := syscall3(tryAcquireSRWLockExclusive, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
	return BOOLEAN(ret1)
}

func TryAcquireSRWLockShared(sRWLock PSRWLOCK) BOOLEAN {
	ret1 := syscall3(tryAcquireSRWLockShared, 1,
		uintptr(unsafe.Pointer(sRWLock)),
		0,
		0)
	return BOOLEAN(ret1)
}

// TODO: Unknown type(s): LPCRITICAL_SECTION
// func TryEnterCriticalSection(lpCriticalSection LPCRITICAL_SECTION) bool

// TODO: Unknown type(s): PTP_CALLBACK_ENVIRON, PTP_SIMPLE_CALLBACK
// func TrySubmitThreadpoolCallback(pfns PTP_SIMPLE_CALLBACK, pv uintptr, pcbe PTP_CALLBACK_ENVIRON) bool

// TODO: Unknown type(s): CONST TIME_ZONE_INFORMATION *
// func TzSpecificLocalTimeToSystemTime(lpTimeZoneInformation /*const*/ CONST TIME_ZONE_INFORMATION *, lpLocalTime /*const*/ *SYSTEMTIME, lpUniversalTime *SYSTEMTIME) bool

// TODO: Unknown type(s): struct _EXCEPTION_POINTERS *
// func UnhandledExceptionFilter(exceptionInfo struct _EXCEPTION_POINTERS *) LONG

func UnlockFile(hFile HANDLE, dwFileOffsetLow uint32, dwFileOffsetHigh uint32, nNumberOfBytesToUnlockLow uint32, nNumberOfBytesToUnlockHigh uint32) bool {
	ret1 := syscall6(unlockFile, 5,
		uintptr(hFile),
		uintptr(dwFileOffsetLow),
		uintptr(dwFileOffsetHigh),
		uintptr(nNumberOfBytesToUnlockLow),
		uintptr(nNumberOfBytesToUnlockHigh),
		0)
	return ret1 != 0
}

func UnlockFileEx(hFile HANDLE, dwReserved uint32, nNumberOfBytesToUnlockLow uint32, nNumberOfBytesToUnlockHigh uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall6(unlockFileEx, 5,
		uintptr(hFile),
		uintptr(dwReserved),
		uintptr(nNumberOfBytesToUnlockLow),
		uintptr(nNumberOfBytesToUnlockHigh),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0)
	return ret1 != 0
}

func UnmapViewOfFile(lpBaseAddress /*const*/ uintptr) bool {
	ret1 := syscall3(unmapViewOfFile, 1,
		lpBaseAddress,
		0,
		0)
	return ret1 != 0
}

func UnregisterApplicationRecoveryCallback() HRESULT {
	ret1 := syscall3(unregisterApplicationRecoveryCallback, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func UnregisterApplicationRestart() HRESULT {
	ret1 := syscall3(unregisterApplicationRestart, 0,
		0,
		0,
		0)
	return HRESULT(ret1)
}

func UnregisterWait(waitHandle HANDLE) bool {
	ret1 := syscall3(unregisterWait, 1,
		uintptr(waitHandle),
		0,
		0)
	return ret1 != 0
}

func UnregisterWaitEx(waitHandle HANDLE, completionEvent HANDLE) bool {
	ret1 := syscall3(unregisterWaitEx, 2,
		uintptr(waitHandle),
		uintptr(completionEvent),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPPROC_THREAD_ATTRIBUTE_LIST, PSIZE_T
// func UpdateProcThreadAttribute(lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST, dwFlags uint32, attribute *uint32, lpValue uintptr, cbSize SIZE_T, lpPreviousValue uintptr, lpReturnSize PSIZE_T) bool

func UpdateResource(hUpdate HANDLE, lpType string, lpName string, wLanguage uint16, lpData LPVOID, cb uint32) bool {
	lpTypeStr := unicode16FromString(lpType)
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(updateResource, 6,
		uintptr(hUpdate),
		uintptr(unsafe.Pointer(&lpTypeStr[0])),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(wLanguage),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(cb))
	return ret1 != 0
}

func VerLanguageName(wLang uint32, szLang LPWSTR, nSize uint32) uint32 {
	ret1 := syscall3(verLanguageName, 3,
		uintptr(wLang),
		uintptr(unsafe.Pointer(szLang)),
		uintptr(nSize))
	return uint32(ret1)
}

func VerifyScripts(dwFlags uint32, lpLocaleScripts string, cchLocaleScripts int32, lpTestScripts string, cchTestScripts int32) bool {
	lpLocaleScriptsStr := unicode16FromString(lpLocaleScripts)
	lpTestScriptsStr := unicode16FromString(lpTestScripts)
	ret1 := syscall6(verifyScripts, 5,
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&lpLocaleScriptsStr[0])),
		uintptr(cchLocaleScripts),
		uintptr(unsafe.Pointer(&lpTestScriptsStr[0])),
		uintptr(cchTestScripts),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): DWORDLONG, LPOSVERSIONINFOEXW
// func VerifyVersionInfo(lpVersionInformation LPOSVERSIONINFOEXW, dwTypeMask uint32, dwlConditionMask DWORDLONG) bool

func VirtualAlloc(lpAddress LPVOID, dwSize SIZE_T, flAllocationType uint32, flProtect uint32) LPVOID {
	ret1 := syscall6(virtualAlloc, 4,
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(flAllocationType),
		uintptr(flProtect),
		0,
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func VirtualAllocEx(hProcess HANDLE, lpAddress LPVOID, dwSize SIZE_T, flAllocationType uint32, flProtect uint32) LPVOID {
	ret1 := syscall6(virtualAllocEx, 5,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(flAllocationType),
		uintptr(flProtect),
		0)
	return (LPVOID)(unsafe.Pointer(ret1))
}

func VirtualAllocExNuma(hProcess HANDLE, lpAddress LPVOID, dwSize SIZE_T, flAllocationType uint32, flProtect uint32, nndPreferred uint32) LPVOID {
	ret1 := syscall6(virtualAllocExNuma, 6,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(flAllocationType),
		uintptr(flProtect),
		uintptr(nndPreferred))
	return (LPVOID)(unsafe.Pointer(ret1))
}

func VirtualFree(lpAddress LPVOID, dwSize SIZE_T, dwFreeType uint32) bool {
	ret1 := syscall3(virtualFree, 3,
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(dwFreeType))
	return ret1 != 0
}

func VirtualFreeEx(hProcess HANDLE, lpAddress LPVOID, dwSize SIZE_T, dwFreeType uint32) bool {
	ret1 := syscall6(virtualFreeEx, 4,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(dwFreeType),
		0,
		0)
	return ret1 != 0
}

func VirtualLock(lpAddress LPVOID, dwSize SIZE_T) bool {
	ret1 := syscall3(virtualLock, 2,
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		0)
	return ret1 != 0
}

func VirtualProtect(lpAddress LPVOID, dwSize SIZE_T, flNewProtect uint32, lpflOldProtect *DWORD) bool {
	ret1 := syscall6(virtualProtect, 4,
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(unsafe.Pointer(lpflOldProtect)),
		0,
		0)
	return ret1 != 0
}

func VirtualProtectEx(hProcess HANDLE, lpAddress LPVOID, dwSize SIZE_T, flNewProtect uint32, lpflOldProtect *DWORD) bool {
	ret1 := syscall6(virtualProtectEx, 5,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(unsafe.Pointer(lpflOldProtect)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PMEMORY_BASIC_INFORMATION
// func VirtualQuery(lpAddress /*const*/ uintptr, lpBuffer PMEMORY_BASIC_INFORMATION, dwLength SIZE_T) SIZE_T

// TODO: Unknown type(s): PMEMORY_BASIC_INFORMATION
// func VirtualQueryEx(hProcess HANDLE, lpAddress /*const*/ uintptr, lpBuffer PMEMORY_BASIC_INFORMATION, dwLength SIZE_T) SIZE_T

func VirtualUnlock(lpAddress LPVOID, dwSize SIZE_T) bool {
	ret1 := syscall3(virtualUnlock, 2,
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(dwSize),
		0)
	return ret1 != 0
}

func WTSGetActiveConsoleSessionId() uint32 {
	ret1 := syscall3(wTSGetActiveConsoleSessionId, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func WaitCommEvent(hFile HANDLE, lpEvtMask *uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall3(waitCommEvent, 3,
		uintptr(hFile),
		uintptr(unsafe.Pointer(lpEvtMask)),
		uintptr(unsafe.Pointer(lpOverlapped)))
	return ret1 != 0
}

// TODO: Unknown type(s): LPDEBUG_EVENT
// func WaitForDebugEvent(lpDebugEvent LPDEBUG_EVENT, dwMilliseconds uint32) bool

func WaitForMultipleObjects(nCount uint32, lpHandles /*const*/ *HANDLE, bWaitAll bool, dwMilliseconds uint32) uint32 {
	ret1 := syscall6(waitForMultipleObjects, 4,
		uintptr(nCount),
		uintptr(unsafe.Pointer(lpHandles)),
		getUintptrFromBool(bWaitAll),
		uintptr(dwMilliseconds),
		0,
		0)
	return uint32(ret1)
}

func WaitForMultipleObjectsEx(nCount uint32, lpHandles /*const*/ *HANDLE, bWaitAll bool, dwMilliseconds uint32, bAlertable bool) uint32 {
	ret1 := syscall6(waitForMultipleObjectsEx, 5,
		uintptr(nCount),
		uintptr(unsafe.Pointer(lpHandles)),
		getUintptrFromBool(bWaitAll),
		uintptr(dwMilliseconds),
		getUintptrFromBool(bAlertable),
		0)
	return uint32(ret1)
}

func WaitForSingleObject(hHandle HANDLE, dwMilliseconds uint32) uint32 {
	ret1 := syscall3(waitForSingleObject, 2,
		uintptr(hHandle),
		uintptr(dwMilliseconds),
		0)
	return uint32(ret1)
}

func WaitForSingleObjectEx(hHandle HANDLE, dwMilliseconds uint32, bAlertable bool) uint32 {
	ret1 := syscall3(waitForSingleObjectEx, 3,
		uintptr(hHandle),
		uintptr(dwMilliseconds),
		getUintptrFromBool(bAlertable))
	return uint32(ret1)
}

// TODO: Unknown type(s): PTP_IO
// func WaitForThreadpoolIoCallbacks(pio PTP_IO, fCancelPendingCallbacks bool)

// TODO: Unknown type(s): PTP_TIMER
// func WaitForThreadpoolTimerCallbacks(pti PTP_TIMER, fCancelPendingCallbacks bool)

// TODO: Unknown type(s): PTP_WAIT
// func WaitForThreadpoolWaitCallbacks(pwa PTP_WAIT, fCancelPendingCallbacks bool)

// TODO: Unknown type(s): PTP_WORK
// func WaitForThreadpoolWorkCallbacks(pwk PTP_WORK, fCancelPendingCallbacks bool)

func WaitNamedPipe(lpNamedPipeName string, nTimeOut uint32) bool {
	lpNamedPipeNameStr := unicode16FromString(lpNamedPipeName)
	ret1 := syscall3(waitNamedPipe, 2,
		uintptr(unsafe.Pointer(&lpNamedPipeNameStr[0])),
		uintptr(nTimeOut),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCONDITION_VARIABLE
// func WakeAllConditionVariable(conditionVariable PCONDITION_VARIABLE)

// TODO: Unknown type(s): PCONDITION_VARIABLE
// func WakeConditionVariable(conditionVariable PCONDITION_VARIABLE)

func WerGetFlags(hProcess HANDLE, pdwFlags *DWORD) HRESULT {
	ret1 := syscall3(werGetFlags, 2,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(pdwFlags)),
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): WER_REGISTER_FILE_TYPE
// func WerRegisterFile(pwzFile string, regFileType WER_REGISTER_FILE_TYPE, dwFlags uint32) HRESULT

func WerRegisterMemoryBlock(pvAddress uintptr, dwSize uint32) HRESULT {
	ret1 := syscall3(werRegisterMemoryBlock, 2,
		pvAddress,
		uintptr(dwSize),
		0)
	return HRESULT(ret1)
}

func WerRegisterRuntimeExceptionModule(pwszOutOfProcessCallbackDll string, pContext uintptr) HRESULT {
	pwszOutOfProcessCallbackDllStr := unicode16FromString(pwszOutOfProcessCallbackDll)
	ret1 := syscall3(werRegisterRuntimeExceptionModule, 2,
		uintptr(unsafe.Pointer(&pwszOutOfProcessCallbackDllStr[0])),
		pContext,
		0)
	return HRESULT(ret1)
}

func WerSetFlags(dwFlags uint32) HRESULT {
	ret1 := syscall3(werSetFlags, 1,
		uintptr(dwFlags),
		0,
		0)
	return HRESULT(ret1)
}

func WerUnregisterFile(pwzFilePath string) HRESULT {
	pwzFilePathStr := unicode16FromString(pwzFilePath)
	ret1 := syscall3(werUnregisterFile, 1,
		uintptr(unsafe.Pointer(&pwzFilePathStr[0])),
		0,
		0)
	return HRESULT(ret1)
}

func WerUnregisterMemoryBlock(pvAddress uintptr) HRESULT {
	ret1 := syscall3(werUnregisterMemoryBlock, 1,
		pvAddress,
		0,
		0)
	return HRESULT(ret1)
}

func WerUnregisterRuntimeExceptionModule(pwszOutOfProcessCallbackDll string, pContext uintptr) HRESULT {
	pwszOutOfProcessCallbackDllStr := unicode16FromString(pwszOutOfProcessCallbackDll)
	ret1 := syscall3(werUnregisterRuntimeExceptionModule, 2,
		uintptr(unsafe.Pointer(&pwszOutOfProcessCallbackDllStr[0])),
		pContext,
		0)
	return HRESULT(ret1)
}

// TODO: Unknown type(s): LPCCH, LPCWCH
// func WideCharToMultiByte(codePage uint32, dwFlags uint32, lpWideCharStr LPCWCH, cchWideChar int32, lpMultiByteStr LPSTR, cbMultiByte int32, lpDefaultChar LPCCH, lpUsedDefaultChar *BOOL) int32

func WinExec(lpCmdLine /*const*/ LPCSTR, uCmdShow uint32) uint32 {
	ret1 := syscall3(winExec, 2,
		uintptr(unsafe.Pointer(lpCmdLine)),
		uintptr(uCmdShow),
		0)
	return uint32(ret1)
}

func Wow64DisableWow64FsRedirection(oldValue *PVOID) bool {
	ret1 := syscall3(wow64DisableWow64FsRedirection, 1,
		uintptr(unsafe.Pointer(oldValue)),
		0,
		0)
	return ret1 != 0
}

func Wow64EnableWow64FsRedirection(wow64FsEnableRedirection BOOLEAN) BOOLEAN {
	ret1 := syscall3(wow64EnableWow64FsRedirection, 1,
		uintptr(wow64FsEnableRedirection),
		0,
		0)
	return BOOLEAN(ret1)
}

// TODO: Unknown type(s): PWOW64_CONTEXT
// func Wow64GetThreadContext(hThread HANDLE, lpContext PWOW64_CONTEXT) bool

// TODO: Unknown type(s): PWOW64_LDT_ENTRY
// func Wow64GetThreadSelectorEntry(hThread HANDLE, dwSelector uint32, lpSelectorEntry PWOW64_LDT_ENTRY) bool

func Wow64RevertWow64FsRedirection(olValue uintptr) bool {
	ret1 := syscall3(wow64RevertWow64FsRedirection, 1,
		olValue,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CONST WOW64_CONTEXT *
// func Wow64SetThreadContext(hThread HANDLE, lpContext /*const*/ CONST WOW64_CONTEXT *) bool

func Wow64SuspendThread(hThread HANDLE) uint32 {
	ret1 := syscall3(wow64SuspendThread, 1,
		uintptr(hThread),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): CONST INPUT_RECORD *
// func WriteConsoleInput(hConsoleInput HANDLE, lpBuffer /*const*/ CONST INPUT_RECORD *, nLength uint32, lpNumberOfEventsWritten *uint32) bool

func WriteConsoleOutputAttribute(hConsoleOutput HANDLE, lpAttribute /*const*/ *WORD, nLength uint32, dwWriteCoord COORD, lpNumberOfAttrsWritten *uint32) bool {
	ret1 := syscall6(writeConsoleOutputAttribute, 5,
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(lpAttribute)),
		uintptr(nLength),
		getUintptrFromCOORD(dwWriteCoord),
		uintptr(unsafe.Pointer(lpNumberOfAttrsWritten)),
		0)
	return ret1 != 0
}

func WriteConsoleOutputCharacter(hConsoleOutput HANDLE, lpCharacter string, nLength uint32, dwWriteCoord COORD, lpNumberOfCharsWritten *uint32) bool {
	lpCharacterStr := unicode16FromString(lpCharacter)
	ret1 := syscall6(writeConsoleOutputCharacter, 5,
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(&lpCharacterStr[0])),
		uintptr(nLength),
		getUintptrFromCOORD(dwWriteCoord),
		uintptr(unsafe.Pointer(lpNumberOfCharsWritten)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CONST CHAR_INFO *, PSMALL_RECT
// func WriteConsoleOutput(hConsoleOutput HANDLE, lpBuffer /*const*/ CONST CHAR_INFO *, dwBufferSize COORD, dwBufferCoord COORD, lpWriteRegion PSMALL_RECT) bool

func WriteConsole(hConsoleOutput HANDLE, lpBuffer /*const*/ uintptr, nNumberOfCharsToWrite uint32, lpNumberOfCharsWritten *uint32, lpReserved LPVOID) bool {
	ret1 := syscall6(writeConsole, 5,
		uintptr(hConsoleOutput),
		lpBuffer,
		uintptr(nNumberOfCharsToWrite),
		uintptr(unsafe.Pointer(lpNumberOfCharsWritten)),
		uintptr(unsafe.Pointer(lpReserved)),
		0)
	return ret1 != 0
}

func WriteFile(hFile HANDLE, lpBuffer /*const*/ uintptr, nNumberOfBytesToWrite uint32, lpNumberOfBytesWritten *uint32, lpOverlapped *OVERLAPPED) bool {
	ret1 := syscall6(writeFile, 5,
		uintptr(hFile),
		lpBuffer,
		uintptr(nNumberOfBytesToWrite),
		uintptr(unsafe.Pointer(lpNumberOfBytesWritten)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPOVERLAPPED_COMPLETION_ROUTINE
// func WriteFileEx(hFile HANDLE, lpBuffer /*const*/ uintptr, nNumberOfBytesToWrite uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) bool

// TODO: Unknown type(s): FILE_SEGMENT_ELEMENT*
// func WriteFileGather(hFile HANDLE, aSegmentArray FILE_SEGMENT_ELEMENT*, nNumberOfBytesToWrite uint32, lpReserved *uint32, lpOverlapped *OVERLAPPED) bool

func WritePrivateProfileSection(lpAppName string, lpString string, lpFileName string) bool {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpStringStr := unicode16FromString(lpString)
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(writePrivateProfileSection, 3,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])))
	return ret1 != 0
}

func WritePrivateProfileString(lpAppName string, lpKeyName string, lpString string, lpFileName string) bool {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpKeyNameStr := unicode16FromString(lpKeyName)
	lpStringStr := unicode16FromString(lpString)
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(writePrivateProfileString, 4,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpKeyNameStr[0])),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return ret1 != 0
}

func WritePrivateProfileStruct(lpszSection string, lpszKey string, lpStruct LPVOID, uSizeStruct uint32, szFile string) bool {
	lpszSectionStr := unicode16FromString(lpszSection)
	lpszKeyStr := unicode16FromString(lpszKey)
	szFileStr := unicode16FromString(szFile)
	ret1 := syscall6(writePrivateProfileStruct, 5,
		uintptr(unsafe.Pointer(&lpszSectionStr[0])),
		uintptr(unsafe.Pointer(&lpszKeyStr[0])),
		uintptr(unsafe.Pointer(lpStruct)),
		uintptr(uSizeStruct),
		uintptr(unsafe.Pointer(&szFileStr[0])),
		0)
	return ret1 != 0
}

func WriteProcessMemory(hProcess HANDLE, lpBaseAddress LPVOID, lpBuffer /*const*/ uintptr, nSize SIZE_T, lpNumberOfBytesWritten *SIZE_T) bool {
	ret1 := syscall6(writeProcessMemory, 5,
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpBaseAddress)),
		lpBuffer,
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpNumberOfBytesWritten)),
		0)
	return ret1 != 0
}

func WriteProfileSection(lpAppName string, lpString string) bool {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(writeProfileSection, 2,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0)
	return ret1 != 0
}

func WriteProfileString(lpAppName string, lpKeyName string, lpString string) bool {
	lpAppNameStr := unicode16FromString(lpAppName)
	lpKeyNameStr := unicode16FromString(lpKeyName)
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(writeProfileString, 3,
		uintptr(unsafe.Pointer(&lpAppNameStr[0])),
		uintptr(unsafe.Pointer(&lpKeyNameStr[0])),
		uintptr(unsafe.Pointer(&lpStringStr[0])))
	return ret1 != 0
}

func WriteTapemark(hDevice HANDLE, dwTapemarkType uint32, dwTapemarkCount uint32, bImmediate bool) uint32 {
	ret1 := syscall6(writeTapemark, 4,
		uintptr(hDevice),
		uintptr(dwTapemarkType),
		uintptr(dwTapemarkCount),
		getUintptrFromBool(bImmediate),
		0,
		0)
	return uint32(ret1)
}

func ZombifyActCtx(hActCtx HANDLE) bool {
	ret1 := syscall3(zombifyActCtx, 1,
		uintptr(hActCtx),
		0,
		0)
	return ret1 != 0
}

func Lstrcat(lpString1 LPWSTR, lpString2 string) LPWSTR {
	lpString2Str := unicode16FromString(lpString2)
	ret1 := syscall3(lstrcat, 2,
		uintptr(unsafe.Pointer(lpString1)),
		uintptr(unsafe.Pointer(&lpString2Str[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func Lstrcmp(lpString1 string, lpString2 string) int32 {
	lpString1Str := unicode16FromString(lpString1)
	lpString2Str := unicode16FromString(lpString2)
	ret1 := syscall3(lstrcmp, 2,
		uintptr(unsafe.Pointer(&lpString1Str[0])),
		uintptr(unsafe.Pointer(&lpString2Str[0])),
		0)
	return int32(ret1)
}

func Lstrcmpi(lpString1 string, lpString2 string) int32 {
	lpString1Str := unicode16FromString(lpString1)
	lpString2Str := unicode16FromString(lpString2)
	ret1 := syscall3(lstrcmpi, 2,
		uintptr(unsafe.Pointer(&lpString1Str[0])),
		uintptr(unsafe.Pointer(&lpString2Str[0])),
		0)
	return int32(ret1)
}

func Lstrcpy(lpString1 LPWSTR, lpString2 string) LPWSTR {
	lpString2Str := unicode16FromString(lpString2)
	ret1 := syscall3(lstrcpy, 2,
		uintptr(unsafe.Pointer(lpString1)),
		uintptr(unsafe.Pointer(&lpString2Str[0])),
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func Lstrcpyn(lpString1 LPWSTR, lpString2 string, iMaxLength int32) LPWSTR {
	lpString2Str := unicode16FromString(lpString2)
	ret1 := syscall3(lstrcpyn, 3,
		uintptr(unsafe.Pointer(lpString1)),
		uintptr(unsafe.Pointer(&lpString2Str[0])),
		uintptr(iMaxLength))
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func Lstrlen(lpString string) int32 {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall3(lstrlen, 1,
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		0,
		0)
	return int32(ret1)
}

// TODO: Unknown type(s): PCUWSTR
// func Uaw_lstrcmpW(string1 PCUWSTR, string2 PCUWSTR) int32

// TODO: Unknown type(s): PCUWSTR
// func Uaw_lstrcmpiW(string1 PCUWSTR, string2 PCUWSTR) int32

// TODO: Unknown type(s): LPCUWSTR
// func Uaw_lstrlenW(string LPCUWSTR) int32

func CloseConsoleHandle(handle HANDLE) bool {
	ret1 := syscall3(closeConsoleHandle, 1,
		uintptr(handle),
		0,
		0)
	return ret1 != 0
}

func CloseProfileUserMapping() bool {
	ret1 := syscall3(closeProfileUserMapping, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func CmdBatNotification(bBatchRunning bool) bool {
	ret1 := syscall3(cmdBatNotification, 1,
		getUintptrFromBool(bBatchRunning),
		0,
		0)
	return ret1 != 0
}

func DelayLoadFailureHook(name /*const*/ LPCSTR, function /*const*/ LPCSTR) FARPROC {
	ret1 := syscall3(delayLoadFailureHook, 2,
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(function)),
		0)
	return func() INT_PTR {
		ret2 := syscall3(ret1, 0,
			0,
			0,
			0)
		return (INT_PTR)(unsafe.Pointer(ret2))
	}
}

func DuplicateConsoleHandle(handle HANDLE, access uint32, inherit bool, options uint32) HANDLE {
	ret1 := syscall6(duplicateConsoleHandle, 4,
		uintptr(handle),
		uintptr(access),
		getUintptrFromBool(inherit),
		uintptr(options),
		0,
		0)
	return HANDLE(ret1)
}

func ExpungeConsoleCommandHistory(unknown string) {
	unknownStr := unicode16FromString(unknown)
	syscall3(expungeConsoleCommandHistory, 1,
		uintptr(unsafe.Pointer(&unknownStr[0])),
		0,
		0)
}

func GetConsoleCommandHistoryLength(unknown string) uint32 {
	unknownStr := unicode16FromString(unknown)
	ret1 := syscall3(getConsoleCommandHistoryLength, 1,
		uintptr(unsafe.Pointer(&unknownStr[0])),
		0,
		0)
	return uint32(ret1)
}

func GetConsoleCommandHistory(unknown1 uint32, unknown2 uint32, unknown3 uint32) uint32 {
	ret1 := syscall3(getConsoleCommandHistory, 3,
		uintptr(unknown1),
		uintptr(unknown2),
		uintptr(unknown3))
	return uint32(ret1)
}

// TODO: Unknown type(s): CONSOLE_FONT_INFO *
// func GetConsoleFontInfo(hConsole HANDLE, maximize bool, numfonts uint32, info CONSOLE_FONT_INFO *) bool

func GetConsoleInputExeName(buflen uint32, buffer LPWSTR) bool {
	ret1 := syscall3(getConsoleInputExeName, 2,
		uintptr(buflen),
		uintptr(unsafe.Pointer(buffer)),
		0)
	return ret1 != 0
}

func GetConsoleInputWaitHandle() HANDLE {
	ret1 := syscall3(getConsoleInputWaitHandle, 0,
		0,
		0,
		0)
	return HANDLE(ret1)
}

func GetConsoleKeyboardLayoutName(layoutName LPWSTR) bool {
	ret1 := syscall3(getConsoleKeyboardLayoutName, 1,
		uintptr(unsafe.Pointer(layoutName)),
		0,
		0)
	return ret1 != 0
}

func GetNumberOfConsoleFonts() uint32 {
	ret1 := syscall3(getNumberOfConsoleFonts, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func K32EmptyWorkingSet(hProcess HANDLE) bool {
	ret1 := syscall3(k32EmptyWorkingSet, 1,
		uintptr(hProcess),
		0,
		0)
	return ret1 != 0
}

func K32EnumDeviceDrivers(image_base uintptr, cb uint32, needed *uint32) bool {
	ret1 := syscall3(k32EnumDeviceDrivers, 3,
		image_base,
		uintptr(cb),
		uintptr(unsafe.Pointer(needed)))
	return ret1 != 0
}

func K32EnumPageFiles(callback PENUM_PAGE_FILE_CALLBACK, context LPVOID) bool {
	callbackCallback := syscall.NewCallback(func(pContextRawArg LPVOID, pPageFileInfoRawArg PENUM_PAGE_FILE_INFORMATION, lpFilenameRawArg /*const*/ *uint16) uintptr {
		lpFilename := stringFromUnicode16(lpFilenameRawArg)
		ret := callback(pContextRawArg, pPageFileInfoRawArg, lpFilename)
		return uintptr(ret)
	})
	ret1 := syscall3(k32EnumPageFiles, 2,
		callbackCallback,
		uintptr(unsafe.Pointer(context)),
		0)
	return ret1 != 0
}

func K32EnumProcessModules(process HANDLE, lphModule *HMODULE, cb uint32, needed *uint32) bool {
	ret1 := syscall6(k32EnumProcessModules, 4,
		uintptr(process),
		uintptr(unsafe.Pointer(lphModule)),
		uintptr(cb),
		uintptr(unsafe.Pointer(needed)),
		0,
		0)
	return ret1 != 0
}

func K32EnumProcessModulesEx(process HANDLE, lphModule *HMODULE, cb uint32, needed *uint32, filter uint32) bool {
	ret1 := syscall6(k32EnumProcessModulesEx, 5,
		uintptr(process),
		uintptr(unsafe.Pointer(lphModule)),
		uintptr(cb),
		uintptr(unsafe.Pointer(needed)),
		uintptr(filter),
		0)
	return ret1 != 0
}

func K32EnumProcesses(lpdwProcessIDs *uint32, cb uint32, lpcbUsed *uint32) bool {
	ret1 := syscall3(k32EnumProcesses, 3,
		uintptr(unsafe.Pointer(lpdwProcessIDs)),
		uintptr(cb),
		uintptr(unsafe.Pointer(lpcbUsed)))
	return ret1 != 0
}

func K32GetDeviceDriverBaseName(image_base uintptr, base_name LPWSTR, size uint32) uint32 {
	ret1 := syscall3(k32GetDeviceDriverBaseName, 3,
		image_base,
		uintptr(unsafe.Pointer(base_name)),
		uintptr(size))
	return uint32(ret1)
}

func K32GetDeviceDriverFileName(image_base uintptr, file_name LPWSTR, size uint32) uint32 {
	ret1 := syscall3(k32GetDeviceDriverFileName, 3,
		image_base,
		uintptr(unsafe.Pointer(file_name)),
		uintptr(size))
	return uint32(ret1)
}

func K32GetMappedFileName(process HANDLE, lpv LPVOID, file_name LPWSTR, size uint32) uint32 {
	ret1 := syscall6(k32GetMappedFileName, 4,
		uintptr(process),
		uintptr(unsafe.Pointer(lpv)),
		uintptr(unsafe.Pointer(file_name)),
		uintptr(size),
		0,
		0)
	return uint32(ret1)
}

func K32GetModuleBaseName(process HANDLE, module HMODULE, base_name LPWSTR, size uint32) uint32 {
	ret1 := syscall6(k32GetModuleBaseName, 4,
		uintptr(process),
		uintptr(module),
		uintptr(unsafe.Pointer(base_name)),
		uintptr(size),
		0,
		0)
	return uint32(ret1)
}

func K32GetModuleFileNameEx(process HANDLE, module HMODULE, file_name LPWSTR, size uint32) uint32 {
	ret1 := syscall6(k32GetModuleFileNameEx, 4,
		uintptr(process),
		uintptr(module),
		uintptr(unsafe.Pointer(file_name)),
		uintptr(size),
		0,
		0)
	return uint32(ret1)
}

func K32GetModuleInformation(process HANDLE, module HMODULE, modinfo *MODULEINFO, cb uint32) bool {
	ret1 := syscall6(k32GetModuleInformation, 4,
		uintptr(process),
		uintptr(module),
		uintptr(unsafe.Pointer(modinfo)),
		uintptr(cb),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PPERFORMANCE_INFORMATION
// func K32GetPerformanceInfo(info PPERFORMANCE_INFORMATION, size uint32) bool

func K32GetProcessImageFileName(process HANDLE, file LPWSTR, size uint32) uint32 {
	ret1 := syscall3(k32GetProcessImageFileName, 3,
		uintptr(process),
		uintptr(unsafe.Pointer(file)),
		uintptr(size))
	return uint32(ret1)
}

func K32GetProcessMemoryInfo(process HANDLE, pmc PPROCESS_MEMORY_COUNTERS, cb uint32) bool {
	ret1 := syscall3(k32GetProcessMemoryInfo, 3,
		uintptr(process),
		uintptr(unsafe.Pointer(pmc)),
		uintptr(cb))
	return ret1 != 0
}

func K32GetWsChanges(process HANDLE, watchinfo PPSAPI_WS_WATCH_INFORMATION, size uint32) bool {
	ret1 := syscall3(k32GetWsChanges, 3,
		uintptr(process),
		uintptr(unsafe.Pointer(watchinfo)),
		uintptr(size))
	return ret1 != 0
}

func K32InitializeProcessForWsWatch(process HANDLE) bool {
	ret1 := syscall3(k32InitializeProcessForWsWatch, 1,
		uintptr(process),
		0,
		0)
	return ret1 != 0
}

func K32QueryWorkingSet(process HANDLE, buffer LPVOID, size uint32) bool {
	ret1 := syscall3(k32QueryWorkingSet, 3,
		uintptr(process),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size))
	return ret1 != 0
}

func K32QueryWorkingSetEx(process HANDLE, buffer LPVOID, size uint32) bool {
	ret1 := syscall3(k32QueryWorkingSetEx, 3,
		uintptr(process),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size))
	return ret1 != 0
}

func OpenConsoleW(name string, access uint32, inherit bool, creation uint32) HANDLE {
	nameStr := unicode16FromString(name)
	ret1 := syscall6(openConsoleW, 4,
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(access),
		getUintptrFromBool(inherit),
		uintptr(creation),
		0,
		0)
	return HANDLE(ret1)
}

func OpenProfileUserMapping() bool {
	ret1 := syscall3(openProfileUserMapping, 0,
		0,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): WAITORTIMERCALLBACK
// func RegisterWaitForSingleObjectEx(hObject HANDLE, callback WAITORTIMERCALLBACK, context uintptr, dwMilliseconds ULONG, dwFlags ULONG) HANDLE

// TODO: Unknown type(s): CONTEXT *
// func RtlCaptureContext(context CONTEXT *)

func RtlCaptureStackBackTrace(skip ULONG, count ULONG, buffer *PVOID, hash *ULONG) USHORT {
	ret1 := syscall6(rtlCaptureStackBackTrace, 4,
		uintptr(skip),
		uintptr(count),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(hash)),
		0,
		0)
	return USHORT(ret1)
}

func RtlCompareMemory(source1 /*const*/ uintptr, source2 /*const*/ uintptr, length SIZE_T) SIZE_T {
	ret1 := syscall3(rtlCompareMemory, 3,
		source1,
		source2,
		uintptr(length))
	return SIZE_T(ret1)
}

func RtlCopyMemory(dest uintptr, src /*const*/ uintptr, aLen SIZE_T) {
	syscall3(rtlCopyMemory, 3,
		dest,
		src,
		uintptr(aLen))
}

func RtlFillMemory(destination uintptr, length SIZE_T, fill byte) {
	syscall3(rtlFillMemory, 3,
		destination,
		uintptr(length),
		uintptr(fill))
}

// TODO: Unknown type(s): PRUNTIME_FUNCTION, UNWIND_HISTORY_TABLE *
// func RtlLookupFunctionEntry(pc *uint32, base *uint32, table UNWIND_HISTORY_TABLE *) PRUNTIME_FUNCTION

func RtlMoveMemory(destination uintptr, source /*const*/ uintptr, length SIZE_T) {
	syscall3(rtlMoveMemory, 3,
		destination,
		source,
		uintptr(length))
}

func RtlPcToFileHeader(pc uintptr, address *PVOID) uintptr {
	ret1 := syscall3(rtlPcToFileHeader, 2,
		pc,
		uintptr(unsafe.Pointer(address)),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): EXCEPTION_RECORD *
// func RtlRaiseException(rec EXCEPTION_RECORD *)

// TODO: Unknown type(s): CONTEXT *, EXCEPTION_RECORD *
// func RtlRestoreContext(context CONTEXT *, rec EXCEPTION_RECORD *)

// TODO: Unknown type(s): EXCEPTION_RECORD *
// func RtlUnwind(endframe uintptr, target_ip uintptr, rec EXCEPTION_RECORD *, retval uintptr)

// TODO: Unknown type(s): CONTEXT *, EXCEPTION_RECORD *, UNWIND_HISTORY_TABLE *
// func RtlUnwindEx(end_frame uintptr, target_ip uintptr, rec EXCEPTION_RECORD *, retval uintptr, context CONTEXT *, table UNWIND_HISTORY_TABLE *)

// TODO: Unknown type(s): CONTEXT *, KNONVOLATILE_CONTEXT_POINTERS *, RUNTIME_FUNCTION *
// func RtlVirtualUnwind(aType ULONG, base ULONG64, pc ULONG64, function RUNTIME_FUNCTION *, context CONTEXT *, data *PVOID, frame_ret *ULONG64, ctx_ptr KNONVOLATILE_CONTEXT_POINTERS *) uintptr

func RtlZeroMemory(destination uintptr, length SIZE_T) {
	syscall3(rtlZeroMemory, 2,
		destination,
		uintptr(length),
		0)
}

func SetCPGlobal(acp uint32) uint32 {
	ret1 := syscall3(setCPGlobal, 1,
		uintptr(acp),
		0,
		0)
	return uint32(ret1)
}

func SetConsoleFont(hConsole HANDLE, index uint32) bool {
	ret1 := syscall3(setConsoleFont, 2,
		uintptr(hConsole),
		uintptr(index),
		0)
	return ret1 != 0
}

func SetConsoleIcon(icon HICON) bool {
	ret1 := syscall3(setConsoleIcon, 1,
		uintptr(icon),
		0,
		0)
	return ret1 != 0
}

func SetConsoleInputExeName(name string) bool {
	nameStr := unicode16FromString(name)
	ret1 := syscall3(setConsoleInputExeName, 1,
		uintptr(unsafe.Pointer(&nameStr[0])),
		0,
		0)
	return ret1 != 0
}

func SetConsoleKeyShortcuts(set bool, keys byte, a uintptr, b uint32) bool {
	ret1 := syscall6(setConsoleKeyShortcuts, 4,
		getUintptrFromBool(set),
		uintptr(keys),
		a,
		uintptr(b),
		0,
		0)
	return ret1 != 0
}

func SetTermsrvAppInstallMode(bInstallMode bool) uint32 {
	ret1 := syscall3(setTermsrvAppInstallMode, 1,
		getUintptrFromBool(bInstallMode),
		0,
		0)
	return uint32(ret1)
}

func TermsrvAppInstallMode() bool {
	ret1 := syscall3(termsrvAppInstallMode, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func UTRegister(hModule HMODULE, lpsz16BITDLL LPSTR, lpszInitName LPSTR, lpszProcName LPSTR, ppfn32Thunk *FARPROC, pfnUT32CallBack FARPROC, lpBuff LPVOID) bool {
	pfnUT32CallBackCallback := syscall.NewCallback(func() uintptr {
		ret := pfnUT32CallBack()
		return uintptr(unsafe.Pointer(ret))
	})
	ret1 := syscall9(uTRegister, 7,
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpsz16BITDLL)),
		uintptr(unsafe.Pointer(lpszInitName)),
		uintptr(unsafe.Pointer(lpszProcName)),
		uintptr(unsafe.Pointer(ppfn32Thunk)),
		pfnUT32CallBackCallback,
		uintptr(unsafe.Pointer(lpBuff)),
		0,
		0)
	return ret1 != 0
}

func UTUnRegister(hModule HMODULE) {
	syscall3(uTUnRegister, 1,
		uintptr(hModule),
		0,
		0)
}

func VerSetConditionMask(dwlConditionMask ULONGLONG, dwTypeBitMask uint32, dwConditionMask byte) ULONGLONG {
	ret1 := syscall3(verSetConditionMask, 3,
		uintptr(dwlConditionMask),
		uintptr(dwTypeBitMask),
		uintptr(dwConditionMask))
	return ULONGLONG(ret1)
}

func VerifyConsoleIoHandle(handle HANDLE) bool {
	ret1 := syscall3(verifyConsoleIoHandle, 1,
		uintptr(handle),
		0,
		0)
	return ret1 != 0
}
