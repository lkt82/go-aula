package models

import "encoding/json"

// SearchResponse represents a top-level search response.
type SearchResponse struct {
	TotalSize      *int                         `json:"totalSize,omitempty"`
	DocTypeCount   []SearchResultCountItem      `json:"docTypeCount,omitempty"`
	GroupTypeCount []SearchResultGroupCountItem `json:"groupTypeCount,omitempty"`
	Results        []SearchResultItem           `json:"results,omitempty"`
}

// SearchResultCountItem represents a doc-type facet count.
type SearchResultCountItem struct {
	Name  *string `json:"name,omitempty"`
	Count *int    `json:"count,omitempty"`
}

// SearchResultGroupCountItem represents a group-type facet count.
type SearchResultGroupCountItem struct {
	Name  *string `json:"name,omitempty"`
	Count *int    `json:"count,omitempty"`
	Key   *string `json:"key,omitempty"`
}

// SearchResultHighlight represents a search result highlight fragment.
type SearchResultHighlight struct {
	Property *string `json:"property,omitempty"`
	Fragment *string `json:"fragment,omitempty"`
}

// SearchResultItem represents a base search result item.
type SearchResultItem struct {
	DocID            *string `json:"docId,omitempty"`
	DocType          *string `json:"docType,omitempty"`
	InstitutionCode  *string `json:"institutionCode,omitempty"`
	InstitutionName  *string `json:"institutionName,omitempty"`
	MunicipalityCode *string `json:"municipalityCode,omitempty"`
	MunicipalityName *string `json:"municipalityName,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
}

// SearchResultCommonFile represents a search result for a common file.
type SearchResultCommonFile struct {
	DocID            *string `json:"docId,omitempty"`
	DocType          *string `json:"docType,omitempty"`
	InstitutionCode  *string `json:"institutionCode,omitempty"`
	InstitutionName  *string `json:"institutionName,omitempty"`
	MunicipalityCode *string `json:"municipalityCode,omitempty"`
	MunicipalityName *string `json:"municipalityName,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	Title            *string `json:"title,omitempty"`
	Created          *string `json:"created,omitempty"`
	FileKey          *string `json:"fileKey,omitempty"`
	FileBucket       *string `json:"fileBucket,omitempty"`
	URL              *string `json:"url,omitempty"`
	FileName         *string `json:"fileName,omitempty"`
	ScanningStatus   *string `json:"scanningStatus,omitempty"`
}

// SearchResultCommonInboxItem represents a search result for a common inbox.
type SearchResultCommonInboxItem struct {
	DocID            *string  `json:"docId,omitempty"`
	DocType          *string  `json:"docType,omitempty"`
	InstitutionCode  *string  `json:"institutionCode,omitempty"`
	InstitutionName  *string  `json:"institutionName,omitempty"`
	MunicipalityCode *string  `json:"municipalityCode,omitempty"`
	MunicipalityName *string  `json:"municipalityName,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Description      *string  `json:"description,omitempty"`
	ID               *int64   `json:"id,omitempty"`
	Score            *float32 `json:"score,omitempty"`
	AulaEmail        *string  `json:"aulaEmail,omitempty"`
}

// SearchResultEventItem represents a search result for a calendar event.
type SearchResultEventItem struct {
	DocID            *string `json:"docId,omitempty"`
	DocType          *string `json:"docType,omitempty"`
	InstitutionCode  *string `json:"institutionCode,omitempty"`
	InstitutionName  *string `json:"institutionName,omitempty"`
	MunicipalityCode *string `json:"municipalityCode,omitempty"`
	MunicipalityName *string `json:"municipalityName,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	ID               *int64  `json:"id,omitempty"`
	Title            *string `json:"title,omitempty"`
	StartDateTime    *string `json:"startDateTime,omitempty"`
	EndDateTime      *string `json:"endDateTime,omitempty"`
	CreatorAulaName  *string `json:"creatorAulaName,omitempty"`
	Location         *string `json:"location,omitempty"`
	EventType        *string `json:"type,omitempty"`
}

