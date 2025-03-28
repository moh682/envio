package invoice

import (
	"time"

	"github.com/google/uuid"
)

type ProductStatus string

type Product struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Serial      string  `json:"serial"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
	// including tax
	Total      float64 `json:"total"`
	IsRefunded bool    `json:"isRefunded"`
}

func NewProduct(name, description, serial string, quantity float64, total float64, isRefunded bool) *Product {
	return &Product{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Time{},
		Name:        name,
		Description: description,
		Serial:      serial,

		UnitPrice:  total / quantity,
		Quantity:   quantity,
		Total:      total,
		IsRefunded: isRefunded,
	}
}
