package models

import "encoding/json"

// OnboardingResponseDto represents the top-level onboarding response.
type OnboardingResponseDto struct {
	Profiles []OnboardingProfileDto `json:"profiles"`
}

// OnboardingProfileDto represents per-profile onboarding information.
type OnboardingProfileDto struct {
	ProfileID                    *ProfileID               `json:"profileId,omitempty"`
	DisplayName                  *string                  `json:"displayName,omitempty"`
	PortalRole                   *string                  `json:"portalRole,omitempty"`
	IsLatestDataPolicyAccepted   bool                     `json:"isLatestDataPolicyAccepted"`
	SupportRole                  bool                     `json:"supportRole"`
	OverConsentAge               *bool                    `json:"overConsentAge,omitempty"`
	ContactInfoEditable          *bool                    `json:"contactInfoEditable,omitempty"`
	Age18AndOlder                *bool                    `json:"age18AndOlder,omitempty"`
	InstitutionProfiles          []LoginInstitutionProfile `json:"institutionProfiles,omitempty"`
	Children                     []LoginChild             `json:"children,omitempty"`
}

// LoginInstitutionProfile represents a rich institution profile as returned by getprofilesbylogin.
type LoginInstitutionProfile struct {
	ID                      InstitutionProfileID `json:"id"`
	ProfileID               ProfileID            `json:"profileId"`
	InstitutionCode         *string              `json:"institutionCode,omitempty"`
	InstitutionName         *string              `json:"institutionName,omitempty"`
	MunicipalityCode        *string              `json:"municipalityCode,omitempty"`
	MunicipalityName        *string              `json:"municipalityName,omitempty"`
	FirstName               *string              `json:"firstName,omitempty"`
	LastName                *string              `json:"lastName,omitempty"`
	FullName                *string              `json:"fullName,omitempty"`
	ShortName               *string              `json:"shortName,omitempty"`
	Gender                  *string              `json:"gender,omitempty"`
	Role                    *string              `json:"role,omitempty"`
	Email                   *string              `json:"email,omitempty"`
	MobilePhoneNumber       *string              `json:"mobilePhoneNumber,omitempty"`
	Metadata                *string              `json:"metadata,omitempty"`
	NewInstitutionProfile   bool                 `json:"newInstitutionProfile"`
	IsPrimary               bool                 `json:"isPrimary"`
	Address                 *Address             `json:"address,omitempty"`
	ProfilePicture          json.RawMessage      `json:"profilePicture,omitempty"`
	Groups                  json.RawMessage      `json:"groups,omitempty"`
	Permissions             json.RawMessage      `json:"permissions,omitempty"`
	MainGroup               json.RawMessage      `json:"mainGroup,omitempty"`
	PageConfiguration       json.RawMessage      `json:"pageConfiguration,omitempty"`
	ModuleConfigurations    json.RawMessage      `json:"moduleConfigurations,omitempty"`
	BlockedCommunication    json.RawMessage      `json:"blockedCommunication,omitempty"`
}

// LoginChild represents a child entry as returned by getprofilesbylogin.
type LoginChild struct {
	ID                          *int64                        `json:"id,omitempty"`
	ProfileID                   *ProfileID                    `json:"profileId,omitempty"`
	UserID                      *string                       `json:"userId,omitempty"`
	Name                        *string                       `json:"name,omitempty"`
	ShortName                   *string                       `json:"shortName,omitempty"`
	InstitutionCode             *string                       `json:"institutionCode,omitempty"`
	HasCustodyOrExtendedAccess  bool                          `json:"hasCustodyOrExtendedAccess"`
	ProfilePicture              json.RawMessage               `json:"profilePicture,omitempty"`
	InstitutionProfile          *LoginChildInstitutionProfile `json:"institutionProfile,omitempty"`
}

// LoginChildInstitutionProfile represents an institution profile for a child within the login response.
type LoginChildInstitutionProfile struct {
	ID              *InstitutionProfileID `json:"id,omitempty"`
	ProfileID       *ProfileID            `json:"profileId,omitempty"`
	InstitutionCode *string               `json:"institutionCode,omitempty"`
	InstitutionName *string               `json:"institutionName,omitempty"`
	FirstName       *string               `json:"firstName,omitempty"`
	LastName        *string               `json:"lastName,omitempty"`
	FullName        *string               `json:"fullName,omitempty"`
	ShortName       *string               `json:"shortName,omitempty"`
	Role            *string               `json:"role,omitempty"`
	Gender          *string               `json:"gender,omitempty"`
	Metadata        *string               `json:"metadata,omitempty"`
	ProfilePicture  json.RawMessage       `json:"profilePicture,omitempty"`
	MainGroup       json.RawMessage       `json:"mainGroup,omitempty"`
}
