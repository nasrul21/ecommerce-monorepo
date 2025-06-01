package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zal logger: %v", err)
	}
}

func L() *zap.Logger {
	return logger
}
