package cache

import (
	"context"

	"server/internal/domain/entity"
)

//go:generate mockgen -source=storage.go -destination=mocks/storage.go -package=mocks

type CacheStorage interface {
	GetItem(ctx context.Context, id string) (entity.Item, error)
	SetItem(ctx context.Context, id string, item entity.Item) error
	DeleteItem(ctx context.Context, id string) error
}

type ItemStorage interface {
	// CreateItem creates new item and returns its ID.
	CreateItem(ctx context.Context, name string, data []byte, description string) (string, error)

	// GetItem retrieves item by ID and returns entity.ErrItemNotFound if not found.
	GetItem(ctx context.Context, id string) (entity.Item, error)

	// UpdateItem updates item or creates new if not found.
	UpdateItem(ctx context.Context, id string, item entity.Item) (updated bool, err error)

	// DeleteItem deletes item and don't returns error if not found.
	DeleteItem(ctx context.Context, id string) (deleted bool, err error)
}
