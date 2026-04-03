package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// UpdateDocumentLockedStatusRequest is the request body for updating a document's locked status.
type UpdateDocumentLockedStatusRequest struct {
	IsLocked bool `json:"isLocked"`
}

// GetSecureDocuments fetches secure documents matching the given filter arguments.
func GetSecureDocuments(ctx context.Context, s *aulaapi.Session, args *models.GetSecureDocumentsArguments) (models.GetSecureDocumentsResult, error) {
	return aulaapi.SessionPost[models.GetSecureDocumentsResult](ctx, s, "?method=documents.getSecureDocuments", args)
}

// GetCommonFiles fetches common (institution-level) files.
func GetCommonFiles(ctx context.Context, s *aulaapi.Session, args *models.GetCommonFilesArguments) ([]models.CommonFileDto, error) {
	var query []string
	if args.Page != nil {
		query = append(query, fmt.Sprintf("page=%d", *args.Page))
	}
	if args.SortType != nil {
		query = append(query, fmt.Sprintf("sortType=%s", *args.SortType))
	}
	if args.SortOrder != nil {
		query = append(query, fmt.Sprintf("sortOrder=%s", *args.SortOrder))
	}
	path := "?method=commonFiles.getCommonFiles"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.CommonFileDto](ctx, s, path)
}

// UpdateSharings updates sharings on one or more secure documents.
func UpdateSharings(ctx context.Context, s *aulaapi.Session, args *models.UpdateSharingArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=documents.updateSharings", args)
}

// RemoveOwnSharings removes the current user's own sharings from documents.
func RemoveOwnSharings(ctx context.Context, s *aulaapi.Session, args *models.RemoveSharingArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=documents.removeOwnSharings", args)
}

// GetImplicitSharings gets implicit sharings for a document.
func GetImplicitSharings(ctx context.Context, s *aulaapi.Session, documentID int64) (models.GetImplicitSharingsDto, error) {
	return aulaapi.SessionGet[models.GetImplicitSharingsDto](ctx, s, fmt.Sprintf("?method=documents.getImplicitSharings&documentId=%d", documentID))
}

// GetDocumentRevisions gets revision history for a document.
func GetDocumentRevisions(ctx context.Context, s *aulaapi.Session, documentID int64, page *int32) (models.DocumentRevisionPageDto, error) {
	path := fmt.Sprintf("?method=documents.getDocumentRevisions&documentId=%d", documentID)
	if page != nil {
		path += fmt.Sprintf("&page=%d", *page)
	}
	return aulaapi.SessionGet[models.DocumentRevisionPageDto](ctx, s, path)
}

// GetExternalDocumentDetails gets details of an external secure document.
func GetExternalDocumentDetails(ctx context.Context, s *aulaapi.Session, documentID int64) (models.ExternalSecureDocumentDetailsDto, error) {
	return aulaapi.SessionGet[models.ExternalSecureDocumentDetailsDto](ctx, s, fmt.Sprintf("?method=documents.getExternalSecureFile&documentId=%d", documentID))
}

// GetExternalDocumentRevision gets a specific revision of an external secure document.
func GetExternalDocumentRevision(ctx context.Context, s *aulaapi.Session, documentID int64) (models.ExternalSecureDocumentDetailsDto, error) {
	return aulaapi.SessionGet[models.ExternalSecureDocumentDetailsDto](ctx, s, fmt.Sprintf("?method=documents.getDocumentRevision&documentId=%d", documentID))
}

// GetInternalDocumentDetails gets details of an internal secure document.
func GetInternalDocumentDetails(ctx context.Context, s *aulaapi.Session, documentID int64) (models.InternalSecureDocumentDetailsDto, error) {
	return aulaapi.SessionGet[models.InternalSecureDocumentDetailsDto](ctx, s, fmt.Sprintf("?method=documents.getInternalSecureDocument&documentId=%d", documentID))
}

// GetInternalDocumentRevision gets a specific revision of an internal secure document.
func GetInternalDocumentRevision(ctx context.Context, s *aulaapi.Session, documentID int64) (models.InternalSecureDocumentDetailsDto, error) {
	return aulaapi.SessionGet[models.InternalSecureDocumentDetailsDto](ctx, s, fmt.Sprintf("?method=documents.getDocumentRevision&documentId=%d", documentID))
}

