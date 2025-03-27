package services

import (
	"github.com/moh682/envio/backend/internal/domain"
)

type InvoiceParser interface {
	Parse(path string) (*domain.Invoice, error)
	ParseDir(dirPath string) ([]*domain.Invoice, error)
}

type InvoiceImporterService interface {
	Import(path string) error
	ImportDir(dirPath string) error
}

type InvoiceImporter struct {
	parser InvoiceParser
	invSvc domain.InvoiceService
}

func NewInvoiceImporterService(invSvc domain.InvoiceService, parser InvoiceParser) InvoiceImporterService {
	return &InvoiceImporter{
		parser: parser,
		invSvc: invSvc,
	}
}

func (s *InvoiceImporter) Import(path string) error {
	inv, err := s.parser.Parse(path)
	if err != nil {
		return err
	}

	return s.invSvc.Store(*inv)
}

func (s *InvoiceImporter) ImportDir(dirPath string) error {
	invs, err := s.parser.ParseDir(dirPath)
	if err != nil {
		return err
	}

	nonPointer := make([]domain.Invoice, len(invs))
	for i, inv := range invs {
		nonPointer[i] = *inv
	}

	return s.invSvc.StoreMany(nonPointer)
}
