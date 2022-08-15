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

func (s itemServer) GetItem(ctx context.Context, request *pb.GetItemRequest) (*pb.GetItemResponse, error) {

	item, err := s.item.GetItem(ctx, request.Id)

	if err != nil {
		if err == entity.ErrItemNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("item with id '%s' not found", request.Id))
		}

		if internal, ok := errors.Convert(err); ok {
			middleware.MustGetLogger(ctx).Error("item service: GetItem", log.Error(internal))
			return nil, status.Error(codes.Internal, fmt.Sprintf("item service: GetItem: %s", internal))
		}

		return nil, status.Error(codes.Aborted, fmt.Sprintf("item service: GetItem: %s", err))
	}

	return &pb.GetItemResponse{
		Item: dto.Item(item).ToProto(),
	}, nil

}
