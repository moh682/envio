package invoice

import "time"

type Statistic struct {
	Date         time.Time `json:"date"`
	Count        int64     `json:"count"`
	TotalExclVat float64   `json:"totalExclVat"`
	TotalInclVat float64   `json:"totalInclVat"`
	VatAmount    float64   `json:"vatAmount"`
}
