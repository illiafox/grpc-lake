package composite

import (
	"github.com/segmentio/kafka-go"
	"server/app/internal/config"
	client "server/app/pkg/client/kafka"
)

var _ = Composite[*kafka.Writer](kafkaComposite{})

type kafkaComposite struct {
	writer *kafka.Writer
}

func (k kafkaComposite) Close() error {
	return k.writer.Close()
}

func (k kafkaComposite) Client() *kafka.Writer {
	return k.writer
}

func NewKafkaComposite(cfg config.RabbitMQ) (Composite[*kafka.Writer], error) {
	writer, err := client.NewWriter(client.Config(cfg))
	if err != nil {
		return nil, err
	}

	return kafkaComposite{writer}, nil
}
