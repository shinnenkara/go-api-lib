package logger

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func (logger *Logger) LogRequest(context *gin.Context) {
	var requestBody map[string]interface{}
	if context.Request.Body != nil {
		bodyBytes, _ := io.ReadAll(context.Request.Body)
		_ = json.Unmarshal(bodyBytes, &requestBody)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	bodyMarshaller := zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		if requestBody != nil {
			for k, v := range requestBody {
				enc.AddReflected(k, v)
			}
		}

		return nil
	})

	queryMarshaller := zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		for k, v := range context.Request.URL.Query() {
			enc.AddString(k, v[0])
		}

		return nil
	})

	requestMarshaller := zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		enc.AddString("url", context.Request.URL.String())
		enc.AddString("method", context.Request.Method)
		enc.AddObject("body", bodyMarshaller)
		enc.AddObject("query", queryMarshaller)

		return nil
	})

	logger.withInfoFields(context).Info("Incoming request",
		zap.Object("request", requestMarshaller),
	)
}
