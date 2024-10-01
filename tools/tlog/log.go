package tlog

import (
	"fmt"
	"log/slog"
)

func Debug(message string, args ...any) {
	slog.Debug(fmt.Sprintf(message, args...))
}

func Info(message string, args ...any) {
	slog.Info(fmt.Sprintf(message, args...))
}

func Warn(message string, args ...any) {
	slog.Warn(fmt.Sprintf(message, args...))
}

func Error(message string, args ...any) {
	slog.Error(fmt.Sprintf(message, args...))
}
