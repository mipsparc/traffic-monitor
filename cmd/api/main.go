package main

import (
	"log"
	"traffic-monitor/internal/app"
	"traffic-monitor/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := server.Run(cfg); err != nil {
		log.Fatalf("error running application: %v", err)
	}
}
