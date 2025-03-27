package domain

import (
	"github.com/google/uuid"
)

type Cost struct {
	ID          ID      `json:"id"`
	ProductNr   string  `json:"product_nr"`
	Description string  `json:"description"`
	Quantity    float32 `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
	Total       float32 `json:"total"`
}

func (c *Cost) Validate() error {
	_, err := uuid.Parse(c.ID.String())
	if err != nil {
		return err
	}

	return nil
}
