package memx

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// AddTask appends a new task to tasks.yaml with an auto-generated ID.
// Returns (taskID, nil) on success so the caller can display confirmation.
func AddTask(cfg *Config, mentalDir, project, title string) (string, error) {
	layout := NewLayout(cfg, mentalDir)

	tasks, err := loadTasks(layout, project)
	if err != nil {
		return "", err
	}

	id := nextTaskID(cfg, tasks)
	tasks = append(tasks, Task{
		ID:     id,
		Title:  title,
		Status: "todo",
	})

	if err := saveTasks(layout, project, tasks); err != nil {
		return "", err
	}
	return id, nil
}

// DoneTask marks the task with the given id as done.
// Returns ErrTaskNotFound if the id does not exist in tasks.yaml.
func DoneTask(cfg *Config, mentalDir, project, id string) error {
	layout := NewLayout(cfg, mentalDir)

	tasks, err := loadTasks(layout, project)
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "done"
			tasks[i].Completed = nowString()[:10] // YYYY-MM-DD
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf(
			"%w: %q in project %q", ErrTaskNotFound, id, project,
		)
	}

	return saveTasks(layout, project, tasks)
}

// ListTasks reads and returns all tasks for a project.
// Returns (nil, nil) when the project has no tasks.
func ListTasks(cfg *Config, mentalDir, project string) ([]Task, error) {
	layout := NewLayout(cfg, mentalDir)
	return loadTasks(layout, project)
}

// loadTasks reads and parses tasks.yaml for a project.
func loadTasks(l *Layout, project string) ([]Task, error) {
	data, err := os.ReadFile(l.TasksFile(project))
	if err != nil {
		return nil, fmt.Errorf(
			"read tasks for %q: %w", project, err,
		)
	}
	var tf TasksFile
	if err := yaml.Unmarshal(data, &tf); err != nil {
		return nil, fmt.Errorf("parse tasks.yaml: %w", err)
	}
	return tf.Tasks, nil
}

// saveTasks writes the task slice back to tasks.yaml.
func saveTasks(l *Layout, project string, tasks []Task) error {
	tf := TasksFile{Tasks: tasks}
	data, err := yaml.Marshal(tf)
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}
	return os.WriteFile(l.TasksFile(project), data, 0o644)
}

// nextTaskID generates the next task ID using the configured prefix
// and zero-padding. Example: "t001", "t002", etc.
func nextTaskID(cfg *Config, tasks []Task) string {
	max := 0
	prefix := cfg.Tasks.IDPrefix

	for _, t := range tasks {
		raw := strings.TrimPrefix(t.ID, prefix)
		n := 0
		_, _ = fmt.Sscanf(raw, "%d", &n)
		if n > max {
			max = n
		}
	}

	return fmt.Sprintf("%s%0*d", prefix, cfg.Tasks.IDPadding, max+1)
}
