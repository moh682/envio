package domain

import (
	"context"
	"time"

	"github.com/moh682/envio/backend/internal/utils"
)

type InvoiceCount struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

type InvoiceStatistics struct {
	Time         time.Time `json:"date"`
	InvoiceCount int       `json:"invoice_count"`
	VatAmount    float32   `json:"vat_amount"`
	TotalInclVat float32   `json:"total_incl_vat"`
	TotalExclVat float32   `json:"total_excl_vat"`
}

type InvoiceRepo interface {
	Store(ctx context.Context, invoice Invoice) error
	StoreMany(ctx context.Context, invoices []Invoice) error
	GetInvoices(ctx context.Context, limit int32, offset int32) ([]*Invoice, error)
	GetAllInvoicesSince(ctx context.Context, since time.Time) ([]*Invoice, error)
	GetYearlyInvoiceCount(ctx context.Context) ([]*InvoiceCount, error)
	GetDailyStatistics(ctx context.Context) ([]*InvoiceStatistics, error)
	GetMonthlyStatistics(ctx context.Context) ([]*InvoiceStatistics, error)
	GetYearlyStatistics(ctx context.Context) ([]*InvoiceStatistics, error)
	GetInvoiceByID(ctx context.Context, id ID) (*Invoice, error)
}

type InvoiceService interface {
	Store(invoice Invoice) error
	StoreMany(invoices []Invoice) error
	GetAll(limit, offset int32) ([]*Invoice, error)
	GetAllInvoicesSince(since time.Time) ([]*Invoice, error)
	GetYearlyInvoiceCount() ([]*InvoiceCount, error)
	GetYearlyStatistics() ([]*InvoiceStatistics, error)
	GetDailyStatistics() ([]*InvoiceStatistics, error)
	GetRevenueComparisonWithPreviousMonth() ([]*InvoiceStatistics, error)
	GetInvoiceByID(id ID) (*Invoice, error)
}

type invoiceService struct {
	repo InvoiceRepo
}

func NewInvoiceService(repo InvoiceRepo) InvoiceService {
	return &invoiceService{
		repo,
	}
}

func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
}
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

func (s *invoiceService) GetYearlyStatistics() ([]*InvoiceStatistics, error) {
	ctx := context.Background()
	return s.repo.GetYearlyStatistics(ctx)
}

func (s *invoiceService) GetInvoiceByID(id ID) (*Invoice, error) {
	ctx := context.Background()
	return s.repo.GetInvoiceByID(ctx, id)
}

func (s *invoiceService) StoreMany(invoices []Invoice) error {
	return s.repo.StoreMany(context.Background(), invoices)
}

func (s *invoiceService) GetRevenueComparisonWithPreviousMonth() ([]*InvoiceStatistics, error) {
	ctx := context.Background()
	stats, err := s.repo.GetMonthlyStatistics(ctx)
	if err != nil {
		return nil, err
	}

	currentMonth := time.Now()
	previousMonth := currentMonth.AddDate(0, -1, 0)

	filteredStats := utils.Filter(stats, func(i int, stat *InvoiceStatistics) bool {
		return utils.IsYearAndMonthEqual(stat.Time, currentMonth) || utils.IsYearAndMonthEqual(stat.Time, previousMonth)
	})

	// given no stats founds for the current and previous month
	if len(filteredStats) == 0 {
		filteredStats = append(filteredStats, &InvoiceStatistics{Time: currentMonth, InvoiceCount: 0, TotalExclVat: 0, TotalInclVat: 0, VatAmount: 0})
	}

	// given only one month found, either current or previous
	if len(filteredStats) == 1 {
		_, hasPreviousMonth := utils.Find(filteredStats, func(i int, stat *InvoiceStatistics) bool { return utils.IsYearAndMonthEqual(stat.Time, previousMonth) })
		if !hasPreviousMonth {
			filteredStats = append(filteredStats, &InvoiceStatistics{Time: StartOfMonth(previousMonth), InvoiceCount: 0, TotalExclVat: 0, TotalInclVat: 0, VatAmount: 0})
		}

		_, hasCurrentMonth := utils.Find(filteredStats, func(i int, stat *InvoiceStatistics) bool { return utils.IsYearAndMonthEqual(stat.Time, currentMonth) })
		if !hasCurrentMonth {
			filteredStats = append(filteredStats, &InvoiceStatistics{Time: StartOfMonth(currentMonth), InvoiceCount: 0, TotalExclVat: 0, TotalInclVat: 0, VatAmount: 0})
		}
	}

	return filteredStats, nil
}

func (s *invoiceService) GetDailyStatistics() ([]*InvoiceStatistics, error) {
	ctx := context.Background()
	return s.repo.GetDailyStatistics(ctx)
}

func (s *invoiceService) Store(invoice Invoice) error {
	err := invoice.Validate()
	if err != nil {
		return err
	}

	ctx := context.Background()
	return s.repo.Store(ctx, invoice)
}

func (s *invoiceService) GetAll(limit, offset int32) ([]*Invoice, error) {
	ctx := context.Background()
	return s.repo.GetInvoices(ctx, limit, offset)
}

func (s *invoiceService) GetYearlyInvoiceCount() ([]*InvoiceCount, error) {
	ctx := context.Background()
	return s.repo.GetYearlyInvoiceCount(ctx)
}

func (s *invoiceService) GetAllInvoicesSince(since time.Time) ([]*Invoice, error) {
	ctx := context.Background()
	return s.repo.GetAllInvoicesSince(ctx, since)
}
