package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"server/internal/domain/entity"
)

func TestNewMessage(t *testing.T) {
	actions := []entity.Action{entity.CreateEvent, entity.DeleteEvent, entity.UpdateEvent}

	for _, action := range actions {
		t.Run(string(action), func(t *testing.T) {
			data, err := NewMessage("id", action)
			require.NoError(t, err, "encode message")

			var msg entity.Message

			err = json.Unmarshal(data, &msg)
			require.NoError(t, err, "decode message")

			require.Equal(t, "id", msg.ItemID)
			require.Equal(t, string(action), msg.Action)
			require.NotZero(t, msg.Timestamp)
		})
	}
}
