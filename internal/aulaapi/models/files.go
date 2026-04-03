package models

// ObjectWithId represents a simple object with an id.
type ObjectWithId struct {
	ID *int64 `json:"id,omitempty"`
}

// MembershipCountResultModel represents membership counts by portal role.
type MembershipCountResultModel struct {
	Employees *int64 `json:"employees,omitempty"`
	Children  *int64 `json:"children,omitempty"`
	Guardians *int64 `json:"guardians,omitempty"`
	Total     *int64 `json:"total,omitempty"`
}

// ShareWithGroupDto represents a group sharing reference with portal roles.
type ShareWithGroupDto struct {
	ID                    *int64                      `json:"id,omitempty"`
	PortalRoles           []string                    `json:"portalRoles,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	ShortName             *string                     `json:"shortName,omitempty"`
	InstitutionCode       *string                     `json:"institutionCode,omitempty"`
	InstitutionName       *string                     `json:"institutionName,omitempty"`
	MembershipCount       *MembershipCountResultModel `json:"membershipCount,omitempty"`
	AllowMembersToBeShown bool                        `json:"allowMembersToBeShown"`
}

// LinkedGroupRequestModel represents a linked group request.
type LinkedGroupRequestModel struct {
	GroupID        *int64   `json:"groupId,omitempty"`
	PortalRolesEnum []string `json:"portalRolesEnum,omitempty"`
}

// FilesAulaFileContent represents file content stored in S3.
type FilesAulaFileContent struct {
	ID             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	URL            *string `json:"url,omitempty"`
	Bucket         *string `json:"bucket,omitempty"`
	Key            *string `json:"key,omitempty"`
	Created        *string `json:"created,omitempty"`
	ScanningStatus *string `json:"scanningStatus,omitempty"`
}

// AulaLinkContent represents an external link content.
type AulaLinkContent struct {
	Service *string `json:"service,omitempty"`
	Name    *string `json:"name,omitempty"`
	URL     *string `json:"url,omitempty"`
}

// AulaDocumentLinkContent represents a document link content.
type AulaDocumentLinkContent struct {
	ID           *int64  `json:"id,omitempty"`
	Title        *string `json:"title,omitempty"`
	CanAccess    bool    `json:"canAccess"`
	DocumentType *string `json:"documentType,omitempty"`
	IsDeleted    bool    `json:"isDeleted"`
}

// AulaFileAlbumDto represents album info embedded in file results.
type AulaFileAlbumDto struct {
	Name         *string             `json:"name,omitempty"`
	SharedGroups []ShareWithGroupDto `json:"sharedGroups,omitempty"`
}

// AulaFileResultProfileDto represents profile info on a file result.
type AulaFileResultProfileDto struct {
	InstProfileID   *int64                `json:"instProfileId,omitempty"`
	InstitutionCode *string               `json:"institutionCode,omitempty"`
	InstitutionName *string               `json:"institutionName,omitempty"`
	Role            *string               `json:"role,omitempty"`
	ProfileID       *int64                `json:"profileId,omitempty"`
	Name            *string               `json:"name,omitempty"`
	ShortName       *string               `json:"shortName,omitempty"`
	Metadata        *string               `json:"metadata,omitempty"`
	ProfilePicture  *FilesAulaFileContent `json:"profilePicture,omitempty"`
}

// AulaMediaFileContent represents media file content with thumbnails, tags, and permissions.
type AulaMediaFileContent struct {
	Title                        *string                    `json:"title,omitempty"`
	Description                  *string                    `json:"description,omitempty"`
	AllowsComments               bool                       `json:"allowsComments"`
	CanViewComments              bool                       `json:"canViewComments"`
	File                         *FilesAulaFileContent      `json:"file,omitempty"`
	MediaType                    *string                    `json:"mediaType,omitempty"`
	Tags                         []AulaFileResultProfileDto `json:"tags,omitempty"`
	DurationNumber               *float64                   `json:"durationNumber,omitempty"`
	Album                        *AulaFileAlbumDto          `json:"album,omitempty"`
	ThumbnailURL                 *string                    `json:"thumbnailUrl,omitempty"`
	LargeThumbnailURL            *string                    `json:"largeThumbnailUrl,omitempty"`
	MediumThumbnailURL           *string                    `json:"mediumThumbnailUrl,omitempty"`
	SmallThumbnailURL            *string                    `json:"smallThumbnailUrl,omitempty"`
	ExtraSmallThumbnailURL       *string                    `json:"extraSmallThumbnailUrl,omitempty"`
	HasVideoThumbnail            bool                       `json:"hasVideoThumbnail"`
	CurrentUserCanDelete         bool                       `json:"currentUserCanDelete"`
	CurrentUserCanEditMetadata   bool                       `json:"currentUserCanEditMetadata"`
	CurrentUserCanReport         bool                       `json:"currentUserCanReport"`
	CurrentUserCanEditTags       bool                       `json:"currentUserCanEditTags"`
	IsUploadingPending           bool                       `json:"isUploadingPending"`
	ConversionStatus             *string                    `json:"conversionStatus,omitempty"`
}

// FilesAulaFileResultDto represents a top-level file result DTO.
type FilesAulaFileResultDto struct {
	ID       *int64                    `json:"id,omitempty"`
	Creator  *AulaFileResultProfileDto `json:"creator,omitempty"`
	File     *FilesAulaFileContent     `json:"file,omitempty"`
	Media    *AulaMediaFileContent     `json:"media,omitempty"`
	Link     *AulaLinkContent          `json:"link,omitempty"`
	Document *AulaDocumentLinkContent  `json:"document,omitempty"`
	Status   *string                   `json:"status,omitempty"`
}

// AulaGalleryMediaFileResultDto represents a gallery media file result.
type AulaGalleryMediaFileResultDto struct {
	Title                        *string                    `json:"title,omitempty"`
	Description                  *string                    `json:"description,omitempty"`
	AllowsComments               bool                       `json:"allowsComments"`
	CanViewComments              bool                       `json:"canViewComments"`
	File                         *FilesAulaFileContent      `json:"file,omitempty"`
	MediaType                    *string                    `json:"mediaType,omitempty"`
	Tags                         []AulaFileResultProfileDto `json:"tags,omitempty"`
	DurationNumber               *float64                   `json:"durationNumber,omitempty"`
	Album                        *AulaFileAlbumDto          `json:"album,omitempty"`
	ThumbnailURL                 *string                    `json:"thumbnailUrl,omitempty"`
	LargeThumbnailURL            *string                    `json:"largeThumbnailUrl,omitempty"`
	MediumThumbnailURL           *string                    `json:"mediumThumbnailUrl,omitempty"`
	SmallThumbnailURL            *string                    `json:"smallThumbnailUrl,omitempty"`
	ExtraSmallThumbnailURL       *string                    `json:"extraSmallThumbnailUrl,omitempty"`
	HasVideoThumbnail            bool                       `json:"hasVideoThumbnail"`
	CurrentUserCanDelete         bool                       `json:"currentUserCanDelete"`
	CurrentUserCanEditMetadata   bool                       `json:"currentUserCanEditMetadata"`
	CurrentUserCanReport         bool                       `json:"currentUserCanReport"`
	CurrentUserCanEditTags       bool                       `json:"currentUserCanEditTags"`
	IsUploadingPending           bool                       `json:"isUploadingPending"`
	ConversionStatus             *string                    `json:"conversionStatus,omitempty"`
	Creator                      *AulaFileResultProfileDto  `json:"creator,omitempty"`
	ID                           *int64                     `json:"id,omitempty"`
	CommentCount                 *int                       `json:"commentCount,omitempty"`
}

// AuthorizedFileFormat represents an authorized file format.
type AuthorizedFileFormat struct {
	ID         *int64  `json:"id,omitempty"`
	FileFormat *string `json:"fileFormat,omitempty"`
	Name       *string `json:"name,omitempty"`
}

// FileConnectionResult represents a file connection result.
type FileConnectionResult struct {
	FileName *string `json:"fileName,omitempty"`
	Mime     *string `json:"mime,omitempty"`
	FilePath *string `json:"filePath,omitempty"`
	Length   *int64  `json:"length,omitempty"`
}

// UploadFileKeyInfo represents S3 key info for uploads.
type UploadFileKeyInfo struct {
	Key    *string `json:"key,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
}

