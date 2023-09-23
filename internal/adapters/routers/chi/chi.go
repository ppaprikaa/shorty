package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"routers.chi",
	fx.Provide(
		NewHandlers,
		fx.Annotate(
			New,
			fx.As(new(http.Handler)),
		),
	),
)

func NewHandlers(server *server.Handler) *handlers {
	handlers := new(handlers)
	handlers.Server = server
	return handlers
}

type handlers struct {
	Server *server.Handler
}

func New(handlers *handlers) *chi.Mux {
	mux := chi.NewMux()

	mux.Get("/api/v1/healthcheck", handlers.Server.Healthcheck())

	return mux
}
