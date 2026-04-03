package enums

// Platform is a platform identifier.
type Platform string

const (
	PlatformAndroid Platform = "android"
	PlatformIOS     Platform = "iOS"
	PlatformUnknown Platform = "unknown"
)

// WeekDay is a day of the week.
type WeekDay string

const (
	WeekDayMonday    WeekDay = "monday"
	WeekDayTuesday   WeekDay = "tuesday"
	WeekDayWednesday WeekDay = "wednesday"
	WeekDayThursday  WeekDay = "thursday"
	WeekDayFriday    WeekDay = "friday"
	WeekDaySaturday  WeekDay = "saturday"
	WeekDaySunday    WeekDay = "sunday"
)

// SortOrderEnum is a sort order direction.
type SortOrderEnum string

const (
	SortOrderEnumUnknown    SortOrderEnum = "unknown"
	SortOrderEnumAscending  SortOrderEnum = "ascending"
	SortOrderEnumDescending SortOrderEnum = "descending"
)

// AdditionalMasterDataResponseAnswerTypeEnum is the answer type for additional master data responses.
type AdditionalMasterDataResponseAnswerTypeEnum string

const (
	AdditionalMasterDataResponseAnswerTypeEnumYesNo       AdditionalMasterDataResponseAnswerTypeEnum = "yesNo"
	AdditionalMasterDataResponseAnswerTypeEnumPhoneNumber AdditionalMasterDataResponseAnswerTypeEnum = "phoneNumber"
	AdditionalMasterDataResponseAnswerTypeEnumText        AdditionalMasterDataResponseAnswerTypeEnum = "text"
)

// AppTypeEnum is the application type.
type AppTypeEnum string

const (
	AppTypeEnumStaff   AppTypeEnum = "staff"
	AppTypeEnumPrivate AppTypeEnum = "private"
	AppTypeEnumUnknown AppTypeEnum = "unknown"
)

// AssociationModeEnum is the association mode.
type AssociationModeEnum string

const (
	AssociationModeEnumNone    AssociationModeEnum = "none"
	AssociationModeEnumSelect  AssociationModeEnum = "select"
	AssociationModeEnumConfirm AssociationModeEnum = "confirm"
)

// AulaFilePickerEnum is the source for file picker.
type AulaFilePickerEnum string

const (
	AulaFilePickerEnumFiles                    AulaFilePickerEnum = "files"
	AulaFilePickerEnumMediaLibrary             AulaFilePickerEnum = "mediaLibrary"
	AulaFilePickerEnumGoogleDrive              AulaFilePickerEnum = "googleDrive"
	AulaFilePickerEnumOneDrive                 AulaFilePickerEnum = "oneDrive"
	AulaFilePickerEnumPhotoCamera              AulaFilePickerEnum = "photoCamera"
	AulaFilePickerEnumVideoCamera              AulaFilePickerEnum = "videoCamera"
	AulaFilePickerEnumAulaGallery              AulaFilePickerEnum = "aulaGallery"
	AulaFilePickerEnumDocument                 AulaFilePickerEnum = "document"
	AulaFilePickerEnumDownloadMediaGoogleDrive AulaFilePickerEnum = "downloadMediaGoogleDrive"
	AulaFilePickerEnumDownloadMediaOneDrive    AulaFilePickerEnum = "downloadMediaOneDrive"
	AulaFilePickerEnumFilesForMedia            AulaFilePickerEnum = "filesForMedia"
	AulaFilePickerEnumAttachFileGoogleDrive    AulaFilePickerEnum = "attachFileGoogleDrive"
	AulaFilePickerEnumAttachFileOneDrive       AulaFilePickerEnum = "attachFileOneDrive"
	AulaFilePickerEnumAll                      AulaFilePickerEnum = "all"
)

// CacheType is a cache size classification.
type CacheType string

const (
	CacheTypeSmall CacheType = "SMALL"
	CacheTypeLarge CacheType = "LARGE"
)

// FilterAndSortType is a filter and sort type for lists.
type FilterAndSortType string

