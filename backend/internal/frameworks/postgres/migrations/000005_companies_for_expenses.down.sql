ALTER TABLE expenses DROP COLUMN company_id;
ALTER TABLE expenses ADD COLUMN company VARCHAR(255);
Drop table companies;
