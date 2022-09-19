package item_service

import (
	"context"
	"fmt"

	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/internal/controller/grpc/interceptor/middleware"
	"server/pkg/errors"
)

func (s itemServer) DeleteItem(ctx context.Context, request *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {

	deleted, err := s.item.DeleteItem(ctx, request.Id)
	if err != nil {
		if internal, ok := errors.Convert(err); ok {
			middleware.MustGetLogger(ctx).Error("item service: DeleteItem", zap.Error(internal))
			return nil, status.Error(codes.Internal, fmt.Sprintf("item service: DeleteItem: %s", internal))
		}
		return nil, status.Error(codes.Aborted, fmt.Sprintf("item service: DeleteItem: %s", err))
	}

	return &pb.DeleteItemResponse{
		Deleted: deleted,
	}, nil

}
