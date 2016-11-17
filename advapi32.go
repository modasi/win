// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libadvapi32 uintptr

	// Functions
	abortSystemShutdown                                 uintptr
	accessCheck                                         uintptr
	accessCheckAndAuditAlarm                            uintptr
	accessCheckByType                                   uintptr
	accessCheckByTypeResultList                         uintptr
	addAccessAllowedAce                                 uintptr
	addAccessAllowedAceEx                               uintptr
	addAccessAllowedObjectAce                           uintptr
	addAccessDeniedAce                                  uintptr
	addAccessDeniedAceEx                                uintptr
	addAccessDeniedObjectAce                            uintptr
	addAce                                              uintptr
	addAuditAccessAce                                   uintptr
	addAuditAccessAceEx                                 uintptr
	addAuditAccessObjectAce                             uintptr
	addConditionalAce                                   uintptr
	addMandatoryAce                                     uintptr
	addUsersToEncryptedFile                             uintptr
	adjustTokenGroups                                   uintptr
	adjustTokenPrivileges                               uintptr
	allocateAndInitializeSid                            uintptr
	allocateLocallyUniqueId                             uintptr
	areAllAccessesGranted                               uintptr
	areAnyAccessesGranted                               uintptr
	auditComputeEffectivePolicyBySid                    uintptr
	auditComputeEffectivePolicyByToken                  uintptr
	auditEnumerateCategories                            uintptr
	auditEnumeratePerUserPolicy                         uintptr
	auditEnumerateSubCategories                         uintptr
	auditFree                                           uintptr
	auditLookupCategoryGuidFromCategoryId               uintptr
	auditLookupCategoryIdFromCategoryGuid               uintptr
	auditLookupCategoryName                             uintptr
	auditLookupSubCategoryName                          uintptr
	auditQueryPerUserPolicy                             uintptr
	auditQuerySecurity                                  uintptr
	auditQuerySystemPolicy                              uintptr
	auditSetPerUserPolicy                               uintptr
	auditSetSecurity                                    uintptr
	auditSetSystemPolicy                                uintptr
	backupEventLog                                      uintptr
	buildExplicitAccessWithName                         uintptr
	buildImpersonateExplicitAccessWithName              uintptr
	buildImpersonateTrustee                             uintptr
	buildSecurityDescriptor                             uintptr
	buildTrusteeWithName                                uintptr
	buildTrusteeWithObjectsAndName                      uintptr
	buildTrusteeWithObjectsAndSid                       uintptr
	buildTrusteeWithSid                                 uintptr
	changeServiceConfig2                                uintptr
	changeServiceConfig                                 uintptr
	checkTokenMembership                                uintptr
	clearEventLog                                       uintptr
	closeEncryptedFileRaw                               uintptr
	closeEventLog                                       uintptr
	closeServiceHandle                                  uintptr
	closeThreadWaitChainSession                         uintptr
	commandLineFromMsiDescriptor                        uintptr
	controlService                                      uintptr
	controlServiceEx                                    uintptr
	convertSecurityDescriptorToStringSecurityDescriptor uintptr
	convertSidToStringSid                               uintptr
	convertStringSecurityDescriptorToSecurityDescriptor uintptr
	convertStringSidToSid                               uintptr
	convertToAutoInheritPrivateObjectSecurity           uintptr
	copySid                                             uintptr
	createPrivateObjectSecurity                         uintptr
	createPrivateObjectSecurityEx                       uintptr
	createPrivateObjectSecurityWithMultipleInheritance  uintptr
	createProcessAsUser                                 uintptr
	createProcessWithLogonW                             uintptr
	createProcessWithTokenW                             uintptr
	createService                                       uintptr
	credDelete                                          uintptr
	credEnumerate                                       uintptr
	credFree                                            uintptr
	credGetSessionTypes                                 uintptr
	credIsMarshaledCredential                           uintptr
	credRename                                          uintptr
	credUnprotect                                       uintptr
	credWrite                                           uintptr
	cryptAcquireContext                                 uintptr
	cryptContextAddRef                                  uintptr
	cryptCreateHash                                     uintptr
	cryptDecrypt                                        uintptr
	cryptDeriveKey                                      uintptr
	cryptDestroyHash                                    uintptr
	cryptDestroyKey                                     uintptr
	cryptDuplicateHash                                  uintptr
	cryptDuplicateKey                                   uintptr
	cryptEncrypt                                        uintptr
	cryptEnumProviderTypes                              uintptr
	cryptEnumProviders                                  uintptr
	cryptExportKey                                      uintptr
	cryptGenKey                                         uintptr
	cryptGenRandom                                      uintptr
	cryptGetDefaultProvider                             uintptr
	cryptGetHashParam                                   uintptr
	cryptGetKeyParam                                    uintptr
	cryptGetProvParam                                   uintptr
	cryptGetUserKey                                     uintptr
	cryptHashData                                       uintptr
	cryptHashSessionKey                                 uintptr
	cryptImportKey                                      uintptr
	cryptReleaseContext                                 uintptr
	cryptSetHashParam                                   uintptr
	cryptSetKeyParam                                    uintptr
	cryptSetProvParam                                   uintptr
	cryptSetProviderEx                                  uintptr
	cryptSetProvider                                    uintptr
	cryptSignHash                                       uintptr
	cryptVerifySignature                                uintptr
	decryptFile                                         uintptr
	deleteAce                                           uintptr
	deleteService                                       uintptr
	deregisterEventSource                               uintptr
	destroyPrivateObjectSecurity                        uintptr
	duplicateEncryptionInfoFile                         uintptr
	duplicateToken                                      uintptr
	encryptFile                                         uintptr
	encryptionDisable                                   uintptr
	equalDomainSid                                      uintptr
	equalPrefixSid                                      uintptr
	equalSid                                            uintptr
	fileEncryptionStatus                                uintptr
	findFirstFreeAce                                    uintptr
	freeSid                                             uintptr
	getAce                                              uintptr
	getEventLogInformation                              uintptr
	getFileSecurity                                     uintptr
	getKernelObjectSecurity                             uintptr
	getLengthSid                                        uintptr
	getLocalManagedApplicationData                      uintptr
	getMultipleTrusteeOperation                         uintptr
	getMultipleTrustee                                  uintptr
	getNumberOfEventLogRecords                          uintptr
	getOldestEventLogRecord                             uintptr
	getPrivateObjectSecurity                            uintptr
	getSecurityDescriptorControl                        uintptr
	getSecurityDescriptorGroup                          uintptr
	getSecurityDescriptorLength                         uintptr
	getSecurityDescriptorOwner                          uintptr
	getServiceDisplayName                               uintptr
	getServiceKeyName                                   uintptr
	getSidIdentifierAuthority                           uintptr
	getSidLengthRequired                                uintptr
	getSidSubAuthority                                  uintptr
	getTrusteeForm                                      uintptr
	getTrusteeName                                      uintptr
	getTrusteeType                                      uintptr
	getUserName                                         uintptr
	getWindowsAccountDomainSid                          uintptr
	impersonateAnonymousToken                           uintptr
	impersonateLoggedOnUser                             uintptr
	impersonateNamedPipeClient                          uintptr
	impersonateSelf                                     uintptr
	initializeAcl                                       uintptr
	initializeSecurityDescriptor                        uintptr
	initializeSid                                       uintptr
	initiateShutdown                                    uintptr
	initiateSystemShutdownEx                            uintptr
	initiateSystemShutdown                              uintptr
	isTextUnicode                                       uintptr
	isTokenRestricted                                   uintptr
	isTokenUntrusted                                    uintptr
	isValidAcl                                          uintptr
	isValidSecurityDescriptor                           uintptr
	isValidSid                                          uintptr
	lockServiceDatabase                                 uintptr
	logonUser                                           uintptr
	lookupPrivilegeDisplayName                          uintptr
	lookupPrivilegeName                                 uintptr
	lookupPrivilegeValue                                uintptr
	makeAbsoluteSD                                      uintptr
	makeSelfRelativeSD                                  uintptr
	mapGenericMask                                      uintptr
	notifyBootConfigStatus                              uintptr
	notifyChangeEventLog                                uintptr
	objectCloseAuditAlarm                               uintptr
	objectDeleteAuditAlarm                              uintptr
	objectOpenAuditAlarm                                uintptr
	objectPrivilegeAuditAlarm                           uintptr
	openBackupEventLog                                  uintptr
	openEncryptedFileRaw                                uintptr
	openEventLog                                        uintptr
	openProcessToken                                    uintptr
	openSCManager                                       uintptr
	openService                                         uintptr
	openThreadToken                                     uintptr
	perfCreateInstance                                  uintptr
	perfDecrementULongCounterValue                      uintptr
	perfDecrementULongLongCounterValue                  uintptr
	perfDeleteInstance                                  uintptr
	perfIncrementULongCounterValue                      uintptr
	perfIncrementULongLongCounterValue                  uintptr
	perfQueryInstance                                   uintptr
	perfSetCounterRefValue                              uintptr
	perfSetULongCounterValue                            uintptr
	perfSetULongLongCounterValue                        uintptr
	perfStopProvider                                    uintptr
	privilegeCheck                                      uintptr
	privilegedServiceAuditAlarm                         uintptr
	querySecurityAccessMask                             uintptr
	queryServiceConfig2                                 uintptr
	queryServiceObjectSecurity                          uintptr
	queryServiceStatus                                  uintptr
	readEventLog                                        uintptr
	regCloseKey                                         uintptr
	regConnectRegistryEx                                uintptr
	regConnectRegistry                                  uintptr
	regCopyTree                                         uintptr
	regCreateKeyEx                                      uintptr
	regCreateKeyTransacted                              uintptr
	regCreateKey                                        uintptr
	regDeleteKeyEx                                      uintptr
	regDeleteKeyTransacted                              uintptr
	regDeleteKeyValue                                   uintptr
	regDeleteKey                                        uintptr
	regDeleteTree                                       uintptr
	regDeleteValue                                      uintptr
	regDisablePredefinedCache                           uintptr
	regDisablePredefinedCacheEx                         uintptr
	regDisableReflectionKey                             uintptr
	regEnableReflectionKey                              uintptr
	regEnumKey                                          uintptr
	regEnumValue                                        uintptr
	regFlushKey                                         uintptr
	regGetKeySecurity                                   uintptr
	regGetValue                                         uintptr
	regLoadAppKey                                       uintptr
	regLoadKey                                          uintptr
	regLoadMUIString                                    uintptr
	regNotifyChangeKeyValue                             uintptr
	regOpenCurrentUser                                  uintptr
	regOpenKeyEx                                        uintptr
	regOpenKeyTransacted                                uintptr
	regOpenKey                                          uintptr
	regOpenUserClassesRoot                              uintptr
	regOverridePredefKey                                uintptr
	regQueryReflectionKey                               uintptr
	regQueryValueEx                                     uintptr
	regQueryValue                                       uintptr
	regReplaceKey                                       uintptr
	regRestoreKey                                       uintptr
	regSaveKeyEx                                        uintptr
	regSaveKey                                          uintptr
	regSetKeySecurity                                   uintptr
	regSetKeyValue                                      uintptr
	regSetValueEx                                       uintptr
	regSetValue                                         uintptr
	regUnLoadKey                                        uintptr
	registerEventSource                                 uintptr
	registerServiceCtrlHandlerEx                        uintptr
	reportEvent                                         uintptr
	revertToSelf                                        uintptr
	saferCloseLevel                                     uintptr
	saferComputeTokenFromLevel                          uintptr
	saferCreateLevel                                    uintptr
	saferRecordEventLogEntry                            uintptr
	saferiIsExecutableFileType                          uintptr
	setFileSecurity                                     uintptr
	setKernelObjectSecurity                             uintptr
	setNamedSecurityInfo                                uintptr
	setPrivateObjectSecurity                            uintptr
	setPrivateObjectSecurityEx                          uintptr
	setSecurityAccessMask                               uintptr
	setSecurityDescriptorControl                        uintptr
	setSecurityDescriptorDacl                           uintptr
	setSecurityDescriptorGroup                          uintptr
	setSecurityDescriptorOwner                          uintptr
	setSecurityDescriptorSacl                           uintptr
	setSecurityInfo                                     uintptr
	setServiceBits                                      uintptr
	setServiceObjectSecurity                            uintptr
	setServiceStatus                                    uintptr
	setThreadToken                                      uintptr
	setUserFileEncryptionKey                            uintptr
	startService                                        uintptr
	uninstallApplication                                uintptr
	unlockServiceDatabase                               uintptr
	wow64Win32ApiEntry                                  uintptr
	eventActivityIdControl                              uintptr
	systemFunction030                                   uintptr
	systemFunction035                                   uintptr
	systemFunction036                                   uintptr
)

