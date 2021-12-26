package queue

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"felixwie.com/savannah/nomad"
	"gopkg.in/src-d/go-git.v4"
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

func getContent(url string, id string) {
	log.Println("creating temp directory")
	err := os.Mkdir(id, 0755)
	checkError(err)

	log.Printf("cloning repository %s", url)
	_, err = git.PlainClone(id, false, &git.CloneOptions{
		URL: url,
	})
	checkError(err)
}

func walkDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	checkError(err)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".nomad") {
			log.Printf("processing file: %s", f.Name())

			decFile, err := os.ReadFile(path.Join(dir, f.Name()))
			checkError(err)

			dispatchFromHCL(string(decFile))
		}
	}
}

func dispatchFromHCL(data string) {
	client := nomad.GetClient()

	job, err := client.Jobs().ParseHCL(data, true)
	checkError(err)

	response, _, err := client.Jobs().Register(job, nil)
	checkError(err)

	log.Printf("adding job %v", response)
}

func checkError(err error) {
	if err != nil {
		log.Printf("error: %v", err)
	}
}
