package models

import "encoding/json"

type RawProduct struct {
	ID          int     `json:"ID"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Product struct {
	ID          int            `json:"ID" validate:"required"`
	Name        string         `json:"name"`
	Quantity    int            `json:"quantity,omitempty"`
	CodeValue   string         `json:"code_value"`
	IsPublished bool           `json:"is_published"`
	Expiration  ExpirationDate `json:"expiration"`
	Price       float64        `json:"price"`
}

type ProductResponse struct {
	Name        string         `json:"name" validate:"required"`
	Quantity    int            `json:"quantity,omitempty" validate:"required"`
	CodeValue   string         `json:"code_value" validate:"required"`
	IsPublished bool           `json:"is_published" validate:"required"`
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
