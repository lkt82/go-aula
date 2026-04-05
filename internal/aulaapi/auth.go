package aulaapi

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// RedirectURI is the redirect URI used by the Aula mobile app.
	RedirectURI = "https://app-private.aula.dk"

	// authorizePath is the SimpleSAMLphp OIDC authorize endpoint.
	authorizePath = "/simplesaml/module.php/oidc/authorize.php"

	// tokenPath is the SimpleSAMLphp OIDC token endpoint.
	tokenPath = "/simplesaml/module.php/oidc/token.php"

	// codeVerifierBytes is the PKCE verifier length (32 bytes = 256 bits).
	codeVerifierBytes = 32
)

// AuthLevel represents the authentication level.
type AuthLevel string

const (
	AuthLevel2 AuthLevel = "level2" // UniLogin
	AuthLevel3 AuthLevel = "level3" // MitID
)

// ClientID returns the OIDC client ID for this authentication level.
func (l AuthLevel) ClientID() string {
	switch l {
	case AuthLevel3:
		return "_99949a54b8b65423862aac1bf629599ed64231607a"
	default:
		return "_742adb5e2759028d86dbadf4af44ef70e8b1f407a6"
	}
}

// Scope returns the OIDC scope for this authentication level.
func (l AuthLevel) Scope() string {
	switch l {
	case AuthLevel3:
		return "aula-sensitive"
	default:
		return "aula"
	}
}

// String returns a human-readable description of the auth level.
func (l AuthLevel) String() string {
	switch l {
	case AuthLevel3:
		return "Level 3 (MitID)"
	default:
		return "Level 2 (UniLogin)"
	}
}

// PkceChallenge is a PKCE (Proof Key for Code Exchange) challenge pair per RFC 7636.
type PkceChallenge struct {
	CodeVerifier  string
	CodeChallenge string
}

// GeneratePkce generates a fresh PKCE challenge pair using a CSPRNG.
func GeneratePkce() (PkceChallenge, error) {
	bytes := make([]byte, codeVerifierBytes)
	if _, err := rand.Read(bytes); err != nil {
		return PkceChallenge{}, fmt.Errorf("generating PKCE verifier: %w", err)
	}
	return PkceFromVerifierBytes(bytes), nil
}

// PkceFromVerifierBytes builds a challenge from raw verifier bytes.
func PkceFromVerifierBytes(bytes []byte) PkceChallenge {
	verifier := base64.RawURLEncoding.EncodeToString(bytes)
	digest := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(digest[:])
	return PkceChallenge{
		CodeVerifier:  verifier,
		CodeChallenge: challenge,
	}
}

// OidcEndpoints holds the OIDC endpoint URLs for an environment.
type OidcEndpoints struct {
	AuthorizeURL string
	TokenURL     string
	Issuer       string
}

// OidcEndpointsForEnvironment constructs endpoints for the given environment.
func OidcEndpointsForEnvironment(env Environment) OidcEndpoints {
	base := "https://" + env.AuthHost()
	return OidcEndpoints{
		AuthorizeURL: base + authorizePath,
		TokenURL:     base + tokenPath,
		Issuer:       base + "/",
	}
}

// AuthorizeParams holds parameters for building an authorization URL.
type AuthorizeParams struct {
	AuthLevel     AuthLevel
	CodeChallenge string
	State         string
	RedirectURI   string // Override; empty uses default.
}

// BuildAuthorizeURL builds the full authorization URL with all required query parameters.
func BuildAuthorizeURL(endpoints *OidcEndpoints, params *AuthorizeParams) string {
	redirect := params.RedirectURI
	if redirect == "" {
		redirect = RedirectURI
	}

	u, _ := url.Parse(endpoints.AuthorizeURL)
	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", params.AuthLevel.ClientID())
	q.Set("scope", params.AuthLevel.Scope())
	q.Set("redirect_uri", redirect)
	q.Set("code_challenge", params.CodeChallenge)
	q.Set("code_challenge_method", "S256")
	q.Set("state", params.State)
	u.RawQuery = q.Encode()
	return u.String()
}

// GenerateState generates a random state string for OIDC CSRF protection.
func GenerateState() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generating state: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

// TokenResponse is the OIDC token endpoint response.
type TokenResponse struct {
	AccessToken  string  `json:"access_token"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    *uint64 `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	IDToken      *string `json:"id_token,omitempty"`
	Scope        *string `json:"scope,omitempty"`
}

