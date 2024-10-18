package main

import (
	"github.com/JTGlez/GoWeb-IT_V2/server"
	"log"
)

func main() {

	config := server.LoadConfig()

	svr, err := server.NewServer(
		server.WithPort(config.Port),
	)
	if err != nil {
		log.Fatalf("Error initializing server: %s", err.Error())
	}

	if err := svr.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
