package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/server"
	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/user"
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

type handlers struct {
	Server *server.Handler
	User   *user.Handler
}

func NewHandlers(server *server.Handler, user *user.Handler) *handlers {
	handlers := new(handlers)
	handlers.Server = server
	handlers.User = user
	return handlers
}

func New(handlers *handlers) *chi.Mux {
	mux := chi.NewMux()

	mux.Get("/api/v1/healthcheck", handlers.Server.Healthcheck())
	mux.Post("/api/v1/user/registration", handlers.User.Registrate())

	return mux
}
