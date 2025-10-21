package core

import (
	"os"

	"go.uber.org/zap/zapcore"
)

// NewConsoleCore 创建命令行输出Core
func NewConsoleCore(encoder zapcore.Encoder, level zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout), // 加锁保证并发安全
		level,
	)
}
