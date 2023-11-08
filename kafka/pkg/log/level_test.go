package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_GetLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name   string
		arg    args
		expect zapcore.Level
	}{
		{
			name: "debugレベルのテスト",
			arg: args{
				level: "debug",
			},
			expect: zap.DebugLevel,
		},
		{
			name: "infoレベルのテスト",
			arg: args{
				level: "info",
			},
			expect: zap.InfoLevel,
		},
		{
			name: "warnレベルのテスト",
			arg: args{
				level: "warn",
			},
			expect: zap.WarnLevel,
		},
		{
			name: "errorレベルのテスト",
			arg: args{
				level: "error",
			},
			expect: zap.ErrorLevel,
		},
		{
			name: "fatalレベルのテスト",
			arg: args{
				level: "fatal",
			},
			expect: zap.FatalLevel,
		},
		{
			name: "panicレベルのテスト",
			arg: args{
				level: "panic",
			},
			expect: zap.PanicLevel,
		},
		{
			name: "存在しないレベルのテスト",
			arg: args{
				level: "",
			},
			expect: zap.DebugLevel,
		},
	}

	for _, tt := range tests {
		actual := GetLevel(tt.arg.level)
		assert.Equal(t, tt.expect, actual)
	}
}
