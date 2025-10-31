package logger

import (
	_ "os"
	"sync"
	"time"

	"github.com/ibiteam/go-core/logger/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerUtil *zap.Logger
	once       sync.Once
)

// Initialize 创建多模式日志实例
func Initialize(cfg Config) {
	once.Do(func() {
		tee := getCoreTee(cfg)
		loggerUtil = zap.New(tee, zap.AddCaller())
	})
}

func getCoreTee(cfg Config) zapcore.Core {
	// 解析日志级别
	level := new(zapcore.Level)
	// 创建编码器
	encoderConfig := getEncoderConfig()
	// 收集所有输出Core
	var cores []zapcore.Core

	// 命令行输出
	switch cfg.OutputMode {
	case "console":
		consoleEncoderConfig := encoderConfig
		if cfg.ConsoleConfig.Colorful {
			consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		consoleCore := core.NewConsoleCore(zapcore.NewConsoleEncoder(consoleEncoderConfig), level)

		cores = append(cores, consoleCore)

	case "file":
		if cfg.FileConfig.Filename == "" {
			panic("配置日志驱动失败,当选择文件驱动时，请配置日志文件名称")
		}
		if cfg.FileConfig.MaxSize <= 0 {
			panic("配置日志驱动失败,当选择文件驱动时，请配置日志文件最大大小")
		}
		if cfg.FileConfig.MaxBackups <= 0 {
			panic("配置日志驱动失败,当选择文件驱动时，请配置日志文件最大备份数")
		}
		if cfg.FileConfig.MaxAge <= 0 {
			panic("配置日志驱动失败,当选择文件驱动时，请配置日志文件最大保留天数")
		}

		fileCore := core.NewFileCore(
			zapcore.NewJSONEncoder(encoderConfig),
			level,
			cfg.FileConfig.Filename,
			cfg.FileConfig.MaxSize,
			cfg.FileConfig.MaxBackups,
			cfg.FileConfig.MaxAge,
			cfg.FileConfig.Compress,
			cfg.FileConfig.LocalTime,
		)

		cores = append(cores, fileCore)

	case "database":
		if cfg.GormConfig.Db == nil {
			panic("配置日志驱动失败,当选择数据库驱动时，请配置数据库连接")
		}

		databaseCore := core.NewGormCore(zapcore.NewJSONEncoder(encoderConfig), cfg.GormConfig.Db, cfg.GormConfig.LogModel, level)

		cores = append(cores, databaseCore)
	}

	if len(cores) == 0 {
		panic("配置日志驱动失败,请选择正确的日志驱动")
	}

	// 合并Core并创建Logger
	return zapcore.NewTee(cores...)
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,   // 每行日志的结尾添加 "\n"
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 日志级别名称大写，如 ERROR、INFO
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05"))
		}, // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}
}

func Error(message string, fields ...zap.Field) {
	if loggerUtil == nil {
		return
	}
	loggerUtil.Error(message, fields...)
}

// Info 告知类日志
func Info(message string, fields ...zap.Field) {
	if loggerUtil == nil {
		return
	}
	loggerUtil.Info(message, fields...)
}

// Warn 警告类日志
func Warn(message string, fields ...zap.Field) {
	if loggerUtil == nil {
		return
	}
	loggerUtil.Warn(message, fields...)
}
