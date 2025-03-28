package invoice_repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/invoice"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

var (
	ErrInvoiceParseTime = errors.New("could not convert date to time")
)

type postgresRepository struct {
	db *sql.DB
}

// GetInvoicesByOrganizationId implements invoice.Repository.
func (p *postgresRepository) GetInvoicesByOrganizationId(ctx context.Context, organizationId uuid.UUID, limit int64, offset int64) ([]*invoice.Invoice, error) {
	queries := db.New(p.db)

	invoiceResults, err := queries.GetAllInvoicesByOrganizationId(ctx, organizationId)
	if err != nil {
		return nil, err
	}

	invoices := make([]*invoice.Invoice, len(invoiceResults))

	for index, value := range invoiceResults {
		products, err := p.getAllProducts(ctx, organizationId, int32(value.Number))
		if err != nil {
			return nil, err
		}

		invoices[index] = &invoice.Invoice{
			Number:   int64(value.Number),
			IssuedAt: value.IssueDate.Time,
			Total:    value.Total,
			IsVat:    value.IsVat,
			Products: products,
		}
	}

	return invoices, nil
}

// Store implements invoice.Repository.
func (p *postgresRepository) Store(ctx context.Context, invoice invoice.Invoice) error {
	panic("unimplemented")
}

func (p *postgresRepository) getAllProducts(ctx context.Context, organizationId uuid.UUID, invoiceNumber int32) ([]*invoice.Product, error) {
	queries := db.New(p.db)

	results, err := queries.GetAllProductsByInvoiceId(ctx, db.GetAllProductsByInvoiceIdParams{OrganizationID: organizationId, InvoiceNumber: invoiceNumber})
	if err != nil {
		return nil, err
	}

	products := make([]*invoice.Product, len(results))

	for index, value := range results {
		products[index] = &invoice.Product{
			ID:          value.ID,
			Serial:      value.Serial.String,
			Description: value.Description,
			Quantity:    float64(value.Quantity),
			Rate:        value.Rate,
			Total:       value.Total,
		}
	}

	return products, nil
}

func NewPostgres(db *sql.DB) invoice.Repository {
	return &postgresRepository{db}
}

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
