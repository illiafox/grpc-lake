package composite

import (
	"server/internal/adapters/api"
	"server/internal/domain/service/cache"
	"server/internal/domain/service/event"
	"server/internal/domain/usecase/item"
)

func NewItemUsecase(itemStorage item.ItemService, cacheStorage cache.CacheStorage, eventStorage event.MessageStorage) api.ItemUsecase {
	cacheService := cache.NewCacheService(itemStorage, cacheStorage)
	eventService := event.NewEventService(eventStorage)

	return item.NewItemUsecase(cacheService, eventService)
}
