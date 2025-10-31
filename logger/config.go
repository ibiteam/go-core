package logger

import (
	"github.com/ibiteam/go-core/logger/model"
	"gorm.io/gorm"
)

// Config 日志全局配置
type Config struct {
	OutputMode    string        // 输出模式: console, file, gorm
	ConsoleConfig ConsoleConfig // 命令行配置
	FileConfig    FileConfig    // 文件配置
	GormConfig    GormConfig
}

// ConsoleConfig 命令行配置
type ConsoleConfig struct {
	Colorful bool // 是否彩色输出
}

// FileConfig 文件配置
type FileConfig struct {
	Filename   string // 日志文件路径
	MaxSize    int    // 单文件最大大小(MB)
	MaxBackups int    // 最大备份数
	MaxAge     int    // 最大保留天数
	Compress   bool   // 是否压缩备份
	LocalTime  bool   // 是否使用本地时间
}

type GormConfig struct {
	Db       *gorm.DB
	LogModel model.CustomModelInterface
}
