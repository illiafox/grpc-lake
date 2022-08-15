package composite

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"server/app/internal/config"
	"server/app/pkg/client/mongodb"
	"time"
)

var _ = Composite[*mongo.Client](mongoComposite{})

type mongoComposite struct {
	client  *mongo.Client
	timeout time.Duration
}

func (m mongoComposite) Client() *mongo.Client {
	if m.client == nil {
		panic("client is nil")
	}

	return m.client
}

func (m mongoComposite) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	return m.client.Disconnect(ctx)
}

func NewMongoComposite(cfg config.MongoDB) (Composite[*mongo.Client], error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	client, err := mongodb.New(ctx, cfg.URI)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	return mongoComposite{
		client:  client,
		timeout: cfg.Timeout,
	}, nil
}
