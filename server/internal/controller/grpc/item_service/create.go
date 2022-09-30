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

func (s ItemServer) CreateItem(ctx context.Context, request *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {

	id, err := s.item.CreateItem(ctx, request.Name, request.Data, request.Description)
	if err != nil {
		internal := new(app_errors.InternalError)
		if ok := errors.As(err, internal); ok {
			middleware.MustGetLogger(ctx).Error("CreateItem", zap.Error(internal))
			return nil, status.Error(codes.Internal, fmt.Sprintf("CreateItem: %s", internal))
		}
		return nil, status.Error(codes.Aborted, fmt.Sprintf("CreateItem: %s", err))
	}

	return &pb.CreateItemResponse{
		Id: id,
	}, nil

}
