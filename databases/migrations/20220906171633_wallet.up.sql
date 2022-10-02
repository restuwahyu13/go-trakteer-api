CREATE TABLE IF NOT EXISTS "wallet" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "no_rek" bigint NOT NULL,
  "bank_name" varchar(50) NOT NULL,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now())
)