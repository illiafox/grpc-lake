package event

import (
	"context"

	"go.uber.org/zap"
	"server/internal/domain/entity"
)

//go:generate mockgen -source=storage.go -destination=mocks/storage.go -package=mocks

type MessageStorage interface {
	// SendMessageJSON sends asynchronously message to the storage.
	// Use HandleReturns to handle errors.
	SendMessageJSON(ctx context.Context, message entity.Message) error

	// HandleReturns handles the return exceptions from storage.
	// Must be called in a separate goroutine.
	HandleReturns(logger *zap.Logger)
}