func init() {
	// Library
	libadvapi32 = doLoadLibrary("advapi32.dll")

	// Functions
	abortSystemShutdown = doGetProcAddress(libadvapi32, "AbortSystemShutdownW")
	accessCheck = doGetProcAddress(libadvapi32, "AccessCheck")
	accessCheckAndAuditAlarm = doGetProcAddress(libadvapi32, "AccessCheckAndAuditAlarmW")
	accessCheckByType = doGetProcAddress(libadvapi32, "AccessCheckByType")
	accessCheckByTypeResultList = doGetProcAddress(libadvapi32, "AccessCheckByTypeResultList")
	addAccessAllowedAce = doGetProcAddress(libadvapi32, "AddAccessAllowedAce")
	addAccessAllowedAceEx = doGetProcAddress(libadvapi32, "AddAccessAllowedAceEx")
	addAccessAllowedObjectAce = doGetProcAddress(libadvapi32, "AddAccessAllowedObjectAce")
	addAccessDeniedAce = doGetProcAddress(libadvapi32, "AddAccessDeniedAce")
	addAccessDeniedAceEx = doGetProcAddress(libadvapi32, "AddAccessDeniedAceEx")
	addAccessDeniedObjectAce = doGetProcAddress(libadvapi32, "AddAccessDeniedObjectAce")
	addAce = doGetProcAddress(libadvapi32, "AddAce")
	addAuditAccessAce = doGetProcAddress(libadvapi32, "AddAuditAccessAce")
	addAuditAccessAceEx = doGetProcAddress(libadvapi32, "AddAuditAccessAceEx")
	addAuditAccessObjectAce = doGetProcAddress(libadvapi32, "AddAuditAccessObjectAce")
	addConditionalAce = doGetProcAddress(libadvapi32, "AddConditionalAce")
	addMandatoryAce = doGetProcAddress(libadvapi32, "AddMandatoryAce")
	addUsersToEncryptedFile = doGetProcAddress(libadvapi32, "AddUsersToEncryptedFile")
	adjustTokenGroups = doGetProcAddress(libadvapi32, "AdjustTokenGroups")
	adjustTokenPrivileges = doGetProcAddress(libadvapi32, "AdjustTokenPrivileges")
	allocateAndInitializeSid = doGetProcAddress(libadvapi32, "AllocateAndInitializeSid")
	allocateLocallyUniqueId = doGetProcAddress(libadvapi32, "AllocateLocallyUniqueId")
	areAllAccessesGranted = doGetProcAddress(libadvapi32, "AreAllAccessesGranted")
	areAnyAccessesGranted = doGetProcAddress(libadvapi32, "AreAnyAccessesGranted")
	auditComputeEffectivePolicyBySid = doGetProcAddress(libadvapi32, "AuditComputeEffectivePolicyBySid")
	auditComputeEffectivePolicyByToken = doGetProcAddress(libadvapi32, "AuditComputeEffectivePolicyByToken")
	auditEnumerateCategories = doGetProcAddress(libadvapi32, "AuditEnumerateCategories")
	auditEnumeratePerUserPolicy = doGetProcAddress(libadvapi32, "AuditEnumeratePerUserPolicy")
	auditEnumerateSubCategories = doGetProcAddress(libadvapi32, "AuditEnumerateSubCategories")
	auditFree = doGetProcAddress(libadvapi32, "AuditFree")
	auditLookupCategoryGuidFromCategoryId = doGetProcAddress(libadvapi32, "AuditLookupCategoryGuidFromCategoryId")
	auditLookupCategoryIdFromCategoryGuid = doGetProcAddress(libadvapi32, "AuditLookupCategoryIdFromCategoryGuid")
	auditLookupCategoryName = doGetProcAddress(libadvapi32, "AuditLookupCategoryNameW")
	auditLookupSubCategoryName = doGetProcAddress(libadvapi32, "AuditLookupSubCategoryNameW")
	auditQueryPerUserPolicy = doGetProcAddress(libadvapi32, "AuditQueryPerUserPolicy")
	auditQuerySecurity = doGetProcAddress(libadvapi32, "AuditQuerySecurity")
	auditQuerySystemPolicy = doGetProcAddress(libadvapi32, "AuditQuerySystemPolicy")
	auditSetPerUserPolicy = doGetProcAddress(libadvapi32, "AuditSetPerUserPolicy")
	auditSetSecurity = doGetProcAddress(libadvapi32, "AuditSetSecurity")
	auditSetSystemPolicy = doGetProcAddress(libadvapi32, "AuditSetSystemPolicy")
	backupEventLog = doGetProcAddress(libadvapi32, "BackupEventLogW")
	buildExplicitAccessWithName = doGetProcAddress(libadvapi32, "BuildExplicitAccessWithNameW")
	buildImpersonateExplicitAccessWithName = doGetProcAddress(libadvapi32, "BuildImpersonateExplicitAccessWithNameW")
	buildImpersonateTrustee = doGetProcAddress(libadvapi32, "BuildImpersonateTrusteeW")
	buildSecurityDescriptor = doGetProcAddress(libadvapi32, "BuildSecurityDescriptorW")
	buildTrusteeWithName = doGetProcAddress(libadvapi32, "BuildTrusteeWithNameW")
	buildTrusteeWithObjectsAndName = doGetProcAddress(libadvapi32, "BuildTrusteeWithObjectsAndNameW")
	buildTrusteeWithObjectsAndSid = doGetProcAddress(libadvapi32, "BuildTrusteeWithObjectsAndSidW")
	buildTrusteeWithSid = doGetProcAddress(libadvapi32, "BuildTrusteeWithSidW")
	changeServiceConfig2 = doGetProcAddress(libadvapi32, "ChangeServiceConfig2W")
	changeServiceConfig = doGetProcAddress(libadvapi32, "ChangeServiceConfigW")
	checkTokenMembership = doGetProcAddress(libadvapi32, "CheckTokenMembership")
	clearEventLog = doGetProcAddress(libadvapi32, "ClearEventLogW")
	closeEncryptedFileRaw = doGetProcAddress(libadvapi32, "CloseEncryptedFileRaw")
	closeEventLog = doGetProcAddress(libadvapi32, "CloseEventLog")
	closeServiceHandle = doGetProcAddress(libadvapi32, "CloseServiceHandle")
	closeThreadWaitChainSession = doGetProcAddress(libadvapi32, "CloseThreadWaitChainSession")
	commandLineFromMsiDescriptor = doGetProcAddress(libadvapi32, "CommandLineFromMsiDescriptor")
	controlService = doGetProcAddress(libadvapi32, "ControlService")
	controlServiceEx = doGetProcAddress(libadvapi32, "ControlServiceExW")
	convertSecurityDescriptorToStringSecurityDescriptor = doGetProcAddress(libadvapi32, "ConvertSecurityDescriptorToStringSecurityDescriptorW")
	convertSidToStringSid = doGetProcAddress(libadvapi32, "ConvertSidToStringSidW")
	convertStringSecurityDescriptorToSecurityDescriptor = doGetProcAddress(libadvapi32, "ConvertStringSecurityDescriptorToSecurityDescriptorW")
	convertStringSidToSid = doGetProcAddress(libadvapi32, "ConvertStringSidToSidW")
	convertToAutoInheritPrivateObjectSecurity = doGetProcAddress(libadvapi32, "ConvertToAutoInheritPrivateObjectSecurity")
	copySid = doGetProcAddress(libadvapi32, "CopySid")
	createPrivateObjectSecurity = doGetProcAddress(libadvapi32, "CreatePrivateObjectSecurity")
	createPrivateObjectSecurityEx = doGetProcAddress(libadvapi32, "CreatePrivateObjectSecurityEx")
	createPrivateObjectSecurityWithMultipleInheritance = doGetProcAddress(libadvapi32, "CreatePrivateObjectSecurityWithMultipleInheritance")
	createProcessAsUser = doGetProcAddress(libadvapi32, "CreateProcessAsUserW")
	createProcessWithLogonW = doGetProcAddress(libadvapi32, "CreateProcessWithLogonW")
	createProcessWithTokenW = doGetProcAddress(libadvapi32, "CreateProcessWithTokenW")
	createService = doGetProcAddress(libadvapi32, "CreateServiceW")
	credDelete = doGetProcAddress(libadvapi32, "CredDeleteW")
	credEnumerate = doGetProcAddress(libadvapi32, "CredEnumerateW")
	credFree = doGetProcAddress(libadvapi32, "CredFree")
	credGetSessionTypes = doGetProcAddress(libadvapi32, "CredGetSessionTypes")
	credIsMarshaledCredential = doGetProcAddress(libadvapi32, "CredIsMarshaledCredentialW")
	credRename = doGetProcAddress(libadvapi32, "CredRenameW")
	credUnprotect = doGetProcAddress(libadvapi32, "CredUnprotectW")
	credWrite = doGetProcAddress(libadvapi32, "CredWriteW")
	cryptAcquireContext = doGetProcAddress(libadvapi32, "CryptAcquireContextW")
	cryptContextAddRef = doGetProcAddress(libadvapi32, "CryptContextAddRef")
	cryptCreateHash = doGetProcAddress(libadvapi32, "CryptCreateHash")
	cryptDecrypt = doGetProcAddress(libadvapi32, "CryptDecrypt")
	cryptDeriveKey = doGetProcAddress(libadvapi32, "CryptDeriveKey")
	cryptDestroyHash = doGetProcAddress(libadvapi32, "CryptDestroyHash")
	cryptDestroyKey = doGetProcAddress(libadvapi32, "CryptDestroyKey")
	cryptDuplicateHash = doGetProcAddress(libadvapi32, "CryptDuplicateHash")
	cryptDuplicateKey = doGetProcAddress(libadvapi32, "CryptDuplicateKey")
	cryptEncrypt = doGetProcAddress(libadvapi32, "CryptEncrypt")
	cryptEnumProviderTypes = doGetProcAddress(libadvapi32, "CryptEnumProviderTypesW")
	cryptEnumProviders = doGetProcAddress(libadvapi32, "CryptEnumProvidersW")
	cryptExportKey = doGetProcAddress(libadvapi32, "CryptExportKey")
	cryptGenKey = doGetProcAddress(libadvapi32, "CryptGenKey")
	cryptGenRandom = doGetProcAddress(libadvapi32, "CryptGenRandom")
	cryptGetDefaultProvider = doGetProcAddress(libadvapi32, "CryptGetDefaultProviderW")
	cryptGetHashParam = doGetProcAddress(libadvapi32, "CryptGetHashParam")
	cryptGetKeyParam = doGetProcAddress(libadvapi32, "CryptGetKeyParam")
	cryptGetProvParam = doGetProcAddress(libadvapi32, "CryptGetProvParam")
	cryptGetUserKey = doGetProcAddress(libadvapi32, "CryptGetUserKey")
	cryptHashData = doGetProcAddress(libadvapi32, "CryptHashData")
	cryptHashSessionKey = doGetProcAddress(libadvapi32, "CryptHashSessionKey")
	cryptImportKey = doGetProcAddress(libadvapi32, "CryptImportKey")
	cryptReleaseContext = doGetProcAddress(libadvapi32, "CryptReleaseContext")
	cryptSetHashParam = doGetProcAddress(libadvapi32, "CryptSetHashParam")
	cryptSetKeyParam = doGetProcAddress(libadvapi32, "CryptSetKeyParam")
	cryptSetProvParam = doGetProcAddress(libadvapi32, "CryptSetProvParam")
	cryptSetProviderEx = doGetProcAddress(libadvapi32, "CryptSetProviderExW")
	cryptSetProvider = doGetProcAddress(libadvapi32, "CryptSetProviderW")
	cryptSignHash = doGetProcAddress(libadvapi32, "CryptSignHashW")
	cryptVerifySignature = doGetProcAddress(libadvapi32, "CryptVerifySignatureW")
	decryptFile = doGetProcAddress(libadvapi32, "DecryptFileW")
	deleteAce = doGetProcAddress(libadvapi32, "DeleteAce")
	deleteService = doGetProcAddress(libadvapi32, "DeleteService")
	deregisterEventSource = doGetProcAddress(libadvapi32, "DeregisterEventSource")
	destroyPrivateObjectSecurity = doGetProcAddress(libadvapi32, "DestroyPrivateObjectSecurity")
	duplicateEncryptionInfoFile = doGetProcAddress(libadvapi32, "DuplicateEncryptionInfoFile")
	duplicateToken = doGetProcAddress(libadvapi32, "DuplicateToken")
	encryptFile = doGetProcAddress(libadvapi32, "EncryptFileW")
	encryptionDisable = doGetProcAddress(libadvapi32, "EncryptionDisable")
	equalDomainSid = doGetProcAddress(libadvapi32, "EqualDomainSid")
	equalPrefixSid = doGetProcAddress(libadvapi32, "EqualPrefixSid")
	equalSid = doGetProcAddress(libadvapi32, "EqualSid")
	fileEncryptionStatus = doGetProcAddress(libadvapi32, "FileEncryptionStatusW")
	findFirstFreeAce = doGetProcAddress(libadvapi32, "FindFirstFreeAce")
	freeSid = doGetProcAddress(libadvapi32, "FreeSid")
	getAce = doGetProcAddress(libadvapi32, "GetAce")
	getEventLogInformation = doGetProcAddress(libadvapi32, "GetEventLogInformation")
	getFileSecurity = doGetProcAddress(libadvapi32, "GetFileSecurityW")
	getKernelObjectSecurity = doGetProcAddress(libadvapi32, "GetKernelObjectSecurity")
	getLengthSid = doGetProcAddress(libadvapi32, "GetLengthSid")
	getLocalManagedApplicationData = doGetProcAddress(libadvapi32, "GetLocalManagedApplicationData")
	getMultipleTrusteeOperation = doGetProcAddress(libadvapi32, "GetMultipleTrusteeOperationW")
	getMultipleTrustee = doGetProcAddress(libadvapi32, "GetMultipleTrusteeW")
	getNumberOfEventLogRecords = doGetProcAddress(libadvapi32, "GetNumberOfEventLogRecords")
	getOldestEventLogRecord = doGetProcAddress(libadvapi32, "GetOldestEventLogRecord")
	getPrivateObjectSecurity = doGetProcAddress(libadvapi32, "GetPrivateObjectSecurity")
	getSecurityDescriptorControl = doGetProcAddress(libadvapi32, "GetSecurityDescriptorControl")
	getSecurityDescriptorGroup = doGetProcAddress(libadvapi32, "GetSecurityDescriptorGroup")
	getSecurityDescriptorLength = doGetProcAddress(libadvapi32, "GetSecurityDescriptorLength")
	getSecurityDescriptorOwner = doGetProcAddress(libadvapi32, "GetSecurityDescriptorOwner")
	getServiceDisplayName = doGetProcAddress(libadvapi32, "GetServiceDisplayNameW")
	getServiceKeyName = doGetProcAddress(libadvapi32, "GetServiceKeyNameW")
	getSidIdentifierAuthority = doGetProcAddress(libadvapi32, "GetSidIdentifierAuthority")
	getSidLengthRequired = doGetProcAddress(libadvapi32, "GetSidLengthRequired")
	getSidSubAuthority = doGetProcAddress(libadvapi32, "GetSidSubAuthority")
	getTrusteeForm = doGetProcAddress(libadvapi32, "GetTrusteeFormW")
	getTrusteeName = doGetProcAddress(libadvapi32, "GetTrusteeNameW")
	getTrusteeType = doGetProcAddress(libadvapi32, "GetTrusteeTypeW")
	getUserName = doGetProcAddress(libadvapi32, "GetUserNameW")
	getWindowsAccountDomainSid = doGetProcAddress(libadvapi32, "GetWindowsAccountDomainSid")
	impersonateAnonymousToken = doGetProcAddress(libadvapi32, "ImpersonateAnonymousToken")
	impersonateLoggedOnUser = doGetProcAddress(libadvapi32, "ImpersonateLoggedOnUser")
	impersonateNamedPipeClient = doGetProcAddress(libadvapi32, "ImpersonateNamedPipeClient")
	impersonateSelf = doGetProcAddress(libadvapi32, "ImpersonateSelf")
	initializeAcl = doGetProcAddress(libadvapi32, "InitializeAcl")
	initializeSecurityDescriptor = doGetProcAddress(libadvapi32, "InitializeSecurityDescriptor")
	initializeSid = doGetProcAddress(libadvapi32, "InitializeSid")
	initiateShutdown = doGetProcAddress(libadvapi32, "InitiateShutdownW")
	initiateSystemShutdownEx = doGetProcAddress(libadvapi32, "InitiateSystemShutdownExW")
	initiateSystemShutdown = doGetProcAddress(libadvapi32, "InitiateSystemShutdownW")
	isTextUnicode = doGetProcAddress(libadvapi32, "IsTextUnicode")
	isTokenRestricted = doGetProcAddress(libadvapi32, "IsTokenRestricted")
	isTokenUntrusted = doGetProcAddress(libadvapi32, "IsTokenUntrusted")
	isValidAcl = doGetProcAddress(libadvapi32, "IsValidAcl")
	isValidSecurityDescriptor = doGetProcAddress(libadvapi32, "IsValidSecurityDescriptor")
	isValidSid = doGetProcAddress(libadvapi32, "IsValidSid")
	lockServiceDatabase = doGetProcAddress(libadvapi32, "LockServiceDatabase")
	logonUser = doGetProcAddress(libadvapi32, "LogonUserW")
	lookupPrivilegeDisplayName = doGetProcAddress(libadvapi32, "LookupPrivilegeDisplayNameW")
	lookupPrivilegeName = doGetProcAddress(libadvapi32, "LookupPrivilegeNameW")
	lookupPrivilegeValue = doGetProcAddress(libadvapi32, "LookupPrivilegeValueW")
	makeAbsoluteSD = doGetProcAddress(libadvapi32, "MakeAbsoluteSD")
	makeSelfRelativeSD = doGetProcAddress(libadvapi32, "MakeSelfRelativeSD")
	mapGenericMask = doGetProcAddress(libadvapi32, "MapGenericMask")
	notifyBootConfigStatus = doGetProcAddress(libadvapi32, "NotifyBootConfigStatus")
	notifyChangeEventLog = doGetProcAddress(libadvapi32, "NotifyChangeEventLog")
	objectCloseAuditAlarm = doGetProcAddress(libadvapi32, "ObjectCloseAuditAlarmW")
	objectDeleteAuditAlarm = doGetProcAddress(libadvapi32, "ObjectDeleteAuditAlarmW")
	objectOpenAuditAlarm = doGetProcAddress(libadvapi32, "ObjectOpenAuditAlarmW")
	objectPrivilegeAuditAlarm = doGetProcAddress(libadvapi32, "ObjectPrivilegeAuditAlarmW")
	openBackupEventLog = doGetProcAddress(libadvapi32, "OpenBackupEventLogW")
	openEncryptedFileRaw = doGetProcAddress(libadvapi32, "OpenEncryptedFileRawW")
	openEventLog = doGetProcAddress(libadvapi32, "OpenEventLogW")
	openProcessToken = doGetProcAddress(libadvapi32, "OpenProcessToken")
	openSCManager = doGetProcAddress(libadvapi32, "OpenSCManagerW")
	openService = doGetProcAddress(libadvapi32, "OpenServiceW")
	openThreadToken = doGetProcAddress(libadvapi32, "OpenThreadToken")
	perfCreateInstance = doGetProcAddress(libadvapi32, "PerfCreateInstance")
	perfDecrementULongCounterValue = doGetProcAddress(libadvapi32, "PerfDecrementULongCounterValue")
	perfDecrementULongLongCounterValue = doGetProcAddress(libadvapi32, "PerfDecrementULongLongCounterValue")
	perfDeleteInstance = doGetProcAddress(libadvapi32, "PerfDeleteInstance")
	perfIncrementULongCounterValue = doGetProcAddress(libadvapi32, "PerfIncrementULongCounterValue")
	perfIncrementULongLongCounterValue = doGetProcAddress(libadvapi32, "PerfIncrementULongLongCounterValue")
	perfQueryInstance = doGetProcAddress(libadvapi32, "PerfQueryInstance")
	perfSetCounterRefValue = doGetProcAddress(libadvapi32, "PerfSetCounterRefValue")
	perfSetULongCounterValue = doGetProcAddress(libadvapi32, "PerfSetULongCounterValue")
	perfSetULongLongCounterValue = doGetProcAddress(libadvapi32, "PerfSetULongLongCounterValue")
	perfStopProvider = doGetProcAddress(libadvapi32, "PerfStopProvider")
	privilegeCheck = doGetProcAddress(libadvapi32, "PrivilegeCheck")
	privilegedServiceAuditAlarm = doGetProcAddress(libadvapi32, "PrivilegedServiceAuditAlarmW")
	querySecurityAccessMask = doGetProcAddress(libadvapi32, "QuerySecurityAccessMask")
	queryServiceConfig2 = doGetProcAddress(libadvapi32, "QueryServiceConfig2W")
	queryServiceObjectSecurity = doGetProcAddress(libadvapi32, "QueryServiceObjectSecurity")
	queryServiceStatus = doGetProcAddress(libadvapi32, "QueryServiceStatus")
	readEventLog = doGetProcAddress(libadvapi32, "ReadEventLogW")
	regCloseKey = doGetProcAddress(libadvapi32, "RegCloseKey")
	regConnectRegistryEx = doGetProcAddress(libadvapi32, "RegConnectRegistryExW")
	regConnectRegistry = doGetProcAddress(libadvapi32, "RegConnectRegistryW")
	regCopyTree = doGetProcAddress(libadvapi32, "RegCopyTreeW")
	regCreateKeyEx = doGetProcAddress(libadvapi32, "RegCreateKeyExW")
	regCreateKeyTransacted = doGetProcAddress(libadvapi32, "RegCreateKeyTransactedW")
	regCreateKey = doGetProcAddress(libadvapi32, "RegCreateKeyW")
	regDeleteKeyEx = doGetProcAddress(libadvapi32, "RegDeleteKeyExW")
	regDeleteKeyTransacted = doGetProcAddress(libadvapi32, "RegDeleteKeyTransactedW")
	regDeleteKeyValue = doGetProcAddress(libadvapi32, "RegDeleteKeyValueW")
	regDeleteKey = doGetProcAddress(libadvapi32, "RegDeleteKeyW")
	regDeleteTree = doGetProcAddress(libadvapi32, "RegDeleteTreeW")
	regDeleteValue = doGetProcAddress(libadvapi32, "RegDeleteValueW")
	regDisablePredefinedCache = doGetProcAddress(libadvapi32, "RegDisablePredefinedCache")
	regDisablePredefinedCacheEx = doGetProcAddress(libadvapi32, "RegDisablePredefinedCacheEx")
	regDisableReflectionKey = doGetProcAddress(libadvapi32, "RegDisableReflectionKey")
	regEnableReflectionKey = doGetProcAddress(libadvapi32, "RegEnableReflectionKey")
	regEnumKey = doGetProcAddress(libadvapi32, "RegEnumKeyW")
	regEnumValue = doGetProcAddress(libadvapi32, "RegEnumValueW")
	regFlushKey = doGetProcAddress(libadvapi32, "RegFlushKey")
	regGetKeySecurity = doGetProcAddress(libadvapi32, "RegGetKeySecurity")
	regGetValue = doGetProcAddress(libadvapi32, "RegGetValueW")
	regLoadAppKey = doGetProcAddress(libadvapi32, "RegLoadAppKeyW")
	regLoadKey = doGetProcAddress(libadvapi32, "RegLoadKeyW")
	regLoadMUIString = doGetProcAddress(libadvapi32, "RegLoadMUIStringW")
	regNotifyChangeKeyValue = doGetProcAddress(libadvapi32, "RegNotifyChangeKeyValue")
	regOpenCurrentUser = doGetProcAddress(libadvapi32, "RegOpenCurrentUser")
	regOpenKeyEx = doGetProcAddress(libadvapi32, "RegOpenKeyExW")
	regOpenKeyTransacted = doGetProcAddress(libadvapi32, "RegOpenKeyTransactedW")
	regOpenKey = doGetProcAddress(libadvapi32, "RegOpenKeyW")
	regOpenUserClassesRoot = doGetProcAddress(libadvapi32, "RegOpenUserClassesRoot")
	regOverridePredefKey = doGetProcAddress(libadvapi32, "RegOverridePredefKey")
	regQueryReflectionKey = doGetProcAddress(libadvapi32, "RegQueryReflectionKey")
	regQueryValueEx = doGetProcAddress(libadvapi32, "RegQueryValueExW")
	regQueryValue = doGetProcAddress(libadvapi32, "RegQueryValueW")
	regReplaceKey = doGetProcAddress(libadvapi32, "RegReplaceKeyW")
	regRestoreKey = doGetProcAddress(libadvapi32, "RegRestoreKeyW")
	regSaveKeyEx = doGetProcAddress(libadvapi32, "RegSaveKeyExW")
	regSaveKey = doGetProcAddress(libadvapi32, "RegSaveKeyW")
	regSetKeySecurity = doGetProcAddress(libadvapi32, "RegSetKeySecurity")
	regSetKeyValue = doGetProcAddress(libadvapi32, "RegSetKeyValueW")
	regSetValueEx = doGetProcAddress(libadvapi32, "RegSetValueExW")
	regSetValue = doGetProcAddress(libadvapi32, "RegSetValueW")
	regUnLoadKey = doGetProcAddress(libadvapi32, "RegUnLoadKeyW")
	registerEventSource = doGetProcAddress(libadvapi32, "RegisterEventSourceW")
	registerServiceCtrlHandlerEx = doGetProcAddress(libadvapi32, "RegisterServiceCtrlHandlerExW")
	reportEvent = doGetProcAddress(libadvapi32, "ReportEventW")
	revertToSelf = doGetProcAddress(libadvapi32, "RevertToSelf")
	saferCloseLevel = doGetProcAddress(libadvapi32, "SaferCloseLevel")
	saferComputeTokenFromLevel = doGetProcAddress(libadvapi32, "SaferComputeTokenFromLevel")
	saferCreateLevel = doGetProcAddress(libadvapi32, "SaferCreateLevel")
	saferRecordEventLogEntry = doGetProcAddress(libadvapi32, "SaferRecordEventLogEntry")
	saferiIsExecutableFileType = doGetProcAddress(libadvapi32, "SaferiIsExecutableFileType")
	setFileSecurity = doGetProcAddress(libadvapi32, "SetFileSecurityW")
	setKernelObjectSecurity = doGetProcAddress(libadvapi32, "SetKernelObjectSecurity")
	setNamedSecurityInfo = doGetProcAddress(libadvapi32, "SetNamedSecurityInfoW")
	setPrivateObjectSecurity = doGetProcAddress(libadvapi32, "SetPrivateObjectSecurity")
	setPrivateObjectSecurityEx = doGetProcAddress(libadvapi32, "SetPrivateObjectSecurityEx")
	setSecurityAccessMask = doGetProcAddress(libadvapi32, "SetSecurityAccessMask")
	setSecurityDescriptorControl = doGetProcAddress(libadvapi32, "SetSecurityDescriptorControl")
	setSecurityDescriptorDacl = doGetProcAddress(libadvapi32, "SetSecurityDescriptorDacl")
	setSecurityDescriptorGroup = doGetProcAddress(libadvapi32, "SetSecurityDescriptorGroup")
	setSecurityDescriptorOwner = doGetProcAddress(libadvapi32, "SetSecurityDescriptorOwner")
	setSecurityDescriptorSacl = doGetProcAddress(libadvapi32, "SetSecurityDescriptorSacl")
	setSecurityInfo = doGetProcAddress(libadvapi32, "SetSecurityInfo")
	setServiceBits = doGetProcAddress(libadvapi32, "SetServiceBits")
	setServiceObjectSecurity = doGetProcAddress(libadvapi32, "SetServiceObjectSecurity")
	setServiceStatus = doGetProcAddress(libadvapi32, "SetServiceStatus")
	setThreadToken = doGetProcAddress(libadvapi32, "SetThreadToken")
	setUserFileEncryptionKey = doGetProcAddress(libadvapi32, "SetUserFileEncryptionKey")
	startService = doGetProcAddress(libadvapi32, "StartServiceW")
	uninstallApplication = doGetProcAddress(libadvapi32, "UninstallApplication")
	unlockServiceDatabase = doGetProcAddress(libadvapi32, "UnlockServiceDatabase")
	wow64Win32ApiEntry = doGetProcAddress(libadvapi32, "Wow64Win32ApiEntry")
	eventActivityIdControl = doGetProcAddress(libadvapi32, "EventActivityIdControl")
	systemFunction030 = doGetProcAddress(libadvapi32, "SystemFunction030")
	systemFunction035 = doGetProcAddress(libadvapi32, "SystemFunction035")
	systemFunction036 = doGetProcAddress(libadvapi32, "SystemFunction036")
}

func AbortSystemShutdown(lpMachineName LPWSTR) bool {
	ret1 := syscall3(abortSystemShutdown, 1,
		uintptr(unsafe.Pointer(lpMachineName)),
		0,
		0)
	return ret1 != 0
}

func AccessCheck(pSecurityDescriptor PSECURITY_DESCRIPTOR, clientToken HANDLE, desiredAccess uint32, genericMapping *GENERIC_MAPPING, privilegeSet *PRIVILEGE_SET, privilegeSetLength *uint32, grantedAccess *uint32, accessStatus *BOOL) bool {
	ret1 := syscall9(accessCheck, 8,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(clientToken),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(genericMapping)),
		uintptr(unsafe.Pointer(privilegeSet)),
		uintptr(unsafe.Pointer(privilegeSetLength)),
		uintptr(unsafe.Pointer(grantedAccess)),
		uintptr(unsafe.Pointer(accessStatus)),
		0)
	return ret1 != 0
}

func AccessCheckAndAuditAlarm(subsystemName string, handleId LPVOID, objectTypeName LPWSTR, objectName LPWSTR, securityDescriptor PSECURITY_DESCRIPTOR, desiredAccess uint32, genericMapping *GENERIC_MAPPING, objectCreation bool, grantedAccess *uint32, accessStatus *BOOL, pfGenerateOnClose *BOOL) bool {
	subsystemNameStr := unicode16FromString(subsystemName)
	ret1 := syscall12(accessCheckAndAuditAlarm, 11,
		uintptr(unsafe.Pointer(&subsystemNameStr[0])),
		uintptr(unsafe.Pointer(handleId)),
		uintptr(unsafe.Pointer(objectTypeName)),
		uintptr(unsafe.Pointer(objectName)),
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(genericMapping)),
		getUintptrFromBool(objectCreation),
		uintptr(unsafe.Pointer(grantedAccess)),
		uintptr(unsafe.Pointer(accessStatus)),
		uintptr(unsafe.Pointer(pfGenerateOnClose)),
		0)
	return ret1 != 0
}

func AccessCheckByType(pSecurityDescriptor PSECURITY_DESCRIPTOR, principalSelfSid PSID, clientToken HANDLE, desiredAccess uint32, objectTypeList *OBJECT_TYPE_LIST, objectTypeListLength uint32, genericMapping *GENERIC_MAPPING, privilegeSet *PRIVILEGE_SET, privilegeSetLength *uint32, grantedAccess *uint32, accessStatus *BOOL) bool {
	ret1 := syscall12(accessCheckByType, 11,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(principalSelfSid),
		uintptr(clientToken),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(objectTypeList)),
		uintptr(objectTypeListLength),
		uintptr(unsafe.Pointer(genericMapping)),
		uintptr(unsafe.Pointer(privilegeSet)),
		uintptr(unsafe.Pointer(privilegeSetLength)),
		uintptr(unsafe.Pointer(grantedAccess)),
		uintptr(unsafe.Pointer(accessStatus)),
		0)
	return ret1 != 0
}

// TODO: Too many syscall arguments: 18
// func AccessCheckByTypeAndAuditAlarm(subsystemName string, handleId LPVOID, objectTypeName string, objectName string, securityDescriptor PSECURITY_DESCRIPTOR, principalSelfSid PSID, desiredAccess uint32, auditType AUDIT_EVENT_TYPE, flags uint32, objectTypeList *OBJECT_TYPE_LIST, objectTypeListLength uint32, genericMapping *GENERIC_MAPPING, objectCreation bool, grantedAccess *uint32, accessStatus *BOOL, pfGenerateOnClose *BOOL) bool

func AccessCheckByTypeResultList(pSecurityDescriptor PSECURITY_DESCRIPTOR, principalSelfSid PSID, clientToken HANDLE, desiredAccess uint32, objectTypeList *OBJECT_TYPE_LIST, objectTypeListLength uint32, genericMapping *GENERIC_MAPPING, privilegeSet *PRIVILEGE_SET, privilegeSetLength *uint32, grantedAccessList *uint32, accessStatusList *uint32) bool {
	ret1 := syscall12(accessCheckByTypeResultList, 11,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(principalSelfSid),
		uintptr(clientToken),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(objectTypeList)),
		uintptr(objectTypeListLength),
		uintptr(unsafe.Pointer(genericMapping)),
		uintptr(unsafe.Pointer(privilegeSet)),
		uintptr(unsafe.Pointer(privilegeSetLength)),
		uintptr(unsafe.Pointer(grantedAccessList)),
		uintptr(unsafe.Pointer(accessStatusList)),
		0)
	return ret1 != 0
}

// TODO: Too many syscall arguments: 18
// func AccessCheckByTypeResultListAndAuditAlarmByHandle(subsystemName string, handleId LPVOID, clientToken HANDLE, objectTypeName string, objectName string, securityDescriptor PSECURITY_DESCRIPTOR, principalSelfSid PSID, desiredAccess uint32, auditType AUDIT_EVENT_TYPE, flags uint32, objectTypeList *OBJECT_TYPE_LIST, objectTypeListLength uint32, genericMapping *GENERIC_MAPPING, objectCreation bool, grantedAccessList *uint32, accessStatusList *uint32, pfGenerateOnClose *BOOL) bool

// TODO: Too many syscall arguments: 18
// func AccessCheckByTypeResultListAndAuditAlarm(subsystemName string, handleId LPVOID, objectTypeName string, objectName string, securityDescriptor PSECURITY_DESCRIPTOR, principalSelfSid PSID, desiredAccess uint32, auditType AUDIT_EVENT_TYPE, flags uint32, objectTypeList *OBJECT_TYPE_LIST, objectTypeListLength uint32, genericMapping *GENERIC_MAPPING, objectCreation bool, grantedAccessList *uint32, accessStatusList *uint32, pfGenerateOnClose *BOOL) bool

func AddAccessAllowedAce(pAcl *ACL, dwAceRevision uint32, accessMask uint32, pSid PSID) bool {
	ret1 := syscall6(addAccessAllowedAce, 4,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(accessMask),
		uintptr(pSid),
		0,
		0)
	return ret1 != 0
}

