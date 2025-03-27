package customer

import (
	"github.com/google/uuid"
)

type Customer struct {
	cars    []Car
	person  Person
	address string
	zipCode string
}

func NewCustomer(id uuid.UUID, person Person, address, zipCode string) Customer {
	return Customer{
		person:  person,
		address: address,
		zipCode: zipCode,
		cars:    []Car{},
	}
}

func (c *Customer) ID() uuid.UUID {
	return c.person.ID()
}

func (c *Customer) Person() Person {
	return c.person
}

func (c *Customer) Address() string {
	return c.address
}

func (c *Customer) ZipCode() string {
	return c.zipCode
}

func (c *Customer) Cars() []valueobjects.Car {
	return c.cars
}

func (c *Customer) AddCar(car valueobjects.Car) {
	c.cars = append(c.cars, car)
}
