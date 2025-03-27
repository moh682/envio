// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createCompany = `-- name: CreateCompany :exec
insert into companies (id, name, cvr) VALUES ($1, $2, $3)
`

type CreateCompanyParams struct {
	ID   uuid.UUID
	Name string
	Cvr  sql.NullInt32
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) error {
	_, err := q.db.ExecContext(ctx, createCompany, arg.ID, arg.Name, arg.Cvr)
	return err
}

const createCost = `-- name: CreateCost :exec
INSERT INTO costs (id, product_number, invoice_id, description, quantity, unit_price, total) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateCostParams struct {
	ID            uuid.UUID
	ProductNumber sql.NullString
	InvoiceID     uuid.UUID
	Description   sql.NullString
	Quantity      float64
	UnitPrice     float64
	Total         float64
}

func (q *Queries) CreateCost(ctx context.Context, arg CreateCostParams) error {
	_, err := q.db.ExecContext(ctx, createCost,
		arg.ID,
		arg.ProductNumber,
		arg.InvoiceID,
		arg.Description,
		arg.Quantity,
		arg.UnitPrice,
		arg.Total,
	)
	return err
}

const createCustomer = `-- name: CreateCustomer :exec
INSERT INTO customers (id, car_registration, name, email, phone, address, zip_code, invoice_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateCustomerParams struct {
	ID              uuid.UUID
	CarRegistration sql.NullString
	Name            sql.NullString
	Email           sql.NullString
	Phone           sql.NullString
	Address         sql.NullString
	ZipCode         sql.NullString
	InvoiceID       uuid.UUID
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, createCustomer,
		arg.ID,
		arg.CarRegistration,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Address,
		arg.ZipCode,
		arg.InvoiceID,
	)
	return err
}

const createExpense = `-- name: CreateExpense :exec
insert into expenses (id, company_id, total_incl_vat, total_excl_vat, vat_amount, vat_rate, issued_at, paid_at, paid_with) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreateExpenseParams struct {
	ID           uuid.UUID
	CompanyID    uuid.UUID
	TotalInclVat float64
	TotalExclVat float64
	VatAmount    float64
	VatRate      float64
	IssuedAt     time.Time
	PaidAt       sql.NullTime
	PaidWith     sql.NullInt32
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) error {
	_, err := q.db.ExecContext(ctx, createExpense,
		arg.ID,
		arg.CompanyID,
		arg.TotalInclVat,
		arg.TotalExclVat,
		arg.VatAmount,
		arg.VatRate,
		arg.IssuedAt,
		arg.PaidAt,
		arg.PaidWith,
	)
	return err
}

