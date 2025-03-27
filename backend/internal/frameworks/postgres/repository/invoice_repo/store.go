package invoicerepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

func toNullString(v string) sql.NullString {
	return sql.NullString{
		Valid:  v != "",
		String: v,
	}
}

func toNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func (r *invoiceRepo) Store(ctx context.Context, i domain.Invoice) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := db.New(r.db)
	q.WithTx(tx)

	invoiceArg := db.CreateInvoiceParams{
		ID:              i.ID.String(),
		IssuedAt:        toNullTime(i.IssuedAt),
		VatPercentage:   float64(i.VatPct),
		VatAmount:       float64(i.VatAmount),
		TotalExcludeVat: float64(i.TotalExclVat),
		TotalIncludeVat: float64(i.Total),
		PaymentStatus:   i.PaymentStatus.String(),
	}
	err = q.CreateInvoice(ctx, invoiceArg)
	if err != nil {
		return err
	}

	custArg := db.CreateCustomerParams{
		ID:              i.Customer.ID.String(),
		InvoiceID:       i.ID.String(),
		Name:            toNullString(i.Customer.Name),
		Address:         toNullString(i.Customer.Address),
		Email:           toNullString(i.Customer.Email),
		CarRegistration: toNullString(i.Customer.CarRegistration),
		Phone:           toNullString(i.Customer.Phone),
		ZipCode:         toNullString(i.Customer.ZipCode),
	}
	err = q.CreateCustomer(ctx, custArg)
	if err != nil {
		return err
	}

	if len(i.Costs) > 0 {
		for _, c := range i.Costs {
			costArg := db.CreateCostParams{
				ID:            c.ID.String(),
				InvoiceID:     i.ID.String(),
				Total:         float64(c.Total),
				ProductNumber: toNullString(c.ProductNr),
				Description:   toNullString(c.Description),
				Quantity:      int64(c.Quantity),
				UnitPrice:     float64(c.UnitPrice),
			}
			err := q.CreateCost(ctx, costArg)
			if err != nil {
				return err
			}

		}
	}

	if len(i.Payments) > 0 {
		for _, p := range i.Payments {
			paymentArg := db.CreatePaymentParams{
				ID:        p.ID.String(),
				InvoiceID: i.ID.String(),
				Method:    p.Method,
				Amount:    float64(p.Amount),
			}
			err := q.CreatePayment(ctx, paymentArg)
			if err != nil {
				return err
			}

		}
	}

	return tx.Commit()
}
