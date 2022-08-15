package zap

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCreateLogger(t *testing.T) {
	logger := NewLogger(os.Stdout)
	require.NotNil(t, logger)
}
