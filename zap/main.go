package main

// sample application for using zap -- https://github.com/uber-go/zap
import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main() {
	colorConsoleExample()
	sugarExample()
}

// example of  using Sugar logging
func sugarExample() {
	url := "http://uber.com"
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	sugar.Debug("this is a debug message")
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// color console output example
func colorConsoleExample() {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zapcore.DebugLevel))
	defer logger.Sync() // flushes buffer, if any
	fail := errors.New("This is an error message")

	logger.Warn("Test Warn", zap.Field{Key: "key", Type: zapcore.ErrorType, Interface: fail})
	logger.Info("Test Info")
	logger.Debug("Test Debug")
	logger.Error("Test Error")
	logger.DPanic("Test Panic")
}
