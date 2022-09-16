package rabbitmq

import (
	"context"
	"server/app/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
	"server/app/internal/domain/service/event"
	"server/app/pkg/errors"
	"server/app/pkg/log"
)

type brokerStorage struct {
	channel  *amqp.Channel
	exchange string
	key      string

	// amqp.Transient (1) or amqp.Persistent (2)
	deliveryMode uint8
}

func NewEventStorage(channel *amqp.Channel, cfg config.RabbitMQ) event.MessageStorage {
	storage := brokerStorage{
		channel:  channel,
		exchange: cfg.Exchange.Name,
		key:      cfg.Key,
		//
		deliveryMode: amqp.Transient,
	}

	if cfg.PersistentDeliveryMode {
		storage.deliveryMode = amqp.Persistent
	}

	return storage
}

func (b brokerStorage) SendMessageJSON(ctx context.Context, data []byte) error {
	err := b.channel.PublishWithContext(ctx,
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

func (b brokerStorage) HandleReturns(logger log.Logger) {
	ch := make(chan string)
	b.channel.NotifyCancel(ch)

	for r := range ch {
		logger.Error("RabbitMQ: cancelled",
			log.String("reason", r),
		)
	}
}
