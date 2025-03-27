package invoice_repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
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

func NewPostgres(db *sql.DB) invoice.Repository {
	return &postgresRepository{db}
}

func (r *postgresRepository) GetCustomerByID(ctx context.Context, id uuid.UUID) (*invoice.Customer, error) {
	q := db.New(r.db)
	customer, err := q.GetCustomerById(ctx, id)
	if err != nil {
		return nil, err
	}

	cust := invoice.Customer{
		ID:      customer.ID,
		Name:    customer.Name.String,
		Email:   customer.Email.String,
		Address: customer.Address.String,
		Phone:   customer.Phone.String,
		Car: &invoice.Car{
			Registration: customer.CarRegistration.String,
		},
		Zip: customer.ZipCode.String,
	}

	return &cust, nil
}

func (r *postgresRepository) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*invoice.Invoice, error) {
	q := db.New(r.db)
	res, err := q.GetInvoiceById(ctx, id)
	if err != nil {
		return nil, err
	}
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

func (r *postgresRepository) GetDailyStatistics(ctx context.Context) ([]*invoice.Statistic, error) {
	q := db.New(r.db)
	result, err := q.GetDailyInvoiceStatistics(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*invoice.Statistic, len(result))
	for i, stat := range result {
		t, err := time.Parse("2006", stat.Date)
		if err != nil {
			return nil, ErrInvoiceParseTime
		}
		results[i] = &invoice.Statistic{
			Date:         t,
			Count:        stat.InvoiceCount,
			TotalExclVat: stat.TotalExcludeVat,
			TotalInclVat: stat.TotalIncludeVat,
			VatAmount:    stat.VatAmount,
		}
	}

	return results, nil
}

func (r *postgresRepository) GetMonthlyStatistics(ctx context.Context) ([]*invoice.Statistic, error) {
	q := db.New(r.db)
	result, err := q.GetMonthlyInvoiceStatistics(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*invoice.Statistic, len(result))
	for i, stat := range result {
		t, err := time.Parse("01/2006", stat.Date)
		if err != nil {
			log.Println(err)
			return nil, ErrInvoiceParseTime
		}
		results[i] = &invoice.Statistic{
			Date:         t,
			Count:        stat.InvoiceCount,
			TotalExclVat: stat.TotalExcludeVat,
			TotalInclVat: stat.TotalIncludeVat,
			VatAmount:    stat.VatAmount,
		}

	}

	return results, nil
}
func (r *postgresRepository) GetYearlyStatistics(ctx context.Context) ([]*invoice.Statistic, error) {
	q := db.New(r.db)
	result, err := q.GetYearlyInvoiceStatistics(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*invoice.Statistic, len(result))
	for i, stat := range result {
		t, err := time.Parse("02/01/2006", stat.Date)
		if err != nil {
			return nil, ErrInvoiceParseTime
		}
		results[i] = &invoice.Statistic{
			Date:         t,
			Count:        stat.InvoiceCount,
			TotalExclVat: stat.TotalExcludeVat,
			TotalInclVat: stat.TotalIncludeVat,
			VatAmount:    stat.VatAmount,
		}

	}

	return results, nil
}

func (r *postgresRepository) GetYearlyInvoiceCount(ctx context.Context) (map[time.Time]int64, error) {
	q := db.New(r.db)
	result, err := q.GetYearlyInvoiceCount(ctx)
	if err != nil {
		return nil, err
	}

	invoiceCounts := make(map[time.Time]int64)
	for _, row := range result {
		year, err := strconv.ParseInt(row.Year, 10, 64)
		if err != nil {
			return nil, errors.New("could not convert year to string")
		}

		date := time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)

		invoiceCounts[date] = int64(row.Count)
	}

	return invoiceCounts, nil
}

