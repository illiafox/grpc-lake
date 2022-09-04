package zap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLogger(t *testing.T) {
	logger := NewLogger(os.Stdout)
	require.NotNil(t, logger)
}
