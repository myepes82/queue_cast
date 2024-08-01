package tests

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetTestLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}
