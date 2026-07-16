package memx

import "errors"

// ErrProjectExists is returned when Init is called for a project
// that already has a directory under MENTAL_DIR.
var ErrProjectExists = errors.New("project already exists")

// ErrProjectMissing is returned when a project directory is not
// found under MENTAL_DIR. Run mental mem init to create it.
var ErrProjectMissing = errors.New("project not initialized — run mental mem init")

// ErrTaskNotFound is returned when a task ID does not exist in tasks.yaml.
var ErrTaskNotFound = errors.New("task not found")
