package app

import (
	"context"

	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/mux"
	"github.com/ppaprikaa/shorty/internal/adapters/servers/http"
	"github.com/ppaprikaa/shorty/internal/core/ports"
	"github.com/ppaprikaa/shorty/internal/log"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

var Module = fx.Module(
	"app",
	http.Module,
	mux.Module,
	fx.Provide(New),
	fx.Invoke(func(*app) {}),
)

type app struct {
	server ports.Server
	log    *slog.Logger
}

func New(lc fx.Lifecycle, server ports.Server, logger *slog.Logger) *app {
	var app = &app{
		server: server,
		log:    logger,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error { return app.Run(ctx) },
			OnStop:  func(ctx context.Context) error { return app.Stop(ctx) },
		},
	)

	return app
}

func (a *app) Run(ctx context.Context) error {
	a.log.Info("Starting application")
	if err := a.server.Run(log.WithContext(ctx, a.log)); err != nil {
		return err
	}

	return nil
}

func (a *app) Stop(ctx context.Context) error {
	a.log.Info("Stopping application")
	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
