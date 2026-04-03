package models

// AdministrativeAuthority represents an administrative authority governing institutions.
type AdministrativeAuthority struct {
	ID               *int64   `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	InstitutionCodes []string `json:"institutionCodes,omitempty"`
}

// InstitutionIdentity represents a lightweight institution reference.
type InstitutionIdentity struct {
	InstitutionCode         *InstitutionCode         `json:"institutionCode,omitempty"`
	InstitutionName         *string                  `json:"institutionName,omitempty"`
	MunicipalityCode        *string                  `json:"municipalityCode,omitempty"`
	MunicipalityName        *string                  `json:"municipalityName,omitempty"`
	AdministrativeAuthority *AdministrativeAuthority `json:"administrativeAuthority,omitempty"`
}

// SimpleInstitution represents a minimal institution reference.
type SimpleInstitution struct {
	InstitutionName *string          `json:"institutionName,omitempty"`
	InstitutionCode *InstitutionCode `json:"institutionCode,omitempty"`
}

// Institution represents a full institution with children, permissions, groups, and metadata.
type Institution struct {
	Children                []ChildProfile           `json:"children,omitempty"`
	InstitutionProfileID    *InstitutionProfileID    `json:"institutionProfileId,omitempty"`
	Name                    *string                  `json:"name,omitempty"`
	InstitutionCode         *InstitutionCode         `json:"institutionCode,omitempty"`
	InstitutionType         *string                  `json:"institutionType,omitempty"`
	MunicipalityCode        *string                  `json:"municipalityCode,omitempty"`
	InstitutionRole         *string                  `json:"institutionRole,omitempty"`
	Permissions             []Permission             `json:"permissions,omitempty"`
	Groups                  []Group                  `json:"groups,omitempty"`
	AdministrativeAuthority *AdministrativeAuthority `json:"administrativeAuthority,omitempty"`
	CommunicationBlock      bool                     `json:"communicationBlock"`
	Selected                bool                     `json:"selected"`
	MailboxID               *int                     `json:"mailboxId,omitempty"`
	ShortName               *string                  `json:"shortName,omitempty"`
}
