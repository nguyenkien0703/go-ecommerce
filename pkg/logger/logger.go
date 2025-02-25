package logger

import (
	"example.com/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Level //"debug" // debug -> info -> warn -> error -> dpanic -> panic -> fatal
	// debug -> infor -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.Filename,   // Path to log file
		MaxSize:    config.MaxSize,    // Max megabytes before rotation
		MaxBackups: config.MaxBackups, // Max number of old log files to keep
		MaxAge:     config.MaxAge,     // Max number of days to retain old files
		Compress:   config.Compress,   // Whether to compress old files
	}
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// /format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.TimeKey = "time"

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
