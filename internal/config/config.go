// Package config resolves MENTAL_DIR and manages application
// configuration using Viper.
//
// # MENTAL_DIR Resolution
//
// The base directory for mental's data follows this priority order:
//
//  1. MENTAL_DIR environment variable (if set and non-empty)
//  2. $XDG_DATA_HOME/mental (if XDG_DATA_HOME is set)
//  3. os.UserDataDir()/mental (~/.local/share/mental on Linux)
//
// This satisfies the XDG Base Directory Specification while allowing
// full override via environment variable for custom setups.
//
// # Usage
//
//	cfg, err := config.Load()
//	if err != nil {
//	    return fmt.Errorf("config: %w", err)
//	}
//	dataDir := cfg.Dir()
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	// EnvDir is the environment variable that overrides the data directory.
	// When set, this takes priority over XDG defaults.
	EnvDir = "MENTAL_DIR"

	// DefaultDirName is the subdirectory name within XDG_DATA_HOME.
	DefaultDirName = "mental"
)

// Config holds resolved application configuration.
// Obtain a Config via Load — do not construct directly.
type Config struct {
	dir string
}

// Load resolves configuration from environment and viper.
// It must be called once at startup before any command runs.
func Load() (*Config, error) {
	viper.SetEnvPrefix("MENTAL")
	viper.AutomaticEnv()

	dir, err := resolveDir()
	if err != nil {
		return nil, fmt.Errorf("resolve data dir: %w", err)
	}

	return &Config{dir: dir}, nil
}

// Dir returns the resolved MENTAL_DIR path.
// This is the root under which all mental data is stored.
func (c *Config) Dir() string {
	return c.dir
}

// resolveDir returns the mental data directory.
// Priority: MENTAL_DIR env → XDG_DATA_HOME/mental → ~/.local/share/mental.
func resolveDir() (string, error) {
	if v := os.Getenv(EnvDir); v != "" {
		return v, nil
	}

	if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
		return filepath.Join(xdg, DefaultDirName), nil
	}

	return xdgDataHome()
}

// xdgDataHome returns the XDG data home directory joined with the
// mental subdirectory. Falls back to ~/.local/share/mental per spec.
func xdgDataHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("UserHomeDir: %w", err)
	}

	return filepath.Join(home, ".local", "share", DefaultDirName), nil
}
