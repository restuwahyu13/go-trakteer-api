CREATE TABLE IF NOT EXISTS "payment_logs" (
  "id" serial PRIMARY KEY,
  "transaction_id" varchar(255) NULL,
  "external_id" varchar(255) NULL,
  "status" varchar(50) NOT NULL,
  "res_payload" jsonb NOT NULL,
  "resource_id" int NULL,
  "resource_type" varchar(50) NOT NULL,
  "created_at" timestamp NULL DEFAULT (now())
)