package log

import "go.uber.org/zap"

type Field = zap.Field

func Bool(key string, value bool) Field {
	return zap.Bool(key, value)
}

func Int64(key string, value int64) Field {
	return zap.Int64(key, value)
}

func Int(key string, value int) Field {
	return zap.Int(key, value)
}

func Any(key string, value any) Field {
	return zap.Any(key, value)
}

func Error(value error) Field {
	return zap.Error(value)
}

func String(key string, value string) Field {
	return zap.String(key, value)
}
