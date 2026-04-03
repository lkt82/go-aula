package models

// SecureDocumentCreatorDto represents a creator of a secure document.
type SecureDocumentCreatorDto struct {
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Alias    bool    `json:"alias"`
	Metadata *string `json:"metadata,omitempty"`
}

// SecureDocumentAssociateInstitutionProfileDto represents an institution profile associated with a secure document.
type SecureDocumentAssociateInstitutionProfileDto struct {
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Alias    bool    `json:"alias"`
	Metadata *string `json:"metadata,omitempty"`
}

// SecureDocumentAssociateGroupDto represents a group associated with a secure document.
type SecureDocumentAssociateGroupDto struct {
	ID              *int64                      `json:"id,omitempty"`
	InstitutionCode *string                     `json:"institutionCode,omitempty"`
	InstitutionName *string                     `json:"institutionName,omitempty"`
	Name            *string                     `json:"name,omitempty"`
	MainGroup       bool                        `json:"mainGroup"`
	MembershipCount *MembershipCountResultModel `json:"membershipCount,omitempty"`
}

// SecureDocumentShareWithGroupDto represents group sharing on a secure document.
type SecureDocumentShareWithGroupDto struct {
	ID                    *int64                      `json:"id,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	CanEdit               bool                        `json:"canEdit"`
	AllowMembersToBeShown bool                        `json:"allowMembersToBeShown"`
	MembershipCount       *MembershipCountResultModel `json:"membershipCount,omitempty"`
}

// SecureDocumentShareWithInstitutionProfileDto represents institution profile sharing on a secure document.
type SecureDocumentShareWithInstitutionProfileDto struct {
	ID              *int64  `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
	CanEdit         bool    `json:"canEdit"`
	Alias           bool    `json:"alias"`
	Metadata        *string `json:"metadata,omitempty"`
	Role            *string `json:"role,omitempty"`
}

// ImplicitSharingProfileDto represents an implicit sharing profile entry.
type ImplicitSharingProfileDto struct {
	SimpleInstitutionProfileDto *SecureDocumentShareWithInstitutionProfileDto `json:"simpleInstitutionProfileDto,omitempty"`
	PermissionOverrideEnum      *string                                       `json:"permissionOverrideEnum,omitempty"`
}

// ImplicitSharingOverride represents an implicit sharing override.
type ImplicitSharingOverride struct {
	InstitutionProfileID   *int64  `json:"institutionProfileId,omitempty"`
	PermissionOverrideEnum *string `json:"permissionOverrideEnum,omitempty"`
}

// SecureDocumentDto represents a core secure document entity.
type SecureDocumentDto struct {
	ID                                *int64                                         `json:"id,omitempty"`
	HasMedia                          bool                                           `json:"hasMedia"`
	CanEdit                           bool                                           `json:"canEdit"`
	CanEditLockedStatus               bool                                           `json:"canEditLockedStatus"`
	IsLocked                          bool                                           `json:"isLocked"`
	JournalingStatus                  *string                                        `json:"journalingStatus,omitempty"`
	Category                          *string                                        `json:"category,omitempty"`
	DocumentTemplateTitle             *string                                        `json:"documentTemplateTitle,omitempty"`
	InstitutionCode                   *string                                        `json:"institutionCode,omitempty"`
	DocumentType                      *string                                        `json:"documentType,omitempty"`
	AssociatedInstitutionProfiles     []SecureDocumentAssociateInstitutionProfileDto  `json:"associatedInstitutionProfiles,omitempty"`
	SharedWithGroups                  []SecureDocumentShareWithGroupDto               `json:"sharedWithGroups,omitempty"`
	SharedWithInstitutionProfiles     []SecureDocumentShareWithInstitutionProfileDto  `json:"sharedWithInstitutionProfiles,omitempty"`
	ImplicitSharings                  []ImplicitSharingProfileDto                    `json:"implicitSharings,omitempty"`
	Creator                           *SecureDocumentCreatorDto                       `json:"creator,omitempty"`
	CreatedAt                         *string                                        `json:"createdAt,omitempty"`
	Title                             *string                                        `json:"title,omitempty"`
	UpdatedAt                         *string                                        `json:"updatedAt,omitempty"`
	UpdatedBy                         *string                                        `json:"updatedBy,omitempty"`
	Version                           *int                                           `json:"version,omitempty"`
	Description                       *string                                        `json:"description,omitempty"`
	IsSharedWithGuardian              bool                                           `json:"isSharedWithGuardian"`
	IsShareable                       bool                                           `json:"isShareable"`
	ShareableGuardianIDs              []int64                                        `json:"shareableGuardianIds,omitempty"`
	TemplateTitle                     *string                                        `json:"templateTitle,omitempty"`
}

