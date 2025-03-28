package invoice

import (
	"github.com/google/uuid"
)

type Car struct {
	Registration string
}

// TODO: This customer should in reality be a reference to a customer in a customer service
// but the customer service is not implemented yet, so we will just use this struct for now
// Customer should have its own domain and service just like the invoice domain
type Customer struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Car     *Car      `json:"car"`
	Address string    `json:"address"`
	Zip     string    `json:"zip"`
}

func NewCustomerReference(id uuid.UUID, name, address, city, zip string, car *Car) Customer {
	return Customer{
		ID:      id,
		Car:     car,
		Address: address,
		Zip:     zip,
		Name:    name,
	}
}
