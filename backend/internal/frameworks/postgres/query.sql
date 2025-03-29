-- name: GetAllInvoicesByOrganizationId :many
SELECT * FROM invoices WHERE organization_id = $1;

-- name: GetAllProductsByInvoiceId :many
SELECT * FROM invoice_products WHERE organization_id = $1 AND invoice_number = $2;

-- name: GetOrganizationByUserId :one
SELECT * FROM users_organizations WHERE user_id =$1;