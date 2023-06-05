package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()

	Logger = logger
}
