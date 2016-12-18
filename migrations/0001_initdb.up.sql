CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS product_inspections (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	serial text unique not null,
	created_at timestamp not null,
	updated_at timestamp not null
);



