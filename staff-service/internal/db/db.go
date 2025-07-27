package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/staff-service/config"
)

func NewConnPool(cfg *config.Config) *pgxpool.Pool {
	dsn := cfg.DB_URL

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("failed to parse DB config: %v", err)
	}

	config.MaxConns = 30
	config.MinConns = 5
	config.MaxConnIdleTime = 10 * time.Minute
	config.MaxConnLifetime = 1 * time.Hour

	config.HealthCheckPeriod = 2 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	log.Println("✅ PostgreSQL connection pool initialized")
	return pool
}
