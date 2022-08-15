package item

import (
	"go.mongodb.org/mongo-driver/mongo"
	"server/app/internal/domain/service"
)

type itemStorage struct {
	collection *mongo.Collection
}

func NewItemStorage(collection *mongo.Collection) service.ItemStorage {
	return &itemStorage{collection: collection}
}
