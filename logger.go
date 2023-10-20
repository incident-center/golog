package golog

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Interface interface {
	Debug(message string, context map[string]any)
	Info(message string, context map[string]any)
	Warn(message string, context map[string]any)
	Error(message string, context map[string]any)
	Fatal(message string, context map[string]any)
}

type Logger struct {
	logger *zap.SugaredLogger
}

func New(level string) (*Logger, error) {
	config := getConfig(level)

	logger := zap.Must(config.Build())
	err := logger.Sync()
	if err != nil && !errors.Is(err, syscall.EINVAL) && !errors.Is(err, syscall.ENOTTY) {
		return nil, err
	}

	sugar := logger.Sugar()

	return &Logger{
		logger: sugar,
	}, nil
}

func (l *Logger) Debug(message string, context map[string]any) {
	l.logger.Debugw(message, zap.Int("goid", l.goid()), zap.Any("context", context))
}

func (l *Logger) Info(message string, context map[string]any) {
	l.logger.Infow(message, zap.Int("goid", l.goid()), zap.Any("context", context))
}

func (l *Logger) Warn(message string, context map[string]any) {
	l.logger.Warnw(message, zap.Int("goid", l.goid()), zap.Any("context", context))
}

func (l *Logger) Error(message string, context map[string]any) {
	l.logger.Errorw(message, zap.Int("goid", l.goid()), zap.Any("context", context))
}

func (l *Logger) Fatal(message string, context map[string]any) {
	l.logger.Fatalw(message, zap.Int("goid", l.goid()), zap.Any("context", context))
}

func getConfig(level string) zap.Config {
	return zap.Config{
		Level:    zap.NewAtomicLevelAt(getLevel(level)),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "__timestamp",
			LevelKey:      "__level",
			MessageKey:    "__message",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}
}

func getLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

func (l *Logger) goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
