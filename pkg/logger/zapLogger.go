package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Logger struct {
	logger *zap.Logger
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
	return &Logger{logger: logger}
}

// Debug выводит сообщение отладки
func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

// Info выводит информационное сообщение
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

// Warn выводит предупреждение
func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

// Error выводит сообщение об ошибке
func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

// Println выводит сообщение с использованием Info логгера
func (l *Logger) Println(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.logger.Info(msg)
}

func (l *Logger) Panic(s string) {
	l.logger.Panic(s)
}

func (l *Logger) Fatalln(s string, err error) {
	l.logger.Fatal(s + " " + err.Error())
}
