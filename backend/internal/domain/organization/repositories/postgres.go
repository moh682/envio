package organization_repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/financial_year"
	"github.com/moh682/envio/backend/internal/domain/organization"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

type postgresRepository struct {
	db *sql.DB
}

// CreateOrganization implements organization.Repository.
func (p *postgresRepository) CreateOrganization(ctx context.Context, userId uuid.UUID, name string, invoiceNumberStart int32) (*organization.Organization, error) {
	now := time.Now()
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	qtx := db.New(p.db).WithTx(tx)

	organizationId := uuid.New()
	err = qtx.CreateOrganization(ctx, db.CreateOrganizationParams{ID: organizationId, Name: name, InvoiceNumberStart: invoiceNumberStart})
	if err != nil {
		return nil, err
	}

	err = qtx.CreateOrganizationUser(ctx, db.CreateOrganizationUserParams{UserID: userId, OrganizationID: organizationId})
	if err != nil {
		return nil, err
	}

	err = qtx.CreateFinancialYear(ctx, db.CreateFinancialYearParams{OrganizationID: organizationId, Year: int32(now.Year())})
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return p.GetOrganizationByUserId(ctx, userId)

}

func (p *postgresRepository) GetOrganizationByUserId(ctx context.Context, userId uuid.UUID) (*organization.Organization, error) {
	queries := db.New(p.db)

	result, err := queries.GetOrganizationByUserId(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	financialYearsResult, err := queries.GetFinancialYearsByUserIdOrganizationId(ctx, db.GetFinancialYearsByUserIdOrganizationIdParams{UserID: userId, OrganizationID: result.OrganizationID})
	if err != nil {
		return nil, err
	}

	financialYears := make([]*financial_year.FinancialYear, len(financialYearsResult))

	for index, value := range financialYearsResult {
		financialYears[index] = &financial_year.FinancialYear{Year: value}
	}

	return &organization.Organization{
		ID:                 result.OrganizationID,
		Name:               result.Name,
		InvoiceNumberStart: result.InvoiceNumberStart,
		FinancialYears:     financialYears,
	}, err
}

func NewPostgres(db *sql.DB) organization.Repository {
	return &postgresRepository{db}
}