func AddAccessAllowedAceEx(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, accessMask uint32, pSid PSID) bool {
	ret1 := syscall6(addAccessAllowedAceEx, 5,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(accessMask),
		uintptr(pSid),
		0)
	return ret1 != 0
}

func AddAccessAllowedObjectAce(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, accessMask uint32, objectTypeGuid *GUID, inheritedObjectTypeGuid *GUID, pSid PSID) bool {
	ret1 := syscall9(addAccessAllowedObjectAce, 7,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(accessMask),
		uintptr(unsafe.Pointer(objectTypeGuid)),
		uintptr(unsafe.Pointer(inheritedObjectTypeGuid)),
		uintptr(pSid),
		0,
		0)
	return ret1 != 0
}

func AddAccessDeniedAce(pAcl *ACL, dwAceRevision uint32, accessMask uint32, pSid PSID) bool {
	ret1 := syscall6(addAccessDeniedAce, 4,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(accessMask),
		uintptr(pSid),
		0,
		0)
	return ret1 != 0
}

func AddAccessDeniedAceEx(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, accessMask uint32, pSid PSID) bool {
	ret1 := syscall6(addAccessDeniedAceEx, 5,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(accessMask),
		uintptr(pSid),
		0)
	return ret1 != 0
}

func AddAccessDeniedObjectAce(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, accessMask uint32, objectTypeGuid *GUID, inheritedObjectTypeGuid *GUID, pSid PSID) bool {
	ret1 := syscall9(addAccessDeniedObjectAce, 7,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(accessMask),
		uintptr(unsafe.Pointer(objectTypeGuid)),
		uintptr(unsafe.Pointer(inheritedObjectTypeGuid)),
		uintptr(pSid),
		0,
		0)
	return ret1 != 0
}

func AddAce(pAcl *ACL, dwAceRevision uint32, dwStartingAceIndex uint32, pAceList LPVOID, nAceListLength uint32) bool {
	ret1 := syscall6(addAce, 5,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(dwStartingAceIndex),
		uintptr(unsafe.Pointer(pAceList)),
		uintptr(nAceListLength),
		0)
	return ret1 != 0
}

func AddAuditAccessAce(pAcl *ACL, dwAceRevision uint32, dwAccessMask uint32, pSid PSID, bAuditSuccess bool, bAuditFailure bool) bool {
	ret1 := syscall6(addAuditAccessAce, 6,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(dwAccessMask),
		uintptr(pSid),
		getUintptrFromBool(bAuditSuccess),
		getUintptrFromBool(bAuditFailure))
	return ret1 != 0
}

func AddAuditAccessAceEx(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, dwAccessMask uint32, pSid PSID, bAuditSuccess bool, bAuditFailure bool) bool {
	ret1 := syscall9(addAuditAccessAceEx, 7,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(dwAccessMask),
		uintptr(pSid),
		getUintptrFromBool(bAuditSuccess),
		getUintptrFromBool(bAuditFailure),
		0,
		0)
	return ret1 != 0
}

func AddAuditAccessObjectAce(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, accessMask uint32, objectTypeGuid *GUID, inheritedObjectTypeGuid *GUID, pSid PSID, bAuditSuccess bool, bAuditFailure bool) bool {
	ret1 := syscall9(addAuditAccessObjectAce, 9,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(accessMask),
		uintptr(unsafe.Pointer(objectTypeGuid)),
		uintptr(unsafe.Pointer(inheritedObjectTypeGuid)),
		uintptr(pSid),
		getUintptrFromBool(bAuditSuccess),
		getUintptrFromBool(bAuditFailure))
	return ret1 != 0
}

func AddConditionalAce(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, aceType UCHAR, accessMask uint32, pSid PSID, conditionStr PWCHAR, returnLength *uint32) bool {
	ret1 := syscall9(addConditionalAce, 8,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(aceType),
		uintptr(accessMask),
		uintptr(pSid),
		uintptr(unsafe.Pointer(conditionStr)),
		uintptr(unsafe.Pointer(returnLength)),
		0)
	return ret1 != 0
}

func AddMandatoryAce(pAcl *ACL, dwAceRevision uint32, aceFlags uint32, mandatoryPolicy uint32, pLabelSid PSID) bool {
	ret1 := syscall6(addMandatoryAce, 5,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceRevision),
		uintptr(aceFlags),
		uintptr(mandatoryPolicy),
		uintptr(pLabelSid),
		0)
	return ret1 != 0
}

func AddUsersToEncryptedFile(lpFileName string, pUsers *ENCRYPTION_CERTIFICATE_LIST) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(addUsersToEncryptedFile, 2,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(pUsers)),
		0)
	return uint32(ret1)
}

func AdjustTokenGroups(tokenHandle HANDLE, resetToDefault bool, newState *TOKEN_GROUPS, bufferLength uint32, previousState *TOKEN_GROUPS, returnLength *DWORD) bool {
	ret1 := syscall6(adjustTokenGroups, 6,
		uintptr(tokenHandle),
		getUintptrFromBool(resetToDefault),
		uintptr(unsafe.Pointer(newState)),
		uintptr(bufferLength),
		uintptr(unsafe.Pointer(previousState)),
		uintptr(unsafe.Pointer(returnLength)))
	return ret1 != 0
}

func AdjustTokenPrivileges(tokenHandle HANDLE, disableAllPrivileges bool, newState *TOKEN_PRIVILEGES, bufferLength uint32, previousState *TOKEN_PRIVILEGES, returnLength *DWORD) bool {
	ret1 := syscall6(adjustTokenPrivileges, 6,
		uintptr(tokenHandle),
		getUintptrFromBool(disableAllPrivileges),
		uintptr(unsafe.Pointer(newState)),
		uintptr(bufferLength),
		uintptr(unsafe.Pointer(previousState)),
		uintptr(unsafe.Pointer(returnLength)))
	return ret1 != 0
}

func AllocateAndInitializeSid(pIdentifierAuthority *SID_IDENTIFIER_AUTHORITY, nSubAuthorityCount byte, nSubAuthority0 uint32, nSubAuthority1 uint32, nSubAuthority2 uint32, nSubAuthority3 uint32, nSubAuthority4 uint32, nSubAuthority5 uint32, nSubAuthority6 uint32, nSubAuthority7 uint32, pSid *PSID) bool {
	ret1 := syscall12(allocateAndInitializeSid, 11,
		uintptr(unsafe.Pointer(pIdentifierAuthority)),
		uintptr(nSubAuthorityCount),
		uintptr(nSubAuthority0),
		uintptr(nSubAuthority1),
		uintptr(nSubAuthority2),
		uintptr(nSubAuthority3),
		uintptr(nSubAuthority4),
		uintptr(nSubAuthority5),
		uintptr(nSubAuthority6),
		uintptr(nSubAuthority7),
		uintptr(unsafe.Pointer(pSid)),
		0)
	return ret1 != 0
}

func AllocateLocallyUniqueId(luid *LUID) bool {
	ret1 := syscall3(allocateLocallyUniqueId, 1,
		uintptr(unsafe.Pointer(luid)),
		0,
		0)
	return ret1 != 0
}

func AreAllAccessesGranted(grantedAccess uint32, desiredAccess uint32) bool {
	ret1 := syscall3(areAllAccessesGranted, 2,
		uintptr(grantedAccess),
		uintptr(desiredAccess),
		0)
	return ret1 != 0
}

func AreAnyAccessesGranted(grantedAccess uint32, desiredAccess uint32) bool {
	ret1 := syscall3(areAnyAccessesGranted, 2,
		uintptr(grantedAccess),
		uintptr(desiredAccess),
		0)
	return ret1 != 0
}

func AuditComputeEffectivePolicyBySid(pSid /*const*/ PSID, pSubCategoryGuids /*const*/ *GUID, policyCount ULONG, ppAuditPolicy *PAUDIT_POLICY_INFORMATION) BOOLEAN {
	ret1 := syscall6(auditComputeEffectivePolicyBySid, 4,
		uintptr(pSid),
		uintptr(unsafe.Pointer(pSubCategoryGuids)),
		uintptr(policyCount),
		uintptr(unsafe.Pointer(ppAuditPolicy)),
		0,
		0)
	return BOOLEAN(ret1)
}

func AuditComputeEffectivePolicyByToken(hTokenHandle HANDLE, pSubCategoryGuids /*const*/ *GUID, policyCount ULONG, ppAuditPolicy *PAUDIT_POLICY_INFORMATION) BOOLEAN {
	ret1 := syscall6(auditComputeEffectivePolicyByToken, 4,
		uintptr(hTokenHandle),
		uintptr(unsafe.Pointer(pSubCategoryGuids)),
		uintptr(policyCount),
		uintptr(unsafe.Pointer(ppAuditPolicy)),
		0,
		0)
	return BOOLEAN(ret1)
}

func AuditEnumerateCategories(ppAuditCategoriesArray uintptr, pCountReturned *uint32) BOOLEAN {
	ret1 := syscall3(auditEnumerateCategories, 2,
		ppAuditCategoriesArray,
		uintptr(unsafe.Pointer(pCountReturned)),
		0)
	return BOOLEAN(ret1)
}

func AuditEnumeratePerUserPolicy(ppAuditSidArray *PPOLICY_AUDIT_SID_ARRAY) BOOLEAN {
	ret1 := syscall3(auditEnumeratePerUserPolicy, 1,
		uintptr(unsafe.Pointer(ppAuditSidArray)),
		0,
		0)
	return BOOLEAN(ret1)
}

func AuditEnumerateSubCategories(pAuditCategoryGuid /*const*/ *GUID, bRetrieveAllSubCategories BOOLEAN, ppAuditSubCategoriesArray uintptr, pCountReturned *uint32) BOOLEAN {
	ret1 := syscall6(auditEnumerateSubCategories, 4,
		uintptr(unsafe.Pointer(pAuditCategoryGuid)),
		uintptr(bRetrieveAllSubCategories),
		ppAuditSubCategoriesArray,
		uintptr(unsafe.Pointer(pCountReturned)),
		0,
		0)
	return BOOLEAN(ret1)
}

func AuditFree(buffer uintptr) {
	syscall3(auditFree, 1,
		buffer,
		0,
		0)
}

func AuditLookupCategoryGuidFromCategoryId(auditCategoryId POLICY_AUDIT_EVENT_TYPE, pAuditCategoryGuid *GUID) BOOLEAN {
	ret1 := syscall3(auditLookupCategoryGuidFromCategoryId, 2,
		uintptr(auditCategoryId),
		uintptr(unsafe.Pointer(pAuditCategoryGuid)),
		0)
	return BOOLEAN(ret1)
}

func AuditLookupCategoryIdFromCategoryGuid(pAuditCategoryGuid /*const*/ *GUID, pAuditCategoryId PPOLICY_AUDIT_EVENT_TYPE) BOOLEAN {
	ret1 := syscall3(auditLookupCategoryIdFromCategoryGuid, 2,
		uintptr(unsafe.Pointer(pAuditCategoryGuid)),
		uintptr(unsafe.Pointer(pAuditCategoryId)),
		0)
	return BOOLEAN(ret1)
}

func AuditLookupCategoryName(pAuditCategoryGuid /*const*/ *GUID, ppszCategoryName *LPWSTR) BOOLEAN {
	ret1 := syscall3(auditLookupCategoryName, 2,
		uintptr(unsafe.Pointer(pAuditCategoryGuid)),
		uintptr(unsafe.Pointer(ppszCategoryName)),
		0)
	return BOOLEAN(ret1)
}

func AuditLookupSubCategoryName(pAuditSubCategoryGuid /*const*/ *GUID, ppszSubCategoryName *LPWSTR) BOOLEAN {
	ret1 := syscall3(auditLookupSubCategoryName, 2,
		uintptr(unsafe.Pointer(pAuditSubCategoryGuid)),
		uintptr(unsafe.Pointer(ppszSubCategoryName)),
		0)
	return BOOLEAN(ret1)
}

func AuditQueryPerUserPolicy(pSid /*const*/ PSID, pSubCategoryGuids /*const*/ *GUID, policyCount ULONG, ppAuditPolicy *PAUDIT_POLICY_INFORMATION) BOOLEAN {
	ret1 := syscall6(auditQueryPerUserPolicy, 4,
		uintptr(pSid),
		uintptr(unsafe.Pointer(pSubCategoryGuids)),
		uintptr(policyCount),
		uintptr(unsafe.Pointer(ppAuditPolicy)),
		0,
		0)
	return BOOLEAN(ret1)
}

func AuditQuerySecurity(securityInformation SECURITY_INFORMATION, ppSecurityDescriptor *PSECURITY_DESCRIPTOR) BOOLEAN {
	ret1 := syscall3(auditQuerySecurity, 2,
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(ppSecurityDescriptor)),
		0)
	return BOOLEAN(ret1)
}

func AuditQuerySystemPolicy(pSubCategoryGuids /*const*/ *GUID, policyCount ULONG, ppAuditPolicy *PAUDIT_POLICY_INFORMATION) BOOLEAN {
	ret1 := syscall3(auditQuerySystemPolicy, 3,
		uintptr(unsafe.Pointer(pSubCategoryGuids)),
		uintptr(policyCount),
		uintptr(unsafe.Pointer(ppAuditPolicy)))
	return BOOLEAN(ret1)
}

func AuditSetPerUserPolicy(pSid /*const*/ PSID, pAuditPolicy /*const*/ PAUDIT_POLICY_INFORMATION, policyCount ULONG) BOOLEAN {
	ret1 := syscall3(auditSetPerUserPolicy, 3,
		uintptr(pSid),
		uintptr(unsafe.Pointer(pAuditPolicy)),
		uintptr(policyCount))
	return BOOLEAN(ret1)
}

func AuditSetSecurity(securityInformation SECURITY_INFORMATION, pSecurityDescriptor PSECURITY_DESCRIPTOR) BOOLEAN {
	ret1 := syscall3(auditSetSecurity, 2,
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		0)
	return BOOLEAN(ret1)
}

func AuditSetSystemPolicy(pAuditPolicy /*const*/ PAUDIT_POLICY_INFORMATION, policyCount ULONG) BOOLEAN {
	ret1 := syscall3(auditSetSystemPolicy, 2,
		uintptr(unsafe.Pointer(pAuditPolicy)),
		uintptr(policyCount),
		0)
	return BOOLEAN(ret1)
}

func BackupEventLog(hEventLog HANDLE, lpBackupFileName string) bool {
	lpBackupFileNameStr := unicode16FromString(lpBackupFileName)
	ret1 := syscall3(backupEventLog, 2,
		uintptr(hEventLog),
		uintptr(unsafe.Pointer(&lpBackupFileNameStr[0])),
		0)
	return ret1 != 0
}

func BuildExplicitAccessWithName(pExplicitAccess *EXPLICIT_ACCESS, pTrusteeName LPWSTR, accessPermissions uint32, accessMode ACCESS_MODE, inheritance uint32) {
	syscall6(buildExplicitAccessWithName, 5,
		uintptr(unsafe.Pointer(pExplicitAccess)),
		uintptr(unsafe.Pointer(pTrusteeName)),
		uintptr(accessPermissions),
		uintptr(accessMode),
		uintptr(inheritance),
		0)
}

func BuildImpersonateExplicitAccessWithName(pExplicitAccess *EXPLICIT_ACCESS, pTrusteeName LPWSTR, pTrustee *TRUSTEE, accessPermissions uint32, accessMode ACCESS_MODE, inheritance uint32) {
	syscall6(buildImpersonateExplicitAccessWithName, 6,
		uintptr(unsafe.Pointer(pExplicitAccess)),
		uintptr(unsafe.Pointer(pTrusteeName)),
		uintptr(unsafe.Pointer(pTrustee)),
		uintptr(accessPermissions),
		uintptr(accessMode),
		uintptr(inheritance))
}

func BuildImpersonateTrustee(pTrustee *TRUSTEE, pImpersonateTrustee *TRUSTEE) {
	syscall3(buildImpersonateTrustee, 2,
		uintptr(unsafe.Pointer(pTrustee)),
		uintptr(unsafe.Pointer(pImpersonateTrustee)),
		0)
}

func BuildSecurityDescriptor(pOwner *TRUSTEE, pGroup *TRUSTEE, cCountOfAccessEntries ULONG, pListOfAccessEntries *EXPLICIT_ACCESS, cCountOfAuditEntries ULONG, pListOfAuditEntries *EXPLICIT_ACCESS, pOldSD PSECURITY_DESCRIPTOR, pSizeNewSD *uint32, pNewSD *PSECURITY_DESCRIPTOR) uint32 {
	ret1 := syscall9(buildSecurityDescriptor, 9,
		uintptr(unsafe.Pointer(pOwner)),
		uintptr(unsafe.Pointer(pGroup)),
		uintptr(cCountOfAccessEntries),
		uintptr(unsafe.Pointer(pListOfAccessEntries)),
		uintptr(cCountOfAuditEntries),
		uintptr(unsafe.Pointer(pListOfAuditEntries)),
		uintptr(unsafe.Pointer(pOldSD)),
		uintptr(unsafe.Pointer(pSizeNewSD)),
		uintptr(unsafe.Pointer(pNewSD)))
	return uint32(ret1)
}

func BuildTrusteeWithName(pTrustee *TRUSTEE, pName LPWSTR) {
	syscall3(buildTrusteeWithName, 2,
		uintptr(unsafe.Pointer(pTrustee)),
		uintptr(unsafe.Pointer(pName)),
		0)
}

func BuildTrusteeWithObjectsAndName(pTrustee *TRUSTEE, pObjName *OBJECTS_AND_NAME, objectType SE_OBJECT_TYPE, objectTypeName LPWSTR, inheritedObjectTypeName LPWSTR, name LPWSTR) {
	syscall6(buildTrusteeWithObjectsAndName, 6,
		uintptr(unsafe.Pointer(pTrustee)),
		uintptr(unsafe.Pointer(pObjName)),
		uintptr(objectType),
		uintptr(unsafe.Pointer(objectTypeName)),
		uintptr(unsafe.Pointer(inheritedObjectTypeName)),
		uintptr(unsafe.Pointer(name)))
}

func BuildTrusteeWithObjectsAndSid(pTrustee *TRUSTEE, pObjSid *OBJECTS_AND_SID, pObjectGuid *GUID, pInheritedObjectGuid *GUID, pSid PSID) {
	syscall6(buildTrusteeWithObjectsAndSid, 5,
		uintptr(unsafe.Pointer(pTrustee)),
		uintptr(unsafe.Pointer(pObjSid)),
		uintptr(unsafe.Pointer(pObjectGuid)),
		uintptr(unsafe.Pointer(pInheritedObjectGuid)),
		uintptr(pSid),
		0)
}

func BuildTrusteeWithSid(pTrustee *TRUSTEE, pSid PSID) {
	syscall3(buildTrusteeWithSid, 2,
		uintptr(unsafe.Pointer(pTrustee)),
		uintptr(pSid),
		0)
}

func ChangeServiceConfig2(hService SC_HANDLE, dwInfoLevel uint32, lpInfo LPVOID) bool {
	ret1 := syscall3(changeServiceConfig2, 3,
		uintptr(hService),
		uintptr(dwInfoLevel),
		uintptr(unsafe.Pointer(lpInfo)))
	return ret1 != 0
}

func ChangeServiceConfig(hService SC_HANDLE, dwServiceType uint32, dwStartType uint32, dwErrorControl uint32, lpBinaryPathName string, lpLoadOrderGroup string, lpdwTagId *uint32, lpDependencies string, lpServiceStartName string, lpPassword string, lpDisplayName string) bool {
	lpBinaryPathNameStr := unicode16FromString(lpBinaryPathName)
	lpLoadOrderGroupStr := unicode16FromString(lpLoadOrderGroup)
	lpDependenciesStr := unicode16FromString(lpDependencies)
	lpServiceStartNameStr := unicode16FromString(lpServiceStartName)
	lpPasswordStr := unicode16FromString(lpPassword)
	lpDisplayNameStr := unicode16FromString(lpDisplayName)
	ret1 := syscall12(changeServiceConfig, 11,
		uintptr(hService),
		uintptr(dwServiceType),
		uintptr(dwStartType),
		uintptr(dwErrorControl),
		uintptr(unsafe.Pointer(&lpBinaryPathNameStr[0])),
		uintptr(unsafe.Pointer(&lpLoadOrderGroupStr[0])),
		uintptr(unsafe.Pointer(lpdwTagId)),
		uintptr(unsafe.Pointer(&lpDependenciesStr[0])),
		uintptr(unsafe.Pointer(&lpServiceStartNameStr[0])),
		uintptr(unsafe.Pointer(&lpPasswordStr[0])),
		uintptr(unsafe.Pointer(&lpDisplayNameStr[0])),
		0)
	return ret1 != 0
}

func CheckTokenMembership(tokenHandle HANDLE, sidToCheck PSID, isMember *BOOL) bool {
	ret1 := syscall3(checkTokenMembership, 3,
		uintptr(tokenHandle),
		uintptr(sidToCheck),
		uintptr(unsafe.Pointer(isMember)))
	return ret1 != 0
}

func ClearEventLog(hEventLog HANDLE, lpBackupFileName string) bool {
	lpBackupFileNameStr := unicode16FromString(lpBackupFileName)
	ret1 := syscall3(clearEventLog, 2,
		uintptr(hEventLog),
		uintptr(unsafe.Pointer(&lpBackupFileNameStr[0])),
		0)
	return ret1 != 0
}

func CloseEncryptedFileRaw(pvContext uintptr) {
	syscall3(closeEncryptedFileRaw, 1,
		pvContext,
		0,
		0)
}

func CloseEventLog(hEventLog HANDLE) bool {
	ret1 := syscall3(closeEventLog, 1,
		uintptr(hEventLog),
		0,
		0)
	return ret1 != 0
}

func CloseServiceHandle(hSCObject SC_HANDLE) bool {
	ret1 := syscall3(closeServiceHandle, 1,
		uintptr(hSCObject),
		0,
		0)
	return ret1 != 0
}

func CloseThreadWaitChainSession(wctHandle HWCT) {
	syscall3(closeThreadWaitChainSession, 1,
		uintptr(wctHandle),
		0,
		0)
}

func CommandLineFromMsiDescriptor(descriptor *WCHAR, commandLine *WCHAR, commandLineLength *uint32) uint32 {
	ret1 := syscall3(commandLineFromMsiDescriptor, 3,
		uintptr(unsafe.Pointer(descriptor)),
		uintptr(unsafe.Pointer(commandLine)),
		uintptr(unsafe.Pointer(commandLineLength)))
	return uint32(ret1)
}

func ControlService(hService SC_HANDLE, dwControl uint32, lpServiceStatus *SERVICE_STATUS) bool {
	ret1 := syscall3(controlService, 3,
		uintptr(hService),
		uintptr(dwControl),
		uintptr(unsafe.Pointer(lpServiceStatus)))
	return ret1 != 0
}

func ControlServiceEx(hService SC_HANDLE, dwControl uint32, dwInfoLevel uint32, pControlParams uintptr) bool {
	ret1 := syscall6(controlServiceEx, 4,
		uintptr(hService),
		uintptr(dwControl),
		uintptr(dwInfoLevel),
		pControlParams,
		0,
		0)
	return ret1 != 0
}

func ConvertSecurityDescriptorToStringSecurityDescriptor(securityDescriptor PSECURITY_DESCRIPTOR, requestedStringSDRevision uint32, securityInformation SECURITY_INFORMATION, stringSecurityDescriptor *LPWSTR, stringSecurityDescriptorLen *uint32) bool {
	ret1 := syscall6(convertSecurityDescriptorToStringSecurityDescriptor, 5,
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(requestedStringSDRevision),
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(stringSecurityDescriptor)),
		uintptr(unsafe.Pointer(stringSecurityDescriptorLen)),
		0)
	return ret1 != 0
}

