package interceptor

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"server/internal/controller/grpc/interceptor/middleware"
	"server/pkg/log"
)

func TestLoggerInterceptor(t *testing.T) {
	const method = "/test"

	var buf bytes.Buffer
	logger := log.NewLogger(io.Discard, &buf)

	interceptor := NewLoggerInterceptor(logger)
	require.NotNil(t, interceptor)

	// fake handler
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		l := middleware.MustGetLogger(ctx)

		l.Info("test")
		require.NoError(t, l.Sync())

		return nil, nil
	}

	_, err := interceptor(
		context.Background(),
		/* empty request */ nil,
		&grpc.UnaryServerInfo{
			Server:     nil,
			FullMethod: method,
		},
		handler,
	)

	require.NoError(t, err)

	var expected = struct {
		Method string `format:"method"`
	}{}

	require.NoError(t, json.NewDecoder(&buf).Decode(&expected)) // decode format logs
	require.Equal(t, method, expected.Method)
}
