package main

import (
	"log"

	"server/internal/app"
	"server/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app.Run(cfg)
}
