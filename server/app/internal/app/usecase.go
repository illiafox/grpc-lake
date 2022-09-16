package app

import (
	"fmt"

	"server/app/internal/adapters/api"
	"server/app/internal/adapters/db/mongodb/item"
	"server/app/internal/adapters/db/redis/cache"
	"server/app/internal/adapters/rabbitmq"
	"server/app/internal/composite"
)

func (app *App) ItemService() (api.ItemUsecase, error) {

	m, err := composite.NewMongoComposite(app.Config.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("mongodb: %w", err)
	}

	app.closers.Add(m, "Closing mongodb connections")

	// //

	r, err := composite.NewRedisComposite(app.Config.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	app.closers.Add(r, "Closing redis connections")

	// //

	k, err := composite.NewKafkaComposite(app.Config.RabbitMQ)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: %w", err)
	}
	app.closers.Add(k, "Closing rabbitmq connections")

	// //

	return composite.NewItemUsecase(
		item.NewItemStorage(m.
			Client().
			Database(app.Config.MongoDB.Database).
			Collection(app.Config.MongoDB.Collection),
		),
		//
		cache.NewCacheStorage(r.Client(), app.Config.Cache.CacheExpire),
		//
		rabbitmq.NewEventStorage(k.Client()),
	), nil
}
