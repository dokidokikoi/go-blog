package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var z *zap.Logger

func init() {
	z = NewDefaultLogger()
}

// NewDefaultLogger 获取默认logger
// 提供两种输出文件以及标准输出
func NewDefaultLogger() *zap.Logger {
	core := zapcore.NewTee(
		NewFileCore("./logs"),
		NewStdCore(),
	)
	return zap.New(core)
}

func Suger() *zap.SugaredLogger {
	return z.Sugar()
}

func L() *zap.Logger {
	return z
}
