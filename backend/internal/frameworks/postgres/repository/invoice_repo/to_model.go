package invoicerepo

import (
	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

func (r *invoiceRepo) invoiceToModel(inv db.Invoice, c db.Customer, co []db.Cost, pa []db.Payment) (*domain.Invoice, error) {
	costs := make([]*domain.Cost, len(co))
	for idx, c := range co {
		id := uuid.MustParse(c.ID)
		costs[idx] = &domain.Cost{
			ID:          domain.ID(id),
			Description: c.Description.String,
			ProductNr:   c.ProductNumber.String,
			Quantity:    float32(c.Quantity),
			UnitPrice:   float32(c.UnitPrice),
			Total:       float32(c.Total),
		}
	}

	payments := make([]*domain.Payment, len(pa))
	for idx, p := range pa {
		id := uuid.MustParse(p.ID)
		payments[idx] = &domain.Payment{
			ID:     domain.ID(id),
			Amount: float32(p.Amount),
			PaidAt: p.PaidAt.Time,
			Method: p.Method,
		}
	}

	customerId := uuid.MustParse(c.ID)
	customer := &domain.Customer{
		ID:              domain.ID(customerId),
		Name:            c.Name.String,
		Email:           c.Email.String,
		Address:         c.Address.String,
		Phone:           c.Phone.String,
		CarRegistration: c.CarRegistration.String,
		ZipCode:         c.ZipCode.String,
		CreatedAt:       c.CreatedAt.Time,
	}

	invoiceId := uuid.MustParse(inv.ID)
	invoice := &domain.Invoice{
		ID:            domain.ID(invoiceId),
		InvoiceNr:     int32(inv.InvoiceNumber),
		IssuedAt:      inv.IssuedAt.Time,
		VatPct:        float32(inv.VatPercentage),
		VatAmount:     float32(inv.VatAmount),
		TotalExclVat:  float32(inv.TotalExcludeVat),
		Total:         float32(inv.TotalIncludeVat),
		PaymentStatus: domain.Paid, // TODO: map from db
		Customer:      customer,
		Costs:         costs,
		Payments:      payments,
	}
	err := invoice.Validate()
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *invoiceRepo) migrationToModel(im db.InvoiceMigration) *domain.Migration {
	id := uuid.MustParse(im.ID)
	return &domain.Migration{
		ID:                      domain.ID(id),
		InvoiceNr:               int32(im.InvoiceNumber),
		FailedAt:                im.FailedAt.Time,
		FilePath:                im.FilePath.String,
		HasMigratedSuccessfully: im.HasMigratedSuccessfully.Int64 == 1,
		LastFailedAt:            im.LastFailedAttempt.Time,
	}
}
