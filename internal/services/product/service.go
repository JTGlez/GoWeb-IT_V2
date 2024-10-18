package product

import (
	"errors"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
)

var (
	ErrorInvalidPrice   = errors.New("price must be greater than 0")
	ErrorNoCoincidences = errors.New("no coincidences found for the desired priceGt target")
)

type serviceProduct struct {
	repo repository.DataInterface
}

func NewServiceProduct(repo repository.DataInterface) ServiceProductInterface {
	return &serviceProduct{
		repo: repo,
	}
}

type ServiceProductInterface interface {
	GetProducts() ([]*models.ProductResponse, error)
	GetProduct(id uint64) (*models.ProductResponse, error)
	GetProductByCodeValue(codeValue string) (*models.ProductResponse, error)
	GetProductsByPrice(priceGt float64) ([]*models.ProductResponse, error)
	CreateProduct(product *models.ProductResponse) (*models.ProductResponse, error)
	PutProduct(product *models.ProductResponse) (*models.ProductResponse, error)
}

func (s serviceProduct) GetProducts() ([]*models.ProductResponse, error) {
	return s.repo.GetProducts()
}

func (s serviceProduct) GetProduct(id uint64) (*models.ProductResponse, error) {
	return s.repo.GetProduct(id)
}

func (s serviceProduct) GetProductByCodeValue(codeValue string) (*models.ProductResponse, error) {
	product, err := s.repo.GetProductByCodeValue(codeValue)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s serviceProduct) GetProductsByPrice(priceGt float64) ([]*models.ProductResponse, error) {

	if priceGt <= 0 {
		return nil, ErrorInvalidPrice
	}

	products, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	var filteredProducts []*models.ProductResponse
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}

	if len(filteredProducts) == 0 {
		return nil, ErrorNoCoincidences
	}

	return filteredProducts, nil
}

func (s serviceProduct) CreateProduct(product *models.ProductResponse) (*models.ProductResponse, error) {
	newProduct, err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (s serviceProduct) PutProduct(product *models.ProductResponse) (*models.ProductResponse, error) {
	updatedProduct, err := s.repo.PutProduct(product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}
