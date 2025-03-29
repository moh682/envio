package organization

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*Organization, error)
}

type Service interface {
	GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*Organization, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*Organization, error) {
	return s.repo.GetOrganizationByUserId(ctx, userId)
}
