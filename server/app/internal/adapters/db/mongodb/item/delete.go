package item

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/app/pkg/errors"
)

func (i itemStorage) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("parse id: %w", err)
	}

	result, err := i.collection.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: objectID},
	})

	if err != nil {
		return false, errors.NewInternal("collection.DeleteOne", err)
	}

	return result.DeletedCount == 1, nil
}
