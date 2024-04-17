package logger

import (
	"io"
	"log/slog"
	"os"
	"shorty/internal/gears/config"
)

func New(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case config.EnvLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	case config.EnvDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelWarn,
		}))
	case config.EnvProd:
		writer := io.MultiWriter(os.Stdout, newFileWriter("logs"))

		logger = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelError,
		}))
	}

	return logger
}
