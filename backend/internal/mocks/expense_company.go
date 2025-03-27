package mocks

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
)

type ExpenseCompanyMock struct {
	company *domain.Company
}

func NewCompany() *ExpenseCompanyMock {
	id := uuid.New()
	company := &domain.Company{
		ID:        domain.ID(id),
		Name:      "Company",
		Cvr:       gofakeit.Number(10000000, 99999999),
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
	}

	return &ExpenseCompanyMock{company: company}
}

func (c *ExpenseCompanyMock) ToCompany() *domain.Company {
	return c.company
}

func (c *ExpenseCompanyMock) WithName(name string) *ExpenseCompanyMock {
	c.company.Name = name
	return c
}

func (c *ExpenseCompanyMock) WithCvr(cvr int) *ExpenseCompanyMock {
	c.company.Cvr = cvr
	return c
}

func (c *ExpenseCompanyMock) WithCreatedAt(createdAt time.Time) *ExpenseCompanyMock {
	c.company.CreatedAt = createdAt
	return c
}

func (c *ExpenseCompanyMock) WithUpdatedAt(updatedAt time.Time) *ExpenseCompanyMock {
	c.company.UpdatedAt = updatedAt
	return c
}

func (c *ExpenseCompanyMock) WithCompany(company *domain.Company) *ExpenseCompanyMock {
	c.company = company
	return c
}
