-- Create the original table without auto-incrementing invoice_number
CREATE TABLE invoices_old (
    id UUID PRIMARY KEY,
    invoice_number INTEGER UNIQUE NOT NULL,
    file_path VARCHAR(255),
    file_created_at TIMESTAMP,
    issued_at TIMESTAMP,
    vat_percentage float NOT NULL,
    vat_amount float NOT NULL,
    total_exclude_vat float NOT NULL,
    total_include_vat float NOT NULL,
    payment_status INTEGER NOT NULL
);

-- Copy data from the modified table back to the original structure
INSERT INTO invoices_old (id, invoice_number, file_path, file_created_at, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status)
SELECT id, invoice_number, file_path, file_created_at, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status
FROM invoices;

-- Drop the modified table
DROP TABLE invoices;

-- Rename the old table back to the original name
ALTER TABLE invoices_old RENAME TO invoices;
