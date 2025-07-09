package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/shinnenkara/go-api-lib/utils"
	"go.uber.org/zap"
)

type Logger struct {
	Engine  *zap.Logger
	appName string
}

func NewLogger(appName string) *Logger {
	return &Logger{Engine: configLogger(), appName: appName}
}

func (logger *Logger) Named(name string) *Logger {
	namedLog := logger.Engine.Named(name)
	return &Logger{Engine: namedLog, appName: logger.appName}
}

func (logger *Logger) Log(message string) {
	logger.Engine.Info(message)
}

func (logger *Logger) Info(context *gin.Context, message string) {
	logger.withInfoFields(context).Info(message)
}

func (logger *Logger) Error(context *gin.Context, message string) {
	logger.withInfoFields(context).Error(message)
}

func (logger *Logger) SyncLogs() {
	err := logger.Engine.Sync()
	utils.FailOnError(err, "Failed to sync logs")
}
