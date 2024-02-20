package logs

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SugarLoggers struct {
	Console *zap.SugaredLogger
	JSON    *zap.SugaredLogger
}

var (
	log        *zap.Logger
	jsonlogger *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	jsonConfig := zap.Config{
		OutputPaths: []string{"logs.json"},
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
	jsonlogger, _ = jsonConfig.Build()
}

func Sugar() SugarLoggers {
	return SugarLoggers{
		Console: log.Sugar(),
		JSON:    jsonlogger.Sugar(),
	}
}

func JSON(key string, value string) zap.Field {
	return zap.String(key, value)
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	jsonlogger.Info(message, tags...)
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	jsonlogger.Error(message, tags...)
}

//	func getOutputLogs() []string {
//		output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))
//		if output == "" {
//			return []string{"stdout"}
//		}
//		if output == "json" {
//			return []string{"logs.json"}
//		}
//		if output == "both" {
//			return []string{"stdout", "logs.json"}
//		}
//		return []string{output}
//	}
func getLogLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}

}
