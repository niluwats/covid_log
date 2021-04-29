package main

import (
	"github.com/niluwats/covid_log/app"
	"github.com/niluwats/covid_log/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()
}
