CREATE TABLE IF NOT EXISTS "balance" (
  "id" serial PRIMARY KEY,
  "amount" bigint NOT NULL,
  "created_at" timestamp NULL DEFAULT (now())
)