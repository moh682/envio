package expenserepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/moh682/envio/backend/internal/frameworks/postgres/db"
)

type expenseRepo struct {
	db *sql.DB
}

func New(db *sql.DB) domain.ExpenseRepo {
	return &expenseRepo{
		db,
	}
}

var (
	ErrNoExpenseEntriesFound = errors.New("no expense entries found")
	ErrNoCompanyFound        = errors.New("no company found")
	ErrorNoExpenseFound      = errors.New("no expense not found")
)

func (r *expenseRepo) GetYearlyExpensesCount(ctx context.Context) ([]*domain.ExpenseCount, error) {
	q := db.New(r.db)
	result, err := q.GetYearlyExpensesCount(ctx)
	if err != nil {
		return nil, err
	}

	expenseCounts := make([]*domain.ExpenseCount, len(result))
	for index, row := range result {
		year, err := time.Parse("2006", row.Year)
		if err != nil {
			return nil, errors.New("could not convert year to string")
		}

		expenseCounts[index] = &domain.ExpenseCount{
			Year:  year.Year(),
			Count: int(row.Count),
		}
	}

	return expenseCounts, nil
}

func (r *expenseRepo) GetAllExpensesSince(ctx context.Context, since time.Time) ([]*domain.Expense, error) {

	q := db.New(r.db)
	expenses, err := q.GetAllExpensesSince(ctx, since)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Expense, len(expenses))
	for idx, e := range expenses {
		company, err := q.GetCompanyByID(ctx, e.CompanyID)
		if err != nil {
			return nil, err
		}
		entries, err := q.GetAllExpenseEntriesByExpenseId(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		exp, err := toExpenseModel(e, entries, company)
		if err != nil {
			return nil, err
		}
		result[idx] = exp
	}

	return result, nil

}

func (r *expenseRepo) GetCompanyByID(ctx context.Context, id string) (*domain.Company, error) {
	q := db.New(r.db)
	company, err := q.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toCompanyModel(company)
}

func (r *expenseRepo) GetAllCompanies(ctx context.Context) ([]*domain.Company, error) {
	q := db.New(r.db)
	companies, err := q.GetAllCompanies(ctx)
	if err != nil {
		return nil, err
	}

	result := []*domain.Company{}
	for _, c := range companies {
		company, err := toCompanyModel(c)

		if err != nil {
			return nil, err
		}
		result = append(result, company)
	}
	return result, nil
}

func (r *expenseRepo) GetCompaniesByName(ctx context.Context, name string) ([]*domain.Company, error) {
	q := db.New(r.db)
	companies, err := q.GetCompaniesByName(ctx, name)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Company, 0)
	for _, c := range companies {
		company, err := toCompanyModel(c)
		if err != nil {
			return nil, err
		}
		result = append(result, company)
	}
	return result, nil
}

func (r *expenseRepo) CreateCompany(ctx context.Context, company *domain.Company) error {
	q := db.New(r.db)

	arg := db.CreateCompanyParams{
		ID:   company.ID.String(),
		Name: company.Name,
		Cvr: sql.NullInt64{
			Int64: int64(company.Cvr),
			Valid: company.Cvr != 0,
		},
	}
	return q.CreateCompany(ctx, arg)
}

func (r *expenseRepo) Create(ctx context.Context, exp *domain.Expense) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := db.New(r.db)
	q.WithTx(tx)

	arg := db.CreateExpenseParams{
		ID:           exp.ID.String(),
		CompanyID:    exp.Company.ID.String(),
		TotalInclVat: float64(exp.TotalInclVat),
		TotalExclVat: float64(exp.TotalExclVat),
		VatAmount:    float64(exp.VatAmount),
		VatRate:      float64(exp.VatPercentage),
		PaidWith: sql.NullInt64{
			Int64: int64(exp.PaidWith.Index()),
			Valid: exp.PaidWith != domain.UNKNOWN,
		},
		IssuedAt: exp.IssuedAt,
		PaidAt: sql.NullTime{
			Valid: !exp.IssuedAt.IsZero(),
			Time:  exp.IssuedAt,
		},
	}
	err = q.CreateExpense(ctx, arg)
	if err != nil {
		return err
	}

	for _, ent := range exp.Entries {
		arg := db.CreateExpenseEntryParams{
			ID:        ent.ID.String(),
			ExpenseID: exp.ID.String(),
			Serial: sql.NullString{
				String: ent.Serial,
				Valid:  ent.Serial != "",
			},
			Description: sql.NullString{
				String: ent.Description,
				Valid:  ent.Description != "",
			},
			UnitPrice:    float64(ent.UnitPrice),
			Quantity:     int64(ent.Quantity),
			TotalExclVat: float64(ent.TotalExclVat),
			TotalInclVat: float64(ent.TotalInclVat),
		}
		err = q.CreateExpenseEntry(ctx, arg)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *expenseRepo) GetAll(ctx context.Context, limit, offset int32) ([]*domain.Expense, error) {

	q := db.New(r.db)
	arg := db.GetAllExpensesParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	result, err := q.GetAllExpenses(ctx, arg)
	if err != nil {
		return nil, err
	}

	expenses := []*domain.Expense{}
	for _, e := range result {
		company, err := q.GetCompanyByID(ctx, e.CompanyID)
		if err != nil {
			return nil, err
		}
		entries, err := q.GetAllExpenseEntriesByExpenseId(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		exp, err := toExpenseModel(e, entries, company)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, exp)
	}

	return expenses, nil
}
