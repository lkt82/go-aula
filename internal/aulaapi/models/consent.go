package models

// ConsentResponsesDto represents an individual consent response.
type ConsentResponsesDto struct {
	ID                      *int64   `json:"id,omitempty"`
	ConsentID               *int     `json:"consentId,omitempty"`
	AllowedAnswers          []string `json:"allowedAnswers,omitempty"`
	ConsentDescription      *string  `json:"consentDescription,omitempty"`
	ConsentResponseAnswer   *string  `json:"consentResponseAnswer,omitempty"`
	ConsentResponseStatus   *string  `json:"consentResponseStatus,omitempty"`
	Editable                bool     `json:"editable"`
	ViewOnlyDependency      *int     `json:"viewOnlyDependency,omitempty"`
	ViewOrder               *int     `json:"viewOrder,omitempty"`
	FromDate                *string  `json:"fromDate,omitempty"`
	ToDate                  *string  `json:"toDate,omitempty"`
}

// InstitutionProfileConsentDto represents a consent profile with consent responses.
type InstitutionProfileConsentDto struct {
	InstitutionProfile *InstitutionProfileConsent `json:"institutionProfile,omitempty"`
	ConsentResponses   []ConsentResponsesDto      `json:"consentResponses,omitempty"`
}

// InstitutionProfileConsent represents a minimal institution profile reference used in consent context.
type InstitutionProfileConsent struct {
	InstitutionProfileID *int64  `json:"institutionProfileId,omitempty"`
	FirstName            *string `json:"firstName,omitempty"`
	LastName             *string `json:"lastName,omitempty"`
	FullName             *string `json:"fullName,omitempty"`
	InstitutionCode      *string `json:"institutionCode,omitempty"`
	InstitutionName      *string `json:"institutionName,omitempty"`
}

// ConsentUpdateDto represents a single consent answer update.
type ConsentUpdateDto struct {
	ConsentID *int64  `json:"consentId,omitempty"`
	Answer    *string `json:"answer,omitempty"`
}

// ProfileConsentUpdatesDto represents a batch consent update.
type ProfileConsentUpdatesDto struct {
	InstitutionProfileID               *int64             `json:"institutionProfileId,omitempty"`
	InstitutionProfileConsentUpdates   []ConsentUpdateDto `json:"institutionProfileConsentUpdates,omitempty"`
}
