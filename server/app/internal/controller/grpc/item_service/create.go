package item_service

import (
	"context"
	"fmt"

	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/app/internal/controller/grpc/interceptor/middleware"
	"server/app/pkg/errors"
	"server/app/pkg/log"
)

func (s itemServer) CreateItem(ctx context.Context, request *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {

	id, err := s.item.CreateItem(ctx, request.Name, request.Data, request.Description)
	if err != nil {
		if internal, ok := errors.Convert(err); ok {
			middleware.MustGetLogger(ctx).Error("item service: CreateItem", log.Error(internal))
			return nil, status.Error(codes.Internal, fmt.Sprintf("item service: CreateItem: %s", internal))
		}
		return nil, status.Error(codes.Aborted, fmt.Sprintf("item service: CreateItem: %s", err))
	}

	return &pb.CreateItemResponse{
		Id: id,
	}, nil

}