package entity

import (
	"time"
)

type Action string

const (
	CreateEvent = Action("create")
	DeleteEvent = Action("delete")
	UpdateEvent = Action("update")
)

type Message struct {
	ItemID string `json:"item_id"`
	Action Action `json:"action"`
	// Timestamp is a Unix timestamp in UTC.
	Timestamp int64 `json:"timestamp"`
}

func NewMessage(id string, action Action) Message {
	return Message{
		ItemID:    id,
		Action:    action,
		Timestamp: time.Now().UTC().Unix(),
	}
}
