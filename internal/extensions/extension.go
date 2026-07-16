// Package extensions defines the contract for mental extensions and
// manages their lifecycle — both built-in and external.
//
// # Architecture
//
// Mental supports two kinds of extensions:
//
//   - Internal: compiled into the binary, registered at startup.
//     Code lives in internal/extensions/<name>/.
//   - External: standalone executables discovered at runtime from
//     $MENTAL_DIR/extensions/<name>/ via an extension.yaml manifest.
//
// All extensions satisfy the Extension interface, enabling uniform
// registration, listing, and dispatch regardless of kind.
//
// # Adding an Internal Extension
//
// 1. Create a package under internal/extensions/<name>/.
// 2. Implement the Extension interface.
// 3. Register via Manager.Register in your package's init or New.
// 4. Call your New function from internal/extensions/manager.go.
//
// # Data Exchange (External Extensions)
//
// External extensions in "structured" mode communicate via JSON on
// stdin/stdout. mental writes a JSON request to the extension's stdin
// and reads a JSON response from its stdout. See runner.go for details.
//
// Extensions in "passthrough" mode own the terminal directly; mental
// does not capture their output.
package extensions

import "context"

// Mode controls how mental communicates with an external extension.
type Mode string

const (
	// ModePassthrough wires the extension directly to the terminal.
	// mental does not capture output. Use for display-only extensions.
	ModePassthrough Mode = "passthrough"

	// ModeStructured captures the extension's stdout as JSON.
	// mental reads the response for further processing or formatting.
	ModeStructured Mode = "structured"
)

// Manifest holds the metadata declared in an extension.yaml file.
// Internal extensions populate this programmatically; external
// extensions have it parsed from their manifest file.
type Manifest struct {
	// Name is the human-readable extension name.
	Name string `yaml:"name"`

	// Type classifies the extension's role: memory, task, search, etc.
	Type string `yaml:"type"`

	// Description is a one-line summary shown in "mental extensions list".
	Description string `yaml:"description"`

	// Executable is the binary name for external extensions.
	// Empty for internal extensions.
	Executable string `yaml:"executable,omitempty"`

	// Author is the extension author (name or org).
	Author string `yaml:"author,omitempty"`

	// Version is the extension version string (semver recommended).
	Version string `yaml:"version,omitempty"`

	// Mode controls how mental communicates with an external extension.
	// Internal extensions ignore this field.
	Mode Mode `yaml:"mode,omitempty"`
}

// Extension is the contract all mental extensions implement.
// Built-in extensions call Register on the global Manager.
type Extension interface {
	// Info returns the extension manifest for listing and introspection.
	Info() Manifest

	// Run executes the extension with the given arguments.
	// ctx carries a deadline; extensions must respect cancellation.
	Run(ctx context.Context, args []string) error
}
