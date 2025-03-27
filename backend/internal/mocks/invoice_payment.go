package mocks

import (
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
)

type PaymentMock struct {
	payment *domain.Payment
}

func NewInvoicePayment() *PaymentMock {
	payment := &domain.Payment{
		ID:     domain.ID(uuid.New()),
		Method: "Card",
		Amount: 1000,
		PaidAt: time.Now(),
	}

	return &PaymentMock{payment: payment}
}

func (p *PaymentMock) ToPayment() *domain.Payment {
	return p.payment
}

func (p *PaymentMock) WithMethod(method string) *PaymentMock {
	p.payment.Method = method
	return p
}

func (p *PaymentMock) WithAmount(amount float32) *PaymentMock {
	p.payment.Amount = amount
	return p
}

func (p *PaymentMock) WithPaidAt(paidAt time.Time) *PaymentMock {
	p.payment.PaidAt = paidAt
	return p
}

func (p *PaymentMock) WithID(id domain.ID) *PaymentMock {
	p.payment.ID = id
	return p
}
