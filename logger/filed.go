package logger

import "go.uber.org/zap"

func Stack() zap.Field {
	return zap.StackSkip("stacktrace", 1)
}

func DictFields(value map[string]string) zap.Field {
	return Dict("fields", value)
}

func Dict(key string, value map[string]string) zap.Field {
	var data = make([]zap.Field, 0, len(value))

	for k, item := range value {
		data = append(data, zap.String(k, item))
	}

	return zap.Dict(key, data...)
}
