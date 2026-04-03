package models

import "encoding/json"

// ProfileApiDto represents profile info in post context.
type ProfileApiDto struct {
	InstitutionProfileID *int64                         `json:"institutionProfileId,omitempty"`
	ProfileID            *int64                         `json:"profileId,omitempty"`
	FirstName            *string                        `json:"firstName,omitempty"`
	LastName             *string                        `json:"lastName,omitempty"`
	FullName             *string                        `json:"fullName,omitempty"`
	Role                 *string                        `json:"role,omitempty"`
	ShortName            *string                        `json:"shortName,omitempty"`
	MainGroupName        *string                        `json:"mainGroupName,omitempty"`
	Metadata             *string                        `json:"metadata,omitempty"`
	Institution          json.RawMessage                `json:"institution,omitempty"`
	ProfilePicture       *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
}

// PostApiDto represents a post on the activity feed.
type PostApiDto struct {
	ID                      *int64                   `json:"id,omitempty"`
	Title                   *string                  `json:"title,omitempty"`
	Content                 *RichTextWrapperDto      `json:"content,omitempty"`
	CommentCount            *int                     `json:"commentCount,omitempty"`
	TimeStamp               *string                  `json:"timeStamp,omitempty"`
	OwnerProfile            *ProfileApiDto           `json:"ownerProfile,omitempty"`
	AllowComments           bool                     `json:"allowComments"`
	IsImportant             bool                     `json:"isImportant"`
	ImportantFrom           *string                  `json:"importantFrom,omitempty"`
	ImportantTo             *string                  `json:"importantTo,omitempty"`
	RelatedProfiles         []ProfileApiDto          `json:"relatedProfiles,omitempty"`
	SharedWithGroups        []ShareWithGroupDto      `json:"sharedWithGroups,omitempty"`
	Attachments             []FilesAulaFileResultDto `json:"attachments,omitempty"`
	CanCurrentUserReport    bool                     `json:"canCurrentUserReport"`
	CanCurrentUserDelete    bool                     `json:"canCurrentUserDelete"`
	CanCurrentUserComment   bool                     `json:"canCurrentUserComment"`
	PublishAt               *string                  `json:"publishAt,omitempty"`
	ExpireAt                *string                  `json:"expireAt,omitempty"`
	EditedAt                *string                  `json:"editedAt,omitempty"`
	IsBookmarked            bool                     `json:"isBookmarked"`
}

// CreatePostApiParameter represents parameters for creating or updating a post.
type CreatePostApiParameter struct {
	ID                            *int64                    `json:"id,omitempty"`
	Title                         *string                   `json:"title,omitempty"`
	Content                       *string                   `json:"content,omitempty"`
	InstitutionCode               *string                   `json:"institutionCode,omitempty"`
	CreatorInstitutionProfileID   *int64                    `json:"creatorInstitutionProfileId,omitempty"`
	AllowComments                 bool                      `json:"allowComments"`
	IsImportant                   bool                      `json:"isImportant"`
	ImportantFrom                 *string                   `json:"importantFrom,omitempty"`
	ImportantTo                   *string                   `json:"importantTo,omitempty"`
	SharedWithGroups              []LinkedGroupRequestModel `json:"sharedWithGroups,omitempty"`
	AttachmentIDs                 []int64                   `json:"attachmentIds,omitempty"`
	PublishAt                     *string                   `json:"publishAt,omitempty"`
	ExpireAt                      *string                   `json:"expireAt,omitempty"`
}

// GetPostApiParameters represents parameters for filtering/querying posts.
type GetPostApiParameters struct {
	Parent                *string  `json:"parent,omitempty"`
	GroupID               *int64   `json:"groupId,omitempty"`
	IsImportant           *bool    `json:"isImportant,omitempty"`
	CreatorPortalRole     *string  `json:"creatorPortalRole,omitempty"`
	InstitutionProfileIDs []int64  `json:"institutionProfileIds,omitempty"`
	RelatedInstitutions   []string `json:"relatedInstitutions,omitempty"`
	OwnPost               bool     `json:"ownPost"`
	IsUnread              bool     `json:"isUnread"`
	IsBookmarked          bool     `json:"isBookmarked"`
	Limit                 *int     `json:"limit,omitempty"`
	Index                 *int     `json:"index,omitempty"`
}

// GetPostApiResult represents a paginated result of post queries.
type GetPostApiResult struct {
	HasMorePosts    bool         `json:"hasMorePosts"`
	PaginationStart *string     `json:"paginationStart,omitempty"`
	PaginationEnd   *string     `json:"paginationEnd,omitempty"`
	Page            *int         `json:"page,omitempty"`
	Posts           []PostApiDto `json:"posts,omitempty"`
}

// CreatePostResult represents a result of creating a post.
type CreatePostResult struct {
	AllImagesHasValidConsents bool `json:"allImagesHasValidConsents"`
}

// ReportApiParameter represents parameters for reporting a post.
type ReportApiParameter struct {
	ID           *int64  `json:"id,omitempty"`
	ReportReason *string `json:"reportReason,omitempty"`
}

// CommentResultModel represents a comment on a post or media item.
type CommentResultModel struct {
	ID           *int64                          `json:"id,omitempty"`
	Creator      *DocumentsSimpleInstitutionProfile `json:"creator,omitempty"`
	Content      *string                         `json:"content,omitempty"`
	CreatedAt    *string                         `json:"createdAt,omitempty"`
	UpdatedAt    *string                         `json:"updatedAt,omitempty"`
	Comments     []CommentResultModel            `json:"comments,omitempty"`
	CommentCount *int                            `json:"commentCount,omitempty"`
	CanDelete    bool                            `json:"canDelete"`
	CanReport    bool                            `json:"canReport"`
	IsDeleted    bool                            `json:"isDeleted"`
	IsReported   bool                            `json:"isReported"`
}

// CommentableInstitutionProfile represents an institution profile eligible to comment.
type CommentableInstitutionProfile struct {
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
	IsSelected           bool                           `json:"isSelected"`
}

// PagedCommentList represents a paginated comment list.
type PagedCommentList struct {
	StartIndex                       *int                            `json:"startIndex,omitempty"`
	Limit                            *int                            `json:"limit,omitempty"`
	TotalResultCount                 *int                            `json:"totalResultCount,omitempty"`
	Comments                         []CommentResultModel            `json:"comments,omitempty"`
	CommentableInstitutionProfiles   []CommentableInstitutionProfile `json:"commentableInstitutionProfiles,omitempty"`
}

// CommentItem represents a comment target reference.
type CommentItem struct {
	CommentType *string `json:"type,omitempty"`
	ID          *int64  `json:"id,omitempty"`
}

// DeleteCommentRequestModel represents a request to delete a comment.
type DeleteCommentRequestModel struct {
	CommentID  *int64  `json:"commentId,omitempty"`
	ParentType *string `json:"parentType,omitempty"`
}

// ReportCommentApiParameters represents parameters for reporting a comment.
type ReportCommentApiParameters struct {
	CommentID    *int64  `json:"commentId,omitempty"`
	ReportReason *string `json:"reportReason,omitempty"`
}

// UpdateCommentRequestModel represents a request to update a comment.
type UpdateCommentRequestModel struct {
	CommentID *int64  `json:"commentId,omitempty"`
	Content   *string `json:"content,omitempty"`
}
