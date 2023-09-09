package main

import (
	"how_to_zap/pkg/log/applog"

	"go.uber.org/zap"
)

func main() {
	applog.AddItem(zap.Any("request_id", "foobar"))
	applog.Debug("show parameter", zap.Any("param1", 1))
	applog.Info("show parameter", zap.Any("param1", 1))
	applog.Warn("show parameter", zap.Any("param1", 1))
	applog.Error("show parameter", zap.Any("param1", 1))
}
