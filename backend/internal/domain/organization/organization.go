package organization

import (
	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/financial_year"
)

type Organization struct {
	ID                 uuid.UUID                       `json:"id"`
	Name               string                          `json:"name"`
	InvoiceNumberStart int32                           `json:"invoiceNumberStart"`
	FinancialYears     []*financial_year.FinancialYear `json:"financialYears"`
}
