package domain

import (
	"time"

	"github.com/google/uuid"
)

type Migration struct {
	ID                      ID        `json:"id"`
	InvoiceNr               int32     `json:"invoice_nr"`
	FailedAt                time.Time `json:"failed_at"`
	FilePath                string    `json:"file_path"`
	HasMigratedSuccessfully bool      `json:"has_migrated_successfully"`
	LastFailedAt            time.Time `json:"last_failed_at"`
}

func (m *Migration) Validate() error {
	_, err := uuid.Parse(m.ID.String())
	if err != nil {
		return err
	}

	return nil
}
