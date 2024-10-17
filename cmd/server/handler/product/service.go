package product

import (
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
	"net/http"
)

type serviceProduct struct {
	db repository.DataInterface
}

type InterfaceProduct interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
}

func NewHandler(db repository.DataInterface) InterfaceProduct {
	return &serviceProduct{
		db: db,
	}
}
