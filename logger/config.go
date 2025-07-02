package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func configLogger() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		NameKey:        "context",
		MessageKey:     "message",
		LevelKey:       "level",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		TimeKey:        "timestamp",
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		CallerKey:      "caller",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
}

func (logger *Logger) withInfoFields(context *gin.Context) *zap.Logger {
	traceID := getTraceId(context)
	context.Set("reqID", traceID)

	return logger.log.With(
		zap.String("trace_id", traceID),
		zap.String("app_service", logger.appName),
	)
}

func getTraceId(context *gin.Context) string {
	traceID := context.Request.Header.Get("x-amzn-trace-id")
	if len(traceID) < 1 {
		traceID = uuid.New().String()
	}

	return traceID
}
