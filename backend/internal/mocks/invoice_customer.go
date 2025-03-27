package mocks

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	domain "github.com/moh682/envio/backend/internal/domain"
)

type CustomerMock struct {
	customer *domain.Customer
}

func (c *CustomerMock) ToCustomer() *domain.Customer {
	return c.customer
}

func NewCustomer() *CustomerMock {
	uuid := uuid.New()
	customer := &domain.Customer{
		ID:              domain.ID(uuid),
		CarRegistration: "AB823746",
		Name:            gofakeit.Name(),
		Email:           gofakeit.Email(),
		Phone:           "9238467190",
		Address:         gofakeit.Address().Address,
		ZipCode:         gofakeit.Address().Zip,
		CreatedAt:       gofakeit.Date(),
	}

	return &CustomerMock{customer: customer}
}

func (c *CustomerMock) WithName(name string) *CustomerMock {
	c.customer.Name = name
	return c
}

func (c *CustomerMock) WithEmail(email string) *CustomerMock {
	c.customer.Email = email
	return c
}

func (c *CustomerMock) WithPhone(phone string) *CustomerMock {

	c.customer.Phone = phone
	return c
}

func (c *CustomerMock) WithAddress(address string) *CustomerMock {
	c.customer.Address = address
	return c
}

func (c *CustomerMock) WithZipCode(zipCode string) *CustomerMock {
	c.customer.ZipCode = zipCode
	return c
}

func (c *CustomerMock) WithCreatedAt(createdAt time.Time) *CustomerMock {
	c.customer.CreatedAt = createdAt
	return c
}