const (
	FilterAndSortTypeFilterAll            FilterAndSortType = "filterAll"
	FilterAndSortTypeFilterUnread         FilterAndSortType = "filterUnread"
	FilterAndSortTypeFilterMarked         FilterAndSortType = "filterMarked"
	FilterAndSortTypeFilterDraft          FilterAndSortType = "filterDraft"
	FilterAndSortTypeSortDate             FilterAndSortType = "sortDate"
	FilterAndSortTypeSortSubject          FilterAndSortType = "sortSubject"
	FilterAndSortTypeSortCreatedDate      FilterAndSortType = "sortCreatedDate"
	FilterAndSortTypeSortMediaCreatedDate FilterAndSortType = "sortMediaCreatedDate"
	FilterAndSortTypeSortMediaCreatedAt   FilterAndSortType = "sortMediaCreatedAt"
	FilterAndSortTypeFilterMyAlbums       FilterAndSortType = "filterMyAlbums"
	FilterAndSortTypeFilterMyMedia        FilterAndSortType = "filterMyMedia"
	FilterAndSortTypeSortAlbumName        FilterAndSortType = "sortAlbumName"
)

// LoadingType is a loading/pagination type.
type LoadingType string

const (
	LoadingTypeLoadMore LoadingType = "loadMore"
	LoadingTypeAction   LoadingType = "action"
	LoadingTypeRefresh  LoadingType = "refresh"
)

// LogLevel is a log level.
type LogLevel string

const (
	LogLevelAll     LogLevel = "all"
	LogLevelTrace   LogLevel = "trace"
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warning"
	LogLevelError   LogLevel = "error"
	LogLevelFatal   LogLevel = "fatal"
)

// ReportEnum is a report target type.
type ReportEnum string

const (
	ReportEnumPost     ReportEnum = "post"
	ReportEnumMedia    ReportEnum = "media"
	ReportEnumComments ReportEnum = "comments"
	ReportEnumUnknown  ReportEnum = "unknown"
)

// ResourceType is a resource type for calendar resources.
type ResourceType string

const (
	ResourceTypeLocation      ResourceType = "location"
	ResourceTypeOther         ResourceType = "other"
	ResourceTypeExtraLocation ResourceType = "extraLocation"
	ResourceTypeElectronics   ResourceType = "electronics"
	ResourceTypeStationery    ResourceType = "stationery"
)

// TimePeriod is a time period selection.
type TimePeriod string

const (
	TimePeriodNone        TimePeriod = "none"
	TimePeriodTwoWeeks    TimePeriod = "twoWeeks"
	TimePeriodOneMonth    TimePeriod = "oneMonth"
	TimePeriodThreeMonths TimePeriod = "threeMonths"
	TimePeriodSixMonths   TimePeriod = "sixMonths"
	TimePeriodOneYear     TimePeriod = "oneYear"
)

// BottomBarLongPressOption is a bottom bar long-press option.
type BottomBarLongPressOption string

const (
	BottomBarLongPressOptionEditShortcuts BottomBarLongPressOption = "editShortcuts"
)

// FrontPageSettingConfigurationEnum is a front page configuration setting.
type FrontPageSettingConfigurationEnum string

const (
	FrontPageSettingConfigurationEnumActivityFeed          FrontPageSettingConfigurationEnum = "activityFeed"
	FrontPageSettingConfigurationEnumMessages              FrontPageSettingConfigurationEnum = "messages"
	FrontPageSettingConfigurationEnumCalendarOverview      FrontPageSettingConfigurationEnum = "calendarOverview"
	FrontPageSettingConfigurationEnumImportantDates        FrontPageSettingConfigurationEnum = "importantDates"
	FrontPageSettingConfigurationEnumDocument              FrontPageSettingConfigurationEnum = "document"
	FrontPageSettingConfigurationEnumComeGo                FrontPageSettingConfigurationEnum = "comeGo"
	FrontPageSettingConfigurationEnumGallery               FrontPageSettingConfigurationEnum = "gallery"
	FrontPageSettingConfigurationEnumContactList           FrontPageSettingConfigurationEnum = "contactList"
	FrontPageSettingConfigurationEnumPersonalReferenceData FrontPageSettingConfigurationEnum = "personalReferenceData"
)