// SearchResultGroupItem represents a search result for a group.
type SearchResultGroupItem struct {
	DocID                               *string                     `json:"docId,omitempty"`
	DocType                             *string                     `json:"docType,omitempty"`
	InstitutionCode                     *string                     `json:"institutionCode,omitempty"`
	InstitutionName                     *string                     `json:"institutionName,omitempty"`
	MunicipalityCode                    *string                     `json:"municipalityCode,omitempty"`
	MunicipalityName                    *string                     `json:"municipalityName,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	Description                         *string                     `json:"description,omitempty"`
	ID                                  *int64                      `json:"id,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	Access                              *string                     `json:"access,omitempty"`
	DashboardEnabled                    *bool                       `json:"dashboardEnabled,omitempty"`
	CurrentUserCanAccessGroupDashboard  bool                        `json:"currentUserCanAccessGroupDashboard"`
	MembershipRole                      *string                     `json:"membershipRole,omitempty"`
	GroupType                           *string                     `json:"type,omitempty"`
	IsGroupMember                       *bool                       `json:"isGroupMember,omitempty"`
	ShortName                           *string                     `json:"shortName,omitempty"`
	MembershipCount                     *MembershipCountResultModel `json:"membershipCount,omitempty"`
	AllowMembersToBeShown               bool                        `json:"allowMembersToBeShown"`
	Admins                              []SearchResultGroupAdmin    `json:"admins,omitempty"`
}

// SearchResultGroupAdmin represents an admin entry within a group search result.
type SearchResultGroupAdmin struct {
	InstitutionProfileID *int64  `json:"institutionProfileId,omitempty"`
	FirstName            *string `json:"firstName,omitempty"`
	LastName             *string `json:"lastName,omitempty"`
	FullName             *string `json:"fullName,omitempty"`
}

// SearchResultMediaItem represents a search result for a media item.
type SearchResultMediaItem struct {
	DocID                *string                    `json:"docId,omitempty"`
	DocType              *string                    `json:"docType,omitempty"`
	InstitutionCode      *string                    `json:"institutionCode,omitempty"`
	InstitutionName      *string                    `json:"institutionName,omitempty"`
	MunicipalityCode     *string                    `json:"municipalityCode,omitempty"`
	MunicipalityName     *string                    `json:"municipalityName,omitempty"`
	Name                 *string                    `json:"name,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	ID                   *int64                     `json:"id,omitempty"`
	Creator              *AulaFileResultProfileDto  `json:"creator,omitempty"`
	Tags                 []AulaFileResultProfileDto `json:"tags,omitempty"`
	Title                *string                    `json:"title,omitempty"`
	AlbumTitle           *string                    `json:"albumTitle,omitempty"`
	AlbumDescription     *string                    `json:"albumDescription,omitempty"`
	ThumbnailURL         *string                    `json:"thumbnailUrl,omitempty"`
	LargeThumbnailURL    *string                    `json:"largeThumbnailUrl,omitempty"`
	MediumThumbnailURL   *string                    `json:"mediumThumbnailUrl,omitempty"`
	SmallThumbnailURL    *string                    `json:"smallThumbnailUrl,omitempty"`
	HasVideoThumbnail    bool                       `json:"hasVideoThumbnail"`
	ExtraSmallThumbnailURL *string                  `json:"extraSmallThumbnailUrl,omitempty"`
	MediaType            *string                    `json:"mediaType,omitempty"`
	File                 *FilesAulaFileContent      `json:"file,omitempty"`
	CurrentUserCanDelete *bool                      `json:"currentUserCanDelete,omitempty"`
	CanComment           bool                       `json:"canComment"`
	CommentCount         *int                       `json:"commentCount,omitempty"`
	ConversionStatus     *string                    `json:"conversionStatus,omitempty"`
}

// SearchResultPostItem represents a search result for a post.
type SearchResultPostItem struct {
	DocID            *string  `json:"docId,omitempty"`
	DocType          *string  `json:"docType,omitempty"`
	InstitutionCode  *string  `json:"institutionCode,omitempty"`
	InstitutionName  *string  `json:"institutionName,omitempty"`
	MunicipalityCode *string  `json:"municipalityCode,omitempty"`
	MunicipalityName *string  `json:"municipalityName,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Description      *string  `json:"description,omitempty"`
	ID               *int64   `json:"id,omitempty"`
	Title            *string  `json:"title,omitempty"`
	Content          *string  `json:"content,omitempty"`
	Timestamp        *string  `json:"timestamp,omitempty"`
	PublishAt        *string  `json:"publishAt,omitempty"`
	EditedAt         *string  `json:"editedAt,omitempty"`
	ReceiverGroups   []string `json:"receiverGroups,omitempty"`
	Creator          *SearchCreator `json:"creator,omitempty"`
}

