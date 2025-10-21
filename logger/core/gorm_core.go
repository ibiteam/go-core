package core

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/ibiteam/go-core/logger/model"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

// GormCore 实时写入数据库的日志Core（无批量缓存）
type GormCore struct {
	db *gorm.DB
	mu sync.Mutex // 并发安全锁（避免并发写入冲突）
}

// NewGormCore 创建实时写入的GORM Core
func NewGormCore(encoder zapcore.Encoder, db *gorm.DB, level zapcore.LevelEnabler) zapcore.Core {
	loggerWriter := zapcore.AddSync(&DBWriteSyncer{DB: db})
	return zapcore.NewCore(encoder, loggerWriter, level)
}

type DBWriteSyncer struct {
	DB *gorm.DB
}

func (d *DBWriteSyncer) Write(p []byte) (n int, err error) {
	var jsonData map[string]interface{}

	if unmarshalErr := json.Unmarshal(p, &jsonData); unmarshalErr != nil {
		return 0, unmarshalErr
	}

	// 将 fields 序列化为 JSON 字符串
	fieldsBytes, _ := json.Marshal(jsonData["fields"])

	message, _ := jsonData["message"].(string)
	stacktrace, _ := jsonData["stacktrace"].(string)
	timestamp, _ := jsonData["time"].(string)
	logTime, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		return 0, err
	}
	level, _ := jsonData["level"].(string)
	var errorLog = model.ErrorLog{
		Message:   message,
		Caller:    stacktrace,
		Level:     level,
		Fields:    string(fieldsBytes),
		Timestamp: logTime,
	}

	err = d.DB.Create(&errorLog).Error
	// 实现将日志数据写入数据库的逻辑
	return int(errorLog.ID), nil
}

func (d *DBWriteSyncer) Sync() error {
	// 实现同步逻辑（如果需要）
	return nil
}
