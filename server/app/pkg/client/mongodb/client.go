package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(ctx context.Context, uri string) (*mongo.Client, error) {

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	// Ping the primary
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return client, nil
}
