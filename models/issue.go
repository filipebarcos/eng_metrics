package models

import (
	"time"

	github "github.com/stoplightio/eng_metrics/github"
)

// Issue model
type Issue struct {
	ID              int
	Title           string
	Number          int
	RepoID          int
	Milestone       string
	Assignees       []string
	AssigneesCount  int
	Labels          []string
	LabelsCount     int
	URL             string
	GithubCreatedAt time.Time
	GithubUpdatedAt time.Time
	GithubClosedAt  *time.Time
}

// BuildIssue builds a model from GH issue
func BuildIssue(issues github.Issue, repoID int) *Issue {
	assigneesCount := issues.Assignees.TotalCount
	labelsCount := issues.Labels.TotalCount

	assignees := make([]string, assigneesCount)
	labels := make([]string, labelsCount)

	for i, assignee := range issues.Assignees.Nodes {
		assignees[i] = assignee.Login
	}

	for i, label := range issues.Labels.Nodes {
		labels[i] = label.Name
	}

	return &Issue{
		RepoID:          repoID,
		Title:           issues.Title,
		Number:          issues.Number,
		Milestone:       issues.Milestone.Title,
		LabelsCount:     labelsCount,
		Labels:          labels,
		AssigneesCount:  assigneesCount,
		Assignees:       assignees,
		URL:             issues.URL,
		GithubCreatedAt: issues.CreatedAt,
		GithubUpdatedAt: issues.UpdatedAt,
		GithubClosedAt:  issues.ClosedAt,
	}
}
