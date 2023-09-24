package http

import (
	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/server"
	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers.http",
	server.Module,
	user.Module,
)
