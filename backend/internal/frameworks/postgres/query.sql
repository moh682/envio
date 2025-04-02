-- name: GetAllInvoicesByOrganizationId :many
SELECT * FROM invoices WHERE organization_id = $1;

-- name: GetAllProductsByInvoiceId :many
SELECT * FROM invoice_products WHERE organization_id = $1 AND invoice_number = $2;

-- name: GetOrganizationByUserId :one
SELECT * FROM organizations JOIN users_organizations ON organizations.id = users_organizations.organization_id AND users_organizations.user_id = $1;

-- name: GetFinancialYearsByUserIdOrganizationId :many
SELECT
  financial_years.year
FROM
  financial_years
  JOIN users_organizations ON users_organizations.user_id = $1
  AND users_organizations.organization_id = financial_years.organization_id
WHERE financial_years.organization_id = $2;

-- name: CreateOrganization :exec
INSERT INTO organizations (
	id,
	name,
	invoice_number_start
) VALUES (
	$1,
	$2,
	$3
);

-- name: CreateOrganizationUser :exec
INSERT INTO users_organizations (
	organization_id,
	user_id
) VALUES (
	$1,
	$2
);

-- name: CreateFinancialYear :exec
INSERT INTO financial_years (
	organization_id,
	year
) VALUES (
	$1,
	$2
);
