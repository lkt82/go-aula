package aulaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/publicsuffix"
)

// Environment represents an Aula deployment environment.
type Environment string

const (
	EnvProduction Environment = "production"
	EnvPreprod    Environment = "preprod"
	EnvHotfix     Environment = "hotfix"
	EnvTest1      Environment = "test1"
	EnvTest3      Environment = "test3"
	EnvDev1       Environment = "dev1"
	EnvDev3       Environment = "dev3"
	EnvDev11      Environment = "dev11"
)

// BackendHost returns the backend host for API calls.
func (e Environment) BackendHost() string {
	switch e {
	case EnvProduction:
		return "www.aula.dk"
	case EnvPreprod:
		return "www1-preprod.aula.dk"
	case EnvHotfix:
		return "www1-hotfix.aula.dk"
	case EnvTest1:
		return "www1-test1.ncaula.com"
	case EnvTest3:
		return "www1-test3.ncaula.com"
	case EnvDev1:
		return "www1-dev1.ncaula.com"
	case EnvDev3:
		return "www1-dev3.ncaula.com"
	case EnvDev11:
		return "www1-dev11.ncaula.com"
	default:
		return "www.aula.dk"
	}
}

// AuthHost returns the auth (login) host.
func (e Environment) AuthHost() string {
	switch e {
	case EnvProduction:
		return "login.aula.dk"
	case EnvPreprod:
		return "login-preprod.aula.dk"
	case EnvHotfix:
		return "login-hotfix.aula.dk"
	default:
		// Non-prod test/dev share the same host for both backend and auth.
		return e.BackendHost()
	}
}

// RequiresBasicAuth returns whether this environment needs HTTP Basic Auth.
func (e Environment) RequiresBasicAuth() bool {
	return e != EnvProduction
}

const (
	csrfCookieName = "Csrfp-Token"
	csrfHeaderName = "csrfp-token"
	basicAuthUser  = "aula-user"
	basicAuthPass  = "Aula-1337"
)

// AulaClientConfig configures an AulaClient.
type AulaClientConfig struct {
	Environment Environment
	APIVersion  int
}

// DefaultClientConfig returns the default configuration (production, API v23).
func DefaultClientConfig() AulaClientConfig {
	return AulaClientConfig{
		Environment: EnvProduction,
		APIVersion:  23,
	}
}

// AulaClient is the HTTP client for the Aula API.
type AulaClient struct {
	http         *http.Client
	jar          *cookiejar.Jar
	baseURL      *url.URL
	environment  Environment
	useBasicAuth bool

	mu          sync.Mutex
	accessToken *string
}

// NewAulaClient creates a client with default configuration.
func NewAulaClient() (*AulaClient, error) {
	return NewAulaClientWithConfig(DefaultClientConfig())
}

// NewAulaClientWithConfig creates a client with the given configuration.
func NewAulaClientWithConfig(config AulaClientConfig) (*AulaClient, error) {
	baseStr := fmt.Sprintf("https://%s/api/v%d/", config.Environment.BackendHost(), config.APIVersion)
	baseURL, err := url.Parse(baseStr)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("creating cookie jar: %w", err)
	}

	client := &http.Client{
		Jar: jar,
	}

	return &AulaClient{
		http:         client,
		jar:          jar,
		baseURL:      baseURL,
		environment:  config.Environment,
		useBasicAuth: config.Environment.RequiresBasicAuth(),
	}, nil
}

// NewAulaClientWithBaseURL creates a client for testing with a custom base URL.
func NewAulaClientWithBaseURL(baseStr string) (*AulaClient, error) {
	baseURL, err := url.Parse(baseStr)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("creating cookie jar: %w", err)
	}

	client := &http.Client{
		Jar: jar,
	}

	return &AulaClient{
		http:         client,
		jar:          jar,
		baseURL:      baseURL,
		environment:  EnvProduction,
		useBasicAuth: false,
	}, nil
}

// Environment returns the target environment.
func (c *AulaClient) Environment() Environment {
	return c.environment
}

// BaseURL returns the base URL for API requests.
func (c *AulaClient) BaseURL() *url.URL {
	return c.baseURL
}

// HTTPClient returns the underlying http.Client.
func (c *AulaClient) HTTPClient() *http.Client {
	return c.http
}

// SetAccessToken sets the OIDC access token.
func (c *AulaClient) SetAccessToken(token *string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.accessToken = token
}

// SetCookie adds a cookie string to the jar for the base URL (for testing).
func (c *AulaClient) SetCookie(cookieStr string) {
	cookies := []*http.Cookie{}
	parts := strings.Split(cookieStr, ";")
	if len(parts) > 0 {
		nameVal := strings.SplitN(strings.TrimSpace(parts[0]), "=", 2)
		if len(nameVal) == 2 {
			cookies = append(cookies, &http.Cookie{
				Name:  nameVal[0],
				Value: nameVal[1],
			})
		}
	}
	c.jar.SetCookies(c.baseURL, cookies)
}

// DebugCookies returns a debug representation of cookies in the jar.
func (c *AulaClient) DebugCookies() string {
	cookies := c.jar.Cookies(c.baseURL)
	if len(cookies) == 0 {
		return "(no cookies)"
	}
	parts := make([]string, len(cookies))
	for i, ck := range cookies {
		parts[i] = fmt.Sprintf("%s=%s", ck.Name, ck.Value)
	}
	return strings.Join(parts, "; ")
}

