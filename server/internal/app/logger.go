package app

import (
	std_log "log"
	"os"

	"server/pkg/log"
)

func (app *App) InitLogger() {
	file, err := os.OpenFile(app.Config.Flags.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		std_log.Fatalf("open log file (%s): %s \n", app.Config.Flags.LogPath, err)
	}

	logger := log.NewLogger(os.Stdout, file)
	if err != nil {
		std_log.Fatalf("init logger: %s \n", err)
	}

	app.closers.Add(func() error {
		_ = logger.Sync()
		return file.Close()
	}, "Closing log file")

	app.Logger = logger
}
