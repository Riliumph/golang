package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_Clone(t *testing.T) {
	zapLogger, _ := zap.NewDevelopment()
	logger := New(zapLogger)
	newLogger := logger.Clone()
	assert.NotEqual(t, &logger, &newLogger)
}
