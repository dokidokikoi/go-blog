package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	infoPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.DebugLevel
	})
)
