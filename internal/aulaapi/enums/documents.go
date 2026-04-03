package enums

// DocumentCategoryEnum is the category of a secure document.
type DocumentCategoryEnum string

const (
	DocumentCategoryEnumAgenda           DocumentCategoryEnum = "agenda"
	DocumentCategoryEnumAgendaAllUser    DocumentCategoryEnum = "agendaAllUser"
	DocumentCategoryEnumPlanOfAction     DocumentCategoryEnum = "planOfAction"
	DocumentCategoryEnumNote             DocumentCategoryEnum = "note"
	DocumentCategoryEnumUnknown          DocumentCategoryEnum = "unknown"
)

// DocumentTypeEnum is the type of document.
type DocumentTypeEnum string

const (
	DocumentTypeEnumUnknown      DocumentTypeEnum = "unknown"
	DocumentTypeEnumExternal     DocumentTypeEnum = "external"
	DocumentTypeEnumInternal     DocumentTypeEnum = "internal"
	DocumentTypeEnumNote         DocumentTypeEnum = "note"
	DocumentTypeEnumRichdocument DocumentTypeEnum = "richdocument"
)

// CommonFileSortEnum is the sort field for common files.
type CommonFileSortEnum string

const (
	CommonFileSortEnumTitle       CommonFileSortEnum = "title"
	CommonFileSortEnumUpdatedTime CommonFileSortEnum = "updatedTime"
)

// ImplicitSharingPermissionOverride is the permission override for implicit sharing.
type ImplicitSharingPermissionOverride string

const (
	ImplicitSharingPermissionOverrideRead     ImplicitSharingPermissionOverride = "read"
	ImplicitSharingPermissionOverrideWrite    ImplicitSharingPermissionOverride = "write"
	ImplicitSharingPermissionOverrideNoAccess ImplicitSharingPermissionOverride = "noAccess"
)

// JournalingStatusEnum is the journaling status.
type JournalingStatusEnum string

const (
	JournalingStatusEnumNotProcessed JournalingStatusEnum = "notProcessed"
	JournalingStatusEnumInProgress   JournalingStatusEnum = "inProgress"
	JournalingStatusEnumFailed       JournalingStatusEnum = "failed"
	JournalingStatusEnumCompleted    JournalingStatusEnum = "completed"
)

// RevisionChangeTypeEnum is the type of change in a document revision.
type RevisionChangeTypeEnum string

const (
	RevisionChangeTypeEnumCreated RevisionChangeTypeEnum = "created"
	RevisionChangeTypeEnumEdited  RevisionChangeTypeEnum = "edited"
	RevisionChangeTypeEnumShared  RevisionChangeTypeEnum = "shared"
	RevisionChangeTypeEnumDeleted RevisionChangeTypeEnum = "deleted"
)

// SecureDocumentExportStatus is the export status of a secure document.
type SecureDocumentExportStatus string

const (
	SecureDocumentExportStatusCreated    SecureDocumentExportStatus = "created"
	SecureDocumentExportStatusProcessing SecureDocumentExportStatus = "processing"
	SecureDocumentExportStatusFailed     SecureDocumentExportStatus = "failed"
	SecureDocumentExportStatusCompleted  SecureDocumentExportStatus = "completed"
	SecureDocumentExportStatusUnknown    SecureDocumentExportStatus = "unknown"
)

// SecureDocumentSortEnum is the sort field for secure documents.
type SecureDocumentSortEnum string

const (
	SecureDocumentSortEnumUnknown       SecureDocumentSortEnum = "unknown"
	SecureDocumentSortEnumTitle         SecureDocumentSortEnum = "title"
	SecureDocumentSortEnumUpdatedAtDate SecureDocumentSortEnum = "updatedAtDate"
)

// FileScanningStatus is the file scanning status.
type FileScanningStatus string

const (
	FileScanningStatusAvailable  FileScanningStatus = "available"
	FileScanningStatusBlocked    FileScanningStatus = "blocked"
	FileScanningStatusProcessing FileScanningStatus = "processing"
	FileScanningStatusBypassed   FileScanningStatus = "bypassed"
)

// FileStatusEnum is the file availability status.
type FileStatusEnum string

const (
	FileStatusEnumAvailable   FileStatusEnum = "available"
	FileStatusEnumPending     FileStatusEnum = "pending"
	FileStatusEnumUnavailable FileStatusEnum = "unavailable"
	FileStatusEnumUnknown     FileStatusEnum = "unknown"
)
