package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"server/app/pkg/errors"
)

type brokerStorage struct {
	channel  *amqp.Channel
	exchange string
	key      string

	// amqp.Transient (1) or amqp.Persistent (2)
	deliveryMode uint8
}

func NewBrokerStorage(channel *amqp.Channel, exchange, key string, PersistentDeliveryMode bool) {
	storage := brokerStorage{
		channel:  channel,
		exchange: exchange,
		key:      exchange,
		//
		deliveryMode: amqp.Transient,
	}

	if PersistentDeliveryMode {
		storage.deliveryMode = amqp.Persistent
	}
}

func (b brokerStorage) SendMessageJSON(ctx context.Context, data []byte) error {
	err := b.channel.Publish(
		// exchange
		b.exchange,
		// key
		b.key,
		// mandatory - if there is no queue, we will get the message back and log the error
		true,
		// immediate - false
		false,

		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: b.deliveryMode,
		},
	)

	if err != nil {
		return errors.NewInternal("publish message", err)
	}

	return nil
}
