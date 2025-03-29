package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migrator interface {
	Up(ctx context.Context) error
}

type migrator struct {
	db          *sql.DB
	migrationFS embed.FS
}

func NewMigrator(db *sql.DB, migrationFS embed.FS) (Migrator, error) {

	return &migrator{
		db,
		migrationFS,
	}, nil
}

func (m *migrator) Up(ctx context.Context) error {

	driver, err := postgres.WithInstance(m.db, &postgres.Config{})
	if err != nil {
		return err
	}

	sourceDriver, err := iofs.New(MigrationFS, "migrations")
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("iofs", sourceDriver, "sqlite3", driver)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("No change in migration")
			return nil
		}

		version, isDirty, err := migration.Version()
		if err != nil {
			return err
		}

		if isDirty {
			err = migration.Force(int(version))
			if err != nil {
				return err
			}
		}

	}

	return nil
}
