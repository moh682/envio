package services

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"strconv"

	invoice "github.com/moh682/envio/backend/internal/domain"
	"github.com/signintech/gopdf"
)

var currencySymbols = map[string]string{
	"DKK": "DKK ",
}

//go:embed "Inter/Inter Variable/Inter.ttf"
var interFont []byte

//go:embed "image.png"
var logoPng []byte

//go:embed "Inter/Inter Hinted for Windows/Desktop/Inter-Bold.ttf"
var interBoldFont []byte

type LabelTranslation struct {
	InvoiceTitleLabel    string
	SubtotalLabel        string
	TaxLabel             string
	TotalLabel           string
	CustomerNameLabel    string
	AddressLabel         string
	PhoneLabel           string
	CarRegistartionLabel string
	EmailLabel           string
	InvoiceNumberLabel   string
	IssuedAtLabel        string
	BankReg              string
	BankAccount          string
	ItemNumberLabel      string
	ItemDescriptionLabel string
	ItemQuantityLabel    string
	ItemUnitPriceLabel   string
	ItemTotalLabel       string
	FromLabel            string
	ToLabel              string
}

type Language string

const (
	descriptionColumnOffset = 100
	quantityColumnOffset    = 360
	rateColumnOffset        = 405
	amountColumnOffset      = 480

	billedToSectionOffset = 360
)

const (
	currency = "DKK"
)

const (
	SubtotalLabel = "Subtotal"
	DiscountLabel = "Discount"
	TaxLabel      = "Tax"
	TotalLabel    = "Total"
)

var Translations = map[string]LabelTranslation{
	"dk": {
		InvoiceTitleLabel: "Faktura",

		FromLabel:            "Fra",
		ToLabel:              "Til",
		CustomerNameLabel:    "Kunde Navn",
		AddressLabel:         "Adresse",
		CarRegistartionLabel: "Bilens reg. nr.",
		PhoneLabel:           "Tlf.",
		EmailLabel:           "Email",
		InvoiceNumberLabel:   "Faktura nr.",
		IssuedAtLabel:        "Dato",
		BankReg:              "Reg. Nr",
		BankAccount:          "Konto Nr",

		// Table Headers Section
		ItemNumberLabel:      "Nr.",
		ItemDescriptionLabel: "Beskrivelse",
		ItemQuantityLabel:    "Antal",
		ItemUnitPriceLabel:   "Pris pr. enhed",
		ItemTotalLabel:       "Total",

		// Sum Section
		SubtotalLabel: "Netto værdi",
		TaxLabel:      "Moms 25%",
		TotalLabel:    "I alt kroner",
	},
}

type pdfService struct {
	Langauge      Language
	Translations  LabelTranslation
	ConfigService config.ConfigService
}

type PDFService interface {
	GenerateInvoice(outputPath string, invoice invoice.Invoice) error
}

func NewPDFService(language string, configService config.ConfigService) PDFService {
	return &pdfService{Langauge: Language(language), ConfigService: configService, Translations: Translations[language]}
}

func (p *pdfService) GenerateInvoice(folderPath string, invoice invoice.Invoice) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()
	err := pdf.AddTTFFontData("Inter", interFont)
	if err != nil {
		return err
	}

	err = pdf.AddTTFFontData("Inter-Bold", interBoldFont)
	if err != nil {
		return err
	}

	// company name should be stored in the database
	writeTitle(&pdf, p.Translations.InvoiceTitleLabel, string(invoice.InvoiceNr), invoice.IssuedAt.Format("02 Jan 2006"))

	o := p.ConfigService.GetOrganisation()
	// Organisation should be stored in the database
	org := Organisation{
		Name:        o.Name,
		Address:     o.Address.Street,
		ZipCode:     o.Address.Zip,
		City:        o.Address.City,
		Cvr:         o.Cvr,
		Email:       o.Email,
		Phone:       o.Phone,
		BankReg:     o.Bank.RegNr,
		BankAccount: o.Bank.AccountNr,
		BankName:    o.Bank.Name,
	}
	err = p.writeFromAndTo(&pdf, org, *invoice.Customer)
	if err != nil {
		return err
	}

	p.writeHeaderRow(&pdf)
	for _, cost := range invoice.Costs {
		writeRow(&pdf, cost.ProductNr, cost.Description, int(cost.Quantity), float64(cost.UnitPrice))
	}

	p.writeTotals(&pdf, float64(invoice.TotalExclVat), float64(invoice.VatAmount), float64(invoice.Total))

	footerText := fmt.Sprintf("%s  ·  %s: %s  ·  %s: %s", org.BankName, p.Translations.BankReg, org.BankReg, p.Translations.BankAccount, org.BankAccount)
	writeFooter(&pdf, footerText)

	outputPath := fmt.Sprintf("%s/invoice_%d.pdf", folderPath, invoice.InvoiceNr)
	return pdf.WritePdf(outputPath)

}

