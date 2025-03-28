-- 1) Drop invoice_products first (depends on invoices).
DROP TABLE IF EXISTS invoice_products;

-- 2) Drop invoices (depends on financial_years and organizations).
DROP TABLE IF EXISTS invoices;

-- 3) Drop expenses (depends on financial_years and organizations).
DROP TABLE IF EXISTS expenses;

-- 4) Drop financial_years (depends on organizations).
DROP TABLE IF EXISTS financial_years;

-- 5) Drop financial_accounts (depends on organizations).
DROP TABLE IF EXISTS financial_accounts;

-- 6) Drop users_organizations (depends on organizations).
DROP TABLE IF EXISTS users_organizations;

-- 7) Finally drop organizations.
DROP TABLE IF EXISTS organizations;
