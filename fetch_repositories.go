package eng_metrics

import (
	"context"
	"log"

	db "github.com/stoplightio/eng_metrics/db"
	gh "github.com/stoplightio/eng_metrics/github"
	models "github.com/stoplightio/eng_metrics/models"
	shared "github.com/stoplightio/eng_metrics/shared"
)

func FetchRepositories(ctx context.Context, _message shared.PubSubMessage) error {
	ghConfig, client := gh.NewGithubClient()
	dbClient, err := db.New()

	if err != nil {
		log.Printf("Well, this is no good, we could not connect to the DB: %s", err)
		return err
	}

	log.Println("DB connection established")

	for _, repo := range ghConfig.Repositories {
		query, variables := gh.RepositoryQuery(ghConfig.Organization, repo)
		log.Println("Querying repo: ", repo)

		err := client.Query(ctx, &query, variables)
		if err != nil {
			log.Printf("We found an error querying %s: %s\n", repo, err)
			continue
		}

		repoID, err := db.CreateRepository(dbClient, models.BuildRepository(query))
		if err != nil {
			log.Printf("We found an error upserting repo (%d): %s\n", repoID, err)
			continue
		}
	}

	log.Println("Execution completed.")

	return nil
}