func ConvertSidToStringSid(sid PSID, stringSid *LPWSTR) bool {
	ret1 := syscall3(convertSidToStringSid, 2,
		uintptr(sid),
		uintptr(unsafe.Pointer(stringSid)),
		0)
	return ret1 != 0
}

func ConvertStringSecurityDescriptorToSecurityDescriptor(stringSecurityDescriptor string, stringSDRevision uint32, securityDescriptor *PSECURITY_DESCRIPTOR, securityDescriptorSize *uint32) bool {
	stringSecurityDescriptorStr := unicode16FromString(stringSecurityDescriptor)
	ret1 := syscall6(convertStringSecurityDescriptorToSecurityDescriptor, 4,
		uintptr(unsafe.Pointer(&stringSecurityDescriptorStr[0])),
		uintptr(stringSDRevision),
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(unsafe.Pointer(securityDescriptorSize)),
		0,
		0)
	return ret1 != 0
}

func ConvertStringSidToSid(stringSid string, sid *PSID) bool {
	stringSidStr := unicode16FromString(stringSid)
	ret1 := syscall3(convertStringSidToSid, 2,
		uintptr(unsafe.Pointer(&stringSidStr[0])),
		uintptr(unsafe.Pointer(sid)),
		0)
	return ret1 != 0
}

func ConvertToAutoInheritPrivateObjectSecurity(parentDescriptor PSECURITY_DESCRIPTOR, currentSecurityDescriptor PSECURITY_DESCRIPTOR, newSecurityDescriptor *PSECURITY_DESCRIPTOR, objectType *GUID, isDirectoryObject BOOLEAN, genericMapping *GENERIC_MAPPING) bool {
	ret1 := syscall6(convertToAutoInheritPrivateObjectSecurity, 6,
		uintptr(unsafe.Pointer(parentDescriptor)),
		uintptr(unsafe.Pointer(currentSecurityDescriptor)),
		uintptr(unsafe.Pointer(newSecurityDescriptor)),
		uintptr(unsafe.Pointer(objectType)),
		uintptr(isDirectoryObject),
		uintptr(unsafe.Pointer(genericMapping)))
	return ret1 != 0
}

func CopySid(nDestinationSidLength uint32, pDestinationSid PSID, pSourceSid PSID) bool {
	ret1 := syscall3(copySid, 3,
		uintptr(nDestinationSidLength),
		uintptr(pDestinationSid),
		uintptr(pSourceSid))
	return ret1 != 0
}

func CreatePrivateObjectSecurity(parentDescriptor PSECURITY_DESCRIPTOR, creatorDescriptor PSECURITY_DESCRIPTOR, newDescriptor *PSECURITY_DESCRIPTOR, isDirectoryObject bool, token HANDLE, genericMapping *GENERIC_MAPPING) bool {
	ret1 := syscall6(createPrivateObjectSecurity, 6,
		uintptr(unsafe.Pointer(parentDescriptor)),
		uintptr(unsafe.Pointer(creatorDescriptor)),
		uintptr(unsafe.Pointer(newDescriptor)),
		getUintptrFromBool(isDirectoryObject),
		uintptr(token),
		uintptr(unsafe.Pointer(genericMapping)))
	return ret1 != 0
}

func CreatePrivateObjectSecurityEx(parentDescriptor PSECURITY_DESCRIPTOR, creatorDescriptor PSECURITY_DESCRIPTOR, newDescriptor *PSECURITY_DESCRIPTOR, objectType *GUID, isContainerObject bool, autoInheritFlags ULONG, token HANDLE, genericMapping *GENERIC_MAPPING) bool {
	ret1 := syscall9(createPrivateObjectSecurityEx, 8,
		uintptr(unsafe.Pointer(parentDescriptor)),
		uintptr(unsafe.Pointer(creatorDescriptor)),
		uintptr(unsafe.Pointer(newDescriptor)),
		uintptr(unsafe.Pointer(objectType)),
		getUintptrFromBool(isContainerObject),
		uintptr(autoInheritFlags),
		uintptr(token),
		uintptr(unsafe.Pointer(genericMapping)),
		0)
	return ret1 != 0
}

func CreatePrivateObjectSecurityWithMultipleInheritance(parentDescriptor PSECURITY_DESCRIPTOR, creatorDescriptor PSECURITY_DESCRIPTOR, newDescriptor *PSECURITY_DESCRIPTOR, objectTypes uintptr, guidCount ULONG, isContainerObject bool, autoInheritFlags ULONG, token HANDLE, genericMapping *GENERIC_MAPPING) bool {
	ret1 := syscall9(createPrivateObjectSecurityWithMultipleInheritance, 9,
		uintptr(unsafe.Pointer(parentDescriptor)),
		uintptr(unsafe.Pointer(creatorDescriptor)),
		uintptr(unsafe.Pointer(newDescriptor)),
		objectTypes,
		uintptr(guidCount),
		getUintptrFromBool(isContainerObject),
		uintptr(autoInheritFlags),
		uintptr(token),
		uintptr(unsafe.Pointer(genericMapping)))
	return ret1 != 0
}

func CreateProcessAsUser(hToken HANDLE, lpApplicationName string, lpCommandLine LPWSTR, lpProcessAttributes *SECURITY_ATTRIBUTES, lpThreadAttributes *SECURITY_ATTRIBUTES, bInheritHandles bool, dwCreationFlags uint32, lpEnvironment LPVOID, lpCurrentDirectory string, lpStartupInfo *STARTUPINFO, lpProcessInformation *PROCESS_INFORMATION) bool {
	lpApplicationNameStr := unicode16FromString(lpApplicationName)
	lpCurrentDirectoryStr := unicode16FromString(lpCurrentDirectory)
	ret1 := syscall12(createProcessAsUser, 11,
		uintptr(hToken),
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
		0)
	return ret1 != 0
}

func CreateProcessWithLogonW(lpUsername string, lpDomain string, lpPassword string, dwLogonFlags uint32, lpApplicationName string, lpCommandLine LPWSTR, dwCreationFlags uint32, lpEnvironment LPVOID, lpCurrentDirectory string, lpStartupInfo *STARTUPINFO, lpProcessInformation *PROCESS_INFORMATION) bool {
	lpUsernameStr := unicode16FromString(lpUsername)
	lpDomainStr := unicode16FromString(lpDomain)
	lpPasswordStr := unicode16FromString(lpPassword)
	lpApplicationNameStr := unicode16FromString(lpApplicationName)
	lpCurrentDirectoryStr := unicode16FromString(lpCurrentDirectory)
	ret1 := syscall12(createProcessWithLogonW, 11,
		uintptr(unsafe.Pointer(&lpUsernameStr[0])),
		uintptr(unsafe.Pointer(&lpDomainStr[0])),
		uintptr(unsafe.Pointer(&lpPasswordStr[0])),
		uintptr(dwLogonFlags),
		uintptr(unsafe.Pointer(&lpApplicationNameStr[0])),
		uintptr(unsafe.Pointer(lpCommandLine)),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpEnvironment)),
		uintptr(unsafe.Pointer(&lpCurrentDirectoryStr[0])),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)),
		0)
	return ret1 != 0
}

func CreateProcessWithTokenW(hToken HANDLE, dwLogonFlags uint32, lpApplicationName string, lpCommandLine LPWSTR, dwCreationFlags uint32, lpEnvironment LPVOID, lpCurrentDirectory string, lpStartupInfo *STARTUPINFO, lpProcessInformation *PROCESS_INFORMATION) bool {
	lpApplicationNameStr := unicode16FromString(lpApplicationName)
	lpCurrentDirectoryStr := unicode16FromString(lpCurrentDirectory)
	ret1 := syscall9(createProcessWithTokenW, 9,
		uintptr(hToken),
		uintptr(dwLogonFlags),
		uintptr(unsafe.Pointer(&lpApplicationNameStr[0])),
		uintptr(unsafe.Pointer(lpCommandLine)),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpEnvironment)),
		uintptr(unsafe.Pointer(&lpCurrentDirectoryStr[0])),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)))
	return ret1 != 0
}

// TODO: Unknown type(s): PLUID_AND_ATTRIBUTES, PSID_AND_ATTRIBUTES
// func CreateRestrictedToken(existingTokenHandle HANDLE, flags uint32, disableSidCount uint32, sidsToDisable PSID_AND_ATTRIBUTES, deletePrivilegeCount uint32, privilegesToDelete PLUID_AND_ATTRIBUTES, restrictedSidCount uint32, sidsToRestrict PSID_AND_ATTRIBUTES, newTokenHandle *HANDLE) bool

func CreateService(hSCManager SC_HANDLE, lpServiceName string, lpDisplayName string, dwDesiredAccess uint32, dwServiceType uint32, dwStartType uint32, dwErrorControl uint32, lpBinaryPathName string, lpLoadOrderGroup string, lpdwTagId *uint32, lpDependencies string, lpServiceStartName string, lpPassword string) SC_HANDLE {
	lpServiceNameStr := unicode16FromString(lpServiceName)
	lpDisplayNameStr := unicode16FromString(lpDisplayName)
	lpBinaryPathNameStr := unicode16FromString(lpBinaryPathName)
	lpLoadOrderGroupStr := unicode16FromString(lpLoadOrderGroup)
	lpDependenciesStr := unicode16FromString(lpDependencies)
	lpServiceStartNameStr := unicode16FromString(lpServiceStartName)
	lpPasswordStr := unicode16FromString(lpPassword)
	ret1 := syscall15(createService, 13,
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(&lpServiceNameStr[0])),
		uintptr(unsafe.Pointer(&lpDisplayNameStr[0])),
		uintptr(dwDesiredAccess),
		uintptr(dwServiceType),
		uintptr(dwStartType),
		uintptr(dwErrorControl),
		uintptr(unsafe.Pointer(&lpBinaryPathNameStr[0])),
		uintptr(unsafe.Pointer(&lpLoadOrderGroupStr[0])),
		uintptr(unsafe.Pointer(lpdwTagId)),
		uintptr(unsafe.Pointer(&lpDependenciesStr[0])),
		uintptr(unsafe.Pointer(&lpServiceStartNameStr[0])),
		uintptr(unsafe.Pointer(&lpPasswordStr[0])),
		0,
		0)
	return SC_HANDLE(ret1)
}

// TODO: Unknown type(s): WELL_KNOWN_SID_TYPE
// func CreateWellKnownSid(wellKnownSidType WELL_KNOWN_SID_TYPE, domainSid PSID, pSid PSID, cbSid *uint32) bool

func CredDelete(targetName string, aType uint32, flags uint32) bool {
	targetNameStr := unicode16FromString(targetName)
	ret1 := syscall3(credDelete, 3,
		uintptr(unsafe.Pointer(&targetNameStr[0])),
		uintptr(aType),
		uintptr(flags))
	return ret1 != 0
}

