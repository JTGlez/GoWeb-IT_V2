package repository

import (
	"errors"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
)

var (
	ErrorUnimplementedAdapter = errors.New("not implemented data source on environment")
	ErrorMissingAdapter       = errors.New("missing adapter option on environment")
)

type DataInterface interface {
	GetProducts() ([]*models.ProductResponse, error)
	GetProduct(id uint64) (*models.ProductResponse, error)
	GetProductByCodeValue(codeValue string) (*models.ProductResponse, error)
	CreateProduct(product *models.ProductResponse) (*models.ProductResponse, error)
	PutProduct(product *models.ProductResponse) (*models.ProductResponse, error)
}
