package invoice

import (
	"errors"
	"time"
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
	// Deprecated: not implemented
	UpdatedAt time.Time `json:"updatedAt"`
	// Deprecated: not implemented
	CreatedAt time.Time `json:"createdAt"`
	// Deprecated: not implemented
	DeletedAt *time.Time `json:"deletedAt"`

	Number   int64     `json:"number"`
	IssuedAt time.Time `json:"issuedAt"`
	Total    float64   `json:"total"`
	IsVat    bool      `json:"isVat"`
	Products []*Product
}

const (
	DanishVAT = 0.25
)
