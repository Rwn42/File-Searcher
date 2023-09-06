package main

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port          int
	Host          string
	SaveDirectory string
}

var cfg Config

func loadConfig() Config {
	bytes, err := os.ReadFile("./config.toml")
	if err != nil {
		log.Fatal(err)
	}

	if err = toml.Unmarshal(bytes, &cfg); err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./public/"+cfg.SaveDirectory, 0755)

	return cfg
}
