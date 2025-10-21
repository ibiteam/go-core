package logger

import (
	"testing"

	"gorm.io/gorm"
)

var DB *gorm.DB

func TestInfo(t *testing.T) {

	logger, _ := NewLogger(Config{
		OutputModes:   []string{"console", "file"},
		ConsoleConfig: ConsoleConfig{Colorful: true},
		GormConfig:    GormConfig{Db: DB},
		FileConfig: FileConfig{
			Filename:   "./logs/app.log",
			MaxSize:    64,
			MaxBackups: 10,
			MaxAge:     10,
			Compress:   false,
			LocalTime:  true,
		},
	})
	data := map[string]string{
		"key": "value",
	}
	logger.Info("info", Stack(), DictFields(data))
}
