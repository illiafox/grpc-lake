package errors

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWrap(t *testing.T) {
	const scope = "test"
	err := errors.New("error")

	internal := NewInternal(err, scope)
	internal = internal.(InternalError).Wrap(scope)

	require.ErrorIs(t, err, errors.Unwrap(internal))
}
