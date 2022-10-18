package interceptor

import (
	"context"
	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewLoggerInterceptor returns a new unary server interceptor with Logger inside Context .
// The logger can be received from logger.GetLogger or logger.MustGetLogger functions.
//
// Interceptor adds 'method' field with full grpc call method
func NewSentryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		defer sentry.Recover()
		span := sentry.StartSpan(ctx, "gRPC call", sentry.TransactionName(info.FullMethod))

		resp, err = handler(span.Context(), req)

		span.Finish()

		if err != nil {
			s := status.Convert(err)

			switch s.Code() {
			case codes.Internal:

				sentry.CaptureMessage(s.Message())
			}
		}

		return resp, err
	}
}
