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
	Name    string   `hcl:"name" json:"name"`
	Branch  string   `hcl:"branch" json:"branch"`
	Folder  string   `hcl:"folder" json:"folder"`
	URL     string   `hcl:"url" json:"url"`
	Polling *Polling `hcl:"polling,block" json:"polling"`
}

type Polling struct {
	Interval int `hcl:"interval" json:"interval"`
}

var cfg Config

func init() {
	pflag.Int("port", 8080, "api server port")
	pflag.String("conf", "./config.hcl", "config file path")
	pflag.Parse()

	v, _ := pflag.CommandLine.GetString("conf")
	log.Printf("reading config from: %#v", v)

	if err := hclsimple.DecodeFile(v, nil, &cfg); err != nil {
		log.Printf("error reading config: %v", err)
		cfg.DataDir = "/data"
		cfg.NomadAddress = "http://localhost:4646"
		cfg.Port = 8080
		cfg.Ui = true
		cfg.Source = []Source{}
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
