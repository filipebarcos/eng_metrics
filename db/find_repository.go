package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stoplightio/eng_metrics/models"
)

// FindRepositoryBySlug will find GH repo data from DB based on repo slug
func FindRepositoryBySlug(client *pgxpool.Pool, slug string) (*models.Repository, error) {
	rows, err := client.Query(context.Background(), `select id, name, is_archived, url, owner from repositories where name = $1 limit 1;`, slug)

	if err != nil {
		return nil, err
	}

	var repo models.Repository

	for rows.Next() {
		err = rows.Scan(&repo.ID, &repo.Name, &repo.IsArchived, &repo.URL, &repo.Owner)
	}

	return &repo, err
}
