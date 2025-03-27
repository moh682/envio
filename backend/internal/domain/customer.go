package domain

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID              ID        `json:"id"`
	CarRegistration string    `json:"car_registration"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	ZipCode         string    `json:"zip_code"`
	CreatedAt       time.Time `json:"created_at"`
}

func (c *Customer) Validate() error {

	// Validate ID
	_, err := uuid.Parse(c.ID.String())
	if err != nil {
		return err
	}

	return nil
}
