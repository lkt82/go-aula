// Package aulaapi provides a Go client for the Aula school platform API.
package aulaapi

import (
	"errors"
	"fmt"
)

// Sentinel errors matching the 13 error handler types from the Aula APK.
var (
	ErrNoNetwork       = errors.New("no network connectivity")
	ErrRequestAborted  = errors.New("request aborted")
	ErrInvalidToken    = errors.New("invalid or expired access token")
	ErrSessionExpired  = errors.New("session expired")
	ErrStepUpRequired  = errors.New("step-up authentication required")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrMaintenance     = errors.New("Aula is under maintenance")
	ErrNotResponding   = errors.New("Aula is not responding")
	ErrHeavyLoad       = errors.New("Aula is under heavy load")
	ErrUserDeactivated = errors.New("user account deactivated")
)

// APIError represents a generic API error with optional status detail.
type APIError struct {
	Message string
	Status  *WebResponseStatus
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %s", e.Message)
}

// AuthError represents an OIDC authentication flow error.
type AuthError struct {
	ErrorCode   string
	Description string
}

func (e *AuthError) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("auth error: %s: %s", e.ErrorCode, e.Description)
	}
	return fmt.Sprintf("auth error: %s", e.ErrorCode)
}
