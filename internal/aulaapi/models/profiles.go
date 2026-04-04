package models

import "encoding/json"

// Domain-specific type aliases.
type InstitutionProfileID = int64
type ProfileID = int64
type InstitutionCode = string

// Address represents a postal address.
type Address struct {
	ID             *int64  `json:"id,omitempty"`
	Street         *string `json:"street,omitempty"`
	PostalCode     json.RawMessage `json:"postalCode,omitempty"`
	PostalDistrict *string `json:"postalDistrict,omitempty"`
}

// ProfilePictureDto represents a profile picture stored in S3/CDN.
type ProfilePictureDto struct {
	ID                    *int64  `json:"id,omitempty"`
	Key                   *string `json:"key,omitempty"`
	Bucket                *string `json:"bucket,omitempty"`
	IsImageScalingPending *bool   `json:"isImageScalingPending,omitempty"`
	URL                   *string `json:"url,omitempty"`
	ScanningStatus        *string `json:"scanningStatus,omitempty"`
}

// AulaFileContent represents a file content reference.
type AulaFileContent struct {
	ID             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	URL            *string `json:"url,omitempty"`
	Bucket         *string `json:"bucket,omitempty"`
	Key            *string `json:"key,omitempty"`
	Created        *string `json:"created,omitempty"`
	ScanningStatus *string `json:"scanningStatus,omitempty"`
}

// BlockedCommunicationInfo represents communication blocking flags per profile type.
type BlockedCommunicationInfo struct {
	Child                    bool `json:"child"`
	Employee                 bool `json:"employee"`
	Guardian                 bool `json:"guardian"`
	IsBlockedAllProfileTypes bool `json:"isBlockedAllProfileTypes"`
}

// InstitutionProfileBase represents base fields shared by all institution profile variants.
type InstitutionProfileBase struct {
	InstitutionProfileID InstitutionProfileID `json:"institutionProfileId"`
	ProfileID            ProfileID            `json:"profileId"`
	UniPersonID          *int64               `json:"uniPersonId,omitempty"`
	MailBoxID            *int64               `json:"mailBoxId,omitempty"`
	FirstName            *string              `json:"firstName,omitempty"`
	LastName             *string              `json:"lastName,omitempty"`
	FullName             *string              `json:"fullName,omitempty"`
	ShortName            *string              `json:"shortName,omitempty"`
	Metadata             *string              `json:"metadata,omitempty"`
	Role                 *string              `json:"role,omitempty"`
	EncryptionKey        *string              `json:"encryptionKey,omitempty"`
	ProfilePicture       *ProfilePictureDto   `json:"profilePicture,omitempty"`
	MainGroup            *string              `json:"mainGroup,omitempty"`
}

// InstitutionProfile represents a full institution profile with contact info and relations.
type InstitutionProfile struct {
	InstitutionProfileID InstitutionProfileID `json:"institutionProfileId"`
	ProfileID            ProfileID            `json:"profileId"`
	UniPersonID          *int64               `json:"uniPersonId,omitempty"`
	MailBoxID            *int64               `json:"mailBoxId,omitempty"`
	FirstName            *string              `json:"firstName,omitempty"`
	LastName             *string              `json:"lastName,omitempty"`
	FullName             *string              `json:"fullName,omitempty"`
	ShortName            *string              `json:"shortName,omitempty"`
	Metadata             *string              `json:"metadata,omitempty"`
	Role                 *string              `json:"role,omitempty"`
	EncryptionKey        *string              `json:"encryptionKey,omitempty"`
	ProfilePicture       *ProfilePictureDto   `json:"profilePicture,omitempty"`
	MainGroup            *string              `json:"mainGroup,omitempty"`
	InstitutionRole      *string              `json:"institutionRole,omitempty"`
	CommunicationBlock   bool                 `json:"communicationBlock"`
	UploadBlock          bool                 `json:"uploadBlock"`
	Email                *string              `json:"email,omitempty"`
	Phone                *string              `json:"phone,omitempty"`
	Address              *Address             `json:"address,omitempty"`
	Birthday             *string              `json:"birthday,omitempty"`
	Relations            []RelationProfile    `json:"relations,omitempty"`
	Alias                bool                 `json:"alias"`
	GroupHomeID          *int64               `json:"groupHomeId,omitempty"`
	Institution          *InstitutionIdentity `json:"institution,omitempty"`
}

