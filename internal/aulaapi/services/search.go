package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GlobalSearch performs a global search across all content types.
func GlobalSearch(ctx context.Context, s *aulaapi.Session, params *models.GlobalSearchParameters) (models.SearchResponse, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.PageLimit != nil {
		query = append(query, fmt.Sprintf("pageLimit=%d", *params.PageLimit))
	}
	if params.PageNumber != nil {
		query = append(query, fmt.Sprintf("pageNumber=%d", *params.PageNumber))
	}
	if params.GroupID != nil {
		query = append(query, fmt.Sprintf("groupId=%d", *params.GroupID))
	}
	if params.DocTypeCount {
		query = append(query, "docTypeCount=true")
	}
	if params.DocType != nil {
		query = append(query, fmt.Sprintf("docType=%s", EncodeValue(*params.DocType)))
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	if params.Offset != nil {
		query = append(query, fmt.Sprintf("offset=%d", *params.Offset))
	}
	for _, id := range params.InstitutionProfileIDs {
		query = append(query, fmt.Sprintf("institutionProfileIds[]=%d", id))
	}
	for _, id := range params.ActiveChildrenInstitutionProfileIDs {
		query = append(query, fmt.Sprintf("activeChildrenInstitutionProfileIds[]=%d", id))
	}
	path := "?method=search.findGeneric"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchResponse](ctx, s, path)
}

// SearchForMessages searches for messages.
func SearchForMessages(ctx context.Context, s *aulaapi.Session, params *models.SearchMessageRequestModel) (models.SearchResultMessagesResponse, error) {
	return aulaapi.SessionPost[models.SearchResultMessagesResponse](ctx, s, "?method=search.findMessage", params)
}

// SearchForProfiles searches for profiles.
func SearchForProfiles(ctx context.Context, s *aulaapi.Session, params *models.SearchForProfilesAndGroupsParameters) (models.SearchResponse, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.OnlyProfiles {
		query = append(query, "onlyProfiles=true")
	}
	if params.Typeahead {
		query = append(query, "typeahead=true")
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	path := "?method=search.findProfiles"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchResponse](ctx, s, path)
}

// SearchForProfilesAndGroups searches for profiles and groups combined.
func SearchForProfilesAndGroups(ctx context.Context, s *aulaapi.Session, params *models.SearchForProfilesAndGroupsParameters) (models.SearchResponse, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.OnlyProfiles {
		query = append(query, "onlyProfiles=true")
	}
	if params.Typeahead {
		query = append(query, "typeahead=true")
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	path := "?method=search.findProfilesAndGroups"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchResponse](ctx, s, path)
}

// SearchForRecipients searches for message recipients.
func SearchForRecipients(ctx context.Context, s *aulaapi.Session, params *models.SearchRecipientParameters) (models.SearchRecipientResponse, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	if params.InstCode != nil {
		query = append(query, fmt.Sprintf("instCode=%s", EncodeValue(*params.InstCode)))
	}
	path := "?method=search.findRecipients"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchRecipientResponse](ctx, s, path)
}

// SearchForRecipientsForPersonalReference searches for recipients for personal reference.
func SearchForRecipientsForPersonalReference(ctx context.Context, s *aulaapi.Session, params *models.SearchRecipientParameters) (models.SearchRecipientResponse, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	path := "?method=search.findRecipientsPersonalReferenceData"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchRecipientResponse](ctx, s, path)
}

// SearchForRecipientsForSecureDocument searches for recipients for secure document sharing.
func SearchForRecipientsForSecureDocument(ctx context.Context, s *aulaapi.Session, params *models.SearchRecipientParameters) (models.SearchRecipientResponse, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	path := "?method=search.findProfilesAndGroupsToShareDocument"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchRecipientResponse](ctx, s, path)
}

// SearchForGroupsToAssociateDocument searches for groups to associate with a document.
func SearchForGroupsToAssociateDocument(ctx context.Context, s *aulaapi.Session, params *models.SearchForAssociateSecureDocumentsParameter) (models.SearchGroupResultModel, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.InstitutionCodes != nil {
		for _, code := range params.InstitutionCodes {
			query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
		}
	}
	path := "?method=search.findProfilesAndGroupsToAssociateDocument"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchGroupResultModel](ctx, s, path)
}

// SearchGroups searches for groups.
func SearchGroups(ctx context.Context, s *aulaapi.Session, params *models.SearchGroupRequestModel) (models.SearchGroupResultModel, error) {
	var query []string
	if params.Text != nil {
		query = append(query, fmt.Sprintf("text=%s", EncodeValue(*params.Text)))
	}
	if params.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *params.Limit))
	}
	if params.Offset != nil {
		query = append(query, fmt.Sprintf("offset=%d", *params.Offset))
	}
	if params.InstitutionCodes != nil {
		for _, code := range params.InstitutionCodes {
			query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
		}
	}
	path := "?method=search.findGroups"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.SearchGroupResultModel](ctx, s, path)
}
