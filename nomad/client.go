package nomad

import (
	"log"

	"felixwie.com/savannah/config"
	"github.com/hashicorp/nomad/api"
)

var client *api.Client

func init() {
	cfg := config.GetConfig()
	addr := cfg.NomadAddress

	var err error
	client, err = api.NewClient(&api.Config{
		Address: addr,
	})

	if err != nil {
		log.Fatalf("could not create nomad client: %v", err)
	}
}

func GetClient() *api.Client {
	return client
}
