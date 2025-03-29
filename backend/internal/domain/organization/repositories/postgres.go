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

func NewPostgres(db *sql.DB) organization.Repository {
	return &postgresRepository{db}
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
		ID: result.OrganizationID,
	}, err
}
