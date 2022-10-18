package logger

import "errors"

// ErrNoLogger is returned when no logger is stored in the context.
var ErrNoLogger = errors.New("logger not provided")
