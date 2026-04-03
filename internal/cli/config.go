// Package cli provides the command-line interface for aula-cli.
package cli

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config holds persistent CLI configuration loaded from disk.
type Config struct {
	// DefaultEnvironment is the environment when --env is not specified.
	DefaultEnvironment string `toml:"default_environment,omitempty"`
	// DefaultFormat is the output format ("json" or "text").
	DefaultFormat string `toml:"default_format,omitempty"`
	// DefaultProfile is the institution profile name for --profile.
	DefaultProfile string `toml:"default_profile,omitempty"`
	// Verbose enables verbose output by default.
	Verbose bool `toml:"verbose,omitempty"`
}

// ConfigPath returns the path to the configuration file.
func ConfigPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return filepath.Join(dir, "aula", "config.toml")
}

// LoadConfig loads configuration from disk. Returns a default config if the
// file does not exist or cannot be parsed.
func LoadConfig() Config {
	path := ConfigPath()
	if path == "" {
		return Config{}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}
	}
	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return Config{}
	}
	return cfg
}
