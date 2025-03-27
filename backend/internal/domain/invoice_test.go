package domain_test

import (
	"testing"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceValidation(t *testing.T) {
	type Test struct {
		name    string
		inv     *domain.Invoice
		wantErr bool
		err     error
	}

	tests := []Test{
		{
			name: "Should pass when the sum of costs is equal to the total",
			inv: mocks.NewInvoice().WithCosts(
				[]*domain.Cost{
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
				}).WithPayments([]*domain.Payment{mocks.NewInvoicePayment().WithAmount(3000).ToPayment()}).
				WithVatAmount(600).WithTotalExclVat(2400).WithTotal(3000).ToInvoice(),
			wantErr: false,
			err:     nil,
		},
		{
			name: "Should fail when the vat Amount is not equal the calculated vat amount",
			inv: mocks.NewInvoice().WithCosts(
				[]*domain.Cost{
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
				}).WithPayments([]*domain.Payment{mocks.NewInvoicePayment().WithAmount(3000).ToPayment()}).
				WithVatAmount(900).WithTotalExclVat(2400).WithTotal(3000).ToInvoice(),
			wantErr: true,
			err:     domain.ErrFinancesDoesNotMatch,
		},
		{
			name: "Should fail when the totalExcludingVat is not equal the calculated totalExcludingVat",
			inv: mocks.NewInvoice().WithCosts(
				[]*domain.Cost{
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
				}).WithPayments([]*domain.Payment{mocks.NewInvoicePayment().WithAmount(3000).ToPayment()}).
				WithVatAmount(600).WithTotalExclVat(2000).WithTotal(3000).ToInvoice(),
			wantErr: true,
			err:     domain.ErrFinancesDoesNotMatch,
		},
		{
			name: "Should fail when the total is not equal the calculated total",
			inv: mocks.NewInvoice().WithCosts(
				[]*domain.Cost{
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
				}).WithPayments([]*domain.Payment{mocks.NewInvoicePayment().WithAmount(3000).ToPayment()}).
				WithVatAmount(600).WithTotalExclVat(2400).WithTotal(4000).ToInvoice(),
			wantErr: true,
			err:     domain.ErrFinancesDoesNotMatch,
		},
		{
			name: "Should fail when the Payment Amount is not equal the Total",
			inv: mocks.NewInvoice().WithCosts(
				[]*domain.Cost{
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
					mocks.NewInvoiceCost().WithQuantity(2).WithUnitPrice(500).WithTotal(1000).ToInvoice(),
				}).WithPayments([]*domain.Payment{mocks.NewInvoicePayment().WithAmount(2000).ToPayment()}).
				WithVatAmount(600).WithTotalExclVat(2400).WithTotal(3000).ToInvoice(),
			wantErr: true,
			err:     domain.ErrFinancesDoesNotMatch,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.inv.Validate()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Test %d: expected error, got nil", i)
				}
				assert.ErrorAs(t, err, &tt.err)
			} else {
				if err != nil {
					t.Errorf("Test %d: unexpected error: %v", i, err)
				}
			}
		})

	}

}
