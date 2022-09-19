package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

const scope = "test"

func TestUnwrap(t *testing.T) {
	err := errors.New("error")

	internal := NewInternal(scope, err)

	require.ErrorIs(t, err, errors.Unwrap(internal))
}

func TestCause(t *testing.T) {
	err := errors.New("error")

	internal := NewInternal(scope, err)

	require.ErrorIs(t, err, internal.(InternalError).Cause())
}
