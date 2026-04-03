package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkt82/go-aula/internal/aulaapi"
)

// AdditionalMasterData represents extra profile information beyond the core profile fields.
type AdditionalMasterData struct {
	InstitutionProfileID *int64          `json:"institutionProfileId,omitempty"`
	Data                 json.RawMessage `json:"data"`
}

// UpdateAdditionalMasterDataRequest is the request body for updating additional master data.
type UpdateAdditionalMasterDataRequest struct {
	InstitutionProfileID int64           `json:"institutionProfileId"`
	Data                 json.RawMessage `json:"data"`
}

// UpdateAdditionalMasterDataEmployeeRequest is the request body for updating employee-specific additional master data.
type UpdateAdditionalMasterDataEmployeeRequest struct {
	InstitutionProfileID int64           `json:"institutionProfileId"`
	Data                 json.RawMessage `json:"data"`
}

// GetAdditionalMasterData gets additional master data for the current user.
func GetAdditionalMasterData(ctx context.Context, s *aulaapi.Session) ([]AdditionalMasterData, error) {
	return aulaapi.SessionGet[[]AdditionalMasterData](ctx, s, "?method=profiles.getAdditionalDataResponsesForOwner")
}

// GetByInstitutionProfileID gets additional master data for a specific institution profile.
func GetByInstitutionProfileID(ctx context.Context, s *aulaapi.Session, institutionProfileID int64) (AdditionalMasterData, error) {
	return aulaapi.SessionGet[AdditionalMasterData](ctx, s, fmt.Sprintf("?method=profiles.getAdditionalDataResponsesByInstitutionProfileIds&institutionProfileId=%d", institutionProfileID))
}

// PostAdditionalMasterData updates additional master data.
func PostAdditionalMasterData(ctx context.Context, s *aulaapi.Session, request *UpdateAdditionalMasterDataRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=profiles.updateAdditionalDataResponses", request)
}

// PostAdditionalMasterDataEmployee updates employee-specific additional master data.
func PostAdditionalMasterDataEmployee(ctx context.Context, s *aulaapi.Session, request *UpdateAdditionalMasterDataEmployeeRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=profiles.updateAdditionalDataResponsesEmployee", request)
}
