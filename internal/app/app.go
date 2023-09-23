package app

import (
	httphandlers "github.com/ppaprikaa/shorty/internal/adapters/handlers/http"
	chirouter "github.com/ppaprikaa/shorty/internal/adapters/routers/chi"
	httpserver "github.com/ppaprikaa/shorty/internal/adapters/servers/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"app",
	httpserver.Module,
	httphandlers.Module,
	chirouter.Module,
)
