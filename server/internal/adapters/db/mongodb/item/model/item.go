package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/internal/domain/entity"
)

type Item struct {
	Name        string             `bson:"name,omitempty"`
	Data        primitive.Binary   `bson:"data,omitempty"`
	Created     primitive.DateTime `bson:"created,omitempty"`
	Description string             `bson:"description,omitempty"`
}

func (i Item) ToEntity() entity.Item {
	return entity.Item{
		Name:        i.Name,
		Data:        i.Data.Data,
		Created:     i.Created.Time(),
		Description: i.Description,
	}
}

func EntityToItem(e entity.Item) Item {
	return Item{
		Name:        e.Name,
		Data:        primitive.Binary{Data: e.Data},
		Created:     primitive.NewDateTimeFromTime(e.Created),
		Description: e.Description,
	}
}
