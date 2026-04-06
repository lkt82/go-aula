package aulaapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// newTestSession creates a session backed by a test HTTP server.
// The handler receives all API requests. The token store uses a temp dir.
func newTestSession(t *testing.T, handler http.Handler) (*Session, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client, err := NewAulaClientWithBaseURL(server.URL)
	if err != nil {
		t.Fatalf("NewAulaClientWithBaseURL: %v", err)
	}

	store := NewTokenStore(t.TempDir())
	session, err := NewSession(client, store, SessionConfig{ExpiryBufferSecs: 60})
	if err != nil {
		t.Fatalf("NewSession: %v", err)
	}

	// Point OIDC endpoints to the test server
	session.endpoints = OidcEndpoints{
		AuthorizeURL: server.URL + "/authorize",
		TokenURL:     server.URL + "/token",
		Issuer:       server.URL,
	}

	return session, server
}

// makeExpiredLoginData creates a LoginData with an expired access token.
func makeExpiredLoginData() *LoginData {
	refreshTok := "test-refresh-token"
	expired := uint64(time.Now().Unix() - 3600) // 1 hour ago
	return &LoginData{
		AccessToken:           "expired-token",
		RefreshToken:          &refreshTok,
		AccessTokenExpiration: &expired,
		AuthLevel:             AuthLevel2,
	}
}

// makeValidLoginData creates a LoginData with a valid (far future) token.
func makeValidLoginData() *LoginData {
	refreshTok := "test-refresh-token"
	future := uint64(time.Now().Unix() + 3600) // 1 hour from now
	return &LoginData{
		AccessToken:           "valid-token",
		RefreshToken:          &refreshTok,
		AccessTokenExpiration: &future,
		AuthLevel:             AuthLevel2,
	}
}

// tokenEndpointHandler returns a handler that responds to token refresh requests.
func tokenEndpointHandler(newToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/token" {
			expiresIn := uint64(3600)
			resp := TokenResponse{
				AccessToken: newToken,
				TokenType:   "Bearer",
				ExpiresIn:   &expiresIn,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
		http.NotFound(w, r)
	}
}

// aulaEnvelope wraps data in the standard Aula API response envelope.
func aulaEnvelope(data interface{}) []byte {
	resp := map[string]interface{}{
		"status": map[string]interface{}{
			"httpCode": 200,
			"code":     0,
		},
		"data": data,
	}
	b, _ := json.Marshal(resp)
	return b
}

func TestEnsureValidToken_NotExpired(t *testing.T) {
	session, _ := newTestSession(t, http.NotFoundHandler())
	session.loginData = makeValidLoginData()

	err := session.EnsureValidToken(context.Background())
	if err != nil {
		t.Fatalf("expected no error for valid token, got: %v", err)
	}
}

func TestEnsureValidToken_NoTokens(t *testing.T) {
	session, _ := newTestSession(t, http.NotFoundHandler())
	// loginData is nil by default

	err := session.EnsureValidToken(context.Background())
	if err == nil {
		t.Fatal("expected error for nil loginData")
	}
	authErr, ok := err.(*AuthError)
	if !ok {
		t.Fatalf("expected AuthError, got: %T", err)
	}
	if authErr.ErrorCode != "no_tokens" {
		t.Fatalf("expected no_tokens error, got: %s", authErr.ErrorCode)
	}
}

func TestEnsureValidToken_ExpiredRefreshes(t *testing.T) {
	var refreshCalled atomic.Int32

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/token" {
			refreshCalled.Add(1)
			expiresIn := uint64(3600)
			resp := TokenResponse{
				AccessToken: "new-token",
				TokenType:   "Bearer",
				ExpiresIn:   &expiresIn,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
		http.NotFound(w, r)
	})

	session, _ := newTestSession(t, handler)
	session.loginData = makeExpiredLoginData()

	err := session.EnsureValidToken(context.Background())
	if err != nil {
		t.Fatalf("expected successful refresh, got: %v", err)
	}
	if refreshCalled.Load() != 1 {
		t.Fatalf("expected 1 refresh call, got %d", refreshCalled.Load())
	}
	if session.loginData.AccessToken != "new-token" {
		t.Fatalf("expected new-token, got: %s", session.loginData.AccessToken)
	}
}

