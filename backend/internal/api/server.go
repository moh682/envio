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
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type Server interface {
	ListenAndServe(port int) error
}

type httpServer struct {
	mux *http.ServeMux
}

func NewHttpServer(db *sql.DB) Server {

	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// We use try.supertokens for demo purposes.
			// At the end of the tutorial we will show you how to create
			// your own SuperTokens core instance and then update your config.
			ConnectionURI: "http://localhost:3567",
			APIKey:        "63b35b7d-2f7b-4e88-b4db-a0b4e8646435",
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "be",
			APIDomain:       "http://localhost:8080",
			WebsiteDomain:   "http://localhost:8080",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}

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
	server := middlewares.Combine(middlewares.Auth(s.mux), middlewares.Logger(s.mux), middlewares.Cors((s.mux)))
	return http.ListenAndServe(addr, server)
}
