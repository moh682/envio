package invoice_http

import (
	"encoding/json"
	"log"
	"net/http"

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

		log.Println("invoices")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("")
	}
}

func (c *httpController) CreateInvoice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("")
	}
}