// BioAuthStatus is a biometric authentication status.
type BioAuthStatus string

const (
	BioAuthStatusCanTryAgain    BioAuthStatus = "canTryAgain"
	BioAuthStatusCanNotTryAgain BioAuthStatus = "canNotTryAgain"
	BioAuthStatusCanceled       BioAuthStatus = "canceled"
	BioAuthStatusAccepted       BioAuthStatus = "accepted"
)

// BlockedLevel is a blocked communication level.
type BlockedLevel string

const (
	BlockedLevelCentral     BlockedLevel = "central"
	BlockedLevelMunicipal   BlockedLevel = "municipal"
	BlockedLevelInstitution BlockedLevel = "institution"
	BlockedLevelUnknown     BlockedLevel = "unknown"
)

// Consent is a consent type.
type Consent string

const (
	ConsentShareContactInformationParent Consent = "SHARE_CONTACT_INFORMATION_PARENT"
	ConsentShareContactInformationChild  Consent = "SHARE_CONTACT_INFORMATION_CHILD"
	ConsentOthers                        Consent = "OTHERS"
)

// ConsentAnswerEnum is a consent answer.
type ConsentAnswerEnum string

const (
	ConsentAnswerEnumAccepted    ConsentAnswerEnum = "accepted"
	ConsentAnswerEnumDeclined    ConsentAnswerEnum = "declined"
	ConsentAnswerEnumClass       ConsentAnswerEnum = "class"
	ConsentAnswerEnumYear        ConsentAnswerEnum = "year"
	ConsentAnswerEnumInstitution ConsentAnswerEnum = "institution"
	ConsentAnswerEnumNotAtAll    ConsentAnswerEnum = "notAtAll"
	ConsentAnswerEnumOther       ConsentAnswerEnum = "other"
)

// ConsentStatus is a consent status.
type ConsentStatus string

const (
	ConsentStatusActive   ConsentStatus = "active"
	ConsentStatusDeactive ConsentStatus = "deactive"
	ConsentStatusPending  ConsentStatus = "pending"
)

// CommentDropDownEnumeration is a comment dropdown action.
type CommentDropDownEnumeration string

const (
	CommentDropDownEnumerationDelete CommentDropDownEnumeration = "delete"
	CommentDropDownEnumerationEdit   CommentDropDownEnumeration = "edit"
	CommentDropDownEnumerationReport CommentDropDownEnumeration = "report"
)

// CommentType is a comment type.
type CommentType string

const (
	CommentTypeComment CommentType = "comment"
	CommentTypeMedia   CommentType = "media"
	CommentTypePost    CommentType = "post"
	CommentTypeUnknown CommentType = "unknown"
)

// PostDetailMoreMenuEnum is a post detail more-menu action.
type PostDetailMoreMenuEnum string

const (
	PostDetailMoreMenuEnumReportPost PostDetailMoreMenuEnum = "reportPost"
)

// PostFilterTypeEnum is a post filter type.
type PostFilterTypeEnum string

const (
	PostFilterTypeEnumAll          PostFilterTypeEnum = "all"
	PostFilterTypeEnumUnread       PostFilterTypeEnum = "unread"
	PostFilterTypeEnumIsImportant  PostFilterTypeEnum = "isImportant"
	PostFilterTypeEnumFromStaff    PostFilterTypeEnum = "fromStaff"
	PostFilterTypeEnumFromParents  PostFilterTypeEnum = "fromParents"
	PostFilterTypeEnumFromStudents PostFilterTypeEnum = "fromStudents"
	PostFilterTypeEnumOwnPost     PostFilterTypeEnum = "ownPost"
	PostFilterTypeEnumBookmarked   PostFilterTypeEnum = "bookmarked"
)

// WidgetPlacementEnum is a widget placement location.
type WidgetPlacementEnum string

