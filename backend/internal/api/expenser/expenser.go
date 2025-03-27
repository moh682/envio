package expenser

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
)

type Expenser interface {
	Create(arg CreateExpenseDTO) CreateExpenseResultDTO
	GetAll(arg GetAllArgDTO) GetAllResultDTO
	GetPayments() GetAllPaymentDTOs
	CreateCompany(arg CreateCompanyDTO) CreateCompanyResultDTO
	GetAllCompanies() GetCompaniesResultDTO
	GetYearlyExpensesCount() GetYearlyExpensesCountResultDTO
}

type expenser struct {
	expenseS domain.ExpenseService
}

func NewExpenser(
	expenseService domain.ExpenseService,
) Expenser {
	return &expenser{
		expenseS: expenseService,
	}
}

func (e *expenser) GetYearlyExpensesCount() GetYearlyExpensesCountResultDTO {
	count, err := e.expenseS.GetYearlyExpensesCount()
	if err != nil {
		return GetYearlyExpensesCountResultDTO{
			Error: &ResultError{Message: err.Error()},
		}
	}

	return GetYearlyExpensesCountResultDTO{
		Error: nil,
		Data:  e.toYearlyExpenseCountDTO(count),
	}
}

func (e *expenser) GetPayments() GetAllPaymentDTOs {
	payments := e.expenseS.GetPaymentMethods()
	paymentDTOs := make([]PaymentDTO, len(payments))
	for idx, p := range payments {
		paymentDTOs[idx] = PaymentDTO{
			Name:  p.String(),
			Index: p.Index(),
		}
	}

	return GetAllPaymentDTOs{
		Payments: paymentDTOs,
		Error:    nil,
	}
}

func (e *expenser) GetAll(arg GetAllArgDTO) GetAllResultDTO {

	exp, err := e.expenseS.GetAll(arg.Limit, arg.Offset)
	if err != nil {
		return GetAllResultDTO{
			Error: &ResultError{Message: err.Error()},
			Data:  make([]ExpenseDTO, 0),
		}
	}

	expenses := []ExpenseDTO{}
	for _, ex := range exp {
		expenses = append(expenses, *e.toExpenseDTO(ex))
	}

	return GetAllResultDTO{
		Error: nil,
		Data:  expenses,
	}

}

func (e *expenser) CreateCompany(arg CreateCompanyDTO) CreateCompanyResultDTO {

	id := uuid.New()
	comp := &domain.Company{ID: domain.ID(id), Name: arg.Name, Cvr: arg.Cvr, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := comp.Validate()
	if err != nil {
		return CreateCompanyResultDTO{
			Error: &ResultError{Message: err.Error()},
		}
	}
	err = e.expenseS.CreateCompany(comp)
	if err != nil {
		return CreateCompanyResultDTO{
			Error: &ResultError{Message: err.Error()},
		}
	}

	return CreateCompanyResultDTO{
		Error: nil,
		Data:  e.toCompanyDTO(comp),
	}
}

func (e *expenser) GetAllCompanies() GetCompaniesResultDTO {
	companies, err := e.expenseS.GetAllCompanies()

	if err != nil {
		return GetCompaniesResultDTO{
			Error: &ResultError{Message: err.Error()},
			Data:  make([]CompanyDTO, 0),
		}
	}

	companyDTOs := make([]CompanyDTO, len(companies))
	for idx, comp := range companies {
		companyDTOs[idx] = *e.toCompanyDTO(comp)
	}

	return GetCompaniesResultDTO{
		Error: nil,
		Data:  companyDTOs,
	}
}

func (e *expenser) Create(arg CreateExpenseDTO) CreateExpenseResultDTO {

	entries := make([]*domain.Entry, len(arg.Entries))
	for idx, ent := range arg.Entries {
		entries[idx] = &domain.Entry{ID: domain.ID(uuid.New()), Serial: ent.Serial, Description: ent.Description, UnitPrice: ent.UnitPrice, Quantity: ent.Quantity, TotalInclVat: ent.TotalInclVat, TotalExclVat: ent.TotalExclVat}
	}

	company, err := e.expenseS.GetCompanyByID(arg.CompanyID)
	if err != nil {
		return CreateExpenseResultDTO{
			Error: &ResultError{Message: fmt.Sprintf("could not find company with id %s", arg.CompanyID)},
		}
	}

	exp := &domain.Expense{
		ID:            domain.ID(uuid.New()),
		Company:       company,
		VatPercentage: arg.VatPercentage, TotalInclVat: arg.TotalInclVat, TotalExclVat: arg.TotalExclVat, VatAmount: arg.VatAmount, IssuedAt: arg.IssuedAt, PaidAt: arg.PaidAt, PaidWith: domain.PaymentMethod(arg.PaidWith), Entries: entries,
	}
	err = exp.Validate()
	if err != nil {
		return CreateExpenseResultDTO{
			Error: &ResultError{Message: err.Error()},
		}
	}

	err = e.expenseS.Create(exp)
	if err != nil {
		return CreateExpenseResultDTO{
			Error: &ResultError{Message: err.Error()},
		}
	}

	return CreateExpenseResultDTO{
		Error: nil,
	}
}
