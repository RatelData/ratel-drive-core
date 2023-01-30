package util

import "go.uber.org/zap"

var loggerInst *zap.Logger

func InitLogger() {
	if loggerInst == nil {
		loggerInst, _ = zap.NewProduction()
	}
}

func GetLogger() *zap.Logger {
	return loggerInst
}
