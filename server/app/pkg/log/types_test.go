package log

import (
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestFuncs(t *testing.T) {
	const key = "key"

	t.Run("Bool", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, zap.Bool(key, false), Bool(key, false))
		require.Equal(t, zap.Bool(key, true), Bool(key, true))
	})

	t.Run("Int64", func(t *testing.T) {
		t.Parallel()

		const value = int64(1234567890)
		require.Equal(t, zap.Int64(key, value), Int64(key, value))
	})

	t.Run("Any", func(t *testing.T) {
		t.Parallel()

		const value = "any"
		require.Equal(t, zap.Any(key, value), Any(key, value))
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		var value = errors.New("test error")
		require.Equal(t, zap.Error(value), Error(value))
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()

		const value = "test"
		require.Equal(t, zap.String(key, value), String(key, value))
	})
}
