package model

import "time"

type ErrorLog struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"column:timestamp"` // 日志时间
	Level     string    `gorm:"column:level"`     // 日志级别
	Message   string    `gorm:"column:message"`   // 日志内容
	Caller    string    `gorm:"column:caller"`    // 调用位置
	Fields    string    `gorm:"column:fields"`    // 额外字段(JSON)
}

func (model ErrorLog) TableName() string {
	return "error_logs"
}
