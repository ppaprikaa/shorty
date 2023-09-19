package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"routers.chi",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(http.Handler)),
		),
	),
)

func New() *chi.Mux {
	mux := chi.NewMux()

	return mux
}
