package item

import (
	"context"

	entity2 "server/internal/domain/entity"
)

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

func (s ItemUsecase) GetItem(ctx context.Context, id string) (entity2.Item, error) {
	return s.item.GetItem(ctx, id)
}

func (s ItemUsecase) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

	id, err := s.item.CreateItem(ctx, name, data, description)
	if err != nil {
		return "", err
	}

	err = s.event.SendItemEvent(ctx, id, entity2.CreateEvent)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s ItemUsecase) UpdateItem(ctx context.Context, id string, item entity2.Item) (created bool, err error) {
	created, err = s.item.UpdateItem(ctx, id, item)
	if err != nil {
		return false, err
	}

	if created {
		err = s.event.SendItemEvent(ctx, id, entity2.CreateEvent)
		if err != nil {
			return false, err
		}
	} else {
		err = s.event.SendItemEvent(ctx, id, entity2.UpdateEvent)
		if err != nil {
			return false, err
		}
	}

	return created, nil
}

func (s ItemUsecase) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {
	deleted, err = s.item.DeleteItem(ctx, id)
	if err != nil {
		return false, err
	}

	if deleted {
		err = s.event.SendItemEvent(ctx, id, entity2.DeleteEvent)
		if err != nil {
			return false, err
		}
	}

	return deleted, nil
}
