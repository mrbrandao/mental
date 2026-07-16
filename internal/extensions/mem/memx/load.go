package memx

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Load reads MEMORY.md and tasks.yaml for a project and returns
// a ProjectContext ready to be printed to LLM context.
func Load(cfg *Config, mentalDir, project string) (*ProjectContext, error) {
	layout := NewLayout(cfg, mentalDir)

	memory, err := readMemory(layout, project)
	if err != nil {
		return nil, fmt.Errorf("read memory: %w", err)
	}

	tasks, err := readTasks(layout, project)
	if err != nil {
		return nil, fmt.Errorf("read tasks: %w", err)
	}

	return &ProjectContext{
		Project: project,
		Memory:  memory,
		Tasks:   tasks,
	}, nil
}

// PrintContext is removed — callers use cmd/ pterm rendering instead.

// readMemory reads MEMORY.md content for a project.
func readMemory(l *Layout, project string) (string, error) {
	data, err := os.ReadFile(l.MemoryFile(project))
	if err != nil {
		return "", fmt.Errorf(
			"read %s: %w — run mental mem init %s first",
			l.MemoryFile(project), err, project,
		)
	}
	return string(data), nil
}

// readTasks reads and parses tasks.yaml for a project.
func readTasks(l *Layout, project string) ([]Task, error) {
	data, err := os.ReadFile(l.TasksFile(project))
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", l.TasksFile(project), err)
	}

	var tf TasksFile
	if err := yaml.Unmarshal(data, &tf); err != nil {
		return nil, fmt.Errorf("parse tasks.yaml: %w", err)
	}
	return tf.Tasks, nil
}
