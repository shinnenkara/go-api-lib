package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppWriter struct {
	logger *zap.SugaredLogger
	level  zapcore.Level
}

func (logger *Logger) Writer() *AppWriter {
	return &AppWriter{
		logger: logger.Engine.Sugar(),
		level:  zapcore.InfoLevel,
	}
}

func (gw *AppWriter) Write(p []byte) (n int, err error) {
	gw.logger.Log(gw.level, string(p))
	return len(p), nil
}
