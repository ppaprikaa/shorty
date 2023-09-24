package user

import (
	"github.com/ppaprikaa/shorty/internal/services/user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers.http.user",
	fx.Provide(New),
)

type Handler struct {
	userService user.Service
}

func New() *Handler {
	handler := new(Handler)
	return handler
}
