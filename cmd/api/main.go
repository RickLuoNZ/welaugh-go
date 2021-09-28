package main

import (
	"github.com/rickluonz/pawsitive/cmd/api/application"
	"github.com/rickluonz/pawsitive/pkg/exithandler"
	"github.com/rickluonz/pawsitive/pkg/logger"
)

func main() {

	// application initialization
	app, err := application.New()
	if err != nil {
		logger.Error.Fatal(err.Error())
	}

	// start the application
	go func() {
		if err := app.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	// handle application exit
	exithandler.Init(func() {
		app.Stop()
	})
}
