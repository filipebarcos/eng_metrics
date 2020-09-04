package eng_metrics

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
	db "github.com/stoplightio/eng_metrics/db"
	gh "github.com/stoplightio/eng_metrics/github"
	models "github.com/stoplightio/eng_metrics/models"
	shared "github.com/stoplightio/eng_metrics/shared"
)

func FetchIssues(ctx context.Context, _message shared.PubSubMessage) error {
	ghConfig, client := gh.NewGithubClient()
	dbClient, err := db.New()

	if err != nil {
		log.Printf("Well, this is no good, we could not connect to the DB: %s", err)
		return err
	}

	log.Println("DB connection established")

	for _, repo := range ghConfig.Repositories {
		query, variables := gh.IssuesQuery(ghConfig.Organization, repo)
		repository, err := db.FindRepositoryBySlug(dbClient, repo)

		if err != nil {
			log.Printf("We could not find the repo in our DB %s: %s. We'll continue with the next repo.\n", repo, err)
			continue
		}

		if repository.IsArchived {
			log.Printf("The repo was archived in GH, we're not going to fetch its issues")
			continue
		}

		log.Printf("Repo (%s) found and not archived", repository.Name)

		for {
			err = client.Query(ctx, &query, variables)
			if err != nil {
				log.Printf("We found an error querying GH for %s: %s\n", repo, err)
				// If there's a problem with GH, we'll wait for next time the function runs
				break
			}

			for _, issue := range query.Repository.Issues.Nodes {
				_, err = db.CreateIssue(dbClient, models.BuildIssue(issue, repository.ID))
				if err != nil {
					log.Printf("We found an error upserting issues for repo(%d): %s\n", repository.ID, err)
					// If there's a problem  upserting the DB, we'll try again for next batch
					continue
				}
			}

			if !query.Repository.Issues.PageInfo.HasNextPage {
				break
			}

			variables["cursor"] = githubv4.String(query.Repository.Issues.PageInfo.EndCursor)
		}
	}

	log.Println("Execution completed.")

	return nil
}
