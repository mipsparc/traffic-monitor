package main

import (
	"log"
	"traffic-monitor/internal/app"
	"traffic-monitor/internal/config"
	"traffic-monitor/internal/repository"
)

func main() {
	cfg, err := config.New()
	if err != nil || cfg == nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := server.Run(cfg); err != nil {
		log.Fatalf("error running application: %v", err)
	}
}
