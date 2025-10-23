package logger

import (
	"testing"
)

func setup() {
	Initialize(Config{
		OutputMode:    "console",
		ConsoleConfig: ConsoleConfig{Colorful: true},
		FileConfig: FileConfig{
			Filename:   "./logs/app.log",
			MaxSize:    64,
			MaxBackups: 10,
			MaxAge:     10,
			Compress:   false,
			LocalTime:  true,
		},
	})
}

func TestInfo(t *testing.T) {
	setup()

	Info("info")
}

func TestError(t *testing.T) {
	setup()

	data := map[string]string{
		"key": "value",
	}
	Error("info", Stack(), DictFields(data))
}

func TestWarn(t *testing.T) {
	setup()

	data := map[string]string{
		"key": "value",
	}
	Warn("Warn", Stack(), DictFields(data))
}
