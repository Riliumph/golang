package applog

import "go.uber.org/zap/zapcore"

// AddItem jsonの項目を追加するヘルパーメソッド
func AddItem(fields ...zapcore.Field) {
	logger.With(fields...)
}

// Debug Debugレベルログを書き出すヘルパーメソッド
func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

// Info Infoレベルログを書き出すヘルパーメソッド
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// Warn Warnレベルログを書き出すヘルパーメソッド
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

// Error Errorレベルログを書き出すヘルパーメソッド
func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}
