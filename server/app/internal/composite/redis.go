package composite

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"server/app/internal/config"
	rdb "server/app/pkg/client/redis"
)

var _ = Composite[redis.UniversalClient](redisComposite{})

type redisComposite struct {
	client redis.UniversalClient
}

func (m redisComposite) Client() redis.UniversalClient {
	if m.client == nil {
		panic("client is nil")
	}

	return m.client
}

func (m redisComposite) Close() error {
	return m.client.Close()
}

func NewRedisComposite(cfg config.Redis) (Composite[redis.UniversalClient], error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.PoolTimeout)
	defer cancel()

	client, err := rdb.New(ctx, rdb.Config{
		Addrs:       cfg.Addrs,
		Password:    cfg.Password,
		PoolTimeout: cfg.PoolTimeout,
		PoolSize:    cfg.PoolSize,
		IdleSize:    cfg.IdleSize,
	})

	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	return redisComposite{
		client: client,
	}, nil

}
