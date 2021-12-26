package config

import (
	"errors"
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/spf13/pflag"
)

type Config struct {
	NomadAddress string   `hcl:"nomad_address"`
	Source       []Source `hcl:"source,block"`
	DataDir      string   `hcl:"data_dir"`
	Ui           bool     `hcl:"ui"`
	Port         int      `hcl:"port"`
}

type Source struct {
	Name   string `hcl:"name"`
	Branch string `hcl:"branch"`
	Folder string `hcl:"folder"`
	URL    string `hcl:"url"`
}

var cfg Config

func init() {
	pflag.Int("port", 8080, "api server port")
	pflag.String("conf", ".", "config file path")
	pflag.Parse()

	v, _ := pflag.CommandLine.GetString("conf")
	log.Printf("reading config from: %#v", v)

	if err := hclsimple.DecodeFile(v, nil, &cfg); err != nil {
		log.Fatalf("error reading config: %v", err)
	}
}

func GetConfig() Config {
	return cfg
}

func GetRepositoryConfig(name string) (*Source, error) {
	for _, v := range cfg.Source {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, errors.New("could not find source config")
}
