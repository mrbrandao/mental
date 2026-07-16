package opencode

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func TestExtractTopics(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		title   string
		wantMin int // minimum expected topics
		notWant []string
	}{
		{
			name:    "typical session title",
			title:   "Renamed ais to mental, added XDG config resolution",
			wantMin: 3,
			notWant: []string{"to", "a", "an"},
		},
		{
			name:    "filters stop words",
			title:   "Fix the bug in the auth service",
			wantMin: 2,
			notWant: []string{"the", "in"},
		},
		{
			name:    "deduplicates",
			title:   "auth migration auth schema auth rollback",
			wantMin: 3,
			notWant: []string{}, // auth should appear once
		},
		{
			name:    "empty title",
			title:   "",
			wantMin: 0,
			notWant: []string{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := extractTopics(tc.title)

			if len(got) < tc.wantMin {
				t.Errorf(
					"topics: got %d, want >= %d: %v",
					len(got), tc.wantMin, got,
				)
			}

			// Check stop words are not present.
			gotSet := make(map[string]bool)
			for _, w := range got {
				gotSet[w] = true
			}
			for _, nw := range tc.notWant {
				if gotSet[nw] {
					t.Errorf("topic %q should be filtered", nw)
				}
			}

			// Check deduplication.
			seen := make(map[string]int)
			for _, w := range got {
				seen[w]++
			}
			for w, count := range seen {
				if count > 1 {
					t.Errorf("topic %q appears %d times", w, count)
				}
			}
		})
	}
}

func TestFetchSession(t *testing.T) {
	t.Parallel()

	t.Run("returns error for missing session", func(t *testing.T) {
		t.Parallel()

		// Create a minimal in-memory SQLite DB with the session table.
		dir := t.TempDir()
		dbPath := filepath.Join(dir, "test.db")

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			t.Fatalf("open: %v", err)
		}
		_, err = db.Exec(
			`CREATE TABLE session (
				id TEXT PRIMARY KEY,
				title TEXT,
				directory TEXT,
				time_updated INTEGER
			)`,
		)
		db.Close()
		if err != nil {
			t.Fatalf("create table: %v", err)
		}

		_, err = fetchSession(dbPath, "nonexistent-id")
		if err == nil {
			t.Error("expected error for missing session")
		}
	})

	t.Run("returns session data", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()
		dbPath := filepath.Join(dir, "test.db")

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			t.Fatalf("open: %v", err)
		}
		_, err = db.Exec(
			`CREATE TABLE session (
				id TEXT PRIMARY KEY,
				title TEXT,
				directory TEXT,
				time_updated INTEGER
			)`,
		)
		if err != nil {
			t.Fatalf("create: %v", err)
		}
		_, err = db.Exec(
			`INSERT INTO session VALUES (?, ?, ?, ?)`,
			"sess-001", "Renamed ais to mental",
			"/home/user/dev/mental", int64(1721084929000),
		)
		db.Close()
		if err != nil {
			t.Fatalf("insert: %v", err)
		}

		row, err := fetchSession(dbPath, "sess-001")
		if err != nil {
			t.Fatalf("fetchSession: %v", err)
		}
		if row.title != "Renamed ais to mental" {
			t.Errorf("title: got %q", row.title)
		}
		if row.directory != "/home/user/dev/mental" {
			t.Errorf("dir: got %q", row.directory)
		}
	})
}

func TestExtractFiles(t *testing.T) {
	t.Parallel()

	t.Run("returns nil when diff file missing", func(t *testing.T) {
		t.Parallel()
		files := extractFiles("/tmp/nonexistent", "sess-001")
		if files != nil {
			t.Errorf("expected nil for missing diff, got %v", files)
		}
	})

	t.Run("parses diff JSON", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()
		diffJSON := `[{"file":"go.mod"},{"file":"cmd/root.go"}]`
		path := filepath.Join(dir, "sess-001.json")
		if err := os.WriteFile(path, []byte(diffJSON), 0o644); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}

		files := extractFiles(dir, "sess-001")
		if len(files) != 2 {
			t.Errorf("files: got %d, want 2", len(files))
		}
		if files[0] != "go.mod" {
			t.Errorf("files[0]: got %q, want go.mod", files[0])
		}
	})
}
