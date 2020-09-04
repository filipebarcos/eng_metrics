package github

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	pemPath = "certs/github.engmetrics.private-key.pem"
)

// Config struct has AccessToken, Organization and Repositories
type Config struct {
	PEMPath        string
	AppID          int64
	InstallationID int64
	Organization   string
	Repositories   []string
}

// GetConfig returns a new GithubConfig struct based of env vars
func GetConfig() *Config {
	return &Config{
		PEMPath:        filepath.Join(os.Getenv("CODEBASE_ROOT"), pemPath),
		AppID:          getIDFromEnv("GITHUB_APP_ID"),
		InstallationID: getIDFromEnv("GITHUB_INSTALLATION_ID"),
		Organization:   os.Getenv("GITHUB_ORGANIZATION"),
		Repositories:   getReposFromEnv("GITHUB_REPOSITORIES", ","),
	}
}

func getReposFromEnv(envvar string, split string) []string {
	value := os.Getenv(envvar)

	if value == "" {
		return []string{value}
	}

	return strings.Split(value, split)
}

func getIDFromEnv(envvar string) int64 {
	res, err := strconv.ParseInt(os.Getenv(envvar), 10, 64)

	if err != nil {
		return int64(0)
	}

	return res
}
