-- name: CreatePayment :exec
INSERT INTO payments (id, invoice_id, method, amount) VALUES ($1, $2, $3, $4);

-- name: CreateInvoice :exec
INSERT INTO invoices (id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status) VALUES ($1, (SELECT COALESCE(MAX(invoice_number), 0) + 1 FROM invoices), $2, $3, $4, $5, $6, $7);

-- name: CreateCustomer :exec
INSERT INTO customers (id, car_registration, name, email, phone, address, zip_code, invoice_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: CreateCost :exec
INSERT INTO costs (id, product_number, invoice_id, description, quantity, unit_price, total) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateInvoiceMigration :exec
INSERT INTO invoice_migrations (id, invoice_number, file_path) VALUES ($1, $2, $3);

-- name: GetInvoiceByNumber :one
SELECT * FROM invoices WHERE invoice_number = $1;

-- name: GetInvoiceById :one
SELECT * FROM invoices WHERE id = $1;

-- name: GetAllInvoices :many
SELECT * FROM invoices LIMIT $1 OFFSET $2;

-- name: GetMaxInvoiceNumber :one
SELECT MAX(invoice_number) FROM invoices;

-- name: GetCustomerById :one
select * from customers where id = $1;

-- name: GetCustomerByInvoiceId :one
select * from customers where invoice_id = $1;

-- name: GetAllCostsByInvoiceId :many
select * from costs where invoice_id = $1;

-- name: GetPaymentsByInvoice :many
select * from payments where invoice_id = $1;

-- name: GetPaymentById :one
select * from payments where id = $1;

-- name: GetAllExpenses :many
select * from expenses LIMIT $1 OFFSET $2;

-- name: GetAllExpenseEntriesByExpenseId :many
select * from expenses_entries where expense_id = $1;

-- name: CreateExpense :exec
insert into expenses (id, company_id, total_incl_vat, total_excl_vat, vat_amount, vat_rate, issued_at, paid_at, paid_with) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: CreateExpenseEntry :exec
insert into expenses_entries (id, expense_id, serial, description, unit_price, quantity, total_excl_vat, total_incl_vat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetCompanyByID :one
select * from companies where id = $1;

-- name: GetCompaniesByName :many
select * from companies where name like $1;

-- name: CreateCompany :exec
insert into companies (id, name, cvr) VALUES ($1, $2, $3);

-- name: GetAllCompanies :many
select * from companies;

-- name: GetAllExpensesSince :many
select * from expenses where issued_at > $1;

-- name: GetAllInvoicesSince :many
select * from invoices where issued_at > $1;

-- name: GetYearlyInvoiceCount :many
SELECT to_char(issued_at, 'YYYY') as year, count(*) as count FROM invoices GROUP BY to_char(issued_at, 'YYYY');

-- name: GetYearlyExpensesCount :many
SELECT to_char(issued_at, 'YYYY') as year, count(*) as count FROM expenses GROUP BY to_char(issued_at, 'YYYY');

-- name: GetDailyInvoiceStatistics :many
SELECT to_char(issued_at, 'DD/MM/YYYY') as date, count(*) as invoice_count, cast(sum(total_exclude_vat) as float) as total_exclude_vat, cast(sum(total_include_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM invoices GROUP BY to_char(issued_at, 'DD/MM/YYYY') order BY date;

-- name: GetMonthlyInvoiceStatistics :many
SELECT to_char(issued_at, 'MM/YYYY') as date, count(*) as invoice_count, cast(sum(total_exclude_vat) as float) as total_exclude_vat, cast(sum(total_include_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM invoices GROUP BY to_char(issued_at, 'MM/YYYY') order BY date;

-- name: GetYearlyInvoiceStatistics :many
SELECT to_char(issued_at, 'YYYY') as date, count(*) as invoice_count, cast(sum(total_exclude_vat)  as float) as total_exclude_vat, cast(sum(total_include_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM invoices GROUP BY to_char(issued_at, 'YYYY') ORDER BY date;

-- name: GetDailyExpenseStatistics :many
SELECT to_char(issued_at, 'DD/MM/YYYY') as date, count(*) as expense_count, cast(sum(total_excl_vat) as float) as total_excl_vat, cast(sum(total_incl_vat) as float) as total_incl_vat, cast(sum(vat_amount) as float) as vat_amount FROM expenses GROUP BY to_char(issued_at, 'YYYY') ORDER BY date;

-- name: GetMonthlyExpenseStatistics :many
SELECT cast(strftime('%m/%Y', issued_at) as TEXT) as date, count(*) as expense_count, cast(sum(total_excl_vat) as float) as total_exclude_vat, cast(sum(total_incl_vat) as float) as total_incl_vat, cast(sum(vat_amount) as float) as vat_amount FROM expenses GROUP BY date order BY date;

-- name: GetYearlyExpenseStatistics :many
SELECT cast(strftime('%Y', issued_at) as TEXT) as date, count(*) as expense_count, cast(sum(total_excl_vat) as float) as total_exclude_vat, cast(sum(total_incl_vat) as float) as total_include_vat, cast(sum(vat_amount) as float) as vat_amount FROM expenses GROUP BY date ORDER BY date;
