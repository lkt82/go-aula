package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GetConsents gets consent status for all institution profiles.
func GetConsents(ctx context.Context, s *aulaapi.Session) ([]models.InstitutionProfileConsentDto, error) {
	return aulaapi.SessionGet[[]models.InstitutionProfileConsentDto](ctx, s, "?method=consents.getConsentResponses")
}

// PostConsents submits consent responses for an institution profile.
func PostConsents(ctx context.Context, s *aulaapi.Session, updates *models.ProfileConsentUpdatesDto) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=consents.updateConsentResponses", updates)
}
