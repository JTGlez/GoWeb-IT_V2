package models

import "encoding/json"

// Product is a basic structure of products with ID stored (must match with the CodeIndex record of the product).
type Product struct {
	ID          uint64         `json:"ID" validate:"required"`
	Name        string         `json:"name"`
	Quantity    int            `json:"quantity,omitempty"`
	CodeValue   string         `json:"code_value"`
	IsPublished bool           `json:"is_published"`
	Expiration  ExpirationDate `json:"expiration"`
	Price       float64        `json:"price"`
}

// ProductResponse is a basic structure for product data sent to and received from the client.
type ProductResponse struct {
	Name        string         `json:"name" validate:"required"`
	Quantity    int            `json:"quantity,omitempty" validate:"required"`
	CodeValue   string         `json:"code_value" validate:"required"`
	IsPublished bool           `json:"is_published"`
	Expiration  ExpirationDate `json:"expiration" validate:"required"`
	Price       float64        `json:"price" validate:"required"`
}

func (p Product) String() string {
	productBytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(productBytes)
}
