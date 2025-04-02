package financial_year

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetFinancialYearsByUserIdOrganizationId(ctx context.Context, userId uuid.UUID, organizationId uuid.UUID) ([]*FinancialYear, error)
}
