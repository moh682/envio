package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/moh682/envio/backend/internal/api/middlewares"
	"github.com/moh682/envio/backend/internal/domain/invoice"
	invoice_http "github.com/moh682/envio/backend/internal/domain/invoice/http"
	invoice_repositories "github.com/moh682/envio/backend/internal/domain/invoice/repositories"
)

type Server interface {
	ListenAndServe(port int) error
}

type httpServer struct {
	mux *http.ServeMux
}

func NewHttpServer(db *sql.DB) Server {

	invoiceRepository := invoice_repositories.NewPostgres(db)
	invoiceService := invoice.NewService(invoiceRepository)
	invoiceController := invoice_http.NewHttpController(invoiceService)

	// TODO: extract invoice routes into their own function
	invoiceRoutes := http.NewServeMux()
	invoiceRoutes.HandleFunc("GET /invoices", invoiceController.ListInvoices())
	invoiceRoutes.HandleFunc("POST /invoices", invoiceController.CreateInvoice())

	return &httpServer{
		mux: invoiceRoutes,
	}
}

func (s *httpServer) ListenAndServe(port int) error {

	if port == 0 {
		return errors.New("port argument cannot be 0")
	}

	addr := ":" + strconv.Itoa(port)
	server := middlewares.Combine(middlewares.Logger(s.mux), middlewares.Cors((s.mux)))
	return http.ListenAndServe(addr, server)
}
