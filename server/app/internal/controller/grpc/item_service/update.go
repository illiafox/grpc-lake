package item_service

import (
	"context"
	"fmt"
	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/app/internal/controller/grpc/interceptor/middleware"
	"server/app/internal/controller/grpc/item_service/dto"
	"server/app/internal/domain/entity"
	"server/app/pkg/errors"
	"server/app/pkg/log"
)

func (s itemServer) UpdateItem(ctx context.Context, request *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {

	item := dto.ProtoToItem(request.Item)
	created, err := s.item.UpdateItem(ctx, request.Id, entity.Item(item))
	if err != nil {
		if internal, ok := errors.Convert(err); ok {
			middleware.MustGetLogger(ctx).Error("item service: UpdateItem", log.Error(internal))
			return nil, status.Error(codes.Internal, fmt.Sprintf("item service: UpdateItem: %s", internal))
		}
		return nil, status.Error(codes.Aborted, fmt.Sprintf("item service: UpdateItem: %s", err))
	}

	return &pb.UpdateItemResponse{
		Created: created,
	}, nil

}
