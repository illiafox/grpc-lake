package main

import (
	"log"

	"server/app/internal/app"
	"server/app/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app.Run(cfg)
}
