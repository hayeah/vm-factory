-- CREATE EXTENSION pgcrypto;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS product_inspections (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v1mc(),
	serial text unique not null,
	created_at timestamp not null default (now() at time zone 'utc'),
	updated_at timestamp not null
);