const (
	WidgetPlacementEnumOwnPage        WidgetPlacementEnum = "ownPage"
	WidgetPlacementEnumRightOfOverview WidgetPlacementEnum = "rightOfOverview"
	WidgetPlacementEnumRightOfCalendar WidgetPlacementEnum = "rightOfCalendar"
	WidgetPlacementEnumBelowCalendar   WidgetPlacementEnum = "belowCalendar"
	WidgetPlacementEnumOnOverview      WidgetPlacementEnum = "onOverview"
	WidgetPlacementEnumOnCalendar      WidgetPlacementEnum = "onCalendar"
)

// PermissionEnum is an institution permission.
type PermissionEnum string

const (
	PermissionEnumAdminModule                                       PermissionEnum = "ADMIN_MODULE"
	PermissionEnumSearchAccessProfiles                              PermissionEnum = "SEARCH_ACCESS_PROFILES"
	PermissionEnumSearchAccessGroups                                PermissionEnum = "SEARCH_ACCESS_GROUPS"
	PermissionEnumHandleGroup                                       PermissionEnum = "HANDLE_GROUP"
	PermissionEnumHandleInterinstitutionalGroups                    PermissionEnum = "HANDLE_INTERINSTITUTIONAL_GROUPS"
	PermissionEnumHandleUserRoles                                   PermissionEnum = "HANDLE_USER_ROLES"
	PermissionEnumHandleReportsOfPosts                              PermissionEnum = "HANDLE_REPORTS_OF_POSTS"
	PermissionEnumDeletePostsMediaComments                          PermissionEnum = "DELETE_POSTS_MEDIA_COMMENTS"
	PermissionEnumHandleCallTimes                                   PermissionEnum = "HANDLE_CALL_TIMES"
	PermissionEnumHandleResourcesInstitution                        PermissionEnum = "HANDLE_RESOURCES_INSTITUTION"
	PermissionEnumHandleAdditionalMasterData                        PermissionEnum = "HANDLE_ADDITIONAL_MASTER_DATA"
	PermissionEnumHandlePhysicalLocation                            PermissionEnum = "HANDLE_PHYSICAL_LOCATION"
	PermissionEnumHandleImportantFiles                              PermissionEnum = "HANDLE_IMPORTANT_FILES"
	PermissionEnumHandleSharedInbox                                 PermissionEnum = "HANDLE_SHARED_INBOX"
	PermissionEnumHandleUserData                                    PermissionEnum = "HANDLE_USER_DATA"
	PermissionEnumHandleAllowedRecipients                           PermissionEnum = "HANDLE_ALLOWED_RECIPIENTS"
	PermissionEnumHandleCommunicationChannelsMunicipality           PermissionEnum = "HANDLE_COMMUNICATION_CHANNELS_MUNICIPALITY"
	PermissionEnumHandleAdministrativeAuthority                     PermissionEnum = "HANDLE_ADMINISTRATIVE_AUTHORITY"
	PermissionEnumHandleGroupingsOfInstitutions                     PermissionEnum = "HANDLE_GROUPINGS_OF_INSTITUTIONS"
	PermissionEnumHandleRightsToPhysicalLocation                    PermissionEnum = "HANDLE_RIGHTS_TO_PHYSICAL_LOCATION"
	PermissionEnumHandleResourceCategory                            PermissionEnum = "HANDLE_RESOURCE_CATEGORY"
	PermissionEnumHandleAdditionalMasterDataBruttoList              PermissionEnum = "HANDLE_ADDITIONAL_MASTER_DATA_BRUTTO_LIST"
	PermissionEnumHandleConsents                                    PermissionEnum = "HANDLE_CONSENTS"
	PermissionEnumHandleLessonImportTime                            PermissionEnum = "HANDLE_LESSON_IMPORT_TIME"
	PermissionEnumHandleMaxFilesize                                 PermissionEnum = "HANDLE_MAX_FILESIZE"
	PermissionEnumHandleFileformats                                 PermissionEnum = "HANDLE_FILEFORMATS"
	PermissionEnumAccessSecureFilesharing                           PermissionEnum = "ACCESS_SECURE_FILESHARING"
	PermissionEnumHandleSecureFiles                                 PermissionEnum = "HANDLE_SECURE_FILES"
	PermissionEnumHandleSecureFilesLimited                          PermissionEnum = "HANDLE_SECURE_FILES_LIMITED"
	PermissionEnumAccessImportantFiles                              PermissionEnum = "ACCESS_IMPORTANT_FILES"
	PermissionEnumAccessSharedInbox                                 PermissionEnum = "ACCESS_SHARED_INBOX"
	PermissionEnumReadMessage                                       PermissionEnum = "READ_MESSAGE"
	PermissionEnumWriteMessage                                      PermissionEnum = "WRITE_MESSAGE"
	PermissionEnumReadPost                                          PermissionEnum = "READ_POST"
	PermissionEnumWritePost                                         PermissionEnum = "WRITE_POST"
	PermissionEnumShareMedia                                        PermissionEnum = "SHARE_MEDIA"
	PermissionEnumSeeMedia                                          PermissionEnum = "SEE_MEDIA"
	PermissionEnumHandleGroupApplication                            PermissionEnum = "HANDLE_GROUP_APPLICATION"
	PermissionEnumWriteComments                                     PermissionEnum = "WRITE_COMMENTS"
	PermissionEnumSeeCalendar                                       PermissionEnum = "SEE_CALENDAR"
	PermissionEnumReadEvents                                        PermissionEnum = "READ_EVENTS"
	PermissionEnumHandleEvents                                      PermissionEnum = "HANDLE_EVENTS"
	PermissionEnumInviteGroupToEvent                                PermissionEnum = "INVITE_GROUP_TO_EVENT"
	PermissionEnumHandleParentalMeetingSchool                       PermissionEnum = "HANDLE_PARENTAL_MEETING_SCHOOL"
	PermissionEnumHandlePerformanceMeeting                          PermissionEnum = "HANDLE_PERFORMANCE_MEETING"
	PermissionEnumBookResources                                     PermissionEnum = "BOOK_RESOURCES"
	PermissionEnumInviteToEvent                                     PermissionEnum = "INVITE_TO_EVENT"
	PermissionEnumAnswerEventWithExtendedAnswer                     PermissionEnum = "ANSWER_EVENT_WITH_EXTENDED_ANSWER"
	PermissionEnumShareSecureFiles                                  PermissionEnum = "SHARE_SECURE_FILES"
	PermissionEnumHandleCommunicationChannelsCentral                PermissionEnum = "HANDLE_COMMUNICATION_CHANNELS_CENTRAL"
	PermissionEnumHandleTransitionYear                              PermissionEnum = "HANDLE_TRANSITION_YEAR"
	PermissionEnumHandleResourcesMunicipality                       PermissionEnum = "HANDLE_RESOURCES_MUNICIPALITY"
	PermissionEnumSendSms                                           PermissionEnum = "SEND_SMS"
	PermissionEnumWriteInfoProfile                                  PermissionEnum = "WRITE_INFO_PROFILE"
	PermissionEnumMessageAttachBccRecipients                        PermissionEnum = "MESSAGE_ATTACH_BCC_RECIPIENTS"
	PermissionEnumHandleParentalMeetingDaycare                      PermissionEnum = "HANDLE_PARENTAL_MEETING_DAYCARE"
	PermissionEnumInboxSetPersonalAutoreply                         PermissionEnum = "INBOX_SET_PERSONAL_AUTOREPLY"
	PermissionEnumInboxFolders                                      PermissionEnum = "INBOX_FOLDERS"
	PermissionEnumMessageSeeSubscribersLastread                     PermissionEnum = "MESSAGE_SEE_SUBSCRIBERS_LASTREAD"
	PermissionEnumHandleConsentAge                                  PermissionEnum = "HANDLE_CONSENT_AGE"
	PermissionEnumHandleDashboard                                   PermissionEnum = "HANDLE_DASHBOARD"
	PermissionEnumReportPostsMediaComments                          PermissionEnum = "REPORT_POSTS_MEDIA_COMMENTS"
	PermissionEnumSeeGuardianChildContactInformation                PermissionEnum = "SEE_GUARDIAN_CHILD_CONTACT_INFORMATION"
	PermissionEnumSeeEmployeeContactInformation                     PermissionEnum = "SEE_EMPLOYEE_CONTACT_INFORMATION"
	PermissionEnumSeeGuardianChildLastLogin                         PermissionEnum = "SEE_GUARDIAN_CHILD_LAST_LOGIN"
	PermissionEnumSeeEmployeeLastLogin                              PermissionEnum = "SEE_EMPLOYEE_LAST_LOGIN"
	PermissionEnumHandleContacts                                    PermissionEnum = "HANDLE_CONTACTS"
	PermissionEnumViewUsersAdditionalData                           PermissionEnum = "VIEW_USERS_ADDITIONAL_DATA"
	PermissionEnumHandleSignature                                   PermissionEnum = "HANDLE_SIGNATURE"
	PermissionEnumViewMediaRegardlessOfConsent                      PermissionEnum = "VIEW_MEDIA_REGARDLESS_OF_CONSENT"
	PermissionEnumViewContactInformationRegardlessOfConsent         PermissionEnum = "VIEW_CONTACT_INFORMATION_REGARDLESS_OF_CONSENT"
	PermissionEnumHandleMaxImageResolution                          PermissionEnum = "HANDLE_MAX_IMAGE_RESOLUTION"
	PermissionEnumHandleMaxVideoLength                              PermissionEnum = "HANDLE_MAX_VIDEO_LENGTH"
	PermissionEnumViewUsersConsents                                 PermissionEnum = "VIEW_USERS_CONSENTS"
	PermissionEnumTagOtherUsersOnOtherMedia                         PermissionEnum = "TAG_OTHER_USERS_ON_OTHER_MEDIA"
	PermissionEnumAttachGoogleDriveFile                             PermissionEnum = "ATTACH_GOOGLE_DRIVE_FILE"
	PermissionEnumImportMediaFromGoogleDrive                        PermissionEnum = "IMPORT_MEDIA_FROM_GOOGLE_DRIVE"
	PermissionEnumAttachOnedriveFile                                PermissionEnum = "ATTACH_ONEDRIVE_FILE"
	PermissionEnumImportMediaFromOnedrive                           PermissionEnum = "IMPORT_MEDIA_FROM_ONEDRIVE"
	PermissionEnumSeeContactParentsContactInfo                      PermissionEnum = "SEE_CONTACT_PARENTS_CONTACT_INFO"
	PermissionEnumHandleMedia                                       PermissionEnum = "HANDLE_MEDIA"
	PermissionEnumImpersonateUser                                   PermissionEnum = "IMPERSONATE_USER"
	PermissionEnumViewEmployeesAdditionalData                       PermissionEnum = "VIEW_EMPLOYEES_ADDITIONAL_DATA"
	PermissionEnumHandleServiceMessages                             PermissionEnum = "HANDLE_SERVICE_MESSAGES"
	PermissionEnumPairingInstitutionAndDevice                       PermissionEnum = "PAIRING_INSTITUTION_AND_DEVICE"
	PermissionEnumViewPresenceStatistics                            PermissionEnum = "VIEW_PRESENCE_STATISTICS"
	PermissionEnumHandleVacationRequests                            PermissionEnum = "HANDLE_VACATION_REQUESTS"
	PermissionEnumHandleOptionsPresenceDashboard                    PermissionEnum = "HANDLE_OPTIONS_PRESENCE_DASHBOARD"
	PermissionEnumUseGroupsAsDistributionLists                      PermissionEnum = "USE_GROUPS_AS_DISTRIBUTION_LISTS"
	PermissionEnumViewNameProtection                                PermissionEnum = "VIEW_NAME_PROTECTION"
	PermissionEnumViewCustody                                       PermissionEnum = "VIEW_CUSTODY"
	PermissionEnumAccessSkoleintraArchive                           PermissionEnum = "ACCESS_SKOLEINTRA_ARCHIVE"
	PermissionEnumSkoleintraAdmin                                   PermissionEnum = "SKOLEINTRA_ADMIN"
	PermissionEnumSecureDocumentsAccessAll                          PermissionEnum = "SECURE_DOCUMENTS_ACCESS_ALL"
	PermissionEnumCreateEventsInInstitutionCalendar                 PermissionEnum = "CREATE_EVENTS_IN_INSTITUTION_CALENDAR"
	PermissionEnumCreateEventsOnlyInInstitutionCalendar             PermissionEnum = "CREATE_EVENTS_ONLY_IN_INSTITUTION_CALENDAR"
	PermissionEnumViewPersonalReferenceDataForAllChildrenAndGuardian PermissionEnum = "VIEW_PERSONAL_REFERENCE_DATA_FOR_ALL_CHILDREN_AND_GUARDIAN"
	PermissionEnumViewContactInformationAll                         PermissionEnum = "VIEW_CONTACT_INFORMATION_ALL"
	PermissionEnumHandleGroupTemplates                              PermissionEnum = "HANDLE_GROUP_TEMPLATES"
	PermissionEnumHandleNoticeBoards                                PermissionEnum = "HANDLE_NOTICE_BOARDS"
	PermissionEnumHandleAccessContactInfo                           PermissionEnum = "HANDLE_ACCESS_CONTACT_INFO"
	PermissionEnumHandleDataPolicy                                  PermissionEnum = "HANDLE_DATA_POLICY"
	PermissionEnumHandleServiceRequest                              PermissionEnum = "HANDLE_SERVICE_REQUEST"
	PermissionEnumHandleCalendarFeedMunicipality                    PermissionEnum = "HANDLE_CALENDAR_FEED_MUNICIPALITY"
	PermissionEnumExportPresenceStatistics                          PermissionEnum = "EXPORT_PRESENCE_STATISTICS"
	PermissionEnumExportSecureFiles                                 PermissionEnum = "EXPORT_SECURE_FILES"
	PermissionEnumReadSecureFiles                                   PermissionEnum = "READ_SECURE_FILES"
	PermissionEnumEditPresenceTemplates                             PermissionEnum = "EDIT_PRESENCE_TEMPLATES"
	PermissionEnumHandleEventCoOrganizer                            PermissionEnum = "HANDLE_EVENT_CO_ORGANIZER"
	PermissionEnumHandleOthersEvents                                PermissionEnum = "HANDLE_OTHERS_EVENTS"
	PermissionEnumEditSharedAlbums                                  PermissionEnum = "EDIT_SHARED_ALBUMS"
	PermissionEnumEditSharedMedia                                   PermissionEnum = "EDIT_SHARED_MEDIA"
	PermissionEnumJournalingToEsdh                                  PermissionEnum = "JOURNALING_TO_ESDH"
)