// SearchCreator represents a creator reference for search results.
type SearchCreator struct {
	ProfileID *int64  `json:"profileId,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// SearchResultProfileItemBase represents a base profile search result.
type SearchResultProfileItemBase struct {
	DocID              *string                        `json:"docId,omitempty"`
	DocType            *string                        `json:"docType,omitempty"`
	InstitutionCode    *string                        `json:"institutionCode,omitempty"`
	InstitutionName    *string                        `json:"institutionName,omitempty"`
	MunicipalityCode   *string                        `json:"municipalityCode,omitempty"`
	MunicipalityName   *string                        `json:"municipalityName,omitempty"`
	Name               *string                        `json:"name,omitempty"`
	Description        *string                        `json:"description,omitempty"`
	ProfileID          *int64                         `json:"profileId,omitempty"`
	Address            *Address                       `json:"address,omitempty"`
	PortalRole         *string                        `json:"portalRole,omitempty"`
	InstitutionRole    *string                        `json:"institutionRole,omitempty"`
	InstitutionProfileID *int64                       `json:"id,omitempty"`
	HomePhoneNumber    *string                        `json:"homePhoneNumber,omitempty"`
	MobilePhoneNumber  *string                        `json:"mobilePhoneNumber,omitempty"`
	WorkPhoneNumber    *string                        `json:"workPhoneNumber,omitempty"`
	FirstName          *string                        `json:"firstName,omitempty"`
	LastName           *string                        `json:"lastName,omitempty"`
	Gender             *string                        `json:"gender,omitempty"`
	AulaEmail          *string                        `json:"aulaEmail,omitempty"`
	ExternalEmail      *string                        `json:"email,omitempty"`
	RoleDefinitions    []SearchRoleDefinition         `json:"roleDefinitions,omitempty"`
	MainGroup          *SearchMainGroup               `json:"mainGroup,omitempty"`
	ShortName          *string                        `json:"shortName,omitempty"`
	Metadata           *string                        `json:"metadata,omitempty"`
	RelatedGroups      []SearchRelatedGroup           `json:"relatedGroups,omitempty"`
	ProfilePicture     *DownloadFileFromAulaArguments  `json:"profilePicture,omitempty"`
}

// SearchResultProfileItemFindRecipients represents a profile search result for "find recipients" context.
type SearchResultProfileItemFindRecipients struct {
	SearchResultProfileItemBase
	RelatedProfiles []SearchRelatedProfile `json:"relatedProfiles,omitempty"`
	GroupHomes      []SearchGroupHome      `json:"groupHomes,omitempty"`
}

// SearchResultProfileItemGlobalSearch represents a profile search result for global search context.
type SearchResultProfileItemGlobalSearch struct {
	SearchResultProfileItemBase
	Relations []SearchRelatedProfile `json:"relations,omitempty"`
}

// SearchRoleDefinition represents a role definition as seen in search results.
type SearchRoleDefinition struct {
	PortalRole      *string `json:"portalRole,omitempty"`
	InstitutionRole *string `json:"institutionRole,omitempty"`
}

// SearchMainGroup represents a main group reference in search.
type SearchMainGroup struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// SearchRelatedGroup represents a related group reference.
type SearchRelatedGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// SearchRelatedProfile represents a related profile reference.
type SearchRelatedProfile struct {
	FirstName            *string          `json:"firstName,omitempty"`
	LastName             *string          `json:"lastName,omitempty"`
	ProfileID            *int64           `json:"profileId,omitempty"`
	InstitutionProfileID *int64           `json:"institutionProfileId,omitempty"`
	RelationType         *string          `json:"relationType,omitempty"`
	AulaEmail            *string          `json:"aulaEmail,omitempty"`
	MainGroup            *SearchMainGroup `json:"mainGroup,omitempty"`
	Metadata             *string          `json:"metadata,omitempty"`
}

// SearchGroupHome represents a group-home reference for recipient search.
type SearchGroupHome struct {
	Name       *string `json:"name,omitempty"`
	OtpInboxID *int64  `json:"otpInboxId,omitempty"`
	ID         *int64  `json:"id,omitempty"`
}

// SearchResultSecureFile represents a search result for a secure file.
type SearchResultSecureFile struct {
	DocID              *string                                      `json:"docId,omitempty"`
	DocType            *string                                      `json:"docType,omitempty"`
	InstitutionCode    *string                                      `json:"institutionCode,omitempty"`
	InstitutionName    *string                                      `json:"institutionName,omitempty"`
	MunicipalityCode   *string                                      `json:"municipalityCode,omitempty"`
	MunicipalityName   *string                                      `json:"municipalityName,omitempty"`
	Name               *string                                      `json:"name,omitempty"`
	Description        *string                                      `json:"description,omitempty"`
	ID                 *int64                                       `json:"id,omitempty"`
	Category           *string                                      `json:"category,omitempty"`
	ChildAssociations  []SearchResultSecureFileChildAssociation     `json:"childAssociations,omitempty"`
	GroupAssociations  []SearchResultSecureFileGroupAssociation     `json:"groupAssociations,omitempty"`
	Created            *string                                      `json:"created,omitempty"`
	Edited             *string                                      `json:"edited,omitempty"`
	CreatorName        *string                                      `json:"creatorName,omitempty"`
	Metada             *string                                      `json:"metada,omitempty"`
	Title              *string                                      `json:"title,omitempty"`
}

// SearchResultSecureFileGroupAssociation represents a group association on a secure file search result.
type SearchResultSecureFileGroupAssociation struct {
	ID             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	InstituionCode *string `json:"instituionCode,omitempty"`
}

// SearchResultSecureFileChildAssociation represents a child association on a secure file search result.
type SearchResultSecureFileChildAssociation struct {
	ProfileID *int64  `json:"profileId,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// BaseSearchResultMessageItem represents a base message search result item.
type BaseSearchResultMessageItem struct {
	DocID                       *string                      `json:"docId,omitempty"`
	DocType                     *string                      `json:"docType,omitempty"`
	InstitutionCode             *string                      `json:"institutionCode,omitempty"`
	InstitutionName             *string                      `json:"institutionName,omitempty"`
	MunicipalityCode            *string                      `json:"municipalityCode,omitempty"`
	MunicipalityName            *string                      `json:"municipalityName,omitempty"`
	Name                        *string                      `json:"name,omitempty"`
	Description                 *string                      `json:"description,omitempty"`
	Marked                      bool                         `json:"marked"`
	Muted                       bool                         `json:"muted"`
	Thread                      *SearchResultMessageThreadItem `json:"thread,omitempty"`
	LeaveTime                   *string                      `json:"leaveTime,omitempty"`
	SensitivityLevel            *int                         `json:"sensitivityLevel,omitempty"`
	Read                        bool                         `json:"read"`
	SelectedInMultiEditMode     bool                         `json:"selectedInMultiEditMode"`
	MessageDraft                *MessageDraft                `json:"messageDraft,omitempty"`
	MailBoxOwner                *RecipientApiModel           `json:"mailBoxOwner,omitempty"`
	FolderID                    *int64                       `json:"folderId,omitempty"`
	FolderName                  *string                      `json:"folderName,omitempty"`
	SubscriptionID              *int64                       `json:"subscriptionId,omitempty"`
	RegardingChildren           []MessageRegardingChildren   `json:"regardingChildren,omitempty"`
}

// SearchResultMessage represents an individual message within a search result.
type SearchResultMessage struct {
	ID                *string         `json:"id,omitempty"`
	Text              json.RawMessage `json:"text,omitempty"`
	SendDateTime      *string         `json:"sendDateTime,omitempty"`
	SenderEmail       *string         `json:"senderEmail,omitempty"`
	SenderDisplayName *string         `json:"senderDisplayName,omitempty"`
	MessageType       *string         `json:"messageType,omitempty"`
	Unread            bool            `json:"unread"`
}

// SearchResultMessageGlobalSearchItem represents a message search result for global search context.
type SearchResultMessageGlobalSearchItem struct {
	BaseSearchResultMessageItem
	Message *SearchResultMessage `json:"message,omitempty"`
}

// SearchResultMessageItemSimple represents a simple message search result.
type SearchResultMessageItemSimple struct {
	DocID                   *string            `json:"docId,omitempty"`
	DocType                 *string            `json:"docType,omitempty"`
	InstitutionCode         *string            `json:"institutionCode,omitempty"`
	InstitutionName         *string            `json:"institutionName,omitempty"`
	MunicipalityCode        *string            `json:"municipalityCode,omitempty"`
	MunicipalityName        *string            `json:"municipalityName,omitempty"`
	Name                    *string            `json:"name,omitempty"`
	Description             *string            `json:"description,omitempty"`
	MessageID               *string            `json:"messageId,omitempty"`
	Message                 *string            `json:"message,omitempty"`
	SubscriptionID          *int64             `json:"subscriptionId,omitempty"`
	Author                  *string            `json:"author,omitempty"`
	Metadata                *string            `json:"metadata,omitempty"`
	ThreadID                *int64             `json:"threadId,omitempty"`
	Title                   *string            `json:"title,omitempty"`
	StepUpRequired          bool               `json:"stepUpRequired"`
	LatestMessageSendTime   *string            `json:"latestMessageSendTime,omitempty"`
	MailBoxOwner            *RecipientApiModel `json:"mailBoxOwner,omitempty"`
}

// SearchResultMessageMessageModuleItem represents a message search result for message module.
type SearchResultMessageMessageModuleItem struct {
	BaseSearchResultMessageItem
	SearchMessage        *SearchResultMessage    `json:"searchMessage,omitempty"`
	Recipients           []MessageParticipantDto `json:"recipients,omitempty"`
	Creator              *MessageParticipantDto  `json:"creator,omitempty"`
	ExtraRecipientsCount *int64                  `json:"extraRecipientsCount,omitempty"`
}

// SearchResultMessageThreadItem represents thread info within a message search result.
type SearchResultMessageThreadItem struct {
	ID               *int64  `json:"id,omitempty"`
	Subject          *string `json:"subject,omitempty"`
	SensitivityLevel *string `json:"sensitivityLevel,omitempty"`
	IsForwarded      bool    `json:"isForwarded"`
	ThreadType       *string `json:"threadType,omitempty"`
}

// SearchResultMessagesResponse represents a response wrapper for message search.
type SearchResultMessagesResponse struct {
	TotalHits *int                                   `json:"totalHits,omitempty"`
	Results   []SearchResultMessageMessageModuleItem `json:"results,omitempty"`
}

// SearchRecipientResponse represents a response wrapper for recipient search.
type SearchRecipientResponse struct {
	TotalHits *int               `json:"totalHits,omitempty"`
	Results   []SearchResultItem `json:"results,omitempty"`
}

// ChildRelationsResponse represents a child relations response for recipient context.
type ChildRelationsResponse struct {
	ChildRelationsProfileList []json.RawMessage        `json:"childRelationsProfileList,omitempty"`
	SearchRecipientGroupList  []SearchResultGroupItem `json:"searchRecipientGroupList,omitempty"`
}

// SearchGroupItemResultModel represents a group search result model.
type SearchGroupItemResultModel struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
	Name            *string `json:"name,omitempty"`
	ID              *int64  `json:"id,omitempty"`
}

