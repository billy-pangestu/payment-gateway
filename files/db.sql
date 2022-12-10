CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.history (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	payload varchar NULL,
	api varchar NULL,
	status varchar NULL,
	qid varchar NULL,
	error_string varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL
);

CREATE TABLE public.merchants (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NULL,
	amount float8 NULL DEFAULT 0,
	CONSTRAINT merchants_pk PRIMARY KEY (id)
);

CREATE TABLE public.user_roles (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar(255) NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);

CREATE TABLE public.users (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	first_name varchar(255) NULL,
	last_name varchar(255) NULL,
	unique_id varchar(255) NULL,
	"password" varchar(255) NULL,
	is_active bool NOT NULL DEFAULT true,
	role_id varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NULL,
	amount float8 NULL DEFAULT 0,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.user_roles(id)
);

CREATE TABLE public.transactions (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	user_id varchar NULL,
	merchant_id varchar NULL,
	amount float8 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (id),
	CONSTRAINT transactions_fk FOREIGN KEY (user_id) REFERENCES public.users(id),
	CONSTRAINT transactions_fk_1 FOREIGN KEY (merchant_id) REFERENCES public.merchants(id)
);