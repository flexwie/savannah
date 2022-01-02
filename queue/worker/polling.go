package worker

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PollingJob struct {
	ID         string
	Repository string
	Branch     string
	Folder     string
	Interval   int
	LatestHash string
	LastPolled time.Time
}

func (p *PollingJob) Process() {
	id := uuid.New().String()
	p.LastPolled = time.Now()

	repo := getContent(p.Repository, id)
	revision, err := repo.ResolveRevision("HEAD")
	checkError(err)

	if p.CompareHash(revision.String()) {
		log.Println("found differences, applying them")
		walkDir(id + "/" + p.Folder)
	}

	defer os.RemoveAll(id)
}

/*
	Compares the given hash to the latest polled hash.
	Sets the polled hash as new if different.
	Returns true if hash is different.
*/
func (p *PollingJob) CompareHash(hash string) bool {
	if !strings.Contains(p.LatestHash, hash) {
		p.LatestHash = hash
		return true
	}

	return false
}
