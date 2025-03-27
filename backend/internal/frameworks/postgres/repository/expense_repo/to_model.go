package expenserepo

import (
	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

func toCompanyModel(c db.Company) (*domain.Company, error) {
	parsedID, err := uuid.Parse(c.ID)
	if err != nil {
		return nil, err
	}
	return &domain.Company{
		ID:        domain.ID(parsedID),
		Name:      c.Name,
		Cvr:       int(c.Cvr.Int64),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt.Time,
	}, nil
}

func toExpenseModel(e db.Expense, entries []db.ExpensesEntry, company db.Company) (*domain.Expense, error) {
	parsedID, err := uuid.Parse(e.ID)
	if err != nil {
		return nil, err
	}

	comp, err := toCompanyModel(company)
	if err != nil {
		return nil, err
	}

	exp := &domain.Expense{
		ID:            domain.ID(parsedID),
		TotalInclVat:  float32(e.TotalInclVat),
		TotalExclVat:  float32(e.TotalExclVat),
		VatAmount:     float32(e.VatAmount),
		VatPercentage: float32(e.VatRate),
		IssuedAt:      e.IssuedAt,
		PaidAt:        e.PaidAt.Time,
		PaidWith:      domain.PaymentMethod(e.PaidWith.Int64),

		Company: comp,
		Entries: make([]*domain.Entry, len(entries)),
	}

	if len(entries) > 0 {
		for idx, ee := range entries {
			exp.Entries[idx] = &domain.Entry{
				ID:           domain.ID(parsedID),
				Serial:       ee.Serial.String,
				Description:  ee.Description.String,
				UnitPrice:    float32(ee.UnitPrice),
				Quantity:     int32(ee.Quantity),
				TotalExclVat: float32(ee.TotalExclVat),
				TotalInclVat: float32(ee.TotalInclVat),
			}
		}
	}

	return exp, nil
}