// TokenErrorResponse is an OAuth error from the token endpoint.
type TokenErrorResponse struct {
	Error            string  `json:"error"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

// ExchangeCode exchanges an authorization code for tokens.
func ExchangeCode(
	httpClient *http.Client,
	endpoints *OidcEndpoints,
	authLevel AuthLevel,
	code, codeVerifier string,
	redirectURI string,
) (*TokenResponse, error) {
	if redirectURI == "" {
		redirectURI = RedirectURI
	}

	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {redirectURI},
		"client_id":     {authLevel.ClientID()},
		"code_verifier": {codeVerifier},
	}

	resp, err := httpClient.PostForm(endpoints.TokenURL, form)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var tokenErr TokenErrorResponse
		if json.Unmarshal(body, &tokenErr) == nil && tokenErr.Error != "" {
			desc := ""
			if tokenErr.ErrorDescription != nil {
				desc = *tokenErr.ErrorDescription
			}
			return nil, &AuthError{
				ErrorCode:   tokenErr.Error,
				Description: desc,
			}
		}
		return nil, &AuthError{
			ErrorCode:   fmt.Sprintf("token_exchange_failed (HTTP %d)", resp.StatusCode),
			Description: string(body),
		}
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("JSON error: %w", err)
	}
	return &tokenResp, nil
}

// RefreshToken refreshes an access token using a refresh token.
func RefreshToken(
	httpClient *http.Client,
	endpoints *OidcEndpoints,
	authLevel AuthLevel,
	refreshTok string,
) (*TokenResponse, error) {
	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshTok},
		"client_id":     {authLevel.ClientID()},
	}

	resp, err := httpClient.PostForm(endpoints.TokenURL, form)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var tokenErr TokenErrorResponse
		if json.Unmarshal(body, &tokenErr) == nil && tokenErr.Error != "" {
			desc := ""
			if tokenErr.ErrorDescription != nil {
				desc = *tokenErr.ErrorDescription
			}
			return nil, &AuthError{
				ErrorCode:   tokenErr.Error,
				Description: desc,
			}
		}
		return nil, &AuthError{
			ErrorCode:   fmt.Sprintf("token_refresh_failed (HTTP %d)", resp.StatusCode),
			Description: string(body),
		}
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("JSON error: %w", err)
	}
	return &tokenResp, nil
}

// LoginData holds persisted login data, mirroring AulaNative.OAuth.LoginData.
type LoginData struct {
	AccessToken           string    `json:"access_token"`
	RefreshToken          *string   `json:"refresh_token,omitempty"`
	ExpiresIn             *uint64   `json:"expires_in,omitempty"`
	AccessTokenExpiration *uint64   `json:"access_token_expiration,omitempty"`
	AuthLevel             AuthLevel `json:"auth_level"`
	Error                 *string   `json:"error,omitempty"`
	ErrorDescription      *string   `json:"error_description,omitempty"`
}

// LoginDataFromTokenResponse creates LoginData from a successful token response.
func LoginDataFromTokenResponse(resp *TokenResponse, authLevel AuthLevel) *LoginData {
	now := uint64(time.Now().Unix())
	var expiration *uint64
	if resp.ExpiresIn != nil {
		exp := now + *resp.ExpiresIn
		expiration = &exp
	}
	return &LoginData{
		AccessToken:           resp.AccessToken,
		RefreshToken:          resp.RefreshToken,
		ExpiresIn:             resp.ExpiresIn,
		AccessTokenExpiration: expiration,
		AuthLevel:             authLevel,
	}
}

// IsExpired checks whether the access token has expired.
func (ld *LoginData) IsExpired() bool {
	return ld.IsExpiredWithBuffer(0)
}

// IsExpiredWithBuffer checks whether the token has expired or will expire
// within bufferSecs seconds.
func (ld *LoginData) IsExpiredWithBuffer(bufferSecs uint64) bool {
	if ld.AccessTokenExpiration == nil {
		return false
	}
	now := uint64(time.Now().Unix())
	return now+bufferSecs >= *ld.AccessTokenExpiration
}

// NewLoginDataError creates an error LoginData (no tokens, just error info).
func NewLoginDataError(errMsg string, description *string, authLevel AuthLevel) *LoginData {
	return &LoginData{
		AuthLevel:        authLevel,
		Error:            &errMsg,
		ErrorDescription: description,
	}
}

// ExtractCodeFromRedirect extracts the authorization code from a redirect URL.
func ExtractCodeFromRedirect(redirectURL string, expectedState *string) (string, error) {
	u, err := url.Parse(redirectURL)
	if err != nil {
		return "", fmt.Errorf("parsing redirect URL: %w", err)
	}

	q := u.Query()

	// Check for OAuth error.
	if errCode := q.Get("error"); errCode != "" {
		return "", &AuthError{
			ErrorCode:   errCode,
			Description: q.Get("error_description"),
		}
	}

	// Verify state if expected.
	if expectedState != nil {
		state := q.Get("state")
		if state == "" {
			return "", &AuthError{
				ErrorCode:   "missing_state",
				Description: "redirect URL is missing the state parameter",
			}
		}
		if state != *expectedState {
			return "", &AuthError{
				ErrorCode:   "state_mismatch",
				Description: "OIDC state parameter does not match",
			}
		}
	}

	code := q.Get("code")
	if code == "" {
		return "", &AuthError{
			ErrorCode:   "missing_code",
			Description: "redirect URL is missing the authorization code",
		}
	}
	return code, nil
}

// ParseEnvironment converts a string to an Environment value.
func ParseEnvironment(s string) (Environment, error) {
	switch strings.ToLower(s) {
	case "production", "prod", "":
		return EnvProduction, nil
	case "preprod":
		return EnvPreprod, nil
	case "hotfix":
		return EnvHotfix, nil
	case "test1":
		return EnvTest1, nil
	case "test3":
		return EnvTest3, nil
	case "dev1":
		return EnvDev1, nil
	case "dev3":
		return EnvDev3, nil
	case "dev11":
		return EnvDev11, nil
	default:
		return "", fmt.Errorf("unknown environment: %s", s)
	}
}
