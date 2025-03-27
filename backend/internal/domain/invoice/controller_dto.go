package invoice

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CreateProductDTO struct {
	Name        string  `json:"name" valdiate:"string"`
	Description string  `json:"description" validate:"string"`
	Serial      string  `json:"serial" validate:"string"`
	Quantity    float64 `json:"quantity" validate:"required,float"`
	UnitPrice   float64 `json:"unit_price" validate:"required,float"`
	// including tax
	Total float64 `json:"total" validate:"required,float"`
}

type CreateInvoiceDTO struct {
	IssuedAt     time.Time          `json:"issued_at" validate:"required,datetime"`
	Total        float64            `json:"total" validate:"required,float"`
	TotalExclVat float64            `json:"total_excl_vat" validate:"required,float"`
	VatAmount    float64            `json:"vat_amount" validate:"required,float"`
	Products     []CreateProductDTO `json:"products" validate:"required"`
	CompanyID    uuid.UUID          `json:"company_id" validate:"required,uuid"`
	CustomerID   uuid.UUID          `json:"customer_id" validate:"required,uuid"`
}

func (c *CreateInvoiceDTO) Validate() error {

	if c.IssuedAt.IsZero() {
		return ErrInvalidIssuedAt
	}

	if c.TotalExclVat+c.VatAmount != c.Total {
		return ErrFinancesDoesNotMatch
	}

	if c.TotalExclVat < 0 {
		return ErrInvalidTotalExclVat
	}

	if len(c.Products) == 0 {
		return ErrInvalidProducts
	}

	for _, p := range c.Products {
		if p.Quantity <= 0 {
			return errors.Join(ErrInvalidProducts, errors.New("quantity must be greater than 0"))
		}

		if p.UnitPrice <= 0 {
			return errors.Join(ErrInvalidProducts, errors.New("unit price must be greater than 0"))
		}

		if p.Total <= 0 {
			return errors.Join(ErrInvalidProducts, errors.New("total must be greater than 0"))
		}

		if p.Total != p.UnitPrice*p.Quantity {
			return errors.Join(ErrInvalidProducts, errors.New("total must be equal to unit price * quantity"))
		}
	}

	return nil
}
