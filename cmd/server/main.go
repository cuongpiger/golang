package main

import (
	"log"

	config "github.com/cuongpiger/golang/config/server"
	"github.com/cuongpiger/golang/internal/server/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	log.Println("config: ", cfg)
	// Run
	app.Run(cfg)
}