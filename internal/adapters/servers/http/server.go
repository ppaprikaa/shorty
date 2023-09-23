package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/core/ports"
	"github.com/ppaprikaa/shorty/internal/log"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

var Module = fx.Module(
	"servers.http",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(ports.Server)),
		),
	),
	fx.Invoke(func(ports.Server) {}),
)

type server struct {
	*http.Server
}

func New(ctx context.Context, lc fx.Lifecycle, cfg *config.HttpServer, handler http.Handler) *server {
	var (
		addr   = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		server = &server{
			Server: &http.Server{
				Addr:         addr,
				Handler:      handler,
				ReadTimeout:  cfg.ReadTimeout,
				WriteTimeout: cfg.WriteTimeout,
				IdleTimeout:  cfg.IdleTimeout,
			},
		}
	)

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			log.FromContext(ctx).Info("HTTP SERVER IS WORKING...")
			log.FromContext(ctx).Info("HTTP SERVER ADDRESS", slog.String("addr", server.Server.Addr))
			return server.Run(startCtx)
		},
		OnStop: func(stopCtx context.Context) error {
			log.FromContext(ctx).Info("HTTP SERVER IS SHUTTING DOWN")
			return server.Shutdown(stopCtx)
		},
	})

	return server
}

func (s *server) Run(ctx context.Context) error {
	go func() { _ = s.Server.ListenAndServe() }()

	return nil
}

func (s *server) Shutdown(ctx context.Context) (err error) {
	return s.Server.Shutdown(ctx)
}
