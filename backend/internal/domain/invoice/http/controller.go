package invoice_http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/invoice"
)

type Controller interface {
	ListInvoices() http.HandlerFunc
	CreateInvoice() http.HandlerFunc
}

type httpController struct {
	invoiceService invoice.Service
}

func NewHttpController(invoiceService invoice.Service) Controller {
	return &httpController{invoiceService}
}

func (c *httpController) ListInvoices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		invoice, err := c.invoiceService.GetAllInvoicesSince(r.Context(), time.Now().AddDate(-10, 0, 0))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(invoice)
	}
}

func (c *httpController) CreateInvoice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inv CreateInvoiceDTO
		if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if inv.Validate() != nil {
			http.Error(w, inv.Validate().Error(), http.StatusBadRequest)
			return
		}

		products := make([]*invoice.Product, len(inv.Products))
		for i, p := range inv.Products {
			products[i] = &invoice.Product{
				ID:          uuid.New(),
				Description: p.Description,
				Serial:      p.Serial,
				Quantity:    p.Quantity,
				UnitPrice:   p.UnitPrice,
				Total:       p.Total,
			}
		}

		customer := invoice.Customer{
			ID:      uuid.New(),
			Name:    inv.Customer.Name,
			Email:   inv.Customer.Email,
			Phone:   inv.Customer.Phone,
			Address: inv.Customer.Address,
			Car: &invoice.Car{
				Registration: inv.Customer.Car.Registration,
			},
		}

		newInvoice := invoice.Invoice{
			ID:           uuid.New(),
			TotalExclVat: inv.TotalExclVat,
			VatAmount:    inv.VatAmount,
			Status:       invoice.FullyPaid,
			Payments: []*invoice.Payment{
				{
					ID:       uuid.New(),
					Amount:   inv.Total,
					PaidAt:   time.Now(),
					Currency: "DKK",
					Method:   "BANK",
				},
			},
			IssuedAt: inv.IssuedAt,
			VatPct:   .25,
			Total:    inv.Total,
			Customer: customer,
			Products: products,
		}

		err := c.invoiceService.Store(r.Context(), newInvoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newInvoice)
	}
}