// SearchGroupResultModel represents a group search response.
type SearchGroupResultModel struct {
	Results []SearchGroupItemResultModel `json:"results,omitempty"`
}

// GlobalSearchParameters represents global search parameters.
type GlobalSearchParameters struct {
	Text                                *string  `json:"text,omitempty"`
	PageLimit                           *int     `json:"pageLimit,omitempty"`
	PageNumber                          *int     `json:"pageNumber,omitempty"`
	Limit                               *int     `json:"limit,omitempty"`
	Offset                              *int     `json:"offset,omitempty"`
	GroupID                             *int64   `json:"groupId,omitempty"`
	DocTypeCount                        bool     `json:"docTypeCount"`
	DocType                             *string  `json:"docType,omitempty"`
	GroupTypes                          []string `json:"groupTypes,omitempty"`
	InstitutionProfileIDs               []int64  `json:"institutionProfileIds,omitempty"`
	ActiveChildrenInstitutionProfileIDs []int64  `json:"activeChildrenInstitutionProfileIds,omitempty"`
}

// SearchForAssociateSecureDocumentsParameter represents secure document association search parameters.
type SearchForAssociateSecureDocumentsParameter struct {
	InstitutionCodes []string `json:"institutionCodes,omitempty"`
	Text             *string  `json:"text,omitempty"`
}

