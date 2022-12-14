package logger

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"server/pkg/log"
)

func TestLogger(t *testing.T) {
	logger := log.NewLogger(io.Discard)

	//

	t.Run("WithLogger", func(t *testing.T) {
		ctx := WithLogger(context.Background(), logger)
		require.NotNil(t, ctx)
	})

	t.Run("GetLogger", func(t *testing.T) {
		ctx := WithLogger(context.Background(), logger)
		logger = GetLogger(ctx)
		require.NotNil(t, logger)
	})

	t.Run("MustGetLogger", func(t *testing.T) {

		t.Run("NoError", func(t *testing.T) {
			ctx := WithLogger(context.Background(), logger)
			logger = MustGetLogger(ctx)
			require.NotNil(t, logger)
		})

		t.Run("Error", func(t *testing.T) {
			defer func() {
				require.Equal(t, ErrNoLogger, recover())
			}()

			logger = MustGetLogger(context.Background())
			require.NotNil(t, logger)
		})

	})

}
