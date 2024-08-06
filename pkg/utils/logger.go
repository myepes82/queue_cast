package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:       "time",
	LevelKey:      "level",
	NameKey:       "logger",
	CallerKey:     "",
	MessageKey:    "message",
	StacktraceKey: "stacktrace",
	EncodeLevel:   zapcore.LowercaseLevelEncoder,
	EncodeTime:    zapcore.RFC3339TimeEncoder,
	EncodeCaller:  zapcore.ShortCallerEncoder,
}

func GetLoggerFactory() *zap.Logger {
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)

	return zap.New(core)
}
