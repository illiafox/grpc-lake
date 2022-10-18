package cache

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"server/internal/domain/entity"
	itemUsecase "server/internal/domain/usecase/item"
	"server/internal/metrics"
	"server/pkg/errors"
)

var _ itemUsecase.ItemService = (*CacheService)(nil)

type CacheService struct {
	item  ItemStorage
	cache CacheStorage
}

func NewCacheService(storage ItemStorage, cache CacheStorage) CacheService {
	return CacheService{
		item:  storage,
		cache: cache,
	}
}

func (c CacheService) GetItem(ctx context.Context, id string) (entity.Item, error) {
	metrics.IncCacheTotalRequests()

	span := sentry.StartSpan(ctx, "Cache.GetItem")
	item, err := c.cache.GetItem(span.Context(), id)
	span.Finish()

	if err != nil {

		// If itemMock not found in cache, try to get it from original storage
		if err == entity.ErrItemNotFound {

			// Call original storage
			span = sentry.StartSpan(ctx, "MainStorage.GetItem")
			item, err = c.item.GetItem(span.Context(), id)
			span.Finish()
			if err != nil {
				return entity.Item{}, err
			}

			// Update cache
			span = sentry.StartSpan(ctx, "Cache.SetItem")
			err = c.cache.SetItem(span.Context(), id, item)
			span.Finish()
			if err != nil {
				return entity.Item{}, fmt.Errorf("cache.SetItem: %w", err)
			}

			return item, nil
		}

		// Capture
		return entity.Item{}, fmt.Errorf("cache.GetItem: %w", err)
	}

	// If found -> Increment cache hit counter
	metrics.IncCacheTotalHits()
	return item, nil
}

func (c CacheService) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

	// Call original storage
	span := sentry.StartSpan(ctx, "MainStorage.CreateItem")
	id, err := c.item.CreateItem(span.Context(), name, data, description)
	span.Finish()
	if err != nil {
		return "", err
	}

	// Invalidate (Delete) cache
	span = sentry.StartSpan(ctx, "Cache.DeleteItem")
	err = c.cache.DeleteItem(span.Context(), id)
	span.Finish()
	if err != nil {
		return "", errors.NewInternal("cache.DeleteItem", err)
	}

	return id, nil
}

func (c CacheService) UpdateItem(ctx context.Context, id string, item entity.Item) (updated bool, err error) {

	// Invalidate (Delete) cache
	span := sentry.StartSpan(ctx, "Cache.DeleteItem")
	err = c.cache.DeleteItem(span.Context(), id)
	span.Finish()
	if err != nil {
		return false, fmt.Errorf("cache.DeleteItem: %w", err)
	}

	// Call original storage
	span = sentry.StartSpan(ctx, "MainStorage.UpdateItem")
	updated, err = c.item.UpdateItem(span.Context(), id, item)
	span.Finish()
	if err != nil {
		return false, err
	}

	return updated, nil
}
func (c CacheService) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {

	// Invalidate (Delete) cache
	span := sentry.StartSpan(ctx, "Cache.DeleteItem")
	err = c.cache.DeleteItem(span.Context(), id)
	span.Finish()
	if err != nil {
		return false, fmt.Errorf("cache.DeleteItem: %w", err)
	}

	// Call original storage
	span = sentry.StartSpan(ctx, "MainStorage.DeleteItem")
	deleted, err = c.item.DeleteItem(span.Context(), id)
	span.Finish()
	if err != nil {
		return false, err
	}

	return deleted, nil
}
