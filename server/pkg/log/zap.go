package log

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(console io.Writer, files ...io.Writer) *zap.Logger {
	pe := zap.NewProductionEncoderConfig()

	// file
	pe.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)
	fileEncoder := zapcore.NewJSONEncoder(pe)

	// console
	pe.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(caller.TrimmedPath())
		encoder.AppendString("|")
	}
	//
	pe.EncodeTime = zapcore.TimeEncoderOfLayout("02/01 15:04:05 -0700") // "02/01/2006 15:04:05 |"
	pe.ConsoleSeparator = " "
	pe.EncodeName = func(n string, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(n)
		enc.AppendString("|")
	}
	//
	pe.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("|")
		enc.AppendString(l.CapitalString())
		enc.AppendString("|")
	}

	//
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	// //
	cores := make([]zapcore.Core, len(files)+1)

	// console
	cores[0] = zapcore.NewCore(consoleEncoder,
		zapcore.AddSync(console),
		zap.DebugLevel,
	)

	// // add syncers
	for i := range files {
		cores[i+1] = zapcore.NewCore(fileEncoder,
			zapcore.AddSync(files[i]),
			zap.DebugLevel,
		)
	}

	return zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
	)
}
