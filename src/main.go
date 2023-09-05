package main

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Port              int
	Host              string
	SaveDirectory     string
	PathFormat        string
	HostSaveDirectory bool
}

func loadConfig() Config {
	var cfg Config

	bytes, err := os.ReadFile("./config.toml")
	if err != nil {
		log.Fatal(err)
	}

	if err = toml.Unmarshal(bytes, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := loadConfig()

	log.Fatal(cfg)
}