const createExpenseEntry = `-- name: CreateExpenseEntry :exec
insert into expenses_entries (id, expense_id, serial, description, unit_price, quantity, total_excl_vat, total_incl_vat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateExpenseEntryParams struct {
	ID           uuid.UUID
	ExpenseID    uuid.UUID
	Serial       sql.NullString
	Description  sql.NullString
	UnitPrice    float64
	Quantity     int32
	TotalExclVat float64
	TotalInclVat float64
}

func (q *Queries) CreateExpenseEntry(ctx context.Context, arg CreateExpenseEntryParams) error {
	_, err := q.db.ExecContext(ctx, createExpenseEntry,
		arg.ID,
		arg.ExpenseID,
		arg.Serial,
		arg.Description,
		arg.UnitPrice,
		arg.Quantity,
		arg.TotalExclVat,
		arg.TotalInclVat,
	)
	return err
}

const createInvoice = `-- name: CreateInvoice :exec
INSERT INTO invoices (id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status) VALUES ($1, (SELECT COALESCE(MAX(invoice_number), 0) + 1 FROM invoices), $2, $3, $4, $5, $6, $7)
`

type CreateInvoiceParams struct {
	ID              uuid.UUID
	IssuedAt        sql.NullTime
	VatPercentage   float64
	VatAmount       float64
	TotalExcludeVat float64
	TotalIncludeVat float64
	PaymentStatus   interface{}
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) error {
	_, err := q.db.ExecContext(ctx, createInvoice,
		arg.ID,
		arg.IssuedAt,
		arg.VatPercentage,
		arg.VatAmount,
		arg.TotalExcludeVat,
		arg.TotalIncludeVat,
		arg.PaymentStatus,
	)
	return err
}

const createInvoiceMigration = `-- name: CreateInvoiceMigration :exec
INSERT INTO invoice_migrations (id, invoice_number, file_path) VALUES ($1, $2, $3)
`

type CreateInvoiceMigrationParams struct {
	ID            uuid.UUID
	InvoiceNumber int32
	FilePath      sql.NullString
}

func (q *Queries) CreateInvoiceMigration(ctx context.Context, arg CreateInvoiceMigrationParams) error {
	_, err := q.db.ExecContext(ctx, createInvoiceMigration, arg.ID, arg.InvoiceNumber, arg.FilePath)
	return err
}

const createPayment = `-- name: CreatePayment :exec
INSERT INTO payments (id, invoice_id, method, amount) VALUES ($1, $2, $3, $4)
`

type CreatePaymentParams struct {
	ID        uuid.UUID
	InvoiceID uuid.UUID
	Method    string
	Amount    float64
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, createPayment,
		arg.ID,
		arg.InvoiceID,
		arg.Method,
		arg.Amount,
	)
	return err
}

const getAllCompanies = `-- name: GetAllCompanies :many
select id, name, cvr, created_at, updated_at from companies
`

func (q *Queries) GetAllCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, getAllCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Cvr,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllCostsByInvoiceId = `-- name: GetAllCostsByInvoiceId :many
select id, invoice_id, product_number, description, quantity, unit_price, total from costs where invoice_id = $1
`

func (q *Queries) GetAllCostsByInvoiceId(ctx context.Context, invoiceID uuid.UUID) ([]Cost, error) {
	rows, err := q.db.QueryContext(ctx, getAllCostsByInvoiceId, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cost
	for rows.Next() {
		var i Cost
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.ProductNumber,
			&i.Description,
			&i.Quantity,
			&i.UnitPrice,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllExpenseEntriesByExpenseId = `-- name: GetAllExpenseEntriesByExpenseId :many
select id, expense_id, serial, description, unit_price, quantity, total_excl_vat, total_incl_vat from expenses_entries where expense_id = $1
`

func (q *Queries) GetAllExpenseEntriesByExpenseId(ctx context.Context, expenseID uuid.UUID) ([]ExpensesEntry, error) {
	rows, err := q.db.QueryContext(ctx, getAllExpenseEntriesByExpenseId, expenseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExpensesEntry
	for rows.Next() {
		var i ExpensesEntry
		if err := rows.Scan(
			&i.ID,
			&i.ExpenseID,
			&i.Serial,
			&i.Description,
			&i.UnitPrice,
			&i.Quantity,
			&i.TotalExclVat,
			&i.TotalInclVat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllExpenses = `-- name: GetAllExpenses :many
select id, total_excl_vat, total_incl_vat, vat_amount, vat_rate, issued_at, paid_at, paid_with, company_id from expenses LIMIT $1 OFFSET $2
`

type GetAllExpensesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllExpenses(ctx context.Context, arg GetAllExpensesParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getAllExpenses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.TotalExclVat,
			&i.TotalInclVat,
			&i.VatAmount,
			&i.VatRate,
			&i.IssuedAt,
			&i.PaidAt,
			&i.PaidWith,
			&i.CompanyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllExpensesSince = `-- name: GetAllExpensesSince :many
select id, total_excl_vat, total_incl_vat, vat_amount, vat_rate, issued_at, paid_at, paid_with, company_id from expenses where issued_at > $1
`

func (q *Queries) GetAllExpensesSince(ctx context.Context, issuedAt time.Time) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getAllExpensesSince, issuedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.TotalExclVat,
			&i.TotalInclVat,
			&i.VatAmount,
			&i.VatRate,
			&i.IssuedAt,
			&i.PaidAt,
			&i.PaidWith,
			&i.CompanyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllInvoices = `-- name: GetAllInvoices :many
