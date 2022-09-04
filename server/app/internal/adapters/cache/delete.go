package cache

import (
	"context"

	"server/app/pkg/errors"
)

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
