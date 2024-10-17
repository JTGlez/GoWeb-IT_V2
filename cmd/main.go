package main

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/ping"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	rt := chi.NewRouter()
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	svcPong := ping.NewHandler()

	rt.Route("/ping", func(r chi.Router) {
		r.Get("/", svcPong.GetPong)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		log.Printf("HOST env not set, using default: %s", port)
	}

	log.Printf("Starting application on %s", port)
	if err := http.ListenAndServe(port, rt); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
