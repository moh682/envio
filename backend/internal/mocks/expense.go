package mocks

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
)

type ExpenseMock struct {
	expense *domain.Expense
}

func NewExpense() *ExpenseMock {
	id := uuid.New()
	exp := &domain.Expense{
		ID:            domain.ID(id),
		TotalExclVat:  800,
		VatAmount:     200,
		VatPercentage: .25,
		TotalInclVat:  1000,
		IssuedAt:      gofakeit.Date(),
		Company:       NewCompany().ToCompany(),
		PaidAt:        gofakeit.Date(),
		PaidWith:      domain.MOBILEPAY,
		Entries:       []*domain.Entry{NewExpenseEntryMock().ToEntry()},
	}

	return &ExpenseMock{expense: exp}
}

func (e *ExpenseMock) ToExpense() *domain.Expense {
	return e.expense
}

func (e *ExpenseMock) WithCompany(company *domain.Company) *ExpenseMock {
	e.expense.Company = company
	return e
}

func (e *ExpenseMock) WithTotalExclVat(totalExclVat float32) *ExpenseMock {
	e.expense.TotalExclVat = totalExclVat
	return e
}

func (e *ExpenseMock) WithVatAmount(vatAmount float32) *ExpenseMock {
	e.expense.VatAmount = vatAmount
	return e
}

func (e *ExpenseMock) WithVatPercentage(vatPercentage float32) *ExpenseMock {
	e.expense.VatPercentage = vatPercentage
	return e
}

func (e *ExpenseMock) WithTotalInclVat(totalInclVat float32) *ExpenseMock {
	e.expense.TotalInclVat = totalInclVat
	return e
}

func (e *ExpenseMock) WithIssuedAt(issuedAt time.Time) *ExpenseMock {
	e.expense.IssuedAt = issuedAt
	return e
}

func (e *ExpenseMock) WithExpense(expense *domain.Expense) *ExpenseMock {
	e.expense = expense
	return e
}
