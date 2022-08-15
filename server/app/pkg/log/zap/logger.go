package zap

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func (l Logger) Named(s string) Logger {
	logger := l.Logger.Named(s)
	return Logger{logger}
}

func (l Logger) With(fields ...Field) Logger {
	logger := l.Logger.With(fields...)
	return Logger{logger}
}

func (l Logger) Sync() error {
	return l.Logger.Sync()
}
