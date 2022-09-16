package event

import (
	"context"
	"server/app/internal/domain/entity"
	"server/app/internal/domain/service/event/model"
	"server/app/pkg/errors"
)

type eventService struct {
	sender MessageStorage
}

func NewEventService(sender MessageStorage) EventService {
	return eventService{
		sender: sender,
	}
}

func (e eventService) SendItemEvent(ctx context.Context, id string, action entity.Action) error {
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
