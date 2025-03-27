package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTimeIsZero = errors.New("time is zero")
)

type ExpenseService interface {
	Create(exp *Expense) error
	GetAll(limit, offset int32) ([]*Expense, error)
	GetPaymentMethods() []PaymentMethod

	GetCompaniesByName(name string) ([]*Company, error)
	GetCompanyByID(id string) (*Company, error)
	CreateCompany(comp *Company) error
	GetAllCompanies() ([]*Company, error)

	GetAllExpensesSince(since time.Time) ([]*Expense, error)
	GetYearlyExpensesCount() ([]*ExpenseCount, error)
}

type expenseService struct {
	repo ExpenseRepo
}

func NewExpenseService(repo ExpenseRepo) ExpenseService {
	return &expenseService{
		repo: repo,
	}
}

func (e *expenseService) GetYearlyExpensesCount() ([]*ExpenseCount, error) {
	ctx := context.Background()
	return e.repo.GetYearlyExpensesCount(ctx)
}

func (e *expenseService) Create(exp *Expense) error {
	ctx := context.Background()
	err := exp.Validate()
	if err != nil {
		return err
	}

	return e.repo.Create(ctx, exp)
}

func (e *expenseService) GetAll(limit, offset int32) ([]*Expense, error) {
	ctx := context.Background()
	return e.repo.GetAll(ctx, limit, offset)
}

func (e *expenseService) GetPaymentMethods() []PaymentMethod {
	methods := make([]PaymentMethod, len(paymentMethods))
	for idx := range paymentMethods {
		method, _ := NewPaymentMethod(idx)
		methods[idx] = method
	}
	return methods
}

func (e *expenseService) GetCompaniesByName(name string) ([]*Company, error) {
	ctx := context.Background()
	return e.repo.GetCompaniesByName(ctx, name)
}

func (e *expenseService) GetAllCompanies() ([]*Company, error) {
	ctx := context.Background()
	return e.repo.GetAllCompanies(ctx)
}

func (e *expenseService) GetCompanyByID(id string) (*Company, error) {
	ctx := context.Background()
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return e.repo.GetCompanyByID(ctx, id)
}

func (e *expenseService) CreateCompany(comp *Company) error {
	ctx := context.Background()
	err := comp.Validate()
	if err != nil {
		return err
	}
	return e.repo.CreateCompany(ctx, comp)

}

func (e *expenseService) GetAllExpensesSince(since time.Time) ([]*Expense, error) {
	ctx := context.Background()
	if since.IsZero() {
		return nil, ErrTimeIsZero
	}

	return e.repo.GetAllExpensesSince(ctx, since)
}
