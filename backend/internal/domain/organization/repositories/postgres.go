package organization_repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/organization"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

type postgresRepository struct {
	db *sql.DB
}

// CreateOrganization implements organization.Repository.
func (p *postgresRepository) CreateOrganization(ctx context.Context, userId uuid.UUID, name string, invoiceNumberStart int32) (*organization.Organization, error) {
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

	return &organization.Organization{
		ID:                 result.OrganizationID,
		Name:               result.Name,
		InvoiceNumberStart: result.InvoiceNumberStart,
	}, err
}

func NewPostgres(db *sql.DB) organization.Repository {
	return &postgresRepository{db}
}
