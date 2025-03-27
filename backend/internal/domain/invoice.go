package domain

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidIssuedAt      = errors.New("invalid issued at")
	ErrInvalidCustomer      = errors.New("invalid customer")
	ErrInvalidCosts         = errors.New("invalid costs")
	ErrInvalidTotalExclVat  = errors.New("invalid total excl vat")
	ErrInvalidPayments      = errors.New("invalid payments")
	ErrFinancesDoesNotMatch = errors.New("finances does not match")
)

type Invoice struct {
	ID            ID            `json:"id"`
	InvoiceNr     int32         `json:"invoice_nr"`
	CompanyName   string        `json:"company_name"`
	FilePath      string        `json:"file_path"`
	FileCreatedAt time.Time     `json:"file_created_at"`
	IssuedAt      time.Time     `json:"issued_at"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	Currency      string        `json:"currency"`
	Total         float32       `json:"total"`
	TotalExclVat  float32       `json:"total_excl_vat"`
	VatPct        float32       `json:"vat_pct"`
	VatAmount     float32       `json:"vat_amount"`

	Customer *Customer  `json:"customer"`
	Costs    []*Cost    `json:"costs"`
	Payments []*Payment `json:"payments"`
}

func (i *Invoice) Validate() error {

	_, err := uuid.Parse(i.ID.String())
	if err != nil {
		return err
	}

	if i.Customer == nil {
		return ErrInvalidCustomer
	}

	if i.Costs == nil {
		return ErrInvalidCosts
	}

	for _, cost := range i.Costs {
		err = cost.Validate()
		if err != nil {
			return err
		}
	}

	if i.Payments == nil {
		return ErrInvalidPayments
	}
	paymentSum := float32(0)
	for _, payment := range i.Payments {
		paymentSum += payment.Amount
		err = payment.Validate()
		if err != nil {
			return err
		}
	}
	if paymentSum != i.Total {
		return fmt.Errorf("invoice.Payment %w: %f != %f", ErrFinancesDoesNotMatch, i.Total, paymentSum)
	}

	err = i.PaymentStatus.Validate()
	if err != nil {
		return err
	}

	err = i.validateFinances()
	if err != nil {
		return err
	}

	if i.IssuedAt.IsZero() {
		return ErrInvalidIssuedAt
	}

	if i.Customer == nil {
		return ErrInvalidCustomer
	}

	if len(i.Costs) == 0 {
		return ErrInvalidCosts
	}

	println(i.TotalExclVat)

	if i.TotalExclVat == 0 {
		return ErrInvalidTotalExclVat
	}

	return nil

}

func (i *Invoice) validateFinances() error {
	var total float32
	for _, cost := range i.Costs {
		total += cost.Total
	}

	if math.Round(float64(i.Total)) != math.Round(float64(total)) {
		return fmt.Errorf("invoice.Total %w: %f != %f", ErrFinancesDoesNotMatch, i.Total, total)
	}

	vatAmount := total * .20
	if math.Round(float64(i.VatAmount)) != math.Round(float64(vatAmount)) {
		return fmt.Errorf("invoice.VatAmount %w: %f != %f", ErrFinancesDoesNotMatch, i.VatAmount, vatAmount)
	}

	TotalExclVat := total * .80
	if math.Round(float64(i.TotalExclVat)) != math.Round(float64(TotalExclVat)) {
		return fmt.Errorf("invoice.TotalExclVat %w: %f != %f", ErrFinancesDoesNotMatch, i.TotalExclVat, TotalExclVat)
	}

	return nil
}
