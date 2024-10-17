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
	ErrorDuplicatedRecord = errors.New("record already exists on DB")
)

type Data struct {
	db        map[int]*models.Product
	CodeIndex map[string]int
	LastID    int
}

func (d Data) GetProducts() ([]*models.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
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
