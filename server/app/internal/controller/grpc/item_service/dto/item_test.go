package dto

import (
	pb_model "github.com/illiafox/grpc-lake/gen/go/api/item_service/model/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestItemToProto(t *testing.T) {
	item := Item{
		Name:        "name",
		Data:        []byte{1, 2, 3},
		Created:     time.Now(),
		Description: "description",
	}

	protoItem := item.ToProto()

	// //

	r := require.New(t)

	r.Equal(item.Name, protoItem.Name)
	r.Equal(item.Data, protoItem.Data)

	{
		c := protoItem.Created.AsTime()
		// Due to timestamppb.Timestamp timezone converting, time must be compared in Unix
		r.Equal(item.Created.UnixNano(), c.UnixNano())
	}

	r.Equal(item.Description, protoItem.Description)
}

func TestProtoToItem(t *testing.T) {
	protoItem := &pb_model.Item{
		Name:        "name",
		Data:        []byte{1, 2, 3},
		Created:     timestamppb.Now(),
		Description: "description",
	}

	item := ProtoToItem(protoItem)

	// //

	r := require.New(t)

	r.Equal(protoItem.Name, item.Name)
	r.Equal(protoItem.Data, item.Data)

	{
		c := protoItem.Created.AsTime()
		// Due to timestamppb.Timestamp timezone converting, time must be compared in Unix
		r.Equal(c.UnixNano(), item.Created.UnixNano())
	}

	r.Equal(protoItem.Description, item.Description)
}
