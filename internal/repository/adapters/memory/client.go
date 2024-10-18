package memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
)

var (
	ErrorRecordExists      = errors.New("record already exists on DB")
	ErrorNonExistentRecord = errors.New("record doesn't exist on DB")
)

type Data struct {
	db        map[uint64]*models.Product
	CodeIndex map[string]uint64
	LastID    uint64
}

// getNextId is used to generate an incremental sequence of ids; just like OracleDB.
func (d *Data) getNextId() uint64 {
	d.LastID++
	return d.LastID
}

// GetProducts returns every product inside the DB.
func (d *Data) GetProducts() ([]*models.ProductResponse, error) {

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

// GetProduct returns a product based on the ID provided.
func (d *Data) GetProduct(id uint64) (*models.ProductResponse, error) {

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

func (d *Data) GetProductByCodeValue(codeValue string) (*models.ProductResponse, error) {

	index, indexExists := d.CodeIndex[codeValue]
	if !indexExists {
		return nil, ErrorNonExistentRecord
	}

	product, exists := d.db[index]
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

func (d *Data) CreateProduct(product *models.ProductResponse) (*models.ProductResponse, error) {

	_, exists := d.CodeIndex[product.CodeValue]
	if exists {
		return nil, ErrorRecordExists
	}

	nextId := d.getNextId()
	newProduct := &models.Product{
		ID:          nextId,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}

	d.db[nextId] = newProduct
	d.CodeIndex[newProduct.CodeValue] = newProduct.ID

	return product, nil
}

func (d *Data) PutProduct(product *models.ProductResponse) (*models.ProductResponse, error) {

	index, indexExists := d.CodeIndex[product.CodeValue]
	if !indexExists {
		return nil, ErrorNonExistentRecord
	}

	existingProduct, exists := d.db[index]
	if !exists {
		return nil, ErrorNonExistentRecord
	}

	if product.NewCodeValue != "" && product.NewCodeValue != existingProduct.CodeValue {
		delete(d.CodeIndex, existingProduct.CodeValue)
		d.CodeIndex[product.NewCodeValue] = existingProduct.ID
		existingProduct.CodeValue = product.NewCodeValue
	}

	existingProduct.Name = product.Name
	existingProduct.Quantity = product.Quantity
	existingProduct.IsPublished = product.IsPublished
	existingProduct.Expiration = product.Expiration
	existingProduct.Price = product.Price
	d.db[existingProduct.ID] = existingProduct

	productResponse := &models.ProductResponse{
		Name:        existingProduct.Name,
		Quantity:    existingProduct.Quantity,
		CodeValue:   existingProduct.CodeValue,
		IsPublished: existingProduct.IsPublished,
		Expiration:  existingProduct.Expiration,
		Price:       existingProduct.Price,
	}

	return productResponse, nil
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

	var rawProducts []*models.Product

	if err := json.Unmarshal(fileData, &rawProducts); err != nil {
		return fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	for _, rawProduct := range rawProducts {
		product := &models.Product{
			ID:          rawProduct.ID,
			Name:        rawProduct.Name,
			Quantity:    rawProduct.Quantity,
			CodeValue:   rawProduct.CodeValue,
			IsPublished: rawProduct.IsPublished,
			Expiration:  rawProduct.Expiration,
			Price:       rawProduct.Price,
		}

		data.db[product.ID] = product
		data.CodeIndex[product.CodeValue] = product.ID
		if product.ID > data.LastID {
			data.LastID = product.ID
		}
	}

	log.Printf("Value for lastID: %d", data.LastID)
	return nil
}

func NewDatabase() repository.DataInterface {
	filePath := os.Getenv("FILEPATH")
	log.Println("Loading In-Memory DB from:", filePath)
	db := make(map[uint64]*models.Product)
	codeIndex := make(map[string]uint64)
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
