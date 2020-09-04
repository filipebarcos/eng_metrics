package eng_metrics

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	db "github.com/stoplightio/eng_metrics/db"
)

// TrackDeploys function will get any request and just create
// a single entry into `deploys` table
func TrackDeploys(w http.ResponseWriter, r *http.Request) {
	deployedAt := time.Now()
	log.Println("Request to track deploys received: ", deployedAt)

	var data struct {
		Numbers []int `json:"numbers"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Println("Could not parse the request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(data.Numbers) == 0 {
		log.Println("No PRs given")
		http.Error(w, "welp, I cant process this", http.StatusBadRequest)
		return
	} else {
		log.Println("PRs being updated are: ", data.Numbers)
	}

	dbClient, err := db.New()

	if err != nil {
		log.Printf("Well, this is no good, we could not connect to the DB: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository, err := db.FindRepositoryBySlug(dbClient, "platform-internal")

	if err != nil {
		log.Printf("We could not find the repo in our DB: %s.", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.UpdatePullRequestsDeployedAt(dbClient, repository.ID, data.Numbers, deployedAt)

	if err != nil {
		log.Printf("Well, we found an error updating the pull requests: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.CreateDeploy(dbClient, deployedAt)

	if err != nil {
		log.Printf("Well, we found an error creating a deploy into DB: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Deploy successfully tracked at: %s", deployedAt)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"ok": true}`)
}
