package main

import (
	"github.com/ppaprikaa/shorty/internal/app"
	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/db/mongo"
	"github.com/ppaprikaa/shorty/internal/db/redis"
	"github.com/ppaprikaa/shorty/internal/log"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		log.Module,
		// Base log context
		// Couldn't initialize, because it is not obvious that module returns base context
		fx.Provide(log.Context),
		mongo.Module,
		redis.Module,
		app.Module,
	).Run()
}