// ExternalSecureDocumentDetailsDto represents an external secure document with attachment.
type ExternalSecureDocumentDetailsDto struct {
	ID                                *int64                                         `json:"id,omitempty"`
	HasMedia                          bool                                           `json:"hasMedia"`
	CanEdit                           bool                                           `json:"canEdit"`
	CanEditLockedStatus               bool                                           `json:"canEditLockedStatus"`
	IsLocked                          bool                                           `json:"isLocked"`
	JournalingStatus                  *string                                        `json:"journalingStatus,omitempty"`
	Category                          *string                                        `json:"category,omitempty"`
	DocumentTemplateTitle             *string                                        `json:"documentTemplateTitle,omitempty"`
	InstitutionCode                   *string                                        `json:"institutionCode,omitempty"`
	DocumentType                      *string                                        `json:"documentType,omitempty"`
	AssociatedInstitutionProfiles     []SecureDocumentAssociateInstitutionProfileDto  `json:"associatedInstitutionProfiles,omitempty"`
	SharedWithGroups                  []SecureDocumentShareWithGroupDto               `json:"sharedWithGroups,omitempty"`
	SharedWithInstitutionProfiles     []SecureDocumentShareWithInstitutionProfileDto  `json:"sharedWithInstitutionProfiles,omitempty"`
	ImplicitSharings                  []ImplicitSharingProfileDto                    `json:"implicitSharings,omitempty"`
	Creator                           *SecureDocumentCreatorDto                       `json:"creator,omitempty"`
	CreatedAt                         *string                                        `json:"createdAt,omitempty"`
	Title                             *string                                        `json:"title,omitempty"`
	UpdatedAt                         *string                                        `json:"updatedAt,omitempty"`
	UpdatedBy                         *string                                        `json:"updatedBy,omitempty"`
	Version                           *int                                           `json:"version,omitempty"`
	Description                       *string                                        `json:"description,omitempty"`
	IsSharedWithGuardian              bool                                           `json:"isSharedWithGuardian"`
	IsShareable                       bool                                           `json:"isShareable"`
	ShareableGuardianIDs              []int64                                        `json:"shareableGuardianIds,omitempty"`
	TemplateTitle                     *string                                        `json:"templateTitle,omitempty"`
	Attachment                        *FilesAulaFileResultDto                        `json:"attachment,omitempty"`
}

