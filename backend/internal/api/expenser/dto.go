package expenser

import (
	"time"
)

type ExpenseDTO struct {
	ID            string            `json:"id"`
	Company       *CompanyDTO       `json:"company"`
	TotalInclVat  float32           `json:"totalInclVat"`
	TotalExclVat  float32           `json:"totalExclVat"`
	VatAmount     float32           `json:"vatAmount"`
	VatPercentage float32           `json:"vatPercentage"`
	IssuedAt      time.Time         `json:"issuedAt"`
	PaidAt        time.Time         `json:"paidAt"`
	PaidWith      string            `json:"paidWith"`
	Entries       []ExpenseEntryDTO `json:"entries"`
}

type CompanyDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cvr   int    `json:"cvr"`
	Phone string `json:"phone"`
}

type ExpenseEntryDTO struct {
	ID            string  `json:"id"`
	Serial        string  `json:"serial"`
	Description   string  `json:"description"`
	UnitPrice     float32 `json:"unitPrice"`
	Quantity      int32   `json:"quantity"`
	TotalInclVat  float32 `json:"totalInclVat"`
	TotalExclVat  float32 `json:"totalExclVat"`
	VatAmount     float32 `json:"vatAmount"`
	VatPercentage float32 `json:"vatPercentage"`
}

type PaymentDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type GetAllPaymentDTOs struct {
	Payments []PaymentDTO `json:"payments"`
	Error    *ResultError `json:"error"`
}

type GetYearlyExpensesCountResultDTO struct {
	Error *ResultError            `json:"error"`
	Data  []YearlyExpenseCountDTO `json:"data"`
}

type YearlyExpenseCountDTO struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

type CreateExpenseDTO struct {
	CompanyID     string                  `json:"companyId"`
	TotalInclVat  float32                 `json:"totalInclVat"`
	TotalExclVat  float32                 `json:"totalExclVat"`
	VatAmount     float32                 `json:"vatAmount"`
	VatPercentage float32                 `json:"vatPercentage"`
	IssuedAt      time.Time               `json:"issuedAt"`
	PaidAt        time.Time               `json:"paidAt"`
	PaidWith      int                     `json:"paidWith"`
	Entries       []CreateExpenseEntryDTO `json:"entries"`
}

type CreateExpenseEntryDTO struct {
	Serial       string  `json:"serial"`
	Description  string  `json:"description"`
	UnitPrice    float32 `json:"unitPrice"`
	Quantity     int32   `json:"quantity"`
	TotalInclVat float32 `json:"totalInclVat"`
	TotalExclVat float32 `json:"totalExclVat"`
}

type CreateCompanyDTO struct {
	Name string `json:"name"`
	Cvr  int    `json:"cvr"`
}

type GetCompanyByNameArgDTO struct {
	Name string `json:"name"`
}

type GetCompanyByNameResultDTO struct {
	Error *ResultError `json:"error"`
	Data  *CompanyDTO  `json:"data"`
}

type GetAllArgDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type CreateExpenseResultDTO struct {
	Error *ResultError `json:"error"`
}

type GetCompaniesResultDTO struct {
	Error *ResultError `json:"error,omitempty"`
	Data  []CompanyDTO `json:"data"`
}

type CreateCompanyResultDTO struct {
	Error *ResultError `json:"error"`
	Data  *CompanyDTO  `json:"data,omitempty"`
}

type ResultError struct {
	Message string `json:"message"`
}

type GetAllResultDTO struct {
	Error *ResultError `json:"error,omitempty"`
	Data  []ExpenseDTO `json:"data"`
}
