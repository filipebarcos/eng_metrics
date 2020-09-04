package db

import (
	"context"
	"os"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// New returns a Postgres connection pool
func New() (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewConn() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
