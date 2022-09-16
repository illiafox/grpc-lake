package cache

import (
	"context"
	"server/app/internal/domain/entity"
	"server/app/internal/metrics"
	"server/app/pkg/errors"
)

type cacheWrapper struct {
	item  ItemStorage
	cache CacheStorage
}

func NewCacheWrapper(storage ItemStorage, cache CacheStorage) any {
	return cacheWrapper{
		item:  storage,
		cache: cache,
	}
}

func (c cacheWrapper) GetItem(ctx context.Context, id string) (entity.Item, error) {

	metrics.IncCacheTotalRequests()
	item, err := c.cache.GetItem(ctx, id)
	if err != nil {

		// If item not found in cache, try to get it from original storage
		if err == entity.ErrItemNotFound {

			// Call original storage
			item, err = c.item.GetItem(ctx, id)
			if err != nil {
				return entity.Item{}, err
			}

			// Update cache
			err = c.cache.SetItem(ctx, id, item)
			if err != nil {
				if internal, ok := errors.Convert(err); ok {
					return entity.Item{}, internal.Wrap("cache.SetItem")
				}
				return entity.Item{}, errors.NewInternal("cache.SetItem", err)
			}

			return item, nil
		}

		// Internal
		return entity.Item{}, errors.NewInternal("cache.GetItem", err)
	}

	// If found -> Increment cache hit counter
	metrics.IncCacheTotalHits()
	return item, nil
}

func (c cacheWrapper) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

	// Call original storage
	id, err := c.item.CreateItem(ctx, name, data, description)
	if err != nil {
		return "", err
	}

	// Invalidate (Delete) cache
	err = c.cache.DeleteItem(ctx, id)
	if err != nil {
		if internal, ok := errors.Convert(err); ok {
			return "", internal.Wrap("cache.DeleteItem")
		}
		return "", errors.NewInternal("cache.DeleteItem", err)
	}

	return id, nil
}

func (c cacheWrapper) UpdateItem(ctx context.Context, id string, item entity.Item) (updated bool, err error) {

	// Call original storage
	updated, err = c.item.UpdateItem(ctx, id, item)
	if err != nil {
		return false, err
	}

	// Invalidate (Delete) cache
	err = c.cache.DeleteItem(ctx, id)
	if err != nil {
		if internal, ok := errors.Convert(err); ok {
			return false, internal.Wrap("cache.DeleteItem")
		}
		return false, errors.NewInternal("cache.DeleteItem", err)
	}

	return updated, nil
}
func (c cacheWrapper) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {

	// Call original storage
	deleted, err = c.item.DeleteItem(ctx, id)
	if err != nil {
		return false, err
	}

	// Invalidate (Delete) cache
	err = c.cache.DeleteItem(ctx, id)
	if err != nil {
		if internal, ok := errors.Convert(err); ok {
			return false, internal.Wrap("cache.DeleteItem")
		}
		return false, errors.NewInternal("cache.DeleteItem", err)
	}

	return deleted, nil
}
