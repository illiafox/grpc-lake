package event

import (
	"context"

	"server/internal/domain/entity"
	"server/internal/domain/service/event/model"
	"server/pkg/errors"
)

type EventService struct {
	sender MessageStorage
}

func NewEventService(sender MessageStorage) EventService {
	return EventService{
		sender: sender,
	}
}

func (e EventService) SendItemEvent(ctx context.Context, id string, action entity.Action) error {
	data, err := model.NewMessage(id, action)
	if err != nil {
		return errors.NewInternal("encode message", err)
	}

	err = e.sender.SendMessageJSON(ctx, data)
	if err != nil {
		return errors.NewInternal("send message", err)
	}

	return nil
}
