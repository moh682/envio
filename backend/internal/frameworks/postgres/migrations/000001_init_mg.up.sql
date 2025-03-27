CREATE TABLE invoices (
	id UUID PRIMARY KEY,
	invoice_number INTEGER unique not null,
	file_path VARCHAR(255),
	file_created_at TIMESTAMP,
	issued_at TIMESTAMP,
	vat_percentage float NOT NULL,
	vat_amount float NOT NULL,
	total_exclude_vat float NOT NULL,
	total_include_vat float NOT NULL,
	payment_status VARCHAR(255) not null
);

CREATE TABLE invoice_migrations (
	id UUID PRIMARY KEY,
	invoice_number INTEGER not null,
	failed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	file_path VARCHAR(255),
	has_migrated_successfully INTEGER DEFAULT 0,
	last_failed_attempt TIMESTAMP
);

CREATE TABLE payments (
	id UUID PRIMARY KEY,
	invoice_id UUID not null,
	amount float not null,
	paid_at TIMESTAMP,
	method VARCHAR(255) not null,

	FOREIGN KEY (invoice_id) REFERENCES invoices(id)
);

CREATE TABLE customers (
	id UUID PRIMARY KEY,
	car_registration VARCHAR(255),
	invoice_id UUID NOT NULL,
	name VARCHAR(255),
	email VARCHAR(255),
	phone VARCHAR(255),
	address VARCHAR(255),
	zip_code VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (invoice_id) REFERENCES invoices(id)
);


CREATE TABLE costs (
	id UUID PRIMARY KEY,
	invoice_id UUID not null,
	product_number VARCHAR(255),
	description VARCHAR(255),
	quantity float NOT NULL,
	unit_price float NOT NULL,
	total float NOT NULL,

	FOREIGN KEY (invoice_id) REFERENCES invoices(id)
);
