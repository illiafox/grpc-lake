package event

import (
	"context"

	"server/app/internal/domain/entity"
	"server/app/pkg/log"
)

type MessageStorage interface {
	// SendMessageJSON sends asynchronously message to the storage.
	// Use HandleReturns to handle errors.
	SendMessageJSON(ctx context.Context, data []byte) error

	// HandleReturns handles the return exceptions from storage.
	// Must be called in a separate goroutine.
	HandleReturns(logger log.Logger)
}

type EventService interface {
	SendItemEvent(ctx context.Context, id string, action entity.Action) error
}
