package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// UpdateMasterDataRequest is the request body for PostMasterData.
type UpdateMasterDataRequest struct {
	ExternalEmail    *string `json:"externalEmail,omitempty"`
	Phonenumber      *string `json:"phonenumber,omitempty"`
	WorkPhonenumber  *string `json:"workPhonenumber,omitempty"`
	HomePhonenumber  *string `json:"homePhonenumber,omitempty"`
	MobilePhonenumber *string `json:"mobilePhonenumber,omitempty"`
}

// UpdateProfilePictureRequest is the request body for updating a profile picture.
type UpdateProfilePictureRequest struct {
	InstitutionProfileID int64  `json:"institutionProfileId"`
	Key                  string `json:"key"`
	Bucket               string `json:"bucket"`
}

// GetProfilesByLogin fetches the logged-in user's profiles after authentication.
func GetProfilesByLogin(ctx context.Context, s *aulaapi.Session) (models.OnboardingResponseDto, error) {
	return aulaapi.SessionGet[models.OnboardingResponseDto](ctx, s, "?method=profiles.getprofilesbylogin")
}

// GetProfileContext fetches the profile context for the current session.
func GetProfileContext(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, "?method=profiles.getProfileContext")
}

// GetProfileMasterData fetches profile master data for the current user.
func GetProfileMasterData(ctx context.Context, s *aulaapi.Session) (models.Profile, error) {
	return aulaapi.SessionGet[models.Profile](ctx, s, "?method=profiles.getProfileMasterData")
}

// GetOnboardingMasterData fetches onboarding master data (first-login flow).
func GetOnboardingMasterData(ctx context.Context, s *aulaapi.Session) (models.OnboardingResponseDto, error) {
	return aulaapi.SessionGet[models.OnboardingResponseDto](ctx, s, "?method=profiles.getProfilesByLogin")
}

// PostMasterData updates profile master data (contact info).
func PostMasterData(ctx context.Context, s *aulaapi.Session, request *UpdateMasterDataRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=profiles.updateProfileMasterData", request)
}

// UpdateProfilePicture updates the profile picture for an institution profile.
func UpdateProfilePicture(ctx context.Context, s *aulaapi.Session, request *UpdateProfilePictureRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=profiles.updateProfilePicture", request)
}

// KeepAlive sends a keep-alive ping to extend the backend session.
func KeepAlive(ctx context.Context, s *aulaapi.Session) error {
	_, err := aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, "?method=session.keepAlive")
	return err
}
