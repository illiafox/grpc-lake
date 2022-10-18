package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/internal/metrics"
)

// NewMetricsInterceptor returns a new unary server interceptor, which handle request and increments metrics.
// Must be used with chaining.
func NewMetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		metrics.IncGrpcTotalRequests()

		if err != nil {
			switch status.Convert(err).Code() {

			case codes.Internal:
				metrics.IncGrpcErrorRequests()

			}
		} else {
			metrics.IncGrpcSuccessRequests()
		}

		return resp, err
	}
}