func CredEnumerate(filter string, flags uint32, count *uint32, credential uintptr) bool {
	filterStr := unicode16FromString(filter)
	ret1 := syscall6(credEnumerate, 4,
		uintptr(unsafe.Pointer(&filterStr[0])),
		uintptr(flags),
		uintptr(unsafe.Pointer(count)),
		credential,
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCREDENTIALW *
// func CredFindBestCredential(targetName string, aType uint32, flags uint32, credential PCREDENTIALW *) bool

func CredFree(buffer uintptr) {
	syscall3(credFree, 1,
		buffer,
		0,
		0)
}

func CredGetSessionTypes(maximumPersistCount uint32, maximumPersist *uint32) bool {
	ret1 := syscall3(credGetSessionTypes, 2,
		uintptr(maximumPersistCount),
		uintptr(unsafe.Pointer(maximumPersist)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCREDENTIAL_TARGET_INFORMATIONW *
// func CredGetTargetInfo(targetName string, flags uint32, targetInfo PCREDENTIAL_TARGET_INFORMATIONW *) bool

func CredIsMarshaledCredential(marshaledCredential string) bool {
	marshaledCredentialStr := unicode16FromString(marshaledCredential)
	ret1 := syscall3(credIsMarshaledCredential, 1,
		uintptr(unsafe.Pointer(&marshaledCredentialStr[0])),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): CRED_PROTECTION_TYPE *
// func CredIsProtected(pszProtectedCredentials LPWSTR, pProtectionType CRED_PROTECTION_TYPE *) bool

// TODO: Unknown type(s): CRED_MARSHAL_TYPE
// func CredMarshalCredential(credType CRED_MARSHAL_TYPE, credential uintptr, marshaledCredential *LPWSTR) bool

// TODO: Unknown type(s): CRED_PROTECTION_TYPE *
// func CredProtect(fAsSelf bool, pszCredentials LPWSTR, cchCredentials uint32, pszProtectedCredentials LPWSTR, pcchMaxChars *uint32, protectionType CRED_PROTECTION_TYPE *) bool

// TODO: Unknown type(s): PCREDENTIAL_TARGET_INFORMATIONW
// func CredReadDomainCredentials(targetInfo PCREDENTIAL_TARGET_INFORMATIONW, flags uint32, count *uint32, credential uintptr) bool

// TODO: Unknown type(s): PCREDENTIALW *
// func CredRead(targetName string, aType uint32, flags uint32, credential PCREDENTIALW *) bool

func CredRename(oldTargetName string, newTargetName string, aType uint32, flags uint32) bool {
	oldTargetNameStr := unicode16FromString(oldTargetName)
	newTargetNameStr := unicode16FromString(newTargetName)
	ret1 := syscall6(credRename, 4,
		uintptr(unsafe.Pointer(&oldTargetNameStr[0])),
		uintptr(unsafe.Pointer(&newTargetNameStr[0])),
		uintptr(aType),
		uintptr(flags),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCRED_MARSHAL_TYPE
// func CredUnmarshalCredential(marshaledCredential string, credType PCRED_MARSHAL_TYPE, credential *PVOID) bool

func CredUnprotect(fAsSelf bool, pszProtectedCredentials LPWSTR, cchCredentials uint32, pszCredentials LPWSTR, pcchMaxChars *uint32) bool {
	ret1 := syscall6(credUnprotect, 5,
		getUintptrFromBool(fAsSelf),
		uintptr(unsafe.Pointer(pszProtectedCredentials)),
		uintptr(cchCredentials),
		uintptr(unsafe.Pointer(pszCredentials)),
		uintptr(unsafe.Pointer(pcchMaxChars)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PCREDENTIAL_TARGET_INFORMATIONW
// func CredWriteDomainCredentials(targetInfo PCREDENTIAL_TARGET_INFORMATIONW, credential *CREDENTIAL, flags uint32) bool

func CredWrite(credential *CREDENTIAL, flags uint32) bool {
	ret1 := syscall3(credWrite, 2,
		uintptr(unsafe.Pointer(credential)),
		uintptr(flags),
		0)
	return ret1 != 0
}

func CryptAcquireContext(phProv *HCRYPTPROV, szContainer string, szProvider string, dwProvType uint32, dwFlags uint32) bool {
	szContainerStr := unicode16FromString(szContainer)
	szProviderStr := unicode16FromString(szProvider)
	ret1 := syscall6(cryptAcquireContext, 5,
		uintptr(unsafe.Pointer(phProv)),
		uintptr(unsafe.Pointer(&szContainerStr[0])),
		uintptr(unsafe.Pointer(&szProviderStr[0])),
		uintptr(dwProvType),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func CryptContextAddRef(hProv HCRYPTPROV, pdwReserved *uint32, dwFlags uint32) bool {
	ret1 := syscall3(cryptContextAddRef, 3,
		uintptr(hProv),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags))
	return ret1 != 0
}

func CryptCreateHash(hProv HCRYPTPROV, algid ALG_ID, hKey HCRYPTKEY, dwFlags uint32, phHash *HCRYPTHASH) bool {
	ret1 := syscall6(cryptCreateHash, 5,
		uintptr(hProv),
		uintptr(algid),
		uintptr(hKey),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(phHash)),
		0)
	return ret1 != 0
}

func CryptDecrypt(hKey HCRYPTKEY, hHash HCRYPTHASH, final bool, dwFlags uint32, pbData *byte, pdwDataLen *uint32) bool {
	ret1 := syscall6(cryptDecrypt, 6,
		uintptr(hKey),
		uintptr(hHash),
		getUintptrFromBool(final),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(unsafe.Pointer(pdwDataLen)))
	return ret1 != 0
}

func CryptDeriveKey(hProv HCRYPTPROV, algid ALG_ID, hBaseData HCRYPTHASH, dwFlags uint32, phKey *HCRYPTKEY) bool {
	ret1 := syscall6(cryptDeriveKey, 5,
		uintptr(hProv),
		uintptr(algid),
		uintptr(hBaseData),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(phKey)),
		0)
	return ret1 != 0
}

func CryptDestroyHash(hHash HCRYPTHASH) bool {
	ret1 := syscall3(cryptDestroyHash, 1,
		uintptr(hHash),
		0,
		0)
	return ret1 != 0
}

func CryptDestroyKey(hKey HCRYPTKEY) bool {
	ret1 := syscall3(cryptDestroyKey, 1,
		uintptr(hKey),
		0,
		0)
	return ret1 != 0
}

func CryptDuplicateHash(hHash HCRYPTHASH, pdwReserved *uint32, dwFlags uint32, phHash *HCRYPTHASH) bool {
	ret1 := syscall6(cryptDuplicateHash, 4,
		uintptr(hHash),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(phHash)),
		0,
		0)
	return ret1 != 0
}

func CryptDuplicateKey(hKey HCRYPTKEY, pdwReserved *uint32, dwFlags uint32, phKey *HCRYPTKEY) bool {
	ret1 := syscall6(cryptDuplicateKey, 4,
		uintptr(hKey),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(phKey)),
		0,
		0)
	return ret1 != 0
}

func CryptEncrypt(hKey HCRYPTKEY, hHash HCRYPTHASH, final bool, dwFlags uint32, pbData *byte, pdwDataLen *uint32, dwBufLen uint32) bool {
	ret1 := syscall9(cryptEncrypt, 7,
		uintptr(hKey),
		uintptr(hHash),
		getUintptrFromBool(final),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(unsafe.Pointer(pdwDataLen)),
		uintptr(dwBufLen),
		0,
		0)
	return ret1 != 0
}

func CryptEnumProviderTypes(dwIndex uint32, pdwReserved *uint32, dwFlags uint32, pdwProvType *uint32, szTypeName LPWSTR, pcbTypeName *uint32) bool {
	ret1 := syscall6(cryptEnumProviderTypes, 6,
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pdwProvType)),
		uintptr(unsafe.Pointer(szTypeName)),
		uintptr(unsafe.Pointer(pcbTypeName)))
	return ret1 != 0
}

func CryptEnumProviders(dwIndex uint32, pdwReserved *uint32, dwFlags uint32, pdwProvType *uint32, szProvName LPWSTR, pcbProvName *uint32) bool {
	ret1 := syscall6(cryptEnumProviders, 6,
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pdwProvType)),
		uintptr(unsafe.Pointer(szProvName)),
		uintptr(unsafe.Pointer(pcbProvName)))
	return ret1 != 0
}

func CryptExportKey(hKey HCRYPTKEY, hExpKey HCRYPTKEY, dwBlobType uint32, dwFlags uint32, pbData *byte, pdwDataLen *uint32) bool {
	ret1 := syscall6(cryptExportKey, 6,
		uintptr(hKey),
		uintptr(hExpKey),
		uintptr(dwBlobType),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(unsafe.Pointer(pdwDataLen)))
	return ret1 != 0
}

func CryptGenKey(hProv HCRYPTPROV, algid ALG_ID, dwFlags uint32, phKey *HCRYPTKEY) bool {
	ret1 := syscall6(cryptGenKey, 4,
		uintptr(hProv),
		uintptr(algid),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(phKey)),
		0,
		0)
	return ret1 != 0
}

func CryptGenRandom(hProv HCRYPTPROV, dwLen uint32, pbBuffer *byte) bool {
	ret1 := syscall3(cryptGenRandom, 3,
		uintptr(hProv),
		uintptr(dwLen),
		uintptr(unsafe.Pointer(pbBuffer)))
	return ret1 != 0
}

func CryptGetDefaultProvider(dwProvType uint32, pdwReserved *uint32, dwFlags uint32, pszProvName LPWSTR, pcbProvName *uint32) bool {
	ret1 := syscall6(cryptGetDefaultProvider, 5,
		uintptr(dwProvType),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pszProvName)),
		uintptr(unsafe.Pointer(pcbProvName)),
		0)
	return ret1 != 0
}

func CryptGetHashParam(hHash HCRYPTHASH, dwParam uint32, pbData *byte, pdwDataLen *uint32, dwFlags uint32) bool {
	ret1 := syscall6(cryptGetHashParam, 5,
		uintptr(hHash),
		uintptr(dwParam),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(unsafe.Pointer(pdwDataLen)),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func CryptGetKeyParam(hKey HCRYPTKEY, dwParam uint32, pbData *byte, pdwDataLen *uint32, dwFlags uint32) bool {
	ret1 := syscall6(cryptGetKeyParam, 5,
		uintptr(hKey),
		uintptr(dwParam),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(unsafe.Pointer(pdwDataLen)),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func CryptGetProvParam(hProv HCRYPTPROV, dwParam uint32, pbData *byte, pdwDataLen *uint32, dwFlags uint32) bool {
	ret1 := syscall6(cryptGetProvParam, 5,
		uintptr(hProv),
		uintptr(dwParam),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(unsafe.Pointer(pdwDataLen)),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func CryptGetUserKey(hProv HCRYPTPROV, dwKeySpec uint32, phUserKey *HCRYPTKEY) bool {
	ret1 := syscall3(cryptGetUserKey, 3,
		uintptr(hProv),
		uintptr(dwKeySpec),
		uintptr(unsafe.Pointer(phUserKey)))
	return ret1 != 0
}

func CryptHashData(hHash HCRYPTHASH, pbData /*const*/ *byte, dwDataLen uint32, dwFlags uint32) bool {
	ret1 := syscall6(cryptHashData, 4,
		uintptr(hHash),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(dwDataLen),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func CryptHashSessionKey(hHash HCRYPTHASH, hKey HCRYPTKEY, dwFlags uint32) bool {
	ret1 := syscall3(cryptHashSessionKey, 3,
		uintptr(hHash),
		uintptr(hKey),
		uintptr(dwFlags))
	return ret1 != 0
}

func CryptImportKey(hProv HCRYPTPROV, pbData /*const*/ *byte, dwDataLen uint32, hPubKey HCRYPTKEY, dwFlags uint32, phKey *HCRYPTKEY) bool {
	ret1 := syscall6(cryptImportKey, 6,
		uintptr(hProv),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(dwDataLen),
		uintptr(hPubKey),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(phKey)))
	return ret1 != 0
}

func CryptReleaseContext(hProv HCRYPTPROV, dwFlags uint32) bool {
	ret1 := syscall3(cryptReleaseContext, 2,
		uintptr(hProv),
		uintptr(dwFlags),
		0)
	return ret1 != 0
}

func CryptSetHashParam(hHash HCRYPTHASH, dwParam uint32, pbData /*const*/ *byte, dwFlags uint32) bool {
	ret1 := syscall6(cryptSetHashParam, 4,
		uintptr(hHash),
		uintptr(dwParam),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func CryptSetKeyParam(hKey HCRYPTKEY, dwParam uint32, pbData /*const*/ *byte, dwFlags uint32) bool {
	ret1 := syscall6(cryptSetKeyParam, 4,
		uintptr(hKey),
		uintptr(dwParam),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func CryptSetProvParam(hProv HCRYPTPROV, dwParam uint32, pbData /*const*/ *byte, dwFlags uint32) bool {
	ret1 := syscall6(cryptSetProvParam, 4,
		uintptr(hProv),
		uintptr(dwParam),
		uintptr(unsafe.Pointer(pbData)),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func CryptSetProviderEx(pszProvName string, dwProvType uint32, pdwReserved *uint32, dwFlags uint32) bool {
	pszProvNameStr := unicode16FromString(pszProvName)
	ret1 := syscall6(cryptSetProviderEx, 4,
		uintptr(unsafe.Pointer(&pszProvNameStr[0])),
		uintptr(dwProvType),
		uintptr(unsafe.Pointer(pdwReserved)),
		uintptr(dwFlags),
		0,
		0)
	return ret1 != 0
}

func CryptSetProvider(pszProvName string, dwProvType uint32) bool {
	pszProvNameStr := unicode16FromString(pszProvName)
	ret1 := syscall3(cryptSetProvider, 2,
		uintptr(unsafe.Pointer(&pszProvNameStr[0])),
		uintptr(dwProvType),
		0)
	return ret1 != 0
}

func CryptSignHash(hHash HCRYPTHASH, dwKeySpec uint32, szDescription string, dwFlags uint32, pbSignature *byte, pdwSigLen *uint32) bool {
	szDescriptionStr := unicode16FromString(szDescription)
	ret1 := syscall6(cryptSignHash, 6,
		uintptr(hHash),
		uintptr(dwKeySpec),
		uintptr(unsafe.Pointer(&szDescriptionStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pbSignature)),
		uintptr(unsafe.Pointer(pdwSigLen)))
	return ret1 != 0
}

func CryptVerifySignature(hHash HCRYPTHASH, pbSignature /*const*/ *byte, dwSigLen uint32, hPubKey HCRYPTKEY, szDescription string, dwFlags uint32) bool {
	szDescriptionStr := unicode16FromString(szDescription)
	ret1 := syscall6(cryptVerifySignature, 6,
		uintptr(hHash),
		uintptr(unsafe.Pointer(pbSignature)),
		uintptr(dwSigLen),
		uintptr(hPubKey),
		uintptr(unsafe.Pointer(&szDescriptionStr[0])),
		uintptr(dwFlags))
	return ret1 != 0
}

func DecryptFile(lpFileName string, dwReserved uint32) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(decryptFile, 2,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(dwReserved),
		0)
	return ret1 != 0
}

func DeleteAce(pAcl *ACL, dwAceIndex uint32) bool {
	ret1 := syscall3(deleteAce, 2,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceIndex),
		0)
	return ret1 != 0
}

func DeleteService(hService SC_HANDLE) bool {
	ret1 := syscall3(deleteService, 1,
		uintptr(hService),
		0,
		0)
	return ret1 != 0
}

func DeregisterEventSource(hEventLog HANDLE) bool {
	ret1 := syscall3(deregisterEventSource, 1,
		uintptr(hEventLog),
		0,
		0)
	return ret1 != 0
}

func DestroyPrivateObjectSecurity(objectDescriptor *PSECURITY_DESCRIPTOR) bool {
	ret1 := syscall3(destroyPrivateObjectSecurity, 1,
		uintptr(unsafe.Pointer(objectDescriptor)),
		0,
		0)
	return ret1 != 0
}

func DuplicateEncryptionInfoFile(srcFileName string, dstFileName string, dwCreationDistribution uint32, dwAttributes uint32, lpSecurityAttributes /*const*/ *SECURITY_ATTRIBUTES) uint32 {
	srcFileNameStr := unicode16FromString(srcFileName)
	dstFileNameStr := unicode16FromString(dstFileName)
	ret1 := syscall6(duplicateEncryptionInfoFile, 5,
		uintptr(unsafe.Pointer(&srcFileNameStr[0])),
		uintptr(unsafe.Pointer(&dstFileNameStr[0])),
		uintptr(dwCreationDistribution),
		uintptr(dwAttributes),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		0)
	return uint32(ret1)
}

func DuplicateToken(existingTokenHandle HANDLE, impersonationLevel SECURITY_IMPERSONATION_LEVEL, duplicateTokenHandle *HANDLE) bool {
	ret1 := syscall3(duplicateToken, 3,
		uintptr(existingTokenHandle),
		uintptr(impersonationLevel),
		uintptr(unsafe.Pointer(duplicateTokenHandle)))
	return ret1 != 0
}

// TODO: Unknown type(s): TOKEN_TYPE
// func DuplicateTokenEx(hExistingToken HANDLE, dwDesiredAccess uint32, lpTokenAttributes *SECURITY_ATTRIBUTES, impersonationLevel SECURITY_IMPERSONATION_LEVEL, tokenType TOKEN_TYPE, phNewToken *HANDLE) bool

func EncryptFile(lpFileName string) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(encryptFile, 1,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return ret1 != 0
}

func EncryptionDisable(dirPath string, disable bool) bool {
	dirPathStr := unicode16FromString(dirPath)
	ret1 := syscall3(encryptionDisable, 2,
		uintptr(unsafe.Pointer(&dirPathStr[0])),
		getUintptrFromBool(disable),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPENUM_SERVICE_STATUSW
// func EnumDependentServices(hService SC_HANDLE, dwServiceState uint32, lpServices LPENUM_SERVICE_STATUSW, cbBufSize uint32, pcbBytesNeeded *uint32, lpServicesReturned *uint32) bool

// TODO: Unknown type(s): SC_ENUM_TYPE
// func EnumServicesStatusEx(hSCManager SC_HANDLE, infoLevel SC_ENUM_TYPE, dwServiceType uint32, dwServiceState uint32, lpServices *byte, cbBufSize uint32, pcbBytesNeeded *uint32, lpServicesReturned *uint32, lpResumeHandle *uint32, pszGroupName string) bool

// TODO: Unknown type(s): LPENUM_SERVICE_STATUSW
// func EnumServicesStatus(hSCManager SC_HANDLE, dwServiceType uint32, dwServiceState uint32, lpServices LPENUM_SERVICE_STATUSW, cbBufSize uint32, pcbBytesNeeded *uint32, lpServicesReturned *uint32, lpResumeHandle *uint32) bool

func EqualDomainSid(pSid1 PSID, pSid2 PSID, pfEqual *BOOL) bool {
	ret1 := syscall3(equalDomainSid, 3,
		uintptr(pSid1),
		uintptr(pSid2),
		uintptr(unsafe.Pointer(pfEqual)))
	return ret1 != 0
}

func EqualPrefixSid(pSid1 PSID, pSid2 PSID) bool {
	ret1 := syscall3(equalPrefixSid, 2,
		uintptr(pSid1),
		uintptr(pSid2),
		0)
	return ret1 != 0
}

func EqualSid(pSid1 PSID, pSid2 PSID) bool {
	ret1 := syscall3(equalSid, 2,
		uintptr(pSid1),
		uintptr(pSid2),
		0)
	return ret1 != 0
}

func FileEncryptionStatus(lpFileName string, lpStatus *uint32) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(fileEncryptionStatus, 2,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(unsafe.Pointer(lpStatus)),
		0)
	return ret1 != 0
}

func FindFirstFreeAce(pAcl *ACL, pAce *LPVOID) bool {
	ret1 := syscall3(findFirstFreeAce, 2,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(unsafe.Pointer(pAce)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PENCRYPTION_CERTIFICATE_HASH_LIST
// func FreeEncryptionCertificateHashList(pHashes PENCRYPTION_CERTIFICATE_HASH_LIST)

// TODO: Unknown type(s): PFN_OBJECT_MGR_FUNCTS, PINHERITED_FROMW
// func FreeInheritedFromArray(pInheritArray PINHERITED_FROMW, aceCnt USHORT, pfnArray PFN_OBJECT_MGR_FUNCTS) uint32

func FreeSid(pSid PSID) uintptr {
	ret1 := syscall3(freeSid, 1,
		uintptr(pSid),
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func GetAce(pAcl *ACL, dwAceIndex uint32, pAce *LPVOID) bool {
	ret1 := syscall3(getAce, 3,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(dwAceIndex),
		uintptr(unsafe.Pointer(pAce)))
	return ret1 != 0
}

// TODO: Unknown type(s): ACL_INFORMATION_CLASS
// func GetAclInformation(pAcl *ACL, pAclInformation LPVOID, nAclInformationLength uint32, dwAclInformationClass ACL_INFORMATION_CLASS) bool

// TODO: Unknown type(s): PACCESS_MASK
// func GetAuditedPermissionsFromAcl(pacl *ACL, pTrustee *TRUSTEE, pSuccessfulAuditedRights PACCESS_MASK, pFailedAuditRights PACCESS_MASK) uint32

// TODO: Unknown type(s): LPHW_PROFILE_INFOW
// func GetCurrentHwProfile(lpHwProfileInfo LPHW_PROFILE_INFOW) bool

// TODO: Unknown type(s): PACCESS_MASK
// func GetEffectiveRightsFromAcl(pacl *ACL, pTrustee *TRUSTEE, pAccessRights PACCESS_MASK) uint32

func GetEventLogInformation(hEventLog HANDLE, dwInfoLevel uint32, lpBuffer LPVOID, cbBufSize uint32, pcbBytesNeeded *uint32) bool {
	ret1 := syscall6(getEventLogInformation, 5,
		uintptr(hEventLog),
		uintptr(dwInfoLevel),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cbBufSize),
		uintptr(unsafe.Pointer(pcbBytesNeeded)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PEXPLICIT_ACCESS_W *
// func GetExplicitEntriesFromAcl(pacl *ACL, pcCountOfExplicitEntries *uint32, pListOfExplicitEntries PEXPLICIT_ACCESS_W *) uint32

func GetFileSecurity(lpFileName string, requestedInformation SECURITY_INFORMATION, pSecurityDescriptor PSECURITY_DESCRIPTOR, nLength uint32, lpnLengthNeeded *uint32) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall6(getFileSecurity, 5,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(requestedInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(nLength),
		uintptr(unsafe.Pointer(lpnLengthNeeded)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PFN_OBJECT_MGR_FUNCTS, PINHERITED_FROMW
// func GetInheritanceSource(pObjectName LPWSTR, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, container bool, pObjectClassGuids uintptr, guidCount uint32, pAcl *ACL, pfnArray PFN_OBJECT_MGR_FUNCTS, pGenericMapping *GENERIC_MAPPING, pInheritArray PINHERITED_FROMW) uint32

func GetKernelObjectSecurity(handle HANDLE, requestedInformation SECURITY_INFORMATION, pSecurityDescriptor PSECURITY_DESCRIPTOR, nLength uint32, lpnLengthNeeded *uint32) bool {
	ret1 := syscall6(getKernelObjectSecurity, 5,
		uintptr(handle),
		uintptr(requestedInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(nLength),
		uintptr(unsafe.Pointer(lpnLengthNeeded)),
		0)
	return ret1 != 0
}

func GetLengthSid(pSid PSID) uint32 {
	ret1 := syscall3(getLengthSid, 1,
		uintptr(pSid),
		0,
		0)
	return uint32(ret1)
}

func GetLocalManagedApplicationData(productCode *WCHAR, displayName *LPWSTR, supportUrl *LPWSTR) {
	syscall3(getLocalManagedApplicationData, 3,
		uintptr(unsafe.Pointer(productCode)),
		uintptr(unsafe.Pointer(displayName)),
		uintptr(unsafe.Pointer(supportUrl)))
}

// TODO: Unknown type(s): PLOCALMANAGEDAPPLICATION *
// func GetLocalManagedApplications(bUserApps bool, pdwApps *uint32, prgLocalApps PLOCALMANAGEDAPPLICATION *) uint32

// TODO: Unknown type(s): APPCATEGORYINFOLIST *
// func GetManagedApplicationCategories(dwReserved uint32, pAppCategory APPCATEGORYINFOLIST *) uint32

// TODO: Unknown type(s): PMANAGEDAPPLICATION *
// func GetManagedApplications(pCategory *GUID, dwQueryFlags uint32, dwInfoLevel uint32, pdwApps *uint32, prgManagedApps PMANAGEDAPPLICATION *) uint32

func GetMultipleTrusteeOperation(pTrustee *TRUSTEE) MULTIPLE_TRUSTEE_OPERATION {
	ret1 := syscall3(getMultipleTrusteeOperation, 1,
		uintptr(unsafe.Pointer(pTrustee)),
		0,
		0)
	return MULTIPLE_TRUSTEE_OPERATION(ret1)
}

func GetMultipleTrustee(pTrustee *TRUSTEE) *TRUSTEE {
	ret1 := syscall3(getMultipleTrustee, 1,
		uintptr(unsafe.Pointer(pTrustee)),
		0,
		0)
	return (*TRUSTEE)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PACL *
// func GetNamedSecurityInfo(pObjectName string, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, ppsidOwner *PSID, ppsidGroup *PSID, ppDacl PACL *, ppSacl PACL *, ppSecurityDescriptor *PSECURITY_DESCRIPTOR) uint32

func GetNumberOfEventLogRecords(hEventLog HANDLE, numberOfRecords *DWORD) bool {
	ret1 := syscall3(getNumberOfEventLogRecords, 2,
		uintptr(hEventLog),
		uintptr(unsafe.Pointer(numberOfRecords)),
		0)
	return ret1 != 0
}

func GetOldestEventLogRecord(hEventLog HANDLE, oldestRecord *DWORD) bool {
	ret1 := syscall3(getOldestEventLogRecord, 2,
		uintptr(hEventLog),
		uintptr(unsafe.Pointer(oldestRecord)),
		0)
	return ret1 != 0
}

func GetPrivateObjectSecurity(objectDescriptor PSECURITY_DESCRIPTOR, securityInformation SECURITY_INFORMATION, resultantDescriptor PSECURITY_DESCRIPTOR, descriptorLength uint32, returnLength *DWORD) bool {
	ret1 := syscall6(getPrivateObjectSecurity, 5,
		uintptr(unsafe.Pointer(objectDescriptor)),
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(resultantDescriptor)),
		uintptr(descriptorLength),
		uintptr(unsafe.Pointer(returnLength)),
		0)
	return ret1 != 0
}

func GetSecurityDescriptorControl(pSecurityDescriptor PSECURITY_DESCRIPTOR, pControl *SECURITY_DESCRIPTOR_CONTROL, lpdwRevision *uint32) bool {
	ret1 := syscall3(getSecurityDescriptorControl, 3,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(unsafe.Pointer(pControl)),
		uintptr(unsafe.Pointer(lpdwRevision)))
	return ret1 != 0
}

// TODO: Unknown type(s): PACL *
// func GetSecurityDescriptorDacl(pSecurityDescriptor PSECURITY_DESCRIPTOR, lpbDaclPresent *BOOL, pDacl PACL *, lpbDaclDefaulted *BOOL) bool

func GetSecurityDescriptorGroup(pSecurityDescriptor PSECURITY_DESCRIPTOR, pGroup *PSID, lpbGroupDefaulted *BOOL) bool {
	ret1 := syscall3(getSecurityDescriptorGroup, 3,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(unsafe.Pointer(pGroup)),
		uintptr(unsafe.Pointer(lpbGroupDefaulted)))
	return ret1 != 0
}

func GetSecurityDescriptorLength(pSecurityDescriptor PSECURITY_DESCRIPTOR) uint32 {
	ret1 := syscall3(getSecurityDescriptorLength, 1,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		0,
		0)
	return uint32(ret1)
}

func GetSecurityDescriptorOwner(pSecurityDescriptor PSECURITY_DESCRIPTOR, pOwner *PSID, lpbOwnerDefaulted *BOOL) bool {
	ret1 := syscall3(getSecurityDescriptorOwner, 3,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(unsafe.Pointer(pOwner)),
		uintptr(unsafe.Pointer(lpbOwnerDefaulted)))
	return ret1 != 0
}

// TODO: Unknown type(s): PUCHAR
// func GetSecurityDescriptorRMControl(securityDescriptor PSECURITY_DESCRIPTOR, rMControl PUCHAR) uint32

// TODO: Unknown type(s): PACL *
// func GetSecurityDescriptorSacl(pSecurityDescriptor PSECURITY_DESCRIPTOR, lpbSaclPresent *BOOL, pSacl PACL *, lpbSaclDefaulted *BOOL) bool

// TODO: Unknown type(s): PACL *
// func GetSecurityInfo(handle HANDLE, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, ppsidOwner *PSID, ppsidGroup *PSID, ppDacl PACL *, ppSacl PACL *, ppSecurityDescriptor *PSECURITY_DESCRIPTOR) uint32

func GetServiceDisplayName(hSCManager SC_HANDLE, lpServiceName string, lpDisplayName LPWSTR, lpcchBuffer *uint32) bool {
	lpServiceNameStr := unicode16FromString(lpServiceName)
	ret1 := syscall6(getServiceDisplayName, 4,
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(&lpServiceNameStr[0])),
		uintptr(unsafe.Pointer(lpDisplayName)),
		uintptr(unsafe.Pointer(lpcchBuffer)),
		0,
		0)
	return ret1 != 0
}

func GetServiceKeyName(hSCManager SC_HANDLE, lpDisplayName string, lpServiceName LPWSTR, lpcchBuffer *uint32) bool {
	lpDisplayNameStr := unicode16FromString(lpDisplayName)
	ret1 := syscall6(getServiceKeyName, 4,
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(&lpDisplayNameStr[0])),
		uintptr(unsafe.Pointer(lpServiceName)),
		uintptr(unsafe.Pointer(lpcchBuffer)),
		0,
		0)
	return ret1 != 0
}

func GetSidIdentifierAuthority(pSid PSID) *SID_IDENTIFIER_AUTHORITY {
	ret1 := syscall3(getSidIdentifierAuthority, 1,
		uintptr(pSid),
		0,
		0)
	return (*SID_IDENTIFIER_AUTHORITY)(unsafe.Pointer(ret1))
}

func GetSidLengthRequired(nSubAuthorityCount UCHAR) uint32 {
	ret1 := syscall3(getSidLengthRequired, 1,
		uintptr(nSubAuthorityCount),
		0,
		0)
	return uint32(ret1)
}

func GetSidSubAuthority(pSid PSID, nSubAuthority uint32) *DWORD {
	ret1 := syscall3(getSidSubAuthority, 2,
		uintptr(pSid),
		uintptr(nSubAuthority),
		0)
	return (*DWORD)(unsafe.Pointer(ret1))
}

// TODO: Unknown type(s): PUCHAR
// func GetSidSubAuthorityCount(pSid PSID) PUCHAR

// TODO: Unknown type(s): PWAITCHAIN_NODE_INFO
// func GetThreadWaitChain(wctHandle HWCT, context *uint32, flags uint32, threadId uint32, nodeCount *uint32, nodeInfoArray PWAITCHAIN_NODE_INFO, isCycle *BOOL) bool

// TODO: Unknown type(s): TOKEN_INFORMATION_CLASS
// func GetTokenInformation(tokenHandle HANDLE, tokenInformationClass TOKEN_INFORMATION_CLASS, tokenInformation LPVOID, tokenInformationLength uint32, returnLength *DWORD) bool

func GetTrusteeForm(pTrustee *TRUSTEE) TRUSTEE_FORM {
	ret1 := syscall3(getTrusteeForm, 1,
		uintptr(unsafe.Pointer(pTrustee)),
		0,
		0)
	return TRUSTEE_FORM(ret1)
}

func GetTrusteeName(pTrustee *TRUSTEE) LPWSTR {
	ret1 := syscall3(getTrusteeName, 1,
		uintptr(unsafe.Pointer(pTrustee)),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func GetTrusteeType(pTrustee *TRUSTEE) TRUSTEE_TYPE {
	ret1 := syscall3(getTrusteeType, 1,
		uintptr(unsafe.Pointer(pTrustee)),
		0,
		0)
	return TRUSTEE_TYPE(ret1)
}

func GetUserName(lpBuffer LPWSTR, pcbBuffer *uint32) bool {
	ret1 := syscall3(getUserName, 2,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(unsafe.Pointer(pcbBuffer)),
		0)
	return ret1 != 0
}

func GetWindowsAccountDomainSid(pSid PSID, pDomainSid PSID, cbDomainSid *uint32) bool {
	ret1 := syscall3(getWindowsAccountDomainSid, 3,
		uintptr(pSid),
		uintptr(pDomainSid),
		uintptr(unsafe.Pointer(cbDomainSid)))
	return ret1 != 0
}

func ImpersonateAnonymousToken(threadHandle HANDLE) bool {
	ret1 := syscall3(impersonateAnonymousToken, 1,
		uintptr(threadHandle),
		0,
		0)
	return ret1 != 0
}

func ImpersonateLoggedOnUser(hToken HANDLE) bool {
	ret1 := syscall3(impersonateLoggedOnUser, 1,
		uintptr(hToken),
		0,
		0)
	return ret1 != 0
}

func ImpersonateNamedPipeClient(hNamedPipe HANDLE) bool {
	ret1 := syscall3(impersonateNamedPipeClient, 1,
		uintptr(hNamedPipe),
		0,
		0)
	return ret1 != 0
}

func ImpersonateSelf(impersonationLevel SECURITY_IMPERSONATION_LEVEL) bool {
	ret1 := syscall3(impersonateSelf, 1,
		uintptr(impersonationLevel),
		0,
		0)
	return ret1 != 0
}

func InitializeAcl(pAcl *ACL, nAclLength uint32, dwAclRevision uint32) bool {
	ret1 := syscall3(initializeAcl, 3,
		uintptr(unsafe.Pointer(pAcl)),
		uintptr(nAclLength),
		uintptr(dwAclRevision))
	return ret1 != 0
}

func InitializeSecurityDescriptor(pSecurityDescriptor PSECURITY_DESCRIPTOR, dwRevision uint32) bool {
	ret1 := syscall3(initializeSecurityDescriptor, 2,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(dwRevision),
		0)
	return ret1 != 0
}

func InitializeSid(sid PSID, pIdentifierAuthority *SID_IDENTIFIER_AUTHORITY, nSubAuthorityCount byte) bool {
	ret1 := syscall3(initializeSid, 3,
		uintptr(sid),
		uintptr(unsafe.Pointer(pIdentifierAuthority)),
		uintptr(nSubAuthorityCount))
	return ret1 != 0
}

func InitiateShutdown(lpMachineName LPWSTR, lpMessage LPWSTR, dwGracePeriod uint32, dwShutdownFlags uint32, dwReason uint32) uint32 {
	ret1 := syscall6(initiateShutdown, 5,
		uintptr(unsafe.Pointer(lpMachineName)),
		uintptr(unsafe.Pointer(lpMessage)),
		uintptr(dwGracePeriod),
		uintptr(dwShutdownFlags),
		uintptr(dwReason),
		0)
	return uint32(ret1)
}

func InitiateSystemShutdownEx(lpMachineName LPWSTR, lpMessage LPWSTR, dwTimeout uint32, bForceAppsClosed bool, bRebootAfterShutdown bool, dwReason uint32) bool {
	ret1 := syscall6(initiateSystemShutdownEx, 6,
		uintptr(unsafe.Pointer(lpMachineName)),
		uintptr(unsafe.Pointer(lpMessage)),
		uintptr(dwTimeout),
		getUintptrFromBool(bForceAppsClosed),
		getUintptrFromBool(bRebootAfterShutdown),
		uintptr(dwReason))
	return ret1 != 0
}

func InitiateSystemShutdown(lpMachineName LPWSTR, lpMessage LPWSTR, dwTimeout uint32, bForceAppsClosed bool, bRebootAfterShutdown bool) bool {
	ret1 := syscall6(initiateSystemShutdown, 5,
		uintptr(unsafe.Pointer(lpMachineName)),
		uintptr(unsafe.Pointer(lpMessage)),
		uintptr(dwTimeout),
		getUintptrFromBool(bForceAppsClosed),
		getUintptrFromBool(bRebootAfterShutdown),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PINSTALLDATA
// func InstallApplication(pInstallInfo PINSTALLDATA) uint32

func IsTextUnicode(lpv /*const*/ uintptr, iSize int32, lpiResult *int32) bool {
	ret1 := syscall3(isTextUnicode, 3,
		lpv,
		uintptr(iSize),
		uintptr(unsafe.Pointer(lpiResult)))
	return ret1 != 0
}

func IsTokenRestricted(tokenHandle HANDLE) bool {
	ret1 := syscall3(isTokenRestricted, 1,
		uintptr(tokenHandle),
		0,
		0)
	return ret1 != 0
}

func IsTokenUntrusted(tokenHandle HANDLE) bool {
	ret1 := syscall3(isTokenUntrusted, 1,
		uintptr(tokenHandle),
		0,
		0)
	return ret1 != 0
}

func IsValidAcl(pAcl *ACL) bool {
	ret1 := syscall3(isValidAcl, 1,
		uintptr(unsafe.Pointer(pAcl)),
		0,
		0)
	return ret1 != 0
}

func IsValidSecurityDescriptor(pSecurityDescriptor PSECURITY_DESCRIPTOR) bool {
	ret1 := syscall3(isValidSecurityDescriptor, 1,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		0,
		0)
	return ret1 != 0
}

func IsValidSid(pSid PSID) bool {
	ret1 := syscall3(isValidSid, 1,
		uintptr(pSid),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): WELL_KNOWN_SID_TYPE
// func IsWellKnownSid(pSid PSID, wellKnownSidType WELL_KNOWN_SID_TYPE) bool

func LockServiceDatabase(hSCManager SC_HANDLE) SC_LOCK {
	ret1 := syscall3(lockServiceDatabase, 1,
		uintptr(hSCManager),
		0,
		0)
	return SC_LOCK(ret1)
}

// TODO: Unknown type(s): PQUOTA_LIMITS
// func LogonUserEx(lpszUsername string, lpszDomain string, lpszPassword string, dwLogonType uint32, dwLogonProvider uint32, phToken *HANDLE, ppLogonSid *PSID, ppProfileBuffer *PVOID, pdwProfileLength *uint32, pQuotaLimits PQUOTA_LIMITS) bool

func LogonUser(lpszUsername string, lpszDomain string, lpszPassword string, dwLogonType uint32, dwLogonProvider uint32, phToken *HANDLE) bool {
	lpszUsernameStr := unicode16FromString(lpszUsername)
	lpszDomainStr := unicode16FromString(lpszDomain)
	lpszPasswordStr := unicode16FromString(lpszPassword)
	ret1 := syscall6(logonUser, 6,
		uintptr(unsafe.Pointer(&lpszUsernameStr[0])),
		uintptr(unsafe.Pointer(&lpszDomainStr[0])),
		uintptr(unsafe.Pointer(&lpszPasswordStr[0])),
		uintptr(dwLogonType),
		uintptr(dwLogonProvider),
		uintptr(unsafe.Pointer(phToken)))
	return ret1 != 0
}

// TODO: Unknown type(s): PSID_NAME_USE
// func LookupAccountName(lpSystemName string, lpAccountName string, sid PSID, cbSid *uint32, referencedDomainName LPWSTR, cchReferencedDomainName *uint32, peUse PSID_NAME_USE) bool

// TODO: Unknown type(s): PSID_NAME_USE
// func LookupAccountSid(lpSystemName string, sid PSID, name LPWSTR, cchName *uint32, referencedDomainName LPWSTR, cchReferencedDomainName *uint32, peUse PSID_NAME_USE) bool

func LookupPrivilegeDisplayName(lpSystemName string, lpName string, lpDisplayName LPWSTR, cchDisplayName *uint32, lpLanguageId *uint32) bool {
	lpSystemNameStr := unicode16FromString(lpSystemName)
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall6(lookupPrivilegeDisplayName, 5,
		uintptr(unsafe.Pointer(&lpSystemNameStr[0])),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(lpDisplayName)),
		uintptr(unsafe.Pointer(cchDisplayName)),
		uintptr(unsafe.Pointer(lpLanguageId)),
		0)
	return ret1 != 0
}

func LookupPrivilegeName(lpSystemName string, lpLuid *LUID, lpName LPWSTR, cchName *uint32) bool {
	lpSystemNameStr := unicode16FromString(lpSystemName)
	ret1 := syscall6(lookupPrivilegeName, 4,
		uintptr(unsafe.Pointer(&lpSystemNameStr[0])),
		uintptr(unsafe.Pointer(lpLuid)),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(cchName)),
		0,
		0)
	return ret1 != 0
}

func LookupPrivilegeValue(lpSystemName string, lpName string, lpLuid *LUID) bool {
	lpSystemNameStr := unicode16FromString(lpSystemName)
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(lookupPrivilegeValue, 3,
		uintptr(unsafe.Pointer(&lpSystemNameStr[0])),
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		uintptr(unsafe.Pointer(lpLuid)))
	return ret1 != 0
}

// TODO: Unknown type(s): PEXPLICIT_ACCESS_W *, PTRUSTEE_W *
// func LookupSecurityDescriptorParts(ppOwner PTRUSTEE_W *, ppGroup PTRUSTEE_W *, pcCountOfAccessEntries *uint32, ppListOfAccessEntries PEXPLICIT_ACCESS_W *, pcCountOfAuditEntries *uint32, ppListOfAuditEntries PEXPLICIT_ACCESS_W *, pSD PSECURITY_DESCRIPTOR) uint32

// TODO: Unknown type(s): PLM_OWF_PASSWORD, PNT_OWF_PASSWORD
// func MSChapSrvChangePassword(serverName LPWSTR, userName LPWSTR, lmOldPresent BOOLEAN, lmOldOwfPassword PLM_OWF_PASSWORD, lmNewOwfPassword PLM_OWF_PASSWORD, ntOldOwfPassword PNT_OWF_PASSWORD, ntNewOwfPassword PNT_OWF_PASSWORD) uint32

// TODO: Unknown type(s): PENCRYPTED_LM_OWF_PASSWORD, PENCRYPTED_NT_OWF_PASSWORD, PSAMPR_ENCRYPTED_USER_PASSWORD
// func MSChapSrvChangePassword2(serverName LPWSTR, userName LPWSTR, newPasswordEncryptedWithOldNt PSAMPR_ENCRYPTED_USER_PASSWORD, oldNtOwfPasswordEncryptedWithNewNt PENCRYPTED_NT_OWF_PASSWORD, lmPresent BOOLEAN, newPasswordEncryptedWithOldLm PSAMPR_ENCRYPTED_USER_PASSWORD, oldLmOwfPasswordEncryptedWithNewLmOrNt PENCRYPTED_LM_OWF_PASSWORD) uint32

func MakeAbsoluteSD(pSelfRelativeSecurityDescriptor PSECURITY_DESCRIPTOR, pAbsoluteSecurityDescriptor PSECURITY_DESCRIPTOR, lpdwAbsoluteSecurityDescriptorSize *uint32, pDacl *ACL, lpdwDaclSize *uint32, pSacl *ACL, lpdwSaclSize *uint32, pOwner PSID, lpdwOwnerSize *uint32, pPrimaryGroup PSID, lpdwPrimaryGroupSize *uint32) bool {
	ret1 := syscall12(makeAbsoluteSD, 11,
		uintptr(unsafe.Pointer(pSelfRelativeSecurityDescriptor)),
		uintptr(unsafe.Pointer(pAbsoluteSecurityDescriptor)),
		uintptr(unsafe.Pointer(lpdwAbsoluteSecurityDescriptorSize)),
		uintptr(unsafe.Pointer(pDacl)),
		uintptr(unsafe.Pointer(lpdwDaclSize)),
		uintptr(unsafe.Pointer(pSacl)),
		uintptr(unsafe.Pointer(lpdwSaclSize)),
		uintptr(pOwner),
		uintptr(unsafe.Pointer(lpdwOwnerSize)),
		uintptr(pPrimaryGroup),
		uintptr(unsafe.Pointer(lpdwPrimaryGroupSize)),
		0)
	return ret1 != 0
}

func MakeSelfRelativeSD(pAbsoluteSecurityDescriptor PSECURITY_DESCRIPTOR, pSelfRelativeSecurityDescriptor PSECURITY_DESCRIPTOR, lpdwBufferLength *uint32) bool {
	ret1 := syscall3(makeSelfRelativeSD, 3,
		uintptr(unsafe.Pointer(pAbsoluteSecurityDescriptor)),
		uintptr(unsafe.Pointer(pSelfRelativeSecurityDescriptor)),
		uintptr(unsafe.Pointer(lpdwBufferLength)))
	return ret1 != 0
}

func MapGenericMask(accessMask *DWORD, genericMapping *GENERIC_MAPPING) {
	syscall3(mapGenericMask, 2,
		uintptr(unsafe.Pointer(accessMask)),
		uintptr(unsafe.Pointer(genericMapping)),
		0)
}

func NotifyBootConfigStatus(bootAcceptable bool) bool {
	ret1 := syscall3(notifyBootConfigStatus, 1,
		getUintptrFromBool(bootAcceptable),
		0,
		0)
	return ret1 != 0
}

func NotifyChangeEventLog(hEventLog HANDLE, hEvent HANDLE) bool {
	ret1 := syscall3(notifyChangeEventLog, 2,
		uintptr(hEventLog),
		uintptr(hEvent),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PSERVICE_NOTIFYW
// func NotifyServiceStatusChange(hService SC_HANDLE, dwNotifyMask uint32, pNotifyBuffer PSERVICE_NOTIFYW) uint32

func ObjectCloseAuditAlarm(subsystemName string, handleId LPVOID, generateOnClose bool) bool {
	subsystemNameStr := unicode16FromString(subsystemName)
	ret1 := syscall3(objectCloseAuditAlarm, 3,
		uintptr(unsafe.Pointer(&subsystemNameStr[0])),
		uintptr(unsafe.Pointer(handleId)),
		getUintptrFromBool(generateOnClose))
	return ret1 != 0
}

func ObjectDeleteAuditAlarm(subsystemName string, handleId LPVOID, generateOnClose bool) bool {
	subsystemNameStr := unicode16FromString(subsystemName)
	ret1 := syscall3(objectDeleteAuditAlarm, 3,
		uintptr(unsafe.Pointer(&subsystemNameStr[0])),
		uintptr(unsafe.Pointer(handleId)),
		getUintptrFromBool(generateOnClose))
	return ret1 != 0
}

func ObjectOpenAuditAlarm(subsystemName string, handleId LPVOID, objectTypeName LPWSTR, objectName LPWSTR, pSecurityDescriptor PSECURITY_DESCRIPTOR, clientToken HANDLE, desiredAccess uint32, grantedAccess uint32, privileges *PRIVILEGE_SET, objectCreation bool, accessGranted bool, generateOnClose *BOOL) bool {
	subsystemNameStr := unicode16FromString(subsystemName)
	ret1 := syscall12(objectOpenAuditAlarm, 12,
		uintptr(unsafe.Pointer(&subsystemNameStr[0])),
		uintptr(unsafe.Pointer(handleId)),
		uintptr(unsafe.Pointer(objectTypeName)),
		uintptr(unsafe.Pointer(objectName)),
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(clientToken),
		uintptr(desiredAccess),
		uintptr(grantedAccess),
		uintptr(unsafe.Pointer(privileges)),
		getUintptrFromBool(objectCreation),
		getUintptrFromBool(accessGranted),
		uintptr(unsafe.Pointer(generateOnClose)))
	return ret1 != 0
}

func ObjectPrivilegeAuditAlarm(subsystemName string, handleId LPVOID, clientToken HANDLE, desiredAccess uint32, privileges *PRIVILEGE_SET, accessGranted bool) bool {
	subsystemNameStr := unicode16FromString(subsystemName)
	ret1 := syscall6(objectPrivilegeAuditAlarm, 6,
		uintptr(unsafe.Pointer(&subsystemNameStr[0])),
		uintptr(unsafe.Pointer(handleId)),
		uintptr(clientToken),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(privileges)),
		getUintptrFromBool(accessGranted))
	return ret1 != 0
}

func OpenBackupEventLog(lpUNCServerName string, lpFileName string) HANDLE {
	lpUNCServerNameStr := unicode16FromString(lpUNCServerName)
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(openBackupEventLog, 2,
		uintptr(unsafe.Pointer(&lpUNCServerNameStr[0])),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0)
	return HANDLE(ret1)
}

func OpenEncryptedFileRaw(lpFileName string, ulFlags ULONG, pvContext *PVOID) uint32 {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(openEncryptedFileRaw, 3,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(ulFlags),
		uintptr(unsafe.Pointer(pvContext)))
	return uint32(ret1)
}

func OpenEventLog(lpUNCServerName string, lpSourceName string) HANDLE {
	lpUNCServerNameStr := unicode16FromString(lpUNCServerName)
	lpSourceNameStr := unicode16FromString(lpSourceName)
	ret1 := syscall3(openEventLog, 2,
		uintptr(unsafe.Pointer(&lpUNCServerNameStr[0])),
		uintptr(unsafe.Pointer(&lpSourceNameStr[0])),
		0)
	return HANDLE(ret1)
}

func OpenProcessToken(processHandle HANDLE, desiredAccess uint32, tokenHandle *HANDLE) bool {
	ret1 := syscall3(openProcessToken, 3,
		uintptr(processHandle),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(tokenHandle)))
	return ret1 != 0
}

func OpenSCManager(lpMachineName string, lpDatabaseName string, dwDesiredAccess uint32) SC_HANDLE {
	lpMachineNameStr := unicode16FromString(lpMachineName)
	lpDatabaseNameStr := unicode16FromString(lpDatabaseName)
	ret1 := syscall3(openSCManager, 3,
		uintptr(unsafe.Pointer(&lpMachineNameStr[0])),
		uintptr(unsafe.Pointer(&lpDatabaseNameStr[0])),
		uintptr(dwDesiredAccess))
	return SC_HANDLE(ret1)
}

func OpenService(hSCManager SC_HANDLE, lpServiceName string, dwDesiredAccess uint32) SC_HANDLE {
	lpServiceNameStr := unicode16FromString(lpServiceName)
	ret1 := syscall3(openService, 3,
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(&lpServiceNameStr[0])),
		uintptr(dwDesiredAccess))
	return SC_HANDLE(ret1)
}

func OpenThreadToken(threadHandle HANDLE, desiredAccess uint32, openAsSelf bool, tokenHandle *HANDLE) bool {
	ret1 := syscall6(openThreadToken, 4,
		uintptr(threadHandle),
		uintptr(desiredAccess),
		getUintptrFromBool(openAsSelf),
		uintptr(unsafe.Pointer(tokenHandle)),
		0,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PWAITCHAINCALLBACK
// func OpenThreadWaitChainSession(flags uint32, callback PWAITCHAINCALLBACK) HWCT

func PerfCreateInstance(hProvider HANDLE, counterSetGuid /*const*/ *GUID, szInstanceName string, dwInstance ULONG) *PERF_COUNTERSET_INSTANCE {
	szInstanceNameStr := unicode16FromString(szInstanceName)
	ret1 := syscall6(perfCreateInstance, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(counterSetGuid)),
		uintptr(unsafe.Pointer(&szInstanceNameStr[0])),
		uintptr(dwInstance),
		0,
		0)
	return (*PERF_COUNTERSET_INSTANCE)(unsafe.Pointer(ret1))
}

func PerfDecrementULongCounterValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, lValue ULONG) ULONG {
	ret1 := syscall6(perfDecrementULongCounterValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		uintptr(lValue),
		0,
		0)
	return ULONG(ret1)
}

func PerfDecrementULongLongCounterValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, llValue ULONGLONG) ULONG {
	ret1 := syscall6(perfDecrementULongLongCounterValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		uintptr(llValue),
		0,
		0)
	return ULONG(ret1)
}

func PerfDeleteInstance(hProvider HANDLE, instanceBlock *PERF_COUNTERSET_INSTANCE) ULONG {
	ret1 := syscall3(perfDeleteInstance, 2,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(instanceBlock)),
		0)
	return ULONG(ret1)
}

func PerfIncrementULongCounterValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, lValue ULONG) ULONG {
	ret1 := syscall6(perfIncrementULongCounterValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		uintptr(lValue),
		0,
		0)
	return ULONG(ret1)
}

func PerfIncrementULongLongCounterValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, llValue ULONGLONG) ULONG {
	ret1 := syscall6(perfIncrementULongLongCounterValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		uintptr(llValue),
		0,
		0)
	return ULONG(ret1)
}

func PerfQueryInstance(hProvider HANDLE, counterSetGuid /*const*/ *GUID, szInstance string, dwInstance ULONG) *PERF_COUNTERSET_INSTANCE {
	szInstanceStr := unicode16FromString(szInstance)
	ret1 := syscall6(perfQueryInstance, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(counterSetGuid)),
		uintptr(unsafe.Pointer(&szInstanceStr[0])),
		uintptr(dwInstance),
		0,
		0)
	return (*PERF_COUNTERSET_INSTANCE)(unsafe.Pointer(ret1))
}

func PerfSetCounterRefValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, lpAddr uintptr) ULONG {
	ret1 := syscall6(perfSetCounterRefValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		lpAddr,
		0,
		0)
	return ULONG(ret1)
}

// TODO: Unknown type(s): PPERF_COUNTERSET_INFO
// func PerfSetCounterSetInfo(hProvider HANDLE, pTemplate PPERF_COUNTERSET_INFO, dwTemplateSize ULONG) ULONG

func PerfSetULongCounterValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, lValue ULONG) ULONG {
	ret1 := syscall6(perfSetULongCounterValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		uintptr(lValue),
		0,
		0)
	return ULONG(ret1)
}

func PerfSetULongLongCounterValue(hProvider HANDLE, pInstance *PERF_COUNTERSET_INSTANCE, counterId ULONG, llValue ULONGLONG) ULONG {
	ret1 := syscall6(perfSetULongLongCounterValue, 4,
		uintptr(hProvider),
		uintptr(unsafe.Pointer(pInstance)),
		uintptr(counterId),
		uintptr(llValue),
		0,
		0)
	return ULONG(ret1)
}

// TODO: Unknown type(s): PERFLIBREQUEST
// func PerfStartProvider(providerGuid *GUID, controlCallback PERFLIBREQUEST, phProvider *HANDLE) ULONG

// TODO: Unknown type(s): PPERF_PROVIDER_CONTEXT
// func PerfStartProviderEx(providerGuid *GUID, providerContext PPERF_PROVIDER_CONTEXT, phProvider *HANDLE) ULONG

func PerfStopProvider(hProvider HANDLE) ULONG {
	ret1 := syscall3(perfStopProvider, 1,
		uintptr(hProvider),
		0,
		0)
	return ULONG(ret1)
}

func PrivilegeCheck(clientToken HANDLE, requiredPrivileges *PRIVILEGE_SET, pfResult *BOOL) bool {
	ret1 := syscall3(privilegeCheck, 3,
		uintptr(clientToken),
		uintptr(unsafe.Pointer(requiredPrivileges)),
		uintptr(unsafe.Pointer(pfResult)))
	return ret1 != 0
}

func PrivilegedServiceAuditAlarm(subsystemName string, serviceName string, clientToken HANDLE, privileges *PRIVILEGE_SET, accessGranted bool) bool {
	subsystemNameStr := unicode16FromString(subsystemName)
	serviceNameStr := unicode16FromString(serviceName)
	ret1 := syscall6(privilegedServiceAuditAlarm, 5,
		uintptr(unsafe.Pointer(&subsystemNameStr[0])),
		uintptr(unsafe.Pointer(&serviceNameStr[0])),
		uintptr(clientToken),
		uintptr(unsafe.Pointer(privileges)),
		getUintptrFromBool(accessGranted),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): PENCRYPTION_CERTIFICATE_HASH_LIST *
// func QueryRecoveryAgentsOnEncryptedFile(lpFileName string, pRecoveryAgents PENCRYPTION_CERTIFICATE_HASH_LIST *) uint32

func QuerySecurityAccessMask(securityInformation SECURITY_INFORMATION, desiredAccess *uint32) {
	syscall3(querySecurityAccessMask, 2,
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(desiredAccess)),
		0)
}

func QueryServiceConfig2(hService SC_HANDLE, dwInfoLevel uint32, lpBuffer *byte, cbBufSize uint32, pcbBytesNeeded *uint32) bool {
	ret1 := syscall6(queryServiceConfig2, 5,
		uintptr(hService),
		uintptr(dwInfoLevel),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cbBufSize),
		uintptr(unsafe.Pointer(pcbBytesNeeded)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): LPQUERY_SERVICE_CONFIGW
// func QueryServiceConfig(hService SC_HANDLE, lpServiceConfig LPQUERY_SERVICE_CONFIGW, cbBufSize uint32, pcbBytesNeeded *uint32) bool

// TODO: Unknown type(s): LPQUERY_SERVICE_LOCK_STATUSW
// func QueryServiceLockStatus(hSCManager SC_HANDLE, lpLockStatus LPQUERY_SERVICE_LOCK_STATUSW, cbBufSize uint32, pcbBytesNeeded *uint32) bool

func QueryServiceObjectSecurity(hService SC_HANDLE, dwSecurityInformation SECURITY_INFORMATION, lpSecurityDescriptor PSECURITY_DESCRIPTOR, cbBufSize uint32, pcbBytesNeeded *uint32) bool {
	ret1 := syscall6(queryServiceObjectSecurity, 5,
		uintptr(hService),
		uintptr(dwSecurityInformation),
		uintptr(unsafe.Pointer(lpSecurityDescriptor)),
		uintptr(cbBufSize),
		uintptr(unsafe.Pointer(pcbBytesNeeded)),
		0)
	return ret1 != 0
}

func QueryServiceStatus(hService SC_HANDLE, lpServiceStatus *SERVICE_STATUS) bool {
	ret1 := syscall3(queryServiceStatus, 2,
		uintptr(hService),
		uintptr(unsafe.Pointer(lpServiceStatus)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): SC_STATUS_TYPE
// func QueryServiceStatusEx(hService SC_HANDLE, infoLevel SC_STATUS_TYPE, lpBuffer *byte, cbBufSize uint32, pcbBytesNeeded *uint32) bool

// TODO: Unknown type(s): PENCRYPTION_CERTIFICATE_HASH_LIST *
// func QueryUsersOnEncryptedFile(lpFileName string, pUsers PENCRYPTION_CERTIFICATE_HASH_LIST *) uint32

// TODO: Unknown type(s): PFE_EXPORT_FUNC
// func ReadEncryptedFileRaw(pfExportCallback PFE_EXPORT_FUNC, pvCallbackContext uintptr, pvContext uintptr) uint32

func ReadEventLog(hEventLog HANDLE, dwReadFlags uint32, dwRecordOffset uint32, lpBuffer LPVOID, nNumberOfBytesToRead uint32, pnBytesRead *uint32, pnMinNumberOfBytesNeeded *uint32) bool {
	ret1 := syscall9(readEventLog, 7,
		uintptr(hEventLog),
		uintptr(dwReadFlags),
		uintptr(dwRecordOffset),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nNumberOfBytesToRead),
		uintptr(unsafe.Pointer(pnBytesRead)),
		uintptr(unsafe.Pointer(pnMinNumberOfBytesNeeded)),
		0,
		0)
	return ret1 != 0
}

func RegCloseKey(hKey HKEY) LONG {
	ret1 := syscall3(regCloseKey, 1,
		uintptr(hKey),
		0,
		0)
	return LONG(ret1)
}

func RegConnectRegistryEx(lpMachineName string, hKey HKEY, flags ULONG, phkResult *HKEY) LONG {
	lpMachineNameStr := unicode16FromString(lpMachineName)
	ret1 := syscall6(regConnectRegistryEx, 4,
		uintptr(unsafe.Pointer(&lpMachineNameStr[0])),
		uintptr(hKey),
		uintptr(flags),
		uintptr(unsafe.Pointer(phkResult)),
		0,
		0)
	return LONG(ret1)
}

func RegConnectRegistry(lpMachineName string, hKey HKEY, phkResult *HKEY) LONG {
	lpMachineNameStr := unicode16FromString(lpMachineName)
	ret1 := syscall3(regConnectRegistry, 3,
		uintptr(unsafe.Pointer(&lpMachineNameStr[0])),
		uintptr(hKey),
		uintptr(unsafe.Pointer(phkResult)))
	return LONG(ret1)
}

func RegCopyTree(hKeySrc HKEY, lpSubKey string, hKeyDest HKEY) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall3(regCopyTree, 3,
		uintptr(hKeySrc),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(hKeyDest))
	return LONG(ret1)
}

func RegCreateKeyEx(hKey HKEY, lpSubKey string, reserved uint32, lpClass LPWSTR, dwOptions uint32, samDesired REGSAM, lpSecurityAttributes *SECURITY_ATTRIBUTES, phkResult *HKEY, lpdwDisposition *uint32) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall9(regCreateKeyEx, 9,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(reserved),
		uintptr(unsafe.Pointer(lpClass)),
		uintptr(dwOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(unsafe.Pointer(phkResult)),
		uintptr(unsafe.Pointer(lpdwDisposition)))
	return LONG(ret1)
}

func RegCreateKeyTransacted(hKey HKEY, lpSubKey string, reserved uint32, lpClass LPWSTR, dwOptions uint32, samDesired REGSAM, lpSecurityAttributes /*const*/ *SECURITY_ATTRIBUTES, phkResult *HKEY, lpdwDisposition *uint32, hTransaction HANDLE, pExtendedParemeter uintptr) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall12(regCreateKeyTransacted, 11,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(reserved),
		uintptr(unsafe.Pointer(lpClass)),
		uintptr(dwOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(unsafe.Pointer(phkResult)),
		uintptr(unsafe.Pointer(lpdwDisposition)),
		uintptr(hTransaction),
		pExtendedParemeter,
		0)
	return LONG(ret1)
}

func RegCreateKey(hKey HKEY, lpSubKey string, phkResult *HKEY) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall3(regCreateKey, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(phkResult)))
	return LONG(ret1)
}

func RegDeleteKeyEx(hKey HKEY, lpSubKey string, samDesired REGSAM, reserved uint32) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall6(regDeleteKeyEx, 4,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(samDesired),
		uintptr(reserved),
		0,
		0)
	return LONG(ret1)
}

func RegDeleteKeyTransacted(hKey HKEY, lpSubKey string, samDesired REGSAM, reserved uint32, hTransaction HANDLE, pExtendedParameter uintptr) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall6(regDeleteKeyTransacted, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(samDesired),
		uintptr(reserved),
		uintptr(hTransaction),
		pExtendedParameter)
	return LONG(ret1)
}

func RegDeleteKeyValue(hKey HKEY, lpSubKey string, lpValueName string) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	lpValueNameStr := unicode16FromString(lpValueName)
	ret1 := syscall3(regDeleteKeyValue, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpValueNameStr[0])))
	return LONG(ret1)
}

func RegDeleteKey(hKey HKEY, lpSubKey string) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall3(regDeleteKey, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		0)
	return LONG(ret1)
}

func RegDeleteTree(hKey HKEY, lpSubKey string) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall3(regDeleteTree, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		0)
	return LONG(ret1)
}

func RegDeleteValue(hKey HKEY, lpValueName string) LONG {
	lpValueNameStr := unicode16FromString(lpValueName)
	ret1 := syscall3(regDeleteValue, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpValueNameStr[0])),
		0)
	return LONG(ret1)
}

func RegDisablePredefinedCache() LONG {
	ret1 := syscall3(regDisablePredefinedCache, 0,
		0,
		0,
		0)
	return LONG(ret1)
}

func RegDisablePredefinedCacheEx() LONG {
	ret1 := syscall3(regDisablePredefinedCacheEx, 0,
		0,
		0,
		0)
	return LONG(ret1)
}

func RegDisableReflectionKey(hBase HKEY) LONG {
	ret1 := syscall3(regDisableReflectionKey, 1,
		uintptr(hBase),
		0,
		0)
	return LONG(ret1)
}

func RegEnableReflectionKey(hBase HKEY) LONG {
	ret1 := syscall3(regEnableReflectionKey, 1,
		uintptr(hBase),
		0,
		0)
	return LONG(ret1)
}

// TODO: Unknown type(s): PFILETIME
// func RegEnumKeyEx(hKey HKEY, dwIndex uint32, lpName LPWSTR, lpcchName *uint32, lpReserved *uint32, lpClass LPWSTR, lpcchClass *uint32, lpftLastWriteTime PFILETIME) LONG

func RegEnumKey(hKey HKEY, dwIndex uint32, lpName LPWSTR, cchName uint32) LONG {
	ret1 := syscall6(regEnumKey, 4,
		uintptr(hKey),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(cchName),
		0,
		0)
	return LONG(ret1)
}

func RegEnumValue(hKey HKEY, dwIndex uint32, lpValueName LPWSTR, lpcchValueName *uint32, lpReserved *uint32, lpType *uint32, lpData *byte, lpcbData *uint32) LONG {
	ret1 := syscall9(regEnumValue, 8,
		uintptr(hKey),
		uintptr(dwIndex),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(unsafe.Pointer(lpcchValueName)),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)),
		0)
	return LONG(ret1)
}

func RegFlushKey(hKey HKEY) LONG {
	ret1 := syscall3(regFlushKey, 1,
		uintptr(hKey),
		0,
		0)
	return LONG(ret1)
}

func RegGetKeySecurity(hKey HKEY, securityInformation SECURITY_INFORMATION, pSecurityDescriptor PSECURITY_DESCRIPTOR, lpcbSecurityDescriptor *uint32) LONG {
	ret1 := syscall6(regGetKeySecurity, 4,
		uintptr(hKey),
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(unsafe.Pointer(lpcbSecurityDescriptor)),
		0,
		0)
	return LONG(ret1)
}

func RegGetValue(hkey HKEY, lpSubKey string, lpValue string, dwFlags uint32, pdwType *uint32, pvData uintptr, pcbData *uint32) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	lpValueStr := unicode16FromString(lpValue)
	ret1 := syscall9(regGetValue, 7,
		uintptr(hkey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpValueStr[0])),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pdwType)),
		pvData,
		uintptr(unsafe.Pointer(pcbData)),
		0,
		0)
	return LONG(ret1)
}

func RegLoadAppKey(lpFile string, phkResult *HKEY, samDesired REGSAM, dwOptions uint32, reserved uint32) LONG {
	lpFileStr := unicode16FromString(lpFile)
	ret1 := syscall6(regLoadAppKey, 5,
		uintptr(unsafe.Pointer(&lpFileStr[0])),
		uintptr(unsafe.Pointer(phkResult)),
		uintptr(samDesired),
		uintptr(dwOptions),
		uintptr(reserved),
		0)
	return LONG(ret1)
}

func RegLoadKey(hKey HKEY, lpSubKey string, lpFile string) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	lpFileStr := unicode16FromString(lpFile)
	ret1 := syscall3(regLoadKey, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpFileStr[0])))
	return LONG(ret1)
}

