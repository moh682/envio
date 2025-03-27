package domain

import "errors"

type PaymentStatus string

const (
	Created       PaymentStatus = "Created"
	Paid          PaymentStatus = "Paid"
	Overdue       PaymentStatus = "Overdue"
	PartiallyPaid PaymentStatus = "PartiallyPaid"
	Refunded      PaymentStatus = "Refunded"
)

var (
	paymentStatuses         = []PaymentStatus{Created, Paid, Overdue, PartiallyPaid, Refunded}
	ErrInvalidPaymentStatus = errors.New("the payment status is invalid")
)

func (status PaymentStatus) Validate() error {
	for _, s := range paymentStatuses {
		if s == status {
			return nil
		}
	}

	return ErrInvalidPaymentStatus
}

func (status PaymentStatus) String() string {
	return string(status)
}
