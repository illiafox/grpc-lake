package cache

import (
	"context"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
)

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
