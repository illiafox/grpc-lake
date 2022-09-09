package config

import (
	"flag"
)

type Flags struct {
	LogPath string
	HTTP    bool
}

func ParseFlags() Flags {
	var (
		LogPath = flag.String("log", "log.txt", "log file path")
		HTTP    = flag.Bool("http", false, "enable http server")
	)

	flag.Parse()

	return Flags{
		LogPath: *LogPath,
		HTTP:    *HTTP,
	}
}
