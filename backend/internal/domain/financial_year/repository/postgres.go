package financial_year_repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/financial_year"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

type postgresRepository struct {
	db *sql.DB
}

// GetFinancialYearsByUserIdAndOrganizationId implements financial_year.Repository.
func (p *postgresRepository) GetFinancialYearsByUserIdOrganizationId(ctx context.Context, userId uuid.UUID, organizationId uuid.UUID) ([]*financial_year.FinancialYear, error) {
	queries := db.New(p.db)

	results, err := queries.GetFinancialYearsByUserIdOrganizationId(ctx, db.GetFinancialYearsByUserIdOrganizationIdParams{UserID: userId, OrganizationID: organizationId})
	if err != nil {
		return nil, err
	}

	financialYears := make([]*financial_year.FinancialYear, len(results))

	for index, value := range results {
		financialYears[index] = &financial_year.FinancialYear{Year: value}
	}

	return financialYears, nil
}

func NewPostgres(db *sql.DB) financial_year.Repository {
	return &postgresRepository{db}
}
