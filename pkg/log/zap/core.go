package zap

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

func NewStdCore() zapcore.Core {
	consoleWriter := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(NewConsoleEncoderConfig())
	return zapcore.NewCore(consoleEncoder, consoleWriter, infoPriority)
}

func NewFileCore(filename string) zapcore.Core {
	writer := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     90,   //days
		Compress:   true, // disabled by default
	}
	fileWriter := zapcore.Lock(zapcore.AddSync(writer))
	fileEncoder := zapcore.NewJSONEncoder(NewFileEncoderConfig())
	return zapcore.NewCore(fileEncoder, fileWriter, infoPriority)
}