// UploadFileInfo represents upload file wrapper.
type UploadFileInfo struct {
	Key *UploadFileKeyInfo `json:"key,omitempty"`
}

// UploadFileData represents S3 pre-signed upload data.
type UploadFileData struct {
	Policy           *string `json:"policy,omitempty"`
	AmzAlgorithm     *string `json:"X-Amz-Algorithm,omitempty"`
	AmzCredential    *string `json:"X-Amz-Credential,omitempty"`
	AmzDate          *string `json:"X-Amz-Date,omitempty"`
	AmzSecurityToken *string `json:"X-Amz-Security-Token,omitempty"`
	AmzSignature     *string `json:"X-Amz-Signature,omitempty"`
	Acl              *string `json:"acl,omitempty"`
	Key              *string `json:"key,omitempty"`
	Bucket           *string `json:"bucket,omitempty"`
	CacheControl     *string `json:"Cache-Control,omitempty"`
}

// UploadLink represents an upload link.
type UploadLink struct {
	Action *string         `json:"action,omitempty"`
	File   *UploadFileInfo `json:"file,omitempty"`
	Data   *UploadFileData `json:"data,omitempty"`
}

// FilePartUploadInformation represents part info for multipart uploads.
type FilePartUploadInformation struct {
	PartIndex    *string `json:"partIndex,omitempty"`
	PreSignedURL *string `json:"preSignedUrl,omitempty"`
}