// SearchForProfilesAndGroupsParameters represents profile and group search parameters.
type SearchForProfilesAndGroupsParameters struct {
	OnlyProfiles bool    `json:"onlyProfiles"`
	Text         *string `json:"text,omitempty"`
	PortalRoles  *string `json:"portalRoles,omitempty"`
	Typeahead    bool    `json:"typeahead"`
	Limit        *int    `json:"limit,omitempty"`
}

// SearchGroupRequestModel represents a group search request.
type SearchGroupRequestModel struct {
	Text            *string  `json:"text,omitempty"`
	InstitutionCodes []string `json:"institutionCodes,omitempty"`
	Limit           *int     `json:"limit,omitempty"`
	Offset          *int     `json:"offset,omitempty"`
	FromModuleValue *string  `json:"fromModuleValue,omitempty"`
}

// SearchMessageRequestModel represents a message search request.
type SearchMessageRequestModel struct {
	Keyword        *string           `json:"keyword,omitempty"`
	ThreadSubject  *string           `json:"threadSubject,omitempty"`
	MessageContent *string           `json:"messageContent,omitempty"`
	HasAttachments *bool             `json:"hasAttachments,omitempty"`
	FromDate       *string           `json:"fromDate,omitempty"`
	ToDate         *string           `json:"toDate,omitempty"`
	ThreadCreators []json.RawMessage `json:"threadCreators,omitempty"`
	Participants   []json.RawMessage `json:"participants,omitempty"`
	Page           *int              `json:"page,omitempty"`
	CommonInboxID  *int64            `json:"commonInboxId,omitempty"`
	FolderID       *int64            `json:"folderId,omitempty"`
	Filter         *string           `json:"filter,omitempty"`
	SortType       *string           `json:"sortType,omitempty"`
	SortOrder      *string           `json:"sortOrder,omitempty"`
}

