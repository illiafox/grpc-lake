package event

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"server/internal/domain/entity"
	"server/internal/domain/usecase/item"
)

var _ item.EventService = (*EventService)(nil)

type EventService struct {
	sender MessageStorage
}

func NewEventService(sender MessageStorage) EventService {
	return EventService{
		sender: sender,
	}
}

func (e EventService) SendItemEvent(ctx context.Context, id string, action entity.Action) error {
	msg := entity.NewMessage(id, action)

	span := sentry.StartSpan(ctx, "SendMessageJSON")
	defer span.Finish()

	err := e.sender.SendMessageJSON(span.Context(), msg)
	if err != nil {
		return fmt.Errorf("SendMessageJSON: %w", err)
	}

	return nil
}
