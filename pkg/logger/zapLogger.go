package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Logger struct {
	ZapLogger *zap.Logger
}

func LoggerZap(botToken string, chatID int64) *Logger {
	telegramWriter := NewTelegramWriter(botToken, chatID)
	discordWriter := NewDiscordWriter("https://discord.com/api/webhooks/1198796243032358973/zI1cqrJg94jEHFS-9rASi-9gwlj4aqu3xz-Fy1RLIP_TAVm7JjJClSuMF3DUHTasyqwT")

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
			), zapcore.NewTee(zapcore.NewCore(
				zapcore.NewConsoleEncoder(cfgNew),
				zapcore.AddSync(discordWriter),
				cfg.Level,
			), zapcore.NewCore(
				zapcore.NewConsoleEncoder(cfgNew),
				fileWriteSyncer,
				cfg.Level,
			)))
		}),
	)

	if err != nil {
		fmt.Printf("Ошибка при создании логгера: %v\n", err)
		return nil
	}

	defer logger.Sync()
	return &Logger{ZapLogger: logger} // LoggerInterface: logger}
}
func LoggerZapDEV() *Logger {
	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"},
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
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}

	defer logger.Sync()

	logger.Info("Develop Running")

	return &Logger{ZapLogger: logger}
}

func (l *Logger) ErrorErr(err error) {
	l.ZapLogger.Error("Произошла ошибка", zap.Error(err))
}
func (l *Logger) Debug(s string) {
	l.ZapLogger.Debug(s)
}
func (l *Logger) Info(s string) {
	l.ZapLogger.Info(s)
}
func (l *Logger) Warn(s string) {
	l.ZapLogger.Warn(s)
}
func (l *Logger) Error(s string) {
	l.ZapLogger.Error(s)
}
func (l *Logger) Panic(s string) {
	l.ZapLogger.Panic(s)
}
func (l *Logger) Fatal(s string) {
	l.ZapLogger.Fatal(s)
}

func (l *Logger) InfoStruct(s string, i interface{}) {
	l.ZapLogger.Info(fmt.Sprintf("%s: %+v \n", s, i))
}

//func (l *Logger) Log(s string,err error)  {
//
//}
//func (l *Logger) Log(s string,err error)  {
//
//}
