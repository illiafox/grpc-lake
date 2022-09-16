package api

import (
	"context"

	"server/app/internal/domain/entity"
)

type ItemUsecase interface {
	// CreateItem creates new item and returns its ID.
	CreateItem(ctx context.Context, name string, data []byte, description string) (string, error)

	// GetItem retrieves item by ID and returns entity.ErrItemNotFound if not found.
	GetItem(ctx context.Context, id string) (entity.Item, error)

	// UpdateItem updates item or creates new if not found.
	UpdateItem(ctx context.Context, id string, item entity.Item) (created bool, err error)

	// DeleteItem deletes item and don't returns error if not found.
	DeleteItem(ctx context.Context, id string) (deleted bool, err error)
}
