package invoice

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Store(ctx context.Context, invoice Invoice) error
	StoreMany(ctx context.Context, invoices []Invoice) error
	GetInvoices(ctx context.Context, limit int64, offset int64) ([]*Invoice, error)
	GetCustomerByID(ctx context.Context, id uuid.UUID) (*Customer, error)
	GetAllInvoicesSince(ctx context.Context, since time.Time) ([]*Invoice, error)
	GetYearlyInvoiceCount(ctx context.Context) (map[time.Time]int64, error)
	GetDailyStatistics(ctx context.Context) ([]*Statistic, error)
	GetMonthlyStatistics(ctx context.Context) ([]*Statistic, error)
	GetYearlyStatistics(ctx context.Context) ([]*Statistic, error)
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error)
}

type Service interface {
	Store(ctx context.Context, invoice Invoice) error
	StoreMany(ctx context.Context, invoices []Invoice) error
	GetInvoices(ctx context.Context, limit int64, offset int64) ([]*Invoice, error)
	GetAllInvoicesSince(ctx context.Context, since time.Time) ([]*Invoice, error)
	GetCustomerByID(ctx context.Context, id uuid.UUID) (*Customer, error)
	GetYearlyInvoiceCount(ctx context.Context) (map[time.Time]int64, error)
	GetDailyStatistics(ctx context.Context) ([]*Statistic, error)
	GetMonthlyStatistics(ctx context.Context) ([]*Statistic, error)
	GetYearlyStatistics(ctx context.Context) ([]*Statistic, error)
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Store(ctx context.Context, invoice Invoice) error {
	return s.repo.Store(ctx, invoice)
}

func (s *service) StoreMany(ctx context.Context, invoices []Invoice) error {
	return s.repo.StoreMany(ctx, invoices)
}

func (s *service) GetInvoices(ctx context.Context, limit int64, offset int64) ([]*Invoice, error) {
	return s.repo.GetInvoices(ctx, limit, offset)
}

func (s *service) GetCustomerByID(ctx context.Context, id uuid.UUID) (*Customer, error) {
	return s.repo.GetCustomerByID(ctx, id)
}

func (s *service) GetAllInvoicesSince(ctx context.Context, since time.Time) ([]*Invoice, error) {
	return s.repo.GetAllInvoicesSince(ctx, since)
}

func (s *service) GetYearlyInvoiceCount(ctx context.Context) (map[time.Time]int64, error) {
	return s.repo.GetYearlyInvoiceCount(ctx)
}

func (s *service) GetDailyStatistics(ctx context.Context) ([]*Statistic, error) {
	return s.repo.GetDailyStatistics(ctx)

}

func (s *service) GetMonthlyStatistics(ctx context.Context) ([]*Statistic, error) {
	return s.repo.GetMonthlyStatistics(ctx)
}

func (s *service) GetYearlyStatistics(ctx context.Context) ([]*Statistic, error) {
	return s.repo.GetYearlyStatistics(ctx)
}

func (s *service) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error) {
	return s.repo.GetInvoiceByID(ctx, id)
}
