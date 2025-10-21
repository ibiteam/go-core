package logger

import (
	"fmt"
	_ "os"
	"time"

	"github.com/ibiteam/go-core/logger/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger 创建多模式日志实例
// gormDB: GORM数据库连接（启用gorm模式时必传）
// modelFactory: 日志模型工厂（可选，默认使用DefaultLogModel）
func NewLogger(cfg Config) (*zap.Logger, error) {
	// 解析日志级别
	level := new(zapcore.Level)

	// 创建编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	// 收集所有输出Core
	var cores []zapcore.Core

	// 命令行输出
	for _, mode := range cfg.OutputModes {
		switch mode {
		case "console":
			consoleEncoderConfig := encoderConfig
			if cfg.ConsoleConfig.Colorful {
				consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			}
			consoleCore := core.NewConsoleCore(
				zapcore.NewConsoleEncoder(consoleEncoderConfig),
				level,
			)
			cores = append(cores, consoleCore)
			break
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

			fmt.Println(cfg.FileConfig)

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
			break

		case "database":
			if cfg.GormConfig.Db == nil {
				panic("配置日志驱动失败,当选择数据库驱动时，请配置数据库连接")
			}
			databaseCore := core.NewGormCore(
				zapcore.NewJSONEncoder(encoderConfig),
				cfg.GormConfig.Db,
				level,
			)
			cores = append(cores, databaseCore)
			break
		}
	}

	// 合并Core并创建Logger
	tee := zapcore.NewTee(cores...)
	return zap.New(tee, zap.AddCaller()), nil // 启用调用位置信息
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
