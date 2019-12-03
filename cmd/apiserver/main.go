package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Pupye/rest-api/internal/app/apiserver"
)

var (
	configpath string
)

func init() {
	flag.StringVar(&configpath, "config-path", "configs/apiserver.toml", "path to config")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configpath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
