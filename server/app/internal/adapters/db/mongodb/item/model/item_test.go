package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/app/internal/domain/entity"
)

func TestItemHasBSONTags(t *testing.T) {
	var item Item

	r := reflect.TypeOf(item)
	for i := 0; i < r.NumField(); i++ {
		if r.Field(i).Tag.Get("bson") == "" {
			t.Fatalf("field %s.%s has empty `bson` tag",
				r.Name(), r.Field(i).Name,
			)
		}
	}
}

func TestEntityToItem(t *testing.T) {

	// // Create a new Item
	e := entity.Item{
		Name:        "name",
		Data:        []byte("data"),
		Created:     time.Unix(time.Now().Unix(), 0), // primitive.DateTime doesn't save nano seconds
		Description: "description",
	}

	// Check if all fields are set
	checkFields(t, e)

	// // Convert to Item
	item := EntityToItem(e)

	// Check if all converted fields are not empty
	checkFields(t, item)

	r := require.New(t)
	// // Manual equality check
	r.Equal(e.Name, item.Name)
	r.Equal(e.Data, item.Data.Data)
	r.Equal(e.Created, item.Created.Time())
	r.Equal(e.Description, item.Description)
}

func TestItemToEntity(t *testing.T) {
	item := Item{
		Name: "name",
		Data: primitive.Binary{
			Data: []byte("data"),
		},
		Created:     primitive.NewDateTimeFromTime(time.Now()),
		Description: "description",
	}

	// Check if all fields are set
	checkFields(t, item)

	e := item.ToEntity()

	// Check if all converted fields are not empty
	checkFields(t, e)

	r := require.New(t)
	r.Equal(item.Name, e.Name)
	r.Equal(item.Data.Data, e.Data)
	r.Equal(item.Created.Time(), e.Created)
	r.Equal(item.Description, e.Description)
}

// checkFields checks whether all fields are set
func checkFields(t *testing.T, item any) {
	v := reflect.ValueOf(item)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			t.Fatalf("field %s.%s is empty (zero)",
				v.Type().Name(), v.Type().Field(i).Name,
			)
		}
	}
}
