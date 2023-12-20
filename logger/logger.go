package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	config.DisableStacktrace = true
	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field){
	log.Info(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field){
	log.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field){
	log.Error(msg, fields...)
}
func Warn(msg string, fields ...zap.Field){
	log.Warn(msg, fields...)
}