package app

import (
	"fmt"
	"server/app/internal/adapters/api"
	"server/app/internal/adapters/db/mongodb/item"
	"server/app/internal/adapters/db/redis/cache"
	"server/app/internal/composite"
)

func (a *App) ItemService() (api.ItemService, error) {

	m, err := composite.NewMongoComposite(a.cfg.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("mongodb: %w", err)
	}

	a.closers.Add(m, "Closing mongodb connections")

	// //

	r, err := composite.NewRedisComposite(a.cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	a.closers.Add(r, "Closing redis connections")

	// //

	return composite.NewItemService(
		item.NewItemStorage(m.
			Client().
			Database(a.cfg.MongoDB.Database).
			Collection(a.cfg.MongoDB.Collection),
		),
		//
		cache.NewCacheStorage(r.Client(), a.cfg.Cache.CacheExpire),
	), nil
}
