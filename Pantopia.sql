CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "company_name" varchar NOT NULL,
  "industry" varchar NOT NULL,
  "target_industry" varchar NOT NULL,
  "target_position" varchar NOT NULL,
  "target_size" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "leads" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "telephone_number" varchar NOT NULL,
  "target_email" varchar NOT NULL,
  "company_name" varchar NOT NULL,
  "conversation" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "leads" ("account_id");

ALTER TABLE "leads" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");