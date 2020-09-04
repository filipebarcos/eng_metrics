package github

import (
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/shurcooL/githubv4"
)

// NewGithubClient returns a githubv4.Client to
// query GH v4 GQL API
func NewGithubClient() (*Config, *githubv4.Client) {
	config := GetConfig()
	itr, err := ghinstallation.NewKeyFromFile(
		http.DefaultTransport,
		config.AppID,
		config.InstallationID,
		config.PEMPath,
	)

	if err != nil {
		log.Fatal(err)
	}

	return config, githubv4.NewClient(&http.Client{Transport: itr})
}