// InternalSecureDocumentDetailsDto represents an internal secure document with rich content and attachments.
type InternalSecureDocumentDetailsDto struct {
	ID                                *int64                                         `json:"id,omitempty"`
	HasMedia                          bool                                           `json:"hasMedia"`
	CanEdit                           bool                                           `json:"canEdit"`
	CanEditLockedStatus               bool                                           `json:"canEditLockedStatus"`
	IsLocked                          bool                                           `json:"isLocked"`
	JournalingStatus                  *string                                        `json:"journalingStatus,omitempty"`
	Category                          *string                                        `json:"category,omitempty"`
	DocumentTemplateTitle             *string                                        `json:"documentTemplateTitle,omitempty"`
	InstitutionCode                   *string                                        `json:"institutionCode,omitempty"`
	DocumentType                      *string                                        `json:"documentType,omitempty"`
	AssociatedInstitutionProfiles     []SecureDocumentAssociateInstitutionProfileDto  `json:"associatedInstitutionProfiles,omitempty"`
	SharedWithGroups                  []SecureDocumentShareWithGroupDto               `json:"sharedWithGroups,omitempty"`
	SharedWithInstitutionProfiles     []SecureDocumentShareWithInstitutionProfileDto  `json:"sharedWithInstitutionProfiles,omitempty"`
	ImplicitSharings                  []ImplicitSharingProfileDto                    `json:"implicitSharings,omitempty"`
	Creator                           *SecureDocumentCreatorDto                       `json:"creator,omitempty"`
	CreatedAt                         *string                                        `json:"createdAt,omitempty"`
	Title                             *string                                        `json:"title,omitempty"`
	UpdatedAt                         *string                                        `json:"updatedAt,omitempty"`
	UpdatedBy                         *string                                        `json:"updatedBy,omitempty"`
	Version                           *int                                           `json:"version,omitempty"`
	Description                       *string                                        `json:"description,omitempty"`
	IsSharedWithGuardian              bool                                           `json:"isSharedWithGuardian"`
	IsShareable                       bool                                           `json:"isShareable"`
	ShareableGuardianIDs              []int64                                        `json:"shareableGuardianIds,omitempty"`
	TemplateTitle                     *string                                        `json:"templateTitle,omitempty"`
	Attachments                       []FilesAulaFileResultDto                       `json:"attachments,omitempty"`
	Content                           *RichTextWrapperDto                            `json:"content,omitempty"`
}

// SecureDocumentExportDto represents secure document export tracking.
type SecureDocumentExportDto struct {
	RequestExportJobID *int64   `json:"requestExportJobId,omitempty"`
	Status             *string  `json:"status,omitempty"`
	Progress           *float32 `json:"progress,omitempty"`
	FileURL            *string  `json:"fileUrl,omitempty"`
	FileName           *string  `json:"fileName,omitempty"`
}

// CommonFileInstitutionDto represents institution info on a common file.
type CommonFileInstitutionDto struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	Name            *string `json:"name,omitempty"`
}

