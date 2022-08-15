package cache

import (
	"github.com/go-redis/redis/v9"
	"server/app/internal/adapters/cache"
	"time"
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
