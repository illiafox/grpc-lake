package item

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/app/internal/adapters/db/mongodb/item/model"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
)

func (i itemStorage) GetItem(ctx context.Context, id string) (entity.Item, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Item{}, fmt.Errorf("parse id: %w", err)
	}

	result := i.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: objectID},
	})

	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Item{}, entity.ErrItemNotFound
		}
		return entity.Item{}, errors.NewInternal("collection.FindOne", err)
	}

	var item model.Item
	if err = result.Decode(&item); err != nil {
		return entity.Item{}, errors.NewInternal("decode item", err)
	}

	return item.ToEntity(), nil
}
