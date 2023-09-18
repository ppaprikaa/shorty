package log

import (
	"context"
	"os"

	"github.com/ppaprikaa/shorty/internal/env"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

var Module = fx.Module(
	"logger",
	fx.Provide(New),
)

func New(ENV env.ENV) *slog.Logger {
	var log *slog.Logger

	switch ENV {
	case env.LOCAL:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case env.DEV:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case env.PROD:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	return log
}

type key struct{}

func WithContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, key{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger := ctx.Value(key{})
	if l, ok := logger.(*slog.Logger); ok {
		return l
	}

	return New(env.DEV)
}
