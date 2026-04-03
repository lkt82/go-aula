package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// AddCommentRequestModel is the request body for AddComment.
type AddCommentRequestModel struct {
	CommentableItem      models.CommentItem `json:"commentableItem"`
	Content              string             `json:"content"`
	CreatorInstProfileID int64              `json:"creatorInstProfileId"`
}

// GetCommentsRequestModel contains query parameters for fetching comments.
type GetCommentsRequestModel struct {
	ParentType string `json:"parentType"`
	ParentID   int64  `json:"parentId"`
	StartIndex *int32 `json:"startIndex,omitempty"`
	Limit      *int32 `json:"limit,omitempty"`
}

// AddComment adds a comment to a post, media item, or other commentable entity.
func AddComment(ctx context.Context, s *aulaapi.Session, request *AddCommentRequestModel) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=comments.addComment", request)
}

// UpdateComment updates an existing comment's content.
func UpdateComment(ctx context.Context, s *aulaapi.Session, commentID int64, request *models.UpdateCommentRequestModel) (json.RawMessage, error) {
	_ = commentID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=comments.updateComment", request)
}

// GetComments fetches comments for a given parent item (post, media, etc.).
func GetComments(ctx context.Context, s *aulaapi.Session, params *GetCommentsRequestModel) (models.PagedCommentList, error) {
	var query []string
	query = append(query, fmt.Sprintf("parentType=%s", params.ParentType))
	query = append(query, fmt.Sprintf("parentId=%d", params.ParentID))
	if params.StartIndex != nil {
		query = append(query, fmt.Sprintf("startIndex=%d", *params.StartIndex))
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	path := "?method=comments.getComments&" + strings.Join(query, "&")
	return aulaapi.SessionGet[models.PagedCommentList](ctx, s, path)
}

// ReportComment reports a comment for moderation.
func ReportComment(ctx context.Context, s *aulaapi.Session, commentID int64, params *models.ReportCommentApiParameters) (json.RawMessage, error) {
	_ = commentID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=comments.reportComment", params)
}

// DeleteComment deletes a comment.
func DeleteComment(ctx context.Context, s *aulaapi.Session, commentID int64, request *models.DeleteCommentRequestModel) (json.RawMessage, error) {
	_ = commentID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=comments.deleteComment", request)
}
