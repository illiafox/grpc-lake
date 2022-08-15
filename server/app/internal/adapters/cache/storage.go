package cache

import (
	"server/app/internal/domain/service"
)

type cacheWrapper struct {
	item  service.ItemStorage
	cache CacheStorage
}

func NewCacheWrapper(storage service.ItemStorage, cache CacheStorage) service.ItemStorage {
	return cacheWrapper{
		item:  storage,
		cache: cache,
	}
}
