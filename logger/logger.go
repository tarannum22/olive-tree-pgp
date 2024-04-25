package logger

import (
	"log/slog"
	"os"
)

func InitGlobalLogger() *slog.Logger {
	logFile, err := os.OpenFile("watermelon.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Create a TextHandler writing to the opened file
	handler := slog.NewTextHandler(logFile, opts)

	// Create a new logger with the TextHandler
	return slog.New(handler)
}
