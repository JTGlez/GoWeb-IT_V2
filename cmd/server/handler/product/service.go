package product

import "net/http"

type serviceProduct struct{}

type InterfaceProduct interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

func NewHandler() InterfaceProduct {
	return &serviceProduct{}
}
