package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// AppDeprecatedResponse is the app deprecation status response.
type AppDeprecatedResponse struct {
	IsDeprecated bool    `json:"isDeprecated"`
	Message      *string `json:"message,omitempty"`
}

// PrivacyPolicyResponse is the privacy policy response.
type PrivacyPolicyResponse struct {
	Content json.RawMessage `json:"content,omitempty"`
	Version *string         `json:"version,omitempty"`
}

// LoginImportantInformationResponse is the login important information response.
type LoginImportantInformationResponse struct {
	Content *string `json:"content,omitempty"`
	Show    bool    `json:"show"`
}

// AdministrativeAuthorityResponse is the administrative authority response.
type AdministrativeAuthorityResponse struct {
	Name    *string `json:"name,omitempty"`
	Contact *string `json:"contact,omitempty"`
}

// GetMaxFileSize gets the maximum allowed file upload size.
func GetMaxFileSize(ctx context.Context, s *aulaapi.Session) (int64, error) {
	return aulaapi.SessionGet[int64](ctx, s, "?method=centralConfiguration.getMaxFileSize")
}

// GetAuthorizedFileFormats gets the list of authorized (allowed) file formats for upload.
func GetAuthorizedFileFormats(ctx context.Context, s *aulaapi.Session) ([]models.AuthorizedFileFormat, error) {
	return aulaapi.SessionGet[[]models.AuthorizedFileFormat](ctx, s, "?method=centralConfiguration.getauthorizedfileformats")
}

// IsAppDeprecated checks whether the current app version is deprecated (force update).
func IsAppDeprecated(ctx context.Context, s *aulaapi.Session) (AppDeprecatedResponse, error) {
	return aulaapi.SessionGet[AppDeprecatedResponse](ctx, s, "?method=centralConfiguration.isAppVersionDeprecated")
}

// GetPrivacyPolicy gets the privacy/data policy content.
func GetPrivacyPolicy(ctx context.Context, s *aulaapi.Session) (PrivacyPolicyResponse, error) {
	return aulaapi.SessionGet[PrivacyPolicyResponse](ctx, s, "?method=centralConfiguration.getDataPolicy")
}

// GetAdministrativeAuthority gets the administrative authority information.
func GetAdministrativeAuthority(ctx context.Context, s *aulaapi.Session) (AdministrativeAuthorityResponse, error) {
	return aulaapi.SessionGet[AdministrativeAuthorityResponse](ctx, s, "?method=municipalConfiguration.getSameAdministrativeAuthorityInstitutions")
}

// GetLoginImportantInformation gets important information to display on the login page.
func GetLoginImportantInformation(ctx context.Context, s *aulaapi.Session) (LoginImportantInformationResponse, error) {
	return aulaapi.SessionGet[LoginImportantInformationResponse](ctx, s, "?method=centralConfiguration.getLoginImportantInformation")
}
