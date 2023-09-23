package http

import (
	"github.com/ppaprikaa/shorty/internal/adapters/handlers/http/server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers.http",
	server.Module,
)
