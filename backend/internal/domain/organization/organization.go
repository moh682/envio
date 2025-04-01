package organization

import "github.com/google/uuid"

type Organization struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	InvoiceNumberStart int32     `json:"invoiceNumberStart"`
}
