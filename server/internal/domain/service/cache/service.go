package cache

import (
	"context"
	"fmt"

	"server/internal/domain/entity"
	"server/internal/domain/usecase/item"
	"server/internal/metrics"
	"server/pkg/errors"
)

var _ item.ItemService = (*CacheService)(nil)

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
	item, err := c.cache.GetItem(ctx, id)
	if err != nil {

		// If itemMock not found in cacheMock, try to get it from original storage
		if err == entity.ErrItemNotFound {

			// Call original storage
			item, err = c.item.GetItem(ctx, id)
			if err != nil {
				return entity.Item{}, err
			}

			// Update cacheMock
			err = c.cache.SetItem(ctx, id, item)
			if err != nil {
				return entity.Item{}, fmt.Errorf("cacheMock.SetItem: %w", err)
			}

			return item, nil
		}

		// Internal
		return entity.Item{}, fmt.Errorf("cacheMock.GetItem: %w", err)
	}

	// If found -> Increment cacheMock hit counter
	metrics.IncCacheTotalHits()
	return item, nil
}

func (c CacheService) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

	// Call original storage
	id, err := c.item.CreateItem(ctx, name, data, description)
	if err != nil {
		return "", err
	}

	// Invalidate (Delete) cacheMock
	err = c.cache.DeleteItem(ctx, id)
	if err != nil {
		return "", errors.NewInternal("cacheMock.DeleteItem", err)
	}

	return id, nil
}

func (c CacheService) UpdateItem(ctx context.Context, id string, item entity.Item) (updated bool, err error) {
	// Invalidate (Delete) cacheMock
	err = c.cache.DeleteItem(ctx, id)
	if err != nil {
		return false, fmt.Errorf("cacheMock.DeleteItem: %w", err)
	}

	// Call original storage
	updated, err = c.item.UpdateItem(ctx, id, item)
	if err != nil {
		return false, err
	}

	return updated, nil
}
func (c CacheService) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {

	// Call original storage
	deleted, err = c.item.DeleteItem(ctx, id)
	if err != nil {
		return false, err
	}

	// Invalidate (Delete) cacheMock
	err = c.cache.DeleteItem(ctx, id)
	if err != nil {
		return false, fmt.Errorf("cacheMock.DeleteItem: %w", err)
	}

	return deleted, nil
}
