package logger

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}