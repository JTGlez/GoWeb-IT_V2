package product

import (
	"errors"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
	"reflect"
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
	PatchProduct(product *models.ProductResponse) (*models.ProductResponse, error)
	DeleteProduct(codeValue string) error
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

func (s serviceProduct) PatchProduct(product *models.ProductResponse) (*models.ProductResponse, error) {
	existingProduct, id, err := s.repo.GetFullProductByCodeValue(product.CodeValue)
	if err != nil {
		return nil, err
	}

	if err := mergeStructs(product, existingProduct); err != nil {
		return nil, err
	}

	if product.NewCodeValue != "" && product.NewCodeValue != existingProduct.CodeValue {
		existingProduct.CodeValue = product.NewCodeValue
	}

	updatedProduct, err := s.repo.PatchProduct(existingProduct, id)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (s serviceProduct) DeleteProduct(codeValue string) error {
	return s.repo.DeleteProduct(codeValue)
}

// mergeStructs actualiza los valores de dst con los de src si no son valores por defecto.
func mergeStructs(src *models.ProductResponse, dst *models.ProductResponse) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.Field(i)

		// Si el valor del campo de src no es el valor por defecto, actualizamos el valor en dst
		if !isZeroValue(srcField) {
			dstField.Set(srcField)
		}
	}
	return nil
}

// isZeroValue determina si un campo tiene el valor por defecto.
func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
