package app

import (
	"context"
	"fmt"

	"server/internal/adapters/api"
	"server/internal/adapters/brokers/rabbitmq"
	"server/internal/adapters/db/mongodb/item"
	"server/internal/adapters/db/redis/cache"
	"server/internal/composite"
)

func (app *App) ItemService() (api.ItemUsecase, error) {

	m, err := composite.NewMongoComposite(app.Config.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("mongodb: %w", err)
	}

	app.closers.Add(func() error {
		return m.Disconnect(context.TODO())
	}, "Closing mongodb connections")

	// //

	r, err := composite.NewRedisComposite(app.Config.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	app.closers.Add(r.Close, "Closing redis connections")

	// //

	ch, err := composite.NewRabbitmqComposite(app.Config.RabbitMQ)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: %w", err)
	}
	app.closers.Add(ch.Close, "Closing rabbitmq connections")

	// //

	return composite.NewItemUsecase(
		item.NewItemStorage(m.
			Database(app.Config.MongoDB.Database).
			Collection(app.Config.MongoDB.Collection),
		),
		//
		cache.NewCacheStorage(r, app.Config.Cache.CacheExpire),
		//
		rabbitmq.NewEventStorage(
			// channel
			ch,
			// config
			app.Config.RabbitMQ,
		),
	), nil
}
