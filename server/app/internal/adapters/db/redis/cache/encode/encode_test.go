package encode

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMsgPack(t *testing.T) {
	const length = 100

	item := Item{
		Name:        string(make([]byte, length)),
		Data:        make([]byte, 2*length),
		Created:     time.Unix(time.Now().Unix(), 0),
		Description: string(make([]byte, length)),
	}

	bts, err := item.MarshalMsg(nil)
	if err != nil {
		require.NoError(t, err, "marshall")
	}

	var msg Item
	left, err := msg.UnmarshalMsg(bts)
	if err != nil {
		require.NoError(t, err, "unmarshall")
	}
	require.Zero(t, len(left), "remaining bytes")

	require.True(t, reflect.DeepEqual(item, msg), "compare original and unmarshalled items")
}
