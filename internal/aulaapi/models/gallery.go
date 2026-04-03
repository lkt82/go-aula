package models

// AlbumCreatorDto represents an album creator profile.
type AlbumCreatorDto struct {
	ID              *int64                         `json:"id,omitempty"`
	InstitutionCode *string                        `json:"institutionCode,omitempty"`
	InstitutionName *string                        `json:"institutionName,omitempty"`
	Name            *string                        `json:"name,omitempty"`
	ShortName       *string                        `json:"shortName,omitempty"`
	Metadata        *string                        `json:"metadata,omitempty"`
	ProfileID       *int64                         `json:"profileId,omitempty"`
	Role            *string                        `json:"role,omitempty"`
	ProfilePicture  *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
}

// AlbumGroupDto represents a group context in album views.
type AlbumGroupDto struct {
	ID              *int64  `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	InstitutionCode *string `json:"institutionCode,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
	Role            *string `json:"role,omitempty"`
	MainGroup       *bool   `json:"mainGroup,omitempty"`
}

// AlbumDto represents a full album with metadata, thumbnails, and permissions.
type AlbumDto struct {
	ID                      *int64              `json:"id,omitempty"`
	Title                   *string             `json:"title,omitempty"`
	Name                    *string             `json:"name,omitempty"`
	Creator                 *AlbumCreatorDto    `json:"creator,omitempty"`
	CreationDate            *string             `json:"creationDate,omitempty"`
	TotalSize               *int                `json:"totalSize,omitempty"`
	Size                    *int                `json:"size,omitempty"`
	From                    *int                `json:"from,omitempty"`
	Description             *string             `json:"description,omitempty"`
	SharedWithGroups        []ShareWithGroupDto `json:"sharedWithGroups,omitempty"`
	ThumbnailsUrls          []string            `json:"thumbnailsUrls,omitempty"`
	CurrentUserCanEdit      bool                `json:"currentUserCanEdit"`
	CurrentUserCanDelete    bool                `json:"currentUserCanDelete"`
	CurrentUserCanAddMedia  bool                `json:"currentUserCanAddMedia"`
}

// MediaCreatorModel represents a media creator profile.
type MediaCreatorModel struct {
	ID              *int64  `json:"id,omitempty"`
	InstitutionCode *string `json:"institutionCode,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
	Role            *string `json:"role,omitempty"`
	ProfileID       *int64  `json:"profileId,omitempty"`
	Name            *string `json:"name,omitempty"`
	ShortName       *string `json:"shortName,omitempty"`
	Metadata        *string `json:"metadata,omitempty"`
}

// MediaListDto represents a list of media items.
type MediaListDto struct {
	Results []FilesAulaFileResultDto `json:"results,omitempty"`
	Album   *AlbumDto                `json:"album,omitempty"`
}

// MediasInAlbumDto represents media items within an album.
type MediasInAlbumDto struct {
	Results    []AulaGalleryMediaFileResultDto `json:"results,omitempty"`
	Album      *AlbumDto                       `json:"album,omitempty"`
	MediaCount *int                            `json:"mediaCount,omitempty"`
}

// CreateAlbumParameters represents parameters for creating an album.
type CreateAlbumParameters struct {
	Title                         *string                   `json:"title,omitempty"`
	AlbumID                       *int64                    `json:"albumId,omitempty"`
	CreatorInstitutionProfileID   *int64                    `json:"creatorInstitutionProfileId,omitempty"`
	SharedWithGroups              []LinkedGroupRequestModel `json:"sharedWithGroups,omitempty"`
	Description                   *string                   `json:"description,omitempty"`
}

// DeleteAlbumParameters represents parameters for deleting albums.
type DeleteAlbumParameters struct {
	AlbumIDs []int64 `json:"albumIds,omitempty"`
}

// GalleryViewFilter represents a filter/sort for gallery view.
type GalleryViewFilter struct {
	SelectedInstitutionCodeForFilter *string `json:"selectedInstitutionCodeForFilter,omitempty"`
	AlbumID                          *int64  `json:"albumId,omitempty"`
	UserSpecificAlbum                *bool   `json:"userSpecificAlbum,omitempty"`
	Limit                            *int    `json:"limit,omitempty"`
	Index                            *int    `json:"index,omitempty"`
	SortOn                           *string `json:"sortOn,omitempty"`
	OrderDirection                   *string `json:"orderDirection,omitempty"`
	FilterBy                         *string `json:"filterBy,omitempty"`
}

// GetMediaInAlbumFilter represents a filter/sort for media-in-album view.
type GetMediaInAlbumFilter struct {
	AlbumID                   *int64  `json:"albumId,omitempty"`
	UserSpecificAlbum         *bool   `json:"userSpecificAlbum,omitempty"`
	Limit                     *int    `json:"limit,omitempty"`
	Index                     *int    `json:"index,omitempty"`
	SortOn                    *string `json:"sortOn,omitempty"`
	OrderDirection            *string `json:"orderDirection,omitempty"`
	FilterBy                  *string `json:"filterBy,omitempty"`
	IsSelectionMode           bool    `json:"isSelectionMode"`
	SelectedInstitutionCode   *string `json:"selectedInstitutionCode,omitempty"`
}
