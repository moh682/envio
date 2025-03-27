package parsers_test

import (
	"testing"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/services/parsers"
	"github.com/stretchr/testify/assert"
)

func TestHBASInvoiceReader(t *testing.T) {
	t.Run("should parse old xlsx file", func(t *testing.T) {
		hbasParser := parsers.NewHBASInvoiceParser()
		inv, err := hbasParser.Parse("examples/Invoice_100.xlsx")
		if err != nil {
			t.Fatal(err)
		}

		// Invoice
		assert.Equal(t, int32(100), inv.InvoiceNr)
		assert.Equal(t, float32(710), inv.TotalExclVat)
		assert.Equal(t, float32(177.5), inv.VatAmount)
		assert.Equal(t, float32(887.5), inv.Total)

		// Costs
		assert.Equal(t, 3, len(inv.Costs))
		assert.Equal(t, "27 ca11742", inv.Costs[0].ProductNr)
		assert.Equal(t, "luftfilter", inv.Costs[0].Description)
		assert.Equal(t, float32(1), inv.Costs[0].Quantity)
		assert.Equal(t, float32(155*1.25), inv.Costs[0].UnitPrice)
		assert.Equal(t, float32(155*1.25), inv.Costs[0].Total)
		assert.Equal(t, "27 0986452028", inv.Costs[1].ProductNr)
		assert.Equal(t, "oliefilter", inv.Costs[1].Description)
		assert.Equal(t, float32(1), inv.Costs[1].Quantity)
		assert.Equal(t, float32(75*1.25), inv.Costs[1].UnitPrice)
		assert.Equal(t, float32(75*1.25), inv.Costs[1].Total)

		// Payments
		assert.Equal(t, 1, len(inv.Payments))
		assert.Equal(t, float32(887.5), inv.Payments[0].Amount)
		assert.Equal(t, "", inv.Payments[0].Method)
		assert.Equal(t, inv.IssuedAt, inv.Payments[0].PaidAt)

		// Customer
		assert.Equal(t, "", inv.Customer.Name)
		assert.Equal(t, "", inv.Customer.Address)
		assert.Equal(t, "", inv.Customer.ZipCode)
		assert.Equal(t, "AM 84 740", inv.Customer.CarRegistration)
	})

	t.Run("should parse one file", func(t *testing.T) {
		hbasParser := parsers.NewHBASInvoiceParser()
		inv, err := hbasParser.Parse("examples/Invoice_3667.xlsx")
		if err != nil {
			t.Fatal(err)
		}

		// Invoice
		assert.Equal(t, int32(3667), inv.InvoiceNr)
		assert.Equal(t, float32(3680), inv.TotalExclVat)
		assert.Equal(t, float32(920), inv.VatAmount)
		assert.Equal(t, float32(4600), inv.Total)

		// Costs
		assert.Equal(t, 7, len(inv.Costs))
		assert.Equal(t, "FLEDER F", inv.Costs[0].Description)
		assert.Equal(t, float32(1), inv.Costs[0].Quantity)
		assert.Equal(t, float32(810*1.25), inv.Costs[0].UnitPrice)
		assert.Equal(t, float32(810*1.25), inv.Costs[0].Total)
		assert.Equal(t, "MONT", inv.Costs[5].Description)
		assert.Equal(t, float32(3), inv.Costs[5].Quantity)
		assert.Equal(t, float32(500*1.25), inv.Costs[5].UnitPrice)
		assert.Equal(t, float32(1500*1.25), inv.Costs[5].Total)

		// Payments
		assert.Equal(t, 1, len(inv.Payments))
		assert.Equal(t, float32(4600), inv.Payments[0].Amount)
		assert.Equal(t, "KONTANT", inv.Payments[0].Method)
		assert.Equal(t, inv.IssuedAt, inv.Payments[0].PaidAt)

		// Customer
		assert.Equal(t, "Peter Testing", inv.Customer.Name)
		assert.Equal(t, "Testing street 1.2.3", inv.Customer.Address)
		assert.Equal(t, "3000", inv.Customer.ZipCode)
		assert.Equal(t, "AB 12 345", inv.Customer.CarRegistration)
	})

	t.Run("should parse many files", func(t *testing.T) {

		invs, err := parsers.NewHBASInvoiceParser().ParseDir("examples")
		if err != nil {
			t.Fatal(err)
		}

		inv, _ := helper.Find(invs, func(i int, inv *domain.Invoice) bool { return inv.InvoiceNr == 3667 })

		// Invoice
		assert.Equal(t, int32(3667), inv.InvoiceNr)
		assert.Equal(t, float32(3680), inv.TotalExclVat)
		assert.Equal(t, float32(920), inv.VatAmount)
		assert.Equal(t, float32(4600), inv.Total)

		// Costs
		assert.Equal(t, 7, len(inv.Costs))
		assert.Equal(t, "FLEDER F", inv.Costs[0].Description)
		assert.Equal(t, float32(1), inv.Costs[0].Quantity)
		assert.Equal(t, float32(810*1.25), inv.Costs[0].UnitPrice)
		assert.Equal(t, float32(810*1.25), inv.Costs[0].Total)
		assert.Equal(t, "MONT", inv.Costs[5].Description)
		assert.Equal(t, float32(3), inv.Costs[5].Quantity)
		assert.Equal(t, float32(500*1.25), inv.Costs[5].UnitPrice)
		assert.Equal(t, float32(1500*1.25), inv.Costs[5].Total)

		// Payments
		assert.Equal(t, 1, len(inv.Payments))
		assert.Equal(t, float32(4600), inv.Payments[0].Amount)
		assert.Equal(t, "KONTANT", inv.Payments[0].Method)
		assert.Equal(t, inv.IssuedAt, inv.Payments[0].PaidAt)

		// Customer
		assert.Equal(t, "Peter Testing", inv.Customer.Name)
		assert.Equal(t, "Testing street 1.2.3", inv.Customer.Address)
		assert.Equal(t, "3000", inv.Customer.ZipCode)
		assert.Equal(t, "AB 12 345", inv.Customer.CarRegistration)

		inv2, _ := helper.Find(invs, func(i int, inv *domain.Invoice) bool { return inv.InvoiceNr == 8293 })

		// Invoice
		assert.Equal(t, int32(8293), inv2.InvoiceNr)
		assert.Equal(t, float32(3180), inv2.TotalExclVat)
		assert.Equal(t, float32(795), inv2.VatAmount)
		assert.Equal(t, float32(3975), inv2.Total)

		// Costs
		assert.Equal(t, 6, len(inv2.Costs))
		assert.Equal(t, "FLEDER F", inv2.Costs[0].Description)
		assert.Equal(t, float32(1), inv2.Costs[0].Quantity)
		assert.Equal(t, float32(810*1.25), inv2.Costs[0].UnitPrice)
		assert.Equal(t, float32(810*1.25), inv2.Costs[0].Total)
		assert.Equal(t, "MONT", inv2.Costs[5].Description)
		assert.Equal(t, float32(3), inv2.Costs[5].Quantity)
		assert.Equal(t, float32(500*1.25), inv2.Costs[5].UnitPrice)
		assert.Equal(t, float32(1500*1.25), inv2.Costs[5].Total)

		// Payments
		assert.Equal(t, 1, len(inv2.Payments))
		assert.Equal(t, float32(3975), inv2.Payments[0].Amount)
		assert.Equal(t, "KONTANT", inv2.Payments[0].Method)
		assert.Equal(t, inv2.IssuedAt, inv2.Payments[0].PaidAt)

		// Customer
		assert.Equal(t, "Carmen Looney", inv2.Customer.Name)
		assert.Equal(t, "Testing street 1.2.3", inv2.Customer.Address)
		assert.Equal(t, "3000", inv2.Customer.ZipCode)
		assert.Equal(t, "AB 12 345", inv2.Customer.CarRegistration)

	})
}
