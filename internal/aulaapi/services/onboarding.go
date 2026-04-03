package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
)

// PolicyFileCreator is the creator of a policy file.
type PolicyFileCreator struct {
	Name            *string `json:"name,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
}

// PolicyFileInfo is file metadata within a policy entry.
type PolicyFileInfo struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// PolicyFile is the file wrapper in a policy entry.
type PolicyFile struct {
	ID      *int64          `json:"id,omitempty"`
	Name    *string         `json:"name,omitempty"`
	Status  *string         `json:"status,omitempty"`
	Creator *PolicyFileCreator `json:"creator,omitempty"`
	File    *PolicyFileInfo `json:"file,omitempty"`
}

// PolicyCommonFile is the commonFile object in a policy entry.
type PolicyCommonFile struct {
	ID    *int64      `json:"id,omitempty"`
	Title *string     `json:"title,omitempty"`
	File  *PolicyFile `json:"file,omitempty"`
}

// PolicyInstitution is the institution in a policy entry.
type PolicyInstitution struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
	MunicipalityName *string `json:"municipalityName,omitempty"`
	Type            *string `json:"type,omitempty"`
}

// PolicyEntry is a single policy entry returned by getPersonalDataPolicies.
type PolicyEntry struct {
	CommonFile  *PolicyCommonFile  `json:"commonFile,omitempty"`
	Institution *PolicyInstitution `json:"institution,omitempty"`
}

// MarkOnboardingComplete marks the onboarding flow as complete for the current profile.
func MarkOnboardingComplete(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, "?method=profiles.markOnboardingCompleted")
}

// GetPolicyLinks gets policy links shown during onboarding (data policy, terms, etc.).
func GetPolicyLinks(ctx context.Context, s *aulaapi.Session) ([]PolicyEntry, error) {
	return aulaapi.SessionGet[[]PolicyEntry](ctx, s, "?method=CommonFiles.getPersonalDataPolicies")
}
