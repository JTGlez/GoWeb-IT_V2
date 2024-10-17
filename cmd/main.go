package main

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/ping"
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/product"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository/adapters"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Router init
	rt := chi.NewRouter()
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// DB Init
	db, err := adapters.NewRepository()
	if err != nil {
		log.Fatalf("could not load repository from the adapter")
	}

	// Service handlers
	svcPong := ping.NewHandler()
	svcProduct := product.NewHandler(db)

	// Routes
	rt.Route("/ping", func(r chi.Router) {
		r.Get("/", svcPong.GetPong)
	})
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", svcProduct.GetProducts)
	})

	// Server configs
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		log.Printf("HOST env not set, using default: %s", port)
	}

	// Startup
	log.Printf("Starting application on %s", port)
	if err := http.ListenAndServe(port, rt); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
