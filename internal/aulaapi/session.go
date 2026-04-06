package aulaapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

const (
	defaultExpiryBufferSecs uint64 = 60
	logoutPath                     = "/auth/logout.php"
)

// SessionConfig configures a Session.
type SessionConfig struct {
	ExpiryBufferSecs uint64
}

// DefaultSessionConfig returns the default session configuration.
func DefaultSessionConfig() SessionConfig {
	return SessionConfig{
		ExpiryBufferSecs: defaultExpiryBufferSecs,
	}
}

// InstitutionProfileID is the primary ID used in most API calls.
type InstitutionProfileID = int64

// ProfileID is the user-level profile ID (spans institutions).
type ProfileID = int64

// InstitutionCode is the institution identifier string.
type InstitutionCode = string

// OnboardingResponseDto holds the profile data from getProfilesByLogin.
// This is a simplified version; the full model is in models/onboarding.go.
type OnboardingResponseDto struct {
	Profiles []OnboardingProfileDto `json:"profiles"`
}

// OnboardingProfileDto represents a single profile from the onboarding response.
type OnboardingProfileDto struct {
	ProfileID   int64                    `json:"profileId"`
	DisplayName string                   `json:"displayName"`
	PortalRole  string                   `json:"portalRole"`
	InstProfiles []LoginInstitutionProfile `json:"institutionProfiles"`
	Children    []LoginChild              `json:"children,omitempty"`
}

// InstitutionProfileIDs returns the institution profile IDs for this profile.
func (p *OnboardingProfileDto) InstitutionProfileIDs() []int64 {
	ids := make([]int64, 0, len(p.InstProfiles))
	for _, ip := range p.InstProfiles {
		if ip.ID != nil {
			ids = append(ids, *ip.ID)
		}
	}
	return ids
}

// MainGroup is a minimal group reference (id + name).
type MainGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// LoginInstitutionProfile is a minimal institution profile from onboarding.
type LoginInstitutionProfile struct {
	ID              *int64     `json:"id,omitempty"`
	ProfileID       *int64     `json:"profileId,omitempty"`
	InstitutionCode *string    `json:"institutionCode,omitempty"`
	InstitutionName *string    `json:"institutionName,omitempty"`
	FirstName       *string    `json:"firstName,omitempty"`
	LastName        *string    `json:"lastName,omitempty"`
	FullName        *string    `json:"fullName,omitempty"`
	ShortName       *string    `json:"shortName,omitempty"`
	Role            *string    `json:"role,omitempty"`
	MainGroup       *MainGroup `json:"mainGroup,omitempty"`
}

// LoginChild represents a child from the onboarding response.
type LoginChild struct {
	ID                          *int64                   `json:"id,omitempty"`
	ProfileID                   *int64                   `json:"profileId,omitempty"`
	Name                        *string                  `json:"name,omitempty"`
	ShortName                   *string                  `json:"shortName,omitempty"`
	InstitutionCode             *string                  `json:"institutionCode,omitempty"`
	HasCustodyOrExtendedAccess  *bool                    `json:"hasCustodyOrExtendedAccess,omitempty"`
	InstitutionProfile          *LoginInstitutionProfile `json:"institutionProfile,omitempty"`
}

// Session provides managed access to the Aula API with automatic token refresh.
type Session struct {
	client    *AulaClient
	store     *TokenStore
	config    SessionConfig
	endpoints OidcEndpoints

	mu                 sync.Mutex
	loginData          *LoginData
	contextInitialized bool
	profileData        *OnboardingResponseDto
}

// NewSession creates a new session. It loads persisted tokens from the store.
func NewSession(client *AulaClient, store *TokenStore, config SessionConfig) (*Session, error) {
	endpoints := OidcEndpointsForEnvironment(client.Environment())
	loginData, err := store.Load()
	if err != nil {
		return nil, fmt.Errorf("loading tokens: %w", err)
	}

	// Sync persisted token to the HTTP client.
	if loginData != nil && loginData.AccessToken != "" {
		client.SetAccessToken(&loginData.AccessToken)
	}

	return &Session{
		client:    client,
		store:     store,
		config:    config,
		endpoints: endpoints,
		loginData: loginData,
	}, nil
}

// Client returns the underlying AulaClient.
func (s *Session) Client() *AulaClient {
	return s.client
}

// LoginData returns the current login data.
func (s *Session) LoginData() *LoginData {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.loginData
}

// HasValidTokens returns whether the session has valid (non-expired) tokens.
func (s *Session) HasValidTokens() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.loginData == nil {
		return false
	}
	return !s.loginData.IsExpiredWithBuffer(s.config.ExpiryBufferSecs)
}

