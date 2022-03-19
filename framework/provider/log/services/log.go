package services

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewNiceLog(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	config := container.MustMake(contract.ConfigKey).(contract.Config)
	return &NiceLog{c: container, Zap: Zap(config)}, nil
}

// NiceLog 的通用实例
type NiceLog struct {
	// 五个必要参数
	Zap *zap.Logger
	c   framework.Container // 容器
}

func (log *NiceLog) Panic(msg string, fields ...zapcore.Field) {
	log.Zap.Panic(msg, fields...)
}

func (log *NiceLog) GetZap() *zap.Logger {
	return log.Zap
}

func (log *NiceLog) Error(msg string, fields ...zapcore.Field) {
	log.Zap.Error(msg, fields...)
}

func (log *NiceLog) Warn(msg string, fields ...zapcore.Field) {
	log.Zap.Warn(msg, fields...)
}

func (log *NiceLog) Info(msg string, fields ...zapcore.Field) {
	log.Zap.Info(msg, fields...)
}

func (log *NiceLog) Debug(msg string, fields ...zapcore.Field) {
	log.Zap.Debug(msg, fields...)
}

func (log *NiceLog) Fatal(msg string, fields ...zapcore.Field) {
	log.Zap.Fatal(msg, fields...)
}

func (log *NiceLog) Trace(msg string, fields ...zapcore.Field) {
	log.Zap.Info(msg, fields...)
}