func RegLoadMUIString(hKey HKEY, pszValue string, pszOutBuf LPWSTR, cbOutBuf uint32, pcbData *uint32, flags uint32, pszDirectory string) LONG {
	pszValueStr := unicode16FromString(pszValue)
	pszDirectoryStr := unicode16FromString(pszDirectory)
	ret1 := syscall9(regLoadMUIString, 7,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&pszValueStr[0])),
		uintptr(unsafe.Pointer(pszOutBuf)),
		uintptr(cbOutBuf),
		uintptr(unsafe.Pointer(pcbData)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&pszDirectoryStr[0])),
		0,
		0)
	return LONG(ret1)
}

func RegNotifyChangeKeyValue(hKey HKEY, bWatchSubtree bool, dwNotifyFilter uint32, hEvent HANDLE, fAsynchronous bool) LONG {
	ret1 := syscall6(regNotifyChangeKeyValue, 5,
		uintptr(hKey),
		getUintptrFromBool(bWatchSubtree),
		uintptr(dwNotifyFilter),
		uintptr(hEvent),
		getUintptrFromBool(fAsynchronous),
		0)
	return LONG(ret1)
}

func RegOpenCurrentUser(samDesired REGSAM, phkResult *HKEY) LONG {
	ret1 := syscall3(regOpenCurrentUser, 2,
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		0)
	return LONG(ret1)
}