// SetLoginData sets login data (after OIDC login) and persists it.
func (s *Session) SetLoginData(data *LoginData) error {
	if err := s.store.Save(data); err != nil {
		return err
	}
	if data.AccessToken != "" {
		s.client.SetAccessToken(&data.AccessToken)
	}
	s.mu.Lock()
	s.loginData = data
	s.mu.Unlock()
	return nil
}

// ProfileData returns the profile data captured during context initialization.
func (s *Session) ProfileData() *OnboardingResponseDto {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.profileData
}

// InstitutionProfileIDs returns the guardian's institution profile IDs.
func (s *Session) InstitutionProfileIDs() []int64 {
	s.mu.Lock()
	pd := s.profileData
	s.mu.Unlock()
	if pd == nil {
		return nil
	}
	var ids []int64
	for _, p := range pd.Profiles {
		ids = append(ids, p.InstitutionProfileIDs()...)
	}
	return ids
}

// AllInstitutionProfileIDs returns guardian + children institution profile IDs (deduped).
func (s *Session) AllInstitutionProfileIDs() []int64 {
	ids := s.InstitutionProfileIDs()
	childIDs := s.ChildrenInstProfileIDs()
	seen := make(map[int64]bool, len(ids))
	for _, id := range ids {
		seen[id] = true
	}
	for _, id := range childIDs {
		if !seen[id] {
			ids = append(ids, id)
			seen[id] = true
		}
	}
	return ids
}

// ChildrenInstProfileIDs returns the children's institution profile IDs.
func (s *Session) ChildrenInstProfileIDs() []int64 {
	s.mu.Lock()
	pd := s.profileData
	s.mu.Unlock()
	if pd == nil {
		return nil
	}
	var ids []int64
	for _, p := range pd.Profiles {
		for _, c := range p.Children {
			if c.InstitutionProfile != nil && c.InstitutionProfile.ID != nil {
				ids = append(ids, *c.InstitutionProfile.ID)
			}
		}
	}
	return ids
}

// ChildrenInstitutionCodes returns deduplicated institution codes from children.
func (s *Session) ChildrenInstitutionCodes() []string {
	s.mu.Lock()
	pd := s.profileData
	s.mu.Unlock()
	if pd == nil {
		return nil
	}
	seen := make(map[string]bool)
	var codes []string
	for _, p := range pd.Profiles {
		for _, c := range p.Children {
			if c.InstitutionCode != nil && !seen[*c.InstitutionCode] {
				codes = append(codes, *c.InstitutionCode)
				seen[*c.InstitutionCode] = true
			}
		}
	}
	return codes
}

// EnsureValidToken checks token expiry and refreshes if needed.
func (s *Session) EnsureValidToken(ctx context.Context) error {
	s.mu.Lock()
	ld := s.loginData
	s.mu.Unlock()

	if ld == nil {
		return &AuthError{
			ErrorCode:   "no_tokens",
			Description: "no persisted tokens; initial login required",
		}
	}

	if !ld.IsExpiredWithBuffer(s.config.ExpiryBufferSecs) {
		return nil
	}

	return s.RefreshTokens(ctx)
}

// RefreshTokens refreshes the access token using the stored refresh token.
func (s *Session) RefreshTokens(ctx context.Context) error {
	s.mu.Lock()
	ld := s.loginData
	s.mu.Unlock()

	if ld == nil {
		return &AuthError{
			ErrorCode:   "no_tokens",
			Description: "cannot refresh: no persisted tokens",
		}
	}

	if ld.RefreshToken == nil {
		return &AuthError{
			ErrorCode:   "no_refresh_token",
			Description: "stored tokens have no refresh token",
		}
	}

	tokenResp, err := RefreshToken(s.client.HTTPClient(), &s.endpoints, ld.AuthLevel, *ld.RefreshToken)
	if err != nil {
		return err
	}

	newData := LoginDataFromTokenResponse(tokenResp, ld.AuthLevel)
	if err := s.store.Save(newData); err != nil {
		return err
	}
	s.client.SetAccessToken(&newData.AccessToken)

	s.mu.Lock()
	s.loginData = newData
	s.mu.Unlock()

	return nil
}

