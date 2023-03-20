package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = getLogger()

func getLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
