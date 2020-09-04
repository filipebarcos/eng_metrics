package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stoplightio/eng_metrics/models"
)

// CreatePullRequest will insert GH repo data into DB
func CreatePullRequest(client *pgxpool.Pool, pullRequest *models.PullRequest) (pgconn.CommandTag, error) {
	return client.Exec(context.Background(), `
		insert into pull_requests (repository_id, number, url, title, milestone,
							labels_count, labels, author, state, review_decision,
							draft, merged_by, github_created_at, github_updated_at,
							github_closed_at, github_published_at, github_merged_at, first_commit_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
		on conflict (repository_id, number)
		do update set url = $3, title = $4, milestone = $5, labels_count = $6, labels = $7,
						author = $8, state = $9, review_decision = $10, draft = $11, merged_by = $12,
						github_updated_at = $14, github_closed_at = $15, github_published_at = $16,
						github_merged_at = $17, updated_at = now(), first_commit_at = $18;
	`, pullRequest.RepoID, pullRequest.Number, pullRequest.URL, pullRequest.Title, pullRequest.Milestone,
		pullRequest.LabelsCount, pullRequest.Labels, pullRequest.Author, pullRequest.State,
		pullRequest.ReviewDecision, pullRequest.Draft, pullRequest.MergedBy, pullRequest.GithubCreatedAt,
		pullRequest.GithubUpdatedAt, pullRequest.GithubClosedAt, pullRequest.GithubPublishedAt,
		pullRequest.GithubMergedAt, pullRequest.FirstCommitAt)
}
