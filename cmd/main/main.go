package main

import (
	"github.com/ppaprikaa/shorty/internal/app"
	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/db/mongo"
	"github.com/ppaprikaa/shorty/internal/db/redis"
	"github.com/ppaprikaa/shorty/internal/log"
	loggercontext "github.com/ppaprikaa/shorty/internal/log/context"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		log.Module,
		loggercontext.Module,
		mongo.Module,
		redis.Module,
		app.Module,
	).Run()
}
