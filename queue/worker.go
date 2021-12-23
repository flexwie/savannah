package queue

import (
	"fmt"

	"github.com/go-git/go-git/storage/memory"
	"github.com/go-git/go-git/v5"
)

type WebhookJob struct {
	ID         string
	Repository string
	Branch     string
	Folder     string
}

func (t *WebhookJob) Process() {
	fmt.Printf("Processing incoming webhook: %s\n", t.ID)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: t.Repository,
	})

	if err != nil {
		return
	}

}
