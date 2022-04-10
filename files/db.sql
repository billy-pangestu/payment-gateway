CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table files (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	type TEXT NOT NULL CHECK (char_length(type) <= 50),
	url TEXT NOT NULL CHECK (char_length(url) <= 255),
	user_upload TEXT NOT NULL CHECK (char_length(user_upload) <= 100),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE 
);

create table roles (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	code TEXT NOT NULL CHECK (char_length(code) <= 50),
	name TEXT NOT NULL CHECK (char_length(name) <= 50),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
INSERT INTO "public"."roles"("id", "code", "name", "status", "created_at", "updated_at", "deleted_at") VALUES ('8d777d38-ecf0-4c86-8054-4fe8bc6b761a', 'superadmin', 'Superadmin', 't', '2020-07-21 06:04:19.166447+07', '2020-07-21 06:04:19.166447+07', NULL);
INSERT INTO "public"."roles"("id", "code", "name", "status", "created_at", "updated_at", "deleted_at") VALUES ('bdee26a1-bad1-4b97-a592-9f86a61e7c37', 'admin', 'Admin', 't', '2020-07-21 06:04:19.168713+07', '2020-07-21 06:04:19.168713+07', NULL);

create table admins (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	code TEXT NOT NULL CHECK (char_length(code) <= 50),
	name TEXT NOT NULL CHECK (char_length(name) <= 50),
	email TEXT NOT NULL CHECK (char_length(email) <= 100),
	password TEXT NOT NULL CHECK (char_length(password) <= 1000),
  role_id uuid NOT NULL REFERENCES roles(id),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
INSERT INTO "public"."admins"("id", "code", "name", "email", "password", "role_id", "status", "created_at", "updated_at", "deleted_at") VALUES ('0bf7761a-f90e-47dd-b7c8-f13a1f328677', 'GO884J28', 'Superadmin 1', 'superadmin1@apollo.com', '459e67abe98800d65ce0eba8102f5987cd72513a3d117d778caa01b1', '8d777d38-ecf0-4c86-8054-4fe8bc6b761a', 't', '2020-07-21 07:08:56.608423+07', '2020-07-21 07:08:56.608423+07', NULL);
INSERT INTO "public"."admins"("id", "code", "name", "email", "password", "role_id", "status", "created_at", "updated_at", "deleted_at") VALUES ('98a26c38-0690-463e-a4e5-fb1ca9ef6fa9', '61ITW3YV', 'Superadmin 2', 'superadmin2@apollo.com', '610375ef544811d054cfce8b8d51739656478d12eb48170166eaf4f5', '8d777d38-ecf0-4c86-8054-4fe8bc6b761a', 't', '2020-07-21 07:09:02.473398+07', '2020-07-21 07:09:02.473398+07', NULL);
INSERT INTO "public"."admins"("id", "code", "name", "email", "password", "role_id", "status", "created_at", "updated_at", "deleted_at") VALUES ('31c38ba3-fe2d-4068-a925-daeb45bbf900', '02PJ7WV7', 'Superadmin 3', 'superadmin3@apollo.com', '7a43bdb81390c30978f7cea4cbb653c9bd491e149254df76ba6d74ec', '8d777d38-ecf0-4c86-8054-4fe8bc6b761a', 't', '2020-07-21 07:09:07.135754+07', '2020-07-21 07:09:07.135754+07', NULL);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "phone_number" text UNIQUE,
  "pin" text,
  "name" text,
  "phone_validated_at" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "other_user_id" uuid,
  "money_in" numeric(18,2),
  "money_out" numeric(18,2),
  "notes" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "tags" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "transaction_id" uuid,
  "user_id" uuid NOT NULL,
  "hash_tags" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "banks" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "name" text,
  "balance" numeric(18,2),
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("other_user_id") REFERENCES "users" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "banks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");