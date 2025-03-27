CREATE TABLE expenses ( 
	id UUID PRIMARY KEY,
	company VARCHAR(255) NOT NULL,
	total_excl_vat float NOT NULL,
	total_incl_vat float NOT NULL,
	vat_amount float NOT NULL,
	vat_rate float NOT NULL,
	issued_at TIMESTAMP NOT NULL,
	paid_at TIMESTAMP,
	paid_with INTEGER
);

CREATE TABLE expenses_entries (
	id UUID PRIMARY KEY,
	expense_id UUID NOT NULL,
	serial VARCHAR(255),
	description VARCHAR(255),
	unit_price float NOT NULL,
	quantity INTEGER NOT NULL,
	total_excl_vat float NOT NULL,
	total_incl_vat float NOT NULL,

	FOREIGN KEY (expense_id) REFERENCES expenses(id)
);