// InstitutionProfileChild represents an institution profile for a child.
type InstitutionProfileChild struct {
	InstitutionProfileID InstitutionProfileID `json:"institutionProfileId"`
	ProfileID            ProfileID            `json:"profileId"`
	UniPersonID          *int64               `json:"uniPersonId,omitempty"`
	MailBoxID            *int64               `json:"mailBoxId,omitempty"`
	FirstName            *string              `json:"firstName,omitempty"`
	LastName             *string              `json:"lastName,omitempty"`
	FullName             *string              `json:"fullName,omitempty"`
	ShortName            *string              `json:"shortName,omitempty"`
	Metadata             *string              `json:"metadata,omitempty"`
	Role                 *string              `json:"role,omitempty"`
	EncryptionKey        *string              `json:"encryptionKey,omitempty"`
	ProfilePicture       *ProfilePictureDto   `json:"profilePicture,omitempty"`
	MainGroup            *string              `json:"mainGroup,omitempty"`
	InstitutionRole      *string              `json:"institutionRole,omitempty"`
	CommunicationBlock   bool                 `json:"communicationBlock"`
	UploadBlock          bool                 `json:"uploadBlock"`
	Email                *string              `json:"email,omitempty"`
	Phone                *string              `json:"phone,omitempty"`
	Address              *Address             `json:"address,omitempty"`
	Birthday             *string              `json:"birthday,omitempty"`
	Relations            []RelationProfile    `json:"relations,omitempty"`
	Alias                bool                 `json:"alias"`
	InstitutionCode      *InstitutionCode     `json:"institutionCode,omitempty"`
	InstitutionName      *string              `json:"institutionName,omitempty"`
	MunicipalityCode     *string              `json:"municipalityCode,omitempty"`
	MunicipalityName     *string              `json:"municipalityName,omitempty"`
}

// ChildProfile represents a simplified child profile.
type ChildProfile struct {
	InstCode                    *string               `json:"instCode,omitempty"`
	Name                        *string               `json:"name,omitempty"`
	ShortName                   *string               `json:"shortName,omitempty"`
	InstProfileID               *InstitutionProfileID `json:"instProfileId,omitempty"`
	ProfileID                   *ProfileID            `json:"profileId,omitempty"`
	ProfilePicture              *ProfilePictureDto    `json:"profilePicture,omitempty"`
	UserID                      *string               `json:"userId,omitempty"`
	HasCustodyOrExtendedAccess  bool                  `json:"hasCustodyOrExtendedAccess"`
	Selected                    bool                  `json:"selected"`
}

// RelationProfile represents a related user as seen from an institution profile.
type RelationProfile struct {
	InstitutionProfileID *InstitutionProfileID `json:"institutionProfileId,omitempty"`
	ProfileID            *ProfileID            `json:"profileId,omitempty"`
	FirstName            *string               `json:"firstName,omitempty"`
	FullName             *string               `json:"fullName,omitempty"`
	LastName             *string               `json:"lastName,omitempty"`
	MailBoxID            *int64                `json:"mailBoxId,omitempty"`
	ShortName            *string               `json:"shortName,omitempty"`
	MainGroupName        *string               `json:"mainGroupName,omitempty"`
	Metadata             *string               `json:"metadata,omitempty"`
	ProfilePicture       *AulaFileContent      `json:"profilePicture,omitempty"`
	Institution          *Institution          `json:"institution,omitempty"`
	Role                 *string               `json:"role,omitempty"`
}

