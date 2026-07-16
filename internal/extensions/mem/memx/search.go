package memx

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// SearchResult is a single match from topics.yaml.
type SearchResult struct {
	Checkpoint string
	Name       string
	Summary    string
	Topics     []string
}

// Search parses topics.yaml and returns entries whose topic list
// contains a case-insensitive substring match for query.
func Search(
	cfg *Config,
	mentalDir,
	project,
	query string,
) ([]SearchResult, error) {
	layout := NewLayout(cfg, mentalDir)

	data, err := os.ReadFile(layout.TopicsFile(project))
	if err != nil {
		return nil, fmt.Errorf(
			"read topics for %q: %w", project, err,
		)
	}

	var tf TopicsFile
	if err := yaml.Unmarshal(data, &tf); err != nil {
		return nil, fmt.Errorf("parse topics.yaml: %w", err)
	}

	q := strings.ToLower(query)
	var results []SearchResult
	for _, entry := range tf.Memory {
		if matchesTopics(entry.Topics, q) ||
			strings.Contains(strings.ToLower(entry.Summary), q) {
			results = append(results, SearchResult{
				Checkpoint: entry.Checkpoint,
				Name:       entry.Name,
				Summary:    entry.Summary,
				Topics:     entry.Topics,
			})
		}
	}
	return results, nil
}

// matchesTopics returns true if any topic contains the query substring.
func matchesTopics(topics []string, query string) bool {
	for _, t := range topics {
		if strings.Contains(strings.ToLower(t), query) {
			return true
		}
	}
	return false
}
