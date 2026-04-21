// Package logger
package logger

import (
	"log/slog"
	"os"
	"server/internal/err/panics"
)

var Logger *slog.Logger

func init() {
	f, err := os.Open("logs.txt")
	panics.PanicErr("failed to open logs file", err)
	slog.SetLogLoggerLevel(slog.LevelDebug)
	consoleHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	fileHandler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	multiHandler := slog.NewMultiHandler(consoleHandler, fileHandler)
	Logger = slog.New(multiHandler)
}
