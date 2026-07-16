package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolveDir(t *testing.T) {

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("UserHomeDir: %v", err)
	}
	base := filepath.Join(home, ".local", "share")

	tests := []struct {
		name       string
		mentalDir  string
		xdgHome    string
		wantSuffix string
	}{
		{
			name:       "MENTAL_DIR takes priority",
			mentalDir:  "/custom/mental",
			xdgHome:    "/xdg/data",
			wantSuffix: "/custom/mental",
		},
		{
			name:       "XDG_DATA_HOME used when MENTAL_DIR unset",
			mentalDir:  "",
			xdgHome:    "/xdg/data",
			wantSuffix: "/xdg/data/mental",
		},
		{
			name:       "default XDG path when both unset",
			mentalDir:  "",
			xdgHome:    "",
			wantSuffix: filepath.Join(base, "mental"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(EnvDir, tc.mentalDir)
			t.Setenv("XDG_DATA_HOME", tc.xdgHome)

			got, err := resolveDir()
			if err != nil {
				t.Fatalf("resolveDir: %v", err)
			}
			if !strings.HasSuffix(got, tc.wantSuffix) &&
				got != tc.wantSuffix {
				t.Errorf(
					"got %q, want suffix %q",
					got, tc.wantSuffix,
				)
			}
		})
	}
}
