package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgresPool(connStr string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
		panic(err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Optional: konfigurasi pooling
	// config.MaxConns = 25 // maksimal koneksi
	// config.MinConns = 5  // minimal koneksi
	// config.MaxConnIdleTime = 5 * time.Minute
	// config.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		panic(err)
	}

	// Test koneksi
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
		panic(err)
	}

	log.Println("Connected to PostgreSQL with connection pooling ✅")
	return pool
}
