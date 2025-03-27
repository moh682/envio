package analser

import (
	"context"

	"time"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/domain/invoice"
)

type Interval int

const (
	YearlyInterval Interval = iota
	MonthlyInterval
	DailyInterval
)

type Statistics struct {

	// Indicates the time start period for which the statistics are calculated
	from time.Time

	// Indicates the time end period for which the statistics are calculated
	to time.Time

	// Financial calculation
	// Revenue is the total amount of money received from selling goods or services
	Revenue float64

	// Profit is the amount left after expenses have been subtracted from the revenue
	Profit float64

	// Expenses are the costs incurred in the process of generating revenue
	Expenses float64

	// Vat that comes from revenue
	RevenueVat float64

	// Vat that comes from expenses
	ExpensesVat float64

	// Is the amount the company either has to pay or is owed to them from the government, for the danish market it is 25%.
	// negative if the company has to pay and positive if the company is owed money
	DeltaVat float64

	// is the total amount of invoices that have been created in the given time period
	invoicesCount int64

	// is the total amount of expenses that have been created in the given time period
	expensesCount int64
}

// TODO: Implement the FinancialCalculator interface
// should be able to calculate the vat based on
type FinancialCalculator interface {
	GetHalfYearlyStatistics(ctx context.Context, from, to time.Time) (*Statistics, error)
}

type analyserService struct {
	invoices []invoice.Invoice
	expenses []domain.Expense
}

// change the domain.Expense to expense.Expense after we have moved the expense into its own package/domain
func NewAnalyserService(invoices []invoice.Invoice, expenses []domain.Expense) FinancialCalculator {
	return &analyserService{
		invoices: invoices,
		expenses: expenses,
	}
}

func (a *analyserService) GetHalfYearlyStatistics(ctx context.Context, from, to time.Time) (*Statistics, error) {

	revenue := float64(0)
	revenueVat := float64(0)
	invoicesCount := int64(0)

	for _, inv := range a.invoices {
		if inv.IssuedAt().After(from) && inv.IssuedAt().Before(to) {
			revenue += inv.Total()
			revenueVat += inv.VatAmount()
			invoicesCount++
		}
	}

	expenses := float64(0)
	expensesVat := float64(0)
	expensesCount := int64(0)

	for _, exp := range a.expenses {
		if exp.IssuedAt.After(from) && exp.IssuedAt.Before(to) {
			expenses += float64(exp.TotalInclVat)
			expensesVat += float64(exp.VatAmount)
			expensesCount++
		}
	}

	profit := revenue - expenses
	deltaVat := revenueVat - expensesVat

	return &Statistics{
		Revenue:       revenue,
		Profit:        profit,
		Expenses:      expenses,
		RevenueVat:    revenueVat,
		ExpensesVat:   expensesVat,
		DeltaVat:      deltaVat,
		invoicesCount: invoicesCount,
		expensesCount: expensesCount,
	}, nil
}
