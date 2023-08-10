package core

import (
	"project/core/internal"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	writeSyncer := internal.Zap.GetLogWriter()
	encoder := internal.Zap.GetEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core)
	return logger
}
