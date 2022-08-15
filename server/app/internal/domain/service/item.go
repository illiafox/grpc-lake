package service

import "server/app/internal/adapters/api"

func NewItemService(storage ItemStorage) api.ItemService {
	return storage
}
