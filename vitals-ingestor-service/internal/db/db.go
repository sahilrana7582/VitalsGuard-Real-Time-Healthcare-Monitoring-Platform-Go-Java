package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/config"
)

func NewDB(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connConfig, err := pgxpool.ParseConfig(cfg.DB_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DB_URL: %w", err)
	}

	connConfig.MaxConns = 25
	connConfig.MinConns = 5
	connConfig.MaxConnLifetime = 5 * time.Minute
	connConfig.MaxConnIdleTime = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database with pgx: %w", err)
	}

	log.Println("Successfully connected to the database using pgx!")
	return pool, nil
}
