package usecase

import (
	"context"
	"fmt"
	"time"

	"server/app/internal/adapters/api"
	"server/app/internal/domain/entity"
)

type itemUsecase struct {
	item  ItemService
	event EventService
}

func NewItemUsecase(item ItemService, event EventService) api.ItemUsecase {
	return itemUsecase{
		item:  item,
		event: event,
	}
}

func (s itemUsecase) GetItem(ctx context.Context, id string) (entity.Item, error) {
	return s.item.GetItem(ctx, id)
}

func (s itemUsecase) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {
	// TODO: REMOVE BEFORE COMMMIT!!!!!!!!!!!!
	t := time.Now()
	id, err := s.item.CreateItem(ctx, name, data, description)
	fmt.Println("CreateItem", time.Since(t))
	if err != nil {
		return "", err
	}

	t = time.Now()
	err = s.event.SendItemEvent(ctx, id, entity.CreateEvent)
	fmt.Println("SendItemEvent", time.Since(t))
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s itemUsecase) UpdateItem(ctx context.Context, id string, item entity.Item) (created bool, err error) {
	created, err = s.item.UpdateItem(ctx, id, item)
	if err != nil {
		return false, err
	}

	if created {
		err = s.event.SendItemEvent(ctx, id, entity.CreateEvent)
		if err != nil {
			return false, err
		}
	} else {
		err = s.event.SendItemEvent(ctx, id, entity.UpdateEvent)
		if err != nil {
			return false, err
		}
	}

	return created, nil
}

func (s itemUsecase) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {
	deleted, err = s.item.DeleteItem(ctx, id)
	if err != nil {
		return false, err
	}

	if deleted {
		err = s.event.SendItemEvent(ctx, id, entity.DeleteEvent)
		if err != nil {
			return false, err
		}
	}

	return deleted, nil
}