type Organisation struct {
	Name        string
	Address     string
	ZipCode     string
	City        string
	Cvr         string
	Phone       string
	Email       string
	BankReg     string
	BankAccount string
	BankName    string
}

func getLogoImage() (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(logoPng))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (s *pdfService) writeFromAndTo(pdf *gopdf.GoPdf, org Organisation, cust invoice.Customer) error {
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont("Inter", "", 9)
	pdf.Br(18)
	pdf.SetTextColor(75, 75, 75)

	logo, err := getLogoImage()
	if err != nil {
		return err
	}

	width := logo.Bounds().Dx()
	height := logo.Bounds().Dy()
	scaledWidth := 100.0
	scaledHeight := float64(height) * scaledWidth / float64(width)
	pdf.SetX(quantityColumnOffset - 10)
	err = pdf.ImageFrom(logo, pdf.GetX(), pdf.GetY(), &gopdf.Rect{W: scaledWidth, H: scaledHeight})
	if err != nil {
		return err
	}
	pdf.Br(scaledHeight + 24)

	_ = pdf.SetFont("Inter", "", 10)
	_ = pdf.Cell(nil, s.Translations.ToLabel)
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, s.Translations.FromLabel, gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(20)

	_ = pdf.Cell(nil, fmt.Sprintf("%s", cust.Name))
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, org.Name, gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(15)

	_ = pdf.Cell(nil, fmt.Sprintf("%s", cust.Address))
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, org.Address, gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(15)

	_ = pdf.Cell(nil, fmt.Sprintf("%s", cust.ZipCode))
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, fmt.Sprintf("%s %s", org.ZipCode, org.City), gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(15)

	_ = pdf.Cell(nil, fmt.Sprintf("%s", cust.CarRegistration))
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, org.Cvr, gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(15)

	_ = pdf.Cell(nil, fmt.Sprintf("%s", cust.Phone))
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, org.Phone, gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(15)

	_ = pdf.Cell(nil, fmt.Sprintf("%s", cust.Email))
	pdf.SetX(billedToSectionOffset)
	_ = pdf.CellWithOption(nil, org.Email, gopdf.CellOption{Align: gopdf.Right})
	pdf.Br(64)

	return nil

}

func writeTitle(pdf *gopdf.GoPdf, title, id, date string) {
	_ = pdf.SetFont("Inter-Bold", "", 24)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Cell(nil, title)
	pdf.Br(36)
	_ = pdf.SetFont("Inter", "", 12)
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, "#")
	_ = pdf.Cell(nil, id)
	pdf.SetTextColor(150, 150, 150)
	_ = pdf.Cell(nil, "  ·  ")
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, date)
	pdf.Br(36)
}

func (s *pdfService) writeHeaderRow(pdf *gopdf.GoPdf) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, s.Translations.ItemNumberLabel)
	pdf.SetX(descriptionColumnOffset)
	_ = pdf.Cell(nil, s.Translations.ItemDescriptionLabel)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, s.Translations.ItemQuantityLabel)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, s.Translations.ItemUnitPriceLabel)
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, s.Translations.ItemTotalLabel)
	pdf.Br(24)
}

func writeFooter(pdf *gopdf.GoPdf, id string) {
	pdf.SetY(800)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, id)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX()+10, pdf.GetY()+6, 550, pdf.GetY()+6)
	pdf.Br(48)
}

func writeRow(pdf *gopdf.GoPdf, serial string, description string, quantity int, rate float64) {
	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(0, 0, 0)

	total := float64(quantity) * rate
	amount := strconv.FormatFloat(total, 'f', 2, 64)

	if len(serial) >= 20 {
		serial = serial[:20] + "..."
	}
	if len(description) >= 40 {
		description = description[:40] + "..."
	}

	_ = pdf.Cell(nil, serial)
	pdf.SetX(descriptionColumnOffset)
	_ = pdf.Cell(nil, description)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, strconv.Itoa(quantity))
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[currency]+strconv.FormatFloat(rate, 'f', 2, 64))
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[currency]+amount)
	pdf.Br(24)
}

func (p *pdfService) writeTotals(pdf *gopdf.GoPdf, totalExVat float64, vatAmount float64, total float64) {
	pdf.SetY(600)
	p.writeTotal(pdf, p.Translations.SubtotalLabel, totalExVat)
	p.writeTotal(pdf, p.Translations.TaxLabel, vatAmount)
	p.writeTotal(pdf, p.Translations.TotalLabel, total)
}

func (p *pdfService) writeTotal(pdf *gopdf.GoPdf, label string, total float64) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, label)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(12)
	pdf.SetX(rateColumnOffset + 35)
	if label == p.Translations.TotalLabel {
		_ = pdf.SetFont("", "", 11.5)
	}
	_ = pdf.Cell(nil, currencySymbols[currency]+strconv.FormatFloat(total, 'f', 2, 64))
	pdf.Br(24)
}
