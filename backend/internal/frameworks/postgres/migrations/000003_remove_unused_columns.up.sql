-- Create a new table without the file_path and file_created_at columns
CREATE TABLE invoices_new (
    id UUID PRIMARY KEY,
    invoice_number INTEGER UNIQUE NOT NULL,
    issued_at TIMESTAMP,
    vat_percentage float NOT NULL,
    vat_amount float NOT NULL,
    total_exclude_vat float NOT NULL,
    total_include_vat float NOT NULL,
    payment_status VARCAR(255) NOT NULL
);

-- Copy the data from the old table to the new table
INSERT INTO invoices_new (id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status)
SELECT id, invoice_number, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status
FROM invoices;

-- Drop the old table
DROP TABLE invoices;

-- Rename the new table to the original name
ALTER TABLE invoices_new RENAME TO invoices;

