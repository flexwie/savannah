package worker

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"felixwie.com/savannah/nomad"
	"gopkg.in/src-d/go-git.v4"
)

func getContent(url string, id string) *git.Repository {
	log.Println("creating temp directory")
	err := os.Mkdir(id, 0755)
	checkError(err)

	log.Printf("cloning repository %s", url)
	repo, err := git.PlainClone(id, false, &git.CloneOptions{
		URL: url,
	})
	checkError(err)

	return repo
}

func walkDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	checkError(err)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".nomad") {
			log.Printf("processing file: %s", f.Name())

			decFile, err := os.ReadFile(path.Join(dir, f.Name()))
			if checkError(err) {
				dispatchFromHCL(string(decFile))
			}

		}
	}
}

func dispatchFromHCL(data string) {
	client := nomad.GetClient()

	job, err := client.Jobs().ParseHCL(data, true)
	if checkError(err) {
		response, _, err := client.Jobs().Register(job, nil)
		if checkError(err) {
			log.Printf("adding job %v", response)
		}
	}
}

func checkError(err error) bool {
	if err != nil {
		log.Printf("error: %v", err)
		return false
	}

	return true
}
