package config

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	NomadAddress string        `hcl:"nomad_address"`
	Destinations []Destination `hcl:"destination,block"`
}

type Destination struct {
	Location string `hcl:"Location"`
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
