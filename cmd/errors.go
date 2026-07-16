// Package cmd wires the mental CLI using Cobra.
// This file holds sentinel errors shared across cmd/ subpackages.
package cmd

import "errors"

// ErrStdinIsTerminal is returned when a command reads from stdin
// but stdin is a terminal (i.e., no data is being piped in).
var ErrStdinIsTerminal = errors.New(
	"stdin is a terminal — pipe input or use provider flags (-a/-s)",
)

// ErrUnknownProvider is returned when an -a/--agent flag value
// does not match any registered session provider.
var ErrUnknownProvider = errors.New("unknown session provider")

// ErrUnknownEngine is returned when a --engine flag value does not
// match any registered memory engine.
var ErrUnknownEngine = errors.New("unknown memory engine")
