package orm

import (
	"context"
	"github.com/ddh-open/gin/framework/contract"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

// LoggerORM orm的日志实现类, 实现了gorm.Logger.Interface
type LoggerORM struct {
	logger contract.Log // 有一个logger对象存放nice的log服务
}

// NewOrmLogger 初始化一个ormLogger,
func NewOrmLogger(logger contract.Log) *LoggerORM {
	return &LoggerORM{logger: logger}
}

// LogMode 什么都不实现，日志级别完全依赖nice的日志定义
func (o *LoggerORM) LogMode(level logger.LogLevel) logger.Interface {
	return o
}

// Info 对接nice的info输出
func (o *LoggerORM) Info(ctx context.Context, s string, i ...interface{}) {
	o.logger.Info(s, zap.Any("message:", i))
}

// Warn 对接nice的Warn输出
func (o *LoggerORM) Warn(ctx context.Context, s string, i ...interface{}) {
	o.logger.Warn(s, zap.Any("message:", i))
}

// Error 对接nice的Error输出
func (o *LoggerORM) Error(ctx context.Context, s string, i ...interface{}) {
	o.logger.Error(s, zap.Any("message:", i))
}

// Trace 对接nice的Trace输出
func (o *LoggerORM) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)
	s := "orm trace sql"
	o.logger.Trace(s, zap.String("begin:", begin.String()), zap.Any("err:", err), zap.String("sql", sql), zap.Int64("rows", rows), zap.String("elapsed", elapsed.String()))
}
