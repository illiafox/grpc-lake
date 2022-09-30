package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"server/internal/adapters/brokers/rabbitmq/encode"
	"server/internal/config"
	"server/internal/domain/entity"
	"server/internal/domain/service/event"
	app_errors "server/pkg/errors"
)

var _ event.MessageStorage = (*BrokerStorage)(nil)

type BrokerStorage struct {
	channel  *amqp.Channel
	exchange string
	key      string

	// amqp.Transient (1) or amqp.Persistent (2)
	deliveryMode uint8
}

func NewEventStorage(channel *amqp.Channel, cfg config.RabbitMQ) BrokerStorage {
	storage := BrokerStorage{
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

func (b BrokerStorage) SendMessageJSON(ctx context.Context, msg entity.Message) error {
	data, err := encode.MessageJSON(msg)
	if err != nil {
		return app_errors.NewInternal("encode message", err)
	}

	err = b.channel.PublishWithContext(ctx,
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
		return app_errors.NewInternal("publish message", err)
	}

	return nil
}

func (b BrokerStorage) HandleReturns(logger *zap.Logger) {

	cancel := make(chan string)
	b.channel.NotifyCancel(cancel)

	ret := make(chan amqp.Return)
	b.channel.NotifyReturn(ret)

	cl := make(chan *amqp.Error)
	b.channel.NotifyClose(cl)

	for {
		select {
		case c := <-cancel: // cancel
			logger.Error("RabbitMQ: cancelled",
				zap.String("reason", c),
			)
		case c := <-cl: // close
			logger.Error("RabbitMQ: closed",
				zap.String("reason", c.Reason),
				zap.Int("code", c.Code),
			)
		case r := <-ret: // return
			logger.Error("RabbitMQ: returned",
				zap.String("exchange", r.Exchange),
				zap.Uint16("replyCode", r.ReplyCode),
			)
		}
	}
}
