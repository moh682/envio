package mocks

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
)

type ExpenseEntryMock struct {
	entry *domain.Entry
}

func NewExpenseEntryMock() *ExpenseEntryMock {
	id := uuid.New()
	entry := &domain.Entry{
		ID:           domain.ID(id),
		Description:  gofakeit.ProductDescription(),
		TotalExclVat: 800.0,
		UnitPrice:    1000.0,
		Quantity:     1,
		TotalInclVat: 1000.0,
	}
	return &ExpenseEntryMock{entry}
}

func (e *ExpenseEntryMock) ToEntry() *domain.Entry {
	return e.entry
}

func (e *ExpenseEntryMock) WithDescription(description string) *ExpenseEntryMock {
	e.entry.Description = description
	return e
}

func (e *ExpenseEntryMock) WithTotalExclVat(totalExclVat float32) *ExpenseEntryMock {
	e.entry.TotalExclVat = totalExclVat
	return e
}

func (e *ExpenseEntryMock) WithUnitPrice(unitPrice float32) *ExpenseEntryMock {
	e.entry.UnitPrice = unitPrice
	return e
}

func (e *ExpenseEntryMock) WithQuantity(quantity int32) *ExpenseEntryMock {
	e.entry.Quantity = quantity
	return e
}

func (e *ExpenseEntryMock) WithTotalInclVat(totalInclVat float32) *ExpenseEntryMock {
	e.entry.TotalInclVat = totalInclVat
	return e
}

func (e *ExpenseEntryMock) WithEntry(entry *domain.Entry) *ExpenseEntryMock {
	e.entry = entry
	return e
}
