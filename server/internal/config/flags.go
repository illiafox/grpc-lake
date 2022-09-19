package config

import (
	"flag"
)

type Flags struct {
	LogPath string
	HTTPS   bool
}

const (
	DefaultLogPath = "log.txt"
	DefaultHTTPS   = false
)

func ParseFlags() Flags {
	var (
		LogPath = flag.String("log", DefaultLogPath, "log file path")
		HTTPS   = flag.Bool("https", DefaultHTTPS, "start server with https")
	)

	flag.Parse()

	return Flags{
		LogPath: *LogPath,
		HTTPS:   *HTTPS,
	}
}
