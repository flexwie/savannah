package config

import (
	"errors"
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	NomadAddress string   `hcl:"nomad_address"`
	Destinations []Source `hcl:"source,block"`
}

type Source struct {
	Name   string `hcl:"name"`
	Branch string `hcl:"branch"`
	Folder string `hcl:"folder"`
}

var cfg Config

func init() {
	if err := hclsimple.DecodeFile("config.hcl", nil, &cfg); err != nil {
		panic(err)
	}
	log.Printf("config is %#v", cfg)

}

func GetConfig() Config {
	return cfg
}

func GetRepositoryConfig(name string) (*Source, error) {
	for _, v := range cfg.Destinations {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, errors.New("could not find source config")
}
