package main

import (
	"flag"
	"os"

	"github.com/Arseniy-Polyakov/GoCourse/cmd/config"
)

func parseFlags(cfg *config.Config) {
	flag.StringVar(&cfg.ServerAddress, "a", ":8080", "SERVER_ADDRESS")
	flag.StringVar(&cfg.BaseURL, "b", ":8080", "BASE_URL")
	flag.Parse()

	if _, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		cfg.ServerAddress = os.Getenv("SERVER_ADDRESS")
	}

	if _, ok := os.LookupEnv("BASE_URL"); ok {
		cfg.BaseURL = os.Getenv("BASE_URL")
	}
}
