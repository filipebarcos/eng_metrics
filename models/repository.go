package models

import (
	"time"

	github "github.com/stoplightio/eng_metrics/github"
)

// Repository model
type Repository struct {
	ID              int
	Name            string
	IsPrivate       bool
	IsFork          bool
	IsTemplate      bool
	IsArchived      bool
	URL             string
	GithubCreatedAt time.Time
	Owner           string
}

// BuildRepository builds a mode from a GH repo
func BuildRepository(repo github.RepositoryType) *Repository {
	return &Repository{
		Name:            repo.Repository.Name,
		IsPrivate:       repo.Repository.IsPrivate,
		IsTemplate:      repo.Repository.IsTemplate,
		IsFork:          repo.Repository.IsFork,
		IsArchived:      repo.Repository.IsArchived,
		URL:             repo.Repository.URL,
		GithubCreatedAt: repo.Repository.CreatedAt,
		Owner:           repo.Repository.Owner.Login,
	}
}
