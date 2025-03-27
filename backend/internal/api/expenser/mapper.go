package expenser

import (
	"github.com/moh682/envio/backend/internal/domain"
)

func (exp *expenser) toYearlyExpenseCountDTO(counts []*domain.ExpenseCount) []YearlyExpenseCountDTO {
	expenseCounts := []YearlyExpenseCountDTO{}
	for _, c := range counts {
		expenseCounts = append(expenseCounts, YearlyExpenseCountDTO{
			Year:  c.Year,
			Count: c.Count,
		})
	}
	return expenseCounts
}

func (exp *expenser) toExpenseEntryDTO(entries []*domain.Entry) []ExpenseEntryDTO {
	entryDTOs := []ExpenseEntryDTO{}
	for _, e := range entries {
		entryDTOs = append(entryDTOs, ExpenseEntryDTO{
			ID:           e.ID.String(),
			Serial:       e.Serial,
			Description:  e.Description,
			UnitPrice:    e.UnitPrice,
			Quantity:     e.Quantity,
			TotalInclVat: e.TotalInclVat,
			TotalExclVat: e.TotalExclVat,
		})
	}
	return entryDTOs
}

func (exp *expenser) toCompanyDTO(c *domain.Company) *CompanyDTO {
	return &CompanyDTO{
		ID:   c.ID.String(),
		Name: c.Name,
		Cvr:  c.Cvr,
	}
}

func (exp *expenser) toExpenseDTO(e *domain.Expense) *ExpenseDTO {
	entries := exp.toExpenseEntryDTO(e.Entries)
	return &ExpenseDTO{
		ID:            e.ID.String(),
		Company:       exp.toCompanyDTO(e.Company),
		TotalInclVat:  e.TotalInclVat,
		TotalExclVat:  e.TotalExclVat,
		VatAmount:     e.VatAmount,
		VatPercentage: e.VatPercentage,

		IssuedAt: e.IssuedAt,
		PaidAt:   e.PaidAt,
		PaidWith: e.PaidWith.String(),
		Entries:  entries,
	}
}
