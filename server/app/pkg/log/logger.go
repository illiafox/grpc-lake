package log

import (
	"io"

	"server/app/pkg/log/zap"
)

// New creates a new logger.
//
// The first argument is the console writer, which logger writes in user-friendly format.
// Others are for optional json formatted logs.
func New(console io.Writer, files ...io.Writer) (Logger, error) {
	return zap.NewLogger(console, files...), nil
}
