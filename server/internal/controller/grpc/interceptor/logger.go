package interceptor

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"server/internal/controller/grpc/interceptor/middleware"
)

// NewLoggerInterceptor returns a new unary server interceptor with Logger inside Context .
// The logger can be received from middleware.GetLogger or middleware.MustGetLogger functions.
//
// Interceptor adds 'method' field with full grpc call method
func NewLoggerInterceptor(l *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (
		resp interface{},
		err error) {

		//

		ctx = middleware.WithLogger(ctx, l.With(
			zap.String("method", info.FullMethod),
		))

		//

		return handler(ctx, req)
	}
}
