package composite

import (
	"context"
	"fmt"
	"runtime"

	"github.com/go-redis/redis/v9"
	"server/internal/config"
)

func NewRedisComposite(cfg config.Redis) (redis.UniversalClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.PoolTimeout)
	defer cancel()

	if cfg.PoolSize <= 0 {
		cfg.PoolSize = 10 * runtime.GOMAXPROCS(0)
	}

	if cfg.IdleSize <= 0 {
		cfg.IdleSize = cfg.PoolSize / 4
	}

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        cfg.Addrs,
		Password:     cfg.Password,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  cfg.PoolTimeout,
		MinIdleConns: cfg.IdleSize,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return client, nil
}
