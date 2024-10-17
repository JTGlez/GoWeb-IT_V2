package repository

import (
	"errors"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
)

var (
	ErrorUnimplementedAdapter = errors.New("not implemented data source on environment")
)

type DataInterface interface {
	GetProducts() ([]*models.ProductResponse, error)
	GetProductById(id int) (*models.ProductResponse, error)
	GetProductsByPrice(priceGt float64) ([]*models.ProductResponse, error)
}
