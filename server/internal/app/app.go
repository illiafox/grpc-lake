package app

import (
	"go.uber.org/zap"
	"server/internal/config"
	"server/internal/metrics"
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

	// Logger
	app.InitLogger()
	defer app.closers.Close(app.Logger)

	// Sentry
	sentry, err := metrics.SetupSentry(app.Config.Sentry)
	if err != nil {
		app.Logger.Error("Setup sentry", zap.Error(err))
		return
	}
	app.closers.Add(sentry, "Flushing sentry data")

	// Item Service
	item, err := app.ItemService()
	if err != nil {
		app.Logger.Error("Setup Item Service", zap.Error(err))
		return
	}

	// Start server
	app.Listen(item)
}
