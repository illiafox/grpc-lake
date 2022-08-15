package app

import (
	"server/app/internal/config"
	"server/app/pkg/closer"
	"server/app/pkg/log"
	"time"
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
