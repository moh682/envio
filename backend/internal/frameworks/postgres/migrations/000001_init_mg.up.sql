-- organization
CREATE TABLE organizations (
    id                      UUID PRIMARY KEY,
    invoice_number_start    VARCHAR(50) NOT NULL
);

-- organization <=> Users
CREATE TABLE users_organizations (
    organization_id     UUID,
    user_id             UUID,
    PRIMARY KEY (organization_id, user_id),
    CONSTRAINT fk_organization_id FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

-- Finance accounts
CREATE TABLE financial_accounts (
    id                      UUID PRIMARY KEY,
    organization_id         UUID NOT NULL,
    name                    VARCHAR(255) NOT NULL,
    key                     VARCHAR(255) NOT NULL,
    is_vat                  BOOLEAN      NOT NULL,
    CONSTRAINT fk_organization_id FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

-- Financial year
CREATE TABLE financial_years (
    organization_id UUID,
    year            INT,
    PRIMARY KEY (organization_id, year),
    CONSTRAINT fk_organization_id FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

-- Invoices
CREATE TABLE invoices (
    organization_id         UUID,
    number                  INT,
    financial_year          INT NOT NULL,
    issue_date              DATE NULL,
    name                    VARCHAR(255),
    email                   VARCHAR(255),
    phone                   VARCHAR(50),
    car_registration        VARCHAR(50),
    total                   FLOAT NOT NULL,
    is_vat                  BOOLEAN DEFAULT TRUE NOT NULL,
    PRIMARY KEY (organization_id, number),
    CONSTRAINT invoice_contact_check CHECK (
        name IS NOT NULL
        OR email IS NOT NULL
        OR phone IS NOT NULL
        OR car_registration IS NOT NULL
    ),
    CONSTRAINT fk_organization_id FOREIGN KEY (organization_id) REFERENCES organizations (id),
    CONSTRAINT fk_financial_year FOREIGN KEY (organization_id, financial_year) REFERENCES financial_years (organization_id, year)
);

-- Products for each invoice
CREATE TABLE invoice_products (
    id              SERIAL,
    invoice_number  INT,
    organization_id UUID NOT NULL,
    serial          VARCHAR(255),
    description     VARCHAR(255) NOT NULL,
    quantity        INT          NOT NULL,
    rate            float NOT NULL,
    total           float NOT NULL,
    PRIMARY KEY (id, invoice_number),
    CONSTRAINT fk_invoice_number FOREIGN KEY (organization_id, invoice_number) REFERENCES invoices (organization_id, number)
);

-- Expenses
CREATE TABLE expenses (
    id                  UUID PRIMARY KEY,
    organization_id     UUID NOT NULL,
    financial_year      INT NOT NULL,
    issue_date          DATE NOT NULL,
    company             VARCHAR(255) NOT NULL,
    payment_option      VARCHAR(100) NOT NULL,
    account             VARCHAR(100) NOT NULL,
    amount              float NOT NULL,
    is_vat              BOOLEAN NOT NULL,
    total               float NOT NULL,
    CONSTRAINT fk_organization_id FOREIGN KEY (organization_id) REFERENCES organizations (id),
    CONSTRAINT fk_financial_year FOREIGN KEY (organization_id, financial_year) REFERENCES financial_years (number, organization_id)
);

