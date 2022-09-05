package app

import (
	"fmt"

	"server/app/internal/adapters/api"
	"server/app/internal/adapters/db/mongodb/item"
	"server/app/internal/adapters/db/redis/cache"
	"server/app/internal/composite"
)

func (app *App) ItemService() (api.ItemService, error) {

	m, err := composite.NewMongoComposite(app.cfg.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("mongodb: %w", err)
	}

	app.closers.Add(m, "Closing mongodb connections")

	// //

	r, err := composite.NewRedisComposite(app.cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	app.closers.Add(r, "Closing redis connections")

	// //

	return composite.NewItemService(
		item.NewItemStorage(m.
			Client().
			Database(app.cfg.MongoDB.Database).
			Collection(app.cfg.MongoDB.Collection),
		),
		//
		cache.NewCacheStorage(r.Client(), app.cfg.Cache.CacheExpire),
	), nil
}
