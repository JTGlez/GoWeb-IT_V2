package models

import "encoding/json"

type Product struct {
	ID          int     `json:"ID" validate:"required"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductResponse struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p Product) String() string {
	productBytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(productBytes)
}
