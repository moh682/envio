package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/moh682/envio/backend/internal/api/middlewares"
	"github.com/moh682/envio/backend/internal/domain/invoice"
	invoice_repositories "github.com/moh682/envio/backend/internal/domain/invoice/repositories"
)

// Server is the API server

type Server interface {
	ListenAndServe(port int) error
}

type httpServer struct {
	mux *http.ServeMux
}

func NewHttpServer(db *sql.DB) Server {

	invoiceRepository := invoice_repositories.NewPostgres(db)
	invoiceService := invoice.NewService(invoiceRepository)
	invoiceController := invoice.NewHttpController(invoiceService)

	invoiceRoutes := http.NewServeMux()
	invoiceRoutes.HandleFunc("GET /invoices", invoiceController.ListInvoices())
	invoiceRoutes.HandleFunc("POST /invoices", invoiceController.CreateInvoice())

	return &httpServer{
		mux: invoiceRoutes,
	}
}

func (s *httpServer) ListenAndServe(port int) error {
	addr := ":" + strconv.Itoa(port)
	server := middlewares.Combine(middlewares.Logger(s.mux))
	return http.ListenAndServe(addr, server)
}
