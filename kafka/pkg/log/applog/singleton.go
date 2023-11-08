// applogger アプリログ用の名前空間
package applog

import (
	"os"
	"sync"

	"kafka/pkg/log"
	"kafka/pkg/log/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// callerWrapNum ロガーが直接実行されるまでの関数コール数
	callerSkipNum = 2
)

var (
	logger      *log.Logger
	logLocker   sync.Once
	logFilePath string
)

// init Go言語の特殊関数
// モジュールロード時に自動的に実行される仕様を利用してSingletonを実現している。
func init() {
	logFilePath = os.Getenv("LOG_DIR") + "app.log"
	logLocker.Do(func() {
		// make config
		config := zap.Config{
			Level:    zap.NewAtomicLevelAt(log.GetLevel(os.Getenv("LOG_LEVEL"))),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				// set default log item
				TimeKey:       log.KeyTime,
				LevelKey:      log.KeyLevel,
				NameKey:       log.KeyName,
				CallerKey:     log.KeyCaller,
				MessageKey:    log.KeyMsg,
				StacktraceKey: log.KeyTrace,
				// set log expression by encoder
				EncodeLevel:    zapcore.CapitalLevelEncoder,   // Log level format: all text is upper case
				EncodeTime:     zapcore.ISO8601TimeEncoder,    // Time format: ISO8601
				EncodeDuration: zapcore.MillisDurationEncoder, // Duration format : seconds for diff calculation
				EncodeCaller:   zapcore.ShortCallerEncoder,    // Stack Trace format: path & line number after the package
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		// make file sink
		fileSink := zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   logFilePath,
				MaxSize:    100, // megabytes
				MaxBackups: 3,
				MaxAge:     28, //days
			},
		)
		fileEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)
		fileCore := zapcore.NewCore(fileEncoder, fileSink, config.Level)

		// make console sink
		consoleSink := zapcore.AddSync(os.Stdout)
		consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
		consoleCore := zapcore.NewCore(consoleEncoder, consoleSink, config.Level)

		// set options
		var options []zap.Option
		options = append(options, zap.AddCaller())
		options = append(options, zap.AddCallerSkip(callerSkipNum))
		options = append(options, zap.AddStacktrace(log.GetLevel("warn")))

		// make core
		core := zapcore.NewTee(
			fileCore,
			consoleCore,
		)

		// make log
		zapLogger := zap.New(core, options...)
		logger = log.NewLogger(zapLogger)
	})
}

// Logger Loggerインスタンスのアクセサ
// func Logger() *log.Logger {
// 	return logger
// }
