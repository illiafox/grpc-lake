package model

import (
	"encoding/json"
	"time"

	"server/internal/domain/entity"
)

// NewMessage serializes the message to JSON.
func NewMessage(id string, action entity.Action) ([]byte, error) {
	msg := entity.Message{
		ItemID:    id,
		Action:    string(action),
		Timestamp: time.Now().Unix(),
	}

	return json.Marshal(msg)
}
