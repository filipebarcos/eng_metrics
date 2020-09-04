package db

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
)

// UpdatePUpdatePullRequestsDeployedAt receives the repoID to whom the PRs belong
// and the actual PR numbers that were deployed
func UpdatePullRequestsDeployedAt(client *pgxpool.Pool, repoID int, numbers []int, deployTime time.Time) error {
	_, err := client.Exec(context.Background(), `
		update pull_requests set deployed_at = $3
		where repository_id = $1 and deployed_at is null and github_merged_at is not null and number = any ($2);
	`, repoID, numbers, deployTime)

	return err
}
