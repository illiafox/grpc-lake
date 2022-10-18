package grpc

import (
	"fmt"
	"net"

	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	service "server/internal/adapters/api"
	"server/internal/controller/grpc/interceptor"
	"server/internal/controller/grpc/item_service"
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

func NewServer(logger *zap.Logger, item service.ItemUsecase) Server {

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewMetricsInterceptor(), // out
			interceptor.NewSentryInterceptor(),
			interceptor.NewLoggerInterceptor(logger), // in
		),
	)

	pb.RegisterItemServiceServer(server, item_service.NewServer(item))

	return Server{srv: server}
}