// csrfToken reads the current CSRF token from the cookie jar.
func (c *AulaClient) csrfToken() string {
	for _, ck := range c.jar.Cookies(c.baseURL) {
		if ck.Name == csrfCookieName {
			return ck.Value
		}
	}
	return ""
}

// buildURL constructs the full URL for a request path.
// Paths starting with '?' are appended directly (RPC-style).
// Other paths are joined as path segments.
func (c *AulaClient) buildURL(path string) string {
	if strings.HasPrefix(path, "?") {
		return c.baseURL.String() + path
	}
	path = strings.TrimPrefix(path, "/")
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return c.baseURL.String() + path
	}
	return u.String()
}

// newRequest creates an HTTP request with common headers and decorations.
func (c *AulaClient) newRequest(ctx context.Context, method, path string) (*http.Request, error) {
	urlStr := c.buildURL(path)

	// Append access token as query parameter.
	c.mu.Lock()
	token := c.accessToken
	c.mu.Unlock()

	if token != nil && *token != "" {
		if strings.Contains(urlStr, "?") {
			urlStr += "&access_token=" + url.QueryEscape(*token)
		} else {
			urlStr += "?access_token=" + url.QueryEscape(*token)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Android")

	// CSRF token.
	if csrf := c.csrfToken(); csrf != "" {
		req.Header.Set(csrfHeaderName, csrf)
	}

	// Basic auth for non-production.
	if c.useBasicAuth {
		req.SetBasicAuth(basicAuthUser, basicAuthPass)
	}

	return req, nil
}

// doRequest executes an HTTP request and parses the response envelope.
func doRequest[T any](c *AulaClient, req *http.Request) (T, error) {
	var zero T

	if os.Getenv("AULA_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "[DEBUG] %s %s\n", req.Method, req.URL)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return zero, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// Map well-known HTTP status codes before parsing body.
	if resp.StatusCode == http.StatusUnauthorized {
		return zero, ErrUnauthorized
	}
	if resp.StatusCode == http.StatusServiceUnavailable {
		return zero, ErrMaintenance
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, fmt.Errorf("reading response body: %w", err)
	}

	// Handle server errors (5xx) early.
	if resp.StatusCode >= 500 {
		return zero, &APIError{
			Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(body)),
		}
	}

	return parseEnvelope[T](body, resp.StatusCode)
}

// Get sends a GET request and deserializes the envelope payload.
func Get[T any](ctx context.Context, c *AulaClient, path string) (T, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path)
	if err != nil {
		var zero T
		return zero, err
	}
	return doRequest[T](c, req)
}

// Post sends a POST request with a JSON body and deserializes the envelope payload.
func Post[T any](ctx context.Context, c *AulaClient, path string, body any) (T, error) {
	req, err := c.newRequest(ctx, http.MethodPost, path)
	if err != nil {
		var zero T
		return zero, err
	}
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			var zero T
			return zero, fmt.Errorf("marshaling request body: %w", err)
		}
		req.Body = io.NopCloser(strings.NewReader(string(data)))
		req.ContentLength = int64(len(data))
		req.Header.Set("Content-Type", "application/json")
	}
	return doRequest[T](c, req)
}

// PostEmpty sends a POST request without a body.
func PostEmpty[T any](ctx context.Context, c *AulaClient, path string) (T, error) {
	req, err := c.newRequest(ctx, http.MethodPost, path)
	if err != nil {
		var zero T
		return zero, err
	}
	return doRequest[T](c, req)
}

// Put sends a PUT request with a JSON body.
func Put[T any](ctx context.Context, c *AulaClient, path string, body any) (T, error) {
	req, err := c.newRequest(ctx, http.MethodPut, path)
	if err != nil {
		var zero T
		return zero, err
	}
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			var zero T
			return zero, fmt.Errorf("marshaling request body: %w", err)
		}
		req.Body = io.NopCloser(strings.NewReader(string(data)))
		req.ContentLength = int64(len(data))
		req.Header.Set("Content-Type", "application/json")
	}
	return doRequest[T](c, req)
}

// Delete sends a DELETE request.
func Delete[T any](ctx context.Context, c *AulaClient, path string) (T, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path)
	if err != nil {
		var zero T
		return zero, err
	}
	return doRequest[T](c, req)
}

// DeleteWithBody sends a DELETE request with a JSON body.
func DeleteWithBody[T any](ctx context.Context, c *AulaClient, path string, body any) (T, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path)
	if err != nil {
		var zero T
		return zero, err
	}
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			var zero T
			return zero, fmt.Errorf("marshaling request body: %w", err)
		}
		req.Body = io.NopCloser(strings.NewReader(string(data)))
		req.ContentLength = int64(len(data))
		req.Header.Set("Content-Type", "application/json")
	}
	return doRequest[T](c, req)
}

// KeepAlive sends a keep-alive ping to extend the current session.
func (c *AulaClient) KeepAlive(ctx context.Context) error {
	_, err := PostEmpty[json.RawMessage](ctx, c, "?method=session.keepAlive")
	return err
}
