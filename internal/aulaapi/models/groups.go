package models

import "encoding/json"

// StubbedGroup represents a minimal group reference.
type StubbedGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GroupMemberGroup represents a stubbed group extended with institution code.
type GroupMemberGroup struct {
	ID              *int64           `json:"id,omitempty"`
	Name            *string          `json:"name,omitempty"`
	InstitutionCode *InstitutionCode `json:"institutionCode,omitempty"`
}

// SimpleGroupDto represents a simple group DTO with institution info.
type SimpleGroupDto struct {
	ID              *int64           `json:"id,omitempty"`
	Name            *string          `json:"name,omitempty"`
	InstitutionCode *InstitutionCode `json:"institutionCode,omitempty"`
	InstitutionName *string          `json:"institutionName,omitempty"`
}

// GroupModule represents a module enabled on a group.
type GroupModule struct {
	ID          *int    `json:"id,omitempty"`
	Module      *string `json:"module,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

// GroupWidget represents a widget enabled on a group.
type GroupWidget struct {
	ID          *int    `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon        *string `json:"icon,omitempty"`
}

// Group represents a full group with memberships, modules, and configuration.
type Group struct {
	ID                                  *int64               `json:"id,omitempty"`
	Name                                *string              `json:"name,omitempty"`
	Description                         *string              `json:"description,omitempty"`
	MembershipType                      json.RawMessage      `json:"membershipType,omitempty"`
	MembershipInstitutions              []string             `json:"membershipInstitutions,omitempty"`
	Access                              *string              `json:"access,omitempty"`
	CurrentUserCanAccessGroupDashBoard  bool                 `json:"currentUserCanAccessGroupDashBoard"`
	Status                              *string              `json:"status,omitempty"`
	Role                                *string              `json:"role,omitempty"`
	DashboardEnabled                    bool                 `json:"dashboardEnabled"`
	InstitutionCode                     *InstitutionCode     `json:"institutionCode,omitempty"`
	GroupType                           *string              `json:"type,omitempty"`
	ValidGroupModules                   []GroupModule        `json:"validGroupModules,omitempty"`
	AllowMembersToBeShown               bool                 `json:"allowMembersToBeShown"`
	ValidGroupWidgets                   []GroupWidget        `json:"validGroupWidgets,omitempty"`
	Memberships                         []GroupMembership    `json:"memberships,omitempty"`
	EndTime                             *string              `json:"endTime,omitempty"`
}

// GroupMembershipInstitutionProfile represents an institution profile as seen in a group membership context.
type GroupMembershipInstitutionProfile struct {
	InstitutionProfileID int64               `json:"institutionProfileId"`
	ProfileID            int64               `json:"profileId"`
	UniPersonID          *int64              `json:"uniPersonId,omitempty"`
	MailBoxID            *int64              `json:"mailBoxId,omitempty"`
	FirstName            *string             `json:"firstName,omitempty"`
	LastName             *string             `json:"lastName,omitempty"`
	FullName             *string             `json:"fullName,omitempty"`
	ShortName            *string             `json:"shortName,omitempty"`
	Metadata             *string             `json:"metadata,omitempty"`
	Role                 *string             `json:"role,omitempty"`
	EncryptionKey        *string             `json:"encryptionKey,omitempty"`
	ProfilePicture       json.RawMessage     `json:"profilePicture,omitempty"`
	MainGroup            *string             `json:"mainGroup,omitempty"`
	MainGroupName        *string             `json:"mainGroupName,omitempty"`
	Relations            []RecipientRelation `json:"relations,omitempty"`
}

// GroupMembership represents a membership of a profile in a group.
type GroupMembership struct {
	ID                 *int64                             `json:"id,omitempty"`
	GroupRole          *string                            `json:"groupRole,omitempty"`
	InactiveDate       *string                            `json:"inactiveDate,omitempty"`
	InstitutionProfile *GroupMembershipInstitutionProfile `json:"institutionProfile,omitempty"`
	GroupID            *int64                             `json:"groupId,omitempty"`
	MemberGroup        *GroupMemberGroup                  `json:"memberGroup,omitempty"`
	InstitutionRole    *string                            `json:"institutionRole,omitempty"`
}

// RecipientRelation represents relation info as seen in group membership contexts.
type RecipientRelation struct {
	InstProfileID *int64  `json:"instProfileId,omitempty"`
	FirstName     *string `json:"firstName,omitempty"`
	LastName      *string `json:"lastName,omitempty"`
	FullName      *string `json:"fullName,omitempty"`
	ShortName     *string `json:"shortName,omitempty"`
	Metadata      *string `json:"metadata,omitempty"`
	Role          *string `json:"role,omitempty"`
	MainGroupName *string `json:"mainGroupName,omitempty"`
}

// GroupByContextDto represents a group context for profile-based views.
type GroupByContextDto struct {
	ID            *int64  `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	ShowAsDefault bool    `json:"showAsDefault"`
}

// GroupMembershipGroupingByProfileType represents a grouping of memberships by portal role.
type GroupMembershipGroupingByProfileType struct {
	Role    *string           `json:"role,omitempty"`
	Members []json.RawMessage `json:"members,omitempty"`
}
