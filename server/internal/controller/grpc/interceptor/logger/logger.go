package logger

import (
	"context"

	"go.uber.org/zap"
)

type loggerKey struct{}

// WithLogger returns a new context with the provided logger.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger returns the logger stored in the context.
//
// If no logger is stored, returns nil.
func GetLogger(ctx context.Context) *zap.Logger {
	l, _ := ctx.Value(loggerKey{}).(*zap.Logger)
	return l
}

// MustGetLogger returns the logger stored in the context.
//
// If no logger is stored, it panics with ErrNoLogger error.
func MustGetLogger(ctx context.Context) *zap.Logger {
	l, _ := ctx.Value(loggerKey{}).(*zap.Logger)

	if l == nil {
		panic(ErrNoLogger)
	}

	return l
}
