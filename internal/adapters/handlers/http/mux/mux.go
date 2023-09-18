package mux

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers.http-mux",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(http.Handler)),
		),
	),
)

func New() *http.ServeMux {
	return http.NewServeMux()
}
