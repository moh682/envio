package invoice_http

import (
	"errors"
	"time"

	"github.com/moh682/envio/backend/internal/domain/invoice"
)

type CreateProductDTO struct {
	Description string  `json:"description" validate:"string"`
	Serial      string  `json:"serial" validate:"string"`
	Quantity    float64 `json:"quantity" validate:"required,float"`
	UnitPrice   float64 `json:"unit_price" validate:"required,float"`
	// including tax
	Total float64 `json:"total" validate:"required,float"`
}

type CreateCarDTO struct {
	Registration string `json:"registration" validate:"required,string"`
}

type CreateCustomerDTO struct {
	Name    string       `json:"name" validate:"string"`
	Email   string       `json:"email" validate:"email"`
	Phone   string       `json:"phone" validate:"string"`
	Car     CreateCarDTO `json:"car"`
	Address string       `json:"address" validate:"string"`
}

type CreateInvoiceDTO struct {
	IssuedAt     time.Time          `json:"issued_at" validate:"required,datetime"`
	Total        float64            `json:"total" validate:"required,float"`
	TotalExclVat float64            `json:"total_excl_vat" validate:"required,float"`
	VatAmount    float64            `json:"vat_amount" validate:"required,float"`
	Products     []CreateProductDTO `json:"products" validate:"required"`
	Customer     CreateCustomerDTO  `json:"customer" validate:"required"`
}

func (customer *CreateCustomerDTO) Validate() error {

	if customer.Car.Registration == "" && customer.Name == "" && customer.Email == "" && customer.Phone == "" && customer.Address == "" {
		return errors.New("customer cannot have all properties empty at the same time")
	}

	return nil
}

func (c *CreateInvoiceDTO) Validate() error {

	if c.IssuedAt.IsZero() {
		return errors.New("issued_at cannot be empty")
	}

	if c.TotalExclVat+c.VatAmount != c.Total {
		return errors.New("total must be equal to total excl vat + vat amount")
	}

	if c.TotalExclVat <= 0 {
		return errors.New("total excl vat must be greater than 0")
	}

	if len(c.Products) == 0 {
		return errors.New("products cannot be empty")
	}

	for _, p := range c.Products {
		if p.Quantity <= 0 {
			return errors.Join(invoice.ErrInvalidProducts, errors.New("quantity must be greater than 0"))
		}

		if p.UnitPrice <= 0 {
			return errors.Join(invoice.ErrInvalidProducts, errors.New("unit price must be greater than 0"))
		}

		if p.Total <= 0 {
			return errors.Join(invoice.ErrInvalidProducts, errors.New("total must be greater than 0"))
		}

		if p.Total != p.UnitPrice*p.Quantity {
			return errors.Join(invoice.ErrInvalidProducts, errors.New("total must be equal to unit price * quantity"))
		}
	}

	if err := c.Customer.Validate(); err != nil {
		return err
	}

	return nil
}
