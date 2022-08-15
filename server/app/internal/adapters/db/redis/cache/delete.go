package cache

import (
	"context"
	"server/app/pkg/errors"
)

func (c cacheStorage) DeleteItem(ctx context.Context, id string) error {

	err := c.client.Del(ctx, id).Err()
	if err != nil {
		return errors.NewInternal("redis.Del", err)
	}

	return nil
}
