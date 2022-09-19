package item

import (
	"context"

	entity2 "server/internal/domain/entity"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks

type ItemService interface {
	// CreateItem creates new item and returns its ID.
	CreateItem(ctx context.Context, name string, data []byte, description string) (string, error)

	// GetItem retrieves item by ID and returns entity.ErrItemNotFound if not found.
	GetItem(ctx context.Context, id string) (entity2.Item, error)

	// UpdateItem updates item or creates new if not found.
	UpdateItem(ctx context.Context, id string, item entity2.Item) (updated bool, err error)

	// DeleteItem deletes item and don't returns error if not found.
	DeleteItem(ctx context.Context, id string) (deleted bool, err error)
}

type EventService interface {
	SendItemEvent(ctx context.Context, id string, action entity2.Action) error
}
