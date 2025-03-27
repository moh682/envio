package parsers

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/services"
	"github.com/xuri/excelize/v2"
)

type hbasInvoiceParser struct{}

var (
	ErrNonExistingDate = errors.New("could not convert date to time")
	ErrInvalidSyntax   = errors.New("invalid syntax")
)

// Parses only xlsx files
func NewHBASInvoiceParser() services.InvoiceParser {
	return &hbasInvoiceParser{}
}

func parseToInt(value string, err error) (int, error) {
	value = strings.TrimSpace(value)
	if err != nil {
		return 0, err
	}

	if value == "" {
		return 0, nil
	}

	if value == "#VALUE!" {
		return 0, ErrInvalidSyntax
	}

	re := regexp.MustCompile(`\D+`)

	// Replace all non-digit characters with an empty string
	value = re.ReplaceAllString(value, "")

	return strconv.Atoi(value)
}

func parseFloat32(value string) (float32, error) {

	value = strings.TrimSpace(value)

	if value == "" {
		return 0, nil
	}

	if value[0] == ',' {
		value = value[1:]
	}
	value = strings.Replace(value, ",", ".", -1)

	if value == "#VALUE!" {
		return 0, ErrInvalidSyntax
	}

	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}

	return float32(f), nil
}

func parseTofloat32(value string, err error) (float32, error) {
	if err != nil {
		return 0, err
	}

	if value == "#VALUE!" {
		return 0, ErrInvalidSyntax
	}

	return parseFloat32(value)

}

