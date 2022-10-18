package item

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"server/internal/adapters/api"
	entity "server/internal/domain/entity"
)

var _ api.ItemUsecase = (*ItemUsecase)(nil)

type ItemUsecase struct {
	item  ItemService
	event EventService
}

func NewItemUsecase(item ItemService, event EventService) ItemUsecase {
	return ItemUsecase{
		item:  item,
		event: event,
	}
}

func (s ItemUsecase) GetItem(ctx context.Context, id string) (entity.Item, error) {
	return s.item.GetItem(ctx, id)
}

func (s ItemUsecase) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

	span := sentry.StartSpan(ctx, "CreateItem")
	id, err := s.item.CreateItem(span.Context(), name, data, description)
	span.Finish()
	if err != nil {
		return "", err
	}

	//

	span = sentry.StartSpan(ctx, "SendItemEvent")
	defer span.Finish()

	err = s.event.SendItemEvent(span.Context(), id, entity.CreateEvent)
	if err != nil {
		return "", fmt.Errorf("SendItemEvent: %w", err)
	}

	return id, nil
}

func (s ItemUsecase) UpdateItem(ctx context.Context, id string, item entity.Item) (created bool, err error) {
	span := sentry.StartSpan(ctx, "UpdateItem")
	created, err = s.item.UpdateItem(span.Context(), id, item)
	span.Finish()
	if err != nil {
		return false, err
	}

	//

	span = sentry.StartSpan(ctx, "SendItemEvent")
	defer span.Finish()

	if created {
		err = s.event.SendItemEvent(span.Context(), id, entity.CreateEvent)
		if err != nil {
			return false, fmt.Errorf("SendItemEvent: %w", err)
		}
	} else {
		err = s.event.SendItemEvent(span.Context(), id, entity.UpdateEvent)
		if err != nil {
			return false, fmt.Errorf("SendItemEvent: %w", err)
		}
	}

	return created, nil
}

func (s ItemUsecase) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {

	span := sentry.StartSpan(ctx, "DeleteItem")
	deleted, err = s.item.DeleteItem(span.Context(), id)
	span.Finish()
	if err != nil {
		return false, err
	}

	//

	if deleted {
		span = sentry.StartSpan(ctx, "SendItemEvent")
		defer span.Finish()

		err = s.event.SendItemEvent(span.Context(), id, entity.DeleteEvent)
		if err != nil {
			return false, fmt.Errorf("SendItemEvent: %w", err)
		}
	}

	return deleted, nil
}
