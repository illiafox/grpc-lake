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
	"server/internal/controller/grpc/item_service/dto"
	"server/internal/domain/entity"
	app_errors "server/pkg/errors"
)

func (s ItemServer) UpdateItem(ctx context.Context, request *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {

	item := dto.ProtoToItem(request.Item)

	created, err := s.item.UpdateItem(ctx, request.Id, entity.Item(item))
	if err != nil {
		internal := new(app_errors.InternalError)
		if ok := errors.As(err, internal); ok {
			middleware.MustGetLogger(ctx).Error("UpdateItem",
				zap.Error(internal),
				zap.String("line", internal.Line),
			)
			return nil, status.Error(codes.Internal, fmt.Sprintf("UpdateItem: %s", internal))
		}

		return nil, status.Error(codes.Aborted, fmt.Sprintf("UpdateItem: %s", err))
	}

	return &pb.UpdateItemResponse{
		Created: created,
	}, nil

}
