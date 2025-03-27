package invoicerepo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

var (
	ErrInvoiceParseTime = errors.New("could not convert date to time")
)

type invoiceRepo struct {
	db *sql.DB
}

func NewInvoiceRepo(db *sql.DB) domain.InvoiceRepo {
	repo := &invoiceRepo{db}
	return repo
}

func (r *invoiceRepo) GetInvoiceByID(ctx context.Context, id domain.ID) (*domain.Invoice, error) {
	q := db.New(r.db)
	res, err := q.GetInvoiceById(ctx, id.String())
	payments, err := q.GetPaymentsByInvoice(ctx, res.ID)
	if err != nil {
		return nil, err
	}

	costs, err := q.GetAllCostsByInvoiceId(ctx, res.ID)
	if err != nil {
		return nil, err
	}

	customer, err := q.GetCustomerByInvoiceId(ctx, res.ID)
	if err != nil {
		return nil, err
	}
	parsed, err := r.invoiceToModel(res, customer, costs, payments)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (r *invoiceRepo) GetDailyStatistics(ctx context.Context) ([]*domain.InvoiceStatistics, error) {
	q := db.New(r.db)
	result, err := q.GetDailyInvoiceStatistics(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*domain.InvoiceStatistics, len(result))
	for i, stat := range result {
		t, err := time.Parse("2006", stat.Date)
		if err != nil {
			return nil, ErrInvoiceParseTime
		}
		results[i] = &domain.InvoiceStatistics{
			Time:         t,
			InvoiceCount: int(stat.InvoiceCount),
			TotalExclVat: float32(stat.TotalExcludeVat.Float64),
			TotalInclVat: float32(stat.TotalIncludeVat.Float64),
			VatAmount:    float32(stat.VatAmount.Float64),
		}
	}

	return results, nil
}

func (r *invoiceRepo) GetMonthlyStatistics(ctx context.Context) ([]*domain.InvoiceStatistics, error) {
	q := db.New(r.db)
	result, err := q.GetMonthlyInvoiceStatistics(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*domain.InvoiceStatistics, len(result))
	for i, stat := range result {
		t, err := time.Parse("01/2006", stat.Date)
		if err != nil {
			log.Println(err)
			return nil, ErrInvoiceParseTime
		}
		results[i] = &domain.InvoiceStatistics{
			Time:         t,
			InvoiceCount: int(stat.InvoiceCount),
			TotalExclVat: float32(stat.TotalExcludeVat.Float64),
			TotalInclVat: float32(stat.TotalIncludeVat.Float64),
			VatAmount:    float32(stat.VatAmount.Float64),
		}

	}

	return results, nil
}
func (r *invoiceRepo) GetYearlyStatistics(ctx context.Context) ([]*domain.InvoiceStatistics, error) {
	q := db.New(r.db)
	result, err := q.GetYearlyInvoiceStatistics(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*domain.InvoiceStatistics, len(result))
	for i, stat := range result {
		t, err := time.Parse("02/01/2006", stat.Date)
		if err != nil {
			return nil, ErrInvoiceParseTime
		}
		results[i] = &domain.InvoiceStatistics{
			Time:         t,
			InvoiceCount: int(stat.InvoiceCount),
			TotalExclVat: float32(stat.TotalExcludeVat.Float64),
			TotalInclVat: float32(stat.TotalIncludeVat.Float64),
			VatAmount:    float32(stat.VatAmount.Float64),
		}

	}

	return results, nil
}

func (r *invoiceRepo) GetYearlyInvoiceCount(ctx context.Context) ([]*domain.InvoiceCount, error) {
	q := db.New(r.db)
	result, err := q.GetYearlyInvoiceCount(ctx)
	if err != nil {
		return nil, err
	}

	invoiceCounts := make([]*domain.InvoiceCount, len(result))
	for index, row := range result {
		year, err := strconv.ParseInt(row.Year, 10, 64)
		if err != nil {
			return nil, errors.New("could not convert year to string")
		}

		invoiceCounts[index] = &domain.InvoiceCount{
			Year:  int(year),
			Count: int(row.Count),
		}
	}

	return invoiceCounts, nil
}

func (r *invoiceRepo) GetAllInvoicesSince(ctx context.Context, since time.Time) ([]*domain.Invoice, error) {
	q := db.New(r.db)
	result, err := q.GetAllInvoicesSince(ctx, sql.NullTime{
		Time:  since,
		Valid: !time.Now().IsZero(),
	})
	if err != nil {
		return nil, err
	}

	invoices := make([]*domain.Invoice, len(result))
	for index, inv := range result {
		payments, err := q.GetPaymentsByInvoice(ctx, inv.ID)
		if err != nil {
			return nil, err
		}

		costs, err := q.GetAllCostsByInvoiceId(ctx, inv.ID)
		if err != nil {
			return nil, err
		}

		customer, err := q.GetCustomerByInvoiceId(ctx, inv.ID)
		if err != nil {
			return nil, err
		}
		parsed, err := r.invoiceToModel(inv, customer, costs, payments)
		if err != nil {
			return nil, err
		}
		invoices[index] = parsed
	}

	return invoices, nil
}

func (r *invoiceRepo) GetInvoices(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]*domain.Invoice, error) {

	q := db.New(r.db)
	arg := db.GetAllInvoicesParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	}
	result, err := q.GetAllInvoices(ctx, arg)
	if err != nil {
		return nil, err
	}

	invoices := make([]*domain.Invoice, len(result))
	for index, inv := range result {
		payments, err := q.GetPaymentsByInvoice(ctx, inv.ID)
		if err != nil {
			return nil, err
		}

		costs, err := q.GetAllCostsByInvoiceId(ctx, inv.ID)
		if err != nil {
			return nil, err
		}

		customer, err := q.GetCustomerByInvoiceId(ctx, inv.ID)
		if err != nil {
			return nil, err
		}
		parsed, err := r.invoiceToModel(inv, customer, costs, payments)
		if err != nil {
			return nil, err
		}
		invoices[index] = parsed
	}

	return invoices, nil

}
