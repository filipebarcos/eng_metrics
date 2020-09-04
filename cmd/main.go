package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"
	metrics "github.com/stoplightio/eng_metrics"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	funcframework.RegisterHTTPFunction("/migrateDb", metrics.MigrateDb)
	funcframework.RegisterHTTPFunction("/trackDeploys", metrics.TrackDeploys)
	funcframework.RegisterEventFunction("/fetchRepositories", metrics.FetchRepositories)
	funcframework.RegisterEventFunction("/fetchIssues", metrics.FetchIssues)
	funcframework.RegisterEventFunction("/fetchPullRequests", metrics.FetchPullRequests)

	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("framework.Start: %v\n", err)
	}
}