SELECT id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status FROM invoices LIMIT $1 OFFSET $2
`

type GetAllInvoicesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllInvoices(ctx context.Context, arg GetAllInvoicesParams) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, getAllInvoices, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceNumber,
			&i.IssuedAt,
			&i.VatPercentage,
			&i.VatAmount,
			&i.TotalExcludeVat,
			&i.TotalIncludeVat,
			&i.PaymentStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllInvoicesSince = `-- name: GetAllInvoicesSince :many
select id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status from invoices where issued_at > $1
`

func (q *Queries) GetAllInvoicesSince(ctx context.Context, issuedAt sql.NullTime) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, getAllInvoicesSince, issuedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceNumber,
			&i.IssuedAt,
			&i.VatPercentage,
			&i.VatAmount,
			&i.TotalExcludeVat,
			&i.TotalIncludeVat,
			&i.PaymentStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompaniesByName = `-- name: GetCompaniesByName :many
select id, name, cvr, created_at, updated_at from companies where name like $1
`

func (q *Queries) GetCompaniesByName(ctx context.Context, name string) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, getCompaniesByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Cvr,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyByID = `-- name: GetCompanyByID :one
select id, name, cvr, created_at, updated_at from companies where id = $1
`

func (q *Queries) GetCompanyByID(ctx context.Context, id uuid.UUID) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompanyByID, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cvr,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCustomerById = `-- name: GetCustomerById :one
select id, car_registration, invoice_id, name, email, phone, address, zip_code, created_at from customers where id = $1
`

func (q *Queries) GetCustomerById(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerById, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CarRegistration,
		&i.InvoiceID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.ZipCode,
		&i.CreatedAt,
	)
	return i, err
}

const getCustomerByInvoiceId = `-- name: GetCustomerByInvoiceId :one
select id, car_registration, invoice_id, name, email, phone, address, zip_code, created_at from customers where invoice_id = $1
`

func (q *Queries) GetCustomerByInvoiceId(ctx context.Context, invoiceID uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByInvoiceId, invoiceID)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CarRegistration,
		&i.InvoiceID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.ZipCode,
		&i.CreatedAt,
	)
	return i, err
}

const getDailyExpenseStatistics = `-- name: GetDailyExpenseStatistics :many
SELECT to_char(issued_at, 'DD/MM/YYYY') as date, count(*) as expense_count, cast(sum(total_excl_vat) as float) as total_excl_vat, cast(sum(total_incl_vat) as float) as total_incl_vat, cast(sum(vat_amount) as float) as vat_amount FROM expenses GROUP BY to_char(issued_at, 'YYYY') ORDER BY date
`

type GetDailyExpenseStatisticsRow struct {
	Date         string
	ExpenseCount int64
	TotalExclVat float64
	TotalInclVat float64
	VatAmount    float64
}

func (q *Queries) GetDailyExpenseStatistics(ctx context.Context) ([]GetDailyExpenseStatisticsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDailyExpenseStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDailyExpenseStatisticsRow
	for rows.Next() {
		var i GetDailyExpenseStatisticsRow
		if err := rows.Scan(
			&i.Date,
			&i.ExpenseCount,
			&i.TotalExclVat,
			&i.TotalInclVat,
			&i.VatAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDailyInvoiceStatistics = `-- name: GetDailyInvoiceStatistics :many
