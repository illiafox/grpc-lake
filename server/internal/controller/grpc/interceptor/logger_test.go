package interceptor

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"server/internal/controller/grpc/interceptor/logger"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"server/pkg/log"
)

func TestLoggerInterceptor(t *testing.T) {
	const method = "/test"

	var buf bytes.Buffer

	interceptor := NewLoggerInterceptor(
		log.NewLogger(io.Discard, &buf),
	)
	require.NotNil(t, interceptor)

	// fake handler
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		l := logger.MustGetLogger(ctx)

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
