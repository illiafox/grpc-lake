package app

import (
	std_log "log"
	"os"

	"server/app/pkg/log"
)

func (a *App) InitLogger() {
	file, err := os.OpenFile(a.flags.log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		std_log.Fatalf("open log file (%s): %s \n", a.flags.log, err)
	}

	logger, err := log.New(os.Stdout, file)
	if err != nil {
		std_log.Fatalf("init logger: %s \n", err)
	}

	a.closers.Add(file, "Closing log file")
	a.logger = logger
}