SELECT to_char(issued_at, 'DD/MM/YYYY') as date, count(*) as invoice_count, cast(sum(total_exclude_vat) as float) as total_exclude_vat, cast(sum(total_include_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM invoices GROUP BY to_char(issued_at, 'DD/MM/YYYY') order BY date
`

type GetDailyInvoiceStatisticsRow struct {
	Date            string
	InvoiceCount    int64
	TotalExcludeVat float64
	TotalIncludeVat float64
	VatAmount       float64
}

func (q *Queries) GetDailyInvoiceStatistics(ctx context.Context) ([]GetDailyInvoiceStatisticsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDailyInvoiceStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDailyInvoiceStatisticsRow
	for rows.Next() {
		var i GetDailyInvoiceStatisticsRow
		if err := rows.Scan(
			&i.Date,
			&i.InvoiceCount,
			&i.TotalExcludeVat,
			&i.TotalIncludeVat,
			&i.VatAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvoiceById = `-- name: GetInvoiceById :one
SELECT id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status FROM invoices WHERE id = $1
`

func (q *Queries) GetInvoiceById(ctx context.Context, id uuid.UUID) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, getInvoiceById, id)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.InvoiceNumber,
		&i.IssuedAt,
		&i.VatPercentage,
		&i.VatAmount,
		&i.TotalExcludeVat,
		&i.TotalIncludeVat,
		&i.PaymentStatus,
	)
	return i, err
}

const getInvoiceByNumber = `-- name: GetInvoiceByNumber :one
SELECT id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status FROM invoices WHERE invoice_number = $1
`

func (q *Queries) GetInvoiceByNumber(ctx context.Context, invoiceNumber int32) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, getInvoiceByNumber, invoiceNumber)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.InvoiceNumber,
		&i.IssuedAt,
		&i.VatPercentage,
		&i.VatAmount,
		&i.TotalExcludeVat,
		&i.TotalIncludeVat,
		&i.PaymentStatus,
	)
	return i, err
}

const getMaxInvoiceNumber = `-- name: GetMaxInvoiceNumber :one
SELECT MAX(invoice_number) FROM invoices
`

func (q *Queries) GetMaxInvoiceNumber(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMaxInvoiceNumber)
	var max interface{}
	err := row.Scan(&max)
	return max, err
}

const getMonthlyExpenseStatistics = `-- name: GetMonthlyExpenseStatistics :many
SELECT cast(strftime('%m/%Y', issued_at) as TEXT) as date, count(*) as expense_count, cast(sum(total_excl_vat) as float) as total_exclude_vat, cast(sum(total_incl_vat) as float) as total_incl_vat, cast(sum(vat_amount) as float) as vat_amount FROM expenses GROUP BY date order BY date
`

type GetMonthlyExpenseStatisticsRow struct {
	Date            string
	ExpenseCount    int64
	TotalExcludeVat float64
	TotalInclVat    float64
	VatAmount       float64
}

func (q *Queries) GetMonthlyExpenseStatistics(ctx context.Context) ([]GetMonthlyExpenseStatisticsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyExpenseStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonthlyExpenseStatisticsRow
	for rows.Next() {
		var i GetMonthlyExpenseStatisticsRow
		if err := rows.Scan(
			&i.Date,
			&i.ExpenseCount,
			&i.TotalExcludeVat,
			&i.TotalInclVat,
			&i.VatAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyInvoiceStatistics = `-- name: GetMonthlyInvoiceStatistics :many
SELECT to_char(issued_at, 'MM/YYYY') as date, count(*) as invoice_count, cast(sum(total_exclude_vat) as float) as total_exclude_vat, cast(sum(total_include_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM invoices GROUP BY to_char(issued_at, 'MM/YYYY') order BY date
`

type GetMonthlyInvoiceStatisticsRow struct {
	Date            string
	InvoiceCount    int64
	TotalExcludeVat float64
	TotalIncludeVat float64
	VatAmount       float64
}

func (q *Queries) GetMonthlyInvoiceStatistics(ctx context.Context) ([]GetMonthlyInvoiceStatisticsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyInvoiceStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonthlyInvoiceStatisticsRow
	for rows.Next() {
		var i GetMonthlyInvoiceStatisticsRow
		if err := rows.Scan(
			&i.Date,
			&i.InvoiceCount,
			&i.TotalExcludeVat,
			&i.TotalIncludeVat,
			&i.VatAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaymentById = `-- name: GetPaymentById :one
select id, invoice_id, amount, paid_at, method from payments where id = $1
`

func (q *Queries) GetPaymentById(ctx context.Context, id uuid.UUID) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPaymentById, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.InvoiceID,
		&i.Amount,
		&i.PaidAt,
		&i.Method,
	)
	return i, err
}

const getPaymentsByInvoice = `-- name: GetPaymentsByInvoice :many
select id, invoice_id, amount, paid_at, method from payments where invoice_id = $1
`

func (q *Queries) GetPaymentsByInvoice(ctx context.Context, invoiceID uuid.UUID) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByInvoice, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.Amount,
			&i.PaidAt,
			&i.Method,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyExpenseStatistics = `-- name: GetYearlyExpenseStatistics :many
SELECT cast(strftime('%Y', issued_at) as TEXT) as date, count(*) as expense_count, cast(sum(total_excl_vat) as float) as total_exclude_vat, cast(sum(total_incl_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM expenses GROUP BY date ORDER BY date
`

type GetYearlyExpenseStatisticsRow struct {
	Date            string
	ExpenseCount    int64
	TotalExcludeVat float64
	TotalIncludeVat float64
	VatAmount       float64
}

func (q *Queries) GetYearlyExpenseStatistics(ctx context.Context) ([]GetYearlyExpenseStatisticsRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyExpenseStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetYearlyExpenseStatisticsRow
	for rows.Next() {
		var i GetYearlyExpenseStatisticsRow
		if err := rows.Scan(
			&i.Date,
			&i.ExpenseCount,
			&i.TotalExcludeVat,
			&i.TotalIncludeVat,
			&i.VatAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyExpensesCount = `-- name: GetYearlyExpensesCount :many
SELECT to_char(issued_at, 'YYYY') as year, count(*) as count FROM expenses GROUP BY to_char(issued_at, 'YYYY')
`

type GetYearlyExpensesCountRow struct {
	Year  string
	Count int64
}

func (q *Queries) GetYearlyExpensesCount(ctx context.Context) ([]GetYearlyExpensesCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyExpensesCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetYearlyExpensesCountRow
	for rows.Next() {
		var i GetYearlyExpensesCountRow
		if err := rows.Scan(&i.Year, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyInvoiceCount = `-- name: GetYearlyInvoiceCount :many
SELECT to_char(issued_at, 'YYYY') as year, count(*) as count FROM invoices GROUP BY to_char(issued_at, 'YYYY')
`

type GetYearlyInvoiceCountRow struct {
	Year  string
	Count int64
}

func (q *Queries) GetYearlyInvoiceCount(ctx context.Context) ([]GetYearlyInvoiceCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyInvoiceCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetYearlyInvoiceCountRow
	for rows.Next() {
		var i GetYearlyInvoiceCountRow
		if err := rows.Scan(&i.Year, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyInvoiceStatistics = `-- name: GetYearlyInvoiceStatistics :many
SELECT to_char(issued_at, 'YYYY') as date, count(*) as invoice_count, cast(sum(total_exclude_vat)  as float) as total_exclude_vat, cast(sum(total_include_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM invoices GROUP BY to_char(issued_at, 'YYYY') ORDER BY date
`

type GetYearlyInvoiceStatisticsRow struct {
	Date            string
	InvoiceCount    int64
	TotalExcludeVat float64
	TotalIncludeVat float64
	VatAmount       float64
}

func (q *Queries) GetYearlyInvoiceStatistics(ctx context.Context) ([]GetYearlyInvoiceStatisticsRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyInvoiceStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetYearlyInvoiceStatisticsRow
	for rows.Next() {
		var i GetYearlyInvoiceStatisticsRow
		if err := rows.Scan(
			&i.Date,
			&i.InvoiceCount,
			&i.TotalExcludeVat,
			&i.TotalIncludeVat,
			&i.VatAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
