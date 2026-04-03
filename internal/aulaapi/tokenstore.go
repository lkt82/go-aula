package aulaapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	tokenFileName = "tokens.json"
	appDirName    = "aula"
)

// TokenStore provides file-based token persistence.
type TokenStore struct {
	dir string
}

// NewTokenStore creates a store at the given directory path.
func NewTokenStore(dir string) *TokenStore {
	return &TokenStore{dir: dir}
}

// DefaultTokenStore creates a store using the platform-appropriate data directory.
// Uses $XDG_DATA_HOME/aula/ on Linux, ~/Library/Application Support/aula/ on macOS,
// or %APPDATA%/aula/ on Windows.
func DefaultTokenStore() (*TokenStore, error) {
	dir, err := userDataDir()
	if err != nil {
		return nil, fmt.Errorf("determining data directory: %w", err)
	}
	return NewTokenStore(filepath.Join(dir, appDirName)), nil
}

// userDataDir returns the platform-appropriate data directory.
func userDataDir() (string, error) {
	// Try XDG_DATA_HOME first (Linux).
	if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
		return dir, nil
	}

	// Fallback to platform defaults.
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// On Windows, use %APPDATA% if set.
	if appdata := os.Getenv("APPDATA"); appdata != "" {
		return appdata, nil
	}

	// Linux/macOS fallback.
	return filepath.Join(home, ".local", "share"), nil
}

// Dir returns the directory where tokens are stored.
func (s *TokenStore) Dir() string {
	return s.dir
}

func (s *TokenStore) tokenPath() string {
	return filepath.Join(s.dir, tokenFileName)
}

// Load reads persisted login data from disk.
// Returns nil, nil if no token file exists.
func (s *TokenStore) Load() (*LoginData, error) {
	data, err := os.ReadFile(s.tokenPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading token file: %w", err)
	}

	var ld LoginData
	if err := json.Unmarshal(data, &ld); err != nil {
		return nil, fmt.Errorf("parsing token file: %w", err)
	}
	return &ld, nil
}

// Save writes login data to disk.
// Creates the storage directory if needed. On Unix, the file is written with mode 0600.
func (s *TokenStore) Save(data *LoginData) error {
	if err := os.MkdirAll(s.dir, 0o700); err != nil {
		return fmt.Errorf("creating token directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling token data: %w", err)
	}

	// Write atomically: temp file then rename.
	tmpPath := filepath.Join(s.dir, ".tokens.json.tmp")
	if err := os.WriteFile(tmpPath, jsonData, 0o600); err != nil {
		return fmt.Errorf("writing token file: %w", err)
	}

	if err := os.Rename(tmpPath, s.tokenPath()); err != nil {
		return fmt.Errorf("renaming token file: %w", err)
	}
	return nil
}

// Clear deletes the token file. Returns nil even if the file does not exist.
func (s *TokenStore) Clear() error {
	err := os.Remove(s.tokenPath())
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing token file: %w", err)
	}
	return nil
}

// Exists checks whether a token file exists on disk.
func (s *TokenStore) Exists() bool {
	_, err := os.Stat(s.tokenPath())
	return err == nil
}
