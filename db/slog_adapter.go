package db

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/lmittmann/tint"
)

type Logger struct {
	l *slog.Logger
}

func InitLogger() *Logger {
	var logger *slog.Logger

	if os.Getenv("APP_ENV") != "production" {
		logger = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		).With("module", "pgx")
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("module", "pgx")
	}

	return &Logger{l: logger}
}

func (l *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	var lvl slog.Level

	switch level {
	case tracelog.LogLevelError:
		lvl = slog.LevelError
	case tracelog.LogLevelWarn:
		lvl = slog.LevelWarn
	case tracelog.LogLevelInfo:
		lvl = slog.LevelInfo
	case tracelog.LogLevelDebug:
		lvl = slog.LevelDebug
	case tracelog.LogLevelTrace:
		lvl = slog.LevelDebug - 1
	case tracelog.LogLevelNone:
		lvl = slog.LevelDebug - 2
	default:
		lvl = slog.LevelDebug
		attrs = append(attrs, slog.Any("INVALID_PGX_LOG_LEVEL", level))
	}

	l.l.LogAttrs(ctx, lvl, msg, attrs...)
}
