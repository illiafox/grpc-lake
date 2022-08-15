package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func New(ctx context.Context, cfg Config) (redis.UniversalClient, error) {
	cfg = cfg.Validated()

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
