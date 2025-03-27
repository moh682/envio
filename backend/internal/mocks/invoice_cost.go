package mocks

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	domain "github.com/moh682/envio/backend/internal/domain"
)

type InvoiceCostMock struct {
	invoiceCost *domain.Cost
}

func (i *InvoiceCostMock) ToInvoiceCost() *domain.Cost {
	return i.invoiceCost
}

func NewInvoiceCost() *InvoiceCostMock {
	cost := &domain.Cost{
		ID:          domain.ID(uuid.New()),
		Description: gofakeit.Sentence(10),
		Total:       float32(1000),
		ProductNr:   fmt.Sprint(gofakeit.Number(10000000, 99999999)),
		Quantity:    float32(gofakeit.Number(1, 10)),
		UnitPrice:   float32(100),
	}

	return &InvoiceCostMock{invoiceCost: cost}
}

func (i *InvoiceCostMock) WithProductNr(productNr string) *InvoiceCostMock {
	i.invoiceCost.ProductNr = productNr
	return i
}

func (i *InvoiceCostMock) WithQuantity(quantity float32) *InvoiceCostMock {
	i.invoiceCost.Quantity = quantity
	return i
}

func (i *InvoiceCostMock) WithUnitPrice(unitPrice float32) *InvoiceCostMock {
	i.invoiceCost.UnitPrice = unitPrice
	return i
}

func (i *InvoiceCostMock) WithTotal(total float32) *InvoiceCostMock {
	i.invoiceCost.Total = total
	return i
}

func (i *InvoiceCostMock) WithDescription(description string) *InvoiceCostMock {
	i.invoiceCost.Description = description
	return i
}

func (i *InvoiceCostMock) WithCost(cost *domain.Cost) *InvoiceCostMock {
	i.invoiceCost = cost
	return i
}

func (i *InvoiceCostMock) ToInvoice() *domain.Cost {
	return i.invoiceCost
}
