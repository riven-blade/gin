package contract

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const LogKey = "nice:log"

// Log 定义了日志服务协议
type Log interface {
	GetZap() *zap.Logger
	Panic(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Debug(msg string, fields ...zapcore.Field)
	Trace(msg string, fields ...zapcore.Field)
}
