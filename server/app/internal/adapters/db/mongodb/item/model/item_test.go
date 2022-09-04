package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/app/internal/domain/entity"
)

func TestEntityToItem(t *testing.T) {
	e := entity.Item{
		Name:        "name",
		Data:        []byte("data"),
		Created:     time.Unix(time.Now().Unix(), 0), // primitive.DateTime doesn't save nano seconds
		Description: "description",
	}

	i := EntityToItem(e)

	r := require.New(t)

	r.Equal(e.Name, i.Name)
	r.Equal(e.Data, i.Data.Data)
	r.Equal(e.Created, i.Created.Time())
	r.Equal(e.Description, i.Description)
}

func TestItemToEntity(t *testing.T) {
	i := Item{
		Name: "name",
		Data: primitive.Binary{
			Data: []byte("data"),
		},
		Created:     primitive.NewDateTimeFromTime(time.Now()),
		Description: "description",
	}

	e := i.ToEntity()

	r := require.New(t)
	r.Equal(i.Name, e.Name)
	r.Equal(i.Data.Data, e.Data)
	r.Equal(i.Created.Time(), e.Created)
	r.Equal(i.Description, e.Description)
}