// SearchRecipientParameters represents recipient search parameters.
type SearchRecipientParameters struct {
	Text                          *string  `json:"text,omitempty"`
	FromModule                    *string  `json:"fromModule,omitempty"`
	DocTypes                      *string  `json:"docTypes,omitempty"`
	PortalRoles                   *string  `json:"portalRoles,omitempty"`
	GroupSearchScope              *string  `json:"groupSearchScope,omitempty"`
	Limit                         *int     `json:"limit,omitempty"`
	ScopeEmployeesToInstitution   *bool    `json:"scopeEmployeesToInstitution,omitempty"`
	GroupID                       *int     `json:"groupId,omitempty"`
	InstCode                      *string  `json:"instCode,omitempty"`
	InstitutionCodes              []string `json:"institutionCodes,omitempty"`
	RegardingChildren             []int64  `json:"regardingChildren,omitempty"`
	MailBoxOwnerType              *string  `json:"mailBoxOwnerType,omitempty"`
	MailBoxOwnerID                *int64   `json:"mailBoxOwnererId,omitempty"`
}

// SearchResourceParameters represents resource search parameters.
type SearchResourceParameters struct {
	Query        *string  `json:"query,omitempty"`
	InstitutionCode []string `json:"institutionCode,omitempty"`
	ExcludeTypes []string `json:"excludeTypes,omitempty"`
	IncludeTypes []string `json:"includeTypes,omitempty"`
}