// CommonFileGroupRestrictionDto represents group restriction on a common file.
type CommonFileGroupRestrictionDto struct {
	ID              *int64  `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	InstitutionCode *string `json:"institutionCode,omitempty"`
}

// CommonFileDto represents a common (institution-level) file.
type CommonFileDto struct {
	ID                       *int64                          `json:"id,omitempty"`
	Attachment               *FilesAulaFileResultDto         `json:"attachment,omitempty"`
	Created                  *string                         `json:"created,omitempty"`
	Institution              *CommonFileInstitutionDto       `json:"institution,omitempty"`
	IsDataPolicy             bool                            `json:"isDataPolicy"`
	Title                    *string                         `json:"title,omitempty"`
	ProfileTypeRestrictions  []string                        `json:"profileTypeRestrictions,omitempty"`
	GroupRestrictions        []CommonFileGroupRestrictionDto `json:"groupRestrictions,omitempty"`
}

// DocumentsSimpleInstitutionProfile represents a simple institution profile reference used in documents/revisions.
type DocumentsSimpleInstitutionProfile struct {
	ProfileID            *int64                         `json:"profileId,omitempty"`
	InstitutionProfileID *int64                         `json:"institutionProfileId,omitempty"`
	InstitutionCode      *string                        `json:"institutionCode,omitempty"`
	InstitutionName      *string                        `json:"institutionName,omitempty"`
	Name                 *string                        `json:"name,omitempty"`
	Role                 *string                        `json:"role,omitempty"`
	MainGroup            *string                        `json:"mainGroup,omitempty"`
	ProfilePicture       *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
	ShortName            *string                        `json:"shortName,omitempty"`
	Metadata             *string                        `json:"metadata,omitempty"`
}

// DocumentRevisionDto represents a document revision entry.
type DocumentRevisionDto struct {
	ID            *int64                              `json:"id,omitempty"`
	CreatedBy     *string                             `json:"createdBy,omitempty"`
	CreatedAt     *string                             `json:"createdAt,omitempty"`
	Title         *string                             `json:"title,omitempty"`
	ChangeType    *string                             `json:"changeType,omitempty"`
	SharedWith    []DocumentsSimpleInstitutionProfile `json:"sharedWith,omitempty"`
	UnsharedWith  []DocumentsSimpleInstitutionProfile `json:"unsharedWith,omitempty"`
	IsAvailable   bool                                `json:"isAvailable"`
	RecipientName *string                             `json:"recipientName,omitempty"`
	ChildrenNames []string                            `json:"childrenNames,omitempty"`
}

// DocumentRevisionPageDto represents a paginated document revision list.
type DocumentRevisionPageDto struct {
	TotalCount           *int                  `json:"totalCount,omitempty"`
	DocumentRevisionDtos []DocumentRevisionDto `json:"documentRevisionDtos,omitempty"`
}

// GetImplicitSharingsDto represents implicit sharings result wrapper.
type GetImplicitSharingsDto struct {
	ImplicitSharings []ImplicitSharingProfileDto `json:"implicitSharings,omitempty"`
}

// CreateDocumentShareGroupArguments represents group sharing argument for document creation.
type CreateDocumentShareGroupArguments struct {
	GroupID *int64 `json:"groupId,omitempty"`
	CanEdit bool   `json:"canEdit"`
}

// CreateDocumentSharedProfileArguments represents profile sharing argument for document creation.
type CreateDocumentSharedProfileArguments struct {
	InstitutionProfileID *int64 `json:"institutionProfileId,omitempty"`
	CanEdit              bool   `json:"canEdit"`
}

// CreateDocumentArguments represents arguments for creating a secure document.
type CreateDocumentArguments struct {
	ID                                *int64                                 `json:"id,omitempty"`
	Category                          *string                                `json:"category,omitempty"`
	CreatorInstitutionProfileID       *int64                                 `json:"creatorInstitutionProfileId,omitempty"`
	RegardingInstitutionProfileIDs    []int64                                `json:"regardingInstitutionProfileIds,omitempty"`
	SharedWithGroups                  []CreateDocumentShareGroupArguments    `json:"sharedWithGroups,omitempty"`
	SharedWithInstitutionProfiles     []CreateDocumentSharedProfileArguments `json:"sharedWithInstitutionProfiles,omitempty"`
	Title                             *string                                `json:"title,omitempty"`
	Version                           *int                                   `json:"version,omitempty"`
	ForceUpdate                       *bool                                  `json:"forceUpdate,omitempty"`
	AttachedThread                    *AttachMessagesToSecureDocumentRequest  `json:"attachedThread,omitempty"`
}

// CreateExternalDocumentArguments represents arguments for creating an external secure document.
type CreateExternalDocumentArguments struct {
	ID                                *int64                                 `json:"id,omitempty"`
	Category                          *string                                `json:"category,omitempty"`
	CreatorInstitutionProfileID       *int64                                 `json:"creatorInstitutionProfileId,omitempty"`
	RegardingInstitutionProfileIDs    []int64                                `json:"regardingInstitutionProfileIds,omitempty"`
	SharedWithGroups                  []CreateDocumentShareGroupArguments    `json:"sharedWithGroups,omitempty"`
	SharedWithInstitutionProfiles     []CreateDocumentSharedProfileArguments `json:"sharedWithInstitutionProfiles,omitempty"`
	Title                             *string                                `json:"title,omitempty"`
	Version                           *int                                   `json:"version,omitempty"`
	ForceUpdate                       *bool                                  `json:"forceUpdate,omitempty"`
	AttachedThread                    *AttachMessagesToSecureDocumentRequest  `json:"attachedThread,omitempty"`
	ExternalFile                      *UploadFileToAulaArguments             `json:"externalFile,omitempty"`
}

// CreateInternalDocumentArguments represents arguments for creating an internal secure document.
type CreateInternalDocumentArguments struct {
	ID                                *int64                                 `json:"id,omitempty"`
	Category                          *string                                `json:"category,omitempty"`
	CreatorInstitutionProfileID       *int64                                 `json:"creatorInstitutionProfileId,omitempty"`
	RegardingInstitutionProfileIDs    []int64                                `json:"regardingInstitutionProfileIds,omitempty"`
	SharedWithGroups                  []CreateDocumentShareGroupArguments    `json:"sharedWithGroups,omitempty"`
	SharedWithInstitutionProfiles     []CreateDocumentSharedProfileArguments `json:"sharedWithInstitutionProfiles,omitempty"`
	Title                             *string                                `json:"title,omitempty"`
	Version                           *int                                   `json:"version,omitempty"`
	ForceUpdate                       *bool                                  `json:"forceUpdate,omitempty"`
	AttachedThread                    *AttachMessagesToSecureDocumentRequest  `json:"attachedThread,omitempty"`
	Content                           *string                                `json:"content,omitempty"`
	AttachmentIDs                     []int64                                `json:"attachmentIds,omitempty"`
	ImplicitSharingOverrides          []ImplicitSharingOverride             `json:"implicitSharingOverrides,omitempty"`
}

// SortingModel represents a sort model for document queries.
type SortingModel struct {
	Field *string `json:"field,omitempty"`
	Order *string `json:"order,omitempty"`
}

// GetCommonFilesArguments represents arguments for querying common files.
type GetCommonFilesArguments struct {
	Page      *int    `json:"page,omitempty"`
	SortType  *string `json:"sortType,omitempty"`
	SortOrder *string `json:"sortOrder,omitempty"`
}

// GetSecureDocumentsArguments represents arguments for querying secure documents.
type GetSecureDocumentsArguments struct {
	FilterInstitutionProfileIDs   []int64        `json:"filterInstitutionProfileIds,omitempty"`
	FilterRegardingGroupIDs       []int64        `json:"filterRegardingGroupIds,omitempty"`
	FilterUnread                  *bool          `json:"filterUnread,omitempty"`
	FilterLocked                  *bool          `json:"filterLocked,omitempty"`
	FilterJournalingStatus        *string        `json:"filterJournalingStatus,omitempty"`
	FilterEditable                bool           `json:"filterEditable"`
	DocumentType                  *string        `json:"documentType,omitempty"`
	Sortings                      []SortingModel `json:"sortings,omitempty"`
	Index                         *int           `json:"index,omitempty"`
	Limit                         *int           `json:"limit,omitempty"`
	FilterRegardingStudentIDs     []int64        `json:"filterRegardingStudentIds,omitempty"`
	FilterDocumentCategories      []string       `json:"filterDocumentCategories,omitempty"`
}

// GetShareableSecureDocumentsArguments represents arguments for querying shareable secure documents.
type GetShareableSecureDocumentsArguments struct {
	FilterInstitutionProfileIDs   []int64        `json:"filterInstitutionProfileIds,omitempty"`
	FilterRegardingGroupIDs       []int64        `json:"filterRegardingGroupIds,omitempty"`
	FilterUnread                  *bool          `json:"filterUnread,omitempty"`
	FilterLocked                  *bool          `json:"filterLocked,omitempty"`
	FilterJournalingStatus        *string        `json:"filterJournalingStatus,omitempty"`
	FilterEditable                bool           `json:"filterEditable"`
	DocumentType                  *string        `json:"documentType,omitempty"`
	Sortings                      []SortingModel `json:"sortings,omitempty"`
	Index                         *int           `json:"index,omitempty"`
	Limit                         *int           `json:"limit,omitempty"`
	FilterRegardingStudentIDs     []int64        `json:"filterRegardingStudentIds,omitempty"`
	FilterDocumentCategories      []string       `json:"filterDocumentCategories,omitempty"`
	ShareToInstitutionProfileIDs  []int64        `json:"shareToInstitutionProfileIds,omitempty"`
}

// RemoveSharingArguments represents arguments for removing sharing from documents.
type RemoveSharingArguments struct {
	DocumentIDs []int64 `json:"documentIds,omitempty"`
}

// UpdateSharingArguments represents arguments for updating sharing on documents.
type UpdateSharingArguments struct {
	DocumentIDs                 []int64                              `json:"documentIds,omitempty"`
	ResetSharings               bool                                 `json:"resetSharings"`
	SharedGroups                []UpdateSharingGroupArguments        `json:"sharedGroups,omitempty"`
	SharedInstitutionProfiles   []UpdateSharingInstProfileArguments  `json:"sharedInstitutionProfiles,omitempty"`
}

// UpdateSharingGroupArguments represents group argument for sharing updates.
type UpdateSharingGroupArguments struct {
	GroupID *int64 `json:"groupId,omitempty"`
	CanEdit bool   `json:"canEdit"`
}

// UpdateSharingInstProfileArguments represents profile argument for sharing updates.
type UpdateSharingInstProfileArguments struct {
	InstitutionProfileID *int64 `json:"institutionProfileId,omitempty"`
	CanEdit              bool   `json:"canEdit"`
}

// CreateExportForMultipleSecureDocumentsRequest represents a request to export multiple secure documents.
type CreateExportForMultipleSecureDocumentsRequest struct {
	SecureDocumentIDs []int64 `json:"secureDocumentIds,omitempty"`
}

// CreatePdfForSingleDocumentRequest represents a request to create a PDF for a single document.
type CreatePdfForSingleDocumentRequest struct {
	SecureDocumentID *int64 `json:"secureDocumentId,omitempty"`
}

// TrackCreatePdfForSingleDocumentRequest represents tracking a single document PDF export.
type TrackCreatePdfForSingleDocumentRequest struct {
	RequestID *int64 `json:"requestId,omitempty"`
}

// TrackExportForMultipleSecureDocumentsRequest represents tracking a multi-document export.
type TrackExportForMultipleSecureDocumentsRequest struct {
	RequestID *int64 `json:"requestId,omitempty"`
}

// GetCommonFilesResult represents a result of querying common files.
type GetCommonFilesResult struct {
	CommonFiles []CommonFileDto `json:"commonFiles,omitempty"`
	TotalAmount *int            `json:"totalAmount,omitempty"`
}

// GetSecureDocumentsRegardingInstitutionProfile represents a profile in secure document filter results.
type GetSecureDocumentsRegardingInstitutionProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetSecureDocumentsFilter represents filter metadata returned with secure document queries.
type GetSecureDocumentsFilter struct {
	RegardingGroups                []SecureDocumentAssociateGroupDto              `json:"regardingGroups,omitempty"`
	RegardingInstitutionProfiles   []GetSecureDocumentsRegardingInstitutionProfile `json:"regardingInstitutionProfiles,omitempty"`
	DocumentCategories             []string                                       `json:"documentCategories,omitempty"`
	SharedGroups                   []SecureDocumentAssociateGroupDto              `json:"sharedGroups,omitempty"`
	SharedInstitutionProfiles      []GetSecureDocumentsRegardingInstitutionProfile `json:"sharedInstitutionProfiles,omitempty"`
}

// GetSecureDocumentsResult represents a result of querying secure documents.
type GetSecureDocumentsResult struct {
	Documents  []SecureDocumentDto       `json:"documents,omitempty"`
	Filters    *GetSecureDocumentsFilter `json:"filters,omitempty"`
	TotalCount *int                      `json:"totalCount,omitempty"`
}
