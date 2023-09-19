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
)

type server struct {
	*http.Server
}

func New(cfg *config.HttpServer, handler http.Handler) *server {
	var (
		addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	)

	return &server{
		Server: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (s *server) Run(ctx context.Context) error {
	log.FromContext(ctx).Info("http server address", slog.String("addr", s.Server.Addr))
	go func() { _ = s.Server.ListenAndServe() }()

	return nil
}

func (s *server) Shutdown(ctx context.Context) (err error) {
	return s.Server.Shutdown(ctx)
}
