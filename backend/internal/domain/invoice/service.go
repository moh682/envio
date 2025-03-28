package invoice

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Store(ctx context.Context, invoice Invoice) error
	GetInvoicesByOrganizationId(ctx context.Context, organizationId uuid.UUID, limit int64, offset int64) ([]*Invoice, error)
}

type Service interface {
	Store(ctx context.Context, invoice Invoice) error
	GetInvoicesByOrganizationId(ctx context.Context, organizationId uuid.UUID, limit int64, offset int64) ([]*Invoice, error)
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

func (s *service) GetInvoicesByOrganizationId(ctx context.Context, organizationId uuid.UUID, limit int64, offset int64) ([]*Invoice, error) {
	return s.repo.GetInvoicesByOrganizationId(ctx,organizationId, limit, offset)
}
