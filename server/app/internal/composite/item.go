package composite

import (
	"server/app/internal/adapters/api"
	"server/app/internal/adapters/cache"
	"server/app/internal/domain/service"
)

func NewItemService(itemStorage service.ItemStorage, cacheStorage cache.CacheStorage) api.ItemService {
	itemStorage = cache.NewCacheWrapper(itemStorage, cacheStorage)

	return service.NewItemService(itemStorage)
}
