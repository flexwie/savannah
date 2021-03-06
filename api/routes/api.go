package routes

import (
	"encoding/json"
	"net/http"

	"felixwie.com/savannah/config"
	q "felixwie.com/savannah/queue"
	"felixwie.com/savannah/queue/worker"
	"github.com/google/uuid"
)

type SyncPayload struct {
	DryRun bool   `json:"dry_run"`
	Name   string `json:"name"`
}

var Cfg config.Config

func init() {
	Cfg = config.GetConfig()
}

func SyncRepository(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data SyncPayload
	if err := decoder.Decode(&data); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("could not decode data"))
		return
	}

	// get the repo from somewhere
	sourceConfig, err := config.GetRepositoryConfig(data.Name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	queue := q.GetQueue()
	queue.Submit(&worker.WebhookJob{
		ID:         uuid.New().String(),
		Repository: sourceConfig.URL,
		Branch:     sourceConfig.Branch,
		Folder:     sourceConfig.Folder,
	})

	w.WriteHeader(http.StatusOK)
}
