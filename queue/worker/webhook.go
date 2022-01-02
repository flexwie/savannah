package worker

import (
	"log"
	"os"
)

type WebhookJob struct {
	ID         string
	Repository string
	Branch     string
	Folder     string
}

func (t *WebhookJob) Process() {
	log.Printf("Processing incoming request: %s\n", t.ID)

	getContent(t.Repository, t.ID)
	walkDir(t.ID + "/" + t.Folder)

	defer os.RemoveAll(t.ID)
}