// ChildRelationsProfile represents a child's relation profile.
type ChildRelationsProfile struct {
	ProfileID       *ProfileID       `json:"profileId,omitempty"`
	FirstName       *string          `json:"firstName,omitempty"`
	LastName        *string          `json:"lastName,omitempty"`
	InstitutionCode *InstitutionCode `json:"institutionCode,omitempty"`
	Role            *string          `json:"role,omitempty"`
	AulaEmail       *string          `json:"aulaEmail,omitempty"`
}

// EditorPluginDetail represents editor plugin detail.
type EditorPluginDetail struct {
	Name            *string `json:"name,omitempty"`
	MunicipalCode   *string `json:"municipalCode,omitempty"`
	InstitutionType *string `json:"institutionType,omitempty"`
}

// MainGroup represents a main group for a profile.
type MainGroup struct {
	ID              *int64           `json:"id,omitempty"`
	Name            *string          `json:"name,omitempty"`
	InstitutionCode *InstitutionCode `json:"institutionCode,omitempty"`
	IsMainGroup     bool             `json:"MainGroup"`
}

// WidgetConfigurationDto represents a widget configuration on a page.
type WidgetConfigurationDto struct {
	ID                    *int       `json:"id,omitempty"`
	Widget                *WidgetDto `json:"widget,omitempty"`
	Placement             *string    `json:"placement,omitempty"`
	AggregatedDisplayMode *string    `json:"aggregatedDisplayMode,omitempty"`
	Order                 *int       `json:"order,omitempty"`
}

// WidgetDto represents a widget definition.
type WidgetDto struct {
	ID               *int    `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Icon             *string `json:"icon,omitempty"`
	IconEmployee     *string `json:"iconEmployee,omitempty"`
	IconHover        *string `json:"iconHover,omitempty"`
	URL              *string `json:"url,omitempty"`
	WidgetType       *string `json:"type,omitempty"`
	UsableForGroups  bool    `json:"usableForGroups"`
	Ordering         *int    `json:"ordering,omitempty"`
	WidgetID         *string `json:"widgetId,omitempty"`
	WidgetVersion    *string `json:"widgetVersion,omitempty"`
	CanAccessOnMobile bool   `json:"canAccessOnMobile"`
}

// PageConfiguration represents a page configuration with widgets and editor plugins.
type PageConfiguration struct {
	WidgetConfigurations []WidgetConfigurationDto `json:"widgetConfigurations,omitempty"`
	EditorPluginDetails  []EditorPluginDetail     `json:"editorPluginDetails,omitempty"`
}

// ModuleDto represents a module definition.
type ModuleDto struct {
	ID                 *int    `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Icon               *string `json:"icon,omitempty"`
	URL                *string `json:"url,omitempty"`
	ModuleType         *string `json:"type,omitempty"`
	Ordering           *int    `json:"ordering,omitempty"`
	CanBePlacedOnGroup bool    `json:"canBePlacedOnGroup"`
}

// ModuleConfigurationDto represents a module configuration.
type ModuleConfigurationDto struct {
	ID                    *int       `json:"id,omitempty"`
	Module                *ModuleDto `json:"module,omitempty"`
	Order                 *int       `json:"order,omitempty"`
	AggregatedDisplayMode *string    `json:"aggregatedDisplayMode,omitempty"`
}

// RoleDefinition represents a role definition.
type RoleDefinition struct {
	ID       *int    `json:"id,omitempty"`
	RoleName *string `json:"roleName,omitempty"`
}

