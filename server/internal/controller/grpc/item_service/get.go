package item_service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/internal/controller/grpc/interceptor/logger"
	"server/internal/controller/grpc/item_service/dto"
	"server/internal/domain/entity"
	app_errors "server/pkg/errors"
)

func (s ItemServer) GetItem(ctx context.Context, request *pb.GetItemRequest) (*pb.GetItemResponse, error) {

	item, err := s.item.GetItem(ctx, request.Id)

	if err != nil {
		if err == entity.ErrItemNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		internal := new(app_errors.InternalError)
		if ok := errors.As(err, internal); ok {
			logger.MustGetLogger(ctx).Error("GetItem",
				zap.Error(internal),
				zap.String("line", internal.Line),
			)
			return nil, status.Error(codes.Internal, fmt.Sprintf("GetItem: %s", internal))
		}

		return nil, status.Error(codes.Aborted, fmt.Sprintf("GetItem: %s", err))
	}

	return &pb.GetItemResponse{
		Item: dto.Item(item).ToProto(),
	}, nil

}
