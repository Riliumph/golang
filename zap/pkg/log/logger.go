package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	details_obj = "details"
)

// Logger ロガークラス
// zapのLoggerを委譲している。
type Logger struct {
	logger *zap.Logger
}

// NewLogger Loggerのコンストラクタ
func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Clone DeepCopyするためのメソッド
func (l *Logger) Clone() *Logger {
	if l == nil {
		return nil
	}
	copy := *l
	return &copy
}

// With 項目を付与する関数
// 実行以降、必ず追加した項目がログに現れるようになる。
// RequestIDなど、並行並列で動作する際にIDを振るために利用される。
func (l *Logger) With(fields ...zap.Field) {
	if l == nil || l.logger == nil {
		// avoid unexpected nil access
		// ex) test code
		return
	}
	l.logger = l.logger.With(fields...)
}

func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	if l == nil || l.logger == nil {
		return
	}
	exe := l.Clone()
	if len(fields) != 0 {
		// 一時的にjsonにセクションを追加する。
		exe.With(zap.Namespace(details_obj))
	}
	defer exe.logger.Sync()
	exe.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	exe := l.Clone()
	if len(fields) != 0 {
		exe.With(zap.Namespace(details_obj))
	}
	defer exe.logger.Sync()
	exe.logger.Debug(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	exe := l.Clone()
	if len(fields) != 0 {
		exe.With(zap.Namespace(details_obj))
	}
	defer exe.logger.Sync()
	exe.logger.Debug(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	exe := l.Clone()
	if len(fields) != 0 {
		exe.With(zap.Namespace(details_obj))
	}
	defer exe.logger.Sync()
	exe.logger.Debug(msg, fields...)
}
