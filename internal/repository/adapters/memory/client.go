package memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
	"io"
	"log"
	"os"
)

var (
	ErrorDuplicatedRecord  = errors.New("record already exists on DB")
	ErrorNonExistentRecord = errors.New("record doesn't exist on DB")
	ErrorNotCoincidences   = errors.New("no coincidences found for the desired priceGt target")
)

type Data struct {
	db        map[int]*models.Product
	CodeIndex map[string]int
	LastID    int
}

func (d Data) GetProducts() ([]*models.ProductResponse, error) {

	var productsResponses []*models.ProductResponse

	for _, product := range d.db {
		productsResponses = append(productsResponses, &models.ProductResponse{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		})
	}
	return productsResponses, nil
}

func (d Data) GetProductById(id int) (*models.ProductResponse, error) {

	product, exists := d.db[id]
	if !exists {
		return nil, ErrorNonExistentRecord
	}

	productResponse := &models.ProductResponse{
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}

	return productResponse, nil

}

func (d Data) GetProductsByPrice(priceGt float64) ([]*models.ProductResponse, error) {

	var productsResponse []*models.ProductResponse

	for _, product := range d.db {
		if product.Price > priceGt {
			productsResponse = append(productsResponse, &models.ProductResponse{
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration,
				Price:       product.Price,
			})
		}
	}

	if len(productsResponse) == 0 {
		return nil, ErrorNotCoincidences
	}

	return productsResponse, nil
}

func LoadProducts(filePath string, data *Data) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	var products []models.Product
	if err := json.Unmarshal(fileData, &products); err != nil {
		return fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	data.LastID = len(products)
	for _, product := range products {
		data.db[product.ID] = &product
		data.CodeIndex[product.CodeValue] = product.ID
	}

	log.Printf("Value for lastID: %d", data.LastID)

	return nil
}

func NewDatabase() repository.DataInterface {
	filePath := os.Getenv("FILEPATH")
	log.Println("Loading In-Memory DB from:", filePath)
	db := make(map[int]*models.Product)
	codeIndex := make(map[string]int)
	data := &Data{
		db:        db,
		CodeIndex: codeIndex,
	}

	if err := LoadProducts(filePath, data); err != nil {
		log.Fatalf("Error loading products: %v", err)
	}

	log.Printf("Total products in database: %d", len(data.db))
	return data
}
