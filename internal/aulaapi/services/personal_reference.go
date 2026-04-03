package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
)

// GetAdditionalAnswerData gets additional answer data for personal references.
func GetAdditionalAnswerData(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, "?method=personalReferenceData.getPersonalReferenceDataAdditionalDataAnswer")
}

// GetConsentAnswerData gets consent answer data for personal references.
func GetConsentAnswerData(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, "?method=personalReferenceData.getPersonalReferenceDataConsentAnswer")
}

// GetQuestionData gets question data for personal references.
func GetQuestionData(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, "?method=personalReferenceData.getPersonalReferenceDataQuestion")
}
