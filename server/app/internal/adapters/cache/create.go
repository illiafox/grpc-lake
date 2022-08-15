package cache

import (
	"context"
	"server/app/pkg/errors"
)

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
