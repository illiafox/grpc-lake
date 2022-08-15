package cache

import (
	"context"
	"server/app/internal/adapters/db/redis/cache/encode"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
)

func (c cacheStorage) SetItem(ctx context.Context, id string, item entity.Item) error {

	e := encode.Item(item)

	data, err := e.MarshalMsg(nil)
	if err != nil {
		return errors.NewInternal("encode.MarshalMsg", err)
	}

	err = c.client.Set(ctx, id, data, c.expire).Err()
	if err != nil {
		return errors.NewInternal("redis.Set", err)
	}

	return nil
}
