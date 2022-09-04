package app

import (
	"server/app/pkg/log/closer"
	"time"

	"server/app/internal/config"
	"server/app/pkg/log"
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

func (a *App) Run() {
	a.ReadConfig()
	a.InitLogger()
	//
	a.Listen()
}
