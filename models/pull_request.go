package models

import (
	"time"

	github "github.com/stoplightio/eng_metrics/github"
)

// Issue model
type PullRequest struct {
	RepoID            int
	ID                int
	Title             string
	Number            int
	Milestone         *string
	Labels            []string
	LabelsCount       int
	URL               string
	Draft             bool
	MergedBy          *string
	ReviewDecision    string
	State             string
	Author            string
	FirstCommitAt     *time.Time
	GithubCreatedAt   time.Time
	GithubUpdatedAt   time.Time
	GithubClosedAt    *time.Time
	GithubMergedAt    *time.Time
	GithubPublishedAt *time.Time
}

// BuildIssue builds a model from GH issue
func BuildPullRequest(pullRequest github.PullRequest, repoID int) *PullRequest {
	labelsCount := pullRequest.Labels.TotalCount

	labels := make([]string, labelsCount)

	for i, label := range pullRequest.Labels.Nodes {
		labels[i] = label.Name
	}
	var commitDate *time.Time

	if len(pullRequest.Commits.Nodes) > 0 {
		commitDate = pullRequest.Commits.Nodes[0].Commit.CommittedDate
	}

	return &PullRequest{
		RepoID:            repoID,
		Title:             pullRequest.Title,
		Number:            pullRequest.Number,
		Milestone:         pullRequest.Milestone.Title,
		LabelsCount:       labelsCount,
		Labels:            labels,
		URL:               pullRequest.URL,
		Draft:             pullRequest.IsDraft,
		MergedBy:          pullRequest.MergedBy.Login,
		Author:            pullRequest.Author.Login,
		ReviewDecision:    pullRequest.ReviewDecision,
		State:             pullRequest.State,
		FirstCommitAt:     commitDate,
		GithubCreatedAt:   pullRequest.CreatedAt,
		GithubUpdatedAt:   pullRequest.UpdatedAt,
		GithubClosedAt:    pullRequest.ClosedAt,
		GithubMergedAt:    pullRequest.MergedAt,
		GithubPublishedAt: pullRequest.PublishedAt,
	}
}
