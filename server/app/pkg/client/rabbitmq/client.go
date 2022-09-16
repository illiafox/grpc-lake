package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewDialChannel(uri string, queue Queue, exchange Exchange) (*amqp.Channel, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		exchange.Name,       // name
		exchange.Kind,       // kind
		exchange.Durable,    // durable
		exchange.AutoDelete, // auto delete
		exchange.Internal,   // internal
		exchange.NoWait,     // no wait
		nil,                 //arguments
	)

	if err != nil {
		return nil, fmt.Errorf("declare exchange: %w", err)
	}

	_, err = ch.QueueDeclare(
		queue.Name,       // name
		queue.Durable,    // durable
		queue.AutoDelete, // delete when unused
		queue.Exclusive,  // exclusive
		queue.NoWait,     // no-wait
		nil,              // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("declare queue: %w", err)
	}

	return ch, nil
}
