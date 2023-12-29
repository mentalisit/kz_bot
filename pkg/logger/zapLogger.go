package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Logger struct {
	ZapLogger *zap.Logger
	LoggerInterface
}

func LoggerZap(botToken string, chatID int64) *Logger {
	telegramWriter := NewTelegramWriter(botToken, chatID)
	// Определяем имя файла с логами, включающее "log", дату и время
	logFileName := fmt.Sprintf("log\\log_%s.log", time.Now().Format("2006-01-02_15-04-05"))

	// Определяем WriteSyncer для файла
	fileWriteSyncer := zapcore.AddSync(createLogFile(logFileName))

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout", logFileName},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	cfgNew := cfg.EncoderConfig
	cfgNew.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := cfg.Build(
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(core, zapcore.NewCore(
				zapcore.NewConsoleEncoder(cfgNew),
				zapcore.AddSync(telegramWriter),
				cfg.Level,
			), zapcore.NewCore(
				zapcore.NewConsoleEncoder(cfgNew),
				fileWriteSyncer,
				cfg.Level,
			))
		}),
	)

	if err != nil {
		fmt.Printf("Ошибка при создании логгера: %v\n", err)
		return nil
	}

	defer logger.Sync()
	return &Logger{ZapLogger: logger, LoggerInterface: logger}
}

type LoggerInterface interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Panic(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

//// Println выводит сообщение с использованием Info логгера
//func (l *Logger) PrintlnINFO(args ...interface{}) {
//	msg := fmt.Sprintln(args...)
//	l.ZapLogger.Info(msg)
//}
