package core

import (
	"queuecast/pkg/errors"
	"queuecast/pkg/utils"

	"go.uber.org/zap"
)

type Logger struct {
	level   string
	printer *zap.Logger
}

var levels = []string{"debug", "info", "warn", "error"}

func NewLogger(level string) (*Logger, error) {

	if !utils.ContainsString(levels, level) {
		return nil, errors.ErrInvalidLogLevel
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return &Logger{
		level:   level,
		printer: logger,
	}, nil
}

func (l *Logger) Close() {
	l.printer.Sync()
}

func (l *Logger) Debug(message string) {
	l.printer.Debug(message)
}
func (l *Logger) Info(message string) {
	l.printer.Info(message)
}
func (l *Logger) Warn(message string) {
	l.printer.Warn(message)
}
func (l *Logger) Error(message string) {
	l.printer.Error(message)
}

func (l *Logger) Log(message string) {
	switch l.level {
	case "debug":
		l.Debug(message)
	case "info":
		l.Info(message)
	case "warn":
		l.Warn(message)
	case "error":
		l.Error(message)
	default:
		l.Info(message)
	}
}
