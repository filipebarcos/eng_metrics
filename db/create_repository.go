package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stoplightio/eng_metrics/models"
)

// CreateRepository will insert GH repo data into DB
func CreateRepository(client *pgxpool.Pool, repo *models.Repository) (int, error) {
	var id int
	err := client.QueryRow(context.Background(), `
		insert into repositories (name, owner, url, is_fork, is_private, is_archived, is_template, github_created_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (name, owner)
		do update set url = $3, is_fork = $4, is_private = $5, is_archived = $6, is_template = $7, updated_at = now()
		returning id;
	`, repo.Name, repo.Owner, repo.URL, repo.IsFork, repo.IsPrivate, repo.IsArchived, repo.IsTemplate, repo.GithubCreatedAt).Scan(&id)

	return id, err
}
