package interceptor

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"io"
	"server/app/internal/controller/grpc/interceptor/middleware"
	"server/app/pkg/log"
	"testing"
)

func TestLoggerInterceptor(t *testing.T) {
	const method = "/test"

	var buf bytes.Buffer
	logger, err := log.New(io.Discard, &buf)

	require.NoError(t, err, "create logger")

	interceptor := NewLoggerInterceptor(logger)
	require.NotNil(t, interceptor)

	// fake handler
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		l := middleware.MustGetLogger(ctx)

		l.Info("test")
		require.NoError(t, l.Sync())

		return nil, nil
	}

	_, err = interceptor(
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
		Method string `json:"method"`
	}{}

	require.NoError(t, json.NewDecoder(&buf).Decode(&expected)) // decode json logs
	require.Equal(t, method, expected.Method)
}
