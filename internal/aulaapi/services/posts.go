package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GetPosts fetches posts matching the given filter/pagination parameters.
func GetPosts(ctx context.Context, s *aulaapi.Session, params *models.GetPostApiParameters) (models.GetPostApiResult, error) {
	var query []string
	if params.Parent != nil {
		query = append(query, fmt.Sprintf("parent=%s", *params.Parent))
	}
	if params.GroupID != nil {
		query = append(query, fmt.Sprintf("groupId=%d", *params.GroupID))
	}
	if params.IsImportant != nil {
		query = append(query, fmt.Sprintf("isImportant=%t", *params.IsImportant))
	}
	if params.CreatorPortalRole != nil {
		query = append(query, fmt.Sprintf("creatorPortalRole=%s", *params.CreatorPortalRole))
	}
	if params.InstitutionProfileIDs != nil {
		for _, id := range params.InstitutionProfileIDs {
			query = append(query, fmt.Sprintf("institutionProfileIds[]=%d", id))
		}
	}
	if params.RelatedInstitutions != nil {
		for _, inst := range params.RelatedInstitutions {
			query = append(query, fmt.Sprintf("relatedInstitutions=%s", inst))
		}
	}
	if params.OwnPost {
		query = append(query, "ownPost=true")
	}
	if params.IsUnread {
		query = append(query, "isUnread=true")
	}
	if params.IsBookmarked {
		query = append(query, "isBookmarked=true")
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	if params.Index != nil {
		query = append(query, fmt.Sprintf("index=%d", *params.Index))
	}

	path := "?method=posts.getAllPosts"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetPostApiResult](ctx, s, path)
}

// GetPostByID fetches a single post by its ID.
func GetPostByID(ctx context.Context, s *aulaapi.Session, postID int64) (models.PostApiDto, error) {
	return aulaapi.SessionGet[models.PostApiDto](ctx, s, fmt.Sprintf("?method=posts.getById&id=%d", postID))
}

// CreatePost creates a new post.
func CreatePost(ctx context.Context, s *aulaapi.Session, params *models.CreatePostApiParameter) (models.CreatePostResult, error) {
	return aulaapi.SessionPost[models.CreatePostResult](ctx, s, "?method=posts.createPost", params)
}

// EditPost edits an existing post.
func EditPost(ctx context.Context, s *aulaapi.Session, postID int64, params *models.CreatePostApiParameter) (json.RawMessage, error) {
	_ = postID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=posts.updatePost", params)
}

// DeletePost deletes a post.
func DeletePost(ctx context.Context, s *aulaapi.Session, postID int64) (json.RawMessage, error) {
	body := map[string]int64{"id": postID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=posts.deletePost", body)
}

// ReportPost reports a post (flag for moderation).
func ReportPost(ctx context.Context, s *aulaapi.Session, postID int64, params *models.ReportApiParameter) (json.RawMessage, error) {
	_ = postID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=posts.reportPost", params)
}

// BookmarkPost bookmarks a post for the current user.
func BookmarkPost(ctx context.Context, s *aulaapi.Session, postID int64) (json.RawMessage, error) {
	body := map[string]int64{"id": postID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=posts.bookmark", body)
}

// UnbookmarkPost removes bookmark from a post.
func UnbookmarkPost(ctx context.Context, s *aulaapi.Session, postID int64) (json.RawMessage, error) {
	body := map[string]int64{"id": postID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=posts.unbookmark", body)
}
