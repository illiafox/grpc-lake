package composite

import (
	"server/internal/adapters/api"
	cache2 "server/internal/domain/service/cache"
	event2 "server/internal/domain/service/event"
	item2 "server/internal/domain/usecase/item"
)

func NewItemUsecase(itemStorage item2.ItemService, cacheStorage cache2.CacheStorage, eventStorage event2.MessageStorage) api.ItemUsecase {
	cacheService := cache2.NewCacheService(itemStorage, cacheStorage)
	eventService := event2.NewEventService(eventStorage)

	return item2.NewItemUsecase(cacheService, eventService)
}
