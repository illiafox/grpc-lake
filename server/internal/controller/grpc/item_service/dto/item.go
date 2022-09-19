package dto

import (
	pb_model "github.com/illiafox/grpc-lake/gen/go/api/item_service/model/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"server/internal/domain/entity"
)

type Item entity.Item

func (i Item) ToProto() *pb_model.Item {
	return &pb_model.Item{
		Name:        i.Name,
		Data:        i.Data,
		Created:     timestamppb.New(i.Created),
		Description: i.Description,
	}
}

func ProtoToItem(proto *pb_model.Item) Item {
	return Item{
		Name:        proto.Name,
		Data:        proto.Data,
		Created:     proto.Created.AsTime(),
		Description: proto.Description,
	}
}
