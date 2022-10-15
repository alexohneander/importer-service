package main

import (
	"alexohneander/importer-service/config"
	"log"
)

func main() {

	// Configure Settings
	app := config.ConfigureServer()

	// Configure Routes
	config.ConfigureRoutes(app)

	// Start App
	startErr := app.Listen(":" + config.PORT)
	if startErr != nil {
		log.Fatal(startErr)
	}
}
