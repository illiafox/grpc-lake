package item

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/app/internal/adapters/db/mongodb/item/model"
	"server/app/pkg/errors"
)

func (i itemStorage) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

	res, err := i.collection.InsertOne(ctx, model.Item{
		Name: name,

		Data: primitive.Binary{
			Data: data,
		},

		Created:     primitive.NewDateTimeFromTime(time.Now()),
		Description: description,
	})

	if err != nil {
		return "", errors.NewInternal("collection.InsertOne", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
