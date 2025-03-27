-- Create the new table with an auto-incrementing invoice_number
CREATE TABLE invoices_new (
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

-- Copy data from the old table to the new table
INSERT INTO invoices_new (id, file_path, file_created_at, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status)
SELECT id, file_path, file_created_at, issued_at, vat_percentage, vat_amount, total_exclude_vat, total_include_vat, payment_status
FROM invoices;

-- Drop the old table
DROP TABLE invoices;

-- Rename the new table to the original name
ALTER TABLE invoices_new RENAME TO invoices;
