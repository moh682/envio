CREATE TABLE companies (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  cvr INTEGER,
  created_at DATETIME NOT NULL default CURRENT_TIMESTAMP,
  updated_at DATETIME
);

ALTER TABLE expenses ADD COLUMN company_id UUID NOT NULL REFERENCES companies(id);

ALTER TABLE expenses DROP COLUMN company;
