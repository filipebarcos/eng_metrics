package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stoplightio/eng_metrics/models"
)

// CreateIssue will insert GH repo data into DB
func CreateIssue(client *pgxpool.Pool, issue *models.Issue) (pgconn.CommandTag, error) {
	return client.Exec(context.Background(), `
		insert into issues (repository_id, number, url, title, milestone,
							labels_count, labels, assignees_count, assignees,
							github_created_at, github_updated_at, github_closed_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		on conflict (repository_id, number)
		do update set url = $3, title = $4, milestone = $5, labels_count = $6, labels = $7,
						assignees_count = $8, assignees = $9, github_updated_at = $11,
						github_closed_at = $12, updated_at = now();
	`, issue.RepoID, issue.Number, issue.URL, issue.Title, issue.Milestone,
		issue.LabelsCount, issue.Labels, issue.AssigneesCount, issue.Assignees,
		issue.GithubCreatedAt, issue.GithubUpdatedAt, issue.GithubClosedAt)
}
