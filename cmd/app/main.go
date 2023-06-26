package main

import (
	"github.com/22Fariz22/analytics/config"
	"github.com/22Fariz22/analytics/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app := app.NewApp(cfg)
	app.Run()
}
