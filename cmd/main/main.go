package main

import (
	"github.com/ppaprikaa/shorty/internal/app"
	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/log"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		log.Module,
		app.Module,
	).Run()
}
