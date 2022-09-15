package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger
// creates a new zap logger interface.
func NewLogger() *zap.Logger {
	level := zapcore.DebugLevel

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), level)
	cores := []zapcore.Core{
		defaultCore,
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger
}
