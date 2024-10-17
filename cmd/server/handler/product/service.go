package product

import (
	"errors"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
	"net/http"
)

var (
	ErrorInvalidID    = errors.New("invalid id input")
	ErrorInvalidPrice = errors.New("invalid priceGt param")
)

type serviceProduct struct {
	db repository.DataInterface
}

type InterfaceProduct interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProductsByPrice(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
}

func NewHandler(db repository.DataInterface) InterfaceProduct {
	return &serviceProduct{
		db: db,
	}
}
