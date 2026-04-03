package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GetAlbums fetches albums matching the given filter.
func GetAlbums(ctx context.Context, s *aulaapi.Session, filter *models.GalleryViewFilter) ([]models.AlbumDto, error) {
	var query []string
	if filter.SelectedInstitutionCodeForFilter != nil {
		query = append(query, fmt.Sprintf("selectedInstitutionCodeForFilter=%s", EncodeValue(*filter.SelectedInstitutionCodeForFilter)))
	}
	if filter.AlbumID != nil {
		query = append(query, fmt.Sprintf("albumId=%d", *filter.AlbumID))
	}
	if filter.UserSpecificAlbum != nil {
		query = append(query, fmt.Sprintf("userSpecificAlbum=%t", *filter.UserSpecificAlbum))
	}
	if filter.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *filter.Limit))
	}
	if filter.Index != nil {
		query = append(query, fmt.Sprintf("index=%d", *filter.Index))
	}
	if filter.SortOn != nil {
		query = append(query, fmt.Sprintf("sortOn=%s", EncodeValue(*filter.SortOn)))
	}
	if filter.OrderDirection != nil {
		query = append(query, fmt.Sprintf("orderDirection=%s", EncodeValue(*filter.OrderDirection)))
	}
	if filter.FilterBy != nil {
		query = append(query, fmt.Sprintf("filterBy=%s", EncodeValue(*filter.FilterBy)))
	}
	path := "?method=gallery.getAlbums"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.AlbumDto](ctx, s, path)
}

// GetAlbumsCached fetches albums with caching hint (same API call as GetAlbums).
func GetAlbumsCached(ctx context.Context, s *aulaapi.Session, filter *models.GalleryViewFilter) ([]models.AlbumDto, error) {
	return GetAlbums(ctx, s, filter)
}

// GetMediasInAlbum fetches media items in a specific album.
func GetMediasInAlbum(ctx context.Context, s *aulaapi.Session, filter *models.GetMediaInAlbumFilter) (models.MediasInAlbumDto, error) {
	albumID := int64(0)
	if filter.AlbumID != nil {
		albumID = *filter.AlbumID
	}
	var query []string
	if filter.UserSpecificAlbum != nil {
		query = append(query, fmt.Sprintf("userSpecificAlbum=%t", *filter.UserSpecificAlbum))
	}
	if filter.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *filter.Limit))
	}
	if filter.Index != nil {
		query = append(query, fmt.Sprintf("index=%d", *filter.Index))
	}
	if filter.SortOn != nil {
		query = append(query, fmt.Sprintf("sortOn=%s", EncodeValue(*filter.SortOn)))
	}
	if filter.OrderDirection != nil {
		query = append(query, fmt.Sprintf("orderDirection=%s", EncodeValue(*filter.OrderDirection)))
	}
	if filter.FilterBy != nil {
		query = append(query, fmt.Sprintf("filterBy=%s", EncodeValue(*filter.FilterBy)))
	}
	if filter.IsSelectionMode {
		query = append(query, "isSelectionMode=true")
	}
	if filter.SelectedInstitutionCode != nil {
		query = append(query, fmt.Sprintf("selectedInstitutionCode=%s", EncodeValue(*filter.SelectedInstitutionCode)))
	}
	path := fmt.Sprintf("?method=gallery.getMedia&albumId=%d", albumID)
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.MediasInAlbumDto](ctx, s, path)
}

// GetMediasInAlbumCached fetches media items in an album with caching hint (same API call).
func GetMediasInAlbumCached(ctx context.Context, s *aulaapi.Session, filter *models.GetMediaInAlbumFilter) (models.MediasInAlbumDto, error) {
	return GetMediasInAlbum(ctx, s, filter)
}

// GetMediaByID fetches a single media item by ID.
func GetMediaByID(ctx context.Context, s *aulaapi.Session, mediaID int64) (models.MediaListDto, error) {
	return aulaapi.SessionGet[models.MediaListDto](ctx, s, fmt.Sprintf("?method=gallery.getMediaById&id=%d", mediaID))
}

// CreateAlbum creates a new album.
func CreateAlbum(ctx context.Context, s *aulaapi.Session, params *models.CreateAlbumParameters) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.createAlbum", params)
}

// UpdateAlbum updates an existing album.
func UpdateAlbum(ctx context.Context, s *aulaapi.Session, albumID int64, params *models.CreateAlbumParameters) (json.RawMessage, error) {
	_ = albumID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.updateAlbum", params)
}

// DeleteAlbum deletes an album.
func DeleteAlbum(ctx context.Context, s *aulaapi.Session, albumID int64) (json.RawMessage, error) {
	body := map[string]int64{"albumId": albumID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.deleteAlbums", body)
}

// DeleteMedia deletes a media item.
func DeleteMedia(ctx context.Context, s *aulaapi.Session, mediaID int64) (json.RawMessage, error) {
	body := map[string]int64{"mediaId": mediaID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.deleteMedia", body)
}

// AddTag adds a tag (person tag) to a media item.
func AddTag(ctx context.Context, s *aulaapi.Session, mediaID int64, params *models.AddOrRemoveTagArguments) (json.RawMessage, error) {
	_ = mediaID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.addTag", params)
}

// RemoveTag removes a tag from a media item.
func RemoveTag(ctx context.Context, s *aulaapi.Session, mediaID, tagID int64) (json.RawMessage, error) {
	body := map[string]int64{"mediaId": mediaID, "tagId": tagID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.removeTag", body)
}

// ReportMedia reports a media item for moderation.
func ReportMedia(ctx context.Context, s *aulaapi.Session, mediaID int64, params *models.ReportApiParameter) (json.RawMessage, error) {
	_ = mediaID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=gallery.reportMedia", params)
}
