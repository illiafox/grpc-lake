package config

import (
	"flag"
)

type Flags struct {
	ConfigPath string
	LogPath    string
	HTTPS      bool
}

const (
	DefaultLogPath = "log.txt"
	DefaultHTTPS   = false
)

func ParseFlags() Flags {
	var (
		// ConfigPath = flag.String("config", "", "Path to config file")
		LogPath = flag.String("log", DefaultLogPath, "Path to log file")
		HTTPS   = flag.Bool("https", DefaultHTTPS, "Start server with https")
	)

	flag.Parse()

	return Flags{
		// ConfigPath: *ConfigPath,
		LogPath: *LogPath,
		HTTPS:   *HTTPS,
	}
}
