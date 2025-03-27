package invoice

import (
	"time"

	"github.com/google/uuid"
)

type ProductStatus string

type Product struct {
	ID        uuid.UUID
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time

	Name        string
	Description string
	Serial      string
	Quantity    float64
	UnitPrice   float64
	// including tax
	Total      float64
	IsRefunded bool
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
