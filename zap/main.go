package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main(){
	colorConsoleExample()
	sugarExample()
}

// example of  using Sugar logging
func sugarExample(){
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

// example of usig color for console output
func colorConsoleExample() {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zapcore.DebugLevel))
	defer logger.Sync() // flushes buffer, if any

	logger.Warn("Test Warn")
	logger.Info("Test Info")
	logger.Debug("Test Debug")
	logger.Error("Test Error")
	logger.DPanic("Test Panic")
}
