package invoice

import (
	"time"
)

type ProductStatus string

type Product struct {
	ID     int32  `json:"id"`
	Serial string `json:"serial"`
	// Deprecated: not implemented
	UpdatedAt time.Time `json:"updatedAt"`
	// Deprecated: not implemented
	CreatedAt time.Time `json:"createdAt"`
	// Deprecated: not implemented
	DeletedAt time.Time `json:"deletedAt"`

	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Rate        float64 `json:"rate"`
	Total       float64 `json:"total"`
}

func NewProduct(name, description, serial string, quantity float64, total float64, isRefunded bool) *Product {
	return &Product{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Time{},
		Description: description,
		Serial:      serial,
		Quantity:    quantity,
		Total:       total,
	}
}
