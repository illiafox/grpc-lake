package middleware

import (
	"context"
	"server/app/pkg/log"
)

type loggerKey struct{}

// WithLogger returns a new context with the provided logger.
func WithLogger(ctx context.Context, logger log.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger returns the logger stored in the context.
//
// If no logger is stored, returns nil.
func GetLogger(ctx context.Context) log.Logger {
	l, _ := ctx.Value(loggerKey{}).(log.Logger)
	return l
}

// MustGetLogger returns the logger stored in the context.
//
// If no logger is stored, it panics with ErrNoLogger error.
func MustGetLogger(ctx context.Context) log.Logger {
	l, _ := ctx.Value(loggerKey{}).(log.Logger)

	if l == nil {
		panic(ErrNoLogger)
	}

	return l
}
