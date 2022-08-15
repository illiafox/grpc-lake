package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"server/app/internal/controller/grpc/interceptor/middleware"
	"server/app/pkg/log"
)

// NewLoggerInterceptor returns a new unary server interceptor with Logger inside Context .
// The logger can be received from middleware.GetLogger or middleware.MustGetLogger functions.
//
// Interceptor adds 'method' field with full grpc call method
func NewLoggerInterceptor(l log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (
		resp interface{},
		err error) {

		//

		ctx = middleware.WithLogger(ctx, l.With(
			log.String("method", info.FullMethod),
		))

		//

		return handler(ctx, req)
	}
}
