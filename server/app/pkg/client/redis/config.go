package redis

import (
	"runtime"
	"time"
)

type Config struct {
	Addrs       []string
	Password    string
	PoolTimeout time.Duration
	PoolSize    int
	IdleSize    int
}

func (cfg Config) Validated() Config {

	if cfg.PoolSize <= 0 {
		cfg.PoolSize = 10 * runtime.GOMAXPROCS(0)
	}

	if cfg.IdleSize <= 0 {
		cfg.IdleSize = cfg.PoolSize / 4
	}

	return cfg
}