// EnsureContextInitialized initializes the server-side profile context.
// Must be called before most API endpoints will work.
func (s *Session) EnsureContextInitialized(ctx context.Context) error {
	s.mu.Lock()
	initialized := s.contextInitialized
	s.mu.Unlock()

	if initialized {
		return nil
	}

	// Step 1: Call getProfilesByLogin to establish the server-side session.
	// Refresh token first if expired, then retry on 401.
	if err := s.EnsureValidToken(ctx); err != nil {
		return fmt.Errorf("token refresh before getProfilesByLogin: %w", err)
	}
	profileResp, err := Get[OnboardingResponseDto](ctx, s.client, "?method=profiles.getprofilesbylogin")
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return fmt.Errorf("getProfilesByLogin: %w", err)
		}
		profileResp, err = Get[OnboardingResponseDto](ctx, s.client, "?method=profiles.getprofilesbylogin")
	}
	if err != nil {
		return fmt.Errorf("getProfilesByLogin: %w", err)
	}

	s.mu.Lock()
	s.profileData = &profileResp
	s.mu.Unlock()

	// Step 2: Call getProfileContext to activate the profile.
	_, err = Get[json.RawMessage](ctx, s.client, "?method=profiles.getProfileContext&portalrole=guardian&deviceId=aula-cli")
	if err != nil {
		return fmt.Errorf("getProfileContext: %w", err)
	}

	s.mu.Lock()
	s.contextInitialized = true
	s.mu.Unlock()

	return nil
}

// preRequest refreshes tokens and initializes context if needed.
func (s *Session) preRequest(ctx context.Context) {
	s.mu.Lock()
	hasTokens := s.loginData != nil
	s.mu.Unlock()

	if hasTokens {
		if err := s.EnsureValidToken(ctx); err != nil {
			fmt.Printf("pre_request: proactive token refresh failed: %v\n", err)
		}
		if err := s.EnsureContextInitialized(ctx); err != nil {
			fmt.Printf("pre_request: profile context initialization failed: %v\n", err)
		}
	}
}

// Get executes a GET request with automatic token refresh and 401 retry.
func (s *Session) Get(ctx context.Context, path string) (json.RawMessage, error) {
	return SessionGet[json.RawMessage](ctx, s, path)
}

// SessionGet executes a typed GET request with automatic token refresh and 401 retry.
func SessionGet[T any](ctx context.Context, s *Session, path string) (T, error) {
	s.preRequest(ctx)
	result, err := Get[T](ctx, s.client, path)
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return result, err
		}
		return Get[T](ctx, s.client, path)
	}
	return result, err
}

// SessionPost executes a typed POST request with automatic token refresh and 401 retry.
func SessionPost[T any](ctx context.Context, s *Session, path string, body any) (T, error) {
	s.preRequest(ctx)
	result, err := Post[T](ctx, s.client, path, body)
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return result, err
		}
		return Post[T](ctx, s.client, path, body)
	}
	return result, err
}

// SessionPostEmpty executes a typed POST request (no body) with auto-refresh.
func SessionPostEmpty[T any](ctx context.Context, s *Session, path string) (T, error) {
	s.preRequest(ctx)
	result, err := PostEmpty[T](ctx, s.client, path)
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return result, err
		}
		return PostEmpty[T](ctx, s.client, path)
	}
	return result, err
}

// SessionPut executes a typed PUT request with auto-refresh.
func SessionPut[T any](ctx context.Context, s *Session, path string, body any) (T, error) {
	s.preRequest(ctx)
	result, err := Put[T](ctx, s.client, path, body)
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return result, err
		}
		return Put[T](ctx, s.client, path, body)
	}
	return result, err
}

// SessionDelete executes a typed DELETE request with auto-refresh.
func SessionDelete[T any](ctx context.Context, s *Session, path string) (T, error) {
	s.preRequest(ctx)
	result, err := Delete[T](ctx, s.client, path)
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return result, err
		}
		return Delete[T](ctx, s.client, path)
	}
	return result, err
}

// SessionDeleteWithBody executes a typed DELETE request with body and auto-refresh.
func SessionDeleteWithBody[T any](ctx context.Context, s *Session, path string, body any) (T, error) {
	s.preRequest(ctx)
	result, err := DeleteWithBody[T](ctx, s.client, path, body)
	if err != nil && (errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrInvalidToken)) {
		if refreshErr := s.RefreshTokens(ctx); refreshErr != nil {
			return result, err
		}
		return DeleteWithBody[T](ctx, s.client, path, body)
	}
	return result, err
}

// KeepAlive sends a keep-alive ping.
func (s *Session) KeepAlive(ctx context.Context) error {
	return s.client.KeepAlive(ctx)
}

// Logout clears stored tokens and hits the OIDC logout endpoint.
func (s *Session) Logout(ctx context.Context) error {
	if err := s.store.Clear(); err != nil {
		return err
	}

	s.mu.Lock()
	s.loginData = nil
	s.mu.Unlock()
	s.client.SetAccessToken(nil)

	// Best-effort: hit the logout endpoint.
	logoutURL := fmt.Sprintf("https://%s%s", s.client.Environment().AuthHost(), logoutPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, logoutURL, nil)
	if err == nil {
		s.client.HTTPClient().Do(req) //nolint:errcheck
	}

	return nil
}
