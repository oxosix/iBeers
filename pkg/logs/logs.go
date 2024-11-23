package logs

import (
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SugarLoggers struct {
	Console *zap.SugaredLogger
	JSON    *zap.SugaredLogger
}

var (
	Log        *zap.Logger
	Jsonlogger *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	// Configuração do log em console (stdout)
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  CustomColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	// Criação do logger para console
	Log, _ = logConfig.Build()

	// Se a variável de ambiente indicar, configura o JSON logger
	if shouldLogToJSON() {
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
		// Criação do logger para JSON
		Jsonlogger, _ = jsonConfig.Build()
	}
}

// Verifica se a condição de gravar em logs.json deve ser atendida
func shouldLogToJSON() bool {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	boolValue, err := strconv.ParseBool(output)

	if err != nil || !boolValue {
		return false
	}
	return true
}

// Função que retorna o nível de log conforme a variável de ambiente LOG_LEVEL
func getLogLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
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

// Função para retornar o logger com sugar (usado em outras partes da aplicação)
func Sugar() SugarLoggers {
	return SugarLoggers{
		Console: Log.Sugar(),
		JSON:    getJSONSugarLogger(),
	}
}

// Função auxiliar para garantir que Jsonlogger não seja nil
func getJSONSugarLogger() *zap.SugaredLogger {
	if Jsonlogger != nil {
		return Jsonlogger.Sugar()
	}
	return nil
}

// Funções para gravar logs em diferentes níveis
func Info(message string, tags ...zap.Field) {
	Log.Info(message, tags...)
	if Jsonlogger != nil {
		Jsonlogger.Info(message, tags...) // Só grava em JSON se Jsonlogger estiver configurado
	}
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	Log.Error(message, tags...)
	if Jsonlogger != nil {
		Jsonlogger.Error(message, tags...) // Só grava em JSON se Jsonlogger estiver configurado
	}
}

// Função auxiliar para criar um campo de log com segurança
func LogWithField(key string, value interface{}) zap.Field {
	if value == nil {
		return zap.Skip() // Ignora a chave se o valor for nil
	}
	return zap.Any(key, value)
}

func CustomColorLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var color string
	switch level {
	case zapcore.DebugLevel:
		color = "\033[32m" // Verde
	case zapcore.InfoLevel:
		color = "\033[34m" // Azul
	case zapcore.WarnLevel:
		color = "\033[33m" // Amarelo
	case zapcore.ErrorLevel:
		color = "\033[31m" // Vermelho
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		color = "\033[35m" // Roxo
	default:
		color = "\033[0m" // Reset (sem cor)
	}

	// Formatar o nível com a cor e resetar depois
	enc.AppendString(color + level.CapitalString() + "\033[0m")
}
