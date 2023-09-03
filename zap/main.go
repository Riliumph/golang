package main

import (
	l "how_to_zap/pkg/logger/applogger"

	"go.uber.org/zap"
)

func main() {
	l.App().With(zap.Any("request_id", "foobar"))
	l.App().Debug("show parameter", zap.Any("param1", 1))
	l.App().Info("show parameter", zap.Any("param1", 1))
	l.App().Warn("show parameter", zap.Any("param1", 1))
	l.App().Error("show parameter", zap.Any("param1", 1))
}
