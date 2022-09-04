package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	const scope = "test"
	err := errors.New("error")

	internal := NewInternal(scope, err)
	internal = internal.(InternalError).Wrap(scope)

	require.ErrorIs(t, err, errors.Unwrap(internal))
}
