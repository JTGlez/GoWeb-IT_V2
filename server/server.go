package server

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/ping"
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler/product"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository/adapters"
	serviceProduct "github.com/JTGlez/GoWeb-IT_V2/internal/services/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Server struct {
	host   string
	port   string
	router *chi.Mux
}

func NewServer(options ...func(server *Server) error) (*Server, error) {
	s := &Server{
		host:   "localhost",
		port:   ":8080",
		router: chi.NewRouter(),
	}

	for _, option := range options {
		if err := option(s); err != nil {
			return nil, err
		}
	}

	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.setupRoutes()

	return s, nil
}

func (s *Server) setupRoutes() {
	db, err := adapters.NewRepository()
	if err != nil {
		log.Fatalf("could not load repository from the adapter: %s", err.Error())
	}

	pongHandler := ping.NewHandler()
	productService := serviceProduct.NewServiceProduct(db)
	productController := product.NewController(productService)

	s.router.Route("/ping", func(r chi.Router) {
		r.Get("/", pongHandler.GetPong)
	})
	s.router.Route("/products", func(r chi.Router) {
		r.Get("/", productController.GetProducts)
		r.Get("/by-id/{id}", productController.GetProductById)
		r.Get("/by-code/{code_value}", productController.GetProductByCodeValue)
		r.Get("/search", productController.GetProductsByPrice)
		r.Post("/", productController.CreateProduct)
		r.Put("/", productController.PutProduct)
		r.Patch("/", productController.PatchProduct)
	})
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s%s", s.host, s.port)
	return http.ListenAndServe(s.host+s.port, s.router)
}
