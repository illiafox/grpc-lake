package app

import (
	std_log "log"
	"os"

	"server/app/pkg/log"
)

func (app *App) InitLogger() {
	file, err := os.OpenFile(app.flags.log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		std_log.Fatalf("open log file (%s): %s \n", app.flags.log, err)
	}

	logger, err := log.New(os.Stdout, file)
	if err != nil {
		std_log.Fatalf("init logger: %s \n", err)
	}

	app.closers.Add(file, "Closing log file")
	app.logger = logger
}