// FileUploadInformation represents multipart upload info.
type FileUploadInformation struct {
	Parts       []FilePartUploadInformation `json:"parts,omitempty"`
	AwsUploadID *string                     `json:"awsUploadId,omitempty"`
}

// BaseResultDto represents a base result DTO with an id.
type BaseResultDto struct {
	ID *int64 `json:"id,omitempty"`
}

// FileResultDto represents a file result with upload info.
type FileResultDto struct {
	ID                              *int64                 `json:"id,omitempty"`
	Name                            *string                `json:"name,omitempty"`
	MultipartFileUploadInformation  *FileUploadInformation `json:"multipartFileUploadInformation,omitempty"`
	UploadID                        *string                `json:"uploadId,omitempty"`
}

// LinkResultDto represents a link result.
type LinkResultDto struct {
	ID      *int64  `json:"id,omitempty"`
	Service *string `json:"service,omitempty"`
	Name    *string `json:"name,omitempty"`
	URL     *string `json:"url,omitempty"`
}

// DocumentLinkResult represents a document link result.
type DocumentLinkResult struct {
	ID         *int64 `json:"id,omitempty"`
	DocumentID *int64 `json:"documentId,omitempty"`
}

// CreateAttachmentsResult represents a result of creating attachments.
type CreateAttachmentsResult struct {
	Media              []FileResultDto `json:"media,omitempty"`
	Files              []FileResultDto `json:"files,omitempty"`
	Documents          []FileResultDto `json:"documents,omitempty"`
	Links              []LinkResultDto `json:"links,omitempty"`
	IsAllConsentsValid bool            `json:"isAllConsentsValid"`
}

// CreateMediaResult represents a result of creating media.
type CreateMediaResult struct {
	AllImagesHasValidConsents bool                     `json:"allImagesHasValidConsents"`
	Media                     []FilesAulaFileResultDto `json:"media,omitempty"`
}

// MediaTagConsentsResult represents a consent validation result for media tags.
type MediaTagConsentsResult struct {
	IsAllConsentsValid bool `json:"isAllConsentsValid"`
}

// DeleteMediaParameters represents parameters for deleting media.
type DeleteMediaParameters struct {
	MediaIDs []int64 `json:"mediaIds,omitempty"`
}