// SearchProfileDocTypeEnum is a search profile document type.
type SearchProfileDocTypeEnum string

const (
	SearchProfileDocTypeEnumProfile     SearchProfileDocTypeEnum = "profile"
	SearchProfileDocTypeEnumGroup       SearchProfileDocTypeEnum = "group"
	SearchProfileDocTypeEnumCommonInbox SearchProfileDocTypeEnum = "commonInbox"
	SearchProfileDocTypeEnumAll         SearchProfileDocTypeEnum = "all"
)

// SearchProfilePortalRoleEnum is a search profile portal role filter.
type SearchProfilePortalRoleEnum string

const (
	SearchProfilePortalRoleEnumEmployee SearchProfilePortalRoleEnum = "employee"
	SearchProfilePortalRoleEnumChild    SearchProfilePortalRoleEnum = "child"
)

// SearchResultItemType is a search result item type.
type SearchResultItemType string

const (
	SearchResultItemTypeNone               SearchResultItemType = "none"
	SearchResultItemTypeGroup              SearchResultItemType = "group"
	SearchResultItemTypeProfile            SearchResultItemType = "profile"
	SearchResultItemTypeChild              SearchResultItemType = "child"
	SearchResultItemTypeEmployee           SearchResultItemType = "employee"
	SearchResultItemTypeGuardian           SearchResultItemType = "guardian"
	SearchResultItemTypeInternalSecureFile SearchResultItemType = "internalSecureFile"
	SearchResultItemTypeExternalSecureFile SearchResultItemType = "externalSecureFile"
	SearchResultItemTypeCommonFile         SearchResultItemType = "commonFile"
	SearchResultItemTypeEvent              SearchResultItemType = "event"
	SearchResultItemTypePost               SearchResultItemType = "post"
	SearchResultItemTypeCommonInbox        SearchResultItemType = "commonInbox"
	SearchResultItemTypeMessage            SearchResultItemType = "message"
	SearchResultItemTypeThread             SearchResultItemType = "thread"
	SearchResultItemTypeThreadMessage      SearchResultItemType = "threadMessage"
	SearchResultItemTypeMedia              SearchResultItemType = "media"
)

