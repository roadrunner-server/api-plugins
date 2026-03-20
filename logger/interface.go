package logger

import "log/slog"

// Named interface provides methods for obtaining a named logger scoped to a
// specific plugin or component.
type Named interface {
	// NamedLogger returns a named logger with the specified name.
	NamedLogger(name string) *slog.Logger
}
