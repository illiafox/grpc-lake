package main

import apps "server/app/internal/app"

func main() {
	app := apps.Init()
	app.Run()
}
