package app

import (
	"time"

	"server/app/internal/config"
	"server/app/pkg/log"
	"server/app/pkg/log/closer"
)

type flags struct {
	log            string
	http           bool
	connectTimeout time.Duration
}

type App struct {
	flags flags
	//
	logger log.Logger
	cfg    config.Config
	//
	closers closer.Closers
}

func (app *App) Run() {
	app.ReadConfig()
	app.InitLogger()
	//
	app.Listen()
}
