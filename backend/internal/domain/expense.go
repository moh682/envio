package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidCompany = errors.New("company is invalid")
	ErrInvalidEntries = errors.New("entries are invalid")
)

type ExpenseCount struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

type ExpenseRepo interface {
	Create(ctx context.Context, exp *Expense) error
	GetAll(ctx context.Context, limit, offset int32) ([]*Expense, error)
	CreateCompany(ctx context.Context, company *Company) error
	GetCompaniesByName(ctx context.Context, name string) ([]*Company, error)
	GetCompanyByID(ctx context.Context, id string) (*Company, error)
	GetAllCompanies(ctx context.Context) ([]*Company, error)
	GetAllExpensesSince(ctx context.Context, since time.Time) ([]*Expense, error)
	GetYearlyExpensesCount(ctx context.Context) ([]*ExpenseCount, error)
}

type Expense struct {
	ID            ID            `json:"id"`
	Company       *Company      `json:"company"`
	TotalInclVat  float32       `json:"total_incl_vat"`
	TotalExclVat  float32       `json:"total_excl_vat"`
	VatAmount     float32       `json:"vat_amount"`
	VatPercentage float32       `json:"vat_percentage"`
	IssuedAt      time.Time     `json:"issued_at"`
	PaidAt        time.Time     `json:"paid_at"`
	PaidWith      PaymentMethod `json:"paid_with"`

	Entries []*Entry `json:"entries"`
}

type Entry struct {
	ID           ID      `json:"id"`
	Serial       string  `json:"serial"`
	Description  string  `json:"description"`
	UnitPrice    float32 `json:"unit_price"`
	Quantity     int32   `json:"quantity"`
	TotalInclVat float32 `json:"total_incl_vat"`
	TotalExclVat float32 `json:"total_excl_vat"`
}

func (e *Expense) Validate() error {
	_, err := uuid.Parse(e.ID.String())
	if err != nil {
		return err
	}

	if e.Company == nil {
		return ErrInvalidCompany
	}

	if e.Entries == nil {
		return ErrInvalidEntries
	}

	for _, entry := range e.Entries {
		if err := entry.Validate(); err != nil {
			return errors.New("entry with uuid:" + entry.ID.String() + " is invalid")
		}
	}

	return nil
}

func (e *Entry) Validate() error {
	_, err := uuid.Parse(e.ID.String())
	if err != nil {
		return err
	}

	return nil
}
