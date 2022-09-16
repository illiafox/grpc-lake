package kafka

import (
	"context"
)

type Message struct {
	Data []byte
}

type MessageWriter interface {
	WriteMessage(ctx context.Context, msg Message) error
}
