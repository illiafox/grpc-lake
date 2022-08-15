package grpc

import (
	"fmt"
	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"google.golang.org/grpc"
	"net"
	service "server/app/internal/adapters/api"
	"server/app/internal/controller/grpc/interceptor"
	"server/app/internal/controller/grpc/item_service"
	"server/app/pkg/log"
)

type Server struct {
	srv *grpc.Server
}

func (s Server) Listen(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	return s.srv.Serve(lis)
}

func (s Server) GracefulStop() {
	s.srv.GracefulStop()
}

func NewServer(logger log.Logger, item service.ItemService) Server {

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewMetricsInterceptor(),      // out
			interceptor.NewLoggerInterceptor(logger), // in
		),
	)

	pb.RegisterItemServiceServer(server, item_service.NewServer(item))

	return Server{srv: server}
}
