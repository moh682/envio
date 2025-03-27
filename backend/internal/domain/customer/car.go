package customer

import "github.com/google/uuid"

type Car struct {
	id uuid.UUID
}

func NewCar() Car {
	return Car{
		id: uuid.New(),
	}
}

func (c *Car) ID() uuid.UUID {
	return c.id
}
