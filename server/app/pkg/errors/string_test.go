package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInternalErrorToString(t *testing.T) {
	const scope = "test"
	var err = errors.New("error")

	internal := NewInternal(err, scope)
	f := fmt.Errorf("%s%s%w", scope, Separator, err)

	require.Equal(t, f.Error(), internal.Error())
}
