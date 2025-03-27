package invoice

import "time"

type Statistic struct {
	Date         time.Time `json:"date"`
	Count        int64     `json:"count"`
	TotalExclVat float64   `json:"total_excl_vat"`
	TotalInclVat float64   `json:"total_incl_vat"`
	VatAmount    float64   `json:"vat_amount"`
}
