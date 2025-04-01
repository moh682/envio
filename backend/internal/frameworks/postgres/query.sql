-- name: GetAllInvoicesByOrganizationId :many
SELECT * FROM invoices WHERE organization_id = $1;

-- name: GetAllProductsByInvoiceId :many
SELECT * FROM invoice_products WHERE organization_id = $1 AND invoice_number = $2;

-- name: GetOrganizationByUserId :one
SELECT * FROM organizations JOIN users_organizations ON organizations.id = users_organizations.organization_id AND users_organizations.user_id = $1;

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