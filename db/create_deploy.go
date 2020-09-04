package db

import (
	"context"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// CreateDeploy
func CreateDeploy(client *pgxpool.Pool, deployTime time.Time) (pgconn.CommandTag, error) {
	return client.Exec(context.Background(), `insert into deploys (created_at) values ($1);`, deployTime)
}
