package cache

import (
	"context"
	"server/app/internal/domain/entity"
	"server/app/internal/metrics"
	"server/app/pkg/errors"
)

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
