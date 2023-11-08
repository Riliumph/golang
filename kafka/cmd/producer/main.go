package main

import (
	"fmt"
	"kafka/pkg/log/applog"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("hello, producer")
	applog.AddItem(zap.Any("request_id", "foobar"))
	applog.Debug("show parameter", zap.Any("param1", 1))
	applog.Info("show parameter", zap.Any("param1", 1))
	applog.Warn("show parameter", zap.Any("param1", 1))
	applog.Error("show parameter", zap.Any("param1", 1))
}
