package app

import (
	"flag"
	"runtime"
	"time"
)

func Init() *App {

	// // flags

	var (
		logPath = flag.String("log", "log.txt", "log file path")
		http    = flag.Bool("http", false, "enable http server")
		connect = flag.Duration("connect", time.Second*5, "database connect timeout")
	)
	flag.Parse()

	// // closer

	app := App{
		flags: flags{
			log:            *logPath,
			http:           *http,
			connectTimeout: *connect,
		},
	}

	defer runtime.GC() // force garbage collector to clear unused flag pointers

	return &app
}
