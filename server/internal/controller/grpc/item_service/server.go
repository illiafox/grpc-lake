package item_service

import (
	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	service "server/internal/adapters/api"
)

// Interface cast check
var _ pb.ItemServiceServer = (*ItemServer)(nil)

type ItemServer struct {
	pb.UnimplementedItemServiceServer
	//
	item service.ItemUsecase
}

func NewServer(item service.ItemUsecase) ItemServer {
	return ItemServer{
		item: item,
	}
}
