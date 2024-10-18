package product

import (
	"errors"
	"github.com/JTGlez/GoWeb-IT_V2/internal/services/product"
	"net/http"
)

var (
	ErrorInvalidID    = errors.New("invalid id input")
	ErrorInvalidPrice = errors.New("invalid priceGt param")
)

type controllerProduct struct {
	productSvc product.ServiceProductInterface
}

type ControllerProductInterface interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProductByCodeValue(w http.ResponseWriter, r *http.Request)
	GetProductsByPrice(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	PutProduct(w http.ResponseWriter, r *http.Request)
}

func NewController(
	productSvc product.ServiceProductInterface) ControllerProductInterface {
	return &controllerProduct{
		productSvc: productSvc,
	}
}