// GroupSearchScopeEnum is a group search scope.
type GroupSearchScopeEnum string

const (
	GroupSearchScopeEnumInstitutional          GroupSearchScopeEnum = "institutional"
	GroupSearchScopeEnumAdministrativeAuthority GroupSearchScopeEnum = "administrativeAuthority"
	GroupSearchScopeEnumCrossInstitutional     GroupSearchScopeEnum = "crossInstitutional"
	GroupSearchScopeEnumMunicipal              GroupSearchScopeEnum = "municipal"
)

// SearchRecipientDocTypeEnum is a search recipient document type.
type SearchRecipientDocTypeEnum string

const (
	SearchRecipientDocTypeEnumProfile     SearchRecipientDocTypeEnum = "profile"
	SearchRecipientDocTypeEnumGroup       SearchRecipientDocTypeEnum = "group"
	SearchRecipientDocTypeEnumCommonInbox SearchRecipientDocTypeEnum = "commonInbox"
	SearchRecipientDocTypeEnumAll         SearchRecipientDocTypeEnum = "all"
)

// SearchRecipientMailBoxOwnerType is a search recipient mailbox owner type.
type SearchRecipientMailBoxOwnerType string

const (
	SearchRecipientMailBoxOwnerTypeInstitutionProfile SearchRecipientMailBoxOwnerType = "institutionProfile"
	SearchRecipientMailBoxOwnerTypeCommonInbox        SearchRecipientMailBoxOwnerType = "commonInbox"
	SearchRecipientMailBoxOwnerTypeOtpInbox           SearchRecipientMailBoxOwnerType = "otpInbox"
)