func parseToDate(value string, err error) (time.Time, error) {
	if err != nil {
		return time.Time{}, err
	}

	if strings.TrimSpace(value) == "" {
		return time.Time{}, ErrNonExistingDate
	}

	formats := []string{"02/01/2006", "02-01-2006", "02.01.2006", "02.01.2006", "02-01-2006", "02/01/2006"}

	for _, format := range formats {
		t, err := time.Parse(format, value)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, ErrNonExistingDate

}

func (h *hbasInvoiceParser) Parse(path string) (*domain.Invoice, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	sheet := f.GetSheetList()[0]

	invoiceNumber, err := parseToInt(f.GetCellValue(sheet, "B7"))
	if err != nil {
		return nil, err
	}

	issuedAt, err := parseToDate(f.GetCellValue(sheet, "B18"))
	if err != nil {
		return nil, err
	}

	customer, err := parseCustomer(f, sheet, issuedAt)
	if err != nil {
		return nil, err
	}

	paymentMethod := ""
	costs := make([]*domain.Cost, 0)
	for i := range 15 {
		y := 23 + i

		cost, err := parseCost(f, sheet, y)
		if err != nil {
			return nil, err
		}

		if cost.ProductNr == "" && cost.Description == "" && cost.Quantity == 0 && cost.UnitPrice == 0 && cost.Total == 0 {
			continue
		}

		if cost.ProductNr == "" && cost.Description != "" && cost.Quantity == 0 && cost.UnitPrice == 0 && cost.Total == 0 {
			paymentMethod = cost.Description
			continue
		}

		costs = append(costs, cost)
	}

	totalExclVat, totalExclVatErr := parseTofloat32(f.GetCellValue(sheet, "A40"))
	vatAmount, vatAmountErr := parseTofloat32(f.GetCellValue(sheet, "B40"))
	total, totalErr := parseTofloat32(f.GetCellValue(sheet, "C40"))
	if totalErr != nil || totalExclVatErr != nil || vatAmountErr != nil {
		if totalErr == ErrInvalidSyntax || totalExclVatErr == ErrInvalidSyntax || vatAmountErr == ErrInvalidSyntax {
			total, totalExclVat, vatAmount = calculateTotal(costs)
		} else {
			if totalErr != nil {
				return nil, totalErr
			}
			if totalExclVatErr != nil {
				return nil, totalExclVatErr
			}
			if vatAmountErr != nil {
				return nil, vatAmountErr
			}
		}
	}

	payments := []*domain.Payment{
		{
			ID:     domain.ID(uuid.New()),
			PaidAt: issuedAt,
			Amount: total,
			Method: paymentMethod,
		}}

	inv := &domain.Invoice{
		FilePath:      path,
		FileCreatedAt: issuedAt,
		ID:            domain.ID(uuid.New()),
		InvoiceNr:     int32(invoiceNumber),
		IssuedAt:      issuedAt,
		TotalExclVat:  totalExclVat,
		VatAmount:     vatAmount,
		Total:         total,
		VatPct:        0.25,
		PaymentStatus: domain.Paid,

		Payments: payments,
		Customer: customer,
		Costs:    costs,
	}

	err = inv.Validate()
	if err != nil {
		return nil, err
	}

	return inv, nil
}

func ContainsAny(s string, substrs []string) bool {
	for _, sub := range substrs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

func (h *hbasInvoiceParser) ParseDir(dirPath string) ([]*domain.Invoice, error) {

	invs := make([]*domain.Invoice, 0)

	root := os.DirFS(dirPath)

	fileNames, err := fs.Glob(root, "*.xlsx")
	if err != nil {
		return invs, err
	}

	paths := make([]string, len(fileNames))
	for i, name := range fileNames {
		paths[i] = fmt.Sprintf("%s/%s", dirPath, name)
	}

	ignoredFiles := []string{"2992 sælger bil.xlsx", "391 JF 25 445.xlsx"}

	for _, file := range paths {

		if ContainsAny(file, ignoredFiles) {
			continue
		}

		inv, err := h.Parse(file)
		log.Printf("Parsing file: %s", file)

		if err != nil {

			if err == domain.ErrInvalidTotalExclVat || err == ErrNonExistingDate {
				log.Println("Skipping file: %s", file)
				continue
			}

			log.Println("Error parsing file: %s", file)
			log.Println(err)
			return invs, err
		}

		invs = append(invs, inv)
	}

	return invs, nil

}

func parseCost(f *excelize.File, sheet string, row int) (*domain.Cost, error) {
	serial, _ := f.GetCellValue(sheet, fmt.Sprintf("A%d", row))
	description, _ := f.GetCellValue(sheet, fmt.Sprintf("B%d", row))
	quantity, err := parseTofloat32(f.GetCellValue(sheet, fmt.Sprintf("C%d", row)))
	if err != nil && err != ErrInvalidSyntax {
		return nil, err
	}
	rate, err := parseTofloat32(f.GetCellValue(sheet, fmt.Sprintf("D%d", row)))
	if err != nil && err != ErrInvalidSyntax {
		return nil, err
	}
	total, err := parseTofloat32(f.GetCellValue(sheet, fmt.Sprintf("E%d", row)))
	if err != nil && err != ErrInvalidSyntax {
		return nil, err
	}

	return &domain.Cost{
		ID:          domain.ID(uuid.New()),
		ProductNr:   serial,
		Description: description,
		Quantity:    quantity,
		UnitPrice:   rate * 1.25,
		Total:       total * 1.25,
	}, nil
}

func calculateTotal(costs []*domain.Cost) (float32, float32, float32) {
	var total float32
	var totalExclVat float32
	var vatAmount float32
	for _, c := range costs {
		total += c.Total
	}

	vatAmount = total * 0.2
	totalExclVat = total - vatAmount
	return total, totalExclVat, vatAmount
}

func parseCustomer(f *excelize.File, sheet string, issuedAt time.Time) (*domain.Customer, error) {

	name, _ := f.GetCellValue(sheet, "B9")
	address, _ := f.GetCellValue(sheet, "B10")
	zip, _ := f.GetCellValue(sheet, "B11")
	carReg, _ := f.GetCellValue(sheet, "B12")
	phone, _ := f.GetCellValue(sheet, "B13")
	email, _ := f.GetCellValue(sheet, "B14")

	return &domain.Customer{
		ID:              domain.ID(uuid.New()),
		Name:            name,
		Address:         address,
		Phone:           phone,
		Email:           email,
		CarRegistration: carReg,
		ZipCode:         zip,
		CreatedAt:       issuedAt,
	}, nil
}