func RegOpenKeyEx(hKey HKEY, lpSubKey string, ulOptions uint32, samDesired REGSAM, phkResult *HKEY) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall6(regOpenKeyEx, 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(ulOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		0)
	return LONG(ret1)
}

func RegOpenKeyTransacted(hKey HKEY, lpSubKey string, ulOptions uint32, samDesired REGSAM, phkResult *HKEY, hTransaction HANDLE, pExtendedParameter uintptr) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall9(regOpenKeyTransacted, 7,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(ulOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		uintptr(hTransaction),
		pExtendedParameter,
		0,
		0)
	return LONG(ret1)
}

func RegOpenKey(hKey HKEY, lpSubKey string, phkResult *HKEY) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall3(regOpenKey, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(phkResult)))
	return LONG(ret1)
}

func RegOpenUserClassesRoot(hToken HANDLE, dwOptions uint32, samDesired REGSAM, phkResult *HKEY) LONG {
	ret1 := syscall6(regOpenUserClassesRoot, 4,
		uintptr(hToken),
		uintptr(dwOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		0,
		0)
	return LONG(ret1)
}

func RegOverridePredefKey(hKey HKEY, hNewHKey HKEY) LONG {
	ret1 := syscall3(regOverridePredefKey, 2,
		uintptr(hKey),
		uintptr(hNewHKey),
		0)
	return LONG(ret1)
}

// TODO: Unknown type(s): PFILETIME
// func RegQueryInfoKey(hKey HKEY, lpClass LPWSTR, lpcchClass *uint32, lpReserved *uint32, lpcSubKeys *uint32, lpcbMaxSubKeyLen *uint32, lpcbMaxClassLen *uint32, lpcValues *uint32, lpcbMaxValueNameLen *uint32, lpcbMaxValueLen *uint32, lpcbSecurityDescriptor *uint32, lpftLastWriteTime PFILETIME) LONG

// TODO: Unknown type(s): PVALENTW
// func RegQueryMultipleValues(hKey HKEY, val_list PVALENTW, num_vals uint32, lpValueBuf LPWSTR, ldwTotsize *uint32) LONG

func RegQueryReflectionKey(hBase HKEY, bIsReflectionDisabled *BOOL) LONG {
	ret1 := syscall3(regQueryReflectionKey, 2,
		uintptr(hBase),
		uintptr(unsafe.Pointer(bIsReflectionDisabled)),
		0)
	return LONG(ret1)
}

func RegQueryValueEx(hKey HKEY, lpValueName string, lpReserved *uint32, lpType *uint32, lpData *byte, lpcbData *uint32) LONG {
	lpValueNameStr := unicode16FromString(lpValueName)
	ret1 := syscall6(regQueryValueEx, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpValueNameStr[0])),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)))
	return LONG(ret1)
}

func RegQueryValue(hKey HKEY, lpSubKey string, lpData LPWSTR, lpcbData *int32) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall6(regQueryValue, 4,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)),
		0,
		0)
	return LONG(ret1)
}

func RegReplaceKey(hKey HKEY, lpSubKey string, lpNewFile string, lpOldFile string) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	lpNewFileStr := unicode16FromString(lpNewFile)
	lpOldFileStr := unicode16FromString(lpOldFile)
	ret1 := syscall6(regReplaceKey, 4,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(unsafe.Pointer(&lpNewFileStr[0])),
		uintptr(unsafe.Pointer(&lpOldFileStr[0])),
		0,
		0)
	return LONG(ret1)
}

func RegRestoreKey(hKey HKEY, lpFile string, dwFlags uint32) LONG {
	lpFileStr := unicode16FromString(lpFile)
	ret1 := syscall3(regRestoreKey, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpFileStr[0])),
		uintptr(dwFlags))
	return LONG(ret1)
}

func RegSaveKeyEx(hKey HKEY, lpFile string, lpSecurityAttributes *SECURITY_ATTRIBUTES, flags uint32) LONG {
	lpFileStr := unicode16FromString(lpFile)
	ret1 := syscall6(regSaveKeyEx, 4,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpFileStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(flags),
		0,
		0)
	return LONG(ret1)
}

func RegSaveKey(hKey HKEY, lpFile string, lpSecurityAttributes *SECURITY_ATTRIBUTES) LONG {
	lpFileStr := unicode16FromString(lpFile)
	ret1 := syscall3(regSaveKey, 3,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpFileStr[0])),
		uintptr(unsafe.Pointer(lpSecurityAttributes)))
	return LONG(ret1)
}

func RegSetKeySecurity(hKey HKEY, securityInformation SECURITY_INFORMATION, pSecurityDescriptor PSECURITY_DESCRIPTOR) LONG {
	ret1 := syscall3(regSetKeySecurity, 3,
		uintptr(hKey),
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)))
	return LONG(ret1)
}

func RegSetKeyValue(hKey HKEY, lpSubKey /*const*/ LPCSTR, lpValueName /*const*/ LPCSTR, dwType uint32, lpData /*const*/ uintptr, cbData uint32) LONG {
	ret1 := syscall6(regSetKeyValue, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpSubKey)),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(dwType),
		lpData,
		uintptr(cbData))
	return LONG(ret1)
}

func RegSetValueEx(hKey HKEY, lpValueName string, reserved uint32, dwType uint32, lpData /*const*/ *byte, cbData uint32) LONG {
	lpValueNameStr := unicode16FromString(lpValueName)
	ret1 := syscall6(regSetValueEx, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpValueNameStr[0])),
		uintptr(reserved),
		uintptr(dwType),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(cbData))
	return LONG(ret1)
}

func RegSetValue(hKey HKEY, lpSubKey string, dwType uint32, lpData string, cbData uint32) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	lpDataStr := unicode16FromString(lpData)
	ret1 := syscall6(regSetValue, 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		uintptr(dwType),
		uintptr(unsafe.Pointer(&lpDataStr[0])),
		uintptr(cbData),
		0)
	return LONG(ret1)
}

func RegUnLoadKey(hKey HKEY, lpSubKey string) LONG {
	lpSubKeyStr := unicode16FromString(lpSubKey)
	ret1 := syscall3(regUnLoadKey, 2,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&lpSubKeyStr[0])),
		0)
	return LONG(ret1)
}

func RegisterEventSource(lpUNCServerName string, lpSourceName string) HANDLE {
	lpUNCServerNameStr := unicode16FromString(lpUNCServerName)
	lpSourceNameStr := unicode16FromString(lpSourceName)
	ret1 := syscall3(registerEventSource, 2,
		uintptr(unsafe.Pointer(&lpUNCServerNameStr[0])),
		uintptr(unsafe.Pointer(&lpSourceNameStr[0])),
		0)
	return HANDLE(ret1)
}

func RegisterServiceCtrlHandlerEx(lpServiceName string, lpHandlerProc HANDLER_FUNCTION_EX, lpContext LPVOID) SERVICE_STATUS_HANDLE {
	lpServiceNameStr := unicode16FromString(lpServiceName)
	lpHandlerProcCallback := syscall.NewCallback(func(dwControlRawArg uint32, dwEventTypeRawArg uint32, lpEventDataRawArg uintptr, lpContextRawArg uintptr) uintptr {
		ret := lpHandlerProc(dwControlRawArg, dwEventTypeRawArg, lpEventDataRawArg, lpContextRawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(registerServiceCtrlHandlerEx, 3,
		uintptr(unsafe.Pointer(&lpServiceNameStr[0])),
		lpHandlerProcCallback,
		uintptr(unsafe.Pointer(lpContext)))
	return SERVICE_STATUS_HANDLE(ret1)
}

// TODO: Unknown type(s): LPHANDLER_FUNCTION
// func RegisterServiceCtrlHandler(lpServiceName string, lpHandlerProc LPHANDLER_FUNCTION) SERVICE_STATUS_HANDLE

// TODO: Unknown type(s): PCOGETACTIVATIONSTATE, PCOGETCALLSTATE
// func RegisterWaitChainCOMCallback(callStateCallback PCOGETCALLSTATE, activationStateCallback PCOGETACTIVATIONSTATE)

// TODO: Unknown type(s): PENCRYPTION_CERTIFICATE_HASH_LIST
// func RemoveUsersFromEncryptedFile(lpFileName string, pHashes PENCRYPTION_CERTIFICATE_HASH_LIST) uint32

func ReportEvent(hEventLog HANDLE, wType uint16, wCategory uint16, dwEventID uint32, lpUserSid PSID, wNumStrings uint16, dwDataSize uint32, lpStrings *LPCWSTR, lpRawData LPVOID) bool {
	ret1 := syscall9(reportEvent, 9,
		uintptr(hEventLog),
		uintptr(wType),
		uintptr(wCategory),
		uintptr(dwEventID),
		uintptr(lpUserSid),
		uintptr(wNumStrings),
		uintptr(dwDataSize),
		uintptr(unsafe.Pointer(lpStrings)),
		uintptr(unsafe.Pointer(lpRawData)))
	return ret1 != 0
}

func RevertToSelf() bool {
	ret1 := syscall3(revertToSelf, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func SaferCloseLevel(hLevelHandle SAFER_LEVEL_HANDLE) bool {
	ret1 := syscall3(saferCloseLevel, 1,
		uintptr(hLevelHandle),
		0,
		0)
	return ret1 != 0
}

func SaferComputeTokenFromLevel(levelHandle SAFER_LEVEL_HANDLE, inAccessToken HANDLE, outAccessToken *HANDLE, dwFlags uint32, lpReserved LPVOID) bool {
	ret1 := syscall6(saferComputeTokenFromLevel, 5,
		uintptr(levelHandle),
		uintptr(inAccessToken),
		uintptr(unsafe.Pointer(outAccessToken)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpReserved)),
		0)
	return ret1 != 0
}

func SaferCreateLevel(dwScopeId uint32, dwLevelId uint32, openFlags uint32, pLevelHandle *SAFER_LEVEL_HANDLE, lpReserved LPVOID) bool {
	ret1 := syscall6(saferCreateLevel, 5,
		uintptr(dwScopeId),
		uintptr(dwLevelId),
		uintptr(openFlags),
		uintptr(unsafe.Pointer(pLevelHandle)),
		uintptr(unsafe.Pointer(lpReserved)),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): SAFER_OBJECT_INFO_CLASS
// func SaferGetLevelInformation(levelHandle SAFER_LEVEL_HANDLE, dwInfoType SAFER_OBJECT_INFO_CLASS, lpQueryBuffer LPVOID, dwInBufferSize uint32, lpdwOutBufferSize *uint32) bool

// TODO: Unknown type(s): SAFER_POLICY_INFO_CLASS
// func SaferGetPolicyInformation(dwScopeId uint32, saferPolicyInfoClass SAFER_POLICY_INFO_CLASS, infoBufferSize uint32, infoBuffer uintptr, infoBufferRetSize *DWORD, lpReserved LPVOID) bool

// TODO: Unknown type(s): PSAFER_CODE_PROPERTIES
// func SaferIdentifyLevel(dwNumProperties uint32, pCodeProperties PSAFER_CODE_PROPERTIES, pLevelHandle *SAFER_LEVEL_HANDLE, lpReserved LPVOID) bool

func SaferRecordEventLogEntry(hLevel SAFER_LEVEL_HANDLE, szTargetPath string, lpReserved LPVOID) bool {
	szTargetPathStr := unicode16FromString(szTargetPath)
	ret1 := syscall3(saferRecordEventLogEntry, 3,
		uintptr(hLevel),
		uintptr(unsafe.Pointer(&szTargetPathStr[0])),
		uintptr(unsafe.Pointer(lpReserved)))
	return ret1 != 0
}

// TODO: Unknown type(s): SAFER_OBJECT_INFO_CLASS
// func SaferSetLevelInformation(levelHandle SAFER_LEVEL_HANDLE, dwInfoType SAFER_OBJECT_INFO_CLASS, lpQueryBuffer LPVOID, dwInBufferSize uint32) bool

// TODO: Unknown type(s): SAFER_POLICY_INFO_CLASS
// func SaferSetPolicyInformation(dwScopeId uint32, saferPolicyInfoClass SAFER_POLICY_INFO_CLASS, infoBufferSize uint32, infoBuffer uintptr, lpReserved LPVOID) bool

func SaferiIsExecutableFileType(szFullPathname string, bFromShellExecute BOOLEAN) bool {
	szFullPathnameStr := unicode16FromString(szFullPathname)
	ret1 := syscall3(saferiIsExecutableFileType, 2,
		uintptr(unsafe.Pointer(&szFullPathnameStr[0])),
		uintptr(bFromShellExecute),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): ACL_INFORMATION_CLASS
// func SetAclInformation(pAcl *ACL, pAclInformation LPVOID, nAclInformationLength uint32, dwAclInformationClass ACL_INFORMATION_CLASS) bool

// TODO: Unknown type(s): PACL *
// func SetEntriesInAcl(cCountOfExplicitEntries ULONG, pListOfExplicitEntries *EXPLICIT_ACCESS, oldAcl *ACL, newAcl PACL *) uint32

func SetFileSecurity(lpFileName string, securityInformation SECURITY_INFORMATION, pSecurityDescriptor PSECURITY_DESCRIPTOR) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(setFileSecurity, 3,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)))
	return ret1 != 0
}

func SetKernelObjectSecurity(handle HANDLE, securityInformation SECURITY_INFORMATION, securityDescriptor PSECURITY_DESCRIPTOR) bool {
	ret1 := syscall3(setKernelObjectSecurity, 3,
		uintptr(handle),
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(securityDescriptor)))
	return ret1 != 0
}

func SetNamedSecurityInfo(pObjectName LPWSTR, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, psidOwner PSID, psidGroup PSID, pDacl *ACL, pSacl *ACL) uint32 {
	ret1 := syscall9(setNamedSecurityInfo, 7,
		uintptr(unsafe.Pointer(pObjectName)),
		uintptr(objectType),
		uintptr(securityInfo),
		uintptr(psidOwner),
		uintptr(psidGroup),
		uintptr(unsafe.Pointer(pDacl)),
		uintptr(unsafe.Pointer(pSacl)),
		0,
		0)
	return uint32(ret1)
}

func SetPrivateObjectSecurity(securityInformation SECURITY_INFORMATION, modificationDescriptor PSECURITY_DESCRIPTOR, objectsSecurityDescriptor *PSECURITY_DESCRIPTOR, genericMapping *GENERIC_MAPPING, token HANDLE) bool {
	ret1 := syscall6(setPrivateObjectSecurity, 5,
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(modificationDescriptor)),
		uintptr(unsafe.Pointer(objectsSecurityDescriptor)),
		uintptr(unsafe.Pointer(genericMapping)),
		uintptr(token),
		0)
	return ret1 != 0
}

func SetPrivateObjectSecurityEx(securityInformation SECURITY_INFORMATION, modificationDescriptor PSECURITY_DESCRIPTOR, objectsSecurityDescriptor *PSECURITY_DESCRIPTOR, autoInheritFlags ULONG, genericMapping *GENERIC_MAPPING, token HANDLE) bool {
	ret1 := syscall6(setPrivateObjectSecurityEx, 6,
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(modificationDescriptor)),
		uintptr(unsafe.Pointer(objectsSecurityDescriptor)),
		uintptr(autoInheritFlags),
		uintptr(unsafe.Pointer(genericMapping)),
		uintptr(token))
	return ret1 != 0
}

