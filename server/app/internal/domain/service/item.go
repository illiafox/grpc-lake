package service

import "server/app/internal/adapters/api"

func NewItemService(storage ItemStorage) api.ItemService {
	// TODO: kafka logs
	return storage
}
