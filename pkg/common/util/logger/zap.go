package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represents logging interface
type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Panic(string)
	Fatal(string)
}

type log struct {
	log *zap.Logger
}

// New instantiates new zap logger
func New() Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger, _ := config.Build()

	return &log{
		log: zapLogger,
	}
}

func (l log) Debug(msg string) {
	l.log.Debug(msg)
}

func (l log) Info(msg string) {
	l.log.Info(msg)
}

func (l log) Warn(msg string) {
	l.log.Warn(msg)
}

func (l log) Error(msg string) {
	l.log.Error(msg)
}

func (l log) Panic(msg string) {
	l.log.Panic(msg)
}

func (l log) Fatal(msg string) {
	l.log.Fatal(msg)
}
