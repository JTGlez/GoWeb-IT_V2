package main

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/ping"
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/product"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository/adapters"
	serviceProduct "github.com/JTGlez/GoWeb-IT_V2/internal/services/product"
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
		log.Fatalf("could not load repository from the adapter: %s", err.Error())
	}

	// Controller, services and repos
	pongHandler := ping.NewHandler()
	productService := serviceProduct.NewServiceProduct(db)
	productController := product.NewController(productService)

	// Routes
	rt.Route("/ping", func(r chi.Router) {
		r.Get("/", pongHandler.GetPong)
	})
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", productController.GetProducts)
		r.Get("/{id}", productController.GetProductById)
		r.Get("/search", productController.GetProductsByPrice)
		r.Post("/", productController.CreateProduct)
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
