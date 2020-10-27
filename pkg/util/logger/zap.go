package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	zap *zap.Logger
}

func newZap() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	logger.Info("aaaaa bbbbb")

	return logger
}

func NewZapLogger() Logger {
	logger := &zapLogger{}

	logger.zap = newZap()
	return logger
}

func (l *zapLogger) Debug(msg string) {
	l.zap.Debug(msg)
}
func (l *zapLogger) Info(msg string) {
	l.zap.Info(msg)
}
func (l *zapLogger) Warn(msg string) {
	l.zap.Warn(msg)
}
func (l *zapLogger) Error(msg string) {
	l.zap.Error(msg)
}
func (l *zapLogger) Panic(msg string) {
	l.zap.Panic(msg)
}
func (l *zapLogger) Fatal(msg string) {
	l.zap.Fatal(msg)
}
