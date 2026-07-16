package extensions

import "errors"

// ErrExtensionNotFound is returned when the manager cannot find
// an extension with the requested name.
var ErrExtensionNotFound = errors.New("extension not found")

// ErrDuplicateName is returned by Register when an extension with
// the same name has already been registered.
var ErrDuplicateName = errors.New("extension already registered")
