package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DSN struct {
	URI      string
	Username string
	Password string
}

func DSNFromEnv() DSN {
	return DSN{
		URI:      os.Getenv("DB_URI"),
		Password: os.Getenv("DB_PASSWORD"),
		Username: os.Getenv("DB_USER"),
	}
}

func New(ctx context.Context, dsn DSN) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI(dsn.URI).
		SetAuth(options.Credential{
			Username: dsn.Username,
			Password: dsn.Password,
		}))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return client, nil
}
