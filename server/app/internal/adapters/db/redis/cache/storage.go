package cache

import (
	"time"

	"github.com/go-redis/redis/v9"
	"server/app/internal/adapters/cache"
)

type cacheStorage struct {
	client redis.UniversalClient
	expire time.Duration
}

func NewCacheStorage(client redis.UniversalClient, expire time.Duration) cache.CacheStorage {
	return &cacheStorage{
		client: client,
		expire: expire,
	}
}
