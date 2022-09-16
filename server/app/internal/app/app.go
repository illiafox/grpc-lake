package app

import (
	"server/app/internal/config"
	"server/app/pkg/log"
	"server/app/pkg/log/closer"
)

type App struct {
	Logger log.Logger
	Config config.Config
	//
	closers closer.Closers
}

func Run(cfg config.Config) {
	app := App{
		Config: cfg,
	}

	app.InitLogger()
	defer app.Logger.Sync() // ignore error
	defer app.closers.Close(app.Logger)

	item, err := app.ItemService()
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}

	app.Listen(item)
}
