package item_service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/internal/controller/grpc/interceptor/middleware"
	app_errors "server/pkg/errors"
)

func (s ItemServer) DeleteItem(ctx context.Context, request *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {

	deleted, err := s.item.DeleteItem(ctx, request.Id)
	if err != nil {
		internal := new(app_errors.InternalError)
		if ok := errors.As(err, internal); ok {
			middleware.MustGetLogger(ctx).Error("DeleteItem", zap.Error(internal))
			return nil, status.Error(codes.Internal, fmt.Sprintf("DeleteItem: %s", internal))
		}

		return nil, status.Error(codes.Aborted, fmt.Sprintf("DeleteItem: %s", err))
	}

	return &pb.DeleteItemResponse{
		Deleted: deleted,
	}, nil

}
