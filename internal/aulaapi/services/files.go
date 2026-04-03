package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// CreateDocumentLinks creates document links (attach secure documents as references).
func CreateDocumentLinks(ctx context.Context, s *aulaapi.Session, documentIDs []int64) ([]models.DocumentLinkResult, error) {
	return aulaapi.SessionPost[[]models.DocumentLinkResult](ctx, s, "?method=files.createDocumentLinks", documentIDs)
}

// CreateAttachments creates file/media/link attachments.
func CreateAttachments(ctx context.Context, s *aulaapi.Session, args *models.CreateAttachmentsArguments) (models.CreateAttachmentsResult, error) {
	return aulaapi.SessionPost[models.CreateAttachmentsResult](ctx, s, "?method=files.createAttachments", args)
}

// GetUploadLinks gets pre-signed upload links for one or more files.
func GetUploadLinks(ctx context.Context, s *aulaapi.Session, args *models.GetUploadLinksArguments) ([]models.UploadLink, error) {
	return aulaapi.SessionPost[[]models.UploadLink](ctx, s, "?method=files.getDownloadUrl", args)
}

// CompleteMultipartUpload completes a multipart upload after all parts have been uploaded to S3.
func CompleteMultipartUpload(ctx context.Context, s *aulaapi.Session, request *models.CompleteMultipartUploadingRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=files.completeMultipartUploading", request)
}
