package app

import (
	"go.uber.org/zap"
	"server/internal/config"
	"server/pkg/log/closer"
)

type App struct {
	Logger *zap.Logger
	Config config.Config
	//
	closers closer.Closers
}

func Run(cfg config.Config) {
	app := App{
		Config: cfg,
	}
	app.InitLogger()

	defer app.closers.Close(app.Logger)

	item, err := app.ItemService()
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}

	app.Listen(item)
}