func SetSecurityAccessMask(securityInformation SECURITY_INFORMATION, desiredAccess *uint32) {
	syscall3(setSecurityAccessMask, 2,
		uintptr(securityInformation),
		uintptr(unsafe.Pointer(desiredAccess)),
		0)
}

func SetSecurityDescriptorControl(pSecurityDescriptor PSECURITY_DESCRIPTOR, controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) bool {
	ret1 := syscall3(setSecurityDescriptorControl, 3,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(controlBitsOfInterest),
		uintptr(controlBitsToSet))
	return ret1 != 0
}

func SetSecurityDescriptorDacl(pSecurityDescriptor PSECURITY_DESCRIPTOR, bDaclPresent bool, pDacl *ACL, bDaclDefaulted bool) bool {
	ret1 := syscall6(setSecurityDescriptorDacl, 4,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		getUintptrFromBool(bDaclPresent),
		uintptr(unsafe.Pointer(pDacl)),
		getUintptrFromBool(bDaclDefaulted),
		0,
		0)
	return ret1 != 0
}

func SetSecurityDescriptorGroup(pSecurityDescriptor PSECURITY_DESCRIPTOR, pGroup PSID, bGroupDefaulted bool) bool {
	ret1 := syscall3(setSecurityDescriptorGroup, 3,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(pGroup),
		getUintptrFromBool(bGroupDefaulted))
	return ret1 != 0
}

func SetSecurityDescriptorOwner(pSecurityDescriptor PSECURITY_DESCRIPTOR, pOwner PSID, bOwnerDefaulted bool) bool {
	ret1 := syscall3(setSecurityDescriptorOwner, 3,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(pOwner),
		getUintptrFromBool(bOwnerDefaulted))
	return ret1 != 0
}

// TODO: Unknown type(s): PUCHAR
// func SetSecurityDescriptorRMControl(securityDescriptor PSECURITY_DESCRIPTOR, rMControl PUCHAR) uint32

func SetSecurityDescriptorSacl(pSecurityDescriptor PSECURITY_DESCRIPTOR, bSaclPresent bool, pSacl *ACL, bSaclDefaulted bool) bool {
	ret1 := syscall6(setSecurityDescriptorSacl, 4,
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		getUintptrFromBool(bSaclPresent),
		uintptr(unsafe.Pointer(pSacl)),
		getUintptrFromBool(bSaclDefaulted),
		0,
		0)
	return ret1 != 0
}

func SetSecurityInfo(handle HANDLE, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, psidOwner PSID, psidGroup PSID, pDacl *ACL, pSacl *ACL) uint32 {
	ret1 := syscall9(setSecurityInfo, 7,
		uintptr(handle),
		uintptr(objectType),
		uintptr(securityInfo),
		uintptr(psidOwner),
		uintptr(psidGroup),
		uintptr(unsafe.Pointer(pDacl)),
		uintptr(unsafe.Pointer(pSacl)),
		0,
		0)
	return uint32(ret1)
}

func SetServiceBits(hServiceStatus SERVICE_STATUS_HANDLE, dwServiceBits uint32, bSetBitsOn bool, bUpdateImmediately bool) bool {
	ret1 := syscall6(setServiceBits, 4,
		uintptr(hServiceStatus),
		uintptr(dwServiceBits),
		getUintptrFromBool(bSetBitsOn),
		getUintptrFromBool(bUpdateImmediately),
		0,
		0)
	return ret1 != 0
}

func SetServiceObjectSecurity(hService SC_HANDLE, dwSecurityInformation SECURITY_INFORMATION, lpSecurityDescriptor PSECURITY_DESCRIPTOR) bool {
	ret1 := syscall3(setServiceObjectSecurity, 3,
		uintptr(hService),
		uintptr(dwSecurityInformation),
		uintptr(unsafe.Pointer(lpSecurityDescriptor)))
	return ret1 != 0
}

func SetServiceStatus(hServiceStatus SERVICE_STATUS_HANDLE, lpServiceStatus *SERVICE_STATUS) bool {
	ret1 := syscall3(setServiceStatus, 2,
		uintptr(hServiceStatus),
		uintptr(unsafe.Pointer(lpServiceStatus)),
		0)
	return ret1 != 0
}

func SetThreadToken(thread *HANDLE, token HANDLE) bool {
	ret1 := syscall3(setThreadToken, 2,
		uintptr(unsafe.Pointer(thread)),
		uintptr(token),
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): TOKEN_INFORMATION_CLASS
// func SetTokenInformation(tokenHandle HANDLE, tokenInformationClass TOKEN_INFORMATION_CLASS, tokenInformation LPVOID, tokenInformationLength uint32) bool

func SetUserFileEncryptionKey(pEncryptionCertificate PENCRYPTION_CERTIFICATE) uint32 {
	ret1 := syscall3(setUserFileEncryptionKey, 1,
		uintptr(unsafe.Pointer(pEncryptionCertificate)),
		0,
		0)
	return uint32(ret1)
}

// TODO: Unknown type(s): CONST SERVICE_TABLE_ENTRYW *
// func StartServiceCtrlDispatcher(lpServiceStartTable /*const*/ CONST SERVICE_TABLE_ENTRYW *) bool

func StartService(hService SC_HANDLE, dwNumServiceArgs uint32, lpServiceArgVectors *LPCWSTR) bool {
	ret1 := syscall3(startService, 3,
		uintptr(hService),
		uintptr(dwNumServiceArgs),
		uintptr(unsafe.Pointer(lpServiceArgVectors)))
	return ret1 != 0
}

// TODO: Unknown type(s): FN_PROGRESS, PROG_INVOKE_SETTING
// func TreeResetNamedSecurityInfo(pObjectName LPWSTR, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, pOwner PSID, pGroup PSID, pDacl *ACL, pSacl *ACL, keepExplicit bool, fnProgress FN_PROGRESS, progressInvokeSetting PROG_INVOKE_SETTING, args uintptr) uint32

// TODO: Unknown type(s): FN_PROGRESS, PROG_INVOKE_SETTING
// func TreeSetNamedSecurityInfo(pObjectName LPWSTR, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, pOwner PSID, pGroup PSID, pDacl *ACL, pSacl *ACL, dwAction uint32, fnProgress FN_PROGRESS, progressInvokeSetting PROG_INVOKE_SETTING, args uintptr) uint32

func UninstallApplication(productCode *WCHAR, dwStatus uint32) uint32 {
	ret1 := syscall3(uninstallApplication, 2,
		uintptr(unsafe.Pointer(productCode)),
		uintptr(dwStatus),
		0)
	return uint32(ret1)
}

func UnlockServiceDatabase(scLock SC_LOCK) bool {
	ret1 := syscall3(unlockServiceDatabase, 1,
		uintptr(scLock),
		0,
		0)
	return ret1 != 0
}

func Wow64Win32ApiEntry(dwFuncNumber uint32, dwFlag uint32, dwRes uint32) LONG {
	ret1 := syscall3(wow64Win32ApiEntry, 3,
		uintptr(dwFuncNumber),
		uintptr(dwFlag),
		uintptr(dwRes))
	return LONG(ret1)
}

// TODO: Unknown type(s): PFE_IMPORT_FUNC
// func WriteEncryptedFileRaw(pfImportCallback PFE_IMPORT_FUNC, pvCallbackContext uintptr, pvContext uintptr) uint32

// TODO: Unknown type(s): PSHA_CTX
// func A_SHAFinal(context PSHA_CTX, result *uint32)

// TODO: Unknown type(s): PSHA_CTX
// func A_SHAInit(context PSHA_CTX)

// TODO: Unknown type(s): PSHA_CTX
// func A_SHAUpdate(context PSHA_CTX, buffer /*const*/ *byte, bufferSize uint32)

// TODO: Unknown type(s): TRACEHANDLE
// func CloseTrace(handle TRACEHANDLE) ULONG

// TODO: Unknown type(s): PEVENT_TRACE_PROPERTIES, TRACEHANDLE
// func ControlTrace(hSession TRACEHANDLE, sessionName string, properties PEVENT_TRACE_PROPERTIES, control ULONG) ULONG

// TODO: Unknown type(s): TRACEHANDLE
// func EnableTrace(enable ULONG, flag ULONG, level ULONG, guid /*const*/ *GUID, hSession TRACEHANDLE) ULONG

// TODO: Unknown type(s): PEVENT_FILTER_DESCRIPTOR, TRACEHANDLE
// func EnableTraceEx(provider /*const*/ *GUID, source /*const*/ *GUID, hSession TRACEHANDLE, enable ULONG, level UCHAR, anykeyword ULONGLONG, allkeyword ULONGLONG, enableprop ULONG, filterdesc PEVENT_FILTER_DESCRIPTOR) ULONG

// TODO: Unknown type(s): PENABLE_TRACE_PARAMETERS, TRACEHANDLE
// func EnableTraceEx2(handle TRACEHANDLE, provider /*const*/ *GUID, control ULONG, level UCHAR, match_any ULONGLONG, match_all ULONGLONG, timeout ULONG, params PENABLE_TRACE_PARAMETERS) ULONG

// TODO: Unknown type(s): PTRACE_GUID_PROPERTIES *
// func EnumerateTraceGuids(propertiesarray PTRACE_GUID_PROPERTIES *, arraycount ULONG, guidcount *uint32) ULONG

func EventActivityIdControl(code ULONG, guid *GUID) ULONG {
	ret1 := syscall3(eventActivityIdControl, 2,
		uintptr(code),
		uintptr(unsafe.Pointer(guid)),
		0)
	return ULONG(ret1)
}

// TODO: Unknown type(s): PCEVENT_DESCRIPTOR, REGHANDLE
// func EventEnabled(handle REGHANDLE, descriptor PCEVENT_DESCRIPTOR) BOOLEAN

// TODO: Unknown type(s): REGHANDLE
// func EventProviderEnabled(handle REGHANDLE, level UCHAR, keyword ULONGLONG) BOOLEAN

// TODO: Unknown type(s): PCEVENT_DESCRIPTOR, PEVENT_DATA_DESCRIPTOR, REGHANDLE
// func EventWrite(handle REGHANDLE, descriptor PCEVENT_DESCRIPTOR, count ULONG, data PEVENT_DATA_DESCRIPTOR) ULONG

// TODO: Unknown type(s): PCEVENT_DESCRIPTOR, PEVENT_DATA_DESCRIPTOR, REGHANDLE
// func EventWriteTransfer(handle REGHANDLE, descriptor PCEVENT_DESCRIPTOR, activity /*const*/ *GUID, related /*const*/ *GUID, count ULONG, data PEVENT_DATA_DESCRIPTOR) ULONG

// TODO: Unknown type(s): PEVENT_TRACE_PROPERTIES, TRACEHANDLE
// func FlushTrace(hSession TRACEHANDLE, sessionName string, properties PEVENT_TRACE_PROPERTIES) ULONG

// TODO: Unknown type(s): PACTRL_ACCESSW *, PACTRL_AUDITW *
// func GetNamedSecurityInfoEx(object string, aType SE_OBJECT_TYPE, info SECURITY_INFORMATION, provider string, property string, access_list PACTRL_ACCESSW *, audit_list PACTRL_AUDITW *, owner *LPWSTR, group *LPWSTR) uint32

// TODO: Unknown type(s): PACTRL_ACCESSW *, PACTRL_AUDITW *
// func GetSecurityInfoEx(hObject HANDLE, objectType SE_OBJECT_TYPE, securityInfo SECURITY_INFORMATION, lpProvider string, lpProperty string, ppAccessList PACTRL_ACCESSW *, ppAuditList PACTRL_AUDITW *, lppOwner *LPWSTR, lppGroup *LPWSTR) uint32

// TODO: Unknown type(s): TRACEHANDLE
// func GetTraceEnableFlags(handle TRACEHANDLE) ULONG

// TODO: Unknown type(s): TRACEHANDLE
// func GetTraceEnableLevel(handle TRACEHANDLE) UCHAR

// TODO: Unknown type(s): TRACEHANDLE
// func GetTraceLoggerHandle(buf uintptr) TRACEHANDLE

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_UNICODE_STRING
// func LsaAddAccountRights(policy LSA_HANDLE, sid PSID, rights PLSA_UNICODE_STRING, count ULONG) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, NTSTATUS
// func LsaClose(objectHandle IN LSA_HANDLE) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_HANDLE, PTRUSTED_DOMAIN_AUTH_INFORMATION, PTRUSTED_DOMAIN_INFORMATION_EX
// func LsaCreateTrustedDomainEx(policy LSA_HANDLE, domain_info PTRUSTED_DOMAIN_INFORMATION_EX, auth_info PTRUSTED_DOMAIN_AUTH_INFORMATION, access ACCESS_MASK, domain PLSA_HANDLE) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS
// func LsaDeleteTrustedDomain(policy LSA_HANDLE, sid PSID) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_UNICODE_STRING *
// func LsaEnumerateAccountRights(policy LSA_HANDLE, sid PSID, rights PLSA_UNICODE_STRING *, count *uint32) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_UNICODE_STRING
// func LsaEnumerateAccountsWithUserRight(policy LSA_HANDLE, rights PLSA_UNICODE_STRING, buffer *PVOID, count *uint32) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN PLSA_ENUMERATION_HANDLE, IN ULONG, NTSTATUS, OUT PULONG, OUT PVOID *
// func LsaEnumerateTrustedDomains(policyHandle IN LSA_HANDLE, enumerationContext IN PLSA_ENUMERATION_HANDLE, buffer OUT PVOID *, preferredMaximumLength IN ULONG, countReturned OUT PULONG) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_ENUMERATION_HANDLE
// func LsaEnumerateTrustedDomainsEx(policy LSA_HANDLE, context PLSA_ENUMERATION_HANDLE, buffer *PVOID, length ULONG, count *uint32) NTSTATUS

// TODO: Unknown type(s): IN PVOID, NTSTATUS
// func LsaFreeMemory(buffer IN PVOID) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN PLSA_UNICODE_STRING, IN ULONG, NTSTATUS, OUT PLSA_REFERENCED_DOMAIN_LIST *, OUT PLSA_TRANSLATED_SID *
// func LsaLookupNames(policyHandle IN LSA_HANDLE, count IN ULONG, names IN PLSA_UNICODE_STRING, referencedDomains OUT PLSA_REFERENCED_DOMAIN_LIST *, sids OUT PLSA_TRANSLATED_SID *) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_REFERENCED_DOMAIN_LIST *, PLSA_TRANSLATED_SID2 *, PLSA_UNICODE_STRING
// func LsaLookupNames2(policy LSA_HANDLE, flags ULONG, count ULONG, names PLSA_UNICODE_STRING, domains PLSA_REFERENCED_DOMAIN_LIST *, sids PLSA_TRANSLATED_SID2 *) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, LSA_REFERENCED_DOMAIN_LIST * *, LSA_TRANSLATED_NAME * *, NTSTATUS
// func LsaLookupSids(policyHandle LSA_HANDLE, count ULONG, sids *PSID, referencedDomains LSA_REFERENCED_DOMAIN_LIST * *, names LSA_TRANSLATED_NAME * *) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func LsaNtStatusToWinError(status NTSTATUS) ULONG

// TODO: Unknown type(s): IN ACCESS_MASK, IN OUT PLSA_HANDLE, IN PLSA_OBJECT_ATTRIBUTES, IN PLSA_UNICODE_STRING, NTSTATUS
// func LsaOpenPolicy(systemName IN PLSA_UNICODE_STRING, objectAttributes IN PLSA_OBJECT_ATTRIBUTES, desiredAccess IN ACCESS_MASK, policyHandle IN OUT PLSA_HANDLE) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_HANDLE, PLSA_UNICODE_STRING
// func LsaOpenTrustedDomainByName(policy LSA_HANDLE, name PLSA_UNICODE_STRING, access ACCESS_MASK, handle PLSA_HANDLE) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN POLICY_INFORMATION_CLASS, NTSTATUS, OUT PVOID *
// func LsaQueryInformationPolicy(policyHandle IN LSA_HANDLE, informationClass IN POLICY_INFORMATION_CLASS, buffer OUT PVOID *) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, TRUSTED_INFORMATION_CLASS
// func LsaQueryTrustedDomainInfo(policy LSA_HANDLE, sid PSID, class TRUSTED_INFORMATION_CLASS, buffer *PVOID) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_UNICODE_STRING, TRUSTED_INFORMATION_CLASS
// func LsaQueryTrustedDomainInfoByName(policy LSA_HANDLE, name PLSA_UNICODE_STRING, class TRUSTED_INFORMATION_CLASS, buffer *PVOID) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_UNICODE_STRING
// func LsaRemoveAccountRights(policy LSA_HANDLE, sid PSID, all BOOLEAN, rights PLSA_UNICODE_STRING, count ULONG) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN PLSA_UNICODE_STRING, NTSTATUS, OUT PLSA_UNICODE_STRING *
// func LsaRetrievePrivateData(policyHandle IN LSA_HANDLE, keyName IN PLSA_UNICODE_STRING, privateData OUT PLSA_UNICODE_STRING *) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN POLICY_INFORMATION_CLASS, IN PVOID, NTSTATUS
// func LsaSetInformationPolicy(policyHandle IN LSA_HANDLE, informationClass IN POLICY_INFORMATION_CLASS, buffer IN PVOID) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN PLSA_UNICODE_STRING, NTSTATUS
// func LsaSetSecret(secretHandle IN LSA_HANDLE, encryptedCurrentValue IN PLSA_UNICODE_STRING, encryptedOldValue IN PLSA_UNICODE_STRING) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, PLSA_UNICODE_STRING, TRUSTED_INFORMATION_CLASS
// func LsaSetTrustedDomainInfoByName(policy LSA_HANDLE, name PLSA_UNICODE_STRING, class TRUSTED_INFORMATION_CLASS, buffer uintptr) NTSTATUS

// TODO: Unknown type(s): LSA_HANDLE, NTSTATUS, TRUSTED_INFORMATION_CLASS
// func LsaSetTrustedDomainInformation(policy LSA_HANDLE, sid PSID, class TRUSTED_INFORMATION_CLASS, buffer uintptr) NTSTATUS

// TODO: Unknown type(s): IN LSA_HANDLE, IN PLSA_UNICODE_STRING, NTSTATUS
// func LsaStorePrivateData(policyHandle IN LSA_HANDLE, keyName IN PLSA_UNICODE_STRING, privateData IN PLSA_UNICODE_STRING) NTSTATUS

// TODO: Unknown type(s): MD4_CTX *
// func MD4Final(ctx MD4_CTX *)

// TODO: Unknown type(s): MD4_CTX *
// func MD4Init(ctx MD4_CTX *)

// TODO: Unknown type(s): MD4_CTX *
// func MD4Update(ctx MD4_CTX *, buf /*const*/ *byte, aLen uint32)

// TODO: Unknown type(s): MD5_CTX *
// func MD5Final(ctx MD5_CTX *)

// TODO: Unknown type(s): MD5_CTX *
// func MD5Init(ctx MD5_CTX *)

// TODO: Unknown type(s): PEVENT_TRACE_LOGFILEW, TRACEHANDLE
// func OpenTrace(logfile PEVENT_TRACE_LOGFILEW) TRACEHANDLE

// TODO: Unknown type(s): PTRACEHANDLE
// func ProcessTrace(handleArray PTRACEHANDLE, handleCount ULONG, startTime *FILETIME, endTime *FILETIME) ULONG

// TODO: Unknown type(s): PEVENT_TRACE_PROPERTIES *
// func QueryAllTraces(parray PEVENT_TRACE_PROPERTIES *, arraycount ULONG, psessioncount *uint32) ULONG

// TODO: Unknown type(s): PEVENT_TRACE_PROPERTIES, TRACEHANDLE
// func QueryTrace(handle TRACEHANDLE, sessionname string, properties PEVENT_TRACE_PROPERTIES) ULONG

// TODO: Unknown type(s): PEVENT_TRACE_PROPERTIES, PTRACEHANDLE
// func StartTrace(pSessionHandle PTRACEHANDLE, sessionName string, properties PEVENT_TRACE_PROPERTIES) ULONG

// TODO: Unknown type(s): PEVENT_TRACE_PROPERTIES, TRACEHANDLE
// func StopTrace(session TRACEHANDLE, session_name string, properties PEVENT_TRACE_PROPERTIES) ULONG

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction001(data /*const*/ *byte, key /*const*/ *byte, output *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction002(data /*const*/ *byte, key /*const*/ *byte, output *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction003(key /*const*/ *byte, output *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS, const struct ustring *, struct ustring *
// func SystemFunction004(in /*const*/ const struct ustring *, key /*const*/ const struct ustring *, out struct ustring *) NTSTATUS

// TODO: Unknown type(s): NTSTATUS, const struct ustring *, struct ustring *
// func SystemFunction005(in /*const*/ const struct ustring *, key /*const*/ const struct ustring *, out struct ustring *) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction006(password /*const*/ LPCSTR, hash LPSTR) NTSTATUS

// TODO: Unknown type(s): NTSTATUS, const UNICODE_STRING *
// func SystemFunction007(string /*const*/ const UNICODE_STRING *, hash *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction008(challenge /*const*/ *byte, hash /*const*/ *byte, response *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction009(challenge /*const*/ *byte, hash /*const*/ *byte, response *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction010(unknown LPVOID, data /*const*/ *byte, hash *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction012(in /*const*/ *byte, key /*const*/ *byte, out *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction013(in /*const*/ *byte, key /*const*/ *byte, out *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction024(in /*const*/ *byte, key /*const*/ *byte, out *byte) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction025(in /*const*/ *byte, key /*const*/ *byte, out *byte) NTSTATUS

func SystemFunction030(b1 /*const*/ uintptr, b2 /*const*/ uintptr) bool {
	ret1 := syscall3(systemFunction030, 2,
		b1,
		b2,
		0)
	return ret1 != 0
}

// TODO: Unknown type(s): NTSTATUS, const struct ustring *, struct ustring *
// func SystemFunction032(data struct ustring *, key /*const*/ const struct ustring *) NTSTATUS

func SystemFunction035(lpszDllFilePath /*const*/ LPCSTR) bool {
	ret1 := syscall3(systemFunction035, 1,
		uintptr(unsafe.Pointer(lpszDllFilePath)),
		0,
		0)
	return ret1 != 0
}

func SystemFunction036(pbBuffer uintptr, dwLen ULONG) BOOLEAN {
	ret1 := syscall3(systemFunction036, 2,
		pbBuffer,
		uintptr(dwLen),
		0)
	return BOOLEAN(ret1)
}

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction040(memory uintptr, length ULONG, flags ULONG) NTSTATUS

// TODO: Unknown type(s): NTSTATUS
// func SystemFunction041(memory uintptr, length ULONG, flags ULONG) NTSTATUS

// TODO: Unknown type(s): PEVENT_TRACE_HEADER, TRACEHANDLE
// func TraceEvent(sessionHandle TRACEHANDLE, eventTrace PEVENT_TRACE_HEADER) ULONG

// TODO: Unknown type(s): TRACEHANDLE
// func UnregisterTraceGuids(registrationHandle TRACEHANDLE) ULONG

// TODO: Unknown type(s): WMIHANDLE *
// func WmiOpenBlock(guid *GUID, access ULONG, handle WMIHANDLE *) ULONG
