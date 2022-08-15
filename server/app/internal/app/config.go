package app

import (
	"log"
	"server/app/internal/config"
)

func (a *App) ReadConfig() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	a.cfg = cfg
}
