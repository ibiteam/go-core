package model

import "time"

type CustomModelInterface interface {
	TableName() string
	SetLogData(message, level, caller, fields string, timestamp time.Time)
}
