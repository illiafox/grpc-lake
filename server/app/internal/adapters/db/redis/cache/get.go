package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"server/app/internal/adapters/db/redis/cache/encode"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
)

func (c cacheStorage) GetItem(ctx context.Context, id string) (entity.Item, error) {

	data, err := c.client.Get(ctx, id).Bytes()
	if err != nil {
		if err == redis.Nil {
			return entity.Item{}, entity.ErrItemNotFound
		}

		return entity.Item{}, errors.NewInternal("redis.Get", err)
	}

	var item encode.Item
	_, err = item.UnmarshalMsg(data)
	if err != nil {
		return entity.Item{}, errors.NewInternal("encode.UnmarshalMsg", err)
	}

	return item.ToEntity(), nil
}
