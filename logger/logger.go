package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/shinnenkara/go-api-lib/utils"
	"go.uber.org/zap"
)

type Logger struct {
	log     *zap.Logger
	appName string
}

func NewLogger(appName string) *Logger {
	return &Logger{log: configLogger(), appName: appName}
}

func (logger *Logger) Named(name string) *Logger {
	namedLog := logger.log.Named(name)
	return &Logger{log: namedLog, appName: logger.appName}
}

func (logger *Logger) Info(context *gin.Context, message string) {
	logger.withInfoFields(context).Info(message)
}

func (logger *Logger) Error(context *gin.Context, message string) {
	logger.withInfoFields(context).Error(message)
}

func (logger *Logger) SyncLogs() {
	err := logger.log.Sync()
	utils.FailOnError(err, "Failed to sync logs")
}
