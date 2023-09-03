package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_GetLevel(t *testing.T) {
	tests := []struct {
		name     string
		logLevel string
		expect   zapcore.Level
	}{
		{
			name:     "debugレベルのテスト",
			logLevel: "debug",
			expect:   zap.DebugLevel,
		},
		{
			name:     "infoレベルのテスト",
			logLevel: "info",
			expect:   zap.InfoLevel,
		},
		{
			name:     "warnレベルのテスト",
			logLevel: "warn",
			expect:   zap.WarnLevel,
		},
		{
			name:     "errorレベルのテスト",
			logLevel: "error",
			expect:   zap.ErrorLevel,
		},
		{
			name:     "fatalレベルのテスト",
			logLevel: "fatal",
			expect:   zap.FatalLevel,
		},
		{
			name:     "panicレベルのテスト",
			logLevel: "panic",
			expect:   zap.PanicLevel,
		},
		{
			name:     "存在しないレベルのテスト",
			logLevel: "",
			expect:   zap.DebugLevel,
		},
	}

	for _, tt := range tests {
		actual := GetLevel(tt.logLevel)
		assert.Equal(t, tt.expect, actual)
	}
}
