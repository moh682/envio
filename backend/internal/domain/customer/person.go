package customer

import "github.com/google/uuid"

type Person struct {
	id uuid.UUID
}

func NewPerson() Person {
	return Person{
		id: uuid.New(),
	}
}

func (p *Person) ID() uuid.UUID {
	return p.id
}
