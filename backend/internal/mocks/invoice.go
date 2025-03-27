package mocks

import (
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"

	"github.com/brianvoe/gofakeit/v7"
)

func NewInvoice() *InvoiceMock {

	invID := uuid.New()
	inv := &domain.Invoice{
		ID:            domain.ID(invID),
		InvoiceNr:     int32(gofakeit.Number(10000000, 99999999)),
		FilePath:      "/path/to/file",
		FileCreatedAt: gofakeit.Date(),
		IssuedAt:      gofakeit.Date(),
		PaymentStatus: domain.Paid,
		Currency:      "DKK",
		Total:         1000,
		TotalExclVat:  800,
		VatPct:        .25,
		VatAmount:     200,
		Customer:      NewCustomer().ToCustomer(),
		Costs:         []*domain.Cost{NewInvoiceCost().ToInvoiceCost()},
		Payments:      []*domain.Payment{{ID: domain.ID(uuid.New()), Method: "Card", Amount: 1000, PaidAt: time.Now()}},
	}
	return &InvoiceMock{invoice: inv}
}

type InvoiceMock struct {
	invoice *domain.Invoice
}

func (i *InvoiceMock) ToInvoice() *domain.Invoice {
	return i.invoice

}

func (i *InvoiceMock) WithInvoiceNr(num int32) *InvoiceMock {
	i.invoice.InvoiceNr = num
	return i
}

func (i *InvoiceMock) WithVatAmount(amount float32) *InvoiceMock {
	i.invoice.VatAmount = amount
	return i
}

func (i *InvoiceMock) WithCustomer(customer *domain.Customer) *InvoiceMock {
	i.invoice.Customer = customer
	return i
}
func (i *InvoiceMock) WithPayments(payments []*domain.Payment) *InvoiceMock {
	i.invoice.Payments = payments
	return i
}

func (i *InvoiceMock) WithCost(cost *domain.Cost) *InvoiceMock {
	i.invoice.Costs = append(i.invoice.Costs, cost)
	return i
}

func (i *InvoiceMock) WithCosts(costs []*domain.Cost) *InvoiceMock {
	i.invoice.Costs = costs
	return i
}

func (i *InvoiceMock) WithPaymentStatus(status domain.PaymentStatus) *InvoiceMock {
	i.invoice.PaymentStatus = status
	return i
}

func (i *InvoiceMock) WithFilePath(path string) *InvoiceMock {
	i.invoice.FilePath = path
	return i
}

func (i *InvoiceMock) WithFileCreatedAt(date time.Time) *InvoiceMock {
	i.invoice.FileCreatedAt = date
	return i
}

func (i *InvoiceMock) WithIssuedAt(date time.Time) *InvoiceMock {
	i.invoice.IssuedAt = date
	return i
}

func (i *InvoiceMock) WithCurrency(currency string) *InvoiceMock {
	i.invoice.Currency = currency
	return i
}

func (i *InvoiceMock) WithTotal(total float32) *InvoiceMock {
	i.invoice.Total = total
	return i
}

func (i *InvoiceMock) WithTotalExclVat(total float32) *InvoiceMock {
	i.invoice.TotalExclVat = total
	return i
}

func (i *InvoiceMock) WithVatPct(pct float32) *InvoiceMock {
	i.invoice.VatPct = pct
	return i
}
