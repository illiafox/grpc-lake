package item_service

import (
	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	service "server/app/internal/adapters/api"
)

// Interface cast check
var _ pb.ItemServiceServer = (*itemServer)(nil)

type itemServer struct {
	pb.UnimplementedItemServiceServer
	//
	item service.ItemService
}

func NewServer(item service.ItemService) pb.ItemServiceServer {
	return itemServer{
		item: item,
	}
}
