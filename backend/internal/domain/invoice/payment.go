package invoice

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PaymentMethod string

const (
	// Cash represents a cash payment method
	Cash PaymentMethod = "cash"
	// Card represents a card payment method
	Card PaymentMethod = "card"
	// Bank represents a bank payment method
	Bank PaymentMethod = "bank"
	// MobilePay represents a mobile payment method
	MobilePay PaymentMethod = "mobilepay"
)

type Currency string

const (
	// DKK represents the Danish Krone currency
	DKK Currency = "DKK"
)

var (
	ErrInvalidPaymentAmount = errors.New("Invalid payment amount")
)

type Payment struct {
	ID       uuid.UUID
	Amount   float64
	PaidAt   time.Time
	Currency Currency
	Method   PaymentMethod
}

func NewPayment(id uuid.UUID, amount float64, paidAt time.Time, currency Currency, method PaymentMethod) (*Payment, error) {

	if amount <= 0 {
		return nil, errors.Join(errors.New("Payment amount must be greater than 0"), ErrInvalidPaymentAmount)
	}

	return &Payment{
		ID:       id,
		Amount:   amount,
		PaidAt:   paidAt,
		Currency: currency,
		Method:   method,
	}, nil
}
