package encode

import (
	"encoding/json"

	"server/internal/domain/entity"
)

type Message struct {
	ItemID    string        `json:"item_id"`
	Action    entity.Action `json:"action"`
	Timestamp int64         `json:"timestamp"`
}

func MessageJSON(msg entity.Message) (data []byte, err error) {
	return json.Marshal(Message(msg))
}
