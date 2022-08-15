package log

import "server/app/pkg/log/zap"

type Logger interface {
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Error(msg string, fields ...Field)

	Named(s string) zap.Logger
	With(fields ...Field) zap.Logger

	Sync() error
}
