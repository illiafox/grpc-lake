package cache

import (
	"context"
	"server/app/internal/domain/entity"
)

type CacheStorage interface {
	GetItem(ctx context.Context, id string) (entity.Item, error)
	SetItem(ctx context.Context, id string, item entity.Item) error
	DeleteItem(ctx context.Context, id string) error
}