func TestRefreshTokens_NoRefreshToken(t *testing.T) {
	session, _ := newTestSession(t, http.NotFoundHandler())
	session.loginData = &LoginData{
		AccessToken: "some-token",
		AuthLevel:   AuthLevel2,
		// RefreshToken is nil
	}

	err := session.RefreshTokens(context.Background())
	if err == nil {
		t.Fatal("expected error for nil RefreshToken")
	}
	authErr, ok := err.(*AuthError)
	if !ok {
		t.Fatalf("expected AuthError, got: %T", err)
	}
	if authErr.ErrorCode != "no_refresh_token" {
		t.Fatalf("expected no_refresh_token, got: %s", authErr.ErrorCode)
	}
}

func TestRefreshTokens_PersistsNewToken(t *testing.T) {
	handler := tokenEndpointHandler("persisted-token")
	session, _ := newTestSession(t, handler)
	session.loginData = makeExpiredLoginData()

	err := session.RefreshTokens(context.Background())
	if err != nil {
		t.Fatalf("RefreshTokens failed: %v", err)
	}

	// Verify token was persisted to store
	loaded, err := session.store.Load()
	if err != nil {
		t.Fatalf("store.Load failed: %v", err)
	}
	if loaded == nil || loaded.AccessToken != "persisted-token" {
		t.Fatalf("expected persisted-token in store, got: %v", loaded)
	}
}

func TestEnsureContextInitialized_RetryOn401(t *testing.T) {
	var callCount atomic.Int32

	var profileCalls atomic.Int32

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)

		// Token refresh endpoint — always succeed
		if r.Method == http.MethodPost && r.URL.Path == "/token" {
			expiresIn := uint64(3600)
			resp := TokenResponse{
				AccessToken: "refreshed-token",
				TokenType:   "Bearer",
				ExpiresIn:   &expiresIn,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}

		// API requests
		if strings.Contains(r.URL.RawQuery, "method=profiles.getprofilesbylogin") {
			n := profileCalls.Add(1)
			// First call: return 401 (invalid token)
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// After refresh: succeed with minimal profile data
			w.Header().Set("Content-Type", "application/json")
			w.Write(aulaEnvelope(map[string]interface{}{
				"profiles": []interface{}{},
			}))
			return
		}

		if strings.Contains(r.URL.RawQuery, "method=profiles.getProfileContext") {
			w.Header().Set("Content-Type", "application/json")
			w.Write(aulaEnvelope(nil))
			return
		}

		http.NotFound(w, r)
	})

	session, _ := newTestSession(t, handler)
	session.loginData = makeValidLoginData()

	err := session.EnsureContextInitialized(context.Background())
	if err != nil {
		t.Fatalf("EnsureContextInitialized should retry and succeed, got: %v", err)
	}
	if !session.contextInitialized {
		t.Fatal("expected contextInitialized to be true")
	}
}

func TestEnsureContextInitialized_OnlyCalledOnce(t *testing.T) {
	var apiCalls atomic.Int32

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiCalls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.RawQuery, "method=profiles.getprofilesbylogin") {
			w.Write(aulaEnvelope(map[string]interface{}{
				"profiles": []interface{}{},
			}))
			return
		}
		if strings.Contains(r.URL.RawQuery, "method=profiles.getProfileContext") {
			w.Write(aulaEnvelope(nil))
			return
		}
		http.NotFound(w, r)
	})

	session, _ := newTestSession(t, handler)
	session.loginData = makeValidLoginData()

	// Call twice
	ctx := context.Background()
	if err := session.EnsureContextInitialized(ctx); err != nil {
		t.Fatalf("first call failed: %v", err)
	}
	calls1 := apiCalls.Load()

	if err := session.EnsureContextInitialized(ctx); err != nil {
		t.Fatalf("second call failed: %v", err)
	}
	calls2 := apiCalls.Load()

	if calls2 != calls1 {
		t.Fatalf("second call should be cached, but got %d extra API calls", calls2-calls1)
	}
}
