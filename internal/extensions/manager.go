package extensions

import (
	"fmt"
	"sync"
)

// Manager holds all registered extensions. There is one global
// instance (Global) populated at startup by internal extension
// packages and by the XDG external extension scanner.
//
// Manager is safe for concurrent reads after the registration phase.
// Do not Register after the application has started serving commands.
type Manager struct {
	mu         sync.RWMutex
	extensions map[string]Extension
}

// Global is the application-wide extension registry.
// Internal extension packages call Global.Register in their init or New.
var Global = &Manager{
	extensions: make(map[string]Extension),
}

// Register adds ext to the manager under its manifest Name.
// Returns an error if the name is already registered.
func (m *Manager) Register(ext Extension) error {
	name := ext.Info().Name
	if name == "" {
		return fmt.Errorf("extension name must not be empty")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.extensions[name]; exists {
		return fmt.Errorf(
			"extension %q already registered", name,
		)
	}

	m.extensions[name] = ext
	return nil
}

// Get returns the extension registered under name, or false if absent.
func (m *Manager) Get(name string) (Extension, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ext, ok := m.extensions[name]
	return ext, ok
}

// List returns all registered extensions in an unordered slice.
func (m *Manager) List() []Extension {
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := make([]Extension, 0, len(m.extensions))
	for _, ext := range m.extensions {
		out = append(out, ext)
	}
	return out
}
