package composite

import (
	"server/app/internal/adapters/api"
	"server/app/internal/domain/service/cache"
	"server/app/internal/domain/service/event"
	"server/app/internal/domain/usecase"
)

func NewItemUsecase(itemStorage cache.ItemStorage, cacheStorage cache.CacheStorage, eventStorage event.MessageStorage) api.ItemUsecase {
	cacheService := cache.NewCacheService(itemStorage, cacheStorage)
	eventService := event.NewEventService(eventStorage)

	return usecase.NewItemUsecase(cacheService, eventService)
}
