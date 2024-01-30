package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/buts00/RestApi.git/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
}

func main() {

	flag.Parse()
	config := apiserver.NewConfig()
	toml.Decode(configPath, config)

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
