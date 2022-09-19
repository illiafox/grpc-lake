package encode

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestItemMsgPackTags(t *testing.T) {
	var item Item

	r := reflect.TypeOf(item)
	for i := 0; i < r.NumField(); i++ {
		tag := r.Field(i).Tag.Get("msgp")

		if tag == "" {
			t.Fatalf("field %s.%s has empty `msgp` tag",
				r.Name(), r.Field(i).Name,
			)
		}

	}
}

func TestItemToEntity(t *testing.T) {

	// // Create a new Item
	e := Item{
		Name:        "name",
		Data:        []byte("data"),
		Created:     time.Unix(time.Now().Unix(), 0),
		Description: "description",
	}

	// Check if all fields are set
	checkFields(t, e)

	// // Convert to Item
	item := e.ToEntity()

	// Check if all converted fields are not empty
	checkFields(t, item)

	r := require.New(t)
	// // Manual equality check
	r.Equal(e.Name, item.Name)
	r.Equal(e.Data, item.Data)
	r.Equal(e.Created, item.Created)
	r.Equal(e.Description, item.Description)
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
