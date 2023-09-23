package server

import "go.uber.org/fx"

var Module = fx.Module(
	"handlers.http.server",
	fx.Provide(New),
)

type Handler struct{}

func New() *Handler {
	handler := new(Handler)

	return handler
}