// SearchRecipientModuleEnum is a search recipient module context.
type SearchRecipientModuleEnum string

const (
	SearchRecipientModuleEnumEvent                 SearchRecipientModuleEnum = "event"
	SearchRecipientModuleEnumMessages              SearchRecipientModuleEnum = "messages"
	SearchRecipientModuleEnumOverview              SearchRecipientModuleEnum = "overview"
	SearchRecipientModuleEnumGallery               SearchRecipientModuleEnum = "gallery"
	SearchRecipientModuleEnumSecureDocument         SearchRecipientModuleEnum = "secureDocument"
	SearchRecipientModuleEnumPersonalReferenceData SearchRecipientModuleEnum = "personalReferenceData"
	SearchRecipientModuleEnumContacts              SearchRecipientModuleEnum = "contacts"
)

// SearchRecipientPortalRoleEnum is a search recipient portal role filter.
type SearchRecipientPortalRoleEnum string

const (
	SearchRecipientPortalRoleEnumEmployee SearchRecipientPortalRoleEnum = "employee"
	SearchRecipientPortalRoleEnumChild    SearchRecipientPortalRoleEnum = "child"
	SearchRecipientPortalRoleEnumGuardian SearchRecipientPortalRoleEnum = "guardian"
	SearchRecipientPortalRoleEnumOtp      SearchRecipientPortalRoleEnum = "otp"
	SearchRecipientPortalRoleEnumAll      SearchRecipientPortalRoleEnum = "all"
)

// SearchResourceTypeEnum is a search resource type.
type SearchResourceTypeEnum string

const (
	SearchResourceTypeEnumLocation SearchResourceTypeEnum = "location"
)
