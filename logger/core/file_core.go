package core

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

// NewFileCore 创建文件输出Core（支持轮转）
func NewFileCore(encoder zapcore.Encoder, level zapcore.LevelEnabler, filename string, maxSize, maxBackups, maxAge int, compress bool, localTime bool) zapcore.Core {
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
		LocalTime:  localTime,
	})
	return zapcore.NewCore(encoder, writeSyncer, level)
}
