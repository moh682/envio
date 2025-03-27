package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var MigrationFS embed.FS

func Connect(user, password, dbname, host string) (*sql.DB, error) {
	dbPath := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, err
	}

	dirs, err := MigrationFS.ReadDir("migrations")
	if err != nil {
		return nil, err
	}
	fmt.Println("Dirs -- > ", dirs)

	mg, err := NewMigrator(db, MigrationFS)
	if err != nil {
		return nil, err
	}
	err = mg.Up(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database", db)

	return db, nil
}
