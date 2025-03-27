package services_test

import (
	"os"
	"testing"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/mocks"
	"github.com/moh682/envio/backend/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestGenerateInvoice(t *testing.T) {

	t.Run("Test Generate Invoice", func(t *testing.T) {

		defer func() {
			err := os.Remove("invoice_1.pdf")
			if err != nil {
				t.Fatal(err)
			}
		}()

		inv := mocks.NewInvoice().WithInvoiceNr(1).WithCosts(
			[]*domain.Cost{
				mocks.NewInvoiceCost().ToInvoice(),
				mocks.NewInvoiceCost().ToInvoice(),
				mocks.NewInvoiceCost().ToInvoice(),
			}).ToInvoice()
		err := services.NewPDFService("dk").GenerateInvoice("./", *inv)
		assert.Nil(t, err)
	})
}
