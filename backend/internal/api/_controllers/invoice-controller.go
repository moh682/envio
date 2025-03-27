package controllers

import (
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/services"
)

type InvoiceController interface {
	Create(arg NewInvoice) error
	GetAll(limit, offset int32) ([]*domain.Invoice, error)
	GetAllInvoicesSince(since time.Time) ([]*domain.Invoice, error)
	GetYearlyInvoiceCount() ([]*domain.InvoiceCount, error)
	GetDailyStatistics() ([]*domain.InvoiceStatistics, error)
	GetRevenueComparisonWithPreviousMonth() ([]*domain.InvoiceStatistics, error)
	GetInvoiceByID(id string) (*domain.Invoice, error)
	ExportInvoice(id string, path string) error
}

type invoiceController struct {
	invoiceService domain.InvoiceService
	pdfService     services.PDFService
}

func NewInvoiceController(invoiceService domain.InvoiceService, pdfService services.PDFService) InvoiceController {
	return &invoiceController{invoiceService, pdfService}
}

func (c *invoiceController) GetInvoiceByID(id string) (*domain.Invoice, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return c.invoiceService.GetInvoiceByID(domain.ID(parsedID))
}

func (c *invoiceController) GetAll(limit, offset int32) ([]*domain.Invoice, error) {
	return c.invoiceService.GetAll(limit, offset)
}

func (c *invoiceController) GetAllInvoicesSince(since time.Time) ([]*domain.Invoice, error) {
	return c.invoiceService.GetAllInvoicesSince(since)
}

func (c *invoiceController) GetYearlyInvoiceCount() ([]*domain.InvoiceCount, error) {
	return c.invoiceService.GetYearlyInvoiceCount()
}

func (c *invoiceController) GetDailyStatistics() ([]*domain.InvoiceStatistics, error) {
	return c.invoiceService.GetDailyStatistics()
}

func (c *invoiceController) GetRevenueComparisonWithPreviousMonth() ([]*domain.InvoiceStatistics, error) {
	return c.invoiceService.GetRevenueComparisonWithPreviousMonth()
}

func (c *invoiceController) ExportInvoice(id, path string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	invoice, err := c.invoiceService.GetInvoiceByID(domain.ID(parsedID))
	if err != nil {
		return err
	}

	return c.pdfService.GenerateInvoice(path, *invoice)
}

func (c *invoiceController) Create(arg NewInvoice) error {

	parsedIssuedAt, err := helper.ParseFromUTC(arg.IssuedAt)
	if err != nil {
		return err
	}

	customer := &domain.Customer{
		ID:        domain.ID(uuid.New()),
		Name:      arg.CustomerName,
		Email:     arg.CustomerEmail,
		Phone:     arg.CustomerPhone,
		Address:   arg.CustomerAddress,
		ZipCode:   arg.CustomerZipCode,
		CreatedAt: parsedIssuedAt,
	}
	err = customer.Validate()
	if err != nil {
		return err
	}

	costs := make([]*domain.Cost, len(arg.Products))
	for idx, p := range arg.Products {
		cost := &domain.Cost{
			ID:          domain.ID(uuid.New()),
			Description: p.Description,
			ProductNr:   p.ProductNr,
			Quantity:    p.Quantity,
			UnitPrice:   p.UnitPrice,
			Total:       p.Total,
		}
		err := cost.Validate()
		if err != nil {
			return err
		}
		costs[idx] = cost

	}

	paymets := []*domain.Payment{{ID: domain.ID(uuid.New()), Amount: arg.InvoiceTotal, PaidAt: parsedIssuedAt}}

	invoice := domain.Invoice{
		ID:            domain.ID(uuid.New()),
		IssuedAt:      parsedIssuedAt,
		TotalExclVat:  arg.InvoiceTotalExclVat,
		VatAmount:     arg.VatAmount,
		FileCreatedAt: parsedIssuedAt,
		Total:         arg.InvoiceTotal,
		InvoiceNr:     0,
		FilePath:      "",
		PaymentStatus: domain.Paid,
		Currency:      "DKK",
		VatPct:        arg.VatPct,
		Customer:      customer,
		Costs:         costs,
		Payments:      paymets,
	}
	err = invoice.Validate()
	if err != nil {
		return err
	}

	return c.invoiceService.Store(invoice)
}
