package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Connection *pgxpool.Pool

func NewConnection(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	Connection, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return Connection, nil
}