// CreateInternalSecureDocument creates a new internal secure document.
func CreateInternalSecureDocument(ctx context.Context, s *aulaapi.Session, args *models.CreateInternalDocumentArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=documents.createInternalSecureDocument", args)
}

// UpdateInternalSecureDocument updates an existing internal secure document.
func UpdateInternalSecureDocument(ctx context.Context, s *aulaapi.Session, documentID int64, args *models.CreateInternalDocumentArguments) (json.RawMessage, error) {
	_ = documentID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=documents.updateInternalSecureDocument", args)
}

// UpdateDocumentLockedStatus locks or unlocks a secure document.
func UpdateDocumentLockedStatus(ctx context.Context, s *aulaapi.Session, documentID int64, isLocked bool) (json.RawMessage, error) {
	_ = documentID
	body := UpdateDocumentLockedStatusRequest{IsLocked: isLocked}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=documents.updateLockedStatus", &body)
}

// SoftDeleteSecureDocument soft deletes a secure document.
func SoftDeleteSecureDocument(ctx context.Context, s *aulaapi.Session, documentID int64) (json.RawMessage, error) {
	body := map[string]int64{"documentId": documentID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=documents.deleteDocument", body)
}

// GetShareableSecureDocuments gets shareable secure documents matching the given filter.
func GetShareableSecureDocuments(ctx context.Context, s *aulaapi.Session, args *models.GetShareableSecureDocumentsArguments) (models.GetSecureDocumentsResult, error) {
	var query []string
	if args.FilterInstitutionProfileIDs != nil {
		for _, id := range args.FilterInstitutionProfileIDs {
			query = append(query, fmt.Sprintf("filterInstitutionProfileIds=%d", id))
		}
	}
	if args.ShareToInstitutionProfileIDs != nil {
		for _, id := range args.ShareToInstitutionProfileIDs {
			query = append(query, fmt.Sprintf("shareToInstitutionProfileIds=%d", id))
		}
	}
	if args.Index != nil {
		query = append(query, fmt.Sprintf("index=%d", *args.Index))
	}
	if args.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *args.Limit))
	}
	path := "?method=documents.getShareableSecureDocuments"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetSecureDocumentsResult](ctx, s, path)
}

// GetMaxDocumentsPerExport gets the maximum number of documents allowed per export.
func GetMaxDocumentsPerExport(ctx context.Context, s *aulaapi.Session) (int32, error) {
	return aulaapi.SessionGet[int32](ctx, s, "?method=documents.getMaxDocumentsPerExport")
}

// CreateExportForMultiple creates a bulk export for multiple secure documents.
func CreateExportForMultiple(ctx context.Context, s *aulaapi.Session, request *models.CreateExportForMultipleSecureDocumentsRequest) (models.SecureDocumentExportDto, error) {
	return aulaapi.SessionPost[models.SecureDocumentExportDto](ctx, s, "?method=documents.createArchiveForMultipleSecureDocuments", request)
}

// TrackExport tracks the status of a multi-document export.
func TrackExport(ctx context.Context, s *aulaapi.Session, exportJobID int64) (models.SecureDocumentExportDto, error) {
	return aulaapi.SessionGet[models.SecureDocumentExportDto](ctx, s, fmt.Sprintf("?method=documents.trackCreateArchiveForMultipleSecureDocumentsRequest&exportJobId=%d", exportJobID))
}

// CreatePDFForSingle creates a PDF for a single secure document.
func CreatePDFForSingle(ctx context.Context, s *aulaapi.Session, documentID int64) (models.SecureDocumentExportDto, error) {
	body := map[string]int64{"documentId": documentID}
	return aulaapi.SessionPost[models.SecureDocumentExportDto](ctx, s, "?method=documents.createPDFForSingleDocument", body)
}

// TrackCreatePDF tracks the status of a single document PDF generation.
func TrackCreatePDF(ctx context.Context, s *aulaapi.Session, documentID int64) (models.SecureDocumentExportDto, error) {
	return aulaapi.SessionGet[models.SecureDocumentExportDto](ctx, s, fmt.Sprintf("?method=documents.trackCreatePDFForSingleDocument&documentId=%d", documentID))
}
