package cache

import (
	"context"
	"server/app/internal/adapters/db/redis/cache/encode"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
	"time"

	"github.com/go-redis/redis/v9"
	"server/app/internal/adapters/cache"
)

type cacheStorage struct {
	client redis.UniversalClient
	expire time.Duration
}

func NewCacheStorage(client redis.UniversalClient, expire time.Duration) cache.CacheStorage {
	return &cacheStorage{
		client: client,
		expire: expire,
	}
}

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

func (c cacheStorage) DeleteItem(ctx context.Context, id string) error {

	err := c.client.Del(ctx, id).Err()
	if err != nil {
		return errors.NewInternal("redis.Del", err)
	}

	return nil
}
