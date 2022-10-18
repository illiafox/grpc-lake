package encode

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"server/internal/domain/entity"
)

func TestEncodeMessage(t *testing.T) {
	actions := []entity.Action{entity.CreateEvent, entity.DeleteEvent, entity.UpdateEvent}
	const id = "id"

	for _, action := range actions {
		t.Run(string(action), func(t *testing.T) {
			expected := entity.NewMessage(id, action)

			data, err := MessageJSON(expected)
			require.NoError(t, err, "encode message")

			var msg Message

			err = json.Unmarshal(data, &msg)
			require.NoError(t, err, "decode message")

			require.Equal(t, expected.ItemID, msg.ItemID)
			require.Equal(t, expected.Action, msg.Action)
			require.Equal(t, expected.Timestamp, msg.Timestamp)
		})
	}
}
