package cache

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v9"
	"server/internal/adapters/db/redis/cache/encode"
	"server/internal/domain/entity"
	"server/internal/domain/service/cache"
	"server/pkg/errors"
)

var _ cache.CacheStorage = (*CacheStorage)(nil)

type CacheStorage struct {
	client redis.UniversalClient
	expire time.Duration
}

func NewCacheStorage(client redis.UniversalClient, expire time.Duration) CacheStorage {
	return CacheStorage{
		client: client,
		expire: expire,
	}
}

func (c CacheStorage) GetItem(ctx context.Context, id string) (entity.Item, error) {

	// Get message
	span := sentry.StartSpan(ctx, "Redis.Get")
	data, err := c.client.Get(span.Context(), id).Bytes()
	span.Finish()
	if err != nil {
		if err == redis.Nil {
			return entity.Item{}, entity.ErrItemNotFound
		}

		return entity.Item{}, errors.NewInternal("redis.Get", err)
	}

	// Decode message
	span = sentry.StartSpan(ctx, "UnmarshalMsg")
	defer span.Finish()

	var item encode.Item
	_, err = item.UnmarshalMsg(data)
	if err != nil {
		return entity.Item{}, errors.NewInternal("encode.UnmarshalMsg", err)
	}

	return item.ToEntity(), nil
}

func (c CacheStorage) SetItem(ctx context.Context, id string, item entity.Item) error {

	// Encode item
	span := sentry.StartSpan(ctx, "Encode Item")
	e := encode.Item(item)
	data, err := e.MarshalMsg(nil)
	span.Finish()
	if err != nil {
		return errors.NewInternal("encode.MarshalMsg", err)
	}

	// Set item to cache
	span = sentry.StartSpan(ctx, "Redis.Set")
	defer span.Finish()

	err = c.client.Set(span.Context(), id, data, c.expire).Err()
	if err != nil {
		return errors.NewInternal("redis.Set", err)
	}

	return nil
}

func (c CacheStorage) DeleteItem(ctx context.Context, id string) error {

	span := sentry.StartSpan(ctx, "Redis.Del")
	defer span.Finish()

	err := c.client.Del(span.Context(), id).Err()
	if err != nil {
		return errors.NewInternal("redis.Del", err)
	}

	return nil
}