// Profile represents a top-level profile.
type Profile struct {
	ID                 *ProfileID         `json:"id,omitempty"`
	InstitutionProfile *InstitutionProfile `json:"institutionProfile,omitempty"`
	Groups             []Group            `json:"groups,omitempty"`
	MunicipalGroups    []Group            `json:"municipalGroups,omitempty"`
	Phonenumber        *string            `json:"phonenumber,omitempty"`
	ExternalEmail      *string            `json:"email,omitempty"`
	WorkPhonenumber    *string            `json:"workPhonenumber,omitempty"`
	HomePhonenumber    *string            `json:"homePhonenumber,omitempty"`
	MobilePhonenumber  *string            `json:"mobilePhonenumber,omitempty"`
	Administrator      json.RawMessage    `json:"administrator,omitempty"`
	FirstName          *string            `json:"firstName,omitempty"`
	LastName           *string            `json:"lastName,omitempty"`
	UserID             *string            `json:"userId,omitempty"`
	PortalRole         *string            `json:"portalRole,omitempty"`
	IsSteppedUp        bool               `json:"isSteppedUp"`
	GroupHomes         []json.RawMessage  `json:"groupHomes,omitempty"`
	IsGroupHomeAdmin   bool               `json:"isGroupHomeAdmin"`
	PageConfiguration  *PageConfiguration `json:"pageConfiguration,omitempty"`
}

// Permission represents a permission held by an institution profile.
type Permission struct {
	PermissionID     *string `json:"permissionId,omitempty"`
	StepUp           bool    `json:"stepUp"`
	GroupScopes      []int   `json:"groupScopes,omitempty"`
	InstitutionScope bool    `json:"institutionScope"`
}

// SimpleInstitutionProfile represents a simplified institution profile.
type SimpleInstitutionProfile struct {
	ProfileID            *ProfileID            `json:"profileId,omitempty"`
	InstitutionProfileID *InstitutionProfileID `json:"institutionProfileId,omitempty"`
	InstitutionCode      *InstitutionCode      `json:"institutionCode,omitempty"`
	InstitutionName      *string               `json:"institutionName,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	Role                 *string               `json:"role,omitempty"`
	MainGroup            *string               `json:"mainGroup,omitempty"`
	ProfilePicture       json.RawMessage       `json:"profilePicture,omitempty"`
	ShortName            *string               `json:"shortName,omitempty"`
	Metadata             *string               `json:"metadata,omitempty"`
}

// StubbedUser represents a minimal user reference.
type StubbedUser struct {
	ProfileID *ProfileID `json:"profileId,omitempty"`
	Name      *string    `json:"name,omitempty"`
}

// ChildMetadata represents child metadata used in stubbed user contexts.
type ChildMetadata struct {
	ProfileID      *ProfileID       `json:"profileId,omitempty"`
	Name           *string          `json:"name,omitempty"`
	ID             *int64           `json:"id,omitempty"`
	Metadata       *string          `json:"metadata,omitempty"`
	ProfilePicture *AulaFileContent `json:"profilePicture,omitempty"`
}

// EmployeeMetadata represents employee metadata.
type EmployeeMetadata struct {
	Name                 *string               `json:"name,omitempty"`
	InstitutionProfileID *InstitutionProfileID `json:"institutionProfileId,omitempty"`
	InstitutionRole      *string               `json:"institutionRole,omitempty"`
}

// ComeGoInstitutionProfile represents a ComeGo-specific institution profile view.
type ComeGoInstitutionProfile struct {
	ProfileID            *ProfileID            `json:"profileId,omitempty"`
	InstitutionProfileID *InstitutionProfileID `json:"institutionProfileId,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	Role                 *string               `json:"role,omitempty"`
	ProfilePicture       json.RawMessage       `json:"profilePicture,omitempty"`
	ShortName            *string               `json:"shortName,omitempty"`
	InstitutionCode      *InstitutionCode      `json:"institutionCode,omitempty"`
}

// ContactListInstitutionProfile represents an institution profile augmented with contact list fields.
type ContactListInstitutionProfile struct {
	InstitutionProfile
	ProfilePictureURL                               *string `json:"profilePictureUrl,omitempty"`
	UserHasGivenConsentToShowContactInformation      bool   `json:"userHasGivenConsentToShowContactInformation"`
	CurrentUserCanViewContactInformation             bool   `json:"currentUserCanViewContactInformation"`
}
