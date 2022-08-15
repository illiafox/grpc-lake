package errors

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnwrap(t *testing.T) {
	const scope = "test"
	err := errors.New("error")

	internal := NewInternal(err, scope)

	require.ErrorIs(t, err, errors.Unwrap(internal))
}

func TestCause(t *testing.T) {
	const scope = "test"
	err := errors.New("error")

	internal := NewInternal(err, scope)

	require.ErrorIs(t, err, internal.(InternalError).Cause())
}
