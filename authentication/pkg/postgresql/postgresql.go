package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DSN struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func DSNFromEnv() DSN {
	return DSN{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func (dsn DSN) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dsn.Host, dsn.Port, dsn.User, dsn.Password, dsn.DBName,
	)
}

func New(ctx context.Context, dsn DSN) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn.String())
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
