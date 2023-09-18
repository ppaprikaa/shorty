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
	fx.Provide(newArgs),
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(ports.Server)),
		),
	),
)

type args struct {
	cfg     *config.HttpServer
	handler http.Handler
}

func newArgs(cfg *config.HttpServer, handler http.Handler) *args {
	return &args{
		cfg:     cfg,
		handler: handler,
	}
}

type server struct {
	*http.Server
}

func New(args *args) *server {
	var (
		addr = fmt.Sprintf("%s:%d", args.cfg.Host, args.cfg.Port)
	)

	return &server{
		Server: &http.Server{
			Addr:         addr,
			Handler:      args.handler,
			ReadTimeout:  args.cfg.ReadTimeout,
			WriteTimeout: args.cfg.WriteTimeout,
			IdleTimeout:  args.cfg.IdleTimeout,
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
