package item

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/internal/adapters/db/mongodb/item/model"
	"server/internal/domain/entity"
	"server/pkg/errors"
)

type ItemStorage struct {
	collection *mongo.Collection
}

func NewItemStorage(collection *mongo.Collection) ItemStorage {
	return ItemStorage{collection: collection}
}

func (i ItemStorage) GetItem(ctx context.Context, id string) (entity.Item, error) {
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

func (i ItemStorage) UpdateItem(ctx context.Context, id string, item entity.Item) (created bool, err error) {
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

func (i ItemStorage) DeleteItem(ctx context.Context, id string) (deleted bool, err error) {
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

func (i ItemStorage) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {

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
