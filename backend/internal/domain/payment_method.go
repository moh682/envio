package domain

import "errors"

type PaymentMethod int

// TODO: change to a string enum instead of int
const (
	UNKNOWN PaymentMethod = iota
	CASH
	MOBILEPAY
	BANK_TRANSFER
)

var (
	paymentMethods          = []string{"UNKNOWN", "CASH", "MOBILEPAY", "BANK_TRANSFER"}
	ErrInvalidPaymentMethod = errors.New("invalid payment status")
)

// Initially starts with status {Created}
func NewPaymentMethod(status int) (PaymentMethod, error) {
	switch {
	case status > len(paymentMethods):
		return 0, ErrInvalidPaymentMethod
	case status < 0:
		return 0, ErrInvalidPaymentMethod
	}
	return PaymentMethod(status), nil
}

func (index PaymentMethod) String() string {
	return paymentMethods[index]
}

func (index PaymentMethod) Index() int {
	return int(index)
}