func (r *postgresRepository) GetAllInvoicesSince(ctx context.Context, since time.Time) ([]*invoice.Invoice, error) {
	q := db.New(r.db)
	result, err := q.GetAllInvoicesSince(ctx, sql.NullTime{
		Time:  since,
		Valid: !time.Now().IsZero(),
	})
	if err != nil {
		return nil, err
	}

	invoices := make([]*invoice.Invoice, len(result))
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

func (r *postgresRepository) GetInvoices(
	ctx context.Context,
	limit int64,
	offset int64,
) ([]*invoice.Invoice, error) {

	q := db.New(r.db)
	arg := db.GetAllInvoicesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	result, err := q.GetAllInvoices(ctx, arg)
	if err != nil {
		return nil, err
	}

	invoices := make([]*invoice.Invoice, len(result))
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

func (r *postgresRepository) Store(ctx context.Context, i invoice.Invoice) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := db.New(r.db)
	q.WithTx(tx)

	invoiceArg := db.CreateInvoiceParams{
		ID:              i.ID,
		IssuedAt:        toNullTime(i.IssuedAt),
		VatPercentage:   i.VatPct,
		VatAmount:       i.VatAmount,
		TotalExcludeVat: i.TotalExclVat,
		TotalIncludeVat: i.Total,
		PaymentStatus:   string(i.Status),
	}
	err = q.CreateInvoice(ctx, invoiceArg)
	if err != nil {
		return err
	}

	custArg := db.CreateCustomerParams{
		ID:              i.Customer.ID,
		InvoiceID:       i.ID,
		Name:            toNullString(i.Customer.Name),
		Address:         toNullString(i.Customer.Address),
		Email:           toNullString(i.Customer.Email),
		CarRegistration: toNullString(i.Customer.Car.Registration),
		Phone:           toNullString(i.Customer.Phone),
		ZipCode:         toNullString(i.Customer.Zip),
	}
	err = q.CreateCustomer(ctx, custArg)
	if err != nil {
		return err
	}

	for _, c := range i.Products {
		costArg := db.CreateCostParams{
			ID:            c.ID,
			InvoiceID:     i.ID,
			Total:         i.Total,
			ProductNumber: toNullString(c.Serial),
			Description:   toNullString(c.Description),
			Quantity:      c.Quantity,
			UnitPrice:     float64(c.UnitPrice),
		}
		err := q.CreateCost(ctx, costArg)
		if err != nil {
			return err
		}

	}

	if len(i.Payments) > 0 {
		for _, p := range i.Payments {
			paymentArg := db.CreatePaymentParams{
				ID:        p.ID,
				InvoiceID: i.ID,
				Method:    string(p.Method),
				Amount:    p.Amount,
			}
			err := q.CreatePayment(ctx, paymentArg)
			if err != nil {
				return err
			}

		}
	}

	return tx.Commit()
}

func (r *postgresRepository) StoreMany(ctx context.Context, il []invoice.Invoice) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := db.New(r.db)
	q.WithTx(tx)

	for _, i := range il {

		invoiceArg := db.CreateInvoiceParams{
			ID:              i.ID,
			IssuedAt:        toNullTime(i.IssuedAt),
			VatPercentage:   i.VatPct,
			VatAmount:       i.VatAmount,
			TotalExcludeVat: i.TotalExclVat,
			TotalIncludeVat: i.Total,
			PaymentStatus:   string(i.Status),
		}
		err = q.CreateInvoice(ctx, invoiceArg)
		if err != nil {
			return err
		}

		custArg := db.CreateCustomerParams{
			ID:              i.Customer.ID,
			InvoiceID:       i.ID,
			Name:            toNullString(i.Customer.Name),
			Address:         toNullString(i.Customer.Address),
			Email:           toNullString(i.Customer.Email),
			CarRegistration: toNullString(i.Customer.Car.Registration),
			Phone:           toNullString(i.Customer.Phone),
			ZipCode:         toNullString(i.Customer.Zip),
		}
		err = q.CreateCustomer(ctx, custArg)
		if err != nil {
			return err
		}

		for _, p := range i.Products {
			costArg := db.CreateCostParams{
				ID:            p.ID,
				InvoiceID:     i.ID,
				Total:         p.Total,
				ProductNumber: toNullString(p.Serial),
				Description:   toNullString(p.Description),
				Quantity:      p.Quantity,
				UnitPrice:     p.UnitPrice,
			}
			err := q.CreateCost(ctx, costArg)
			if err != nil {
				return err
			}

		}

		for _, p := range i.Payments {
			paymentArg := db.CreatePaymentParams{
				ID:        p.ID,
				InvoiceID: i.ID,
				Method:    string(p.Method),
				Amount:    p.Amount,
			}
			err := q.CreatePayment(ctx, paymentArg)
			if err != nil {
				return err
			}

		}
	}

	return tx.Commit()
}

func (r *postgresRepository) invoiceToModel(inv db.Invoice, c db.Customer, co []db.Cost, pa []db.Payment) (*invoice.Invoice, error) {
	products := make([]*invoice.Product, len(co))
	for idx, c := range co {
		products[idx] = &invoice.Product{
			ID:          c.ID,
			Description: c.Description.String,
			Serial:      c.ProductNumber.String,
			Quantity:    float64(c.Quantity),
			UnitPrice:   c.UnitPrice,
			Total:       c.Total,
		}
	}

	payments := make([]*invoice.Payment, len(pa))
	for idx, p := range pa {
		payments[idx] = &invoice.Payment{
			ID:     c.ID,
			Amount: p.Amount,
			PaidAt: p.PaidAt.Time,
			Method: invoice.PaymentMethod(p.Method),
		}
	}

	customer := invoice.Customer{
		ID:      c.ID,
		Name:    c.Name.String,
		Email:   c.Email.String,
		Address: c.Address.String,
		Phone:   c.Phone.String,
		Car: &invoice.Car{
			Registration: c.CarRegistration.String,
		},
		Zip: c.ZipCode.String,
	}

	invoice := &invoice.Invoice{
		ID:           inv.ID,
		Number:       int64(inv.InvoiceNumber),
		IssuedAt:     inv.IssuedAt.Time,
		VatPct:       inv.VatPercentage,
		VatAmount:    inv.VatAmount,
		TotalExclVat: inv.TotalExcludeVat,
		Total:        inv.TotalIncludeVat,
		Status:       invoice.FullyPaid, // TODO: map from db
		Customer:     customer,
		Products:     products,
		Payments:     payments,
	}

	return invoice, nil
}