// UploadFileContentArguments represents content arguments for file upload.
type UploadFileContentArguments struct {
	Key    *string `json:"key,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	Name   *string `json:"name,omitempty"`
	ID     *int64  `json:"id,omitempty"`
}

// UploadMediaContentArguments represents content arguments for media upload.
type UploadMediaContentArguments struct {
	Duration    *float64                    `json:"duration,omitempty"`
	Tags        []ObjectWithId              `json:"tags,omitempty"`
	Title       *string                     `json:"title,omitempty"`
	Description *string                     `json:"description,omitempty"`
	MediaType   *string                     `json:"mediaType,omitempty"`
	File        *UploadFileContentArguments `json:"file,omitempty"`
}

// UploadLinkContentArguments represents content arguments for link upload.
type UploadLinkContentArguments struct {
	ExternalFileID *string `json:"externalFileId,omitempty"`
	AccessToken    *string `json:"accessToken,omitempty"`
	Service        *string `json:"service,omitempty"`
}

// UploadFileToAulaArguments represents arguments for uploading a file to Aula.
type UploadFileToAulaArguments struct {
	Size      *float32                    `json:"size,omitempty"`
	Creator   *ObjectWithId               `json:"creator,omitempty"`
	File      *UploadFileContentArguments `json:"file,omitempty"`
	Media     *UploadMediaContentArguments `json:"media,omitempty"`
	Link      *UploadLinkContentArguments `json:"link,omitempty"`
	ID        *int64                      `json:"id,omitempty"`
	Name      *string                     `json:"name,omitempty"`
	IsLoading bool                        `json:"isLoading"`
}

// CreateAttachmentsArguments represents arguments for creating attachments.
type CreateAttachmentsArguments struct {
	InstitutionCode             *string                              `json:"institutionCode,omitempty"`
	OwnerInstitutionProfileID   *int64                               `json:"ownerInstitutionProfileId,omitempty"`
	Media                       []AttachmentMediaFileUploadArguments `json:"media,omitempty"`
	Links                       []AttachmentLinkUploadArguments      `json:"links,omitempty"`
	Files                       []AttachmentFileUploadArguments      `json:"files,omitempty"`
	AttachedSecureDocumentIDs   []int64                              `json:"attachedSecureDocumentIds,omitempty"`
}

// GetUploadLinksArguments represents getting upload link names.
type GetUploadLinksArguments struct {
	UploadNames     []string `json:"upload_names,omitempty"`
	InstitutionCode *string  `json:"institutionCode,omitempty"`
}

// AddOrRemoveTagArguments represents adding or removing a tag on a media item.
type AddOrRemoveTagArguments struct {
	InstProfileID *int64 `json:"instProfileId,omitempty"`
	MediaID       *int64 `json:"mediaId,omitempty"`
}

// CompleteMultipartUploadPartRequest represents completing a multipart upload part.
type CompleteMultipartUploadPartRequest struct {
	ETag       *string `json:"eTag,omitempty"`
	PartNumber *string `json:"partNumber,omitempty"`
}

// CompleteMultipartUploadingRequest represents completing a multipart upload.
type CompleteMultipartUploadingRequest struct {
	FileID *int64                               `json:"fileId,omitempty"`
	Parts  []CompleteMultipartUploadPartRequest `json:"parts,omitempty"`
}

// MultipartUploadingInfoArguments represents multipart upload info arguments.
type MultipartUploadingInfoArguments struct {
	NumberOfPart *int `json:"numberOfPart,omitempty"`
}

// BaseFileUploadArguments represents base file upload arguments.
type BaseFileUploadArguments struct {
	UploadID              *string                          `json:"uploadId,omitempty"`
	MultipartUploadingInfo *MultipartUploadingInfoArguments `json:"multipartUploadingInfo,omitempty"`
}

// AttachmentFileUploadArguments represents file attachment upload arguments.
type AttachmentFileUploadArguments struct {
	UploadID              *string                          `json:"uploadId,omitempty"`
	MultipartUploadingInfo *MultipartUploadingInfoArguments `json:"multipartUploadingInfo,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
}

// AttachmentLinkUploadArguments represents link attachment upload arguments.
type AttachmentLinkUploadArguments struct {
	ExternalFileID *string `json:"externalFileId,omitempty"`
	AccessToken    *string `json:"accessToken,omitempty"`
	Service        *string `json:"service,omitempty"`
}

// AttachmentMediaFileUploadArguments represents media attachment upload arguments.
type AttachmentMediaFileUploadArguments struct {
	UploadID              *string                          `json:"uploadId,omitempty"`
	MultipartUploadingInfo *MultipartUploadingInfoArguments `json:"multipartUploadingInfo,omitempty"`
	ID                    *int64                           `json:"id,omitempty"`
	AlbumID               *int64                           `json:"albumId,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	MediaType             *string                          `json:"mediaType,omitempty"`
	Tags                  []int64                          `json:"tags,omitempty"`
	Title                 *string                          `json:"title,omitempty"`
	Description           *string                          `json:"description,omitempty"`
}

// UpdateAttachmentsArguments represents update attachment arguments.
type UpdateAttachmentsArguments struct {
	Media []AttachmentMediaFileUploadArguments `json:"media,omitempty"`
}

// UploadAttachmentServiceResult represents upload attachment service result.
type UploadAttachmentServiceResult struct {
	MediaIDs          []int64 `json:"mediaIds,omitempty"`
	AllConsentIsValid bool    `json:"allConsentIsValid"`
	IsSuccess         bool    `json:"isSuccess"`
}

// CommonFileModel represents a common file model.
type CommonFileModel struct {
	ID    *int64                         `json:"id,omitempty"`
	Title *string                        `json:"title,omitempty"`
	File  *DownloadFileFromAulaArguments `json:"file,omitempty"`
}
