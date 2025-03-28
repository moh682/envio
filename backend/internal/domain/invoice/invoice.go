package invoice

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidIssuedAt      = errors.New("invalid issued at")
	ErrInvalidCustomer      = errors.New("invalid customer")
	ErrInvalidProducts      = errors.New("invalid costs")
	ErrInvalidTotalExclVat  = errors.New("invalid total excl vat")
	ErrInvalidPayments      = errors.New("invalid payments")
	ErrFinancesDoesNotMatch = errors.New("finances does not match")
	ErrorCustomerNotFound   = errors.New("customer not found")
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

var (
	// Paid represents a paid payment status
	FullyPaid PaymentStatus = "fully paid"
	// PartiallyPaid represents a partially paid payment status
	PartiallyPaid PaymentStatus = "missing payment"
	// Unpaid represents an unpaid payment status
	UnPaid PaymentStatus = "unpaid"
)

type Invoice struct {
	ID        uuid.UUID  `json:"id"`
	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`

	Number   int64     `json:"number"`
	IssuedAt time.Time `json:"issuedAt"`

	Total        float64       `json:"total"`
	TotalExclVat float64       `json:"totalExclVat"`
	VatPct       float64       `json:"vatPct"`
	VatAmount    float64       `json:"vatAmount"`
	Status       PaymentStatus `json:"status"`

	Customer Customer   `json:"customer"`
	Products []*Product `json:"products"`
	Payments []*Payment `json:"payments"`
}

const (
	DanishVAT = 0.25
)

func NewInvoice(number int64, issuedAt time.Time, customer Customer, products []*Product, payments []*Payment) (*Invoice, error) {

	total := float64(0)
	for _, product := range products {
		total += product.Total
	}

	vatAmount := total * .20
	totalExclVat := total - vatAmount

	if customer.ID == uuid.Nil {
		return nil, ErrorCustomerNotFound
	}

	return &Invoice{

		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,

		Number:   number,
		IssuedAt: issuedAt,

		VatPct:       0.25,
		TotalExclVat: totalExclVat,
		Total:        total,
		VatAmount:    vatAmount,

		Customer: customer,
		Products: products,
		Payments: payments,
	}, nil
}

// func (i *Invoice) Validate() error {
//
// 	if i.customer == nil {
// 		return ErrInvalidCustomer
// 	}
//
// 	if i.products == nil {
// 		return ErrInvalidCosts
// 	}
//
// 	for _, cost := range i.products {
// 		err = cost.Validate()
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	if i.Payments == nil {
// 		return ErrInvalidPayments
// 	}
// 	paymentSum := float32(0)
// 	for _, payment := range i.Payments {
// 		paymentSum += payment.Amount
// 		err = payment.Validate()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	if paymentSum != i.Total {
// 		return fmt.Errorf("invoice.Payment %w: %f != %f", ErrFinancesDoesNotMatch, i.Total, paymentSum)
// 	}
//
// 	err = i.PaymentStatus.Validate()
// 	if err != nil {
// 		return err
// 	}
//
// 	err = i.validateFinances()
// 	if err != nil {
// 		return err
// 	}
//
// 	if i.IssuedAt.IsZero() {
// 		return ErrInvalidIssuedAt
// 	}
//
// 	if i.Customer == nil {
// 		return ErrInvalidCustomer
// 	}
//
// 	if len(i.products) == 0 {
// 		return ErrInvalidCosts
// 	}
//
// 	println(i.TotalExclVat)
//
// 	if i.TotalExclVat == 0 {
// 		return ErrInvalidTotalExclVat
// 	}
//
// 	return nil
// }
//
// func (i *Invoice) validateFinances() error {
// 	var total float32
// 	for _, cost := range i.products {
// 		total += cost.Total
// 	}
//
// 	if math.Round(float64(i.Total)) != math.Round(float64(total)) {
// 		return fmt.Errorf("invoice.Total %w: %f != %f", ErrFinancesDoesNotMatch, i.Total, total)
// 	}
//
// 	vatAmount := total * .20
// 	if math.Round(float64(i.VatAmount)) != math.Round(float64(vatAmount)) {
// 		return fmt.Errorf("invoice.VatAmount %w: %f != %f", ErrFinancesDoesNotMatch, i.VatAmount, vatAmount)
// 	}
//
// 	TotalExclVat := total * .80
// 	if math.Round(float64(i.TotalExclVat)) != math.Round(float64(TotalExclVat)) {
// 		return fmt.Errorf("invoice.TotalExclVat %w: %f != %f", ErrFinancesDoesNotMatch, i.TotalExclVat, TotalExclVat)
// 	}
//
// 	return nil
// }
