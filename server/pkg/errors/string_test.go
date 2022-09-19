package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInternalErrorToString(t *testing.T) {
	const scope = "test"
	var err = errors.New("error")

	internal := NewInternal(scope, err)
	f := fmt.Errorf("%s%s%w", scope, Separator, err)

	require.Equal(t, f.Error(), internal.Error())
}
