package app

import (
	"log"

	"server/app/internal/config"
)

func (app *App) ReadConfig() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	app.cfg = cfg
}
