package event

import (
	"context"

	"go.uber.org/zap"
)

//go:generate mockgen -source=storage.go -destination=mocks/storage.go -package=mocks

type MessageStorage interface {
	// SendMessageJSON sends asynchronously message to the storage.
	// Use HandleReturns to handle errors.
	SendMessageJSON(ctx context.Context, data []byte) error

	// HandleReturns handles the return exceptions from storage.
	// Must be called in a separate goroutine.
	HandleReturns(logger *zap.Logger)
}
