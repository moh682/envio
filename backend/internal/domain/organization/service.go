package organization

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*Organization, error)
	CreateOrganization(ctx context.Context, userId uuid.UUID, name string, invoiceNumberStart int32) (*Organization, error)
}

type Service interface {
	GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*Organization, error)
	CreateOrganization(ctx context.Context, userId uuid.UUID, name string, invoiceNumberStart int32) (*Organization, error)
}

type service struct {
	repo Repository
}

// CreateOrganization implements Service.
func (s *service) CreateOrganization(ctx context.Context, userId uuid.UUID, name string, invoiceNumberStart int32) (*Organization, error) {
	return s.repo.CreateOrganization(ctx, userId, name, invoiceNumberStart)
}

func (s *service) GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*Organization, error) {
	return s.repo.GetOrganizationByUserId(ctx, userId)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
