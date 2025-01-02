package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello %s age %d", "Long", 26)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Long"), zap.Int("age", 26))

	encoder := getEncoderLog()
	writeSync := getWriteSync()
	core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Hello", zap.String("name", "Long"), zap.Int("age", 26))
	logger.Error("error log", zap.Int("hehe", 123))
}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile(("./log/log.txt"), os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
