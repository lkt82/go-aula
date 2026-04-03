package models

import "encoding/json"

// ProfileContext represents a user profile context containing institution memberships.
type ProfileContext struct {
	ProfileID    *ProfileID    `json:"profileId,omitempty"`
	PortalRole   *string       `json:"portalRole,omitempty"`
	Institutions []Institution `json:"institutions,omitempty"`
}

// User represents a user search result model.
type User struct {
	Address     *string           `json:"address,omitempty"`
	DisplayName *string           `json:"displayName,omitempty"`
	Email       *string           `json:"email,omitempty"`
	FirstName   *string           `json:"firstName,omitempty"`
	LastName    *string           `json:"lastName,omitempty"`
	Highlights  []json.RawMessage `json:"highlights,omitempty"`
}

// UserRelationship represents a relationship between a user and children/institutions.
type UserRelationship struct {
	ProfileID                 *ProfileID `json:"profileId,omitempty"`
	ChildRelationships        []string   `json:"childRelationships,omitempty"`
	InstitutionRelationships  []string   `json:"institutionRelationships,omitempty"`
}
