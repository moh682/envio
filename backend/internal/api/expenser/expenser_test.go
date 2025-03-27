package expenser_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/moh682/envio/backend/internal/api/expenser"
	domain "github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/frameworks/postgres"
	expenserepo "github.com/moh682/envio/backend/internal/frameworks/postgres/repository/expense_repo"
	"github.com/stretchr/testify/assert"
)

func setupDatabase(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory database: %v", err)
	}

	migrator, err := postgres.NewMigrator(db, postgres.MigrationFS)
	if err != nil {
		t.Fatalf("Failed to create migrator: %v", err)
	}
	err = migrator.Up(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestExpenserE2E_GetPayments(t *testing.T) {

	db := setupDatabase(t)
	defer db.Close()

	repo := expenserepo.New(db)
	service := domain.NewExpenseService(repo)
	expenser := expenser.NewExpenser(service)

	result := expenser.GetPayments()
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(result.Payments))
	assert.Equal(t, "UNKNOWN", result.Payments[0].Name)
	assert.Equal(t, "CASH", result.Payments[1].Name)
	assert.Equal(t, "MOBILEPAY", result.Payments[2].Name)
	assert.Equal(t, "BANK_TRANSFER", result.Payments[3].Name)

}

func TestExpenserE2E_Create(t *testing.T) {

	db := setupDatabase(t)
	defer db.Close()

	repo := expenserepo.New(db)
	service := domain.NewExpenseService(repo)
	api := expenser.NewExpenser(service)

	_company := expenser.CreateCompanyDTO{
		Name: "Test Company",
		Cvr:  12345678,
	}
	createCompanyResult := api.CreateCompany(_company)
	assert.Nil(t, createCompanyResult.Error)

	getCompaniesResult := api.GetAllCompanies()
	assert.Nil(t, getCompaniesResult.Error)
	assert.Equal(t, 1, len(getCompaniesResult.Data))
	assert.Equal(t, getCompaniesResult.Data[0].Name, _company.Name)
	assert.Equal(t, getCompaniesResult.Data[0].Cvr, _company.Cvr)

	arg := expenser.CreateExpenseDTO{
		CompanyID:     getCompaniesResult.Data[0].ID,
		TotalInclVat:  100,
		TotalExclVat:  80,
		VatAmount:     20,
		VatPercentage: .25,
		IssuedAt:      gofakeit.PastDate(),
		PaidAt:        gofakeit.PastDate(),
		PaidWith:      domain.CASH.Index(),
		Entries: []expenser.CreateExpenseEntryDTO{
			{Serial: gofakeit.ProductUPC(), Description: gofakeit.ProductDescription(), UnitPrice: 10, Quantity: 2, TotalInclVat: 20, TotalExclVat: 16},
			{Serial: gofakeit.ProductUPC(), Description: gofakeit.ProductDescription(), UnitPrice: 10, Quantity: 2, TotalInclVat: 20, TotalExclVat: 16},
		},
	}
	result := api.Create(arg)
	assert.Nil(t, result.Error)

	getExpensesResult := api.GetAll(expenser.GetAllArgDTO{Limit: 10, Offset: 0})
	assert.Nil(t, getExpensesResult.Error)
	assert.Equal(t, 1, len(getExpensesResult.Data))
	assert.Equal(t, getExpensesResult.Data[0].Company.ID, arg.CompanyID)
	assert.Equal(t, getExpensesResult.Data[0].TotalInclVat, arg.TotalInclVat)
	assert.Equal(t, getExpensesResult.Data[0].TotalExclVat, arg.TotalExclVat)
	assert.Equal(t, getExpensesResult.Data[0].VatAmount, arg.VatAmount)
	assert.Equal(t, getExpensesResult.Data[0].VatPercentage, arg.VatPercentage)
	assert.Equal(t, getExpensesResult.Data[0].IssuedAt.Format("2006-01-02"), arg.IssuedAt.Format("2006-01-02"))
	assert.Equal(t, getExpensesResult.Data[0].PaidAt.Format("2006-01-02"), arg.PaidAt.Format("2006-01-02"))

	payment, _ := domain.NewPaymentMethod(arg.PaidWith)
	assert.Equal(t, getExpensesResult.Data[0].PaidWith, payment.String())

	assert.Equal(t, len(getExpensesResult.Data[0].Entries), 2)
	assert.Equal(t, getExpensesResult.Data[0].Entries[0].Serial, arg.Entries[0].Serial)
	assert.Equal(t, getExpensesResult.Data[0].Entries[0].Description, arg.Entries[0].Description)
	assert.Equal(t, getExpensesResult.Data[0].Entries[0].UnitPrice, arg.Entries[0].UnitPrice)
	assert.Equal(t, getExpensesResult.Data[0].Entries[0].Quantity, arg.Entries[0].Quantity)
	assert.Equal(t, getExpensesResult.Data[0].Entries[0].TotalInclVat, arg.Entries[0].TotalInclVat)
	assert.Equal(t, getExpensesResult.Data[0].Entries[0].TotalExclVat, arg.Entries[0].TotalExclVat)
	assert.Equal(t, getExpensesResult.Data[0].Entries[1].Serial, arg.Entries[1].Serial)
	assert.Equal(t, getExpensesResult.Data[0].Entries[1].Description, arg.Entries[1].Description)
	assert.Equal(t, getExpensesResult.Data[0].Entries[1].UnitPrice, arg.Entries[1].UnitPrice)
	assert.Equal(t, getExpensesResult.Data[0].Entries[1].Quantity, arg.Entries[1].Quantity)
	assert.Equal(t, getExpensesResult.Data[0].Entries[1].TotalInclVat, arg.Entries[1].TotalInclVat)
	assert.Equal(t, getExpensesResult.Data[0].Entries[1].TotalExclVat, arg.Entries[1].TotalExclVat)

}
