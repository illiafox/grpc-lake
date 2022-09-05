package item

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/app/internal/adapters/db/mongodb/item/model"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
)

func (i itemStorage) UpdateItem(ctx context.Context, id string, item entity.Item) (created bool, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("parse id: %w", err)
	}

	opts := options.Update().SetUpsert(true)

	result, err := i.collection.UpdateByID(ctx, objectID, bson.D{
		{Key: "$set", Value: model.EntityToItem(item)},
	}, opts)

	if err != nil {
		return false, errors.NewInternal("collection.UpdateByID", err)
	}

	return result.UpsertedCount == 1, nil
}
