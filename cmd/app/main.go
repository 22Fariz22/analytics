package main

import (
	"github.com/22Fariz22/analytics/internal/app"
	"github.com/22Fariz22/analytics/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg)
	app.Run()

}
