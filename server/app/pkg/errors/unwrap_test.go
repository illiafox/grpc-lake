package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnwrap(t *testing.T) {
	const scope = "test"
	err := errors.New("error")

	internal := NewInternal(scope, err)

	require.ErrorIs(t, err, errors.Unwrap(internal))
}

func TestCause(t *testing.T) {
	const scope = "test"
	err := errors.New("error")

	internal := NewInternal(scope, err)

	require.ErrorIs(t, err, internal.(InternalError).Cause())
}
