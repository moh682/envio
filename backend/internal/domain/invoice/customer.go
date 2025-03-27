package invoice

import (
	"github.com/google/uuid"
)

type Car struct {
	Registration string
}

type Customer struct {
	ID      uuid.UUID
	Name    string
	Email   string
	Phone   string
	Car     *Car
	Address string
	Zip     string
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
