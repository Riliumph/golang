package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	levelMap = map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
		"fatal": zap.FatalLevel,
		"panic": zap.PanicLevel,
	}
)

func GetLevel(level_key string) zapcore.Level {
	val, ok := levelMap[level_key]
	if !ok {
		return zap.DebugLevel
	}
	return val
}
