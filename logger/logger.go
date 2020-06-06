package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogLevel string

const LevelDebug LogLevel = "DEBUG"
const LevelInfo LogLevel = "INFO"
const LevelWarn LogLevel = "WARN"
const LevelError LogLevel = "ERROR"
const LevelFatal LogLevel = "FATAL"

type Logger interface {
	Log(ctx context.Context, level LogLevel, message string, data map[string]string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Log(ctx context.Context, level LogLevel, message string, data map[string]string) {
	var logImpl *zerolog.Event
	switch level {
	case LevelDebug:
		logImpl = log.Debug()
	case LevelInfo:
		logImpl = log.Info()
	case LevelWarn:
		logImpl = log.Warn()
	case LevelError:
		logImpl = log.Error()
	case LevelFatal:
		logImpl = log.Fatal()
	}

	if requestID := ctx.Value("request-id"); requestID != nil {
		if rid, ok := requestID.(string); ok {
			logImpl = logImpl.Str("request_id", rid)
		}
	}

	for key, value := range data {
		logImpl = logImpl.Str(key, value)
	}
	logImpl.Msg(message)
}
