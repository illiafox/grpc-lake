package item

import (
	"context"
	"server/app/internal/adapters/api"
	"server/app/internal/domain/entity"
	event2 "server/app/internal/domain/service/event"
)

type itemService struct {
	storage ItemStorage
	event   event2.EventStorage
}

func NewItemService(storage ItemStorage, event event2.EventStorage) api.ItemService {
	return itemService{
		storage: storage,
		event:   event,
	}
}

func (s itemService) GetItem(ctx context.Context, id string) (entity.Item, error) {
	// TODO: add get event
	return s.storage.GetItem(ctx, id)
}

func (s itemService) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {
	id, err := s.storage.CreateItem(ctx, name, data, description)
	if err != nil {
		return "", err
	}

	err = s.event.(ctx, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s itemService) UpdateItem(ctx context.Context, id string, item entity.Item) (created bool, err error) {
	created, err = s.storage.UpdateItem(ctx, id, item)
	if err != nil {
		return false, err
	}

	if created {
		err = s.event.SendCreateItemEvent(ctx, id)
		if err != nil {
			return false, err
		}
	} else {
		err = s.event.SendUpdateItemEvent(ctx, id)
		if err != nil {
			return false, err
		}
	}

	return created, nil
}

func (s itemService) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {
	deleted, err = s.storage.DeleteItem(ctx, id)
	if err != nil {
		return false, err
	}

	if deleted {
		err = s.event.SendDeleteItemEvent(ctx, id)
		if err != nil {
			return false, err
		}
	}

	return deleted, nil
}
