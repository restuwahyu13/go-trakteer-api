CREATE TABLE IF NOT EXISTS "balance" (
  "id" serial PRIMARY KEY,
  "amount" bigint NOT NULL UNSIGNED,
  "created_at" timestamp NULL DEFAULT (now())
)