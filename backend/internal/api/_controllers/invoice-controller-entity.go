package controllers

type product struct {
	ProductNr   string  `json:"product_nr"`
	Description string  `json:"description"`
	Quantity    float32 `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
	Total       float32 `json:"total"`
}

type NewInvoice struct {
	InvoiceTotal            float32   `json:"invoice_total"`
	InvoiceTotalExclVat     float32   `json:"invoice_total_excl_vat"`
	IssuedAt                string    `json:"issued_at"`
	VatPct                  float32   `json:"vat_pct"`
	VatAmount               float32   `json:"vat_amount"`
	CustomerName            string    `json:"customer_name"`
	CustomerAddress         string    `json:"customer_address"`
	CustomerZipCode         string    `json:"customer_zip_code"`
	CustomerEmail           string    `json:"customer_email"`
	CustomerPhone           string    `json:"customer_phone"`
	CustomerCarRegistration string    `json:"customer_car_registration"`
	Products                []product `json:"products"`
}
