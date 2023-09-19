package context

import (
	"context"

	"github.com/ppaprikaa/shorty/internal/log"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

var Module = fx.Module(
	"log.context",
	fx.Provide(Context),
)

func Context(logger *slog.Logger) context.Context {
	return log.WithContext(context.Background(), logger)
}
