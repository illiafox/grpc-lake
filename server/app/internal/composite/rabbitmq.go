package composite

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"server/app/internal/config"
	client "server/app/pkg/client/rabbitmq"
)

var _ = Composite[*amqp.Channel](rabbitmqComposite{})

type rabbitmqComposite struct {
	channel *amqp.Channel
}

func (k rabbitmqComposite) Close() error {
	return k.channel.Close()
}

func (k rabbitmqComposite) Client() *amqp.Channel {
	return k.channel
}

func NewRabbitmqComposite(cfg config.RabbitMQ) (Composite[*amqp.Channel], error) {
	channel, err := client.NewDialChannel(cfg.URI, client.Queue(cfg.Queue), client.Exchange(cfg.Exchange))
	if err != nil {
		return nil, err
	}
	return rabbitmqComposite{channel}, nil
}
